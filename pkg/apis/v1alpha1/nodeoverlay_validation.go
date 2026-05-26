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

package v1alpha1

import (
	"context"
)

// RuntimeValidate will be used to validate any part of the CRD that can not be validated at CRD creation
func (in *NodeOverlay) RuntimeValidate(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// This function is used by the NodeOverlay validation webhook to verify the nodeoverlay requirements.
// When this function is called, the nodeoverlay's requirements do not include the requirements from labels.
// NodeOverlay requirements only support well known labels.
func (in *NodeOverlaySpec) validateRequirements(ctx context.Context) (errs error) {
	_ = "STUB: not implemented"
	return nil
}

func (in *NodeOverlaySpec) validateCapacity() (errs error) { _ = "STUB: not implemented"; return nil }
