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

package nodepool

import (
	"context"

	"github.com/awslabs/operatorpkg/option"
	"github.com/awslabs/operatorpkg/status"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	v1 "sigs.k8s.io/karpenter/pkg/apis/v1"
	"sigs.k8s.io/karpenter/pkg/cloudprovider"
)

func IsManaged(nodePool *v1.NodePool, cp cloudprovider.CloudProvider) bool {
	_ = "STUB: not implemented"
	return false
}

func IsStatic(np *v1.NodePool) bool { _ = "STUB: not implemented"; return false }

func GetNodeClass(ctx context.Context, c client.Client, nodePool *v1.NodePool, cp cloudprovider.CloudProvider) (status.Object, error) {
	_ = "STUB: not implemented"
	return *new(status.Object), nil
}

// IsManagedPredicateFuncs is used to filter controller-runtime NodeClaim watches to NodeClaims managed by the given cloudprovider.
func IsManagedPredicateFuncs(cp cloudprovider.CloudProvider) predicate.Funcs {
	_ = "STUB: not implemented"
	return *new(predicate.Funcs)
}

// IsStaticPredicateFunc is used to filter controller-runtime NodePool watches to Static NodePools
func IsStaticPredicateFuncs() predicate.Funcs {
	_ = "STUB: not implemented"
	return *new(predicate.Funcs)
}

func ForNodeClass(nc status.Object) client.ListOption {
	_ = "STUB: not implemented"
	return *new(client.ListOption)
}

func ListManaged(ctx context.Context, c client.Client, cloudProvider cloudprovider.CloudProvider, opts ...client.ListOption) ([]*v1.NodePool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type NodeClaimHandlerOption struct {
	staticOnly bool
	client     client.Client // used only if staticOnly && nameFilter == nil
}

func WithStaticOnly(o *NodeClaimHandlerOption) { _ = "STUB: not implemented"; return }

func WithClient(c client.Client) func(*NodeClaimHandlerOption) {
	_ = "STUB: not implemented"
	return nil
}

func NodeClaimEventHandler(opts ...option.Function[NodeClaimHandlerOption]) handler.EventHandler {
	_ = "STUB: not implemented"
	return *new(handler.EventHandler)
}

func NodeEventHandler() handler.EventHandler {
	_ = "STUB: not implemented"
	return *new(handler.EventHandler)
}

// NodeClassEventHandler is a watcher on v1.NodePool that maps NodeClass to NodePools based
// on the nodeClassRef and enqueues reconcile.Requests for the NodePool
func NodeClassEventHandler(c client.Client) handler.EventHandler {
	_ = "STUB: not implemented"
	return *new(handler.EventHandler)
}

// OrderByWeight orders the NodePools in the provided slice by their priority weight in-place. This priority evaluates
// the following things in precedence order:
//  1. NodePools that have a larger weight are ordered first
//  2. If two NodePools have the same weight, then the NodePool with the name later in the alphabet will come first
func OrderByWeight(nps []*v1.NodePool) { _ = "STUB: not implemented"; return }

// Order NodePools by name for a consistent ordering when sorting equal weight
