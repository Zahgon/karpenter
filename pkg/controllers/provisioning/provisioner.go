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

package provisioning

import (
	"context"
	"errors"
	"fmt"

	"github.com/awslabs/operatorpkg/option"
	"github.com/awslabs/operatorpkg/reconciler"
	"github.com/awslabs/operatorpkg/serrors"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/utils/clock"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	v1 "sigs.k8s.io/karpenter/pkg/apis/v1"
	"sigs.k8s.io/karpenter/pkg/cloudprovider"
	scheduler "sigs.k8s.io/karpenter/pkg/controllers/provisioning/scheduling"
	"sigs.k8s.io/karpenter/pkg/controllers/state"
	"sigs.k8s.io/karpenter/pkg/events"
	"sigs.k8s.io/karpenter/pkg/scheduling"
	"sigs.k8s.io/karpenter/pkg/utils/pretty"
)

// LaunchOptions are the set of options that can be used to trigger certain
// actions and configuration during scheduling
type LaunchOptions struct {
	RecordPodNomination bool
	Reason              string
}

// RecordPodNomination causes nominate pod events to be recorded against the node.
func RecordPodNomination(o *LaunchOptions) { _ = "STUB: not implemented"; return }

func WithReason(reason string) func(*LaunchOptions) { _ = "STUB: not implemented"; return nil }

// Provisioner waits for enqueued pods, batches them, creates capacity and binds the pods to the capacity.
type Provisioner struct {
	cloudProvider  cloudprovider.CloudProvider
	kubeClient     client.Client
	batcher        *Batcher[types.UID]
	volumeTopology *scheduler.VolumeTopology
	cluster        *state.Cluster
	recorder       events.Recorder
	cm             *pretty.ChangeMonitor
	clock          clock.Clock
}

func NewProvisioner(kubeClient client.Client, recorder events.Recorder,
	cloudProvider cloudprovider.CloudProvider, cluster *state.Cluster,
	clock clock.Clock,
) *Provisioner {
	_ = "STUB: not implemented"
	return nil
}

func (p *Provisioner) Trigger(uid types.UID) { _ = "STUB: not implemented"; return }

func (p *Provisioner) Name() string { _ = "STUB: not implemented"; return "" }

func (p *Provisioner) Register(_ context.Context, m manager.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Provisioner) Reconcile(ctx context.Context) (result reconciler.Result, err error) {
	_ = "STUB: not implemented"
	return *new(reconciler.Result), nil
}

// Batch pods

// We need to ensure that our internal cluster state mechanism is synced before we proceed
// with making any scheduling decision off of our state nodes. Otherwise, we have the potential to make
// a scheduling decision based on a smaller subset of nodes in our cluster state than actually exist.

// Schedule pods to potential nodes, exit if nothing to do

// CreateNodeClaims launches nodes passed into the function in parallel. It returns a slice of the successfully created node
// names as well as a multierr of any errors that occurred while launching nodes
func (p *Provisioner) CreateNodeClaims(ctx context.Context, nodeClaims []*scheduler.NodeClaim, opts ...option.Function[LaunchOptions]) ([]string, error) {
	_ = "STUB: not implemented"
	// Create capacity and bind pods
	return nil, nil
}

// create a new context to avoid a data race on the ctx variable

// Regardless of if we successfully created the NodeClaim or not, we should release the reservation. If the NodeClaim
// was successfully created, we've updated the active node count in Create already. If we failed, we should release
// the reservation and allow the provisioner to create new NodeClaims in a subsequent attempt.
// NOTE: Only applies to static NodePools since node limits are not supported for dynamic NodePools

func (p *Provisioner) GetPendingPods(ctx context.Context) ([]*corev1.Pod, error) {
	_ = "STUB: not implemented"
	// filter for provisionable pods first, so we don't check for validity/PVCs on pods we won't provision anyway
	// (e.g. those owned by daemonsets)
	return nil, nil
}

// Mark in memory that this pod is unschedulable

// Don't create pod events for pods that are specifically avoiding scheduling to Karpenter-managed capacity

// consolidationWarnings potentially writes logs warning about possible unexpected interactions
// between scheduling constraints and consolidation
// nolint:gocyclo
func (p *Provisioner) consolidationWarnings(ctx context.Context, pods []*corev1.Pod) {
	_ = "STUB: not implemented"
	// We have pending pods that have preferred anti-affinity or topology spread constraints.  These can interact
	// unexpectedly with consolidation, so we warn once per hour when we see these pods.
	return
}

// We reduce the amount of logging that we do per-pod by grouping log lines like this together

var ErrNodePoolsNotFound = errors.New("no nodepools found")

