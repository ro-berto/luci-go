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

package admin

import (
	"context"
	"fmt"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"time"

	"golang.org/x/sync/errgroup"

	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/common/retry/transient"
	"go.chromium.org/luci/common/sync/parallel"
	"go.chromium.org/luci/gae/service/datastore"
	"go.chromium.org/luci/grpc/appstatus"
	"go.chromium.org/luci/server/auth"
	"go.chromium.org/luci/server/tq"

	commonpb "go.chromium.org/luci/cv/api/common/v1"
	"go.chromium.org/luci/cv/internal/changelist"
	"go.chromium.org/luci/cv/internal/common"
	"go.chromium.org/luci/cv/internal/common/eventbox"
	"go.chromium.org/luci/cv/internal/gerrit/poller"
	"go.chromium.org/luci/cv/internal/gerrit/updater"
	"go.chromium.org/luci/cv/internal/prjmanager"
	"go.chromium.org/luci/cv/internal/prjmanager/prjpb"
	adminpb "go.chromium.org/luci/cv/internal/rpc/admin/api"
	"go.chromium.org/luci/cv/internal/rpc/pagination"
	"go.chromium.org/luci/cv/internal/run"
	"go.chromium.org/luci/cv/internal/run/eventpb"
)

// allowGroup is a Chrome Infra Auth group, members of which are allowed to call
// admin API. See https://crbug.com/1183616.
const allowGroup = "service-luci-change-verifier-admins"

type AdminServer struct {
	TQDispatcher  *tq.Dispatcher
	GerritUpdater *updater.Updater
	PMNotifier    *prjmanager.Notifier
	RunNotifier   *run.Notifier

	adminpb.UnimplementedAdminServer
}

func (d *AdminServer) GetProject(ctx context.Context, req *adminpb.GetProjectRequest) (resp *adminpb.GetProjectResponse, err error) {
	defer func() { err = appstatus.GRPCifyAndLog(ctx, err) }()
	if err = checkAllowed(ctx, "GetProject"); err != nil {
		return
	}
	if req.GetProject() == "" {
		return nil, appstatus.Error(codes.InvalidArgument, "project is required")
	}

	eg, ctx := errgroup.WithContext(ctx)

	var p *prjmanager.Project
	eg.Go(func() (err error) {
		p, err = prjmanager.Load(ctx, req.GetProject())
		return
	})

	resp = &adminpb.GetProjectResponse{}
	eg.Go(func() error {
		list, err := eventbox.List(ctx, prjmanager.EventboxRecipient(ctx, req.GetProject()))
		if err != nil {
			return errors.Annotate(err, "failed to fetch Project Events").Err()
		}
		events := make([]*prjpb.Event, len(list))
		for i, item := range list {
			events[i] = &prjpb.Event{}
			if err = proto.Unmarshal(item.Value, events[i]); err != nil {
				return errors.Annotate(err, "failed to unmarshal Event %q", item.ID).Err()
			}
		}
		resp.Events = events
		return nil
	})

	switch err = eg.Wait(); {
	case err != nil:
		return nil, err
	case p == nil:
		return nil, appstatus.Error(codes.NotFound, "project not found")
	default:
		resp.State = p.State
		resp.State.LuciProject = req.GetProject()
		return resp, nil
	}
}

