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
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"golang.org/x/net/context/ctxhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	gerritpb "go.chromium.org/luci/common/proto/gerrit"
)

const (
	// OAuthScope is the OAuth 2.0 scope that must be included when acquiring an
	// access token for Gerrit RPCs.
	OAuthScope = "https://www.googleapis.com/auth/gerritcodereview"

	// contentTypeJSON is the http header content-type value for json encoded
	// objects in the body.
	contentTypeJSON = "application/json; charset=UTF-8"

	// contentTypeText is the http header content-type value for plain text body.
	contentTypeText = "application/x-www-form-urlencoded; charset=UTF-8"

	// defaultQueryLimit is the default limit for ListChanges, and
	// maxQueryLimit is the maximum allowed limit. If either of these are
	// changed, the proto comments should also be updated.
	defaultQueryLimit = 25
	maxQueryLimit     = 1000
)

// This file implements Gerrit proto service client on top of Gerrit REST API.
// WARNING: The returned client is incomplete, so if you want access to a
// particular field from this API, you may need to update the relevant struct
// and add an unmarshalling of that field.

// NewRESTClient creates a new Gerrit client based on Gerrit's REST API.
//
// The host must be a full Gerrit host, e.g. "chromium-review.googlesource.com".
//
// If `auth` is true, this indicates that the given HTTP client sends
// authenticated requests. If so, the requests to Gerrit will include "/a/" URL
// path prefix.
//
// RPC methods of the returned client return an error if a grpc.CallOption is
// passed.
func NewRESTClient(httpClient *http.Client, host string, auth bool) (gerritpb.GerritClient, error) {
	if strings.Contains(host, "/") {
		return nil, errors.Reason("invalid host %q", host).Err()
	}
	baseURL := "https://" + host
	if auth {
		baseURL += "/a"
	}
	return &client{hClient: httpClient, baseURL: baseURL}, nil
}

// Implementation.

// jsonPrefix is expected in all JSON responses from the Gerrit REST API.
var jsonPrefix = []byte(")]}'")

// client implements gerritpb.GerritClient.
type client struct {
	hClient *http.Client
	// baseURL is the base URL for all API requests,
	// for example "https://chromium-review.googlesource.com/a".
	baseURL string
}

func (c *client) ListChanges(ctx context.Context, req *gerritpb.ListChangesRequest, opts ...grpc.CallOption) (*gerritpb.ListChangesResponse, error) {
	limit := req.Limit
	if req.Limit < 0 {
		return nil, errors.Reason("field Limit %d must be nonnegative", req.Limit).Err()
	} else if req.Limit == 0 {
		limit = defaultQueryLimit
	} else if req.Limit > maxQueryLimit {
		return nil, errors.Reason("field Limit %d should be at most %d", req.Limit, maxQueryLimit).Err()
	}
	if req.Offset < 0 {
		return nil, errors.Reason("field Offset %d must be nonnegative", req.Offset).Err()
	}

	params := url.Values{}
	params.Add("q", req.Query)
	for _, o := range req.Options {
		params.Add("o", o.String())
	}
	params.Add("n", strconv.FormatInt(limit, 10))
	params.Add("S", strconv.FormatInt(req.Offset, 10))

	var changes []*changeInfo
	if _, err := c.call(ctx, "GET", "/changes/", params, nil, &changes); err != nil {
		return nil, err
	}

	resp := &gerritpb.ListChangesResponse{}
	for _, c := range changes {
		p, err := c.ToProto()
		if err != nil {
			return nil, err
		}
		resp.Changes = append(resp.Changes, p)
	}
	if len(changes) > 0 {
		resp.MoreChanges = changes[len(changes)-1].MoreChanges
	}

	return resp, nil
}

