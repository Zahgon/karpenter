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

package common

import (
	"context"
	_ "embed"
	"flag"
	"testing"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"

	v1 "sigs.k8s.io/karpenter/pkg/apis/v1"
	. "sigs.k8s.io/karpenter/pkg/utils/testing" //nolint:staticcheck
	"sigs.k8s.io/karpenter/test/pkg/debug"
)

type ContextKey string

const GitRefContextKey = ContextKey("gitRef")

// I need to add the the default kwok nodeclass path
// That way it's not defined in code but we use it when we initialize the nodeclass
var (
	//go:embed default_kowknodeclass.yaml
	defaultNodeClass []byte
	//go:embed default_nodepool.yaml
	defaultNodePool []byte
	nodeClassPath   = flag.String("default-nodeclass", "", "Pass in a default cloud specific node class")
	nodePoolPath    = flag.String("default-nodepool", "", "Pass in a default karpenter nodepool")
)

type Environment struct {
	context.Context
	cancel context.CancelFunc

	TimeIntervalCollector *debug.TimeIntervalCollector
	Client                client.Client
	Config                *rest.Config
	KubeClient            kubernetes.Interface
	Monitor               *Monitor
	DefaultNodeClass      *unstructured.Unstructured

	OutputDir         string
	StartingNodeCount int
}

func NewEnvironment(t *testing.T) *Environment { _ = "STUB: not implemented"; return nil }

// Get the output dir if it's set

func (env *Environment) Stop() { _ = "STUB: not implemented"; return }

func NewConfig() *rest.Config { _ = "STUB: not implemented"; return nil }

func NewClient(ctx context.Context, config *rest.Config) client.Client {
	_ = "STUB: not implemented"
	return *new(client.Client)
}

func (env *Environment) DefaultNodePool(nodeClass *unstructured.Unstructured) *v1.NodePool {
	_ = "STUB: not implemented"
	return nil
}

// Update to use the provided default nodeclass

func (env *Environment) IsDefaultNodeClassKWOK() bool { _ = "STUB: not implemented"; return false }

func decodeNodeClass() *unstructured.Unstructured {
	_ = "STUB: not implemented"
	// Open the file
	return nil
}
