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

package dsmapper

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

	"go.chromium.org/luci/common/clock"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/common/retry/transient"
	"go.chromium.org/luci/gae/service/datastore"

	"go.chromium.org/luci/server/dsmapper/dsmapperpb"
	"go.chromium.org/luci/server/dsmapper/internal/splitter"
)

// ErrNoSuchJob is returned by GetJob if there's no Job with requested ID.
var ErrNoSuchJob = errors.New("no such mapping job")

// Query is a representation of datastore queries supported by the
// mapper.
//
// A query defines a set of entities the mapper operates on.
//
// This struct can be embedded into entities as is.
type Query struct {
	Kind     string         // entity kind to limit the query, "" for kindless
	Ancestor *datastore.Key // entity group to limit the query to (or nil)
}

// ToDatastoreQuery returns corresponding datastore.Query.
func (q *Query) ToDatastoreQuery() *datastore.Query {
	dq := datastore.NewQuery(q.Kind)
	if q.Ancestor != nil {
		dq = dq.Ancestor(q.Ancestor)
	}
	return dq
}

// JobConfig defines what a new mapping job should do.
//
// It should be supplied by the users of the mapper library.
type JobConfig struct {
	Query      Query  // a query identifying a set of entities
	Mapper     ID     // ID of a registered mapper to apply to entities
	Params     []byte // arbitrary user-provided data to pass to the mapper
	ShardCount int    // number of shards to split the key range into
	PageSize   int    // how many entities to process at once in each shard

	// Optional parameters below for fine tunning. They have reasonable defaults,
	// and should generally be not touched.

	// PagesPerTask is how many pages (each of PageSize entities) to process
	// inside a TQ task.
	//
	// Default is unlimited: process until the deadline.
	PagesPerTask int

	// TaskDuration is how long to run a single mapping TQ task before
	// checkpointing the state and launching the next mapping TQ task.
	//
	// Small values (e.g. 1 min) makes each processing TQ task relatively small,
	// so it doesn't eat a lot of memory, or produces gigantic unreadable logs.
	// It also makes TQ's "Pause queue" button more handy.
	//
	// Default is 1 min.
	TaskDuration time.Duration

	// TrackProgress enables calculating number of entities per shard before
	// launching mappers, and using it to calculate completion ETA.
	//
	// May be VERY slow if processing large amount of entities. Slowness manifests
	// as a delay between job's launch and it actual start of shards processing.
	//
	// Enable only if shards are relatively small (< 100K entities per shard).
	TrackProgress bool
}

// Validate returns an error of the config is invalid.
//
// Mapper existence is not checked.
func (jc *JobConfig) Validate() error {
	switch {
	case jc.ShardCount < 1:
		return errors.Reason("ShardCount should be >= 1, try 8").Err()
	case jc.PageSize <= 0:
		return errors.Reason("PageSize should be > 0, try 256").Err()
	case jc.PagesPerTask < 0:
		return errors.Reason("PagesPerTask should be >= 0, keep 0 for default").Err()
	case jc.TaskDuration < 0:
		return errors.Reason("TaskDuration should be >= 0, keep 0 for default").Err()
	}
	return nil
}

// JobID identifies a mapping job.
type JobID int64

// Job is datastore representation of a mapping job (either active or not).
//
// It is a root entity with autogenerated key.
//
// Use Controller and Job methods to work with jobs. Attempting to use datastore
// API directly results in an undefined behavior.
type Job struct {
	_kind  string                `gae:"$kind,mapper.Job"`
	_extra datastore.PropertyMap `gae:"-,extra"`

	// ID is auto-generated unique identifier of the job.
	ID JobID `gae:"$id"`
	// Config is the configuration of this job. Doesn't change once set.
	Config JobConfig `gae:",noindex"`
	// State is used to track job's lifecycle, see the enum.
	State dsmapperpb.State
	// Created is when the job was created, FYI.
	Created time.Time
	// Updated is when the job was last touched, FYI.
	Updated time.Time
}

// shardList is an entity with a list of shard IDs associated with a job.
//
// A "static" singleton child entity of Job. Stored separately to allow callers
// to skip loading (potentially huge) list of shards if they are not interested
// in it.
type shardList struct {
	_kind  string                `gae:"$kind,mapper.ShardList"`
	_id    int64                 `gae:"$id,1"`
	_extra datastore.PropertyMap `gae:"-,extra"`

	Parent *datastore.Key `gae:"$parent"`
	Shards []int64        `gae:",noindex"`
}

