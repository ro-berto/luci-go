// Copyright 2018 The LUCI Authors.
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

package gerrit

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/timestamppb"

	gerritpb "go.chromium.org/luci/common/proto/gerrit"
	"go.chromium.org/luci/grpc/grpcutil"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"
)

func TestBuildURL(t *testing.T) {
	t.Parallel()

	Convey("buildURL works correctly", t, func() {
		cPB, err := NewRESTClient(nil, "x-review.googlesource.com", true)
		So(err, ShouldBeNil)
		c, ok := cPB.(*client)
		So(ok, ShouldBeTrue)

		So(c.buildURL("/changes/project~123", nil, nil), ShouldResemble,
			"https://x-review.googlesource.com/a/changes/project~123")
		So(c.buildURL("/changes/project~123", url.Values{"o": []string{"ONE", "TWO"}}, nil), ShouldResemble,
			"https://x-review.googlesource.com/a/changes/project~123?o=ONE&o=TWO")

		opt := UseGerritMirror(func(host string) string { return "mirror-" + host })
		So(c.buildURL("/changes/project~123", nil, []grpc.CallOption{opt}), ShouldResemble,
			"https://mirror-x-review.googlesource.com/a/changes/project~123")

		c.auth = false
		So(c.buildURL("/path", nil, nil), ShouldResemble,
			"https://x-review.googlesource.com/path")
	})
}

func TestListChanges(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	Convey("ListChanges", t, func() {
		Convey("Validates Limit number", func() {
			srv, c := newMockPbClient(func(w http.ResponseWriter, r *http.Request) {})
			defer srv.Close()

			_, err := c.ListChanges(ctx, &gerritpb.ListChangesRequest{
				Query: "label:Commit-Queue",
				Limit: -1,
			})
			So(err, ShouldErrLike, "must be nonnegative")

			_, err = c.ListChanges(ctx, &gerritpb.ListChangesRequest{
				Query: "label:Commit-Queue",
				Limit: 1001,
			})
			So(err, ShouldErrLike, "should be at most")
		})

		req := &gerritpb.ListChangesRequest{
			Query: "label:Code-Review",
			Limit: 1,
		}

		Convey("OK case with one change, _more_changes set in response", func() {
			expectedResponse := &gerritpb.ListChangesResponse{
				Changes: []*gerritpb.ChangeInfo{
					{
						Number: 1,
						Owner: &gerritpb.AccountInfo{
							AccountId: 1000096,
							Name:      "John Doe",
							Email:     "jdoe@example.com",
							Username:  "jdoe",
						},
						Project: "example/repo",
						Ref:     "refs/heads/master",
						Created: timestamppb.New(parseTime("2014-05-05T07:15:44.639000000Z")),
						Updated: timestamppb.New(parseTime("2014-05-05T07:15:44.639000000Z")),
					},
				},
				MoreChanges: true,
			}
			var actualRequest *http.Request
			srv, c := newMockPbClient(func(w http.ResponseWriter, r *http.Request) {
				actualRequest = r
				w.WriteHeader(200)
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprint(w, `)]}'[
					{
						"_number": 1,
						"owner": {
							"_account_id":      1000096,
							"name":             "John Doe",
							"email":            "jdoe@example.com",
							"username":         "jdoe"
						},
						"project": "example/repo",
						"branch":  "master",
						"created":   "2014-05-05 07:15:44.639000000",
						"updated":   "2014-05-05 07:15:44.639000000",
						"_more_changes": true
					}
				]`)
			})
			defer srv.Close()

			Convey("Response and request are as expected", func() {
				res, err := c.ListChanges(ctx, req)
				So(err, ShouldBeNil)
				So(res, ShouldResemble, expectedResponse)
				So(actualRequest.URL.Query()["q"], ShouldResemble, []string{"label:Code-Review"})
				So(actualRequest.URL.Query()["S"], ShouldResemble, []string{"0"})
				So(actualRequest.URL.Query()["n"], ShouldResemble, []string{"1"})
			})

			Convey("Options are included in the request", func() {
				req.Options = append(req.Options, gerritpb.QueryOption_DETAILED_ACCOUNTS, gerritpb.QueryOption_ALL_COMMITS)
				_, err := c.ListChanges(ctx, req)
				So(err, ShouldBeNil)
				So(
					actualRequest.URL.Query()["o"],
					ShouldResemble,
					[]string{"DETAILED_ACCOUNTS", "ALL_COMMITS"},
				)
			})
		})
	})
}

