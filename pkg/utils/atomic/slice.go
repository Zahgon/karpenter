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

package atomic

import (
	"sync"
)

// Slice exposes a slice of a type in a race-free manner.
type Slice[T any] struct {
	mu     sync.RWMutex
	values []T
}

func (a *Slice[T]) Reset() { _ = "STUB: not implemented"; return }

func (a *Slice[T]) Add(input T) { _ = "STUB: not implemented"; return }

func (a *Slice[T]) Range(f func(pool T) bool) { _ = "STUB: not implemented"; return }

func (a *Slice[T]) Set(values []T) { _ = "STUB: not implemented"; return }

func (a *Slice[T]) Len() int { _ = "STUB: not implemented"; return 0 }