func (d *AdminServer) GetProjectLogs(ctx context.Context, req *adminpb.GetProjectLogsRequest) (resp *adminpb.GetProjectLogsResponse, err error) {
	defer func() { err = appstatus.GRPCifyAndLog(ctx, err) }()
	if err = checkAllowed(ctx, "GetProjectLogs"); err != nil {
		return
	}
	switch {
	case req.GetPageToken() != "":
		return nil, appstatus.Error(codes.Unimplemented, "not implemented yet")
	case req.GetPageSize() < 0:
		return nil, appstatus.Error(codes.InvalidArgument, "negative page size not allowed")
	case req.GetProject() == "":
		return nil, appstatus.Error(codes.InvalidArgument, "project is required")
	case req.GetEversionMin() < 0:
		return nil, appstatus.Error(codes.InvalidArgument, "eversion_min must be non-negative")
	case req.GetEversionMax() < 0:
		return nil, appstatus.Error(codes.InvalidArgument, "eversion_max must be non-negative")
	}

	q := datastore.NewQuery(prjmanager.ProjectLogKind)
	q = q.Ancestor(datastore.MakeKey(ctx, prjmanager.ProjectKind, req.GetProject()))

	if m := req.GetEversionMin(); m > 0 {
		q.Gte("EVersion", m)
	}
	if m := req.GetEversionMax(); m > 0 {
		q.Lte("EVersion", m)
	}
	switch s := req.GetPageSize(); {
	case s > 1024:
		q.Limit(1024)
	case s > 0:
		q.Limit(s)
	default:
		q.Limit(128)
	}

	var out []*prjmanager.ProjectLog
	if err = datastore.GetAll(ctx, q, &out); err != nil {
		return nil, err
	}

	resp = &adminpb.GetProjectLogsResponse{Logs: make([]*adminpb.ProjectLog, len(out))}
	for i, l := range out {
		resp.Logs[i] = &adminpb.ProjectLog{
			Eversion:   l.EVersion,
			State:      l.State,
			UpdateTime: tspbNillable(l.UpdateTime),
			Reasons:    &prjpb.LogReasons{Reasons: l.Reasons},
		}
	}
	return resp, nil
}

func (d *AdminServer) GetRun(ctx context.Context, req *adminpb.GetRunRequest) (resp *adminpb.GetRunResponse, err error) {
	defer func() { err = appstatus.GRPCifyAndLog(ctx, err) }()
	if err = checkAllowed(ctx, "GetRun"); err != nil {
		return
	}
	if req.GetRun() == "" {
		return nil, appstatus.Error(codes.InvalidArgument, "run ID is required")
	}
	return loadRunAndEvents(ctx, common.RunID(req.GetRun()), nil)
}

func (d *AdminServer) GetCL(ctx context.Context, req *adminpb.GetCLRequest) (resp *adminpb.GetCLResponse, err error) {
	defer func() { err = appstatus.GRPCifyAndLog(ctx, err) }()
	if err = checkAllowed(ctx, "GetCL"); err != nil {
		return
	}

	cl, err := loadCL(ctx, req)
	if err != nil {
		return nil, err
	}
	runs := make([]string, len(cl.IncompleteRuns))
	for i, id := range cl.IncompleteRuns {
		runs[i] = string(id)
	}
	resp = &adminpb.GetCLResponse{
		Id:               int64(cl.ID),
		Eversion:         int64(cl.EVersion),
		ExternalId:       string(cl.ExternalID),
		UpdateTime:       timestamppb.New(cl.UpdateTime),
		Snapshot:         cl.Snapshot,
		ApplicableConfig: cl.ApplicableConfig,
		Access:           cl.Access,
		IncompleteRuns:   runs,
	}
	return resp, nil
}

func (d *AdminServer) GetPoller(ctx context.Context, req *adminpb.GetPollerRequest) (resp *adminpb.GetPollerResponse, err error) {
	defer func() { err = appstatus.GRPCifyAndLog(ctx, err) }()
	if err = checkAllowed(ctx, "GetPoller"); err != nil {
		return
	}
	if req.GetProject() == "" {
		return nil, appstatus.Error(codes.InvalidArgument, "project is required")
	}

	s := poller.State{LuciProject: req.GetProject()}
	switch err := datastore.Get(ctx, &s); {
	case err == datastore.ErrNoSuchEntity:
		return nil, appstatus.Error(codes.NotFound, "poller not found")
	case err != nil:
		return nil, errors.Annotate(err, "failed to fetch Poller state").Tag(transient.Tag).Err()
	}
	resp = &adminpb.GetPollerResponse{
		Project:     s.LuciProject,
		Eversion:    s.EVersion,
		ConfigHash:  s.ConfigHash,
		UpdateTime:  timestamppb.New(s.UpdateTime),
		QueryStates: s.QueryStates,
	}
	return resp, nil
}

