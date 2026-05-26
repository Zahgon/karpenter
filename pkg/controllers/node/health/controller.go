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

package health

import (
	"context"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"

	"k8s.io/utils/clock"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	v1 "sigs.k8s.io/karpenter/pkg/apis/v1"
	"sigs.k8s.io/karpenter/pkg/cloudprovider"
	"sigs.k8s.io/karpenter/pkg/events"
)

var allowedUnhealthyPercent = intstr.FromString("20%")

// Controller for the resource
type Controller struct {
	clock         clock.Clock
	recorder      events.Recorder
	kubeClient    client.Client
	cloudProvider cloudprovider.CloudProvider
}

// NewController constructs a controller instance
func NewController(kubeClient client.Client, cloudProvider cloudprovider.CloudProvider, clock clock.Clock, recorder events.Recorder) *Controller {
	_ = "STUB: not implemented"
	return nil
}

func (c *Controller) Name() string { _ = "STUB: not implemented"; return "" }

func (c *Controller) Register(ctx context.Context, m manager.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

// Return true if any of these conditions are met:
// 1. Condition type no longer exists in new node
// 2. Transition time has changed
// 3. Status has changed

func (c *Controller) Reconcile(ctx context.Context, node *corev1.Node) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// Validate that the node is owned by us

// If the Node is unhealthy, but has not reached its full toleration disruption
// requeue at the termination time of the unhealthy node

// If a nodeclaim does have a nodepool label, validate the nodeclaims inside the nodepool are healthy (i.e bellow the allowed threshold)
// In the case of standalone nodeclaim, validate the nodes inside the cluster are healthy before proceeding
// to repair the nodes

// For unhealthy past the tolerationDisruption window we can forcefully terminate the node

// deleteNodeClaim removes the NodeClaim from the api-server
func (c *Controller) deleteNodeClaim(ctx context.Context, nodeClaim *v1.NodeClaim, node *corev1.Node, unhealthyNodeCondition *corev1.NodeCondition) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// The deletion timestamp has successfully been set for the Node, update relevant metrics.

// Find a node with a condition that matches one of the unhealthy conditions defined by the cloud provider
// If there are multiple unhealthy status condition we will requeue based on the condition closest to its terminationDuration
func (c *Controller) findUnhealthyConditions(node *corev1.Node) (nc *corev1.NodeCondition, cpTerminationDuration time.Duration) {
	_ = "STUB: not implemented"
	return nil, *new(time.Duration)
}

// check the status and the type on the condition

// Determine requeue time

func (c *Controller) annotateTerminationGracePeriod(ctx context.Context, nodeClaim *v1.NodeClaim) error {
	_ = "STUB: not implemented"
	return nil
}

// isNodePoolHealthy checks if the number of unhealthy nodes managed by the given NodePool exceeds the health threshold.
// defined by the cloud provider
// Up to 20% of Nodes may be unhealthy before the NodePool becomes unhealthy (or the nearest whole number, rounding up).
// For example, given a NodePool with three nodes, one may be unhealthy without rendering the NodePool unhealthy, even though that's 33% of the total nodes.
// This is analogous to how minAvailable and maxUnavailable work for PodDisruptionBudgets: https://kubernetes.io/docs/tasks/run-application/configure-pdb/#rounding-logic-when-specifying-percentages.
func (c *Controller) isNodePoolHealthy(ctx context.Context, nodePoolName string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (c *Controller) isClusterHealthy(ctx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (c *Controller) areNodesHealthy(ctx context.Context, opts ...client.ListOption) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (c *Controller) publishNodePoolHealthEvent(ctx context.Context, node *corev1.Node, nodeClaim *v1.NodeClaim, npName string) error {
	_ = "STUB: not implemented"
	return nil
}
