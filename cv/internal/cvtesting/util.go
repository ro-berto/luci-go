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

package cvtesting

import (
	"context"
	cryptorand "crypto/rand"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net/mail"
	"os"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"time"

	nativeDatastore "cloud.google.com/go/datastore"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"

	"go.chromium.org/luci/auth"
	"go.chromium.org/luci/auth/identity"
	"go.chromium.org/luci/common/clock"
	"go.chromium.org/luci/common/clock/testclock"
	"go.chromium.org/luci/common/data/stringset"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/common/logging/gologger"
	"go.chromium.org/luci/common/tsmon"
	"go.chromium.org/luci/common/tsmon/distribution"
	"go.chromium.org/luci/common/tsmon/store"
	"go.chromium.org/luci/common/tsmon/target"
	"go.chromium.org/luci/common/tsmon/types"
	"go.chromium.org/luci/gae/filter/txndefer"
	"go.chromium.org/luci/gae/impl/cloud"
	"go.chromium.org/luci/gae/impl/memory"
	"go.chromium.org/luci/gae/service/datastore"
	"go.chromium.org/luci/gae/service/info"
	serverauth "go.chromium.org/luci/server/auth"
	"go.chromium.org/luci/server/auth/authtest"
	"go.chromium.org/luci/server/caching"
	"go.chromium.org/luci/server/secrets"
	"go.chromium.org/luci/server/tq"
	"go.chromium.org/luci/server/tq/tqtesting"
	_ "go.chromium.org/luci/server/tq/txn/datastore"

	migrationpb "go.chromium.org/luci/cv/api/migration"
	listenerpb "go.chromium.org/luci/cv/settings/listener"

	bbfake "go.chromium.org/luci/cv/internal/buildbucket/fake"
	"go.chromium.org/luci/cv/internal/common"
	"go.chromium.org/luci/cv/internal/common/bq"
	"go.chromium.org/luci/cv/internal/common/tree"
	"go.chromium.org/luci/cv/internal/common/tree/treetest"
	"go.chromium.org/luci/cv/internal/configs/srvcfg"
	"go.chromium.org/luci/cv/internal/gerrit"
	gf "go.chromium.org/luci/cv/internal/gerrit/gerritfake"

	. "github.com/smartystreets/goconvey/convey"
)

const gaeTopLevelDomain = ".appspot.com"

// TODO(tandrii): add fake config generation facilities.

// Test encapsulates typical setup for CV test.
//
// Typical use:
//
//	ct := cvtesting.Test{}
//	ctx, cancel := ct.SetUp()
//	defer cancel()
type Test struct {
	// Env simulates CV environment.
	Env *common.Env
	// GFake is a Gerrit fake. Defaults to an empty one.
	GFake *gf.Fake
	// BuildbucketFake is a Buildbucket fake. Defaults to an empty one.
	BuildbucketFake *bbfake.Fake
	// TreeFake is a fake Tree. Defaults to an open Tree.
	TreeFake *treetest.Fake
	// BQFake is a fake BQ client.
	BQFake *bq.Fake
	// TQDispatcher is a dispatcher with which task classes must be registered.
	//
	// Must not be set.
	TQDispatcher *tq.Dispatcher
	// TQ allows to run TQ tasks.
	TQ *tqtesting.Scheduler
	// SucceededTQTasks is a list of the TQ tasks that were executed successfully.
	SucceededTQTasks tqtesting.TaskList

	// Clock allows to move time forward.
	// By default, the time is moved automatically is something waits on it.
	Clock testclock.TestClock
	// TSMonStore store keeps all metrics in memory and allows examination.
	TSMonStore store.Store

	// MaxDuration limits how long a test can run as a fail safe.
	//
	// Defaults to 10s to most likely finish in pre/post submit tests,
	// with limited CPU resources.
	// Set to ~10ms when debugging a hung test.
	MaxDuration time.Duration

	// InitialMigrationSettings controls CQD -> CV migration.
	//
	// If you need to change them, do call UpdateMigrationConfig().
	InitialMigrationSettings *migrationpb.Settings

	// cleanupFuncs are executed in reverse order in cleanup().
	cleanupFuncs []func()

	// authDB is used to mock CrIA memberships.
	authDB *authtest.FakeDB
}

type testingContextKeyType struct{}

