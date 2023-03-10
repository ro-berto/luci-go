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

package buganizer

import (
	"context"

	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/third_party/google.golang.org/genproto/googleapis/devtools/issuetracker/v1"
	issuetrackerclient "go.chromium.org/luci/third_party/google.golang.org/google/devtools/issuetracker/v1"
)

// An implementation of the client wrapper that uses the Client provided
// by issuetrackerclient package. This client acts as a delegate and
// a proxy to the actual implementation.
type RPCClient struct {
	Client *issuetrackerclient.Client
}

// NewRPCClient returns a new ClientWrapper.
func NewRPCClient(ctx context.Context) (*RPCClient, error) {
	client, err := issuetrackerclient.NewClient(ctx)
	if err != nil {
		return nil, errors.Annotate(err, "create new wrapper client").Err()
	}
	return &RPCClient{
		Client: client,
	}, nil
}

func (w *RPCClient) Close() {
	w.Client.Close()
}

// BatchGetIssues delegates a call to Client.BatchGetIssues and
// returns the list of issues returned or the error that occured.
func (w *RPCClient) BatchGetIssues(ctx context.Context, in *issuetracker.BatchGetIssuesRequest) (*issuetracker.BatchGetIssuesResponse, error) {
	return w.Client.BatchGetIssues(ctx, in)
}

// GetIssue delegates a call to Client.GetIssue and returns the issue
// returned or the error that occured.
func (w *RPCClient) GetIssue(ctx context.Context, in *issuetracker.GetIssueRequest) (*issuetracker.Issue, error) {
	return w.Client.GetIssue(ctx, in)
}

// CreateIssue delegates a call to Client.CreateIssue and returns the
// issue that was created or the error that occured.
func (w *RPCClient) CreateIssue(ctx context.Context, in *issuetracker.CreateIssueRequest) (*issuetracker.Issue, error) {
	return w.Client.CreateIssue(ctx, in)
}

// ModifyIssue delegates a call to Client.ModifyIssue and returns the
// modified issue or the error that occured.
func (w *RPCClient) ModifyIssue(ctx context.Context, in *issuetracker.ModifyIssueRequest) (*issuetracker.Issue, error) {
	return w.Client.ModifyIssue(ctx, in)
}

// ListIssueUpdates delegates a call to Client.ListIssueUpdates and returns the
// issue updates iterator delegate.
func (w *RPCClient) ListIssueUpdates(ctx context.Context, in *issuetracker.ListIssueUpdatesRequest) IssueUpdateIterator {
	it := w.Client.ListIssueUpdates(ctx, in)
	return it
}

// CreateIssueComment delegates a call to Client.CreateIssueComment and returns
// the comment that was created or the error that occured.
func (w *RPCClient) CreateIssueComment(ctx context.Context, in *issuetracker.CreateIssueCommentRequest) (*issuetracker.IssueComment, error) {
	return w.Client.CreateIssueComment(ctx, in)
}

// ListIssueComments delegates a call to Client.ListIssueComments and returns
// the issue comment iterator delegate.
func (w *RPCClient) ListIssueComments(ctx context.Context, in *issuetracker.ListIssueCommentsRequest) IssueCommentIterator {
	it := w.Client.ListIssueComments(ctx, in)
	return it
}
