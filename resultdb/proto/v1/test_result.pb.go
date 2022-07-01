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
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: go.chromium.org/luci/resultdb/proto/v1/test_result.proto

package resultpb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Machine-readable status of a test result.
type TestStatus int32

const (
	// Status was not specified.
	// Not to be used in actual test results; serves as a default value for an
	// unset field.
	TestStatus_STATUS_UNSPECIFIED TestStatus = 0
	// The test case has passed.
	TestStatus_PASS TestStatus = 1
	// The test case has failed.
	// Suggests that the code under test is incorrect, but it is also possible
	// that the test is incorrect or it is a flake.
	TestStatus_FAIL TestStatus = 2
	// The test case has crashed during execution.
	// The outcome is inconclusive: the code under test might or might not be
	// correct, but the test+code is incorrect.
	TestStatus_CRASH TestStatus = 3
	// The test case has started, but was aborted before finishing.
	// A common reason: timeout.
	TestStatus_ABORT TestStatus = 4
	// The test case did not execute.
	// Examples:
	// - The execution of the collection of test cases, such as a test
	//   binary, was aborted prematurely and execution of some test cases was
	//   skipped.
	// - The test harness configuration specified that the test case MUST be
	//   skipped.
	TestStatus_SKIP TestStatus = 5
)

// Enum value maps for TestStatus.
var (
	TestStatus_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "PASS",
		2: "FAIL",
		3: "CRASH",
		4: "ABORT",
		5: "SKIP",
	}
	TestStatus_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"PASS":               1,
		"FAIL":               2,
		"CRASH":              3,
		"ABORT":              4,
		"SKIP":               5,
	}
)

func (x TestStatus) Enum() *TestStatus {
	p := new(TestStatus)
	*p = x
	return p
}

func (x TestStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TestStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_go_chromium_org_luci_resultdb_proto_v1_test_result_proto_enumTypes[0].Descriptor()
}

func (TestStatus) Type() protoreflect.EnumType {
	return &file_go_chromium_org_luci_resultdb_proto_v1_test_result_proto_enumTypes[0]
}

func (x TestStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TestStatus.Descriptor instead.
func (TestStatus) EnumDescriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_resultdb_proto_v1_test_result_proto_rawDescGZIP(), []int{0}
}

// Reason why a test variant was exonerated.
type ExonerationReason int32

const (
	// Reason was not specified.
	// Not to be used in actual test exonerations; serves as a default value for
	// an unset field.
	ExonerationReason_EXONERATION_REASON_UNSPECIFIED ExonerationReason = 0
	// Similar unexpected results were observed on a mainline branch
	// (i.e. against a build without unsubmitted changes applied).
	// (For avoidance of doubt, this includes both flakily and
	// deterministically occurring unexpected results.)
	// Applies to unexpected results in presubmit/CQ runs only.
	ExonerationReason_OCCURS_ON_MAINLINE ExonerationReason = 1
	// Similar unexpected results were observed in presubmit run(s) for other,
	// unrelated CL(s). (This is suggestive of the issue being present
	// on mainline but is not confirmed as there are possible confounding
	// factors, like how tests are run on CLs vs how tests are run on
	// mainline branches.)
	// Applies to unexpected results in presubmit/CQ runs only.
	ExonerationReason_OCCURS_ON_OTHER_CLS ExonerationReason = 2
	// The tests are not critical to the test subject (e.g. CL) passing.
	// This could be because more data is being collected to determine if
	// the tests are stable enough to be made critical (as is often the
	// case for experimental test suites).
	// If information exists indicating the tests are producing unexpected
	// results, and the tests are not critical for that reason,
	// prefer more specific reasons OCCURS_ON_MAINLINE or OCCURS_ON_OTHER_CLS.
	ExonerationReason_NOT_CRITICAL ExonerationReason = 3
	// The test result was an unexpected pass. (Note that such an exoneration is
	// not automatically created for unexpected passes, unless the option is
	// specified to ResultSink or the project manually creates one).
	ExonerationReason_UNEXPECTED_PASS ExonerationReason = 4
)

