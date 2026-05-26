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

	"k8s.io/utils/clock"
	"sigs.k8s.io/controller-runtime/pkg/client"

	v1 "sigs.k8s.io/karpenter/pkg/apis/v1"
	"sigs.k8s.io/karpenter/pkg/cloudprovider"
	"sigs.k8s.io/karpenter/pkg/controllers/provisioning"
	"sigs.k8s.io/karpenter/pkg/controllers/state"
	"sigs.k8s.io/karpenter/pkg/events"
)

type ValidationError struct {
	error
}

func NewValidationError(err error) *ValidationError { _ = "STUB: not implemented"; return nil }

func IsValidationError(err error) bool { _ = "STUB: not implemented"; return false }

// BudgetValidationError indicates validation failed due to disruption budget constraints
type BudgetValidationError struct {
	*ValidationError
}

func NewBudgetValidationError(err error) *BudgetValidationError {
	_ = "STUB: not implemented"
	return nil
}

func (e *BudgetValidationError) Unwrap() error { _ = "STUB: not implemented"; return nil }

// SchedulingValidationError indicates validation failed due to scheduling constraints
type SchedulingValidationError struct {
	*ValidationError
}

func NewSchedulingValidationError(err error) *SchedulingValidationError {
	_ = "STUB: not implemented"
	return nil
}

func (e *SchedulingValidationError) Unwrap() error { _ = "STUB: not implemented"; return nil }

// ChurnValidationError indicates validation failed due to churn detection
type ChurnValidationError struct {
	*ValidationError
}

func NewChurnValidationError(err error) *ChurnValidationError {
	_ = "STUB: not implemented"
	return nil
}

func (e *ChurnValidationError) Unwrap() error { _ = "STUB: not implemented"; return nil }

type Validator interface {
	Validate(context.Context, Command, time.Duration) (Command, error)
}

// Validation is used to perform validation on a consolidation command.  It makes an assumption that when re-used, all
// of the commands passed to IsValid were constructed based off of the same consolidation state.  This allows it to
// skip the validation TTL for all but the first command.
type validation struct {
	clock         clock.Clock
	cluster       *state.Cluster
	kubeClient    client.Client
	cloudProvider cloudprovider.CloudProvider
	provisioner   *provisioning.Provisioner
	recorder      events.Recorder
	queue         *Queue
	reason        v1.DisruptionReason
}

type EmptinessValidator struct {
	validation
	filter         CandidateFilter
	validationType string
}

func NewEmptinessValidator(c consolidation) *EmptinessValidator {
	_ = "STUB: not implemented"
	return nil
}

func (e *EmptinessValidator) Validate(ctx context.Context, cmd Command, validationPeriod time.Duration) (Command, error) {
	_ = "STUB: not implemented"
	return *new(Command), nil
}

type ConsolidationValidator struct {
	validation
	filter         CandidateFilter
	validationType string
}

func NewSingleConsolidationValidator(c consolidation) *ConsolidationValidator {
	_ = "STUB: not implemented"
	return nil
}

func NewMultiConsolidationValidator(c consolidation) *ConsolidationValidator {
	_ = "STUB: not implemented"
	return nil
}

func (c *ConsolidationValidator) Validate(ctx context.Context, cmd Command, validationPeriod time.Duration) (Command, error) {
	_ = "STUB: not implemented"
	return *new(Command), nil
}

func (c *ConsolidationValidator) isValid(ctx context.Context, cmd Command, validationPeriod time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// Revalidate candidates after validating the command. This mitigates the chance of a race condition outlined in
// the following GitHub issue: https://github.com/kubernetes-sigs/karpenter/issues/1167.

func (e *EmptinessValidator) validateCandidates(ctx context.Context, candidates ...*Candidate) ([]*Candidate, error) {
	_ = "STUB: not implemented"
	// This GetCandidates call filters out nodes that were nominated
	return nil, nil
}

// ValidateCandidates gets the current representation of the provided candidates and ensures that they are all still valid.
// For a candidate to still be valid, the following conditions must be met:
//
//	a. It must pass the global candidate filtering logic (no blocking PDBs, no do-not-disrupt annotation, etc)
//	b. It must not have any pods nominated for it
//	c. It must still be disruptable without violating node disruption budgets
//
// If these conditions are met for all candidates, ValidateCandidates returns a slice with the updated representations.
func (c *ConsolidationValidator) validateCandidates(ctx context.Context, candidates ...*Candidate) ([]*Candidate, error) {
	_ = "STUB: not implemented"
	// GracefulDisruptionClass is hardcoded here because ValidateCandidates is only used for consolidation disruption. All consolidation disruption is graceful disruption.
	return nil, nil
}

// If we filtered out any candidates, return nil as some NodeClaims in the consolidation decision have changed.

// Return nil if any candidate meets either of the following conditions:
//  a. A pod was nominated to the candidate
//  b. Disrupting the candidate would violate node disruption budgets

// ValidateCommand validates a command for a Method
func (v *validation) validateCommand(ctx context.Context, cmd Command, candidates []*Candidate) error {
	_ = "STUB: not implemented"
	// None of the chosen candidate are valid for execution, so retry
	return nil
}

// We want to ensure that the re-simulated scheduling using the current cluster state produces the same result.
// There are three possible options for the number of new candidates that we need to handle:
// len(NewNodeClaims) == 0, as long as we weren't expecting a new node, this is valid
// len(NewNodeClaims) > 1, something in the cluster changed so that the candidates we were going to delete can no longer
//                    be deleted without producing more than one node
// len(NewNodeClaims) == 1, as long as the noe looks like what we were expecting, this is valid

// scheduling produced zero new NodeClaims and we weren't expecting any, so this is valid.

// if it produced no new NodeClaims, but we were expecting one we should re-simulate as there is likely a better
// consolidation option now

// we need more than one replacement node which is never valid currently (all of our node replacement is m->1, never m->n)

// we now know that scheduling simulation wants to create one new node

// but we weren't expecting any new NodeClaims, so this is invalid

// We know that the scheduling simulation wants to create a new node and that the command we are verifying wants
// to create a new node. The scheduling simulation doesn't apply any filtering to instance types, so it may include
// instance types that we don't want to launch which were filtered out when the lifecycleCommand was created.  To
// check if our lifecycleCommand is valid, we just want to ensure that the list of instance types we are considering
// creating are a subset of what scheduling says we should create.  We check for a subset since the scheduling
// simulation here does no price filtering, so it will include more expensive types.
//
// This is necessary since consolidation only wants cheaper NodeClaims.  Suppose consolidation determined we should delete
// a 4xlarge and replace it with a 2xlarge. If things have changed and the scheduling simulation we just performed
// now says that we need to launch a 4xlarge. It's still launching the correct number of NodeClaims, but it's just
// as expensive or possibly more so we shouldn't validate.

// Now we know:
// - current scheduling simulation says to create a new node with types T = {T_0, T_1, ..., T_n}
// - our lifecycle command says to create a node with types {U_0, U_1, ..., U_n} where U is a subset of T

// getValidationFailureReason categorizes validation errors into specific failure types
func getValidationFailureReason(err error) string { _ = "STUB: not implemented"; return "" }
