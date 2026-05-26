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
	"github.com/awslabs/operatorpkg/option"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/sets"

	v1 "sigs.k8s.io/karpenter/pkg/apis/v1"
)

// Requirements are an efficient set representation under the hood. Since its underlying
// types are slices and maps, this type should not be used as a pointer.
type Requirements map[string]*Requirement

func NewRequirements(requirements ...*Requirement) Requirements {
	_ = "STUB: not implemented"
	return *new(Requirements)
}

// NewRequirements constructs requirements from NodeSelectorRequirementWithMinValues
func NewNodeSelectorRequirementsWithMinValues(requirements ...v1.NodeSelectorRequirementWithMinValues) Requirements {
	_ = "STUB: not implemented"
	return *new(Requirements)
}

// NewRequirements constructs requirements from NodeSelectorRequirements
func NewNodeSelectorRequirements(requirements ...corev1.NodeSelectorRequirement) Requirements {
	_ = "STUB: not implemented"
	return *new(Requirements)
}

// NewLabelRequirements constructs requirements from labels
func NewLabelRequirements(labels map[string]string) Requirements {
	_ = "STUB: not implemented"
	return *new(Requirements)
}

// NewPodRequirements constructs requirements from a pod and treats any preferred requirements as required.
func NewPodRequirements(pod *corev1.Pod) Requirements {
	_ = "STUB: not implemented"
	return *new(Requirements)
}

// NewStrictPodRequirements constructs requirements from a pod and only includes true requirements (not preferences).
func NewStrictPodRequirements(pod *corev1.Pod) Requirements {
	_ = "STUB: not implemented"
	return *new(Requirements)
}

type podRequirementType byte

const (
	podRequirementTypeAll = iota
	podRequirementTypeRequiredOnly
)

func newPodRequirements(pod *corev1.Pod, typ podRequirementType) Requirements {
	_ = "STUB: not implemented"
	return *new(Requirements)
}

// The legal operators for pod affinity and anti-affinity are In, NotIn, Exists, DoesNotExist.
// Select heaviest preference and treat as a requirement. An outer loop will iteratively unconstrain them if unsatisfiable.

// Select first requirement. An outer loop will iteratively remove OR requirements if unsatisfiable

// HasPreferredNodeAffinity returns true if the pod has a preferred node affinity term
func HasPreferredNodeAffinity(p *corev1.Pod) bool { _ = "STUB: not implemented"; return false }

func (r Requirements) NodeSelectorRequirements() []v1.NodeSelectorRequirementWithMinValues {
	_ = "STUB: not implemented"
	return nil
}

// Add requirements to provided requirements. Mutates existing requirements
func (r Requirements) Add(requirements ...*Requirement) { _ = "STUB: not implemented"; return }

// Keys returns unique set of the label keys from the requirements
func (r Requirements) Keys() sets.Set[string] { _ = "STUB: not implemented"; return nil }

func (r Requirements) Values() []*Requirement { _ = "STUB: not implemented"; return nil }

func (r Requirements) Has(key string) bool { _ = "STUB: not implemented"; return false }

func (r Requirements) Get(key string) *Requirement { _ = "STUB: not implemented"; return nil }

// If not defined, allow any values with the exists operator

type CompatibilityOptions struct {
	AllowUndefined sets.Set[string]
}

var AllowUndefinedWellKnownLabels = func(options *CompatibilityOptions) {
	options.AllowUndefined = v1.WellKnownLabels
}

func (r Requirements) IsCompatible(requirements Requirements, options ...option.Function[CompatibilityOptions]) bool {
	_ = "STUB: not implemented"
	return false
}

// Compatible ensures the provided requirements can loosely be met.
func (r Requirements) Compatible(requirements Requirements, options ...option.Function[CompatibilityOptions]) error {
	_ = "STUB: not implemented"
	return nil
}

// Custom Labels must intersect, but if not defined are denied.

// break early so we only report the first error

// Well Known Labels must intersect, but if not defined, are allowed.

func getSuffix(key string) string { _ = "STUB: not implemented"; return "" }

func labelHint(r Requirements, key string, allowedUndefined sets.Set[string]) string {
	_ = "STUB: not implemented"
	return ""
}

// badKeyError allows lazily generating the error string in the case of a bad key error. When requirements fail
// to match, we are most often interested in the failure and not why it fails.
type badKeyError struct {
	key      string
	incoming *Requirement
	existing *Requirement
}

func (b badKeyError) Error() string { _ = "STUB: not implemented"; return "" }

// intersectKeys is much faster and allocates less han getting the two key sets separately and intersecting them
func (r Requirements) intersectKeys(rhs Requirements) sets.Set[string] {
	_ = "STUB: not implemented"
	return nil
}

// Intersects returns errors if the requirements don't have overlapping values, undefined keys are allowed
func (r Requirements) Intersects(requirements Requirements) (errs error) {
	_ = "STUB: not implemented"
	return nil
}

// where the incoming requirement has operator { NotIn, DoesNotExist }

// and the existing requirement has operator { NotIn, DoesNotExist }

func (r Requirements) HasMinValues() bool { _ = "STUB: not implemented"; return false }

func (r Requirements) String() string { _ = "STUB: not implemented"; return "" }