// Enum value maps for ExonerationReason.
var (
	ExonerationReason_name = map[int32]string{
		0: "EXONERATION_REASON_UNSPECIFIED",
		1: "OCCURS_ON_MAINLINE",
		2: "OCCURS_ON_OTHER_CLS",
		3: "NOT_CRITICAL",
		4: "UNEXPECTED_PASS",
	}
	ExonerationReason_value = map[string]int32{
		"EXONERATION_REASON_UNSPECIFIED": 0,
		"OCCURS_ON_MAINLINE":             1,
		"OCCURS_ON_OTHER_CLS":            2,
		"NOT_CRITICAL":                   3,
		"UNEXPECTED_PASS":                4,
	}
)

func (x ExonerationReason) Enum() *ExonerationReason {
	p := new(ExonerationReason)
	*p = x
	return p
}

func (x ExonerationReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExonerationReason) Descriptor() protoreflect.EnumDescriptor {
	return file_go_chromium_org_luci_resultdb_proto_v1_test_result_proto_enumTypes[1].Descriptor()
}

func (ExonerationReason) Type() protoreflect.EnumType {
	return &file_go_chromium_org_luci_resultdb_proto_v1_test_result_proto_enumTypes[1]
}

func (x ExonerationReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExonerationReason.Descriptor instead.
func (ExonerationReason) EnumDescriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_resultdb_proto_v1_test_result_proto_rawDescGZIP(), []int{1}
}

// A result of a functional test case.
// Often a single test case is executed multiple times and has multiple results,
// a single test suite has multiple test cases,
// and the same test suite can be executed in different variants
// (OS, GPU, compile flags, etc).
//
// This message does not specify the test id.
// It should be available in the message that embeds this message.
type TestResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Can be used to refer to this test result, e.g. in ResultDB.GetTestResult
	// RPC.
	// Format:
	// "invocations/{INVOCATION_ID}/tests/{URL_ESCAPED_TEST_ID}/results/{RESULT_ID}".
	// where URL_ESCAPED_TEST_ID is test_id escaped with
	// https://golang.org/pkg/net/url/#PathEscape See also https://aip.dev/122.
	//
	// Output only.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Test id, a unique identifier of the test in a LUCI project.
	// Regex: ^[[::print::]]{1,512}$
	//
	// If two tests have a common test id prefix that ends with a
	// non-alphanumeric character, they considered a part of a group. Examples:
	// - "a/b/c"
	// - "a/b/d"
	// - "a/b/e:x"
	// - "a/b/e:y"
	// - "a/f"
	// This defines the following groups:
	// - All items belong to one group because of the common prefix "a/"
	// - Within that group, the first 4 form a sub-group because of the common
	//   prefix "a/b/"
	// - Within that group, "a/b/e:x" and "a/b/e:y" form a sub-group because of
	//   the common prefix "a/b/e:".
	// This can be used in UI.
	// LUCI does not interpret test ids in any other way.
	TestId string `protobuf:"bytes,2,opt,name=test_id,json=testId,proto3" json:"test_id,omitempty"`
	// Identifies a test result in a given invocation and test id.
	// Regex: ^[a-z0-9\-_.]{1,32}$
	ResultId string `protobuf:"bytes,3,opt,name=result_id,json=resultId,proto3" json:"result_id,omitempty"`
	// Description of one specific way of running the test,
	// e.g. a specific bucket, builder and a test suite.
	Variant *Variant `protobuf:"bytes,4,opt,name=variant,proto3" json:"variant,omitempty"`
	// Whether the result of test case execution is expected.
	// In a typical Chromium CL, 99%+ of test results are expected.
	// Users are typically interested only in the unexpected results.
	//
	// An unexpected result != test case failure. There are test cases that are
	// expected to fail/skip/crash. The test harness compares the actual status
	// with the expected one(s) and this field is the result of the comparison.
	Expected bool `protobuf:"varint,5,opt,name=expected,proto3" json:"expected,omitempty"`
	// Machine-readable status of the test case.
	// MUST NOT be STATUS_UNSPECIFIED.
	Status TestStatus `protobuf:"varint,6,opt,name=status,proto3,enum=luci.resultdb.v1.TestStatus" json:"status,omitempty"`
	// Human-readable explanation of the result, in HTML.
	// MUST be sanitized before rendering in the browser.
	//
	// The size of the summary must be equal to or smaller than 4096 bytes in
	// UTF-8.
	//
	// Supports artifact embedding using custom tags:
	// * <text-artifact> renders contents of an artifact as text.
	//   Usage:
	//   * To embed result level artifact: <text-artifact
	//   artifact-id="<artifact_id>">
	//   * To embed invocation level artifact: <text-artifact
	//   artifact-id="<artifact_id>" inv-level>
	SummaryHtml string `protobuf:"bytes,7,opt,name=summary_html,json=summaryHtml,proto3" json:"summary_html,omitempty"`
	// The point in time when the test case started to execute.
	StartTime *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// Duration of the test case execution.
	// MUST be equal to or greater than 0.
	Duration *durationpb.Duration `protobuf:"bytes,9,opt,name=duration,proto3" json:"duration,omitempty"`
	// Metadata for this test result.
	// It might describe this particular execution or the test case.
	// A key can be repeated.
	Tags []*StringPair `protobuf:"bytes,10,rep,name=tags,proto3" json:"tags,omitempty"`
	// Hash of the variant.
	// hex(sha256(sorted(''.join('%s:%s\n' for k, v in variant.items())))).
	//
	// Output only.
	VariantHash string `protobuf:"bytes,12,opt,name=variant_hash,json=variantHash,proto3" json:"variant_hash,omitempty"`
	// Information about the test at the time of its execution.
	TestMetadata *TestMetadata `protobuf:"bytes,13,opt,name=test_metadata,json=testMetadata,proto3" json:"test_metadata,omitempty"`
	// Information about the test failure. Only present if the test failed.
	FailureReason *FailureReason `protobuf:"bytes,14,opt,name=failure_reason,json=failureReason,proto3" json:"failure_reason,omitempty"`
}

