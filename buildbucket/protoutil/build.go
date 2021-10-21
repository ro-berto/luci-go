// Copyright 2019 The LUCI Authors.
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

package protoutil

import (
	"time"

	"github.com/golang/protobuf/ptypes"

	"go.chromium.org/luci/common/data/strpair"

	pb "go.chromium.org/luci/buildbucket/proto"
)

// BuildMediaType is a media type for a binary-encoded Build message.
const BuildMediaType = "application/luci+proto; message=buildbucket.v2.Build"

// RunDuration returns duration between build start and end.
func RunDuration(b *pb.Build) (duration time.Duration, ok bool) {
	start, startErr := ptypes.Timestamp(b.StartTime)
	end, endErr := ptypes.Timestamp(b.EndTime)
	if startErr != nil || start.IsZero() || endErr != nil || end.IsZero() {
		return 0, false
	}

	return end.Sub(start), true
}

// SchedulingDuration returns duration between build creation and start.
func SchedulingDuration(b *pb.Build) (duration time.Duration, ok bool) {
	create, createErr := ptypes.Timestamp(b.CreateTime)
	start, startErr := ptypes.Timestamp(b.StartTime)
	if createErr != nil || create.IsZero() || startErr != nil || start.IsZero() {
		return 0, false
	}

	return start.Sub(create), true
}

// Tags parses b.Tags as a strpair.Map.
func Tags(b *pb.Build) strpair.Map {
	m := make(strpair.Map, len(b.Tags))
	for _, t := range b.Tags {
		m.Add(t.Key, t.Value)
	}
	return m
}

// SetStatus sets the status field on `b`, and also correctly adjusts
// StartTime and EndTime.
//
// Will set UpdateTime iff `st` is different than `b.Status`.
func SetStatus(now time.Time, b *pb.Build, st pb.Status) error {
	if b.Status == st {
		return nil
	}
	b.Status = st

	nowPb, err := ptypes.TimestampProto(now.UTC())
	if err != nil {
		return err
	}

	b.UpdateTime = nowPb
	switch {
	case st == pb.Status_STARTED:
		if b.StartTime == nil {
			b.StartTime = nowPb
		}

	case IsEnded(st):
		if b.EndTime == nil {
			b.EndTime = nowPb
		}
	}

	return nil
}
