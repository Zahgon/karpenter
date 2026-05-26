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

package lifecycle

import (
	"context"

	"k8s.io/utils/clock"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	"sigs.k8s.io/karpenter/pkg/state/nodepoolhealth"

	v1 "sigs.k8s.io/karpenter/pkg/apis/v1"
	"sigs.k8s.io/karpenter/pkg/cloudprovider"
	"sigs.k8s.io/karpenter/pkg/events"
)

const (
	minReconciles = 1000
	maxReconciles = 5000
)

// Controller is a NodeClaim Lifecycle controller that manages the lifecycle of the NodeClaim up until its termination
// The controller is responsible for ensuring that new Nodes get launched, that they have properly registered with
// the cluster as nodes and that they are properly initialized, ensuring that nodeclaims that do not have matching nodes
// after some liveness TTL are removed
type Controller struct {
	clock         clock.Clock
	kubeClient    client.Client
	cloudProvider cloudprovider.CloudProvider
	recorder      events.Recorder
	nodePoolState *nodepoolhealth.State

	launch         *Launch
	registration   *Registration
	initialization *Initialization
	liveness       *Liveness
}

func NewController(clk clock.Clock, kubeClient client.Client, cloudProvider cloudprovider.CloudProvider, recorder events.Recorder, nodePoolState *nodepoolhealth.State, registrationHooks []cloudprovider.NodeLifecycleHook) *Controller {
	_ = "STUB: not implemented"
	return nil
}

func (c *Controller) Register(ctx context.Context, m manager.Manager) error {
	_ = "STUB: not implemented"
	// higher concurrency limit since we want fast reaction to node syncing and launch
	return nil
}

// back off until last attempt occurs ~90 seconds before nodeclaim expiration

// qps scales linearly at 1% of concurrentReconciles, bucket size is 10 * qps

func (c *Controller) Name() string { _ = "STUB: not implemented"; return "" }

// nolint:gocyclo
func (c *Controller) Reconcile(ctx context.Context, nodeClaim *v1.NodeClaim) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// Add the finalizer immediately since we shouldn't launch if we don't yet have the finalizer.
// Otherwise, we could leak resources

// We use client.MergeFromWithOptimisticLock because patching a list with a JSON merge patch
// can cause races due to the fact that it fully replaces the list on a change
// Here, we are updating the finalizer list

// We sleep here after a patch operation since we want to ensure that we are able to read our own writes
// so that we avoid duplicating metrics and log lines due to quick re-queues from our node watcher
// USE CAUTION when determining whether to increase this timeout or remove this line

//nolint:gocyclo
func (c *Controller) finalize(ctx context.Context, nodeClaim *v1.NodeClaim) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// Only delete Nodes if the NodeClaim has been registered. Deleting Nodes without the termination finalizer
// may result in leaked leases due to a kubelet bug until k8s 1.29. The Node should be garbage collected after the
// instance is terminated by CCM.
// Upstream Kubelet Fix: https://github.com/kubernetes/kubernetes/pull/119661

// If we still get the Node, but it's already marked as terminating, we don't need to call Delete again

// We delete nodes to trigger the node finalization and deletion flow

// We wait until all the nodes associated with this nodeClaim have completed their deletion before triggering the finalization of the nodeClaim

// We can expect ProviderID to be empty when there is a failure while launching the nodeClaim

// We use client.MergeFromWithOptimisticLock because patching a list with a JSON merge patch
// can cause races due to the fact that it fully replaces the list on a change
// Here, we are updating the status condition list

// The NodeClaim may have been modified in the EnsureTerminated function

// We use client.MergeFromWithOptimisticLock because patching a list with a JSON merge patch
// can cause races due to the fact that it fully replaces the list on a change
// Here, we are updating the finalizer list
// https://github.com/kubernetes/kubernetes/issues/111643#issuecomment-2016489732

func (c *Controller) ensureTerminationGracePeriodTerminationTimeAnnotation(ctx context.Context, nodeClaim *v1.NodeClaim) error {
	_ = "STUB: not implemented"
	// if the expiration annotation is already set, we don't need to do anything
	return nil
}

// In Kubernetes, every object has a terminationGracePeriodSeconds, defaulted to and un-changeable from 0. There is an additional TerminationGracePeriodSeconds in the PodSpec which can be configured.
// We use the kubernetes object TerminationGracePeriod to infer that the DeletionTimestamp is always equal to the time the NodeClaim is deleted.
// This should not be confused with the NodeClaim.spec.terminationGracePeriod field introduced in Karpenter Custom Resources.

func (c *Controller) annotateTerminationGracePeriodTerminationTime(ctx context.Context, nodeClaim *v1.NodeClaim, terminationTime string) error {
	_ = "STUB: not implemented"
	return nil
}

// We use client.MergeFromWithOptimisticLock because patching a terminationGracePeriod annotation
// can cause races with the health controller, as that controller sets the current time as the terminationGracePeriod annotation
// Here, We want to resolve any conflict and not overwrite the terminationGracePeriod annotation
