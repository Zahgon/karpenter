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

package common

import (
	"context"
	"regexp"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	v1 "sigs.k8s.io/karpenter/pkg/apis/v1"
)

func (env *Environment) ExpectCreated(objects ...client.Object) { _ = "STUB: not implemented"; return }

func (env *Environment) ExpectDeleted(objects ...client.Object) { _ = "STUB: not implemented"; return }

// ExpectUpdated will update objects in the cluster to match the inputs.
// WARNING: This ignores the resource version check, which can result in
// overwriting changes made by other controllers in the cluster.
// This is useful in ensuring that we can clean up resources by patching
// out finalizers.
// Grab the object before making the updates to reduce the chance of this race.
func (env *Environment) ExpectUpdated(objects ...client.Object) { _ = "STUB: not implemented"; return }

// ExpectStatusUpdated will update objects in the cluster to match the inputs.
// WARNING: This ignores the resource version check, which can result in
// overwriting changes made by other controllers in the cluster.
// This is useful in ensuring that we can clean up resources by patching
// out finalizers.
// Grab the object before making the updates to reduce the chance of this race.
func (env *Environment) ExpectStatusUpdated(objects ...client.Object) {
	_ = "STUB: not implemented"
	return
}

func (env *Environment) ExpectReplaceNodeClassCondition(nodeclass *unstructured.Unstructured, condition metav1.Condition) *unstructured.Unstructured {
	_ = "STUB: not implemented"
	return nil
}

func (env *Environment) ExpectParsedProviderID(providerID string) string {
	_ = "STUB: not implemented"
	return ""
}

// ExpectCreatedOrUpdated can update objects in the cluster to match the inputs.
// WARNING: ExpectUpdated ignores the resource version check, which can result in
// overwriting changes made by other controllers in the cluster.
// This is useful in ensuring that we can clean up resources by patching
// out finalizers.
// Grab the object before making the updates to reduce the chance of this race.
func (env *Environment) ExpectCreatedOrUpdated(objects ...client.Object) {
	_ = "STUB: not implemented"
	return
}

func (env *Environment) ReplaceNodeConditions(node *corev1.Node, conds ...corev1.NodeCondition) *corev1.Node {
	_ = "STUB: not implemented"
	return nil
}

// ConsistentlyExpectDisruptionsUntilNoneLeft consistently ensures a max on number of concurrently disrupting and non-terminating nodes.
// This actually uses an Eventually() under the hood so that when we reach 0 tainted nodes we exit early.
// We use the StopTrying() so that we can exit the Eventually() if we've breached an assertion on total concurrency of disruptions.
// For example: if we have 5 nodes, with a budget of 2 nodes, we ensure that `disruptingNodes <= maxNodesDisrupting=2`
// We use nodesAtStart+maxNodesDisrupting to assert that we're not creating too many instances in replacement.
func (env *Environment) ConsistentlyExpectDisruptionsUntilNoneLeft(nodesAtStart, maxNodesDisrupting int, timeout time.Duration) {
	_ = "STUB: not implemented"
	return
}

// We use an eventually to exit when we detect the number of tainted/disrupted nodes matches our target.

// Grab Nodes and NodeClaims

// Don't include NodeClaims with the `Terminating` status condition, as they're not included in budgets

// Don't include Nodes whose NodeClaims have been ignored

// Filter further by the number of tainted nodes to get the number of nodes that are disrupting

func (env *Environment) EventuallyExpectLaunchedNodeClaimCount(comparator string, count int) []*v1.NodeClaim {
	_ = "STUB: not implemented"
	return nil
}

func (env *Environment) EventuallyExpectConsolidatable(nodeClaims ...*v1.NodeClaim) {
	_ = "STUB: not implemented"
	return
}

func (env *Environment) ConsistentlyExpectHealthyPods(duration time.Duration, pods ...*corev1.Pod) {
	_ = "STUB: not implemented"
	return
}

func (env *Environment) EventuallyExpectTerminating(pods ...*corev1.Pod) {
	_ = "STUB: not implemented"
	return
}

func (env *Environment) eventuallyExpectTerminatingWithTimeout(timeout time.Duration, pods ...*corev1.Pod) {
	_ = "STUB: not implemented"
	return
}

func (env *Environment) ConsistentlyExpectActivePods(duration time.Duration, pods ...*corev1.Pod) {
	_ = "STUB: not implemented"
	return
}

func (env *Environment) ExpectSettings() (res []corev1.EnvVar) {
	_ = "STUB: not implemented"
	return nil
}

func (env *Environment) ExpectSettingsReplaced(vars ...corev1.EnvVar) {
	_ = "STUB: not implemented"
	return
}

