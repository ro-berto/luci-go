// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/results/proto/v1/invocation.proto

package resultspb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// A container of test results.
// Composable: can include other invocations, see inclusions field.
// Immutable once finalized.
type Invocation struct {
	// A unique identifier of the invocation.
	// LUCI systems MAY create invocations with nicely formatted IDs, such as
	// "build/1234567890". All other clients MUST use GUIDs.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Test results in this invocation.
	//
	// When reading an invocation, by default this field includes tests
	// from the included invocations (see includes field), so the client
	// does not have to load the transitive closure of included invocations.
	Tests []*Invocation_Test `protobuf:"bytes,3,rep,name=tests,proto3" json:"tests,omitempty"`
	// All variant definitions used in this invocation message.
	// Maps a variant id to its definition.
	VariantDefs map[string]*VariantDef `protobuf:"bytes,4,rep,name=variant_defs,json=variantDefs,proto3" json:"variant_defs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Invocation-level string key-value pairs.
	// A key can be repeated.
	Tags []*StringPair `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`
	// The contents of the invocation are incomplete.
	// For example, test harness stopped before executing all tests, or a
	// Swarming bot died.
	Incomplete bool `protobuf:"varint,6,opt,name=incomplete,proto3" json:"incomplete,omitempty"`
	// List of invocations to include in this invocation.
	// When fetching an invocation, test results of the recursively included
	// invocations are presented as a part of the returned invocation.
	// If the same test variant was executed in multiple invocations, all
	// results are unioned.
	Includes []*Invocation_Include `protobuf:"bytes,7,rep,name=includes,proto3" json:"includes,omitempty"`
	// A secret token required in UpdateInvocation.
	// Present only in InsertInvocation response.
	UpdateToken string `protobuf:"bytes,8,opt,name=update_token,json=updateToken,proto3" json:"update_token,omitempty"`
	// This invocation is immutable.
	// Once finalized, fetching the invocation via API is guaranteed to return
	// the same data.
	Final bool `protobuf:"varint,9,opt,name=final,proto3" json:"final,omitempty"`
	// Timestamp when the invocation will be forcefully finalized.
	// Can be extended with UpdateInvocation until finalized.
	Deadline *timestamp.Timestamp `protobuf:"bytes,10,opt,name=deadline,proto3" json:"deadline,omitempty"`
	// When the invocation was created. Read-only.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,11,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// When the invocation was finalized. Read-only.
	FinalizeTime         *timestamp.Timestamp `protobuf:"bytes,12,opt,name=finalize_time,json=finalizeTime,proto3" json:"finalize_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Invocation) Reset()         { *m = Invocation{} }
func (m *Invocation) String() string { return proto.CompactTextString(m) }
func (*Invocation) ProtoMessage()    {}
func (*Invocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_6acaa1bf180a2e75, []int{0}
}

func (m *Invocation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Invocation.Unmarshal(m, b)
}
func (m *Invocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Invocation.Marshal(b, m, deterministic)
}
func (m *Invocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Invocation.Merge(m, src)
}
func (m *Invocation) XXX_Size() int {
	return xxx_messageInfo_Invocation.Size(m)
}
func (m *Invocation) XXX_DiscardUnknown() {
	xxx_messageInfo_Invocation.DiscardUnknown(m)
}

var xxx_messageInfo_Invocation proto.InternalMessageInfo

func (m *Invocation) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Invocation) GetTests() []*Invocation_Test {
	if m != nil {
		return m.Tests
	}
	return nil
}

func (m *Invocation) GetVariantDefs() map[string]*VariantDef {
	if m != nil {
		return m.VariantDefs
	}
	return nil
}

func (m *Invocation) GetTags() []*StringPair {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Invocation) GetIncomplete() bool {
	if m != nil {
		return m.Incomplete
	}
	return false
}

func (m *Invocation) GetIncludes() []*Invocation_Include {
	if m != nil {
		return m.Includes
	}
	return nil
}

