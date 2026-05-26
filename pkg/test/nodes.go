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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1 "sigs.k8s.io/karpenter/pkg/apis/v1"
)

type NodeOptions struct {
	metav1.ObjectMeta
	ReadyStatus   corev1.ConditionStatus
	ReadyReason   string
	Conditions    []corev1.NodeCondition
	Unschedulable bool
	ProviderID    string
	Taints        []corev1.Taint
	Allocatable   corev1.ResourceList
	Capacity      corev1.ResourceList
}

func Node(overrides ...NodeOptions) *corev1.Node { _ = "STUB: not implemented"; return nil }

func NodeClaimLinkedNode(nodeClaim *v1.NodeClaim) *corev1.Node {
	_ = "STUB: not implemented"
	return nil
}
