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
	"iter"
	"sync"
	"sync/atomic"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/utils/clock"
	"sigs.k8s.io/controller-runtime/pkg/client"

	v1 "sigs.k8s.io/karpenter/pkg/apis/v1"
	"sigs.k8s.io/karpenter/pkg/cloudprovider"
)

// Cluster maintains cluster state that is often needed but expensive to compute.
type Cluster struct {
	kubeClient    client.Client
	cloudProvider cloudprovider.CloudProvider
	clock         clock.Clock
	hasSynced     atomic.Bool

	mu                        sync.RWMutex
	nodes                     map[string]*StateNode           // provider id -> cached node
	bindings                  map[types.NamespacedName]string // pod namespaced named -> node name
	nodeNameToProviderID      map[string]string               // node name -> provider id
	nodeClaimNameToProviderID map[string]string               // node claim name -> provider id
	nodePoolResources         map[string]corev1.ResourceList  // node pool name -> resource list
	daemonSetPods             sync.Map                        // daemonSet -> existing pod

	NodePoolState *NodePoolState

	podAcks                         sync.Map // pod namespaced name -> time when Karpenter first saw the pod as pending
	podsSchedulingAttempted         sync.Map // pod namespaced name -> time when Karpenter tried to schedule a pod
	podsSchedulableTimes            sync.Map // pod namespaced name -> time when it was first marked as able to fit to a node
	podHealthyNodePoolScheduledTime sync.Map // pod namespaced name -> time when pod scheduled to a nodePool that has NodeRegistrationHealthy=true, is marked as able to fit to a node
	podToNodeClaim                  sync.Map // pod namespaced name -> nodeClaim name

	clusterStateMu sync.RWMutex // Separate mutex as this is called in some places that mu is held
	// A monotonically increasing timestamp representing the time state of the
	// cluster with respect to consolidation. This increases when something has
	// changed about the cluster that might make consolidation possible. By recording
	// the state, interested disruption methods can check to see if this has changed to
	// optimize and not try to disrupt if nothing about the cluster has changed.
	clusterState time.Time

	unsyncedTimeMu      sync.Mutex
	unsyncedStartTime   time.Time
	lastUnsyncedLogTime time.Time
	antiAffinityPods    sync.Map // pod namespaced name -> *corev1.Pod of pods that have required anti affinities
}

func NewCluster(clk clock.Clock, client client.Client, cloudProvider cloudprovider.CloudProvider) *Cluster {
	_ = "STUB: not implemented"
	return nil
}

// Synced validates that the NodeClaims and the Nodes that are stored in the apiserver
// have the same representation in the cluster state. This is to ensure that our view
// of the cluster is as close to correct as it can be when we begin to perform operations
// utilizing the cluster state as our source of truth
//
//nolint:gocyclo
func (c *Cluster) Synced(ctx context.Context) (synced bool) {
	_ = "STUB: not implemented"
	// Set the metric depending on the result of the Synced() call
	return false
}

// We want to log every 10s when the cluster hasn't synced for 30s which is long enough for us to think there is an issue

// Set the metric to whatever the result of the Synced() call is

// If the cluster state has already synced once, then we assume that objects are kept internally consistent
// with each other to avoid having to continually re-check that we have fully captured the same view
// of cluster state that controller-runtime has

// Check to see if any node claim doesn't have a provider ID. If it doesn't, then the nodeclaim hasn't been
// launched, and we need to wait to see what the resolved values are before continuing.

// If we haven't synced before, then we need to make sure that our internal cache is fully hydrated
// before we start doing operations against the state
// Because we get so many NodeClaims from this response, we are not DeepCopying the cached data here
// DO NOT MUTATE NodeClaims in this function as this will affect the underlying cached NodeClaim

// Because we get so many Nodes from this response, we are not DeepCopying the cached data here
// DO NOT MUTATE Nodes in this function as this will affect the underlying cached Node

// Check to see if any node claim doesn't have a provider ID. If it doesn't, then the nodeclaim hasn't been
// launched, and we need to wait to see what the resolved values are before continuing.

