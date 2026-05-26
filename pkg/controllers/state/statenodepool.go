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

package state

import (
	"sync"
	"sync/atomic"

	"k8s.io/apimachinery/pkg/util/sets"

	v1 "sigs.k8s.io/karpenter/pkg/apis/v1"
)

// Currently NodeClaims be in one of these states
type NodeClaimState struct {
	// NodeClaims that have been launched
	Active sets.Set[string]
	// NodeClaims that are pending disruption
	PendingDisruption sets.Set[string]
	// NodeClaims are marked for Deletion
	Deleting sets.Set[string]
}

// StateNodePool is a cached version of a NodePool in the cluster that maintains state which is expensive to compute every time it's needed.
type NodePoolState struct {
	mu sync.RWMutex

	nodePoolNameToNodeClaimState map[string]NodeClaimState // node pool name -> node claim state (Active and Deleting node claim names)
	nodeClaimNameToNodePoolName  map[string]string         // node claim name -> node pool name
	nodePoolNameToNodePoolLimit  map[string]*atomic.Int64  // node pool -> nodepool limit
}

func NewNodePoolState() *NodePoolState { _ = "STUB: not implemented"; return nil }

// Helper methods that hold the lock

// Sets up the NodePoolState to track NodeClaim
func (n *NodePoolState) SetNodeClaimMapping(npName, ncName string) {
	_ = "STUB: not implemented"
	return
}

// Marks the given NodeClaim as active in NodePoolState
func (n *NodePoolState) MarkNodeClaimActive(npName, ncName string) {
	_ = "STUB: not implemented"
	return
}

// Marks the given NodeClaim as Deleting in NodePoolState
func (n *NodePoolState) MarkNodeClaimDeleting(npName, ncName string) {
	_ = "STUB: not implemented"
	return
}

// Marks the given NodeClaim as Deleting in NodePoolState
func (n *NodePoolState) MarkNodeClaimPendingDisruption(npName, ncName string) {
	_ = "STUB: not implemented"
	return
}

// Cleans up the NodeClaim in NodePoolState and NodePool keys if NodePool is deleted or its sized down to 0
func (n *NodePoolState) Cleanup(ncName string) { _ = "STUB: not implemented"; return }

// Returns the current NodeClaims for a NodePool by its state (active, deleting, pendingdisruption)
func (n *NodePoolState) GetNodeCount(npName string) (active, deleting, pendingdisruption int) {
	_ = "STUB: not implemented"
	return 0, 0, 0
}

// ReserveNodeCount attempts to reserve nodes against a NodePool's limit.
// It ensures that the total of active nodes + deleting nodes + reserved nodes doesn't exceed the limit.
func (n *NodePoolState) ReserveNodeCount(np string, limit int64, wantedLimit int64) int64 {
	_ = "STUB: not implemented"
	return 0
}

// We retry until CompareAndSwap is successful

// ReleaseNodeCount releases the NodePoolTracker ReservedNodeLimit
func (n *NodePoolState) ReleaseNodeCount(npName string, count int64) {
	_ = "STUB: not implemented"
	return
}

// We retry until CompareAndSwap is successful

// Methods that expect the caller to hold the lock

func (n *NodePoolState) nodeCounts(npName string) (active, deleting, pendingdisruption int) {
	_ = "STUB: not implemented"
	return 0, 0, 0
}

// Updates the NodeClaim state and releases the Limit if NodeClaim transitions from not Active to Active
func (n *NodePoolState) UpdateNodeClaim(nodeClaim *v1.NodeClaim, markedForDeletion bool) {
	_ = "STUB: not implemented"
	// If we are launching/deleting a NodeClaim we need to track the state of the NodeClaim and its limits
	return
}

// If our node/nodeclaim is marked for deletion, we need to make sure that we delete it

func (n *NodePoolState) ensureNodePoolEntry(np string) { _ = "STUB: not implemented"; return }

func (n *NodePoolState) Reset() { _ = "STUB: not implemented"; return }
