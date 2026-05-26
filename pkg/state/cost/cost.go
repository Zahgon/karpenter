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

package cost

import (
	"context"
	"sync"

	opmetrics "github.com/awslabs/operatorpkg/metrics"
	"github.com/prometheus/client_golang/prometheus"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	crmetrics "sigs.k8s.io/controller-runtime/pkg/metrics"

	v1 "sigs.k8s.io/karpenter/pkg/apis/v1"
	"sigs.k8s.io/karpenter/pkg/cloudprovider"
	"sigs.k8s.io/karpenter/pkg/metrics"
)

// NecessaryLabels defines the set of required Kubernetes labels that must be present
// on NodeClaim objects for cost tracking to function properly.
var NecessaryLabels = []string{corev1.LabelInstanceTypeStable, v1.CapacityTypeLabelKey, corev1.LabelTopologyZone, v1.NodePoolLabelKey}

var (
	CostTrackingErrorsTotal = opmetrics.NewPrometheusCounter(
		crmetrics.Registry,
		prometheus.CounterOpts{
			Namespace: metrics.Namespace,
			Subsystem: metrics.NodePoolSubsystem,
			Name:      "cost_tracker_errors_total",
			Help:      "Number of errors encountered during cost tracking operations. Labeled by nodepool and nodeclaim.",
		},
		[]string{
			metrics.NodePoolLabel,
		},
	)
)

// ClusterCost tracks the cost of compute resources across all NodePools in a cluster.
// This is an alpha-level component and its API may change without notice.
//
// The ClusterCost maintains real-time cost information by:
// - Tracking NodeClaim additions and removals
// - Managing instance type offerings and their prices
// - Calculating aggregate costs per NodePool and cluster-wide
//
// All operations are thread-safe through internal locking mechanisms.
type ClusterCost struct {
	sync.RWMutex
	npCostMap map[string]*NodePoolCost // nodepool.Name -> NodePoolCost
	// nodeClaimSet tracks which NodeClaims are currently being monitored for cost
	nodeClaimMap map[types.NamespacedName]NodeClaimMetaData // nodeClaim object key -> NodeClaimMetaData

	cloudProvider cloudprovider.CloudProvider
	client        client.Client
}

// NodePoolCost represents the cost tracking information for a single NodePool.
// It maintains the current cost, available instance types, and count of active offerings.
type NodePoolCost struct {
	cost float64
	// offeringCounts tracks how many instances of each offering type are currently active
	offeringCounts map[OfferingKey]OfferingCount
}

// OfferingKey uniquely identifies a specific compute offering by its zone,
// capacity type (e.g., spot/on-demand), and instance type name.
// This is not a hard invariant
type OfferingKey struct {
	Zone, CapacityType, InstanceName string
}

// OfferingCount tracks the number and cost of instances for a specific offering.
type OfferingCount struct {
	Count int
	Price float64 // Price of the offering, not Price * count
}

type NodeClaimMetaData struct {
	NodePoolName string
	NodeClaimKey OfferingKey
}

// NewClusterCost creates and initializes a new ClusterCost instance for tracking
// compute costs across the cluster. It requires a cloud provider for accessing
// instance type and pricing information, and a Kubernetes client for NodePool loofferingKeyups.
func NewClusterCost(ctx context.Context, cloudProvider cloudprovider.CloudProvider, client client.Client) *ClusterCost {
	_ = "STUB: not implemented"
	return nil
}

// UpdateOfferings updates the available instance types and their pricing information
// for a specific NodePool. This method is typically called when NodePool configurations
// change or when cloud provider pricing information is refreshed.
//
// Returns an error if instance type information cannot be updated or if cost
// recalculation fails.
func (cc *ClusterCost) UpdateOfferings(ctx context.Context, np *v1.NodePool, instanceTypes []*cloudprovider.InstanceType) {
	_ = "STUB: not implemented"
	return
}

func (cc *ClusterCost) internalNodepoolUpdate(ctx context.Context, np *v1.NodePool) error {
	_ = "STUB: not implemented"
	return nil
}

func (cc *ClusterCost) internalUpdateOfferings(np *v1.NodePool, instanceTypes []*cloudprovider.InstanceType) {
	_ = "STUB: not implemented"
	return
}

// Add back all of the offering counts that don't exist in the new instance types
// This can't occur on container restart, so we may lose cost data from offerings that are no longer returned
// from the cloud provider but still have nodeclaims.

// re-calculate the cost as the instances have changed

func (npc *NodePoolCost) updateCost() float64 { _ = "STUB: not implemented"; return 0 }

// add the new price times the count of that offering

func (cc *ClusterCost) createNewNodePoolCost(npName string, instanceTypes []*cloudprovider.InstanceType) {
	_ = "STUB: not implemented"
	// create the new npc
	return
}

// UpdateNodeClaim adds a NodeClaim to cost tracking. The NodeClaim must have
// all required labels or it will be ignored and logged as an error.
func (cc *ClusterCost) UpdateNodeClaim(ctx context.Context, nodeClaim *v1.NodeClaim) error {
	_ = "STUB: not implemented"
	return nil
}

// First lets check if the right labels are there

// not technically a failure mode as we expect to retry once the
// labels are propagated

// DeleteNodeClaim removes a NodeClaim from cost tracking. If the NodeClaim
// was not being tracked, this operation is a no-op.
func (cc *ClusterCost) DeleteNodeClaim(ctx context.Context, nn types.NamespacedName) error {
	_ = "STUB: not implemented"
	return nil
}

// If it succeeds, we can remove the metadata

func (cc *ClusterCost) DeleteNodePool(ctx context.Context, npName string) {
	_ = "STUB: not implemented"
	return
}

// internalAddOffering updates the internal clusterCost state to include a new offering for a given nodepool.
// It is used to increment the overall cost when a node joins the cluster. It is only called by UpdateNodeClaim
// after that function has determined if a nodeclaim is new.
func (cc *ClusterCost) internalAddOffering(ctx context.Context, npName string, offeringKey OfferingKey) error {
	_ = "STUB: not implemented"
	return nil
}

// create the new npc

// our offerings must be out of date, we should update and retry

// Start at 0; the unconditional oc.Count += 1 below accounts for this add.

// internalRemoveOffering updates the internal clusterCost state to remove an existing offering for a given nodepool.
// It is used to decrement the overall cost when a node leeaves the cluster. It is only called by DeleteNodeClaim
// after that function has determined if a nodeclaim is already being accounted for.
func (cc *ClusterCost) internalRemoveOffering(npName string, offeringKey OfferingKey) error {
	_ = "STUB: not implemented"
	return nil
}

func (cc *ClusterCost) Reset() { _ = "STUB: not implemented"; return }

// GetClusterCost returns the total cost of all compute resources across
// all NodePools in the cluster.
func (cc *ClusterCost) GetClusterCost() float64 { _ = "STUB: not implemented"; return 0 }

// GetNodepoolCost returns the total cost of compute resources for a specific
// NodePool. Returns 0 if the NodePool is not being tracked.
func (cc *ClusterCost) GetNodepoolCost(np *v1.NodePool) float64 {
	_ = "STUB: not implemented"
	return 0
}

func nodeClaimMissingLabels(nc v1.NodeClaim) bool { _ = "STUB: not implemented"; return false }
