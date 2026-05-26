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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"

	"sigs.k8s.io/karpenter/pkg/scheduling"
)

type TopologyType byte

const (
	TopologyTypeSpread TopologyType = iota
	TopologyTypePodAffinity
	TopologyTypePodAntiAffinity
)

func (t TopologyType) String() string { _ = "STUB: not implemented"; return "" }

// TopologyGroup is used to track pod counts that match a selector by the topology domain (e.g. SELECT COUNT(*) FROM pods GROUP BY(topology_ke
type TopologyGroup struct {
	// Hashed Fields
	Key        string
	Type       TopologyType
	maxSkew    int32
	minDomains *int32
	namespaces sets.Set[string]
	selector   labels.Selector

	// NOTE: This is actually nillable since there's no API server validation to require it on affinity / TSC terms. A term without it is a no-op.
	rawSelector *metav1.LabelSelector
	nodeFilter  TopologyNodeFilter

	// Index
	owners       map[types.UID]struct{} // Pods that have this topology as a scheduling rule
	domains      map[string]int32       // TODO(ellistarn) explore replacing with a minheap
	emptyDomains sets.Set[string]       // domains for which we know that no pod exists
}

func NewTopologyGroup(
	topologyType TopologyType,
	topologyKey string,
	pod *corev1.Pod,
	namespaces sets.Set[string],
	labelSelector *metav1.LabelSelector,
	maxSkew int32,
	minDomains *int32,
	taintPolicy *corev1.NodeInclusionPolicy,
	affinityPolicy *corev1.NodeInclusionPolicy,
	domainGroup TopologyDomainGroup,
) *TopologyGroup {
	_ = "STUB: not implemented"
	// the nil *TopologyNodeFilter always passes which is what we need for affinity/anti-affinity
	return nil
}

func (t *TopologyGroup) Get(pod *corev1.Pod, podDomains, nodeDomains *scheduling.Requirement) (*scheduling.Requirement, sets.Set[string]) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *TopologyGroup) Record(domains ...string) { _ = "STUB: not implemented"; return }

// Counts returns true if the pod would count for the topology, given that it schedule to a node with the provided
// requirements
func (t *TopologyGroup) Counts(pod *corev1.Pod, taints []corev1.Taint, requirements scheduling.Requirements, compatibilityOptions ...option.Function[scheduling.CompatibilityOptions]) bool {
	_ = "STUB: not implemented"
	return false
}

// Register ensures that the topology is aware of the given domain names.
func (t *TopologyGroup) Register(domains ...string) { _ = "STUB: not implemented"; return }

func (t *TopologyGroup) Unregister(domains ...string) { _ = "STUB: not implemented"; return }

func (t *TopologyGroup) AddOwner(key types.UID) { _ = "STUB: not implemented"; return }

func (t *TopologyGroup) RemoveOwner(key types.UID) { _ = "STUB: not implemented"; return }

func (t *TopologyGroup) IsOwnedBy(key types.UID) bool { _ = "STUB: not implemented"; return false }

// Hash is used so we can track single topologies that affect multiple groups of pods.  If a deployment has 100x pods
// with self anti-affinity, we track that as a single topology with 100 owners instead of 100x topologies.
func (t *TopologyGroup) Hash() uint64 { _ = "STUB: not implemented"; return 0 }

// hashSelector is a specialized hash function for a metav1.LabelSelector. Due to https://github.com/mitchellh/hashstructure/issues/36
// repeated requirements inside a label selector can result in hash collisions when using SlicesAsSets. This function provides the same
// behavior while avoiding that bug by storing the individual expression hashes in a set, ensuring there aren't repeated elements.
//
// NOTE: Although repeated elements typically won't occur, they can occur on k8s 1.34+ when using matchLabelKeys since both Karpenter
// and the API server inject an expression.
func hashSelector(selector *metav1.LabelSelector) uint64 { _ = "STUB: not implemented"; return 0 }

