// Copyright 2015 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package teelogger

import (
	"context"

	"go.chromium.org/luci/common/logging"
)

type leveledLogger struct {
	minLevel logging.Level
	logger   logging.Logger
}

type teeImpl struct {
	l []leveledLogger
}

func (t *teeImpl) Debugf(fmt string, args ...interface{}) {
	t.LogCall(logging.Debug, 1, fmt, args)
}

func (t *teeImpl) Infof(fmt string, args ...interface{}) {
	t.LogCall(logging.Info, 1, fmt, args)
}

func (t *teeImpl) Warningf(fmt string, args ...interface{}) {
	t.LogCall(logging.Warning, 1, fmt, args)
}

func (t *teeImpl) Errorf(fmt string, args ...interface{}) {
	t.LogCall(logging.Error, 1, fmt, args)
}

func (t *teeImpl) LogCall(level logging.Level, calldepth int, f string, args []interface{}) {
	for _, l := range t.l {
		if level >= l.minLevel {
			l.logger.LogCall(level, calldepth+1, f, args)
		}
	}
}

// Use adds a tee logger to the context, using the logger factory in
// the context, as well as the other loggers produced by given factories.
//
// We use factories (instead of logging.Logger instances), since we must be able
// to produce logging.Logger instances bound to contexts to be able to use
// logging levels are fields (they are part of the context state).
func Use(ctx context.Context, factories ...logging.Factory) context.Context {
	if cur := logging.GetFactory(ctx); cur != nil {
		factories = append([]logging.Factory{cur}, factories...)
	}
	return logging.SetFactory(ctx, func(ic context.Context) logging.Logger {
		ll := make([]leveledLogger, len(factories))
		for i, f := range factories {
			logger := f(ic)
			ll[i] = leveledLogger{
				logger:   logger,
				minLevel: logging.GetLevel(ic),
			}
		}
		return &teeImpl{ll}
	})
}

// Filtered is a static representation of a single entry to filter messages to
// loggers by provided level.
type Filtered struct {
	Factory logging.Factory
	Level   logging.Level
}

// UseFiltered adds a tee logger to the context, using the logger factory in
// the context, as well as the other provided by given filtereds.
// Filtered loggers ignore the current logging level in the context.
//
// We use factories (instead of logging.Logger instances), since we must be able
// to produce logging.Logger instances bound to contexts to be able to use
// logging levels are fields (they are part of the context state).
// The logger instance bound to context is used with level provided by context.
func UseFiltered(ctx context.Context, filtereds ...Filtered) context.Context {
	cur := logging.GetFactory(ctx)
	count := len(filtereds)
	if cur != nil {
		count += 1
	}
	return logging.SetFactory(ctx, func(ic context.Context) logging.Logger {
		ll := make([]leveledLogger, count)
		for i, f := range filtereds {
			ll[i] = leveledLogger{
				logger:   f.Factory(ic),
				minLevel: f.Level,
			}
		}
		if cur != nil {
			ll[count-1] = leveledLogger{
				logger:   cur(ic),
				minLevel: logging.GetLevel(ic),
			}
		}
		return &teeImpl{ll}
	})
}
