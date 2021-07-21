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
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.0
// source: go.chromium.org/luci/resultdb/sink/proto/v1/test_result.proto

package sinkpb

import (
	v1 "go.chromium.org/luci/resultdb/proto/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
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

// A result file format.
type TestResultFile_Format int32

const (
	// The file is a sequence of TestResult JSON objects (not a JSON Array).
	// The default format.
	TestResultFile_LUCI TestResultFile_Format = 0
	// Chromium's JSON Test Results format
	// https://chromium.googlesource.com/chromium/src/+/master/docs/testing/json_test_results_format.md
	TestResultFile_CHROMIUM_JSON_TEST_RESULTS TestResultFile_Format = 1
	// GTest format.
	// Not well documented.
	// Implementation:
	// https://cs.chromium.org/chromium/src/base/test/launcher/test_results_tracker.cc
	TestResultFile_GOOGLE_TEST TestResultFile_Format = 2
)

// Enum value maps for TestResultFile_Format.
var (
	TestResultFile_Format_name = map[int32]string{
		0: "LUCI",
		1: "CHROMIUM_JSON_TEST_RESULTS",
		2: "GOOGLE_TEST",
	}
	TestResultFile_Format_value = map[string]int32{
		"LUCI":                       0,
		"CHROMIUM_JSON_TEST_RESULTS": 1,
		"GOOGLE_TEST":                2,
	}
)

func (x TestResultFile_Format) Enum() *TestResultFile_Format {
	p := new(TestResultFile_Format)
	*p = x
	return p
}

func (x TestResultFile_Format) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TestResultFile_Format) Descriptor() protoreflect.EnumDescriptor {
	return file_go_chromium_org_luci_resultdb_sink_proto_v1_test_result_proto_enumTypes[0].Descriptor()
}

func (TestResultFile_Format) Type() protoreflect.EnumType {
	return &file_go_chromium_org_luci_resultdb_sink_proto_v1_test_result_proto_enumTypes[0]
}

func (x TestResultFile_Format) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TestResultFile_Format.Descriptor instead.
func (TestResultFile_Format) EnumDescriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_resultdb_sink_proto_v1_test_result_proto_rawDescGZIP(), []int{2, 0}
}

