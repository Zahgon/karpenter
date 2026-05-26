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
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/utils/clock"

	"sigs.k8s.io/karpenter/pkg/events"
)

// IsActive checks if Karpenter should consider this pod as running by ensuring that the pod:
// - Isn't a terminal pod (Failed or Succeeded)
// - Isn't actively terminating
func IsActive(pod *corev1.Pod) bool { _ = "STUB: not implemented"; return false }

// IsReschedulable checks if a Karpenter should consider this pod when re-scheduling to new capacity by ensuring that the pod:
// - Is an active pod (isn't terminal or actively terminating) OR Is owned by a StatefulSet and Is Terminating
// - Isn't owned by a DaemonSet
// - Isn't a mirror pod (https://kubernetes.io/docs/tasks/configure-pod-container/static-pod/)
func IsReschedulable(pod *corev1.Pod) bool {
	_ = "STUB: not implemented"
	// StatefulSet pods can be handled differently here because we know that StatefulSet pods MUST
	// get deleted before new pods are re-created. This means that we can model terminating pods for StatefulSets
	// differently for higher availability by considering terminating pods for scheduling
	return false
}

// IsEvictable checks if a pod is evictable by Karpenter by ensuring that the pod:
// - Is an active pod (isn't terminal or actively terminating)
// - Doesn't tolerate the "karpenter.sh/disruption=disrupting" taint
// - Isn't a mirror pod (https://kubernetes.io/docs/tasks/configure-pod-container/static-pod/)
// - Does not have an active "karpenter.sh/do-not-disrupt" annotation (https://karpenter.sh/docs/concepts/disruption/#pod-level-controls)
func IsEvictable(pod *corev1.Pod, clk clock.Clock, recorder events.Recorder) bool {
	_ = "STUB: not implemented"
	return false
}

// IsWaitingEviction checks if this is a pod that we are waiting to be removed from the node by ensuring that the pod:
// - Isn't a terminal pod (Failed or Succeeded)
// - Can be drained by Karpenter (See IsDrainable)
func IsWaitingEviction(pod *corev1.Pod, clk clock.Clock) bool {
	_ = "STUB: not implemented"
	return false
}

// IsDrainable checks if a pod can be drained by Karpenter by ensuring that the pod:
// - Doesn't tolerate the "karpenter.sh/disruption=disrupting" taint
// - Isn't a pod that has been terminating past its terminationGracePeriodSeconds
// - Isn't a mirror pod (https://kubernetes.io/docs/tasks/configure-pod-container/static-pod/)
// Note: pods with the `karpenter.sh/do-not-disrupt` annotation are included since node drain should stall until these pods are evicted or become terminal, even though Karpenter won't orchestrate the eviction.
func IsDrainable(pod *corev1.Pod, clk clock.Clock) bool { _ = "STUB: not implemented"; return false }

// Mirror pods cannot be deleted through the API server since they are created and managed by kubelet
// This means they are effectively read-only and can't be controlled by API server calls
// https://kubernetes.io/docs/reference/generated/kubectl/kubectl-commands#drain

// IsPodEligibleForForcedEviction checks if a pod needs to be deleted with a reduced grace period ensuring that the pod:
// - Is terminating
// - Has a deletion timestamp that is after the node's grace period expiration time
func IsPodEligibleForForcedEviction(pod *corev1.Pod, nodeGracePeriodExpirationTime *time.Time) bool {
	_ = "STUB: not implemented"
	return false
}

// IsProvisionable checks if a pod needs to be scheduled to new capacity by Karpenter by ensuring that the pod:
// - Has been marked as "Unschedulable" in the PodScheduled reason by the kube-scheduler
// - Has not been bound to a node
// - Isn't currently preempting other pods on the cluster and about to schedule
// - Isn't owned by a DaemonSet
// - Isn't a mirror pod (https://kubernetes.io/docs/tasks/configure-pod-container/static-pod/)
func IsProvisionable(pod *corev1.Pod) bool { _ = "STUB: not implemented"; return false }

