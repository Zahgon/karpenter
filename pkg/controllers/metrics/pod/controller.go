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

package pod

import (
	"context"
	"time"

	opmetrics "github.com/awslabs/operatorpkg/metrics"
	"github.com/prometheus/client_golang/prometheus"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	crmetrics "sigs.k8s.io/controller-runtime/pkg/metrics"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	"sigs.k8s.io/karpenter/pkg/controllers/state"
	"sigs.k8s.io/karpenter/pkg/metrics"
)

const (
	podName             = "name"
	podNamespace        = "namespace"
	ownerSelfLink       = "owner"
	podHostName         = "node"
	podHostZone         = "zone"
	podHostArchitecture = "arch"
	podHostInstanceType = "instance_type"
	podPhase            = "phase"
	podScheduled        = "scheduled"
	podReady            = "ready"
)

var (
	PodState = opmetrics.NewPrometheusGauge(
		crmetrics.Registry,
		prometheus.GaugeOpts{
			Namespace: metrics.Namespace,
			Subsystem: metrics.PodSubsystem,
			Name:      "state",
			Help:      "Pod state is the current state of pods. This metric can be used several ways as it is labeled by the pod name, namespace, owner, node, nodepool name, zone, architecture, capacity type, instance type, pod phase, and pod readiness.",
		},
		labelNames(),
	)
	PodStartupDurationSeconds = opmetrics.NewPrometheusSummary(
		crmetrics.Registry,
		prometheus.SummaryOpts{
			Namespace:  metrics.Namespace,
			Subsystem:  metrics.PodSubsystem,
			Name:       "startup_duration_seconds",
			Help:       "The time from pod creation until the pod is running.",
			Objectives: metrics.SummaryObjectives(),
		},
		[]string{},
	)
	PodUnstartedTimeSeconds = opmetrics.NewPrometheusGauge(
		crmetrics.Registry,
		prometheus.GaugeOpts{
			Namespace: metrics.Namespace,
			Subsystem: metrics.PodSubsystem,
			Name:      "unstarted_time_seconds",
			Help:      "The time from pod creation until the pod is running.",
		},
		[]string{podName, podNamespace},
	)
	PodBoundDurationSeconds = opmetrics.NewPrometheusHistogram(
		crmetrics.Registry,
		prometheus.HistogramOpts{
			Namespace: metrics.Namespace,
			Subsystem: metrics.PodSubsystem,
			Name:      "bound_duration_seconds",
			Help:      "The time from pod creation until the pod is bound.",
			Buckets:   metrics.DurationBuckets(),
		},
		[]string{},
	)
	PodUnboundTimeSeconds = opmetrics.NewPrometheusGauge(
		crmetrics.Registry,
		prometheus.GaugeOpts{
			Namespace: metrics.Namespace,
			Subsystem: metrics.PodSubsystem,
			Name:      "unbound_time_seconds",
			Help:      "The time from pod creation until the pod is bound.",
		},
		[]string{podName, podNamespace},
	)
	// Stage: alpha
	PodProvisioningBoundDurationSeconds = opmetrics.NewPrometheusHistogram(
		crmetrics.Registry,
		prometheus.HistogramOpts{
			Namespace: metrics.Namespace,
			Subsystem: metrics.PodSubsystem,
			Name:      "provisioning_bound_duration_seconds",
			Help:      "The time from when Karpenter first thinks the pod can schedule until it binds. Note: this calculated from a point in memory, not by the pod creation timestamp.",
			Buckets:   metrics.DurationBuckets(),
		},
		[]string{},
	)
	// Stage: alpha
	PodProvisioningUnboundTimeSeconds = opmetrics.NewPrometheusGauge(
		crmetrics.Registry,
		prometheus.GaugeOpts{
			Namespace: metrics.Namespace,
			Subsystem: metrics.PodSubsystem,
			Name:      "provisioning_unbound_time_seconds",
			Help:      "The time from when Karpenter first thinks the pod can schedule until it binds. Note: this calculated from a point in memory, not by the pod creation timestamp.",
		},
		[]string{podName, podNamespace},
	)
	// Stage: alpha
	PodProvisioningStartupDurationSeconds = opmetrics.NewPrometheusHistogram(
		crmetrics.Registry,
		prometheus.HistogramOpts{
			Namespace: metrics.Namespace,
			Subsystem: metrics.PodSubsystem,
			Name:      "provisioning_startup_duration_seconds",
			Help:      "The time from when Karpenter first thinks the pod can schedule until the pod is running. Note: this calculated from a point in memory, not by the pod creation timestamp.",
			Buckets:   metrics.DurationBuckets(),
		},
		[]string{},
	)
	// Stage: alpha
	PodProvisioningUnstartedTimeSeconds = opmetrics.NewPrometheusGauge(
		crmetrics.Registry,
		prometheus.GaugeOpts{
			Namespace: metrics.Namespace,
			Subsystem: metrics.PodSubsystem,
			Name:      "provisioning_unstarted_time_seconds",
			Help:      "The time from when Karpenter first thinks the pod can schedule until the pod is running. Note: this calculated from a point in memory, not by the pod creation timestamp.",
		},
		[]string{podName, podNamespace},
	)
	// Stage: alpha
	PodSchedulingUndecidedTimeSeconds = opmetrics.NewPrometheusGauge(
		crmetrics.Registry,
		prometheus.GaugeOpts{
			Namespace: metrics.Namespace,
			Subsystem: metrics.PodSubsystem,
			Name:      "provisioning_scheduling_undecided_time_seconds",
			Help:      "The time from when Karpenter has seen a pod without making a scheduling decision for the pod. Note: this calculated from a point in memory, not by the pod creation timestamp.",
		},
		[]string{podName, podNamespace},
	)
)

