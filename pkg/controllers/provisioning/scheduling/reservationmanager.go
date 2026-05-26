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
	"k8s.io/apimachinery/pkg/util/sets"

	"sigs.k8s.io/karpenter/pkg/cloudprovider"
)

type ReservationManager struct {
	reservations map[string]sets.Set[string] // hostname -> set[reservation id]
	capacity     map[string]int              // reservation id -> count
}

func NewReservationManager(instanceTypes map[string][]*cloudprovider.InstanceType) *ReservationManager {
	_ = "STUB: not implemented"
	return nil
}

// If we have multiple offerings with the same reservation ID, track the one with the least capacity. This could be
// the result of multiple nodepools referencing the same capacity reservation, and there being an update to the
// capacity between calls to GetInstanceTypes.

// Should always be idempotent
func (rm *ReservationManager) CanReserve(hostname string, offering *cloudprovider.Offering) bool {
	_ = "STUB: not implemented"
	return false
}

// Note: this panic should never occur, and would indicate a serious bug in the scheduling code.

// Should always be idempotent
func (rm *ReservationManager) Reserve(hostname string, offerings ...*cloudprovider.Offering) {
	_ = "STUB: not implemented"
	return
}

func (rm *ReservationManager) Release(hostname string, offerings ...*cloudprovider.Offering) {
	_ = "STUB: not implemented"
	return
}

func (rm *ReservationManager) HasReservation(hostname string, offering *cloudprovider.Offering) bool {
	_ = "STUB: not implemented"
	return false
}

func (rm *ReservationManager) RemainingCapacity(offering *cloudprovider.Offering) int {
	_ = "STUB: not implemented"
	return 0
}