func (m *Invocation) GetUpdateToken() string {
	if m != nil {
		return m.UpdateToken
	}
	return ""
}

func (m *Invocation) GetFinal() bool {
	if m != nil {
		return m.Final
	}
	return false
}

func (m *Invocation) GetDeadline() *timestamp.Timestamp {
	if m != nil {
		return m.Deadline
	}
	return nil
}

func (m *Invocation) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Invocation) GetFinalizeTime() *timestamp.Timestamp {
	if m != nil {
		return m.FinalizeTime
	}
	return nil
}

// A container for a given test case.
type Invocation_Test struct {
	// Test path, a unique identifier of the test in a LUCI project.
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
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// Results of specific variants of the test.
	Variants             []*Invocation_TestVariant `protobuf:"bytes,2,rep,name=variants,proto3" json:"variants,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *Invocation_Test) Reset()         { *m = Invocation_Test{} }
func (m *Invocation_Test) String() string { return proto.CompactTextString(m) }
func (*Invocation_Test) ProtoMessage()    {}
func (*Invocation_Test) Descriptor() ([]byte, []int) {
	return fileDescriptor_6acaa1bf180a2e75, []int{0, 0}
}

func (m *Invocation_Test) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Invocation_Test.Unmarshal(m, b)
}
func (m *Invocation_Test) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Invocation_Test.Marshal(b, m, deterministic)
}
func (m *Invocation_Test) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Invocation_Test.Merge(m, src)
}
func (m *Invocation_Test) XXX_Size() int {
	return xxx_messageInfo_Invocation_Test.Size(m)
}
func (m *Invocation_Test) XXX_DiscardUnknown() {
	xxx_messageInfo_Invocation_Test.DiscardUnknown(m)
}

var xxx_messageInfo_Invocation_Test proto.InternalMessageInfo

func (m *Invocation_Test) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Invocation_Test) GetVariants() []*Invocation_TestVariant {
	if m != nil {
		return m.Variants
	}
	return nil
}

// Test-variant-specific information.
type Invocation_TestVariant struct {
	// ID of the variant, a key in Invocation.variant_defs.
	VariantId string `protobuf:"bytes,1,opt,name=variant_id,json=variantId,proto3" json:"variant_id,omitempty"`
	// Individual results of this test variant.
	Results []*TestResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	// Exonerations for this test variant, see Exoneration message comments.
	//
	// Exonerations transcend inclusion edges: it does not matter if an
	// unexpected test result is a part of this invocation directly or
	// indirectly. A Chromium tryjob can include a swarming task
	// invocation with unexpected results, and then add an exoneration if the
	// same test variant fails without patch.
	Exonerations         []*Exoneration `protobuf:"bytes,3,rep,name=exonerations,proto3" json:"exonerations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Invocation_TestVariant) Reset()         { *m = Invocation_TestVariant{} }
func (m *Invocation_TestVariant) String() string { return proto.CompactTextString(m) }
func (*Invocation_TestVariant) ProtoMessage()    {}
func (*Invocation_TestVariant) Descriptor() ([]byte, []int) {
	return fileDescriptor_6acaa1bf180a2e75, []int{0, 1}
}

func (m *Invocation_TestVariant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Invocation_TestVariant.Unmarshal(m, b)
}
func (m *Invocation_TestVariant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Invocation_TestVariant.Marshal(b, m, deterministic)
}
func (m *Invocation_TestVariant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Invocation_TestVariant.Merge(m, src)
}
func (m *Invocation_TestVariant) XXX_Size() int {
	return xxx_messageInfo_Invocation_TestVariant.Size(m)
}
func (m *Invocation_TestVariant) XXX_DiscardUnknown() {
	xxx_messageInfo_Invocation_TestVariant.DiscardUnknown(m)
}

var xxx_messageInfo_Invocation_TestVariant proto.InternalMessageInfo

func (m *Invocation_TestVariant) GetVariantId() string {
	if m != nil {
		return m.VariantId
	}
	return ""
}

func (m *Invocation_TestVariant) GetResults() []*TestResult {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *Invocation_TestVariant) GetExonerations() []*Exoneration {
	if m != nil {
		return m.Exonerations
	}
	return nil
}

// One directed inclusion edge, see also Invocation.inclusions.
type Invocation_Include struct {
	// ID of the included invocation.
	InvocationId string `protobuf:"bytes,1,opt,name=invocation_id,json=invocationId,proto3" json:"invocation_id,omitempty"`
	// The included invocation does not influence the final outcome
	// of this invocation. A typical example is a retry where the last
	// attempt wins and previous attempts no longer matter, hence
	// inconsequential.
	//
	// An existing inclusion edge can be marked inconsequential after insertion,
	// but before the parent invocation is finalized:
	// Call UpdateInvocation with an inclusion of the same invocation id, but
	// inconsequential set to true.
	// It is not possible to change from inconsequential to consequential.
	Inconsequential bool `protobuf:"varint,2,opt,name=inconsequential,proto3" json:"inconsequential,omitempty"`
	// Whether the included invocation is finalized before the parent invocation.
	// The formula for the field is
	//   included_inv.finalize_time < parent_inv.finalize_time
	// If the included invocation is final, but the parent is not yet, the
	// edge is ready. If both are not final yet, the edge is not ready *yet*,
	// but its value may change over time, until the parent invocation is
	// finalized.
	//
	// In practice, either
	// - an edge is ready because the parent is expected to wait for its
	//   children to conclude its own result, OR
	// - it does not matter e.g. if the parent was canceled and finalized
	//   prematurely.
	//
	// When fetching an invocation via API, unready edges are ignored.
	//
	// This field is read-only: values provided by clients are ignored.
	Ready                bool     `protobuf:"varint,3,opt,name=ready,proto3" json:"ready,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Invocation_Include) Reset()         { *m = Invocation_Include{} }
func (m *Invocation_Include) String() string { return proto.CompactTextString(m) }
func (*Invocation_Include) ProtoMessage()    {}
func (*Invocation_Include) Descriptor() ([]byte, []int) {
	return fileDescriptor_6acaa1bf180a2e75, []int{0, 3}
}

func (m *Invocation_Include) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Invocation_Include.Unmarshal(m, b)
}
func (m *Invocation_Include) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Invocation_Include.Marshal(b, m, deterministic)
}
func (m *Invocation_Include) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Invocation_Include.Merge(m, src)
}
func (m *Invocation_Include) XXX_Size() int {
	return xxx_messageInfo_Invocation_Include.Size(m)
}
func (m *Invocation_Include) XXX_DiscardUnknown() {
	xxx_messageInfo_Invocation_Include.DiscardUnknown(m)
}

var xxx_messageInfo_Invocation_Include proto.InternalMessageInfo

func (m *Invocation_Include) GetInvocationId() string {
	if m != nil {
		return m.InvocationId
	}
	return ""
}

func (m *Invocation_Include) GetInconsequential() bool {
	if m != nil {
		return m.Inconsequential
	}
	return false
}

func (m *Invocation_Include) GetReady() bool {
	if m != nil {
		return m.Ready
	}
	return false
}

func init() {
	proto.RegisterType((*Invocation)(nil), "luci.resultsdb.Invocation")
	proto.RegisterMapType((map[string]*VariantDef)(nil), "luci.resultsdb.Invocation.VariantDefsEntry")
	proto.RegisterType((*Invocation_Test)(nil), "luci.resultsdb.Invocation.Test")
	proto.RegisterType((*Invocation_TestVariant)(nil), "luci.resultsdb.Invocation.TestVariant")
	proto.RegisterType((*Invocation_Include)(nil), "luci.resultsdb.Invocation.Include")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/results/proto/v1/invocation.proto", fileDescriptor_6acaa1bf180a2e75)
}

var fileDescriptor_6acaa1bf180a2e75 = []byte{
	// 565 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x5f, 0x6b, 0xdb, 0x3e,
	0x14, 0xc5, 0xf9, 0xd3, 0xb8, 0xd7, 0x6e, 0x7f, 0x45, 0xfc, 0x1e, 0x84, 0xc7, 0xd6, 0xac, 0x83,
	0x11, 0x18, 0xd8, 0x6b, 0xb7, 0x95, 0xb1, 0xc2, 0x0a, 0x65, 0x7d, 0xc8, 0xcb, 0x18, 0x5a, 0xd8,
	0x43, 0x1f, 0x56, 0x14, 0x5b, 0x71, 0x45, 0x6d, 0xc9, 0xb3, 0xe4, 0xb0, 0xec, 0xeb, 0xec, 0x2b,
	0xee, 0x03, 0x0c, 0xcb, 0x72, 0xdc, 0x66, 0xd0, 0xf4, 0xcd, 0x3a, 0x3e, 0xe7, 0x9e, 0xab, 0x7b,
	0x0f, 0x82, 0xd3, 0x54, 0x86, 0xf1, 0x4d, 0x29, 0x73, 0x5e, 0xe5, 0xa1, 0x2c, 0xd3, 0x28, 0xab,
	0x62, 0x1e, 0x95, 0x4c, 0x55, 0x99, 0x56, 0x51, 0x51, 0x4a, 0x2d, 0xa3, 0xe5, 0x71, 0xc4, 0xc5,
	0x52, 0xc6, 0x54, 0x73, 0x29, 0x42, 0x83, 0xa1, 0xfd, 0x9a, 0x17, 0x5a, 0x5e, 0x32, 0x0f, 0x0e,
	0x53, 0x29, 0xd3, 0x8c, 0x35, 0x8a, 0x79, 0xb5, 0x88, 0x34, 0xcf, 0x99, 0xd2, 0x34, 0x2f, 0x1a,
	0x41, 0x70, 0xf2, 0x38, 0xa3, 0x58, 0xe6, 0x79, 0x6b, 0x72, 0xf4, 0x67, 0x04, 0x30, 0x5d, 0x3b,
	0xa3, 0x7d, 0xe8, 0xf1, 0x04, 0x3b, 0x63, 0x67, 0xb2, 0x4b, 0x7a, 0x3c, 0x41, 0xef, 0x60, 0xa8,
	0x99, 0xd2, 0x0a, 0xf7, 0xc7, 0xfd, 0x89, 0x77, 0x72, 0x18, 0xde, 0xef, 0x29, 0xec, 0xa4, 0xe1,
	0x8c, 0x29, 0x4d, 0x1a, 0x36, 0xfa, 0x0c, 0xfe, 0x92, 0x96, 0x9c, 0x0a, 0x7d, 0x9d, 0xb0, 0x85,
	0xc2, 0x03, 0xa3, 0x7e, 0xf5, 0x80, 0xfa, 0x5b, 0x43, 0xff, 0xc4, 0x16, 0xea, 0x52, 0xe8, 0x72,
	0x45, 0xbc, 0x65, 0x87, 0xa0, 0x10, 0x06, 0x9a, 0xa6, 0x0a, 0x0f, 0x4d, 0x9d, 0x60, 0xb3, 0xce,
	0x57, 0x5d, 0x72, 0x91, 0x7e, 0xa1, 0xbc, 0x24, 0x86, 0x87, 0x9e, 0x01, 0x70, 0x11, 0xcb, 0xbc,
	0xc8, 0x98, 0x66, 0x78, 0x67, 0xec, 0x4c, 0x5c, 0x72, 0x07, 0x41, 0x1f, 0xc1, 0xe5, 0x22, 0xce,
	0xaa, 0x84, 0x29, 0x3c, 0x32, 0x35, 0x8f, 0x1e, 0xe8, 0x6d, 0xda, 0x50, 0xc9, 0x5a, 0x83, 0x9e,
	0x83, 0x5f, 0x15, 0x09, 0xd5, 0xec, 0x5a, 0xcb, 0x5b, 0x26, 0xb0, 0x6b, 0x06, 0xe6, 0x35, 0xd8,
	0xac, 0x86, 0xd0, 0xff, 0x30, 0x5c, 0x70, 0x41, 0x33, 0xbc, 0x6b, 0xdc, 0x9b, 0x03, 0x3a, 0x05,
	0x37, 0x61, 0x34, 0xc9, 0xb8, 0x60, 0x18, 0xc6, 0x8e, 0xb9, 0x4c, 0xb3, 0xd6, 0xb0, 0x5d, 0x6b,
	0x38, 0x6b, 0xd7, 0x4a, 0xd6, 0x5c, 0x74, 0x06, 0x5e, 0x5c, 0x32, 0x63, 0xc8, 0x73, 0x86, 0xbd,
	0xad, 0x52, 0x68, 0xe8, 0x35, 0x80, 0xce, 0x61, 0xcf, 0xb8, 0xf3, 0x5f, 0x56, 0xee, 0x6f, 0x95,
	0xfb, 0xad, 0xa0, 0x86, 0x82, 0xef, 0x30, 0xa8, 0xb7, 0x8b, 0x10, 0x0c, 0x0a, 0xaa, 0x6f, 0x6c,
	0x3e, 0xcc, 0x37, 0xba, 0x00, 0xd7, 0x6e, 0x4a, 0xe1, 0x9e, 0x19, 0xe5, 0xcb, 0x2d, 0x21, 0xb1,
	0xab, 0x26, 0x6b, 0x5d, 0xf0, 0xdb, 0x01, 0xef, 0xce, 0x1f, 0xf4, 0x14, 0xa0, 0x8d, 0xcf, 0x3a,
	0x8d, 0xbb, 0x16, 0x99, 0x26, 0xe8, 0x2d, 0x8c, 0x6c, 0x71, 0xeb, 0xf8, 0x4f, 0x20, 0x4c, 0x16,
	0xcd, 0x89, 0xb4, 0x54, 0x74, 0x0e, 0x3e, 0xfb, 0x29, 0x05, 0x2b, 0x4d, 0x27, 0x6d, 0xa2, 0x9f,
	0x6c, 0x4a, 0x2f, 0x3b, 0x0e, 0xb9, 0x27, 0x08, 0xae, 0xe0, 0x60, 0x33, 0xa5, 0xe8, 0x00, 0xfa,
	0xb7, 0x6c, 0x65, 0x5b, 0xac, 0x3f, 0xd1, 0x6b, 0x18, 0x2e, 0x69, 0x56, 0x31, 0xdc, 0xb3, 0x43,
	0xde, 0xa8, 0xdf, 0x95, 0x20, 0x0d, 0xf1, 0x43, 0xef, 0xbd, 0x13, 0x08, 0x18, 0xd9, 0x94, 0xa1,
	0x17, 0xb0, 0xd7, 0x3d, 0x05, 0xdd, 0xfd, 0xfd, 0x0e, 0x9c, 0x26, 0x68, 0x02, 0xff, 0xd5, 0x71,
	0x16, 0x8a, 0xfd, 0xa8, 0x98, 0xd0, 0x9c, 0x66, 0xc6, 0xcf, 0x25, 0x9b, 0x70, 0x9d, 0xc3, 0x92,
	0xd1, 0x64, 0x85, 0xfb, 0x4d, 0x0e, 0xcd, 0xe1, 0xe2, 0xf8, 0x2a, 0x7a, 0xd4, 0x63, 0x71, 0x66,
	0x81, 0x62, 0x3e, 0xdf, 0x31, 0xd8, 0x9b, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x26, 0x0f, 0xad,
	0xe9, 0xcf, 0x04, 0x00, 0x00,
}
