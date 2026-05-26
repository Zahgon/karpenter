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

package nodepoolhealth

import (
	"sync"

	"k8s.io/apimachinery/pkg/types"

	"sigs.k8s.io/karpenter/pkg/utils/ringbuffer"
)

const (
	BufferSize     = 4
	ThresholdFalse = 0.5 // 50% of 0s for NodeRegistrationHealthy=False
)

type Status int

const (
	StatusUnknown Status = iota
	StatusHealthy
	StatusUnhealthy
)

type Tracker struct {
	sync.RWMutex
	buffer ringbuffer.RingBuffer[bool]
}

func NewTracker(capacity int) *Tracker { _ = "STUB: not implemented"; return nil }

func (t *Tracker) Update(success bool) { _ = "STUB: not implemented"; return }

func (t *Tracker) Reset() { _ = "STUB: not implemented"; return }

func (t *Tracker) Status() Status { _ = "STUB: not implemented"; return *new(Status) }

// Count number of true and false

// Determine health status based on threshold

func (t *Tracker) SetStatus(status Status) { _ = "STUB: not implemented"; return }

type State struct {
	sync.RWMutex
	trackers map[types.UID]*Tracker
}

func NewState() *State { _ = "STUB: not implemented"; return nil }

func (s *State) nodePoolNodeRegistration(nodePoolUID types.UID) *Tracker {
	_ = "STUB: not implemented"
	return nil
}

// Double-check after acquiring write lock

func (s *State) Status(nodePoolUID types.UID) Status {
	_ = "STUB: not implemented"
	return *new(Status)
}

func (s *State) Update(nodePoolUID types.UID, launchStatus bool) { _ = "STUB: not implemented"; return }

func (s *State) SetStatus(nodePoolUID types.UID, status Status) { _ = "STUB: not implemented"; return }

func (s *State) DryRun(nodePoolUID types.UID, launchStatus bool) *Tracker {
	_ = "STUB: not implemented"
	return nil
}