func TestGetChange(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	Convey("GetChange", t, func() {
		Convey("Validate args", func() {
			srv, c := newMockPbClient(func(w http.ResponseWriter, r *http.Request) {})
			defer srv.Close()

			_, err := c.GetChange(ctx, &gerritpb.GetChangeRequest{})
			So(err, ShouldErrLike, "number must be positive")
		})

		req := &gerritpb.GetChangeRequest{Number: 1}

		Convey("OK", func() {
			expectedChange := &gerritpb.ChangeInfo{
				Number: 1,
				Owner: &gerritpb.AccountInfo{
					AccountId:       1000096,
					Name:            "John Doe",
					Email:           "jdoe@example.com",
					SecondaryEmails: []string{"johndoe@chromium.org"},
					Username:        "jdoe",
				},
				Project:         "example/repo",
				Ref:             "refs/heads/master",
				Status:          gerritpb.ChangeStatus_NEW,
				CurrentRevision: "deadbeef",
				Submittable:     true,
				IsPrivate:       true,
				Revisions: map[string]*gerritpb.RevisionInfo{
					"deadbeef": {
						Number:      1,
						Kind:        gerritpb.RevisionInfo_REWORK,
						Ref:         "refs/changes/123",
						Created:     timestamppb.New(parseTime("2016-03-29T17:47:23.751000000Z")),
						Description: "first upload",
						Files: map[string]*gerritpb.FileInfo{
							"go/to/file.go": {
								LinesInserted: 32,
								LinesDeleted:  44,
								SizeDelta:     -567,
								Size:          11984,
							},
						},
						Commit: &gerritpb.CommitInfo{
							Id:      "", // Gerrit doesn't set it, as it duplicates key in revisions map.
							Message: "Title.\n\nBody is here.\n\nChange-Id: I100deadbeef",
							Parents: []*gerritpb.CommitInfo_Parent{
								{Id: "deadbeef00"},
							},
						},
					},
				},
				Labels: map[string]*gerritpb.LabelInfo{
					"Code-Review": {
						Approved: &gerritpb.AccountInfo{
							Name:  "Rubber Stamper",
							Email: "rubberstamper@example.com",
						},
					},
					"Commit-Queue": {
						Optional:     true,
						DefaultValue: 0,
						Values:       map[int32]string{0: "Not ready", 1: "Dry run", 2: "Commit"},
						All: []*gerritpb.ApprovalInfo{
							{
								User: &gerritpb.AccountInfo{
									AccountId: 1010101,
									Name:      "Dry Runner",
									Email:     "dry-runner@example.com",
								},
								Value:                1,
								PermittedVotingRange: &gerritpb.VotingRangeInfo{Min: 0, Max: 2},
								Date:                 timestamppb.New(parseTime("2020-12-13T18:32:35.000000000Z")),
							},
						},
					},
				},
				Created: timestamppb.New(parseTime("2014-05-05T07:15:44.639000000Z")),
				Updated: timestamppb.New(parseTime("2014-05-05T07:15:44.639000000Z")),
				Messages: []*gerritpb.ChangeMessageInfo{
					{
						Id: "YH-egE",
						Author: &gerritpb.AccountInfo{
							AccountId: 1000096,
							Name:      "John Doe",
							Email:     "john.doe@example.com",
							Username:  "jdoe",
						},
						Date:    timestamppb.New(parseTime("2013-03-23T21:34:02.419000000Z")),
						Message: "Patch Set 1:\n\nThis is the first message.",
					},
					{
						Id: "WEEdhU",
						Author: &gerritpb.AccountInfo{
							AccountId: 1000097,
							Name:      "Jane Roe",
							Email:     "jane.roe@example.com",
							Username:  "jroe",
						},
						Date:    timestamppb.New(parseTime("2013-03-23T21:36:52.332000000Z")),
						Message: "Patch Set 1:\n\nThis is the second message.\n\nWith a line break.",
					},
				},
				Requirements: []*gerritpb.Requirement{
					{
						Status:       gerritpb.Requirement_REQUIREMENT_STATUS_OK,
						FallbackText: "nothing more required",
						Type:         "alpha-numer1c-type",
					},
				},
			}
			var actualRequest *http.Request
			srv, c := newMockPbClient(func(w http.ResponseWriter, r *http.Request) {
				actualRequest = r
				w.WriteHeader(200)
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprint(w, `)]}'{
					"_number": 1,
					"status": "NEW",
					"owner": {
						"_account_id":      1000096,
						"name":             "John Doe",
						"email":            "jdoe@example.com",
						"secondary_emails": ["johndoe@chromium.org"],
						"username":         "jdoe"
					},
					"created":   "2014-05-05 07:15:44.639000000",
					"updated":   "2014-05-05 07:15:44.639000000",
					"project": "example/repo",
					"branch":  "master",
					"current_revision": "deadbeef",
					"submittable": true,
					"is_private": true,
					"revisions": {
						"deadbeef": {
							"_number": 1,
							"kind": "REWORK",
							"ref": "refs/changes/123",
							"created": "2016-03-29 17:47:23.751000000",
							"description": "first upload",
							"files": {
								"go/to/file.go": {
									"lines_inserted": 32,
									"lines_deleted": 44,
									"size_delta": -567,
									"size": 11984
								}
							},
							"commit": {
								"parents": [{"commit": "deadbeef00"}],
								"author": {
									"name": "John Doe",
									"email": "jdoe@example.com",
									"date": "2014-05-05 07:15:44.639000000",
									"tz": 60
								},
								"committer": {
									"name": "John Doe",
									"email": "jdoe@example.com",
									"date": "2014-05-05 07:15:44.639000000",
									"tz": 60
								},
								"subject": "Title.",
								"message": "Title.\n\nBody is here.\n\nChange-Id: I100deadbeef"
							}
						}
					},
					"labels": {
						"Code-Review": {
							"approved": {
								"name": "Rubber Stamper",
								"email": "rubberstamper@example.com"
							}
						},
						"Commit-Queue": {
							"all": [
								{
									"value": 1,
									"date": "2020-12-13 18:32:35.000000000",
									"permitted_voting_range": {
										"min": 0,
										"max": 2
									},
									"_account_id": 1010101,
									"name": "Dry Runner",
									"email": "dry-runner@example.com",
									"avatars": [
										{
											"url": "https://example.com/photo.jpg",
											"height": 32
										}
									]
								}
							],
							"values": {
								" 0": "Not ready",
								"+1": "Dry run",
								"+2": "Commit"
							},
							"default_value": 0,
							"optional": true
						}
					},
					"messages": [
						{
							"id": "YH-egE",
							"author": {
								"_account_id": 1000096,
								"name": "John Doe",
								"email": "john.doe@example.com",
								"username": "jdoe"
							},
							"date": "2013-03-23 21:34:02.419000000",
							"message": "Patch Set 1:\n\nThis is the first message.",
							"_revision_number": 1
						},
						{
							"id": "WEEdhU",
							"author": {
								"_account_id": 1000097,
								"name": "Jane Roe",
								"email": "jane.roe@example.com",
								"username": "jroe"
							},
							"date": "2013-03-23 21:36:52.332000000",
							"message": "Patch Set 1:\n\nThis is the second message.\n\nWith a line break.",
							"_revision_number": 1
						}
					],
					"requirements": [
						{
							"status": "OK",
							"fallback_text": "nothing more required",
							"type": "alpha-numer1c-type"
						}
					]
				}`)
			})
			defer srv.Close()

			Convey("Basic", func() {
				res, err := c.GetChange(ctx, req)
				So(err, ShouldBeNil)
				So(res, ShouldResemble, expectedChange)
			})

			Convey("With project", func() {
				req.Project = "infra/luci"
				res, err := c.GetChange(ctx, req)
				So(err, ShouldBeNil)
				So(res, ShouldResembleProto, expectedChange)
				So(actualRequest.URL.EscapedPath(), ShouldEqual, "/changes/infra%2Fluci~1")
			})

			Convey("Options", func() {
				req.Options = append(req.Options, gerritpb.QueryOption_DETAILED_ACCOUNTS, gerritpb.QueryOption_ALL_COMMITS)
				_, err := c.GetChange(ctx, req)
				So(err, ShouldBeNil)
				So(
					actualRequest.URL.Query()["o"],
					ShouldResemble,
					[]string{"DETAILED_ACCOUNTS", "ALL_COMMITS"},
				)
			})
		})
	})
}

