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

package scheduling

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/sets"

	v1 "sigs.k8s.io/karpenter/pkg/apis/v1"
)

//go:generate go tool -modfile=../../go.tools.mod controller-gen object:headerFile="../../hack/boilerplate.go.txt" paths="."

// Requirement is an efficient represenatation of corev1.NodeSelectorRequirement
// +k8s:deepcopy-gen=true
type Requirement struct {
	Key        string
	complement bool
	values     sets.Set[string]
	gte        *int // inclusive lower bound (Gt is converted to Gte)
	lte        *int // inclusive upper bound (Lt is converted to Lte)
	MinValues  *int
}

// NewRequirementWithFlexibility constructs new requirement from the combination of key, values, minValues and the operator that
// connects the keys and values. GT and LT operators are canonicalized to GTE and LTE respectively.
// nolint:gocyclo
func NewRequirementWithFlexibility(key string, operator corev1.NodeSelectorOperator, minValues *int, values ...string) *Requirement {
	_ = "STUB: not implemented"
	return nil
}

// This is a super-common case, so optimize for it an inline everything.

// prevalidated

// Gt MaxInt matches nothing

// canonicalize GT N to GTE N+1

// prevalidated
// canonicalize LT N to LTE N-1

// prevalidated

// prevalidated

func NewRequirement(key string, operator corev1.NodeSelectorOperator, values ...string) *Requirement {
	_ = "STUB: not implemented"
	return nil
}

// BoundedNodeSelectorRequirements handles the case where both gte and lte exist.
// Unlike other operators which intersect into a single values set, bounds are stored
// separately and must be serialized as two distinct NodeSelectorRequirements.
// This method is separate from NodeSelectorRequirement() to avoid slice allocations
// in the common case where only one bound exists.
func (r *Requirement) BoundedNodeSelectorRequirements() []v1.NodeSelectorRequirementWithMinValues {
	_ = "STUB: not implemented"
	return nil
}

func (r *Requirement) NodeSelectorRequirement() v1.NodeSelectorRequirementWithMinValues {
	_ = "STUB: not implemented"
	return *new(v1.NodeSelectorRequirementWithMinValues)
}

// Intersection constraints the Requirement from the incoming requirements
// nolint:gocyclo
func (r *Requirement) Intersection(requirement *Requirement) *Requirement {
	_ = "STUB: not implemented"
	// Complement
	return nil
}

// Boundaries

// Values

// Remove boundaries for concrete sets

// nolint:gocyclo
// HasIntersection is a more efficient implementation of Intersection
// It validates whether there is an intersection between the two requirements without actually creating the sets
// This prevents the garbage collector from having to spend cycles cleaning up all of these created set objects
func (r *Requirement) HasIntersection(requirement *Requirement) bool {
	_ = "STUB: not implemented"
	return false
}

// Both requirements have a complement

// Only one requirement has a complement

// Both requirements are non-complement requirements

func (r *Requirement) Any() string { _ = "STUB: not implemented"; return "" }

//nolint:gosec

// Has returns true if the requirement allows the value
func (r *Requirement) Has(value string) bool { _ = "STUB: not implemented"; return false }

func (r *Requirement) Values() []string { _ = "STUB: not implemented"; return nil }

func (r *Requirement) Insert(items ...string) { _ = "STUB: not implemented"; return }

func (r *Requirement) Operator() corev1.NodeSelectorOperator {
	_ = "STUB: not implemented"
	return *new(corev1.NodeSelectorOperator)
}

// corev1.NodeSelectorOpGt and corev1.NodeSelectorOpLt are treated as "Exists" with bounds

func (r *Requirement) Len() int { _ = "STUB: not implemented"; return 0 }

func (r *Requirement) String() string { _ = "STUB: not implemented"; return "" }

func withinBounds(valueAsString string, gte, lte *int) bool {
	_ = "STUB: not implemented"
	return false
}

// If bounds are set, non integer values are invalid

func minIntPtr(a, b *int) *int { _ = "STUB: not implemented"; return nil }

func maxIntPtr(a, b *int) *int { _ = "STUB: not implemented"; return nil }
