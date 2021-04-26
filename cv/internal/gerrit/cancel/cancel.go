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

// Package cancel implements cancelling triggers of Run by removing CQ Votes
// on a CL.
package cancel

import (
	"context"
	"sort"
	"strings"
	"time"

	"google.golang.org/grpc/codes"

	"go.chromium.org/luci/common/clock"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	gerritpb "go.chromium.org/luci/common/proto/gerrit"
	"go.chromium.org/luci/common/retry/transient"
	"go.chromium.org/luci/gae/service/datastore"
	"go.chromium.org/luci/grpc/grpcutil"

	"go.chromium.org/luci/cv/internal/changelist"
	"go.chromium.org/luci/cv/internal/common"
	"go.chromium.org/luci/cv/internal/gerrit"
	"go.chromium.org/luci/cv/internal/gerrit/botdata"
	"go.chromium.org/luci/cv/internal/gerrit/trigger"
	"go.chromium.org/luci/cv/internal/lease"
	"go.chromium.org/luci/cv/internal/run"
)

// ErrPreconditionFailedTag is an error tag indicating that Cancel precondition
// failed.
var ErrPreconditionFailedTag = errors.BoolTag{
	Key: errors.NewTagKey("cancel precondition not met"),
}

// ErrPermanentTag is an error tag indicating that error occurs during the
// cancellation is permanent (e.g. lack of vote permission).
var ErrPermanentTag = errors.BoolTag{
	Key: errors.NewTagKey("permanent error while cancelling triggers"),
}

// Notify defines whom to notify for the cancellation.
//
// Note: REVIEWERS or VOTERS must be used together with OWNERS.
// TODO(yiwzhang): Remove this restriction if necessary.
type Notify int32

const (
	// NONE notifies no one.
	NONE Notify = 0
	// OWNER notifies the owner of the CL.
	OWNER Notify = 1
	// REVIEWERS notifies all reviewers of the CL.
	REVIEWERS Notify = 2
	// VOTERS notifies all users that have voted on CQ label when cancelling.
	VOTERS Notify = 4
)

func (notify Notify) validate() error {
	if (notify&VOTERS == VOTERS || notify&REVIEWERS == REVIEWERS) && notify&OWNER != OWNER {
		return errors.New("must notify OWNER when notifying REVIEWERS or VOTERS")
	}
	return nil
}

func (notify Notify) toGerritNotify(voters []*gerritpb.ApprovalInfo) (n gerritpb.Notify, nd *gerritpb.NotifyDetails) {
	n = gerritpb.Notify_NOTIFY_NONE
	if len(voters) > 0 && notify&VOTERS == VOTERS {
		accounts := make([]int64, len(voters))
		for i, v := range voters {
			accounts[i] = v.GetUser().GetAccountId()
		}
		sort.Slice(accounts, func(i, j int) bool {
			return accounts[i] < accounts[j]
		})
		nd = &gerritpb.NotifyDetails{
			Recipients: []*gerritpb.NotifyDetails_Recipient{
				{
					RecipientType: gerritpb.NotifyDetails_RECIPIENT_TYPE_TO,
					Info: &gerritpb.NotifyDetails_Info{
						Accounts: accounts,
					},
				},
			},
		}
	}
	switch {
	case notify&REVIEWERS == REVIEWERS:
		n = gerritpb.Notify_NOTIFY_OWNER_REVIEWERS
	case notify&OWNER == OWNER:
		n = gerritpb.Notify_NOTIFY_OWNER
	}
	return
}