func TestRestCreateChange(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	Convey("CreateChange basic", t, func() {
		var actualBody []byte
		srv, c := newMockPbClient(func(w http.ResponseWriter, r *http.Request) {
			// ignore errors here, but verify body later.
			actualBody, _ = ioutil.ReadAll(r.Body)
			w.WriteHeader(201)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `)]}'`)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"_number":   1,
				"project":   "example/repo",
				"branch":    "master",
				"change_id": "c1",
				"status":    "NEW",
				"created":   "2014-05-05 07:15:44.639000000",
				"updated":   "2014-05-05 07:15:44.639000000",
			})
		})
		defer srv.Close()

		req := gerritpb.CreateChangeRequest{
			Project:    "example/repo",
			Ref:        "refs/heads/master",
			Subject:    "example subject",
			BaseCommit: "someOpaqueHash",
		}
		res, err := c.CreateChange(ctx, &req)
		So(err, ShouldBeNil)
		So(res, ShouldResemble, &gerritpb.ChangeInfo{
			Number:      1,
			Project:     "example/repo",
			Ref:         "refs/heads/master",
			Status:      gerritpb.ChangeStatus_NEW,
			Submittable: false,
			Created:     timestamppb.New(parseTime("2014-05-05T07:15:44.639000000Z")),
			Updated:     timestamppb.New(parseTime("2014-05-05T07:15:44.639000000Z")),
		})

		var ci changeInput
		err = json.Unmarshal(actualBody, &ci)
		So(err, ShouldBeNil)
		So(ci, ShouldResemble, changeInput{
			Project:    "example/repo",
			Branch:     "refs/heads/master",
			Subject:    "example subject",
			BaseCommit: "someOpaqueHash",
		})
	})
}

func TestSubmitRevision(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	Convey("SubmitRevision", t, func() {
		var actualURL *url.URL
		srv, c := newMockPbClient(func(w http.ResponseWriter, r *http.Request) {
			actualURL = r.URL
			w.WriteHeader(200)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `)]}'`)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status": "MERGED",
			})
		})
		defer srv.Close()

		req := &gerritpb.SubmitRevisionRequest{
			Number:     42,
			RevisionId: "someRevision",
			Project:    "someProject",
		}
		res, err := c.SubmitRevision(ctx, req)
		So(err, ShouldBeNil)
		So(res, ShouldResembleProto, &gerritpb.SubmitInfo{
			Status: gerritpb.ChangeStatus_MERGED,
		})
		So(actualURL.Path, ShouldEqual, "/changes/someProject~42/revisions/someRevision/submit")
	})
}

