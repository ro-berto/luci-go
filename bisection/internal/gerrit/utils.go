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

package gerrit

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"time"

	"go.chromium.org/luci/common/clock"
	"go.chromium.org/luci/common/errors"
	gerritpb "go.chromium.org/luci/common/proto/gerrit"
	"go.chromium.org/luci/gae/service/info"
)

// getGerritEmailAddress is a helper function to get the email address
// that LUCI Bisection uses to perform Gerrit actions
func getGerritEmailAddress(ctx context.Context) (string, error) {
	emailAddress, err := info.ServiceAccount(ctx)
	if err != nil {
		return "", errors.Annotate(err,
			"failed to get LUCI Bisection email address used for Gerrit").Err()
	}
	return emailAddress, nil
}

// GetHost extracts the Gerrit host from the given Gerrit review URL
func GetHost(ctx context.Context, rawReviewURL string) (string, error) {
	reviewURL := strings.TrimSpace(rawReviewURL)
	pattern := regexp.MustCompile("https://([^/]+)")
	matches := pattern.FindStringSubmatch(reviewURL)
	if matches == nil {
		return "", fmt.Errorf("could not find Gerrit host from review URL = '%s'",
			reviewURL)
	}
	return matches[1], nil
}

// HasLUCIBisectionComment returns whether LUCI Bisection has previously commented
// on the change
func HasLUCIBisectionComment(ctx context.Context, change *gerritpb.ChangeInfo) (bool, error) {
	lbAccount, err := getGerritEmailAddress(ctx)
	if err != nil {
		return false, err
	}

	for _, message := range change.Messages {
		if message.Author != nil {
			if message.Author.Email == lbAccount {
				return true, nil
			}
		}
	}

	return false, nil
}

// IsOwnedByLUCIBisection returns whether the change is owned by LUCI Bisection
func IsOwnedByLUCIBisection(ctx context.Context, change *gerritpb.ChangeInfo) (bool, error) {
	if change.Owner == nil {
		return false, nil
	}

	lbAccount, err := getGerritEmailAddress(ctx)
	if err != nil {
		return false, err
	}

	return change.Owner.Email == lbAccount, nil
}

// IsRecentSubmit returns whether the change was submitted recently, as defined
// by the maximum age duration given relative to now.
func IsRecentSubmit(ctx context.Context, change *gerritpb.ChangeInfo, maxAge time.Duration) bool {
	earliest := clock.Now(ctx).Add(-maxAge)
	submitted := change.Submitted.AsTime()
	return submitted.Equal(earliest) || submitted.After(earliest)
}