// Controller for the resource
type Controller struct {
	kubeClient  client.Client
	metricStore *metrics.Store
	cluster     *state.Cluster

	pendingPods     sets.Set[string]
	unscheduledPods sets.Set[string]
}

func labelNames() []string { _ = "STUB: not implemented"; return nil }

// NewController constructs a podController instance
func NewController(kubeClient client.Client, cluster *state.Cluster) *Controller {
	_ = "STUB: not implemented"
	return nil
}

// Reconcile executes a termination control loop for the resource
func (c *Controller) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// Delete the unstarted metric since the pod is deleted

// Delete the unbound metric since the pod is deleted

// Get the time for when we Karpenter first thought the pod was schedulable. This should be zero if we didn't simulate for this pod.

// Requeue every 30s for pods that are stuck without a state change

func (c *Controller) recordPodSchedulingUndecidedMetric(pod *corev1.Pod) {
	_ = "STUB: not implemented"
	return
}

// If we've made a decision on this pod or the pod is already bound, delete the metric idempotently and return

// If we haven't made a decision, get the time that we ACK'd the pod and emit the metric based on that

func (c *Controller) recordPodStartupMetric(pod *corev1.Pod, schedulableTime time.Time) {
	_ = "STUB: not implemented"
	return
}

// Idempotently delete the unstarted_time_seconds metric if the schedulable time is zero

// Delete the unstarted metric since the pod is now started

// Clear cluster state's representation of these pods as we don't need to keep track of them anymore

// We do not emit the startup duration metric for pods that are terminal because such pods will have
// Ready status condition set to False which will cause the metric to take negative values.

// Delete the unstarted metric since the pod is now started

// Clear cluster state's representation of these pods as we don't need to keep track of them anymore

// Idempotently delete the unstarted_time_seconds metric if the schedulable time is zero

func (c *Controller) recordPodBoundMetric(pod *corev1.Pod, schedulableTime time.Time) {
	_ = "STUB: not implemented"
	return
}

// If the podScheduled condition does not exist, or it exists and is not set to true, we emit pod_current_unbound_time_seconds metric.

// Idempotently delete the unbound_time_seconds metric if the schedulable time is zero

// Delete the unbound metric since the pod is now bound

// makeLabels creates the makeLabels using the current state of the pod
func (c *Controller) makeLabels(ctx context.Context, pod *corev1.Pod) (prometheus.Labels, error) {
	_ = "STUB: not implemented"
	return *new(prometheus.Labels), nil
}

// Selflink has been deprecated after v.1.20
// Manually generate the selflink for the first owner reference
// Currently we do not support multiple owner references

func (c *Controller) Name() string { _ = "STUB: not implemented"; return "" }

func (c *Controller) Register(_ context.Context, m manager.Manager) error {
	_ = "STUB: not implemented"
	return nil
}
