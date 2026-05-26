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

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	"sigs.k8s.io/karpenter/pkg/controllers/state"
)

const (
	minReconciles = 10
	maxReconciles = 1000
)

// PodController for the resource
type PodController struct {
	kubeClient  client.Client
	provisioner *Provisioner
	cluster     *state.Cluster
}

// NewPodController constructs a controller instance
func NewPodController(kubeClient client.Client, provisioner *Provisioner, cluster *state.Cluster) *PodController {
	_ = "STUB: not implemented"
	return nil
}

// Reconcile the resource
func (c *PodController) Name() string { _ = "STUB: not implemented"; return "" }

func (c *PodController) Reconcile(ctx context.Context, p *corev1.Pod) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

//nolint:ineffassign,staticcheck

// ACK the pending pod when first observed so that total time spent pending due to Karpenter is tracked.

// Continue to requeue until the pod is no longer provisionable. Pods may
// not be scheduled as expected if new pods are created while nodes are
// coming online. Even if a provisioning loop is successful, the pod may
// require another provisioning loop to become schedulable.

func (c *PodController) Register(ctx context.Context, m manager.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

// NodeController for the resource
type NodeController struct {
	kubeClient  client.Client
	provisioner *Provisioner
}

// NewNodeController constructs a controller instance
func NewNodeController(kubeClient client.Client, provisioner *Provisioner) *NodeController {
	_ = "STUB: not implemented"
	return nil
}

// Reconcile the resource
func (c *NodeController) Name() string { _ = "STUB: not implemented"; return "" }

func (c *NodeController) Reconcile(ctx context.Context, n *corev1.Node) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	//nolint:ineffassign
	return *new(reconcile.Result), nil
}

//nolint:ineffassign,staticcheck

// If the disruption taint doesn't exist and the deletion timestamp isn't set, it's not being disrupted.
// We don't check the deletion timestamp here, as we expect the termination controller to eventually set
// the taint when it picks up the node from being deleted.

// Continue to requeue until the node is no longer provisionable. Pods may
// not be scheduled as expected if new pods are created while nodes are
// coming online. Even if a provisioning loop is successful, the pod may
// require another provisioning loop to become schedulable.

func (c *NodeController) Register(ctx context.Context, m manager.Manager) error {
	_ = "STUB: not implemented"
	return nil
}
