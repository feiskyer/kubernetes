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
	"time"

	"github.com/Azure/azure-sdk-for-go/arm/compute"
	"github.com/Azure/azure-sdk-for-go/arm/disk"
	"github.com/Azure/azure-sdk-for-go/arm/network"
	"github.com/Azure/azure-sdk-for-go/arm/storage"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/adal"
)

// VirtualMachinesClient defines needed functions for azure compute.VirtualMachinesClient
type VirtualMachinesClient interface {
	CreateOrUpdate(resourceGroupName string, VMName string, parameters compute.VirtualMachine, cancel <-chan struct{}) (<-chan compute.VirtualMachine, <-chan error)
	Get(resourceGroupName string, VMName string, expand compute.InstanceViewTypes) (result compute.VirtualMachine, err error)
	List(resourceGroupName string) (result compute.VirtualMachineListResult, err error)
	ListNextResults(lastResults compute.VirtualMachineListResult) (result compute.VirtualMachineListResult, err error)
}

// InterfacesClient defines needed functions for azure network.InterfacesClient
type InterfacesClient interface {
	CreateOrUpdate(resourceGroupName string, networkInterfaceName string, parameters network.Interface, cancel <-chan struct{}) (<-chan network.Interface, <-chan error)
	Get(resourceGroupName string, networkInterfaceName string, expand string) (result network.Interface, err error)
	GetVirtualMachineScaleSetNetworkInterface(resourceGroupName string, virtualMachineScaleSetName string, virtualmachineIndex string, networkInterfaceName string, expand string) (result network.Interface, err error)
}

// LoadBalancersClient defines needed functions for azure network.LoadBalancersClient
type LoadBalancersClient interface {
	CreateOrUpdate(resourceGroupName string, loadBalancerName string, parameters network.LoadBalancer, cancel <-chan struct{}) (<-chan network.LoadBalancer, <-chan error)
	Delete(resourceGroupName string, loadBalancerName string, cancel <-chan struct{}) (<-chan autorest.Response, <-chan error)
	Get(resourceGroupName string, loadBalancerName string, expand string) (result network.LoadBalancer, err error)
	List(resourceGroupName string) (result network.LoadBalancerListResult, err error)
	ListNextResults(lastResult network.LoadBalancerListResult) (result network.LoadBalancerListResult, err error)
}

// PublicIPAddressesClient defines needed functions for azure network.PublicIPAddressesClient
type PublicIPAddressesClient interface {
	CreateOrUpdate(resourceGroupName string, publicIPAddressName string, parameters network.PublicIPAddress, cancel <-chan struct{}) (<-chan network.PublicIPAddress, <-chan error)
	Delete(resourceGroupName string, publicIPAddressName string, cancel <-chan struct{}) (<-chan autorest.Response, <-chan error)
	Get(resourceGroupName string, publicIPAddressName string, expand string) (result network.PublicIPAddress, err error)
	List(resourceGroupName string) (result network.PublicIPAddressListResult, err error)
	ListNextResults(lastResults network.PublicIPAddressListResult) (result network.PublicIPAddressListResult, err error)
}

// SubnetsClient defines needed functions for azure network.SubnetsClient
type SubnetsClient interface {
	CreateOrUpdate(resourceGroupName string, virtualNetworkName string, subnetName string, subnetParameters network.Subnet, cancel <-chan struct{}) (<-chan network.Subnet, <-chan error)
	Delete(resourceGroupName string, virtualNetworkName string, subnetName string, cancel <-chan struct{}) (<-chan autorest.Response, <-chan error)
	Get(resourceGroupName string, virtualNetworkName string, subnetName string, expand string) (result network.Subnet, err error)
	List(resourceGroupName string, virtualNetworkName string) (result network.SubnetListResult, err error)
}

// SecurityGroupsClient defines needed functions for azure network.SecurityGroupsClient
type SecurityGroupsClient interface {
	CreateOrUpdate(resourceGroupName string, networkSecurityGroupName string, parameters network.SecurityGroup, cancel <-chan struct{}) (<-chan network.SecurityGroup, <-chan error)
	Delete(resourceGroupName string, networkSecurityGroupName string, cancel <-chan struct{}) (<-chan autorest.Response, <-chan error)
	Get(resourceGroupName string, networkSecurityGroupName string, expand string) (result network.SecurityGroup, err error)
	List(resourceGroupName string) (result network.SecurityGroupListResult, err error)
}

