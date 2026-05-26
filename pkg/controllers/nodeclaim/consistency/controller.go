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

package consistency

import (
	"context"
	"time"

	"github.com/patrickmn/go-cache"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/utils/clock"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	v1 "sigs.k8s.io/karpenter/pkg/apis/v1"
	"sigs.k8s.io/karpenter/pkg/cloudprovider"
	"sigs.k8s.io/karpenter/pkg/events"
)

type Controller struct {
	clock         clock.Clock
	kubeClient    client.Client
	cloudProvider cloudprovider.CloudProvider
	checks        []Check
	recorder      events.Recorder
	lastScanned   *cache.Cache
}

type Issue string

type Check interface {
	// Check performs the consistency check, this should return a list of slice discovered, or an empty
	// slice if no issues were found
	Check(context.Context, *corev1.Node, *v1.NodeClaim) ([]Issue, error)
}

// scanPeriod is how often we inspect and report issues that are found.
const scanPeriod = 10 * time.Minute

func NewController(clk clock.Clock, kubeClient client.Client, cloudProvider cloudprovider.CloudProvider, recorder events.Recorder) *Controller {
	_ = "STUB: not implemented"
	return nil
}

func (c *Controller) Name() string { _ = "STUB: not implemented"; return "" }

func (c *Controller) Reconcile(ctx context.Context, nodeClaim *v1.NodeClaim) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// If we get an event before we should check for consistency checks, we ignore and wait

// the above should always succeed

// We assume the invariant that there is a single node for a single nodeClaim. If this invariant is violated,
// then we assume this is bubbled up through the nodeClaim lifecycle controller and don't perform consistency checks

// We use client.MergeFromWithOptimisticLock because patching a list with a JSON merge patch
// can cause races due to the fact that it fully replaces the list on a change
// Here, we are updating the status condition list

func (c *Controller) checkConsistency(ctx context.Context, nodeClaim *v1.NodeClaim, node *corev1.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// If status condition for consistent state is not true and no issues are found, set the status condition to true

// If there are issues then set the status condition for consistent state as false

func (c *Controller) Register(ctx context.Context, m manager.Manager) error {
	_ = "STUB: not implemented"
	return nil
}
