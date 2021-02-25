// Copyright 2017 The LUCI Authors.
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

package engine

import (
	"bytes"
	"context"
	"fmt"
	"math"
	"strings"
	"time"

	"go.chromium.org/luci/gae/service/datastore"

	"go.chromium.org/luci/auth/identity"
	"go.chromium.org/luci/common/clock"
	"go.chromium.org/luci/common/data/rand/mathrand"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/common/retry/transient"
	"go.chromium.org/luci/common/tsmon/distribution"
	"go.chromium.org/luci/common/tsmon/field"
	"go.chromium.org/luci/common/tsmon/metric"
	"go.chromium.org/luci/common/tsmon/types"

	"go.chromium.org/luci/scheduler/appengine/internal"
	"go.chromium.org/luci/scheduler/appengine/task"
)

// errInvocationIDConflict is returned by generateInvocationID.
var errInvocationIDConflict = errors.New("could not find available invocationID", transient.Tag)

const (
	// debugLogSizeLimit is how many bytes the invocation debug log can be before
	// it gets trimmed. See 'trimDebugLog'. The debug log isn't supposed to be
	// huge.
	debugLogSizeLimit = 200000

	// debugLogTailLines is how many last log lines to keep when trimming the log.
	debugLogTailLines = 100
)

// Jan 1 2015, in UTC.
var invocationIDEpoch time.Time

func init() {
	var err error
	invocationIDEpoch, err = time.Parse(time.RFC822, "01 Jan 15 00:00 UTC")
	if err != nil {
		panic(err)
	}
}

var (
	// distributionBucketerSecToDay 1s to 1 day.
	//
	// 0.05037 is math.log10(math.exp(math.log(24*60*60)/98)), which means
	// the last bucket will contain overflow of everything >=1 day.
	distributionBucketerSecToDay = distribution.GeometricBucketer(math.Pow(10.0, 0.05037), 100)

	metricInvocationsDurations = metric.NewCumulativeDistribution(
		"luci/scheduler/invocations/durations",
		"Durations of completed invocations (sec).",
		&types.MetricMetadata{Units: types.Seconds},
		distributionBucketerSecToDay,
		field.String("jobID"),
		field.String("status"), // one of final statuses of task.Status enum.
	)

	metricInvocationsOverrun = metric.NewCounter(
		"luci/scheduler/invocations/overrun",
		"Number of invocations that were not run due to overrun of prior one.",
		nil,
		field.String("jobID"),
	)
)

// generateInvocationID is called within a transaction to pick a new Invocation
// ID and ensure it isn't taken yet.
//
// This function essentially pick root key for a new entity group, checking
// that it hasn't been taken yet.
//
// Format of the invocation ID:
//   - highest order bit set to 0 to keep the value positive.
//   - next 43 bits set to negated time since some predefined epoch, in ms.
//   - next 16 bits are generated by math.Rand
//   - next 4 bits set to 0. They indicate ID format.
//
// Makes one attempt at allocating an ID. If it fails (should be extremely
// rare), the entire transaction should be retried. We do it to avoid
// unnecessarily hitting multiple entity groups from a single transaction.
//
// Returns only transient errors.
func generateInvocationID(c context.Context) (int64, error) {
	// See http://play.golang.org/p/POpQzpT4Up.
	invTs := int64(clock.Now(c).UTC().Sub(invocationIDEpoch) / time.Millisecond)
	invTs = ^invTs & 8796093022207 // 0b111....1, 42 bits (clear highest bit)
	invTs = invTs << 20

	randSuffix := mathrand.Int63n(c, 65536)
	invID := invTs | (randSuffix << 4)
	exists, err := datastore.Exists(c, datastore.NewKey(c, "Invocation", "", invID, nil))
	if err != nil {
		return 0, transient.Tag.Apply(err)
	}
	if !exists.All() {
		return invID, nil
	}

	return 0, errInvocationIDConflict
}