// Input contains info to cancel triggers of Run on a CL.
type Input struct {
	// CL is a Gerrit CL entity.
	//
	// Must have CL.Snapshot set.
	CL *changelist.CL
	// Trigger identifies the triggering vote. Required.
	//
	// Removed only after all other votes on CQ label are removed.
	Trigger *run.Trigger
	// LUCIProject is the project that initiates this cancellation.
	//
	// The project scoped account of this LUCI project SHOULD have the permission
	// to set the CQ label on behalf of other users in Gerrit.
	LUCIProject string
	// Message to be posted along with the triggering vote removal
	Message string
	// Requester describes the caller (e.g. Project Manager, Run Manager).
	Requester string
	// Notify describes whom to notify regarding the cancellation.
	//
	// Example: OWNER|REVIEWERS|VOTERS
	Notify Notify
	// LeaseDuration is how long a lease will be held for this cancellation.
	//
	// If the passed context has a closer deadline, uses that deadline as lease
	// `ExpireTime`.
	LeaseDuration time.Duration
	// RunCLExternalIDs are IDs of all CLs involved in the Run.
	//
	// It will be included in `botdata.BotData` and posted to Gerrit as part of
	// the message in "unhappy path". See doc for `Cancel()`
	RunCLExternalIDs []changelist.ExternalID
}

// Cancel removes all votes on CQ label and posts the given message.
//
// If the patcheset of the passed CL is not current, returns error tagged with
// `ErrPreconditionFailedTag`.
//
// The triggering votes will be removed last and all other votes will be removed
// in chronological order (latest to earliest). Message will posted in the same
// rpc call to Gerrit that removes the triggering votes.
//
// Returns nil if all votes are successfully removed or no action was even
// necessary.
//
// If non-recoverable error occurs(e.g. lack of permission to remove votes),
// attempts to post the message including original message in the input,
// failure message and special `botdata.BotData` s.t. user is clearly
// communicated with the error and CV won't create new Run by parsing the
// `BotData` message. The returned error will be tagged with `ErrPermanentTag`.
func Cancel(ctx context.Context, in Input) error {
	switch {
	case in.CL.Snapshot == nil:
		panic("cl.Snapshot must be non-nil")
	case in.Trigger == nil:
		panic("trigger must be non-nil")
	case in.LUCIProject != in.CL.Snapshot.GetLuciProject():
		panic(errors.Reason("mismatched LUCI Project: got %q in input and %q in CL snapshot", in.LUCIProject, in.CL.Snapshot.GetLuciProject()).Err())
	}
	if err := in.Notify.validate(); err != nil {
		panic(err)
	}

	if err := ensurePSLatestInCV(ctx, in.CL); err != nil {
		return err
	}

	c := change{
		Host:        in.CL.Snapshot.GetGerrit().GetHost(),
		LUCIProject: in.LUCIProject,
		Project:     in.CL.Snapshot.GetGerrit().GetInfo().GetProject(),
		Number:      in.CL.Snapshot.GetGerrit().GetInfo().GetNumber(),
		Revision:    in.CL.Snapshot.GetGerrit().GetInfo().GetCurrentRevision(),
	}
	if err := c.initGerritClient(ctx); err != nil {
		return err
	}

	leaseCtx, close, err := applyLeaseForCL(ctx, in.CL.ID, in.LeaseDuration, in.Requester)
	if err != nil {
		return err
	}
	defer close()
	return cancelLeased(leaseCtx, c, in)
}

// TODO(tandrii): merge with prjmanager/purger's error messages.
var failMessage = "CV failed to unset the " + trigger.CQLabelName +
	" label on your behalf. Please unvote and revote on the " +
	trigger.CQLabelName + " label to retry."

func cancelLeased(ctx context.Context, c change, in Input) error {
	ci, err := c.getLatest(ctx)
	switch {
	case err != nil:
		return err
	case ci.GetUpdated().AsTime().Before(in.CL.Snapshot.GetGerrit().GetInfo().GetUpdated().AsTime()):
		return errors.Reason("got stale change info from gerrit for %s/%d", c.Host, c.Number).Tag(transient.Tag).Err()
	case ci.GetCurrentRevision() != c.Revision:
		return errors.Reason("failed to cancel because ps %d is not current for %s/%d", in.CL.Snapshot.GetPatchset(), c.Host, c.Number).Tag(ErrPreconditionFailedTag).Err()
	}

	removeErr := c.removeVotesAndPostMsg(ctx, ci, in.Trigger, in.Message, in.Notify)
	if removeErr == nil || !ErrPermanentTag.In(removeErr) {
		return removeErr
	}

	// Received permanent error, try posting message.
	var msgBuilder strings.Builder
	if in.Message != "" {
		msgBuilder.WriteString(in.Message)
		msgBuilder.WriteString("\n\n")
	}
	msgBuilder.WriteString(failMessage)
	if postMsgErr := c.postCancelMessage(ctx, ci, msgBuilder.String(), in.Trigger, in.RunCLExternalIDs, in.Notify); postMsgErr != nil {
		logging.WithError(postMsgErr).Errorf(ctx, "failed to post cancellation failure message to Gerrit")
	}
	return removeErr
}

