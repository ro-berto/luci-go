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

package tasks

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"path/filepath"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"google.golang.org/api/googleapi"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"go.chromium.org/luci/common/api/swarming/swarming/v1"
	"go.chromium.org/luci/common/clock/testclock"
	"go.chromium.org/luci/common/data/caching/lru"
	"go.chromium.org/luci/common/retry/transient"
	"go.chromium.org/luci/gae/filter/txndefer"
	"go.chromium.org/luci/gae/impl/memory"
	"go.chromium.org/luci/gae/service/datastore"
	"go.chromium.org/luci/server/caching"
	"go.chromium.org/luci/server/caching/cachingtest"
	"go.chromium.org/luci/server/tq"

	"go.chromium.org/luci/buildbucket/appengine/internal/clients"
	"go.chromium.org/luci/buildbucket/appengine/internal/metrics"
	"go.chromium.org/luci/buildbucket/appengine/model"
	taskdefs "go.chromium.org/luci/buildbucket/appengine/tasks/defs"
	"go.chromium.org/luci/buildbucket/cmd/bbagent/bbinput"
	pb "go.chromium.org/luci/buildbucket/proto"
	"go.chromium.org/luci/buildbucket/protoutil"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"
)

func TestTaskDef(t *testing.T) {
	Convey("compute task slice", t, func() {
		b := &model.Build{
			ID: 123,
			Proto: &pb.Build{
				Id: 123,
				SchedulingTimeout: &durationpb.Duration{
					Seconds: 3600,
				},
				ExecutionTimeout: &durationpb.Duration{
					Seconds: 4800,
				},
				GracePeriod: &durationpb.Duration{
					Seconds: 60,
				},
				Infra: &pb.BuildInfra{
					Swarming: &pb.BuildInfra_Swarming{},
					Bbagent: &pb.BuildInfra_BBAgent{
						CacheDir:    "cache",
						PayloadPath: "kitchen-checkout",
					},
					Buildbucket: &pb.BuildInfra_Buildbucket{
						Agent: &pb.BuildInfra_Buildbucket_Agent{
							Source: &pb.BuildInfra_Buildbucket_Agent_Source{
								DataType: &pb.BuildInfra_Buildbucket_Agent_Source_Cipd{
									Cipd: &pb.BuildInfra_Buildbucket_Agent_Source_CIPD{
										Package: "infra/tools/luci/bbagent/${platform}",
										Version: "canary-version",
										Server:  "cipd server",
									},
								},
							},
						},
					},
				},
			},
		}
		Convey("only base slice", func() {
			b.Proto.Infra.Swarming = &pb.BuildInfra_Swarming{
				Caches: []*pb.BuildInfra_Swarming_CacheEntry{
					{Name: "shared_builder_cache", Path: "builder"},
				},
				TaskDimensions: []*pb.RequestedDimension{
					{Key: "pool", Value: "Chrome"},
				},
			}
			slices, err := computeTaskSlice(b)
			So(err, ShouldBeNil)
			So(len(slices), ShouldEqual, 1)
			So(slices[0].Properties.Caches, ShouldResemble, []*swarming.SwarmingRpcsCacheEntry{{
				Path: filepath.Join("cache", "builder"),
				Name: "shared_builder_cache",
			}})
			So(slices[0].Properties.Dimensions, ShouldResemble, []*swarming.SwarmingRpcsStringPair{
				{
					Key:   "pool",
					Value: "Chrome",
				},
			})
		})

		Convey("multiple dimensions and cache fallback", func() {
			// Creates 4 task_slices by modifying the buildercfg in 2 ways:
			//  - Add two named caches, one expiring at 60 seconds, one at 360 seconds.
			//  - Add an optional builder dimension, expiring at 120 seconds.
			//
			// This ensures the combination of these features works correctly, and that
			// multiple 'caches' dimensions can be injected.
			b.Proto.Infra.Swarming = &pb.BuildInfra_Swarming{
				Caches: []*pb.BuildInfra_Swarming_CacheEntry{
					{Name: "shared_builder_cache", Path: "builder", WaitForWarmCache: &durationpb.Duration{Seconds: 60}},
					{Name: "second_cache", Path: "second", WaitForWarmCache: &durationpb.Duration{Seconds: 360}},
				},
				TaskDimensions: []*pb.RequestedDimension{
					{Key: "a", Value: "1", Expiration: &durationpb.Duration{Seconds: 120}},
					{Key: "a", Value: "2", Expiration: &durationpb.Duration{Seconds: 120}},
					{Key: "pool", Value: "Chrome"},
				},
			}
			slices, err := computeTaskSlice(b)
			So(err, ShouldBeNil)
			So(len(slices), ShouldEqual, 4)

			// All slices properties fields have the same value except dimensions.
			for _, tSlice := range slices {
				So(tSlice.Properties.ExecutionTimeoutSecs, ShouldEqual, 4800)
				So(tSlice.Properties.GracePeriodSecs, ShouldEqual, 240)
				So(tSlice.Properties.Caches, ShouldResemble, []*swarming.SwarmingRpcsCacheEntry{
					{Path: filepath.Join("cache", "builder"), Name: "shared_builder_cache"},
					{Path: filepath.Join("cache", "second"), Name: "second_cache"},
				})
				So(tSlice.Properties.Env, ShouldResemble, []*swarming.SwarmingRpcsStringPair{
					{Key: "BUILDBUCKET_EXPERIMENTAL", Value: "FALSE"},
				})
			}

			So(slices[0].ExpirationSecs, ShouldEqual, 60)
			// The dimensions are different. 'a' and 'caches' are injected.
			So(slices[0].Properties.Dimensions, ShouldResemble, []*swarming.SwarmingRpcsStringPair{
				{Key: "a", Value: "1"},
				{Key: "a", Value: "2"},
				{Key: "caches", Value: "second_cache"},
				{Key: "caches", Value: "shared_builder_cache"},
				{Key: "pool", Value: "Chrome"},
			})

			// 120 - 60
			So(slices[1].ExpirationSecs, ShouldEqual, 60)
			// The dimensions are different. 'a' and 'caches' are injected.
			So(slices[1].Properties.Dimensions, ShouldResemble, []*swarming.SwarmingRpcsStringPair{
				{Key: "a", Value: "1"},
				{Key: "a", Value: "2"},
				{Key: "caches", Value: "second_cache"},
				{Key: "pool", Value: "Chrome"},
			})

			// 360 - 120
			So(slices[2].ExpirationSecs, ShouldEqual, 240)
			// 'a' expired, one 'caches' remains.
			So(slices[2].Properties.Dimensions, ShouldResemble, []*swarming.SwarmingRpcsStringPair{
				{Key: "caches", Value: "second_cache"},
				{Key: "pool", Value: "Chrome"},
			})

			// 3600-360
			So(slices[3].ExpirationSecs, ShouldEqual, 3240)
			// # The cold fallback; the last 'caches' expired.
			So(slices[3].Properties.Dimensions, ShouldResemble, []*swarming.SwarmingRpcsStringPair{
				{Key: "pool", Value: "Chrome"},
			})
		})
	})

	Convey("compute bbagent command", t, func() {
		b := &model.Build{
			ID: 123,
			Proto: &pb.Build{
				Infra: &pb.BuildInfra{
					Buildbucket: &pb.BuildInfra_Buildbucket{
						Hostname: "bbhost.com",
					},
				},
			},
		}
		Convey("bbagent_getbuild experiment", func() {
			b.Experiments = []string{"+luci.buildbucket.bbagent_getbuild"}
			bbagentCmd := computeCommand(b)
			So(bbagentCmd, ShouldResemble, []string{
				"bbagent${EXECUTABLE_SUFFIX}",
				"-host",
				"bbhost.com",
				"-build-id",
				"123",
			})
		})

		Convey("no bbagent_getbuild experiment", func() {
			b.Experiments = []string{"-luci.buildbucket.bbagent_getbuild"}
			b.Proto.Infra.Bbagent = &pb.BuildInfra_BBAgent{
				CacheDir:    "cache",
				PayloadPath: "payload_path",
			}
			bbagentCmd := computeCommand(b)
			expectedEncoded := bbinput.Encode(&pb.BBAgentArgs{
				Build:       b.Proto,
				CacheDir:    "cache",
				PayloadPath: "payload_path",
			})
			So(bbagentCmd, ShouldResemble, []string{
				"bbagent${EXECUTABLE_SUFFIX}",
				expectedEncoded,
			})
		})
	})

	Convey("compute env_prefixes", t, func() {
		b := &model.Build{
			ID: 123,
			Proto: &pb.Build{
				Infra: &pb.BuildInfra{
					Swarming: &pb.BuildInfra_Swarming{},
				},
			},
		}
		Convey("empty swarming cache", func() {
			prefixes := computeEnvPrefixes(b)
			So(prefixes, ShouldResemble, []*swarming.SwarmingRpcsStringListPair{})
		})

		Convey("normal", func() {
			b.Proto.Infra.Swarming.Caches = []*pb.BuildInfra_Swarming_CacheEntry{
				{Path: "vpython", Name: "vpython", EnvVar: "VPYTHON_VIRTUALENV_ROOT"},
				{Path: "abc", Name: "abc", EnvVar: "ABC"},
			}
			prefixes := computeEnvPrefixes(b)
			So(prefixes, ShouldResemble, []*swarming.SwarmingRpcsStringListPair{
				{Key: "ABC", Value: []string{filepath.Join("cache", "abc")}},
				{Key: "VPYTHON_VIRTUALENV_ROOT", Value: []string{filepath.Join("cache", "vpython")}},
			})
		})
	})

	Convey("compute swarming new task req", t, func() {
		ctx := memory.UseWithAppID(context.Background(), "dev~app-id")
		ctx, _ = testclock.UseTime(ctx, time.Unix(1444945245, 0).UTC())
		b := &model.Build{
			ID:        123,
			Project:   "project",
			BucketID:  "bucket",
			BuilderID: "builder",
			Proto: &pb.Build{
				Id:     123,
				Number: 1,
				Builder: &pb.BuilderID{
					Project: "project",
					Bucket:  "bucket",
					Builder: "builder",
				},
				Infra: &pb.BuildInfra{
					Swarming: &pb.BuildInfra_Swarming{
						Priority:           20,
						TaskServiceAccount: "abc",
						Hostname:           "swarm.com",
					},
					Bbagent: &pb.BuildInfra_BBAgent{},
					Buildbucket: &pb.BuildInfra_Buildbucket{
						Agent: &pb.BuildInfra_Buildbucket_Agent{
							Source: &pb.BuildInfra_Buildbucket_Agent_Source{
								DataType: &pb.BuildInfra_Buildbucket_Agent_Source_Cipd{
									Cipd: &pb.BuildInfra_Buildbucket_Agent_Source_CIPD{
										Package: "infra/tools/luci/bbagent/${platform}",
										Version: "canary-version",
										Server:  "cipd server",
									},
								},
							},
						},
					},
				},
			},
		}

		req, err := computeSwarmingNewTaskReq(ctx, b)
		// Strip out TaskSlices. It has been tested in other tests
		req.TaskSlices = []*swarming.SwarmingRpcsTaskSlice(nil)
		So(err, ShouldBeNil)
		ud, _ := json.Marshal(&userdata{
			BuildID:          123,
			CreatedTS:        1444945245000000,
			SwarmingHostname: "swarm.com",
		})
		expected := &swarming.SwarmingRpcsNewTaskRequest{
			RequestUuid:    "203882df-ce4b-5012-b32a-2c1d29c321a7",
			Name:           "bb-123-builder-1",
			Realm:          "project:bucket",
			Tags:           []string{"buildbucket_bucket:bucket", "buildbucket_build_id:123", "buildbucket_hostname:app-id.appspot.com", "buildbucket_template_canary:0", "luci_project:project"},
			Priority:       int64(20),
			PubsubTopic:    "projects/app-id/topics/swarming-go",
			PubsubUserdata: string(ud),
			ServiceAccount: "abc",
		}
		So(req, ShouldResemble, expected)
	})
}

