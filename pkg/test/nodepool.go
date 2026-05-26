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

package test

import (
	v1 "sigs.k8s.io/karpenter/pkg/apis/v1"
)

// NodePool creates a test NodePool with defaults that can be overridden by overrides.
// Overrides are applied in order, with a last write wins semantic.
func NodePool(overrides ...v1.NodePool) *v1.NodePool { _ = "STUB: not implemented"; return nil }

// NodePools creates homogeneous groups of NodePools
// based on the passed in options, evenly divided by the total NodePools requested
func NodePools(total int, options ...v1.NodePool) []*v1.NodePool {
	_ = "STUB: not implemented"
	return nil
}

// ReplaceRequirements any current requirements on the passed through NodePool with the passed in requirements
// If any of the keys match between the existing requirements and the new requirements, the new requirement with the same
// key will replace the old requirement with that key
func ReplaceRequirements(nodePool *v1.NodePool, reqs ...v1.NodeSelectorRequirementWithMinValues) *v1.NodePool {
	_ = "STUB: not implemented"
	return nil
}

// StaticNodePool creates a test NodePool suitable for static provisioning
// It will keep limits.nodes if provided in overrides, otherwise limits will be nil
func StaticNodePool(overrides ...v1.NodePool) *v1.NodePool {
	_ = "STUB: not implemented"
	// First create the NodePool with all overrides
	return nil
}

// Set limits based on whether nodes limit was provided
