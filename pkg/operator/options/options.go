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

package options

import (
	"context"
	"flag"
	"time"
)

type PreferencePolicy string

const (
	PreferencePolicyIgnore  PreferencePolicy = "Ignore"
	PreferencePolicyRespect PreferencePolicy = "Respect"
)

type MinValuesPolicy string

const (
	MinValuesPolicyStrict     MinValuesPolicy = "Strict"
	MinValuesPolicyBestEffort MinValuesPolicy = "BestEffort"
)

var (
	validLogLevels          = []string{"", "debug", "info", "error"}
	validPreferencePolicies = []PreferencePolicy{PreferencePolicyIgnore, PreferencePolicyRespect}

	Injectables = []Injectable{&Options{}}
)

type optionsKey struct{}

type FeatureGates struct {
	inputStr string

	NodeRepair              bool
	ReservedCapacity        bool
	SpotToSpotConsolidation bool
	NodeOverlay             bool
	StaticCapacity          bool
}

// Options contains all CLI flags / env vars for karpenter-core. It adheres to the options.Injectable interface.
type Options struct {
	ServiceName                      string
	MetricsPort                      int
	HealthProbePort                  int
	KubeClientQPS                    int
	KubeClientBurst                  int
	EnableProfiling                  bool
	DisableControllerWarmup          bool
	DisableLeaderElection            bool
	DisableClusterStateObservability bool
	LeaderElectionName               string
	LeaderElectionNamespace          string
	MemoryLimit                      int64
	CPURequests                      int64
	LogLevel                         string
	LogOutputPaths                   string
	LogErrorOutputPaths              string
	BatchMaxDuration                 time.Duration
	BatchIdleDuration                time.Duration
	preferencePolicyRaw              string
	PreferencePolicy                 PreferencePolicy
	minValuesPolicyRaw               string
	MinValuesPolicy                  MinValuesPolicy
	IgnoreDRARequests                bool // NOTE: This flag will be removed once formal DRA support is GA in Karpenter.
	FeatureGates                     FeatureGates
}

type FlagSet struct {
	*flag.FlagSet
}

// BoolVarWithEnv defines a bool flag with a specified name, default value, usage string, and fallback environment
// variable.
func (fs *FlagSet) BoolVarWithEnv(p *bool, name string, envVar string, val bool, usage string) {
	_ = "STUB: not implemented"
	return
}

func (o *Options) AddFlags(fs *FlagSet) { _ = "STUB: not implemented"; return }

func (o *Options) Parse(fs *FlagSet, args ...string) error { _ = "STUB: not implemented"; return nil }

func (o *Options) ToContext(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func DefaultFeatureGates() FeatureGates { _ = "STUB: not implemented"; return *new(FeatureGates) }

func ParseFeatureGates(gateStr string) (FeatureGates, error) {
	_ = "STUB: not implemented"
	return *new(FeatureGates), nil
}

// Parses feature gates with the upstream mechanism. This is meant to be used with flag directly but this enables
// simple merging with environment vars.

func ToContext(ctx context.Context, opts *Options) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func FromContext(ctx context.Context) *Options { _ = "STUB: not implemented"; return nil }

// This is a developer error if this happens, so we should panic
