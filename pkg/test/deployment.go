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
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type DeploymentOptions struct {
	metav1.ObjectMeta
	Labels     map[string]string
	Replicas   int32
	PodOptions PodOptions
}

func Deployment(overrides ...DeploymentOptions) *appsv1.Deployment {
	_ = "STUB: not implemented"
	return nil
}

// DeploymentOptionModifier is a function that modifies DeploymentOptions
type DeploymentOptionModifier func(*DeploymentOptions)

// CreateDeploymentOptions creates a test.DeploymentOptions with the specified parameters
// and applies any provided modifiers. This provides a clean, extensible way to create
// deployment configurations for tests.
//
// Example usage:
//
//	// Simple deployment
//	opts := CreateDeploymentOptions("my-app", 10, "100m", "128Mi")
//
//	// With modifiers
//	opts := CreateDeploymentOptions("my-app", 10, "100m", "128Mi",
//	    WithHostnameSpread(),
//	    WithDoNotDisrupt(),
//	    WithLabels(map[string]string{"tier": "frontend"}))
//
//	// Use with test.Deployment
//	deployment := test.Deployment(opts)
func CreateDeploymentOptions(name string, replicas int32, cpuRequest, memoryRequest string, opts ...DeploymentOptionModifier) DeploymentOptions {
	_ = "STUB: not implemented"
	// Create base deployment options
	return *new(DeploymentOptions)
}

// Apply all modifiers

// WithLabels adds or merges labels to the pod template
func WithLabels(labels map[string]string) DeploymentOptionModifier {
	_ = "STUB: not implemented"
	return *new(DeploymentOptionModifier)
}

// WithAnnotations adds or merges annotations to the pod template
func WithAnnotations(annotations map[string]string) DeploymentOptionModifier {
	_ = "STUB: not implemented"
	return *new(DeploymentOptionModifier)
}

// WithDoNotDisrupt adds the do-not-disrupt annotation to prevent Karpenter disruption
func WithDoNotDisrupt() DeploymentOptionModifier {
	_ = "STUB: not implemented"
	return *new(DeploymentOptionModifier)
}

// WithHostnameSpread adds topology spread constraints to spread pods across hostnames
func WithHostnameSpread() DeploymentOptionModifier {
	_ = "STUB: not implemented"
	return *new(DeploymentOptionModifier)
}

// Get the current labels to use in the label selector

// WithZoneSpread adds topology spread constraints to spread pods across zones
func WithZoneSpread() DeploymentOptionModifier {
	_ = "STUB: not implemented"
	return *new(DeploymentOptionModifier)
}

// Get the current labels to use in the label selector

// WithPodAntiAffinity adds pod anti-affinity to prevent pods from being scheduled on the same topology
func WithPodAntiAffinity(topologyKey string) DeploymentOptionModifier {
	_ = "STUB: not implemented"
	return *new(DeploymentOptionModifier)
}

// Get the current labels to use in the label selector

// WithTopologySpreadConstraints adds custom topology spread constraints
func WithTopologySpreadConstraints(constraints []v1.TopologySpreadConstraint) DeploymentOptionModifier {
	_ = "STUB: not implemented"
	return *new(DeploymentOptionModifier)
}

// WithResourceLimits adds resource limits to the pod containers
func WithResourceLimits(cpuLimit, memoryLimit string) DeploymentOptionModifier {
	_ = "STUB: not implemented"
	return *new(DeploymentOptionModifier)
}

// WithPriorityClass sets the priority class for the pods
func WithPriorityClass(priorityClassName string) DeploymentOptionModifier {
	_ = "STUB: not implemented"
	return *new(DeploymentOptionModifier)
}

// WithNodeSelector adds node selector requirements
func WithNodeSelector(nodeSelector map[string]string) DeploymentOptionModifier {
	_ = "STUB: not implemented"
	return *new(DeploymentOptionModifier)
}

// WithTolerations adds tolerations to the pod spec
func WithTolerations(tolerations []v1.Toleration) DeploymentOptionModifier {
	_ = "STUB: not implemented"
	return *new(DeploymentOptionModifier)
}

// WithTerminationGracePeriod sets the termination grace period for pods
func WithTerminationGracePeriod(seconds int64) DeploymentOptionModifier {
	_ = "STUB: not implemented"
	return *new(DeploymentOptionModifier)
}

// WithNoResourceRequests removes resource requirements entirely (for backward compatibility)
func WithNoResourceRequests() DeploymentOptionModifier {
	_ = "STUB: not implemented"
	return *new(DeploymentOptionModifier)
}

// WithImage sets a custom container image
func WithImage(image string) DeploymentOptionModifier {
	_ = "STUB: not implemented"
	return *new(DeploymentOptionModifier)
}

// WithRestartPolicy sets the pod restart policy
func WithRestartPolicy(policy v1.RestartPolicy) DeploymentOptionModifier {
	_ = "STUB: not implemented"
	return *new(DeploymentOptionModifier)
}

// WithReadinessProbe adds a readiness probe to the pod
func WithReadinessProbe(probe *v1.Probe) DeploymentOptionModifier {
	_ = "STUB: not implemented"
	return *new(DeploymentOptionModifier)
}

// WithLivenessProbe adds a liveness probe to the pod
func WithLivenessProbe(probe *v1.Probe) DeploymentOptionModifier {
	_ = "STUB: not implemented"
	return *new(DeploymentOptionModifier)
}

// WithPodAntiAffinityHostname is a convenience function for hostname anti-affinity
func WithPodAntiAffinityHostname() DeploymentOptionModifier {
	_ = "STUB: not implemented"
	return *new(DeploymentOptionModifier)
}

// WithCommand sets a custom command for the container
func WithCommand(command []string) DeploymentOptionModifier {
	_ = "STUB: not implemented"
	return *new(DeploymentOptionModifier)
}

// WithPreStopSleep adds a pre-stop sleep hook to the container
func WithPreStopSleep(seconds int64) DeploymentOptionModifier {
	_ = "STUB: not implemented"
	return *new(DeploymentOptionModifier)
}