func TestRestChangeEditFileContent(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	Convey("ChangeEditFileContent basic", t, func() {
		// large enough?
		var actualBody []byte
		var actualURL *url.URL
		srv, c := newMockPbClient(func(w http.ResponseWriter, r *http.Request) {
			actualURL = r.URL
			// ignore errors here, but verify body later.
			actualBody, _ = ioutil.ReadAll(r.Body)
			// API returns 204 on success.
			w.WriteHeader(204)
		})
		defer srv.Close()

		_, err := c.ChangeEditFileContent(ctx, &gerritpb.ChangeEditFileContentRequest{
			Number:   42,
			Project:  "someproject",
			FilePath: "some/path",
			Content:  []byte("changed file"),
		})
		So(err, ShouldBeNil)
		So(actualURL.Path, ShouldEqual, "/changes/someproject~42/edit/some/path")
		So(actualBody, ShouldResemble, []byte("changed file"))
	})
}

// TODO (yulanlin): Assert body verbatim without decoding
func TestAddReviewer(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	Convey("Add reviewer to cc basic", t, func() {
		var actualURL *url.URL
		var actualBody []byte
		srv, c := newMockPbClient(func(w http.ResponseWriter, r *http.Request) {
			actualURL = r.URL
			// ignore the error because body contents will be checked
			actualBody, _ = ioutil.ReadAll(r.Body)
			w.WriteHeader(200)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `)]}'
			{
				"input": "ccer@test.com",
				"ccs": [
					{
						"_account_id": 10001,
						"name": "Reviewer Review",
						"approvals": {
							"Code-Review": " 0"
						}
					}
				]
			}`)
		})
		defer srv.Close()

		req := &gerritpb.AddReviewerRequest{
			Number:    42,
			Project:   "someproject",
			Reviewer:  "ccer@test.com",
			State:     gerritpb.AddReviewerRequest_ADD_REVIEWER_STATE_CC,
			Confirmed: true,
			Notify:    gerritpb.Notify_NOTIFY_OWNER,
		}
		res, err := c.AddReviewer(ctx, req)
		So(err, ShouldBeNil)

		// assert the request was as expected
		So(actualURL.Path, ShouldEqual, "/changes/someproject~42/reviewers")
		var body addReviewerRequest
		err = json.Unmarshal(actualBody, &body)
		if err != nil {
			t.Logf("failed to decode req body: %v\n", err)
		}
		So(body, ShouldResemble, addReviewerRequest{
			Reviewer:  "ccer@test.com",
			State:     "CC",
			Confirmed: true,
			Notify:    "OWNER",
		})

		// assert the result was as expected
		So(res, ShouldResemble, &gerritpb.AddReviewerResult{
			Input:     "ccer@test.com",
			Reviewers: []*gerritpb.ReviewerInfo{},
			Ccs: []*gerritpb.ReviewerInfo{
				{
					Account: &gerritpb.AccountInfo{
						Name:      "Reviewer Review",
						AccountId: 10001,
					},
					Approvals: map[string]int32{
						"Code-Review": 0,
					},
				},
			},
		})
	})
}

func TestDeleteReviewer(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	Convey("Delete reviewer", t, func() {
		var actualURL *url.URL
		srv, c := newMockPbClient(func(w http.ResponseWriter, r *http.Request) {
			actualURL = r.URL
			// API returns 204 on success.
			w.WriteHeader(204)
		})
		defer srv.Close()

		_, err := c.DeleteReviewer(ctx, &gerritpb.DeleteReviewerRequest{
			Number:    42,
			Project:   "someproject",
			AccountId: "jdoe@example.com",
		})
		So(err, ShouldBeNil)
		So(actualURL.Path, ShouldEqual, "/changes/someproject~42/reviewers/jdoe@example.com/delete")
	})
}

func TestSetReview(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	Convey("Set Review", t, func() {
		var actualURL *url.URL
		var actualRawBody []byte
		srv, c := newMockPbClient(func(w http.ResponseWriter, r *http.Request) {
			actualURL = r.URL
			actualRawBody, _ = ioutil.ReadAll(r.Body)
			w.WriteHeader(200)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `)]}'
			{
				"labels": {
					"Code-Review": -1
				}
			}`)
		})
		defer srv.Close()

		res, err := c.SetReview(ctx, &gerritpb.SetReviewRequest{
			Number:     42,
			Project:    "someproject",
			RevisionId: "somerevision",
			Message:    "This is a message",
			Labels: map[string]int32{
				"Code-Review": -1,
			},
			Tag:    "autogenerated:cq",
			Notify: gerritpb.Notify_NOTIFY_OWNER,
			NotifyDetails: &gerritpb.NotifyDetails{
				Recipients: []*gerritpb.NotifyDetails_Recipient{
					{
						RecipientType: gerritpb.NotifyDetails_RECIPIENT_TYPE_TO,
						Info: &gerritpb.NotifyDetails_Info{
							Accounts: []int64{4, 5, 3},
						},
					},
					{
						RecipientType: gerritpb.NotifyDetails_RECIPIENT_TYPE_TO,
						Info: &gerritpb.NotifyDetails_Info{
							Accounts: []int64{2, 3, 1},
							// 3 is overlapping with the first recipient,
						},
					},
					{
						RecipientType: gerritpb.NotifyDetails_RECIPIENT_TYPE_BCC,
						Info: &gerritpb.NotifyDetails_Info{
							Accounts: []int64{6, 1},
						},
					},
				},
			},
			OnBehalfOf: 10001,
			Ready:      true,
			AddToAttentionSet: []*gerritpb.AttentionSetInput{
				{User: "10002", Reason: "passed presubmit"},
			},
			RemoveFromAttentionSet: []*gerritpb.AttentionSetInput{
				{User: "10001", Reason: "passed presubmit"},
			},
			IgnoreAutomaticAttentionSetRules: true,
		})
		So(err, ShouldBeNil)
		So(actualURL.Path, ShouldEqual, "/changes/someproject~42/revisions/somerevision/review")

		var actualBody, expectedBody map[string]interface{}
		So(json.Unmarshal(actualRawBody, &actualBody), ShouldBeNil)
		So(json.Unmarshal([]byte(`{
			"message": "This is a message",
			"labels": {
				"Code-Review": -1
			},
			"tag": "autogenerated:cq",
			"notify": "OWNER",
			"notify_details": {
				"TO":  {"accounts": [1, 2, 3, 4, 5]},
				"BCC": {"accounts": [1, 6]}
			},
			"on_behalf_of": 10001,
			"ready": true,
			"add_to_attention_set": [
				{"user": "10002", "reason": "passed presubmit"}
			],
			"remove_from_attention_set": [
				{"user": "10001", "reason": "passed presubmit"}
			],
			"ignore_automatic_attention_set_rules": true
		}`), &expectedBody), ShouldBeNil)
		So(actualBody, ShouldResemble, expectedBody)

		So(res, ShouldResembleProto, &gerritpb.ReviewResult{
			Labels: map[string]int32{
				"Code-Review": -1,
			},
		})
	})
}

