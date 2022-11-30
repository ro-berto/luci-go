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

package rpc

import (
	"context"
	"fmt"
	"regexp"
	"sort"
	"strings"
	"time"

	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"

	cipdCommon "go.chromium.org/luci/cipd/common"
	"go.chromium.org/luci/common/clock"
	"go.chromium.org/luci/common/data/stringset"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/common/proto/protowalk"
	"go.chromium.org/luci/common/sync/parallel"
	"go.chromium.org/luci/gae/service/datastore"
	"go.chromium.org/luci/gae/service/info"
	"go.chromium.org/luci/grpc/appstatus"
	"go.chromium.org/luci/server/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/timestamppb"

	bb "go.chromium.org/luci/buildbucket"
	"go.chromium.org/luci/buildbucket/appengine/internal/buildid"
	"go.chromium.org/luci/buildbucket/appengine/internal/config"
	"go.chromium.org/luci/buildbucket/appengine/internal/metrics"
	"go.chromium.org/luci/buildbucket/appengine/internal/perm"
	"go.chromium.org/luci/buildbucket/appengine/internal/resultdb"
	"go.chromium.org/luci/buildbucket/appengine/internal/search"
	"go.chromium.org/luci/buildbucket/appengine/model"
	"go.chromium.org/luci/buildbucket/appengine/tasks"
	taskdefs "go.chromium.org/luci/buildbucket/appengine/tasks/defs"
	"go.chromium.org/luci/buildbucket/bbperms"
	pb "go.chromium.org/luci/buildbucket/proto"
	"go.chromium.org/luci/buildbucket/protoutil"
)

var casInstanceRe = regexp.MustCompile(`^projects/[^/]*/instances/[^/]*$`)

type CreateBuildChecker struct{}

var _ protowalk.FieldProcessor = (*CreateBuildChecker)(nil)

func (*CreateBuildChecker) Process(field protoreflect.FieldDescriptor, msg protoreflect.Message) (data protowalk.ResultData, applied bool) {
	cbfb := proto.GetExtension(field.Options().(*descriptorpb.FieldOptions), pb.E_CreateBuildFieldOption).(*pb.CreateBuildFieldOption)
	switch cbfb.FieldBehavior {
	case annotations.FieldBehavior_OUTPUT_ONLY:
		msg.Clear(field)
		return protowalk.ResultData{Message: "cleared OUTPUT_ONLY field"}, true
	case annotations.FieldBehavior_REQUIRED:
		return protowalk.ResultData{Message: "required", IsErr: true}, true
	default:
		panic("unsupported field behavior")
	}
}

func init() {
	protowalk.RegisterFieldProcessor(&CreateBuildChecker{}, func(field protoreflect.FieldDescriptor) protowalk.ProcessAttr {
		if fo := field.Options().(*descriptorpb.FieldOptions); fo != nil {
			if cbfb := proto.GetExtension(fo, pb.E_CreateBuildFieldOption).(*pb.CreateBuildFieldOption); cbfb != nil {
				switch cbfb.FieldBehavior {
				case annotations.FieldBehavior_OUTPUT_ONLY:
					return protowalk.ProcessIfSet
				case annotations.FieldBehavior_REQUIRED:
					return protowalk.ProcessIfUnset
				default:
					panic("unsupported field behavior")
				}
			}
		}
		return protowalk.ProcessNever
	})
}

func validateBucketConstraints(ctx context.Context, b *pb.Build) error {
	bck := &model.Bucket{
		Parent: model.ProjectKey(ctx, b.Builder.Project),
		ID:     b.Builder.Bucket,
	}
	bckStr := fmt.Sprintf("%s:%s", b.Builder.Project, b.Builder.Bucket)
	if err := datastore.Get(ctx, bck); err != nil {
		return errors.Annotate(err, "failed to fetch bucket config %s", bckStr).Err()
	}

	constraints := bck.Proto.GetConstraints()
	if constraints == nil {
		return errors.Reason("constraints for %s not found", bckStr).Err()
	}

	// want to return early if swarming is not set and backend is set.
	if b.GetInfra().GetSwarming() == nil {
		return nil
	}
	allowedPools := stringset.NewFromSlice(constraints.GetPools()...)
	allowedSAs := stringset.NewFromSlice(constraints.GetServiceAccounts()...)
	poolAllowed := false
	var pool string
	for _, dim := range b.GetInfra().GetSwarming().GetTaskDimensions() {
		if dim.Key != "pool" {
			continue
		}
		pool = dim.Value
		if allowedPools.Has(dim.Value) {
			poolAllowed = true
			break
		}
	}
	if !poolAllowed {
		return errors.Reason("build.infra.swarming.dimension['pool']: %s not allowed", pool).Err()
	}

	sa := b.GetInfra().GetSwarming().GetTaskServiceAccount()
	if sa == "" || !allowedSAs.Has(sa) {
		return errors.Reason("build.infra.swarming.task_service_account: %s not allowed", sa).Err()
	}
	return nil
}

