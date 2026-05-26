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
	"k8s.io/apimachinery/pkg/types"
)

// Queue is a queue of pods that is scheduled.  It's used to attempt to schedule pods as long as we are making progress
// in scheduling. This is sometimes required to maintain zonal topology spreads with constrained pods, and can satisfy
// pod affinities that occur in a batch of pods if there are enough constraints provided.
type Queue struct {
	pods    []*v1.Pod
	lastLen map[types.UID]int
}

// NewQueue constructs a new queue given the input pods, sorting them to optimize for bin-packing into nodes.
func NewQueue(pods []*v1.Pod, podData map[types.UID]*PodData) *Queue {
	_ = "STUB: not implemented"
	return nil
}

// Pop returns the next pod or false if no longer making progress
func (q *Queue) Pop() (*v1.Pod, bool) { _ = "STUB: not implemented"; return nil, false }

// If we are about to pop a pod when it was last pushed with the same number of pods in the queue, then
// we've cycled through all pods in the queue without making progress and can stop

// Push a pod onto the queue, counting each time a pod is immediately requeued. This is used to detect staleness.
func (q *Queue) Push(pod *v1.Pod) { _ = "STUB: not implemented"; return }

func (q *Queue) List() []*v1.Pod { _ = "STUB: not implemented"; return nil }

func byCPUAndMemoryDescending(pods []*v1.Pod, podData map[types.UID]*PodData) func(i int, j int) bool {
	_ = "STUB: not implemented"
	return nil
}

// LHS has less CPU, so it should be sorted after

// If all else is equal, give a consistent ordering. This reduces the number of NominatePod events as we
// de-duplicate those based on identical content.

// unfortunately creation timestamp only has a 1-second resolution, so we would still re-order pods created
// during a deployment scale-up if we only looked at creation time

// pod UIDs aren't in any order, but since we first sort by creation time this only serves to consistently order
// pods created within the same second