func TestSyncBuild(t *testing.T) {
	t.Parallel()
	Convey("SyncBuild", t, func() {
		ctl := gomock.NewController(t)
		defer ctl.Finish()
		now := testclock.TestRecentTimeUTC
		mockSwarm := clients.NewMockSwarmingClient(ctl)
		ctx, _ := testclock.UseTime(context.Background(), now)
		ctx = context.WithValue(ctx, &clients.MockSwarmingClientKey, mockSwarm)
		ctx = memory.UseWithAppID(ctx, "dev~app-id")
		ctx = txndefer.FilterRDS(ctx)
		ctx = metrics.WithServiceInfo(ctx, "svc", "job", "ins")
		datastore.GetTestable(ctx).AutoIndex(true)
		datastore.GetTestable(ctx).Consistent(true)
		ctx, sch := tq.TestingContext(ctx, nil)

		b := &model.Build{
			ID: 123,
			Proto: &pb.Build{
				Id:         123,
				Status:     pb.Status_SCHEDULED,
				CreateTime: &timestamppb.Timestamp{Seconds: now.UnixNano() / 1000000000},
				SchedulingTimeout: &durationpb.Duration{
					Seconds: 3600,
				},
				ExecutionTimeout: &durationpb.Duration{
					Seconds: 4800,
				},
				GracePeriod: &durationpb.Duration{
					Seconds: 60,
				},
				Builder: &pb.BuilderID{
					Project: "proj",
					Bucket:  "bucket",
					Builder: "builder",
				},
			},
		}
		So(datastore.Put(ctx, b), ShouldBeNil)
		inf := &model.BuildInfra{
			ID:    1,
			Build: datastore.KeyForObj(ctx, &model.Build{ID: 123}),
			Proto: &pb.BuildInfra{
				Swarming: &pb.BuildInfra_Swarming{
					Hostname: "swarm",
					Caches: []*pb.BuildInfra_Swarming_CacheEntry{
						{Name: "shared_builder_cache", Path: "builder", WaitForWarmCache: &durationpb.Duration{Seconds: 60}},
						{Name: "second_cache", Path: "second", WaitForWarmCache: &durationpb.Duration{Seconds: 360}},
					},
					TaskDimensions: []*pb.RequestedDimension{
						{Key: "a", Value: "1", Expiration: &durationpb.Duration{Seconds: 120}},
						{Key: "a", Value: "2", Expiration: &durationpb.Duration{Seconds: 120}},
						{Key: "pool", Value: "Chrome"},
					},
				},
				Bbagent: &pb.BuildInfra_BBAgent{
					CacheDir:    "cache",
					PayloadPath: "kitchen-checkout",
				},
				Buildbucket: &pb.BuildInfra_Buildbucket{
					Agent: &pb.BuildInfra_Buildbucket_Agent{
						Source: &pb.BuildInfra_Buildbucket_Agent_Source{
							DataType: &pb.BuildInfra_Buildbucket_Agent_Source_Cipd{
								Cipd: &pb.BuildInfra_Buildbucket_Agent_Source_CIPD{
									Package: "infra/tools/luci/bbagent/${platform}",
									Version: "canary-version",
									Server:  "cipd server",
								},
							},
						},
					},
				},
			},
		}
		So(datastore.Put(ctx, inf), ShouldBeNil)
		Convey("swarming-build-create", func() {

			Convey("build not found", func() {
				err := SyncBuild(ctx, 789, 0)
				So(err, ShouldErrLike, "build 789 or buildInfra not found")
			})

			Convey("build too old", func() {
				So(datastore.Put(ctx, &model.Build{
					ID:         111,
					CreateTime: now.AddDate(0, 0, -3),
					Proto: &pb.Build{
						Builder: &pb.BuilderID{},
					},
				}), ShouldBeNil)

				So(datastore.Put(ctx, &model.BuildInfra{
					ID:    1,
					Build: datastore.KeyForObj(ctx, &model.Build{ID: 111}),
				}), ShouldBeNil)
				err := SyncBuild(ctx, 111, 0)
				So(err, ShouldBeNil)
				So(sch.Tasks(), ShouldHaveLength, 0)
			})

			Convey("build ended", func() {
				So(datastore.Put(ctx, &model.Build{
					ID:     111,
					Status: pb.Status_SUCCESS,
					Proto: &pb.Build{
						Builder: &pb.BuilderID{},
					},
				}), ShouldBeNil)
				So(datastore.Put(ctx, &model.BuildInfra{
					ID:    1,
					Build: datastore.KeyForObj(ctx, &model.Build{ID: 111}),
				}), ShouldBeNil)
				err := SyncBuild(ctx, 111, 0)
				So(err, ShouldBeNil)
				So(sch.Tasks(), ShouldHaveLength, 0)
			})

			Convey("create swarming success", func() {
				mockSwarm.EXPECT().CreateTask(gomock.Any(), gomock.Any()).Return(&swarming.SwarmingRpcsTaskRequestMetadata{
					TaskId: "task123",
				}, nil)
				err := SyncBuild(ctx, 123, 0)
				So(err, ShouldBeNil)
				updatedBuild := &model.Build{ID: 123}
				updatedInfra := &model.BuildInfra{Build: datastore.KeyForObj(ctx, updatedBuild)}
				So(datastore.Get(ctx, updatedBuild), ShouldBeNil)
				So(datastore.Get(ctx, updatedInfra), ShouldBeNil)
				So(updatedBuild.UpdateToken, ShouldNotBeEmpty)
				So(updatedInfra.Proto.Swarming.TaskId, ShouldEqual, "task123")
				So(sch.Tasks(), ShouldHaveLength, 1)
			})

			Convey("create swarming http 400 err", func() {
				mockSwarm.EXPECT().CreateTask(gomock.Any(), gomock.Any()).Return(nil, &googleapi.Error{Code: 400})
				err := SyncBuild(ctx, 123, 0)
				So(err, ShouldBeNil)
				failedBuild := &model.Build{ID: 123}
				So(datastore.Get(ctx, failedBuild), ShouldBeNil)
				So(failedBuild.Status, ShouldEqual, pb.Status_INFRA_FAILURE)
				So(failedBuild.Proto.SummaryMarkdown, ShouldContainSubstring, "failed to create a swarming task: googleapi: got HTTP response code 400")
				So(sch.Tasks(), ShouldHaveLength, 3)
			})

			Convey("create swarming http 500 err", func() {
				mockSwarm.EXPECT().CreateTask(gomock.Any(), gomock.Any()).Return(nil, &googleapi.Error{Code: 500})
				err := SyncBuild(ctx, 123, 0)
				So(err, ShouldErrLike, "failed to create a swarming task")
				So(transient.Tag.In(err), ShouldBeTrue)
				bld := &model.Build{ID: 123}
				So(datastore.Get(ctx, bld), ShouldBeNil)
				So(bld.Status, ShouldEqual, pb.Status_SCHEDULED)
			})

			Convey("create swarming http 500 err give up", func() {
				ctx1, _ := testclock.UseTime(ctx, now.Add(swarmingCreateTaskGiveUpTimeout))
				mockSwarm.EXPECT().CreateTask(gomock.Any(), gomock.Any()).Return(nil, &googleapi.Error{Code: 500})
				err := SyncBuild(ctx1, 123, 0)
				So(err, ShouldBeNil)
				failedBuild := &model.Build{ID: 123}
				So(datastore.Get(ctx, failedBuild), ShouldBeNil)
				So(failedBuild.Status, ShouldEqual, pb.Status_INFRA_FAILURE)
				So(failedBuild.Proto.SummaryMarkdown, ShouldContainSubstring, "failed to create a swarming task: googleapi: got HTTP response code 500")
				So(sch.Tasks(), ShouldHaveLength, 3)
			})

			Convey("swarming task creation success but update build fail", func() {
				mockSwarm.EXPECT().CreateTask(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, req *swarming.SwarmingRpcsNewTaskRequest) (*swarming.SwarmingRpcsTaskRequestMetadata, error) {
					// Hack to make the build update fail when trying to update build with the new task ID.
					inf.Proto.Swarming.TaskId = "old task ID"
					So(datastore.Put(ctx, inf), ShouldBeNil)
					return &swarming.SwarmingRpcsTaskRequestMetadata{TaskId: "new task ID"}, nil
				})

				err := SyncBuild(ctx, 123, 0)
				So(err, ShouldErrLike, "failed to update build 123: build already has a task old task ID")
				currentInfra := &model.BuildInfra{Build: datastore.KeyForObj(ctx, &model.Build{
					ID: 123,
				})}
				So(datastore.Get(ctx, currentInfra), ShouldBeNil)
				So(currentInfra.Proto.Swarming.TaskId, ShouldEqual, "old task ID")
				// One task should be pushed into CancelSwarmingTask queue since update build failed.
				So(sch.Tasks(), ShouldHaveLength, 1)
				So(sch.Tasks().Payloads()[0], ShouldResembleProto, &taskdefs.CancelSwarmingTask{
					Hostname: "swarm",
					TaskId:   "new task ID",
					Realm:    "proj:bucket",
				})
			})

			Convey("create swarming - cancel swarming task fail", func() {
				mockSwarm.EXPECT().CreateTask(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, req *swarming.SwarmingRpcsNewTaskRequest) (*swarming.SwarmingRpcsTaskRequestMetadata, error) {
					// Hack to make the build update and the equeueing CancelSwarmingTask fail.
					inf.Proto.Swarming.TaskId = "old task ID"
					So(datastore.Put(ctx, inf), ShouldBeNil)
					return &swarming.SwarmingRpcsTaskRequestMetadata{TaskId: ""}, nil
				})

				err := SyncBuild(ctx, 123, 0)
				So(err, ShouldErrLike, "failed to enqueue swarming task cancellation task for build 123: task_id is required")
			})
		})

		Convey("swarming sync", func() {
			inf.Proto.Swarming.TaskId = "task_id"
			So(datastore.Put(ctx, inf), ShouldBeNil)

			Convey("non-existing task ID", func() {
				mockSwarm.EXPECT().GetTaskResult(ctx, "task_id").Return(nil, &googleapi.Error{Code: 404})
				err := syncBuildWithTaskResult(ctx, 123, "task_id", mockSwarm)
				So(err, ShouldBeNil)
				failedBuild := &model.Build{ID: 123}
				So(datastore.Get(ctx, failedBuild), ShouldBeNil)
				So(failedBuild.Status, ShouldEqual, pb.Status_INFRA_FAILURE)
				So(failedBuild.Proto.SummaryMarkdown, ShouldContainSubstring, "invalid swarming task task_id")
			})

			Convey("swarming server 500", func() {
				mockSwarm.EXPECT().GetTaskResult(ctx, "task_id").Return(nil, &googleapi.Error{Code: 500})
				err := syncBuildWithTaskResult(ctx, 123, "task_id", mockSwarm)
				So(transient.Tag.In(err), ShouldBeTrue)
			})

			Convey("empty task result", func() {
				mockSwarm.EXPECT().GetTaskResult(ctx, "task_id").Return(nil, nil)
				err := syncBuildWithTaskResult(ctx, 123, "task_id", mockSwarm)
				So(err, ShouldBeNil)
				failedBuild := &model.Build{ID: 123}
				So(datastore.Get(ctx, failedBuild), ShouldBeNil)
				So(failedBuild.Status, ShouldEqual, pb.Status_INFRA_FAILURE)
				So(failedBuild.Proto.SummaryMarkdown, ShouldContainSubstring, "Swarming task task_id unexpectedly disappeared")
			})

			Convey("invalid task result state", func() {
				// syncBuildWithTaskResult should return Fatal error
				mockSwarm.EXPECT().GetTaskResult(ctx, "task_id").Return(&swarming.SwarmingRpcsTaskResult{State: "INVALID"}, nil)
				err := syncBuildWithTaskResult(ctx, 123, "task_id", mockSwarm)
				So(tq.Fatal.In(err), ShouldBeTrue)
				bb := &model.Build{ID: 123}
				So(datastore.Get(ctx, bb), ShouldBeNil)
				So(bb.Status, ShouldEqual, pb.Status_SCHEDULED) // build status should not been impacted

				// The swarming-build-sync flow shouldn't bubble up the Fatal error.
				// It should ignore and enqueue the next generation of sync task.
				mockSwarm.EXPECT().GetTaskResult(ctx, "task_id").Return(&swarming.SwarmingRpcsTaskResult{State: "INVALID"}, nil)
				err = SyncBuild(ctx, 123, 1)
				So(err, ShouldBeNil)
				bb = &model.Build{ID: 123}
				So(datastore.Get(ctx, bb), ShouldBeNil)
				So(bb.Status, ShouldEqual, pb.Status_SCHEDULED)
				So(sch.Tasks(), ShouldHaveLength, 1)
				So(sch.Tasks().Payloads()[0], ShouldResembleProto, &taskdefs.SyncSwarmingBuildTask{
					BuildId:    123,
					Generation: 2,
				})
			})

			Convey("cancel incomplete steps for an ended build", func() {
				mockSwarm.EXPECT().GetTaskResult(ctx, "task_id").Return(&swarming.SwarmingRpcsTaskResult{
					State:       "BOT_DIED",
					StartedTs:   "2018-01-29T21:15:02.649750",
					CompletedTs: "2018-01-30T00:15:18.162860"}, nil)
				steps := model.BuildSteps{
					ID:    1,
					Build: datastore.KeyForObj(ctx, &model.Build{ID: 123}),
				}
				So(steps.FromProto([]*pb.Step{
					{Name: "step1", Status: pb.Status_SUCCESS},
					{Name: "step2", Status: pb.Status_STARTED},
				}), ShouldBeNil)
				So(datastore.Put(ctx, &steps), ShouldBeNil)

				err := syncBuildWithTaskResult(ctx, 123, "task_id", mockSwarm)
				So(err, ShouldBeNil)
				failedBuild := &model.Build{ID: 123}
				So(datastore.Get(ctx, failedBuild), ShouldBeNil)
				So(failedBuild.Status, ShouldEqual, pb.Status_INFRA_FAILURE)
				allSteps := &model.BuildSteps{
					ID:    1,
					Build: datastore.KeyForObj(ctx, &model.Build{ID: 123}),
				}
				So(datastore.Get(ctx, allSteps), ShouldBeNil)
				mSteps, err := allSteps.ToProto(ctx)
				So(err, ShouldBeNil)
				So(mSteps, ShouldResembleProto, []*pb.Step{
					{
						Name:   "step1",
						Status: pb.Status_SUCCESS,
					},
					{
						Name:    "step2",
						Status:  pb.Status_CANCELED,
						EndTime: &timestamppb.Timestamp{Seconds: 1517271318, Nanos: 162860000},
					},
				})
			})

			var cases = []struct {
				fakeTaskResult *swarming.SwarmingRpcsTaskResult
				expected       *expectedBuildFields
			}{
				{
					fakeTaskResult: &swarming.SwarmingRpcsTaskResult{State: "PENDING"},
					expected: &expectedBuildFields{
						status: pb.Status_SCHEDULED,
					},
				},
				{
					fakeTaskResult: &swarming.SwarmingRpcsTaskResult{
						State:     "RUNNING",
						StartedTs: "2018-01-29T21:15:02.649750",
					},
					expected: &expectedBuildFields{
						status: pb.Status_STARTED,
						startT: &timestamppb.Timestamp{Seconds: 1517260502, Nanos: 649750000},
					},
				},
				{
					fakeTaskResult: &swarming.SwarmingRpcsTaskResult{
						State:       "COMPLETED",
						StartedTs:   "2018-01-29T21:15:02.649750",
						CompletedTs: "2018-01-30T00:15:18.162860",
					},
					expected: &expectedBuildFields{
						status: pb.Status_SUCCESS,
						startT: &timestamppb.Timestamp{Seconds: 1517260502, Nanos: 649750000},
						endT:   &timestamppb.Timestamp{Seconds: 1517271318, Nanos: 162860000},
					},
				},
				{
					fakeTaskResult: &swarming.SwarmingRpcsTaskResult{
						State:       "COMPLETED",
						StartedTs:   "2018-01-29T21:15:02.649750",
						CompletedTs: "2018-01-30T00:15:18.162860",
						BotDimensions: []*swarming.SwarmingRpcsStringListPair{
							{
								Key:   "os",
								Value: []string{"Ubuntu", "Trusty"},
							},
							{
								Key:   "pool",
								Value: []string{"luci.chromium.try"},
							},
							{
								Key:   "id",
								Value: []string{"bot1"},
							},
							{
								Key: "empty",
							},
						},
					},
					expected: &expectedBuildFields{
						status: pb.Status_SUCCESS,
						startT: &timestamppb.Timestamp{Seconds: 1517260502, Nanos: 649750000},
						endT:   &timestamppb.Timestamp{Seconds: 1517271318, Nanos: 162860000},
						botDimensions: []*pb.StringPair{
							{Key: "id", Value: "bot1"},
							{Key: "os", Value: "Trusty"},
							{Key: "os", Value: "Ubuntu"},
							{Key: "pool", Value: "luci.chromium.try"},
						},
					},
				},
				{
					fakeTaskResult: &swarming.SwarmingRpcsTaskResult{
						State:       "COMPLETED",
						Failure:     true,
						StartedTs:   "2018-01-29T21:15:02.649750",
						CompletedTs: "2018-01-30T00:15:18.162860",
					},
					expected: &expectedBuildFields{
						status: pb.Status_INFRA_FAILURE,
						startT: &timestamppb.Timestamp{Seconds: 1517260502, Nanos: 649750000},
						endT:   &timestamppb.Timestamp{Seconds: 1517271318, Nanos: 162860000},
					},
				},
				{
					fakeTaskResult: &swarming.SwarmingRpcsTaskResult{
						State:       "BOT_DIED",
						StartedTs:   "2018-01-29T21:15:02.649750",
						CompletedTs: "2018-01-30T00:15:18.162860",
					},
					expected: &expectedBuildFields{
						status: pb.Status_INFRA_FAILURE,
						startT: &timestamppb.Timestamp{Seconds: 1517260502, Nanos: 649750000},
						endT:   &timestamppb.Timestamp{Seconds: 1517271318, Nanos: 162860000},
					},
				},
				{
					fakeTaskResult: &swarming.SwarmingRpcsTaskResult{
						State:       "BOT_DIED",
						StartedTs:   "2018-01-29T21:15:02.649750",
						CompletedTs: "2018-01-30T00:15:17Z07:00", // invalid format, So sync flow fallbacks to use AbandonedTs
						AbandonedTs: "2018-01-30T00:15:18.162860",
					},
					expected: &expectedBuildFields{
						status: pb.Status_INFRA_FAILURE,
						startT: &timestamppb.Timestamp{Seconds: 1517260502, Nanos: 649750000},
						endT:   &timestamppb.Timestamp{Seconds: 1517271318, Nanos: 162860000},
					},
				},
				{
					fakeTaskResult: &swarming.SwarmingRpcsTaskResult{
						State:       "TIMED_OUT",
						StartedTs:   "2018-01-29T21:15:02.649750",
						CompletedTs: "2018-01-30T00:15:18.162860",
					},
					expected: &expectedBuildFields{
						status:    pb.Status_INFRA_FAILURE,
						isTimeOut: true,
						startT:    &timestamppb.Timestamp{Seconds: 1517260502, Nanos: 649750000},
						endT:      &timestamppb.Timestamp{Seconds: 1517271318, Nanos: 162860000},
					},
				},
				{
					fakeTaskResult: &swarming.SwarmingRpcsTaskResult{
						State:       "EXPIRED",
						AbandonedTs: "2018-01-30T00:15:18.162860",
					},
					expected: &expectedBuildFields{
						status:               pb.Status_INFRA_FAILURE,
						isResourceExhaustion: true,
						isTimeOut:            true,
						endT:                 &timestamppb.Timestamp{Seconds: 1517271318, Nanos: 162860000},
					},
				},
				{
					fakeTaskResult: &swarming.SwarmingRpcsTaskResult{
						State:       "KILLED",
						AbandonedTs: "2018-01-30T00:15:18.162860",
					},
					expected: &expectedBuildFields{
						status: pb.Status_CANCELED,
						endT:   &timestamppb.Timestamp{Seconds: 1517271318, Nanos: 162860000},
					},
				},
				{
					fakeTaskResult: &swarming.SwarmingRpcsTaskResult{
						State:       "NO_RESOURCE",
						AbandonedTs: "2018-01-30T00:15:18.162860",
					},
					expected: &expectedBuildFields{
						status:               pb.Status_INFRA_FAILURE,
						isResourceExhaustion: true,
						endT:                 &timestamppb.Timestamp{Seconds: 1517271318, Nanos: 162860000},
					},
				},
				{
					fakeTaskResult: &swarming.SwarmingRpcsTaskResult{
						State:       "NO_RESOURCE",
						AbandonedTs: "2015-11-29T00:15:18.162860",
					},
					expected: &expectedBuildFields{
						status:               pb.Status_INFRA_FAILURE,
						isResourceExhaustion: true,
						endT:                 &timestamppb.Timestamp{Seconds: now.UnixNano() / 1000000000},
					},
				},
			}
			for i, tCase := range cases {
				Convey(fmt.Sprintf("test %d - task %s", i, tCase.fakeTaskResult.State), func() {
					mockSwarm.EXPECT().GetTaskResult(ctx, "task_id").Return(tCase.fakeTaskResult, nil)
					err := SyncBuild(ctx, 123, 1)
					So(err, ShouldBeNil)
					syncedBuild := &model.Build{ID: 123}
					So(datastore.Get(ctx, syncedBuild), ShouldBeNil)
					So(syncedBuild.Status, ShouldEqual, tCase.expected.status)
					if tCase.expected.isResourceExhaustion {
						So(syncedBuild.Proto.StatusDetails.ResourceExhaustion, ShouldResembleProto, &pb.StatusDetails_ResourceExhaustion{})
					} else {
						So(syncedBuild.Proto.StatusDetails.GetResourceExhaustion(), ShouldBeNil)
					}
					if tCase.expected.isTimeOut {
						So(syncedBuild.Proto.StatusDetails.Timeout, ShouldResembleProto, &pb.StatusDetails_Timeout{})
					} else {
						So(syncedBuild.Proto.StatusDetails.GetTimeout(), ShouldBeNil)
					}
					if tCase.expected.startT != nil {
						So(syncedBuild.Proto.StartTime, ShouldResembleProto, tCase.expected.startT)
					}
					if tCase.expected.endT != nil {
						So(syncedBuild.Proto.EndTime, ShouldResembleProto, tCase.expected.endT)
					}
					if tCase.expected.botDimensions != nil {
						syncedInfra := &model.BuildInfra{Build: datastore.KeyForObj(ctx, syncedBuild)}
						So(datastore.Get(ctx, syncedInfra), ShouldBeNil)
						So(syncedInfra.Proto.Swarming.BotDimensions, ShouldResembleProto, tCase.expected.botDimensions)
					}
					if protoutil.IsEnded(syncedBuild.Status) {
						// FinalizeResultDB, ExportBigQuery, NotifyPubSub and a continuation sync task.
						So(sch.Tasks(), ShouldHaveLength, 4)
					} else if syncedBuild.Status == pb.Status_STARTED {
						// NotifyPubSub and a continuation sync task.
						So(sch.Tasks(), ShouldHaveLength, 2)
					}
				})
			}
		})
	})

}

