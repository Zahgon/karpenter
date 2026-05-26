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

	"github.com/awslabs/operatorpkg/reconciler"
	"k8s.io/apimachinery/pkg/types"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	"sigs.k8s.io/karpenter/pkg/cloudprovider"
	"sigs.k8s.io/karpenter/pkg/state/cost"
)

// This whole pricing controller only exists because CP's surface InstanceType information
// via a poll-method. Long term, this controller should likely be removed in favor of refactoring the
// cloudprovider interface to a push model for instance type information:
// https://github.com/kubernetes-sigs/karpenter/issues/2605
type PricingController struct {
	client        client.Client
	cloudProvider cloudprovider.CloudProvider
	clusterCost   *cost.ClusterCost
	npOfMap       map[types.NamespacedName]map[cost.OfferingKey]float64
}

func NewPricingController(client client.Client, cloudProvider cloudprovider.CloudProvider, clusterCost *cost.ClusterCost) *PricingController {
	_ = "STUB: not implemented"
	return nil
}

func (c *PricingController) Reconcile(ctx context.Context) (reconciler.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconciler.Result), nil
}

func equal(oldOfs map[cost.OfferingKey]float64, newOfs map[cost.OfferingKey]float64) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *PricingController) Name() string { _ = "STUB: not implemented"; return "" }

func (c *PricingController) Register(_ context.Context, m manager.Manager) error {
	_ = "STUB: not implemented"
	return nil
}