// IsTestingContext checks if the given context was derived from one created by
// cvtesting.Test.SetUp().
func IsTestingContext(ctx context.Context) bool {
	return ctx.Value(testingContextKeyType{}) != nil
}

func (t *Test) SetUp() (context.Context, func()) {
	if t.Env == nil {
		t.Env = &common.Env{
			LogicalHostname: "luci-change-verifier" + gaeTopLevelDomain,
			HTTPAddressBase: "https://luci-change-verifier" + gaeTopLevelDomain,
			GAEInfo: struct {
				CloudProject string
				ServiceName  string
				InstanceID   string
			}{
				CloudProject: "luci-change-verifier",
				ServiceName:  "test-service",
				InstanceID:   "test-instance",
			},
		}
	}

	t.setMaxDuration()
	ctxShared := context.WithValue(context.Background(), testingContextKeyType{}, struct{}{})
	// Don't set the deadline (timeout) into the context given to the test,
	// as it may interfere with test clock.
	ctx, cancel := context.WithCancel(ctxShared)
	ctxTimed, cancelTimed := context.WithTimeout(ctxShared, t.MaxDuration)
	go func(ctx context.Context) {
		// Instead, watch for expiry of ctxTimed and cancel test `ctx`.
		select {
		case <-ctxTimed.Done():
			cancel()
		case <-ctx.Done():
			// Normal test termination.
			cancelTimed()
		}
	}(ctx)
	t.cleanupFuncs = append(t.cleanupFuncs, func() {
		// Fail the test if the test has timed out.
		So(ctxTimed.Err(), ShouldBeNil)
		cancel()
		cancelTimed()
	})

	// setup the test clock first so that logger can use test clock timestamp.
	ctx = t.setUpTestClock(ctx)
	if testing.Verbose() {
		// TODO(crbug/1282023): make this logger emit testclock-based timestamps.
		ctx = logging.SetLevel(gologger.StdConfig.Use(ctx), logging.Debug)
	}
	// setup timerCallback after setup logger so that the logging in the
	// callback function can honor the verbose mode.
	ctx = t.setTestClockTimerCB(ctx)
	ctx = caching.WithEmptyProcessCache(ctx)
	ctx = secrets.GeneratePrimaryTinkAEADForTest(ctx)

	if t.TQDispatcher != nil {
		panic("TQDispatcher must not be set")
	}
	t.TQDispatcher = &tq.Dispatcher{}
	ctx, t.TQ = tq.TestingContext(ctx, t.TQDispatcher)
	t.TQ.TaskSucceeded = tqtesting.TasksCollector(&t.SucceededTQTasks)

	if t.GFake == nil {
		t.GFake = &gf.Fake{}
	}
	if t.BuildbucketFake == nil {
		t.BuildbucketFake = &bbfake.Fake{}
	}
	if t.TreeFake == nil {
		t.TreeFake = treetest.NewFake(ctx, tree.Open)
	}
	if t.BQFake == nil {
		t.BQFake = &bq.Fake{}
	}

	ctx = t.installDS(ctx)
	ctx = txndefer.FilterRDS(ctx)
	t.authDB = authtest.NewFakeDB()
	ctx = serverauth.WithState(ctx, &authtest.FakeState{FakeDB: t.authDB})

	ctx, _, _ = tsmon.WithFakes(ctx)
	t.TSMonStore = store.NewInMemory(&target.Task{})
	tsmon.GetState(ctx).SetStore(t.TSMonStore)

	if t.InitialMigrationSettings == nil {
		t.InitialMigrationSettings = &migrationpb.Settings{
			ApiHosts: []*migrationpb.Settings_ApiHost{{
				Host:          t.Env.LogicalHostname,
				Prod:          true,
				ProjectRegexp: []string{".+"},
			}},
			UseCvStartMessage: &migrationpb.Settings_UseCVStartMessage{
				ProjectRegexp: []string{".+"},
			},
		}
	}
	if err := srvcfg.SetTestMigrationConfig(ctx, proto.Clone(t.InitialMigrationSettings).(*migrationpb.Settings)); err != nil {
		panic(err)
	}
	if err := srvcfg.SetTestListenerConfig(ctx, &listenerpb.Settings{}, nil); err != nil {
		panic(err)
	}

	return ctx, t.cleanup
}

func (t *Test) cleanup() {
	for i := len(t.cleanupFuncs) - 1; i >= 0; i-- {
		t.cleanupFuncs[i]()
	}
}

