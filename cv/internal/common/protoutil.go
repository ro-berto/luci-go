// Copyright 2021 The LUCI Authors.
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

package common

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// Time2PBNillable is like timestamppb.New() but returns nil on zero time.
func Time2PBNillable(t time.Time) *timestamppb.Timestamp {
	if t.IsZero() {
		return nil
	}
	return timestamppb.New(t)
}

// PB2TimeNillable is the opposite of Time2PBNillable.
//
// ie same as pb.AsTime() but returns zero time if pb is nil.
func PB2TimeNillable(pb *timestamppb.Timestamp) time.Time {
	if pb == nil {
		return time.Time{}
	}
	return pb.AsTime()
}
