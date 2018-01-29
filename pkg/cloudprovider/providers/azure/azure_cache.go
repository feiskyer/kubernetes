/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package azure

import (
	"fmt"
	"sync"
	"time"

	"github.com/Azure/azure-sdk-for-go/arm/compute"
	"github.com/Azure/azure-sdk-for-go/arm/network"
	"github.com/golang/glog"

	"k8s.io/client-go/tools/cache"
	"k8s.io/kubernetes/pkg/cloudprovider"
)

var (
	defaultCacheExpiration = 15 * time.Second
	lbCacheExpiration      = 2 * time.Minute
)

type timedcacheEntry struct {
	key  string
	data interface{}
}

type timedcache struct {
	store cache.Store
	lock  sync.Mutex
}

// ttl time.Duration
func newTimedcache(ttl time.Duration) timedcache {
	return timedcache{
		store: cache.NewTTLStore(cacheKeyFunc, ttl),
	}
}

func cacheKeyFunc(obj interface{}) (string, error) {
	return obj.(*timedcacheEntry).key, nil
}

func (t *timedcache) GetOrCreate(key string, createFunc func() interface{}) (interface{}, error) {
	entry, exists, err := t.store.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if exists {
		return (entry.(*timedcacheEntry)).data, nil
	}

	t.lock.Lock()
	defer t.lock.Unlock()
	entry, exists, err = t.store.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if exists {
		return (entry.(*timedcacheEntry)).data, nil
	}

	if createFunc == nil {
		return nil, nil
	}
	created := createFunc()
	t.store.Add(&timedcacheEntry{
		key:  key,
		data: created,
	})
	return created, nil
}

func (t *timedcache) Delete(key string) {
	_ = t.store.Delete(&timedcacheEntry{
		key: key,
	})
}

func (t *timedcache) List() []interface{} {
	return t.store.List()
}

type vmCacheEntry struct {
	lock *sync.Mutex
	vm   *compute.VirtualMachine
}

type vmCache struct {
	cache         timedcache
	vmClient      VirtualMachinesClient
	resourceGroup string
}

func newVMCache(resourceGroup string, vmClient VirtualMachinesClient) *vmCache {
	return &vmCache{
		vmClient:      vmClient,
		resourceGroup: resourceGroup,
		cache:         newTimedcache(defaultCacheExpiration),
	}
}

func (vmc *vmCache) get(name string) (vm compute.VirtualMachine, err error) {
	entry, err := vmc.cache.GetOrCreate(name, func() interface{} {
		return &vmCacheEntry{
			lock: &sync.Mutex{},
			vm:   nil,
		}
	})
	if err != nil {
		return compute.VirtualMachine{}, err
	}

	vmEntry := entry.(*vmCacheEntry)
	if vmEntry.vm == nil {
		vmEntry.lock.Lock()
		defer vmEntry.lock.Unlock()
		if vmEntry.vm == nil {
			// Currently InstanceView request are used by azure_zones, while the calls come after non-InstanceView
			// request. If we first send an InstanceView request and then a non InstanceView request, the second
			// request will still hit throttling. This is what happens now for cloud controller manager: In this
			// case we do get instance view every time to fulfill the azure_zones requirement without hitting
			// throttling.
			// Consider adding separate parameter for controlling 'InstanceView' once node update issue #56276 is fixed
			vm, err = vmc.vmClient.Get(vmc.resourceGroup, name, compute.InstanceView)
			exists, realErr := checkResourceExistsFromError(err)
			if realErr != nil {
				return vm, realErr
			}

			if !exists {
				return vm, cloudprovider.InstanceNotFound
			}

			vmEntry.vm = &vm
		}
		return *vmEntry.vm, nil
	}

	glog.V(6).Infof("vmCache.get() hits cache for(%s)", name)
	return *vmEntry.vm, nil
}

func (vmc *vmCache) delete(name string) {
	vmc.cache.Delete(name)
}

// lbCache holds a list of cached loadbalancers.
type lbCache struct {
	resourceGroup string
	lbClient      LoadBalancersClient

	lock        sync.Mutex
	cache       map[string]*network.LoadBalancer
	lastRefresh time.Time
}

// newLBCache creates a new lbCache.
func newLBCache(resourceGroup string, lbClient LoadBalancersClient) *lbCache {
	return &lbCache{
		lbClient:      lbClient,
		resourceGroup: resourceGroup,
		cache:         make(map[string]*network.LoadBalancer),
	}
}

// listLBFull gets a list of loadbalancers by calling Azure API.
func (lbc *lbCache) listLBFull() (map[string]*network.LoadBalancer, error) {
	allLBs := map[string]*network.LoadBalancer{}
	var result network.LoadBalancerListResult

	result, err := lbc.lbClient.List(lbc.resourceGroup)
	if err != nil {
		glog.Errorf("LoadBalancerClient.List(%s) failed: %v", lbc.resourceGroup, err)
		return nil, err
	}

	moreResults := (result.Value != nil && len(*result.Value) > 0)
	for moreResults {
		for idx := range *result.Value {
			lb := (*result.Value)[idx]
			allLBs[*lb.Name] = &lb
		}
		moreResults = false

		// follow the next link to get all the vms for resource group
		if result.NextLink != nil {
			result, err = lbc.lbClient.ListNextResults(lbc.resourceGroup, result)
			if err != nil {
				glog.Errorf("LoadBalancerClient.ListNextResults(%s) failed: %v", lbc.resourceGroup, err)
				return nil, err
			}

			moreResults = (result.Value != nil && len(*result.Value) > 0)
		}
	}

	return allLBs, nil
}

// list gets a list of loadbalancers from cache.
func (lbc *lbCache) list() (map[string]*network.LoadBalancer, error) {
	lbc.lock.Lock()
	defer lbc.lock.Unlock()

	if lbc.lastRefresh.Add(lbCacheExpiration).After(time.Now()) {
		return lbc.cache, nil
	}

	if err := lbc.forceRefresh(); err != nil {
		return nil, err
	}

	return lbc.cache, nil
}

// forceRefresh refreshes cache by calling Azure API.
func (lbc *lbCache) forceRefresh() error {
	glog.V(5).Infof("Refreshing lbCache")
	newLBs, err := lbc.listLBFull()
	if err != nil {
		return err
	}

	lbc.cache = newLBs
	lbc.lastRefresh = time.Now()
	return nil
}

// get gets a loadbalancer by name.
func (lbc *lbCache) get(name string) (*network.LoadBalancer, error) {
	lbc.lock.Lock()
	defer lbc.lock.Unlock()

	if lb, ok := lbc.cache[name]; ok {
		return lb, nil
	}

	return nil, fmt.Errorf("loadbalancer %q not found", name)
}

// update updates a loadbalancer's cache.
func (lbc *lbCache) update(name string, newLB *network.LoadBalancer) {
	lbc.lock.Lock()
	defer lbc.lock.Unlock()

	lbc.cache[name] = newLB
	lbc.lastRefresh = time.Now()
}

// delete removes a loadbalancer from cache.
func (lbc *lbCache) delete(name string) {
	lbc.lock.Lock()
	defer lbc.lock.Unlock()

	delete(lbc.cache, name)
	lbc.lastRefresh = time.Now().Add(-24 * time.Hour)
}
