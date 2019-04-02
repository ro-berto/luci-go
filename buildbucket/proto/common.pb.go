// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/buildbucket/proto/common.proto

package buildbucketpb

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// Status of a build or a step.
type Status int32

const (
	// Unspecified state. Meaning depends on the context.
	Status_STATUS_UNSPECIFIED Status = 0
	// Build was scheduled, but did not start or end yet.
	Status_SCHEDULED Status = 1
	// Build/step has started.
	Status_STARTED Status = 2
	// A union of all terminal statuses.
	// Can be used in BuildPredicate.status.
	// A concrete build/step cannot have this status.
	// Can be used as a bitmask to check that a build/step ended.
	Status_ENDED_MASK Status = 4
	// A build/step ended successfully.
	// This is a terminal status. It may not transition to another status.
	Status_SUCCESS Status = 12
	// A build/step ended unsuccessfully due to its Build.Input,
	// e.g. tests failed, and NOT due to a build infrastructure failure.
	// This is a terminal status. It may not transition to another status.
	Status_FAILURE Status = 20
	// A build/step ended unsuccessfully due to a failure independent of the
	// input, e.g. swarming failed, not enough capacity or the recipe was unable
	// to read the patch from gerrit.
	// start_time is not required for this status.
	// This is a terminal status. It may not transition to another status.
	Status_INFRA_FAILURE Status = 36
	// A build was cancelled explicitly, e.g. via an RPC.
	// This is a terminal status. It may not transition to another status.
	Status_CANCELED Status = 68
)

var Status_name = map[int32]string{
	0:  "STATUS_UNSPECIFIED",
	1:  "SCHEDULED",
	2:  "STARTED",
	4:  "ENDED_MASK",
	12: "SUCCESS",
	20: "FAILURE",
	36: "INFRA_FAILURE",
	68: "CANCELED",
}

var Status_value = map[string]int32{
	"STATUS_UNSPECIFIED": 0,
	"SCHEDULED":          1,
	"STARTED":            2,
	"ENDED_MASK":         4,
	"SUCCESS":            12,
	"FAILURE":            20,
	"INFRA_FAILURE":      36,
	"CANCELED":           68,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a1a0c34bd7fcf0dc, []int{0}
}

// A boolean with an undefined value.
type Trinary int32

const (
	Trinary_UNSET Trinary = 0
	Trinary_YES   Trinary = 1
	Trinary_NO    Trinary = 2
)

var Trinary_name = map[int32]string{
	0: "UNSET",
	1: "YES",
	2: "NO",
}

var Trinary_value = map[string]int32{
	"UNSET": 0,
	"YES":   1,
	"NO":    2,
}

func (x Trinary) String() string {
	return proto.EnumName(Trinary_name, int32(x))
}

func (Trinary) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a1a0c34bd7fcf0dc, []int{1}
}

