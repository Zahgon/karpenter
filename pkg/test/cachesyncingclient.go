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
	"time"

	"github.com/avast/retry-go"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// CacheSyncingClient exists for tests that need to use custom fieldSelectors (thus, they need a client cache)
// and also need consistency in their testing by waiting for caches to sync after performing WRITE operations
// NOTE: This cache sync doesn't sync with third-party operations on the api-server
type CacheSyncingClient struct {
	client.Client
}

// If we timeout on polling, the assumption is that the cache updated to a newer version
// and we missed the current WRITE operation that we just performed
var pollingOptions = []retry.Option{
	retry.Attempts(100), // This whole poll should take ~1s
	retry.Delay(time.Millisecond * 10),
	retry.DelayType(retry.FixedDelay),
}

func (c *CacheSyncingClient) Create(ctx context.Context, obj client.Object, opts ...client.CreateOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *CacheSyncingClient) Delete(ctx context.Context, obj client.Object, opts ...client.DeleteOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *CacheSyncingClient) Update(ctx context.Context, obj client.Object, opts ...client.UpdateOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *CacheSyncingClient) Patch(ctx context.Context, obj client.Object, patch client.Patch, opts ...client.PatchOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *CacheSyncingClient) DeleteAllOf(ctx context.Context, obj client.Object, opts ...client.DeleteAllOfOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *CacheSyncingClient) Status() client.StatusWriter {
	_ = "STUB: not implemented"
	return *new(client.StatusWriter)
}

type cacheSyncingStatusWriter struct {
	client client.Client
}

func (c *cacheSyncingStatusWriter) Create(_ context.Context, _ client.Object, _ client.Object, _ ...client.SubResourceCreateOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *cacheSyncingStatusWriter) Update(ctx context.Context, obj client.Object, opts ...client.SubResourceUpdateOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *cacheSyncingStatusWriter) Patch(ctx context.Context, obj client.Object, patch client.Patch, opts ...client.SubResourcePatchOption) error {
	_ = "STUB: not implemented"
	return nil
}

func objectSynced(ctx context.Context, c client.Client, obj client.Object) error {
	_ = "STUB: not implemented"
	return nil
}

// If the object isn't found, we assume that the cache was synced since the Update operation must have caused
// the object to get completely removed (like a finalizer update)
