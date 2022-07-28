// Copyright 2020 The LUCI Authors.
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

package bqexporter

import (
	"context"

	"cloud.google.com/go/bigquery"
	desc "github.com/golang/protobuf/protoc-gen-go/descriptor"

	"go.chromium.org/luci/common/bq"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/proto/google/descutil"
	"go.chromium.org/luci/resultdb/internal/invocations"
	pb "go.chromium.org/luci/resultdb/proto/v1"
	"go.chromium.org/luci/server/tq"
)

func generateSchema(fdset *desc.FileDescriptorSet, message string) (schema bigquery.Schema, err error) {
	conv := bq.SchemaConverter{
		Desc:           fdset,
		SourceCodeInfo: make(map[*desc.FileDescriptorProto]bq.SourceCodeInfoMap, len(fdset.File)),
	}
	for _, f := range fdset.File {
		conv.SourceCodeInfo[f], err = descutil.IndexSourceCodeInfo(f)
		if err != nil {
			return nil, errors.Annotate(err, "failed to index source code info in file %q", f.GetName()).Err()
		}
	}
	schema, _, err = conv.Schema(message)
	return schema, err
}

func getInvocationIDSet(ctx context.Context, invID invocations.ID, processor func(context.Context, invocations.IDSet) error) error {
	inv, err := invocations.Read(ctx, invID)
	if err != nil {
		return err
	}
	if inv.State != pb.Invocation_FINALIZED {
		return errors.Reason("%s is not finalized yet", invID.Name()).Err()
	}

	// Get the invocation set.
	if err := invocations.BatchedReachable(ctx, invocations.NewIDSet(invID), processor); err != nil {
		if invocations.TooManyTag.In(err) {
			err = tq.Fatal.Apply(err)
		}
		return err
	}
	return nil
}
