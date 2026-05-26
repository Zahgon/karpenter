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

package events

import (
	corev1 "k8s.io/api/core/v1"

	v1 "sigs.k8s.io/karpenter/pkg/apis/v1"
	"sigs.k8s.io/karpenter/pkg/events"
)

func Launching(nodeClaim *v1.NodeClaim, reason string) events.Event {
	_ = "STUB: not implemented"
	return *new(events.Event)
}

func WaitingOnReadiness(nodeClaim *v1.NodeClaim) events.Event {
	_ = "STUB: not implemented"
	return *new(events.Event)
}

func Terminating(node *corev1.Node, nodeClaim *v1.NodeClaim, reason string) []events.Event {
	_ = "STUB: not implemented"
	return nil
}

// Unconsolidatable is an event that informs the user that a NodeClaim/Node combination cannot be consolidated
// due to the state of the NodeClaim/Node or due to some state of the pods that are scheduled to the NodeClaim/Node
func Unconsolidatable(node *corev1.Node, nodeClaim *v1.NodeClaim, msg string) []events.Event {
	_ = "STUB: not implemented"
	return nil
}

// Blocked is an event that informs the user that a NodeClaim/Node combination is blocked on deprovisioning
// due to the state of the NodeClaim/Node or due to some state of the pods that are scheduled to the NodeClaim/Node
func Blocked(node *corev1.Node, nodeClaim *v1.NodeClaim, msg string) (evs []events.Event) {
	_ = "STUB: not implemented"
	return nil
}

func NodePoolBlockedForDisruptionReason(nodePool *v1.NodePool, reason v1.DisruptionReason) events.Event {
	_ = "STUB: not implemented"
	return *new(events.Event)
}

func NodePoolBlocked(nodePool *v1.NodePool) events.Event {
	_ = "STUB: not implemented"
	return *new(events.Event)
}

// Set a small timeout as a NodePool's disruption budget can change every minute.

// ConsolidationCandidate is an event that informs the user that a consolidation candidate has been generated
func ConsolidationCandidate(node *corev1.Node, nodeClaim *v1.NodeClaim, command string, savings float64) []events.Event {
	_ = "STUB: not implemented"
	return nil
}

// ConsolidationRejected is an event that informs the user that a consolidation candidate was rejected during validation
func ConsolidationRejected(node *corev1.Node, nodeClaim *v1.NodeClaim, command string, reason string, savings float64) []events.Event {
	_ = "STUB: not implemented"
	return nil
}
