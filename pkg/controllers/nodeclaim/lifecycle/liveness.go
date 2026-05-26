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

package lifecycle

import (
	"context"
	"time"

	"k8s.io/utils/clock"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	v1 "sigs.k8s.io/karpenter/pkg/apis/v1"
	"sigs.k8s.io/karpenter/pkg/state/nodepoolhealth"
)

type Liveness struct {
	clock      clock.Clock
	kubeClient client.Client
	npState    *nodepoolhealth.State
}

// registrationTimeout is a heuristic time that we expect the node to register within
// launchTimeout is a heuristic time that we expect to be able to launch within
// If we don't see the node within this time, then we should delete the NodeClaim and try again

const (
	registrationTimeout       = time.Minute * 15
	registrationTimeoutReason = "registration_timeout"
	launchTimeout             = time.Minute * 5
	launchTimeoutReason       = "launch_timeout"
)

type NodeClaimTimeout struct {
	duration time.Duration
	reason   string
}

var (
	RegistrationTimeout = NodeClaimTimeout{
		duration: registrationTimeout,
		reason:   registrationTimeoutReason,
	}
	LaunchTimeout = NodeClaimTimeout{
		duration: launchTimeout,
		reason:   launchTimeoutReason,
	}
)

//nolint:gocyclo
func (l *Liveness) Reconcile(ctx context.Context, nodeClaim *v1.NodeClaim) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// This should never occur because if we failed to launch we requeue the object with error instead of this requeueAfter

// If the Registered statusCondition hasn't gone True during the timeout since we first updated it, we should terminate the NodeClaim
// NOTE: Timeout has to be stored and checked in the same place since l.clock can advance after the check causing a race

// Delete the NodeClaim if we believe the NodeClaim won't register since we haven't seen the node

// updateNodePoolRegistrationHealth sets the NodeRegistrationHealthy=False
// on the NodePool if the nodeClaim fails to launch/register
func (l *Liveness) updateNodePoolRegistrationHealth(ctx context.Context, nodeClaim *v1.NodeClaim) error {
	_ = "STUB: not implemented"
	return nil
}

// If the nodeClaim failed to register during the timeout set NodeRegistrationHealthy status condition on
// NodePool to False. If the launch failed get the launch failure reason and message from nodeClaim.

// We use client.MergeFromWithOptimisticLock because patching a list with a JSON merge patch
// can cause races due to the fact that it fully replaces the list on a change
// Here, we are updating the status condition list

func (l *Liveness) deleteNodeClaimForTimeout(ctx context.Context, timeout NodeClaimTimeout, nodeClaim *v1.NodeClaim) error {
	_ = "STUB: not implemented"
	return nil
}