// The names tracked in-memory should at least have all the data that is in the api-server
// This doesn't ensure that the two states are exactly aligned (we could still not be tracking a node
// that exists in the cluster state but not in the apiserver) but it ensures that we have a state
// representation for every node/nodeClaim that exists on the apiserver

// ForPodsWithAntiAffinity calls the supplied function once for each pod with required anti affinity terms that is
// currently bound to a node. The pod returned may not be up-to-date with respect to status, however since the
// anti-affinity terms can't be modified, they will be correct.
func (c *Cluster) ForPodsWithAntiAffinity(fn func(p *corev1.Pod, n *corev1.Node) bool) {
	_ = "STUB: not implemented"
	return
}

// if we receive the node deletion event before the pod deletion event, this can happen

// Nodes returns an iterator which iterates over the state nodes in the cluster under a read-lock. It is not safe to
// store the state.StateNode object and it should only be accessed while the iterator is active.
func (c *Cluster) Nodes() iter.Seq[*StateNode] { _ = "STUB: not implemented"; return nil }

// DeepCopyNodes creates a DeepCopy of all state nodes.
// NOTE: This is very inefficient so this should only be used when DeepCopying is absolutely necessary
func (c *Cluster) DeepCopyNodes() StateNodes { _ = "STUB: not implemented"; return *new(StateNodes) }

// IsNodeNominated returns true if the given node was expected to have a pod bound to it during a recent scheduling
// batch
func (c *Cluster) IsNodeNominated(providerID string) bool { _ = "STUB: not implemented"; return false }

// NominateNodeForPod records that a node was the target of a pending pod during a scheduling batch
func (c *Cluster) NominateNodeForPod(ctx context.Context, providerID string) {
	_ = "STUB: not implemented"
	return
}

// extends nomination window if already nominated

// UnmarkForDeletion removes the marking on the node as a node the controller intends to delete
func (c *Cluster) UnmarkForDeletion(providerIDs ...string) { _ = "STUB: not implemented"; return }

// MarkForDeletion marks the node as pending deletion in the internal cluster state
func (c *Cluster) MarkForDeletion(providerIDs ...string) { _ = "STUB: not implemented"; return }

func (c *Cluster) UpdateNodeClaim(nodeClaim *v1.NodeClaim) { _ = "STUB: not implemented"; return }

// If the nodeclaim has a providerID, create a StateNode for it, and populate the data.
// We only need to do this for a nodeclaim with a providerID as nodeclaims without provider IDs haven't
// been launched yet.

// Update nodepool state with NodeClaim

// If the nodeclaim hasn't launched yet, we want to add it into cluster state to ensure
// that we're not racing with the internal cache for the cluster, assuming the node doesn't exist.

func (c *Cluster) DeleteNodeClaim(name string) { _ = "STUB: not implemented"; return }

