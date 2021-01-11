// Copyright 2020 The LUCI Authors.
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

package common

import (
	"context"
	"strings"

	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/common/retry/transient"
	"go.chromium.org/luci/server/tq"
)

// MostSevereError returns the most severe error in order of
// non-transient => transient => nil.
//
// Walks over potentially recursive errors.MultiError errors only.
//
// Returns only singular errors or nil if input was nil.
func MostSevereError(err error) error {
	if err == nil {
		return nil
	}
	errs, ok := err.(errors.MultiError)
	if !ok {
		return err
	}
	var firstTrans error
	for _, err := range errs {
		switch err = MostSevereError(err); {
		case err == nil:
		case !transient.Tag.In(err):
			return err
		case firstTrans == nil:
			firstTrans = err
		}
	}
	return firstTrans
}

// TQifyError does final error processing before returning from a TQ handler.
//
// * logs error stack,
// * non-transient errors are tagged with tq.Fatal to avoid retries.
func TQifyError(ctx context.Context, err error) error {
	if err == nil {
		return nil
	}
	LogError(
		ctx,
		err,

		// These packages are not useful in CV tests:
		"github.com/smartystreets/goconvey/convey",
		"github.com/jtolds/gls",
		// These packages are not useful in production:
		"go.chromium.org/luci/server",
		"go.chromium.org/luci/server/tq",
		"go.chromium.org/luci/server/router",
	)
	if !transient.Tag.In(err) {
		err = tq.Fatal.Apply(err)
	}
	return err
}

func LogError(ctx context.Context, err error, excludePackages ...string) {
	const maxBytesPerEntry = 64 * 1024 // Cloud Logging hard limit is 256KB.
	logError(ctx, err, maxBytesPerEntry, excludePackages...)
}

func logError(ctx context.Context, err error, maxBytesPerEntry int, excludePackages ...string) {
	// TODO(tandrii): upstream this fork of errors.Log.
	log := logging.Get(ctx)
	lines := errors.RenderStack(err, excludePackages...)
	for len(lines) > 0 {
		bytes := 0
		batch := lines
		for i, l := range lines {
			bytes += len(l)
			// Always allow first line into the batch to ensure progress, even if it's
			// too big. Practically, this should never happen, but if it doesn't, we
			// shouldn't be looping forever.
			if bytes > maxBytesPerEntry && i > 0 {
				batch = lines[:i] // excluding i-th.
				break
			}
		}
		lines = lines[len(batch):]
		log.Errorf("%s", strings.Join(batch, "\n"))
	}
}