func (env *Environment) ExpectSettingsOverridden(vars ...corev1.EnvVar) {
	_ = "STUB: not implemented"
	return
}

func (env *Environment) ExpectSettingsRemoved(vars ...corev1.EnvVar) {
	_ = "STUB: not implemented"
	return
}

func (env *Environment) ExpectConfigMapExists(key types.NamespacedName) *corev1.ConfigMap {
	_ = "STUB: not implemented"
	return nil
}

func (env *Environment) ExpectConfigMapDataReplaced(key types.NamespacedName, data ...map[string]string) (changed bool) {
	_ = "STUB: not implemented"
	return false
}

// Completely replace the data

// If the data hasn't changed, we can just return and not update anything

// Update the configMap to update the settings

func (env *Environment) ExpectConfigMapDataOverridden(key types.NamespacedName, data ...map[string]string) (changed bool) {
	_ = "STUB: not implemented"
	return false
}

// If the data hasn't changed, we can just return and not update anything

// Update the configMap to update the settings

func (env *Environment) ExpectExists(obj client.Object) client.Object {
	_ = "STUB: not implemented"
	return *new(client.Object)
}

func (env *Environment) EventuallyExpectBound(pods ...*corev1.Pod) {
	_ = "STUB: not implemented"
	return
}

func (env *Environment) EventuallyExpectHealthy(pods ...*corev1.Pod) {
	_ = "STUB: not implemented"
	return
}

func (env *Environment) EventuallyExpectHealthyWithTimeout(timeout time.Duration, pods ...*corev1.Pod) {
	_ = "STUB: not implemented"
	return
}

func (env *Environment) EventuallyExpectKarpenterRestarted() { _ = "STUB: not implemented"; return }

func (env *Environment) ExpectKarpenterLeaseOwnerChanged() { _ = "STUB: not implemented"; return }

func (env *Environment) EventuallyExpectRollout(name, namespace string) {
	_ = "STUB: not implemented"
	return
}

func (env *Environment) ExpectKarpenterPods() []*corev1.Pod { _ = "STUB: not implemented"; return nil }

func (env *Environment) ExpectActiveKarpenterPodName() string { _ = "STUB: not implemented"; return "" }

// Holder identity for lease is always in the format "<pod-name>_<pseudo-random-value>

func (env *Environment) ExpectActiveKarpenterPod() *corev1.Pod {
	_ = "STUB: not implemented"
	return nil
}

func (env *Environment) EventuallyExpectPendingPodCount(selector labels.Selector, numPods int) {
	_ = "STUB: not implemented"
	return
}

func (env *Environment) EventuallyExpectBoundPodCount(selector labels.Selector, numPods int) []*corev1.Pod {
	_ = "STUB: not implemented"
	return nil
}

func (env *Environment) EventuallyExpectHealthyPodCount(selector labels.Selector, numPods int) []*corev1.Pod {
	_ = "STUB: not implemented"
	return nil
}

func (env *Environment) EventuallyExpectHealthyPodCountWithTimeout(timeout time.Duration, selector labels.Selector, numPods int) []*corev1.Pod {
	_ = "STUB: not implemented"
	return nil
}

func (env *Environment) ExpectPodsMatchingSelector(selector labels.Selector) []*corev1.Pod {
	_ = "STUB: not implemented"
	return nil
}

func (env *Environment) EventuallyExpectUniqueNodeNames(selector labels.Selector, uniqueNames int) {
	_ = "STUB: not implemented"
	return
}

func (env *Environment) eventuallyExpectScaleDown() { _ = "STUB: not implemented"; return }

// expect the current node count to be what it was when the test started

func (env *Environment) EventuallyExpectNotFound(objects ...client.Object) {
	_ = "STUB: not implemented"
	return
}

func (env *Environment) EventuallyExpectNotFoundAssertion(objects ...client.Object) AsyncAssertion {
	_ = "STUB: not implemented"
	return *new(AsyncAssertion)
}

func (env *Environment) ExpectCreatedNodeCount(comparator string, count int) []*corev1.Node {
	_ = "STUB: not implemented"
	return nil
}

func (env *Environment) ExpectNodeCount(comparator string, count int) {
	_ = "STUB: not implemented"
	return
}

func (env *Environment) ExpectNodeClaimCount(comparator string, count int) {
	_ = "STUB: not implemented"
	return
}

func NodeClaimNames(nodeClaims []*v1.NodeClaim) []string { _ = "STUB: not implemented"; return nil }

func NodeNames(nodes []*corev1.Node) []string { _ = "STUB: not implemented"; return nil }

func (env *Environment) ConsistentlyExpectNodeCount(comparator string, count int, duration time.Duration) []*corev1.Node {
	_ = "STUB: not implemented"
	return nil
}

