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
	"time"

	"sigs.k8s.io/karpenter/pkg/operator/options"
)

type OptionsFields struct {
	// Vendor Neutral
	ServiceName                      *string
	MetricsPort                      *int
	HealthProbePort                  *int
	KubeClientQPS                    *int
	KubeClientBurst                  *int
	EnableProfiling                  *bool
	DisableControllerWarmup          *bool
	DisableLeaderElection            *bool
	DisableClusterStateObservability *bool
	LeaderElectionName               *string
	LeaderElectionNamespace          *string
	MemoryLimit                      *int64
	CPURequests                      *int64
	LogLevel                         *string
	LogOutputPaths                   *string
	LogErrorOutputPaths              *string
	PreferencePolicy                 *options.PreferencePolicy
	MinValuesPolicy                  *options.MinValuesPolicy
	BatchMaxDuration                 *time.Duration
	BatchIdleDuration                *time.Duration
	IgnoreDRARequests                *bool
	FeatureGates                     FeatureGates
}

type FeatureGates struct {
	NodeRepair              *bool
	ReservedCapacity        *bool
	SpotToSpotConsolidation *bool
	NodeOverlay             *bool
	StaticCapacity          *bool
}

func Options(overrides ...OptionsFields) *options.Options { _ = "STUB: not implemented"; return nil }

// use 5 threads to enforce parallelism
