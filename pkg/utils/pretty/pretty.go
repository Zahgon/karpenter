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

package pretty

import (
	"cmp"
	"regexp"

	v1 "k8s.io/api/core/v1"
)

func Concise(o any) string { _ = "STUB: not implemented"; return "" }

// Slice truncates a slice after a certain number of max items to ensure
// that the Slice isn't too long
func Slice[T any](s []T, maxItems int) string { _ = "STUB: not implemented"; return "" }

// Map truncates a map after a certain number of max items to ensure that the
// description in a log doesn't get too long
func Map[K cmp.Ordered, V any](values map[K]V, maxItems int) string {
	_ = "STUB: not implemented"
	return ""
}

func Taint(t v1.Taint) string { _ = "STUB: not implemented"; return "" }

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string { _ = "STUB: not implemented"; return "" }

func Sentence(str string) string { _ = "STUB: not implemented"; return "" }
