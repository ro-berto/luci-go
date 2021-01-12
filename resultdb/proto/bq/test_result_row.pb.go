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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.1
// source: go.chromium.org/luci/resultdb/proto/bq/test_result_row.proto

package resultpb

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

// TestResultRow represents a row in a BigQuery table for result of a functional
// test case.
type TestResultRow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Exported contains info of the exported invocation.
	//
	// Note: it's possible that this invocation is not the result's
	// immediate parent invocation, but the including invocation.
	// For example if the BigQuery table is for all test results of Chromium CI
	// builds, then the exported invocation is for a CI build, which includes
	// multiple invocations for swarming tasks within that build.
	Exported *InvocationRecord `protobuf:"bytes,1,opt,name=exported,proto3" json:"exported,omitempty"`
	// Parent contains info of the result's immediate parent invocation.
	Parent *InvocationRecord `protobuf:"bytes,2,opt,name=parent,proto3" json:"parent,omitempty"`
	// Is a unique identifier of the test in a LUCI project.
	// Refer to TestResult.test_id for details.
	TestId string `protobuf:"bytes,3,opt,name=test_id,json=testId,proto3" json:"test_id,omitempty"`
	// Identifies a test result in a given invocation and test id.
	ResultId string `protobuf:"bytes,4,opt,name=result_id,json=resultId,proto3" json:"result_id,omitempty"`
	// Describes one specific way of running the test,
	// e.g. a specific bucket, builder and a test suite.
	Variant []*v1.StringPair `protobuf:"bytes,5,rep,name=variant,proto3" json:"variant,omitempty"`
	// A hex-encoded sha256 of concatenated "<key>:<value>\n" variant pairs.
	VariantHash string `protobuf:"bytes,6,opt,name=variant_hash,json=variantHash,proto3" json:"variant_hash,omitempty"`
	// Expected is a flag indicating whether the result of test case execution is expected.
	// Refer to TestResult.Expected for details.
	Expected bool `protobuf:"varint,7,opt,name=expected,proto3" json:"expected,omitempty"`
	// Status of the test result.
	// See TestStatus for possible values.
	Status string `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
	// A human-readable explanation of the result, in HTML.
	SummaryHtml string `protobuf:"bytes,9,opt,name=summary_html,json=summaryHtml,proto3" json:"summary_html,omitempty"`
	// The point in time when the test case started to execute.
	StartTime *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// Duration of the test case execution in seconds.
	Duration *durationpb.Duration `protobuf:"bytes,11,opt,name=duration,proto3" json:"duration,omitempty"`
	// Tags contains metadata for this test result.
	// It might describe this particular execution or the test case.
	Tags []*v1.StringPair `protobuf:"bytes,12,rep,name=tags,proto3" json:"tags,omitempty"`
	// If the failures of the test variant are exonerated.
	// Note: the exoneration is at the test variant level, not result level.
	Exonerated bool `protobuf:"varint,13,opt,name=exonerated,proto3" json:"exonerated,omitempty"`
	// Partition_time is used to partition the table.
	// It is the time when exported invocation was created in Spanner.
	// https://cloud.google.com/bigquery/docs/creating-column-partitions#limitations
	// mentions "The partitioning column must be a top-level field."
	// So we keep this column here instead of adding the CreateTime to Invocation.
	PartitionTime *timestamppb.Timestamp `protobuf:"bytes,14,opt,name=partition_time,json=partitionTime,proto3" json:"partition_time,omitempty"`
	// The location of the test definition.
	// Deprecated. Use test_metadata instead.
	//
	// Deprecated: Do not use.
	TestLocation *v1.TestLocation `protobuf:"bytes,15,opt,name=test_location,json=testLocation,proto3" json:"test_location,omitempty"`
	// Metadata of the test case,
	// e.g. the original test name, test_location, monorail_component and team_email.
	TestMetadata *v1.TestMetadata `protobuf:"bytes,16,opt,name=test_metadata,json=testMetadata,proto3" json:"test_metadata,omitempty"`
}

func (x *TestResultRow) Reset() {
	*x = TestResultRow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_resultdb_proto_bq_test_result_row_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestResultRow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestResultRow) ProtoMessage() {}

func (x *TestResultRow) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_resultdb_proto_bq_test_result_row_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestResultRow.ProtoReflect.Descriptor instead.
func (*TestResultRow) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_resultdb_proto_bq_test_result_row_proto_rawDescGZIP(), []int{0}
}

func (x *TestResultRow) GetExported() *InvocationRecord {
	if x != nil {
		return x.Exported
	}
	return nil
}

func (x *TestResultRow) GetParent() *InvocationRecord {
	if x != nil {
		return x.Parent
	}
	return nil
}

func (x *TestResultRow) GetTestId() string {
	if x != nil {
		return x.TestId
	}
	return ""
}

func (x *TestResultRow) GetResultId() string {
	if x != nil {
		return x.ResultId
	}
	return ""
}

func (x *TestResultRow) GetVariant() []*v1.StringPair {
	if x != nil {
		return x.Variant
	}
	return nil
}

func (x *TestResultRow) GetVariantHash() string {
	if x != nil {
		return x.VariantHash
	}
	return ""
}

func (x *TestResultRow) GetExpected() bool {
	if x != nil {
		return x.Expected
	}
	return false
}

func (x *TestResultRow) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *TestResultRow) GetSummaryHtml() string {
	if x != nil {
		return x.SummaryHtml
	}
	return ""
}

func (x *TestResultRow) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *TestResultRow) GetDuration() *durationpb.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

func (x *TestResultRow) GetTags() []*v1.StringPair {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *TestResultRow) GetExonerated() bool {
	if x != nil {
		return x.Exonerated
	}
	return false
}

func (x *TestResultRow) GetPartitionTime() *timestamppb.Timestamp {
	if x != nil {
		return x.PartitionTime
	}
	return nil
}

// Deprecated: Do not use.
func (x *TestResultRow) GetTestLocation() *v1.TestLocation {
	if x != nil {
		return x.TestLocation
	}
	return nil
}

func (x *TestResultRow) GetTestMetadata() *v1.TestMetadata {
	if x != nil {
		return x.TestMetadata
	}
	return nil
}

var File_go_chromium_org_luci_resultdb_proto_bq_test_result_row_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_resultdb_proto_bq_test_result_row_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x64, 0x62, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x71, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x5f, 0x72, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x6c, 0x75, 0x63, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x64, 0x62, 0x2e, 0x62, 0x71,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x33, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x64, 0x62,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x71, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x33, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d,
	0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x64, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3a, 0x67, 0x6f, 0x2e,
	0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63,
	0x69, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x64, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x06, 0x0a, 0x0d, 0x54, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x6f, 0x77, 0x12, 0x3e, 0x0a, 0x08, 0x65, 0x78, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6c, 0x75,
	0x63, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x64, 0x62, 0x2e, 0x62, 0x71, 0x2e, 0x49,
	0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52,
	0x08, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x12, 0x3a, 0x0a, 0x06, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6c, 0x75, 0x63, 0x69,
	0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x64, 0x62, 0x2e, 0x62, 0x71, 0x2e, 0x49, 0x6e, 0x76,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x06, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x07, 0x76,
	0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6c,
	0x75, 0x63, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x69, 0x72, 0x52, 0x07, 0x76, 0x61, 0x72, 0x69,
	0x61, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x5f, 0x68,
	0x61, 0x73, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x76, 0x61, 0x72, 0x69, 0x61,
	0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x68, 0x74, 0x6d, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x48, 0x74, 0x6d, 0x6c, 0x12, 0x39, 0x0a,
	0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x30, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x6c, 0x75, 0x63, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x64, 0x62, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x69, 0x72, 0x52, 0x04, 0x74, 0x61, 0x67,
	0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x6f, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x65, 0x78, 0x6f, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x64, 0x12, 0x41, 0x0a, 0x0e, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x47, 0x0a, 0x0d, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6c, 0x75,
	0x63, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x65, 0x73, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x02, 0x18, 0x01, 0x52,
	0x0c, 0x74, 0x65, 0x73, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x43, 0x0a,
	0x0d, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x10,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6c, 0x75, 0x63, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x52, 0x0c, 0x74, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75,
	0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x64, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x71, 0x3b, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_resultdb_proto_bq_test_result_row_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_resultdb_proto_bq_test_result_row_proto_rawDescData = file_go_chromium_org_luci_resultdb_proto_bq_test_result_row_proto_rawDesc
)

func file_go_chromium_org_luci_resultdb_proto_bq_test_result_row_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_resultdb_proto_bq_test_result_row_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_resultdb_proto_bq_test_result_row_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_resultdb_proto_bq_test_result_row_proto_rawDescData)
	})
	return file_go_chromium_org_luci_resultdb_proto_bq_test_result_row_proto_rawDescData
}

var file_go_chromium_org_luci_resultdb_proto_bq_test_result_row_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_go_chromium_org_luci_resultdb_proto_bq_test_result_row_proto_goTypes = []interface{}{
	(*TestResultRow)(nil),         // 0: luci.resultdb.bq.TestResultRow
	(*InvocationRecord)(nil),      // 1: luci.resultdb.bq.InvocationRecord
	(*v1.StringPair)(nil),         // 2: luci.resultdb.v1.StringPair
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),   // 4: google.protobuf.Duration
	(*v1.TestLocation)(nil),       // 5: luci.resultdb.v1.TestLocation
	(*v1.TestMetadata)(nil),       // 6: luci.resultdb.v1.TestMetadata
}
var file_go_chromium_org_luci_resultdb_proto_bq_test_result_row_proto_depIdxs = []int32{
	1, // 0: luci.resultdb.bq.TestResultRow.exported:type_name -> luci.resultdb.bq.InvocationRecord
	1, // 1: luci.resultdb.bq.TestResultRow.parent:type_name -> luci.resultdb.bq.InvocationRecord
	2, // 2: luci.resultdb.bq.TestResultRow.variant:type_name -> luci.resultdb.v1.StringPair
	3, // 3: luci.resultdb.bq.TestResultRow.start_time:type_name -> google.protobuf.Timestamp
	4, // 4: luci.resultdb.bq.TestResultRow.duration:type_name -> google.protobuf.Duration
	2, // 5: luci.resultdb.bq.TestResultRow.tags:type_name -> luci.resultdb.v1.StringPair
	3, // 6: luci.resultdb.bq.TestResultRow.partition_time:type_name -> google.protobuf.Timestamp
	5, // 7: luci.resultdb.bq.TestResultRow.test_location:type_name -> luci.resultdb.v1.TestLocation
	6, // 8: luci.resultdb.bq.TestResultRow.test_metadata:type_name -> luci.resultdb.v1.TestMetadata
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_resultdb_proto_bq_test_result_row_proto_init() }
func file_go_chromium_org_luci_resultdb_proto_bq_test_result_row_proto_init() {
	if File_go_chromium_org_luci_resultdb_proto_bq_test_result_row_proto != nil {
		return
	}
	file_go_chromium_org_luci_resultdb_proto_bq_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_resultdb_proto_bq_test_result_row_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestResultRow); i {
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
			RawDescriptor: file_go_chromium_org_luci_resultdb_proto_bq_test_result_row_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_resultdb_proto_bq_test_result_row_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_resultdb_proto_bq_test_result_row_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_resultdb_proto_bq_test_result_row_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_resultdb_proto_bq_test_result_row_proto = out.File
	file_go_chromium_org_luci_resultdb_proto_bq_test_result_row_proto_rawDesc = nil
	file_go_chromium_org_luci_resultdb_proto_bq_test_result_row_proto_goTypes = nil
	file_go_chromium_org_luci_resultdb_proto_bq_test_result_row_proto_depIdxs = nil
}
