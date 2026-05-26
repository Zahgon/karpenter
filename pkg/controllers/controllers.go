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

package controllers

import (
	"context"

	"github.com/awslabs/operatorpkg/controller"
	"github.com/awslabs/operatorpkg/option"
	"k8s.io/utils/clock"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	"sigs.k8s.io/karpenter/pkg/cloudprovider"
	"sigs.k8s.io/karpenter/pkg/controllers/nodeoverlay"
	"sigs.k8s.io/karpenter/pkg/controllers/state"
	"sigs.k8s.io/karpenter/pkg/events"
)

type ControllerOptions struct {
	registrationHooks []cloudprovider.NodeLifecycleHook
}

// WithRegistrationHook registers a hook that blocks Karpenter from marking a node as registered
// until the hook's preconditions are satisfied. This is useful when a cloud provider needs to
// apply well-known labels asynchronously after instance launch (e.g., capacity reservation labels
// used by topology spread constraints).
func WithRegistrationHook(hook cloudprovider.NodeLifecycleHook) option.Function[ControllerOptions] {
	_ = "STUB: not implemented"
	return nil
}

func NewControllers(
	ctx context.Context,
	mgr manager.Manager,
	clock clock.Clock,
	kubeClient client.Client,
	recorder events.Recorder,
	cloudProvider cloudprovider.CloudProvider,
	overlayUndecoratedCloudProvider cloudprovider.CloudProvider,
	cluster *state.Cluster,
	instanceTypeStore *nodeoverlay.InstanceTypeStore,
	opts ...option.Function[ControllerOptions],
) []controller.Controller {
	_ = "STUB: not implemented"
	return nil
}

// 0.5, 1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192

// 0.5, 1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192

// 0.5, 1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192

// The cloud provider must define status conditions for the node repair controller to use to detect unhealthy nodes