func (x *TestResult) Reset() {
	*x = TestResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_resultdb_proto_v1_test_result_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestResult) ProtoMessage() {}

func (x *TestResult) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_resultdb_proto_v1_test_result_proto_msgTypes[0]
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
	return file_go_chromium_org_luci_resultdb_proto_v1_test_result_proto_rawDescGZIP(), []int{0}
}

func (x *TestResult) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
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

func (x *TestResult) GetVariant() *Variant {
	if x != nil {
		return x.Variant
	}
	return nil
}

func (x *TestResult) GetExpected() bool {
	if x != nil {
		return x.Expected
	}
	return false
}

func (x *TestResult) GetStatus() TestStatus {
	if x != nil {
		return x.Status
	}
	return TestStatus_STATUS_UNSPECIFIED
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

func (x *TestResult) GetTags() []*StringPair {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *TestResult) GetVariantHash() string {
	if x != nil {
		return x.VariantHash
	}
	return ""
}

func (x *TestResult) GetTestMetadata() *TestMetadata {
	if x != nil {
		return x.TestMetadata
	}
	return nil
}

func (x *TestResult) GetFailureReason() *FailureReason {
	if x != nil {
		return x.FailureReason
	}
	return nil
}

// Indicates the test subject (e.g. a CL) is absolved from blame
// for an unexpected result of a test variant.
// For example, the test variant fails both with and without CL, so it is not
// CL's fault.
type TestExoneration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Can be used to refer to this test exoneration, e.g. in
	// ResultDB.GetTestExoneration RPC.
	// Format:
	// invocations/{INVOCATION_ID}/tests/{URL_ESCAPED_TEST_ID}/exonerations/{EXONERATION_ID}.
	// URL_ESCAPED_TEST_ID is test_variant.test_id escaped with
	// https://golang.org/pkg/net/url/#PathEscape See also https://aip.dev/122.
	//
	// Output only.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Test identifier, see TestResult.test_id.
	TestId string `protobuf:"bytes,2,opt,name=test_id,json=testId,proto3" json:"test_id,omitempty"`
	// Description of the variant of the test, see Variant type.
	// Unlike TestResult.extra_variant_pairs, this one must be a full definition
	// of the variant, i.e. it is not combined with Invocation.base_test_variant.
	Variant *Variant `protobuf:"bytes,3,opt,name=variant,proto3" json:"variant,omitempty"`
	// Identifies an exoneration in a given invocation and test id.
	// It is server-generated.
	ExonerationId string `protobuf:"bytes,4,opt,name=exoneration_id,json=exonerationId,proto3" json:"exoneration_id,omitempty"`
	// Reasoning behind the exoneration, in HTML.
	// MUST be sanitized before rendering in the browser.
	ExplanationHtml string `protobuf:"bytes,5,opt,name=explanation_html,json=explanationHtml,proto3" json:"explanation_html,omitempty"`
	// Hash of the variant.
	// hex(sha256(sorted(''.join('%s:%s\n' for k, v in variant.items())))).
	VariantHash string `protobuf:"bytes,6,opt,name=variant_hash,json=variantHash,proto3" json:"variant_hash,omitempty"`
	// Reasoning behind the exoneration, in machine-readable form.
	// Used to assist downstream analyses, such as automatic bug-filing.
	// This allow detection of e.g. critical tests failing in presubmit,
	// even if they are being exonerated because they fail on other CLs.
	Reason ExonerationReason `protobuf:"varint,7,opt,name=reason,proto3,enum=luci.resultdb.v1.ExonerationReason" json:"reason,omitempty"`
}

