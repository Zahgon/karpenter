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
	"math/rand"
	"sync"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1 "sigs.k8s.io/karpenter/pkg/apis/v1"
)

const DiscoveryLabel = "testing/cluster"

var (
	sequentialNumber     = 0
	randomizer           = rand.New(rand.NewSource(time.Now().UnixNano())) //nolint
	sequentialNumberLock = new(sync.Mutex)
)

func RandomName() string { _ = "STUB: not implemented"; return "" }

func NamespacedObjectMeta(overrides ...metav1.ObjectMeta) metav1.ObjectMeta {
	_ = "STUB: not implemented"
	return *new(metav1.ObjectMeta)
}

func ObjectMeta(overrides ...metav1.ObjectMeta) metav1.ObjectMeta {
	_ = "STUB: not implemented"
	return *new(metav1.ObjectMeta)
}

// For cleanup discovery

func TemplateObjectMeta(overrides ...v1.ObjectMeta) v1.ObjectMeta {
	_ = "STUB: not implemented"
	return *new(v1.ObjectMeta)
}

// For cleanup discovery

func MustMerge[T any](dest T, srcs ...T) T { _ = "STUB: not implemented"; return *new(T) }

func RandomProviderID() string { _ = "STUB: not implemented"; return "" }

func ProviderID(base string) string { _ = "STUB: not implemented"; return "" }
