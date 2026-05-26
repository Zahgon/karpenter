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

package metrics

import (
	opmetrics "github.com/awslabs/operatorpkg/metrics"
)

const (
	// Common namespace for application metrics.
	Namespace = "karpenter"

	NodePoolLabel         = "nodepool"
	ReasonLabel           = "reason"
	ResourceTypeLabel     = "resource_type"
	CapacityTypeLabel     = "capacity_type"
	MinValuesRelaxedLabel = "min_values_relaxed"

	// Reasons for CREATE/DELETE shared metrics
	ProvisionedReason = "provisioned"
	ExpiredReason     = "expired"
	UnhealthyReason   = "unhealthy"
)

// DurationBuckets returns a []float64 of default threshold values for duration histograms.
// Each returned slice is new and may be modified without impacting other bucket definitions.
func DurationBuckets() []float64 {
	_ = "STUB: not implemented"
	// Use same bucket thresholds as controller-runtime.
	// https://github.com/kubernetes-sigs/controller-runtime/blob/v0.10.0/pkg/internal/controller/metrics/metrics.go#L47-L48
	// Add in values larger than 60 for singleton controllers that do not have a timeout.
	return nil
}

// Returns a map of summary objectives (quantile-error pairs)
func SummaryObjectives() map[float64]float64 { _ = "STUB: not implemented"; return nil }

// Measure returns a deferrable function that observes the duration between the
// defer statement and the end of the function.
func Measure(observer opmetrics.ObservationMetric, labels map[string]string) func() {
	_ = "STUB: not implemented"
	return nil
}
