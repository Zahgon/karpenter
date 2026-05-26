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

package state

import (
	"context"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/utils/clock"

	"sigs.k8s.io/controller-runtime/pkg/client"

	"sigs.k8s.io/karpenter/pkg/events"

	v1 "sigs.k8s.io/karpenter/pkg/apis/v1"
	"sigs.k8s.io/karpenter/pkg/scheduling"
	"sigs.k8s.io/karpenter/pkg/utils/pdb"
)

type PodBlockEvictionError struct {
	error
}

func NewPodBlockEvictionError(err error) *PodBlockEvictionError {
	_ = "STUB: not implemented"
	return nil
}

func IsPodBlockEvictionError(err error) bool { _ = "STUB: not implemented"; return false }

func IgnorePodBlockEvictionError(err error) error { _ = "STUB: not implemented"; return nil }

//go:generate go tool -modfile=../../../go.tools.mod controller-gen object:headerFile="../../../hack/boilerplate.go.txt" paths="."

// StateNodes is a typed version of a list of *Node
// nolint: revive
type StateNodes []*StateNode

// Active filters StateNodes that are not in a MarkedForDeletion state
func (n StateNodes) Active() StateNodes { _ = "STUB: not implemented"; return *new(StateNodes) }

// Deleting filters StateNodes that are in a MarkedForDeletion state
func (n StateNodes) Deleting() StateNodes { _ = "STUB: not implemented"; return *new(StateNodes) }

// Pods gets the pods assigned to all StateNodes based on the kubernetes api-server bindings
func (n StateNodes) Pods(ctx context.Context, kubeClient client.Client) ([]*corev1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (n StateNodes) CurrentlyReschedulablePods(ctx context.Context, kubeClient client.Client, clk clock.Clock, recorder events.Recorder) ([]*corev1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// StateNode is a cached version of a node in the cluster that maintains state which is expensive to compute every time it's
// needed.  This currently contains node utilization across all the allocatable resources, but will soon be used to
// compute topology information.
// +k8s:deepcopy-gen=true
// nolint: revive
type StateNode struct {
	Node      *corev1.Node
	NodeClaim *v1.NodeClaim

	// daemonSetRequests is the total amount of resources that have been requested by daemon sets. This allows users
	// of the Node to identify the remaining resources that we expect future daemonsets to consume.
	daemonSetRequests map[types.NamespacedName]corev1.ResourceList
	daemonSetLimits   map[types.NamespacedName]corev1.ResourceList

	podRequests map[types.NamespacedName]corev1.ResourceList
	podLimits   map[types.NamespacedName]corev1.ResourceList

	hostPortUsage *scheduling.HostPortUsage
	volumeUsage   *scheduling.VolumeUsage

	// TODO remove this when v1alpha5 APIs are deprecated. With v1 APIs Karpenter relies on the existence
	// of the karpenter.sh/disruption taint to know when a node is marked for deletion.
	markedForDeletion bool
	nominatedUntil    metav1.Time
}

func NewNode() *StateNode { _ = "STUB: not implemented"; return nil }

func (in *StateNode) ShallowCopy() *StateNode { _ = "STUB: not implemented"; return nil }

func (in *StateNode) Name() string { _ = "STUB: not implemented"; return "" }

// ProviderID is the key that is used to map this StateNode
// If the Node and NodeClaim have a providerID, this should map to a real providerID
// If the Node does not have a providerID, this will map to the node name
func (in *StateNode) ProviderID() string { _ = "STUB: not implemented"; return "" }

// Pods gets the pods assigned to the Node based on the kubernetes api-server bindings
func (in *StateNode) Pods(ctx context.Context, kubeClient client.Client) ([]*corev1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ValidateNodeDisruptable returns an error if the StateNode cannot be disrupted
// This checks all associated StateNode internals, node labels, and do-not-disrupt annotations on the node.
// ValidateNodeDisruptable takes in a recorder to emit events on the nodeclaims when the state node is not a candidate
//
//nolint:gocyclo
func (in *StateNode) ValidateNodeDisruptable(clk clock.Clock) error {
	_ = "STUB: not implemented"
	return nil
}

// skip the node if it is nominated by a recent provisioning pass to be the target of a pending pod.

// check whether the node has the NodePool label

// ValidatePodDisruptable returns an error if the StateNode contains a pod that cannot be disrupted
// This checks associated PDBs and do-not-disrupt annotations for each pod on the node.
// ValidatePodDisruptable takes in a recorder to emit events on the nodeclaims when the state node is not a candidate
//
//nolint:gocyclo
func (in *StateNode) ValidatePodsDisruptable(ctx context.Context, kubeClient client.Client, pdbs pdb.Limits, clk clock.Clock, recorder events.Recorder) ([]*corev1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// We only consider pods that are actively running for "karpenter.sh/do-not-disrupt"
// This means that we will allow Mirror Pods and DaemonSets to block disruption using this annotation

// CurrentlyReschedulablePods gets the pods assigned to the Node that are currently reschedulable based on the kubernetes api-server bindings
func (in *StateNode) CurrentlyReschedulablePods(ctx context.Context, kubeClient client.Client, clk clock.Clock, recorder events.Recorder) ([]*corev1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (in *StateNode) HostName() string { _ = "STUB: not implemented"; return "" }

func (in *StateNode) Annotations() map[string]string {
	_ = "STUB: not implemented"
	// If the nodeclaim exists and the state node isn't initialized
	// use the nodeclaim representation of the annotations
	return nil
}

func (in *StateNode) Labels() map[string]string {
	_ = "STUB: not implemented"
	// If the nodeclaim exists and the state node isn't registered
	// use the nodeclaim representation of the labels
	return nil
}

func (in *StateNode) Taints() []corev1.Taint {
	_ = "STUB: not implemented"
	// If we have a managed node that isn't registered, we should use its NodeClaim
	// representation of taints. Likewise, if we don't have a Node representation for this
	// providerID in our state, we should also just use the NodeClaim since this is all that we have
	return nil
}

// We reject any well-known ephemeral taints and startup taints attached to this node until
// the node is initialized. Without this, if the taint is generic and re-appears on the node for a
// different reason (e.g. the node is cordoned) we will assume that pods can schedule against the
// node in the future incorrectly.

func (in *StateNode) Registered() bool {
	_ = "STUB: not implemented"
	// Node is managed by Karpenter, so we can check for the Registered label
	return false
}

// Nodes not managed by Karpenter are always considered Registered

func (in *StateNode) Initialized() bool {
	_ = "STUB: not implemented"
	// Node is managed by Karpenter, so we can check for the Initialized label
	return false
}

// Nodes not managed by Karpenter are always considered Initialized

func (in *StateNode) Capacity() corev1.ResourceList {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList)
}

// Override any zero quantity values in the node status

// A StateNode will always have a capacity of 1 node.

func (in *StateNode) Allocatable() corev1.ResourceList {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList)
}

// Override any zero quantity values in the node status

// Available is allocatable minus anything allocated to pods.
func (in *StateNode) Available() corev1.ResourceList {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList)
}

func (in *StateNode) DaemonSetRequests() corev1.ResourceList {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList)
}

func (in *StateNode) DaemonSetLimits() corev1.ResourceList {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList)
}

