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

package resources

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

var Node = v1.ResourceName("nodes")

// RequestsForPods returns the total resources of a variadic list of podspecs.
func RequestsForPods(pods ...*v1.Pod) v1.ResourceList {
	_ = "STUB: not implemented"
	return *new(v1.ResourceList)
}

// LimitsForPods returns the total resources of a variadic list of podspecs
func LimitsForPods(pods ...*v1.Pod) v1.ResourceList {
	_ = "STUB: not implemented"
	return *new(v1.ResourceList)
}

// Merge the resources from the variadic into a single v1.ResourceList
func Merge(resources ...v1.ResourceList) v1.ResourceList {
	_ = "STUB: not implemented"
	return *new(v1.ResourceList)
}

// MergeInto sums the resources from src into dest, modifying dest. If you need to repeatedly sum
// multiple resource lists, it allocates less to continually sum into an existing list as opposed to
// constructing a new one for each sum like Merge
func MergeInto(dest v1.ResourceList, src v1.ResourceList) v1.ResourceList {
	_ = "STUB: not implemented"
	return *new(v1.ResourceList)
}

func Subtract(lhs, rhs v1.ResourceList) v1.ResourceList {
	_ = "STUB: not implemented"
	return *new(v1.ResourceList)
}

// SubtractFrom subtracts the src v1.ResourceList from the dest v1.ResourceList in-place
func SubtractFrom(dest v1.ResourceList, src v1.ResourceList) { _ = "STUB: not implemented"; return }

// Ceiling computes the effective resource requirements for a given Pod,
// using the same logic as the scheduler.
func Ceiling(pod *v1.Pod) v1.ResourceRequirements {
	_ = "STUB: not implemented"
	return *new(v1.ResourceRequirements)
}

// MaxResources returns the maximum quantities for a given list of resources
func MaxResources(resources ...v1.ResourceList) v1.ResourceList {
	_ = "STUB: not implemented"
	return *new(v1.ResourceList)
}

// Quantity parses the string value into a *Quantity
func Quantity(value string) *resource.Quantity { _ = "STUB: not implemented"; return nil }

// IsZero implements r.IsZero(). This method is provided to make some code a bit cleaner as the Quantity.IsZero() takes
// a pointer receiver and map index expressions aren't addressable, so it can't be called directly.
func IsZero(r resource.Quantity) bool { _ = "STUB: not implemented"; return false }

func Cmp(lhs resource.Quantity, rhs resource.Quantity) int { _ = "STUB: not implemented"; return 0 }

// Fits returns true if the candidate set of resources is less than or equal to the total set of resources.
func Fits(candidate, total v1.ResourceList) bool {
	_ = "STUB: not implemented"
	// If any of the total resource values are negative then the resource will never fit
	return false
}

// String returns a string version of the resource list suitable for presenting in a log
func String(list v1.ResourceList) string { _ = "STUB: not implemented"; return "" }
