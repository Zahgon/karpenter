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

package events

import (
	"time"

	corev1 "k8s.io/api/core/v1"

	v1 "sigs.k8s.io/karpenter/pkg/apis/v1"
	"sigs.k8s.io/karpenter/pkg/events"

	storagev1 "k8s.io/api/storage/v1"
)

func EvictPod(pod *corev1.Pod, reason string) events.Event {
	_ = "STUB: not implemented"
	return *new(events.Event)
}

func DisruptPodDelete(pod *corev1.Pod, gracePeriodSeconds *int64, nodeGracePeriodTerminationTime *time.Time) events.Event {
	_ = "STUB: not implemented"
	return *new(events.Event)
}

func NodeFailedToDrain(node *corev1.Node, err error) events.Event {
	_ = "STUB: not implemented"
	return *new(events.Event)
}

func NodeAwaitingVolumeDetachmentEvent(node *corev1.Node, volumeAttachments ...*storagev1.VolumeAttachment) events.Event {
	_ = "STUB: not implemented"
	return *new(events.Event)
}

func NodeTerminationGracePeriodExpiring(node *corev1.Node, terminationTime string) events.Event {
	_ = "STUB: not implemented"
	return *new(events.Event)
}

func NodeClaimTerminationGracePeriodExpiring(nodeClaim *v1.NodeClaim, terminationTime string) events.Event {
	_ = "STUB: not implemented"
	return *new(events.Event)
}
