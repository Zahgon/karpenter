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

	"github.com/awslabs/operatorpkg/option"

	v1 "sigs.k8s.io/karpenter/pkg/apis/v1"
)

// Emptiness is a subreconciler that deletes empty candidates.
type Emptiness struct {
	consolidation
	validator Validator
}

func NewEmptiness(c consolidation, opts ...option.Function[MethodOptions]) *Emptiness {
	_ = "STUB: not implemented"
	return nil
}

// ShouldDisrupt is a predicate used to filter candidates
func (e *Emptiness) ShouldDisrupt(_ context.Context, c *Candidate) bool {
	_ = "STUB: not implemented"
	return false
}

// If consolidation is disabled, don't do anything. This emptiness should run for both WhenEmpty and WhenEmptyOrUnderutilized

// return true if there are no pods and the nodeclaim is consolidatable

// ComputeCommand generates a disruption command given candidates
//
//nolint:gocyclo
func (e *Emptiness) ComputeCommands(ctx context.Context, disruptionBudgetMapping map[string]int, candidates ...*Candidate) ([]Command, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// set constrainedByBudgets to true if any node was a candidate but was constrained by a budget

// If there's disruptions allowed for the candidate's nodepool,
// add it to the list of candidates, and decrement the budget.

// none empty, so do nothing

// if there are no candidates, but a nodepool had a fully blocking budget,
// don't mark the cluster as consolidated, as it's possible this nodepool
// should be consolidated the next time we try to disrupt.

func (e *Emptiness) Reason() v1.DisruptionReason {
	_ = "STUB: not implemented"
	return *new(v1.DisruptionReason)
}

func (e *Emptiness) Class() string { _ = "STUB: not implemented"; return "" }

func (e *Emptiness) ConsolidationType() string { _ = "STUB: not implemented"; return "" }