func (t *Test) RoundTestClock(multiple time.Duration) {
	t.Clock.Set(t.Clock.Now().Add(multiple).Truncate(multiple))
}

func (t *Test) GFactory() gerrit.Factory {
	return gerrit.CachingFactory(16, gerrit.TimeLimitedFactory(gerrit.InstrumentedFactory(t.GFake)))
}

// TSMonSentValue returns the latest value of the given metric.
//
// If not set, returns nil.
func (t *Test) TSMonSentValue(ctx context.Context, m types.Metric, fieldVals ...interface{}) interface{} {
	resetTime := time.Time{}
	return t.TSMonStore.Get(ctx, m, resetTime, fieldVals)
}

// TSMonSentDistr returns the latest distr value of the given metric.
//
// If not set, returns nil.
// Panics if metric's value is not a distribution.
func (t *Test) TSMonSentDistr(ctx context.Context, m types.Metric, fieldVals ...interface{}) *distribution.Distribution {
	v := t.TSMonSentValue(ctx, m, fieldVals...)
	if v == nil {
		return nil
	}
	d, ok := v.(*distribution.Distribution)
	if !ok {
		panic(fmt.Errorf("metric %q value is not a %T, but %T", m.Info().Name, d, v))
	}
	return d
}

// UpdateMigrationConfig updates the config and due to its caching, pushes time
// forward to invalidate old cache.
func (t *Test) UpdateMigrationConfig(ctx context.Context, cfg *migrationpb.Settings) {
	if err := srvcfg.SetTestMigrationConfig(ctx, cfg); err != nil {
		panic(err)
	}
	t.Clock.Add(time.Minute)
}

func (t *Test) setMaxDuration() {
	// Can't use Go's test timeout because it is per TestXYZ func,
	// which typically instantiates & runs several `cvtesting.Test`s.
	switch s := os.Getenv("CV_TEST_TIMEOUT_SEC"); {
	case s != "":
		v, err := strconv.ParseInt(s, 10, 31)
		if err != nil {
			panic(err)
		}
		t.MaxDuration = time.Duration(v) * time.Second
	case t.MaxDuration != time.Duration(0):
		// TODO(tandrii): remove the possibility to override this per test in favor
		// of CV_TEST_TIMEOUT_SEC env var.
	case raceDetectionEnabled:
		t.MaxDuration = 90 * time.Second
	default:
		t.MaxDuration = 20 * time.Second
	}
}

// DisableProjectInGerritListener updates the cached config to disable LUCI
// projects matching a given regexp in Listener.
func (t *Test) DisableProjectInGerritListener(ctx context.Context, projectRE string) {
	cfg, err := srvcfg.GetListenerConfig(ctx, nil)
	if err != nil {
		panic(err)
	}
	existing := stringset.NewFromSlice(cfg.DisabledProjectRegexps...)
	existing.Add(projectRE)
	cfg.DisabledProjectRegexps = existing.ToSortedSlice()
	if err := srvcfg.SetTestListenerConfig(ctx, cfg, nil); err != nil {
		panic(err)
	}
}

func (t *Test) installDS(ctx context.Context) context.Context {
	if !strings.HasSuffix(t.Env.LogicalHostname, gaeTopLevelDomain) {
		panic(fmt.Errorf("Env.LogicalHostname %q doesn't end with %q", t.Env.LogicalHostname, gaeTopLevelDomain))
	}
	appID := t.Env.LogicalHostname[:len(t.Env.LogicalHostname)-len(gaeTopLevelDomain)]

	if ctx, ok := t.installDSReal(ctx); ok {
		return memory.UseInfo(ctx, appID)
	}
	if ctx, ok := t.installDSEmulator(ctx); ok {
		return memory.UseInfo(ctx, appID)
	}

	ctx = memory.UseWithAppID(ctx, appID)
	// CV runs against Firestore backend, which is consistent.
	datastore.GetTestable(ctx).Consistent(true)
	// Intentionally not enabling AutoIndex so that new code accidentally needing
	// a new index adds it both here (for the rest of CV tests to work, notably
	// e2e ones) and into appengine/index.yaml.
	datastore.GetTestable(ctx).AutoIndex(false)
	return ctx
}