func (c *client) GetChange(ctx context.Context, req *gerritpb.GetChangeRequest, opts ...grpc.CallOption) (
	*gerritpb.ChangeInfo, error) {
	if err := checkArgs(opts, req); err != nil {
		return nil, err
	}

	path := fmt.Sprintf("/changes/%s", gerritChangeIDForRouting(req.Number, req.Project))

	params := url.Values{}
	for _, o := range req.Options {
		params.Add("o", o.String())
	}

	var resp changeInfo
	if _, err := c.call(ctx, "GET", path, params, nil, &resp); err != nil {
		return nil, err
	}
	return resp.ToProto()
}

type changeInput struct {
	Project    string `json:"project"`
	Branch     string `json:"branch"`
	Subject    string `json:"subject"`
	BaseCommit string `json:"base_commit"`
}

func (c *client) CreateChange(ctx context.Context, req *gerritpb.CreateChangeRequest, opts ...grpc.CallOption) (*gerritpb.ChangeInfo, error) {
	var resp changeInfo
	data := &changeInput{
		Project:    req.Project,
		Branch:     req.Ref,
		Subject:    req.Subject,
		BaseCommit: req.BaseCommit,
	}

	if _, err := c.call(ctx, "POST", "/changes/", url.Values{}, data, &resp, http.StatusCreated); err != nil {
		return nil, errors.Annotate(err, "create empty change").Err()
	}

	switch ci, err := resp.ToProto(); {
	case err != nil:
		return nil, err
	case ci.Status != gerritpb.ChangeStatus_NEW:
		return nil, errors.Reason("unknown status %s for newly created change", ci.Status).Err()
	default:
		return ci, nil
	}
}

func (c *client) ChangeEditFileContent(ctx context.Context, req *gerritpb.ChangeEditFileContentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	path := fmt.Sprintf("/changes/%s/edit/%s", gerritChangeIDForRouting(req.Number, req.Project), url.PathEscape(req.FilePath))
	if _, _, err := c.callRaw(ctx, "PUT", path, url.Values{}, textInputHeaders(), req.Content, http.StatusNoContent); err != nil {
		return nil, errors.Annotate(err, "change edit file content").Err()
	}
	return &emptypb.Empty{}, nil
}

func (c *client) DeleteEditFileContent(ctx context.Context, req *gerritpb.DeleteEditFileContentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	path := fmt.Sprintf("/changes/%s/edit/%s", gerritChangeIDForRouting(req.Number, req.Project), url.PathEscape(req.FilePath))
	var data struct{}
	// The response cannot be JSON-deserialized.
	if _, err := c.call(ctx, "DELETE", path, url.Values{}, &data, nil, http.StatusNoContent); err != nil {
		return nil, errors.Annotate(err, "delete edit file content").Err()
	}
	return &emptypb.Empty{}, nil
}

func (c *client) ChangeEditPublish(ctx context.Context, req *gerritpb.ChangeEditPublishRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	path := fmt.Sprintf("/changes/%s/edit:publish", gerritChangeIDForRouting(req.Number, req.Project))
	if _, _, err := c.callRaw(ctx, "POST", path, url.Values{}, textInputHeaders(), []byte{}, http.StatusNoContent); err != nil {
		return nil, errors.Annotate(err, "change edit publish").Err()
	}
	return &emptypb.Empty{}, nil
}

func (c *client) AddReviewer(ctx context.Context, req *gerritpb.AddReviewerRequest, opts ...grpc.CallOption) (*gerritpb.AddReviewerResult, error) {
	var resp addReviewerResult
	data := &addReviewerRequest{
		Reviewer:  req.Reviewer,
		State:     enumToString(int32(req.State.Number()), gerritpb.AddReviewerRequest_State_name),
		Confirmed: req.Confirmed,
		Notify:    enumToString(int32(req.Notify.Number()), gerritpb.Notify_name),
	}
	path := fmt.Sprintf("/changes/%s/reviewers", gerritChangeIDForRouting(req.Number, req.Project))
	if _, err := c.call(ctx, "POST", path, url.Values{}, data, &resp); err != nil {
		return nil, errors.Annotate(err, "add reviewers").Err()
	}
	rr, err := resp.ToProto()
	if err != nil {
		return nil, errors.Annotate(err, "decoding response").Err()
	}
	return rr, nil
}

