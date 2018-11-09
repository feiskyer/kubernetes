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
	"net/http"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2017-09-01/network"

	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog"
)

// Event creates a event for the specified object.
func (az *Cloud) Event(obj runtime.Object, eventtype, reason, message string) {
	if obj != nil && reason != "" {
		az.eventRecorder.Event(obj, eventtype, reason, message)
	}
}

// ListVirtualMachines invokes az.VirtualMachinesClient.List with exponential backoff retry
func (az *Cloud) ListVirtualMachines(resourceGroup string) ([]compute.VirtualMachine, error) {
	ctx, cancel := getContextWithCancel()
	defer cancel()

	allNodes, err := az.VirtualMachinesClient.List(ctx, resourceGroup)
	if err != nil {
		klog.Errorf("VirtualMachinesClient.List(%v) failure with err=%v", resourceGroup, err)
		return nil, err
	}
	klog.V(2).Infof("VirtualMachinesClient.List(%v) success", resourceGroup)
	return allNodes, nil
}

// CreateOrUpdateSecurityGroup invokes az.SecurityGroupsClient.CreateOrUpdate with exponential backoff retry
func (az *Cloud) CreateOrUpdateSecurityGroup(service *v1.Service, sg network.SecurityGroup) error {
	ctx, cancel := getContextWithCancel()
	defer cancel()

	resp, err := az.SecurityGroupsClient.CreateOrUpdate(ctx, az.ResourceGroup, *sg.Name, sg)
	klog.V(10).Infof("SecurityGroupsClient.CreateOrUpdate(%s): end", *sg.Name)
	if err == nil {
		if isSuccessHTTPResponse(resp) {
			// Invalidate the cache right after updating
			az.nsgCache.Delete(*sg.Name)
		} else if resp != nil {
			return fmt.Errorf("HTTP response %q", resp.Status)
		}
	}
	return err
}

// CreateOrUpdateLB invokes az.LoadBalancerClient.CreateOrUpdate with exponential backoff retry
func (az *Cloud) CreateOrUpdateLB(service *v1.Service, lb network.LoadBalancer) error {
	ctx, cancel := getContextWithCancel()
	defer cancel()

	resp, err := az.LoadBalancerClient.CreateOrUpdate(ctx, az.ResourceGroup, *lb.Name, lb)
	klog.V(10).Infof("LoadBalancerClient.CreateOrUpdate(%s): end", *lb.Name)
	if err == nil {
		if isSuccessHTTPResponse(resp) {
			// Invalidate the cache right after updating
			az.lbCache.Delete(*lb.Name)
		} else if resp != nil {
			return fmt.Errorf("HTTP response %q", resp.Status)
		}
	}
	return err
}

// ListLB invokes az.LoadBalancerClient.List with exponential backoff retry
func (az *Cloud) ListLB(service *v1.Service) ([]network.LoadBalancer, error) {
	ctx, cancel := getContextWithCancel()
	defer cancel()

	allLBs, err := az.LoadBalancerClient.List(ctx, az.ResourceGroup)
	if err != nil {
		az.Event(service, v1.EventTypeWarning, "ListLoadBalancers", err.Error())
		klog.Errorf("LoadBalancerClient.List(%v) failure with err=%v", az.ResourceGroup, err)
		return nil, err
	}
	klog.V(2).Infof("LoadBalancerClient.List(%v) success", az.ResourceGroup)
	return allLBs, nil
}

// ListPIP list the PIP resources in the given resource group
func (az *Cloud) ListPIP(service *v1.Service, pipResourceGroup string) ([]network.PublicIPAddress, error) {
	ctx, cancel := getContextWithCancel()
	defer cancel()

	allPIPs, err := az.PublicIPAddressesClient.List(ctx, pipResourceGroup)
	if err != nil {
		az.Event(service, v1.EventTypeWarning, "ListPublicIPs", err.Error())
		klog.Errorf("PublicIPAddressesClient.List(%v) failure with err=%v", pipResourceGroup, err)
		return nil, err
	}
	klog.V(2).Infof("PublicIPAddressesClient.List(%v) success", pipResourceGroup)
	return allPIPs, nil
}