func TestAddToAttentionSet(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	Convey("Add to attention set", t, func() {
		var actualURL *url.URL
		var actualBody []byte
		srv, c := newMockPbClient(func(w http.ResponseWriter, r *http.Request) {
			actualURL = r.URL
			// ignore the error because body contents will be checked
			actualBody, _ = ioutil.ReadAll(r.Body)
			w.WriteHeader(200)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `)]}'
				{
					"_account_id": 10001,
					"name": "FYI reviewer",
					"email": "fyi@test.com",
					"username": "fyi"
				}`)
		})
		defer srv.Close()

		req := &gerritpb.AttentionSetRequest{
			Project: "someproject",
			Number:  42,
			Input: &gerritpb.AttentionSetInput{
				User:   "fyi@test.com",
				Reason: "For awareness",
				Notify: gerritpb.Notify_NOTIFY_ALL,
			},
		}
		res, err := c.AddToAttentionSet(ctx, req)
		So(err, ShouldBeNil)

		// assert the request was as expected
		So(actualURL.Path, ShouldEqual, "/changes/someproject~42/attention")
		expectedBody, err := json.Marshal(attentionSetInput{
			User:   "fyi@test.com",
			Reason: "For awareness",
			Notify: "ALL",
		})
		if err != nil {
			t.Logf("failed to encode expected body: %v\n", err)
		}
		So(actualBody, ShouldResemble, expectedBody)

		// assert the result was as expected
		So(res, ShouldResemble, &gerritpb.AccountInfo{
			AccountId: 10001,
			Name:      "FYI reviewer",
			Email:     "fyi@test.com",
			Username:  "fyi",
		})
	})
}

func TestGetMergeable(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	Convey("GetMergeable basic", t, func() {
		var actualURL *url.URL
		srv, c := newMockPbClient(func(w http.ResponseWriter, r *http.Request) {
			actualURL = r.URL
			w.WriteHeader(200)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `)]}'
        {
          "submit_type": "CHERRY_PICK",
          "strategy": "simple-two-way-in-core",
          "mergeable": true,
          "commit_merged": false,
          "content_merged": false,
          "conflicts": [
            "conflict1",
            "conflict2"
          ],
          "mergeable_into": [
            "my_branch_1"
          ]
        }`)
		})
		defer srv.Close()

		mi, err := c.GetMergeable(ctx, &gerritpb.GetMergeableRequest{
			Number:     42,
			Project:    "someproject",
			RevisionId: "somerevision",
		})
		So(err, ShouldBeNil)
		So(actualURL.Path, ShouldEqual, "/changes/someproject~42/revisions/somerevision/mergeable")
		So(mi, ShouldResemble, &gerritpb.MergeableInfo{
			SubmitType:    gerritpb.MergeableInfo_CHERRY_PICK,
			Strategy:      gerritpb.MergeableStrategy_SIMPLE_TWO_WAY_IN_CORE,
			Mergeable:     true,
			CommitMerged:  false,
			ContentMerged: false,
			Conflicts:     []string{"conflict1", "conflict2"},
			MergeableInto: []string{"my_branch_1"},
		})
	})
}

func TestListFiles(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	Convey("ListFiles basic", t, func() {
		var actualURL *url.URL
		srv, c := newMockPbClient(func(w http.ResponseWriter, r *http.Request) {
			actualURL = r.URL
			w.WriteHeader(200)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `)]}'
				{
					"gerrit-server/src/main/java/com/google/gerrit/server/project/RefControl.java": {
						"lines_inserted": 123456
					},
					"file2": {
						"size": 7
					}
				}`)
		})
		defer srv.Close()

		mi, err := c.ListFiles(ctx, &gerritpb.ListFilesRequest{
			Number:     42,
			Project:    "someproject",
			RevisionId: "somerevision",
			Parent:     999,
		})
		So(err, ShouldBeNil)
		So(actualURL.Path, ShouldEqual, "/changes/someproject~42/revisions/somerevision/files/")
		So(actualURL.Query().Get("parent"), ShouldEqual, "999")
		So(mi, ShouldResemble, &gerritpb.ListFilesResponse{
			Files: map[string]*gerritpb.FileInfo{
				"gerrit-server/src/main/java/com/google/gerrit/server/project/RefControl.java": {
					LinesInserted: 123456,
				},
				"file2": {
					Size: 7,
				},
			},
		})
	})
}

