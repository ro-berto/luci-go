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

package gerritfake

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	gerritutil "go.chromium.org/luci/common/api/gerrit"
	"go.chromium.org/luci/common/clock"
	"go.chromium.org/luci/common/clock/testclock"
	"go.chromium.org/luci/common/data/stringset"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	gerritpb "go.chromium.org/luci/common/proto/gerrit"

	"go.chromium.org/luci/cv/internal/gerrit"
)

// Client implements client for Fake Gerrit.
type Client struct {
	f           *Fake
	luciProject string // used in ACL checks.
	host        string
}

var _ gerrit.Client = (*Client)(nil)

///////////////////////////////////////////////////////////////////////////////
// Read RPCs

// Lists changes that match a query.
//
// Note, although the Gerrit API supports multiple queries, for which
// it can return multiple lists of changes, this is not a foreseen use-case
// so this API just includes one query with one returned list of changes.
//
// https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#list-changes
func (client *Client) ListChanges(ctx context.Context, in *gerritpb.ListChangesRequest, opts ...grpc.CallOption) (*gerritpb.ListChangesResponse, error) {
	client.f.recordRequest(in)
	if in.GetOffset() != 0 {
		return nil, status.New(codes.Unimplemented, "Offset is not supported by GerritFake").Err()
	}
	q, err := parseListChangesQuery(in.GetQuery())
	if err != nil {
		return nil, status.New(codes.InvalidArgument, err.Error()).Err()
	}
	client.f.m.Lock()
	defer client.f.m.Unlock()

	changes := make([]*gerritpb.ChangeInfo, 0, len(client.f.cs))
	for _, ch := range client.f.cs {
		switch {
		case ch.Host != client.host:
		case ch.ACLs(OpRead, client.luciProject).Code() != codes.OK:
		case !q.matches(ch):
		default:
			changes = append(changes, applyChangeOpts(ch, in.GetOptions()))
		}
	}
	// Sort from the most recently to least recently updated,
	// and if equal, deterministically disambiguate on change number to avoid
	// flaky tests.
	sort.Slice(changes, func(i, j int) bool {
		l := changes[i].GetUpdated().AsTime()
		r := changes[j].GetUpdated().AsTime()
		switch {
		case l.Before(r):
			return false
		case l.After(r):
			return true
		default:
			return changes[i].GetNumber() > changes[j].GetNumber()
		}
	})
	res := &gerritpb.ListChangesResponse{Changes: changes}
	if in.GetLimit() > 0 && int64(len(changes)) > in.GetLimit() {
		res.Changes = changes[:in.GetLimit()]
		res.MoreChanges = true
	}
	return res, nil
}

// Loads a change by id.
//
// https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#get-change
func (client *Client) GetChange(ctx context.Context, in *gerritpb.GetChangeRequest, opts ...grpc.CallOption) (*gerritpb.ChangeInfo, error) {
	client.f.m.Lock()
	defer client.f.m.Unlock()
	client.f.recordRequest(in)

	change, err := client.getChangeEnforceACLsLocked(in.GetNumber())
	if err != nil {
		return nil, err
	}
	return applyChangeOpts(change, in.GetOptions()), nil
}

