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
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/utils/clock"
	"sigs.k8s.io/controller-runtime/pkg/client"

	v1 "sigs.k8s.io/karpenter/pkg/apis/v1"
	"sigs.k8s.io/karpenter/pkg/cloudprovider"
	"sigs.k8s.io/karpenter/pkg/controllers/state"
	"sigs.k8s.io/karpenter/pkg/events"
	karpopts "sigs.k8s.io/karpenter/pkg/operator/options"
	"sigs.k8s.io/karpenter/pkg/scheduling"
)

type ReservedOfferingMode int

// TODO: Evaluate if another mode should be created for drift. The problem with strict is that it assumes we can run
// multiple scheduling loops to make progress, but if scheduling all pods from the drifted node in a single iteration
// requires fallback, we're at a stalemate. This makes strict a non-starter for drift IMO.
// On the other hand, fallback will result in non-ideal launches when there's constrained capacity. This should be
// rectified by consolidation, but if we can be "right" the at initial launch that would be preferable.
// One potential improvement is a "preferences" type strategy, where we attempt to schedule the pod without fallback
// first. This is an improvement over the current fallback strategy since it ensures all new nodeclaims are attempted,
// before then attempting all nodepools, but it still doesn't address the case when offerings are reserved pessimistically.
// I don't believe there's a solution to this short of the max-flow based instance selection algorithm, which has its own
// drawbacks.
const (
	// ReservedOfferingModeFallbackAlways indicates to the scheduler that the addition of a pod to a nodeclaim which
	// results in all potential reserved offerings being filtered out is allowed (e.g. on-demand / spot fallback).
	ReservedOfferingModeFallback ReservedOfferingMode = iota
	// ReservedOfferingModeStrict indicates that the scheduler should fail to add a pod to a nodeclaim if doing so would
	// prevent it from scheduling to reserved capacity, when it would have otherwise.
	ReservedOfferingModeStrict
)

type PreferencePolicy int

const (
	// PreferencePolicyRespect indicates to the scheduler that it should attempt to respect all preference requirements
	// and topologies. The scheduler will treat all preferences as required at first and then will slowly relax
	// these requirements one at a time until it is able to schedule the pod
	PreferencePolicyRespect PreferencePolicy = iota
	// PreferencePolicyIgnore indicates to the scheduler that it should ignore all preference requirements and
	// topologies. Preferences include preferredDuringSchedulingIgnoredDuringExecution affinities and ScheduleAnyways
	// topologySpreadConstraints
	PreferencePolicyIgnore
)

type options struct {
	reservedOfferingMode    ReservedOfferingMode
	preferencePolicy        PreferencePolicy
	minValuesPolicy         karpopts.MinValuesPolicy
	numConcurrentReconciles int
}

type Options = option.Function[options]

var DisableReservedCapacityFallback = func(opts *options) {
	opts.reservedOfferingMode = ReservedOfferingModeStrict
}

var IgnorePreferences = func(opts *options) {
	opts.preferencePolicy = PreferencePolicyIgnore
}

var NumConcurrentReconciles = func(numConcurrentReconciles int) func(*options) {
	return func(opts *options) {
		opts.numConcurrentReconciles = numConcurrentReconciles
	}
}

var MinValuesPolicy = func(policy karpopts.MinValuesPolicy) func(*options) {
	return func(opts *options) {
		opts.minValuesPolicy = policy
	}
}

func NewScheduler(
	ctx context.Context,
	kubeClient client.Client,
	nodePools []*v1.NodePool,
	cluster *state.Cluster,
	stateNodes []*state.StateNode,
	topology *Topology,
	instanceTypes map[string][]*cloudprovider.InstanceType,
	daemonSetPods []*corev1.Pod,
	recorder events.Recorder,
	clock clock.Clock,
	volumeReqsByPod map[types.UID][]scheduling.Requirements,
	opts ...Options,
) *Scheduler {
	_ = "STUB: not implemented"
	return nil
}

// if any of the nodePools add a taint with a prefer no schedule effect, we add a toleration for the taint
// during preference relaxation

// Pre-filter instance types eligible for NodePools to reduce work done during scheduling loops for pods
// if no templates remain, we still want to build the scheduler so that Karpenter can ack pods which can schedule to existing and in-flight capacity

// cache pod data to avoid having to continually recompute it
// Volume requirements per pod

type PodData struct {
	Requests                 corev1.ResourceList
	Requirements             scheduling.Requirements
	StrictRequirements       scheduling.Requirements
	HasResourceClaimRequests bool
	VolumeRequirements       []scheduling.Requirements // Volume topology requirement alternatives
}

