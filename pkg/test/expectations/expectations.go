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

//nolint:revive
package expectations

import (
	"context"
	"time"

	opmetrics "github.com/awslabs/operatorpkg/metrics"
	"github.com/awslabs/operatorpkg/singleton"
	"github.com/awslabs/operatorpkg/status"
	. "github.com/onsi/ginkgo/v2" //nolint:revive
	. "github.com/onsi/gomega"    //nolint:revive
	"github.com/prometheus/client_golang/prometheus"
	prometheusmodel "github.com/prometheus/client_model/go"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	v1 "sigs.k8s.io/karpenter/pkg/apis/v1"
	"sigs.k8s.io/karpenter/pkg/cloudprovider"
	"sigs.k8s.io/karpenter/pkg/controllers/provisioning"
	"sigs.k8s.io/karpenter/pkg/controllers/provisioning/scheduling"
	"sigs.k8s.io/karpenter/pkg/controllers/state"
	"sigs.k8s.io/karpenter/pkg/controllers/state/informer"
)

const (
	ReconcilerPropagationTime = 10 * time.Second
	RequestInterval           = 1 * time.Second
)

type Bindings map[*corev1.Pod]*Binding

type Binding struct {
	NodeClaim *v1.NodeClaim
	Node      *corev1.Node
}

func (b Bindings) Get(p *corev1.Pod) *Binding { _ = "STUB: not implemented"; return nil }

func ExpectExists[T client.Object](ctx context.Context, c client.Client, obj T) T {
	_ = "STUB: not implemented"
	return *new(T)
}

func ExpectPodExists(ctx context.Context, c client.Client, name string, namespace string) *corev1.Pod {
	_ = "STUB: not implemented"
	return nil
}

func ExpectNodeExists(ctx context.Context, c client.Client, name string) *corev1.Node {
	_ = "STUB: not implemented"
	return nil
}

func ExpectNotFound(ctx context.Context, c client.Client, objects ...client.Object) {
	_ = "STUB: not implemented"
	return
}

func ExpectScheduled(ctx context.Context, c client.Client, pod *corev1.Pod) *corev1.Node {
	_ = "STUB: not implemented"
	return nil
}

func ExpectPodsScheduled(ctx context.Context, c client.Client, pods ...*corev1.Pod) {
	_ = "STUB: not implemented"
	return
}

func ExpectNotScheduled(ctx context.Context, c client.Client, pod *corev1.Pod) *corev1.Pod {
	_ = "STUB: not implemented"
	return nil
}

func ExpectApplied(ctx context.Context, c client.Client, objects ...client.Object) {
	_ = "STUB: not implemented"
	return
}

// Snapshot the status, since create/update may override

// Create or Update

// Update status

// Some objects do not have a status

// Re-get the object to grab the updated spec and status

// Set the deletion timestamp by adding a finalizer and deleting

func ExpectDeleted(ctx context.Context, c client.Client, objects ...client.Object) {
	_ = "STUB: not implemented"
	return
}

func ExpectReconciled(ctx context.Context, reconciler reconcile.Reconciler, request reconcile.Request) reconcile.Result {
	_ = "STUB: not implemented"
	return *new(reconcile.Result)
}

func ExpectReconciledFailed(ctx context.Context, reconciler reconcile.Reconciler, request reconcile.Request) reconcile.Result {
	_ = "STUB: not implemented"
	return *new(reconcile.Result)
}

func ExpectSingletonReconciled(ctx context.Context, reconciler singleton.Reconciler) reconcile.Result {
	_ = "STUB: not implemented"
	return *new(reconcile.Result)
}

func ExpectSingletonReconcileFailed(ctx context.Context, reconciler singleton.Reconciler) error {
	_ = "STUB: not implemented"
	return nil
}

func ExpectObjectReconciled[T client.Object](ctx context.Context, c client.Client, reconciler reconcile.ObjectReconciler[T], object T) reconcile.Result {
	_ = "STUB: not implemented"
	return *new(reconcile.Result)
}

func ExpectObjectReconcileFailed[T client.Object](ctx context.Context, c client.Client, reconciler reconcile.ObjectReconciler[T], object T) error {
	_ = "STUB: not implemented"
	return nil
}

