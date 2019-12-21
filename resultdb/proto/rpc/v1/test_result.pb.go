// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/resultdb/proto/rpc/v1/test_result.proto

package rpcpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_type "go.chromium.org/luci/resultdb/proto/type"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

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

var TestStatus_name = map[int32]string{
	0: "STATUS_UNSPECIFIED",
	1: "PASS",
	2: "FAIL",
	3: "CRASH",
	4: "ABORT",
	5: "SKIP",
}

var TestStatus_value = map[string]int32{
	"STATUS_UNSPECIFIED": 0,
	"PASS":               1,
	"FAIL":               2,
	"CRASH":              3,
	"ABORT":              4,
	"SKIP":               5,
}

func (x TestStatus) String() string {
	return proto.EnumName(TestStatus_name, int32(x))
}

func (TestStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8253a2b79929a2cf, []int{0}
}

// A result of a functional test case.
// Often a single test case is executed multiple times and has multiple results,
// a single test suite has multiple test cases,
// and the same test suite can be executed in different variants
// (OS, GPU, compile flags, etc).
//
// This message does not specify the test path.
// It should be available in the message that embeds this message.
type TestResult struct {
	// Can be used to refer to this test result, e.g. in ResultDB.GetTestResult
	// RPC.
	// Format:
	// "invocations/{INVOCATION_ID}/tests/{URL_ESCAPED_TEST_PATH}/results/{RESULT_ID}".
	// URL_ESCAPED_TEST_PATH is test_path escaped with
	// https://golang.org/pkg/net/url/#PathEscape See also https://aip.dev/122.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Test path, a unique identifier of the test in a LUCI project.
	// Regex: ^[[::print::]]+$.
	//
	// If two tests have a common test path prefix that ends with a
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
	// LUCI does not interpret test paths in any other way.
	TestPath string `protobuf:"bytes,2,opt,name=test_path,json=testPath,proto3" json:"test_path,omitempty"`
	// Identifies a test result in a given invocation and test path.
	// Regex: ^[[:ascii:]]{1,32}$.
	ResultId string `protobuf:"bytes,3,opt,name=result_id,json=resultId,proto3" json:"result_id,omitempty"`
	// Description of one specific way of running the test,
	// e.g. a specific bucket, builder and a test suite.
	Variant *_type.Variant `protobuf:"bytes,4,opt,name=variant,proto3" json:"variant,omitempty"`
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
	Status TestStatus `protobuf:"varint,6,opt,name=status,proto3,enum=luci.resultdb.rpc.v1.TestStatus" json:"status,omitempty"`
	// Human-readable explanation of the result, in HTML.
	// MUST be sanitized before rendering in the browser.
	SummaryHtml string `protobuf:"bytes,7,opt,name=summary_html,json=summaryHtml,proto3" json:"summary_html,omitempty"`
	// The point in time when the test case started to execute.
	StartTime *timestamp.Timestamp `protobuf:"bytes,8,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// Duration of the test case execution.
	Duration *duration.Duration `protobuf:"bytes,9,opt,name=duration,proto3" json:"duration,omitempty"`
	// Metadata for this test result.
	// It might describe this particular execution or the test case.
	Tags []*_type.StringPair `protobuf:"bytes,10,rep,name=tags,proto3" json:"tags,omitempty"`
	// Artifacts consumed by this test result.
	//
	// Example: building a Chrome OS image is expensive and non-deterministic, so
	// they are retained and used as input artifact to a test case.
	InputArtifacts []*Artifact `protobuf:"bytes,11,rep,name=input_artifacts,json=inputArtifacts,proto3" json:"input_artifacts,omitempty"`
	// Artifacts produced by this test result.
	// Examples: traces, logs, screenshots, memory dumps, profiler output.
	OutputArtifacts      []*Artifact `protobuf:"bytes,12,rep,name=output_artifacts,json=outputArtifacts,proto3" json:"output_artifacts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *TestResult) Reset()         { *m = TestResult{} }
func (m *TestResult) String() string { return proto.CompactTextString(m) }
func (*TestResult) ProtoMessage()    {}
func (*TestResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_8253a2b79929a2cf, []int{0}
}

func (m *TestResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestResult.Unmarshal(m, b)
}
func (m *TestResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestResult.Marshal(b, m, deterministic)
}
func (m *TestResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestResult.Merge(m, src)
}
func (m *TestResult) XXX_Size() int {
	return xxx_messageInfo_TestResult.Size(m)
}
func (m *TestResult) XXX_DiscardUnknown() {
	xxx_messageInfo_TestResult.DiscardUnknown(m)
}

var xxx_messageInfo_TestResult proto.InternalMessageInfo

func (m *TestResult) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TestResult) GetTestPath() string {
	if m != nil {
		return m.TestPath
	}
	return ""
}

