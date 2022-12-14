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

package tasks

import (
	"context"
	"sort"
	"testing"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"

	"go.chromium.org/luci/common/retry/transient"
	"go.chromium.org/luci/gae/filter/txndefer"
	"go.chromium.org/luci/gae/impl/memory"
	"go.chromium.org/luci/gae/service/datastore"
	"go.chromium.org/luci/server/auth"
	"go.chromium.org/luci/server/auth/authtest"
	"go.chromium.org/luci/server/tq"
	"go.chromium.org/luci/server/tq/tqtesting"

	"go.chromium.org/luci/buildbucket/appengine/internal/clients"
	"go.chromium.org/luci/buildbucket/appengine/internal/compression"
	"go.chromium.org/luci/buildbucket/appengine/model"
	taskdefs "go.chromium.org/luci/buildbucket/appengine/tasks/defs"
	pb "go.chromium.org/luci/buildbucket/proto"

	. "github.com/smartystreets/goconvey/convey"

	. "go.chromium.org/luci/common/testing/assertions"
)

func TestNotification(t *testing.T) {
	t.Parallel()

	Convey("notifyPubsub", t, func() {
		ctx := auth.WithState(memory.Use(context.Background()), &authtest.FakeState{})
		ctx = txndefer.FilterRDS(ctx)
		ctx, sch := tq.TestingContext(ctx, nil)
		datastore.GetTestable(ctx).AutoIndex(true)
		datastore.GetTestable(ctx).Consistent(true)
		So(datastore.Put(ctx, &model.Project{
			ID: "project_with_external_topics",
			CommonConfig: &pb.BuildbucketCfg_CommonConfig{
				BuildsNotificationTopics: []*pb.BuildbucketCfg_Topic{
					{
						Name: "projects/my-cloud-project/topics/my-topic",
					},
				},
			},
		}), ShouldBeNil)

		sortTasksByClassName := func(tasks tqtesting.TaskList) {
			sort.Slice(tasks, func(i, j int) bool {
				return tasks[i].Class < tasks[j].Class
			})
		}

		Convey("w/o callback", func() {
			txErr := datastore.RunInTransaction(ctx, func(ctx context.Context) error {
				return NotifyPubSub(ctx, &model.Build{
					ID: 123,
					Proto: &pb.Build{
						Builder: &pb.BuilderID{
							Project: "project_no_external_topics",
						},
					},
				})
			}, nil)
			So(txErr, ShouldBeNil)
			tasks := sch.Tasks()
			So(tasks, ShouldHaveLength, 1)
			So(tasks[0].Payload.(*taskdefs.NotifyPubSub).GetBuildId(), ShouldEqual, 123)
			So(tasks[0].Payload.(*taskdefs.NotifyPubSub).GetCallback(), ShouldBeFalse)
		})

		Convey("w/ callback", func() {
			txErr := datastore.RunInTransaction(ctx, func(ctx context.Context) error {
				cb := model.PubSubCallback{
					AuthToken: "token",
					Topic:     "topic",
					UserData:  []byte("user_data"),
				}
				return NotifyPubSub(ctx, &model.Build{
					ID: 123,
					PubSubCallback: cb,
					Proto: &pb.Build{
						Builder: &pb.BuilderID{
							Project: "project_no_external_topics",
						},
					},
				})
			}, nil)
			So(txErr, ShouldBeNil)
			tasks := sch.Tasks()
			So(tasks, ShouldHaveLength, 2)

			n1 := tasks[0].Payload.(*taskdefs.NotifyPubSub)
			n2 := tasks[1].Payload.(*taskdefs.NotifyPubSub)
			So(n1.GetBuildId(), ShouldEqual, 123)
			So(n2.GetBuildId(), ShouldEqual, 123)
			// One w/ callback and one w/o callback.
			So(n1.GetCallback() != n2.GetCallback(), ShouldBeTrue)
		})

		Convey("has external topics", func() {
			txErr := datastore.RunInTransaction(ctx, func(ctx context.Context) error {
				return NotifyPubSub(ctx, &model.Build{
					ID: 123,
					Proto: &pb.Build{
						Builder: &pb.BuilderID{
							Project: "project_with_external_topics",
						},
					},
				})
			}, nil)
			So(txErr, ShouldBeNil)
			tasks := sch.Tasks()
			sortTasksByClassName(tasks)
			So(tasks, ShouldHaveLength, 2)
			So(tasks[0].Payload.(*taskdefs.NotifyPubSub).GetBuildId(), ShouldEqual, 123)
			So(tasks[0].Payload.(*taskdefs.NotifyPubSub).GetCallback(), ShouldBeFalse)
			taskGo1 := tasks[1].Payload.(*taskdefs.NotifyPubSubGo)
			So(taskGo1.BuildId, ShouldEqual, 123)
			So(taskGo1.Topic, ShouldResembleProto, &pb.BuildbucketCfg_Topic{Name: "projects/my-cloud-project/topics/my-topic"})
		})

		Convey("empty project.common_config", func() {
			So(datastore.Put(ctx, &model.Project{
				ID: "project_empty",
				CommonConfig: &pb.BuildbucketCfg_CommonConfig{},
			}), ShouldBeNil)
			txErr := datastore.RunInTransaction(ctx, func(ctx context.Context) error {
				return NotifyPubSub(ctx, &model.Build{
					ID: 123,
					Proto: &pb.Build{
						Builder: &pb.BuilderID{
							Project: "project_empty",
						},
					},
				})
			}, nil)
			So(txErr, ShouldBeNil)
			tasks := sch.Tasks()
			So(tasks, ShouldHaveLength, 1)
			So(tasks[0].Payload.(*taskdefs.NotifyPubSub).GetBuildId(), ShouldEqual, 123)
			So(tasks[0].Payload.(*taskdefs.NotifyPubSub).GetCallback(), ShouldBeFalse)
		})
	})

	Convey("PublishBuildsV2Notification", t, func() {
		ctx := auth.WithState(memory.Use(context.Background()), &authtest.FakeState{})
		ctx = txndefer.FilterRDS(ctx)
		ctx, sch := tq.TestingContext(ctx, nil)
		datastore.GetTestable(ctx).AutoIndex(true)
		datastore.GetTestable(ctx).Consistent(true)

		b := &model.Build{
			ID: 123,
			Proto: &pb.Build{
				Id: 123,
				Builder: &pb.BuilderID{
					Project: "project",
					Bucket:  "bucket",
					Builder: "builder",
				},
				Status: pb.Status_CANCELED,
			},
		}
		bk := datastore.KeyForObj(ctx, b)
		bsBytes, err := proto.Marshal(&pb.Build{
			Steps: []*pb.Step{
				{
					Name:            "step",
					SummaryMarkdown: "summary",
					Logs: []*pb.Log{{
						Name:    "log1",
						Url:     "url",
						ViewUrl: "view_url",
					},
					},
				},
			},
		})
		So(err, ShouldBeNil)
		bs := &model.BuildSteps{ID: 1, Build: bk, Bytes: bsBytes}
		bi := &model.BuildInfra{
			ID:    1,
			Build: bk,
			Proto: &pb.BuildInfra{
				Buildbucket: &pb.BuildInfra_Buildbucket{
					Hostname: "hostname",
				},
			},
		}
		bo := &model.BuildOutputProperties{
			Build: bk,
			Proto: &structpb.Struct{
				Fields: map[string]*structpb.Value{
					"output": {
						Kind: &structpb.Value_StringValue{
							StringValue: "output value",
						},
					},
				},
			},
		}
		binpProp := &model.BuildInputProperties{
			Build: bk,
			Proto: &structpb.Struct{
				Fields: map[string]*structpb.Value{
					"input": {
						Kind: &structpb.Value_StringValue{
							StringValue: "input value",
						},
					},
				},
			},
		}
		So(datastore.Put(ctx, b, bi, bs, bo, binpProp), ShouldBeNil)

		Convey("build not exist", func() {
			err := PublishBuildsV2Notification(ctx, 999, nil)
			So(err, ShouldBeNil)
			tasks := sch.Tasks()
			So(tasks, ShouldHaveLength, 0)
		})

		Convey("To internal topic", func() {

			Convey("success", func() {
				err := PublishBuildsV2Notification(ctx, 123, nil)
				So(err, ShouldBeNil)

				tasks := sch.Tasks()
				So(tasks, ShouldHaveLength, 1)
				So(tasks[0].Message.Attributes["project"], ShouldEqual, "project")
				So(tasks[0].Message.Attributes["is_completed"], ShouldEqual, "true")
				So(tasks[0].Payload.(*taskdefs.BuildsV2PubSub).GetBuild(), ShouldResembleProto, &pb.Build{
					Id: 123,
					Builder: &pb.BuilderID{
						Project: "project",
						Bucket:  "bucket",
						Builder: "builder",
					},
					Status: pb.Status_CANCELED,
					Infra: &pb.BuildInfra{
						Buildbucket: &pb.BuildInfra_Buildbucket{
							Hostname: "hostname",
						},
					},
					Input:  &pb.Build_Input{},
					Output: &pb.Build_Output{},
				})
				So(tasks[0].Payload.(*taskdefs.BuildsV2PubSub).GetBuildLargeFields(), ShouldNotBeNil)
				bLargeBytes := tasks[0].Payload.(*taskdefs.BuildsV2PubSub).GetBuildLargeFields()
				buildLarge, err := zlibUncompressBuild(bLargeBytes)
				So(err, ShouldBeNil)
				So(buildLarge, ShouldResembleProto, &pb.Build{
					Steps: []*pb.Step{
						{
							Name:            "step",
							SummaryMarkdown: "summary",
							Logs: []*pb.Log{{
								Name:    "log1",
								Url:     "url",
								ViewUrl: "view_url",
							},
							},
						},
					},
					Input: &pb.Build_Input{
						Properties: &structpb.Struct{
							Fields: map[string]*structpb.Value{
								"input": {
									Kind: &structpb.Value_StringValue{
										StringValue: "input value",
									},
								},
							},
						},
					},
					Output: &pb.Build_Output{
						Properties: &structpb.Struct{
							Fields: map[string]*structpb.Value{
								"output": {
									Kind: &structpb.Value_StringValue{
										StringValue: "output value",
									},
								},
							},
						},
					},
				})
			})

			Convey("success - no large fields", func() {
				b := &model.Build{
					ID: 456,
					Proto: &pb.Build{
						Id: 456,
						Builder: &pb.BuilderID{
							Project: "project",
							Bucket:  "bucket",
							Builder: "builder",
						},
						Status: pb.Status_CANCELED,
					},
				}
				bk := datastore.KeyForObj(ctx, b)
				bi := &model.BuildInfra{
					ID:    1,
					Build: bk,
					Proto: &pb.BuildInfra{
						Buildbucket: &pb.BuildInfra_Buildbucket{
							Hostname: "hostname",
						},
					},
				}
				So(datastore.Put(ctx, b, bi), ShouldBeNil)

				err := PublishBuildsV2Notification(ctx, 456, nil)
				So(err, ShouldBeNil)

				tasks := sch.Tasks()
				So(tasks, ShouldHaveLength, 1)
				So(tasks[0].Payload.(*taskdefs.BuildsV2PubSub).GetBuild(), ShouldResembleProto, &pb.Build{
					Id: 456,
					Builder: &pb.BuilderID{
						Project: "project",
						Bucket:  "bucket",
						Builder: "builder",
					},
					Status: pb.Status_CANCELED,
					Infra: &pb.BuildInfra{
						Buildbucket: &pb.BuildInfra_Buildbucket{
							Hostname: "hostname",
						},
					},
					Input:  &pb.Build_Input{},
					Output: &pb.Build_Output{},
				})
				So(tasks[0].Payload.(*taskdefs.BuildsV2PubSub).GetBuildLargeFields(), ShouldNotBeNil)
				bLargeBytes := tasks[0].Payload.(*taskdefs.BuildsV2PubSub).GetBuildLargeFields()
				buildLarge, err := zlibUncompressBuild(bLargeBytes)
				So(err, ShouldBeNil)
				So(buildLarge, ShouldResembleProto, &pb.Build{
					Input:  &pb.Build_Input{},
					Output: &pb.Build_Output{},
				})
			})
		})

		Convey("To external topic", func() {
			ctx, psserver, psclient, err := clients.SetupTestPubsub(ctx, "my-cloud-project")
			So(err, ShouldBeNil)
			defer func() {
				psclient.Close()
				psserver.Close()
			}()
			_, err = psclient.CreateTopic(ctx, "my-topic")
			So(err, ShouldBeNil)

			Convey("success (zlib compression)", func() {
				err := PublishBuildsV2Notification(ctx, 123, &pb.BuildbucketCfg_Topic{Name: "projects/my-cloud-project/topics/my-topic"})
				So(err, ShouldBeNil)

				tasks := sch.Tasks()
				So(tasks, ShouldHaveLength, 0)
				So(psserver.Messages(), ShouldHaveLength, 1)
				publishedMsg := psserver.Messages()[0]

				So(publishedMsg.Attributes["project"], ShouldEqual, "project")
				So(publishedMsg.Attributes["bucket"], ShouldEqual, "bucket")
				So(publishedMsg.Attributes["builder"], ShouldEqual, "builder")
				So(publishedMsg.Attributes["is_completed"], ShouldEqual, "true")
				buildMsg := &taskdefs.BuildsV2PubSub{}
				err = protojson.Unmarshal(publishedMsg.Data, buildMsg)
				So(err, ShouldBeNil)
				So(buildMsg.Build, ShouldResembleProto, &pb.Build{
					Id: 123,
					Builder: &pb.BuilderID{
						Project: "project",
						Bucket:  "bucket",
						Builder: "builder",
					},
					Status: pb.Status_CANCELED,
					Infra: &pb.BuildInfra{
						Buildbucket: &pb.BuildInfra_Buildbucket{
							Hostname: "hostname",
						},
					},
					Input:  &pb.Build_Input{},
					Output: &pb.Build_Output{},
				})
				So(buildMsg.BuildLargeFields, ShouldNotBeNil)
				buildLarge, err := zlibUncompressBuild(buildMsg.BuildLargeFields)
				So(err, ShouldBeNil)
				So(buildLarge, ShouldResembleProto, &pb.Build{
					Steps: []*pb.Step{
						{
							Name:            "step",
							SummaryMarkdown: "summary",
							Logs: []*pb.Log{{
								Name:    "log1",
								Url:     "url",
								ViewUrl: "view_url",
							},
							},
						},
					},
					Input: &pb.Build_Input{
						Properties: &structpb.Struct{
							Fields: map[string]*structpb.Value{
								"input": {
									Kind: &structpb.Value_StringValue{
										StringValue: "input value",
									},
								},
							},
						},
					},
					Output: &pb.Build_Output{
						Properties: &structpb.Struct{
							Fields: map[string]*structpb.Value{
								"output": {
									Kind: &structpb.Value_StringValue{
										StringValue: "output value",
									},
								},
							},
						},
					},
				})
			})

			Convey("success (zstd compression)", func() {
				err := PublishBuildsV2Notification(ctx, 123, &pb.BuildbucketCfg_Topic{
					Name:        "projects/my-cloud-project/topics/my-topic",
					Compression: pb.Compression_ZSTD,
				})
				So(err, ShouldBeNil)

				tasks := sch.Tasks()
				So(tasks, ShouldHaveLength, 0)
				So(psserver.Messages(), ShouldHaveLength, 1)
				publishedMsg := psserver.Messages()[0]

				So(publishedMsg.Attributes["project"], ShouldEqual, "project")
				So(publishedMsg.Attributes["bucket"], ShouldEqual, "bucket")
				So(publishedMsg.Attributes["builder"], ShouldEqual, "builder")
				So(publishedMsg.Attributes["is_completed"], ShouldEqual, "true")
				buildMsg := &taskdefs.BuildsV2PubSub{}
				err = protojson.Unmarshal(publishedMsg.Data, buildMsg)
				So(err, ShouldBeNil)
				So(buildMsg.Build, ShouldResembleProto, &pb.Build{
					Id: 123,
					Builder: &pb.BuilderID{
						Project: "project",
						Bucket:  "bucket",
						Builder: "builder",
					},
					Status: pb.Status_CANCELED,
					Infra: &pb.BuildInfra{
						Buildbucket: &pb.BuildInfra_Buildbucket{
							Hostname: "hostname",
						},
					},
					Input:  &pb.Build_Input{},
					Output: &pb.Build_Output{},
				})
				So(buildMsg.BuildLargeFields, ShouldNotBeNil)
				So(buildMsg.Compression, ShouldEqual, pb.Compression_ZSTD)
				buildLarge, err := zstdUncompressBuild(buildMsg.BuildLargeFields)
				So(err, ShouldBeNil)
				So(buildLarge, ShouldResembleProto, &pb.Build{
					Steps: []*pb.Step{
						{
							Name:            "step",
							SummaryMarkdown: "summary",
							Logs: []*pb.Log{{
								Name:    "log1",
								Url:     "url",
								ViewUrl: "view_url",
							},
							},
						},
					},
					Input: &pb.Build_Input{
						Properties: &structpb.Struct{
							Fields: map[string]*structpb.Value{
								"input": {
									Kind: &structpb.Value_StringValue{
										StringValue: "input value",
									},
								},
							},
						},
					},
					Output: &pb.Build_Output{
						Properties: &structpb.Struct{
							Fields: map[string]*structpb.Value{
								"output": {
									Kind: &structpb.Value_StringValue{
										StringValue: "output value",
									},
								},
							},
						},
					},
				})
			})

			Convey("non-exist topic", func() {
				err := PublishBuildsV2Notification(ctx, 123, &pb.BuildbucketCfg_Topic{
					Name: "projects/my-cloud-project/topics/non-exist-topic",
				})
				So(err, ShouldNotBeNil)
				So(transient.Tag.In(err), ShouldBeTrue)
			})
		})
	})
}

func zlibUncompressBuild(compressed []byte) (*pb.Build, error) {
	originalData, err := compression.ZlibDecompress(compressed)
	if err != nil {
		return nil, err
	}
	b := &pb.Build{}
	if err := proto.Unmarshal(originalData, b); err != nil {
		return nil, err
	}
	return b, nil
}

func zstdUncompressBuild(compressed []byte) (*pb.Build, error) {
	originalData, err := compression.ZstdDecompress(compressed, nil)
	if err != nil {
		return nil, err
	}
	b := &pb.Build{}
	if err := proto.Unmarshal(originalData, b); err != nil {
		return nil, err
	}
	return b, nil
}