func validateHostName(host string) error {
	if strings.Contains(host, "://") {
		return errors.Reason(`must not contain "://"`).Err()
	}
	return nil
}

func validateCipdPackage(pkg string, mustWithSuffix bool) error {
	pkgSuffix := "/${platform}"
	if mustWithSuffix && !strings.HasSuffix(pkg, pkgSuffix) {
		return errors.Reason("expected to end with %s", pkgSuffix).Err()
	}
	return cipdCommon.ValidatePackageName(strings.TrimSuffix(pkg, pkgSuffix))
}

func validateAgentInput(in *pb.BuildInfra_Buildbucket_Agent_Input) error {
	for path, ref := range in.GetData() {
		for i, spec := range ref.GetCipd().GetSpecs() {
			if err := validateCipdPackage(spec.GetPackage(), false); err != nil {
				return errors.Annotate(err, "[%s]: [%d]: cipd.package", path, i).Err()
			}
			if err := cipdCommon.ValidateInstanceVersion(spec.GetVersion()); err != nil {
				return errors.Annotate(err, "[%s]: [%d]: cipd.version", path, i).Err()
			}
		}

		cas := ref.GetCas()
		if cas != nil {
			switch {
			case !casInstanceRe.MatchString(cas.GetCasInstance()):
				return errors.Reason("[%s]: cas.cas_instance: does not match %s", path, casInstanceRe).Err()
			case cas.GetDigest() == nil:
				return errors.Reason("[%s]: cas.digest: not specified", path).Err()
			case cas.Digest.GetSizeBytes() < 0:
				return errors.Reason("[%s]: cas.digest.size_bytes: must be greater or equal to 0", path).Err()
			}
		}
	}
	return nil
}

func validateAgentSource(src *pb.BuildInfra_Buildbucket_Agent_Source) error {
	cipd := src.GetCipd()
	if err := validateCipdPackage(cipd.GetPackage(), true); err != nil {
		return errors.Annotate(err, "cipd.package:").Err()
	}
	if err := cipdCommon.ValidateInstanceVersion(cipd.GetVersion()); err != nil {
		return errors.Annotate(err, "cipd.version").Err()
	}
	return nil
}

func validateAgentPurposes(purposes map[string]pb.BuildInfra_Buildbucket_Agent_Purpose, in *pb.BuildInfra_Buildbucket_Agent_Input) error {
	if len(purposes) == 0 {
		return nil
	}

	for path := range purposes {
		if _, ok := in.GetData()[path]; !ok {
			return errors.Reason("Invalid path %s - not in input dataRef", path).Err()
		}
	}
	return nil
}

func validateAgent(agent *pb.BuildInfra_Buildbucket_Agent) error {
	var err error
	switch {
	case teeErr(validateAgentInput(agent.GetInput()), &err) != nil:
		return errors.Annotate(err, "input").Err()
	case teeErr(validateAgentSource(agent.GetSource()), &err) != nil:
		return errors.Annotate(err, "source").Err()
	case teeErr(validateAgentPurposes(agent.GetPurposes(), agent.GetInput()), &err) != nil:
		return errors.Annotate(err, "purposes").Err()
	default:
		return nil
	}
}

