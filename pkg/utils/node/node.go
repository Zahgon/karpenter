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

package node

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	storagev1 "k8s.io/api/storage/v1"

	"k8s.io/utils/clock"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	"sigs.k8s.io/karpenter/pkg/events"

	v1 "sigs.k8s.io/karpenter/pkg/apis/v1"
	"sigs.k8s.io/karpenter/pkg/cloudprovider"
)

// NodeClaimNotFoundError is an error returned when no v1.NodeClaims are found matching the passed providerID
type NodeClaimNotFoundError struct {
	error
}

func NewNodeClaimNotFoundError(providerID string) NodeClaimNotFoundError {
	_ = "STUB: not implemented"
	return *new(NodeClaimNotFoundError)
}

func IsNodeClaimNotFoundError(err error) bool { _ = "STUB: not implemented"; return false }

func IgnoreNodeClaimNotFoundError(err error) error { _ = "STUB: not implemented"; return nil }

// DuplicateNodeClaimError is an error returned when multiple v1.NodeClaims are found matching the passed providerID
type DuplicateNodeClaimError struct {
	error
}

func NewDuplicateNodeClaimError(providerID string, nodeClaims ...*v1.NodeClaim) DuplicateNodeClaimError {
	_ = "STUB: not implemented"
	return *new(DuplicateNodeClaimError)
}

func IsDuplicateNodeClaimError(err error) bool { _ = "STUB: not implemented"; return false }

func IgnoreDuplicateNodeClaimError(err error) error { _ = "STUB: not implemented"; return nil }

// GetPods grabs all pods that are currently bound to the passed nodes
func GetPods(ctx context.Context, kubeClient client.Client, nodes ...*corev1.Node) ([]*corev1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetNodeClaims grabs all NodeClaims with a providerID that matches the provided Node
func GetNodeClaims(ctx context.Context, kubeClient client.Client, node *corev1.Node) ([]*v1.NodeClaim, error) {
	_ = "STUB: not implemented"
	// Nodes without providerID should not match any NodeClaims to prevent false positives
	// with NodeClaims that also have empty providerIDs (e.g., during NodeClaim creation)
	return nil, nil
}

// NodeClaimForNode is a helper function that takes a corev1.Node and attempts to find the matching v1.NodeClaim by its providerID
// This function will return errors if:
//  1. No v1.NodeClaims match the corev1.Node's providerID
//  2. Multiple v1.NodeClaims match the corev1.Node's providerID
func NodeClaimForNode(ctx context.Context, c client.Client, node *corev1.Node) (*v1.NodeClaim, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetCurrentlyReschedulablePods grabs all pods from the passed nodes that satisfy the IsReschedulable criteria
func GetCurrentlyReschedulablePods(ctx context.Context, kubeClient client.Client, clk clock.Clock, recorder events.Recorder, nodes ...*corev1.Node) ([]*corev1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetProvisionablePods grabs all the pods from the passed nodes that satisfy the IsProvisionable criteria
func GetProvisionablePods(ctx context.Context, kubeClient client.Client) ([]*corev1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetVolumeAttachments grabs all volumeAttachments associated with the passed node
func GetVolumeAttachments(ctx context.Context, kubeClient client.Client, node *corev1.Node) ([]*storagev1.VolumeAttachment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetCondition(n *corev1.Node, match corev1.NodeConditionType) corev1.NodeCondition {
	_ = "STUB: not implemented"
	return *new(corev1.NodeCondition)
}

func IsManaged(node *corev1.Node, cp cloudprovider.CloudProvider) bool {
	_ = "STUB: not implemented"
	return false
}

// IsManagedPredicateFuncs is used to filter controller-runtime NodeClaim watches to NodeClaims managed by the given cloudprovider.
func IsManagedPredicateFuncs(cp cloudprovider.CloudProvider) predicate.Funcs {
	_ = "STUB: not implemented"
	return *new(predicate.Funcs)
}

func NodeClaimEventHandler(c client.Client) handler.EventHandler {
	_ = "STUB: not implemented"
	return *new(handler.EventHandler)
}
