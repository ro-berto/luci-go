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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.0
// source: go.chromium.org/luci/resultdb/proto/v1/invocation.proto

package resultpb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Invocation_State int32

const (
	// The default value. This value is used if the state is omitted.
	Invocation_STATE_UNSPECIFIED Invocation_State = 0
	// The invocation was created and accepts new results.
	Invocation_ACTIVE Invocation_State = 1
	// The invocation is in the process of transitioning into FINALIZED state.
	// This will happen automatically soon after all of its directly or
	// indirectly included invocations become inactive.
	Invocation_FINALIZING Invocation_State = 2
	// The invocation is immutable and no longer accepts new results nor
	// inclusions directly or indirectly.
	Invocation_FINALIZED Invocation_State = 3
)

// Enum value maps for Invocation_State.
var (
	Invocation_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "ACTIVE",
		2: "FINALIZING",
		3: "FINALIZED",
	}
	Invocation_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"ACTIVE":            1,
		"FINALIZING":        2,
		"FINALIZED":         3,
	}
)

func (x Invocation_State) Enum() *Invocation_State {
	p := new(Invocation_State)
	*p = x
	return p
}

func (x Invocation_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Invocation_State) Descriptor() protoreflect.EnumDescriptor {
	return file_go_chromium_org_luci_resultdb_proto_v1_invocation_proto_enumTypes[0].Descriptor()
}

func (Invocation_State) Type() protoreflect.EnumType {
	return &file_go_chromium_org_luci_resultdb_proto_v1_invocation_proto_enumTypes[0]
}

func (x Invocation_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Invocation_State.Descriptor instead.
func (Invocation_State) EnumDescriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_resultdb_proto_v1_invocation_proto_rawDescGZIP(), []int{0, 0}
}