func (c *client) DeleteReviewer(ctx context.Context, req *gerritpb.DeleteReviewerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	path := fmt.Sprintf("/changes/%s/reviewers/%s/delete", gerritChangeIDForRouting(req.Number, req.Project), url.PathEscape(req.AccountId))
	if _, err := c.call(ctx, "POST", path, url.Values{}, nil, nil, http.StatusNoContent); err != nil {
		return nil, errors.Annotate(err, "delete reviewer").Err()
	}
	return &emptypb.Empty{}, nil
}

func (c *client) SetReview(ctx context.Context, in *gerritpb.SetReviewRequest, opts ...grpc.CallOption) (*gerritpb.ReviewResult, error) {
	if err := checkArgs(opts, in); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/changes/%s/revisions/%s/review", gerritChangeIDForRouting(in.Number, in.Project), in.RevisionId)
	data := reviewInput{
		Message:                          in.Message,
		Tag:                              in.Tag,
		Notify:                           enumToString(int32(in.Notify.Number()), gerritpb.Notify_name),
		NotifyDetails:                    toNotifyDetails(in.GetNotifyDetails()),
		OnBehalfOf:                       in.OnBehalfOf,
		Ready:                            in.Ready,
		WorkInProgress:                   in.WorkInProgress,
		AddToAttentionSet:                toAttentionSetInputs(in.GetAddToAttentionSet()),
		RemoveFromAttentionSet:           toAttentionSetInputs(in.GetRemoveFromAttentionSet()),
		IgnoreAutomaticAttentionSetRules: in.IgnoreAutomaticAttentionSetRules,
	}
	if in.Labels != nil {
		data.Labels = make(map[string]int32)
		for k, v := range in.Labels {
			data.Labels[k] = v
		}
	}
	var resp gerritpb.ReviewResult
	if _, err := c.call(ctx, "POST", path, url.Values{}, &data, &resp); err != nil {
		return nil, errors.Annotate(err, "set review").Err()
	}
	return &resp, nil
}

func (c *client) AddToAttentionSet(ctx context.Context, req *gerritpb.AttentionSetRequest, opts ...grpc.CallOption) (*gerritpb.AccountInfo, error) {
	path := fmt.Sprintf("/changes/%s/attention", gerritChangeIDForRouting(req.Number, req.Project))
	var data *attentionSetInput
	if i := req.GetInput(); i != nil {
		data = toAttentionSetInput(i)
	} else {
		// TODO(tandrii): migrate current users and remove.
		data = &attentionSetInput{
			User:   req.User,
			Reason: req.Reason,
			Notify: enumToString(int32(req.Notify.Number()), gerritpb.Notify_name),
		}
	}
	var resp accountInfo
	if _, err := c.call(ctx, "POST", path, url.Values{}, data, &resp); err != nil {
		return nil, errors.Annotate(err, "add to attention set").Err()
	}
	return resp.ToProto(), nil
}

func (c *client) SubmitChange(ctx context.Context, req *gerritpb.SubmitChangeRequest, opts ...grpc.CallOption) (*gerritpb.ChangeInfo, error) {
	var resp changeInfo
	path := fmt.Sprintf("/changes/%s/submit", gerritChangeIDForRouting(req.Number, req.Project))
	var data struct{}
	if _, err := c.call(ctx, "POST", path, url.Values{}, &data, &resp); err != nil {
		return nil, errors.Annotate(err, "submit change").Err()
	}
	return resp.ToProto()
}

