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

package scheduling

import (
	v1 "k8s.io/api/core/v1"

	"sigs.k8s.io/karpenter/pkg/controllers/state"
	"sigs.k8s.io/karpenter/pkg/scheduling"
)

type ExistingNode struct {
	*state.StateNode
	cachedAvailable v1.ResourceList // Cache so we don't have to re-subtract resources on the StateNode every time
	cachedTaints    []v1.Taint      // Cache so we don't hae to re-construct the taints each time we attempt to schedule a pod

	Pods               []*v1.Pod
	topology           *Topology
	remainingResources v1.ResourceList
	requirements       scheduling.Requirements
}

func NewExistingNode(n *state.StateNode, topology *Topology, taints []v1.Taint, daemonResources v1.ResourceList) *ExistingNode {
	_ = "STUB: not implemented"
	// The state node passed in here must be a deep copy from cluster state as we modify it
	// the remaining daemonResources to schedule are the total daemonResources minus what has already scheduled
	return nil
}

// If unexpected daemonset pods schedule to the node due to labels appearing on the node which cause the
// DS to be able to schedule, we need to ensure that we don't let our remainingDaemonResources go negative as
// it will cause us to mis-calculate the amount of remaining resources

// CanAdd returns whether the pod can be added to the ExistingNode
// based on the taints/tolerations, volume requirements, host port compatibility,
// requirements, resources, and topology requirements
func (n *ExistingNode) CanAdd(pod *v1.Pod, podData *PodData, volumes scheduling.Volumes) (updatedRequirements scheduling.Requirements, err error) {
	_ = "STUB: not implemented"
	// Check Taints
	return *new(scheduling.Requirements), nil
}

// determine the host ports that will be used if the pod schedules

// check resource requests first since that's a pretty likely reason the pod won't schedule on an in-flight
// node, which at this point can't be increased in size

// Check NodeClaim Affinity Requirements

// avoid creating our temp set of requirements until after we've ensured that at least
// the pod is compatible

// Build the list of volume requirement alternatives to try.

// Try each volume topology alternative. The selected constraints affect topology checks.

// tryVolumeAlternative attempts to add a pod with a specific set of volume requirements,
// checking topology compatibility against the existing node.
func (n *ExistingNode) tryVolumeAlternative(pod *v1.Pod, podData *PodData, baseRequirements scheduling.Requirements, volReqs scheduling.Requirements) (scheduling.Requirements, error) {
	_ = "STUB: not implemented"
	return *new(scheduling.Requirements), nil
}

// Add volume requirements to nodeRequirements ONLY (not to pod's affinity).
// This ensures the existing node satisfies the selected volume topology constraints,
// while TSC counting uses pod's original affinity.

// Check Topology Requirements
// NOTE: podData.StrictRequirements does NOT include volume requirements,
// ensuring TSC counting uses pod's original affinity.

// Add updates the ExistingNode to schedule the pod to this ExistingNode, updating
// the ExistingNode with new requirements and volumes based on the pod scheduling
func (n *ExistingNode) Add(pod *v1.Pod, podData *PodData, nodeRequirements scheduling.Requirements, volumes scheduling.Volumes) {
	_ = "STUB: not implemented"
	// Update node
	return
}
