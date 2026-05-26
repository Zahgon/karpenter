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

package metrics

import (
	"sync"

	opmetrics "github.com/awslabs/operatorpkg/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

// Store is a mapping from a key to a list of Metrics
// Each time Update() is called for a key on Store, the metric store ensures that all metrics are "refreshed"
// for all currently tracked metrics assigned to the key. This means that any metric that contains the same labels
// as a previous metric will be updated through the standard prometheus.Gauge metric Set() call while any metric with
// different labels than the recently fired metrics will be removed from the prometheus client response and the Store
type Store struct {
	sync.Mutex
	store map[string][]*StoreMetric
}

func NewStore() *Store { _ = "STUB: not implemented"; return nil }

// StoreMetric is a single state metric associated with a metrics.Gauge
type StoreMetric struct {
	opmetrics.GaugeMetric
	Value  float64
	Labels prometheus.Labels
}

// update is an internal non-thread-safe method for updating metrics given a key in the Store
func (s *Store) update(key string, metrics []*StoreMetric) { _ = "STUB: not implemented"; return }

// Cleanup old metrics if the old metric family has metrics that weren't updated by this round of metrics

// Update calls the update() method internally
func (s *Store) Update(key string, metrics []*StoreMetric) { _ = "STUB: not implemented"; return }

// ReplaceAll replaces all metrics in the store with the new metrics passes into the ReplaceAll function. This calls
// the update method as normal for any keys that match existing keys while removing any keys that existed in the old
// store but don't exist in the new store.
func (s *Store) ReplaceAll(newStore map[string][]*StoreMetric) { _ = "STUB: not implemented"; return }

// delete is an internal non-thread-safe method for deleting metrics given a key in the Store
func (s *Store) delete(key string) { _ = "STUB: not implemented"; return }

// Delete calls the delete() method internally
func (s *Store) Delete(key string) { _ = "STUB: not implemented"; return }
