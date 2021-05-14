// Copyright 2021 The LUCI Authors.
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

// Package bq provides functions for preparing a row to send to BigQuery upon
// completion of a Run.
package bq

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"

	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/common/retry/transient"
	"go.chromium.org/luci/gae/service/datastore"

	cvbqpb "go.chromium.org/luci/cv/api/bigquery/v1"
	cvbq "go.chromium.org/luci/cv/internal/bq"
	"go.chromium.org/luci/cv/internal/common"
	"go.chromium.org/luci/cv/internal/migration"
	"go.chromium.org/luci/cv/internal/run"
)

func SendRun(ctx context.Context, id common.RunID) error {
	a, err := makeAttempt(ctx, id)
	if err != nil {
		return errors.Annotate(err, "failed to make Attempt").Err()
	}

	// During the migration period when CQDaemon does most checks and triggers
	// builds, CV can't populate all of the fields of Attempt without the
	// information from CQDaemon; so for finished Attempts reported by
	// CQDaemon, we can fill in the remaining fields.
	switch cqda, err := fetchCQDAttempt(ctx, id); {
	case err != nil:
		return err
	case cqda != nil:
		a = reconcileAttempts(a, cqda)
	}

	// TODO(crbug/1173168): Change the destination table name after
	// CQDaemon stops sending rows; and sent to "commit-queue" project.
	// When we first start sending rows, we want to send to a separate table
	// name.
	return cvbq.SendRow(ctx, "raw", "attempts_cv", string(id), a)
}

func makeAttempt(ctx context.Context, id common.RunID) (*cvbqpb.Attempt, error) {
	r := run.Run{ID: id}
	switch err := datastore.Get(ctx, &r); {
	case err == datastore.ErrNoSuchEntity:
		return nil, errors.Reason("Run not found").Err()
	case err != nil:
		return nil, errors.Annotate(err, "failed to fetch Run").Tag(transient.Tag).Err()
	}
	if !run.IsEnded(r.Status) {
		panic("Run status must be final before sending to BQ.")
	}

	// Load CLs and convert them to GerritChanges including submit status.
	runCLs, err := run.LoadRunCLs(ctx, id, r.CLs)
	if err != nil {
		return nil, err
	}
	submittedSet := make(map[int64]struct{}, len(r.Submission.GetSubmittedCls()))
	for _, clid := range r.Submission.GetSubmittedCls() {
		submittedSet[clid] = struct{}{}
	}
	gerritChanges := make([]*cvbqpb.GerritChange, len(runCLs))
	for i, cl := range runCLs {
		gerritChanges[i] = toGerritChange(cl, submittedSet, r.Mode)
	}

	// TODO(crbug/1173168, crbug/1105669): We want to change the BQ
	// schema so that StartTime is processing start time and CreateTime is
	// trigger time.
	a := &cvbqpb.Attempt{
		Key:                  r.ID.AttemptKey(),
		LuciProject:          r.ID.LUCIProject(),
		ConfigGroup:          r.ConfigGroupID.Name(),
		ClGroupKey:           clGroupKey(runCLs),
		EquivalentClGroupKey: equivalentClGroupKey(runCLs),
		// Run.CreateTime is trigger time, which corresponds to what CQD sends for
		// StartTime.
		StartTime:     timestamppb.New(r.CreateTime),
		EndTime:       timestamppb.New(r.EndTime),
		GerritChanges: gerritChanges,
		// Builds, Substatus and HasCustomRequirement are not known to CV yet
		// during the migration state, so they should be filled in with Attempt
		// from CQD if possible.
		Builds: nil,
		Status: attemptStatus(ctx, &r),
		// TODO(crbug/1114686): Add a new FAILED_SUBMIT substatus, which
		// should be used in the case that some CLs failed to submit after
		// passing checks. (In this case, for backwards compatibility, we
		// will set status = SUCCESS, substatus = FAILED_SUBMIT.)
		Substatus: cvbqpb.AttemptSubstatus_NO_SUBSTATUS,
	}
	return a, nil
}

