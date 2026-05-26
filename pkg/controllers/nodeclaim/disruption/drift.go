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

package disruption

import (
	"context"

	"github.com/patrickmn/go-cache"
	"k8s.io/utils/clock"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	v1 "sigs.k8s.io/karpenter/pkg/apis/v1"
	"sigs.k8s.io/karpenter/pkg/cloudprovider"
)

const (
	NodePoolDrifted      cloudprovider.DriftReason = "NodePoolDrifted"
	RequirementsDrifted  cloudprovider.DriftReason = "RequirementsDrifted"
	InstanceTypeNotFound cloudprovider.DriftReason = "InstanceTypeNotFound"
)

// Drift is a nodeclaim sub-controller that adds or removes status conditions on drifted nodeclaims
type Drift struct {
	clock                          clock.Clock
	cloudProvider                  cloudprovider.CloudProvider
	instanceTypeNotFoundCheckCache *cache.Cache
}

func (d *Drift) Reconcile(ctx context.Context, nodePool *v1.NodePool, nodeClaim *v1.NodeClaim) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// From here there are three scenarios to handle:
// 1. If NodeClaim is not launched, remove the drift status condition

// 2. Otherwise, if the NodeClaim isn't drifted, but has the status condition, remove it.

// 3. Finally, if the NodeClaim is drifted, but doesn't have status condition, add it.

// Requeue after 5 minutes for the cache TTL

// isDrifted will check if a NodeClaim is drifted from the fields in the NodePool Spec and the CloudProvider
func (d *Drift) isDrifted(ctx context.Context, nodePool *v1.NodePool, nodeClaim *v1.NodeClaim) (cloudprovider.DriftReason, error) {
	_ = "STUB: not implemented"
	// First check for static drift or node requirements have drifted to save on API calls.
	return *new(cloudprovider.DriftReason), nil
}

// To reduce the amount of GetInstanceTypes() calls that we make per-NodeClaim, only check this for a NodeClaim once every 30m and don't start checking it until 1h after creation
// It's alright to be more delayed with instance type drift since this is a cloudprovider-generated set of options rather than a user-defined field

// Include instance type checking separate from the other two to reduce the amount of times we grab the instance types.

// Only add a cache entry once we've validated that an instance type exists. We only cache a successful check rather
// that the result to ensure we respond quickly to transient abnormalities in the GetInstanceTypes response.

// Then check if it's drifted from the cloud provider side.

// InstanceType Offerings should return the full list of allowed instance types, even if they're temporarily
// unavailable. If we can't find the instance type that the NodeClaim is running with, or if we don't find
// a compatible offering for that given instance type (zone and capacity type being the only added in requirements),
// then we consider it drifted, since we don't have the data for this instance type anymore.
// Note this is different than RequirementDrift, where we only compare if the NodePool is compatible with the NodeClaim
// based on the requirements constructed from the NodeClaim Labels.
// InstanceTypeNotFoundDrift is computed by looking directly at the list of instance types and comparing the NodeClaim
// to the instance types and its offerings.
// 1. The NodeClaim doesn't have the instance type label
// 2. The NodeClaim has an instance type that doesn't exist in the cloudprovider instance types
// 3. There are no offerings that match the requirements
func instanceTypeNotFound(its []*cloudprovider.InstanceType, nodeClaim *v1.NodeClaim) cloudprovider.DriftReason {
	_ = "STUB: not implemented"
	return *new(cloudprovider.DriftReason)
}

// The reserved capacity type is special because a NodeClaim can be demoted from reserved to on-demand after creation.
// For this reason, when evaluating drift due to unavailable offerings, we should check both reserved and on-demand for
// reserved nodeclaims. This ensures we don't drift a nodeclaim whoes label hasn't been updated yet. If the NodePool
// isn't compatible with on-demand, this will be caught in subsequent iterations by requirements drift. For a similar
// reason we don't compare against the reservation ID and leave that to the provider to implement.

// Eligible fields for drift are described in the docs
// https://karpenter.sh/docs/concepts/disruption/#drift
func areStaticFieldsDrifted(nodePool *v1.NodePool, nodeClaim *v1.NodeClaim) cloudprovider.DriftReason {
	_ = "STUB: not implemented"
	return *new(cloudprovider.DriftReason)
}

// validate that the hash version on the NodePool is the same as the NodeClaim before evaluating for static drift

func areRequirementsDrifted(nodePool *v1.NodePool, nodeClaim *v1.NodeClaim) cloudprovider.DriftReason {
	_ = "STUB: not implemented"
	return *new(cloudprovider.DriftReason)
}

// Every nodepool requirement is compatible with the NodeClaim label set