func ensurePSLatestInCV(ctx context.Context, cl *changelist.CL) error {
	curCLInCV := &changelist.CL{ID: cl.ID}
	switch err := datastore.Get(ctx, curCLInCV); {
	case err == datastore.ErrNoSuchEntity:
		return errors.Reason("cl(id=%d) doesn't exist in datastore", cl.ID).Err()
	case err != nil:
		return errors.Annotate(err, "failed to load cl: %d", cl.ID).Tag(transient.Tag).Err()
	case curCLInCV.Snapshot.GetPatchset() > cl.Snapshot.GetPatchset():
		return errors.Reason("failed to cancel because ps %d is not current for cl(%d)", cl.Snapshot.GetPatchset(), cl.ID).Tag(ErrPreconditionFailedTag).Err()
	}
	return nil
}

func applyLeaseForCL(ctx context.Context, clid common.CLID, duration time.Duration, requester string) (context.Context, func(), error) {
	leaseExpireTime := clock.Now(ctx).UTC().Add(duration)
	if d, ok := ctx.Deadline(); ok && d.Before(leaseExpireTime) {
		leaseExpireTime = d // Honor the deadline imposed by context
	}
	leaseExpireTime = leaseExpireTime.Truncate(time.Millisecond)
	l, err := lease.Apply(ctx, lease.Application{
		ResourceID: lease.MakeCLResourceID(clid),
		Holder:     requester,
		ExpireTime: leaseExpireTime,
	})
	switch {
	case err == lease.ErrConflict:
		return nil, nil, errors.Annotate(err, "CL is currently being mutated").Tag(transient.Tag).Err()
	case err != nil:
		return nil, nil, err
	}

	dctx, cancel := context.WithDeadline(ctx, leaseExpireTime)
	close := func() {
		cancel()
		// Best-effort termination since lease will expire naturally.
		l.Terminate(ctx)
	}
	return dctx, close, nil
}

type change struct {
	Host        string
	LUCIProject string
	Project     string
	Number      int64
	Revision    string

	gc gerrit.Client
}

func (c *change) initGerritClient(ctx context.Context) (err error) {
	if c.gc == nil {
		c.gc, err = gerrit.CurrentClient(ctx, c.Host, c.LUCIProject)
	}
	return err
}

func (c *change) getLatest(ctx context.Context) (*gerritpb.ChangeInfo, error) {
	ci, err := c.gc.GetChange(ctx, &gerritpb.GetChangeRequest{
		Number:  c.Number,
		Project: c.Project,
		Options: []gerritpb.QueryOption{
			gerritpb.QueryOption_CURRENT_REVISION,
			gerritpb.QueryOption_DETAILED_LABELS,
			gerritpb.QueryOption_DETAILED_ACCOUNTS,
			gerritpb.QueryOption_MESSAGES,
		},
	})

	return ci, c.annotateGerritErr(ctx, err, "get")
}

func (c *change) removeVotesAndPostMsg(ctx context.Context, ci *gerritpb.ChangeInfo, t *run.Trigger, msg string, notify Notify) error {
	// TODO(tandrii): support cg.GetAdditionalModes().
	votes := ci.GetLabels()[trigger.CQLabelName].GetAll()
	sort.Slice(votes, func(i, j int) bool {
		return votes[i].GetDate().AsTime().After(votes[j].GetDate().AsTime())
	})

	errs := errors.NewLazyMultiError(len(votes))
	needRemoveTriggerVote := false
	for i, vote := range votes {
		switch accountID := vote.GetUser().GetAccountId(); {
		case vote.GetValue() == 0:
			// no-op
		case accountID == t.GetGerritAccountId():
			needRemoveTriggerVote = true
		default:
			errs.Assign(i, c.removeVote(ctx, accountID, "", gerritpb.Notify_NOTIFY_NONE, nil))
		}
	}

	switch n, nd := notify.toGerritNotify(votes); {
	case errs.Get() != nil:
		return common.MostSevereError(errs.Get())
	case !needRemoveTriggerVote:
		// No need to remove triggering votes, post message only.
		return c.annotateGerritErr(ctx, c.postGerritMsg(ctx, ci, msg, t, n, nd), "post message")
	default:
		return c.removeVote(ctx, t.GetGerritAccountId(), msg, n, nd)
	}
}

