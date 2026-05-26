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

package env

import (
	"time"
)

// WithDefaultInt returns the int value of the supplied environment variable or, if not present,
// the supplied default value. If the int conversion fails, returns the default
func WithDefaultInt(key string, def int) int { _ = "STUB: not implemented"; return 0 }

// WithDefaultInt64 returns the int value of the supplied environment variable or, if not present,
// the supplied default value. If the int conversion fails, returns the default
func WithDefaultInt64(key string, def int64) int64 { _ = "STUB: not implemented"; return 0 }

// WithDefaultString returns the string value of the supplied environment variable or, if not present,
// the supplied default value.
func WithDefaultString(key string, def string) string { _ = "STUB: not implemented"; return "" }

// WithDefaultBool returns the boolean value of the supplied environment variable or, if not present,
// the supplied default value.
func WithDefaultBool(key string, def bool) bool { _ = "STUB: not implemented"; return false }

// WithDefaultDuration returns the duration value of the supplied environment variable or, if not present,
// the supplied default value.
func WithDefaultDuration(key string, def time.Duration) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// GetRevision function is based on the function defined under https://pkg.go.dev/knative.dev/pkg@v0.0.0-20240815051656-89743d9bbf7c/changeset
// at https://github.com/knative/pkg/blob/89743d9bbf7c/changeset/commit.go#L51
func GetRevision() string { _ = "STUB: not implemented"; return "" }
