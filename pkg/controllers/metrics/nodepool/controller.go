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

package nodepool

import (
	"context"

	opmetrics "github.com/awslabs/operatorpkg/metrics"
	"github.com/prometheus/client_golang/prometheus"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	crmetrics "sigs.k8s.io/controller-runtime/pkg/metrics"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	v1 "sigs.k8s.io/karpenter/pkg/apis/v1"
	"sigs.k8s.io/karpenter/pkg/cloudprovider"
	"sigs.k8s.io/karpenter/pkg/metrics"
	"sigs.k8s.io/karpenter/pkg/state/cost"
)

var (
	Limit = opmetrics.NewPrometheusGauge(
		crmetrics.Registry,
		prometheus.GaugeOpts{
			Namespace: metrics.Namespace,
			Subsystem: metrics.NodePoolSubsystem,
			Name:      "limit",
			Help:      "Limits specified on the nodepool that restrict the quantity of resources provisioned. Labeled by nodepool name and resource type.",
		},
		[]string{
			metrics.ResourceTypeLabel,
			metrics.NodePoolLabel,
		},
	)
	Usage = opmetrics.NewPrometheusGauge(
		crmetrics.Registry,
		prometheus.GaugeOpts{
			Namespace: metrics.Namespace,
			Subsystem: metrics.NodePoolSubsystem,
			Name:      "usage",
			Help:      "The amount of resources that have been provisioned for a nodepool. Labeled by nodepool name and resource type.",
		},
		[]string{
			metrics.ResourceTypeLabel,
			metrics.NodePoolLabel,
		},
	)
	ClusterCost = opmetrics.NewPrometheusGauge(
		crmetrics.Registry,
		prometheus.GaugeOpts{
			Namespace: metrics.Namespace,
			Subsystem: metrics.NodePoolSubsystem,
			Name:      "cost_total",
			Help:      "Total cost of the nodepool from Karpenter's perspective. Units are determined by the cloud provider. Not an authoritative source for billing. Includes modifications due to NodeOverlays",
		},
		[]string{metrics.NodePoolLabel},
	)
)

type Controller struct {
	kubeClient    client.Client
	cloudProvider cloudprovider.CloudProvider
	clusterCost   *cost.ClusterCost
	metricStore   *metrics.Store
}

// NewController constructs a controller instance
func NewController(kubeClient client.Client, cloudProvider cloudprovider.CloudProvider, clusterCost *cost.ClusterCost) *Controller {
	_ = "STUB: not implemented"
	return nil
}

// Reconcile executes a termination control loop for the resource
func (c *Controller) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// periodically update our metrics per nodepool even if nothing has changed

func (c *Controller) buildMetrics(nodePool *v1.NodePool) (res []*metrics.StoreMetric) {
	_ = "STUB: not implemented"
	return nil
}

func getLimits(nodePool *v1.NodePool) corev1.ResourceList {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList)
}

func makeLabels(nodePool *v1.NodePool, resourceTypeName string) prometheus.Labels {
	_ = "STUB: not implemented"
	return *new(prometheus.Labels)
}

func (c *Controller) Name() string { _ = "STUB: not implemented"; return "" }

func (c *Controller) Register(_ context.Context, m manager.Manager) error {
	_ = "STUB: not implemented"
	return nil
}
