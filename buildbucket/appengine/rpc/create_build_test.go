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
	"testing"

	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/structpb"

	"go.chromium.org/luci/common/data/stringset"
	"go.chromium.org/luci/gae/impl/memory"
	"go.chromium.org/luci/gae/service/datastore"

	"go.chromium.org/luci/buildbucket/appengine/model"
	pb "go.chromium.org/luci/buildbucket/proto"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"
)

func TestValidateCreateBuildRequest(t *testing.T) {
	t.Parallel()
	Convey("validateCreateBuildRequest", t, func() {
		req := &pb.CreateBuildRequest{
			Build: &pb.Build{
				Builder: &pb.BuilderID{
					Project: "project",
					Bucket:  "bucket",
					Builder: "builder",
				},
				Infra: &pb.BuildInfra{
					Bbagent: &pb.BuildInfra_BBAgent{
						PayloadPath: "kitchen-checkout",
						CacheDir:    "cache",
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
							Input: &pb.BuildInfra_Buildbucket_Agent_Input{
								Data: map[string]*pb.InputDataRef{
									"path_a": {
										DataType: &pb.InputDataRef_Cipd{
											Cipd: &pb.InputDataRef_CIPD{
												Specs: []*pb.InputDataRef_CIPD_PkgSpec{{Package: "pkg_a", Version: "latest"}},
											},
										},
										OnPath: []string{"path_a/bin", "path_a"},
									},
									"path_b": {
										DataType: &pb.InputDataRef_Cipd{
											Cipd: &pb.InputDataRef_CIPD{
												Specs: []*pb.InputDataRef_CIPD_PkgSpec{{Package: "pkg_b", Version: "latest"}},
											},
										},
										OnPath: []string{"path_b/bin", "path_b"},
									},
								},
							},
						},
					},
					Swarming: &pb.BuildInfra_Swarming{
						Hostname: "host",
						Priority: 25,

						TaskDimensions: []*pb.RequestedDimension{
							{
								Key:   "pool",
								Value: "example.pool",
							},
						},
						TaskServiceAccount: "example@account.com",
						Caches: []*pb.BuildInfra_Swarming_CacheEntry{
							{
								Name: "builder_1809c38861a9996b1748e4640234fbd089992359f6f23f62f68deb98528f5f2b_v2",
								Path: "builder",
								WaitForWarmCache: &durationpb.Duration{
									Seconds: 240,
								},
							},
						},
					},
					Logdog: &pb.BuildInfra_LogDog{
						Hostname: "host",
						Project:  "project",
					},
					Resultdb: &pb.BuildInfra_ResultDB{
						Hostname: "host",
					},
				},
				Input: &pb.Build_Input{
					Properties: &structpb.Struct{
						Fields: map[string]*structpb.Value{
							"key": {
								Kind: &structpb.Value_StringValue{
									StringValue: "value",
								},
							},
						},
					},
					GerritChanges: []*pb.GerritChange{
						{
							Host:     "h1",
							Project:  "b",
							Change:   1,
							Patchset: 1,
						},
					},
					Experiments: []string{"customized.exp.name", "luci.wellknown.exp"},
				},
				Exe: &pb.Executable{
					Cmd: []string{"recipes"},
				},
			},
			RequestId: "request_id",
		}
		wellknownExps := stringset.NewFromSlice("luci.wellknown.exp")
		ctx := memory.Use(context.Background())
		datastore.GetTestable(ctx).AutoIndex(true)
		datastore.GetTestable(ctx).Consistent(true)

		So(datastore.Put(ctx, &model.Bucket{
			ID:     "bucket",
			Parent: model.ProjectKey(ctx, "project"),
			Proto: &pb.Bucket{
				Acls: []*pb.Acl{
					{
						Identity: "user:caller@example.com",
						Role:     pb.Acl_SCHEDULER,
					},
				},
				Name: "bucket",
				Constraints: &pb.Bucket_Constraints{
					Pools:           []string{"example.pool"},
					ServiceAccounts: []string{"example@account.com"},
				},
			},
		}), ShouldBeNil)

		Convey("works", func() {
			_, err := validateCreateBuildRequest(ctx, wellknownExps, req)
			So(err, ShouldBeNil)
		})

		Convey("mask", func() {
			req.Mask = &pb.BuildMask{
				Fields: &fieldmaskpb.FieldMask{
					Paths: []string{
						"invalid",
					},
				},
			}
			_, err := validateCreateBuildRequest(ctx, wellknownExps, req)
			So(err, ShouldErrLike, `invalid mask`)
		})

		Convey("RequestID", func() {
			req.RequestId = "request/id"
			_, err := validateCreateBuildRequest(ctx, wellknownExps, req)
			So(err, ShouldErrLike, `request_id cannot contain '/'`)
		})

		Convey("Build", func() {
			Convey("Builder", func() {
				Convey("invalid Builder", func() {
					req.Build.Builder = &pb.BuilderID{
						Project: "project",
					}
					_, err := validateCreateBuildRequest(ctx, wellknownExps, req)
					So(err, ShouldErrLike, `build: builder: bucket is required`)
				})
				Convey("w/o Builder", func() {
					req.Build.Builder = nil
					_, err := validateCreateBuildRequest(ctx, wellknownExps, req)
					So(err, ShouldErrLike, `.build.builder: required`)
				})
			})

			Convey("Exe", func() {
				Convey("cipd_package", func() {
					req.Build.Exe = &pb.Executable{
						CipdPackage: "Invalid",
					}
					_, err := validateCreateBuildRequest(ctx, wellknownExps, req)
					So(err, ShouldErrLike, `build: exe: cipd_package`)
				})
				Convey("cipd_version", func() {
					req.Build.Exe = &pb.Executable{
						CipdPackage: "valid/package/name",
						CipdVersion: "+100",
					}
					_, err := validateCreateBuildRequest(ctx, wellknownExps, req)
					So(err, ShouldErrLike, `build: exe: cipd_version`)
				})
				Convey("exe doesn't match agent", func() {
					req.Build.Exe = &pb.Executable{
						CipdPackage: "valid/package/name",
						CipdVersion: "version",
					}
					req.Build.Infra.Buildbucket.Agent.Purposes = map[string]pb.BuildInfra_Buildbucket_Agent_Purpose{
						"payload_path": pb.BuildInfra_Buildbucket_Agent_PURPOSE_EXE_PAYLOAD,
					}
					Convey("payload in agentInput is not a CIPD Package", func() {
						req.Build.Infra.Buildbucket.Agent.Input = &pb.BuildInfra_Buildbucket_Agent_Input{
							Data: map[string]*pb.InputDataRef{
								"payload_path": {
									DataType: &pb.InputDataRef_Cas{
										Cas: &pb.InputDataRef_CAS{
											CasInstance: "projects/project/instances/instance",
											Digest: &pb.InputDataRef_CAS_Digest{
												Hash:      "hash",
												SizeBytes: 1,
											},
										},
									},
								},
							},
						}
						_, err := validateCreateBuildRequest(ctx, wellknownExps, req)
						So(err, ShouldErrLike, `build: exe: not match build.infra.buildbucket.agent`)
					})
					Convey("different package", func() {
						req.Build.Infra.Buildbucket.Agent.Input = &pb.BuildInfra_Buildbucket_Agent_Input{
							Data: map[string]*pb.InputDataRef{
								"payload_path": {
									DataType: &pb.InputDataRef_Cipd{
										Cipd: &pb.InputDataRef_CIPD{
											Specs: []*pb.InputDataRef_CIPD_PkgSpec{{Package: "another", Version: "latest"}},
										},
									},
								},
							},
						}
						_, err := validateCreateBuildRequest(ctx, wellknownExps, req)
						So(err, ShouldErrLike, `build: exe: cipd_package does not match build.infra.buildbucket.agent`)
					})
					Convey("different version", func() {
						req.Build.Infra.Buildbucket.Agent.Input = &pb.BuildInfra_Buildbucket_Agent_Input{
							Data: map[string]*pb.InputDataRef{
								"payload_path": {
									DataType: &pb.InputDataRef_Cipd{
										Cipd: &pb.InputDataRef_CIPD{
											Specs: []*pb.InputDataRef_CIPD_PkgSpec{{Package: "valid/package/name", Version: "latest"}},
										},
									},
								},
							},
						}
						_, err := validateCreateBuildRequest(ctx, wellknownExps, req)
						So(err, ShouldErrLike, `build: exe: cipd_version does not match build.infra.buildbucket.agent`)
					})
				})
			})

			Convey("Input", func() {
				Convey("gerrit_changes", func() {
					req.Build.Input.GerritChanges = []*pb.GerritChange{
						{
							Host: "host",
						},
					}
					_, err := validateCreateBuildRequest(ctx, wellknownExps, req)
					So(err, ShouldErrLike, `build: input: gerrit_changes`)
				})
				Convey("gitiles_commit", func() {
					req.Build.Input.GitilesCommit = &pb.GitilesCommit{
						Host: "host",
					}
					_, err := validateCreateBuildRequest(ctx, wellknownExps, req)
					So(err, ShouldErrLike, `build: input: gitiles_commit`)
				})
				Convey("properties", func() {
					req.Build.Input.Properties = &structpb.Struct{
						Fields: map[string]*structpb.Value{
							"$recipe_engine/buildbucket": {
								Kind: &structpb.Value_StringValue{
									StringValue: "value",
								},
							},
						},
					}
					_, err := validateCreateBuildRequest(ctx, wellknownExps, req)
					So(err, ShouldErrLike, `build: input: properties`)
				})
				Convey("experiments", func() {
					req.Build.Input.Experiments = []string{"luci.not.wellknown"}
					_, err := validateCreateBuildRequest(ctx, wellknownExps, req)
					So(err, ShouldErrLike, `build: input: experiment "luci.not.wellknown"`)
				})
			})

			Convey("Tags", func() {
				req.Build.Tags = []*pb.StringPair{
					{
						Key:   "build_address",
						Value: "value2",
					},
				}
				_, err := validateCreateBuildRequest(ctx, wellknownExps, req)
				So(err, ShouldErrLike, `build: tags`)
			})

			Convey("Infra", func() {

				Convey("buildbucket", func() {
					Convey("agent", func() {
						Convey("input", func() {
							Convey("package", func() {
								req.Build.Infra.Buildbucket.Agent.Input = &pb.BuildInfra_Buildbucket_Agent_Input{
									Data: map[string]*pb.InputDataRef{
										"path_a": {
											DataType: &pb.InputDataRef_Cipd{
												Cipd: &pb.InputDataRef_CIPD{
													Specs: []*pb.InputDataRef_CIPD_PkgSpec{{Package: "", Version: "latest"}},
												},
											},
										},
									},
								}
								_, err := validateCreateBuildRequest(ctx, wellknownExps, req)
								So(err, ShouldErrLike, `build: infra: buildbucket: agent: input: [path_a]: [0]: cipd.package`)
							})
							Convey("version", func() {
								req.Build.Infra.Buildbucket.Agent.Input = &pb.BuildInfra_Buildbucket_Agent_Input{
									Data: map[string]*pb.InputDataRef{
										"path_a": {
											DataType: &pb.InputDataRef_Cipd{
												Cipd: &pb.InputDataRef_CIPD{
													Specs: []*pb.InputDataRef_CIPD_PkgSpec{{Package: "a/package", Version: ""}},
												},
											},
										},
									},
								}
								_, err := validateCreateBuildRequest(ctx, wellknownExps, req)
								So(err, ShouldErrLike, `build: infra: buildbucket: agent: input: [path_a]: [0]: cipd.version`)
							})
						})
						Convey("source", func() {
							Convey("package", func() {
								req.Build.Infra.Buildbucket.Agent.Source.GetCipd().Package = "cipd/package"
								_, err := validateCreateBuildRequest(ctx, wellknownExps, req)
								So(err, ShouldErrLike, `build: infra: buildbucket: agent: source: cipd.package`)
							})
							Convey("version", func() {
								req.Build.Infra.Buildbucket.Agent.Source.GetCipd().Version = "+100"
								_, err := validateCreateBuildRequest(ctx, wellknownExps, req)
								So(err, ShouldErrLike, `build: infra: buildbucket: agent: source: cipd.version`)
							})
						})
						Convey("purposes", func() {
							req.Build.Infra.Buildbucket.Agent.Input = &pb.BuildInfra_Buildbucket_Agent_Input{
								Data: map[string]*pb.InputDataRef{
									"path_a": {
										DataType: &pb.InputDataRef_Cipd{
											Cipd: &pb.InputDataRef_CIPD{
												Specs: []*pb.InputDataRef_CIPD_PkgSpec{{Package: "pkg_a", Version: "latest"}},
											},
										},
										OnPath: []string{"path_a/bin", "path_a"},
									},
								},
							}
							req.Build.Infra.Buildbucket.Agent.Purposes = map[string]pb.BuildInfra_Buildbucket_Agent_Purpose{
								"path_b": pb.BuildInfra_Buildbucket_Agent_PURPOSE_BBAGENT_UTILITY,
							}
							_, err := validateCreateBuildRequest(ctx, wellknownExps, req)
							So(err, ShouldErrLike, `build: infra: buildbucket: agent: purposes`)
						})
					})
				})

				Convey("swarming", func() {
					Convey("hostname", func() {
						req.Build.Infra.Swarming.Hostname = "https://host"
						_, err := validateCreateBuildRequest(ctx, wellknownExps, req)
						So(err, ShouldErrLike, `build: infra: swarming: hostname: must not contain "://"`)
					})
					Convey("priority", func() {
						req.Build.Infra.Swarming.Priority = 500
						_, err := validateCreateBuildRequest(ctx, wellknownExps, req)
						So(err, ShouldErrLike, `build: infra: swarming: priority must be in [0, 255]`)
					})
					Convey("task_dimensions", func() {
						Convey("key", func() {
							req.Build.Infra.Swarming.TaskDimensions = []*pb.RequestedDimension{
								{
									Key: "",
								},
							}
							_, err := validateCreateBuildRequest(ctx, wellknownExps, req)
							So(err, ShouldErrLike, `build: infra: swarming: task_dimensions: [0]: key must be specified`)
						})
						Convey("value", func() {
							req.Build.Infra.Swarming.TaskDimensions = []*pb.RequestedDimension{
								{
									Key:   "key",
									Value: "",
								},
							}
							_, err := validateCreateBuildRequest(ctx, wellknownExps, req)
							So(err, ShouldErrLike, `build: infra: swarming: task_dimensions: [0]: value must be specified`)
						})
						Convey("expiration", func() {
							req.Build.Infra.Swarming.TaskDimensions = []*pb.RequestedDimension{
								{
									Key:   "key",
									Value: "value",
									Expiration: &durationpb.Duration{
										Seconds: 200,
									},
								},
							}
							_, err := validateCreateBuildRequest(ctx, wellknownExps, req)
							So(err, ShouldErrLike, `build: infra: swarming: task_dimensions: [0]: expiration:`)
						})
					})
				})

				Convey("logdog", func() {
					Convey("hostname", func() {
						req.Build.Infra.Logdog.Hostname = "https://host"
						_, err := validateCreateBuildRequest(ctx, wellknownExps, req)
						So(err, ShouldErrLike, `build: infra: logdog: hostname: must not contain "://"`)
					})
				})

				Convey("resultdb", func() {
					Convey("hostname", func() {
						req.Build.Infra.Resultdb.Hostname = "https://host"
						_, err := validateCreateBuildRequest(ctx, wellknownExps, req)
						So(err, ShouldErrLike, `build: infra: resultdb: hostname: must not contain "://"`)
					})
				})

				Convey("backend", func() {
					Convey("should not specified", func() {
						req.Build.Infra.Backend = &pb.BuildInfra_Backend{}
						_, err := validateCreateBuildRequest(ctx, wellknownExps, req)
						So(err, ShouldErrLike, `build: infra: backend: should not be specified`)
					})
				})
			})

			Convey("output_only fields are cleared", func() {
				req.Build.Id = 87654321
				_, err := validateCreateBuildRequest(ctx, wellknownExps, req)
				So(err, ShouldBeNil)
				So(req.Build.Id, ShouldEqual, 0)
			})
		})

		Convey("CreateBuild specified output_only fields are cleared", func() {
			req.Build.Status = pb.Status_SCHEDULED
			req.Build.SummaryMarkdown = "random string"
			_, err := validateCreateBuildRequest(ctx, wellknownExps, req)
			So(err, ShouldBeNil)
			So(req.Build.Status, ShouldEqual, pb.Status_STATUS_UNSPECIFIED)
			So(req.Build.SummaryMarkdown, ShouldEqual, "")
		})

		Convey("CreateBuild ensures required fields", func() {
			Convey("top level required fields are ensured", func() {
				req.Build.Infra = nil
				_, err := validateCreateBuildRequest(ctx, wellknownExps, req)
				So(err, ShouldErrLike, ".build.infra: required")
			})

			Convey("sub fields are required if their upper level is non nil", func() {
				req.Build.Infra.Resultdb = nil
				_, err := validateCreateBuildRequest(ctx, wellknownExps, req)
				So(err, ShouldBeNil)

				req.Build.Infra.Resultdb = &pb.BuildInfra_ResultDB{}
				_, err = validateCreateBuildRequest(ctx, wellknownExps, req)
				So(err, ShouldErrLike, ".build.infra.resultdb.hostname: required")
			})
		})
	})
}