// An executable to run when the build is ready to start.
//
// The following describes the protocol between buildbucket and the executable.
// The executable MUST NOT assume/rely on anything that is not explicitly
// specified below.
//
// The executable will be started with the following input:
// - empty current working directory (CWD).
// - stdin: binary-encoded buildbucket.v2.Build message.
// - $TMPDIR, $TEMP, $TMP: env variables point to an empty directory, on the
//   same file system as CWD.
// - $LOGDOG_COORDINATOR_HOST, $LOGDOG_STREAM_PROJECT, $LOGDOG_STREAM_PREFIX,
//   $LOGDOG_NAMESPACE, $LOGDOG_STREAM_SERVER_PATH: env variables describing
//   LogDog context, part of the LogDog Butler protocol.
//   The executable MAY create logdog streams using these variables.
// - LUCI_CONTEXT["run_build"]["cache_dir"]: a LUCI context value pointing to
//   the root directory of Swarming named caches.
//   MAY be on the same file system as CWD.
//   For example, builder cache is available at "<cache_dir>/builder".
//   See also CacheEntry.
//
// The executable MUST write "$LOGDOG_NAMESPACE/build.proto" stream using the
// LogDog streamserver at $LOGDOG_STREAM_SERVER_PATH.
// The stream MUST have a binary-encoded buildbucket.v2.Build message datagrams,
// with content type "application/luci+proto; message=buildbucket.v2.Build".
// Server-side build will be updated with the values from the latest message.
//
// A build step S without children in the datagram MAY have a Step.Log named
// "$build.proto".
// It MUST point to a LogDog datagram stream with binary-encoded
// buildbucket.v2.Build messages and content type
// "application/luci+proto; message=buildbucket.v2.Build".
// The step tree from the second datagram will appear as substeps of step S.
// This rule applies recursively, i.e. a leaf step in the datagram MAY also
// have a "$build.proto" log.
// The graph of datagram streams MUST be a tree, i.e. acyclic.
//
// All build step log urls of all Build messages MUST be relative to
// $LOGDOG_STREAM_PREFIX (only stream name) and start with $LOGDOG_NAMESPACE.
type Executable struct {
	// The CIPD package containing the executable.
	// On Linux/Mac, the executable MUST be named "run_build".
	// On Windows, it MUST be named "run_build.exe" or "run_build.bat",
	// in this order of precedence.
	CipdPackage string `protobuf:"bytes,1,opt,name=cipd_package,json=cipdPackage,proto3" json:"cipd_package,omitempty"`
	// The CIPD version to fetch. Defaults to `latest`.
	CipdVersion          string   `protobuf:"bytes,2,opt,name=cipd_version,json=cipdVersion,proto3" json:"cipd_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Executable) Reset()         { *m = Executable{} }
func (m *Executable) String() string { return proto.CompactTextString(m) }
func (*Executable) ProtoMessage()    {}
func (*Executable) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1a0c34bd7fcf0dc, []int{0}
}

func (m *Executable) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Executable.Unmarshal(m, b)
}
func (m *Executable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Executable.Marshal(b, m, deterministic)
}
func (m *Executable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Executable.Merge(m, src)
}
func (m *Executable) XXX_Size() int {
	return xxx_messageInfo_Executable.Size(m)
}
func (m *Executable) XXX_DiscardUnknown() {
	xxx_messageInfo_Executable.DiscardUnknown(m)
}

var xxx_messageInfo_Executable proto.InternalMessageInfo

func (m *Executable) GetCipdPackage() string {
	if m != nil {
		return m.CipdPackage
	}
	return ""
}

func (m *Executable) GetCipdVersion() string {
	if m != nil {
		return m.CipdVersion
	}
	return ""
}

// Machine-readable details of a status.
// Human-readble details are present in a sibling summary_markdown field.
type StatusDetails struct {
	// If set, indicates that the failure was due to a resource exhaustion / quota
	// denial.
	// Applicable in FAILURE and INFRA_FAILURE statuses.
	ResourceExhaustion *StatusDetails_ResourceExhaustion `protobuf:"bytes,3,opt,name=resource_exhaustion,json=resourceExhaustion,proto3" json:"resource_exhaustion,omitempty"`
	// If set, indicates that the failure was due to a timeout.
	// Applicable in FAILURE and INFRA_FAILURE statuses.
	Timeout              *StatusDetails_Timeout `protobuf:"bytes,4,opt,name=timeout,proto3" json:"timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *StatusDetails) Reset()         { *m = StatusDetails{} }
func (m *StatusDetails) String() string { return proto.CompactTextString(m) }
func (*StatusDetails) ProtoMessage()    {}
func (*StatusDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1a0c34bd7fcf0dc, []int{1}
}

func (m *StatusDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusDetails.Unmarshal(m, b)
}
func (m *StatusDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusDetails.Marshal(b, m, deterministic)
}
func (m *StatusDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusDetails.Merge(m, src)
}
func (m *StatusDetails) XXX_Size() int {
	return xxx_messageInfo_StatusDetails.Size(m)
}
func (m *StatusDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusDetails.DiscardUnknown(m)
}

var xxx_messageInfo_StatusDetails proto.InternalMessageInfo

func (m *StatusDetails) GetResourceExhaustion() *StatusDetails_ResourceExhaustion {
	if m != nil {
		return m.ResourceExhaustion
	}
	return nil
}

func (m *StatusDetails) GetTimeout() *StatusDetails_Timeout {
	if m != nil {
		return m.Timeout
	}
	return nil
}

type StatusDetails_ResourceExhaustion struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusDetails_ResourceExhaustion) Reset()         { *m = StatusDetails_ResourceExhaustion{} }
func (m *StatusDetails_ResourceExhaustion) String() string { return proto.CompactTextString(m) }
func (*StatusDetails_ResourceExhaustion) ProtoMessage()    {}
func (*StatusDetails_ResourceExhaustion) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1a0c34bd7fcf0dc, []int{1, 0}
}

