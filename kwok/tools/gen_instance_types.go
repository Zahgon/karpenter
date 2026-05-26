/*
Copyright The Kubernetes Authors.

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

package main

import (
	"encoding/json"
	"fmt"
	"os"

	corev1 "k8s.io/api/core/v1"

	kwok "sigs.k8s.io/karpenter/kwok/cloudprovider"
)

var (
	KwokZones = []string{"test-zone-a", "test-zone-b", "test-zone-c", "test-zone-d"}
)

func makeGenericInstanceTypeName(cpu, memFactor int, arch string, os corev1.OSName) string {
	_ = "STUB: not implemented"
	return ""
}

// cpu

// standard

// memory

// exotic

func priceFromResources(resources corev1.ResourceList) float64 { _ = "STUB: not implemented"; return 0 }

// case ResourceGPUVendorA, ResourceGPUVendorB:
// 	price += 1.0

func constructGenericInstanceTypes() []kwok.InstanceTypeOptions {
	_ = "STUB: not implemented"
	return nil
}

// Construct instance type details, then construct offerings.

func main() {
	opts := constructGenericInstanceTypes()
	output, err := json.MarshalIndent(opts, "", "    ")
	if err != nil {
		fmt.Printf("could not marshal generated instance types to JSON: %v\n", err)
		os.Exit(1)
	}
	fmt.Print(string(output))
}
