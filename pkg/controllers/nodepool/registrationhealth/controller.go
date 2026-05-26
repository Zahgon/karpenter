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

package registrationhealth

import (
	"context"

	"sigs.k8s.io/karpenter/pkg/state/nodepoolhealth"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	v1 "sigs.k8s.io/karpenter/pkg/apis/v1"
	"sigs.k8s.io/karpenter/pkg/cloudprovider"
)

// Controller for the resource
type Controller struct {
	kubeClient    client.Client
	cloudProvider cloudprovider.CloudProvider
	npState       *nodepoolhealth.State
}

// NewController will create a controller to reset NodePool's registration health when there is an update to NodePool/NodeClass spec
func NewController(kubeClient client.Client, cloudProvider cloudprovider.CloudProvider, npState *nodepoolhealth.State) *Controller {
	_ = "STUB: not implemented"
	return nil
}

func (c *Controller) Name() string { _ = "STUB: not implemented"; return "" }

//nolint:gocyclo
func (c *Controller) Reconcile(ctx context.Context, nodePool *v1.NodePool) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// Ignore NodePools which aren't using a supported NodeClass

// If Karpenter restarts i.e. if the buffer for the nodePool is empty and the NodeRegistrationHealthy status condition
// is set to either true/false then we pre-hydrate the buffer with the existing state of the status condition

// If NodeClass/NodePool have been updated then NodeRegistrationHealthy = Unknown and reset the buffer

// We use client.MergeFromWithOptimisticLock because patching a list with a JSON merge patch
// can cause races due to the fact that it fully replaces the list on a change
// Here, we are updating the status condition list

func (c *Controller) Register(ctx context.Context, m manager.Manager) error {
	_ = "STUB: not implemented"
	return nil
}
