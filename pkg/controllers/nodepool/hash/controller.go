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

package hash

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	v1 "sigs.k8s.io/karpenter/pkg/apis/v1"
	"sigs.k8s.io/karpenter/pkg/cloudprovider"
)

// Controller is hash controller that constructs a hash based on the fields that are considered for static drift.
// The hash is placed in the metadata for increased observability and should be found on each object.
type Controller struct {
	kubeClient    client.Client
	cloudProvider cloudprovider.CloudProvider
}

func NewController(kubeClient client.Client, cloudProvider cloudprovider.CloudProvider) *Controller {
	_ = "STUB: not implemented"
	return nil
}

// Reconcile the resource
func (c *Controller) Name() string { _ = "STUB: not implemented"; return "" }

func (c *Controller) Reconcile(ctx context.Context, np *v1.NodePool) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

func (c *Controller) Register(ctx context.Context, m manager.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

// Updating `nodepool-hash-version` annotation inside the karpenter controller means a breaking change has been made to the hash calculation.
// The `nodepool-hash` annotation on the NodePool will be updated, due to the breaking change, making the `nodepool-hash` on the NodeClaim different from
// NodePool. Since, we cannot rely on the `nodepool-hash` on the NodeClaims, due to the breaking change, we will need to re-calculate the hash and update the annotation.
// For more information on the Drift Hash Versioning: https://github.com/kubernetes-sigs/karpenter/blob/main/designs/drift-hash-versioning.md
func (c *Controller) updateNodeClaimHash(ctx context.Context, np *v1.NodePool) error {
	_ = "STUB: not implemented"
	return nil
}

// Any NodeClaim that is already drifted will remain drifted if the karpenter.sh/nodepool-hash-version doesn't match
// Since the hashing mechanism has changed we will not be able to determine if the drifted status of the NodeClaim has changed
