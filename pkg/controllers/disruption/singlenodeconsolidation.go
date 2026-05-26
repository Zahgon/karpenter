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
	"time"

	"github.com/awslabs/operatorpkg/option"
	"k8s.io/apimachinery/pkg/util/sets"

	v1 "sigs.k8s.io/karpenter/pkg/apis/v1"
)

var SingleNodeConsolidationTimeoutDuration = 3 * time.Minute

const SingleNodeConsolidationType = "single"

// SingleNodeConsolidation is the consolidation controller that performs single-node consolidation.
type SingleNodeConsolidation struct {
	consolidation
	PreviouslyUnseenNodePools sets.Set[string]
	validator                 Validator
}

func NewSingleNodeConsolidation(c consolidation, opts ...option.Function[MethodOptions]) *SingleNodeConsolidation {
	_ = "STUB: not implemented"
	return nil
}

// ComputeCommand generates a disruption command given candidates
// nolint:gocyclo
func (s *SingleNodeConsolidation) ComputeCommands(ctx context.Context, disruptionBudgetMapping map[string]int, candidates ...*Candidate) ([]Command, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Set a timeout

// Track that we've seen this nodepool

// If the disruption budget doesn't allow this candidate to be disrupted,
// continue to the next candidate. We don't need to decrement any budget
// counter since single node consolidation commands can only have one candidate.

// Filter out empty candidates. If there was an empty node that wasn't consolidated before this, we should
// assume that it was due to budgets. If we don't filter out budgets, users who set a budget for `empty`
// can find their nodes disrupted here.

// compute a possible consolidation option

// if there are no candidates because of a budget, don't mark
// as consolidated, as it's possible it should be consolidatable
// the next time we try to disrupt.

func (s *SingleNodeConsolidation) Reason() v1.DisruptionReason {
	_ = "STUB: not implemented"
	return *new(v1.DisruptionReason)
}

func (s *SingleNodeConsolidation) Class() string { _ = "STUB: not implemented"; return "" }

func (s *SingleNodeConsolidation) ConsolidationType() string { _ = "STUB: not implemented"; return "" }

// sortCandidates interweaves candidates from different nodepools and prioritizes nodepools
// that timed out in previous runs
func (s *SingleNodeConsolidation) SortCandidates(ctx context.Context, candidates []*Candidate) []*Candidate {
	_ = "STUB: not implemented"

	// First sort by disruption cost as the base ordering
	return nil
}

func (s *SingleNodeConsolidation) shuffleCandidates(ctx context.Context, nodePoolCandidates map[string][]*Candidate) []*Candidate {
	_ = "STUB: not implemented"
	return nil

	// Log any timed out nodepools that we're prioritizing
}

// Find the maximum number of candidates in any nodepool

// Interweave candidates from different nodepools
