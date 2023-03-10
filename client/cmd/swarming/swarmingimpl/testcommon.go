// Copyright 2017 The LUCI Authors.
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

package swarmingimpl

import (
	"context"
	"flag"
	"net/http"

	rbeclient "github.com/bazelbuild/remote-apis-sdks/go/pkg/client"

	"go.chromium.org/luci/common/api/swarming/swarming/v1"
	"google.golang.org/api/googleapi"
)

var _ swarmingService = (*testService)(nil)

type testService struct {
	newTask      func(context.Context, *swarming.SwarmingRpcsNewTaskRequest) (*swarming.SwarmingRpcsTaskRequestMetadata, error)
	countTasks   func(context.Context, float64, string, ...string) (*swarming.SwarmingRpcsTasksCount, error)
	countBots    func(context.Context, ...string) (*swarming.SwarmingRpcsBotsCount, error)
	listTasks    func(context.Context, int64, float64, string, []string, []googleapi.Field) ([]*swarming.SwarmingRpcsTaskResult, error)
	cancelTask   func(context.Context, string, *swarming.SwarmingRpcsTaskCancelRequest) (*swarming.SwarmingRpcsCancelResponse, error)
	taskRequest  func(context.Context, string) (*swarming.SwarmingRpcsTaskRequest, error)
	taskResult   func(context.Context, string, bool) (*swarming.SwarmingRpcsTaskResult, error)
	taskOutput   func(context.Context, string) (*swarming.SwarmingRpcsTaskOutput, error)
	filesFromCAS func(context.Context, string, *rbeclient.Client, *swarming.SwarmingRpcsCASReference) ([]string, error)
	listBots     func(context.Context, []string, []googleapi.Field) ([]*swarming.SwarmingRpcsBotInfo, error)
	deleteBot    func(context.Context, string) (*swarming.SwarmingRpcsDeletedResponse, error)
	terminateBot func(context.Context, string) (*swarming.SwarmingRpcsTerminateResponse, error)
	listBotTasks func(context.Context, string, int64, float64, string, []googleapi.Field) ([]*swarming.SwarmingRpcsTaskResult, error)
}

func (testService) Client() *http.Client {
	return nil
}

func (s testService) NewTask(ctx context.Context, req *swarming.SwarmingRpcsNewTaskRequest) (*swarming.SwarmingRpcsTaskRequestMetadata, error) {
	return s.newTask(ctx, req)
}

func (s testService) CountTasks(ctx context.Context, start float64, state string, tags ...string) (*swarming.SwarmingRpcsTasksCount, error) {
	return s.countTasks(ctx, start, state, tags...)
}

func (s testService) ListTasks(ctx context.Context, limit int64, start float64, state string, tags []string, fields []googleapi.Field) ([]*swarming.SwarmingRpcsTaskResult, error) {
	return s.listTasks(ctx, limit, start, state, tags, fields)
}

func (s testService) CancelTask(ctx context.Context, taskID string, req *swarming.SwarmingRpcsTaskCancelRequest) (*swarming.SwarmingRpcsCancelResponse, error) {
	return s.cancelTask(ctx, taskID, req)
}

func (s testService) TaskRequest(ctx context.Context, taskID string) (*swarming.SwarmingRpcsTaskRequest, error) {
	return s.taskRequest(ctx, taskID)
}

func (s testService) TaskResult(ctx context.Context, taskID string, perf bool) (*swarming.SwarmingRpcsTaskResult, error) {
	return s.taskResult(ctx, taskID, perf)
}

func (s testService) TaskOutput(ctx context.Context, taskID string) (*swarming.SwarmingRpcsTaskOutput, error) {
	return s.taskOutput(ctx, taskID)
}

func (s testService) FilesFromCAS(ctx context.Context, outdir string, cascli *rbeclient.Client, casRef *swarming.SwarmingRpcsCASReference) ([]string, error) {
	return s.filesFromCAS(ctx, outdir, cascli, casRef)
}

func (s testService) CountBots(ctx context.Context, dimensions ...string) (*swarming.SwarmingRpcsBotsCount, error) {
	return s.countBots(ctx, dimensions...)
}

func (s *testService) ListBots(ctx context.Context, dimensions []string, fields []googleapi.Field) ([]*swarming.SwarmingRpcsBotInfo, error) {
	return s.listBots(ctx, dimensions, fields)
}

func (s testService) DeleteBot(ctx context.Context, botID string) (*swarming.SwarmingRpcsDeletedResponse, error) {
	return s.deleteBot(ctx, botID)
}

func (s testService) TerminateBot(ctx context.Context, botID string) (*swarming.SwarmingRpcsTerminateResponse, error) {
	return s.terminateBot(ctx, botID)
}

func (s testService) ListBotTasks(ctx context.Context, botID string, limit int64, start float64, state string, fields []googleapi.Field) ([]*swarming.SwarmingRpcsTaskResult, error) {
	return s.listBotTasks(ctx, botID, limit, start, state, fields)
}

var _ AuthFlags = (*testAuthFlags)(nil)

type testAuthFlags struct{}

func (af *testAuthFlags) Register(_ *flag.FlagSet) {}

func (af *testAuthFlags) Parse() error { return nil }

func (af *testAuthFlags) NewHTTPClient(_ context.Context) (*http.Client, error) {
	return nil, nil
}

func (af *testAuthFlags) NewRBEClient(_ context.Context, _ string, _ string) (*rbeclient.Client, error) {
	return nil, nil
}
