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

package terminator

import (
	"context"
	"sync"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	"sigs.k8s.io/karpenter/pkg/events"
)

const (
	evictionQueueBaseDelay = 100 * time.Millisecond
	evictionQueueMaxDelay  = 10 * time.Second
	minReconciles          = 100
	maxReconciles          = 5000

	multiplePodDisruptionBudgetsError = "This pod has more than one PodDisruptionBudget, which the eviction subresource does not support."
)

type NodeDrainError struct {
	error
}

func NewNodeDrainError(err error) *NodeDrainError { _ = "STUB: not implemented"; return nil }

func IsNodeDrainError(err error) bool { _ = "STUB: not implemented"; return false }

type QueueKey struct {
	types.NamespacedName
	UID types.UID
}

func NewQueueKey(pod *corev1.Pod) QueueKey { _ = "STUB: not implemented"; return *new(QueueKey) }

type Queue struct {
	sync.Mutex

	source chan event.TypedGenericEvent[*corev1.Pod]
	set    sets.Set[QueueKey]

	kubeClient client.Client
	recorder   events.Recorder
}

func NewQueue(kubeClient client.Client, recorder events.Recorder) *Queue {
	_ = "STUB: not implemented"
	return nil
}

func (q *Queue) Name() string { _ = "STUB: not implemented"; return "" }

func (q *Queue) Register(ctx context.Context, m manager.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

// qps scales linearly with concurrentReconciles, bucket size is 10 * qps

// Add adds pods to the Queue
func (q *Queue) Add(pods ...*corev1.Pod) { _ = "STUB: not implemented"; return }

func (q *Queue) Has(pod *corev1.Pod) bool { _ = "STUB: not implemented"; return false }

func (q *Queue) Reconcile(ctx context.Context, pod *corev1.Pod) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

//This is a different pod than the one the queue, we should exit without evicting
//This race happens when a pod is replaced with one that has the same namespace and name
//but a different UID after the original pod is added to the queue but before the
//controller can reconcile on it

// Evict the pod

// status codes for the eviction API are defined here:
// https://kubernetes.io/docs/concepts/scheduling-eviction/api-eviction/#how-api-initiated-eviction-works

// 404 - The pod no longer exists
// https://github.com/kubernetes/kubernetes/blob/ad19beaa83363de89a7772f4d5af393b85ce5e61/pkg/registry/core/pod/storage/eviction.go#L160
// 409 - The pod exists, but it is not the same pod that we initiated the eviction on
// https://github.com/kubernetes/kubernetes/blob/ad19beaa83363de89a7772f4d5af393b85ce5e61/pkg/registry/core/pod/storage/eviction.go#L318

// The pod exists and is the same pod, we need to continue
// 429 - PDB violation
// Regardless of whether the PDBs allow disruptions, Kubernetes doesn't support multiple PDBs on a single pod:
// https://github.com/kubernetes/kubernetes/blob/84cacae7046df93c1f6f8ea97c912d948e1ad06a/pkg/registry/core/pod/storage/eviction.go#L226

// Its not a PDB, we should requeue

func evictionReason(ctx context.Context, pod *corev1.Pod, kubeClient client.Client) string {
	_ = "STUB: not implemented"
	return ""
}
