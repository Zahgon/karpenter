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
	corev1 "k8s.io/api/core/v1"

	v1 "sigs.k8s.io/karpenter/pkg/apis/v1"
)

// NodeClaim creates a test NodeClaim with defaults that can be overridden by overrides.
// Overrides are applied in order, with a last write wins semantic.
func NodeClaim(overrides ...v1.NodeClaim) *v1.NodeClaim { _ = "STUB: not implemented"; return nil }

func NodeClaimAndNode(overrides ...v1.NodeClaim) (*v1.NodeClaim, *corev1.Node) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NodeClaimsAndNodes creates homogeneous groups of NodeClaims and Nodes based on the passed in options, evenly divided by the total nodeclaims requested
func NodeClaimsAndNodes(total int, options ...v1.NodeClaim) ([]*v1.NodeClaim, []*corev1.Node) {
	_ = "STUB: not implemented"
	return nil, nil
}
