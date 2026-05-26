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

package controller

import (
	"context"
)

// CPUCount calculates CPU count (in cores) from context options (in millicores)
func CPUCount(ctx context.Context) float64 { _ = "STUB: not implemented"; return 0 }

// LinearScaleReconciles calculates maxConcurrentReconciles using linear scaling
func LinearScaleReconciles(cpuCount float64, minReconciles int, maxReconciles int) int {
	_ = "STUB: not implemented"
	// At 1 core: minReconciles; At 60 cores: maxReconciles
	return 0
}

// Clamp to ensure we stay within bounds

func GetTypedBucketConfigs(minQPS int, minReconciles int, concurrentReconciles int) (int, int) {
	_ = "STUB: not implemented"
	return 0, 0
}
