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
	"sync"
	"time"

	"github.com/Azure/azure-sdk-for-go/arm/compute"
	"github.com/golang/glog"

	"k8s.io/client-go/tools/cache"
	"k8s.io/kubernetes/pkg/cloudprovider"
)

var (
	cacheExpiration = 15 * time.Second
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
		cache:         newTimedcache(cacheExpiration),
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
