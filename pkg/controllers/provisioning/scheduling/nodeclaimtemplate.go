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

package scheduling

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	v1 "sigs.k8s.io/karpenter/pkg/apis/v1"
	"sigs.k8s.io/karpenter/pkg/cloudprovider"
	"sigs.k8s.io/karpenter/pkg/scheduling"
)

// DefaultTerminationGracePeriod is used as runtime defaulting for TerminationGracePeriod on the NodeClaim
// This would be a mechanism to allow cloud providers to enforce a TerminationGracePeriod on all node
// provisioned by Karpenter
var DefaultTerminationGracePeriod *metav1.Duration = nil

// MaxInstanceTypes is a constant that restricts the number of instance types to be sent for launch. Note that this
// is intentionally changed to var just to help in testing the code.
var MaxInstanceTypes = 600

// NodeClaimTemplate encapsulates the fields required to create a node and mirrors
// the fields in NodePool. These structs are maintained separately in order
// for fields like Requirements to be able to be stored more efficiently.
type NodeClaimTemplate struct {
	v1.NodeClaim

	NodePoolName        string
	NodePoolUUID        types.UID
	NodePoolWeight      int32
	InstanceTypeOptions cloudprovider.InstanceTypes
	Requirements        scheduling.Requirements
	IsStaticNodeClaim   bool
}

func NewNodeClaimTemplate(nodePool *v1.NodePool) *NodeClaimTemplate {
	_ = "STUB: not implemented"
	return nil
}

// resolveCustomLabelsFromRequirements resolves the concrete values for user-defined labels from a NodeClaimTemplate's
// requirements.
func (i *NodeClaimTemplate) resolveCustomLabelsFromRequirements() map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (i *NodeClaimTemplate) ToNodeClaim() *v1.NodeClaim {
	_ = "STUB: not implemented"
	// Inject instanceType requirements for NodeClaims belonging to dynamic NodePool
	// For static we let cloudprovider.Create()
	return nil
}

// Order the instance types by price and only take up to MaxInstanceTypes of them to decrease the instance type size in the requirements

// Collect available capacity types from the selected instance types

// We'll assign any labels with known, concrete values at NodeClaim creation time. This includes any labels from the
// NodeClaimTemplate (since there's a single possible value), and any resolved values for custom labels in the
// NodeClaimTemplate's requirements. The latter **cannot** be instance type dependent (like well-known labels) since
// Karpenter can't reason about which label domains would belong to each instance type.
