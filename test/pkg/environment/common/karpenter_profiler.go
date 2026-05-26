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

package common

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
)

const CPUProfileSeconds = 20

// KarpenterProfiler polls pprof for Karpenter memory and CPU usage and captures profiles at peak
type KarpenterProfiler struct {
	env               *Environment
	peakMemoryMB      float64
	peakMemoryProfile []byte
	peakCPUNanos      int64
	peakCPUProfile    []byte
	cancel            context.CancelFunc
	done              chan struct{}
	pollCount         int
	lastError         string
}

// StartKarpenterProfiler begins profiling Karpenter resource usage in the background
func StartKarpenterProfiler(env *Environment) *KarpenterProfiler {
	_ = "STUB: not implemented"
	return nil
}

// Stop stops the profiler and returns peak memory (MB), memory profile, peak CPU (nanoseconds), and CPU profile
func (kp *KarpenterProfiler) Stop() (float64, []byte, int64, []byte) {
	_ = "STUB: not implemented"
	return 0, nil, 0, nil
}

func (kp *KarpenterProfiler) run(ctx context.Context) { _ = "STUB: not implemented"; return }

func (kp *KarpenterProfiler) captureProfiles(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

// Capture heap profile (instant)

// Capture CPU profile (20 second sample)

func (kp *KarpenterProfiler) fetchHeapProfile(port int) (float64, []byte) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (kp *KarpenterProfiler) fetchCPUProfile(port int) (int64, []byte) {
	_ = "STUB: not implemented"
	return 0, nil
}

func parseProfileValue(data []byte, sampleType string) int64 { _ = "STUB: not implemented"; return 0 }
