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

package podevents

import (
	"context"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/utils/clock"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	"sigs.k8s.io/karpenter/pkg/cloudprovider"
)

// dedupeTimeout is 10 seconds to reduce the number of writes to the APIServer, since pod scheduling and deletion events are very frequent.
// The smaller this value is, the more writes Karpenter will make in a busy cluster. This timeout is intentionally smaller than the consolidation
// 15 second validation period, so that we can ensure that we invalidate consolidation commands that are decided while we're de-duping pod events.
const dedupeTimeout = 10 * time.Second

// Podevents is a nodeclaim controller that deletes adds the lastPodEvent status onto the nodeclaim
type Controller struct {
	clock         clock.Clock
	kubeClient    client.Client
	cloudProvider cloudprovider.CloudProvider
}

// NewController constructs a nodeclaim disruption controller
func NewController(clk clock.Clock, kubeClient client.Client, cloudProvider cloudprovider.CloudProvider) *Controller {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gocyclo
func (c *Controller) Reconcile(ctx context.Context, pod *corev1.Pod) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	// If the pod doesn't have a node name, we don't know which node this pod refers to.
	// or if this is a daemonset
	return *new(reconcile.Result), nil
}

// If there's no associated node claim, it's not a karpenter owned node.

// if the nodeclaim doesn't exist, or has duplicates, ignore.

// If we've set the lastPodEvent before and it hasn't been before the timeout, don't do anything

// otherwise, set the pod event time to now

func (c *Controller) Name() string { _ = "STUB: not implemented"; return "" }

func (c *Controller) Register(ctx context.Context, m manager.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

// If a pod is bound to a node or goes terminal

// if this is a newly bound pod

// if this is a newly terminal pod

// if this is a newly terminating pod

// return true if it was bound to a node, went terminal, or went terminating