// toGerritChange creates a GerritChange for the given RunCL.
//
// This includes the submit status of the CL.
func toGerritChange(cl *run.RunCL, submitted map[int64]struct{}, mode run.Mode) *cvbqpb.GerritChange {
	detail := cl.Detail
	ci := detail.GetGerrit().GetInfo()
	gc := &cvbqpb.GerritChange{
		Host:                       detail.GetGerrit().Host,
		Project:                    ci.Project,
		Change:                     ci.Number,
		Patchset:                   int64(detail.Patchset),
		EarliestEquivalentPatchset: int64(detail.MinEquivalentPatchset),
		TriggerTime:                cl.Trigger.Time,
		Mode:                       mode.BQAttemptMode(),
		SubmitStatus:               cvbqpb.GerritChange_PENDING,
	}

	if mode == run.FullRun {
		// Mark the CL submit status as success if it appears in the submitted CLs
		// list, and failure if it does not.
		if _, ok := submitted[int64(cl.ID)]; ok {
			gc.SubmitStatus = cvbqpb.GerritChange_SUCCESS
		} else {
			gc.SubmitStatus = cvbqpb.GerritChange_FAILURE
		}
	}

	return gc
}

// fetchCQDAttempt fetches an Attempt from CQDaemon if available.
//
// Returns nil if no Attempt is available.
func fetchCQDAttempt(ctx context.Context, id common.RunID) (*cvbqpb.Attempt, error) {
	v := migration.VerifiedCQDRun{ID: id}
	switch err := datastore.Get(ctx, &v); {
	case err == datastore.ErrNoSuchEntity:
		// A Run may end without a VerifiedCQDRun stored if the Run is canceled.
		logging.Debugf(ctx, "no VerifiedCQDRun found for Run %q", id)
	case err != nil:
		return nil, errors.Annotate(err, "failed to fetch VerifiedCQDRun").Tag(transient.Tag).Err()
	}
	return v.Payload.GetRun().GetAttempt(), nil
}

// reconcileAttempts merges the CV Attempt and CQDaemon Attempt.
//
// Modifies and returns the CV Attempt.
//
// Once CV does the relevant work (keeping track of builds, reading the CL
// description footers, and performing checks) these will no longer have to be
// filled in with the CQDaemon Attempt values.
func reconcileAttempts(a, cqda *cvbqpb.Attempt) *cvbqpb.Attempt {
	// The list of Builds will be known to CV after it starts triggering
	// and tracking builds; until then CQD is the source of truth.
	a.Builds = cqda.Builds
	// Substatus generally indicates a failure reason, which is
	// known once one of the checks fails. CQDaemon may specify
	// a substatus in the case of abort (substatus: MANUAL_CANCEL)
	// or failure (FAILED_TRYJOBS etc.).
	if a.Status == cvbqpb.AttemptStatus_ABORTED || a.Status == cvbqpb.AttemptStatus_FAILURE {
		a.Status = cqda.Status
		a.Substatus = cqda.Substatus
	}
	a.Status = cqda.Status
	a.Substatus = cqda.Substatus
	// The HasCustomRequirement is determined by CL description footers.
	a.HasCustomRequirement = cqda.HasCustomRequirement
	return a
}

// attemptStatus converts a Run status to Attempt status.
func attemptStatus(ctx context.Context, r *run.Run) cvbqpb.AttemptStatus {
	switch r.Status {
	case run.Status_SUCCEEDED:
		return cvbqpb.AttemptStatus_SUCCESS
	case run.Status_FAILED:
		// In the case that the checks passed but not all CLs were submitted
		// successfully, the Attempt will still have status set to SUCCESS for
		// backwards compatibility. Note that r.Submission is expected to be
		// set only if a submission is attempted, meaning all checks passed.
		if r.Submission != nil && len(r.Submission.Cls) != len(r.Submission.SubmittedCls) {
			return cvbqpb.AttemptStatus_SUCCESS
		}
		return cvbqpb.AttemptStatus_FAILURE
	case run.Status_CANCELLED:
		return cvbqpb.AttemptStatus_ABORTED
	default:
		logging.Errorf(ctx, "Unexpected attempt status %q", r.Status)
		return cvbqpb.AttemptStatus_ATTEMPT_STATUS_UNSPECIFIED
	}
}

// clGroupKey constructs an opaque key unique to this set of CLs and patchsets.
func clGroupKey(cls []*run.RunCL) string {
	// TODO(crbug/1173168): Implement. In CQDaemon this is based on the sorted
	// CL triples (host, number, patchset). But it should be changed, e.g.
	// using sha2.
	// TODO(crbug/1090123): This should not be equal with equivalentClGroupKey.
	return ""
}

// equivalentClGroupKey constructs an opaque key unique to a set of CLs and
// their min equivalent patchsets.
func equivalentClGroupKey(cls []*run.RunCL) string {
	// TODO(crbug/1173168): Implement. In CQDaemon this is based on the sorted
	// CL triples (hostname, number, min equivalent patchset). But it should be
	// changed, e.g. using sha2.
	// TODO(crbug/1090123): This should not be equal with CL group key.
	return ""
}