func (c *client) AbandonChange(ctx context.Context, req *gerritpb.AbandonChangeRequest, opts ...grpc.CallOption) (*gerritpb.ChangeInfo, error) {
	var resp changeInfo
	path := fmt.Sprintf("/changes/%s/abandon", gerritChangeIDForRouting(req.Number, req.Project))
	data := map[string]string{
		"message": req.Message,
	}
	if _, err := c.call(ctx, "POST", path, url.Values{}, &data, &resp); err != nil {
		return nil, errors.Annotate(err, "abandon change").Err()
	}
	return resp.ToProto()
}

func (c *client) SubmitRevision(ctx context.Context, req *gerritpb.SubmitRevisionRequest, opts ...grpc.CallOption) (*gerritpb.SubmitInfo, error) {
	var resp submitInfo
	path := fmt.Sprintf("/changes/%s/revisions/%s/submit", gerritChangeIDForRouting(req.Number, req.Project), req.RevisionId)
	var data struct{}
	if _, err := c.call(ctx, "POST", path, url.Values{}, &data, &resp); err != nil {
		return nil, errors.Annotate(err, "submit revision").Err()
	}
	return resp.ToProto(), nil
}

func (c *client) GetMergeable(ctx context.Context, in *gerritpb.GetMergeableRequest, opts ...grpc.CallOption) (*gerritpb.MergeableInfo, error) {
	var resp mergeableInfo
	path := fmt.Sprintf("/changes/%s/revisions/%s/mergeable", gerritChangeIDForRouting(in.Number, in.Project), in.RevisionId)
	if _, err := c.call(ctx, "GET", path, url.Values{}, nil, &resp); err != nil {
		return nil, errors.Annotate(err, "get mergeable").Err()
	}
	return resp.ToProto()
}

func (c *client) ListFiles(ctx context.Context, req *gerritpb.ListFilesRequest, opts ...grpc.CallOption) (*gerritpb.ListFilesResponse, error) {
	var resp map[string]fileInfo
	params := url.Values{}
	if req.Parent > 0 {
		params.Add("parent", strconv.FormatInt(req.Parent, 10))
	}
	if req.SubstringQuery != "" {
		params.Add("q", req.SubstringQuery)
	}
	if req.Base != "" {
		params.Add("base", req.Base)
	}
	path := fmt.Sprintf("/changes/%s/revisions/%s/files/", gerritChangeIDForRouting(req.Number, req.Project), req.RevisionId)
	if _, err := c.call(ctx, "GET", path, params, nil, &resp); err != nil {
		return nil, errors.Annotate(err, "list files").Err()
	}
	lfr := &gerritpb.ListFilesResponse{
		Files: make(map[string]*gerritpb.FileInfo, len(resp)),
	}
	for k, v := range resp {
		lfr.Files[k] = v.ToProto()
	}
	return lfr, nil
}

func (c *client) GetRelatedChanges(ctx context.Context, req *gerritpb.GetRelatedChangesRequest, opts ...grpc.CallOption) (*gerritpb.GetRelatedChangesResponse, error) {
	// Example:
	// https://chromium-review.googlesource.com/changes/1563638/revisions/2/related
	path := fmt.Sprintf("/changes/%s/revisions/%s/related",
		gerritChangeIDForRouting(req.Number, req.Project), req.RevisionId)
	out := struct {
		Changes []relatedChangeAndCommitInfo `json:"changes"`
	}{}
	if _, err := c.call(ctx, "GET", path, nil, nil, &out); err != nil {
		return nil, errors.Annotate(err, "related changes").Err()
	}
	changes := make([]*gerritpb.GetRelatedChangesResponse_ChangeAndCommit, len(out.Changes))
	for i, c := range out.Changes {
		changes[i] = c.ToProto()
	}
	return &gerritpb.GetRelatedChangesResponse{Changes: changes}, nil
}