// VirtualMachineScaleSetsClient defines needed functions for azure compute.VirtualMachineScaleSetsClient
type VirtualMachineScaleSetsClient interface {
	CreateOrUpdate(resourceGroupName string, VMScaleSetName string, parameters compute.VirtualMachineScaleSet, cancel <-chan struct{}) (<-chan compute.VirtualMachineScaleSet, <-chan error)
	Get(resourceGroupName string, VMScaleSetName string) (result compute.VirtualMachineScaleSet, err error)
	List(resourceGroupName string) (result compute.VirtualMachineScaleSetListResult, err error)
	ListNextResults(lastResults compute.VirtualMachineScaleSetListResult) (result compute.VirtualMachineScaleSetListResult, err error)
	UpdateInstances(resourceGroupName string, VMScaleSetName string, VMInstanceIDs compute.VirtualMachineScaleSetVMInstanceRequiredIDs, cancel <-chan struct{}) (<-chan compute.OperationStatusResponse, <-chan error)
}

// VirtualMachineScaleSetVMsClient defines needed functions for azure compute.VirtualMachineScaleSetVMsClient
type VirtualMachineScaleSetVMsClient interface {
	Get(resourceGroupName string, VMScaleSetName string, instanceID string) (result compute.VirtualMachineScaleSetVM, err error)
	GetInstanceView(resourceGroupName string, VMScaleSetName string, instanceID string) (result compute.VirtualMachineScaleSetVMInstanceView, err error)
	List(resourceGroupName string, virtualMachineScaleSetName string, filter string, selectParameter string, expand string) (result compute.VirtualMachineScaleSetVMListResult, err error)
	ListNextResults(lastResults compute.VirtualMachineScaleSetVMListResult) (result compute.VirtualMachineScaleSetVMListResult, err error)
}

// RoutesClient defines needed functions for azure network.RoutesClient
type RoutesClient interface {
	CreateOrUpdate(resourceGroupName string, routeTableName string, routeName string, routeParameters network.Route, cancel <-chan struct{}) (<-chan network.Route, <-chan error)
	Delete(resourceGroupName string, routeTableName string, routeName string, cancel <-chan struct{}) (<-chan autorest.Response, <-chan error)
}

// RouteTablesClient defines needed functions for azure network.RouteTablesClient
type RouteTablesClient interface {
	CreateOrUpdate(resourceGroupName string, routeTableName string, parameters network.RouteTable, cancel <-chan struct{}) (<-chan network.RouteTable, <-chan error)
	Get(resourceGroupName string, routeTableName string, expand string) (result network.RouteTable, err error)
}

// StorageAccountClient defines needed functions for azure storage.AccountsClient
type StorageAccountClient interface {
	Create(resourceGroupName string, accountName string, parameters storage.AccountCreateParameters, cancel <-chan struct{}) (<-chan storage.Account, <-chan error)
	Delete(resourceGroupName string, accountName string) (result autorest.Response, err error)
	ListKeys(resourceGroupName string, accountName string) (result storage.AccountListKeysResult, err error)
	ListByResourceGroup(resourceGroupName string) (result storage.AccountListResult, err error)
	GetProperties(resourceGroupName string, accountName string) (result storage.Account, err error)
}

// DisksClient defines needed functions for azure disk.DisksClient
type DisksClient interface {
	CreateOrUpdate(resourceGroupName string, diskName string, diskParameter disk.Model, cancel <-chan struct{}) (<-chan disk.Model, <-chan error)
	Delete(resourceGroupName string, diskName string, cancel <-chan struct{}) (<-chan disk.OperationStatusResponse, <-chan error)
	Get(resourceGroupName string, diskName string) (result disk.Model, err error)
}

// azVirtualMachinesClient implements VirtualMachinesClient.
type azVirtualMachinesClient struct {
	client compute.VirtualMachinesClient
}