func (c *Cluster) UpdateNode(ctx context.Context, node *corev1.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// If we know that we own this node, we shouldn't allow the providerID to be empty

// If we have a managed node with no instance type label that hasn't been initialized,
// we need to wait until the instance type label gets propagated on it

func (c *Cluster) DeleteNode(name string) { _ = "STUB: not implemented"; return }

func (c *Cluster) UpdatePod(ctx context.Context, pod *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Cluster) NodeClaimExists(nodeClaimName string) bool {
	_ = "STUB: not implemented"
	return false
}

// AckPods marks the pod as acknowledged for scheduling from the provisioner. This is only done once per-pod.
func (c *Cluster) AckPods(pods ...*corev1.Pod) { _ = "STUB: not implemented"; return }

// store the value as now only if it doesn't exist.

// PodAckTime will return the time the pod was first seen in our cache.
func (c *Cluster) PodAckTime(podKey types.NamespacedName) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

// MarkPodSchedulingDecisions keeps track of when we first tried to schedule a pod to a node.
// It updates podHealthyNodePoolScheduledTime for pods scheduled against nodePool that have
// NodeRegistrationHealthy=true. This also marks when the pod is first seen as schedulable for pod metrics.
// We'll only emit a metric for a pod if we haven't done it before.
// nolint:gocyclo
func (c *Cluster) MarkPodSchedulingDecisions(ctx context.Context, podErrors map[*corev1.Pod]error, npPods map[string][]*corev1.Pod, ncPods map[string][]*corev1.Pod) {
	_ = "STUB: not implemented"
	return
}

// delete podsSchedulableTimes and podHealthyNodePoolScheduledTime for pods that have pod errors

// If we already attempted this, we don't need to emit another metric.

// We should have ACK'd the pod.

// Swallow errors if we can't get the nodepool

// Skip pods that are already bound to a node (e.g. pods from deleting nodes
// included in the scheduling simulation for capacity planning). Storing a new
// timestamp for already-bound pods would cause negative metric values since
// their PodScheduled LastTransitionTime is in the past.

// If we already attempted this, we don't need to emit another metric.

// We should have ACK'd the pod.

// If the pod is scheduled to a nodePool and if the nodePool has NodeRegistrationHealthy=true
// then mark the time when we thought it can schedule to now.

// If the pod was scheduled to a healthy nodePool earlier but is now getting scheduled to an
// unhealthy one then we need to delete its entry from the map because it will not schedule successfully

func (c *Cluster) UpdatePodToNodeClaimMapping(ncPods map[string][]*corev1.Pod) {
	_ = "STUB: not implemented"
	return
}

// PodSchedulingDecisionTime returns when Karpenter first decided if a pod could schedule a pod in scheduling simulations.
// This returns 0, false if Karpenter never made a decision on the pod.
func (c *Cluster) PodSchedulingDecisionTime(podKey types.NamespacedName) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

// PodSchedulingSuccessTime returns when Karpenter first thought it could schedule a pod in its scheduling simulation.
// This returns 0, false if the pod was never considered in scheduling as a pending pod.
func (c *Cluster) PodSchedulingSuccessTime(podKey types.NamespacedName) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

// PodNodeClaimMapping returns the nodeClaim against which the pod is simulated to get scheduled
func (c *Cluster) PodNodeClaimMapping(podKey types.NamespacedName) string {
	_ = "STUB: not implemented"
	return ""
}

// PodSchedulingSuccessTimeRegistrationHealthyCheck returns when Karpenter first thought it could schedule a pod in its scheduling simulation.
// This returns 0, false if the pod was never considered in scheduling as a pending pod.
func (c *Cluster) PodSchedulingSuccessTimeRegistrationHealthyCheck(podKey types.NamespacedName) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func (c *Cluster) DeletePod(podKey types.NamespacedName) { _ = "STUB: not implemented"; return }

func (c *Cluster) ClearPodSchedulingMappings(podKey types.NamespacedName) {
	_ = "STUB: not implemented"
	return
}

// MarkUnconsolidated marks the cluster state as being unconsolidated.  This should be called in any situation where
// something in the cluster has changed such that the cluster may have moved from a non-consolidatable to a consolidatable
// state.
func (c *Cluster) MarkUnconsolidated() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// ConsolidationState returns a timestamp of the last time that the cluster state with respect to consolidation changed.
// If nothing changes, this timestamp resets after five minutes to force watchers that use this to defer work to
// occasionally revalidate that nothing external (e.g. an instance type becoming available) has changed that now makes
// it possible for them to operate. Time was chosen as the type here as it allows comparisons using the built-in
// monotonic clock.
func (c *Cluster) ConsolidationState() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// time.Time uses a monotonic clock for these comparisons

// This ensures that at least once every 5 minutes we consider consolidating our cluster in case something else has
// changed (e.g. instance type availability) that we can't detect which would allow consolidation to occur.

func (c *Cluster) NodePoolResourcesFor(nodePoolName string) corev1.ResourceList {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList)
}

// Reset the cluster state for unit testing
func (c *Cluster) Reset() { _ = "STUB: not implemented"; return }

// sets the cluster to be synced or unsynced for unit testing
func (c *Cluster) SetSynced(state bool) { _ = "STUB: not implemented"; return }

func (c *Cluster) GetDaemonSetPod(daemonset *appsv1.DaemonSet) *corev1.Pod {
	_ = "STUB: not implemented"
	return nil
}