func validateInfraBuildbucket(ctx context.Context, ib *pb.BuildInfra_Buildbucket) error {
	var err error
	bbHost := fmt.Sprintf("%s.appspot.com", info.AppID(ctx))
	switch {
	case teeErr(validateHostName(ib.GetHostname()), &err) != nil:
		return errors.Annotate(err, "hostname").Err()
	case ib.GetHostname() != "" && ib.Hostname != bbHost:
		return errors.Reason("incorrect hostname, want: %s, got: %s", bbHost, ib.Hostname).Err()
	case teeErr(validateAgent(ib.GetAgent()), &err) != nil:
		return errors.Annotate(err, "agent").Err()
	case teeErr(validateRequestedDimensions(ib.RequestedDimensions), &err) != nil:
		return errors.Annotate(err, "requested_dimensions").Err()
	case teeErr(validateProperties(ib.RequestedProperties), &err) != nil:
		return errors.Annotate(err, "requested_properties").Err()
	}
	for _, host := range ib.GetKnownPublicGerritHosts() {
		if err = validateHostName(host); err != nil {
			return errors.Annotate(err, "known_public_gerrit_hosts").Err()
		}
	}
	return nil
}

func validateCaches(caches []*pb.BuildInfra_Swarming_CacheEntry) error {
	names := stringset.New(len(caches))
	paths := stringset.New(len(caches))
	for i, cache := range caches {
		switch {
		case cache.Name == "":
			return errors.Reason(fmt.Sprintf("%dth cache: name unspecified", i)).Err()
		case len(cache.Name) > 128:
			return errors.Reason(fmt.Sprintf("%dth cache: name too long (limit is 128)", i)).Err()
		case !names.Add(cache.Name):
			return errors.Reason(fmt.Sprintf("duplicated cache name: %s", cache.Name)).Err()
		case cache.Path == "":
			return errors.Reason(fmt.Sprintf("%dth cache: path unspecified", i)).Err()
		case strings.Contains(cache.Path, "\\"):
			return errors.Reason(fmt.Sprintf("%dth cache: path must use POSIX format", i)).Err()
		case !paths.Add(cache.Path):
			return errors.Reason(fmt.Sprintf("duplicated cache path: %s", cache.Path)).Err()
		case cache.WaitForWarmCache.AsDuration()%(60*time.Second) != 0:
			return errors.Reason(fmt.Sprintf("%dth cache: wait_for_warm_cache must be multiples of 60 seconds.", i)).Err()
		}
	}
	return nil
}

// validateDimensions validates the task dimension.
func validateDimension(dim *pb.RequestedDimension) error {
	var err error
	switch {
	case teeErr(validateExpirationDuration(dim.GetExpiration()), &err) != nil:
		return errors.Annotate(err, "expiration").Err()
	case dim.GetKey() == "":
		return errors.Reason("key must be specified").Err()
	case dim.Value == "":
		return errors.Reason("value must be specified").Err()
	default:
		return nil
	}
}

// validateDimensions validates the task dimensions.
func validateDimensions(dims []*pb.RequestedDimension) error {
	for i, dim := range dims {
		if err := validateDimension(dim); err != nil {
			return errors.Annotate(err, "[%d]", i).Err()
		}
	}
	return nil
}

func validateInfraBackendTaskTarget(target string) error {
	invalidKeywordsForTaskBackendTarget := []string{"http", "rpc"}
	for _, word := range invalidKeywordsForTaskBackendTarget {
		if strings.Contains(target, word) {
			return errors.Reason("backend task target contains invalid keyword: %s.", word).Err()
		}
	}
	split := strings.Split(target, "://")
	if len(split) != 2 {
		return errors.Reason("backend task target was not properly formatted.").Err()
	}
	return nil
}

func validateInfraBackend(ctx context.Context, ib *pb.BuildInfra_Backend) error {
	if ib == nil {
		return nil
	}
	return validateInfraBackendTaskTarget(ib.GetTask().GetId().GetTarget())
}

func validateInfraSwarming(is *pb.BuildInfra_Swarming) error {
	var err error
	if is == nil {
		return nil
	}
	switch {
	case teeErr(validateHostName(is.GetHostname()), &err) != nil:
		return errors.Annotate(err, "hostname").Err()
	case is.GetPriority() < 0 || is.GetPriority() > 255:
		return errors.Reason("priority must be in [0, 255]").Err()
	case teeErr(validateDimensions(is.GetTaskDimensions()), &err) != nil:
		return errors.Annotate(err, "task_dimensions").Err()
	case teeErr(validateCaches(is.GetCaches()), &err) != nil:
		return errors.Annotate(err, "caches").Err()
	default:
		return nil
	}
}

