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

package injection

import (
	"context"

	"sigs.k8s.io/karpenter/pkg/operator/options"
)

type controllerNameKeyType struct{}

var controllerNameKey = controllerNameKeyType{}

func WithControllerName(ctx context.Context, name string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func GetControllerName(ctx context.Context) string { _ = "STUB: not implemented"; return "" }

func WithOptionsOrDie(ctx context.Context, opts ...options.Injectable) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}