// Retrieves related changes of a revision.
//
// Related changes are changes that either depend on, or are dependencies of
// the revision.
//
// https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#get-related-changes
func (client *Client) GetRelatedChanges(ctx context.Context, in *gerritpb.GetRelatedChangesRequest, opts ...grpc.CallOption) (*gerritpb.GetRelatedChangesResponse, error) {
	client.f.m.Lock()
	defer client.f.m.Unlock()
	client.f.recordRequest(in)

	change, err := client.getChangeEnforceACLsLocked(in.GetNumber())
	if err != nil {
		return nil, err
	}
	ps, _, err := change.resolveRevision(in.GetRevisionId())
	if err != nil {
		return nil, err
	}
	start := psKey(client.host, int(in.GetNumber()), ps)

	res := &gerritpb.GetRelatedChangesResponse{}
	added := stringset.New(10)
	add := func(psk string) error {
		if !added.Add(psk) {
			return nil
		}
		change, rev, ri, err := client.f.resolvePSKeyLocked(psk)
		if err != nil {
			return err
		}
		cc := &gerritpb.GetRelatedChangesResponse_ChangeAndCommit{
			Project:         change.Info.GetProject(),
			Number:          change.Info.GetNumber(),
			Patchset:        int64(ri.GetNumber()),
			CurrentPatchset: int64(change.Info.GetRevisions()[change.Info.GetCurrentRevision()].GetNumber()),
			Commit:          &gerritpb.CommitInfo{Id: rev},
		}
		for _, parentPSkey := range client.f.parentsOf[psk] {
			_, parentRev, _, err := client.f.resolvePSKeyLocked(parentPSkey)
			if err != nil {
				return err
			}
			cc.GetCommit().Parents = append(cc.GetCommit().Parents, &gerritpb.CommitInfo_Parent{Id: parentRev})
		}
		res.Changes = append(res.Changes, cc)
		return nil
	}
	// NOTE: Gerrit actually guarantees specific order. For simplicity, this fake
	// doesn't. We just recurse in direction of both child->parent and
	// parent->child and add visited changes to the list.
	if err := visitNodesDFS(start, client.f.childrenOf, add); err != nil {
		return nil, err
	}
	if err := visitNodesDFS(start, client.f.parentsOf, add); err != nil {
		return nil, err
	}
	if len(res.GetChanges()) == 1 {
		// Just the starting change itself, emulate Gerrit by returning empty list.
		res.Changes = nil
	}
	return res, nil
}

// Lists the files that were modified, added or deleted in a revision.
//
// https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#list-files
func (client *Client) ListFiles(ctx context.Context, in *gerritpb.ListFilesRequest, opts ...grpc.CallOption) (*gerritpb.ListFilesResponse, error) {
	client.f.m.Lock()
	defer client.f.m.Unlock()
	client.f.recordRequest(in)

	change, err := client.getChangeEnforceACLsLocked(in.GetNumber())
	if err != nil {
		return nil, err
	}
	_, ri, err := change.resolveRevision(in.GetRevisionId())
	if err != nil {
		return nil, err
	}
	// Note: for simplicity of fake, use files inside a revision, even though it
	// differs from what ListFiles may return for merge commit in Gerrit.
	ret := &gerritpb.ListFilesResponse{}
	// Deep copy before returning.
	proto.Merge(ret, &gerritpb.ListFilesResponse{Files: ri.GetFiles()})
	return ret, nil
}

///////////////////////////////////////////////////////////////////////////////
// Write RPCs

// Set various review bits on a change.
//
// Currently, only support following functionalities:
//  - Post Message.
//  - Set votes on a label (by project itself or on behalf of other user)
//
// https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#set-review
func (client *Client) SetReview(ctx context.Context, in *gerritpb.SetReviewRequest, opts ...grpc.CallOption) (*gerritpb.ReviewResult, error) {
	client.f.m.Lock()
	defer client.f.m.Unlock()
	client.f.recordRequest(in)

	ch, found := client.f.cs[key(client.host, int(in.GetNumber()))]
	if !found {
		return nil, status.Errorf(codes.NotFound, "change %s/%d not found", client.host, in.GetNumber())
	}
	if err := client.setReviewEnforceACLs(in, ch); err != nil {
		return nil, err
	}
	ch.Info.Updated = calcUpdatedTime(ctx, ch.Info.GetUpdated())
	now := clock.Now(ctx).UTC()
	if in.Message != "" {
		ch.Info.Messages = append(ch.Info.Messages, &gerritpb.ChangeMessageInfo{
			Id:      strconv.Itoa(len(ch.Info.Messages)),
			Author:  U(client.luciProject),
			Date:    timestamppb.New(now),
			Message: in.Message,
		})
	}

	if len(in.Labels) > 0 {
		for label, val := range in.Labels {
			if in.OnBehalfOf == 0 {
				Vote(label, int(val), now, U(client.luciProject))(ch.Info)
			} else {
				Vote(label, int(val), now, U(fmt.Sprintf("user-%d", in.OnBehalfOf)))(ch.Info)
			}
		}
	}

	return &gerritpb.ReviewResult{Labels: in.GetLabels()}, nil
}