// ExpectDeletionTimestampSetWithOffset ensures that the deletion timestamp is set on the objects by adding a finalizer
// and then deleting the object immediately after. This holds the object until the finalizer is patched out in the DeferCleanup
func ExpectDeletionTimestampSet(ctx context.Context, c client.Client, objects ...client.Object) {
	_ = "STUB: not implemented"
	return
}

func ExpectCleanedUp(ctx context.Context, c client.Client) { _ = "STUB: not implemented"; return }

// Fail open for CRDs that don't exist on this k8s version (e.g. ResourceClaim on < 1.34)

func ExpectFinalizersRemovedFromList(ctx context.Context, c client.Client, objectLists ...client.ObjectList) {
	_ = "STUB: not implemented"
	return
}

func ExpectFinalizersRemoved(ctx context.Context, c client.Client, objs ...client.Object) {
	_ = "STUB: not implemented"
	return
}

func ExpectProvisioned(ctx context.Context, c client.Client, cluster *state.Cluster, cloudProvider cloudprovider.CloudProvider, provisioner *provisioning.Provisioner, pods ...*corev1.Pod) Bindings {
	_ = "STUB: not implemented"
	return *new(Bindings)
}

// Only bind the pods that are passed through

// We have to manually bind the pod to the node when using a fakeClient by setting the value for pod.Spec.NodeName

// track pod bindings

//nolint:gocyclo
func ExpectProvisionedNoBinding(ctx context.Context, c client.Client, cluster *state.Cluster, cloudProvider cloudprovider.CloudProvider, provisioner *provisioning.Provisioner, pods ...*corev1.Pod) Bindings {
	_ = "STUB: not implemented"

	// Persist objects
	return *new(Bindings)
}

// TODO: Check the error on the provisioner scheduling round

// TODO: Check the error on the provisioner launch

func ExpectProvisionedResults(ctx context.Context, c client.Client, cluster *state.Cluster, cloudProvider cloudprovider.CloudProvider, provisioner *provisioning.Provisioner, pods ...*corev1.Pod) scheduling.Results {
	_ = "STUB: not implemented"

	// Persist objects
	return *new(scheduling.Results)
}