func newAzVirtualMachinesClient(subscriptionID string, endpoint, servicePrincipalToken *adal.ServicePrincipalToken) *azVirtualMachinesClient {
	virtualMachinesClient := compute.NewVirtualMachinesClient(subscriptionID)
	virtualMachinesClient.BaseURI = endpoint
	virtualMachinesClient.Authorizer = autorest.NewBearerAuthorizer(servicePrincipalToken)
	virtualMachinesClient.PollingDelay = 5 * time.Second
	configureUserAgent(&virtualMachinesClient.Client)

	return &azVirtualMachinesClient{
		client: virtualMachinesClient,
	}
}

func (az *azVirtualMachinesClient) CreateOrUpdate(resourceGroupName string, VMName string, parameters compute.VirtualMachine, cancel <-chan struct{}) (<-chan compute.VirtualMachine, <-chan error) {

}

func (az *azVirtualMachinesClient) Get(resourceGroupName string, VMName string, expand compute.InstanceViewTypes) (result compute.VirtualMachine, err error) {

}

func (az *azVirtualMachinesClient) List(resourceGroupName string) (result compute.VirtualMachineListResult, err error) {

}

func (az *azVirtualMachinesClient) ListNextResults(lastResults compute.VirtualMachineListResult) (result compute.VirtualMachineListResult, err error) {

}

// azInterfacesClient implements InterfacesClient.
type azInterfacesClient struct {
	client network.InterfacesClient
}

func newAzInterfacesClient(subscriptionID string, endpoint, servicePrincipalToken *adal.ServicePrincipalToken) *azInterfacesClient {
	interfacesClient := network.NewInterfacesClient(subscriptionID)
	interfacesClient.BaseURI = endpoint
	interfacesClient.Authorizer = autorest.NewBearerAuthorizer(servicePrincipalToken)
	interfacesClient.PollingDelay = 5 * time.Second
	configureUserAgent(&interfacesClient.Client)

	return &azInterfacesClient{
		client: interfacesClient,
	}
}

func (az *azInterfacesClient) CreateOrUpdate(resourceGroupName string, networkInterfaceName string, parameters network.Interface, cancel <-chan struct{}) (<-chan network.Interface, <-chan error) {

}

func (az *azInterfacesClient) Get(resourceGroupName string, networkInterfaceName string, expand string) (result network.Interface, err error) {

}

func (az *azInterfacesClient) GetVirtualMachineScaleSetNetworkInterface(resourceGroupName string, virtualMachineScaleSetName string, virtualmachineIndex string, networkInterfaceName string, expand string) (result network.Interface, err error) {

}

// azLoadBalancersClient implements LoadBalancersClient.
type azLoadBalancersClient struct {
	client network.LoadBalancersClient
}

func newAzLoadBalancersClient(subscriptionID string, endpoint, servicePrincipalToken *adal.ServicePrincipalToken) *azLoadBalancersClient {
	loadBalancerClient := network.NewLoadBalancersClient(subscriptionID)
	loadBalancerClient.BaseURI = endpoint
	loadBalancerClient.Authorizer = autorest.NewBearerAuthorizer(servicePrincipalToken)
	loadBalancerClient.PollingDelay = 5 * time.Second
	configureUserAgent(&loadBalancerClient.Client)

	return &azLoadBalancersClient{
		client: loadBalancerClient,
	}
}

func (az *azLoadBalancersClient) CreateOrUpdate(resourceGroupName string, loadBalancerName string, parameters network.LoadBalancer, cancel <-chan struct{}) (<-chan network.LoadBalancer, <-chan error) {

}

func (az *azLoadBalancersClient) Delete(resourceGroupName string, loadBalancerName string, cancel <-chan struct{}) (<-chan autorest.Response, <-chan error) {

}

func (az *azLoadBalancersClient) Get(resourceGroupName string, loadBalancerName string, expand string) (result network.LoadBalancer, err error) {

}

func (az *azLoadBalancersClient) List(resourceGroupName string) (result network.LoadBalancerListResult, err error) {

}

func (az *azLoadBalancersClient) ListNextResults(lastResult network.LoadBalancerListResult) (result network.LoadBalancerListResult, err error) {

}

// azPublicIPAddressesClient implements PublicIPAddressesClient.
type azPublicIPAddressesClient struct {
	client network.PublicIPAddressesClient
}

