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

package buildbucket

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gomodule/redigo/redis"
	"google.golang.org/genproto/protobuf/field_mask"

	buildbucketpb "go.chromium.org/luci/buildbucket/proto"
	bbv1 "go.chromium.org/luci/common/api/buildbucket/buildbucket/v1"
	"go.chromium.org/luci/common/clock"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/common/retry/transient"
	"go.chromium.org/luci/common/sync/parallel"
	"go.chromium.org/luci/common/tsmon/metric"
	"go.chromium.org/luci/gae/service/datastore"
	"go.chromium.org/luci/milo/common"
	"go.chromium.org/luci/milo/common/model"
	"go.chromium.org/luci/server/auth"
	"go.chromium.org/luci/server/redisconn"
	"go.chromium.org/luci/server/router"
)

// BuildSummaryStorageDuration is the maximum lifetime of a BuildSummary.
//
// Lifetime is the time elapsed since the Build creation time.
// Cron runs periodically to scan and remove all the Builds of which lifetime
// exceeded this duration.
//
// BuildSummaries are kept alive longer than builds in buildbuckets. So we can
// compute blamelist for builds that are at the end of their lifetime.
//
// TODO(weiweilin): expose BuildStorageDuration from buildbucket and compute
// BuildSummaryStorageDuration base on that (e.g. add two months). So we can
// ensure BuildSummaries are kept alive longer than builds.
const BuildSummaryStorageDuration = time.Hour * 24 * 30 * 20 // ~20 months

var (
	deletedBuildsCounter = metric.NewCounter(
		"luci/milo/cron/delete-builds/delete-count",
		"The number of BuildSummaries deleted by Milo delete-builds cron job",
		nil,
	)
)

// PubSubHandler is a webhook that stores the builds coming in from pubsub.
func PubSubHandler(ctx *router.Context) {
	err := pubSubHandlerImpl(ctx.Context, ctx.Request)
	if err != nil {
		logging.Errorf(ctx.Context, "error while handling pubsub event")
		errors.Log(ctx.Context, err)
	}
	if transient.Tag.In(err) {
		// Transient errors are 4xx so that PubSub retries them.
		// TODO(crbug.com/1099036): Address High traffic builders causing errors.
		ctx.Writer.WriteHeader(http.StatusTooEarly)
		return
	}
	// No errors or non-transient errors are 200s so that PubSub does not retry
	// them.
	ctx.Writer.WriteHeader(http.StatusOK)
}

// DeleteOldBuilds is a cron job that deletes BuildSummaries that are older than
// BuildSummaryStorageDuration.
func DeleteOldBuilds(c context.Context) error {
	const (
		batchSize = 128
		nWorkers  = 8
	)

	buildPurgeCutoffTime := clock.Now(c).Add(-BuildSummaryStorageDuration)
	q := datastore.NewQuery("BuildSummary").
		Lt("Created", buildPurgeCutoffTime).
		Order("Created").
		// Apply a limit so the call won't timeout.
		Limit(batchSize * nWorkers * 512).
		KeysOnly(true)

	return parallel.FanOutIn(func(taskC chan<- func() error) {
		buildsC := make(chan []*datastore.Key, nWorkers)
		statsC := make(chan int, nWorkers)

		// Collect and log stats.
		taskC <- func() error {
			start := clock.Now(c)
			totalDeletedCount := 0
			for deletedCount := range statsC {
				totalDeletedCount += deletedCount
				deletedBuildsCounter.Add(c, int64(totalDeletedCount))
			}

			logging.Infof(c, "took %v to delete %d build summaries", clock.Since(c, start), totalDeletedCount)
			return nil
		}

		// Find builds to delete.
		taskC <- func() error {
			defer close(buildsC)

			bsKeys := make([]*datastore.Key, 0, batchSize)
			err := datastore.RunBatch(c, batchSize, q, func(key *datastore.Key) error {
				bsKeys = append(bsKeys, key)
				if len(bsKeys) == batchSize {
					buildsC <- bsKeys
					bsKeys = make([]*datastore.Key, 0, batchSize)
				}
				return nil
			})
			if err != nil {
				return err
			}

			if len(bsKeys) > 0 {
				buildsC <- bsKeys
			}
			return nil
		}

		// Spawn workers to delete builds.
		taskC <- func() error {
			defer close(statsC)

			return parallel.WorkPool(nWorkers, func(workC chan<- func() error) {
				for bks := range buildsC {
					// Bind to a local variable so each worker can have their own copy.
					bks := bks
					workC <- func() error {
						// Flatten first w/o filtering to calculate how many builds were
						// actually removed.
						err := errors.Flatten(datastore.Delete(c, bks))

						allErrs := 0
						badErrs := 0
						if err != nil {
							allErrs = len(err.(errors.MultiError))
							err = errors.Flatten(errors.Filter(err, datastore.ErrNoSuchEntity))
							badErrs = len(err.(errors.MultiError))
						}

						logging.Infof(c, "Removed %d out of %d build summaries (%d errors, %d already gone)",
							len(bks)-allErrs, len(bks), badErrs, allErrs-badErrs)
						statsC <- len(bks) - allErrs

						return err
					}
				}
			})
		}
	})
}

