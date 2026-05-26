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

package terminator

import (
	"context"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/utils/clock"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"sigs.k8s.io/karpenter/pkg/events"
)

type Terminator struct {
	clock         clock.Clock
	kubeClient    client.Client
	evictionQueue *Queue
	recorder      events.Recorder
}

func NewTerminator(clk clock.Clock, kubeClient client.Client, eq *Queue, recorder events.Recorder) *Terminator {
	_ = "STUB: not implemented"
	return nil
}

// Taint idempotently adds a given taint to a node with a NodeClaim
func (t *Terminator) Taint(ctx context.Context, node *corev1.Node, taint corev1.Taint) error {
	_ = "STUB: not implemented"
	return nil

	// If the node already has the correct taint (key and effect), do nothing.
}

// Otherwise, if the taint key exists (but with a different effect), remove it.

// Adding this label to the node ensures that the node is removed from the load-balancer target group
// while it is draining and before it is terminated. This prevents 500s coming prior to health check
// when the load balancer controller hasn't yet determined that the node and underlying connections are gone
// https://github.com/aws/aws-node-termination-handler/issues/316
// https://github.com/aws/karpenter/pull/2518

// We use client.MergeFromWithOptimisticLock because patching a list with a JSON merge patch
// can cause races due to the fact that it fully replaces the list on a change
// Here, we are updating the taint list

// Drain evicts pods from the node and returns true when all pods are evicted
// https://kubernetes.io/docs/concepts/architecture/nodes/#graceful-node-shutdown
func (t *Terminator) Drain(ctx context.Context, node *corev1.Node, nodeGracePeriodExpirationTime *time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

// Monitor pods in pod groups that either haven't been evicted or are actively evicting

// Only add pods to the eviction queue that haven't been evicted yet

func (t *Terminator) groupPodsByPriority(pods []*corev1.Pod) [][]*corev1.Pod {
	_ = "STUB: not implemented"
	// 1. Prioritize noncritical pods, non-daemon pods https://kubernetes.io/docs/concepts/architecture/nodes/#graceful-node-shutdown
	return nil
}

func (t *Terminator) DeleteExpiringPods(ctx context.Context, pods []*corev1.Pod, nodeGracePeriodTerminationTime *time.Time) error {
	_ = "STUB: not implemented"
	return nil

	// check if the node has an expiration time and the pod needs to be deleted
}

// delete pod proactively to give as much of its terminationGracePeriodSeconds as possible for deletion
// ensure that we clamp the maximum pod terminationGracePeriodSeconds to the node's remaining expiration time in the delete command

// ignore 404, not a problem
// otherwise, bubble up the error

// if a pod should be deleted to give it the full terminationGracePeriodSeconds of time before the node will shut down, return the time the pod should be deleted
func (t *Terminator) podDeleteTimeWithGracePeriod(nodeGracePeriodExpirationTime *time.Time, pod *corev1.Pod) *time.Time {
	_ = "STUB: not implemented"
	return nil
}

// k8s defaults to 30s, so we should never see a nil TerminationGracePeriodSeconds

// calculate the time the pod should be deleted to allow it's full grace period for termination, equal to its terminationGracePeriodSeconds before the node's expiration time
// eg: if a node will be force terminated in 30m, but the current pod has a grace period of 45m, we return a time of 15m ago
