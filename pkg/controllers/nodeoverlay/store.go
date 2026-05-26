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

package nodeoverlay

import (
	"sync/atomic"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/sets"

	"sigs.k8s.io/karpenter/pkg/apis/v1alpha1"
	"sigs.k8s.io/karpenter/pkg/cloudprovider"
)

type priceUpdate struct {
	OverlayUpdate *string
	lowestWeight  *int32
}

type capacityUpdate struct {
	OverlayUpdate                 corev1.ResourceList
	lowestWeightCapacityResources corev1.ResourceList
	lowestWeight                  *int32
}

type instanceTypeUpdate struct {
	Price    map[string]*priceUpdate
	Capacity *capacityUpdate
}
type InstanceTypeStore struct {
	store atomic.Pointer[internalInstanceTypeStore]
}

func NewInstanceTypeStore() *InstanceTypeStore { _ = "STUB: not implemented"; return nil }

func (s *InstanceTypeStore) UpdateStore(updatedStore *internalInstanceTypeStore) {
	_ = "STUB: not implemented"
	return
}

func (s *InstanceTypeStore) ApplyAll(nodePoolName string, its []*cloudprovider.InstanceType) ([]*cloudprovider.InstanceType, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *InstanceTypeStore) Apply(nodePoolName string, it *cloudprovider.InstanceType) (*cloudprovider.InstanceType, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// InstanceTypeStore manages instance type updates for node pools.
// It maintains a nested mapping structure where:
//   - First level:  nodePoolName -> map of instance updates
//   - Second level: instanceName -> specific update configurations
//
// The store is used to:
//   - Track instance type modifications per node pool
//   - Validate instance configurations
//   - Update instance properties for scheduling decisions
type internalInstanceTypeStore struct {
	updates            map[string]map[string]*instanceTypeUpdate // nodePoolName -> (instanceName -> updates)
	evaluatedNodePools sets.Set[string]                          // The set of NodePools that were evaluated to construct this InstanceTypeStore instance
}

func newInternalInstanceTypeStore() *internalInstanceTypeStore {
	_ = "STUB: not implemented"
	return nil
}

// Apply takes a node pool name and instance type, and returns a modified copy of the instance type
// with any stored updates applied. It uses a selective copy-on-write strategy to minimize memory usage:
// - Shared: Requirements and Overhead (never modified, safe to share)
// - Selective copy: Offerings (only copied if price overlay applied)
// - Selective copy: Capacity (only copied if capacity overlay applied)
func (s *internalInstanceTypeStore) apply(nodePoolName string, it *cloudprovider.InstanceType) *cloudprovider.InstanceType {
	_ = "STUB: not implemented"
	return nil
}

// Create a shallow copy of the instance type, sharing immutable fields

// Shared - never modified
// Shared - never modified

// Handle capacity overlay - only deep copy if we're modifying it

// ApplyCapacityOverlay replaces Capacity with a new merged map (original untouched)

// Handle offerings - copy-on-write only for offerings that need price overlay

// Shared - not modified

// applyPriceOverlays creates a new offerings slice with selective copying:
// - Offerings that need price overlay are copied and mutated
// - Offerings without overlay share the original pointer
// This minimizes allocations while ensuring each node pool has independent pricing.
func (s *internalInstanceTypeStore) applyPriceOverlays(offerings cloudprovider.Offerings, priceUpdates map[string]*priceUpdate) cloudprovider.Offerings {
	_ = "STUB: not implemented"
	return *new(cloudprovider.Offerings)
}

// This offering needs modification - create a copy

// Shared - requirements are immutable

// Not modified - share the pointer

// updateInstanceTypeCapacity add a new Capacity overlay update to the associated instance type.
// NOTE: This method does not perform conflict validation. The callee must check for conflicts first.
func (i *internalInstanceTypeStore) updateInstanceTypeCapacity(nodePoolName string, instanceTypeName string, nodeOverlay v1alpha1.NodeOverlay) {
	_ = "STUB: not implemented"
	return
}

func (i *internalInstanceTypeStore) isCapacityUpdateConflicting(nodePoolName string, instanceTypeName string, nodeOverlay v1alpha1.NodeOverlay) bool {
	_ = "STUB: not implemented"
	return false
}

// IMPORTANT: This logic assumes NodeOverlays are processed in descending order by weight.

// updateInstanceTypeOffering add a new Price overlay update to the associated instance type.
// NOTE: This method does not perform conflict validation. The callee must check for conflicts first.
func (i *internalInstanceTypeStore) updateInstanceTypeOffering(nodePoolName string, instanceTypeName string, nodeOverlay v1alpha1.NodeOverlay, offerings cloudprovider.Offerings) {
	_ = "STUB: not implemented"
	return
}

func (i *internalInstanceTypeStore) isOfferingUpdateConflicting(nodePoolName string, instanceTypeName string, of *cloudprovider.Offering, nodeOverlay v1alpha1.NodeOverlay) bool {
	_ = "STUB: not implemented"
	return false
}

// IMPORTANT: This logic assumes NodeOverlays are processed in descending order by weight.

func (s *InstanceTypeStore) Reset() { _ = "STUB: not implemented"; return }