func validateInfraLogDog(il *pb.BuildInfra_LogDog) error {
	var err error
	switch {
	case teeErr(validateHostName(il.GetHostname()), &err) != nil:
		return errors.Annotate(err, "hostname").Err()
	default:
		return nil
	}
}

func validateInfraResultDB(irdb *pb.BuildInfra_ResultDB) error {
	var err error
	switch {
	case irdb == nil:
		return nil
	case teeErr(validateHostName(irdb.GetHostname()), &err) != nil:
		return errors.Annotate(err, "hostname").Err()
	default:
		return nil
	}
}

func validateInfra(ctx context.Context, infra *pb.BuildInfra) error {
	var err error
	switch {
	case infra.GetBackend() == nil && infra.GetSwarming() == nil:
		return errors.Reason("backend or swarming is needed in build infra").Err()
	case infra.GetBackend() != nil && infra.GetSwarming() != nil:
		return errors.Reason("can only have one of backend or swarming in build infra. both were provided").Err()
	case teeErr(validateInfraBackend(ctx, infra.GetBackend()), &err) != nil:
		return errors.Annotate(err, "backend").Err()
	case teeErr(validateInfraSwarming(infra.GetSwarming()), &err) != nil:
		return errors.Annotate(err, "swarming").Err()
	case teeErr(validateInfraBuildbucket(ctx, infra.GetBuildbucket()), &err) != nil:
		return errors.Annotate(err, "buildbucket").Err()
	case teeErr(validateInfraLogDog(infra.GetLogdog()), &err) != nil:
		return errors.Annotate(err, "logdog").Err()
	case teeErr(validateInfraResultDB(infra.GetResultdb()), &err) != nil:
		return errors.Annotate(err, "resultdb").Err()
	default:
		return nil
	}
}

func validateInput(wellKnownExperiments stringset.Set, in *pb.Build_Input) error {
	var err error
	switch {
	case teeErr(validateGerritChanges(in.GerritChanges), &err) != nil:
		return errors.Annotate(err, "gerrit_changes").Err()
	case in.GetGitilesCommit() != nil && teeErr(validateCommitWithRef(in.GitilesCommit), &err) != nil:
		return errors.Annotate(err, "gitiles_commit").Err()
	case in.Properties != nil && teeErr(validateProperties(in.Properties), &err) != nil:
		return errors.Annotate(err, "properties").Err()
	}
	for _, expName := range in.Experiments {
		if err := config.ValidateExperimentName(expName, wellKnownExperiments); err != nil {
			return errors.Annotate(err, "experiment %q", expName).Err()
		}
	}
	return nil
}

func validateExe(exe *pb.Executable, agent *pb.BuildInfra_Buildbucket_Agent) error {
	var err error
	switch {
	case exe.GetCipdPackage() == "":
		return nil
	case teeErr(validateCipdPackage(exe.CipdPackage, false), &err) != nil:
		return errors.Annotate(err, "cipd_package").Err()
	case exe.GetCipdVersion() != "" && teeErr(cipdCommon.ValidateInstanceVersion(exe.CipdVersion), &err) != nil:
		return errors.Annotate(err, "cipd_version").Err()
	}

	// Validate exe matches with agent.
	var payloadPath string
	for dir, purpose := range agent.GetPurposes() {
		if purpose == pb.BuildInfra_Buildbucket_Agent_PURPOSE_EXE_PAYLOAD {
			payloadPath = dir
			break
		}
	}
	if payloadPath == "" {
		return nil
	}

	if pkgs, ok := agent.GetInput().GetData()[payloadPath]; ok {
		cipdPkgs := pkgs.GetCipd()
		if cipdPkgs == nil {
			return errors.Reason("not match build.infra.buildbucket.agent").Err()
		}

		packageMatches := false
		for _, spec := range cipdPkgs.Specs {
			if spec.Package != exe.CipdPackage {
				continue
			}
			packageMatches = true
			if spec.Version != exe.CipdVersion {
				return errors.Reason("cipd_version does not match build.infra.buildbucket.agent").Err()
			}
			break
		}
		if !packageMatches {
			return errors.Reason("cipd_package does not match build.infra.buildbucket.agent").Err()
		}
	}
	return nil
}

