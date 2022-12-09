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

package clients

import (
	"context"
	"regexp"
	"strings"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
	"google.golang.org/grpc"

	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/grpc/grpcmon"
	"go.chromium.org/luci/server/auth"
)

var (
	mockPubsubClientKey = "mock pubsub clients key for testing only"

	// cloudProjectIDRE is the cloud project identifier regex derived from
	// https://cloud.google.com/resource-manager/docs/creating-managing-projects#before_you_begin
	cloudProjectIDRE = regexp.MustCompile(`^[a-z]([a-z0-9-]){4,28}[a-z0-9]$`)
	// topicNameRE is the full topic name regex derived from https://cloud.google.com/pubsub/docs/admin#resource_names
	topicNameRE = regexp.MustCompile(`^projects/(.*)/topics/(.*)$`)
	// topicIDRE is the topic id regex derived from https://cloud.google.com/pubsub/docs/admin#resource_names
	topicIDRE = regexp.MustCompile(`^[A-Za-z]([0-9A-Za-z\._\-~+%]){3,255}$`)
)

func NewPubsubClient(ctx context.Context, cloudProject, luciProject string) (*pubsub.Client, error) {
	if mockClients, ok := ctx.Value(&mockPubsubClientKey).(map[string]*pubsub.Client); ok {
		if mockClient, exist := mockClients[cloudProject]; exist {
			return mockClient, nil
		}
		return nil, errors.Reason("couldn't find mock pubsub client for %s", cloudProject).Err()
	}

	creds, err := auth.GetPerRPCCredentials(ctx, auth.AsProject, auth.WithProject(luciProject), auth.WithScopes(auth.CloudOAuthScopes...))
	if err != nil {
		return nil, err
	}
	client, err := pubsub.NewClient(
		ctx, cloudProject,
		option.WithGRPCDialOption(grpcmon.WithClientRPCStatsMonitor()),
		option.WithGRPCDialOption(grpc.WithPerRPCCredentials(creds)),
	)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// ValidatePubSubTopicName validates the format of topic, extract the cloud project and topic id, and return them.
func ValidatePubSubTopicName(topic string) (string, string, error) {
	matches := topicNameRE.FindAllStringSubmatch(topic, -1)
	if matches == nil || len(matches[0]) != 3 {
		return "", "", errors.Reason("topic %q does not match %q", topic, topicNameRE).Err()
	}

	cloudProj := matches[0][1]
	topicID := matches[0][2]
	// Only internal App Engine projects start "google.com:" with go/gae4g-setup#choosing-the-right-app-engine-version,
	// all other project ids conform to cloudProjectIDRE.
	if !strings.HasPrefix(cloudProj, "google.com:") && !cloudProjectIDRE.MatchString(cloudProj) {
		return "", "", errors.Reason("cloud project id %q does not match %q", cloudProj, cloudProjectIDRE).Err()
	}
	if strings.HasPrefix(topicID, "goog") {
		return "", "", errors.Reason("topic id %q shouldn't begin with the string goog", topicID).Err()
	}
	if !topicIDRE.MatchString(topicID) {
		return "", "", errors.Reason("topic id %q does not match %q", topicID, topicIDRE).Err()
	}
	return cloudProj, topicID, nil
}