func (c *Cluster) UpdateDaemonSet(ctx context.Context, daemonset *appsv1.DaemonSet) error {
	_ = "STUB: not implemented"
	return nil

	// Scope down this call to only select the pods in this namespace that specifically match the DaemonSet
	// Because we get so many pods from this response, we are not DeepCopying the cached data here
	// DO NOT MUTATE pods in this function as this will affect the underlying cached pod
}

func (c *Cluster) DeleteDaemonSet(key types.NamespacedName) { _ = "STUB: not implemented"; return }

// WARNING
// Everything under this section of code assumes that you have already held a lock when you are calling into these functions
// and explicitly modifying the cluster state. If you do not hold the cluster state lock before calling any of these helpers
// you will hit race conditions and data corruption

func (c *Cluster) newStateFromNodeClaim(nodeClaim *v1.NodeClaim, oldNode *StateNode) *StateNode {
	_ = "STUB: not implemented"
	return nil
}

// Cleanup the old nodeClaim with its old providerID if its providerID changes
// This can happen since nodes don't get created with providerIDs. Rather, CCM picks up the
// created node and injects the providerID into the spec.providerID

func (c *Cluster) cleanupNodeClaim(name string) { _ = "STUB: not implemented"; return }

// Delete the node claim from the nodeClaimNameToProviderID in the case that the provider ID hasn't resolved
// yet. This ensures that if a nodeClaim is created and then deleted before it was able to launch that
// this is cleaned up.

// Delete the NodeClaim that is tracked in NodePoolState

func (c *Cluster) newStateFromNode(ctx context.Context, node *corev1.Node, oldNode *StateNode) (*StateNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Cleanup the old node with its old providerID if its providerID changes
// This can happen since nodes don't get created with providerIDs. Rather, CCM picks up the
// created node and injects the providerID into the spec.providerID

func (c *Cluster) cleanupNode(name string) { _ = "STUB: not implemented"; return }

// nolint:gocyclo
func (c *Cluster) updateNodePoolResources(oldNode, newNode *StateNode) {
	_ = "STUB: not implemented"
	return
}

// Garbage collect any NodePool keys that no longer have any resources assigned to them.
// We do this when there are no longer any NodeClaims that map to this NodePool
// so that we don't leak NodePool keys in our nodePoolResources map

func (c *Cluster) populateVolumeLimits(ctx context.Context, n *StateNode) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Cluster) populateResourceRequests(ctx context.Context, n *StateNode) error {
	_ = "STUB: not implemented"
	return nil
}

// updateNodeUsageFromPod is called every time a reconcile event occurs for the pod. If the pods binding has changed
// (unbound to bound), we need to update the resource requests on the node.
func (c *Cluster) updateNodeUsageFromPod(ctx context.Context, pod *corev1.Pod) error {
	_ = "STUB: not implemented"
	// nothing to do if the pod isn't bound, checking early allows avoiding unnecessary locking
	return nil
}

// the node must exist for us to update the resource requests on the node

func (c *Cluster) updateNodeUsageFromPodCompletion(podKey types.NamespacedName) {
	_ = "STUB: not implemented"
	return
}

// we didn't think the pod was bound, so we weren't tracking it and don't need to do anything

// we weren't tracking the node yet, so nothing to do

func (c *Cluster) cleanupOldBindings(pod *corev1.Pod) { _ = "STUB: not implemented"; return }

// we are already tracking the pod binding, so nothing to update

// the pod has switched nodes, this can occur if a pod name was re-used, and it was deleted/re-created rapidly,
// binding to a different node the second time

// we were tracking the old node, so we need to reduce its capacity by the amount of the pod that left

// new pod binding has occurred

func (c *Cluster) updatePodAntiAffinities(pod *corev1.Pod) {
	_ = "STUB: not implemented"
	// We intentionally don't track inverse anti-affinity preferences. We're not
	// required to enforce them so it just adds complexity for very little
	// value. The problem with them comes from the relaxation process, the pod
	// we are relaxing is not the pod with the anti-affinity term.
	return
}

func (c *Cluster) triggerConsolidationOnChange(old, new *StateNode) {
	_ = "STUB: not implemented"
	return
}

// If either the old node or new node are mocked

// HasSynced returns whether the cluster state has been synchronized at least once.
func (c *Cluster) HasSynced() bool { _ = "STUB: not implemented"; return false }
