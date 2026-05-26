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

package termination

import (
	"context"
	"time"

	corev1 "k8s.io/api/core/v1"
	storagev1 "k8s.io/api/storage/v1"
	"k8s.io/utils/clock"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	v1 "sigs.k8s.io/karpenter/pkg/apis/v1"
	"sigs.k8s.io/karpenter/pkg/cloudprovider"
	"sigs.k8s.io/karpenter/pkg/controllers/node/termination/terminator"
	"sigs.k8s.io/karpenter/pkg/events"
)

const (
	minReconciles = 100
	maxReconciles = 5000
	MinDrainTime  = 5 * time.Second
)

// Controller for the resource
type Controller struct {
	clock         clock.Clock
	kubeClient    client.Client
	cloudProvider cloudprovider.CloudProvider
	terminator    *terminator.Terminator
	recorder      events.Recorder
}

// NewController constructs a controller instance
func NewController(clk clock.Clock, kubeClient client.Client, cloudProvider cloudprovider.CloudProvider, terminator *terminator.Terminator, recorder events.Recorder) *Controller {
	_ = "STUB: not implemented"
	return nil
}

func (c *Controller) Reconcile(ctx context.Context, n *corev1.Node) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

//nolint:gocyclo
func (c *Controller) finalize(ctx context.Context, node *corev1.Node) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// We're not guaranteed to find a NodeClaim for the node (e.g. if we failed to persist the provider ID to the nodeclaim
// at launch). If there are duplicate NodeClaims, we will treat it as though there is no NodeClaim since there is no
// longer a single source of truth.

// If the underlying instance no longer exists, we want to delete to avoid trying to gracefully draining the
// associated node. We do a check on the Ready condition of the node since, even though the CloudProvider says the
// instance is not around, we know that the kubelet process is still running if the Node Ready condition is true.
// Similar logic to: https://github.com/kubernetes/kubernetes/blob/3a75a8c8d9e6a1ebd98d8572132e675d4980f184/staging/src/k8s.io/cloud-provider/controllers/nodelifecycle/node_lifecycle_controller.go#L144

// If we don't have a NodeClaim, then there's nothing for us to patch here

// We use client.MergeFromWithOptimisticLock because patching a list with a JSON merge patch
// can cause races due to the fact that it fully replaces the list on a change
// Here, we are updating the status condition list

// We only increment the drained metric after we have ensured that we have patched the status condition onto the NodeClaim

// We'll only increment this metric if there is a NodeClaim present for the node, but this prevents us from double
// counting over multiple reconciles.

// We sleep here after a patch operation since we want to ensure that we are able to read our own writes
// so that we avoid duplicating metrics and log lines due to quick re-queues from our node watcher
// USE CAUTION when determining whether to increase this timeout or remove this line

type terminationFunc func(context.Context, *v1.NodeClaim, *corev1.Node, *time.Time) (reconcile.Result, error)

// awaitDrain initiates the drain of the node and will continue to requeue until the node has been drained and the minimum drain time has passed.
// If the nodeClaim has a terminationGracePeriod set, pods will be deleted to ensure this function does not requeue past the
// nodeTerminationTime.
func (c *Controller) awaitDrain(
	ctx context.Context,
	nodeClaim *v1.NodeClaim,
	node *corev1.Node,
	nodeTerminationTime *time.Time,
) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// If the nodeclaim exists, check if minDrainTime has elapsed. If it hasn't we should requeue.
// This check helps to ensure that we drain pods scheduled to the Node immediately after we taint it, which
// can occur when the scheduler has not seen the taint yet.

// awaitVolumeDetachment will continue to requeue until all volume attachments associated with the node have been
// deleted. The deletion is performed by the upstream attach-detach controller, Karpenter just needs to await deletion.
// This will be skipped once the nodeClaim's terminationGracePeriod has elapsed at nodeTerminationTime.
//
//nolint:gocyclo
func (c *Controller) awaitVolumeDetachment(
	ctx context.Context,
	nodeClaim *v1.NodeClaim,
	node *corev1.Node,
	nodeTerminationTime *time.Time,
) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	// In order for Pods associated with PersistentVolumes to smoothly migrate from the terminating Node, we wait
	// for VolumeAttachments of drain-able Pods to be cleaned up before terminating Node and removing its finalizer.
	// However, if TerminationGracePeriod is configured for Node, and we are past that period, we will skip waiting.
	return *new(reconcile.Result), nil
}

// There are no remaining volume attachments blocking instance termination. If we've already updated the status
// condition, fall through. Otherwise, update the status condition and requeue.

// There are volume attachments blocking instance termination remaining. We should set the status condition to
// unknown (if not already) and requeue. This case should never fall through, to continue to instance termination
// one of two conditions must be met: all blocking volume attachment objects must be deleted or the nodeclaim's TGP
// must have expired.

// There are volume attachments blocking instance termination remaining, but the nodeclaim's TGP has expired. In this
// case we should set the status condition to false (requeing if it wasn't already) and then fall through to instance
// termination.

// awaitInstanceTermination will initiate instance termination and continue to requeue until the cloudprovider indicates
// the instance is no longer found. Once gone, the node's finalizer will be removed, unblocking the NodeClaim lifecycle
// controller.
func (c *Controller) awaitInstanceTermination(
	ctx context.Context,
	nodeClaim *v1.NodeClaim,
	_ *corev1.Node,
	_ *time.Time,
) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

func (c *Controller) hasTerminationGracePeriodElapsed(nodeTerminationTime *time.Time) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *Controller) pendingVolumeAttachments(ctx context.Context, node *corev1.Node) ([]*storagev1.VolumeAttachment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Filter out VolumeAttachments associated with not drain-able Pods

// filterVolumeAttachments filters out storagev1.VolumeAttachments that should not block the termination
// of the passed corev1.Node
func filterVolumeAttachments(ctx context.Context, kubeClient client.Client, node *corev1.Node, volumeAttachments []*storagev1.VolumeAttachment, clk clock.Clock) ([]*storagev1.VolumeAttachment, error) {
	_ = "STUB: not implemented"
	// No need to filter empty VolumeAttachments list
	return nil, nil
}

// Create list of non-drain-able Pods associated with Node

// Filter out VolumeAttachments associated with non-drain-able Pods
// Match on Pod -> PersistentVolumeClaim -> PersistentVolume Name <- VolumeAttachment

func (c *Controller) removeFinalizer(ctx context.Context, n *corev1.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// We use client.StrategicMergeFrom here since the node object supports it and
// a strategic merge patch represents the finalizer list as a keyed "set" so removing
// an item from the list doesn't replace the full list
// https://github.com/kubernetes/kubernetes/issues/111643#issuecomment-2016489732

// We use stored.DeletionTimestamp since the api-server may give back a node after the patch without a deletionTimestamp

func (c *Controller) nodeTerminationTime(node *corev1.Node, nodeClaim *v1.NodeClaim) (*time.Time, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Controller) Name() string { _ = "STUB: not implemented"; return "" }

func (c *Controller) Register(ctx context.Context, m manager.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

// qps scales linearly at 10% of concurrentReconciles, bucket size is 10 * qps
