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

package kwok

import (
	"context"
	_ "embed"

	"github.com/awslabs/operatorpkg/status"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"sigs.k8s.io/karpenter/kwok/apis/v1alpha1"
	v1 "sigs.k8s.io/karpenter/pkg/apis/v1"
	"sigs.k8s.io/karpenter/pkg/cloudprovider"
)

func NewCloudProvider(ctx context.Context, kubeClient client.Client, instanceTypes []*cloudprovider.InstanceType) *CloudProvider {
	_ = "STUB: not implemented"
	return nil
}

type CloudProvider struct {
	kubeClient    client.Client
	instanceTypes []*cloudprovider.InstanceType
}

func (c CloudProvider) Create(ctx context.Context, nodeClaim *v1.NodeClaim) (*v1.NodeClaim, error) {
	_ = "STUB: not implemented"
	// Create the Node because KwoK nodes don't have a kubelet, which is what Karpenter normally relies on to create the node.
	return nil, nil
}

// Kick-off a goroutine to allow us to asynchronously register nodes
// We're fine to leak this because failed registration can also happen in real providers

// convert the node back into a node claim to get the chosen resolved requirement values.

func (c CloudProvider) resolveNodeClassFromNodeClaim(ctx context.Context, nodeClaim *v1.NodeClaim) (*v1alpha1.KWOKNodeClass, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c CloudProvider) Delete(ctx context.Context, nodeClaim *v1.NodeClaim) error {
	_ = "STUB: not implemented"
	return nil
}

func (c CloudProvider) Get(ctx context.Context, providerID string) (*v1.NodeClaim, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c CloudProvider) List(ctx context.Context) ([]*v1.NodeClaim, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Return the hard-coded instance types.
func (c CloudProvider) GetInstanceTypes(ctx context.Context, nodePool *v1.NodePool) ([]*cloudprovider.InstanceType, error) {
	_ = "STUB: not implemented"
	return nil,

		// Return nothing since there's no cloud provider drift.
		nil
}

func (c CloudProvider) IsDrifted(ctx context.Context, nodeClaim *v1.NodeClaim) (cloudprovider.DriftReason, error) {
	_ = "STUB: not implemented"
	return *new(cloudprovider.DriftReason), nil
}

func (c CloudProvider) Name() string { _ = "STUB: not implemented"; return "" }

func (c CloudProvider) GetSupportedNodeClasses() []status.Object {
	_ = "STUB: not implemented"
	return nil
}

func (c CloudProvider) RepairPolicies() []cloudprovider.RepairPolicy {
	_ = "STUB: not implemented"
	return nil
}

// Supported Kubelet Node Conditions

func (c CloudProvider) getInstanceType(instanceTypeName string) (*cloudprovider.InstanceType, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c CloudProvider) toNode(nodeClaim *v1.NodeClaim) (*corev1.Node, error) {
	_ = "STUB: not implemented"
	//nolint
	return nil, nil
}

// Loop through instance type values, as the node claim will only have the In operator.

// KWOK nodes don't support overriding Karpenter's WellKnownResources,
// so we only apply resource requests, since NodeOverlay will not apply.
// If this changes in the future, we'll need to update capacity and allocatable values for KWOK nodes.

func addInstanceLabels(labels map[string]string, instanceType *cloudprovider.InstanceType, nodeClaim *v1.NodeClaim, offering *cloudprovider.Offering) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// start with labels on the nodeclaim

// add the derived nodeclaim requirement labels

// ensure we have an instance type and then any instance type requirements

// add in github.com/awslabs/eks-node-viewer label so that it shows up.

// Kwok has some scalability limitations.
// Randomly add each new node to one of the pre-created kwokPartitions.

func addKwokAnnotation(annotations map[string]string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (c CloudProvider) toNodeClaim(node *corev1.Node) (*v1.NodeClaim, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
