// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/cq/api/bigquery/attempt.proto

package bigquery

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

type Mode int32

const (
	// Default, never set.
	Mode_MODE_UNSPECIFIED Mode = 0
	// Run all tests but do not submit.
	Mode_DRY_RUN Mode = 1
	// Run all tests and potentially submit.
	Mode_FULL_RUN Mode = 2
)

var Mode_name = map[int32]string{
	0: "MODE_UNSPECIFIED",
	1: "DRY_RUN",
	2: "FULL_RUN",
}

var Mode_value = map[string]int32{
	"MODE_UNSPECIFIED": 0,
	"DRY_RUN":          1,
	"FULL_RUN":         2,
}

func (x Mode) String() string {
	return proto.EnumName(Mode_name, int32(x))
}

func (Mode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8792fe122a6ce934, []int{0}
}

type AttemptStatus int32

const (
	// Default, never set.
	AttemptStatus_ATTEMPT_STATUS_UNSPECIFIED AttemptStatus = 0
	// Started but not completed. Used by CQ API, TBD.
	AttemptStatus_STARTED AttemptStatus = 1
	// Ready to submit, all checks passed.
	AttemptStatus_SUCCESS AttemptStatus = 2
	// Attempt stopped before completion, due to some external event and not
	// a failure of the CLs to pass all tests. For example, this may happen
	// when a new patchset is uploaded, a CL is deleted, etc.
	AttemptStatus_ABORTED AttemptStatus = 3
	// Completed and failed some check. This may happen when a build failed,
	// footer syntax was incorrect, or CL was not approved.
	AttemptStatus_FAILURE AttemptStatus = 4
	// Failure in CQ itself caused the Attempt to be dropped.
	AttemptStatus_INFRA_FAILURE AttemptStatus = 5
)

var AttemptStatus_name = map[int32]string{
	0: "ATTEMPT_STATUS_UNSPECIFIED",
	1: "STARTED",
	2: "SUCCESS",
	3: "ABORTED",
	4: "FAILURE",
	5: "INFRA_FAILURE",
}

var AttemptStatus_value = map[string]int32{
	"ATTEMPT_STATUS_UNSPECIFIED": 0,
	"STARTED":                    1,
	"SUCCESS":                    2,
	"ABORTED":                    3,
	"FAILURE":                    4,
	"INFRA_FAILURE":              5,
}

func (x AttemptStatus) String() string {
	return proto.EnumName(AttemptStatus_name, int32(x))
}

func (AttemptStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8792fe122a6ce934, []int{1}
}

type GerritChange_SubmitStatus int32

const (
	// Default value.
	GerritChange_SUBMIT_STATUS_UNSPECIFIED GerritChange_SubmitStatus = 0
	// CQ didn't try submitting this CL.
	//
	// Includes a case where CQ tried submitting the CL, but submission failed
	// due to transient error leaving CL as is, and CQ didn't try again.
	GerritChange_PENDING GerritChange_SubmitStatus = 1
	// CQ tried to submit, but got presumably transient errors and couldn't
	// ascertain whether submission was successful.
	//
	// It's possible that change was actually submitted, but CQ didn't receive
	// a confirmation from Gerrit and follow up checks of the change status
	// failed, too.
	GerritChange_UNKNOWN GerritChange_SubmitStatus = 2
	// CQ tried to submit, but Gerrit rejected the submission because this
	// Change can't be submitted.
	// Typically, this is because a rebase conflict needs to be resolved,
	// or rarely because the change needs some kind of approval.
	GerritChange_FAILURE GerritChange_SubmitStatus = 3
	// CQ submitted this change (aka "merged" in Gerrit jargon).
	//
	// Submission of Gerrit CLs in an Attempt is not an atomic operation,
	// so it's possible that only some of the GerritChanges are submitted.
	GerritChange_SUCCESS GerritChange_SubmitStatus = 4
)

