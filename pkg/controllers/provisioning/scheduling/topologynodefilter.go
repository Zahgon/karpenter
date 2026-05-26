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

	"sigs.k8s.io/karpenter/pkg/scheduling"
)

// TopologyNodeFilter is used to determine if a given actual node or scheduling node matches the pod's node selectors
// and required node affinity terms.  This is used with topology spread constraints to determine if the node should be
// included for topology counting purposes. This is only used with topology spread constraints as affinities/anti-affinities
// always count across all nodes. A nil or zero-value TopologyNodeFilter behaves well and the filter returns true for
// all nodes.
type TopologyNodeFilter struct {
	Requirements   []scheduling.Requirements
	TaintPolicy    corev1.NodeInclusionPolicy
	AffinityPolicy corev1.NodeInclusionPolicy
	Tolerations    []corev1.Toleration
}

func MakeTopologyNodeFilter(p *corev1.Pod, taintPolicy corev1.NodeInclusionPolicy, affinityPolicy corev1.NodeInclusionPolicy) TopologyNodeFilter {
	_ = "STUB: not implemented"
	return *new(TopologyNodeFilter)
}

// if we only have a label selector, that's the only requirement that must match

// otherwise, we need to match the combination of label selector and any term of the required node affinities since
// those terms are OR'd together

// Matches returns true if the TopologyNodeFilter doesn't prohibit node from the participating in the topology
func (t TopologyNodeFilter) Matches(taints []corev1.Taint, requirements scheduling.Requirements, compatibilityOptions ...option.Function[scheduling.CompatibilityOptions]) bool {
	_ = "STUB: not implemented"
	return false
}

// MatchesRequirements returns true if the TopologyNodeFilter doesn't prohibit a node with the requirements from
// participating in the topology. This method allows checking the requirements from a scheduling.NodeClaim to see if the
// node we will soon create participates in this topology.
func (t TopologyNodeFilter) matchesRequirements(requirements scheduling.Requirements, compatibilityOptions ...option.Function[scheduling.CompatibilityOptions]) bool {
	_ = "STUB: not implemented"
	// no requirements, so it always matches
	return false
}

// these are an OR, so if any passes the filter passes
