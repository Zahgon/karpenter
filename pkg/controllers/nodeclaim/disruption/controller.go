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

package disruption

import (
	"context"

	"k8s.io/utils/clock"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	v1 "sigs.k8s.io/karpenter/pkg/apis/v1"
	"sigs.k8s.io/karpenter/pkg/cloudprovider"
)

type nodeClaimReconciler interface {
	Reconcile(context.Context, *v1.NodePool, *v1.NodeClaim) (reconcile.Result, error)
}

// Controller is a disruption controller that adds StatusConditions to nodeclaims when they meet certain disruption conditions
// e.g. When the NodeClaim has become empty, then it is marked as "Empty" in the StatusConditions
type Controller struct {
	kubeClient    client.Client
	cloudProvider cloudprovider.CloudProvider

	drift         *Drift
	consolidation *Consolidation
}

// NewController constructs a nodeclaim disruption controller. Note that every sub-controller has a dependency on its nodepool.
// Disruption mechanisms that don't depend on the nodepool (like expiration), should live elsewhere.
func NewController(clk clock.Clock, kubeClient client.Client, cloudProvider cloudprovider.CloudProvider) *Controller {
	_ = "STUB: not implemented"
	return nil
}

// Reconcile executes a control loop for the resource
func (c *Controller) Name() string { _ = "STUB: not implemented"; return "" }

func (c *Controller) Reconcile(ctx context.Context, nodeClaim *v1.NodeClaim) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// We use client.MergeFromWithOptimisticLock because patching a list with a JSON merge patch
// can cause races due to the fact that it fully replaces the list on a change
// Here, we are updating the status condition list

func (c *Controller) Register(ctx context.Context, m manager.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Controller) Reset() { _ = "STUB: not implemented"; return }

func (c *Controller) runReconcilers(
	ctx context.Context,
	np *v1.NodePool,
	nc *v1.NodeClaim,
) ([]reconcile.Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NodeClaims belonging to static NodePools are never eligible for consolidation, so we shouldn't mark them as consolidatable