func TestSubNotify(t *testing.T) {
	t.Parallel()
	Convey("SubNotify", t, func() {
		ctl := gomock.NewController(t)
		defer ctl.Finish()
		now := testclock.TestRecentTimeUTC
		mockSwarm := clients.NewMockSwarmingClient(ctl)
		ctx, _ := testclock.UseTime(context.Background(), now)
		ctx = context.WithValue(ctx, &clients.MockSwarmingClientKey, mockSwarm)
		ctx = memory.UseWithAppID(ctx, "dev~app-id")
		ctx = txndefer.FilterRDS(ctx)
		ctx = metrics.WithServiceInfo(ctx, "svc", "job", "ins")
		datastore.GetTestable(ctx).AutoIndex(true)
		datastore.GetTestable(ctx).Consistent(true)
		ctx, _ = tq.TestingContext(ctx, nil)
		ctx = cachingtest.WithGlobalCache(ctx, map[string]caching.BlobCache{
			"swarming-pubsub-msg-id": &cachingtest.BlobCache{
				LRU: lru.New(0),
			},
		})
		ctx, sch := tq.TestingContext(ctx, nil)

		b := &model.Build{
			ID: 123,
			Proto: &pb.Build{
				Id:         123,
				Status:     pb.Status_SCHEDULED,
				CreateTime: &timestamppb.Timestamp{Seconds: now.UnixNano() / 1000000000},
				SchedulingTimeout: &durationpb.Duration{
					Seconds: 3600,
				},
				ExecutionTimeout: &durationpb.Duration{
					Seconds: 4800,
				},
				GracePeriod: &durationpb.Duration{
					Seconds: 60,
				},
				Builder: &pb.BuilderID{
					Project: "proj",
					Bucket:  "bucket",
					Builder: "builder",
				},
			},
		}
		So(datastore.Put(ctx, b), ShouldBeNil)
		inf := &model.BuildInfra{
			ID:    1,
			Build: datastore.KeyForObj(ctx, &model.Build{ID: 123}),
			Proto: &pb.BuildInfra{
				Swarming: &pb.BuildInfra_Swarming{
					Hostname: "swarm",
					TaskId:   "task123",
					Caches: []*pb.BuildInfra_Swarming_CacheEntry{
						{Name: "shared_builder_cache", Path: "builder", WaitForWarmCache: &durationpb.Duration{Seconds: 60}},
						{Name: "second_cache", Path: "second", WaitForWarmCache: &durationpb.Duration{Seconds: 360}},
					},
					TaskDimensions: []*pb.RequestedDimension{
						{Key: "a", Value: "1", Expiration: &durationpb.Duration{Seconds: 120}},
						{Key: "a", Value: "2", Expiration: &durationpb.Duration{Seconds: 120}},
						{Key: "pool", Value: "Chrome"},
					},
				},
				Bbagent: &pb.BuildInfra_BBAgent{
					CacheDir:    "cache",
					PayloadPath: "kitchen-checkout",
				},
				Buildbucket: &pb.BuildInfra_Buildbucket{
					Agent: &pb.BuildInfra_Buildbucket_Agent{
						Source: &pb.BuildInfra_Buildbucket_Agent_Source{
							DataType: &pb.BuildInfra_Buildbucket_Agent_Source_Cipd{
								Cipd: &pb.BuildInfra_Buildbucket_Agent_Source_CIPD{
									Package: "infra/tools/luci/bbagent/${platform}",
									Version: "canary-version",
									Server:  "cipd server",
								},
							},
						},
					},
				},
			},
		}
		So(datastore.Put(ctx, inf), ShouldBeNil)

		Convey("bad msg data", func() {
			body := makeSwarmingPubsubMsg(&userdata{
				BuildID:          999,
				CreatedTS:        1448841600000,
				SwarmingHostname: "swarm",
			}, "", "msg1")
			err := SubNotify(ctx, body)
			So(err, ShouldErrLike, "task_id not found in message data")
			So(transient.Tag.In(err), ShouldBeFalse)

			body = makeSwarmingPubsubMsg(&userdata{
				CreatedTS:        1448841600000,
				SwarmingHostname: "swarm",
			}, "task123", "msg1")
			err = SubNotify(ctx, body)
			So(err, ShouldErrLike, "invalid build_id 0")

			body = makeSwarmingPubsubMsg(&userdata{
				BuildID:          999,
				SwarmingHostname: "swarm",
			}, "task123", "msg1")
			err = SubNotify(ctx, body)
			So(err, ShouldErrLike, "invalid created_ts 0")

			body = makeSwarmingPubsubMsg(&userdata{
				BuildID:          999,
				CreatedTS:        1448841600000,
				SwarmingHostname: " ",
			}, "task123", "msg1")
			err = SubNotify(ctx, body)
			So(err, ShouldErrLike, "swarming hostname not found in userdata")

			body = makeSwarmingPubsubMsg(&userdata{
				BuildID:          999,
				CreatedTS:        1448841600000,
				SwarmingHostname: "https://swarm.com",
			}, "task123", "msg1")
			err = SubNotify(ctx, body)
			So(err, ShouldErrLike, "swarming hostname https://swarm.com must not contain '://'")
		})

		Convey("build not found", func() {
			old := now.Add(-time.Minute).UnixNano() / int64(time.Microsecond)
			body := makeSwarmingPubsubMsg(&userdata{
				BuildID:          999,
				CreatedTS:        old,
				SwarmingHostname: "swarm",
			}, "task123", "msg1")
			err := SubNotify(ctx, body)
			So(err, ShouldErrLike, "Build 999 or BuildInfra for task https://swarm/task?id=task123 not found")
			So(transient.Tag.In(err), ShouldBeFalse)

			recent := now.Add(-50*time.Second).UnixNano() / int64(time.Microsecond)
			body = makeSwarmingPubsubMsg(&userdata{
				BuildID:          999,
				CreatedTS:        recent,
				SwarmingHostname: "swarm",
			}, "task123", "msg1")
			err = SubNotify(ctx, body)
			So(err, ShouldErrLike, "Build 999 or BuildInfra for task https://swarm/task?id=task123 not found yet")
			So(transient.Tag.In(err), ShouldBeTrue)
		})

		Convey("different swarming hostname", func() {

			body := makeSwarmingPubsubMsg(&userdata{
				BuildID:          123,
				CreatedTS:        1517260502000000,
				SwarmingHostname: "swarm2",
			}, "task123", "msg1")
			err := SubNotify(ctx, body)
			So(err, ShouldErrLike, "swarming_hostname swarm of build 123 does not match swarm2")
			So(transient.Tag.In(err), ShouldBeFalse)
		})

		Convey("different task id", func() {
			body := makeSwarmingPubsubMsg(&userdata{
				BuildID:          123,
				CreatedTS:        1517260502000000,
				SwarmingHostname: "swarm",
			}, "task345", "msg1")
			err := SubNotify(ctx, body)
			So(err, ShouldErrLike, "swarming_task_id task123 of build 123 does not match task345")
			So(transient.Tag.In(err), ShouldBeFalse)
		})

		Convey("swarming 500s error", func() {
			body := makeSwarmingPubsubMsg(&userdata{
				BuildID:          123,
				CreatedTS:        1517260502000000,
				SwarmingHostname: "swarm",
			}, "task123", "msg1")
			mockSwarm.EXPECT().GetTaskResult(ctx, "task123").Return(nil, &googleapi.Error{Code: 500, Message: "swarming internal error"})
			err := SubNotify(ctx, body)
			So(err, ShouldErrLike, "googleapi: Error 500: swarming internal error")
			So(transient.Tag.In(err), ShouldBeTrue)

			cache := caching.GlobalCache(ctx, "swarming-pubsub-msg-id")
			_, err = cache.Get(ctx, "msg1")
			So(err, ShouldEqual, caching.ErrCacheMiss)
		})

		Convey("status already ended", func() {
			b.Proto.Status = pb.Status_SUCCESS
			So(datastore.Put(ctx, b), ShouldBeNil)

			body := makeSwarmingPubsubMsg(&userdata{
				BuildID:          123,
				CreatedTS:        1517260502000000,
				SwarmingHostname: "swarm",
			}, "task123", "msg1")
			err := SubNotify(ctx, body)
			So(err, ShouldBeNil)
			mockSwarm.EXPECT().CreateTask(gomock.Any(), gomock.Any()).Times(0)

			So(sch.Tasks(), ShouldHaveLength, 0)
		})

		Convey("status changed to success", func() {
			body := makeSwarmingPubsubMsg(&userdata{
				BuildID:          123,
				CreatedTS:        1517260502000000,
				SwarmingHostname: "swarm",
			}, "task123", "msg1")
			mockSwarm.EXPECT().GetTaskResult(ctx, "task123").Return(&swarming.SwarmingRpcsTaskResult{
				State:       "COMPLETED",
				StartedTs:   "2018-01-29T21:15:02.649750",
				CompletedTs: "2018-01-30T00:15:18.162860",
				BotDimensions: []*swarming.SwarmingRpcsStringListPair{
					{
						Key:   "new_key",
						Value: []string{"new_val"},
					},
				},
			}, nil)
			err := SubNotify(ctx, body)
			So(err, ShouldBeNil)
			syncedBuild := &model.Build{ID: 123}
			So(datastore.Get(ctx, syncedBuild), ShouldBeNil)
			So(syncedBuild.Status, ShouldEqual, pb.Status_SUCCESS)
			So(syncedBuild.Proto.StartTime, ShouldResembleProto, &timestamppb.Timestamp{Seconds: 1517260502, Nanos: 649750000})
			So(syncedBuild.Proto.EndTime, ShouldResembleProto, &timestamppb.Timestamp{Seconds: 1517271318, Nanos: 162860000})

			syncedInfra := &model.BuildInfra{Build: datastore.KeyForObj(ctx, syncedBuild)}
			So(datastore.Get(ctx, syncedInfra), ShouldBeNil)
			So(syncedInfra.Proto.Swarming.BotDimensions, ShouldResembleProto, []*pb.StringPair{
				{
					Key:   "new_key",
					Value: "new_val",
				},
			})

			cache := caching.GlobalCache(ctx, "swarming-pubsub-msg-id")
			cached, err := cache.Get(ctx, "msg1")
			So(err, ShouldBeNil)
			So(cached, ShouldResemble, []byte{1})

			// FinalizeResultDB, ExportBigQuery, NotifyPubSub tasks.
			So(sch.Tasks(), ShouldHaveLength, 3)
		})

		Convey("status unchanged(in STARTED) while bot dimensions changed", func() {
			b.Proto.Status = pb.Status_STARTED
			So(datastore.Put(ctx, b), ShouldBeNil)
			body := makeSwarmingPubsubMsg(&userdata{
				BuildID:          123,
				CreatedTS:        1517260502000000,
				SwarmingHostname: "swarm",
			}, "task123", "msg1")
			mockSwarm.EXPECT().GetTaskResult(ctx, "task123").Return(&swarming.SwarmingRpcsTaskResult{
				State:     "RUNNING",
				StartedTs: "2018-01-29T21:15:02.649750",
				BotDimensions: []*swarming.SwarmingRpcsStringListPair{
					{
						Key:   "new_key",
						Value: []string{"new_val"},
					},
				},
			}, nil)
			err := SubNotify(ctx, body)
			So(err, ShouldBeNil)
			syncedBuild := &model.Build{ID: 123}
			So(datastore.Get(ctx, syncedBuild), ShouldBeNil)
			So(syncedBuild.Status, ShouldEqual, pb.Status_STARTED)

			syncedInfra := &model.BuildInfra{Build: datastore.KeyForObj(ctx, syncedBuild)}
			So(datastore.Get(ctx, syncedInfra), ShouldBeNil)
			So(syncedInfra.Proto.Swarming.BotDimensions, ShouldResembleProto, []*pb.StringPair{{
				Key:   "new_key",
				Value: "new_val",
			}})

			So(sch.Tasks(), ShouldHaveLength, 0)
		})

		Convey("status unchanged(not in STARTED) while bot dimensions changed", func() {
			b.Proto.Status = pb.Status_STARTED
			So(datastore.Put(ctx, b), ShouldBeNil)
			body := makeSwarmingPubsubMsg(&userdata{
				BuildID:          123,
				CreatedTS:        1517260502000000,
				SwarmingHostname: "swarm",
			}, "task123", "msg1")
			mockSwarm.EXPECT().GetTaskResult(ctx, "task123").Return(&swarming.SwarmingRpcsTaskResult{
				State: "PENDING",
				BotDimensions: []*swarming.SwarmingRpcsStringListPair{
					{
						Key:   "new_key",
						Value: []string{"new_val"},
					},
				},
			}, nil)
			err := SubNotify(ctx, body)
			So(err, ShouldBeNil)
			syncedBuild := &model.Build{ID: 123}
			So(datastore.Get(ctx, syncedBuild), ShouldBeNil)
			So(syncedBuild.Status, ShouldEqual, pb.Status_STARTED)

			currentInfra := &model.BuildInfra{Build: datastore.KeyForObj(ctx, syncedBuild)}
			So(datastore.Get(ctx, currentInfra), ShouldBeNil)
			So(currentInfra.Proto.Swarming.BotDimensions, ShouldBeEmpty)

			So(sch.Tasks(), ShouldHaveLength, 0)
		})

		Convey("duplicate message", func() {
			cache := caching.GlobalCache(ctx, "swarming-pubsub-msg-id")
			err := cache.Set(ctx, "msg123", []byte{1}, 0*time.Second)
			So(err, ShouldBeNil)

			body := makeSwarmingPubsubMsg(&userdata{
				BuildID:          123,
				CreatedTS:        1517260502000000,
				SwarmingHostname: "swarm",
			}, "task123", "msg123")
			mockSwarm.EXPECT().GetTaskResult(ctx, gomock.Any()).Times(0)
			err = SubNotify(ctx, body)
			So(err, ShouldBeNil)
		})
	})
}

func makeSwarmingPubsubMsg(userdata *userdata, taskID string, msgID string) io.Reader {
	ud, _ := json.Marshal(userdata)
	data := struct {
		TaskID   string `json:"task_id"`
		Userdata string `json:"userdata"`
	}{TaskID: taskID, Userdata: string(ud)}
	bd, _ := json.Marshal(data)
	msg := struct {
		Message struct {
			Data      string
			MessageID string
		}
	}{struct {
		Data      string
		MessageID string
	}{Data: base64.StdEncoding.EncodeToString(bd), MessageID: msgID}}
	jmsg, _ := json.Marshal(msg)
	return bytes.NewReader(jmsg)
}

type expectedBuildFields struct {
	status               pb.Status
	startT               *timestamppb.Timestamp
	endT                 *timestamppb.Timestamp
	isTimeOut            bool
	isResourceExhaustion bool
	botDimensions        []*pb.StringPair
}