// nextDomainTopologySpread returns a scheduling.Requirement that includes a node domain that a pod should be scheduled to,
// along with the number of valid domains that satisfied the maxSkew constraint.
// If there are multiple eligible domains, we return any random domain that satisfies the `maxSkew` configuration.
// If there are no eligible domains, we return a `DoesNotExist` requirement, implying that we could not satisfy the topologySpread requirement.
// nolint:gocyclo
func (t *TopologyGroup) nextDomainTopologySpread(pod *corev1.Pod, podDomains, nodeDomains *scheduling.Requirement) (*scheduling.Requirement, sets.Set[string]) {
	_ = "STUB: not implemented"
	// min count is calculated across all domains
	return nil, nil
}

// We special-case kubernetes.io/hostname primarily for new NodeClaims since their domain won't be registered until we Add() them

// t.domains[hostName] produces a 0 value for new NodeClaims

// Because Karpenter can always create a new domain for hostname, we assume the global miniumum is always zero
// This means we can just check whether our current count is less than or equal to the skew to check if the domain is valid

// If we are explicitly selecting on specific node domains ("In" requirement),
// this is going to be more efficient to iterate through
// This is particularly useful when considering the hostname topology key that can have a
// lot of t.domains but only a single nodeDomain

// but we can only choose from the node domains

// comment from kube-scheduler regarding the viable choices to schedule to based on skew is:
// 'existing matching num' + 'if self-match (1 or 0)' - 'global min matching num' <= 'maxSkew'

// avoids an error message about 'zone in [""]', preferring 'zone in []'

func (t *TopologyGroup) domainMinCount(domains *scheduling.Requirement) int32 {
	_ = "STUB: not implemented"
	// hostname based topologies always have a min pod count of zero since we can create one
	return 0
}

// determine our current min count

// nolint:gocyclo
func (t *TopologyGroup) nextDomainAffinity(pod *corev1.Pod, podDomains *scheduling.Requirement, nodeDomains *scheduling.Requirement) *scheduling.Requirement {
	_ = "STUB: not implemented"
	return nil
}

// We special-case kubernetes.io/hostname primarily for new NodeClaims since their domain won't be registered until we Add() them

// t.domains[hostName] produces a 0 value for new NodeClaims

// If we are explicitly selecting on specific node domains ("In" requirement),
// this is going to be more efficient to iterate through
// This is particularly useful when considering the hostname topology key that can have a
// lot of t.domains but only a single nodeDomain

// If pod is self-selecting and no pod has been scheduled yet OR the pods that have scheduled are
// incompatible with our podDomains, we can pick a domain at random to bootstrap scheduling.

// First try to find a domain that is within the intersection of pod/node domains. In the case of an in-flight node
// this causes us to pick the domain that the existing in-flight node is already in if possible instead of picking
// a random viable domain.

// and if there are no node domains, just return the first random domain that is viable

// anyCompatiblePodDomain validates whether any t.domain is compatible with our podDomains
// This is only useful in affinity checking because it tells us whether we can schedule the pod
// to the current node since it is the first pod that exists in the TopologyGroup OR all other domains
// in the TopologyGroup are incompatible with the podDomains
func (t *TopologyGroup) anyCompatiblePodDomain(podDomains *scheduling.Requirement) bool {
	_ = "STUB: not implemented"
	return false
}

// nolint:gocyclo
func (t *TopologyGroup) nextDomainAntiAffinity(podDomains, nodeDomains *scheduling.Requirement) *scheduling.Requirement {
	_ = "STUB: not implemented"
	return nil
}

// pods with anti-affinity must schedule to a domain where there are currently none of those pods (an empty
// domain). If there are none of those domains, then the pod can't schedule and we don't need to walk this
// list of domains.  The use case where this optimization is really great is when we are launching nodes for
// a deployment of pods with self anti-affinity.  The domains map here continues to grow, and we continue to
// fully scan it each iteration.

// We special-case kubernetes.io/hostname primarily for new NodeClaims since their domain won't be registered until we Add() them

// t.domains[hostName] produces a 0 value for new NodeClaims

// If we are explicitly selecting on specific node domains ("In" requirement) and the number of node domains
// is less than our empty domains, this is going to be more efficient to iterate through
// This is particularly useful when considering the hostname topology key that can have a
// lot of t.domains but only a single nodeDomain

// selects returns true if the given pod is selected by this topology
func (t *TopologyGroup) selects(pod *corev1.Pod) bool { _ = "STUB: not implemented"; return false }