func (env *Environment) ConsistentlyExpectNoDisruptions(nodeCount int, duration time.Duration) (taintedNodes []*corev1.Node) {
	_ = "STUB: not implemented"
	return nil
}

// ConsistentlyExpectDisruptionsWithNodeCount will continually ensure that there are exactly disruptingNodes with totalNodes (including replacements and existing nodes)
func (env *Environment) ConsistentlyExpectDisruptionsWithNodeCount(disruptingNodes, totalNodes int, duration time.Duration) (taintedNodes []*corev1.Node) {
	_ = "STUB: not implemented"
	return nil
}

// Ensure we don't change our NodeClaims

func (env *Environment) EventuallyExpectTaintedNodeCount(comparator string, count int) []*corev1.Node {
	_ = "STUB: not implemented"
	return nil
}

func (env *Environment) EventuallyExpectNodesUntaintedWithTimeout(timeout time.Duration, nodes ...*corev1.Node) {
	_ = "STUB: not implemented"
	return
}

func (env *Environment) EventuallyExpectNodeClaimCount(comparator string, count int) []*v1.NodeClaim {
	_ = "STUB: not implemented"
	return nil
}

func (env *Environment) EventuallyExpectNodeCount(comparator string, count int) []*corev1.Node {
	_ = "STUB: not implemented"
	return nil
}

func (env *Environment) EventuallyExpectNodeCountWithSelector(comparator string, count int, selector labels.Selector) []*corev1.Node {
	_ = "STUB: not implemented"
	return nil
}

func (env *Environment) EventuallyExpectCreatedNodeCount(comparator string, count int) []*corev1.Node {
	_ = "STUB: not implemented"
	return nil
}

func (env *Environment) EventuallyExpectDeletedNodeCount(comparator string, count int) []*corev1.Node {
	_ = "STUB: not implemented"
	return nil
}

func (env *Environment) EventuallyExpectDeletedNodeCountWithSelector(comparator string, count int, selector labels.Selector) []*corev1.Node {
	_ = "STUB: not implemented"
	return nil
}

func (env *Environment) EventuallyExpectInitializedNodeCount(comparator string, count int) []*corev1.Node {
	_ = "STUB: not implemented"
	return nil
}

func (env *Environment) EventuallyExpectCreatedNodeClaimCount(comparator string, count int) []*v1.NodeClaim {
	_ = "STUB: not implemented"
	return nil
}

func (env *Environment) EventuallyExpectNodeClaimsReady(nodeClaims ...*v1.NodeClaim) {
	_ = "STUB: not implemented"
	return
}

func (env *Environment) EventuallyExpectDrifted(nodeClaims ...*v1.NodeClaim) {
	_ = "STUB: not implemented"
	return
}

// ExpectBlockNodeRegistration sets up and validates a node registration blocking mechanism using ValidatingAdmissionPolicy.
// It creates a policy that prevents nodes from registering if they have the label 'registration: fail'.
//
// The function performs the following steps:
// 1. Verifies the cluster version is 1.28 or higher (requirement for ValidatingAdmissionPolicy)
// 2. Creates an admission policy that specifically targets node creation
// 3. Creates a binding for the admission policy to enforce the validation
// 4. Ensures the policy is active through validation testing
//
// Note: Requires Kubernetes version 1.28+ to function properly.
func (env *Environment) ExpectBlockNodeRegistration() { _ = "STUB: not implemented"; return }

// Define the ValidatingAdmissionPolicy that will inspect node creation requests
// The policy's validation expression checks if the 'registration' label equals 'fail'

// Create the binding that connects the admission policy to the cluster's admission chain

// Create both the policy and binding in the cluster

// Wait for the admission policy to become active
// Note: There can be a delay between resource creation and policy enforcement
// We use a dry-run node creation attempt to verify the policy is active

// ExpectBlockNodeClassStatus sets up a nodeclass status update blocking mechanism using ValidatingAdmissionPolicy.
// It creates a policy that prevents nodeclassess from updating their status
//
// The function performs the following steps:
// 1. Verifies the cluster version is 1.28 or higher (requirement for ValidatingAdmissionPolicy)
// 2. Creates an admission policy that specifically targets nodeclass status updates
// 3. Creates a binding for the admission policy to enforce the validation
//
// Note: Requires Kubernetes version 1.28+ to function properly.
func (env *Environment) ExpectBlockNodeClassStatus(nodeClass *unstructured.Unstructured) {
	_ = "STUB: not implemented"
	return
}

// Define the ValidatingAdmissionPolicy that will inspect node creation requests
// The policy's validation expression checks if the 'registration' label equals 'fail'