func TestGetRelatedChanges(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	Convey("GetRelatedChanges works", t, func() {
		var actualURL *url.URL
		srv, c := newMockPbClient(func(w http.ResponseWriter, r *http.Request) {
			actualURL = r.URL
			w.WriteHeader(200)
			w.Header().Set("Content-Type", "application/json")
			// Taken from
			// https://chromium-review.googlesource.com/changes/playground%2Fgerrit-cq~1563638/revisions/2/related
			fmt.Fprint(w, `)]}'
				{
				  "changes": [
				    {
				      "project": "playground/gerrit-cq",
				      "change_id": "If00fa4f207440d7f12fbfff8c05afa9077ab0c21",
				      "commit": {
				        "commit": "4d048b016cb4df4d5d2805f0d3d1042cb1d80671",
				        "parents": [
				          {
				            "commit": "cd7db096c014399c369ddddd319708c3f46752f5"
				          }
				        ],
				        "author": {
				          "name": "Andrii Shyshkalov",
				          "email": "tandrii@chromium.org",
				          "date": "2019-04-11 06:41:01.000000000",
				          "tz": -420
				        },
				        "subject": "p3 change"
				      },
				      "_change_number": 1563639,
				      "_revision_number": 1,
				      "_current_revision_number": 1,
				      "status": "NEW"
				    },
				    {
				      "project": "playground/gerrit-cq",
				      "change_id": "I80bf05eb9124dc126490820ec192c77a24938622",
				      "commit": {
				        "commit": "bce1f3beea01b8b282001b01bd9ea442730d578e",
				        "parents": [
				          {
				            "commit": "fdd1f6d3875e68c99303ebfb25dd5d097e91c83f"
				          }
				        ],
				        "author": {
				          "name": "Andrii Shyshkalov",
				          "email": "tandrii@chromium.org",
				          "date": "2019-04-11 06:40:28.000000000",
				          "tz": -420
				        },
				        "subject": "p2 change"
				      },
				      "_change_number": 1563638,
				      "_revision_number": 2,
				      "_current_revision_number": 2,
				      "status": "NEW"
				    },
				    {
				      "project": "playground/gerrit-cq",
				      "change_id": "Icf12c110abc0cbc0c7d01a40dc047683634a62d7",
				      "commit": {
				        "commit": "fdd1f6d3875e68c99303ebfb25dd5d097e91c83f",
				        "parents": [
				          {
				            "commit": "f8e5384ee591cd5105113098d24c60e750b6c4f6"
				          }
				        ],
				        "author": {
				          "name": "Andrii Shyshkalov",
				          "email": "tandrii@chromium.org",
				          "date": "2019-04-11 06:40:18.000000000",
				          "tz": -420
				        },
				        "subject": "p1 change"
				      },
				      "_change_number": 1563637,
				      "_revision_number": 1,
				      "_current_revision_number": 1,
				      "status": "NEW"
				    }
				  ]
				}
			`)
		})
		defer srv.Close()

		rcs, err := c.GetRelatedChanges(ctx, &gerritpb.GetRelatedChangesRequest{
			Number:     1563638,
			Project:    "playground/gerrit-cq",
			RevisionId: "2",
		})
		So(err, ShouldBeNil)
		So(actualURL.EscapedPath(), ShouldEqual, "/changes/playground%2Fgerrit-cq~1563638/revisions/2/related")
		So(rcs, ShouldResembleProto, &gerritpb.GetRelatedChangesResponse{
			Changes: []*gerritpb.GetRelatedChangesResponse_ChangeAndCommit{
				{
					Project: "playground/gerrit-cq",
					Commit: &gerritpb.CommitInfo{
						Id:      "4d048b016cb4df4d5d2805f0d3d1042cb1d80671",
						Parents: []*gerritpb.CommitInfo_Parent{{Id: "cd7db096c014399c369ddddd319708c3f46752f5"}},
					},
					Number:          1563639,
					Patchset:        1,
					CurrentPatchset: 1,
				},
				{
					Project: "playground/gerrit-cq",
					Commit: &gerritpb.CommitInfo{
						Id:      "bce1f3beea01b8b282001b01bd9ea442730d578e",
						Parents: []*gerritpb.CommitInfo_Parent{{Id: "fdd1f6d3875e68c99303ebfb25dd5d097e91c83f"}},
					},
					Number:          1563638,
					Patchset:        2,
					CurrentPatchset: 2,
				},
				{
					Project: "playground/gerrit-cq",
					Commit: &gerritpb.CommitInfo{
						Id:      "fdd1f6d3875e68c99303ebfb25dd5d097e91c83f",
						Parents: []*gerritpb.CommitInfo_Parent{{Id: "f8e5384ee591cd5105113098d24c60e750b6c4f6"}},
					},
					Number:          1563637,
					Patchset:        1,
					CurrentPatchset: 1,
				},
			},
		})
	})
}