func (m *TestResult) GetResultId() string {
	if m != nil {
		return m.ResultId
	}
	return ""
}

func (m *TestResult) GetVariant() *_type.Variant {
	if m != nil {
		return m.Variant
	}
	return nil
}

func (m *TestResult) GetExpected() bool {
	if m != nil {
		return m.Expected
	}
	return false
}

func (m *TestResult) GetStatus() TestStatus {
	if m != nil {
		return m.Status
	}
	return TestStatus_STATUS_UNSPECIFIED
}

func (m *TestResult) GetSummaryHtml() string {
	if m != nil {
		return m.SummaryHtml
	}
	return ""
}

func (m *TestResult) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *TestResult) GetDuration() *duration.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func (m *TestResult) GetTags() []*_type.StringPair {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *TestResult) GetInputArtifacts() []*Artifact {
	if m != nil {
		return m.InputArtifacts
	}
	return nil
}

func (m *TestResult) GetOutputArtifacts() []*Artifact {
	if m != nil {
		return m.OutputArtifacts
	}
	return nil
}

// A file produced/consumed by a test case.
// See TestResult.output_artifacts for examples.
type Artifact struct {
	// A slash-separated relative path, identifies the artifact.
	// Example: "traces/a.txt".
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Machine-readable URL to fetch the contents of the artifact.
	// Valid schemes: "isolate", "gs", "logdog", "rbe-cas".
	FetchUrl string `protobuf:"bytes,2,opt,name=fetch_url,json=fetchUrl,proto3" json:"fetch_url,omitempty"`
	// Human-consumable URL to the file content.
	// Typically a URL of a page where the user can view/download the arficact.
	ViewUrl string `protobuf:"bytes,3,opt,name=view_url,json=viewUrl,proto3" json:"view_url,omitempty"`
	// Media type of the artifact.
	// Logs are typically "plain/text" and screenshots are typically "image/png".
	ContentType string `protobuf:"bytes,4,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	// Size of the file, in bytes.
	// Can be used in UI to decide whether to fetch an artifact and display it
	// inline, or only show a link if it is too large.
	Size int64 `protobuf:"varint,5,opt,name=size,proto3" json:"size,omitempty"`
	// Contents of the artifact if it is stored inline with the test result.
	// Empty for artifacts stored elsewhere. To fetch such artifacts, use
	// fetch_url.
	// Size MUST be <= 8KB.
	Contents             []byte   `protobuf:"bytes,6,opt,name=contents,proto3" json:"contents,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Artifact) Reset()         { *m = Artifact{} }
func (m *Artifact) String() string { return proto.CompactTextString(m) }
func (*Artifact) ProtoMessage()    {}
func (*Artifact) Descriptor() ([]byte, []int) {
	return fileDescriptor_8253a2b79929a2cf, []int{1}
}

func (m *Artifact) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Artifact.Unmarshal(m, b)
}
func (m *Artifact) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Artifact.Marshal(b, m, deterministic)
}
func (m *Artifact) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Artifact.Merge(m, src)
}
func (m *Artifact) XXX_Size() int {
	return xxx_messageInfo_Artifact.Size(m)
}
func (m *Artifact) XXX_DiscardUnknown() {
	xxx_messageInfo_Artifact.DiscardUnknown(m)
}

var xxx_messageInfo_Artifact proto.InternalMessageInfo

func (m *Artifact) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Artifact) GetFetchUrl() string {
	if m != nil {
		return m.FetchUrl
	}
	return ""
}

func (m *Artifact) GetViewUrl() string {
	if m != nil {
		return m.ViewUrl
	}
	return ""
}

func (m *Artifact) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *Artifact) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *Artifact) GetContents() []byte {
	if m != nil {
		return m.Contents
	}
	return nil
}