// Submit a specific revision of a change.
//
// https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#submit-revision
func (client *Client) SubmitRevision(ctx context.Context, in *gerritpb.SubmitRevisionRequest, opts ...grpc.CallOption) (*gerritpb.SubmitInfo, error) {
	client.f.m.Lock()
	defer client.f.m.Unlock()
	client.f.recordRequest(in)

	ch, found := client.f.cs[key(client.host, int(in.GetNumber()))]
	if !found {
		return nil, status.Errorf(codes.NotFound, "change %s/%d not found", client.host, in.GetNumber())
	}
	if status := ch.ACLs(OpSubmit, client.luciProject); status.Code() != codes.OK {
		return nil, status.Err()
	}

	rev := in.GetRevisionId()
	if _, ok := ch.Info.GetRevisions()[rev]; !ok {
		return nil, status.Errorf(codes.NotFound, "revision %s not found", rev)
	}
	if rev != ch.Info.GetCurrentRevision() {
		return nil, status.Errorf(codes.Internal, "revision %s is not current revision", rev)
	}

	switch ch.Info.GetStatus() {
	case gerritpb.ChangeStatus_NEW:
		ch.Info.Status = gerritpb.ChangeStatus_MERGED
		// Most projects use a submit strategy which always creates a new patchset.
		// simulate the behavior here.
		PS(int(ch.Info.GetRevisions()[rev].GetNumber() + 1))(ch.Info)
		ch.Info.Updated = calcUpdatedTime(ctx, ch.Info.Updated)
		return &gerritpb.SubmitInfo{Status: gerritpb.ChangeStatus_MERGED}, nil
	case gerritpb.ChangeStatus_MERGED:
		return nil, status.Errorf(codes.Internal, "change is merged")
	case gerritpb.ChangeStatus_ABANDONED:
		return nil, status.Errorf(codes.Internal, "change is abandoned")
	default:
		panic(fmt.Errorf("unrecognized status %s", ch.Info.GetStatus()))
	}
}

///////////////////////////////////////////////////////////////////////////////
// Helper methods

// visitNodesDFS visits all nodes reachable from the current node via depth
// first search.
//
// Calls clbk for each node visited. If clbk returns error, visitNodesDFS aborts
// immediatey and returns the same error.
func visitNodesDFS(node string, edges map[string][]string, clbk func(node string) error) error {
	visited := stringset.New(1)

	var visit func(n string) error
	visit = func(n string) error {
		if !visited.Add(n) {
			return nil
		}
		for _, m := range edges[n] {
			if err := visit(m); err != nil {
				return err
			}
		}
		return clbk(n)
	}
	return visit(node)
}

func (client *Client) getChangeEnforceACLsLocked(change int64) (*Change, error) {
	ch, found := client.f.cs[key(client.host, int(change))]
	if !found {
		return nil, status.Errorf(codes.NotFound, "change %s/%d not found", client.host, change)
	}
	if status := ch.ACLs(OpRead, client.luciProject); status.Code() != codes.OK {
		return nil, status.Err()
	}
	return ch, nil
}

func (client *Client) setReviewEnforceACLs(in *gerritpb.SetReviewRequest, ch *Change) error {
	if in.Message != "" {
		if status := ch.ACLs(OpReview, client.luciProject); status.Code() != codes.OK {
			return status.Err()
		}
	}
	if len(in.Labels) > 0 {
		if in.OnBehalfOf == 0 {
			if status := ch.ACLs(OpReview, client.luciProject); status.Code() != codes.OK {
				return status.Err()
			}
		} else {
			if status := ch.ACLs(OpAlterVotesOfOthers, client.luciProject); status.Code() != codes.OK {
				return status.Err()
			}
		}
	}
	return nil
}

