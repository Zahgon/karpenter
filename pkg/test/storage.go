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

package test

import (
	v1 "k8s.io/api/core/v1"
	storagev1 "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PersistentVolumeOptions struct {
	metav1.ObjectMeta
	// Zones builds a single topology term with multiple zone values unless NodeSelectorTerms is set.
	Zones []string
	// NodeSelectorTerms, when set, are used verbatim instead of deriving a single term from Zones.
	NodeSelectorTerms  []v1.NodeSelectorTerm
	StorageClassName   string
	Driver             string
	UseAWSInTreeDriver bool
	UseLocal           bool
	UseHostPath        bool
}

func PersistentVolume(overrides ...PersistentVolumeOptions) *v1.PersistentVolume {
	_ = "STUB: not implemented"
	return nil
}

// Determine the PersistentVolumeSource based on the options

type PersistentVolumeClaimOptions struct {
	metav1.ObjectMeta
	StorageClassName *string
	VolumeName       string
	Resources        v1.VolumeResourceRequirements
}

func PersistentVolumeClaim(overrides ...PersistentVolumeClaimOptions) *v1.PersistentVolumeClaim {
	_ = "STUB: not implemented"
	return nil
}

type StorageClassOptions struct {
	metav1.ObjectMeta
	// Zones builds a single allowed topology term with multiple zone values unless AllowedTopologies is set.
	Zones []string
	// AllowedTopologies, when set, are used verbatim instead of deriving a single term from Zones.
	AllowedTopologies []v1.TopologySelectorTerm
	Provisioner       *string
	VolumeBindingMode *storagev1.VolumeBindingMode
}

func StorageClass(overrides ...StorageClassOptions) *storagev1.StorageClass {
	_ = "STUB: not implemented"
	return nil
}

type VolumeAttachmentOptions struct {
	metav1.ObjectMeta
	NodeName   string
	VolumeName string
}

func VolumeAttachment(overrides ...VolumeAttachmentOptions) *storagev1.VolumeAttachment {
	_ = "STUB: not implemented"
	return nil
}
