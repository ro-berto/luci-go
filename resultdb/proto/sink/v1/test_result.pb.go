// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/resultdb/proto/sink/v1/test_result.proto

package sinkpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_type "go.chromium.org/luci/resultdb/proto/type"
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
//
// This is a copy of luci.resultdb.rpc.v1.TestStatus in
// ../../rpc/v1/test_result.proto, because of https://aip.dev/215.
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
	return fileDescriptor_c45f29128cb6d695, []int{0}
}

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

var TestResultFile_Format_name = map[int32]string{
	0: "LUCI",
	1: "CHROMIUM_JSON_TEST_RESULTS",
	2: "GOOGLE_TEST",
}

var TestResultFile_Format_value = map[string]int32{
	"LUCI":                       0,
	"CHROMIUM_JSON_TEST_RESULTS": 1,
	"GOOGLE_TEST":                2,
}

func (x TestResultFile_Format) String() string {
	return proto.EnumName(TestResultFile_Format_name, int32(x))
}

func (TestResultFile_Format) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c45f29128cb6d695, []int{2, 0}
}

// A local equivalent of luci.resultdb.TestResult message
// in ../../rpc/v1/test_result.proto.
// See its comments for details.
type TestResult struct {
	// Equivalent of luci.resultpb.TestResult.TestPath.
	TestPath string `protobuf:"bytes,1,opt,name=test_path,json=testPath,proto3" json:"test_path,omitempty"`
	// Equivalent of luci.resultpb.TestResult.result_id.
	ResultId string `protobuf:"bytes,2,opt,name=result_id,json=resultId,proto3" json:"result_id,omitempty"`
	// Equivalent of luci.resultpb.TestResult.extra_variant_pairs.
	ExtraVariantPairs *_type.Variant `protobuf:"bytes,3,opt,name=extra_variant_pairs,json=extraVariantPairs,proto3" json:"extra_variant_pairs,omitempty"`
	// Equivalent of luci.resultpb.TestResult.expected.
	Expected bool `protobuf:"varint,4,opt,name=expected,proto3" json:"expected,omitempty"`
	// Equivalent of luci.resultpb.TestResult.status.
	Status TestStatus `protobuf:"varint,5,opt,name=status,proto3,enum=luci.resultdb.sink.TestStatus" json:"status,omitempty"`
	// Equivalent of luci.resultpb.TestResult.summary_markdown.
	SummaryMarkdown string `protobuf:"bytes,6,opt,name=summary_markdown,json=summaryMarkdown,proto3" json:"summary_markdown,omitempty"`
	// Equivalent of luci.resultpb.TestResult.start_time.
	StartTime *timestamp.Timestamp `protobuf:"bytes,7,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// Equivalent of luci.resultpb.TestResult.run_duration.
	RunDuration *duration.Duration `protobuf:"bytes,8,opt,name=run_duration,json=runDuration,proto3" json:"run_duration,omitempty"`
	// Equivalent of luci.resultpb.TestResult.tags.
	Tags []*_type.StringPair `protobuf:"bytes,9,rep,name=tags,proto3" json:"tags,omitempty"`
	// Equivalent of luci.resultpb.TestResult.input_artifacts.
	// The map key is an artifact name.
	InputArtifacts map[string]*Artifact `protobuf:"bytes,10,rep,name=input_artifacts,json=inputArtifacts,proto3" json:"input_artifacts,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Equivalent of luci.resultpb.TestResult.output_artifacts.
	// The map key is an artifact name.
	OutputArtifacts      map[string]*Artifact `protobuf:"bytes,11,rep,name=output_artifacts,json=outputArtifacts,proto3" json:"output_artifacts,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TestResult) Reset()         { *m = TestResult{} }
func (m *TestResult) String() string { return proto.CompactTextString(m) }
func (*TestResult) ProtoMessage()    {}
func (*TestResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_c45f29128cb6d695, []int{0}
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

func (m *TestResult) GetExtraVariantPairs() *_type.Variant {
	if m != nil {
		return m.ExtraVariantPairs
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

func (m *TestResult) GetSummaryMarkdown() string {
	if m != nil {
		return m.SummaryMarkdown
	}
	return ""
}

func (m *TestResult) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *TestResult) GetRunDuration() *duration.Duration {
	if m != nil {
		return m.RunDuration
	}
	return nil
}

func (m *TestResult) GetTags() []*_type.StringPair {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *TestResult) GetInputArtifacts() map[string]*Artifact {
	if m != nil {
		return m.InputArtifacts
	}
	return nil
}

func (m *TestResult) GetOutputArtifacts() map[string]*Artifact {
	if m != nil {
		return m.OutputArtifacts
	}
	return nil
}

// A local equivalent of luci.resultdb.Artifact message
// in ../../rpc/v1/test_result.proto.
// See its comments for details.
// Does not have a name because the latter is avialable as a map key in
// TestResult.input_artifacts and TestResult.output_artifacts.
type Artifact struct {
	// Types that are valid to be assigned to Body:
	//	*Artifact_FilePath
	//	*Artifact_Contents
	Body isArtifact_Body `protobuf_oneof:"body"`
	// Equivalent of luci.resultpb.Artifact.content_type.
	ContentType          string   `protobuf:"bytes,3,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Artifact) Reset()         { *m = Artifact{} }
func (m *Artifact) String() string { return proto.CompactTextString(m) }
func (*Artifact) ProtoMessage()    {}
func (*Artifact) Descriptor() ([]byte, []int) {
	return fileDescriptor_c45f29128cb6d695, []int{1}
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

type isArtifact_Body interface {
	isArtifact_Body()
}

type Artifact_FilePath struct {
	FilePath string `protobuf:"bytes,1,opt,name=file_path,json=filePath,proto3,oneof"`
}

type Artifact_Contents struct {
	Contents []byte `protobuf:"bytes,2,opt,name=contents,proto3,oneof"`
}

func (*Artifact_FilePath) isArtifact_Body() {}

func (*Artifact_Contents) isArtifact_Body() {}

func (m *Artifact) GetBody() isArtifact_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *Artifact) GetFilePath() string {
	if x, ok := m.GetBody().(*Artifact_FilePath); ok {
		return x.FilePath
	}
	return ""
}

func (m *Artifact) GetContents() []byte {
	if x, ok := m.GetBody().(*Artifact_Contents); ok {
		return x.Contents
	}
	return nil
}

func (m *Artifact) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Artifact) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Artifact_FilePath)(nil),
		(*Artifact_Contents)(nil),
	}
}

