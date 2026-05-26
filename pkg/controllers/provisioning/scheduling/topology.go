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
	"context"

	"github.com/awslabs/operatorpkg/option"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/selection"
	"k8s.io/apimachinery/pkg/util/sets"
	"sigs.k8s.io/controller-runtime/pkg/client"

	v1 "sigs.k8s.io/karpenter/pkg/apis/v1"
	"sigs.k8s.io/karpenter/pkg/cloudprovider"
	"sigs.k8s.io/karpenter/pkg/controllers/state"
	"sigs.k8s.io/karpenter/pkg/scheduling"
)

type Topology struct {
	kubeClient       client.Client
	preferencePolicy PreferencePolicy
	// Both the topologyGroups and inverseTopologies are maps of the hash from TopologyGroup.Hash() to the topology group
	// itself. This is used to allow us to store one topology group that tracks the topology of many pods instead of
	// having a 1<->1 mapping between topology groups and pods owned/selected by that group.
	topologyGroups map[uint64]*TopologyGroup
	// Anti-affinity works both ways (if a zone has a pod foo with anti-affinity to a pod bar, we can't schedule bar to
	// that zone, even though bar has no anti affinity terms on it. For this to work, we need to separately track the
	// topologies of pods with anti-affinity terms, so we can prevent scheduling the pods they have anti-affinity to
	// in some cases.
	inverseTopologyGroups map[uint64]*TopologyGroup
	// The universe of domains by topology key
	domainGroups map[string]TopologyDomainGroup
	// excludedPods are the pod UIDs of pods that are excluded from counting.  This is used so we can simulate
	// moving pods to prevent them from being double counted.
	excludedPods sets.Set[string]
	cluster      *state.Cluster
	stateNodes   []*state.StateNode
}

