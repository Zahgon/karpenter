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

package v1

import (
	"context"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/sets"
)

var (
	SupportedNodeSelectorOps = sets.NewString(
		string(v1.NodeSelectorOpIn),
		string(v1.NodeSelectorOpNotIn),
		string(v1.NodeSelectorOpExists),
		string(v1.NodeSelectorOpDoesNotExist),
		string(v1.NodeSelectorOpGt),
		string(v1.NodeSelectorOpLt),
		string(NodeSelectorOpGte),
		string(NodeSelectorOpLte),
	)

	SupportedReservedResources = sets.NewString(
		v1.ResourceCPU.String(),
		v1.ResourceMemory.String(),
		v1.ResourceEphemeralStorage.String(),
		"pid",
	)

	SupportedEvictionSignals = sets.NewString(
		"memory.available",
		"nodefs.available",
		"nodefs.inodesFree",
		"imagefs.available",
		"imagefs.inodesFree",
		"pid.available",
	)
)

type taintKeyEffect struct {
	OwnerKey string         //nolint:kubeapilinter
	Effect   v1.TaintEffect //nolint:kubeapilinter
}

func (in *NodeClaimTemplateSpec) validateTaints() (errs error) {
	_ = "STUB: not implemented"
	return nil
}

func validateTaintsField(taints []v1.Taint, existing map[taintKeyEffect]struct{}, fieldName string) error {
	_ = "STUB: not implemented"
	return nil
}

// Validate OwnerKey

// Validate Value

// Validate effect

// Check for duplicate OwnerKey/Effect pairs

// This function is used by the NodeClaim validation webhook to verify the nodepool requirements.
// When this function is called, the nodepool's requirements do not include the requirements from labels.
// NodeClaim requirements only support well known labels.
func (in *NodeClaimTemplateSpec) validateRequirements(ctx context.Context) (errs error) {
	_ = "STUB: not implemented"
	return nil
}

func ValidateRequirement(ctx context.Context, requirement NodeSelectorRequirementWithMinValues) error {
	_ = "STUB: not implemented" //nolint:gocyclo
	return nil
}

// Validate that at least one value is valid for well-known labels with known values

// ValidateWellKnownValues checks if the requirement has well known values.
// An error will cause a NodePool's Readiness to transition to False.
// It returns an error if all values are invalid.
// It returns an error if there are not enough valid values to satisfy min values for a requirement with known values.
// It logs if invalid values are found but valid values can be used.
func validateWellKnownValues(ctx context.Context, requirement NodeSelectorRequirementWithMinValues) error {
	_ = "STUB: not implemented"
	// If the key doesn't have well-known values or the operator is not In, nothing to validate
	return nil
}

// If the key doesn't have well-known values defined, nothing to validate

// If there are only invalid values, set an error to transition the nodepool's readiness to false

// If there are valid values, but the minimum number of values is not met, set an error to prevent the nodepool from going ready

// If there are valid and invalid values, log the invalid values and proceed with valid values
