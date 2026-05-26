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

package disruption

import (
	"context"
	"fmt"

	"k8s.io/utils/clock"
	"sigs.k8s.io/controller-runtime/pkg/client"

	v1 "sigs.k8s.io/karpenter/pkg/apis/v1"
	"sigs.k8s.io/karpenter/pkg/cloudprovider"
	"sigs.k8s.io/karpenter/pkg/controllers/provisioning"
	"sigs.k8s.io/karpenter/pkg/controllers/provisioning/scheduling"
	"sigs.k8s.io/karpenter/pkg/controllers/state"
	"sigs.k8s.io/karpenter/pkg/events"
)

var errCandidateDeleting = fmt.Errorf("candidate is deleting")

//nolint:gocyclo
func SimulateScheduling(ctx context.Context, kubeClient client.Client, cluster *state.Cluster, provisioner *provisioning.Provisioner, clk clock.Clock, recorder events.Recorder,
	candidates ...*Candidate,
) (scheduling.Results, error) {
	_ = "STUB: not implemented"
	return *new(scheduling.Results), nil
}

// We do one final check to ensure that the node that we are attempting to consolidate isn't
// already handled for deletion by some other controller. This could happen if the node was markedForDeletion
// between returning the candidates and getting the stateNodes above

// start by getting all pending pods

// Don't provision capacity for pods which will not get evicted due to fully blocking PDBs.
// Since Karpenter doesn't know when these pods will be successfully evicted, spinning up capacity until
// these pods are evicted is wasteful.

// We get the pods that are on nodes that are deleting

// We consider existing nodes for scheduling. When these nodes are unmanaged, their taint logic should
// tell us if we can schedule to them or not; however, if these nodes are managed, we will still schedule to them
// even if they are still in the middle of their initialization loop. In the case of disruption, we don't want
// to proceed disrupting if our scheduling decision relies on nodes that haven't entered a terminal state.

// Only add a pod scheduling error if it isn't on an already deleting node.
// If the pod is on a deleting node, we assume one of two things has already happened:
// 1. The node was manually terminated, at which the provisioning controller has scheduled or is scheduling a node
//    for the pod.
// 2. The node was chosen for a previous disruption command, we assume that the uninitialized node will come up
//    for this command, and we assume it will be successful. If it is not successful, the node will become
//    not terminating, and we will no longer need to consider these pods.

// UninitializedNodeError tracks a special pod error for disruption where pods schedule to a node
// that hasn't been initialized yet, meaning that we can't be confident to make a disruption decision based off of it
type UninitializedNodeError struct {
	*scheduling.ExistingNode
}

func NewUninitializedNodeError(node *scheduling.ExistingNode) *UninitializedNodeError {
	_ = "STUB: not implemented"
	return nil
}

func (u *UninitializedNodeError) Error() string { _ = "STUB: not implemented"; return "" }

// instanceTypesAreSubset returns true if the lhs slice of instance types are a subset of the rhs.
func instanceTypesAreSubset(lhs []*cloudprovider.InstanceType, rhs []*cloudprovider.InstanceType) bool {
	_ = "STUB: not implemented"
	return false
}

// GetCandidates returns nodes that appear to be currently deprovisionable based off of their nodePool
func GetCandidates(ctx context.Context, cluster *state.Cluster, kubeClient client.Client, recorder events.Recorder, clk clock.Clock,
	cloudProvider cloudprovider.CloudProvider, shouldDisrupt CandidateFilter, disruptionClass string, queue *Queue,
) ([]*Candidate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Filter only the valid candidates that we should disrupt

// BuildNodePoolMap builds a provName -> nodePool map and a provName -> instanceName -> instance type map
func BuildNodePoolMap(ctx context.Context, kubeClient client.Client, cloudProvider cloudprovider.CloudProvider) (map[string]*v1.NodePool, map[string]map[string]*cloudprovider.InstanceType, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// don't error out on building the node pool, we just won't be able to handle any nodes that
// were created by it

// BuildDisruptionBudgets prepares our disruption budget mapping. The disruption budget maps each disruption reason to the number of allowed disruptions.
// We calculate allowed disruptions by taking the max disruptions allowed by disruption reason and subtracting the number of nodes that are NotReady and already being deleted by that disruption reason.
//
//nolint:gocyclo
func BuildDisruptionBudgetMapping(ctx context.Context, cluster *state.Cluster, clk clock.Clock, kubeClient client.Client, cloudProvider cloudprovider.CloudProvider, recorder events.Recorder, reason v1.DisruptionReason) (map[string]int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// map[nodepool] -> node count in nodepool
// map[nodepool] -> nodes undergoing disruption

// We only consider nodes that we own and are initialized towards the total.
// If a node is launched/registered, but not initialized, pods aren't scheduled
// to the node, and these are treated as unhealthy until they're cleaned up.
// This prevents odd roundup cases with percentages where replacement nodes that
// aren't initialized could be counted towards the total, resulting in more disruptions
// to active nodes than desired, where Karpenter should wait for these nodes to be
// healthy before continuing.

// Additionally, don't consider nodeclaims that have the terminating condition. A nodeclaim should have
// the Terminating condition only when the node is drained and cloudprovider.Delete() was successful
// on the underlying cloud provider machine.

// If the node satisfies one of the following, we subtract it from the allowed disruptions.
// 1. Has a NotReady conditiion
// 2. Is marked as disrupting

// mapCandidates maps the list of proposed candidates with the current state
func mapCandidates(proposed, current []*Candidate) []*Candidate {
	_ = "STUB: not implemented"
	return nil
}
