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

	v1 "sigs.k8s.io/karpenter/pkg/apis/v1"
)

const MultiNodeConsolidationTimeoutDuration = 1 * time.Minute
const MultiNodeConsolidationType = "multi"

type MultiNodeConsolidation struct {
	consolidation
	validator Validator
}

func NewMultiNodeConsolidation(c consolidation, opts ...option.Function[MethodOptions]) *MultiNodeConsolidation {
	_ = "STUB: not implemented"
	return nil
}

// nolint:gocyclo
func (m *MultiNodeConsolidation) ComputeCommands(ctx context.Context, disruptionBudgetMapping map[string]int, candidates ...*Candidate) ([]Command, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// In order, filter out all candidates that would violate the budget.
// Since multi-node consolidation relies on the ordering of
// these candidates, and does computation in batches of these nodes by
// simulateScheduling(nodes[0, n]), doing a binary search on n to find
// the optimal consolidation command, this pre-filters out nodes that
// would have violated the budget anyway, preserving the ordering
// and only considering a number of nodes that can be disrupted.

// If there's disruptions allowed for the candidate's nodepool,
// add it to the list of candidates, and decrement the budget.

// Filter out empty candidates. If there was an empty node that wasn't consolidated before this, we should
// assume that it was due to budgets. If we don't filter out budgets, users who set a budget for `empty`
// can find their nodes disrupted here.

// set constrainedByBudgets to true if any node was a candidate but was constrained by a budget

// Only consider a maximum batch of 100 NodeClaims to save on computation.
// This could be further configurable in the future.

// if there are no candidates because of a budget, don't mark
// as consolidated, as it's possible it should be consolidatable
// the next time we try to disrupt.

// firstNConsolidationOption looks at the first N NodeClaims to determine if they can all be consolidated at once.  The
// NodeClaims are sorted by increasing disruption order which correlates to likelihood of being able to consolidate the node
// nolint:gocyclo
func (m *MultiNodeConsolidation) firstNConsolidationOption(ctx context.Context, candidates []*Candidate, max int) (Command, error) {
	_ = "STUB: not implemented"
	// we always operate on at least two NodeClaims at once, for single NodeClaims standard consolidation will find all solutions
	return *new(Command), nil
}

// Set a timeout

// Pass the timeout context to ensure sub-operations can be canceled

// context deadline exceeded will return to the top of the loop and either return nothing or the last saved command

// ensure that the action is sensical for replacements, see explanation on filterOutSameType for why this is
// required

// we check the error before the replacement instanceTypeOptions since we return nil for the replacement if we get an error

// We can consolidate NodeClaims [0,mid]

// filterOutSameInstanceType filters out instance types that are more expensive than the cheapest instance type that is being
// consolidated if the list of replacement instance types include one of the instance types that is being removed
//
// This handles the following potential consolidation result:
// NodeClaims=[t3a.2xlarge, t3a.2xlarge, t3a.small] -> 1 of t3a.small, t3a.xlarge, t3a.2xlarge
//
// In this case, we shouldn't perform this consolidation at all.  This is equivalent to just
// deleting the 2x t3a.xlarge NodeClaims.  This code will identify that t3a.small is in both lists and filter
// out any instance type that is the same or more expensive than the t3a.small
//
// For another scenario:
// NodeClaims=[t3a.2xlarge, t3a.2xlarge, t3a.small] -> 1 of t3a.nano, t3a.small, t3a.xlarge, t3a.2xlarge
//
// This code sees that t3a.small is the cheapest type in both lists and filters it and anything more expensive out
// leaving the valid consolidation:
// NodeClaims=[t3a.2xlarge, t3a.2xlarge, t3a.small] -> 1 of t3a.nano
func filterOutSameInstanceType(replacement *Replacement, consolidate []*Candidate) (*Replacement, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// get the price of the cheapest node that we currently are considering deleting indexed by instance type

// we are considering replacing multiple NodeClaims with a single NodeClaim of one of the same types, so the replacement
// node must be cheaper than the price of the existing node, or we should just keep that one and do a
// deletion only to reduce cluster disruption (fewer pods will re-schedule).

func (m *MultiNodeConsolidation) Reason() v1.DisruptionReason {
	_ = "STUB: not implemented"
	return *new(v1.DisruptionReason)
}

func (m *MultiNodeConsolidation) Class() string { _ = "STUB: not implemented"; return "" }

func (m *MultiNodeConsolidation) ConsolidationType() string { _ = "STUB: not implemented"; return "" }