func (d *AdminServer) SearchRuns(ctx context.Context, req *adminpb.SearchRunsRequest) (resp *adminpb.RunsResponse, err error) {
	defer func() { err = appstatus.GRPCifyAndLog(ctx, err) }()
	if err = checkAllowed(ctx, "SearchRuns"); err != nil {
		return
	}
	if req.PageSize, err = pagination.ValidatePageSize(req, 16, 128); err != nil {
		return nil, err
	}
	cursor := &run.Cursor{}
	if err := pagination.ValidatePageToken(req, cursor); err != nil {
		return nil, err
	}

	// Compute potentially interesting run keys using the most efficient query.
	var runKeys []*datastore.Key
	var cl *changelist.CL
	switch {
	case req.GetCl() != nil:
		cl, runKeys, err = searchRunsByCL(ctx, req, cursor)
	case req.GetProject() != "":
		runKeys, err = searchRunsByProject(ctx, req, cursor)
	default:
		return nil, appstatus.Error(codes.Unimplemented, "search across projects is not implemented")
	}
	if err != nil {
		return nil, err
	}

	var nextCursor *run.Cursor
	if l := len(runKeys); int32(l) == req.GetPageSize() {
		// For Admin API, it's OK to return StringID as is, as Admins can see any
		// LUCI project.
		nextCursor = &run.Cursor{Run: runKeys[l-1].StringID()}
	}

	// Fetch individual runs in parallel and apply final filtering.
	shouldSkip := func(r *run.Run) bool {
		if req.GetProject() != "" && req.GetProject() != r.ID.LUCIProject() {
			return true
		}
		if req.GetCl() != nil && !r.CLs.Contains(cl.ID) {
			return true
		}
		switch s := req.GetStatus(); s {
		case commonpb.Run_STATUS_UNSPECIFIED:
		case commonpb.Run_ENDED_MASK:
			if !run.IsEnded(r.Status) {
				return true
			}
		default:
			if s != r.Status {
				return true
			}
		}
		return false
	}
	runs := make([]*adminpb.GetRunResponse, len(runKeys))
	errs := parallel.WorkPool(16, func(work chan<- func() error) {
		for i, key := range runKeys {
			i, key := i, key
			work <- func() (err error) {
				runs[i], err = loadRunAndEvents(ctx, common.RunID(key.StringID()), shouldSkip)
				return
			}
		}
	})
	if errs != nil {
		return nil, common.MostSevereError(errs)
	}

	resp = &adminpb.RunsResponse{}

	// Remove nil runs, which were skipped above.
	resp.Runs = runs[:0]
	for _, r := range runs {
		if r != nil {
			resp.Runs = append(resp.Runs, r)
		}
	}

	resp.NextPageToken, err = pagination.TokenString(nextCursor)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// searchRunsByCL returns CL & Run IDs as Datastore keys, using CL to limit
// results.
func searchRunsByCL(ctx context.Context, req *adminpb.SearchRunsRequest, cursor *run.Cursor) (*changelist.CL, []*datastore.Key, error) {
	cl, err := loadCL(ctx, req.GetCl())
	if err != nil {
		return nil, nil, err
	}

	q := datastore.NewQuery(run.RunCLKind).Eq("IndexedID", cl.ID).Limit(req.GetPageSize()).KeysOnly(true)
	if excl := cursor.GetRun(); excl != "" {
		q = q.Gt("__key__", datastore.MakeKey(ctx, run.RunKind, excl, run.RunCLKind, int64(cl.ID)))
	}
	// TODO(tandrii): if req.GetProject() is given, restrict __key__ here.
	var runCLKeys []*datastore.Key
	if err := datastore.GetAll(ctx, q, &runCLKeys); err != nil {
		return nil, nil, errors.Annotate(err, "failed to fetch Runs IDs").Tag(transient.Tag).Err()
	}

	runKeys := make([]*datastore.Key, len(runCLKeys))
	for i, k := range runCLKeys {
		runKeys[i] = k.Parent()
	}
	return cl, runKeys, nil
}

// searchRunsByProject returns Run IDs as Datastore keys, using LUCI Project to
// limit results.
func searchRunsByProject(ctx context.Context, req *adminpb.SearchRunsRequest, cursor *run.Cursor) ([]*datastore.Key, error) {
	// Prepare queries.
	baseQ := run.NewQueryWithLUCIProject(ctx, req.GetProject()).Limit(req.GetPageSize()).KeysOnly(true)
	var queries []*datastore.Query
	switch s := req.GetStatus(); s {
	case commonpb.Run_STATUS_UNSPECIFIED:
		queries = append(queries, baseQ)
	case commonpb.Run_ENDED_MASK:
		for _, s := range []commonpb.Run_Status{commonpb.Run_SUCCEEDED, commonpb.Run_CANCELLED, commonpb.Run_FAILED} {
			queries = append(queries, baseQ.Eq("Status", s))
		}
	default:
		queries = append(queries, baseQ.Eq("Status", s))
	}
	if excl := cursor.GetRun(); excl != "" {
		for i, q := range queries {
			queries[i] = q.Gt("__key__", datastore.MakeKey(ctx, run.RunKind, excl))
		}
	}

	// Run all queries at once.
	var runKeys []*datastore.Key
	err := datastore.RunMulti(ctx, queries, func(k *datastore.Key) error {
		runKeys = append(runKeys, k)
		if len(runKeys) == int(req.GetPageSize()) {
			return datastore.Stop
		}
		return nil
	})
	if err != nil {
		return nil, errors.Annotate(err, "failed to fetch Runs").Tag(transient.Tag).Err()
	}
	return runKeys, nil
}

// Copy from dsset.
type itemEntity struct {
	_kind string `gae:"$kind,dsset.Item"`

	ID     string         `gae:"$id"`
	Parent *datastore.Key `gae:"$parent"`
	Value  []byte         `gae:",noindex"`
}

func (d *AdminServer) DeleteProjectEvents(ctx context.Context, req *adminpb.DeleteProjectEventsRequest) (resp *adminpb.DeleteProjectEventsResponse, err error) {
	defer func() { err = appstatus.GRPCifyAndLog(ctx, err) }()
	if err = checkAllowed(ctx, "DeleteProjectEvents"); err != nil {
		return
	}

	switch {
	case req.GetProject() == "":
		return nil, appstatus.Error(codes.InvalidArgument, "project is required")
	case req.GetLimit() <= 0:
		return nil, appstatus.Error(codes.InvalidArgument, "limit must be >0")
	}

	parent := datastore.MakeKey(ctx, prjmanager.ProjectKind, req.GetProject())
	q := datastore.NewQuery("dsset.Item").Ancestor(parent).Limit(req.GetLimit())
	var entities []*itemEntity
	if err := datastore.GetAll(ctx, q, &entities); err != nil {
		return nil, errors.Annotate(err, "failed to fetch up to %d events", req.GetLimit()).Tag(transient.Tag).Err()
	}

	stats := make(map[string]int64, 10)
	for _, e := range entities {
		pb := &prjpb.Event{}
		if err := proto.Unmarshal(e.Value, pb); err != nil {
			stats["<unknown>"]++
		} else {
			stats[fmt.Sprintf("%T", pb.GetEvent())]++
		}
	}
	if err := datastore.Delete(ctx, entities); err != nil {
		return nil, errors.Annotate(err, "failed to delete %d events", len(entities)).Tag(transient.Tag).Err()
	}
	return &adminpb.DeleteProjectEventsResponse{Events: stats}, nil
}

func (d *AdminServer) RefreshProjectCLs(ctx context.Context, req *adminpb.RefreshProjectCLsRequest) (resp *adminpb.RefreshProjectCLsResponse, err error) {
	defer func() { err = appstatus.GRPCifyAndLog(ctx, err) }()
	if err = checkAllowed(ctx, "RefreshProjectCLs"); err != nil {
		return
	}
	if req.GetProject() == "" {
		return nil, appstatus.Error(codes.InvalidArgument, "project is required")
	}

	p, err := prjmanager.Load(ctx, req.GetProject())
	if err != nil {
		return nil, errors.Annotate(err, "failed to fetch Project %q", req.GetProject()).Tag(transient.Tag).Err()
	}

	cls := make([]*changelist.CL, len(p.State.GetPcls()))
	errs := parallel.WorkPool(20, func(work chan<- func() error) {
		for i, pcl := range p.State.GetPcls() {
			i := i
			id := pcl.GetClid()
			work <- func() error {
				// Load individual CL to avoid OOMs.
				cl := changelist.CL{ID: common.CLID(id)}
				if err := datastore.Get(ctx, &cl); err != nil {
					return errors.Annotate(err, "failed to fetch CL %d", id).Tag(transient.Tag).Err()
				}
				cls[i] = &changelist.CL{ID: cl.ID, EVersion: cl.EVersion}

				host, change, err := cl.ExternalID.ParseGobID()
				if err != nil {
					return err
				}
				payload := &updater.RefreshGerritCL{
					LuciProject: req.GetProject(),
					Host:        host,
					Change:      change,
					ClidHint:    id,
				}
				return d.GerritUpdater.Schedule(ctx, payload)
			}
		}
	})
	if err := common.MostSevereError(errs); err != nil {
		return nil, err
	}

	if err := d.PMNotifier.NotifyCLsUpdated(ctx, req.GetProject(), changelist.ToUpdatedEvents(cls...)); err != nil {
		return nil, err
	}

	clvs := make(map[int64]int64, len(p.State.GetPcls()))
	for _, cl := range cls {
		clvs[int64(cl.ID)] = int64(cl.EVersion)
	}
	return &adminpb.RefreshProjectCLsResponse{ClVersions: clvs}, nil
}

func (d *AdminServer) SendProjectEvent(ctx context.Context, req *adminpb.SendProjectEventRequest) (_ *emptypb.Empty, err error) {
	defer func() { err = appstatus.GRPCifyAndLog(ctx, err) }()
	if err = checkAllowed(ctx, "SendProjectEvent"); err != nil {
		return
	}
	switch {
	case req.GetProject() == "":
		return nil, appstatus.Error(codes.InvalidArgument, "project is required")
	case req.GetEvent().GetEvent() == nil:
		return nil, appstatus.Error(codes.InvalidArgument, "event with a specific inner event is required")
	}

	switch p, err := prjmanager.Load(ctx, req.GetProject()); {
	case err != nil:
		return nil, errors.Annotate(err, "failed to fetch Project").Err()
	case p == nil:
		return nil, appstatus.Error(codes.NotFound, "project not found")
	}

	if err := d.PMNotifier.SendNow(ctx, req.GetProject(), req.GetEvent()); err != nil {
		return nil, errors.Annotate(err, "failed to send event").Err()
	}
	return &emptypb.Empty{}, nil
}

func (d *AdminServer) SendRunEvent(ctx context.Context, req *adminpb.SendRunEventRequest) (_ *emptypb.Empty, err error) {
	defer func() { err = appstatus.GRPCifyAndLog(ctx, err) }()
	if err = checkAllowed(ctx, "SendRunEvent"); err != nil {
		return
	}
	switch {
	case req.GetRun() == "":
		return nil, appstatus.Error(codes.InvalidArgument, "Run is required")
	case req.GetEvent().GetEvent() == nil:
		return nil, appstatus.Error(codes.InvalidArgument, "event with a specific inner event is required")
	}

	switch err := datastore.Get(ctx, &run.Run{ID: common.RunID(req.GetRun())}); {
	case err == datastore.ErrNoSuchEntity:
		return nil, appstatus.Error(codes.NotFound, "Run not found")
	case err != nil:
		return nil, errors.Annotate(err, "failed to fetch Run").Tag(transient.Tag).Err()
	}

	if err := d.RunNotifier.SendNow(ctx, common.RunID(req.GetRun()), req.GetEvent()); err != nil {
		return nil, errors.Annotate(err, "failed to send event").Err()
	}
	return &emptypb.Empty{}, nil
}

func (d *AdminServer) ScheduleTask(ctx context.Context, req *adminpb.ScheduleTaskRequest) (_ *emptypb.Empty, err error) {
	defer func() { err = appstatus.GRPCifyAndLog(ctx, err) }()
	if err = checkAllowed(ctx, "ScheduleTask"); err != nil {
		return
	}

	const trans = true
	var possiblePayloads = []struct {
		inTransaction bool
		payload       proto.Message
	}{
		{trans, req.GetBatchRefreshGerritCl()},
		{false, req.GetExportRunToBq()},
		{trans, req.GetKickManageProject()},
		{trans, req.GetKickManageRun()},
		{false, req.GetManageProject()},
		{false, req.GetManageRun()},
		{false, req.GetPollGerrit()},
		{trans, req.GetPurgeCl()},
		{false, req.GetRefreshGerritCl()},
		{false, req.GetRefreshProjectConfig()},
	}

	chosen := possiblePayloads[0]
	for _, another := range possiblePayloads[1:] {
		switch {
		case reflect.ValueOf(another.payload).IsNil():
		case !reflect.ValueOf(chosen.payload).IsNil():
			return nil, appstatus.Error(codes.InvalidArgument, "exactly one task payload required, but 2+ given")
		default:
			chosen = another
		}
	}

	if reflect.ValueOf(chosen.payload).IsNil() {
		return nil, appstatus.Error(codes.InvalidArgument, "exactly one task payload required, but none given")
	}
	kind := chosen.payload.ProtoReflect().Type().Descriptor().Name()
	if chosen.inTransaction && req.GetDeduplicationKey() != "" {
		return nil, appstatus.Errorf(codes.InvalidArgument, "task %q is transactional, so the deduplication_key is not allowed", kind)
	}

	t := &tq.Task{
		Payload:          chosen.payload,
		DeduplicationKey: req.GetDeduplicationKey(),
		Title:            fmt.Sprintf("admin/%s/%s/%s", auth.CurrentIdentity(ctx), kind, req.GetDeduplicationKey()),
	}

	if chosen.inTransaction {
		err = datastore.RunInTransaction(ctx, func(ctx context.Context) error {
			return d.TQDispatcher.AddTask(ctx, t)
		}, nil)
	} else {
		err = d.TQDispatcher.AddTask(ctx, t)
	}

	if err != nil {
		return nil, errors.Annotate(err, "failed to schedule task").Err()
	}
	return &emptypb.Empty{}, nil
}

func checkAllowed(ctx context.Context, name string) error {
	switch yes, err := auth.IsMember(ctx, allowGroup); {
	case err != nil:
		return errors.Annotate(err, "failed to check ACL").Err()
	case !yes:
		return appstatus.Errorf(codes.PermissionDenied, "not a member of %s", allowGroup)
	default:
		logging.Warningf(ctx, "%s is calling admin.%s", auth.CurrentIdentity(ctx), name)
		return nil
	}
}

var regexCrRevPath = regexp.MustCompile(`/([ci])/(\d+)(/(\d+))?`)
var regexGoB = regexp.MustCompile(`((\w+-)+review\.googlesource\.com)/(#/)?(c/)?(([^\+]+)/\+/)?(\d+)(/(\d+)?)?`)

func parseGerritURL(s string) (changelist.ExternalID, error) {
	u, err := url.Parse(s)
	if err != nil {
		return "", err
	}
	var host string
	var change int64
	if u.Host == "crrev.com" {
		m := regexCrRevPath.FindStringSubmatch(u.Path)
		if m == nil {
			return "", errors.New("invalid crrev.com URL")
		}
		switch m[1] {
		case "c":
			host = "chromium-review.googlesource.com"
		case "i":
			host = "chrome-internal-review.googlesource.com"
		default:
			panic("impossible")
		}
		if change, err = strconv.ParseInt(m[2], 10, 64); err != nil {
			return "", errors.Reason("invalid crrev.com URL change number /%s/", m[2]).Err()
		}
	} else {
		m := regexGoB.FindStringSubmatch(s)
		if m == nil {
			return "", errors.Reason("Gerrit URL didn't match regexp %q", regexGoB.String()).Err()
		}
		if host = m[1]; host == "" {
			return "", errors.New("invalid Gerrit host")
		}
		if change, err = strconv.ParseInt(m[7], 10, 64); err != nil {
			return "", errors.Reason("invalid Gerrit URL change number /%s/", m[7]).Err()
		}
	}
	return changelist.GobID(host, change)
}

func loadRunAndEvents(ctx context.Context, rid common.RunID, shouldSkip func(r *run.Run) bool) (*adminpb.GetRunResponse, error) {
	r := &run.Run{ID: rid}
	switch err := datastore.Get(ctx, r); {
	case err == datastore.ErrNoSuchEntity:
		return nil, appstatus.Error(codes.NotFound, "run not found")
	case err != nil:
		return nil, errors.Annotate(err, "failed to fetch Run").Tag(transient.Tag).Err()
	case shouldSkip != nil && shouldSkip(r):
		return nil, nil
	}

	eg, ctx := errgroup.WithContext(ctx)
	var externalIDs []string
	eg.Go(func() error {
		switch rcls, err := run.LoadRunCLs(ctx, r.ID, r.CLs); {
		case err != nil:
			return errors.Annotate(err, "failed to fetch RunCLs").Err()
		default:
			externalIDs = make([]string, len(rcls))
			for i, rcl := range rcls {
				externalIDs[i] = string(rcl.ExternalID)
			}
			return nil
		}
	})

	var logEntries []*run.LogEntry
	eg.Go(func() error {
		var err error
		logEntries, err = run.LoadRunLogEntries(ctx, r.ID)
		if err != nil {
			return errors.Annotate(err, "failed to fetch RunCLs").Err()
		}
		return nil
	})

	var events []*eventpb.Event
	eg.Go(func() error {
		list, err := eventbox.List(ctx, run.EventboxRecipient(ctx, rid))
		if err != nil {
			return errors.Annotate(err, "failed to fetch Run Events").Err()
		}
		events = make([]*eventpb.Event, len(list))
		for i, item := range list {
			events[i] = &eventpb.Event{}
			if err = proto.Unmarshal(item.Value, events[i]); err != nil {
				return errors.Annotate(err, "failed to unmarshal Event %q", item.ID).Err()
			}
		}
		return nil
	})

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	return &adminpb.GetRunResponse{
		Id:               string(rid),
		Eversion:         int64(r.EVersion),
		Mode:             string(r.Mode),
		Status:           r.Status,
		CreateTime:       tspbNillable(r.CreateTime),
		StartTime:        tspbNillable(r.StartTime),
		UpdateTime:       tspbNillable(r.UpdateTime),
		EndTime:          tspbNillable(r.EndTime),
		Owner:            string(r.Owner),
		ConfigGroupId:    string(r.ConfigGroupID),
		Cls:              common.CLIDsAsInt64s(r.CLs),
		ExternalCls:      externalIDs,
		Options:          r.Options,
		Tryjobs:          r.Tryjobs,
		Submission:       r.Submission,
		FinalizedByCqd:   r.FinalizedByCQD,
		LatestClsRefresh: tspbNillable(r.LatestCLsRefresh),

		LogEntries: logEntries,
		Events:     events,
	}, nil
}

func loadCL(ctx context.Context, req *adminpb.GetCLRequest) (*changelist.CL, error) {
	var err error
	var cl *changelist.CL
	var eid changelist.ExternalID
	switch {
	case req.GetId() != 0:
		cl = &changelist.CL{ID: common.CLID(req.GetId())}
		err = datastore.Get(ctx, cl)
	case req.GetExternalId() != "":
		eid = changelist.ExternalID(req.GetExternalId())
		cl, err = eid.Get(ctx)
	case req.GetGerritUrl() != "":
		eid, err = parseGerritURL(req.GetGerritUrl())
		if err != nil {
			return nil, appstatus.Errorf(codes.InvalidArgument, "invalid Gerrit URL %q: %s", req.GetGerritUrl(), err)
		}
		cl, err = eid.Get(ctx)
	default:
		return nil, appstatus.Error(codes.InvalidArgument, "id or external_id or gerrit_url is required")
	}

	switch {
	case err == datastore.ErrNoSuchEntity:
		if req.GetId() == 0 {
			return nil, appstatus.Errorf(codes.NotFound, "CL %d not found", req.GetId())
		}
		return nil, appstatus.Errorf(codes.NotFound, "CL %s not found", eid)
	case err != nil:
		return nil, err
	}
	return cl, nil
}

func tspbNillable(t time.Time) *timestamppb.Timestamp {
	if t.IsZero() {
		return nil
	}
	return timestamppb.New(t)
}
