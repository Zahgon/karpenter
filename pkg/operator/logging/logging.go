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

package logging

import (
	"context"

	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
)

// NopLogger is used to throw away logs when we don't actually want to log in
// certain portions of the code since logging would be too noisy
var NopLogger = zapr.NewLogger(zap.NewNop())

const (
	Unknown = "unknown"
	Commit  = "commit"
)

func DefaultZapConfig(ctx context.Context, component string) zap.Config {
	_ = "STUB: not implemented"
	return *new(zap.Config)
}

// Webhook log level can only be configured directly through the zap-config
// Webhooks are deprecated, so support for changing their log level is also deprecated

// NewLogger returns a configured *zap.SugaredLogger
func NewLogger(ctx context.Context, component string) *zap.Logger {
	_ = "STUB: not implemented"
	return nil
}

func WithCommit(logger *zap.Logger) *zap.Logger { _ = "STUB: not implemented"; return nil }

// Enrich logs with the components git revision.

type ignoreDebugEventsSink struct {
	name string
	sink logr.LogSink
}

func (i ignoreDebugEventsSink) Init(ri logr.RuntimeInfo) { _ = "STUB: not implemented"; return }

func (i ignoreDebugEventsSink) Enabled(level int) bool { _ = "STUB: not implemented"; return false }
func (i ignoreDebugEventsSink) Info(level int, msg string, keysAndValues ...any) {
	_ = "STUB: not implemented"
	// ignore debug "events" logs
	return
}

func (i ignoreDebugEventsSink) Error(err error, msg string, keysAndValues ...any) {
	_ = "STUB: not implemented"
	return
}

func (i ignoreDebugEventsSink) WithValues(keysAndValues ...any) logr.LogSink {
	_ = "STUB: not implemented"
	return *new(logr.LogSink)
}

func (i ignoreDebugEventsSink) WithName(name string) logr.LogSink {
	_ = "STUB: not implemented"
	return *new(logr.LogSink)
}

// IgnoreDebugEvents wraps the logger with one that ignores any debug logs coming from a logger named "events".  This
// prevents every event we write from creating a debug log which spams the log file during scale-ups due to recording
// pod scheduling decisions as events for visibility.
func IgnoreDebugEvents(logger logr.Logger) logr.Logger {
	_ = "STUB: not implemented"
	return *new(logr.Logger)
}