// CreateOrUpdatePIP invokes az.PublicIPAddressesClient.CreateOrUpdate with exponential backoff retry
func (az *Cloud) CreateOrUpdatePIP(service *v1.Service, pipResourceGroup string, pip network.PublicIPAddress) error {
	ctx, cancel := getContextWithCancel()
	defer cancel()

	resp, err := az.PublicIPAddressesClient.CreateOrUpdate(ctx, pipResourceGroup, *pip.Name, pip)
	klog.V(10).Infof("PublicIPAddressesClient.CreateOrUpdate(%s, %s): end", pipResourceGroup, *pip.Name)
	return az.processHTTPRetryResponse(service, "CreateOrUpdatePublicIPAddress", resp, err)
}

// CreateOrUpdateInterface invokes az.PublicIPAddressesClient.CreateOrUpdate with exponential backoff retry
func (az *Cloud) CreateOrUpdateInterface(service *v1.Service, nic network.Interface) error {
	ctx, cancel := getContextWithCancel()
	defer cancel()

	resp, err := az.InterfacesClient.CreateOrUpdate(ctx, az.ResourceGroup, *nic.Name, nic)
	klog.V(10).Infof("InterfacesClient.CreateOrUpdate(%s): end", *nic.Name)
	return az.processHTTPRetryResponse(service, "CreateOrUpdateInterface", resp, err)
}

// DeletePublicIP invokes az.PublicIPAddressesClient.Delete with exponential backoff retry
func (az *Cloud) DeletePublicIP(service *v1.Service, pipResourceGroup string, pipName string) error {
	ctx, cancel := getContextWithCancel()
	defer cancel()

	resp, err := az.PublicIPAddressesClient.Delete(ctx, pipResourceGroup, pipName)
	return az.processHTTPRetryResponse(service, "DeletePublicIPAddress", resp, err)
}

// DeleteLB invokes az.LoadBalancerClient.Delete with exponential backoff retry
func (az *Cloud) DeleteLB(service *v1.Service, lbName string) error {
	ctx, cancel := getContextWithCancel()
	defer cancel()

	resp, err := az.LoadBalancerClient.Delete(ctx, az.ResourceGroup, lbName)
	if err == nil {
		if isSuccessHTTPResponse(resp) {
			// Invalidate the cache right after updating
			az.lbCache.Delete(lbName)
		} else if resp != nil {
			return fmt.Errorf("HTTP response %q", resp.Status)
		}
	}
	return err
}

// CreateOrUpdateRouteTable invokes az.RouteTablesClient.CreateOrUpdate with exponential backoff retry
func (az *Cloud) CreateOrUpdateRouteTable(routeTable network.RouteTable) error {
	ctx, cancel := getContextWithCancel()
	defer cancel()

	resp, err := az.RouteTablesClient.CreateOrUpdate(ctx, az.ResourceGroup, az.RouteTableName, routeTable)
	return az.processHTTPRetryResponse(nil, "", resp, err)
}

// CreateOrUpdateRoute invokes az.RoutesClient.CreateOrUpdate with exponential backoff retry
func (az *Cloud) CreateOrUpdateRoute(route network.Route) error {
	ctx, cancel := getContextWithCancel()
	defer cancel()

	resp, err := az.RoutesClient.CreateOrUpdate(ctx, az.ResourceGroup, az.RouteTableName, *route.Name, route)
	klog.V(10).Infof("RoutesClient.CreateOrUpdate(%s): end", *route.Name)
	return az.processHTTPRetryResponse(nil, "", resp, err)
}

// isSuccessHTTPResponse determines if the response from an HTTP request suggests success
func isSuccessHTTPResponse(resp *http.Response) bool {
	if resp == nil {
		return false
	}

	// HTTP 2xx suggests a successful response
	if 199 < resp.StatusCode && resp.StatusCode < 300 {
		return true
	}

	return false
}

func (az *Cloud) processHTTPRetryResponse(service *v1.Service, reason string, resp *http.Response, err error) error {
	if isSuccessHTTPResponse(resp) {
		// HTTP 2xx suggests a successful response
		return nil
	}

	if err != nil {
		az.Event(service, v1.EventTypeWarning, reason, err.Error())
		klog.Errorf("processHTTPRetryResponse failure with err: %v", err)
	} else if resp != nil {
		az.Event(service, v1.EventTypeWarning, reason, fmt.Sprintf("Azure HTTP response %d", resp.StatusCode))
		klog.Errorf("processHTTPRetryResponse failure with HTTP response %q", resp.Status)
	}

	return err
}