func newAzPublicIPAddressesClient(subscriptionID string, endpoint, servicePrincipalToken *adal.ServicePrincipalToken) *azPublicIPAddressesClient {
	publicIPAddressClient := network.NewPublicIPAddressesClient(subscriptionID)
	publicIPAddressClient.BaseURI = endpoint
	publicIPAddressClient.Authorizer = autorest.NewBearerAuthorizer(servicePrincipalToken)
	publicIPAddressClient.PollingDelay = 5 * time.Second
	configureUserAgent(&publicIPAddressClient.Client)

	return &azPublicIPAddressesClient{
		client: publicIPAddressClient,
	}
}

func (az *azPublicIPAddressesClient) CreateOrUpdate(resourceGroupName string, publicIPAddressName string, parameters network.PublicIPAddress, cancel <-chan struct{}) (<-chan network.PublicIPAddress, <-chan error) {

}

func (az *azPublicIPAddressesClient) Delete(resourceGroupName string, publicIPAddressName string, cancel <-chan struct{}) (<-chan autorest.Response, <-chan error) {

}

func (az *azPublicIPAddressesClient) Get(resourceGroupName string, publicIPAddressName string, expand string) (result network.PublicIPAddress, err error) {

}

func (az *azPublicIPAddressesClient) List(resourceGroupName string) (result network.PublicIPAddressListResult, err error) {

}

func (az *azPublicIPAddressesClient) ListNextResults(lastResults network.PublicIPAddressListResult) (result network.PublicIPAddressListResult, err error) {

}

// azSubnetsClient implements SubnetsClient.
type azSubnetsClient struct {
	client network.SubnetsClient
}

func newAzSubnetsClient(subscriptionID string, endpoint, servicePrincipalToken *adal.ServicePrincipalToken) *azSubnetsClient {
	subnetsClient := network.NewSubnetsClient(subscriptionID)
	subnetsClient.BaseURI = endpoint
	subnetsClient.Authorizer = autorest.NewBearerAuthorizer(servicePrincipalToken)
	subnetsClient.PollingDelay = 5 * time.Second
	configureUserAgent(&subnetsClient.Client)

	return &azSubnetsClient{
		client: subnetsClient,
	}
}

func (az *azSubnetsClient) CreateOrUpdate(resourceGroupName string, virtualNetworkName string, subnetName string, subnetParameters network.Subnet, cancel <-chan struct{}) (<-chan network.Subnet, <-chan error) {

}

func (az *azSubnetsClient) Delete(resourceGroupName string, virtualNetworkName string, subnetName string, cancel <-chan struct{}) (<-chan autorest.Response, <-chan error) {

}

func (az *azSubnetsClient) Get(resourceGroupName string, virtualNetworkName string, subnetName string, expand string) (result network.Subnet, err error) {

}

func (az *azSubnetsClient) List(resourceGroupName string, virtualNetworkName string) (result network.SubnetListResult, err error) {

}

// azSecurityGroupsClient implements SecurityGroupsClient.
type azSecurityGroupsClient struct {
	client network.SecurityGroupsClient
}

func newAzSecurityGroupsClient(subscriptionID string, endpoint, servicePrincipalToken *adal.ServicePrincipalToken) *azSecurityGroupsClient {
	securityGroupsClient := network.NewSecurityGroupsClient(subscriptionID)
	securityGroupsClient.BaseURI = endpoint
	securityGroupsClient.Authorizer = autorest.NewBearerAuthorizer(servicePrincipalToken)
	securityGroupsClient.PollingDelay = 5 * time.Second
	configureUserAgent(&securityGroupsClient.Client)

	return &azSecurityGroupsClient{
		client: securityGroupsClient,
	}
}

func (az *azSecurityGroupsClient) CreateOrUpdate(resourceGroupName string, networkSecurityGroupName string, parameters network.SecurityGroup, cancel <-chan struct{}) (<-chan network.SecurityGroup, <-chan error) {

}

func (az *azSecurityGroupsClient) Delete(resourceGroupName string, networkSecurityGroupName string, cancel <-chan struct{}) (<-chan autorest.Response, <-chan error) {

}

func (az *azSecurityGroupsClient) Get(resourceGroupName string, networkSecurityGroupName string, expand string) (result network.SecurityGroup, err error) {

}

func (az *azSecurityGroupsClient) List(resourceGroupName string) (result network.SecurityGroupListResult, err error) {

}

