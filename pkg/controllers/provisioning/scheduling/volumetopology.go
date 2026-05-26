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

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"sigs.k8s.io/karpenter/pkg/scheduling"
)

// UnsupportedProvisioners is a map of volume plugins that are not supported. When a pod requests storage using a PersistentVolumeClaim (PVC)
// that uses a StorageClass with any of these unsupported provisioners, Karpenter will skip scheduling that pod.
var UnsupportedProvisioners = sets.New[string]()

// UnsupportedTopologyKeys is a set of topology keys that are not supported. When a StorageClass has AllowedTopologies
// containing any of these keys, Karpenter will skip scheduling pods that reference a PVC with that StorageClass, since nodes
// created by Karpenter will never satisfy the topology requirement.
var UnsupportedTopologyKeys = sets.New[string]()

func NewVolumeTopology(kubeClient client.Client) *VolumeTopology {
	_ = "STUB: not implemented"
	return nil
}

type VolumeTopology struct {
	kubeClient client.Client
}

// GetRequirements returns the volume topology requirements for the pod as a list of alternatives.
// Each alternative is a scheduling.Requirements representing one valid combination of volume topology constraints.
// When a volume has multiple allowed topology terms (OR'd NodeSelectorTerms or AllowedTopologies),
// each term becomes a separate alternative. For pods with multiple volumes, the cross product of all
// per-volume alternatives is computed.
//
// These requirements should be:
//   - Added to nodeRequirements (for NodeClaim topology filtering)
//   - NOT added to pod's NodeAffinity (to preserve correct TSC counting)
func (v *VolumeTopology) GetRequirements(ctx context.Context, pod *v1.Pod) ([]scheduling.Requirements, error) {
	_ = "STUB: not implemented"
	// Start with a single empty alternative (matches everything)
	return nil, nil
}

// If we still have just the initial empty alternative, there are no volume requirements

func mergeVolumeRequirementAlternatives(alternatives, volAlts []scheduling.Requirements) []scheduling.Requirements {
	_ = "STUB: not implemented"
	return nil
}

// Prefer only compatible cross-product branches, but preserve the old merged result when every
// branch is incompatible. Treating the all-pruned case as unschedulable requires separate
// provisioning metrics and scheduling-decision handling.

func mergeCompatibleVolumeRequirementAlternatives(alternatives, volAlts []scheduling.Requirements) []scheduling.Requirements {
	_ = "STUB: not implemented"
	return nil
}

func mergeAllVolumeRequirementAlternatives(alternatives, volAlts []scheduling.Requirements) []scheduling.Requirements {
	_ = "STUB: not implemented"
	return nil
}

func volumeRequirementsCompatible(existing, volReq scheduling.Requirements) bool {
	_ = "STUB: not implemented"
	return false
}

func mergeVolumeRequirements(existing, volReq scheduling.Requirements) scheduling.Requirements {
	_ = "STUB: not implemented"
	return *new(scheduling.Requirements)
}

func (v *VolumeTopology) getRequirements(ctx context.Context, pod *v1.Pod, volume v1.Volume) ([]scheduling.Requirements, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Not all volume types have PVCs, e.g. emptyDir, hostPath, etc.

// Persistent Volume Requirements

// Storage Class Requirements

func (v *VolumeTopology) getStorageClassRequirements(ctx context.Context, storageClassName string) ([]scheduling.Requirements, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Each TopologySelectorTerm is OR'd — each becomes a separate alternative

func (v *VolumeTopology) getPersistentVolumeRequirements(ctx context.Context, pod *v1.Pod, volumeName string) ([]scheduling.Requirements, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Each NodeSelectorTerm is OR'd — each becomes a separate alternative

// If we are using a Local volume or a HostPath volume, then we should ignore the Hostname affinity
// on it because re-scheduling this pod to a new node means not using the same Hostname affinity that we currently have

// Preserve hostname-only terms as unconstrained alternatives, since hostname affinity
// is intentionally ignored for Local and HostPath volumes.

// ValidatePersistentVolumeClaims returns an error if the pod doesn't appear to be valid with respect to
// PVCs (e.g. the PVC is not found or references an unknown storage class).
// nolint:gocyclo
func (v *VolumeTopology) ValidatePersistentVolumeClaims(ctx context.Context, pod *v1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

// Not all volume types have PVCs, e.g. emptyDir, hostPath, etc.

// Handle cases specifically rejected by kube-scheduler
// https://github.com/kubernetes/kubernetes/blob/56f6358c11b78e8e3d39e8cd8ff016ff7c70c56b/pkg/scheduler/framework/plugins/volumebinding/volume_binding.go#L333

// kube-scheduler treats PVCs that have a volumeName as Immediate volumes
// Any PVC that does not contain the "pv.kubernetes.io/bind-completed" annotation is not considered bound
// https://github.com/kubernetes/kubernetes/blob/ecf2c52f756461cfb7ffd5469975ecd635e5feeb/pkg/scheduler/framework/plugins/volumebinding/binder.go#L770

// PVC is unbound, we can't schedule unless the pod defines a valid storage class

// Ignore pods than have unbound pvc for volumeBindingMode immediate

// Reject pods whose StorageClass has AllowedTopologies with unsupported topology keys,
// since Karpenter-created nodes will never have matching labels for these keys.

// Finally, validate that the driver is in the set of supported drivers

func (v *VolumeTopology) validateVolume(ctx context.Context, volumeName string) error {
	_ = "STUB: not implemented"
	// we have a volume name, so ensure that it exists
	return nil
}
