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

package test

import (
	"sync"

	"sigs.k8s.io/karpenter/pkg/events"
)

var _ events.Recorder = (*EventRecorder)(nil)

// EventRecorder is a mock event recorder that is used to facilitate testing.
type EventRecorder struct {
	mu     sync.RWMutex
	calls  map[string]int
	events []events.Event
}

func NewEventRecorder() *EventRecorder { _ = "STUB: not implemented"; return nil }

func (e *EventRecorder) Publish(evts ...events.Event) { _ = "STUB: not implemented"; return }

func (e *EventRecorder) Calls(reason string) int { _ = "STUB: not implemented"; return 0 }

func (e *EventRecorder) Reset() { _ = "STUB: not implemented"; return }

func (e *EventRecorder) Events() (res []events.Event) { _ = "STUB: not implemented"; return nil }

func (e *EventRecorder) ForEachEvent(f func(evt events.Event)) { _ = "STUB: not implemented"; return }

func (e *EventRecorder) DetectedEvent(msg string) bool { _ = "STUB: not implemented"; return false }
