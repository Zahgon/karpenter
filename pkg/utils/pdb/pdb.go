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

package pdb

import (
	"context"

	v1 "k8s.io/api/core/v1"
	policyv1 "k8s.io/api/policy/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/utils/clock"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"sigs.k8s.io/karpenter/pkg/events"
)

type evictionBlocker int

const (
	zeroDisruptions evictionBlocker = iota
	fullyBlockingPDBs
)

// Limits is used to evaluate if evicting a list of pods is possible.
type Limits []*pdbItem

func NewLimits(ctx context.Context, kubeClient client.Client) (Limits, error) {
	_ = "STUB: not implemented"
	return *new(Limits), nil
}

// CanEvictPods returns true if every pod in the list is evictable. They may not all be evictable simultaneously, but
// for every PDB that controls the pods at least one pod can be evicted.
// nolint:gocyclo
func (l Limits) CanEvictPods(pods []*v1.Pod, clk clock.Clock, recorder events.Recorder) ([]client.ObjectKey, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// isFullyBlocked returns true if the given pod is fully blocked by a PDB.
func (l Limits) isFullyBlocked(pod *v1.Pod, clk clock.Clock, recorder events.Recorder) ([]client.ObjectKey, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// nolint:gocyclo
func (l Limits) isEvictable(pod *v1.Pod, clk clock.Clock, recorder events.Recorder, evictionBlocker evictionBlocker) ([]client.ObjectKey, bool) {
	_ = "STUB: not implemented"
	// If the pod isn't eligible for being evicted, then the predicate doesn't matter
	// This is due to the fact that we won't call the eviction API on these pods when we are disrupting the node
	return nil, false
}

// Regardless of whether the PDBs allow disruptions, Kubernetes doesn't support multiple PDBs on a single pod:
// https://github.com/kubernetes/kubernetes/blob/84cacae7046df93c1f6f8ea97c912d948e1ad06a/pkg/registry/core/pod/storage/eviction.go#L226

// if the PDB policy is set to allow evicting unhealthy pods, then it won't stop us from
// evicting unhealthy pods

// IsCurrentlyReschedulable checks if a Karpenter should consider this pod when re-scheduling to new capacity by ensuring that the pod:
// - Is reschedulable as per the checks in IsReschedulable(...)
// - Does not have an active "karpenter.sh/do-not-disrupt" annotation (https://karpenter.sh/docs/concepts/disruption/#pod-level-controls)
// - Does not have fully blocking PDBs which would prevent the pod from being evicted
// The way this is different from IsReschedulable is that this also considers non-permanent conditions which prevent a pod from being rescheduled
// to a different node like the "do-not-disrupt" annotation or fully blocking PDBs.
func (l Limits) IsCurrentlyReschedulable(pod *v1.Pod, clk clock.Clock, recorder events.Recorder) bool {
	_ = "STUB: not implemented"
	// Don't provision capacity for pods which will not get evicted due to fully blocking PDBs.
	// Since Karpenter doesn't know when these pods will be successfully evicted, spinning up capacity until these pods are evicted is wasteful.
	return false
}

type pdbItem struct {
	key                         client.ObjectKey
	selector                    labels.Selector
	disruptionsAllowed          int32
	isFullyBlocking             bool
	canAlwaysEvictUnhealthyPods bool
}

// nolint:gocyclo
func newPdb(pdb policyv1.PodDisruptionBudget) (*pdbItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
