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
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"
	csitranslation "k8s.io/csi-translation-lib"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

//go:generate go tool -modfile=../../go.tools.mod controller-gen object:headerFile="../../hack/boilerplate.go.txt" paths="."

// translator is a CSI Translator that translates in-tree plugin names to their out-of-tree CSI driver names
var translator = csitranslation.New()

// +k8s:deepcopy-gen=true
type Volumes map[string]sets.Set[string]

func (u Volumes) Add(provisioner string, pvcID string) { _ = "STUB: not implemented"; return }

func (u Volumes) Union(vol Volumes) Volumes { _ = "STUB: not implemented"; return *new(Volumes) }

func (u Volumes) Insert(volumes Volumes) { _ = "STUB: not implemented"; return }

//nolint:gocyclo
func GetVolumes(ctx context.Context, kubeClient client.Client, pod *v1.Pod) (Volumes, error) {
	_ = "STUB: not implemented"
	return *new(Volumes), nil
}

// If the PVC is not found it was manually deleted and its finalizer removed. We should ignore this volume when
// computing limits, otherwise Karpenter may never be able to update its cluster state.

// Not all volume types have PVCs, e.g. emptyDir, hostPath, etc.

// might be a non-CSI driver, something we don't currently handle

// ResolveDriver resolves the storage driver name in the following order:
//  1. If the PV associated with the pod volume is using CSI.driver in its spec, then use that name
//  2. If the StorageClass associated with the PV has a Provisioner
func ResolveDriver(ctx context.Context, kubeClient client.Client, pod *v1.Pod, volumeName string, pvc *v1.PersistentVolumeClaim, storageClassName string) (string, error) {
	_ = "STUB: not implemented"
	// We can track the volume usage by the CSI Driver name which is pulled from the storage class for dynamic
	// volumes, or if it's bound/static we can pull the volume name
	return "", nil
}

// The PVC is bound, but not to a volume managed by a CSI driver or a known in-tree equivalent. This PVC can be ignored for the purposes of volume limit tracking.

// This can occur in two scenarios:
//  1. The storage class was explicitly set to "" to disable dynamic provisioning
//  2. The storage class was not set but the cluster doesn't have a default storage class
// In either of these cases, a PV must have been previously bound to the PVC and has since been removed. We can
// ignore this PVC while computing limits and continue.

// There are two scenarios where a StorageClass may be defined but not found:
//  1. The StorageClass was manually deleted and the finalizer removed
//  2. The StorageClass never existed and was used to bind the PVC to an existing PV, but that PV was removed
// In either of these cases, we should ignore the PVC while computing limits and continue.

// driverFromSC resolves the storage driver name by getting the Provisioner name from the StorageClass
func driverFromSC(ctx context.Context, kubeClient client.Client, storageClassName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Check if the provisioner name is an in-tree plugin name

// driverFromVolume resolves the storage driver name by getting the CSI spec from inside the PersistentVolume
func driverFromVolume(ctx context.Context, kubeClient client.Client, volumeName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// VolumeUsage tracks volume limits on a per node basis.  The number of volumes that can be mounted varies by instance
// type. We need to be aware and track the mounted volume usage to inform our awareness of which pods can schedule to
// which nodes.
// +k8s:deepcopy-gen=true
type VolumeUsage struct {
	volumes    Volumes
	podVolumes map[types.NamespacedName]Volumes
	limits     map[string]int
}

func NewVolumeUsage() *VolumeUsage { _ = "STUB: not implemented"; return nil }

func (v *VolumeUsage) ExceedsLimits(vols Volumes) error { _ = "STUB: not implemented"; return nil }

func (v *VolumeUsage) AddLimit(storageDriver string, value int) { _ = "STUB: not implemented"; return }

func (v *VolumeUsage) Add(pod *v1.Pod, volumes Volumes) { _ = "STUB: not implemented"; return }

func (v *VolumeUsage) DeletePod(key types.NamespacedName) { _ = "STUB: not implemented"; return }

// volume names could be duplicated, so we re-create our volumes