func (m *StatusDetails_ResourceExhaustion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusDetails_ResourceExhaustion.Unmarshal(m, b)
}
func (m *StatusDetails_ResourceExhaustion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusDetails_ResourceExhaustion.Marshal(b, m, deterministic)
}
func (m *StatusDetails_ResourceExhaustion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusDetails_ResourceExhaustion.Merge(m, src)
}
func (m *StatusDetails_ResourceExhaustion) XXX_Size() int {
	return xxx_messageInfo_StatusDetails_ResourceExhaustion.Size(m)
}
func (m *StatusDetails_ResourceExhaustion) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusDetails_ResourceExhaustion.DiscardUnknown(m)
}

var xxx_messageInfo_StatusDetails_ResourceExhaustion proto.InternalMessageInfo

type StatusDetails_Timeout struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusDetails_Timeout) Reset()         { *m = StatusDetails_Timeout{} }
func (m *StatusDetails_Timeout) String() string { return proto.CompactTextString(m) }
func (*StatusDetails_Timeout) ProtoMessage()    {}
func (*StatusDetails_Timeout) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1a0c34bd7fcf0dc, []int{1, 1}
}

func (m *StatusDetails_Timeout) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusDetails_Timeout.Unmarshal(m, b)
}
func (m *StatusDetails_Timeout) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusDetails_Timeout.Marshal(b, m, deterministic)
}
func (m *StatusDetails_Timeout) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusDetails_Timeout.Merge(m, src)
}
func (m *StatusDetails_Timeout) XXX_Size() int {
	return xxx_messageInfo_StatusDetails_Timeout.Size(m)
}
func (m *StatusDetails_Timeout) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusDetails_Timeout.DiscardUnknown(m)
}

var xxx_messageInfo_StatusDetails_Timeout proto.InternalMessageInfo

