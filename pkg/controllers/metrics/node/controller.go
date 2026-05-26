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

package node

import (
	"context"

	opmetrics "github.com/awslabs/operatorpkg/metrics"
	"github.com/awslabs/operatorpkg/reconciler"

	"github.com/prometheus/client_golang/prometheus"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	"sigs.k8s.io/karpenter/pkg/controllers/state"
	"sigs.k8s.io/karpenter/pkg/metrics"
)

const (
	nodeName  = "node_name"
	nodePhase = "phase"
)

var (
	Allocatable         opmetrics.GaugeMetric
	TotalPodRequests    opmetrics.GaugeMetric
	TotalPodLimits      opmetrics.GaugeMetric
	TotalDaemonRequests opmetrics.GaugeMetric
	TotalDaemonLimits   opmetrics.GaugeMetric
	SystemOverhead      opmetrics.GaugeMetric
	Lifetime            opmetrics.GaugeMetric
	ClusterUtilization  opmetrics.GaugeMetric
)

// Initialize metrics at runtime to ensure cloud provider's well-known labels are properly
// injected, preventing race conditions in dependency ordering during label injection for global variable. .
func initializeMetrics() { _ = "STUB: not implemented"; return }

func nodeLabelNamesWithResourceType() []string { _ = "STUB: not implemented"; return nil }

func nodeLabelNames() []string {
	_ = "STUB: not implemented"

	// WellKnownLabels includes the nodepool label, so we don't need to add it as its own item here.
	// If we do, prometheus will panic since there would be duplicate labels.
	return nil
}

type Controller struct {
	cluster     *state.Cluster
	metricStore *metrics.Store
}

func NewController(cluster *state.Cluster) *Controller { _ = "STUB: not implemented"; return nil }

func (c *Controller) Reconcile(ctx context.Context) (reconciler.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconciler.Result), nil
}

//nolint:ineffassign,staticcheck

// Build per-node metrics

// Build cluster level metric

func (c *Controller) Name() string { _ = "STUB: not implemented"; return "" }

func (c *Controller) Register(_ context.Context, m manager.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

func buildClusterUtilizationMetric(nodes state.StateNodes) []*metrics.StoreMetric {
	_ = "STUB: not implemented"

	// Aggregate resources allocated/utilized for all the nodes and pods inside the nodes
	return nil
}

// This zero check may be unnecessary. I'm erring towards caution.

// Typecast to float before the calculation to maximize resolution

func buildMetrics(n *state.StateNode) (res []*metrics.StoreMetric) {
	_ = "STUB: not implemented"
	return nil
}

func getNodeLabelsWithResourceType(node *corev1.Node, resourceTypeName string) prometheus.Labels {
	_ = "STUB: not implemented"
	return *new(prometheus.Labels)
}

func getNodeLabels(node *corev1.Node) prometheus.Labels {
	_ = "STUB: not implemented"
	return *new(prometheus.Labels)
}

// Populate well known labels

func getWellKnownLabels() map[string]string { _ = "STUB: not implemented"; return nil }

// Reformat label names to be consistent with Prometheus naming conventions (snake_case)

func resourceNameToString(resourceName corev1.ResourceName) string {
	_ = "STUB: not implemented"
	return ""
}