func (c *client) GetPureRevert(ctx context.Context, req *gerritpb.GetPureRevertRequest, opts ...grpc.CallOption) (*gerritpb.PureRevertInfo, error) {
	var resp gerritpb.PureRevertInfo
	path := fmt.Sprintf("/changes/%s/pure_revert", gerritChangeIDForRouting(req.Number, req.Project))
	if _, err := c.call(ctx, "GET", path, url.Values{}, nil, &resp); err != nil {
		return nil, errors.Annotate(err, "pure revert").Err()
	}
	return &resp, nil
}

func (c *client) ListFileOwners(ctx context.Context, req *gerritpb.ListFileOwnersRequest, opts ...grpc.CallOption) (*gerritpb.ListOwnersResponse, error) {
	var resp []ownerInfo
	params := url.Values{}
	if req.Options.Details {
		params.Add("o", "DETAILS")
	}
	if req.Options.AllEmails {
		params.Add("o", "ALL_EMAILS")
	}

	path := fmt.Sprintf("/projects/%s/branches/%s/code_owners/%s", url.PathEscape(req.Project), url.PathEscape(req.Ref), url.PathEscape(req.Path))
	if _, err := c.call(ctx, "GET", path, params, nil, &resp); err != nil {
		return nil, errors.Annotate(err, "list file owners").Err()
	}
	owners := make([]*gerritpb.OwnerInfo, len(resp))
	for i, owner := range resp {
		owners[i] = &gerritpb.OwnerInfo{Account: owner.Account.ToProto()}
	}
	return &gerritpb.ListOwnersResponse{
		Owners: owners,
	}, nil
}

func (c *client) ListProjects(ctx context.Context, req *gerritpb.ListProjectsRequest, opts ...grpc.CallOption) (*gerritpb.ListProjectsResponse, error) {
	resp := map[string]*projectInfo{}
	params := url.Values{}
	for _, ref := range req.Refs {
		params.Add("b", ref)
	}
	if _, err := c.call(ctx, "GET", "/projects/", params, nil, &resp); err != nil {
		return nil, errors.Annotate(err, "list projects").Err()
	}
	projectProtos := make(map[string]*gerritpb.ProjectInfo, len(resp))
	for id, p := range resp {
		projectInfo, err := p.ToProto()
		if err != nil {
			return nil, errors.Annotate(err, "decoding response").Err()
		}
		projectProtos[id] = projectInfo
	}
	return &gerritpb.ListProjectsResponse{
		Projects: projectProtos,
	}, nil
}

func (c *client) GetRefInfo(ctx context.Context, req *gerritpb.RefInfoRequest, opts ...grpc.CallOption) (*gerritpb.RefInfo, error) {
	var resp gerritpb.RefInfo
	path := fmt.Sprintf("/projects/%s/branches/%s", url.PathEscape(req.Project), url.PathEscape(req.Ref))
	if _, err := c.call(ctx, "GET", path, url.Values{}, nil, &resp); err != nil {
		return nil, errors.Annotate(err, "get branch info").Err()
	}
	resp.Ref = branchToRef(resp.Ref)
	return &resp, nil
}

// call executes a request to Gerrit REST API with JSON input/output.
// If data is nil, request will be made without a body.
//
// call returns HTTP status code and gRPC error.
// If error happens before HTTP status code was determined, HTTP status code
// will be -1.
func (c *client) call(ctx context.Context, method, urlPath string, params url.Values, data, dest interface{}, expectedHTTPCodes ...int) (int, error) {
	headers := make(map[string]string)

	var rawData []byte
	if data != nil {
		if method == "GET" {
			// This error prevents a more cryptic HTTP error that would otherwise be
			// returned by Gerrit if you try to call it with a GET request and a request
			// body.
			panic("data cannot be provided for a GET request")
		}
		var err error
		rawData, err = json.Marshal(data)
		if err != nil {
			return -1, status.Errorf(codes.Internal, "failed to serialize request message: %s", err)
		}
		headers["Content-Type"] = contentTypeJSON
	}

	ret, body, err := c.callRaw(ctx, method, urlPath, params, headers, rawData, expectedHTTPCodes...)
	body = bytes.TrimPrefix(body, jsonPrefix)
	if err == nil && dest != nil {
		if err = json.Unmarshal(body, dest); err != nil {
			logging.Errorf(ctx, "failed to deserialize response %s; body:\n\n%s", err, string(body))
			return ret, status.Errorf(codes.Internal, "failed to deserialize response: %s", err)
		}
	}
	return ret, err
}

