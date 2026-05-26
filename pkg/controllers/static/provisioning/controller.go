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

package static

import (
	"context"

	"k8s.io/utils/clock"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	"sigs.k8s.io/karpenter/pkg/controllers/provisioning"
	"sigs.k8s.io/karpenter/pkg/controllers/state"
	"sigs.k8s.io/karpenter/pkg/events"

	"sigs.k8s.io/karpenter/pkg/cloudprovider"

	v1 "sigs.k8s.io/karpenter/pkg/apis/v1"
)

type Controller struct {
	kubeClient    client.Client
	cloudProvider cloudprovider.CloudProvider
	provisioner   *provisioning.Provisioner
	cluster       *state.Cluster
}

func NewController(kubeClient client.Client, cluster *state.Cluster, recorder events.Recorder, cloudProvider cloudprovider.CloudProvider, provisioner *provisioning.Provisioner, clock clock.Clock) *Controller {
	_ = "STUB: not implemented"
	return nil
}

// Reconcile the resource
// Requeue after computing Static NodePool to ensure we don't miss any events
func (c *Controller) Name() string { _ = "STUB: not implemented"; return "" }

func (c *Controller) Reconcile(ctx context.Context, np *v1.NodePool) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// We need to wait until our representation of cluster is populated accurately with what we have in api-server or else
// we would end up over provisioning due to misrepresentation of NodePoolState in our cluster state.
// This usually happens when there is controller crash, so we check if the cluster has synced atleast once.

// Size down of replicas will be handled in deprovisioning controller to drain nodes and delete NodeClaims
// If there are drifting NodeClaims we need to count them as Active as disruption controller is in the process of creating replacements

func (c *Controller) Register(_ context.Context, m manager.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

// Reoncile on NodePool Create and Update (when replicas change or when NodePool status moves from NotReady to Ready)

// We care about Static NodeClaims Deleting (Delete and Update event that has DeletionTimeStamp newly set) as we might have to provision

func HasNodePoolReplicaOrStatusChanged(oldNP, newNP *v1.NodePool) bool {
	_ = "STUB: not implemented"
	return false
}

// np.StatusConditions.Root() may mutate the np which is unsafe when operating directly against the informer cache
