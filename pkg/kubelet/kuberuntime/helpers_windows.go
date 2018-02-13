// +build windows

/*
Copyright 2018 The Kubernetes Authors.

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

package kuberuntime

const (
	// Taken from https://docs.microsoft.com/en-us/virtualization/windowscontainers/manage-containers/resource-controls
	minShares     = 1
	maxShares     = 10000
	milliCPUToCPU = 1000

	// The default shares on Process isolation is 5000.
	defaultSharesPerCPUProcess = 5000
	// The default shares for Hyper-V isolation is 10.
	defaultSharesPerCPUHyperV = 10
)

// milliCPUToShares converts milliCPU to CPU shares
func milliCPUToShares(milliCPU int64, hyperv bool) int64 {
	var defaultSharesPerCPU int64 = defaultSharesPerCPUProcess
	if hyperv {
		defaultSharesPerCPU = defaultSharesPerCPUHyperV
	}

	if milliCPU == 0 {
		// Return here to really match kernel default for zero milliCPU.
		return defaultSharesPerCPU
	}

	// Conceptually (milliCPU / milliCPUToCPU) * sharesPerCPU, but factored to improve rounding.
	shares := (milliCPU * defaultSharesPerCPU) / milliCPUToCPU
	if shares < minShares {
		return minShares
	}
	if shares > maxShares {
		return maxShares
	}
	return shares
}
