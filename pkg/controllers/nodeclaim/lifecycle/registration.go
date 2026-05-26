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

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	v1 "sigs.k8s.io/karpenter/pkg/apis/v1"
	"sigs.k8s.io/karpenter/pkg/cloudprovider"
	"sigs.k8s.io/karpenter/pkg/events"
	"sigs.k8s.io/karpenter/pkg/state/nodepoolhealth"
)

type Registration struct {
	kubeClient        client.Client
	recorder          events.Recorder
	npState           *nodepoolhealth.State
	registrationHooks []cloudprovider.NodeLifecycleHook
}

//nolint:gocyclo
func (r *Registration) Reconcile(ctx context.Context, nodeClaim *v1.NodeClaim) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// Ensure that we always set the status condition to the latest generation

// if the sync hasn't happened yet and the race protecting startup taint isn't present then log it as missing and proceed
// if the sync has happened then the startup taint has been removed if it was present

// Sync labels, annotations, taints, finalizer, and owner references onto the node.

// Check all registration hooks before completing registration.
// If any hook is not ready, registration is deferred and the unregistered taint remains.

// Re-sync the node after hooks complete since hooks may have mutated the nodeClaim.

// We use client.MergeFromWithOptimisticLock because patching a list with a JSON merge patch
// can cause races due to the fact that it fully replaces the list on a change

// If hooks were not ready, return.

// checkRegistrationHooks evaluates all registration hooks in parallel. If any hook returns an error,
// it is returned. If any hook signals it is not ready (non-empty result), the status condition is
// updated to list all pending hooks and the shortest requeue interval is returned.
//
//nolint:gocyclo
func (r *Registration) checkRegistrationHooks(ctx context.Context, nodeClaim *v1.NodeClaim) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// Collect pending and errored hook names and compute the shortest requeue interval

//nolint:staticcheck

// updateNodePoolRegistrationHealth adds a positive value to the nodepool buffer that stores node
// registration results and sets NodeRegistrationHealthy=True on the NodePool if IsHealthy() > 0
func (r *Registration) updateNodePoolRegistrationHealth(ctx context.Context, nodeClaim *v1.NodeClaim) error {
	_ = "STUB: not implemented"
	return nil
}

// We use client.MergeFromWithOptimisticLock because patching a list with a JSON merge patch
// can cause races due to the fact that it fully replaces the list on a change
// Here, we are updating the status condition list

// syncNode mutates the node in memory to sync labels, annotations, taints, finalizer, and owner references
// from the NodeClaim.
func (r *Registration) syncNode(nodeClaim *v1.NodeClaim, node *corev1.Node) {
	_ = "STUB: not implemented"
	return
}

// We do not sync the taints if this label is present. We instead assume that the karpenter provider
// is managing taints. We still manage/remove the unregistered taint to signal the end of syncing.

// Sync all taints inside NodeClaim into the Node taints