var summaryBuildMask = &field_mask.FieldMask{
	Paths: []string{
		"id",
		"builder",
		"number",
		"create_time",
		"start_time",
		"end_time",
		"update_time",
		"status",
		"summary_markdown",
		"tags",
		"infra.swarming",
		"input.experimental",
		"input.gitiles_commit",
		"output.properties",
		"critical",
	},
}

// pubSubHandlerImpl takes the http.Request, expects to find
// a common.PubSubSubscription JSON object in the Body, containing a bbPSEvent,
// and handles the contents with generateSummary.
func pubSubHandlerImpl(c context.Context, r *http.Request) error {
	msg := common.PubSubSubscription{}
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		// This might be a transient error, e.g. when the json format changes
		// and Milo isn't updated yet.
		return errors.Annotate(err, "could not decode message").Tag(transient.Tag).Err()
	}
	if v, ok := msg.Message.Attributes["version"].(string); ok && v != "v1" {
		// TODO(nodir): switch to v2, crbug.com/826006
		logging.Debugf(c, "unsupported pubsub message version %q. Ignoring", v)
		return nil
	}
	bData, err := msg.GetData()
	if err != nil {
		return errors.Annotate(err, "could not parse pubsub message string").Err()
	}

	event := struct {
		Build    bbv1.LegacyApiCommonBuildMessage `json:"build"`
		Hostname string                           `json:"hostname"`
	}{}
	if err := json.Unmarshal(bData, &event); err != nil {
		return errors.Annotate(err, "could not parse pubsub message data").Err()
	}

	client, err := buildbucketBuildsClient(c, event.Hostname, auth.AsSelf)
	if err != nil {
		return err
	}
	build, err := client.GetBuild(c, &buildbucketpb.GetBuildRequest{
		Id:     event.Build.Id,
		Fields: summaryBuildMask,
	})
	if err != nil {
		return err
	}

	// TODO(iannucci,nodir): get the bot context too
	// TODO(iannucci,nodir): support manifests/got_revision
	bs, err := model.BuildSummaryFromBuild(c, event.Hostname, build)
	if err != nil {
		return err
	}
	if err := bs.AddManifestKeysFromBuildSets(c); err != nil {
		return err
	}

	logging.Debugf(c, "Received from %s: build %s (%s)\n%v",
		event.Hostname, bs.ProjectID, bs.BuildID, bs.Summary.Status, bs)

	return updateBuild(c, bs)
}

func updateBuild(c context.Context, bs *model.BuildSummary) error {
	now := time.Now()
	updateBuilderSummary, err := shouldUpdateBuilderSummary(c, bs)
	if err != nil {
		updateBuilderSummary = true
		logging.WithError(err).Warningf(c, "failed to determine whether the builder summary from %s should be updated. Fallback to always update.", bs.BuilderID)
	}
	logging.Infof(c, "took %v to determine whether the builder summary from %s should be updated", time.Since(now), bs.BuilderID)

	now = time.Now()
	err = transient.Tag.Apply(datastore.RunInTransaction(c, func(c context.Context) error {
		curBS := &model.BuildSummary{BuildKey: bs.BuildKey}
		switch err := datastore.Get(c, curBS); err {
		case datastore.ErrNoSuchEntity:
			// continue
		default:
			return errors.Annotate(err, "reading current BuildSummary").Err()
		}

		if bs.Version <= curBS.Version {
			logging.Warningf(c, "current BuildSummary is newer: %d <= %d",
				bs.Version, curBS.Version)
			return nil
		}

		if err := datastore.Put(c, bs); err != nil {
			return err
		}

		if !updateBuilderSummary {
			logging.Infof(c, "skipping builder summary update for builder: %s, with build status: %s", bs.BuilderID, bs.Summary.Status)
			return nil
		}

		return model.UpdateBuilderForBuild(c, bs)
	}, nil))

	logging.Infof(c, "took %v to update builder summary from %s", time.Since(now), bs.BuilderID)
	return err
}