// fetchShardIDs fetches IDs of the job shards.
func (j *Job) fetchShardIDs(ctx context.Context) ([]int64, error) {
	l := shardList{Parent: datastore.KeyForObj(ctx, j)}
	switch err := datastore.Get(ctx, &l); {
	case err == datastore.ErrNoSuchEntity:
		return nil, errors.Annotate(err, "broken state, no ShardList entity for job %d", j.ID).Err()
	case err != nil:
		return nil, errors.Annotate(err, "when fetching list of shards of job %d", j.ID).Tag(transient.Tag).Err()
	default:
		return l.Shards, nil
	}
}

// fetchShards fetches all job shards.
func (j *Job) fetchShards(ctx context.Context) ([]shard, error) {
	ids, err := j.fetchShardIDs(ctx)
	if err != nil {
		return nil, err
	}

	shards := make([]shard, len(ids))
	for idx, sid := range ids {
		shards[idx].ID = sid
	}

	if err := datastore.Get(ctx, shards); err != nil {
		return nil, errors.Annotate(err, "failed to fetch some shards of job %d", j.ID).Tag(transient.Tag).Err()
	}
	return shards, nil
}

// FetchInfo fetches information about the job (including all shards).
func (j *Job) FetchInfo(ctx context.Context) (*dsmapperpb.JobInfo, error) {
	info := &dsmapperpb.JobInfo{
		Id:            int64(j.ID),
		State:         j.State,
		Created:       timestamppb.New(j.Created),
		Updated:       timestamppb.New(j.Updated),
		TotalEntities: -1, // assume unknown, will be replaced below if known
	}

	// Jobs in STARTING state have no shards yet.
	if j.State == dsmapperpb.State_STARTING {
		return info, nil
	}

	shards, err := j.fetchShards(ctx)
	if err != nil {
		return nil, err
	}

	haveProgress := true // false if at least one shard has unknown ETA
	updated := j.Updated // will be max(Updated of each shard)

	info.Shards = make([]*dsmapperpb.ShardInfo, len(shards))
	for i, s := range shards {
		sh := s.info()
		info.Shards[i] = sh
		info.ProcessedEntities += sh.ProcessedEntities
		if ts := sh.Updated.AsTime(); ts.After(updated) {
			updated = ts
		}
		if sh.TotalEntities == -1 {
			haveProgress = false
		}
	}

	// Calculate the overall rate from scratch, do NOT sum rates of shards,
	// since it will also sum estimation errors too (which can be wild).
	info.Updated = timestamppb.New(updated)
	if runtime := updated.Sub(j.Created); runtime > 0 {
		info.EntitiesPerSec = float32(float64(info.ProcessedEntities) / runtime.Seconds())
	}

	if haveProgress {
		maxETA := time.Time{}

		info.TotalEntities = 0
		for _, s := range info.Shards {
			info.TotalEntities += s.TotalEntities
			if s.Eta != nil {
				if ts := s.Eta.AsTime(); maxETA.IsZero() || ts.After(maxETA) {
					maxETA = ts
				}
			}
		}

		// The job completes when its longest shard does. Shards do not pass work
		// to each other.
		if !maxETA.IsZero() {
			info.Eta = timestamppb.New(maxETA)
		}
	}

	return info, nil
}

// getJob fetches a Job entity.
//
// Recognizes and tags transient errors.
func getJob(ctx context.Context, id JobID) (*Job, error) {
	job := &Job{ID: id}
	switch err := datastore.Get(ctx, job); {
	case err == datastore.ErrNoSuchEntity:
		return nil, ErrNoSuchJob
	case err != nil:
		return nil, errors.Annotate(err, "transient datastore error").Tag(transient.Tag).Err()
	default:
		return job, nil
	}
}

// getJobInState fetches a Job entity and checks its state.
//
// Returns:
//   (*Job, nil) if the job is there and its state matches one of given states.
//   (nil, nil) if the job is there, but in a different state.
//   (nil, transient error) on datastore fetch errors.
//   (nil, fatal error) if there's no such job at all.
func getJobInState(ctx context.Context, id JobID, states ...dsmapperpb.State) (*Job, error) {
	job, err := getJob(ctx, id)
	if err != nil {
		return nil, errors.Reason("failed to fetch job with ID %d", id).Err()
	}
	for _, s := range states {
		if job.State == s {
			return job, nil
		}
	}
	logging.Warningf(ctx, "Skipping the job: its state is %s, expecting one of %q", job.State, states)
	return nil, nil
}

