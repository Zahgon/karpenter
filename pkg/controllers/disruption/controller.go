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

	"github.com/awslabs/operatorpkg/option"
	"github.com/awslabs/operatorpkg/reconciler"

	"k8s.io/utils/clock"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	"sigs.k8s.io/karpenter/pkg/cloudprovider"
	"sigs.k8s.io/karpenter/pkg/controllers/provisioning"
	"sigs.k8s.io/karpenter/pkg/controllers/state"
	"sigs.k8s.io/karpenter/pkg/events"
)

type Controller struct {
	queue         *Queue
	kubeClient    client.Client
	cluster       *state.Cluster
	provisioner   *provisioning.Provisioner
	recorder      events.Recorder
	clock         clock.Clock
	cloudProvider cloudprovider.CloudProvider
	methods       []Method
	mu            sync.Mutex
	lastRun       map[string]time.Time
}

// pollingPeriod that we inspect cluster to look for opportunities to disrupt
const pollingPeriod = 10 * time.Second

type ControllerOptions struct {
	methods []Method
}

func WithMethods(methods ...Method) option.Function[ControllerOptions] {
	_ = "STUB: not implemented"
	return nil
}

func NewController(clk clock.Clock, kubeClient client.Client, provisioner *provisioning.Provisioner,
	cp cloudprovider.CloudProvider, recorder events.Recorder, cluster *state.Cluster, queue *Queue, opts ...option.Function[ControllerOptions]) *Controller {
	_ = "STUB: not implemented"
	return nil
}

func NewMethods(clk clock.Clock, cluster *state.Cluster, kubeClient client.Client, provisioner *provisioning.Provisioner, cp cloudprovider.CloudProvider, recorder events.Recorder, queue *Queue) []Method {
	_ = "STUB: not implemented"
	return nil
}

// Delete any empty NodeClaims as there is zero cost in terms of disruption.

// Terminate and create replacement for drifted NodeClaims in Static NodePool

// Terminate any NodeClaims that have drifted from provisioning specifications, allowing the pods to reschedule.

// Attempt to identify multiple NodeClaims that we can consolidate simultaneously to reduce pod churn

// And finally fall back our single NodeClaim consolidation to further reduce cluster cost.

func (c *Controller) Name() string { _ = "STUB: not implemented"; return "" }

func (c *Controller) Register(_ context.Context, m manager.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Controller) Reconcile(ctx context.Context) (reconciler.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconciler.Result), nil
}

// this won't catch if the reconciler loop hangs forever, but it will catch other issues

// Log if there are any budgets that are misconfigured that weren't caught by validation.
// Only validate the first reason, since CEL validation will catch invalid disruption reasons

// We need to ensure that our internal cluster state mechanism is synced before we proceed
// with making any scheduling decision off of our state nodes. Otherwise, we have the potential to make
// a scheduling decision based on a smaller subset of nodes in our cluster state than actually exist.

// Karpenter taints nodes with a karpenter.sh/disruption taint as part of the disruption process while it progresses in memory.
// If Karpenter restarts or fails with an error during a disruption action, some nodes can be left tainted.
// Idempotently remove this taint from candidates that are not in the orchestration queue before continuing.

// Attempt different disruption methods. We'll only let one method perform an action

// All methods did nothing, so return nothing to do

func (c *Controller) disrupt(ctx context.Context, disruption Method) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// If there are no candidates, move to the next disruption

// Determine the disruption action

// Assign common fields

// Attempt to disrupt

func (c *Controller) recordRun(s string) { _ = "STUB: not implemented"; return }

func (c *Controller) logAbnormalRuns(ctx context.Context) { _ = "STUB: not implemented"; return }

// logInvalidBudgets will log if there are any invalid schedules detected
func (c *Controller) logInvalidBudgets(ctx context.Context) { _ = "STUB: not implemented"; return }

// Use a dummy value of 100 since we only care if this errors.

// Prevent duplicate error message