func TestGetFileOwners(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	Convey("Get Owners: ", t, func() {
		Convey("Details", func() {
			var actualURL *url.URL
			srv, c := newMockPbClient(func(w http.ResponseWriter, r *http.Request) {
				actualURL = r.URL
				w.WriteHeader(200)
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprint(w, `)]}'
				{"code_owners":[{
					"account":{
						"_account_id":1000096,
						"name":"User Name",
						"email":"user@test.com",
						"avatars":[{"url":"https://test.com/photo.jpg","height":32}]
					}}]}`)
			})
			defer srv.Close()

			resp, err := c.ListFileOwners(ctx, &gerritpb.ListFileOwnersRequest{
				Project: "projectName",
				Ref:     "main",
				Path:    "path/to/file",
				Options: &gerritpb.AccountOptions{
					Details: true,
				},
			})
			So(err, ShouldBeNil)
			So(actualURL.Path, ShouldEqual, "/projects/projectName/branches/main/code_owners/path/to/file")
			So(actualURL.Query().Get("o"), ShouldEqual, "DETAILS")
			So(resp, ShouldResemble, &gerritpb.ListOwnersResponse{
				Owners: []*gerritpb.OwnerInfo{
					{
						Account: &gerritpb.AccountInfo{
							AccountId: 1000096,
							Name:      "User Name",
							Email:     "user@test.com",
						},
					},
				},
			})
		})
		Convey("All Emails", func() {
			var actualURL *url.URL
			srv, c := newMockPbClient(func(w http.ResponseWriter, r *http.Request) {
				actualURL = r.URL
				w.WriteHeader(200)
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprint(w, `)]}'
				{"code_owners": [{
					"account": {
						"_account_id": 1000096,
						"email": "test@test.com",
						"secondary_emails": ["alt@test.com"]
					}}]}`)
			})
			defer srv.Close()

			resp, err := c.ListFileOwners(ctx, &gerritpb.ListFileOwnersRequest{
				Project: "projectName",
				Ref:     "main",
				Path:    "path/to/file",
				Options: &gerritpb.AccountOptions{
					AllEmails: true,
				},
			})
			So(err, ShouldBeNil)
			So(actualURL.Path, ShouldEqual, "/projects/projectName/branches/main/code_owners/path/to/file")
			So(actualURL.Query().Get("o"), ShouldEqual, "ALL_EMAILS")
			So(resp, ShouldResemble, &gerritpb.ListOwnersResponse{
				Owners: []*gerritpb.OwnerInfo{
					{
						Account: &gerritpb.AccountInfo{
							AccountId:       1000096,
							Email:           "test@test.com",
							SecondaryEmails: []string{"alt@test.com"},
						},
					},
				},
			})
		})
	})
}

func TestListProjects(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	Convey("List Projects", t, func() {
		Convey("...works for a single ref", func() {
			var actualURL *url.URL
			srv, c := newMockPbClient(func(w http.ResponseWriter, r *http.Request) {
				actualURL = r.URL
				w.WriteHeader(200)
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprint(w, `)]}'
				{
					"android_apks": {
					  "id": "android_apks",
					  "state": "ACTIVE",
					  "branches": {
						"main": "82264ea131fcc2a386b83e38b962b370315c7c93"
					  },
					  "web_links": [
						{
						  "name": "gitiles",
						  "url": "https://chromium.googlesource.com/android_apks/",
						  "target": "_blank"
						}
					  ]
					}
				  }`)
			})
			defer srv.Close()

			projects, err := c.ListProjects(ctx, &gerritpb.ListProjectsRequest{
				Refs: []string{"refs/heads/main"},
			})
			So(err, ShouldBeNil)
			So(actualURL.Path, ShouldEqual, "/projects/")
			So(actualURL.Query().Get("b"), ShouldEqual, "refs/heads/main")
			So(projects, ShouldResemble, &gerritpb.ListProjectsResponse{
				Projects: map[string]*gerritpb.ProjectInfo{
					"android_apks": {
						Name:  "android_apks",
						State: gerritpb.ProjectInfo_PROJECT_STATE_ACTIVE,
						Refs: map[string]string{
							"refs/heads/main": "82264ea131fcc2a386b83e38b962b370315c7c93",
						},
						WebLinks: []*gerritpb.WebLinkInfo{
							{
								Name: "gitiles",
								Url:  "https://chromium.googlesource.com/android_apks/",
							},
						},
					},
				},
			})
		})

		Convey("...works for multiple refs", func() {
			var actualURL *url.URL
			srv, c := newMockPbClient(func(w http.ResponseWriter, r *http.Request) {
				actualURL = r.URL
				w.WriteHeader(200)
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprint(w, `)]}'
				{
					"android_apks": {
					  "id": "android_apks",
					  "state": "ACTIVE",
					  "branches": {
						"main": "82264ea131fcc2a386b83e38b962b370315c7c93",
						"master": "82264ea131fcc2a386b83e38b962b370315c7c93"
					  },
					  "web_links": [
						{
						  "name": "gitiles",
						  "url": "https://chromium.googlesource.com/android_apks/",
						  "target": "_blank"
						}
					  ]
					}
				  }`)
			})
			defer srv.Close()

			projects, err := c.ListProjects(ctx, &gerritpb.ListProjectsRequest{
				Refs: []string{"refs/heads/main", "refs/heads/master"},
			})
			So(err, ShouldBeNil)
			So(actualURL.Path, ShouldEqual, "/projects/")
			So(actualURL.Query()["b"], ShouldResemble, []string{"refs/heads/main", "refs/heads/master"})
			So(projects, ShouldResemble, &gerritpb.ListProjectsResponse{
				Projects: map[string]*gerritpb.ProjectInfo{
					"android_apks": {
						Name:  "android_apks",
						State: gerritpb.ProjectInfo_PROJECT_STATE_ACTIVE,
						Refs: map[string]string{
							"refs/heads/main":   "82264ea131fcc2a386b83e38b962b370315c7c93",
							"refs/heads/master": "82264ea131fcc2a386b83e38b962b370315c7c93",
						},
						WebLinks: []*gerritpb.WebLinkInfo{
							{
								Name: "gitiles",
								Url:  "https://chromium.googlesource.com/android_apks/",
							},
						},
					},
				},
			})
		})
	})
}