type Scheduler struct {
	uuid                    types.UID // Unique UUID attached to this scheduling loop
	newNodeClaims           []*NodeClaim
	existingNodes           []*ExistingNode
	nodeClaimTemplates      []*NodeClaimTemplate
	remainingResources      map[string]corev1.ResourceList // (NodePool name) -> remaining resources for that NodePool
	daemonOverhead          map[*NodeClaimTemplate]corev1.ResourceList
	daemonHostPortUsage     map[*NodeClaimTemplate]*scheduling.HostPortUsage
	cachedPodData           map[types.UID]*PodData                  // (Pod Namespace/Name) -> pre-computed data for pods to avoid re-computation and memory usage
	volumeReqsByPod         map[types.UID][]scheduling.Requirements // Volume topology requirement alternatives per pod
	preferences             *Preferences
	topology                *Topology
	cluster                 *state.Cluster
	recorder                events.Recorder
	kubeClient              client.Client
	clock                   clock.Clock
	reservationManager      *ReservationManager
	reservedOfferingMode    ReservedOfferingMode
	preferencePolicy        PreferencePolicy
	minValuesPolicy         karpopts.MinValuesPolicy
	numConcurrentReconciles int
}

// DRAError indicates a pod will not be attempted to be scheduled because it has Dynamic Resource Allocation requirements
// that are not yet supported by Karpenter
type DRAError struct {
	error
}

func NewDRAError(err error) DRAError { _ = "STUB: not implemented"; return *new(DRAError) }

func IsDRAError(err error) bool { _ = "STUB: not implemented"; return false }

func (e DRAError) Unwrap() error {
	_ = "STUB: not implemented"

	// Results contains the results of the scheduling operation
	return nil
}

type Results struct {
	NewNodeClaims []*NodeClaim
	ExistingNodes []*ExistingNode
	PodErrors     map[*corev1.Pod]error
}

// Record sends eventing and log messages back for the results that were produced from a scheduling run
// It also nominates nodes in the cluster state based on the scheduling run to signal to other components
// leveraging the cluster state that a previous scheduling run that was recorded is relying on these nodes
func (r Results) Record(ctx context.Context, recorder events.Recorder, cluster *state.Cluster) {
	_ = "STUB: not implemented"
	// Report failures and nominations
	return
}

// Report new nodes, or exit to avoid log spam

// Report in flight newNodes, or exit to avoid log spam

func (r Results) ReservedOfferingErrors() map[*corev1.Pod]error {
	_ = "STUB: not implemented"
	return nil
}

func (r Results) DRAErrors() map[*corev1.Pod]error { _ = "STUB: not implemented"; return nil }

func (r Results) NodePoolToPodMapping() map[string][]*corev1.Pod {
	_ = "STUB: not implemented"
	return nil
}

func (r Results) ExistingNodeToPodMapping() map[string][]*corev1.Pod {
	_ = "STUB: not implemented"
	return nil
}

// Filter out nodes that are not managed

// AllNonPendingPodsScheduled returns true if all pods scheduled.
// We don't care if a pod was pending before consolidation and will still be pending after. It may be a pod that we can't
// schedule at all and don't want it to block consolidation.
func (r Results) AllNonPendingPodsScheduled() bool { _ = "STUB: not implemented"; return false }

// NonPendingPodSchedulingErrors creates a string that describes why pods wouldn't schedule that is suitable for presentation
func (r Results) NonPendingPodSchedulingErrors() string { _ = "STUB: not implemented"; return "" }

// TruncateInstanceTypes filters the result based on the maximum number of instanceTypes that needs
// to be considered. This filters all instance types generated in NewNodeClaims in the Results
func (r Results) TruncateInstanceTypes(ctx context.Context, maxInstanceTypes int) Results {
	_ = "STUB: not implemented"
	return *new(Results)
}

// The InstanceTypeOptions are truncated due to limitations in sending the number of instances to launch API.

// Check if the truncated InstanceTypeOptions in each NewNodeClaim from the results still satisfy the minimum requirements
// If number of InstanceTypes in the NodeClaim cannot satisfy the minimum requirements, add its Pods to error map with reason.

func (s *Scheduler) Solve(ctx context.Context, pods []*corev1.Pod) (Results, error) {
	_ = "STUB: not implemented"
	return *new(Results), nil
}