// A local equivalent of luci.resultdb.TestResult message
// in ../../v1/test_result.proto.
// See its comments for details.
type TestResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Equivalent of luci.resultdb.v1.TestResult.TestId.
	TestId string `protobuf:"bytes,1,opt,name=test_id,json=testId,proto3" json:"test_id,omitempty"`
	// Equivalent of luci.resultdb.v1.TestResult.result_id.
	//
	// If omitted, a random, unique ID is generated..
	ResultId string `protobuf:"bytes,2,opt,name=result_id,json=resultId,proto3" json:"result_id,omitempty"`
	// Equivalent of luci.resultdb.v1.TestResult.expected.
	Expected bool `protobuf:"varint,3,opt,name=expected,proto3" json:"expected,omitempty"`
	// Equivalent of luci.resultdb.v1.TestResult.status.
	Status v1.TestStatus `protobuf:"varint,4,opt,name=status,proto3,enum=luci.resultdb.v1.TestStatus" json:"status,omitempty"`
	// Equivalent of luci.resultdb.v1.TestResult.summary_html.
	SummaryHtml string `protobuf:"bytes,5,opt,name=summary_html,json=summaryHtml,proto3" json:"summary_html,omitempty"`
	// Equivalent of luci.resultdb.v1.TestResult.start_time.
	StartTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// Equivalent of luci.resultdb.v1.TestResult.duration.
	Duration *durationpb.Duration `protobuf:"bytes,7,opt,name=duration,proto3" json:"duration,omitempty"`
	// Equivalent of luci.resultdb.v1.TestResult.tags.
	Tags []*v1.StringPair `protobuf:"bytes,8,rep,name=tags,proto3" json:"tags,omitempty"`
	// Artifacts to upload and associate with this test result.
	// The map key is an artifact id.
	Artifacts map[string]*Artifact `protobuf:"bytes,9,rep,name=artifacts,proto3" json:"artifacts,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Equivalent of luci.resultdb.v1.TestResult.test_metadata.
	TestMetadata *v1.TestMetadata `protobuf:"bytes,11,opt,name=test_metadata,json=testMetadata,proto3" json:"test_metadata,omitempty"`
	// Equivalent of luci.resultdb.v1.TestResult.failure_reason.
	FailureReason *v1.FailureReason `protobuf:"bytes,12,opt,name=failure_reason,json=failureReason,proto3" json:"failure_reason,omitempty"`
}

func (x *TestResult) Reset() {
	*x = TestResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_resultdb_sink_proto_v1_test_result_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestResult) ProtoMessage() {}

func (x *TestResult) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_resultdb_sink_proto_v1_test_result_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestResult.ProtoReflect.Descriptor instead.
func (*TestResult) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_resultdb_sink_proto_v1_test_result_proto_rawDescGZIP(), []int{0}
}

func (x *TestResult) GetTestId() string {
	if x != nil {
		return x.TestId
	}
	return ""
}

func (x *TestResult) GetResultId() string {
	if x != nil {
		return x.ResultId
	}
	return ""
}

func (x *TestResult) GetExpected() bool {
	if x != nil {
		return x.Expected
	}
	return false
}

func (x *TestResult) GetStatus() v1.TestStatus {
	if x != nil {
		return x.Status
	}
	return v1.TestStatus_STATUS_UNSPECIFIED
}

func (x *TestResult) GetSummaryHtml() string {
	if x != nil {
		return x.SummaryHtml
	}
	return ""
}

func (x *TestResult) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *TestResult) GetDuration() *durationpb.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

func (x *TestResult) GetTags() []*v1.StringPair {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *TestResult) GetArtifacts() map[string]*Artifact {
	if x != nil {
		return x.Artifacts
	}
	return nil
}

func (x *TestResult) GetTestMetadata() *v1.TestMetadata {
	if x != nil {
		return x.TestMetadata
	}
	return nil
}

func (x *TestResult) GetFailureReason() *v1.FailureReason {
	if x != nil {
		return x.FailureReason
	}
	return nil
}

// A local equivalent of luci.resultdb.Artifact message
// in ../../rpc/v1/artifact.proto.
// See its comments for details.
// Does not have a name or artifact_id because they are represented by the
// TestResult.artifacts map key.
type Artifact struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Body:
	//	*Artifact_FilePath
	//	*Artifact_Contents
	Body isArtifact_Body `protobuf_oneof:"body"`
	// Equivalent of luci.resultdb.v1.Artifact.content_type.
	ContentType string `protobuf:"bytes,3,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
}

func (x *Artifact) Reset() {
	*x = Artifact{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_resultdb_sink_proto_v1_test_result_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Artifact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Artifact) ProtoMessage() {}

func (x *Artifact) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_resultdb_sink_proto_v1_test_result_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Artifact.ProtoReflect.Descriptor instead.
func (*Artifact) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_resultdb_sink_proto_v1_test_result_proto_rawDescGZIP(), []int{1}
}

func (m *Artifact) GetBody() isArtifact_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (x *Artifact) GetFilePath() string {
	if x, ok := x.GetBody().(*Artifact_FilePath); ok {
		return x.FilePath
	}
	return ""
}

func (x *Artifact) GetContents() []byte {
	if x, ok := x.GetBody().(*Artifact_Contents); ok {
		return x.Contents
	}
	return nil
}

func (x *Artifact) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

type isArtifact_Body interface {
	isArtifact_Body()
}

type Artifact_FilePath struct {
	// Absolute path to the artifact file on the same machine as the
	// ResultSink server.
	FilePath string `protobuf:"bytes,1,opt,name=file_path,json=filePath,proto3,oneof"`
}

type Artifact_Contents struct {
	// Contents of the artifact. Useful when sending a file from a different
	// machine.
	// TODO(nodir, sajjadm): allow sending contents in chunks.
	Contents []byte `protobuf:"bytes,2,opt,name=contents,proto3,oneof"`
}

func (*Artifact_FilePath) isArtifact_Body() {}

func (*Artifact_Contents) isArtifact_Body() {}

// A file with test results.
type TestResultFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Absolute OS-native path to the results file on the same machine as the
	// ResultSink server.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// Format of the file.
	Format TestResultFile_Format `protobuf:"varint,2,opt,name=format,proto3,enum=luci.resultsink.v1.TestResultFile_Format" json:"format,omitempty"`
}