// Blocks status condition updates that lack the 'TestingNotReady' reason field.
// This prevents the Karpenter controller from modifying status conditions
// while allowing our test suite to make updates. This provides a deterministic
// mechanism for E2E tests to update NodeClass conditions.

// Create the binding that connects the admission policy to the cluster's admission chain

// Create both the policy and binding in the cluster

// Wait for the admission policy to become active
// Note: There can be a delay between resource creation and policy enforcement
// We use a dry-run nodeclass status update attempt to verify the policy is active

func (env *Environment) ConsistentlyExpectNodeClaimsNotDrifted(duration time.Duration, nodeClaims ...*v1.NodeClaim) {
	_ = "STUB: not implemented"
	return
}

func (env *Environment) ConsistentlyExpectNodeClaimCountNotExceed(duration time.Duration, count int) {
	_ = "STUB: not implemented"
	return
}

func (env *Environment) EventuallyExpectEmpty(nodeClaims ...*v1.NodeClaim) {
	_ = "STUB: not implemented"
	return
}

func (env *Environment) GetNode(nodeName string) corev1.Node {
	_ = "STUB: not implemented"
	return *new(corev1.Node)
}

func (env *Environment) ExpectNoCrashes() { _ = "STUB: not implemented"; return }

var (
	lastLogged = metav1.Now()
)

func (env *Environment) printControllerLogs(options *corev1.PodLogOptions) {
	_ = "STUB: not implemented"
	return
}

// local version of the log options

func (env *Environment) EventuallyExpectMinUtilization(resource corev1.ResourceName, comparator string, value float64) {
	_ = "STUB: not implemented"
	return
}

func (env *Environment) EventuallyExpectAvgUtilization(resource corev1.ResourceName, comparator string, value float64) {
	_ = "STUB: not implemented"
	return
}

func (env *Environment) ExpectDaemonSetEnvironmentVariableUpdated(obj client.ObjectKey, name, value string, containers ...string) {
	_ = "STUB: not implemented"
	return
}

// If the env var already exists, update its value. Otherwise, create a new var.

// ForcePodsToSpread ensures that currently scheduled pods get spread evenly across all passed nodes by deleting pods off of existing
// nodes and waiting them to reschedule. This is useful for scenarios where you want to force the nodes be underutilized
// but you want to keep a consistent count of nodes rather than leaving around empty ones.
func (env *Environment) ForcePodsToSpread(nodes ...*corev1.Node) {
	_ = "STUB: not implemented"

	// Get the total count of pods across
	return
}

// Set the nodes to unschedulable so that the pods won't reschedule.

// TODO: Consider moving this time check to an Eventually poll. This gets a little tricker with helper functions
// since you need to make sure that your Expectation helper functions are scoped to to your "g Gomega" scope
// so that you don't fail the first time you get a failure on your expectation

func (env *Environment) ExpectActivePodsForNode(nodeName string) []*corev1.Pod {
	_ = "STUB: not implemented"
	return nil
}

func (env *Environment) ExpectCABundle() string {
	_ = "STUB: not implemented"
	// Discover CA Bundle from the REST client. We could alternatively
	// have used the simpler client-go InClusterConfig() method.
	// However, that only works when Karpenter is running as a Pod
	// within the same cluster it's managing.
	return ""
}

// fills in CAData!

func (env *Environment) GetDaemonSetCount(np *v1.NodePool) int {
	_ = "STUB: not implemented"

	// Performs the same logic as the scheduler to get the number of daemonset
	// pods that we estimate we will need to schedule as overhead to each node
	return 0
}

func (env *Environment) GetDaemonSetOverhead(np *v1.NodePool) corev1.ResourceList {
	_ = "STUB: not implemented"

	// Performs the same logic as the scheduler to get the number of daemonset
	// pods that we estimate we will need to schedule as overhead to each node
	return *new(corev1.ResourceList)
}

type PrometheusMetric struct {
	Name   string
	Labels map[string]string
	Value  float64
}

func (env *Environment) ExpectPodMetrics() (res []PrometheusMetric) {
	_ = "STUB: not implemented"
	return nil
}

//nolint:errorlint

var (
	prometheusMetricRegex = regexp.MustCompile(`(?P<Name>.*){(?P<Labels>.*)} (?P<Value>\d*(?:\.\d*)?)`)
)

func parseMetricsLine(line string) (metric PrometheusMetric, err error) {
	_ = "STUB: not implemented"
	return *new(PrometheusMetric), nil
}

func (env *Environment) ExpectPodPortForwarded(ctx context.Context, pod *corev1.Pod, podPort, localPort int) {
	_ = "STUB: not implemented"
	return
}

func (env *Environment) GetNodePoolCost(nodePoolName string) float64 {
	_ = "STUB: not implemented"
	return 0
}
