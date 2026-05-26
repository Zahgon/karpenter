//go:build random_test_delay

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

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
)

// If the random_test_delay build flag is used, every Expect() call gets an additional random delay added to it.  This
// is intended to attempt to make tests more robust by eliminating tests that depend on timing.
func init() {
	gomega.Default = &gomegaWrapper{
		inner: gomega.Default,
		r:     rand.New(rand.NewSource(ginkgo.GinkgoRandomSeed())),
	}
}

type gomegaWrapper struct {
	inner gomega.Gomega
	mu    sync.Mutex
	r     *rand.Rand
}

func (g *gomegaWrapper) randomDelay() { _ = "STUB: not implemented"; return }

func (g *gomegaWrapper) Ω(actual interface{}, extra ...interface{}) types.Assertion {
	_ = "STUB: not implemented"
	return *new(types.Assertion)
}

func (g *gomegaWrapper) Expect(actual interface{}, extra ...interface{}) types.Assertion {
	_ = "STUB: not implemented"
	return *new(types.Assertion)
}

func (g *gomegaWrapper) ExpectWithOffset(offset int, actual interface{}, extra ...interface{}) types.Assertion {
	_ = "STUB: not implemented"
	return *new(types.Assertion)
}

func (g *gomegaWrapper) Eventually(actualOrCtx interface{}, args ...interface{}) types.AsyncAssertion {
	_ = "STUB: not implemented"
	return *new(types.AsyncAssertion)
}

func (g *gomegaWrapper) EventuallyWithOffset(offset int, actual interface{}, args ...interface{}) types.AsyncAssertion {
	_ = "STUB: not implemented"
	return *new(types.AsyncAssertion)
}

func (g *gomegaWrapper) Consistently(actualOrCtx interface{}, args ...interface{}) types.AsyncAssertion {
	_ = "STUB: not implemented"
	return *new(types.AsyncAssertion)
}

func (g *gomegaWrapper) ConsistentlyWithOffset(offset int, actualOrCtx interface{}, args ...interface{}) types.AsyncAssertion {
	_ = "STUB: not implemented"
	return *new(types.AsyncAssertion)
}

func (g *gomegaWrapper) SetDefaultEventuallyTimeout(duration time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (g *gomegaWrapper) SetDefaultEventuallyPollingInterval(duration time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (g *gomegaWrapper) SetDefaultConsistentlyDuration(duration time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (g *gomegaWrapper) SetDefaultConsistentlyPollingInterval(duration time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (g *gomegaWrapper) Inner() gomega.Gomega {
	_ = "STUB: not implemented"
	return *new(gomega.Gomega)
}
