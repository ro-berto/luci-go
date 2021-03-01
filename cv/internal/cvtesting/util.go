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

// Package cvtesting reduces boilerplate in tests.
package cvtesting

import (
	"context"
	cryptorand "crypto/rand"
	"encoding/hex"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"testing"
	"time"

	nativeDatastore "cloud.google.com/go/datastore"
	"google.golang.org/api/option"

	"go.chromium.org/luci/auth"
	"go.chromium.org/luci/common/clock"
	"go.chromium.org/luci/common/clock/testclock"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/common/logging/gologger"
	"go.chromium.org/luci/gae/filter/txndefer"
	"go.chromium.org/luci/gae/impl/cloud"
	"go.chromium.org/luci/gae/impl/memory"
	"go.chromium.org/luci/gae/service/datastore"
	"go.chromium.org/luci/gae/service/info"
	serverauth "go.chromium.org/luci/server/auth"
	"go.chromium.org/luci/server/tq"
	"go.chromium.org/luci/server/tq/tqtesting"

	"go.chromium.org/luci/cv/internal/config"
	gf "go.chromium.org/luci/cv/internal/gerrit/gerritfake"

	. "github.com/smartystreets/goconvey/convey"
)

// TODO(tandrii): add fake config generation facilities.

// Test encapsulates typical setup for CV test.
//
// Typical use:
//   ct := cvtesting.Test{}
//   ctx, cancel := ct.SetUp()
//   defer cancel()
type Test struct {
	// Cfg manipulates CV config.
	Cfg config.TestController
	// GFake is a Gerrit fake. Defaults to an empty one.
	GFake *gf.Fake
	// TQ allows to run TQ tasks.
	TQ *tqtesting.Scheduler
	// Clock allows to move time forward.
	// By default, the time is moved automatically is something waits on it.
	Clock testclock.TestClock

	// MaxDuration limits how long a test can run as a fail safe.
	//
	// Defaults to 10s to most likely finish in pre/post submit tests,
	// with limited CPU resources.
	// Set to ~10ms when debugging a hung test.
	MaxDuration time.Duration

	// cleanups are executed in reverse order in cleanup().
	cleanups []func()
}

func (t *Test) SetUp() (ctx context.Context, deferme func()) {
	// Set defaults.
	if t.MaxDuration == time.Duration(0) {
		if os.Getenv("DATASTORE_PROJECT") == "" {
			t.MaxDuration = 10 * time.Second
		} else {
			t.MaxDuration = 60 * time.Second
		}
	}
	if t.GFake == nil {
		t.GFake = &gf.Fake{}
	}

	ctx = context.Background()
	if testing.Verbose() {
		ctx = logging.SetLevel(gologger.StdConfig.Use(ctx), logging.Debug)
	}

	topCtx, cancel := context.WithTimeout(ctx, t.MaxDuration)
	t.cleanups = append(t.cleanups, func() {
		// Fail the test if the topCtx has timed out.
		So(topCtx.Err(), ShouldBeNil)
		cancel()
	})

	// Use a date-time that is easy to eyeball in logs.
	utc := time.Date(2020, time.February, 2, 10, 30, 00, 0, time.UTC)
	// But set it up in a clock as a local time to expose incorrect assumptions of UTC.
	now := time.Date(2020, time.February, 2, 13, 30, 00, 0, time.FixedZone("Fake local", 3*60*60))
	So(now.Equal(utc), ShouldBeTrue)
	ctx, t.Clock = testclock.UseTime(topCtx, now)
	t.Clock.SetTimerCallback(func(dur time.Duration, _ clock.Timer) {
		// Move fake time forward whenever someone's waiting for it.
		t.Clock.Add(dur)
	})

	ctx = t.installDS(ctx)
	ctx = txndefer.FilterRDS(ctx)

	ctx = t.GFake.Install(ctx)
	ctx, t.TQ = tq.TestingContext(ctx, nil)
	return ctx, t.cleanup
}

func (t *Test) cleanup() {
	for i := len(t.cleanups) - 1; i >= 0; i-- {
		t.cleanups[i]()
	}
}

func (t *Test) installDS(ctx context.Context) context.Context {
	if ctx, ok := t.installDSReal(ctx); ok {
		return ctx
	}
	if ctx, ok := t.installDSEmulator(ctx); ok {
		return ctx
	}
	ctx = memory.Use(ctx)
	// CV runs against Firestore backend, which is consistent.
	datastore.GetTestable(ctx).Consistent(true)
	datastore.GetTestable(ctx).AutoIndex(true)
	return ctx
}

