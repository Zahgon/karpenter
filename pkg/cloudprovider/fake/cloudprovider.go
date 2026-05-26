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

package fake

import (
	"context"
	"sync"

	"github.com/awslabs/operatorpkg/status"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "sigs.k8s.io/karpenter/pkg/apis/v1"
	"sigs.k8s.io/karpenter/pkg/cloudprovider"
	"sigs.k8s.io/karpenter/pkg/test/v1alpha1"
)

func init() {
	v1.WellKnownLabels = v1.WellKnownLabels.Insert(v1alpha1.LabelReservationID)
	cloudprovider.ReservationIDLabel = v1alpha1.LabelReservationID
	cloudprovider.ReservedCapacityLabels.Insert(v1alpha1.LabelReservationID)
}

var _ cloudprovider.CloudProvider = (*CloudProvider)(nil)

type CloudProvider struct {
	InstanceTypes            []*cloudprovider.InstanceType
	InstanceTypesForNodePool map[string][]*cloudprovider.InstanceType
	ErrorsForNodePool        map[string]error

	mu sync.RWMutex
	// CreateCalls contains the arguments for every create call that was made since it was cleared
	CreateCalls        []*v1.NodeClaim
	AllowedCreateCalls int
	NextCreateErr      error
	NextGetErr         error
	NextDeleteErr      error
	DeleteCalls        []*v1.NodeClaim
	GetCalls           []string

	CreatedNodeClaims         map[string]*v1.NodeClaim
	Drifted                   cloudprovider.DriftReason
	NodeClassGroupVersionKind []schema.GroupVersionKind
	RepairPolicy              []cloudprovider.RepairPolicy
}

func NewCloudProvider() *CloudProvider { _ = "STUB: not implemented"; return nil }

// Reset is for BeforeEach calls in testing to reset the tracking of CreateCalls
func (c *CloudProvider) Reset() { _ = "STUB: not implemented"; return }

//nolint:gocyclo
func (c *CloudProvider) Create(ctx context.Context, nodeClaim *v1.NodeClaim) (*v1.NodeClaim, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Order instance types so that we get the cheapest instance types of the available offerings

// Labels

// Find offering, prioritizing reserved instances

// Propagate labels dictated by offering requirements - e.g. zone, capacity-type, and reservation-id

func (c *CloudProvider) Get(_ context.Context, id string) (*v1.NodeClaim, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CloudProvider) List(_ context.Context) ([]*v1.NodeClaim, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CloudProvider) GetInstanceTypes(_ context.Context, np *v1.NodePool) ([]*cloudprovider.InstanceType, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CloudProvider) Delete(_ context.Context, nc *v1.NodeClaim) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *CloudProvider) IsDrifted(context.Context, *v1.NodeClaim) (cloudprovider.DriftReason, error) {
	_ = "STUB: not implemented"
	return *new(cloudprovider.DriftReason), nil
}

func (c *CloudProvider) RepairPolicies() []cloudprovider.RepairPolicy {
	_ = "STUB: not implemented"
	return nil

	// Name returns the CloudProvider implementation name.
}

func (c *CloudProvider) Name() string { _ = "STUB: not implemented"; return "" }

func (c *CloudProvider) GetSupportedNodeClasses() []status.Object {
	_ = "STUB: not implemented"
	return nil
}
