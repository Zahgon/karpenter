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
	"context"

	v1 "k8s.io/api/core/v1"
)

type Preferences struct {
	// ToleratePreferNoSchedule controls if preference relaxation adds a toleration for PreferNoSchedule taints.  This only
	// helps if there is a corresponding taint, so we don't always add it.
	ToleratePreferNoSchedule bool
}

func (p *Preferences) Relax(ctx context.Context, pod *v1.Pod) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *Preferences) removePreferredNodeAffinityTerm(pod *v1.Pod) *string {
	_ = "STUB: not implemented"
	return nil
}

// Remove the terms if there are any (terms are an OR semantic)

// Sort descending by weight to remove heaviest preferences to try lighter ones

func (p *Preferences) removeRequiredNodeAffinityTerm(pod *v1.Pod) *string {
	_ = "STUB: not implemented"
	return nil
}

// Remove the first term if there's more than one (terms are an OR semantic), Unlike preferred affinity, we cannot remove all terms

func (p *Preferences) removeTopologySpreadScheduleAnyway(pod *v1.Pod) *string {
	_ = "STUB: not implemented"
	return nil
}

func (p *Preferences) removePreferredPodAffinityTerm(pod *v1.Pod) *string {
	_ = "STUB: not implemented"
	return nil
}

// Remove the all the terms

// Sort descending by weight to remove heaviest preferences to try lighter ones

func (p *Preferences) removePreferredPodAntiAffinityTerm(pod *v1.Pod) *string {
	_ = "STUB: not implemented"
	return nil
}

// Remove the all the terms

// Sort descending by weight to remove heaviest preferences to try lighter ones

func (p *Preferences) toleratePreferNoScheduleTaints(pod *v1.Pod) *string {
	_ = "STUB: not implemented"
	// Tolerate all Taints with PreferNoSchedule effect
	return nil
}