func TestGetBranchInfo(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	Convey("Get Branch Info", t, func() {
		var actualURL *url.URL
		srv, c := newMockPbClient(func(w http.ResponseWriter, r *http.Request) {
			actualURL = r.URL
			w.WriteHeader(200)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `)]}'
			{
				"web_links": [
				  {
					"name": "gitiles",
					"url": "https://chromium.googlesource.com/infra/experimental/+/refs/heads/main",
					"target": "_blank"
				  }
				],
				"ref": "refs/heads/main",
				"revision": "10e5c33f63a843440cbe6c9c6cbc1bf513c598eb",
				"can_delete": true
			  }`)
		})
		defer srv.Close()

		bi, err := c.GetRefInfo(ctx, &gerritpb.RefInfoRequest{
			Project: "infra/experimental",
			Ref:     "refs/heads/main",
		})
		So(err, ShouldBeNil)

		So(actualURL.Path, ShouldEqual, "/projects/infra/experimental/branches/refs/heads/main")
		So(bi, ShouldResemble, &gerritpb.RefInfo{
			Ref:      "refs/heads/main",
			Revision: "10e5c33f63a843440cbe6c9c6cbc1bf513c598eb",
		})
	})
}

func TestGetPureRevert(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	Convey("Get Pure Revert", t, func() {
		var actualURL *url.URL
		srv, c := newMockPbClient(func(w http.ResponseWriter, r *http.Request) {
			actualURL = r.URL
			w.WriteHeader(200)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `)]}'
			{
				"is_pure_revert" : false
			}`)
		})
		defer srv.Close()

		req := &gerritpb.GetPureRevertRequest{
			Number:  42,
			Project: "someproject",
		}
		res, err := c.GetPureRevert(ctx, req)
		So(err, ShouldBeNil)
		So(actualURL.Path, ShouldEqual, "/changes/someproject~42/pure_revert")
		So(res, ShouldResemble, &gerritpb.PureRevertInfo{
			IsPureRevert: false,
		})
	})
}

func TestGerritError(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	Convey("Gerrit returns", t, func() {
		// All APIs share the same error handling code path, so use SubmitChange as
		// an example.
		req := &gerritpb.SubmitChangeRequest{Number: 1}
		Convey("HTTP 403", func() {
			srv, c := newMockPbClient(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(403)
			})
			defer srv.Close()
			_, err := c.SubmitChange(ctx, req)
			So(grpcutil.Code(err), ShouldEqual, codes.PermissionDenied)
		})
		Convey("HTTP 404 ", func() {
			srv, c := newMockPbClient(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(404)
			})
			defer srv.Close()
			_, err := c.SubmitChange(ctx, req)
			So(grpcutil.Code(err), ShouldEqual, codes.NotFound)
		})
		Convey("HTTP 409 ", func() {
			srv, c := newMockPbClient(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(409)
				w.Header().Set("Content-Type", "text/plain")
				w.Write([]byte("block by Verified"))
			})
			defer srv.Close()
			_, err := c.SubmitChange(ctx, req)
			So(grpcutil.Code(err), ShouldEqual, codes.FailedPrecondition)
			So(err, ShouldErrLike, "block by Verified")
		})
		Convey("HTTP 429 ", func() {
			srv, c := newMockPbClient(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(429)
			})
			defer srv.Close()
			_, err := c.SubmitChange(ctx, req)
			So(grpcutil.Code(err), ShouldEqual, codes.ResourceExhausted)
		})
	})
}

func newMockPbClient(handler func(w http.ResponseWriter, r *http.Request)) (*httptest.Server, gerritpb.GerritClient) {
	// TODO(tandrii): rename this func once newMockClient name is no longer used in the same package.
	srv := httptest.NewServer(http.HandlerFunc(handler))
	return srv, &client{testBaseURL: srv.URL}
}

// parseTime parses a RFC3339Nano formatted timestamp string.
// Panics when error occurs during parse.
func parseTime(t string) time.Time {
	ret, err := time.Parse(time.RFC3339Nano, t)
	if err != nil {
		panic(err)
	}
	return ret
}
