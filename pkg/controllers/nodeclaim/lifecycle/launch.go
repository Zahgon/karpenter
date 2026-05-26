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

	"github.com/patrickmn/go-cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	v1 "sigs.k8s.io/karpenter/pkg/apis/v1"
	"sigs.k8s.io/karpenter/pkg/cloudprovider"
	"sigs.k8s.io/karpenter/pkg/events"
)

type Launch struct {
	kubeClient    client.Client
	cloudProvider cloudprovider.CloudProvider
	cache         *cache.Cache // exists due to eventual consistency on the cache
	recorder      events.Recorder
}

func (l *Launch) Reconcile(ctx context.Context, nodeClaim *v1.NodeClaim) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// Ensure that we always set the status condition to the latest generation

// Once the NodeClaim has successfully marked as launched, we no longer need to store it

// One of the following scenarios can happen with a NodeClaim that isn't marked as launched:
//  1. It was already launched by the CloudProvider but the client-go cache wasn't updated quickly enough or
//     patching failed on the status. In this case, we use the in-memory cached value for the created NodeClaim.
//  2. It is a standard NodeClaim launch where we should call CloudProvider Create() and fill in details of the launched
//     NodeClaim into the NodeClaim CR.

// Either the Node launch failed or the Node was deleted due to InsufficientCapacity/NodeClassNotReady/NotFound

func (l *Launch) launchNodeClaim(ctx context.Context, nodeClaim *v1.NodeClaim) (*v1.NodeClaim, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func PopulateNodeClaimDetails(nodeClaim, retrieved *v1.NodeClaim) *v1.NodeClaim {
	_ = "STUB: not implemented"
	// These are ordered in priority order so that user-defined nodeClaim labels and requirements trump retrieved labels
	// or the static nodeClaim labels
	return nil
}

// CloudProvider-resolved labels
// User-defined labels

func truncateMessage(msg string) string { _ = "STUB: not implemented"; return "" }
