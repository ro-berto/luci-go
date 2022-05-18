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

package bbfacade

import (
	"context"
	"strconv"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"

	bbpb "go.chromium.org/luci/buildbucket/proto"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/config/validation"

	"go.chromium.org/luci/cv/api/recipe/v1"
	"go.chromium.org/luci/cv/internal/tryjob"
)

func parseStatusAndResult(ctx context.Context, b *bbpb.Build) (tryjob.Status, *tryjob.Result, error) {
	s := tryjob.Status_STATUS_UNSPECIFIED
	r := &tryjob.Result{
		CreateTime: b.CreateTime,
		UpdateTime: b.UpdateTime,
		Backend: &tryjob.Result_Buildbucket_{
			Buildbucket: &tryjob.Result_Buildbucket{
				Id:              b.Id,
				Status:          b.Status,
				SummaryMarkdown: b.SummaryMarkdown,
			},
		},
	}

	buildResult := parseBuildResult(ctx, b)
	r.Output = buildResult.output
	if buildResult.err != nil {
		logging.Debugf(ctx, "errors parsing recipe output: %s", buildResult.err)
		if buildResult.err.WithSeverity(validation.Blocking) != nil {
			r.Output = &recipe.Output{}
			logging.Debugf(ctx, "ignoring recipe output due to blocking parsing errors")
		}
	}

	switch buildStatus := b.Status; {
	case buildStatus == bbpb.Status_SUCCESS:
		s = tryjob.Status_ENDED
		r.Status = tryjob.Result_SUCCEEDED
	case b.GetStatusDetails().GetTimeout() != nil:
		s = tryjob.Status_ENDED
		r.Status = tryjob.Result_TIMEOUT
	case buildStatus == bbpb.Status_FAILURE:
		s = tryjob.Status_ENDED
		if buildResult.isTransFailure {
			r.Status = tryjob.Result_FAILED_TRANSIENTLY
		} else {
			r.Status = tryjob.Result_FAILED_PERMANENTLY
		}
	case buildStatus == bbpb.Status_CANCELED:
		// For consistency with existing CQD behavior, non-timeout
		// cancellations are treated as transient failures.
		//
		// This behavior is probably a bug in CQD, but it's become expected.
		//
		// TODO(crbug.com/1317392): Revisit the handling of explicitly cancelled
		// tryjobs.
		fallthrough
	case buildStatus == bbpb.Status_INFRA_FAILURE:
		s = tryjob.Status_ENDED
		r.Status = tryjob.Result_FAILED_TRANSIENTLY
	case buildStatus == bbpb.Status_STARTED:
		fallthrough
	case buildStatus == bbpb.Status_SCHEDULED:
		s = tryjob.Status_TRIGGERED
		r.Status = tryjob.Result_UNKNOWN
	default:
		return s, nil, errors.Reason("unexpected buildbucket status %q", b.Status).Err()
	}
	return s, r, nil
}

type outputParsingResult struct {
	// output is the protobuf with the details of the build's output.
	// See https://pkg.go.dev/go.chromium.org/luci/cv/api/recipe/v1#Output documentation.
	output *recipe.Output

	// isTransFailure indicates that based on the properties, the build's
	// failure should be treated as transient.
	isTransFailure bool

	// err indicates issues parsing the build output properties.
	err *validation.Error
}

const transientFailureType = "TRANSIENT_FAILURE"

// outputPropKeys are the keys in the output properties that CV is interested
// in.
var outputPropKeys = []string{
	// New protobuf-based property.
	"$recipe_engine/cq/output",
	// Legacy.
	"do_not_retry",
	"failure_type",
	"triggered_build_ids",
}

func parseBuildResult(ctx context.Context, b *bbpb.Build) *outputParsingResult {
	pr := &outputParsingResult{}
	vc := validation.Context{Context: ctx}
	defer func() {
		if err := vc.Finalize(); err != nil {
			pr.err = err.(*validation.Error)
		}
	}()

	props := b.GetOutput().GetProperties()
	if !hasCVRelatedPropKey(props) {
		return pr
	}

	pr.output = &recipe.Output{}
	if outputVal, ok := props.GetFields()["$recipe_engine/cq/output"]; ok {
		vc.Enter("parsing $recipe_engine/cq/output")
		if output, err := protojson.Marshal(outputVal); err != nil {
			vc.Error(err)
		} else if err := protojson.Unmarshal(output, pr.output); err != nil {
			vc.Error(err)
		}
		vc.Exit()
	}

	vc.Enter("<parsing legacy properties>")
	if dnr, dnrPropertySet := props.GetFields()["do_not_retry"]; dnrPropertySet {
		vc.Enter("do_not_retry")
		switch v, ok := dnr.GetKind().(*structpb.Value_BoolValue); {
		case !ok:
			vc.Errorf("expected a boolean value for field do_not_retry; got %+v", dnr)
		case pr.output.Retry != recipe.Output_OUTPUT_RETRY_UNSPECIFIED:
			// If it has been set by the protobuf field, do not change it.
		case v.BoolValue:
			pr.output.Retry = recipe.Output_OUTPUT_RETRY_DENIED
		default:
			pr.output.Retry = recipe.Output_OUTPUT_RETRY_ALLOWED
		}
		vc.Exit()
	}

	if failureType := props.GetFields()["failure_type"]; failureType.GetStringValue() == transientFailureType {
		pr.isTransFailure = true
	}

	// If this has been set by the protobuf field (there's at least one
	// triggered build id already), do not change it.
	if triggeredBuilds, ok := props.GetFields()["triggered_build_ids"]; ok && len(pr.output.TriggeredBuildIds) == 0 {
		vc.Enter("triggered_build_ids")
		if _, ok := triggeredBuilds.GetKind().(*structpb.Value_ListValue); ok {
			for _, v := range triggeredBuilds.GetListValue().GetValues() {
				// Support both str and int values for robustness.
				switch v := v.GetKind().(type) {
				case *structpb.Value_NumberValue:
					pr.output.TriggeredBuildIds = append(pr.output.TriggeredBuildIds, int64(v.NumberValue))
				case *structpb.Value_StringValue:
					// These may be encoded as string to avoid loss of precision
					// (structpb encodes numeric values as float64).
					intVal, err := strconv.ParseInt(v.StringValue, 10, 64)
					if err != nil {
						vc.Errorf("unable to parse %q as a build_id", v.StringValue)
						continue
					}
					pr.output.TriggeredBuildIds = append(pr.output.TriggeredBuildIds, intVal)
				default:
					vc.Errorf("value of unexpected type %+v", v)
				}
			}
		} else {
			vc.Errorf("expected a list value instead of %+v", triggeredBuilds)
		}
		vc.Exit()
	}
	vc.Exit()
	return pr
}

func hasCVRelatedPropKey(props *structpb.Struct) bool {
	for _, key := range outputPropKeys {
		if _, ok := props.GetFields()[key]; ok {
			return true
		}
	}
	return false
}