func (in *StateNode) HostPortUsage() *scheduling.HostPortUsage {
	_ = "STUB: not implemented"
	return nil
}

func (in *StateNode) VolumeUsage() *scheduling.VolumeUsage { _ = "STUB: not implemented"; return nil }

func (in *StateNode) PodRequests() corev1.ResourceList {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList)
}

func (in *StateNode) PodLimits() corev1.ResourceList {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList)
}

func (in *StateNode) MarkedForDeletion() bool {
	_ = "STUB: not implemented"
	// The Node is marked for deletion if:
	//  1. The Node has MarkedForDeletion set
	//  2. The Node has a NodeClaim counterpart and is actively deleting (or the nodeclaim is marked as terminating)
	//  3. The Node has no NodeClaim counterpart and is actively deleting
	return false
}

func (in *StateNode) Deleted() bool { _ = "STUB: not implemented"; return false }

func (in *StateNode) Nominate(ctx context.Context, clk clock.Clock) {
	_ = "STUB: not implemented"
	return
}

func (in *StateNode) Nominated(clk clock.Clock) bool { _ = "STUB: not implemented"; return false }

func (in *StateNode) Managed() bool { _ = "STUB: not implemented"; return false }

func (in *StateNode) updateForPod(ctx context.Context, kubeClient client.Client, pod *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

// if it's a daemonset, we track what it has requested separately

func (in *StateNode) cleanupForPod(podKey types.NamespacedName) { _ = "STUB: not implemented"; return }

func nominationWindow(ctx context.Context) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// RequireNoScheduleTaint will add/remove the karpenter.sh/disruption:NoSchedule taint from the candidates.
// This is used to enforce no taints at the beginning of disruption, and
// to add/remove taints while executing a disruption action.
// nolint:gocyclo
func RequireNoScheduleTaint(ctx context.Context, kubeClient client.Client, addTaint bool, nodes ...*StateNode) error {
	_ = "STUB: not implemented"
	return nil
}

// If the StateNode is Karpenter owned and only has a nodeclaim, or is not owned by
// Karpenter, thus having no nodeclaim, don't touch the node.

// If the node already has the taint, continue to the next

// Node is being deleted, so no need to remove taint as the node will be gone soon.
// This ensures that the disruption controller doesn't modify taints that the Termination
// controller is also modifying

// If the taint is present and we want to remove the taint, remove it.

// otherwise, add it.

// If the taint key is present (but with a different value or effect), remove it.

// We use client.MergeFromWithOptimisticLock because patching a list with a JSON merge patch
// can cause races due to the fact that it fully replaces the list on a change
// Here, we are updating the taint list

// ClearNodeClaimsCondition will remove the conditionType from the NodeClaim status of the provided statenodes
func ClearNodeClaimsCondition(ctx context.Context, kubeClient client.Client, conditionType string, nodes ...*StateNode) error {
	_ = "STUB: not implemented"
	return nil
}