// Invocation entity stores single invocation of a job (with perhaps multiple
// attempts due retries if the invocation fails to start).
//
// Root entity. ID is generated based on time by generateInvocationID()
// function.
type Invocation struct {
	_kind  string                `gae:"$kind,Invocation"`
	_extra datastore.PropertyMap `gae:"-,extra"`

	// ID is identifier of this particular attempt to run a job.
	ID int64 `gae:"$id"`

	// JobID is '<ProjectID>/<JobName>' string of a parent job.
	//
	// Set when the invocation is created and never changes.
	JobID string `gae:",noindex"`

	// IndexedJobID is '<ProjectID>/<JobName>' string of a parent job, but it is
	// set only for finished invocations.
	//
	// It is used to make the invocations appear in the listings of finished
	// invocations.
	//
	// We can't use JobID field for this since the invocation launch procedure can
	// potentially generate orphaned "garbage" invocations in some edge cases (if
	// Invocation transaction lands, but separate Job transaction doesn't). They
	// are harmless, but we don't want them to show up in listings.
	IndexedJobID string

	// RealmID is a global realm name (i.e. "<ProjectID>:...") the invocation
	// belongs to.
	//
	// It is copied from the Job entity when the invocation is created. May be
	// empty for old invocations.
	RealmID string `gae:",noindex"`

	// Started is time when this invocation was created.
	Started time.Time `gae:",noindex"`

	// Finished is time when this invocation transitioned to a terminal state.
	Finished time.Time `gae:",noindex"`

	// TriggeredBy is identity of whoever triggered the invocation, if it was
	// triggered via a single trigger submitted by some external user (not by the
	// service itself).
	//
	// Empty identity string if it was triggered by the service itself.
	TriggeredBy identity.Identity

	// PropertiesRaw is a blob with serialized task.Request.Properties supplied
	// when the invocation was created.
	//
	// Task managers use it to prepare the parameters for tasks.
	PropertiesRaw []byte `gae:",noindex"`

	// Tags is a sorted list of indexed "key:value" pairs supplied via
	// task.Request.Tags when the invocation was created.
	//
	// May be passed down the stack by task managers.
	Tags []string

	// IncomingTriggersRaw is a serialized list of triggers that the invocation
	// consumed.
	//
	// They are popped from job's pending triggers set when the invocation
	// starts.
	//
	// Use IncomingTriggers() function to grab them in deserialized form.
	IncomingTriggersRaw []byte `gae:",noindex"`

	// OutgoingTriggersRaw is a serialized list of triggers that the invocation
	// produced.
	//
	// They are fanned out into pending trigger sets of corresponding triggered
	// jobs (specified by TriggeredJobIDs).
	//
	// Use OutgoingTriggers() function to grab them in deserialized form.
	OutgoingTriggersRaw []byte `gae:",noindex"`

	// PendingTimersRaw is a serialized list of pending invocation timers.
	//
	// Timers are emitted by Controller's AddTimer call.
	//
	// Use PendingTimers() function to grab them in deserialized form.
	PendingTimersRaw []byte `gae:",noindex"`

	// Revision is revision number of config.cfg when this invocation was created.
	// For informational purpose.
	Revision string `gae:",noindex"`

	// RevisionURL is URL to human readable page with config file at
	// an appropriate revision. For informational purpose.
	RevisionURL string `gae:",noindex"`

	// Task is the job payload for this invocation in binary serialized form.
	// For informational purpose. See Catalog.UnmarshalTask().
	Task []byte `gae:",noindex"`

	// TriggeredJobIDs is a list of jobIDs of jobs which this job triggers.
	// The list is sorted and without duplicates.
	TriggeredJobIDs []string `gae:",noindex"`

	// DebugLog is short free form text log with debug messages.
	DebugLog string `gae:",noindex"`

	// RetryCount is 0 on a first attempt to launch the task. Increased with each
	// retry. For informational purposes.
	RetryCount int64 `gae:",noindex"`

	// Status is current status of the invocation (e.g. "RUNNING"), see the enum.
	Status task.Status

	// ViewURL is optional URL to a human readable page with task status, e.g.
	// Swarming task page. Populated by corresponding TaskManager.
	ViewURL string `gae:",noindex"`

	// TaskData is a storage where TaskManager can keep task-specific state
	// between calls.
	TaskData []byte `gae:",noindex"`

	// MutationsCount is used for simple compare-and-swap transaction control.
	//
	// It is incremented on each change to the entity.
	MutationsCount int64 `gae:",noindex"`
}

// isEqual returns true iff 'e' is equal to 'other'
func (e *Invocation) isEqual(other *Invocation) bool {
	return e == other || (e.ID == other.ID &&
		e.MutationsCount == other.MutationsCount && // compare it first, it changes most often
		e.JobID == other.JobID &&
		e.IndexedJobID == other.IndexedJobID &&
		e.Started.Equal(other.Started) &&
		e.Finished.Equal(other.Finished) &&
		e.TriggeredBy == other.TriggeredBy &&
		bytes.Equal(e.PropertiesRaw, other.PropertiesRaw) &&
		equalSortedLists(e.Tags, other.Tags) &&
		bytes.Equal(e.IncomingTriggersRaw, other.IncomingTriggersRaw) &&
		bytes.Equal(e.OutgoingTriggersRaw, other.OutgoingTriggersRaw) &&
		bytes.Equal(e.PendingTimersRaw, other.PendingTimersRaw) &&
		e.Revision == other.Revision &&
		e.RevisionURL == other.RevisionURL &&
		bytes.Equal(e.Task, other.Task) &&
		equalSortedLists(e.TriggeredJobIDs, other.TriggeredJobIDs) &&
		e.DebugLog == other.DebugLog &&
		e.RetryCount == other.RetryCount &&
		e.Status == other.Status &&
		e.ViewURL == other.ViewURL &&
		bytes.Equal(e.TaskData, other.TaskData))
}

// GetProjectID parses the ProjectID from the JobID and returns it.
func (e *Invocation) GetProjectID() string {
	parts := strings.Split(e.JobID, "/")
	return parts[0]
}

// debugLog appends a line to DebugLog field.
func (e *Invocation) debugLog(c context.Context, format string, args ...interface{}) {
	debugLog(c, &e.DebugLog, format, args...)
}

