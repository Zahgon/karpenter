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

package nodeclaim

import (
	"context"

	"github.com/awslabs/operatorpkg/status"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	v1 "sigs.k8s.io/karpenter/pkg/apis/v1"
	"sigs.k8s.io/karpenter/pkg/cloudprovider"
)

func IsManaged(nodeClaim *v1.NodeClaim, cp cloudprovider.CloudProvider) bool {
	_ = "STUB: not implemented"
	return false
}

// IsManagedPredicateFuncs is used to filter controller-runtime NodeClaim watches to NodeClaims managed by the given cloudprovider.
func IsManagedPredicateFuncs(cp cloudprovider.CloudProvider) predicate.Funcs {
	_ = "STUB: not implemented"
	return *new(predicate.Funcs)
}

func ForProviderID(providerID string) client.ListOption {
	_ = "STUB: not implemented"
	return *new(client.ListOption)
}

func ForNodePool(nodePoolName string) client.ListOption {
	_ = "STUB: not implemented"
	return *new(client.ListOption)
}

func ForNodeClass(nodeClass status.Object) client.ListOption {
	_ = "STUB: not implemented"
	return *new(client.ListOption)
}

func ListManaged(ctx context.Context, c client.Client, cloudProvider cloudprovider.CloudProvider, opts ...client.ListOption) ([]*v1.NodeClaim, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PodEventHandler is a watcher on corev1.Pods that maps Pods to NodeClaim based on the node names
// and enqueues reconcile.Requests for the NodeClaims
func PodEventHandler(c client.Client, cloudProvider cloudprovider.CloudProvider) handler.EventHandler {
	_ = "STUB: not implemented"
	return *new(handler.EventHandler)
}

// Because we get so many NodeClaims from this response, we are not DeepCopying the cached data here
// DO NOT MUTATE NodeClaims in this function as this will affect the underlying cached NodeClaim

// NodeEventHandler is a watcher on corev1.Node that maps Nodes to NodeClaims based on provider ids
// and enqueues reconcile.Requests for the NodeClaims
func NodeEventHandler(c client.Client, cloudProvider cloudprovider.CloudProvider) handler.EventHandler {
	_ = "STUB: not implemented"
	return *new(handler.EventHandler)
}

// Because we get so many NodeClaims from this response, we are not DeepCopying the cached data here
// DO NOT MUTATE NodeClaims in this function as this will affect the underlying cached NodeClaim

// NodePoolEventHandler is a watcher on v1.NodeClaim that maps NodePool to NodeClaims based
// on the v1.NodePoolLabelKey and enqueues reconcile.Requests for the NodeClaim
func NodePoolEventHandler(c client.Client, cloudProvider cloudprovider.CloudProvider) handler.EventHandler {
	_ = "STUB: not implemented"
	return *new(handler.EventHandler)
}

// Because we get so many NodeClaims from this response, we are not DeepCopying the cached data here
// DO NOT MUTATE NodeClaims in this function as this will affect the underlying cached NodeClaim

// NodeClassEventHandler is a watcher on v1.NodeClaim that maps NodeClass to NodeClaims based
// on the nodeClassRef and enqueues reconcile.Requests for the NodeClaim
func NodeClassEventHandler(c client.Client) handler.EventHandler {
	_ = "STUB: not implemented"
	return *new(handler.EventHandler)
}

// Because we get so many NodeClaims from this response, we are not DeepCopying the cached data here
// DO NOT MUTATE NodeClaims in this function as this will affect the underlying cached NodeClaim

// NodeNotFoundError is an error returned when no corev1.Nodes are found matching the passed providerID
type NodeNotFoundError struct {
	ProviderID string
}

func (e *NodeNotFoundError) Error() string { _ = "STUB: not implemented"; return "" }

func IsNodeNotFoundError(err error) bool { _ = "STUB: not implemented"; return false }

func IgnoreNodeNotFoundError(err error) error { _ = "STUB: not implemented"; return nil }

// DuplicateNodeError is an error returned when multiple corev1.Nodes are found matching the passed providerID
type DuplicateNodeError struct {
	ProviderID string
}

func (e *DuplicateNodeError) Error() string { _ = "STUB: not implemented"; return "" }

func IsDuplicateNodeError(err error) bool { _ = "STUB: not implemented"; return false }

func IgnoreDuplicateNodeError(err error) error { _ = "STUB: not implemented"; return nil }

// NodeForNodeClaim is a helper function that takes a v1.NodeClaim and attempts to find the matching corev1.Node by its providerID
// This function will return errors if:
//  1. No corev1.Nodes match the v1.NodeClaim providerID
//  2. Multiple corev1.Nodes match the v1.NodeClaim providerID
func NodeForNodeClaim(ctx context.Context, c client.Client, nodeClaim *v1.NodeClaim) (*corev1.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AllNodesForNodeClaim is a helper function that takes a v1.NodeClaim and finds ALL matching corev1.Nodes by their providerID
// If the providerID is not resolved for a NodeClaim, then no Nodes will map to it
func AllNodesForNodeClaim(ctx context.Context, c client.Client, nodeClaim *v1.NodeClaim) ([]*corev1.Node, error) {
	_ = "STUB: not implemented"
	// NodeClaims that have no resolved providerID have no nodes mapped to them
	return nil, nil
}

func UpdateNodeOwnerReferences(nodeClaim *v1.NodeClaim, node *corev1.Node) *corev1.Node {
	_ = "STUB: not implemented"
	return nil
}