// We loop trying to schedule unschedulable pods as long as we are making progress.  This solves a few
// issues including pods with affinity to another pod in the batch. We could topo-sort to solve this, but it wouldn't
// solve the problem of scheduling pods where a particular order is needed to prevent a max-skew violation. E.g. if we
// had 5xA pods and 5xB pods were they have a zonal topology spread, but A can only go in one zone and B in another.
// We need to schedule them alternating, A, B, A, B, .... and this solution also solves that as well.

// Reset the metric for the controller, so we don't keep old ids around

// Try the next pod

// We relax the pod all the way the first time we see it
// If we don't schedule it, we store the original pod (with preferences)
// in the queue and give ourselves another chance to schedule it later

// Update the cached podData since the pod was relaxed, and it could have changed its requirement set

func (s *Scheduler) trySchedule(ctx context.Context, p *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

// We should only relax the pod's requirements when the error is not a reserved offering error because the pod may be
// able to schedule later without relaxing constraints. This could occur in this scheduling run, if other NodeClaims
// release the required reservations when constrained, or in subsequent runs. For an example, reference the following
// test: "shouldn't relax preferences when a pod fails to schedule due to a reserved offering error".

// DRA errors are permanent while the IgnoreDRARequests flag is enabled, so we shouldn't attempt to relax
// pod requirements as we don't want to schedule the pod.

// Eventually we won't be able to relax anymore and this while loop will exit

// Update the cached podData since the pod was relaxed, and it could have changed its requirement set

func (s *Scheduler) updateCachedPodData(p *corev1.Pod) { _ = "STUB: not implemented"; return }

// strictPodRequirements is important as it ensures we don't inadvertently restrict the possible pod domains by a
// preferred node affinity.  Only required node affinities can actually reduce pod domains.

// Volume requirements

func (s *Scheduler) add(ctx context.Context, pod *corev1.Pod) error {
	_ = "STUB: not implemented"
	// Check if pod has DRA requirements - if so, return DRA error when IgnoreDRARequests is enabled
	return nil
}

// first try to schedule against an in-flight real node

// Consider using https://pkg.go.dev/container/heap

// Pick existing node that we are about to create

func (s *Scheduler) addToExistingNode(ctx context.Context, pod *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

// determine the volumes that will be mounted if the pod schedules

// Ensure that we always take an earlier successful schedule to keep consistent ordering

// If we set the existingNode to something valid, this means that we successfully scheduled to one of these nodes

func (s *Scheduler) addToInflightNode(ctx context.Context, pod *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

// Ensure that we always take an earlier successful schedule to keep consistent ordering

//nolint:gocyclo
func (s *Scheduler) addToNewNodeClaim(ctx context.Context, pod *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

// if limits have been applied to the nodepool, ensure we filter instance types to avoid violating those limits

// Node limits can be enforced early, since we know exactly how much capacity in nodes will be consumed by any instance type (1 node).

// If the pod is compatible with a NodePool with reserved offerings available, we shouldn't fall back to a NodePool
// with a lower weight. We could consider allowing "fallback" to NodePools with equal weight if they also have
// reserved capacity in the future if scheduling latency becomes an issue.

// A reserved offering error means that any subsequent successful after this NodeClaimTemplate isn't valid

// Ensure that we always take an earlier successful schedule to keep consistent ordering
// We care about this particularly with NewNodeClaims because NodeClaims should be evaluated by weight

// we will launch this nodeClaim and need to track its maximum possible resource usage against our remaining resources

func (s *Scheduler) calculateExistingNodeClaims(ctx context.Context, stateNodes []*state.StateNode, daemonSetPods []*corev1.Pod) {
	_ = "STUB: not implemented"
	// create our existing nodes
	return
}

// getCompatibleDaemonPods filters daemon pods that can schedule to the given node
func (s *Scheduler) getCompatibleDaemonPods(ctx context.Context, node *state.StateNode, taints []corev1.Taint, daemonSetPods []*corev1.Pod) []*corev1.Pod {
	_ = "STUB: not implemented"
	return nil
}

// shouldSkipDaemonPod checks if a daemon pod should be skipped due to DRA requirements
func (s *Scheduler) shouldSkipDaemonPod(ctx context.Context, p *corev1.Pod) bool {
	_ = "STUB: not implemented"
	return false
}

// isDaemonPodCompatibleWithNode checks if a daemon pod is compatible with the node
func (s *Scheduler) isDaemonPodCompatibleWithNode(p *corev1.Pod, taints []corev1.Taint, nodeLabels map[string]string) bool {
	_ = "STUB: not implemented"
	return false
}

// updateRemainingResources updates the remaining resources for the node's nodepool
func (s *Scheduler) updateRemainingResources(node *state.StateNode) {
	_ = "STUB: not implemented"
	// We don't use the status field and instead recompute the remaining resources to ensure we have a consistent view
	// of the cluster during scheduling.  Depending on how node creation falls out, this will also work for cases where
	// we don't create NodeClaim resources.
	return
}

// sortExistingNodes sorts existing nodes with initialized nodes first
func (s *Scheduler) sortExistingNodes() {
	_ = "STUB: not implemented"
	// Order the existing nodes for scheduling with initialized nodes first
	// This is done specifically for consolidation where we want to make sure we schedule to initialized nodes
	// before we attempt to schedule uninitialized ones
	return
}

// computeEffectiveZoneFromPod calculates the effective zone constraint by intersecting
// pod-level zone signals, PVC volume zones, and TSC valid domains. This can be the
// specific zone name if exactly one zone, "flexible" if multiple zones, "none" if no intersection.
//
//nolint:gocyclo
func (s *Scheduler) computeEffectiveZoneFromPod(pod *corev1.Pod) string {
	_ = "STUB: not implemented"
	return ""
}

// volumeZoneReq returns a single Requirement representing the union of zone constraints
// across all volume alternatives. Returns nil if volumes don't constrain zones.
func volumeZoneReq(volumeReqs []scheduling.Requirements) *scheduling.Requirement {
	_ = "STUB: not implemented"
	return nil
}

// parallelizeUntil is an implementation of workqueue.ParallelizeUntil that modifies the
// doWorkPiece so that a worker always finishes its work when it pulls a piece off of pieces
// The function returns a bool that represents whether the worker should continue doing work
// or whether the worker should finish
func parallelizeUntil(workers, pieces int, doWorkPiece func(int) bool) {
	_ = "STUB: not implemented"
	return
}

// getDaemonOverhead determines the overhead for each NodeClaimTemplate required for daemons to schedule for any node provisioned by the NodeClaimTemplate
func getDaemonOverhead(ctx context.Context, nodeClaimTemplates []*NodeClaimTemplate, daemonSetPods []*corev1.Pod) map[*NodeClaimTemplate]corev1.ResourceList {
	_ = "STUB: not implemented"
	return nil
}

// Exclude daemon pods with DRA requirements when IgnoreDRARequests is enabled

// getDaemonHostPortUsage determines requested host ports for DaemonSet pods, given a NodeClaimTemplate
func getDaemonHostPortUsage(ctx context.Context, nodeClaimTemplates []*NodeClaimTemplate, daemonSetPods []*corev1.Pod) map[*NodeClaimTemplate]*scheduling.HostPortUsage {
	_ = "STUB: not implemented"
	return nil
}

// gather compatible DaemonSet pods for the NodeClaimTemplate

// Exclude daemon pods with DRA requirements when IgnoreDRARequests is enabled

// isDaemonPodCompatible determines if the daemon pod is compatible with the NodeClaimTemplate for daemon scheduling
func isDaemonPodCompatible(nodeClaimTemplate *NodeClaimTemplate, pod *corev1.Pod) bool {
	_ = "STUB: not implemented"
	return false

	// Add a toleration for PreferNoSchedule since a daemon pod shouldn't respect the preference
}

// We don't consider pod preferences for scheduling requirements since we know that pod preferences won't matter with Daemonset scheduling

// If relaxing the Node Affinity term didn't succeed, then this DaemonSet can't schedule to this NodePool
// We don't consider other forms of relaxation here since we don't consider pod affinities/anti-affinities
// when considering DaemonSet schedulability

// subtractMax returns the remaining resources after subtracting the max resource quantity per instance type. To avoid
// overshooting out, we need to pessimistically assume that if e.g. we request a 2, 4 or 8 CPU instance type
// that the 8 CPU instance type is all that will be available.  This could cause a batch of pods to take multiple rounds
// to schedule.
func subtractMax(remaining corev1.ResourceList, instanceTypes []*cloudprovider.InstanceType) corev1.ResourceList {
	_ = "STUB: not implemented"
	// shouldn't occur, but to be safe
	return *new(corev1.ResourceList)
}

// filterByRemainingResources is used to filter out instance types that if launched would exceed the nodepool limits
func filterByRemainingResources(instanceTypes []*cloudprovider.InstanceType, remaining corev1.ResourceList) []*cloudprovider.InstanceType {
	_ = "STUB: not implemented"
	return nil
}

// if the instance capacity is greater than the remaining quantity for this resource
