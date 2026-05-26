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
	"sync"
	"time"

	"k8s.io/utils/clock"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	v1 "sigs.k8s.io/karpenter/pkg/apis/v1"
	"sigs.k8s.io/karpenter/pkg/controllers/provisioning"
	"sigs.k8s.io/karpenter/pkg/controllers/state"
	"sigs.k8s.io/karpenter/pkg/events"
)

const (
	queueBaseDelay          = 1 * time.Second
	queueMaxDelay           = 10 * time.Second
	minRetryDuration        = 10 * time.Minute
	maxRetryDuration        = 1 * time.Hour
	maxConcurrentReconciles = 100
	retryDurationScale      = 80 * time.Millisecond
)

type UnrecoverableError struct {
	error
}

func NewUnrecoverableError(err error) *UnrecoverableError { _ = "STUB: not implemented"; return nil }

func IsUnrecoverableError(err error) bool { _ = "STUB: not implemented"; return false }

func (q *Queue) GetMaxRetryDuration() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

type Queue struct {
	sync.RWMutex
	ProviderIDToCommand map[string]*Command // providerID -> command, maps a candidate to its command
	source              chan event.TypedGenericEvent[*v1.NodeClaim]
	kubeClient          client.Client
	recorder            events.Recorder
	cluster             *state.Cluster
	clock               clock.Clock
	provisioner         *provisioning.Provisioner
}

// NewQueue creates a queue that will asynchronously orchestrate disruption commands
func NewQueue(kubeClient client.Client, recorder events.Recorder, cluster *state.Cluster, clock clock.Clock,
	provisioner *provisioning.Provisioner,
) *Queue {
	_ = "STUB: not implemented"

	// nolint:staticcheck
	// We need to implement a deprecated interface since Command currently doesn't implement "comparable"
	return nil
}

func (q *Queue) Name() string { _ = "STUB: not implemented"; return "" }

func (q *Queue) Register(ctx context.Context, m manager.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

func (q *Queue) Reconcile(ctx context.Context, nodeClaim *v1.NodeClaim) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// If recoverable, re-queue and try again.

// If the command failed, bail on the action.
// 1. Emit metrics for launch failures
// 2. Ensure cluster state no longer thinks these nodes are deleting
// 3. Remove it from the Queue's internal data structure

// Log the error

// waitOrTerminate will wait until launched nodeclaims are ready.
// Once the replacements are ready, it will terminate the candidates.
// nolint:gocyclo
func (q *Queue) waitOrTerminate(ctx context.Context, cmd *Command) (err error) {
	_ = "STUB: not implemented"
	// We use the number of commands in the queue as a proxy for cloud provider traffic.
	// As the number of commands increase, we expect more delays and scale the retry duration accordingly.
	return nil
}

// Wrap an error in an unrecoverable error if it timed out

// If we know the node claim is Initialized, no need to check again.

// Get the nodeclaim

// The NodeClaim got deleted after an initial eventual consistency delay
// This means that there was an ICE error or the Node initializationTTL expired
// In this case, the error is unrecoverable, so don't requeue.

// We emitted this event when disruption was blocked on launching/termination.
// This does not block other forms of deprovisioning, but we should still emit this.

// If we have any errors, don't continue

// All replacements have been provisioned.
// All we need to do now is get a successful delete call for each node claim,
// then the termination controller will handle the eventual deletion of the nodes.

// If there were any deletion failures, we should requeue.
// In the case where we requeue, but the timeout for the command is reached, we'll mark this as a failure.

// markDisrupted taints the node and adds the Disrupted condition to the NodeClaim for a candidate that is about to be disrupted
// For static NodeClaims, we mark NodeClaims as pendingdisruption in statenodepool
func (q *Queue) markDisrupted(ctx context.Context, cmd *Command) ([]*Candidate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// refresh nodeclaim before updating status

// Mark all StaticNodeClaims as pendingdisruption in nodepoolstate

// createReplacementNodeClaims creates replacement NodeClaims
func (q *Queue) createReplacementNodeClaims(ctx context.Context, cmd *Command) error {
	_ = "STUB: not implemented"
	return nil
}

// shouldn't ever occur since a partially failed CreateNodeClaims should return an error

// StartCommand will do the following:
// 1. Taint candidate nodes
// 2. Spin up replacement nodes
// 3. Add Command to the queue to wait to delete the candidates.
func (q *Queue) StartCommand(ctx context.Context, cmd *Command) error {
	_ = "STUB: not implemented"
	// First check if we can add the command.
	return nil
}

// Cordon the old nodes before we launch the replacements to prevent new pods from scheduling to the old nodes

// If we get a failure marking some nodes as disrupted, if we are launching replacements, we shouldn't continue
// with disrupting the candidates. If it's just a delete operation, we can proceed

// Update the command to only consider the successfully MarkDisrupted candidates

// If we failed to launch the replacement, don't disrupt.  If this is some permanent failure,
// we don't want to disrupt workloads with no way to provision new nodes for them.

// IMPORTANT
// We must MarkForDeletion AFTER we launch the replacements and not before
// The reason for this is to avoid producing double-launches
// If we MarkForDeletion before we create replacements, it's possible for the provisioner
// to recognize that it needs to launch capacity for terminating pods, causing us to launch
// capacity for these pods twice instead of just once

// Nominate each node for scheduling and emit pod nomination events
// We emit all nominations before we exit the disruption loop as
// we want to ensure that nodes that are nominated are respected in the subsequent
// disruption reconciliation. This is essential in correctly modeling multiple
// disruption commands in parallel.
// This will only nominate nodes for 2 * batchingWindow. Once the candidates are
// tainted with the Karpenter taint, the provisioning controller will continue
// to do scheduling simulations and nominate the pods on the candidate nodes until
// the node is cleaned up.

// IMPORTANT
// We are adding the first nodeclaim in the list of candidates into the reconciliation queue
// This invariant SHOULD NOT be relied on anywhere else besides within this file.

// An action is only performed and pods/nodes are only disrupted after a successful add to the queue

// HasAny checks to see if the candidate is part of an currently executing command.
func (q *Queue) HasAny(ids ...string) bool { _ = "STUB: not implemented"; return false }

// If the mapping has at least one of the candidates' providerIDs, return true.

// For TESTING ONLY
// This function is not thread safe as it returns pointers to commands.
// If you edit these commands returned, you can create race conditions.
func (q *Queue) GetCommands() []*Command { _ = "STUB: not implemented"; return nil }

// CompleteCommand fully clears the queue of all references of a hash/command
func (q *Queue) CompleteCommand(cmd *Command) { _ = "STUB: not implemented"; return }

// Remove all candidates linked to the command

func (q *Queue) IsEmpty() bool { _ = "STUB: not implemented"; return false }
