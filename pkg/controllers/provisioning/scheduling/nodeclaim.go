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
	"context"

	corev1 "k8s.io/api/core/v1"

	"sigs.k8s.io/karpenter/pkg/cloudprovider"
	"sigs.k8s.io/karpenter/pkg/scheduling"
)

// NodeClaim is a set of constraints, compatible pods, and possible instance types that could fulfill these constraints. This
// will be turned into one or more actual node instances within the cluster after bin packing.
type NodeClaim struct {
	NodeClaimTemplate

	Pods               []*corev1.Pod
	reservationManager *ReservationManager
	topology           *Topology
	hostPortUsage      *scheduling.HostPortUsage
	daemonResources    corev1.ResourceList
	hostname           string

	// We store the reserved offerings rather than appending reservation ID labels for two reasons:
	// - We need to release any reservations that were made in previous iterations and are no longer compatible with the
	//   NodeClaim.
	// - Since other NodeClaims may have released reservations which are compatible with this NodeClaim since the last
	//   time a pod was scheduled, it's possible for the set of reserved offerings to expand as well as contract over
	//   multiple iterations. This has the benefit of maximizing the flexibility of an in-flight NodeClaim, maximizing
	//   the scheduler's binpacking efficiency. Tightening the NodeClaim's requirements before finalization would prevent
	//   this expansion.
	reservedOfferings    cloudprovider.Offerings
	reservedOfferingMode ReservedOfferingMode
}

// ReservedOfferingError indicates a NodeClaim couldn't be created or a pod couldn't be added to an exxisting NodeClaim
// due to
type ReservedOfferingError struct {
	error
}

func NewReservedOfferingError(err error) ReservedOfferingError {
	_ = "STUB: not implemented"
	return *new(ReservedOfferingError)
}

func IsReservedOfferingError(err error) bool { _ = "STUB: not implemented"; return false }

func (e ReservedOfferingError) Unwrap() error { _ = "STUB: not implemented"; return nil }

var nodeID int64

func NewNodeClaim(
	nodeClaimTemplate *NodeClaimTemplate,
	topology *Topology,
	daemonResources corev1.ResourceList,
	hostPortUsage *scheduling.HostPortUsage,
	instanceTypes []*cloudprovider.InstanceType,
	reservationManager *ReservationManager,
	reservedOfferingMode ReservedOfferingMode,
) *NodeClaim {
	_ = "STUB: not implemented"
	return nil
}

// CanAdd returns whether the pod can be added to the NodeClaim
// based on the taints/tolerations, host port compatibility,
// requirements, resources, reserved capacity reservations, and topology requirements
func (n *NodeClaim) CanAdd(ctx context.Context, pod *corev1.Pod, podData *PodData, relaxMinValues bool) (updatedRequirements scheduling.Requirements, updatedInstanceTypes []*cloudprovider.InstanceType, offeringsToReserve []*cloudprovider.Offering, err error) {
	_ = "STUB: not implemented"
	// Check Taints
	return *new(scheduling.Requirements), nil, nil, nil
}

// exposed host ports on the node

// Check NodeClaim Affinity Requirements

// Build the list of volume requirement alternatives to try.
// Each alternative represents one valid combination of topology requirements for the pod's volumes.
// If there are no volume requirements, use a single nil entry with no additional topology constraint.

// Try each volume topology alternative. We need to iterate here because the selected
// volume topology constraints affect downstream topology checks (e.g., pod anti-affinity).

// tryVolumeAlternative attempts to add a pod with a specific set of volume requirements,
// checking topology, instance types, and offerings compatibility.
func (n *NodeClaim) tryVolumeAlternative(ctx context.Context, pod *corev1.Pod, podData *PodData, baseRequirements scheduling.Requirements, volReqs scheduling.Requirements, relaxMinValues bool) (scheduling.Requirements, []*cloudprovider.InstanceType, []*cloudprovider.Offering, error) {
	_ = "STUB: not implemented"
	return *new(scheduling.Requirements), nil, nil, nil
}

// Add volume requirements to nodeClaimRequirements ONLY (not to pod's affinity).
// This ensures the NodeClaim satisfies the selected volume topology constraints,
// while TSC counting uses pod's original affinity.

// Check Topology Requirements
// NOTE: podData.StrictRequirements does NOT include volume requirements,
// ensuring TSC counting uses pod's original affinity.

// Check instance type combinations

// Update min values on the requirements if they are relaxed

// We avoid wrapping this err because calling String() on InstanceTypeFilterError is an expensive operation
// due to calls to resources.Merge and stringifying the nodeClaimRequirements

// Add updates the NodeClaim to schedule the pod to this NodeClaim, updating
// the NodeClaim with new requirements, instance types, and offerings to reserve
// based on the pod scheduling
func (n *NodeClaim) Add(pod *corev1.Pod, podData *PodData, nodeClaimRequirements scheduling.Requirements, instanceTypes []*cloudprovider.InstanceType, offeringsToReserve []*cloudprovider.Offering) {
	_ = "STUB: not implemented"
	// Update node
	return
}

// releaseReservedOfferings releases all offerings which are present in the current reserved offerings, but are not
// present in the updated reserved offerings.
func (n *NodeClaim) releaseReservedOfferings(current, updated cloudprovider.Offerings) {
	_ = "STUB: not implemented"
	return
}

// reserveOfferings handles the reservation of `karpenter.sh/capacity-type: reserved` offerings, returning the set of
// reserved offerings. If the ReservedOfferingMode is set to strict, this function may also return an error if it failed
// to reserve compatible offerings when some were available.
//
//nolint:gocyclo
func (n *NodeClaim) offeringsToReserve(
	ctx context.Context,
	instanceTypes []*cloudprovider.InstanceType,
	nodeClaimRequirements scheduling.Requirements,
) (cloudprovider.Offerings, error) {
	_ = "STUB: not implemented"
	return *new(cloudprovider.Offerings), nil
}