func validateBuild(ctx context.Context, wellKnownExperiments stringset.Set, b *pb.Build) error {
	var err error
	switch {
	case teeErr(protoutil.ValidateRequiredBuilderID(b.Builder), &err) != nil:
		return errors.Annotate(err, "builder").Err()
	case teeErr(validateExe(b.Exe, b.GetInfra().GetBuildbucket().GetAgent()), &err) != nil:
		return errors.Annotate(err, "exe").Err()
	case teeErr(validateInput(wellKnownExperiments, b.Input), &err) != nil:
		return errors.Annotate(err, "input").Err()
	case teeErr(validateInfra(ctx, b.Infra), &err) != nil:
		return errors.Annotate(err, "infra").Err()
	case teeErr(validateBucketConstraints(ctx, b), &err) != nil:
		return err
	case teeErr(validateTags(b.Tags, TagNew), &err) != nil:
		return errors.Annotate(err, "tags").Err()
	default:
		return nil
	}
}

func validateCreateBuildRequest(ctx context.Context, wellKnownExperiments stringset.Set, req *pb.CreateBuildRequest) (*model.BuildMask, error) {
	if procRes := protowalk.Fields(req, &protowalk.DeprecatedProcessor{}, &protowalk.OutputOnlyProcessor{}, &protowalk.RequiredProcessor{}, &CreateBuildChecker{}); procRes != nil {
		if resStrs := procRes.Strings(); len(resStrs) > 0 {
			logging.Infof(ctx, strings.Join(resStrs, ". "))
		}
		if err := procRes.Err(); err != nil {
			return nil, err
		}
	}

	if err := validateBuild(ctx, wellKnownExperiments, req.GetBuild()); err != nil {
		return nil, errors.Annotate(err, "build").Err()
	}

	if strings.Contains(req.GetRequestId(), "/") {
		return nil, errors.Reason("request_id cannot contain '/'").Err()
	}

	m, err := model.NewBuildMask("", nil, req.Mask)
	if err != nil {
		return nil, errors.Annotate(err, "invalid mask").Err()
	}

	return m, nil
}

type buildCreator struct {
	// Valid builds to be saved in datastore. The len(blds) <= len(reqIDs)
	blds []*model.Build
	// idxMapBldToReq is an index map of index of blds -> index of reqIDs.
	idxMapBldToReq []int
	// RequestIDs of each request.
	reqIDs []string
	// errors when creating the builds.
	merr errors.MultiError
}