// installDSProd configures CV tests to run with actual DS.
//
// If DATASTORE_PROJECT ENV var isn't set, returns false.
//
// To use, first
//
//    $ luci-auth context -- bash
//    $ export DATASTORE_PROJECT=my-cloud-project-with-datastore
//
// and then run go tests the usual way, e.g.:
//
//    $ go test ./...
func (t *Test) installDSReal(ctx context.Context) (context.Context, bool) {
	project := os.Getenv("DATASTORE_PROJECT")
	if project == "" {
		return ctx, false
	}
	if project == "luci-change-verifier" {
		panic("Don't use production CV project. Using -dev is OK.")
	}

	at := auth.NewAuthenticator(ctx, auth.SilentLogin, auth.Options{
		Scopes: serverauth.CloudOAuthScopes,
	})
	ts, err := at.TokenSource()
	if err != nil {
		err = errors.Annotate(err, "failed to initialize the token source (are you in `$ luci-auth context`?)").Err()
		So(err, ShouldBeNil)
	}

	logging.Debugf(ctx, "Using DS of project %q", project)
	client, err := nativeDatastore.NewClient(ctx, project, option.WithTokenSource(ts))
	So(err, ShouldBeNil)
	return t.installDSshared(ctx, project, client), true
}

// installDSEmulator configures CV tests to run with DS emulator.
//
// If DATASTORE_EMULATOR_HOST ENV var isn't set, returns false.
//
// To use, run
//
//     $ gcloud beta emulators datastore start --consistency=1.0
//
// and export DATASTORE_EMULATOR_HOST as printed by above command.
//
// NOTE: as of Feb 2021, emulator runs in legacy Datastore mode,
// not Firestore.
func (t *Test) installDSEmulator(ctx context.Context) (context.Context, bool) {
	emulatorHost := os.Getenv("DATASTORE_EMULATOR_HOST")
	if emulatorHost == "" {
		return ctx, false
	}

	logging.Debugf(ctx, "Using DS emulator at %q", emulatorHost)
	client, err := nativeDatastore.NewClient(ctx, "luci-gae-emulator-test")
	So(err, ShouldBeNil)
	return t.installDSshared(ctx, "luci-gae-emulator-test", client), true
}

func (t *Test) installDSshared(ctx context.Context, cloudProject string, client *nativeDatastore.Client) context.Context {
	t.cleanups = append(t.cleanups, func() {
		if err := client.Close(); err != nil {
			logging.Errorf(ctx, "failed to close DS client: %s", err)
		}
	})
	ctx = (&cloud.Config{ProjectID: cloudProject, DS: client}).Use(ctx, nil)
	maybeCleanupOldDSNamespaces(ctx)

	// Enter a namespace for this tests.
	ns := genDSNamespaceName(time.Now())
	logging.Debugf(ctx, "Using %q DS namespace", ns)
	ctx = info.MustNamespace(ctx, ns)
	// Failure to clear is hard before the test,
	// ignored after the test.
	So(clearDS(ctx), ShouldBeNil)
	t.cleanups = append(t.cleanups, func() {
		if err := clearDS(ctx); err != nil {
			logging.Errorf(ctx, "failed to clean DS namespace %s: %s", ns, err)
		}
	})
	return ctx
}

func genDSNamespaceName(t time.Time) string {
	rnd := make([]byte, 8)
	if _, err := cryptorand.Read(rnd); err != nil {
		panic(err)
	}
	return fmt.Sprintf("testing-%s-%s", time.Now().Format("2006-01-02"), hex.EncodeToString(rnd))
}

var dsNamespaceRegexp = regexp.MustCompile(`^testing-(\d{4}-\d\d-\d\d)-[0-9a-f]+$`)

func isOldTestDSNamespace(ns string, now time.Time) bool {
	m := dsNamespaceRegexp.FindSubmatch([]byte(ns))
	if len(m) == 0 {
		return false
	}
	// Anything up ~2 days old should be kept to avoid accidentally removing
	// currently under test namespace in presence of timezones and out of sync
	// clocks.
	const maxAge = 2 * 24 * time.Hour
	t, err := time.Parse("2006-01-02", string(m[1]))
	if err != nil {
		panic(err)
	}
	return now.Sub(t) > maxAge
}

func clearDS(ctx context.Context) error {
	// Execute a kindless query to clear entire namespace.
	q := datastore.NewQuery("").KeysOnly(true)
	var allKeys []*datastore.Key
	if err := datastore.GetAll(ctx, q, &allKeys); err != nil {
		return errors.Annotate(err, "failed to get entities").Err()
	}
	if err := datastore.Delete(ctx, allKeys); err != nil {
		return errors.Annotate(err, "failed to delete %d entities", len(allKeys)).Err()
	}
	return nil
}

func maybeCleanupOldDSNamespaces(ctx context.Context) {
	if rand.Intn(1024) < 1020 { // ~99% of cases.
		return
	}
	q := datastore.NewQuery("__namespace__").KeysOnly(true)
	var allKeys []*datastore.Key
	if err := datastore.GetAll(ctx, q, &allKeys); err != nil {
		logging.Warningf(ctx, "failed to query all namespaces: %s", err)
		return
	}
	now := time.Now()
	var toDelete []string
	for _, k := range allKeys {
		ns := k.StringID()
		if isOldTestDSNamespace(ns, now) {
			toDelete = append(toDelete, ns)
		}
	}
	logging.Debugf(ctx, "cleaning up %d old namespaces", len(toDelete))
	for _, ns := range toDelete {
		logging.Debugf(ctx, "cleaning up %s", ns)
		if err := clearDS(info.MustNamespace(ctx, ns)); err != nil {
			logging.Errorf(ctx, "failed to clean old DS namespace %s: %s", ns, err)
		}
	}
}
