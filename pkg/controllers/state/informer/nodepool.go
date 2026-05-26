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

	"sigs.k8s.io/karpenter/pkg/cloudprovider"
	"sigs.k8s.io/karpenter/pkg/controllers/state"
	"sigs.k8s.io/karpenter/pkg/state/cost"
)

// NodePoolController reconciles NodePools to re-trigger consolidation on change.
type NodePoolController struct {
	kubeClient    client.Client
	cloudProvider cloudprovider.CloudProvider
	cluster       *state.Cluster
	clusterCost   *cost.ClusterCost
}

func NewNodePoolController(kubeClient client.Client, cloudProvider cloudprovider.CloudProvider, cluster *state.Cluster, clusterCost *cost.ClusterCost) *NodePoolController {
	_ = "STUB: not implemented"
	return nil
}

func (c *NodePoolController) Name() string { _ = "STUB: not implemented"; return "" }

func (c *NodePoolController) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

//nolint:ineffassign,staticcheck

// notify cluster state of the nodepool deletion

// Something changed in the NodePool so we should re-consider consolidation

func (c *NodePoolController) Register(ctx context.Context, m manager.Manager) error {
	_ = "STUB: not implemented"
	return nil
}