// A conceptual container of results. Immutable once finalized.
// It represents all results of some computation; examples: swarming task,
// buildbucket build, CQ attempt.
// Composable: can include other invocations, see inclusion.proto.
//
// Next id: 14.
type Invocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Can be used to refer to this invocation, e.g. in ResultDB.GetInvocation
	// RPC.
	// Format: invocations/{INVOCATION_ID}
	// See also https://aip.dev/122.
	//
	// Output only.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Current state of the invocation.
	//
	// At creation time this can be set to FINALIZING e.g. if this invocation is
	// a simple wrapper of another and will itself not be modified.
	//
	// Otherwise this is an output only field.
	State Invocation_State `protobuf:"varint,2,opt,name=state,proto3,enum=luci.resultdb.v1.Invocation_State" json:"state,omitempty"`
	// When the invocation was created.
	// Output only.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Invocation-level string key-value pairs.
	// A key can be repeated.
	Tags []*StringPair `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`
	// When the invocation was finalized, i.e. transitioned to FINALIZED state.
	// If this field is set, implies that the invocation is finalized.
	//
	// Output only.
	FinalizeTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=finalize_time,json=finalizeTime,proto3" json:"finalize_time,omitempty"`
	// Timestamp when the invocation will be forcefully finalized.
	// Can be extended with UpdateInvocation until finalized.
	Deadline *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=deadline,proto3" json:"deadline,omitempty"`
	// Names of invocations included into this one. Overall results of this
	// invocation is a UNION of results directly included into this invocation
	// and results from the included invocations, recursively.
	// For example, a Buildbucket build invocation may include invocations of its
	// child swarming tasks and represent overall result of the build,
	// encapsulating the internal structure of the build.
	//
	// The graph is directed.
	// There can be at most one edge between a given pair of invocations.
	// The shape of the graph does not matter. What matters is only the set of
	// reachable invocations. Thus cycles are allowed and are noop.
	//
	// QueryTestResults returns test results from the transitive closure of
	// invocations.
	//
	// This field can be set under Recorder.CreateInvocationsRequest to include
	// existing invocations at the moment of invocation creation.
	// New invocations created in the same batch (via
	// Recorder.BatchCreateInvocationsRequest) are also allowed.
	// Otherwise, this field is to be treated as Output only.
	//
	// To modify included invocations, use Recorder.UpdateIncludedInvocations in
	// all other cases.
	IncludedInvocations []string `protobuf:"bytes,8,rep,name=included_invocations,json=includedInvocations,proto3" json:"included_invocations,omitempty"`
	// bigquery_exports indicates what BigQuery table(s) that results in this
	// invocation should export to.
	BigqueryExports []*BigQueryExport `protobuf:"bytes,9,rep,name=bigquery_exports,json=bigqueryExports,proto3" json:"bigquery_exports,omitempty"`
	// LUCI identity (e.g. "user:<email>") who created the invocation.
	// Typically, a LUCI service account (e.g.
	// "user:cr-buildbucket@appspot.gserviceaccount.com"), but can also be a user
	// (e.g. "user:johndoe@example.com").
	//
	// Output only.
	CreatedBy string `protobuf:"bytes,10,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	// Full name of the resource that produced results in this invocation.
	// See also https://aip.dev/122#full-resource-names
	// Typical examples:
	// - Swarming task: "//chromium-swarm.appspot.com/tasks/deadbeef"
	// - Buildbucket build: "//cr-buildbucket.appspot.com/builds/1234567890".
	ProducerResource string `protobuf:"bytes,11,opt,name=producer_resource,json=producerResource,proto3" json:"producer_resource,omitempty"`
	// Realm that the invocation exists under.
	// See https://chromium.googlesource.com/infra/luci/luci-py/+/refs/heads/master/appengine/auth_service/proto/realms_config.proto
	Realm string `protobuf:"bytes,12,opt,name=realm,proto3" json:"realm,omitempty"`
	// Specifies if/how to index the contents of this invocation.
	HistoryOptions *HistoryOptions `protobuf:"bytes,13,opt,name=history_options,json=historyOptions,proto3" json:"history_options,omitempty"`
}

func (x *Invocation) Reset() {
	*x = Invocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_resultdb_proto_v1_invocation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Invocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Invocation) ProtoMessage() {}

func (x *Invocation) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_resultdb_proto_v1_invocation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Invocation.ProtoReflect.Descriptor instead.
func (*Invocation) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_resultdb_proto_v1_invocation_proto_rawDescGZIP(), []int{0}
}

func (x *Invocation) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Invocation) GetState() Invocation_State {
	if x != nil {
		return x.State
	}
	return Invocation_STATE_UNSPECIFIED
}

func (x *Invocation) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Invocation) GetTags() []*StringPair {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Invocation) GetFinalizeTime() *timestamppb.Timestamp {
	if x != nil {
		return x.FinalizeTime
	}
	return nil
}

func (x *Invocation) GetDeadline() *timestamppb.Timestamp {
	if x != nil {
		return x.Deadline
	}
	return nil
}

func (x *Invocation) GetIncludedInvocations() []string {
	if x != nil {
		return x.IncludedInvocations
	}
	return nil
}

func (x *Invocation) GetBigqueryExports() []*BigQueryExport {
	if x != nil {
		return x.BigqueryExports
	}
	return nil
}

func (x *Invocation) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *Invocation) GetProducerResource() string {
	if x != nil {
		return x.ProducerResource
	}
	return ""
}

func (x *Invocation) GetRealm() string {
	if x != nil {
		return x.Realm
	}
	return ""
}

func (x *Invocation) GetHistoryOptions() *HistoryOptions {
	if x != nil {
		return x.HistoryOptions
	}
	return nil
}

// BigQueryExport indicates that results in this invocation should be exported
// to BigQuery after finalization.
type BigQueryExport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the BigQuery project.
	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	// Name of the BigQuery Dataset.
	Dataset string `protobuf:"bytes,2,opt,name=dataset,proto3" json:"dataset,omitempty"`
	// Name of the BigQuery Table.
	Table string `protobuf:"bytes,3,opt,name=table,proto3" json:"table,omitempty"`
	// Types that are assignable to ResultType:
	//	*BigQueryExport_TestResults_
	//	*BigQueryExport_TextArtifacts_
	ResultType isBigQueryExport_ResultType `protobuf_oneof:"result_type"`
}

func (x *BigQueryExport) Reset() {
	*x = BigQueryExport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_resultdb_proto_v1_invocation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BigQueryExport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BigQueryExport) ProtoMessage() {}

func (x *BigQueryExport) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_resultdb_proto_v1_invocation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BigQueryExport.ProtoReflect.Descriptor instead.
func (*BigQueryExport) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_resultdb_proto_v1_invocation_proto_rawDescGZIP(), []int{1}
}

func (x *BigQueryExport) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *BigQueryExport) GetDataset() string {
	if x != nil {
		return x.Dataset
	}
	return ""
}

func (x *BigQueryExport) GetTable() string {
	if x != nil {
		return x.Table
	}
	return ""
}

func (m *BigQueryExport) GetResultType() isBigQueryExport_ResultType {
	if m != nil {
		return m.ResultType
	}
	return nil
}

func (x *BigQueryExport) GetTestResults() *BigQueryExport_TestResults {
	if x, ok := x.GetResultType().(*BigQueryExport_TestResults_); ok {
		return x.TestResults
	}
	return nil
}

func (x *BigQueryExport) GetTextArtifacts() *BigQueryExport_TextArtifacts {
	if x, ok := x.GetResultType().(*BigQueryExport_TextArtifacts_); ok {
		return x.TextArtifacts
	}
	return nil
}

type isBigQueryExport_ResultType interface {
	isBigQueryExport_ResultType()
}

type BigQueryExport_TestResults_ struct {
	TestResults *BigQueryExport_TestResults `protobuf:"bytes,4,opt,name=test_results,json=testResults,proto3,oneof"`
}

type BigQueryExport_TextArtifacts_ struct {
	TextArtifacts *BigQueryExport_TextArtifacts `protobuf:"bytes,6,opt,name=text_artifacts,json=textArtifacts,proto3,oneof"`
}

func (*BigQueryExport_TestResults_) isBigQueryExport_ResultType() {}

func (*BigQueryExport_TextArtifacts_) isBigQueryExport_ResultType() {}

// HistoryOptions indicates how the invocations should be indexed, so that their
// results can be queried over a range of time or of commits.
type HistoryOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Set this to index the results by the containing invocation's create_time.
	UseInvocationTimestamp bool `protobuf:"varint,1,opt,name=use_invocation_timestamp,json=useInvocationTimestamp,proto3" json:"use_invocation_timestamp,omitempty"`
	// Set this to index by commit position.
	// It's up to the creator of the invocation to set this consistently over
	// time across the same test variant.
	Commit *CommitPosition `protobuf:"bytes,2,opt,name=commit,proto3" json:"commit,omitempty"`
}

func (x *HistoryOptions) Reset() {
	*x = HistoryOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_resultdb_proto_v1_invocation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistoryOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoryOptions) ProtoMessage() {}

func (x *HistoryOptions) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_resultdb_proto_v1_invocation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoryOptions.ProtoReflect.Descriptor instead.
func (*HistoryOptions) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_resultdb_proto_v1_invocation_proto_rawDescGZIP(), []int{2}
}

func (x *HistoryOptions) GetUseInvocationTimestamp() bool {
	if x != nil {
		return x.UseInvocationTimestamp
	}
	return false
}

func (x *HistoryOptions) GetCommit() *CommitPosition {
	if x != nil {
		return x.Commit
	}
	return nil
}

// TestResults indicates that test results should be exported.
type BigQueryExport_TestResults struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Use predicate to query test results that should be exported to
	// BigQuery table.
	Predicate *TestResultPredicate `protobuf:"bytes,1,opt,name=predicate,proto3" json:"predicate,omitempty"`
}

func (x *BigQueryExport_TestResults) Reset() {
	*x = BigQueryExport_TestResults{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_resultdb_proto_v1_invocation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BigQueryExport_TestResults) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BigQueryExport_TestResults) ProtoMessage() {}

func (x *BigQueryExport_TestResults) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_resultdb_proto_v1_invocation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BigQueryExport_TestResults.ProtoReflect.Descriptor instead.
func (*BigQueryExport_TestResults) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_resultdb_proto_v1_invocation_proto_rawDescGZIP(), []int{1, 0}
}

func (x *BigQueryExport_TestResults) GetPredicate() *TestResultPredicate {
	if x != nil {
		return x.Predicate
	}
	return nil
}

// TextArtifacts indicates that text artifacts should be exported.
type BigQueryExport_TextArtifacts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Use predicate to query artifacts that should be exported to
	// BigQuery table.
	//
	// Sub-field predicate.content_type_regexp defaults to "text/.*".
	Predicate *ArtifactPredicate `protobuf:"bytes,1,opt,name=predicate,proto3" json:"predicate,omitempty"`
}

func (x *BigQueryExport_TextArtifacts) Reset() {
	*x = BigQueryExport_TextArtifacts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_resultdb_proto_v1_invocation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BigQueryExport_TextArtifacts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BigQueryExport_TextArtifacts) ProtoMessage() {}

func (x *BigQueryExport_TextArtifacts) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_resultdb_proto_v1_invocation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BigQueryExport_TextArtifacts.ProtoReflect.Descriptor instead.
func (*BigQueryExport_TextArtifacts) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_resultdb_proto_v1_invocation_proto_rawDescGZIP(), []int{1, 1}
}

func (x *BigQueryExport_TextArtifacts) GetPredicate() *ArtifactPredicate {
	if x != nil {
		return x.Predicate
	}
	return nil
}

var File_go_chromium_org_luci_resultdb_proto_v1_invocation_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_resultdb_proto_v1_invocation_proto_rawDesc = []byte{
	0x0a, 0x37, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x64, 0x62, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x6c, 0x75, 0x63, 0x69, 0x2e,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65,
	0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x33, 0x67,
	0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c,
	0x75, 0x63, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x64, 0x62, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x36, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e,
	0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x64,
	0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x65, 0x64, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xda, 0x05, 0x0a, 0x0a, 0x49,
	0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe0, 0x41, 0x03, 0xe0, 0x41, 0x05, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x6c, 0x75, 0x63, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x43, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x42, 0x06, 0xe0, 0x41, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6c, 0x75, 0x63, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x64, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x69, 0x72,
	0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x44, 0x0a, 0x0d, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x69,
	0x7a, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0c,
	0x66, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x36, 0x0a, 0x08,
	0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x64, 0x65, 0x61, 0x64,
	0x6c, 0x69, 0x6e, 0x65, 0x12, 0x31, 0x0a, 0x14, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64,
	0x5f, 0x69, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x08, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x13, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x49, 0x6e, 0x76, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x4b, 0x0a, 0x10, 0x62, 0x69, 0x67, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x5f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x6c, 0x75, 0x63, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x64,
	0x62, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x45, 0x78, 0x70,
	0x6f, 0x72, 0x74, 0x52, 0x0f, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x45, 0x78, 0x70,
	0x6f, 0x72, 0x74, 0x73, 0x12, 0x22, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x62, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x2b, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x12, 0x49, 0x0a, 0x0f, 0x68,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6c, 0x75, 0x63, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0e, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x49, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45,
	0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x46, 0x49, 0x4e, 0x41, 0x4c, 0x49, 0x5a, 0x49, 0x4e, 0x47,
	0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x46, 0x49, 0x4e, 0x41, 0x4c, 0x49, 0x5a, 0x45, 0x44, 0x10,
	0x03, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x04, 0x22, 0xcc, 0x03, 0x0a, 0x0e, 0x42, 0x69, 0x67, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1d, 0x0a, 0x07, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1d, 0x0a, 0x07, 0x64, 0x61, 0x74,
	0x61, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52,
	0x07, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x12, 0x19, 0x0a, 0x05, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x05, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x12, 0x51, 0x0a, 0x0c, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6c, 0x75, 0x63, 0x69,
	0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x69, 0x67,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x54, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x48, 0x00, 0x52, 0x0b, 0x74, 0x65, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x57, 0x0a, 0x0e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x61,
	0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e,
	0x2e, 0x6c, 0x75, 0x63, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x64, 0x62, 0x2e, 0x76,
	0x31, 0x2e, 0x42, 0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x54, 0x65, 0x78, 0x74, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x48, 0x00,
	0x52, 0x0d, 0x74, 0x65, 0x78, 0x74, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x1a,
	0x52, 0x0a, 0x0b, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x43,
	0x0a, 0x09, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x25, 0x2e, 0x6c, 0x75, 0x63, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x64,
	0x62, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x50,
	0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x09, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x1a, 0x52, 0x0a, 0x0d, 0x54, 0x65, 0x78, 0x74, 0x41, 0x72, 0x74, 0x69, 0x66,
	0x61, 0x63, 0x74, 0x73, 0x12, 0x41, 0x0a, 0x09, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6c, 0x75, 0x63, 0x69, 0x2e, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66,
	0x61, 0x63, 0x74, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x09, 0x70, 0x72,
	0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x84, 0x01, 0x0a, 0x0e, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x38, 0x0a, 0x18, 0x75, 0x73, 0x65,
	0x5f, 0x69, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x75, 0x73, 0x65,
	0x49, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x12, 0x38, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6c, 0x75, 0x63, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x50, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x42, 0x31, 0x5a,
	0x2f, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x64, 0x62, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x3b, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_resultdb_proto_v1_invocation_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_resultdb_proto_v1_invocation_proto_rawDescData = file_go_chromium_org_luci_resultdb_proto_v1_invocation_proto_rawDesc
)

func file_go_chromium_org_luci_resultdb_proto_v1_invocation_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_resultdb_proto_v1_invocation_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_resultdb_proto_v1_invocation_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_resultdb_proto_v1_invocation_proto_rawDescData)
	})
	return file_go_chromium_org_luci_resultdb_proto_v1_invocation_proto_rawDescData
}

var file_go_chromium_org_luci_resultdb_proto_v1_invocation_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_go_chromium_org_luci_resultdb_proto_v1_invocation_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_go_chromium_org_luci_resultdb_proto_v1_invocation_proto_goTypes = []interface{}{
	(Invocation_State)(0),                // 0: luci.resultdb.v1.Invocation.State
	(*Invocation)(nil),                   // 1: luci.resultdb.v1.Invocation
	(*BigQueryExport)(nil),               // 2: luci.resultdb.v1.BigQueryExport
	(*HistoryOptions)(nil),               // 3: luci.resultdb.v1.HistoryOptions
	(*BigQueryExport_TestResults)(nil),   // 4: luci.resultdb.v1.BigQueryExport.TestResults
	(*BigQueryExport_TextArtifacts)(nil), // 5: luci.resultdb.v1.BigQueryExport.TextArtifacts
	(*timestamppb.Timestamp)(nil),        // 6: google.protobuf.Timestamp
	(*StringPair)(nil),                   // 7: luci.resultdb.v1.StringPair
	(*CommitPosition)(nil),               // 8: luci.resultdb.v1.CommitPosition
	(*TestResultPredicate)(nil),          // 9: luci.resultdb.v1.TestResultPredicate
	(*ArtifactPredicate)(nil),            // 10: luci.resultdb.v1.ArtifactPredicate
}
var file_go_chromium_org_luci_resultdb_proto_v1_invocation_proto_depIdxs = []int32{
	0,  // 0: luci.resultdb.v1.Invocation.state:type_name -> luci.resultdb.v1.Invocation.State
	6,  // 1: luci.resultdb.v1.Invocation.create_time:type_name -> google.protobuf.Timestamp
	7,  // 2: luci.resultdb.v1.Invocation.tags:type_name -> luci.resultdb.v1.StringPair
	6,  // 3: luci.resultdb.v1.Invocation.finalize_time:type_name -> google.protobuf.Timestamp
	6,  // 4: luci.resultdb.v1.Invocation.deadline:type_name -> google.protobuf.Timestamp
	2,  // 5: luci.resultdb.v1.Invocation.bigquery_exports:type_name -> luci.resultdb.v1.BigQueryExport
	3,  // 6: luci.resultdb.v1.Invocation.history_options:type_name -> luci.resultdb.v1.HistoryOptions
	4,  // 7: luci.resultdb.v1.BigQueryExport.test_results:type_name -> luci.resultdb.v1.BigQueryExport.TestResults
	5,  // 8: luci.resultdb.v1.BigQueryExport.text_artifacts:type_name -> luci.resultdb.v1.BigQueryExport.TextArtifacts
	8,  // 9: luci.resultdb.v1.HistoryOptions.commit:type_name -> luci.resultdb.v1.CommitPosition
	9,  // 10: luci.resultdb.v1.BigQueryExport.TestResults.predicate:type_name -> luci.resultdb.v1.TestResultPredicate
	10, // 11: luci.resultdb.v1.BigQueryExport.TextArtifacts.predicate:type_name -> luci.resultdb.v1.ArtifactPredicate
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_resultdb_proto_v1_invocation_proto_init() }
func file_go_chromium_org_luci_resultdb_proto_v1_invocation_proto_init() {
	if File_go_chromium_org_luci_resultdb_proto_v1_invocation_proto != nil {
		return
	}
	file_go_chromium_org_luci_resultdb_proto_v1_common_proto_init()
	file_go_chromium_org_luci_resultdb_proto_v1_predicate_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_resultdb_proto_v1_invocation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Invocation); i {
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
		file_go_chromium_org_luci_resultdb_proto_v1_invocation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BigQueryExport); i {
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
		file_go_chromium_org_luci_resultdb_proto_v1_invocation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HistoryOptions); i {
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
		file_go_chromium_org_luci_resultdb_proto_v1_invocation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BigQueryExport_TestResults); i {
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
		file_go_chromium_org_luci_resultdb_proto_v1_invocation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BigQueryExport_TextArtifacts); i {
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
	file_go_chromium_org_luci_resultdb_proto_v1_invocation_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*BigQueryExport_TestResults_)(nil),
		(*BigQueryExport_TextArtifacts_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_go_chromium_org_luci_resultdb_proto_v1_invocation_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_resultdb_proto_v1_invocation_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_resultdb_proto_v1_invocation_proto_depIdxs,
		EnumInfos:         file_go_chromium_org_luci_resultdb_proto_v1_invocation_proto_enumTypes,
		MessageInfos:      file_go_chromium_org_luci_resultdb_proto_v1_invocation_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_resultdb_proto_v1_invocation_proto = out.File
	file_go_chromium_org_luci_resultdb_proto_v1_invocation_proto_rawDesc = nil
	file_go_chromium_org_luci_resultdb_proto_v1_invocation_proto_goTypes = nil
	file_go_chromium_org_luci_resultdb_proto_v1_invocation_proto_depIdxs = nil
}
