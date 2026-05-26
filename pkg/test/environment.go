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
	"context"

	"github.com/awslabs/operatorpkg/option"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/util/version"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
)

type Environment struct {
	envtest.Environment

	Client              client.Client
	KubernetesInterface kubernetes.Interface
	Version             *version.Version
	Done                chan struct{}
	Cancel              context.CancelFunc
}

type EnvironmentOptions struct {
	crds          []*apiextensionsv1.CustomResourceDefinition
	fieldIndexers []func(cache.Cache) error
	configOptions []func(*rest.Config)
}

// WithCRDs registers the specified CRDs to the apiserver for use in testing
func WithCRDs(crds ...*apiextensionsv1.CustomResourceDefinition) option.Function[EnvironmentOptions] {
	_ = "STUB: not implemented"
	return nil
}

// WithFieldIndexers expects a function that indexes fields against the cache such as cache.IndexField(...).
//
// Note: Only use when necessary, the use of field indexers in functional tests requires the use of the cache syncing
// client, which can have significant drawbacks for test performance.
func WithFieldIndexers(fieldIndexers ...func(cache.Cache) error) option.Function[EnvironmentOptions] {
	_ = "STUB: not implemented"
	return nil
}

// WithConfigOptions allows customization of the rest.Config before client creation
func WithConfigOptions(options ...func(*rest.Config)) option.Function[EnvironmentOptions] {
	_ = "STUB: not implemented"
	return nil
}

func NodeProviderIDFieldIndexer(ctx context.Context) func(cache.Cache) error {
	_ = "STUB: not implemented"
	return nil
}

func NodeClaimProviderIDFieldIndexer(ctx context.Context) func(cache.Cache) error {
	_ = "STUB: not implemented"
	return nil
}

func NodeClaimNodeClassRefFieldIndexer(ctx context.Context) func(cache.Cache) error {
	_ = "STUB: not implemented"
	return nil
}

func NodePoolNodeClassRefFieldIndexer(ctx context.Context) func(cache.Cache) error {
	_ = "STUB: not implemented"
	return nil
}

func VolumeAttachmentFieldIndexer(ctx context.Context) func(cache.Cache) error {
	_ = "STUB: not implemented"
	return nil
}

func NewEnvironment(options ...option.Function[EnvironmentOptions]) *Environment {
	_ = "STUB: not implemented"
	return nil
}

// PodAffinityNamespaceSelector is used for label selectors in pod affinities.  If the feature-gate is turned off,
// the api-server just clears out the label selector so we never see it.  If we turn it on, the label selectors
// are passed to us and we handle them. This feature is alpha in apiextensionsv1.21, beta in apiextensionsv1.22 and will be GA in 1.24. See
// https://github.com/kubernetes/enhancements/issues/2249 for more info.

//MinDomains got promoted to stable in 1.32

// MinDomainsInPodTopologySpread enforces a minimum number of eligible node domains for pod scheduling
// See https://kubernetes.io/docs/concepts/scheduling-eviction/topology-spread-constraints/#spread-constraint-definition
// Ref: https://github.com/aws/karpenter-core/pull/330

// Apply any config overrides

// We use a modified client if we need field indexers

func (e *Environment) Stop() error { _ = "STUB: not implemented"; return nil }
