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
)

type Initialization struct {
	kubeClient client.Client
}

// Reconcile checks for initialization based on if:
// a) its current status is set to Ready
// b) all the startup taints have been removed from the node
// c) all extended resources have been registered
// This method handles both nil nodepools and nodes without extended resources gracefully.
func (i *Initialization) Reconcile(ctx context.Context, nodeClaim *v1.NodeClaim) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// Ensure that we always set the status condition to the latest generation

//nolint:nilerr

// KnownEphemeralTaintsRemoved validates whether all the ephemeral taints are removed
func KnownEphemeralTaintsRemoved(node *corev1.Node) (*corev1.Taint, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// StartupTaintsRemoved returns true if there are no startup taints registered for the nodepool, or if all startup
// taints have been removed from the node
func StartupTaintsRemoved(node *corev1.Node, nodeClaim *v1.NodeClaim) (*corev1.Taint, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// if the node still has a startup taint applied, it's not ready

// RequestedResourcesRegistered returns true if there are no extended resources on the node, or they have all been
// registered by device plugins
func RequestedResourcesRegistered(node *corev1.Node, nodeClaim *v1.NodeClaim) (corev1.ResourceName, bool) {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceName), false
}

// kubelet will zero out both the capacity and allocatable for an extended resource on startup, so if our
// annotation says the resource should be there, but it's zero'd in both then the device plugin hasn't
// registered it yet.
// We wait on allocatable since this is the value that is used in scheduling

func formatTaint(taint *corev1.Taint) string { _ = "STUB: not implemented"; return "" }
