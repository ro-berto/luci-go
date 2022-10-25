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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: go.chromium.org/luci/buildbucket/proto/step.proto

package buildbucketpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A build step.
//
// A step may have children, see name field.
type Step struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the step, unique within the build.
	// Identifies the step.
	//
	// Pipe character ("|") is reserved to separate parent and child step names.
	// For example, value "a|b" indicates step "b" under step "a".
	// If this is a child step, a parent MUST exist and MUST precede this step in
	// the list of steps.
	// All step names, including child and parent names recursively,
	// MUST NOT be an empty string.
	// For example, all of the below names are invalid.
	// - |a
	// - a|
	// - a||b
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The timestamp when the step started.
	//
	// MUST NOT be specified, if status is SCHEDULED.
	// MUST be specified, if status is STARTED, SUCCESS, FAILURE, or INFRA_FAILURE
	// MAY be specified, if status is CANCELED.
	StartTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// The timestamp when the step ended.
	// Present iff status is terminal.
	// MUST NOT be before start_time.
	EndTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// Status of the step.
	// Must be specified, i.e. not STATUS_UNSPECIFIED.
	Status Status `protobuf:"varint,4,opt,name=status,proto3,enum=buildbucket.v2.Status" json:"status,omitempty"`
	// Logs produced by the step.
	// Log order is up to the step.
	//
	// BigQuery: excluded from rows.
	Logs []*Log `protobuf:"bytes,5,rep,name=logs,proto3" json:"logs,omitempty"`
	// MergeBuild is used for go.chromium.org/luci/luciexe to indicate to the
	// luciexe host process if some Build stream should be merged under this step.
	//
	// BigQuery: excluded from rows.
	MergeBuild *Step_MergeBuild `protobuf:"bytes,6,opt,name=merge_build,json=mergeBuild,proto3" json:"merge_build,omitempty"`
	// Human-readable summary of the step provided by the step itself,
	// in Markdown format (https://spec.commonmark.org/0.28/).
	//
	// V1 equivalent: combines and supersedes Buildbot's step_text and step links and also supports
	// other formatted text.
	//
	// BigQuery: excluded from rows.
	SummaryMarkdown string `protobuf:"bytes,7,opt,name=summary_markdown,json=summaryMarkdown,proto3" json:"summary_markdown,omitempty"`
	// Arbitrary annotations for the step.
	//
	// One key may have multiple values, which is why this is not a map<string,string>.
	//
	// These are NOT interpreted by Buildbucket.
	//
	// Tag keys SHOULD indicate the domain/system that interprets them, e.g.:
	//
	//	my_service.category = COMPILE
	//
	// Rather than
	//
	//	is_compile = true
	//
	// This will help contextualize the tag values when looking at a build (who
	// set this tag? who will interpret this tag?))
	//
	// The 'luci.' key prefix is reserved for LUCI's own usage.
	//
	// The Key may not exceed 256 bytes.
	// The Value may not exceed 1024 bytes.
	//
	// Key and Value may not be empty.
	Tags []*StringPair `protobuf:"bytes,8,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *Step) Reset() {
	*x = Step{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_buildbucket_proto_step_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Step) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Step) ProtoMessage() {}

func (x *Step) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_buildbucket_proto_step_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Step.ProtoReflect.Descriptor instead.
func (*Step) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_buildbucket_proto_step_proto_rawDescGZIP(), []int{0}
}

func (x *Step) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Step) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *Step) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *Step) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_STATUS_UNSPECIFIED
}

func (x *Step) GetLogs() []*Log {
	if x != nil {
		return x.Logs
	}
	return nil
}

func (x *Step) GetMergeBuild() *Step_MergeBuild {
	if x != nil {
		return x.MergeBuild
	}
	return nil
}

func (x *Step) GetSummaryMarkdown() string {
	if x != nil {
		return x.SummaryMarkdown
	}
	return ""
}

func (x *Step) GetTags() []*StringPair {
	if x != nil {
		return x.Tags
	}
	return nil
}

type Step_MergeBuild struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If set, then this stream is expected to be a datagram stream
	// containing Build messages.
	//
	// This should be the stream name relative to the current build's
	// $LOGDOG_NAMESPACE.
	FromLogdogStream string `protobuf:"bytes,1,opt,name=from_logdog_stream,json=fromLogdogStream,proto3" json:"from_logdog_stream,omitempty"`
	// If set, then this stream will be merged "in line" with this step.
	//
	// Properties emitted by the merge build stream will overwrite global
	// outputs with the same top-level key.
	//
	// Steps emitted by the merge build stream will NOT have their names
	// namespaced (though the log stream names are still expected to
	// adhere to the regular luciexe rules).
	//
	// Because this is a legacy feature, this intentionally omits other fields
	// which "could be" merged, because there was no affordance to emit them
	// under the legacy annotator scheme:
	//   - output.gitiles_commit will not be merged.
	//   - output.logs will not be merged.
	//   - summary_markdown will not be merged.
	//
	// This is NOT a recommended mode of operation, but legacy ChromeOS
	// builders rely on this behavior.
	//
	// See crbug.com/1310155.
	LegacyGlobalNamespace bool `protobuf:"varint,2,opt,name=legacy_global_namespace,json=legacyGlobalNamespace,proto3" json:"legacy_global_namespace,omitempty"`
}

func (x *Step_MergeBuild) Reset() {
	*x = Step_MergeBuild{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_buildbucket_proto_step_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Step_MergeBuild) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Step_MergeBuild) ProtoMessage() {}

func (x *Step_MergeBuild) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_buildbucket_proto_step_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Step_MergeBuild.ProtoReflect.Descriptor instead.
func (*Step_MergeBuild) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_buildbucket_proto_step_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Step_MergeBuild) GetFromLogdogStream() string {
	if x != nil {
		return x.FromLogdogStream
	}
	return ""
}

func (x *Step_MergeBuild) GetLegacyGlobalNamespace() bool {
	if x != nil {
		return x.LegacyGlobalNamespace
	}
	return false
}

var File_go_chromium_org_luci_buildbucket_proto_step_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_buildbucket_proto_step_proto_rawDesc = []byte{
	0x0a, 0x31, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x65, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x2e, 0x76, 0x32, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x33, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75,
	0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf6, 0x03, 0x0a, 0x04, 0x53, 0x74,
	0x65, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x27, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x04, 0x6c, 0x6f, 0x67,
	0x73, 0x12, 0x40, 0x0a, 0x0b, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x74, 0x65, 0x70, 0x2e, 0x4d, 0x65, 0x72,
	0x67, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x0a, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x42, 0x75,
	0x69, 0x6c, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x6d,
	0x61, 0x72, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x4d, 0x61, 0x72, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x12, 0x2e,
	0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x69, 0x72, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x1a, 0x72,
	0x0a, 0x0a, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x12, 0x2c, 0x0a, 0x12,
	0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x6c, 0x6f, 0x67, 0x64, 0x6f, 0x67, 0x5f, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x66, 0x72, 0x6f, 0x6d, 0x4c, 0x6f,
	0x67, 0x64, 0x6f, 0x67, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x36, 0x0a, 0x17, 0x6c, 0x65,
	0x67, 0x61, 0x63, 0x79, 0x5f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x6c, 0x65, 0x67,
	0x61, 0x63, 0x79, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75,
	0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_buildbucket_proto_step_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_buildbucket_proto_step_proto_rawDescData = file_go_chromium_org_luci_buildbucket_proto_step_proto_rawDesc
)

func file_go_chromium_org_luci_buildbucket_proto_step_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_buildbucket_proto_step_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_buildbucket_proto_step_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_buildbucket_proto_step_proto_rawDescData)
	})
	return file_go_chromium_org_luci_buildbucket_proto_step_proto_rawDescData
}

var file_go_chromium_org_luci_buildbucket_proto_step_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_go_chromium_org_luci_buildbucket_proto_step_proto_goTypes = []interface{}{
	(*Step)(nil),                  // 0: buildbucket.v2.Step
	(*Step_MergeBuild)(nil),       // 1: buildbucket.v2.Step.MergeBuild
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
	(Status)(0),                   // 3: buildbucket.v2.Status
	(*Log)(nil),                   // 4: buildbucket.v2.Log
	(*StringPair)(nil),            // 5: buildbucket.v2.StringPair
}
var file_go_chromium_org_luci_buildbucket_proto_step_proto_depIdxs = []int32{
	2, // 0: buildbucket.v2.Step.start_time:type_name -> google.protobuf.Timestamp
	2, // 1: buildbucket.v2.Step.end_time:type_name -> google.protobuf.Timestamp
	3, // 2: buildbucket.v2.Step.status:type_name -> buildbucket.v2.Status
	4, // 3: buildbucket.v2.Step.logs:type_name -> buildbucket.v2.Log
	1, // 4: buildbucket.v2.Step.merge_build:type_name -> buildbucket.v2.Step.MergeBuild
	5, // 5: buildbucket.v2.Step.tags:type_name -> buildbucket.v2.StringPair
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_buildbucket_proto_step_proto_init() }
func file_go_chromium_org_luci_buildbucket_proto_step_proto_init() {
	if File_go_chromium_org_luci_buildbucket_proto_step_proto != nil {
		return
	}
	file_go_chromium_org_luci_buildbucket_proto_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_buildbucket_proto_step_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Step); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_go_chromium_org_luci_buildbucket_proto_step_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Step_MergeBuild); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_go_chromium_org_luci_buildbucket_proto_step_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_buildbucket_proto_step_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_buildbucket_proto_step_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_buildbucket_proto_step_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_buildbucket_proto_step_proto = out.File
	file_go_chromium_org_luci_buildbucket_proto_step_proto_rawDesc = nil
	file_go_chromium_org_luci_buildbucket_proto_step_proto_goTypes = nil
	file_go_chromium_org_luci_buildbucket_proto_step_proto_depIdxs = nil
}