// IsDisruptable checks if a pod can be disrupted using clock-aware logic for time-based do-not-disrupt annotations.
// It considers both boolean ("true") and duration-based values (e.g., "5m", "1h").
// For duration-based values, it checks if the pod has been running longer than the specified duration.
// Invalid annotation formats are treated as if the annotation doesn't exist and an event is emitted.
// Non-active pods are always considered disruptable.
func IsDisruptable(pod *corev1.Pod, clk clock.Clock, recorder events.Recorder) bool {
	_ = "STUB: not implemented"
	return false
}

// FailedToSchedule ensures that the kube-scheduler has seen this pod and has intentionally
// marked this pod with a condition, noting that it thinks that the pod can't schedule anywhere
// It does this by marking the pod status condition "PodScheduled" as "Unschedulable"
// Note that it's possible that other schedulers may be scheduling another pod and may have a different
// semantic (e.g. Fargate on AWS marks with MATCH_NODE_SELECTOR_FAILED). If that's the case, Karpenter
// won't react to this pod because the scheduler didn't add this specific condition.
func FailedToSchedule(pod *corev1.Pod) bool { _ = "STUB: not implemented"; return false }

func IsScheduled(pod *corev1.Pod) bool { _ = "STUB: not implemented"; return false }

func IsPreempting(pod *corev1.Pod) bool { _ = "STUB: not implemented"; return false }

func IsTerminal(pod *corev1.Pod) bool { _ = "STUB: not implemented"; return false }

func IsTerminating(pod *corev1.Pod) bool { _ = "STUB: not implemented"; return false }

func IsStuckTerminating(pod *corev1.Pod, clk clock.Clock) bool {
	_ = "STUB: not implemented"
	// The pod DeletionTimestamp will be set to the time the pod was deleted plus its
	// grace period in seconds. We give an additional minute as a buffer to allow
	// pods to force delete off the node before we actually go and terminate the node
	// so that we get less pod leaking on the cluster.
	return false
}

func IsOwnedByStatefulSet(pod *corev1.Pod) bool { _ = "STUB: not implemented"; return false }

func IsOwnedByDaemonSet(pod *corev1.Pod) bool { _ = "STUB: not implemented"; return false }

// IsOwnedByNode returns true if the pod is a static pod owned by a specific node
func IsOwnedByNode(pod *corev1.Pod) bool { _ = "STUB: not implemented"; return false }

func IsOwnedBy(pod *corev1.Pod, gvks []schema.GroupVersionKind) bool {
	_ = "STUB: not implemented"
	return false
}

// parseDoNotDisrupt parses the do-not-disrupt annotation value as a duration.
// Returns the parsed duration or an error if the value is not a valid positive duration.
func parseDoNotDisrupt(value string) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

// IsDoNotDisruptActive checks if the do-not-disrupt protection is still active for a pod
// It considers both boolean ("true") and duration-based values (e.g., "5m", "1h")
// For duration-based values, it checks if the pod has been running longer than the specified duration
// Invalid annotation formats are treated as if the annotation doesn't exist (returns false)
// and an event is emitted if a recorder is provided
func IsDoNotDisruptActive(pod *corev1.Pod, clk clock.Clock, recorder events.Recorder) bool {
	_ = "STUB: not implemented"
	return false
}

// Invalid format - emit event and treat as if annotation doesn't exist

// Check if the pod has been running longer than the grace period

// If we can't determine start time, fail safe

// Emit event when duration-based protection is still active

// Emit event when grace period has elapsed

// ToleratesDisruptedNoScheduleTaint returns true if the pod tolerates karpenter.sh/disruption:NoSchedule taint
func ToleratesDisruptedNoScheduleTaint(pod *corev1.Pod) bool {
	_ = "STUB: not implemented"
	return false
}

// HasRequiredPodAntiAffinity returns true if a non-empty PodAntiAffinity/RequiredDuringSchedulingIgnoredDuringExecution
// is defined in the pod spec
func HasRequiredPodAntiAffinity(pod *corev1.Pod) bool { _ = "STUB: not implemented"; return false }

// HasPodAntiAffinity returns true if a non-empty PodAntiAffinity is defined in the pod spec
func HasPodAntiAffinity(pod *corev1.Pod) bool { _ = "STUB: not implemented"; return false }

// HasDRARequirements returns true if any pod containers consume ResourceClaims
func HasDRARequirements(pod *corev1.Pod) bool { _ = "STUB: not implemented"; return false }
