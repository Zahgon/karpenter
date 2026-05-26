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

package operator

import (
	"context"
	"runtime"

	"github.com/awslabs/operatorpkg/controller"
	opmetrics "github.com/awslabs/operatorpkg/metrics"
	"github.com/awslabs/operatorpkg/option"
	"github.com/prometheus/client_golang/prometheus"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/utils/clock"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	crmetrics "sigs.k8s.io/controller-runtime/pkg/metrics"

	"sigs.k8s.io/karpenter/pkg/controllers/nodeoverlay"
	"sigs.k8s.io/karpenter/pkg/events"
	"sigs.k8s.io/karpenter/pkg/metrics"
	"sigs.k8s.io/karpenter/pkg/utils/env"
)

var AppName = "karpenter"

var (
	BuildInfo = opmetrics.NewPrometheusGauge(
		crmetrics.Registry,
		prometheus.GaugeOpts{
			Namespace: metrics.Namespace,
			Name:      "build_info",
			Help:      "A metric with a constant '1' value labeled by version from which karpenter was built.",
		},
		[]string{"version", "goversion", "goarch", "commit"},
	)
)

// Version is the karpenter app version injected during compilation
// when using the Makefile
var Version = "unspecified"

func init() {
	opmetrics.RegisterClientMetrics(crmetrics.Registry)

	BuildInfo.Set(1, map[string]string{
		"version":   Version,
		"goversion": runtime.Version(),
		"goarch":    runtime.GOARCH,
		"commit":    env.GetRevision(),
	})
}

type Operator struct {
	manager.Manager

	KubernetesInterface kubernetes.Interface
	EventRecorder       events.Recorder
	Clock               clock.Clock
	InstanceTypeStore   *nodeoverlay.InstanceTypeStore
}

type Options struct {
	LeaderElectionLabels map[string]string
	LeaderElectionConfig *rest.Config // Optional separate config for leader election
}

// Adds LeaderElectionLabels to the underlying manager's LeaderElectionOptions
func WithLeaderElectionLabels(labels map[string]string) option.Function[Options] {
	_ = "STUB: not implemented"
	return nil
}

// Adds LeaderElectionConfig to the underlying manager's LeaderElectionOptions allowing for custom client config
func WithLeaderElectionConfig(config *rest.Config) option.Function[Options] {
	_ = "STUB: not implemented"
	return nil
}

// NewOperator instantiates a controller manager or panics
func NewOperator(o ...option.Function[Options]) (context.Context, *Operator) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

// Root Context

// Options

// Make the Karpenter binary aware of the container memory limit
// https://pkg.go.dev/runtime/debug#SetMemoryLimit

// Logging

// Client Config

// Copy the leader config for lower QPS/Burst
// We changed this from explicitly setting the RateLimiter on the config and not creating
// a separate leaderConfig ourselves because this caused a subtle bug when copying the leaderConfig
// for the leader election client. The leaderConfig would use the same RateLimiter, so client-side rate
// limiting on the regular config would also cause client-side rate limiting on the leader election client,
// often leading to leader loss during large scale-ups or periods of high churn

// Client

// Manager

// EnableWarmup allows controllers to start their sources (watches/informers) before leader election
// is won. This pre-populates caches and improves leader failover time. Only effective when leader
// election is enabled, so we only set it when both conditions are true.

// TODO @joinnis: Investigate the mgrOpts.PprofBindAddress that would allow native support for pprof
// On initial look, it seems like this native pprof doesn't support some of the routes that we have here
// like "/debug/pprof/heap" or "/debug/pprof/block"

func (o *Operator) WithControllers(ctx context.Context, controllers ...controller.Controller) *Operator {
	_ = "STUB: not implemented"
	return nil
}

func (o *Operator) Start(ctx context.Context) { _ = "STUB: not implemented"; return }

func setupIndexers(ctx context.Context, mgr manager.Manager) { _ = "STUB: not implemented"; return }

// If the CRD does not exist, we should fail open when setting up indexers. This ensures controllers that aren't reliant on those CRDs may continue to function

// lo.Must0 also does a panic