// shard represents a key range being worked on by a single worker (Start, End].
//
// Shard entities are written to when workers checkpoint progress or finish.
// They are read when calculating overall progress of the job.
//
// It is a root entity with autogenerated key. Shards are associated with jobs
// via ShardList entity (owned by Job, for Job -> [Shard] queries), and via
// JobID property (for Shard -> Job queries). They are purposefully not a part
// of Job entity group, to avoid exceeding O(1) entity group write limit.
type shard struct {
	_kind  string                `gae:"$kind,mapper.Shard"`
	_extra datastore.PropertyMap `gae:"-,extra"`

	// ID is auto-generated unique identifier of the shard.
	ID int64 `gae:"$id"`
	// JobID is ID of a job that owns this shard.
	JobID JobID
	// Index is the index of the shard in the job's shards list.
	Index int `gae:",noindex"`
	// State is used to track shard's lifecycle, see the enum.
	State dsmapperpb.State
	// Error is an error message for failed shards.
	Error string `gae:",noindex"`
	// ProcessTaskNum is next expected ProcessShard task number.
	ProcessTaskNum int64 `gae:",noindex"`
	// Range is an entity key range covered by this shard.
	Range splitter.Range `gae:",noindex"`
	// ExpectedCount is expected number of entities in the shard, -1 if unknown.
	ExpectedCount int64 `gae:",noindex"`
	// ProcessedCount is number entities processed by the shard thus far.
	ProcessedCount int64 `gae:",noindex"`
	// ResumeFrom is the last processed key or nil if just starting.
	ResumeFrom *datastore.Key `gae:",noindex"`
	// Created is when the shard was created, FYI.
	Created time.Time
	// Updated is when the shard was last touched, FYI.
	Updated time.Time
}

// info returns a proto message with information about the shard.
func (s *shard) info() *dsmapperpb.ShardInfo {
	var rate float64
	var eta *timestamp.Timestamp

	if runtime := s.Updated.Sub(s.Created); runtime > 0 {
		rate = float64(s.ProcessedCount) / runtime.Seconds()
		if s.ExpectedCount != -1 && rate > 0.0001 {
			secs := float64(s.ExpectedCount) / rate
			eta = timestamppb.New(s.Created.Add(time.Duration(float64(time.Second) * secs)))
		}
	}

	return &dsmapperpb.ShardInfo{
		Index:             int32(s.Index),
		State:             s.State,
		Error:             s.Error,
		Created:           timestamppb.New(s.Created),
		Updated:           timestamppb.New(s.Updated),
		Eta:               eta, // nil if unknown
		ProcessedEntities: s.ProcessedCount,
		TotalEntities:     s.ExpectedCount, // -1 if unknown
		EntitiesPerSec:    float32(rate),   // 0 if unknown
	}
}

// getActiveShard returns shard entity with given ID if its still in active
// state and its ProcessTaskNum matches the given taskNum.
//
// Returns:
//   (*shard, nil) if the shard is there and matches the criteria.
//   (nil, nil) if the shard is there, but it doesn't match the criteria.
//   (nil, transient error) on datastore fetch errors.
//   (nil, fatal error) if there's no such shard at all.
func getActiveShard(ctx context.Context, shardID, taskNum int64) (*shard, error) {
	sh := &shard{ID: shardID}
	switch err := datastore.Get(ctx, sh); {
	case err == datastore.ErrNoSuchEntity:
		return nil, errors.Annotate(err, "no such shard, aborting").Err() // fatal, no retries
	case err != nil:
		return nil, errors.Annotate(err, "failed to fetch the shard").Tag(transient.Tag).Err()
	case isFinalState(sh.State):
		logging.Warningf(ctx, "The shard is finished already")
		return nil, nil
	case sh.ProcessTaskNum != taskNum:
		logging.Warningf(ctx, "The task is stale (shard's task_num is %d, but task's is %d). Skipping it", sh.ProcessTaskNum, taskNum)
		return nil, nil
	default:
		return sh, nil
	}
}

// shardTxnCb examines and optionally mutates the shard.
//
// It returns (true, nil) to instruct shardTxn to store the shard, (false, nil)
// to skip storing, and (..., err) to return the error.
type shardTxnCb func(ctx context.Context, sh *shard) (save bool, err error)

// shardTxn fetches the shard and calls the callback to examine or mutate it.
//
// Silently skips finished shards.
func shardTxn(ctx context.Context, shardID int64, cb shardTxnCb) error {
	return runTxn(ctx, func(ctx context.Context) error {
		sh := shard{ID: shardID}
		switch err := datastore.Get(ctx, &sh); {
		case err == datastore.ErrNoSuchEntity:
			return err
		case err != nil:
			return transient.Tag.Apply(err)
		case isFinalState(sh.State):
			return nil // the shard is already marked as done
		}
		switch save, err := cb(ctx, &sh); {
		case err != nil:
			return err
		case !save:
			return nil
		default:
			sh.Updated = clock.Now(ctx).UTC()
			return transient.Tag.Apply(datastore.Put(ctx, &sh))
		}
	})
}