// Indicates the test subject (e.g. a CL) is absolved from blame
// for an unexpected result of a test variant.
// For example, the test variant fails both with and without CL, so it is not
// CL's fault.
type TestExoneration struct {
	// Can be used to refer to this test exoneration, e.g. in
	// ResultDB.GetTestExoneration RPC.
	// Format:
	// invocations/{INVOCATION_ID}/tests/{URL_ESCAPED_TEST_PATH}/exonerations/{EXONERATION_ID}.
	// URL_ESCAPED_TEST_PATH is test_variant.test_path escaped with
	// https://golang.org/pkg/net/url/#PathEscape See also https://aip.dev/122.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Test identifier, see TestResult.test_path.
	TestPath string `protobuf:"bytes,2,opt,name=test_path,json=testPath,proto3" json:"test_path,omitempty"`
	// Description of the variant of the test, see Variant type.
	// Unlike TestResult.extra_variant_pairs, this one must be a full definition
	// of the variant, i.e. it is not combined with Invocation.base_test_variant.
	Variant *_type.Variant `protobuf:"bytes,3,opt,name=variant,proto3" json:"variant,omitempty"`
	// Identifies an exoneration in a given invocation and test path.
	// It is server-generated.
	ExonerationId string `protobuf:"bytes,4,opt,name=exoneration_id,json=exonerationId,proto3" json:"exoneration_id,omitempty"`
	// Reasoning behind the exoneration, in HTML.
	// MUST be sanitized before rendering in the browser.
	ExplanationHtml      string   `protobuf:"bytes,5,opt,name=explanation_html,json=explanationHtml,proto3" json:"explanation_html,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestExoneration) Reset()         { *m = TestExoneration{} }
func (m *TestExoneration) String() string { return proto.CompactTextString(m) }
func (*TestExoneration) ProtoMessage()    {}
func (*TestExoneration) Descriptor() ([]byte, []int) {
	return fileDescriptor_8253a2b79929a2cf, []int{2}
}

func (m *TestExoneration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestExoneration.Unmarshal(m, b)
}
func (m *TestExoneration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestExoneration.Marshal(b, m, deterministic)
}
func (m *TestExoneration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestExoneration.Merge(m, src)
}
func (m *TestExoneration) XXX_Size() int {
	return xxx_messageInfo_TestExoneration.Size(m)
}
func (m *TestExoneration) XXX_DiscardUnknown() {
	xxx_messageInfo_TestExoneration.DiscardUnknown(m)
}

var xxx_messageInfo_TestExoneration proto.InternalMessageInfo

func (m *TestExoneration) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TestExoneration) GetTestPath() string {
	if m != nil {
		return m.TestPath
	}
	return ""
}

func (m *TestExoneration) GetVariant() *_type.Variant {
	if m != nil {
		return m.Variant
	}
	return nil
}

func (m *TestExoneration) GetExonerationId() string {
	if m != nil {
		return m.ExonerationId
	}
	return ""
}

func (m *TestExoneration) GetExplanationHtml() string {
	if m != nil {
		return m.ExplanationHtml
	}
	return ""
}

func init() {
	proto.RegisterEnum("luci.resultdb.rpc.v1.TestStatus", TestStatus_name, TestStatus_value)
	proto.RegisterType((*TestResult)(nil), "luci.resultdb.rpc.v1.TestResult")
	proto.RegisterType((*Artifact)(nil), "luci.resultdb.rpc.v1.Artifact")
	proto.RegisterType((*TestExoneration)(nil), "luci.resultdb.rpc.v1.TestExoneration")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/resultdb/proto/rpc/v1/test_result.proto", fileDescriptor_8253a2b79929a2cf)
}

var fileDescriptor_8253a2b79929a2cf = []byte{
	// 691 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xe9, 0x6a, 0xdb, 0x4c,
	0x14, 0xfd, 0x14, 0x2f, 0x91, 0xc7, 0xfe, 0x62, 0x31, 0x94, 0x56, 0x75, 0x20, 0x31, 0x29, 0x04,
	0x53, 0xa8, 0x44, 0x52, 0x42, 0x21, 0xb4, 0x05, 0x65, 0x23, 0xa6, 0x21, 0x35, 0x96, 0x53, 0xfa,
	0x4f, 0x8c, 0xe5, 0xb1, 0x3d, 0xa0, 0x8d, 0xd1, 0x95, 0x9b, 0xf4, 0x01, 0xfa, 0x6a, 0x7d, 0x0d,
	0x3f, 0x4a, 0x99, 0x45, 0x75, 0xdc, 0x84, 0x36, 0xff, 0xac, 0x7b, 0xce, 0xb9, 0x73, 0x97, 0x73,
	0x8d, 0xde, 0xcf, 0x52, 0x27, 0x9c, 0xf3, 0x34, 0x66, 0x45, 0xec, 0xa4, 0x7c, 0xe6, 0x46, 0x45,
	0xc8, 0x5c, 0x4e, 0xf3, 0x22, 0x82, 0xc9, 0xd8, 0xcd, 0x78, 0x0a, 0xa9, 0xcb, 0xb3, 0xd0, 0x5d,
	0x1c, 0xb8, 0x40, 0x73, 0x08, 0x14, 0xe4, 0x48, 0x00, 0x3f, 0x13, 0x6c, 0xa7, 0x64, 0x3b, 0x3c,
	0x0b, 0x9d, 0xc5, 0x41, 0x67, 0x77, 0x96, 0xa6, 0xb3, 0x88, 0xba, 0x24, 0x63, 0xee, 0x94, 0xd1,
	0x68, 0x12, 0x8c, 0xe9, 0x9c, 0x2c, 0x58, 0xca, 0x95, 0xac, 0xb3, 0xa3, 0x09, 0xf2, 0x6b, 0x5c,
	0x4c, 0xdd, 0x49, 0xc1, 0x09, 0xb0, 0x34, 0xd1, 0xf8, 0xee, 0x9f, 0x38, 0xb0, 0x98, 0xe6, 0x40,
	0xe2, 0x4c, 0x13, 0x8e, 0x9e, 0x52, 0x35, 0xdc, 0x65, 0xd4, 0x0d, 0xd3, 0x38, 0x2e, 0xf3, 0xee,
	0xfd, 0xa8, 0x21, 0x34, 0xa2, 0x39, 0x0c, 0x25, 0x11, 0x77, 0x50, 0x35, 0x21, 0x31, 0xb5, 0x8d,
	0xae, 0xd1, 0x6b, 0x9c, 0xd4, 0x97, 0x5e, 0x65, 0xe9, 0xd5, 0x86, 0x32, 0x86, 0xbb, 0xa8, 0x21,
	0xdb, 0xcd, 0x08, 0xcc, 0xed, 0x0d, 0x49, 0x90, 0xa8, 0x29, 0xa2, 0x03, 0x02, 0x73, 0xfc, 0x0a,
	0x35, 0xd4, 0x83, 0x01, 0x9b, 0xd8, 0x95, 0x32, 0x45, 0x6d, 0xe9, 0x6d, 0x0c, 0x4d, 0x05, 0xf4,
	0x27, 0xf8, 0x18, 0x6d, 0x2e, 0x08, 0x67, 0x24, 0x01, 0xbb, 0xda, 0x35, 0x7a, 0xcd, 0xc3, 0x6d,
	0x67, 0x7d, 0x64, 0xa2, 0x48, 0xe7, 0x8b, 0xa2, 0xa8, 0x17, 0x4a, 0x01, 0xde, 0x45, 0x26, 0xbd,
	0xcd, 0x68, 0x08, 0x74, 0x62, 0xd7, 0xba, 0x46, 0xcf, 0xd4, 0x15, 0x94, 0x41, 0xfc, 0x01, 0xd5,
	0x73, 0x20, 0x50, 0xe4, 0x76, 0xbd, 0x6b, 0xf4, 0xb6, 0x0e, 0xbb, 0xce, 0x63, 0xeb, 0x70, 0x44,
	0xc7, 0xbe, 0xe4, 0xa9, 0x04, 0x5a, 0x84, 0xf7, 0x51, 0x2b, 0x2f, 0xe2, 0x98, 0xf0, 0xbb, 0x60,
	0x0e, 0x71, 0x64, 0x6f, 0xae, 0xba, 0x6c, 0x6a, 0xe0, 0x12, 0xe2, 0x08, 0x7f, 0x44, 0x28, 0x07,
	0xc2, 0x21, 0x10, 0x5b, 0xb0, 0x4d, 0xd9, 0x46, 0xc7, 0x51, 0x2b, 0x72, 0xca, 0x15, 0x39, 0xa3,
	0x72, 0x45, 0x2a, 0x43, 0x43, 0x4a, 0x44, 0x10, 0x1f, 0x23, 0xb3, 0xdc, 0xaf, 0xdd, 0x90, 0xea,
	0x97, 0x0f, 0xd4, 0x67, 0x9a, 0xa0, 0x5b, 0x2c, 0xf9, 0xf8, 0x1d, 0xaa, 0x02, 0x99, 0xe5, 0x36,
	0xea, 0x56, 0x7a, 0xcd, 0xc3, 0x9d, 0xc7, 0x86, 0xe7, 0x03, 0x67, 0xc9, 0x6c, 0x40, 0x18, 0x57,
	0x62, 0x29, 0xc0, 0x57, 0xa8, 0xcd, 0x92, 0xac, 0x80, 0x80, 0x70, 0x60, 0x53, 0x12, 0x42, 0x6e,
	0x37, 0x1f, 0xcd, 0xa1, 0x87, 0xe4, 0x69, 0x9a, 0xca, 0xb1, 0x25, 0xb5, 0x65, 0x2c, 0xc7, 0xd7,
	0xc8, 0x4a, 0x0b, 0x58, 0x4f, 0xd7, 0x7a, 0x7a, 0xba, 0xb6, 0x12, 0xff, 0xce, 0xb7, 0xf7, 0xd3,
	0x40, 0x66, 0xf9, 0x85, 0x5f, 0xac, 0xd9, 0x70, 0xdd, 0x83, 0x53, 0x0a, 0xe1, 0x3c, 0x28, 0x78,
	0xb4, 0xe6, 0x41, 0x19, 0xbd, 0xe1, 0x11, 0xde, 0x41, 0xe6, 0x82, 0xd1, 0x6f, 0x92, 0x50, 0x59,
	0x11, 0x36, 0x45, 0x50, 0xe0, 0xfb, 0xa8, 0x15, 0xa6, 0x09, 0xd0, 0x04, 0x02, 0x31, 0x2b, 0xe9,
	0xc1, 0x72, 0xc5, 0x1a, 0x18, 0xdd, 0x65, 0x54, 0x94, 0x90, 0xb3, 0xef, 0x54, 0xda, 0xac, 0xa2,
	0x4b, 0x10, 0x01, 0xe1, 0x41, 0xcd, 0x53, 0x26, 0x6b, 0xe9, 0x0a, 0xca, 0xe0, 0xde, 0xd2, 0x40,
	0x6d, 0x61, 0xb0, 0xf3, 0xdb, 0x34, 0xa1, 0x7a, 0x69, 0x7f, 0xbb, 0xab, 0xed, 0x07, 0x77, 0x75,
	0xef, 0xa4, 0x8e, 0x56, 0xd7, 0x52, 0xf9, 0xe7, 0xb5, 0xac, 0x0e, 0xe5, 0x0d, 0xda, 0xa2, 0xab,
	0xe7, 0xc5, 0x39, 0x56, 0xd7, 0x5e, 0xfe, 0xff, 0x1e, 0xda, 0x9f, 0x60, 0x07, 0x59, 0xf4, 0x36,
	0x8b, 0x48, 0xa2, 0xe8, 0xd2, 0xfb, 0xb5, 0xd5, 0x60, 0xda, 0xf7, 0x40, 0xe1, 0xff, 0xd7, 0x5f,
	0xd5, 0x9f, 0x86, 0x3a, 0x21, 0xfc, 0x1c, 0x61, 0x7f, 0xe4, 0x8d, 0x6e, 0xfc, 0xe0, 0xe6, 0xda,
	0x1f, 0x9c, 0x9f, 0xf6, 0x2f, 0xfa, 0xe7, 0x67, 0xd6, 0x7f, 0xd8, 0x44, 0xd5, 0x81, 0xe7, 0xfb,
	0x96, 0x21, 0x7e, 0x5d, 0x78, 0xfd, 0x2b, 0x6b, 0x03, 0x37, 0x50, 0xed, 0x74, 0xe8, 0xf9, 0x97,
	0x56, 0x45, 0xfc, 0xf4, 0x4e, 0x3e, 0x0f, 0x47, 0x56, 0x55, 0xe0, 0xfe, 0xa7, 0xfe, 0xc0, 0xaa,
	0x8d, 0xeb, 0xd2, 0xff, 0x6f, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0xee, 0x66, 0x0b, 0xf0, 0x85,
	0x05, 0x00, 0x00,
}