//nolint:gocyclo
func (p *Provisioner) NewScheduler(
	ctx context.Context,
	pods []*corev1.Pod,
	stateNodes []*state.StateNode,
	opts ...scheduler.Options,
) (*scheduler.Scheduler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// nodeTemplates generated from NodePools are ordered by weight
// since they are stored within a slice and scheduling
// will always attempt to schedule on the first nodeTemplate

// Get volume topology requirements WITHOUT modifying pods.
// Volume requirements are passed separately and added to nodeRequirements only.
// Pods that fail volume topology lookup are excluded from scheduling.

// Calculate cluster topology, if a context error occurs, it is wrapped and returned

// Pass volumeReqs to scheduler - added to nodeRequirements for NodeClaim zone selection

func (p *Provisioner) Schedule(ctx context.Context) (scheduler.Results, error) {
	_ = "STUB: not implemented"
	return *new(scheduler.Results), nil
}

// We collect the nodes with their used capacities before we get the list of pending pods. This ensures that
// the node capacities we schedule against are always >= what the actual capacity is at any given instance. This
// prevents over-provisioning at the cost of potentially under-provisioning which will self-heal during the next
// scheduling loop when we launch a new node.  When this order is reversed, our node capacity may be reduced by pods
// that have bound which we then provision new un-needed capacity for.
// -------
// We don't consider the nodes that are MarkedForDeletion since this capacity shouldn't be considered
// as persistent capacity for the cluster (since it will soon be removed). Additionally, we are scheduling for
// the pods that are on these nodes so the MarkedForDeletion node capacity can't be considered.

// Get pods, exit if nothing to do

// Get pods from nodes that are preparing for deletion
// We do this after getting the pending pods so that we undershoot if pods are
// actively migrating from a node that is being deleted
// NOTE: The assumption is that these nodes are cordoned and no additional pods will schedule to them

// nothing to schedule, so just return success

// Timeout the Solve() method after 1m to ensure that we move faster through provisioning

// context errors are ignored because we want to finish provisioning for what has already been scheduled

// A reserved offering error doesn't indicate a pod is unschedulable, just that the scheduling decision was deferred.

// Mark in memory when these pods were marked as schedulable or when we made a decision on the pods

// Only passing existing nodes here and not new nodeClaims because
// these nodeClaims don't have a name until they are created

func (p *Provisioner) Create(ctx context.Context, n *scheduler.NodeClaim, opts ...option.Function[LaunchOptions]) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Update pod to nodeClaim mapping for newly created nodeClaims. We do
// this here because nodeClaim does not have a name until it is created.

// If annotation is missing for any reason, assume that min values wasn't relaxed.

// Update the nodeclaim manually in state to avoid eventual consistency delay races with our watcher.
// This is essential to avoiding races where disruption can create a replacement node, then immediately
// requeue. This can race with controller-runtime's internal cache as it watches events on the cluster
// to then trigger cluster state updates. Triggering it manually ensures that Karpenter waits for the
// internal cache to sync before moving onto another disruption loop.

func instanceTypeList(names []string) string { _ = "STUB: not implemented"; return "" }

// print the first 5 instance types only (indices 0-4)

func (p *Provisioner) getDaemonSetPods(ctx context.Context) ([]*corev1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Replacing retrieved pod affinity with daemonset pod template required node affinity since this is overridden
// by the daemonset controller during pod creation
// https://github.com/kubernetes/kubernetes/blob/c5cf0ac1889f55ab51749798bec684aed876709d/pkg/controller/daemon/util/daemonset_util.go#L176

func (p *Provisioner) Validate(ctx context.Context, pod *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

var KarpenterManagedLabelDoesNotExistError = serrors.Wrap(fmt.Errorf("configured to not run on a Karpenter provisioned node"), "requirement", fmt.Sprintf("%s %s", v1.NodePoolLabelKey, corev1.NodeSelectorOpDoesNotExist))

// validateKarpenterManagedLabelCanExist provides a more clear error message in the event of scheduling a pod that specifically doesn't
// want to run on a Karpenter node (e.g. a Karpenter controller replica).
func validateKarpenterManagedLabelCanExist(p *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

// getVolumeTopologyRequirements collects volume topology requirements for each pod
// WITHOUT modifying the pods. These requirements will be added to nodeRequirements
// (for NodeClaim zone selection) but NOT to pod affinities (for correct TSC counting).
func (p *Provisioner) getVolumeTopologyRequirements(ctx context.Context, pods []*corev1.Pod) ([]*corev1.Pod, map[types.UID][]scheduling.Requirements, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func validateNodeSelector(ctx context.Context, p *corev1.Pod) (errs error) {
	_ = "STUB: not implemented"
	return nil
}

func validateAffinity(ctx context.Context, p *corev1.Pod) (errs error) {
	_ = "STUB: not implemented"
	return nil
}

func validateNodeSelectorTerm(ctx context.Context, term corev1.NodeSelectorTerm) (errs error) {
	_ = "STUB: not implemented"
	return nil
}