// createBuilds saves the builds to datastore and triggers swarming task creation
// tasks for each saved build.
// A single returned error means a top-level error.
// Otherwise, it would be a MultiError where len(MultiError) equals to len(bc.reqIDs).
func (bc *buildCreator) createBuilds(ctx context.Context) ([]*model.Build, error) {
	now := clock.Now(ctx).UTC()
	user := auth.CurrentIdentity(ctx)
	appID := info.AppID(ctx) // e.g. cr-buildbucket
	ids := buildid.NewBuildIDs(ctx, now, len(bc.blds))
	nums := make([]*model.Build, 0, len(bc.blds))
	var idxMapNums []int

	for i := range bc.blds {
		bc.blds[i].ID = ids[i]
		bc.blds[i].CreatedBy = user
		bc.blds[i].CreateTime = now

		// Set proto field values which can only be determined at creation-time.
		bc.blds[i].Proto.CreatedBy = string(user)
		bc.blds[i].Proto.CreateTime = timestamppb.New(now)
		bc.blds[i].Proto.Id = ids[i]
		if bc.blds[i].Proto.Infra.Buildbucket.Hostname == "" {
			bc.blds[i].Proto.Infra.Buildbucket.Hostname = fmt.Sprintf("%s.appspot.com", appID)
		}
		bc.blds[i].Proto.Infra.Logdog.Prefix = fmt.Sprintf("buildbucket/%s/%d", appID, bc.blds[i].Proto.Id)
		protoutil.SetStatus(now, bc.blds[i].Proto, pb.Status_SCHEDULED)

		if bc.blds[i].Proto.GetInfra().GetBuildbucket().GetBuildNumber() {
			idxMapNums = append(idxMapNums, bc.idxMapBldToReq[i])
			nums = append(nums, bc.blds[i])
		}
	}

	if err := generateBuildNumbers(ctx, nums); err != nil {
		me := err.(errors.MultiError)
		bc.merr = mergeErrs(bc.merr, me, "error generating build numbers", func(idx int) int { return idxMapNums[idx] })
	}

	validBlds, idxMapValidBlds := getValidBlds(bc.blds, bc.merr, bc.idxMapBldToReq)
	err := parallel.FanOutIn(func(work chan<- func() error) {
		work <- func() error { return model.UpdateBuilderStat(ctx, validBlds, now) }
		work <- func() error { return resultdb.CreateInvocations(ctx, validBlds) }
		work <- func() error { return search.UpdateTagIndex(ctx, validBlds) }
	})
	if err != nil {
		errs := err.(errors.MultiError)
		for _, e := range errs {
			if me, ok := e.(errors.MultiError); ok {
				bc.merr = mergeErrs(bc.merr, me, "", func(idx int) int { return idxMapValidBlds[idx] })
			} else {
				return nil, e // top-level error
			}
		}
	}

	// This parallel work isn't combined with the above parallel work to ensure build entities and Swarming (or Backend)
	// task creation tasks are only created if everything else has succeeded (since everything can't be done
	// in one transaction).
	_ = parallel.WorkPool(min(64, len(validBlds)), func(work chan<- func() error) {
		for i, b := range validBlds {
			i := i
			b := b
			origI := idxMapValidBlds[i]
			if bc.merr[origI] != nil {
				validBlds[i] = nil
				continue
			}

			reqID := bc.reqIDs[origI]
			work <- func() error {
				toPut := []interface{}{
					b,
					&model.BuildInfra{
						Build: datastore.KeyForObj(ctx, b),
						Proto: b.Proto.Infra,
					},
					&model.BuildInputProperties{
						Build: datastore.KeyForObj(ctx, b),
						Proto: b.Proto.Input.Properties,
					},
				}
				r := model.NewRequestID(ctx, b.ID, now, reqID)

				// Write the entities and trigger a task queue task to create the Swarming task.
				err := datastore.RunInTransaction(ctx, func(ctx context.Context) error {
					// Deduplicate by request ID.
					if reqID != "" {
						switch err := datastore.Get(ctx, r); {
						case err == datastore.ErrNoSuchEntity:
							toPut = append(toPut, r)
						case err != nil:
							return errors.Annotate(err, "failed to deduplicate request ID: %d", b.ID).Err()
						default:
							b.ID = r.BuildID
							if err := datastore.Get(ctx, b); err != nil {
								return errors.Annotate(err, "failed to fetch deduplicated build: %d", b.ID).Err()
							}
							return nil
						}
					}

					// Request was not a duplicate.
					switch err := datastore.Get(ctx, &model.Build{ID: b.ID}); {
					case err == nil:
						return appstatus.Errorf(codes.AlreadyExists, "build already exists: %d", b.ID)
					case err != datastore.ErrNoSuchEntity:
						return errors.Annotate(err, "failed to fetch build: %d", b.ID).Err()
					}
					if err := datastore.Put(ctx, toPut...); err != nil {
						return errors.Annotate(err, "failed to store build: %d", b.ID).Err()
					}

					// If there is a backend set, lets use it and return to not use swarming.
					if b.Proto.Infra.GetBackend() != nil {
						if err := tasks.CreateBackendBuildTask(ctx, &taskdefs.CreateBackendBuildTask{
							BuildId: b.ID,
						}); err != nil {
							return errors.Annotate(err, "failed to enqueue CreateBackendTask").Err()
						}
						return nil
					}

					if stringset.NewFromSlice(b.Proto.Input.Experiments...).Has(bb.ExperimentBackendGo) {
						if err := tasks.CreateSwarmingBuildTask(ctx, &taskdefs.CreateSwarmingBuildTask{
							BuildId: b.ID,
						}); err != nil {
							return errors.Annotate(err, "failed to enqueue CreateSwarmingBuildTask: %d", b.ID).Err()
						}
					} else {
						if err := tasks.CreateSwarmingTask(ctx, &taskdefs.CreateSwarmingTask{
							BuildId: b.ID,
						}); err != nil {
							return errors.Annotate(err, "failed to enqueue CreateSwarmingTask: %d", b.ID).Err()
						}
					}
					return nil
				}, nil)

				// Record any error happened in the above transaction.
				if err != nil {
					validBlds[i] = nil
					bc.merr[origI] = err
					return nil
				}
				metrics.BuildCreated(ctx, b)
				return nil
			}
		}
	})

	if bc.merr.First() == nil {
		return validBlds, nil
	}
	// Map back to final results to make sure len(resBlds) always equal to len(reqs).
	resBlds := make([]*model.Build, len(bc.reqIDs))
	for i, bld := range validBlds {
		origI := idxMapValidBlds[i]
		if bc.merr[origI] == nil {
			resBlds[origI] = bld
		}
	}
	return resBlds, bc.merr
}

