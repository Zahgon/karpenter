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

package counter

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	v1 "sigs.k8s.io/karpenter/pkg/apis/v1"
	"sigs.k8s.io/karpenter/pkg/cloudprovider"
	"sigs.k8s.io/karpenter/pkg/controllers/state"
	"sigs.k8s.io/karpenter/pkg/utils/resources"
)

// Controller for the resource
type Controller struct {
	kubeClient    client.Client
	cloudProvider cloudprovider.CloudProvider
	cluster       *state.Cluster
}

var BaseResources = corev1.ResourceList{
	corev1.ResourceCPU:              resource.MustParse("0"),
	corev1.ResourceMemory:           resource.MustParse("0"),
	corev1.ResourcePods:             resource.MustParse("0"),
	corev1.ResourceEphemeralStorage: resource.MustParse("0"),
	resources.Node:                  resource.MustParse("0"),
}

// NewController is a constructor
func NewController(kubeClient client.Client, cloudProvider cloudprovider.CloudProvider, cluster *state.Cluster) *Controller {
	_ = "STUB: not implemented"
	return nil
}

// Reconcile a control loop for the resource
func (c *Controller) Name() string { _ = "STUB: not implemented"; return "" }

func (c *Controller) Reconcile(ctx context.Context, nodePool *v1.NodePool) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// We need to ensure that our internal cluster state mechanism is synced before we proceed
// Otherwise, we have the potential to patch over the status with a lower value for the nodepool resource
// counts on startup

// Determine resource usage and update nodepool.status.resources and nodepool.status.Nodes

func (c *Controller) Register(ctx context.Context, m manager.Manager) error {
	_ = "STUB: not implemented"
	return nil
}