// azVirtualMachineScaleSetsClient implements VirtualMachineScaleSetsClient.
type azVirtualMachineScaleSetsClient struct {
	client compute.VirtualMachineScaleSetsClient
}

func newAzVirtualMachineScaleSetsClient(subscriptionID string, endpoint, servicePrincipalToken *adal.ServicePrincipalToken) *azVirtualMachineScaleSetsClient {
	virtualMachineScaleSetsClient := compute.NewVirtualMachineScaleSetsClient(subscriptionID)
	virtualMachineScaleSetsClient.BaseURI = endpoint
	virtualMachineScaleSetsClient.Authorizer = autorest.NewBearerAuthorizer(servicePrincipalToken)
	virtualMachineScaleSetsClient.PollingDelay = 5 * time.Second
	configureUserAgent(&virtualMachineScaleSetsClient.Client)

	return &azVirtualMachineScaleSetsClient{
		client: virtualMachineScaleSetsClient,
	}
}

func (az *azVirtualMachineScaleSetsClient) CreateOrUpdate(resourceGroupName string, VMScaleSetName string, parameters compute.VirtualMachineScaleSet, cancel <-chan struct{}) (<-chan compute.VirtualMachineScaleSet, <-chan error) {

}

func (az *azVirtualMachineScaleSetsClient) Get(resourceGroupName string, VMScaleSetName string) (result compute.VirtualMachineScaleSet, err error) {

}

func (az *azVirtualMachineScaleSetsClient) List(resourceGroupName string) (result compute.VirtualMachineScaleSetListResult, err error) {

}

func (az *azVirtualMachineScaleSetsClient) ListNextResults(lastResults compute.VirtualMachineScaleSetListResult) (result compute.VirtualMachineScaleSetListResult, err error) {

}

func (az *azVirtualMachineScaleSetsClient) UpdateInstances(resourceGroupName string, VMScaleSetName string, VMInstanceIDs compute.VirtualMachineScaleSetVMInstanceRequiredIDs, cancel <-chan struct{}) (<-chan compute.OperationStatusResponse, <-chan error) {

}

// azVirtualMachineScaleSetVMsClient implements VirtualMachineScaleSetVMsClient.
type azVirtualMachineScaleSetVMsClient struct {
	client compute.VirtualMachineScaleSetVMsClient
}

func newAzVirtualMachineScaleSetVMsClient(subscriptionID string, endpoint, servicePrincipalToken *adal.ServicePrincipalToken) *azVirtualMachineScaleSetVMsClient {
	virtualMachineScaleSetVMsClient := compute.NewVirtualMachineScaleSetVMsClient(subscriptionID)
	virtualMachineScaleSetVMsClient.BaseURI = endpoint
	virtualMachineScaleSetVMsClient.Authorizer = autorest.NewBearerAuthorizer(servicePrincipalToken)
	virtualMachineScaleSetVMsClient.PollingDelay = 5 * time.Second
	configureUserAgent(&virtualMachineScaleSetVMsClient.Client)

	return &azVirtualMachineScaleSetVMsClient{
		client: virtualMachineScaleSetVMsClient,
	}
}

func (az *azVirtualMachineScaleSetVMsClient) Get(resourceGroupName string, VMScaleSetName string, instanceID string) (result compute.VirtualMachineScaleSetVM, err error) {

}

func (az *azVirtualMachineScaleSetVMsClient) GetInstanceView(resourceGroupName string, VMScaleSetName string, instanceID string) (result compute.VirtualMachineScaleSetVMInstanceView, err error) {

}

func (az *azVirtualMachineScaleSetVMsClient) List(resourceGroupName string, virtualMachineScaleSetName string, filter string, selectParameter string, expand string) (result compute.VirtualMachineScaleSetVMListResult, err error) {

}

func (az *azVirtualMachineScaleSetVMsClient) ListNextResults(lastResults compute.VirtualMachineScaleSetVMListResult) (result compute.VirtualMachineScaleSetVMListResult, err error) {

}

// azRoutesClient implements RoutesClient.
type azRoutesClient struct {
	client compute.RoutesClient
}