// getValidBlds returns a list of valid builds where its corresponding error is nil.
// as well as an index map where idxMap[returnedIndex] == originalIndex.
func getValidBlds(blds []*model.Build, origErrs errors.MultiError, idxMapBldToReq []int) ([]*model.Build, []int) {
	if len(blds) != len(idxMapBldToReq) {
		panic("The length of blds and the length of idxMapBldToReq must be the same.")
	}
	var validBlds []*model.Build
	var idxMap []int
	for i, bld := range blds {
		origI := idxMapBldToReq[i]
		if origErrs[origI] == nil {
			idxMap = append(idxMap, origI)
			validBlds = append(validBlds, bld)
		}
	}
	return validBlds, idxMap
}

// generateBuildNumbers mutates the given builds, setting build numbers and
// build address tags.
//
// It would return a MultiError (if any) where len(MultiError) equals to len(reqs).
func generateBuildNumbers(ctx context.Context, builds []*model.Build) error {
	merr := make(errors.MultiError, len(builds))
	seq := make(map[string][]*model.Build)
	idxMap := make(map[string][]int) // BuilderID -> a list of index
	for i, b := range builds {
		name := protoutil.FormatBuilderID(b.Proto.Builder)
		seq[name] = append(seq[name], b)
		idxMap[name] = append(idxMap[name], i)
	}
	_ = parallel.WorkPool(min(64, len(builds)), func(work chan<- func() error) {
		for name, blds := range seq {
			name := name
			blds := blds
			work <- func() error {
				n, err := model.GenerateSequenceNumbers(ctx, name, len(blds))
				if err != nil {
					for _, idx := range idxMap[name] {
						merr[idx] = err
					}
					return nil
				}
				for i, b := range blds {
					b.Proto.Number = n + int32(i)
					addr := fmt.Sprintf("build_address:luci.%s.%s/%s/%d", b.Proto.Builder.Project, b.Proto.Builder.Bucket, b.Proto.Builder.Builder, b.Proto.Number)
					b.Tags = append(b.Tags, addr)
					sort.Strings(b.Tags)
				}
				return nil
			}
		}
	})

	if merr.First() == nil {
		return nil
	}
	return merr.AsError()
}

// CreateBuild handles a request to create a build. Implements pb.BuildsServer.
func (*Builds) CreateBuild(ctx context.Context, req *pb.CreateBuildRequest) (*pb.Build, error) {
	if err := perm.HasInBucket(ctx, bbperms.BuildsCreate, req.Build.Builder.Project, req.Build.Builder.Bucket); err != nil {
		return nil, err
	}

	globalCfg, err := config.GetSettingsCfg(ctx)
	if err != nil {
		return nil, errors.Annotate(err, "error fetching service config").Err()
	}
	wellKnownExperiments := protoutil.WellKnownExperiments(globalCfg)

	m, err := validateCreateBuildRequest(ctx, wellKnownExperiments, req)
	if err != nil {
		return nil, appstatus.BadRequest(err)
	}

	bld := &model.Build{
		Proto: req.Build,
	}
	bc := &buildCreator{
		blds:           []*model.Build{bld},
		idxMapBldToReq: []int{0},
		reqIDs:         []string{req.RequestId},
		merr:           make(errors.MultiError, 1),
	}
	blds, err := bc.createBuilds(ctx)
	if err != nil {
		return nil, errors.Annotate(err, "error creating build").Err()
	}

	return blds[0].ToProto(ctx, m, nil)
}