var (
	// Sustained datastore updates should not be higher than once per second per
	// entity.
	// See https://cloud.google.com/datastore/docs/concepts/limits.
	entityUpdateIntervalInS int64 = 2

	// A Redis LUA script that update the key with the new integer value if and
	// only if the provided value is greater than the recorded value (0 if none
	// were recorded). It returns 1 if the value is updated and 0 otherwise.
	updateIfLargerScript = redis.NewScript(1, `
local newValueStr = ARGV[1]
local newValue = tonumber(newValueStr)
local existingValueStr = redis.call('GET', KEYS[1])
local existingValue = tonumber(existingValueStr) or 0
if newValue < existingValue then
	return 0
elseif newValue == existingValue then
	-- large u64/i64 (>2^53) may lose precision after being converted to f64.
	-- Compare the last 5 digits if the f64 presentation of the integers are the
	-- same.
	local newValue = tonumber(string.sub(newValueStr, -5)) or 0
	local existingValue = tonumber(string.sub(existingValueStr, -5)) or 0
	if newValue <= existingValue then
		return 0
	end
end

redis.call('SET', KEYS[1], newValueStr, 'EX', ARGV[2])
return 1
`)
)

// shouldUpdateBuilderSummary determines whether the builder summary should be
// updated with the provided build summary.
//
// If the function is called with builds from the same builder multiple times
// within a time bucket, it will block the function until the start of the next
// time bucket and only return true for the call with the lastest created build
// with a terminal status, and return false for other calls occurred within the
// same time bucket.
func shouldUpdateBuilderSummary(c context.Context, buildSummary *model.BuildSummary) (bool, error) {
	if !buildSummary.Summary.Status.Terminal() {
		return false, nil
	}

	conn, err := redisconn.Get(c)
	if err != nil {
		return true, err
	}
	defer conn.Close()

	createdAt := buildSummary.Created.UnixNano()

	now := time.Now().Unix()
	thisTimeBucket := time.Unix(now-now%entityUpdateIntervalInS, 0)
	thisTimeBucketKey := fmt.Sprintf("%s:%v", buildSummary.BuilderID, thisTimeBucket)

	// Check if there's a builder summary update occurred in this time bucket.
	_, err = redis.String(conn.Do("SET", thisTimeBucketKey, createdAt, "NX", "EX", entityUpdateIntervalInS+1))
	switch err {
	case redis.ErrNil:
		// continue
	case nil:
		// There's no builder summary update occurred in this time bucket yet, we
		// should run the update.
		return true, nil
	default:
		return true, err
	}

	// There's already a builder summary update occurred in this time bucket. Try
	// scheduling the update for the next time bucket instead.
	nextTimeBucket := thisTimeBucket.Add(time.Duration(entityUpdateIntervalInS * int64(time.Second)))
	nextTimeBucketKey := fmt.Sprintf("%s:%v", buildSummary.BuilderID, nextTimeBucket)
	replaced, err := redis.Int(updateIfLargerScript.Do(conn, nextTimeBucketKey, createdAt, entityUpdateIntervalInS+1))
	if err != nil {
		return true, err
	}

	if replaced == 0 {
		// There's already an update with a newer build scheduled for the next time
		// bucket, skip the update.
		return false, nil
	}

	// Wait until the start of the next time bucket.
	time.Sleep(time.Until(nextTimeBucket))
	newCreatedAt, err := redis.Int64(conn.Do("GET", nextTimeBucketKey))
	if err != nil {
		return true, err
	}

	// If the update event was not replaced by event by a newer build, we should
	// run the update with this build.
	if newCreatedAt == createdAt {
		return true, nil
	}

	// Otherwise, skip the update.
	logging.Infof(c, "skipping BuilderSummary update from builder %s because a newer build was found", buildSummary.BuilderID)
	return false, nil
}