// callRaw executes a request to Gerrit REST API with raw bytes input/output.
//
// callRaw returns HTTP status code and gRPC error.
// If error happens before HTTP status code was determined, HTTP status code
// will be -1.
func (c *client) callRaw(ctx context.Context, method, urlPath string, params url.Values, headers map[string]string, data []byte, expectedHTTPCodes ...int) (int, []byte, error) {
	url := c.baseURL + urlPath
	if len(params) > 0 {
		url += "?" + params.Encode()
	}

	var requestBody io.Reader
	if data != nil {
		requestBody = bytes.NewBuffer(data)
	}
	req, err := http.NewRequest(method, url, requestBody)
	if err != nil {
		return 0, []byte{}, status.Errorf(codes.Internal, "failed to create an HTTP request: %s", err)
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	res, err := ctxhttp.Do(ctx, c.hClient, req)
	if err != nil {
		return -1, []byte{}, status.Errorf(codes.Internal, "failed to execute %s HTTP request: %s", method, err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return res.StatusCode, []byte{}, status.Errorf(codes.Internal, "failed to read response: %s", err)
	}
	logging.Debugf(ctx, "HTTP %d %d <= %s %s", res.StatusCode, len(body), method, url)

	expectedHTTPCodes = append(expectedHTTPCodes, http.StatusOK)
	for _, s := range expectedHTTPCodes {
		if res.StatusCode == s {
			return res.StatusCode, body, nil
		}
	}

	switch res.StatusCode {
	case http.StatusTooManyRequests:
		logging.Warningf(ctx, "Gerrit quota error.\nResponse headers: %v\nResponse body: %s",
			res.Header, body)
		return res.StatusCode, body, status.Errorf(codes.ResourceExhausted, "insufficient Gerrit quota")

	case http.StatusForbidden:
		return res.StatusCode, body, status.Errorf(codes.PermissionDenied, "permission denied")

	case http.StatusNotFound:
		return res.StatusCode, body, status.Errorf(codes.NotFound, "not found")

	case http.StatusConflict:
		// Gerrit returns error message in the response body.
		return res.StatusCode, body, status.Errorf(codes.FailedPrecondition, string(body))

	default:
		logging.Errorf(ctx, "gerrit: unexpected HTTP %d response.\nResponse headers: %v\nResponse body: %s",
			res.StatusCode,
			res.Header, body)
		return res.StatusCode, body, status.Errorf(codes.Internal, "unexpected HTTP %d from Gerrit", res.StatusCode)
	}
}

type validatable interface {
	Validate() error
}

func checkArgs(opts []grpc.CallOption, req validatable) error {
	if len(opts) > 0 {
		return errors.New("gerrit.client does not support grpc options")
	}
	if err := req.Validate(); err != nil {
		return errors.Annotate(err, "request is invalid").Err()
	}
	return nil
}

func branchToRef(ref string) string {
	if strings.HasPrefix(ref, "refs/") {
		// Assume it's actually an explicit full ref.
		// Also assume that ref "refs/heads/refs/XXX will be passed as is by Gerrit
		// instead of "refs/XXX"
		return ref
	}
	return "refs/heads/" + ref
}

func gerritChangeIDForRouting(number int64, project string) string {
	if project != "" {
		return fmt.Sprintf("%s~%d", url.PathEscape(project), number)
	}
	return strconv.Itoa(int(number))
}

func textInputHeaders() map[string]string {
	return map[string]string{
		"Content-Type": contentTypeText,
		"Accept":       contentTypeJSON,
	}
}
