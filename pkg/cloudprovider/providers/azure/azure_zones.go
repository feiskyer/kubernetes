/*
Copyright 2016 The Kubernetes Authors.

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
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"

	"github.com/golang/glog"

	"k8s.io/apimachinery/pkg/types"
	"k8s.io/kubernetes/pkg/cloudprovider"
)

const (
	instanceInfoURL = "http://169.254.169.254/metadata/v1/InstanceInfo"
	zoneMetadataURI = "instance/compute/zone"
)

var faultMutex = &sync.Mutex{}
var faultDomain *string

type instanceInfo struct {
	ID           string `json:"ID"`
	UpdateDomain string `json:"UD"`
	FaultDomain  string `json:"FD"`
}

func (az *Cloud) makeZone(zoneNo int) string {
	return fmt.Sprintf("%s-%d", az.Location, zoneNo)
}

// GetZone returns the Zone containing the current failure zone and locality region that the program is running in
func (az *Cloud) GetZone(ctx context.Context) (cloudprovider.Zone, error) {
	zone, err := az.metadata.Text(zoneMetadataURI)
	if err != nil {
		return cloudprovider.Zone{}, err
	}

	if zone == "" {
		glog.V(3).Infof("Availability zone is not enabled for current instance, falling back to fault domain")
		return az.getZoneFromFaultDomain(instanceInfoURL)
	}

	zoneNo, err := strconv.Atoi(zone)
	if err != nil {
		return cloudprovider.Zone{}, fmt.Errorf("failed to parse zone number %q: %v", zone, err)
	}

	return cloudprovider.Zone{
		FailureDomain: az.makeZone(zoneNo),
		Region:        az.Location,
	}, nil
}

// getZoneFromFaultDomain gets fault domain for the instance.
// This is the fallback when availability zone is not enabled for the instance.
func (az *Cloud) getZoneFromFaultDomain(url string) (cloudprovider.Zone, error) {
	faultMutex.Lock()
	defer faultMutex.Unlock()
	if faultDomain == nil {
		var err error
		faultDomain, err = fetchFaultDomain(url)
		if err != nil {
			return cloudprovider.Zone{}, err
		}
	}
	zone := cloudprovider.Zone{
		FailureDomain: *faultDomain,
		Region:        az.Location,
	}
	return zone, nil
}

// GetZoneByProviderID implements Zones.GetZoneByProviderID
// This is particularly useful in external cloud providers where the kubelet
// does not initialize node data.
func (az *Cloud) GetZoneByProviderID(ctx context.Context, providerID string) (cloudprovider.Zone, error) {
	nodeName, err := az.vmSet.GetNodeNameByProviderID(providerID)
	if err != nil {
		return cloudprovider.Zone{}, err
	}

	return az.GetZoneByNodeName(ctx, nodeName)
}

// GetZoneByNodeName implements Zones.GetZoneByNodeName
// This is particularly useful in external cloud providers where the kubelet
// does not initialize node data.
func (az *Cloud) GetZoneByNodeName(ctx context.Context, nodeName types.NodeName) (cloudprovider.Zone, error) {
	return az.vmSet.GetZoneByNodeName(string(nodeName))
}

func fetchFaultDomain(url string) (*string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return readFaultDomain(resp.Body)
}

func readFaultDomain(reader io.Reader) (*string, error) {
	var instanceInfo instanceInfo
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &instanceInfo)
	if err != nil {
		return nil, err
	}
	return &instanceInfo.FaultDomain, nil
}
