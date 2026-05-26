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
	"time"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	"sigs.k8s.io/karpenter/pkg/controllers/state"
)

const (
	minReconciles = 10
	maxReconciles = 3000
)

var stateRetryPeriod = 1 * time.Minute

// PodController reconciles pods for the purpose of maintaining state regarding pods that is expensive to compute.
type PodController struct {
	kubeClient client.Client
	cluster    *state.Cluster
}

func NewPodController(kubeClient client.Client, cluster *state.Cluster) *PodController {
	_ = "STUB: not implemented"
	return nil
}

func (c *PodController) Name() string { _ = "STUB: not implemented"; return "" }

func (c *PodController) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// notify cluster state of the node deletion

// We requeue here since the NotFound error is from finding the node for the binding

func (c *PodController) Register(ctx context.Context, m manager.Manager) error {
	_ = "STUB: not implemented"
	return nil
}