// Track every incompatible reserved offering for release. Since releasing a reservation is a no-op when there is no
// reservation for the given host, there's no need to check that a reservation actually exists for the offering.

// Note that reservation is an idempotent operation - if we have previously successfully reserved an offering for
// this host, this operation is guaranteed to succeed. We may also succeed to make reservations for offerings which
// failed in previous iterations if other NodeClaims have released them since the last attempt.

// If an instance type with a compatible reserved offering exists, but we failed to make any reservations, we should
// fail. This could occur when all of the capacity for compatible instances has been reserved by previously created
// nodeclaims. Since we reserve offering pessimistically, i.e. we will reserve any offering that the instance could
// be launched with, we should fall back and attempt to schedule this pod in a subsequent scheduling simulation once
// reservation capacity is available again.

// If the nodeclaim previously had compatible reserved offerings, but the additional requirements filtered those out,
// we should fail to add the pod to this nodeclaim.

// FinalizeScheduling is called once all scheduling has completed and allows the node to perform any cleanup
// necessary before its requirements are used for instance launching
func (n *NodeClaim) FinalizeScheduling() {
	_ = "STUB: not implemented"
	// We need nodes to have hostnames for topology purposes, but we don't want to pass that node name on to consumers
	// of the node as it will be displayed in error messages
	return
}

// If there are any reserved offerings tracked, inject those requirements onto the NodeClaim. This ensures that if
// there are multiple reserved offerings for an instance type, we don't attempt to overlaunch into a single offering.

// Tightening constraint to reserved ensures that we get automatic drift handling when the Node / NodeClaim's capacity
// type label is dynamically updated by the cloudprovider.

func (n *NodeClaim) RemoveInstanceTypeOptionsByPriceAndMinValues(reqs scheduling.Requirements, maxPrice float64) (*NodeClaim, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func InstanceTypeList(instanceTypeOptions []*cloudprovider.InstanceType) string {
	_ = "STUB: not implemented"
	return ""
}

// print the first 5 instance types only (indices 0-4)

type InstanceTypeFilterError struct {
	// Each of these three flags indicates if that particular criteria was met by at least one instance type
	requirementsMet bool
	fits            bool
	hasOffering     bool

	// requirementsAndFits indicates if a single instance type met the scheduling requirements and had enough resources
	requirementsAndFits bool
	// requirementsAndOffering indicates if a single instance type met the scheduling requirements and was a required offering
	requirementsAndOffering bool
	// fitsAndOffering indicates if a single instance type had enough resources and was a required offering
	fitsAndOffering          bool
	minValuesIncompatibleErr error

	// We capture requirements so that we can know what the requirements were when evaluating instance type compatibility
	requirements scheduling.Requirements
	// We capture podRequests here since when a pod can't schedule due to requests, it's because the pod
	// was on its own on the simulated Node and exceeded the available resources for any instance type for this NodePool
	podRequests corev1.ResourceList
	// We capture daemonRequests since this contributes to the resources that are required to schedule to this NodePool
	daemonRequests corev1.ResourceList
}

//nolint:gocyclo
func (e InstanceTypeFilterError) Error() string {
	_ = "STUB: not implemented"
	// minValues is specified in the requirements and is not met
	return ""
}

// no instance type met any of the three criteria, meaning each criteria was enough to completely prevent
// this pod from scheduling

// check the other pairwise criteria

// and then each individual criteria. These are sort of the same as above in that each one indicates that no
// instance type matched that criteria at all, so it was enough to exclude all instance types.  I think it's
// helpful to have these separate, since we can report the multiple excluding criteria above.

// special case for a user typo I saw reported once

// see if any pair of criteria was enough to exclude all instances

// finally all instances were filtered out, but we had at least one instance that met each criteria, and met each
// pairwise set of criteria, so the only thing that remains is no instance which met all three criteria simultaneously

//nolint:gocyclo
func filterInstanceTypesByRequirements(instanceTypes []*cloudprovider.InstanceType, requirements scheduling.Requirements, podRequests, daemonRequests, totalRequests corev1.ResourceList, relaxMinValues bool) (cloudprovider.InstanceTypes, map[string]int, error) {
	_ = "STUB: not implemented"
	return *new(cloudprovider.InstanceTypes), nil, nil
}

// We hold the results of our scheduling simulation inside of this InstanceTypeFilterError struct
// to reduce the CPU load of having to generate the error string for a failed scheduling simulation

// the tradeoff to not short-circuiting on the filtering is that we can report much better error messages
// about why scheduling failed

// By using this iterative approach vs. the Available() function it prevents allocations
// which have to be garbage collected and slow down Karpenter's scheduling algorithm

// track if any single instance type met a single criteria

// track if any single instance type met the three pairs of criteria

// and if it met all criteria, we keep the instance type and continue filtering.  We now won't be reporting
// any errors.

// We don't care about the minimum number of instance types that meet our requirements here, we only care if they meet our requirements.

// If MinValuesPolicy is set to Strict, return empty InstanceTypeOptions as we cannot launch with the remaining InstanceTypes when min values is violated.

func compatible(instanceType *cloudprovider.InstanceType, requirements scheduling.Requirements) bool {
	_ = "STUB: not implemented"
	return false
}

func fits(instanceType *cloudprovider.InstanceType, requests corev1.ResourceList) bool {
	_ = "STUB: not implemented"
	return false
}