func (x *TestExoneration) Reset() {
	*x = TestExoneration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_resultdb_proto_v1_test_result_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestExoneration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestExoneration) ProtoMessage() {}

func (x *TestExoneration) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_resultdb_proto_v1_test_result_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestExoneration.ProtoReflect.Descriptor instead.
func (*TestExoneration) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_resultdb_proto_v1_test_result_proto_rawDescGZIP(), []int{1}
}

func (x *TestExoneration) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TestExoneration) GetTestId() string {
	if x != nil {
		return x.TestId
	}
	return ""
}

func (x *TestExoneration) GetVariant() *Variant {
	if x != nil {
		return x.Variant
	}
	return nil
}

func (x *TestExoneration) GetExonerationId() string {
	if x != nil {
		return x.ExonerationId
	}
	return ""
}

func (x *TestExoneration) GetExplanationHtml() string {
	if x != nil {
		return x.ExplanationHtml
	}
	return ""
}

func (x *TestExoneration) GetVariantHash() string {
	if x != nil {
		return x.VariantHash
	}
	return ""
}

func (x *TestExoneration) GetReason() ExonerationReason {
	if x != nil {
		return x.Reason
	}
	return ExonerationReason_EXONERATION_REASON_UNSPECIFIED
}

var File_go_chromium_org_luci_resultdb_proto_v1_test_result_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_resultdb_proto_v1_test_result_proto_rawDesc = []byte{
	0x0a, 0x38, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x64, 0x62, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x6c, 0x75, 0x63, 0x69,
	0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62,
	0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x33,
	0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x6c, 0x75, 0x63, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x64, 0x62, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x3a, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d,
	0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x64, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74,
	0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x3b, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x64, 0x62, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f,
	0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9a, 0x05, 0x0a,
	0x0a, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1a, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe0, 0x41, 0x03, 0xe0, 0x41,
	0x05, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x07, 0x74, 0x65, 0x73, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x06, 0x74,
	0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe0, 0x41, 0x05, 0xe0, 0x41, 0x02,
	0x52, 0x08, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x07, 0x76, 0x61,
	0x72, 0x69, 0x61, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6c, 0x75,
	0x63, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x56,
	0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x07, 0x76, 0x61, 0x72,
	0x69, 0x61, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x08, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x08, 0x65, 0x78, 0x70,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x39, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x6c, 0x75, 0x63, 0x69, 0x2e, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x26, 0x0a, 0x0c, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x68, 0x74, 0x6d, 0x6c,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x0b, 0x73, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x48, 0x74, 0x6d, 0x6c, 0x12, 0x3e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x09, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x0a, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6c, 0x75, 0x63, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x64, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x69, 0x72,
	0x42, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x29, 0x0a, 0x0c, 0x76,
	0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x06, 0xe0, 0x41, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x0b, 0x76, 0x61, 0x72, 0x69, 0x61,
	0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x43, 0x0a, 0x0d, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x6c, 0x75, 0x63, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x64, 0x62, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x0c, 0x74,
	0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x46, 0x0a, 0x0e, 0x66,
	0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6c, 0x75, 0x63, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x52, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x52, 0x0d, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x52, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x4a, 0x04, 0x08, 0x0b, 0x10, 0x0c, 0x22, 0xc4, 0x02, 0x0a, 0x0f, 0x54, 0x65,
	0x73, 0x74, 0x45, 0x78, 0x6f, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe0, 0x41, 0x03,
	0xe0, 0x41, 0x05, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x73,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x73, 0x74,
	0x49, 0x64, 0x12, 0x33, 0x0a, 0x07, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6c, 0x75, 0x63, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x52, 0x07,
	0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x12, 0x2d, 0x0a, 0x0e, 0x65, 0x78, 0x6f, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x06, 0xe0, 0x41, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x0d, 0x65, 0x78, 0x6f, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x10, 0x65, 0x78, 0x70, 0x6c, 0x61, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x68, 0x74, 0x6d, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x0f, 0x65, 0x78, 0x70, 0x6c, 0x61, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x48, 0x74, 0x6d, 0x6c, 0x12, 0x26, 0x0a, 0x0c, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e,
	0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x05, 0x52, 0x0b, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x40,
	0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23,
	0x2e, 0x6c, 0x75, 0x63, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x64, 0x62, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x78, 0x6f, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x2a, 0x58, 0x0a, 0x0a, 0x54, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16,
	0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x41, 0x53, 0x53, 0x10, 0x01,
	0x12, 0x08, 0x0a, 0x04, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x52,
	0x41, 0x53, 0x48, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x42, 0x4f, 0x52, 0x54, 0x10, 0x04,
	0x12, 0x08, 0x0a, 0x04, 0x53, 0x4b, 0x49, 0x50, 0x10, 0x05, 0x2a, 0x8f, 0x01, 0x0a, 0x11, 0x45,
	0x78, 0x6f, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x12, 0x22, 0x0a, 0x1e, 0x45, 0x58, 0x4f, 0x4e, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x4f, 0x43, 0x43, 0x55, 0x52, 0x53, 0x5f, 0x4f,
	0x4e, 0x5f, 0x4d, 0x41, 0x49, 0x4e, 0x4c, 0x49, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13,
	0x4f, 0x43, 0x43, 0x55, 0x52, 0x53, 0x5f, 0x4f, 0x4e, 0x5f, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x5f,
	0x43, 0x4c, 0x53, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x4e, 0x4f, 0x54, 0x5f, 0x43, 0x52, 0x49,
	0x54, 0x49, 0x43, 0x41, 0x4c, 0x10, 0x03, 0x12, 0x13, 0x0a, 0x0f, 0x55, 0x4e, 0x45, 0x58, 0x50,
	0x45, 0x43, 0x54, 0x45, 0x44, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x10, 0x04, 0x42, 0x31, 0x5a, 0x2f,
	0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x6c, 0x75, 0x63, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x64, 0x62, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x3b, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_resultdb_proto_v1_test_result_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_resultdb_proto_v1_test_result_proto_rawDescData = file_go_chromium_org_luci_resultdb_proto_v1_test_result_proto_rawDesc
)