func (x *TestResultFile) Reset() {
	*x = TestResultFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_resultdb_sink_proto_v1_test_result_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestResultFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestResultFile) ProtoMessage() {}

func (x *TestResultFile) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_resultdb_sink_proto_v1_test_result_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestResultFile.ProtoReflect.Descriptor instead.
func (*TestResultFile) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_resultdb_sink_proto_v1_test_result_proto_rawDescGZIP(), []int{2}
}

func (x *TestResultFile) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *TestResultFile) GetFormat() TestResultFile_Format {
	if x != nil {
		return x.Format
	}
	return TestResultFile_LUCI
}

var File_go_chromium_org_luci_resultdb_sink_proto_v1_test_result_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_resultdb_sink_proto_v1_test_result_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x64, 0x62, 0x2f,
	0x73, 0x69, 0x6e, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65,
	0x73, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x12, 0x6c, 0x75, 0x63, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x69, 0x6e, 0x6b,
	0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x33, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75,
	0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x64, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3b, 0x67, 0x6f, 0x2e, 0x63, 0x68,
	0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x64, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76,
	0x31, 0x2f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3a, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d,
	0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x64, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x74,
	0x65, 0x73, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x38, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e,
	0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x64,
	0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x05, 0x0a,
	0x0a, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74,
	0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65,
	0x73, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x34, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e,
	0x6c, 0x75, 0x63, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x64, 0x62, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x68,
	0x74, 0x6d, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x75, 0x6d, 0x6d, 0x61,
	0x72, 0x79, 0x48, 0x74, 0x6d, 0x6c, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x35, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73,
	0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6c, 0x75, 0x63, 0x69, 0x2e, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x50, 0x61, 0x69, 0x72, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x4b, 0x0a, 0x09, 0x61, 0x72,
	0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e,
	0x6c, 0x75, 0x63, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x69, 0x6e, 0x6b, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x41, 0x72,
	0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x61, 0x72,
	0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x12, 0x43, 0x0a, 0x0d, 0x74, 0x65, 0x73, 0x74, 0x5f,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x6c, 0x75, 0x63, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x64, 0x62, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x0c,
	0x74, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x46, 0x0a, 0x0e,
	0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6c, 0x75, 0x63, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x52,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x0d, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x52, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x1a, 0x5a, 0x0a, 0x0e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x32, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6c, 0x75, 0x63, 0x69, 0x2e, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x69, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72, 0x74,
	0x69, 0x66, 0x61, 0x63, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x4a, 0x04, 0x08, 0x0a, 0x10, 0x0b, 0x52, 0x0d, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x72, 0x0a, 0x08, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63,
	0x74, 0x12, 0x1d, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68,
	0x12, 0x1c, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x48, 0x00, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x21,
	0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x42, 0x06, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0xac, 0x01, 0x0a, 0x0e, 0x54, 0x65,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x12, 0x41, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x29, 0x2e, 0x6c, 0x75, 0x63, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x69,
	0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x46, 0x69, 0x6c, 0x65, 0x2e, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x52, 0x06, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x22, 0x43, 0x0a, 0x06, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x08, 0x0a,
	0x04, 0x4c, 0x55, 0x43, 0x49, 0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a, 0x43, 0x48, 0x52, 0x4f, 0x4d,
	0x49, 0x55, 0x4d, 0x5f, 0x4a, 0x53, 0x4f, 0x4e, 0x5f, 0x54, 0x45, 0x53, 0x54, 0x5f, 0x52, 0x45,
	0x53, 0x55, 0x4c, 0x54, 0x53, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x47, 0x4f, 0x4f, 0x47, 0x4c,
	0x45, 0x5f, 0x54, 0x45, 0x53, 0x54, 0x10, 0x02, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x6f, 0x2e, 0x63,
	0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69,
	0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x64, 0x62, 0x2f, 0x73, 0x69, 0x6e, 0x6b, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x69, 0x6e, 0x6b, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_resultdb_sink_proto_v1_test_result_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_resultdb_sink_proto_v1_test_result_proto_rawDescData = file_go_chromium_org_luci_resultdb_sink_proto_v1_test_result_proto_rawDesc
)

