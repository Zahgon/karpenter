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
	"math/rand"

	v1 "k8s.io/api/core/v1"
	policyv1 "k8s.io/api/policy/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// PodOptions customizes a Pod.
type PodOptions struct {
	metav1.ObjectMeta
	Image                         string
	NodeName                      string
	Overhead                      v1.ResourceList
	PriorityClassName             string
	InitContainers                []v1.Container
	PodResourceRequirements       v1.ResourceRequirements
	ResourceRequirements          v1.ResourceRequirements
	NodeSelector                  map[string]string
	NodeRequirements              []v1.NodeSelectorRequirement
	NodePreferences               []v1.NodeSelectorRequirement
	PodRequirements               []v1.PodAffinityTerm
	PodPreferences                []v1.WeightedPodAffinityTerm
	PodAntiRequirements           []v1.PodAffinityTerm
	PodAntiPreferences            []v1.WeightedPodAffinityTerm
	TopologySpreadConstraints     []v1.TopologySpreadConstraint
	Tolerations                   []v1.Toleration
	PersistentVolumeClaims        []string
	EphemeralVolumeTemplates      []EphemeralVolumeTemplateOptions
	HostPorts                     []int32
	Conditions                    []v1.PodCondition
	Phase                         v1.PodPhase
	RestartPolicy                 v1.RestartPolicy
	TerminationGracePeriodSeconds *int64
	ReadinessProbe                *v1.Probe
	LivenessProbe                 *v1.Probe
	PreStopSleep                  *int64
	Command                       []string
	ResourceClaims                []v1.PodResourceClaim
	ContainerResourceClaims       []v1.ResourceClaim
	InitContainerResourceClaims   []v1.ResourceClaim
}

type PDBOptions struct {
	metav1.ObjectMeta
	Labels                     map[string]string
	MinAvailable               *intstr.IntOrString
	MaxUnavailable             *intstr.IntOrString
	UnhealthyPodEvictionPolicy *policyv1.UnhealthyPodEvictionPolicyType
	Status                     *policyv1.PodDisruptionBudgetStatus
}

type EphemeralVolumeTemplateOptions struct {
	StorageClassName *string
}

var (
	DefaultImage        = "public.ecr.aws/eks-distro/kubernetes/pause:3.2"
	KWOKDelayAnnotation = "pod-delete.stage.kwok.x-k8s.io/delay"
)

// Pod creates a test pod with defaults that can be overridden by PodOptions.
// Overrides are applied in order, with a last write wins semantic.
// nolint:gocyclo
func Pod(overrides ...PodOptions) *v1.Pod { _ = "STUB: not implemented"; return nil }

// If PreStopSleep is enabled, add it to each of the containers.
// Can't use v1.LifecycleHandler == v1.SleepAction as that's a feature gate in Alpha 1.29.
// https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#hook-handler-implementations

// Add init container resource claims if specified

func buildResourceRequirements(options PodOptions) v1.ResourceRequirements {
	_ = "STUB: not implemented"
	return *new(v1.ResourceRequirements)
}

// Pods creates homogeneous groups of pods based on the passed in options, evenly divided by the total pods requested
func Pods(total int, options ...PodOptions) []*v1.Pod { _ = "STUB: not implemented"; return nil }

func UnscheduleablePodOptions(overrides ...PodOptions) PodOptions {
	_ = "STUB: not implemented"
	return *new(PodOptions)
}

// UnschedulablePod creates a test pod with a pending scheduling status condition
func UnschedulablePod(options ...PodOptions) *v1.Pod { _ = "STUB: not implemented"; return nil }

// UnschedulablePods returns slice of configurable length of identical test pods with a pending scheduling status condition
func UnschedulablePods(options PodOptions, num int) []*v1.Pod {
	_ = "STUB: not implemented"
	return nil
}

// PodDisruptionBudget creates a PodDisruptionBudget.  To function properly, it should have its status applied
func PodDisruptionBudget(overrides ...PDBOptions) *policyv1.PodDisruptionBudget {
	_ = "STUB: not implemented"
	return nil
}

// To be considered for application by eviction, the Status.ObservedGeneration must be >= the PDB generation.
// kube-controller-manager normally sets ObservedGeneration, but we don't have one when running under
// EnvTest. If this isn't modified the eviction controller assumes that the PDB hasn't been processed
// by the disruption controller yet and adds a 10 second retry to our evict() call

func buildAffinity(options PodOptions) *v1.Affinity { _ = "STUB: not implemented"; return nil }

func buildPodAffinity(podRequirements []v1.PodAffinityTerm, podPreferences []v1.WeightedPodAffinityTerm) *v1.PodAffinity {
	_ = "STUB: not implemented"
	return nil
}

func buildPodAntiAffinity(podAntiRequirements []v1.PodAffinityTerm, podAntiPreferences []v1.WeightedPodAffinityTerm) *v1.PodAntiAffinity {
	_ = "STUB: not implemented"
	return nil
}

func buildNodeAffinity(nodeRequirements []v1.NodeSelectorRequirement, nodePreferences []v1.NodeSelectorRequirement) *v1.NodeAffinity {
	_ = "STUB: not implemented"
	return nil
}

func MakePodAntiAffinityPodOptions(key string) PodOptions {
	_ = "STUB: not implemented"
	// all of these pods have anti-affinity to each other
	return *new(PodOptions)
}

func MakePodAffinityPodOptions(key string) PodOptions {
	_ = "STUB: not implemented"
	return *new(PodOptions)
}

func MakeTopologySpreadPodOptions(key string) PodOptions {
	_ = "STUB: not implemented"
	return *new(PodOptions)
}

func MakeGenericPodOptions() PodOptions { _ = "STUB: not implemented"; return *new(PodOptions) }

func MakeDiversePodOptions() []PodOptions { _ = "STUB: not implemented"; return nil }

func MakeDRAPodOptions(claimName string) PodOptions {
	_ = "STUB: not implemented"
	return *new(PodOptions)
}

func RandomAffinityLabels() map[string]string { _ = "STUB: not implemented"; return nil }

func RandomLabels() map[string]string { _ = "STUB: not implemented"; return nil }

//nolint:gosec
var r = rand.New(rand.NewSource(42))

func RandomLabelValue() string { _ = "STUB: not implemented"; return "" }

func RandomMemory() resource.Quantity { _ = "STUB: not implemented"; return *new(resource.Quantity) }

func RandomCPU() resource.Quantity { _ = "STUB: not implemented"; return *new(resource.Quantity) }
