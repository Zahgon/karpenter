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

package garbagecollection

import (
	"context"

	"github.com/awslabs/operatorpkg/reconciler"

	"k8s.io/utils/clock"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	"sigs.k8s.io/karpenter/pkg/cloudprovider"
)

type Controller struct {
	clock         clock.Clock
	kubeClient    client.Client
	cloudProvider cloudprovider.CloudProvider
}

func NewController(c clock.Clock, kubeClient client.Client, cloudProvider cloudprovider.CloudProvider) *Controller {
	_ = "STUB: not implemented"
	return nil
}

func (c *Controller) Name() string { _ = "STUB: not implemented"; return "" }

func (c *Controller) Reconcile(ctx context.Context) (reconciler.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconciler.Result), nil
}

// Only consider NodeClaims that are Registered since we don't want to fully rely on the CloudProvider
// API to trigger deletion of the Node. Instead, we'll wait for our registration timeout to trigger

// Ignore these errors since a registered NodeClaim should only have a NotFound node when
// the Node was deleted out from under us and a Duplicate Node is an invalid state

// We do a check on the Ready condition of the node since, even though the CloudProvider says the instance
// is not around, we know that the kubelet process is still running if the Node Ready condition is true
// Similar logic to: https://github.com/kubernetes/kubernetes/blob/3a75a8c8d9e6a1ebd98d8572132e675d4980f184/staging/src/k8s.io/cloud-provider/controllers/nodelifecycle/node_lifecycle_controller.go#L144

func (c *Controller) Register(_ context.Context, m manager.Manager) error {
	_ = "STUB: not implemented"
	return nil
}
