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

package fake

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/sets"

	v1 "sigs.k8s.io/karpenter/pkg/apis/v1"
	"sigs.k8s.io/karpenter/pkg/cloudprovider"
	"sigs.k8s.io/karpenter/pkg/scheduling"
)

const (
	LabelInstanceSize                           = "size"
	ExoticInstanceLabelKey                      = "special"
	IntegerInstanceLabelKey                     = "integer"
	ResourceGPUVendorA      corev1.ResourceName = "fake.com/vendor-a"
	ResourceGPUVendorB      corev1.ResourceName = "fake.com/vendor-b"
)

func init() {
	v1.WellKnownLabels.Insert(
		LabelInstanceSize,
		ExoticInstanceLabelKey,
		IntegerInstanceLabelKey,
	)
}

func NewInstanceType(options InstanceTypeOptions) *cloudprovider.InstanceType {
	_ = "STUB: not implemented"
	return nil
}

func NewInstanceTypeWithCustomRequirement(options InstanceTypeOptions, customReq *scheduling.Requirement) *cloudprovider.InstanceType {
	_ = "STUB: not implemented"
	return nil
}

// InstanceTypesAssorted create many unique instance types with varying CPU/memory/architecture/OS/zone/capacity type.
func InstanceTypesAssorted() []*cloudprovider.InstanceType { _ = "STUB: not implemented"; return nil }

// InstanceTypes creates instance types with incrementing resources
// 2Gi of RAM and 10 pods for every 1vcpu
// i.e. 1vcpu, 2Gi mem, 10 pods
//
//	2vcpu, 4Gi mem, 20 pods
//	3vcpu, 6Gi mem, 30 pods
func InstanceTypes(total int) []*cloudprovider.InstanceType { _ = "STUB: not implemented"; return nil }

type InstanceTypeOptions struct {
	Name             string
	Offerings        cloudprovider.Offerings
	Architecture     string
	OperatingSystems sets.Set[string]
	Resources        corev1.ResourceList
}

func PriceFromResources(resources corev1.ResourceList) float64 { _ = "STUB: not implemented"; return 0 }