// A Gerrit patchset.
type GerritChange struct {
	// Gerrit hostname, e.g. "chromium-review.googlesource.com".
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	// Gerrit project, e.g. "chromium/src".
	Project string `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
	// Change number, e.g. 12345.
	Change int64 `protobuf:"varint,3,opt,name=change,proto3" json:"change,omitempty"`
	// Patch set number, e.g. 1.
	Patchset             int64    `protobuf:"varint,4,opt,name=patchset,proto3" json:"patchset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GerritChange) Reset()         { *m = GerritChange{} }
func (m *GerritChange) String() string { return proto.CompactTextString(m) }
func (*GerritChange) ProtoMessage()    {}
func (*GerritChange) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1a0c34bd7fcf0dc, []int{2}
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

// A landed Git commit hosted on Gitiles.
type GitilesCommit struct {
	// Gitiles hostname, e.g. "chromium.googlesource.com".
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	// Repository name on the host, e.g. "chromium/src".
	Project string `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
	// Commit HEX SHA1.
	Id string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// Commit ref, e.g. "refs/heads/master".
	// NOT a branch name: if specified, must start with "refs/".
	Ref string `protobuf:"bytes,4,opt,name=ref,proto3" json:"ref,omitempty"`
	// Defines a total order of commits on the ref. Requires ref field.
	// Typically 1-based, monotonically increasing, contiguous integer
	// defined by a Gerrit plugin, goto.google.com/git-numberer.
	// TODO(tandrii): make it a public doc.
	Position             uint32   `protobuf:"varint,5,opt,name=position,proto3" json:"position,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GitilesCommit) Reset()         { *m = GitilesCommit{} }
func (m *GitilesCommit) String() string { return proto.CompactTextString(m) }
func (*GitilesCommit) ProtoMessage()    {}
func (*GitilesCommit) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1a0c34bd7fcf0dc, []int{3}
}

func (m *GitilesCommit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GitilesCommit.Unmarshal(m, b)
}
func (m *GitilesCommit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GitilesCommit.Marshal(b, m, deterministic)
}
func (m *GitilesCommit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GitilesCommit.Merge(m, src)
}
func (m *GitilesCommit) XXX_Size() int {
	return xxx_messageInfo_GitilesCommit.Size(m)
}
func (m *GitilesCommit) XXX_DiscardUnknown() {
	xxx_messageInfo_GitilesCommit.DiscardUnknown(m)
}

var xxx_messageInfo_GitilesCommit proto.InternalMessageInfo

func (m *GitilesCommit) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *GitilesCommit) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *GitilesCommit) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GitilesCommit) GetRef() string {
	if m != nil {
		return m.Ref
	}
	return ""
}

func (m *GitilesCommit) GetPosition() uint32 {
	if m != nil {
		return m.Position
	}
	return 0
}

// A key-value pair of strings.
type StringPair struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringPair) Reset()         { *m = StringPair{} }
func (m *StringPair) String() string { return proto.CompactTextString(m) }
func (*StringPair) ProtoMessage()    {}
func (*StringPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1a0c34bd7fcf0dc, []int{4}
}

func (m *StringPair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringPair.Unmarshal(m, b)
}
func (m *StringPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringPair.Marshal(b, m, deterministic)
}
func (m *StringPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringPair.Merge(m, src)
}
func (m *StringPair) XXX_Size() int {
	return xxx_messageInfo_StringPair.Size(m)
}
func (m *StringPair) XXX_DiscardUnknown() {
	xxx_messageInfo_StringPair.DiscardUnknown(m)
}

var xxx_messageInfo_StringPair proto.InternalMessageInfo

func (m *StringPair) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *StringPair) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// Half-open time range.
type TimeRange struct {
	// Inclusive lower boundary. Optional.
	StartTime *timestamp.Timestamp `protobuf:"bytes,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// Exclusive upper boundary. Optional.
	EndTime              *timestamp.Timestamp `protobuf:"bytes,2,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TimeRange) Reset()         { *m = TimeRange{} }
func (m *TimeRange) String() string { return proto.CompactTextString(m) }
func (*TimeRange) ProtoMessage()    {}
func (*TimeRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1a0c34bd7fcf0dc, []int{5}
}

func (m *TimeRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeRange.Unmarshal(m, b)
}
func (m *TimeRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeRange.Marshal(b, m, deterministic)
}
func (m *TimeRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeRange.Merge(m, src)
}
func (m *TimeRange) XXX_Size() int {
	return xxx_messageInfo_TimeRange.Size(m)
}
func (m *TimeRange) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeRange.DiscardUnknown(m)
}

var xxx_messageInfo_TimeRange proto.InternalMessageInfo

func (m *TimeRange) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *TimeRange) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

