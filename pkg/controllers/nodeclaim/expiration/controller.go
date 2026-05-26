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

package expiration

import (
	"context"

	"k8s.io/utils/clock"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	v1 "sigs.k8s.io/karpenter/pkg/apis/v1"
	"sigs.k8s.io/karpenter/pkg/cloudprovider"
)

// Expiration is a nodeclaim controller that deletes expired nodeclaims based on expireAfter
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

func (c *Controller) Reconcile(ctx context.Context, nodeClaim *v1.NodeClaim) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// From here there are three scenarios to handle:
// 1. If ExpireAfter is not configured, exit expiration loop

// 2. If the NodeClaim isn't expired leave the reconcile loop.

// Use t.Sub(clock.Now()) instead of time.Until() to ensure we're using the injected clock.

// 3. Otherwise, if the NodeClaim is expired we can forcefully expire the nodeclaim (by deleting it)

// 4. The deletion timestamp has successfully been set for the NodeClaim, update relevant metrics.

// We sleep here after the delete operation since we want to ensure that we are able to read our own writes so that
// we avoid duplicating metrics and log lines due to quick re-queues.
// USE CAUTION when determining whether to increase this timeout or remove this line

func (c *Controller) Name() string { _ = "STUB: not implemented"; return "" }

func (c *Controller) Register(_ context.Context, m manager.Manager) error {
	_ = "STUB: not implemented"
	return nil
}