func (c *change) removeVote(ctx context.Context, accountID int64, msg string, n gerritpb.Notify, nd *gerritpb.NotifyDetails) error {
	_, err := c.gc.SetReview(ctx, &gerritpb.SetReviewRequest{
		Number:     c.Number,
		Project:    c.Project,
		RevisionId: c.Revision,
		Labels: map[string]int32{
			trigger.CQLabelName: 0,
		},
		Message: gerrit.TruncateMessage(msg),
		// TODO(yiwzhang): implement subtag like "autogenerated:cv~dry-run" if
		// we want to display more than one message from CV in Gerrit UI.
		// Gerrit will show the latest message for each unique tag.
		Tag:           "autogenerated:cv",
		Notify:        n,
		NotifyDetails: nd,
		OnBehalfOf:    accountID,
	})
	return c.annotateGerritErr(ctx, err, "remove vote")
}

func (c *change) postCancelMessage(ctx context.Context, ci *gerritpb.ChangeInfo, msg string, t *run.Trigger, runCLExternalIDs []changelist.ExternalID, notify Notify) (err error) {
	bd := botdata.BotData{
		Action:      botdata.Cancel,
		TriggeredAt: t.GetTime().AsTime(),
		Revision:    c.Revision,
		CLs:         make([]botdata.ChangeID, len(runCLExternalIDs)),
	}
	for i, eid := range runCLExternalIDs {
		bd.CLs[i].Host, bd.CLs[i].Number, err = eid.ParseGobID()
		if err != nil {
			return
		}
	}
	if msg, err = botdata.Append(msg, bd); err != nil {
		return err
	}
	n, nd := notify.toGerritNotify(ci.GetLabels()[trigger.CQLabelName].GetAll())
	if err := c.postGerritMsg(ctx, ci, msg, t, n, nd); err != nil {
		return c.annotateGerritErr(ctx, err, "post message")
	}
	return nil
}

// postGerritMsg posts the given message to Gerrit.
//
// Skips if duplicate message is found after triggering time.
func (c *change) postGerritMsg(ctx context.Context, ci *gerritpb.ChangeInfo, msg string, t *run.Trigger, n gerritpb.Notify, nd *gerritpb.NotifyDetails) (err error) {
	for _, m := range ci.GetMessages() {
		switch {
		case m.GetDate().AsTime().Before(t.GetTime().AsTime()):
		case strings.Contains(m.Message, strings.TrimSpace(msg)):
			return nil
		}
	}

	_, err = c.gc.SetReview(ctx, &gerritpb.SetReviewRequest{
		Number:        c.Number,
		Project:       c.Project,
		RevisionId:    ci.GetCurrentRevision(),
		Message:       gerrit.TruncateMessage(msg),
		Tag:           "autogenerated:cv",
		Notify:        n,
		NotifyDetails: nd,
	})
	return c.annotateGerritErr(ctx, err, "post message")
}

func (c *change) annotateGerritErr(ctx context.Context, err error, action string) error {
	switch grpcutil.Code(err) {
	case codes.OK:
		return nil
	case codes.PermissionDenied:
		return errors.Reason("no permission to %s %s/%d", action, c.Host, c.Number).Tag(ErrPermanentTag).Err()
	case codes.NotFound:
		return errors.Reason("change %s/%d not found", c.Host, c.Number).Tag(ErrPermanentTag).Err()
	default:
		return gerrit.UnhandledError(ctx, err, "failed to %s %s/%d", action, c.Host, c.Number)
	}
}
