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

package disruption

import (
	"context"

	v1 "sigs.k8s.io/karpenter/pkg/apis/v1"
	"sigs.k8s.io/karpenter/pkg/cloudprovider"
	"sigs.k8s.io/karpenter/pkg/controllers/provisioning"
	"sigs.k8s.io/karpenter/pkg/controllers/state"
)

// StaticDrift is a subreconciler that deletes drifted static candidates.
type StaticDrift struct {
	cluster       *state.Cluster
	provisioner   *provisioning.Provisioner
	cloudprovider cloudprovider.CloudProvider
}

func NewStaticDrift(cluster *state.Cluster, provisioner *provisioning.Provisioner, cloudprovider cloudprovider.CloudProvider) *StaticDrift {
	_ = "STUB: not implemented"
	return nil
}

// ShouldDisrupt is a predicate used to filter candidates
func (d *StaticDrift) ShouldDisrupt(_ context.Context, c *Candidate) bool {
	_ = "STUB: not implemented"
	return false
}

func (d *StaticDrift) ComputeCommands(ctx context.Context, disruptionBudgetMapping map[string]int, candidates ...*Candidate) ([]Command, error) {
	_ = "STUB: not implemented"
	// Group candidates by nodepool name
	return nil, nil
}

// Current nodes (includes in‑flight per your cluster state)

// We dont want to disrupt nodes until scale down is complete

// Acquire limits from cluster state without bursting over

// We will not get a negative value here

// Select candidates up to maxAllowedDrifts

func (d *StaticDrift) Reason() v1.DisruptionReason {
	_ = "STUB: not implemented"
	return *new(v1.DisruptionReason)
}

func (d *StaticDrift) Class() string { _ = "STUB: not implemented"; return "" }

func (d *StaticDrift) ConsolidationType() string { _ = "STUB: not implemented"; return "" }