// A requested dimension. Looks like StringPair, but also has an expiration.
type RequestedDimension struct {
	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	// If set, ignore this dimension after this duraiton.
	Expiration           *duration.Duration `protobuf:"bytes,3,opt,name=expiration,proto3" json:"expiration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *RequestedDimension) Reset()         { *m = RequestedDimension{} }
func (m *RequestedDimension) String() string { return proto.CompactTextString(m) }
func (*RequestedDimension) ProtoMessage()    {}
func (*RequestedDimension) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1a0c34bd7fcf0dc, []int{6}
}

func (m *RequestedDimension) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestedDimension.Unmarshal(m, b)
}
func (m *RequestedDimension) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestedDimension.Marshal(b, m, deterministic)
}
func (m *RequestedDimension) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestedDimension.Merge(m, src)
}
func (m *RequestedDimension) XXX_Size() int {
	return xxx_messageInfo_RequestedDimension.Size(m)
}
func (m *RequestedDimension) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestedDimension.DiscardUnknown(m)
}

var xxx_messageInfo_RequestedDimension proto.InternalMessageInfo

func (m *RequestedDimension) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *RequestedDimension) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *RequestedDimension) GetExpiration() *duration.Duration {
	if m != nil {
		return m.Expiration
	}
	return nil
}

func init() {
	proto.RegisterEnum("buildbucket.v2.Status", Status_name, Status_value)
	proto.RegisterEnum("buildbucket.v2.Trinary", Trinary_name, Trinary_value)
	proto.RegisterType((*Executable)(nil), "buildbucket.v2.Executable")
	proto.RegisterType((*StatusDetails)(nil), "buildbucket.v2.StatusDetails")
	proto.RegisterType((*StatusDetails_ResourceExhaustion)(nil), "buildbucket.v2.StatusDetails.ResourceExhaustion")
	proto.RegisterType((*StatusDetails_Timeout)(nil), "buildbucket.v2.StatusDetails.Timeout")
	proto.RegisterType((*GerritChange)(nil), "buildbucket.v2.GerritChange")
	proto.RegisterType((*GitilesCommit)(nil), "buildbucket.v2.GitilesCommit")
	proto.RegisterType((*StringPair)(nil), "buildbucket.v2.StringPair")
	proto.RegisterType((*TimeRange)(nil), "buildbucket.v2.TimeRange")
	proto.RegisterType((*RequestedDimension)(nil), "buildbucket.v2.RequestedDimension")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/buildbucket/proto/common.proto", fileDescriptor_a1a0c34bd7fcf0dc)
}

var fileDescriptor_a1a0c34bd7fcf0dc = []byte{
	// 668 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x41, 0x6f, 0xdb, 0x36,
	0x18, 0x8d, 0x64, 0xc7, 0xb6, 0x3e, 0xdb, 0x81, 0xc6, 0x05, 0x81, 0xe7, 0xc3, 0x96, 0x19, 0x0b,
	0x10, 0xe4, 0x20, 0x0f, 0x49, 0x36, 0x60, 0xd8, 0xa1, 0x70, 0x25, 0x25, 0x75, 0x9a, 0xba, 0x01,
	0x65, 0x17, 0x68, 0x2f, 0x06, 0x2d, 0x31, 0x36, 0x1b, 0x4b, 0x54, 0x29, 0x2a, 0x4d, 0x50, 0xf4,
	0xdc, 0xbf, 0xd9, 0x9f, 0x52, 0x90, 0x92, 0xdd, 0xa4, 0x01, 0xda, 0xf4, 0xc6, 0xf7, 0xf8, 0x1e,
	0x1f, 0xbf, 0x8f, 0x9f, 0x04, 0x47, 0x73, 0xee, 0x84, 0x0b, 0xc1, 0x63, 0x96, 0xc7, 0x0e, 0x17,
	0xf3, 0xfe, 0x32, 0x0f, 0x59, 0x7f, 0x96, 0xb3, 0x65, 0x34, 0xcb, 0xc3, 0x2b, 0x2a, 0xfb, 0xa9,
	0xe0, 0x92, 0xf7, 0x43, 0x1e, 0xc7, 0x3c, 0x71, 0x34, 0x40, 0x5b, 0x77, 0xf6, 0x9d, 0xeb, 0xc3,
	0xee, 0xef, 0x73, 0xce, 0xe7, 0x4b, 0x5a, 0x48, 0x67, 0xf9, 0x65, 0x3f, 0xca, 0x05, 0x91, 0x6c,
	0xa5, 0xef, 0xfe, 0xf1, 0xed, 0xbe, 0x64, 0x31, 0xcd, 0x24, 0x89, 0xd3, 0x42, 0xd0, 0xc3, 0x00,
	0xfe, 0x0d, 0x0d, 0x73, 0x49, 0x66, 0x4b, 0x8a, 0xfe, 0x84, 0x56, 0xc8, 0xd2, 0x68, 0x9a, 0x92,
	0xf0, 0x8a, 0xcc, 0x69, 0xc7, 0xd8, 0x35, 0xf6, 0x2d, 0xdc, 0x54, 0xdc, 0x45, 0x41, 0xad, 0x25,
	0xd7, 0x54, 0x64, 0x8c, 0x27, 0x1d, 0xf3, 0xab, 0xe4, 0x55, 0x41, 0xf5, 0x3e, 0x1b, 0xd0, 0x0e,
	0x24, 0x91, 0x79, 0xe6, 0x51, 0x49, 0xd8, 0x32, 0x43, 0x04, 0x7e, 0x15, 0x34, 0xe3, 0xb9, 0x08,
	0xe9, 0x94, 0xde, 0x2c, 0x48, 0x9e, 0xa9, 0x3b, 0x76, 0x2a, 0xbb, 0xc6, 0x7e, 0xf3, 0xf0, 0x6f,
	0xe7, 0x7e, 0x51, 0xce, 0x3d, 0xaf, 0x83, 0x4b, 0xa3, 0xbf, 0xf6, 0x61, 0x24, 0x1e, 0x70, 0xe8,
	0x09, 0xd4, 0x55, 0x6d, 0x3c, 0x97, 0x9d, 0xaa, 0x3e, 0x76, 0xef, 0xfb, 0xc7, 0x8e, 0x0b, 0x31,
	0x5e, 0xb9, 0xba, 0xdb, 0x80, 0x1e, 0x46, 0x75, 0x2d, 0xa8, 0x97, 0xca, 0xb3, 0x6a, 0xc3, 0xb0,
	0xcd, 0xb3, 0x6a, 0xc3, 0xb4, 0x2b, 0xbd, 0x14, 0x5a, 0xa7, 0x54, 0x08, 0x26, 0xdd, 0x05, 0x49,
	0xe6, 0x14, 0x21, 0xa8, 0x2e, 0x78, 0x26, 0xcb, 0x86, 0xe9, 0x35, 0xea, 0x40, 0x3d, 0x15, 0xfc,
	0x2d, 0x0d, 0x65, 0xd9, 0xa4, 0x15, 0x44, 0x3b, 0x50, 0x0b, 0xb5, 0x4f, 0x77, 0xa0, 0x82, 0x4b,
	0x84, 0xba, 0xd0, 0x48, 0x89, 0x0c, 0x17, 0x19, 0x2d, 0x8a, 0xa8, 0xe0, 0x35, 0xee, 0x7d, 0x80,
	0xf6, 0x29, 0x93, 0x6c, 0x49, 0x33, 0x97, 0xc7, 0x31, 0x93, 0x3f, 0x19, 0xb9, 0x05, 0x26, 0x8b,
	0x74, 0x9c, 0x85, 0x4d, 0x16, 0x21, 0x1b, 0x2a, 0x82, 0x5e, 0xea, 0x14, 0x0b, 0xab, 0xa5, 0x0e,
	0xe7, 0x19, 0xd3, 0x0f, 0xb3, 0xb9, 0x6b, 0xec, 0xb7, 0xf1, 0x1a, 0xf7, 0x8e, 0x01, 0x02, 0x29,
	0x58, 0x32, 0xbf, 0x20, 0x4c, 0x28, 0xef, 0x15, 0xbd, 0x2d, 0x83, 0xd5, 0x12, 0x6d, 0xc3, 0xe6,
	0x35, 0x59, 0xe6, 0xb4, 0x4c, 0x2d, 0x40, 0xef, 0x23, 0x58, 0xaa, 0x77, 0x58, 0xd7, 0xf6, 0x1f,
	0x40, 0x26, 0x89, 0x90, 0x53, 0xd5, 0x6f, 0xed, 0x6d, 0x1e, 0x76, 0x9d, 0x62, 0x3c, 0x9d, 0xd5,
	0x78, 0xea, 0x57, 0xd1, 0xe3, 0x89, 0x2d, 0xad, 0x56, 0x18, 0xfd, 0x03, 0x0d, 0x9a, 0x44, 0x85,
	0xd1, 0xfc, 0xa1, 0xb1, 0x4e, 0x93, 0x48, 0xa1, 0xde, 0x7b, 0xf5, 0xa0, 0xef, 0x72, 0x9a, 0x49,
	0x1a, 0x79, 0x2c, 0xa6, 0x89, 0x1a, 0xce, 0xc7, 0x5e, 0x5e, 0xdd, 0x97, 0xde, 0xa4, 0xac, 0xf8,
	0x9a, 0xca, 0x49, 0xfd, 0xed, 0x41, 0xac, 0x57, 0x7e, 0x6e, 0xf8, 0x8e, 0xf8, 0xe0, 0x93, 0x01,
	0xb5, 0x62, 0xd8, 0xd0, 0x0e, 0xa0, 0x60, 0x3c, 0x18, 0x4f, 0x82, 0xe9, 0x64, 0x14, 0x5c, 0xf8,
	0xee, 0xf0, 0x64, 0xe8, 0x7b, 0xf6, 0x06, 0x6a, 0x83, 0x15, 0xb8, 0xcf, 0x7c, 0x6f, 0x72, 0xee,
	0x7b, 0xb6, 0x81, 0x9a, 0x50, 0x0f, 0xc6, 0x03, 0x3c, 0xf6, 0x3d, 0xdb, 0x44, 0x5b, 0x00, 0xfe,
	0xc8, 0xf3, 0xbd, 0xe9, 0x8b, 0x41, 0xf0, 0xdc, 0xae, 0xea, 0xcd, 0x89, 0xeb, 0xfa, 0x41, 0x60,
	0xb7, 0x14, 0x38, 0x19, 0x0c, 0xcf, 0x27, 0xd8, 0xb7, 0xb7, 0xd1, 0x2f, 0xd0, 0x1e, 0x8e, 0x4e,
	0xf0, 0x60, 0xba, 0xa2, 0xfe, 0x42, 0x2d, 0x68, 0xb8, 0x83, 0x91, 0xeb, 0xab, 0x73, 0xbd, 0x83,
	0x3d, 0xa8, 0x8f, 0x05, 0x4b, 0x88, 0xb8, 0x45, 0x16, 0x6c, 0x4e, 0x46, 0x81, 0x3f, 0xb6, 0x37,
	0x50, 0x1d, 0x2a, 0xaf, 0xfd, 0xc0, 0x36, 0x50, 0x0d, 0xcc, 0xd1, 0x4b, 0xdb, 0x7c, 0xfa, 0xef,
	0x9b, 0xe3, 0xc7, 0xfd, 0x8c, 0xfe, 0xbf, 0xc3, 0xa4, 0xb3, 0x59, 0x4d, 0x93, 0x47, 0x5f, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x67, 0xe4, 0xe1, 0x47, 0xcb, 0x04, 0x00, 0x00,
}