// trimDebugLog makes sure DebugLog field doesn't exceed limits.
//
// It cuts the middle of the log. We need to do this to keep the entity small
// enough to fit the datastore limits.
func (e *Invocation) trimDebugLog() {
	if len(e.DebugLog) <= debugLogSizeLimit {
		return
	}

	const cutMsg = "--- the log has been cut here ---"
	giveUp := func() {
		e.DebugLog = e.DebugLog[:debugLogSizeLimit-len(cutMsg)-2] + "\n" + cutMsg + "\n"
	}

	// We take last debugLogTailLines lines of log and move them "up", so that
	// the total log size is less than debugLogSizeLimit. We then put a line with
	// the message that some log lines have been cut. If these operations are not
	// possible (e.g. we have some giant lines or something), we give up and just
	// cut the end of the log.

	// Find debugLogTailLines-th "\n" from the end, e.DebugLog[tailStart:] is the
	// log tail.
	tailStart := len(e.DebugLog)
	for i := 0; i < debugLogTailLines; i++ {
		tailStart = strings.LastIndex(e.DebugLog[:tailStart-1], "\n")
		if tailStart <= 0 {
			giveUp()
			return
		}
	}
	tailStart++

	// Figure out how many bytes of head we can keep to make trimmed log small
	// enough.
	tailLen := len(e.DebugLog) - tailStart + len(cutMsg) + 1
	headSize := debugLogSizeLimit - tailLen
	if headSize <= 0 {
		giveUp()
		return
	}

	// Find last "\n" in the head.
	headEnd := strings.LastIndex(e.DebugLog[:headSize], "\n")
	if headEnd <= 0 {
		giveUp()
		return
	}

	// We want to keep 50 lines of the head no matter what.
	headLines := strings.Count(e.DebugLog[:headEnd], "\n")
	if headLines < 50 {
		giveUp()
		return
	}

	// Remove duplicated 'cutMsg' lines. They may appear if 'debugLog' (followed
	// by 'trimDebugLog') is called on already trimmed log multiple times.
	lines := strings.Split(e.DebugLog[:headEnd], "\n")
	lines = append(lines, cutMsg)
	lines = append(lines, strings.Split(e.DebugLog[tailStart:], "\n")...)
	trimmed := make([]byte, 0, debugLogSizeLimit)
	trimmed = append(trimmed, lines[0]...)
	for i := 1; i < len(lines); i++ {
		if !(lines[i-1] == cutMsg && lines[i] == cutMsg) {
			trimmed = append(trimmed, '\n')
			trimmed = append(trimmed, lines[i]...)
		}
	}
	e.DebugLog = string(trimmed)
}

// IncomingTriggers is a list of triggers that the invocation consumed.
//
// It is deserialized on the fly from IncomingTriggersRaw.
func (e *Invocation) IncomingTriggers() ([]*internal.Trigger, error) {
	return unmarshalTriggersList(e.IncomingTriggersRaw)
}

// OutgoingTriggers is a list of triggers that the invocation produced.
//
// It is deserialized on the fly from OutgoingTriggersRaw.
func (e *Invocation) OutgoingTriggers() ([]*internal.Trigger, error) {
	return unmarshalTriggersList(e.OutgoingTriggersRaw)
}

// PendingTimers is a list of not-yet-consumed invocation timers.
//
// It is deserialized on the fly from PendingTimersRaw.
func (e *Invocation) PendingTimers() ([]*internal.Timer, error) {
	return unmarshalTimersList(e.PendingTimersRaw)
}

// cleanupUnreferencedInvocations tries to delete given invocations.
//
// This is best effort cleanup after failures. It logs errors, but doesn't
// return them, to indicate that there's nothing we can actually do.
//
// 'invs' is allowed to have nils, they are skipped. Allowed to be called
// within a transaction, ignores it.
func cleanupUnreferencedInvocations(c context.Context, invs []*Invocation) {
	keysToKill := make([]*datastore.Key, 0, len(invs))
	for _, inv := range invs {
		if inv != nil {
			logging.Warningf(c, "Cleaning up inv %d of job %q", inv.ID, inv.JobID)
			keysToKill = append(keysToKill, datastore.KeyForObj(c, inv))
		}
	}
	if err := datastore.Delete(datastore.WithoutTransaction(c), keysToKill); err != nil {
		logging.WithError(err).Warningf(c, "Invocation cleanup failed")
	}
}

// reportOverrunMetrics reports overrun to monitoring.
// Should be called after transaction to save this invocation is completed.
func (e *Invocation) reportOverrunMetrics(c context.Context) {
	metricInvocationsOverrun.Add(c, 1, e.JobID)
}

// reportCompletionMetrics reports invocation stats to monitoring.
// Should be called after transaction to save this invocation is completed.
func (e *Invocation) reportCompletionMetrics(c context.Context) {
	if !e.Status.Final() || e.Finished.IsZero() {
		panic(fmt.Errorf("reportCompletionMetrics on incomplete invocation: %v", e))
	}
	duration := e.Finished.Sub(e.Started)
	metricInvocationsDurations.Add(c, duration.Seconds(), e.JobID, string(e.Status))
}