func applyChangeOpts(change *Change, opts []gerritpb.QueryOption) *gerritpb.ChangeInfo {
	qopts := make(map[gerritpb.QueryOption]struct{}, len(opts))
	for _, qopt := range opts {
		qopts[qopt] = struct{}{}
	}
	has := func(o gerritpb.QueryOption) bool {
		_, yes := qopts[o]
		return yes
	}

	// First, deep copy.
	ci := &gerritpb.ChangeInfo{}
	proto.Merge(ci, change.Info)

	// Second, mutate obeying query options.
	// TODO(tandrii): support more options as needed.
	switch {
	case has(gerritpb.QueryOption_ALL_REVISIONS):
		// Nothing to remove.
	case has(gerritpb.QueryOption_CURRENT_REVISION):
		// Remove all but current.
		for rev := range ci.GetRevisions() {
			if rev != ci.GetCurrentRevision() {
				delete(ci.GetRevisions(), rev)
			}
		}
	default:
		ci.CurrentRevision = "" // Yeah, weirdly, Gerrit doesn't set this unconditionally.
		ci.Revisions = nil
	}

	return ci
}

// calcUpdatedTime always push Updated time forward.
func calcUpdatedTime(ctx context.Context, curUpdatedTime *timestamppb.Timestamp) *timestamppb.Timestamp {
	now := clock.Now(ctx).UTC() // UTC is for easy to read logs
	switch u := curUpdatedTime.AsTime(); {
	case now.Before(u):
		panic(fmt.Errorf("clock's time [%s] is before the Updated time [%s]", now, u))
	case u.Equal(now):
		if tclock, ok := clock.Get(ctx).(testclock.TestClock); ok {
			logging.Debugf(ctx, "testclock.Time += 1second to ensure increasing Updated time")
			tclock.Add(time.Second)
			return timestamppb.New(tclock.Now())
		}
		panic(fmt.Errorf("clock's time [%s] is equal to the Updated time [%s] and not running in test", now, u))
	}
	return timestamppb.New(now)
}

type parsedListChangesQuery struct {
	after         time.Time
	before        time.Time
	status        gerritpb.ChangeStatus
	projectPrefix string
	projects      stringset.Set
	label         struct {
		name              string
		minValueExclusive int
	}
}