// installDSProd configures CV tests to run with actual DS.
//
// If DATASTORE_PROJECT ENV var isn't set, returns false.
//
// To use, first
//
//	$ luci-auth context -- bash
//	$ export DATASTORE_PROJECT=my-cloud-project-with-datastore
//
// and then run go tests the usual way, e.g.:
//
//	$ go test ./...
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
//	$ gcloud beta emulators datastore start --consistency=1.0
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
	t.cleanupFuncs = append(t.cleanupFuncs, func() {
		if err := client.Close(); err != nil {
			logging.Errorf(ctx, "failed to close DS client: %s", err)
		}
	})
	ctx = (&cloud.ConfigLite{ProjectID: cloudProject, DS: client}).Use(ctx)
	maybeCleanupOldDSNamespaces(ctx)

	// Enter a namespace for this tests.
	ns := genDSNamespaceName(time.Now())
	logging.Debugf(ctx, "Using %q DS namespace", ns)
	ctx = info.MustNamespace(ctx, ns)
	// Failure to clear is hard before the test,
	// ignored after the test.
	So(clearDS(ctx), ShouldBeNil)
	t.cleanupFuncs = append(t.cleanupFuncs, func() {
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

// setUpTestClock simulates passage of time w/o idling CPU.
func (t *Test) setUpTestClock(ctx context.Context) context.Context {
	if t.Clock != nil {
		return clock.Set(ctx, t.Clock)
	}
	// Use a date-time that is easy to eyeball in logs.
	utc := time.Date(2020, time.February, 2, 10, 30, 00, 0, time.UTC)
	// But set it up in a clock as a local time to expose incorrect assumptions of UTC.
	now := time.Date(2020, time.February, 2, 13, 30, 00, 0, time.FixedZone("Fake local", 3*60*60))
	So(now.Equal(utc), ShouldBeTrue)
	ctx, t.Clock = testclock.UseTime(ctx, now)
	return ctx
}

// setTestClockTimerCB moves test time forward if something we recognize waits
// for it.
func (t *Test) setTestClockTimerCB(ctx context.Context) context.Context {
	// Testclock calls this callback every time something is waiting.
	// To avoid getting stuck tests, we need to move testclock forward by the
	// requested duration in most cases but not all.
	moveIf := stringset.NewFromSlice(
		// Used by tqtesting to wait until ETA of the next task.
		tqtesting.ClockTag,
		// Used to retry launching buildbucket tryjob.
		"retry-launch",
	)
	ignoreIf := stringset.NewFromSlice(
		// Used in clock.WithTimeout(ctx) | clock.WithDeadline(ctx).
		clock.ContextDeadlineTag,
		// Used by CQDFake to wait until the next loop.
		// NOTE: can't import cqdfake package const here due to circular import,
		// and while it's possible to refactor this, we expect to delete cqdfake
		// relatively soon.
		"cqdfake",
	)
	t.Clock.SetTimerCallback(func(dur time.Duration, timer clock.Timer) {
		tags := testclock.GetTags(timer)
		move, ignore := 0, 0
		for _, tag := range tags {
			switch {
			case moveIf.Has(tag):
				move++
			case ignoreIf.Has(tag):
				ignore++
			default:
				// Ignore by default, but log it to help fix the test if it gets stuck.
				logging.Warningf(ctx, "ignoring unexpected timer tag: %q. If test is stuck, add tag to `moveIf` above this log line", tag)
			}
		}
		// In ~all cases, there is exactly 1 tag, but be future proof.
		switch {
		case move > 0:
			logging.Debugf(ctx, "moving test clock %s by %s forward for %s", t.Clock.Now(), dur, tags)
			t.Clock.Add(dur)
		case ignore == 0:
			logging.Warningf(ctx, "ignoring timer without tags. If test is stuck, tag the waits via `clock` library")
		}
	})
	return ctx
}

// AddMember adds a given member into a given luci auth group.
//
// The email may omit domain. In that case, this method will add "@example.com"
// as the domain name.
func (t *Test) AddMember(email, group string) {
	if _, err := mail.ParseAddress(email); err != nil {
		email = fmt.Sprintf("%s@example.com", email)
	}
	id, err := identity.MakeIdentity(fmt.Sprintf("user:%s", email))
	if err != nil {
		panic(err)
	}
	t.authDB.AddMocks(authtest.MockMembership(id, group))
}

func (t *Test) ResetMockedAuthDB(ctx context.Context) {
	t.authDB = authtest.NewFakeDB()
	serverauth.GetState(ctx).(*authtest.FakeState).FakeDB = t.authDB
}
