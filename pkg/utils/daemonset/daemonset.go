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

package daemonset

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

func PodForDaemonSet(daemonSet *appsv1.DaemonSet) *corev1.Pod {
	_ = "STUB: not implemented"
	return nil
}

// The API server performs defaulting to merge limits into requests for pods. However, this is not performed for higher
// level objects (e.g. deployments, daemonsets, etc). We should perform this defaulting ourselves when we create a fake
// pod for a daemonset.

// mergeResourceLimitsIntoRequests merges resource limits into requests if no request exists for the given resource.
// This is performed in place on the provided container.
func mergeResourceLimitsIntoRequests(container *corev1.Container) {
	_ = "STUB: not implemented"
	return
}