var GerritChange_SubmitStatus_name = map[int32]string{
	0: "SUBMIT_STATUS_UNSPECIFIED",
	1: "PENDING",
	2: "UNKNOWN",
	3: "FAILURE",
	4: "SUCCESS",
}

var GerritChange_SubmitStatus_value = map[string]int32{
	"SUBMIT_STATUS_UNSPECIFIED": 0,
	"PENDING":                   1,
	"UNKNOWN":                   2,
	"FAILURE":                   3,
	"SUCCESS":                   4,
}

func (x GerritChange_SubmitStatus) String() string {
	return proto.EnumName(GerritChange_SubmitStatus_name, int32(x))
}

func (GerritChange_SubmitStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8792fe122a6ce934, []int{1, 0}
}

// Attempt includes the state of one CQ attempt.
//
// An attempt involves doing checks for one or more CLs that could
// potentially be submitted together.
type Attempt struct {
	// The opaque key unique to this Attempt.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// The LUCI project that this attempt belongs to.
	Project string `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
	// An opaque key that is unique for a given set of Gerrit change patchsets
	// (or, equivalently, buildsets). The same cl_group_key will be used if
	// another attempt is made for the same set of changes at a different time.
	ClGroupKey string `protobuf:"bytes,3,opt,name=cl_group_key,json=clGroupKey,proto3" json:"cl_group_key,omitempty"`
	// Similar to cl_group_key, except the key will be the same as long as
	// the earliest_equivalent_patchset values are the same, even if the patchset
	// values are different, e.g. when a new "trivial" patchset is uploaded.
	EquivalentClGroupKey string `protobuf:"bytes,4,opt,name=equivalent_cl_group_key,json=equivalentClGroupKey,proto3" json:"equivalent_cl_group_key,omitempty"`
	// Time when the attempt started (trigger time of the last CL triggered)
	// and ended (released by CQ).
	StartTime *timestamp.Timestamp `protobuf:"bytes,5,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime   *timestamp.Timestamp `protobuf:"bytes,6,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// Gerrit changes, with specific patchsets, in this Attempt.
	// There should be one or more.
	GerritChanges []*GerritChange `protobuf:"bytes,7,rep,name=gerrit_changes,json=gerritChanges,proto3" json:"gerrit_changes,omitempty"`
	// Builds checked as part of this attempt, whether triggered or reused.
	Builds []*Build `protobuf:"bytes,8,rep,name=builds,proto3" json:"builds,omitempty"`
	// Final status of the Attempt.
	Status               AttemptStatus `protobuf:"varint,9,opt,name=status,proto3,enum=bigquery.AttemptStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Attempt) Reset()         { *m = Attempt{} }
func (m *Attempt) String() string { return proto.CompactTextString(m) }
func (*Attempt) ProtoMessage()    {}
func (*Attempt) Descriptor() ([]byte, []int) {
	return fileDescriptor_8792fe122a6ce934, []int{0}
}

func (m *Attempt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Attempt.Unmarshal(m, b)
}
func (m *Attempt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Attempt.Marshal(b, m, deterministic)
}
func (m *Attempt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Attempt.Merge(m, src)
}
func (m *Attempt) XXX_Size() int {
	return xxx_messageInfo_Attempt.Size(m)
}
func (m *Attempt) XXX_DiscardUnknown() {
	xxx_messageInfo_Attempt.DiscardUnknown(m)
}

var xxx_messageInfo_Attempt proto.InternalMessageInfo

func (m *Attempt) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Attempt) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *Attempt) GetClGroupKey() string {
	if m != nil {
		return m.ClGroupKey
	}
	return ""
}

func (m *Attempt) GetEquivalentClGroupKey() string {
	if m != nil {
		return m.EquivalentClGroupKey
	}
	return ""
}

func (m *Attempt) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *Attempt) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *Attempt) GetGerritChanges() []*GerritChange {
	if m != nil {
		return m.GerritChanges
	}
	return nil
}

func (m *Attempt) GetBuilds() []*Build {
	if m != nil {
		return m.Builds
	}
	return nil
}

func (m *Attempt) GetStatus() AttemptStatus {
	if m != nil {
		return m.Status
	}
	return AttemptStatus_ATTEMPT_STATUS_UNSPECIFIED
}

// GerritChange represents one revision (patchset) of one Gerrit change
// in an attempt.
//
// See also: GerritChange in buildbucket/proto/common.proto.
type GerritChange struct {
	// Gerrit hostname, e.g. "chromium-review.googlesource.com".
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	// Gerrit project, e.g. "chromium/src".
	Project string `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
	// Change number, e.g. 12345.
	Change int64 `protobuf:"varint,3,opt,name=change,proto3" json:"change,omitempty"`
	// Patch set number, e.g. 1.
	Patchset int64 `protobuf:"varint,4,opt,name=patchset,proto3" json:"patchset,omitempty"`
	// The earliest patchset of the CL that is considered
	// equivalent to the patchset above.
	EarliestEquivalentPatchset int64 `protobuf:"varint,5,opt,name=earliest_equivalent_patchset,json=earliestEquivalentPatchset,proto3" json:"earliest_equivalent_patchset,omitempty"`
	// The time that the CQ was triggered for this CL in this attempt.
	TriggerTime *timestamp.Timestamp `protobuf:"bytes,6,opt,name=trigger_time,json=triggerTime,proto3" json:"trigger_time,omitempty"`
	// CQ Mode for this CL, e.g. dry run or full run.
	Mode Mode `protobuf:"varint,7,opt,name=mode,proto3,enum=bigquery.Mode" json:"mode,omitempty"`
	// Whether CQ tried to submit this change and the result of the operation.
	SubmitStatus         GerritChange_SubmitStatus `protobuf:"varint,8,opt,name=submit_status,json=submitStatus,proto3,enum=bigquery.GerritChange_SubmitStatus" json:"submit_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *GerritChange) Reset()         { *m = GerritChange{} }
func (m *GerritChange) String() string { return proto.CompactTextString(m) }
func (*GerritChange) ProtoMessage()    {}
func (*GerritChange) Descriptor() ([]byte, []int) {
	return fileDescriptor_8792fe122a6ce934, []int{1}
}

func (m *GerritChange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GerritChange.Unmarshal(m, b)
}
func (m *GerritChange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GerritChange.Marshal(b, m, deterministic)
}
func (m *GerritChange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GerritChange.Merge(m, src)
}
func (m *GerritChange) XXX_Size() int {
	return xxx_messageInfo_GerritChange.Size(m)
}
func (m *GerritChange) XXX_DiscardUnknown() {
	xxx_messageInfo_GerritChange.DiscardUnknown(m)
}

var xxx_messageInfo_GerritChange proto.InternalMessageInfo

func (m *GerritChange) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *GerritChange) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *GerritChange) GetChange() int64 {
	if m != nil {
		return m.Change
	}
	return 0
}

func (m *GerritChange) GetPatchset() int64 {
	if m != nil {
		return m.Patchset
	}
	return 0
}

func (m *GerritChange) GetEarliestEquivalentPatchset() int64 {
	if m != nil {
		return m.EarliestEquivalentPatchset
	}
	return 0
}

func (m *GerritChange) GetTriggerTime() *timestamp.Timestamp {
	if m != nil {
		return m.TriggerTime
	}
	return nil
}

func (m *GerritChange) GetMode() Mode {
	if m != nil {
		return m.Mode
	}
	return Mode_MODE_UNSPECIFIED
}

func (m *GerritChange) GetSubmitStatus() GerritChange_SubmitStatus {
	if m != nil {
		return m.SubmitStatus
	}
	return GerritChange_SUBMIT_STATUS_UNSPECIFIED
}

// Build represents one tryjob Buildbucket build.
//
// See also: Build in buildbucket/proto/build.proto.
type Build struct {
	// Buildbucket build ID, unique per Buildbucket instance.
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Buildbucket host, e.g. "cr-buildbucket.appspot.com".
	Host string `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	// True if this build was pre-existing, or false if it was triggered as
	// part of this CQ attempt.
	Reused bool `protobuf:"varint,3,opt,name=reused,proto3" json:"reused,omitempty"`
	// YES if this build must pass in order for the CLs to be considered
	// ready to submit; NO if the build status should not be used to assess
	// correctness the CLs in the Attempt.
	Critical             bool     `protobuf:"varint,4,opt,name=critical,proto3" json:"critical,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Build) Reset()         { *m = Build{} }
func (m *Build) String() string { return proto.CompactTextString(m) }
func (*Build) ProtoMessage()    {}
func (*Build) Descriptor() ([]byte, []int) {
	return fileDescriptor_8792fe122a6ce934, []int{2}
}

func (m *Build) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Build.Unmarshal(m, b)
}
func (m *Build) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Build.Marshal(b, m, deterministic)
}
func (m *Build) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Build.Merge(m, src)
}
func (m *Build) XXX_Size() int {
	return xxx_messageInfo_Build.Size(m)
}
func (m *Build) XXX_DiscardUnknown() {
	xxx_messageInfo_Build.DiscardUnknown(m)
}

var xxx_messageInfo_Build proto.InternalMessageInfo

func (m *Build) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Build) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *Build) GetReused() bool {
	if m != nil {
		return m.Reused
	}
	return false
}

func (m *Build) GetCritical() bool {
	if m != nil {
		return m.Critical
	}
	return false
}

func init() {
	proto.RegisterEnum("bigquery.Mode", Mode_name, Mode_value)
	proto.RegisterEnum("bigquery.AttemptStatus", AttemptStatus_name, AttemptStatus_value)
	proto.RegisterEnum("bigquery.GerritChange_SubmitStatus", GerritChange_SubmitStatus_name, GerritChange_SubmitStatus_value)
	proto.RegisterType((*Attempt)(nil), "bigquery.Attempt")
	proto.RegisterType((*GerritChange)(nil), "bigquery.GerritChange")
	proto.RegisterType((*Build)(nil), "bigquery.Build")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/cq/api/bigquery/attempt.proto", fileDescriptor_8792fe122a6ce934)
}

var fileDescriptor_8792fe122a6ce934 = []byte{
	// 658 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xdb, 0x6e, 0xda, 0x4c,
	0x18, 0xfc, 0xc1, 0x1c, 0x9c, 0x8f, 0xc3, 0xef, 0xae, 0xa2, 0xc4, 0x45, 0x3d, 0x20, 0x7a, 0x51,
	0x94, 0x0b, 0x5b, 0xa2, 0x8a, 0xaa, 0x5e, 0x44, 0x2a, 0x01, 0x93, 0xa2, 0x04, 0x07, 0xad, 0x6d,
	0x55, 0xbd, 0xb2, 0x8c, 0xbd, 0x35, 0xdb, 0x1a, 0x4c, 0xd6, 0xeb, 0x48, 0x79, 0x94, 0xbe, 0x4e,
	0x9f, 0xac, 0xf2, 0x62, 0x03, 0x51, 0x8f, 0x77, 0x3b, 0xdf, 0xcc, 0xd8, 0x3b, 0xdf, 0x2c, 0x0c,
	0xc2, 0x58, 0xf3, 0x97, 0x2c, 0x5e, 0xd1, 0x74, 0xa5, 0xc5, 0x2c, 0xd4, 0xa3, 0xd4, 0xa7, 0xba,
	0x7f, 0xa7, 0x7b, 0x1b, 0xaa, 0x2f, 0x68, 0x78, 0x97, 0x12, 0xf6, 0xa0, 0x7b, 0x9c, 0x93, 0xd5,
	0x86, 0x6b, 0x1b, 0x16, 0xf3, 0x18, 0xc9, 0xc5, 0xbc, 0xf3, 0x32, 0x8c, 0xe3, 0x30, 0x22, 0xba,
	0x98, 0x2f, 0xd2, 0xcf, 0x3a, 0xa7, 0x2b, 0x92, 0x70, 0x6f, 0xb5, 0xd9, 0x4a, 0x7b, 0xdf, 0x24,
	0xa8, 0x0f, 0xb7, 0x66, 0xa4, 0x80, 0xf4, 0x95, 0x3c, 0xa8, 0xa5, 0x6e, 0xa9, 0x7f, 0x84, 0xb3,
	0x23, 0x52, 0xa1, 0xbe, 0x61, 0xf1, 0x17, 0xe2, 0x73, 0xb5, 0x2c, 0xa6, 0x05, 0x44, 0x5d, 0x68,
	0xfa, 0x91, 0x1b, 0xb2, 0x38, 0xdd, 0xb8, 0x99, 0x49, 0x12, 0x34, 0xf8, 0xd1, 0x55, 0x36, 0xba,
	0x26, 0x0f, 0xe8, 0x1c, 0x4e, 0xc9, 0x5d, 0x4a, 0xef, 0xbd, 0x88, 0xac, 0xb9, 0xfb, 0x48, 0x5c,
	0x11, 0xe2, 0xe3, 0x3d, 0x3d, 0xda, 0xdb, 0xde, 0x01, 0x24, 0xdc, 0x63, 0xdc, 0xcd, 0x6e, 0xaa,
	0x56, 0xbb, 0xa5, 0x7e, 0x63, 0xd0, 0xd1, 0xb6, 0x31, 0xb4, 0x22, 0x86, 0x66, 0x17, 0x31, 0xf0,
	0x91, 0x50, 0x67, 0x18, 0x9d, 0x83, 0x4c, 0xd6, 0xc1, 0xd6, 0x58, 0xfb, 0xab, 0xb1, 0x4e, 0xd6,
	0x81, 0xb0, 0x5d, 0x40, 0x3b, 0x24, 0x8c, 0x51, 0xee, 0xfa, 0x4b, 0x6f, 0x1d, 0x92, 0x44, 0xad,
	0x77, 0xa5, 0x7e, 0x63, 0x70, 0xa2, 0x15, 0x6b, 0xd4, 0xae, 0x04, 0x3f, 0x12, 0x34, 0x6e, 0x85,
	0x07, 0x28, 0x41, 0xaf, 0xa1, 0xb6, 0x48, 0x69, 0x14, 0x24, 0xaa, 0x2c, 0x6c, 0xff, 0xef, 0x6d,
	0x97, 0xd9, 0x1c, 0xe7, 0x34, 0xd2, 0xa1, 0x96, 0x70, 0x8f, 0xa7, 0x89, 0x7a, 0xd4, 0x2d, 0xf5,
	0xdb, 0x83, 0xd3, 0xbd, 0x30, 0x6f, 0xc0, 0x12, 0x34, 0xce, 0x65, 0xbd, 0xef, 0x12, 0x34, 0x0f,
	0xff, 0x8c, 0x10, 0x54, 0x96, 0x71, 0xc2, 0xf3, 0x86, 0xc4, 0xf9, 0x0f, 0x15, 0x9d, 0x40, 0x6d,
	0x1b, 0x48, 0x94, 0x23, 0xe1, 0x1c, 0xa1, 0x0e, 0xc8, 0x1b, 0x8f, 0xfb, 0xcb, 0x84, 0x70, 0xd1,
	0x84, 0x84, 0x77, 0x18, 0xbd, 0x87, 0x67, 0xc4, 0x63, 0x11, 0x25, 0x09, 0x77, 0x0f, 0xda, 0xdb,
	0xe9, 0xab, 0x42, 0xdf, 0x29, 0x34, 0xc6, 0x4e, 0x32, 0x2f, 0xbe, 0x70, 0x01, 0x4d, 0xce, 0x68,
	0x18, 0x12, 0xf6, 0xaf, 0x45, 0x34, 0x72, 0xbd, 0x28, 0xa3, 0x07, 0x95, 0x55, 0x1c, 0x10, 0xb5,
	0x2e, 0x56, 0xd4, 0xde, 0xaf, 0x68, 0x16, 0x07, 0x04, 0x0b, 0x0e, 0x7d, 0x80, 0x56, 0x92, 0x2e,
	0x56, 0x94, 0xbb, 0xf9, 0x3e, 0x65, 0x21, 0x7e, 0xf5, 0xeb, 0xbe, 0x34, 0x4b, 0x68, 0xf3, 0xdd,
	0x36, 0x93, 0x03, 0xd4, 0xf3, 0xa0, 0x79, 0xc8, 0xa2, 0xe7, 0xf0, 0xd4, 0x72, 0x2e, 0x67, 0x53,
	0xdb, 0xb5, 0xec, 0xa1, 0xed, 0x58, 0xae, 0x63, 0x5a, 0x73, 0x63, 0x34, 0x9d, 0x4c, 0x8d, 0xb1,
	0xf2, 0x1f, 0x6a, 0x40, 0x7d, 0x6e, 0x98, 0xe3, 0xa9, 0x79, 0xa5, 0x94, 0x32, 0xe0, 0x98, 0xd7,
	0xe6, 0xed, 0x47, 0x53, 0x29, 0x67, 0x60, 0x32, 0x9c, 0xde, 0x38, 0xd8, 0x50, 0xa4, 0x0c, 0x58,
	0xce, 0x68, 0x64, 0x58, 0x96, 0x52, 0xe9, 0xb9, 0x50, 0x15, 0xcf, 0x00, 0xb5, 0xa1, 0x4c, 0x03,
	0x51, 0x9d, 0x84, 0xcb, 0x34, 0xd8, 0x95, 0x59, 0x3e, 0x28, 0xf3, 0x04, 0x6a, 0x8c, 0xa4, 0x09,
	0x09, 0x44, 0x65, 0x32, 0xce, 0x51, 0x56, 0x99, 0xcf, 0x28, 0xa7, 0xbe, 0x17, 0x89, 0xca, 0x64,
	0xbc, 0xc3, 0x67, 0x6f, 0xa1, 0x92, 0xed, 0x06, 0x1d, 0x83, 0x32, 0xbb, 0x1d, 0x1b, 0x3f, 0x5f,
	0x79, 0x8c, 0x3f, 0xb9, 0xd8, 0x31, 0x95, 0x12, 0x6a, 0x82, 0x3c, 0x71, 0x6e, 0x6e, 0x04, 0x2a,
	0x9f, 0xdd, 0x43, 0xeb, 0xd1, 0xbb, 0x43, 0x2f, 0xa0, 0x33, 0xb4, 0x6d, 0x63, 0x36, 0xff, 0x7d,
	0x7c, 0xcb, 0x1e, 0x62, 0xdb, 0x18, 0x6f, 0xe3, 0x17, 0x21, 0x45, 0xfc, 0xe1, 0xe5, 0xad, 0x60,
	0xa4, 0xc3, 0x5d, 0x54, 0xd0, 0x13, 0x68, 0x4d, 0xcd, 0x09, 0x1e, 0xba, 0xc5, 0xa8, 0xba, 0xa8,
	0x89, 0x37, 0xf0, 0xe6, 0x47, 0x00, 0x00, 0x00, 0xff, 0xff, 0x28, 0x3f, 0x32, 0x43, 0xda, 0x04,
	0x00, 0x00,
}