func file_go_chromium_org_luci_resultdb_sink_proto_v1_test_result_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_resultdb_sink_proto_v1_test_result_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_resultdb_sink_proto_v1_test_result_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_resultdb_sink_proto_v1_test_result_proto_rawDescData)
	})
	return file_go_chromium_org_luci_resultdb_sink_proto_v1_test_result_proto_rawDescData
}

var file_go_chromium_org_luci_resultdb_sink_proto_v1_test_result_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_go_chromium_org_luci_resultdb_sink_proto_v1_test_result_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_go_chromium_org_luci_resultdb_sink_proto_v1_test_result_proto_goTypes = []interface{}{
	(TestResultFile_Format)(0),    // 0: luci.resultsink.v1.TestResultFile.Format
	(*TestResult)(nil),            // 1: luci.resultsink.v1.TestResult
	(*Artifact)(nil),              // 2: luci.resultsink.v1.Artifact
	(*TestResultFile)(nil),        // 3: luci.resultsink.v1.TestResultFile
	nil,                           // 4: luci.resultsink.v1.TestResult.ArtifactsEntry
	(v1.TestStatus)(0),            // 5: luci.resultdb.v1.TestStatus
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),   // 7: google.protobuf.Duration
	(*v1.StringPair)(nil),         // 8: luci.resultdb.v1.StringPair
	(*v1.TestMetadata)(nil),       // 9: luci.resultdb.v1.TestMetadata
	(*v1.FailureReason)(nil),      // 10: luci.resultdb.v1.FailureReason
}
var file_go_chromium_org_luci_resultdb_sink_proto_v1_test_result_proto_depIdxs = []int32{
	5,  // 0: luci.resultsink.v1.TestResult.status:type_name -> luci.resultdb.v1.TestStatus
	6,  // 1: luci.resultsink.v1.TestResult.start_time:type_name -> google.protobuf.Timestamp
	7,  // 2: luci.resultsink.v1.TestResult.duration:type_name -> google.protobuf.Duration
	8,  // 3: luci.resultsink.v1.TestResult.tags:type_name -> luci.resultdb.v1.StringPair
	4,  // 4: luci.resultsink.v1.TestResult.artifacts:type_name -> luci.resultsink.v1.TestResult.ArtifactsEntry
	9,  // 5: luci.resultsink.v1.TestResult.test_metadata:type_name -> luci.resultdb.v1.TestMetadata
	10, // 6: luci.resultsink.v1.TestResult.failure_reason:type_name -> luci.resultdb.v1.FailureReason
	0,  // 7: luci.resultsink.v1.TestResultFile.format:type_name -> luci.resultsink.v1.TestResultFile.Format
	2,  // 8: luci.resultsink.v1.TestResult.ArtifactsEntry.value:type_name -> luci.resultsink.v1.Artifact
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_resultdb_sink_proto_v1_test_result_proto_init() }
func file_go_chromium_org_luci_resultdb_sink_proto_v1_test_result_proto_init() {
	if File_go_chromium_org_luci_resultdb_sink_proto_v1_test_result_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_resultdb_sink_proto_v1_test_result_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestResult); i {
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
		file_go_chromium_org_luci_resultdb_sink_proto_v1_test_result_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Artifact); i {
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
		file_go_chromium_org_luci_resultdb_sink_proto_v1_test_result_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestResultFile); i {
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
	file_go_chromium_org_luci_resultdb_sink_proto_v1_test_result_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Artifact_FilePath)(nil),
		(*Artifact_Contents)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_go_chromium_org_luci_resultdb_sink_proto_v1_test_result_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_resultdb_sink_proto_v1_test_result_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_resultdb_sink_proto_v1_test_result_proto_depIdxs,
		EnumInfos:         file_go_chromium_org_luci_resultdb_sink_proto_v1_test_result_proto_enumTypes,
		MessageInfos:      file_go_chromium_org_luci_resultdb_sink_proto_v1_test_result_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_resultdb_sink_proto_v1_test_result_proto = out.File
	file_go_chromium_org_luci_resultdb_sink_proto_v1_test_result_proto_rawDesc = nil
	file_go_chromium_org_luci_resultdb_sink_proto_v1_test_result_proto_goTypes = nil
	file_go_chromium_org_luci_resultdb_sink_proto_v1_test_result_proto_depIdxs = nil
}