func ExpectNodeClaimDeployedNoNode(ctx context.Context, c client.Client, cloudProvider cloudprovider.CloudProvider, nc *v1.NodeClaim) (*v1.NodeClaim, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO @joinnis: Check this error rather than swallowing it. This is swallowed right now due to how we are doing some testing in the cloudprovider

// Make the nodeclaim ready in the status conditions

func ExpectNodeClaimDeployed(ctx context.Context, c client.Client, cloudProvider cloudprovider.CloudProvider, nc *v1.NodeClaim) (*v1.NodeClaim, *corev1.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Mock the nodeclaim launch and node joining at the apiserver

func ExpectNodeClaimDeployedAndStateUpdated(ctx context.Context, c client.Client, cluster *state.Cluster, cloudProvider cloudprovider.CloudProvider, nc *v1.NodeClaim) (*v1.NodeClaim, *corev1.Node) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ExpectNodeClaimsCascadeDeletion(ctx context.Context, c client.Client, nodeClaims ...*v1.NodeClaim) {
	_ = "STUB: not implemented"
	return
}

func ExpectMakeNodeClaimsInitialized(ctx context.Context, c client.Client, nodeClaims ...*v1.NodeClaim) {
	_ = "STUB: not implemented"
	return
}

func ExpectMakeNodesInitialized(ctx context.Context, c client.Client, nodes ...*corev1.Node) {
	_ = "STUB: not implemented"
	return
}

func ExpectMakeNodesNotReady(ctx context.Context, c client.Client, nodes ...*corev1.Node) {
	_ = "STUB: not implemented"
	return
}

func ExpectMakeNodesReady(ctx context.Context, c client.Client, nodes ...*corev1.Node) {
	_ = "STUB: not implemented"
	return
}

// Remove any of the known ephemeral taints to make the Node ready

func ExpectReconcileSucceeded(ctx context.Context, reconciler reconcile.Reconciler, key client.ObjectKey) reconcile.Result {
	_ = "STUB: not implemented"
	return *new(reconcile.Result)
}

func ExpectStatusConditionExists(obj status.Object, t string) status.Condition {
	_ = "STUB: not implemented"
	return *new(status.Condition)
}

func ExpectOwnerReferenceExists(obj, owner client.Object) metav1.OwnerReference {
	_ = "STUB: not implemented"
	return *new(metav1.OwnerReference)
}

// ExpectMetricName attempts to resolve a metric name from a collector. This function will work so long as the fully
// qualified name is a single metric name. This holds true for the built in types, but may not for custom collectors.
func ExpectMetricName(collector prometheus.Collector) string {
	_ = "STUB: not implemented"

	// Prometheus defines an async method to resolve the description for a collector. This is simpler than it looks,
	// Describe just returns a string through the provided channel.
	return ""
}

// Add a timeout so a failure doesn't result in stalling the entire test suite. This should never occur.

// Extract the fully qualified name from the description string. This is just different enough from json that we
// need to parse with regex.

// FindMetricWithLabelValues attempts to find a metric with a name with a set of label values
// If no metric is found, the *prometheusmodel.Metric will be nil
func FindMetricWithLabelValues(name string, labelValues map[string]string) (*prometheusmodel.Metric, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func ExpectMetricGaugeValue(collector opmetrics.GaugeMetric, expectedValue float64, labels map[string]string) {
	_ = "STUB: not implemented"
	return
}

func ExpectMetricCounterValue(collector opmetrics.CounterMetric, expectedValue float64, labels map[string]string) {
	_ = "STUB: not implemented"
	return
}

func ExpectMetricHistogramSampleCountValue(metricName string, expectedValue uint64, labels map[string]string) {
	_ = "STUB: not implemented"
	return
}

func ExpectManualBinding(ctx context.Context, c client.Client, pod *corev1.Pod, node *corev1.Node) {
	_ = "STUB: not implemented"
	return
}

func ExpectSkew(ctx context.Context, c client.Client, namespace string, constraint *corev1.TopologySpreadConstraint) Assertion {
	_ = "STUB: not implemented"
	return *new(Assertion)
}

// Check node name since hostname labels aren't applied

// ExpectResources expects all the resources in expected to exist in real with the same values
func ExpectResources(expected, real corev1.ResourceList) { _ = "STUB: not implemented"; return }

func ExpectNodes(ctx context.Context, c client.Client) []*corev1.Node {
	_ = "STUB: not implemented"
	return nil
}

func ExpectNodeClaims(ctx context.Context, c client.Client) []*v1.NodeClaim {
	_ = "STUB: not implemented"
	return nil
}

func ExpectStateNodeExists(cluster *state.Cluster, node *corev1.Node) *state.StateNode {
	_ = "STUB: not implemented"
	return nil
}

func ExpectStateNodeExistsForNodeClaim(cluster *state.Cluster, nodeClaim *v1.NodeClaim) *state.StateNode {
	_ = "STUB: not implemented"
	return nil
}

func ExpectMakeNodesAndNodeClaimsInitializedAndStateUpdated(ctx context.Context, c client.Client, nodeStateController *informer.NodeController, nodeClaimStateController *informer.NodeClaimController, nodes []*corev1.Node, nodeClaims []*v1.NodeClaim) {
	_ = "STUB: not implemented"
	return
}

// Inform cluster state about node and nodeclaim readiness

// ExpectEvicted triggers an eviction call for all the passed pods
func ExpectEvicted(ctx context.Context, c client.Client, pods ...*corev1.Pod) {
	_ = "STUB: not implemented"
	return
}

// EventuallyExpectTerminating ensures that the deletion timestamp is eventually set
// We need this since there is some propagation time for the eviction API to set the deletionTimestamp
func EventuallyExpectTerminating(ctx context.Context, c client.Client, objs ...client.Object) {
	_ = "STUB: not implemented"
	return
}

// ConsistentlyExpectNotTerminating ensures that the deletion timestamp is not set
// We need this since there is some propagation time for the eviction API to set the deletionTimestamp
func ConsistentlyExpectNotTerminating(ctx context.Context, c client.Client, objs ...client.Object) {
	_ = "STUB: not implemented"
	return
}

func ExpectParallelized(fs ...func()) { _ = "STUB: not implemented"; return }

func ExpectStateNodePoolCount(cluster *state.Cluster, npName string, r, d, pd int) {
	_ = "STUB: not implemented"
	return
}