func newAzRoutesClient(subscriptionID string, endpoint, servicePrincipalToken *adal.ServicePrincipalToken) *azRoutesClient {
	routesClient := network.NewRoutesClient(subscriptionID)
	routesClient.BaseURI = endpoint
	routesClient.Authorizer = autorest.NewBearerAuthorizer(servicePrincipalToken)
	routesClient.PollingDelay = 5 * time.Second
	configureUserAgent(&routesClient.Client)

	return &azRoutesClient{
		client: routesClient,
	}
}

func (az *azRoutesClient) CreateOrUpdate(resourceGroupName string, routeTableName string, routeName string, routeParameters network.Route, cancel <-chan struct{}) (<-chan network.Route, <-chan error) {

}

func (az *azRoutesClient) Delete(resourceGroupName string, routeTableName string, routeName string, cancel <-chan struct{}) (<-chan autorest.Response, <-chan error) {

}

// azRouteTablesClient implements RouteTablesClient.
type azRouteTablesClient struct {
	client compute.RouteTablesClient
}

func newAzRouteTablesClient(subscriptionID string, endpoint, servicePrincipalToken *adal.ServicePrincipalToken) *azRouteTablesClient {
	routeTablesClient := network.NewRouteTablesClient(subscriptionID)
	routeTablesClient.BaseURI = endpoint
	routeTablesClient.Authorizer = autorest.NewBearerAuthorizer(servicePrincipalToken)
	routeTablesClient.PollingDelay = 5 * time.Second
	configureUserAgent(&routeTablesClient.Client)

	return &azRouteTablesClient{
		client: routeTablesClient,
	}
}

func (az *azRouteTablesClient) CreateOrUpdate(resourceGroupName string, routeTableName string, parameters network.RouteTable, cancel <-chan struct{}) (<-chan network.RouteTable, <-chan error) {

}

func (az *azRouteTablesClient) Get(resourceGroupName string, routeTableName string, expand string) (result network.RouteTable, err error) {

}

// azStorageAccountClient implements StorageAccountClient.
type azStorageAccountClient struct {
	client compute.StorageAccountClient
}

func newAzStorageAccountClient(subscriptionID string, endpoint, servicePrincipalToken *adal.ServicePrincipalToken) *azStorageAccountClient {
	storageAccountClient := storage.NewAccountsClientWithBaseURI(endpoint, subscriptionID)
	storageAccountClient.Authorizer = autorest.NewBearerAuthorizer(servicePrincipalToken)
	storageAccountClient.PollingDelay = 5 * time.Second
	configureUserAgent(&storageAccountClient.Client)

	return &azStorageAccountClient{
		client: storageAccountClient,
	}
}

func (az *azStorageAccountClient) Create(resourceGroupName string, accountName string, parameters storage.AccountCreateParameters, cancel <-chan struct{}) (<-chan storage.Account, <-chan error) {

}

func (az *azStorageAccountClient) Delete(resourceGroupName string, accountName string) (result autorest.Response, err error) {

}

func (az *azStorageAccountClient) ListKeys(resourceGroupName string, accountName string) (result storage.AccountListKeysResult, err error) {

}

func (az *azStorageAccountClient) ListByResourceGroup(resourceGroupName string) (result storage.AccountListResult, err error) {

}

func (az *azStorageAccountClient) GetProperties(resourceGroupName string, accountName string) (result storage.Account, err error) {

}

// azDisksClient implements DisksClient.
type azDisksClient struct {
	client compute.DisksClient
}

func newAzDisksClient(subscriptionID string, endpoint, servicePrincipalToken *adal.ServicePrincipalToken) *azDisksClient {
	disksClient := disk.NewDisksClientWithBaseURI(endpoint, subscriptionID)
	disksClient.Authorizer = autorest.NewBearerAuthorizer(servicePrincipalToken)
	disksClient.PollingDelay = 5 * time.Second
	configureUserAgent(&disksClient.Client)

	return &azDisksClient{
		client: disksClient,
	}
}

func (az *azDisksClient) CreateOrUpdate(resourceGroupName string, diskName string, diskParameter disk.Model, cancel <-chan struct{}) (<-chan disk.Model, <-chan error) {

}

func (az *azDisksClient) Delete(resourceGroupName string, diskName string, cancel <-chan struct{}) (<-chan disk.OperationStatusResponse, <-chan error) {

}

func (az *azDisksClient) Get(resourceGroupName string, diskName string) (result disk.Model, err error) {

}