func NewTopology(
	ctx context.Context,
	kubeClient client.Client,
	cluster *state.Cluster,
	stateNodes []*state.StateNode,
	nodePools []*v1.NodePool,
	instanceTypes map[string][]*cloudprovider.InstanceType,
	pods []*corev1.Pod,
	opts ...Options,
) (*Topology, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// these are the pods that we intend to schedule, so if they are currently in the cluster we shouldn't count them for
// topology purposes

func buildDomainGroups(nodePools []*v1.NodePool, instanceTypes map[string][]*cloudprovider.InstanceType) map[string]TopologyDomainGroup {
	_ = "STUB: not implemented"
	return nil
}

// We need to intersect the instance type requirements with the current nodePool requirements.  This
// ensures that something like zones from an instance type don't expand the universe of valid domains.

// topologyError allows lazily generating the error string in the topology error.  If a pod fails to schedule, most often
// we are only interested in the fact that it failed to schedule and not why.
type topologyError struct {
	topology    *TopologyGroup
	podDomains  *scheduling.Requirement
	nodeDomains *scheduling.Requirement
}

func (t topologyError) Error() string { _ = "STUB: not implemented"; return "" }

// Update unregisters the pod as the owner of all affinities and then creates any new topologies based on the pod spec
// registered the pod as the owner of all associated affinities, new or old.  This allows Update() to be called after
// relaxation of a preference to properly break the topology <-> owner relationship so that the preferred topology will
// no longer influence scheduling.
func (t *Topology) Update(ctx context.Context, p *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

// Avoid recomputing topology counts if we've already seen this group

// Record records the topology changes given that pod p schedule on a node with the given requirements
func (t *Topology) Record(p *corev1.Pod, taints []corev1.Taint, requirements scheduling.Requirements, compatibilityOptions ...option.Function[scheduling.CompatibilityOptions]) {
	_ = "STUB: not implemented"
	// once we've committed to a domain, we record the usage in every topology that cares about it
	return
}

// for anti-affinity topologies we need to block out all possible domains that the pod could land in

// but for affinity & topology spread, we can only record the domain if we know the specific domain we land in

// for anti-affinities, we record where the pods could be, even if
// requirements haven't collapsed to a single value.

// AddRequirements tightens the input requirements by adding additional requirements that are being enforced by topology spreads
// affinities, anti-affinities or inverse anti-affinities.  The nodeHostname is the hostname that we are currently considering
// placing the pod on.  It returns these newly tightened requirements, or an error in the case of a set of requirements that
// cannot be satisfied.
func (t *Topology) AddRequirements(p *corev1.Pod, taints []corev1.Taint, podRequirements, nodeRequirements scheduling.Requirements, compatibilityOptions ...option.Function[scheduling.CompatibilityOptions]) (scheduling.Requirements, error) {
	_ = "STUB: not implemented"
	return *new(scheduling.Requirements), nil
}

// GetTopologyZoneConstraints returns the set of valid zones from all topology constraints
// that use the zone topology key for the given pod, along with whether the constraints are satisfiable.
func (t *Topology) GetTopologyZoneConstraints(p *corev1.Pod, podRequirements scheduling.Requirements) (sets.Set[string], bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Register is used to register a domain as available across topologies for the given topology key.
func (t *Topology) Register(topologyKey string, domain string) { _ = "STUB: not implemented"; return }

// Unregister is used to unregister a domain as available across topologies for the given topology key.
func (t *Topology) Unregister(topologyKey string, domain string) { _ = "STUB: not implemented"; return }

// updateInverseAffinities is used to identify pods with anti-affinity terms so we can track those topologies.  We
// have to look at every pod in the cluster as there is no way to query for a pod with anti-affinity terms.
func (t *Topology) updateInverseAffinities(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// don't count the pod we are excluding

// updateInverseAntiAffinity is used to track topologies of inverse anti-affinities. Here the domains & counts track the
// pods with the anti-affinity.
func (t *Topology) updateInverseAntiAffinity(ctx context.Context, pod *corev1.Pod, domains map[string]string) error {
	_ = "STUB: not implemented"
	// We intentionally don't track inverse anti-affinity preferences. We're not
	// required to enforce them so it just adds complexity for very little
	// value.  The problem with them comes from the relaxation process, the pod
	// we are relaxing is not the pod with the anti-affinity term.
	return nil
}

// countDomains initializes the topology group by registereding any well known domains and performing pod counts
// against the cluster for any existing pods.
//
//nolint:gocyclo
func (t *Topology) countDomains(ctx context.Context, tg *TopologyGroup) error {
	_ = "STUB: not implemented"
	return nil
}

// collect the pods from all the specified namespaces (don't see a way to query multiple namespaces
// simultaneously)

// capture new domain values from existing nodes that may not have any pods selected by the topology group
// scheduled to them already
// Note: long term we should handle this when constructing the domain groups, but that would require domain groups
// to handle affinity in addition to taints / tolerations.

// ignore state nodes which are tracking in-flight NodeClaims

// ignore the node if it doesn't match the topology group

// sort our pods by the node they are scheduled to

// pod is excluded for counting purposes

// no need to look up the node since we already have it

// Pods that cannot be evicted can be leaked in the API Server after
// a Node is removed. Since pod bindings are immutable, these pods
// cannot be recovered, and will be deleted by the pod lifecycle
// garbage collector. These pods are not running, and should not
// impact future topology calculations.

// assign back to previous node so we can hopefully re-use these in the next iteration

// Kubelet sets the hostname label, but the node may not be ready yet so there is no label.  We fall back and just
// treat the node name as the label.  It probably is in most cases, but even if not we at least count the existence
// of the pods in some domain, even if not in the correct one.  This is needed to handle the case of pods with
// self-affinity only fulfilling that affinity if all domains are empty.

// Don't include pods if node doesn't contain domain https://kubernetes.io/docs/concepts/workloads/pods/pod-topology-spread-constraints/#conventions

// nodes may or may not be considered for counting purposes for topology spread constraints depending on if they
// are selected by the pod's node selectors and required node affinities.  If these are unset, the node always counts.

func (t *Topology) newForTopologies(p *corev1.Pod) []*TopologyGroup {
	_ = "STUB: not implemented"
	return nil
}

// newForAffinities returns a list of topology groups that have been constructed based on the input pod and required/preferred affinity terms
func (t *Topology) newForAffinities(ctx context.Context, p *corev1.Pod) ([]*TopologyGroup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// No affinity defined

// include both soft and hard affinity terms

// include both soft and hard antiaffinity terms

// build topologies

// buildNamespaceList constructs a unique list of namespaces consisting of the pod's namespace and the optional list of
// namespaces and those selected by the namespace selector
func (t *Topology) buildNamespaceList(ctx context.Context, namespace string, namespaces []string, selector *metav1.LabelSelector) (sets.Set[string], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getMatchingTopologies returns a sorted list of topologies that either control the scheduling of pod p, or for which
// the topology selects pod p and the scheduling of p affects the count per topology domain
func (t *Topology) getMatchingTopologies(p *corev1.Pod, taints []corev1.Taint, requirements scheduling.Requirements, compatibilityOptions ...option.Function[scheduling.CompatibilityOptions]) []*TopologyGroup {
	_ = "STUB: not implemented"
	return nil
}

func TopologyListOptions(namespace string, labelSelector *metav1.LabelSelector) *client.ListOptions {
	_ = "STUB: not implemented"
	return nil
}

func mapOperator(operator metav1.LabelSelectorOperator) selection.Operator {
	_ = "STUB: not implemented"
	return *new(selection.Operator)
}

// this shouldn't occur as we cover all valid cases of LabelSelectorOperator that the API allows.  If it still
// does occur somehow we'll panic just later when the requirement throws an error.,

func IgnoredForTopology(p *corev1.Pod) bool { _ = "STUB: not implemented"; return false }