// A file with test results.
type TestResultFile struct {
	// Absolute OS-native path to the results file on the same machine as the
	// ResultSink server.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// Format of the file.
	Format               TestResultFile_Format `protobuf:"varint,2,opt,name=format,proto3,enum=luci.resultdb.sink.TestResultFile_Format" json:"format,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *TestResultFile) Reset()         { *m = TestResultFile{} }
func (m *TestResultFile) String() string { return proto.CompactTextString(m) }
func (*TestResultFile) ProtoMessage()    {}
func (*TestResultFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_c45f29128cb6d695, []int{2}
}

func (m *TestResultFile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestResultFile.Unmarshal(m, b)
}
func (m *TestResultFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestResultFile.Marshal(b, m, deterministic)
}
func (m *TestResultFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestResultFile.Merge(m, src)
}
func (m *TestResultFile) XXX_Size() int {
	return xxx_messageInfo_TestResultFile.Size(m)
}
func (m *TestResultFile) XXX_DiscardUnknown() {
	xxx_messageInfo_TestResultFile.DiscardUnknown(m)
}

var xxx_messageInfo_TestResultFile proto.InternalMessageInfo

func (m *TestResultFile) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *TestResultFile) GetFormat() TestResultFile_Format {
	if m != nil {
		return m.Format
	}
	return TestResultFile_LUCI
}

func init() {
	proto.RegisterEnum("luci.resultdb.sink.TestStatus", TestStatus_name, TestStatus_value)
	proto.RegisterEnum("luci.resultdb.sink.TestResultFile_Format", TestResultFile_Format_name, TestResultFile_Format_value)
	proto.RegisterType((*TestResult)(nil), "luci.resultdb.sink.TestResult")
	proto.RegisterMapType((map[string]*Artifact)(nil), "luci.resultdb.sink.TestResult.InputArtifactsEntry")
	proto.RegisterMapType((map[string]*Artifact)(nil), "luci.resultdb.sink.TestResult.OutputArtifactsEntry")
	proto.RegisterType((*Artifact)(nil), "luci.resultdb.sink.Artifact")
	proto.RegisterType((*TestResultFile)(nil), "luci.resultdb.sink.TestResultFile")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/resultdb/proto/sink/v1/test_result.proto", fileDescriptor_c45f29128cb6d695)
}

var fileDescriptor_c45f29128cb6d695 = []byte{
	// 732 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x59, 0x6f, 0xf3, 0x44,
	0x14, 0xad, 0xb3, 0xe1, 0xdc, 0x54, 0x89, 0x99, 0x22, 0x64, 0xd2, 0x52, 0x42, 0x9e, 0x52, 0x1e,
	0x6c, 0xe1, 0x02, 0x62, 0x7d, 0x48, 0xd3, 0xa4, 0x35, 0x4d, 0x9b, 0x68, 0xec, 0x20, 0x04, 0x12,
	0x66, 0x92, 0x4c, 0x52, 0xab, 0xf1, 0xa2, 0xf1, 0xb8, 0x34, 0xff, 0x89, 0xff, 0xc0, 0x5f, 0x43,
	0x33, 0xb6, 0xa1, 0x69, 0xfb, 0x6d, 0xd2, 0xf7, 0x94, 0x99, 0x73, 0xcf, 0x39, 0x73, 0x97, 0x5c,
	0xc3, 0x4f, 0xeb, 0xc8, 0x58, 0xdc, 0xb2, 0x28, 0xf0, 0xd3, 0xc0, 0x88, 0xd8, 0xda, 0xdc, 0xa4,
	0x0b, 0xdf, 0x64, 0x34, 0x49, 0x37, 0x7c, 0x39, 0x37, 0x63, 0x16, 0xf1, 0xc8, 0x4c, 0xfc, 0xf0,
	0xce, 0xbc, 0xff, 0xd2, 0xe4, 0x34, 0xe1, 0x5e, 0x16, 0x33, 0x64, 0x04, 0x21, 0x41, 0x37, 0x0a,
	0xba, 0x21, 0x88, 0xed, 0xe3, 0x75, 0x14, 0xad, 0x37, 0x34, 0xd3, 0xce, 0xd3, 0x95, 0xb9, 0x4c,
	0x19, 0xe1, 0x7e, 0x14, 0x66, 0x9a, 0xf6, 0x67, 0x4f, 0xe3, 0xdc, 0x0f, 0x68, 0xc2, 0x49, 0x10,
	0xe7, 0x84, 0xaf, 0xdf, 0x26, 0x27, 0xbe, 0x8d, 0xa9, 0xb9, 0x88, 0x82, 0xa0, 0xf0, 0xed, 0xfe,
	0x53, 0x03, 0x70, 0x69, 0xc2, 0xb1, 0x24, 0xa2, 0x43, 0xa8, 0xcb, 0x7c, 0x63, 0xc2, 0x6f, 0x75,
	0xa5, 0xa3, 0xf4, 0xea, 0x58, 0x15, 0xc0, 0x94, 0xf0, 0x5b, 0x11, 0xcc, 0xfc, 0x3c, 0x7f, 0xa9,
	0x97, 0xb2, 0x60, 0x06, 0xd8, 0x4b, 0x74, 0x05, 0x07, 0xf4, 0x81, 0x33, 0xe2, 0xdd, 0x13, 0xe6,
	0x93, 0x50, 0x58, 0xf8, 0x2c, 0xd1, 0xcb, 0x1d, 0xa5, 0xd7, 0xb0, 0x0e, 0x8d, 0xdd, 0x92, 0x45,
	0x1e, 0xc6, 0x2f, 0x19, 0x11, 0x7f, 0x28, 0x75, 0xf9, 0x6d, 0x2a, 0x54, 0xa8, 0x0d, 0x2a, 0x7d,
	0x88, 0xe9, 0x82, 0xd3, 0xa5, 0x5e, 0xe9, 0x28, 0x3d, 0x15, 0xff, 0x77, 0x47, 0xdf, 0x40, 0x2d,
	0xe1, 0x84, 0xa7, 0x89, 0x5e, 0xed, 0x28, 0xbd, 0xa6, 0x75, 0x6c, 0x3c, 0x6f, 0xa7, 0x21, 0x4a,
	0x72, 0x24, 0x0b, 0xe7, 0x6c, 0x74, 0x02, 0x5a, 0x92, 0x06, 0x01, 0x61, 0x5b, 0x2f, 0x20, 0xec,
	0x6e, 0x19, 0xfd, 0x15, 0xea, 0x35, 0x59, 0x44, 0x2b, 0xc7, 0xaf, 0x73, 0x18, 0x7d, 0x07, 0x90,
	0x70, 0xc2, 0xb8, 0x27, 0x9a, 0xac, 0x7f, 0x20, 0x4b, 0x68, 0x1b, 0xd9, 0x04, 0x8c, 0x62, 0x02,
	0x86, 0x5b, 0x4c, 0x00, 0xd7, 0x25, 0x5b, 0xdc, 0xd1, 0x8f, 0xb0, 0xcf, 0xd2, 0xd0, 0x2b, 0xa6,
	0xa7, 0xab, 0x52, 0xfc, 0xc9, 0x33, 0xf1, 0x79, 0x4e, 0xc0, 0x0d, 0x96, 0x86, 0xc5, 0x05, 0x59,
	0x50, 0xe1, 0x64, 0x9d, 0xe8, 0xf5, 0x4e, 0xb9, 0xd7, 0x78, 0x56, 0x99, 0xec, 0x9a, 0xc3, 0x99,
	0x1f, 0xae, 0x45, 0x9b, 0xb0, 0xe4, 0xa2, 0xdf, 0xa1, 0xe5, 0x87, 0x71, 0xca, 0x3d, 0xc2, 0xb8,
	0xbf, 0x22, 0x0b, 0x9e, 0xe8, 0x20, 0xe5, 0xd6, 0xab, 0x1a, 0x93, 0xcd, 0xda, 0xb0, 0x85, 0xaa,
	0x5f, 0x88, 0x86, 0x21, 0x67, 0x5b, 0xdc, 0xf4, 0x77, 0x40, 0xf4, 0x07, 0x68, 0x51, 0xca, 0x77,
	0xdd, 0x1b, 0xd2, 0xfd, 0xf4, 0x0d, 0xee, 0x13, 0x29, 0x7b, 0x62, 0xdf, 0x8a, 0x76, 0xd1, 0xb6,
	0x07, 0x07, 0x2f, 0xa4, 0x81, 0x34, 0x28, 0xdf, 0xd1, 0x6d, 0xfe, 0x07, 0x14, 0x47, 0x64, 0x41,
	0xf5, 0x9e, 0x6c, 0x52, 0x2a, 0xff, 0x77, 0x0d, 0xeb, 0xe8, 0xa5, 0xd7, 0x0b, 0x13, 0x9c, 0x51,
	0xbf, 0x2f, 0x7d, 0xab, 0xb4, 0xff, 0x84, 0x8f, 0x5e, 0xca, 0xe4, 0xfd, 0xbd, 0xd0, 0x65, 0xa0,
	0x16, 0x30, 0xfa, 0x14, 0xea, 0x2b, 0x7f, 0x43, 0x1f, 0xad, 0xcf, 0xe5, 0x1e, 0x56, 0x05, 0x24,
	0x17, 0xe8, 0x08, 0xd4, 0x45, 0x14, 0x72, 0x1a, 0xf2, 0x44, 0xbe, 0xb2, 0x2f, 0xa2, 0x05, 0x82,
	0x3e, 0x87, 0xfd, 0xfc, 0xec, 0x89, 0x49, 0xcb, 0xd5, 0xa9, 0xe3, 0x46, 0x8e, 0xb9, 0xdb, 0x98,
	0x9e, 0xd5, 0xa0, 0x32, 0x8f, 0x96, 0xdb, 0xee, 0xdf, 0x0a, 0x34, 0xff, 0xef, 0xf5, 0xc8, 0xdf,
	0x50, 0x84, 0xa0, 0xf2, 0x68, 0x69, 0xe5, 0x19, 0xf5, 0xa1, 0xb6, 0x8a, 0x58, 0x40, 0xb8, 0x7c,
	0xad, 0x69, 0x9d, 0xbc, 0x7e, 0x66, 0xc2, 0xc7, 0x18, 0x49, 0x01, 0xce, 0x85, 0xdd, 0x01, 0xd4,
	0x32, 0x04, 0xa9, 0x50, 0x19, 0xcf, 0x06, 0xb6, 0xb6, 0x87, 0x8e, 0xa1, 0x3d, 0xb8, 0xc4, 0x93,
	0x6b, 0x7b, 0x76, 0xed, 0xfd, 0xec, 0x4c, 0x6e, 0x3c, 0x77, 0xe8, 0xb8, 0x1e, 0x1e, 0x3a, 0xb3,
	0xb1, 0xeb, 0x68, 0x0a, 0x6a, 0x41, 0xe3, 0x62, 0x32, 0xb9, 0x18, 0x0f, 0x65, 0x40, 0x2b, 0x7d,
	0xf1, 0x6b, 0xf6, 0x8d, 0xc9, 0x16, 0x12, 0x7d, 0x0c, 0xc8, 0x71, 0xfb, 0xee, 0xcc, 0xf1, 0x66,
	0x37, 0xce, 0x74, 0x38, 0xb0, 0x47, 0xf6, 0xf0, 0x5c, 0xdb, 0x13, 0x0f, 0x4c, 0xfb, 0x8e, 0x30,
	0x50, 0xa1, 0x32, 0xea, 0xdb, 0x63, 0xad, 0x84, 0xea, 0x50, 0x1d, 0xe0, 0xbe, 0x73, 0xa9, 0x95,
	0xc5, 0xb1, 0x7f, 0x36, 0xc1, 0xae, 0x56, 0x11, 0x71, 0xe7, 0xca, 0x9e, 0x6a, 0xd5, 0xb3, 0xaf,
	0x7e, 0xb3, 0xde, 0xe1, 0x5b, 0xfc, 0x83, 0xf8, 0x8d, 0xe7, 0xf3, 0x9a, 0x44, 0x4f, 0xff, 0x0d,
	0x00, 0x00, 0xff, 0xff, 0x31, 0x9d, 0x35, 0x6a, 0xc8, 0x05, 0x00, 0x00,
}