// parseListChangesQuery parses ListChangesRequest.Query for CV needs.
//
// It has lots of shortcomings:
//  * silently allows to repeat and overwrite prior instance of predicate,
//    e.g. "status:new status:merged" is treated as "status:merged".
//  * restricts (.. OR .. ) clauses only to project: predicates
//  * doesn't support OR without ()
//  * and many others.
//
// TODO(tandrii): this should be replaced by a proper library solution,
// perhaps the only implementing parsing & evaluation of https://aip.dev/160
// filtering proposal, which should suffice.
func parseListChangesQuery(query string) (p *parsedListChangesQuery, err error) {
	defer func() {
		if err != nil {
			err = errors.Annotate(err, "invalid query argument %q", query).Err()
			p = nil
		}
	}()

	mustUnquote := func(quoted string) string {
		l := len(quoted)
		if l <= 2 || quoted[0] != '"' || quoted[l-1] != '"' {
			err = errors.Reason("expected quoted string, but got %q", quoted).Err()
		}
		return quoted[1 : l-1]
	}
	inClause := false
	mustBeInClause := func(tok string) {
		if !inClause {
			err = errors.Reason("%q must be inside ()", tok).Err()
		}
	}
	mustBeOutClause := func(tok string) {
		if inClause {
			err = errors.Reason("%q must be outside of ()", tok).Err()
		}
	}

	p = &parsedListChangesQuery{}
	tokenizer := queryTokenizer{query}
	for {
		switch tok := tokenizer.next(); tok {
		case "":
			mustBeOutClause(tok)
			return
		case "(":
			mustBeOutClause(tok)
			inClause = true
		case ")":
			mustBeInClause(tok)
			inClause = false
		case "OR":
			mustBeInClause(tok)

		// TODO(tandrii): check for duplicate predicates here and below.
		case "project:":
			if p.projects.Len() > 0 {
				mustBeInClause(tok)
			} else {
				p.projects = stringset.New(1)
			}
			p.projects.Add(mustUnquote(tokenizer.next()))
		case "projects:":
			mustBeOutClause(tok)
			p.projectPrefix = mustUnquote(tokenizer.next())
		case "after:":
			mustBeOutClause(tok)
			// gerritutil.ParseTime checks quotes.
			p.after, err = gerritutil.ParseTime(tokenizer.next())
		case "before:":
			mustBeOutClause(tok)
			p.before, err = gerritutil.ParseTime(tokenizer.next())
		case "status:":
			mustBeOutClause(tok)
			tok = tokenizer.next()
			if v, ok := gerritpb.ChangeStatus_value[strings.ToUpper(tok)]; !ok {
				err = errors.Reason("unrecognized status %q", tok).Err()
			} else {
				p.status = gerritpb.ChangeStatus(v)
			}
		case "label:":
			tok = tokenizer.next()
			switch parts := strings.SplitN(tok, ">", 2); {
			case len(parts) != 2 || parts[0] == "" || parts[1] == "":
				err = errors.Reason("invalid label: %s", tok).Err()
			default:
				p.label.name = parts[0]
				p.label.minValueExclusive, err = strconv.Atoi(parts[1])
			}
		default:
			err = errors.Reason("unrecognized token %q", tok).Err()
		}
		if err != nil {
			return
		}
	}
}

func (p *parsedListChangesQuery) matches(c *Change) bool {
	switch {
	// after/before are inclusive in Gerrit.
	case !p.after.IsZero() && p.after.After(c.Info.GetUpdated().AsTime()):
	case !p.before.IsZero() && p.before.Before(c.Info.GetUpdated().AsTime()):
	case p.projects.Len() > 0 && !p.projects.Has(c.Info.GetProject()):
	case p.projectPrefix != "" && !strings.HasPrefix(c.Info.GetProject(), p.projectPrefix):
	case p.status != gerritpb.ChangeStatus_CHANGE_STATUS_INVALID && c.Info.GetStatus() != p.status:
	case !p.matchesLabel(c):
	default:
		return true
	}
	return false
}

func (p *parsedListChangesQuery) matchesLabel(c *Change) bool {
	switch li, exists := c.Info.GetLabels()[p.label.name]; {
	case p.label.name == "":
		return true
	case !exists:
		return false
	default:
		// In theory, we could use aggregated `li.GetValue()`, but this requires all
		// ChangeInfos to be faked correctly.
		for _, vote := range li.GetAll() {
			if vote.GetValue() > int32(p.label.minValueExclusive) {
				return true
			}
		}
		return false
	}
}

type queryTokenizer struct {
	remaining string
}

func (q *queryTokenizer) next() (token string) {
	consume := func(l int) {
		token, q.remaining = q.remaining[:l], q.remaining[l:]
	}

	q.remaining = strings.TrimLeft(q.remaining, " ")
	switch {
	case q.remaining == "":
	case q.remaining[0] == '(' || q.remaining[0] == ')':
		consume(1)
	case q.remaining[0] == '"':
		if endQuote := strings.IndexRune(q.remaining[1:], '"'); endQuote == -1 {
			// No matching closing ", so consume the rest of the string.
			consume(len(q.remaining))
		} else {
			consume(1 + endQuote + 1)
		}
	case len(q.remaining) == 1:
		consume(1)
	case q.remaining[:2] == "OR":
		consume(2)
	default:
		for i, c := range q.remaining {
			switch c {
			case ':':
				consume(i + 1)
				return
			case ' ':
				consume(i)
				return
			}
		}
		consume(len(q.remaining))
	}
	return
}
