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

package informer

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	"sigs.k8s.io/karpenter/pkg/controllers/state"
)

// NodeController reconciles nodes for the purpose of maintaining state regarding nodes that is expensive to compute.
type NodeController struct {
	kubeClient client.Client
	cluster    *state.Cluster
}

// NewNodeController constructs a controller instance
func NewNodeController(kubeClient client.Client, cluster *state.Cluster) *NodeController {
	_ = "STUB: not implemented"
	return nil
}

func (c *NodeController) Name() string { _ = "STUB: not implemented"; return "" }

func (c *NodeController) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// notify cluster state of the node deletion

// ensure it's aware of any nodes we discover, this is a no-op if the node is already known to our cluster state

func (c *NodeController) Register(ctx context.Context, m manager.Manager) error {
	_ = "STUB: not implemented"
	return nil
}
