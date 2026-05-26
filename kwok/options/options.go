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

package options

import (
	"context"

	"sigs.k8s.io/karpenter/pkg/operator/options"
)

func init() {
	options.Injectables = append(options.Injectables, &Options{})
}

type optionsKey struct{}

// Options contains all CLI flags / env vars for the KWOK cloudprovider.
type Options struct {
	InstanceTypesFilePath string
}

func (o *Options) AddFlags(fs *options.FlagSet) { _ = "STUB: not implemented"; return }

func (o *Options) Parse(fs *options.FlagSet, args ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (o *Options) ToContext(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func ToContext(ctx context.Context, opts *Options) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func FromContext(ctx context.Context) *Options { _ = "STUB: not implemented"; return nil }
