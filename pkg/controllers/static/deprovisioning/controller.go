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

package static

import (
	"context"

	"k8s.io/utils/clock"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	v1 "sigs.k8s.io/karpenter/pkg/apis/v1"
	"sigs.k8s.io/karpenter/pkg/cloudprovider"
	"sigs.k8s.io/karpenter/pkg/controllers/state"
	"sigs.k8s.io/karpenter/pkg/events"
)

const (
	TerminationReason = "overprovisioned"
)

type Controller struct {
	kubeClient    client.Client
	cloudProvider cloudprovider.CloudProvider
	cluster       *state.Cluster
	clock         clock.Clock
	recorder      events.Recorder
}

func NewController(kubeClient client.Client, cluster *state.Cluster, cloudProvider cloudprovider.CloudProvider, clock clock.Clock, recorder events.Recorder) *Controller {
	_ = "STUB: not implemented"
	return nil
}

// Reconcile the resource
// Requeue after computing Static NodePool to ensure we don't miss any events
func (c *Controller) Name() string { _ = "STUB: not implemented"; return "" }

func (c *Controller) Reconcile(ctx context.Context, np *v1.NodePool) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// We dont have to wait for cluster sync as we cannot really have internal state representing more NodeClaims than actual
// During controller crashes we gradually populate our cluster/NodePoolState, as and when we populate we delete NC if we are over-provisioned

// To avoid race conditions between deprovisioning and the disruption controller,
// we only include running NodeClaims when counting for deprovisioning purposes.
// Including both active NodeClaims and those pending disruption could cause us
// to temporarily exceed the desired replica count while replacements are being created.

// Only handle scale down - scale up is handled by provisioning controller

// Get deprovisioning candidates

// Terminate selected NodeClaims

// Mark the NodeClaim as Deleting in StateNodePool

func (c *Controller) Register(_ context.Context, m manager.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

// Reoncile on NodePool Create and Update (when replicas change)

// We care about Static NodeClaims creating as we might have over provisioned and need to deprovision

func HasNodePoolReplicaCountChanged(oldNP, newNP *v1.NodePool) bool {
	_ = "STUB: not implemented"
	return false
}

// Returns NodeClaims suitable for deprovisioning, prioritizing:
// 1. Unresolved NodeClaims (no ProviderID yet - haven't launched)
// 2. Empty nodes (nodes with no pods or only DaemonSet pods without do-not-disrupt annotation)
// 3. If more nodes needed, nodes with lowest disruption cost (nodes with pods that have do-not-disrupt will have highest cost)
func (c *Controller) getDeprovisioningCandidates(ctx context.Context, np *v1.NodePool, count int) []*v1.NodeClaim {
	_ = "STUB: not implemented"
	return nil
}

// Unresolved NodeClaims (haven't launched yet or that failed Create call)

// Get all StateNodes for this NodePool

// Resolved nodes (empty first, then by disruption cost)

// unResolvedDeprovisioningCandidates returns unresolved NodeClaims (those without ProviderID) up to the specified count
func (c *Controller) unresolvedDeprovisioningCandidates(ctx context.Context, nodePoolName string, count int) []*v1.NodeClaim {
	_ = "STUB: not implemented"
	return nil
}

// resolvedDeprovisioningCandidates returns resolved NodeClaims (those with ProviderID) up to the specified count,
// prioritizing empty nodes first, then nodes with lowest disruption cost
func (c *Controller) resolvedDeprovisioningCandidates(ctx context.Context, nodes []*state.StateNode, np *v1.NodePool, count int) []*v1.NodeClaim {
	_ = "STUB: not implemented"
	return nil
}

// Priority 1: Empty nodes

// Get non-empty nodes with their costs

// If one node has do-not-disrupt pods and the other doesn't, the one without should come first

// If neither has do-not-disrupt pods, compare their costs

// Take the remaining needed nodes with lowest cost
