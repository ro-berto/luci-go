// Copyright 2022 The LUCI Authors.
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

package listener

import (
	"context"
	"reflect"
	"sync"

	"cloud.google.com/go/pubsub"

	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/common/retry/transient"
	listenerpb "go.chromium.org/luci/cv/settings/listener"
)

const (
	defaultNumGoroutines          = 10
	defaultMaxOutstandingMessages = 1000
)

type processor interface {
	// process processes a given pubsub message.
	process(context.Context, *pubsub.Message) error
}

// subscriber receives and processes messages from a given subscription.
type subscriber struct {
	sub *pubsub.Subscription
	// The message processor
	proc processor

	// protect cancelFunc and done
	mu sync.Mutex
	// nil before start
	cancelFunc context.CancelFunc
	// nil before start
	done chan struct{}
}

// start starts a goroutine to receive and process messages from
// the subscription continuously.
//
// The goroutine stops in any of the following occurrences.
// - the context, passed to start, is done
// - stop() is called
//
// Cannot be called while the subscriber is running.
func (s *subscriber) start(ctx context.Context) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.done != nil {
		select {
		case <-s.done:
		default:
			return errors.Reason("cannot start again, while the subscriber is running").Err()
		}
	}
	switch ex, err := s.sub.Exists(ctx); {
	case err != nil:
		return errors.Annotate(err, "pubsub.Exists(%s)", s.sub.ID()).Err()
	case !ex:
		return errors.Reason("subscription %q doesn't exist", s.sub.ID()).Err()
	}

	cctx, cancel := context.WithCancel(ctx)
	cctx = logging.SetField(cctx, "subscriptionID", s.sub.ID())
	s.cancelFunc = cancel
	s.done = make(chan struct{})
	ch := make(chan struct{})

	var procName string
	switch t := reflect.TypeOf(s.proc); {
	case t.Kind() == reflect.Ptr:
		procName = t.Elem().Name()
	default:
		procName = t.Name()
	}

	go func() {
		close(ch)
		// cancel the context on exit.
		defer cancel()
		defer close(s.done)
		err := s.sub.Receive(cctx, func(ctx context.Context, m *pubsub.Message) {
			switch err := s.proc.process(ctx, m); {
			case err == nil:
				m.Ack()
			case transient.Tag.In(err):
				m.Nack()
				logging.Warningf(cctx, "%s.process: transient error %s", procName, err)
			default:
				// Ack the message, if there is a permanent error, as retry
				// will unlikely fix the error.
				//
				// Full poll should rediscover the lost event.
				m.Ack()
				logging.Errorf(cctx, "%s.process: permanent error %s", procName, err)
			}
		})
		if err != nil {
			// cctx may be no longer valid at this moment.
			// use ctx for logging, instead.
			logging.Errorf(ctx, "subscriber.start.Receive(%s): %s", s.sub.ID(), err)
		}
	}()

	select {
	case <-ch:
	case <-ctx.Done():
		// if the given context is done before the new goroutine starts,
		// cancels the goroutine context so that it will be terminated
		// after the start.
		return ctx.Err()
	}
	return nil
}

func (s *subscriber) stop(ctx context.Context) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.cancelFunc != nil {
		s.cancelFunc()
		select {
		case <-s.done:
		case <-ctx.Done():
		}
	}
}

// sameReceiveSettings returns true if the current receive settings are the same
// as given ones.
func (s *subscriber) sameReceiveSettings(in *listenerpb.Settings_ReceiveSettings) bool {
	nGoroutines := defaultNumGoroutines
	maxOutstandingMessages := defaultMaxOutstandingMessages

	if in != nil {
		nGoroutines = int(in.NumGoroutines)
		maxOutstandingMessages = int(in.MaxOutstandingMessages)
	}
	return (s.sub.ReceiveSettings.NumGoroutines == nGoroutines &&
		s.sub.ReceiveSettings.MaxOutstandingMessages == maxOutstandingMessages)
}

func (s *subscriber) isStopped() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.done != nil {
		select {
		case <-s.done:
		default:
			return false
		}
	}
	return true
}