func file_go_chromium_org_luci_resultdb_proto_v1_test_result_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_resultdb_proto_v1_test_result_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_resultdb_proto_v1_test_result_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_resultdb_proto_v1_test_result_proto_rawDescData)
	})
	return file_go_chromium_org_luci_resultdb_proto_v1_test_result_proto_rawDescData
}

var file_go_chromium_org_luci_resultdb_proto_v1_test_result_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_go_chromium_org_luci_resultdb_proto_v1_test_result_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_go_chromium_org_luci_resultdb_proto_v1_test_result_proto_goTypes = []interface{}{
	(TestStatus)(0),               // 0: luci.resultdb.v1.TestStatus
	(ExonerationReason)(0),        // 1: luci.resultdb.v1.ExonerationReason
	(*TestResult)(nil),            // 2: luci.resultdb.v1.TestResult
	(*TestExoneration)(nil),       // 3: luci.resultdb.v1.TestExoneration
	(*Variant)(nil),               // 4: luci.resultdb.v1.Variant
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),   // 6: google.protobuf.Duration
	(*StringPair)(nil),            // 7: luci.resultdb.v1.StringPair
	(*TestMetadata)(nil),          // 8: luci.resultdb.v1.TestMetadata
	(*FailureReason)(nil),         // 9: luci.resultdb.v1.FailureReason
}
var file_go_chromium_org_luci_resultdb_proto_v1_test_result_proto_depIdxs = []int32{
	4, // 0: luci.resultdb.v1.TestResult.variant:type_name -> luci.resultdb.v1.Variant
	0, // 1: luci.resultdb.v1.TestResult.status:type_name -> luci.resultdb.v1.TestStatus
	5, // 2: luci.resultdb.v1.TestResult.start_time:type_name -> google.protobuf.Timestamp
	6, // 3: luci.resultdb.v1.TestResult.duration:type_name -> google.protobuf.Duration
	7, // 4: luci.resultdb.v1.TestResult.tags:type_name -> luci.resultdb.v1.StringPair
	8, // 5: luci.resultdb.v1.TestResult.test_metadata:type_name -> luci.resultdb.v1.TestMetadata
	9, // 6: luci.resultdb.v1.TestResult.failure_reason:type_name -> luci.resultdb.v1.FailureReason
	4, // 7: luci.resultdb.v1.TestExoneration.variant:type_name -> luci.resultdb.v1.Variant
	1, // 8: luci.resultdb.v1.TestExoneration.reason:type_name -> luci.resultdb.v1.ExonerationReason
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_resultdb_proto_v1_test_result_proto_init() }
func file_go_chromium_org_luci_resultdb_proto_v1_test_result_proto_init() {
	if File_go_chromium_org_luci_resultdb_proto_v1_test_result_proto != nil {
		return
	}
	file_go_chromium_org_luci_resultdb_proto_v1_common_proto_init()
	file_go_chromium_org_luci_resultdb_proto_v1_test_metadata_proto_init()
	file_go_chromium_org_luci_resultdb_proto_v1_failure_reason_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_resultdb_proto_v1_test_result_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_go_chromium_org_luci_resultdb_proto_v1_test_result_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestExoneration); i {
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
			RawDescriptor: file_go_chromium_org_luci_resultdb_proto_v1_test_result_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_resultdb_proto_v1_test_result_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_resultdb_proto_v1_test_result_proto_depIdxs,
		EnumInfos:         file_go_chromium_org_luci_resultdb_proto_v1_test_result_proto_enumTypes,
		MessageInfos:      file_go_chromium_org_luci_resultdb_proto_v1_test_result_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_resultdb_proto_v1_test_result_proto = out.File
	file_go_chromium_org_luci_resultdb_proto_v1_test_result_proto_rawDesc = nil
	file_go_chromium_org_luci_resultdb_proto_v1_test_result_proto_goTypes = nil
	file_go_chromium_org_luci_resultdb_proto_v1_test_result_proto_depIdxs = nil
}
