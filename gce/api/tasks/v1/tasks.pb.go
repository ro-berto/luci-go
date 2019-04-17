// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/gce/api/tasks/v1/tasks.proto

package tasks

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	v1 "go.chromium.org/luci/gce/api/config/v1"
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

// A task to count the VMs in a config.
type CountVMs struct {
	// The ID of the config whose VMs to count.
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CountVMs) Reset()         { *m = CountVMs{} }
func (m *CountVMs) String() string { return proto.CompactTextString(m) }
func (*CountVMs) ProtoMessage()    {}
func (*CountVMs) Descriptor() ([]byte, []int) {
	return fileDescriptor_f63d8744087b0bbc, []int{0}
}

func (m *CountVMs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CountVMs.Unmarshal(m, b)
}
func (m *CountVMs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CountVMs.Marshal(b, m, deterministic)
}
func (m *CountVMs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CountVMs.Merge(m, src)
}
func (m *CountVMs) XXX_Size() int {
	return xxx_messageInfo_CountVMs.Size(m)
}
func (m *CountVMs) XXX_DiscardUnknown() {
	xxx_messageInfo_CountVMs.DiscardUnknown(m)
}

var xxx_messageInfo_CountVMs proto.InternalMessageInfo

func (m *CountVMs) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// A task to create a GCE instance from a VM.
type CreateInstance struct {
	// The ID of the VM to create a GCE instance from.
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateInstance) Reset()         { *m = CreateInstance{} }
func (m *CreateInstance) String() string { return proto.CompactTextString(m) }
func (*CreateInstance) ProtoMessage()    {}
func (*CreateInstance) Descriptor() ([]byte, []int) {
	return fileDescriptor_f63d8744087b0bbc, []int{1}
}

func (m *CreateInstance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateInstance.Unmarshal(m, b)
}
func (m *CreateInstance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateInstance.Marshal(b, m, deterministic)
}
func (m *CreateInstance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateInstance.Merge(m, src)
}
func (m *CreateInstance) XXX_Size() int {
	return xxx_messageInfo_CreateInstance.Size(m)
}
func (m *CreateInstance) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateInstance.DiscardUnknown(m)
}

var xxx_messageInfo_CreateInstance proto.InternalMessageInfo

func (m *CreateInstance) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// A task to delete a Swarming bot associated with a VM.
type DeleteBot struct {
	// The ID of the VM to delete a Swarming bot for.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The hostname of the Swarming bot to delete.
	Hostname             string   `protobuf:"bytes,2,opt,name=hostname,proto3" json:"hostname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteBot) Reset()         { *m = DeleteBot{} }
func (m *DeleteBot) String() string { return proto.CompactTextString(m) }
func (*DeleteBot) ProtoMessage()    {}
func (*DeleteBot) Descriptor() ([]byte, []int) {
	return fileDescriptor_f63d8744087b0bbc, []int{2}
}

func (m *DeleteBot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteBot.Unmarshal(m, b)
}
func (m *DeleteBot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteBot.Marshal(b, m, deterministic)
}
func (m *DeleteBot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteBot.Merge(m, src)
}
func (m *DeleteBot) XXX_Size() int {
	return xxx_messageInfo_DeleteBot.Size(m)
}
func (m *DeleteBot) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteBot.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteBot proto.InternalMessageInfo

func (m *DeleteBot) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DeleteBot) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

// A task to destroy a GCE instance created from a VM.
type DestroyInstance struct {
	// The ID of the VM to destroy a GCE instance for.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The URL of the GCE instance to destroy.
	Url                  string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DestroyInstance) Reset()         { *m = DestroyInstance{} }
func (m *DestroyInstance) String() string { return proto.CompactTextString(m) }
func (*DestroyInstance) ProtoMessage()    {}
func (*DestroyInstance) Descriptor() ([]byte, []int) {
	return fileDescriptor_f63d8744087b0bbc, []int{3}
}

func (m *DestroyInstance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DestroyInstance.Unmarshal(m, b)
}
func (m *DestroyInstance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DestroyInstance.Marshal(b, m, deterministic)
}
func (m *DestroyInstance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DestroyInstance.Merge(m, src)
}
func (m *DestroyInstance) XXX_Size() int {
	return xxx_messageInfo_DestroyInstance.Size(m)
}
func (m *DestroyInstance) XXX_DiscardUnknown() {
	xxx_messageInfo_DestroyInstance.DiscardUnknown(m)
}

var xxx_messageInfo_DestroyInstance proto.InternalMessageInfo

func (m *DestroyInstance) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DestroyInstance) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

// A task to drain a particular VM.
type DrainVM struct {
	// The ID of the VM to drain.
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DrainVM) Reset()         { *m = DrainVM{} }
func (m *DrainVM) String() string { return proto.CompactTextString(m) }
func (*DrainVM) ProtoMessage()    {}
func (*DrainVM) Descriptor() ([]byte, []int) {
	return fileDescriptor_f63d8744087b0bbc, []int{4}
}

func (m *DrainVM) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DrainVM.Unmarshal(m, b)
}
func (m *DrainVM) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DrainVM.Marshal(b, m, deterministic)
}
func (m *DrainVM) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DrainVM.Merge(m, src)
}
func (m *DrainVM) XXX_Size() int {
	return xxx_messageInfo_DrainVM.Size(m)
}
func (m *DrainVM) XXX_DiscardUnknown() {
	xxx_messageInfo_DrainVM.DiscardUnknown(m)
}

var xxx_messageInfo_DrainVM proto.InternalMessageInfo

func (m *DrainVM) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// A task to create a particular VM.
type CreateVM struct {
	// The index of the VM to create.
	Index int32 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	// The attributes of the VM.
	Attributes *v1.VM `protobuf:"bytes,2,opt,name=attributes,proto3" json:"attributes,omitempty"`
	// The ID of the config this VM belongs to.
	Config string `protobuf:"bytes,3,opt,name=config,proto3" json:"config,omitempty"`
	// The lifetime of the VM in seconds.
	Lifetime int64 `protobuf:"varint,4,opt,name=lifetime,proto3" json:"lifetime,omitempty"`
	// The prefix to use when naming this VM.
	Prefix string `protobuf:"bytes,5,opt,name=prefix,proto3" json:"prefix,omitempty"`
	// The config revision this VM is created from.
	Revision string `protobuf:"bytes,6,opt,name=revision,proto3" json:"revision,omitempty"`
	// The hostname of the Swarming server this VM connects to.
	Swarming string `protobuf:"bytes,7,opt,name=swarming,proto3" json:"swarming,omitempty"`
	// The timeout of the VM in seconds.
	Timeout              int64    `protobuf:"varint,8,opt,name=timeout,proto3" json:"timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateVM) Reset()         { *m = CreateVM{} }
func (m *CreateVM) String() string { return proto.CompactTextString(m) }
func (*CreateVM) ProtoMessage()    {}
func (*CreateVM) Descriptor() ([]byte, []int) {
	return fileDescriptor_f63d8744087b0bbc, []int{5}
}

func (m *CreateVM) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateVM.Unmarshal(m, b)
}
func (m *CreateVM) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateVM.Marshal(b, m, deterministic)
}
func (m *CreateVM) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateVM.Merge(m, src)
}
func (m *CreateVM) XXX_Size() int {
	return xxx_messageInfo_CreateVM.Size(m)
}
func (m *CreateVM) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateVM.DiscardUnknown(m)
}

var xxx_messageInfo_CreateVM proto.InternalMessageInfo

func (m *CreateVM) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *CreateVM) GetAttributes() *v1.VM {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *CreateVM) GetConfig() string {
	if m != nil {
		return m.Config
	}
	return ""
}

func (m *CreateVM) GetLifetime() int64 {
	if m != nil {
		return m.Lifetime
	}
	return 0
}

func (m *CreateVM) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *CreateVM) GetRevision() string {
	if m != nil {
		return m.Revision
	}
	return ""
}

func (m *CreateVM) GetSwarming() string {
	if m != nil {
		return m.Swarming
	}
	return ""
}

func (m *CreateVM) GetTimeout() int64 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

// A task to expand a config.
type ExpandConfig struct {
	// The ID of the config to expand.
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExpandConfig) Reset()         { *m = ExpandConfig{} }
func (m *ExpandConfig) String() string { return proto.CompactTextString(m) }
func (*ExpandConfig) ProtoMessage()    {}
func (*ExpandConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_f63d8744087b0bbc, []int{6}
}

func (m *ExpandConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExpandConfig.Unmarshal(m, b)
}
func (m *ExpandConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExpandConfig.Marshal(b, m, deterministic)
}
func (m *ExpandConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExpandConfig.Merge(m, src)
}
func (m *ExpandConfig) XXX_Size() int {
	return xxx_messageInfo_ExpandConfig.Size(m)
}
func (m *ExpandConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ExpandConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ExpandConfig proto.InternalMessageInfo

func (m *ExpandConfig) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// A task to manage a Swarming bot associated with a VM.
type ManageBot struct {
	// The ID of the VM to manage a Swarming bot for.
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ManageBot) Reset()         { *m = ManageBot{} }
func (m *ManageBot) String() string { return proto.CompactTextString(m) }
func (*ManageBot) ProtoMessage()    {}
func (*ManageBot) Descriptor() ([]byte, []int) {
	return fileDescriptor_f63d8744087b0bbc, []int{7}
}

func (m *ManageBot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ManageBot.Unmarshal(m, b)
}
func (m *ManageBot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ManageBot.Marshal(b, m, deterministic)
}
func (m *ManageBot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManageBot.Merge(m, src)
}
func (m *ManageBot) XXX_Size() int {
	return xxx_messageInfo_ManageBot.Size(m)
}
func (m *ManageBot) XXX_DiscardUnknown() {
	xxx_messageInfo_ManageBot.DiscardUnknown(m)
}

var xxx_messageInfo_ManageBot proto.InternalMessageInfo

func (m *ManageBot) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// A task to report GCE quota utilization.
type ReportQuota struct {
	// The ID of the project to report quota utilization for.
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReportQuota) Reset()         { *m = ReportQuota{} }
func (m *ReportQuota) String() string { return proto.CompactTextString(m) }
func (*ReportQuota) ProtoMessage()    {}
func (*ReportQuota) Descriptor() ([]byte, []int) {
	return fileDescriptor_f63d8744087b0bbc, []int{8}
}

func (m *ReportQuota) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReportQuota.Unmarshal(m, b)
}
func (m *ReportQuota) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReportQuota.Marshal(b, m, deterministic)
}
func (m *ReportQuota) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportQuota.Merge(m, src)
}
func (m *ReportQuota) XXX_Size() int {
	return xxx_messageInfo_ReportQuota.Size(m)
}
func (m *ReportQuota) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportQuota.DiscardUnknown(m)
}

var xxx_messageInfo_ReportQuota proto.InternalMessageInfo

func (m *ReportQuota) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// A task to terminate a Swarming bot associated with a VM.
type TerminateBot struct {
	// The ID of the VM to terminate a Swarming bot for.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The hostname of the Swarming bot to terminate.
	Hostname             string   `protobuf:"bytes,2,opt,name=hostname,proto3" json:"hostname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TerminateBot) Reset()         { *m = TerminateBot{} }
func (m *TerminateBot) String() string { return proto.CompactTextString(m) }
func (*TerminateBot) ProtoMessage()    {}
func (*TerminateBot) Descriptor() ([]byte, []int) {
	return fileDescriptor_f63d8744087b0bbc, []int{9}
}

func (m *TerminateBot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TerminateBot.Unmarshal(m, b)
}
func (m *TerminateBot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TerminateBot.Marshal(b, m, deterministic)
}
func (m *TerminateBot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TerminateBot.Merge(m, src)
}
func (m *TerminateBot) XXX_Size() int {
	return xxx_messageInfo_TerminateBot.Size(m)
}
func (m *TerminateBot) XXX_DiscardUnknown() {
	xxx_messageInfo_TerminateBot.DiscardUnknown(m)
}

var xxx_messageInfo_TerminateBot proto.InternalMessageInfo

func (m *TerminateBot) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TerminateBot) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func init() {
	proto.RegisterType((*CountVMs)(nil), "tasks.CountVMs")
	proto.RegisterType((*CreateInstance)(nil), "tasks.CreateInstance")
	proto.RegisterType((*DeleteBot)(nil), "tasks.DeleteBot")
	proto.RegisterType((*DestroyInstance)(nil), "tasks.DestroyInstance")
	proto.RegisterType((*DrainVM)(nil), "tasks.DrainVM")
	proto.RegisterType((*CreateVM)(nil), "tasks.CreateVM")
	proto.RegisterType((*ExpandConfig)(nil), "tasks.ExpandConfig")
	proto.RegisterType((*ManageBot)(nil), "tasks.ManageBot")
	proto.RegisterType((*ReportQuota)(nil), "tasks.ReportQuota")
	proto.RegisterType((*TerminateBot)(nil), "tasks.TerminateBot")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/gce/api/tasks/v1/tasks.proto", fileDescriptor_f63d8744087b0bbc)
}

var fileDescriptor_f63d8744087b0bbc = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0xab, 0xd3, 0x40,
	0x10, 0xc6, 0x49, 0x6b, 0xda, 0x74, 0xde, 0xe3, 0x29, 0x8b, 0xc8, 0x5a, 0x51, 0x4a, 0x4e, 0x0f,
	0x0f, 0x09, 0xb5, 0x07, 0xc1, 0xa3, 0xad, 0x07, 0x0f, 0x39, 0x18, 0x24, 0xf7, 0x6d, 0x32, 0x4d,
	0x17, 0x93, 0xdd, 0xb0, 0x3b, 0xa9, 0xf5, 0x5f, 0xf6, 0xaf, 0x90, 0xcd, 0xb6, 0x45, 0x5a, 0x15,
	0xbc, 0xcd, 0x2f, 0x5f, 0xbe, 0x99, 0x2f, 0x33, 0x81, 0x65, 0xad, 0x93, 0x72, 0x6f, 0x74, 0x2b,
	0xfb, 0x36, 0xd1, 0xa6, 0x4e, 0x9b, 0xbe, 0x94, 0x69, 0x5d, 0x62, 0x2a, 0x3a, 0x99, 0x92, 0xb0,
	0xdf, 0x6c, 0x7a, 0x58, 0xfa, 0x22, 0xe9, 0x8c, 0x26, 0xcd, 0xc2, 0x01, 0xe6, 0xab, 0x7f, 0x3a,
	0x4b, 0xad, 0x76, 0xb2, 0x76, 0x56, 0x5f, 0x79, 0x6f, 0x3c, 0x87, 0x68, 0xad, 0x7b, 0x45, 0x45,
	0x66, 0xd9, 0x03, 0x8c, 0x64, 0xc5, 0x83, 0x45, 0xf0, 0x38, 0xcb, 0x47, 0xb2, 0x8a, 0x17, 0xf0,
	0xb0, 0x36, 0x28, 0x08, 0x3f, 0x2b, 0x4b, 0x42, 0x95, 0x78, 0xf3, 0xc6, 0x7b, 0x98, 0x6d, 0xb0,
	0x41, 0xc2, 0x8f, 0x9a, 0xae, 0x45, 0x36, 0x87, 0x68, 0xaf, 0x2d, 0x29, 0xd1, 0x22, 0x1f, 0x0d,
	0x4f, 0x2f, 0x1c, 0xaf, 0xe0, 0xe9, 0x06, 0x2d, 0x19, 0xfd, 0xe3, 0x6f, 0xbd, 0xd9, 0x33, 0x18,
	0xf7, 0xa6, 0x39, 0x39, 0x5d, 0x19, 0xbf, 0x84, 0xe9, 0xc6, 0x08, 0xa9, 0x8a, 0xec, 0x26, 0xc8,
	0xcf, 0x00, 0x22, 0x9f, 0xb5, 0xc8, 0xd8, 0x73, 0x08, 0xa5, 0xaa, 0xf0, 0x38, 0xe8, 0x61, 0xee,
	0x81, 0xbd, 0x05, 0x10, 0x44, 0x46, 0x6e, 0x7b, 0x42, 0x3b, 0xb4, 0xbd, 0x7b, 0x07, 0xc9, 0x69,
	0x19, 0x45, 0x96, 0xff, 0xa6, 0xb2, 0x17, 0x30, 0xf1, 0x02, 0x1f, 0x0f, 0x23, 0x4e, 0xe4, 0x3e,
	0xa9, 0x91, 0x3b, 0x24, 0xd9, 0x22, 0x7f, 0xb2, 0x08, 0x1e, 0xc7, 0xf9, 0x85, 0x9d, 0xa7, 0x33,
	0xb8, 0x93, 0x47, 0x1e, 0x7a, 0x8f, 0x27, 0xe7, 0x31, 0x78, 0x90, 0x56, 0x6a, 0xc5, 0x27, 0x7e,
	0x0d, 0x67, 0x76, 0x9a, 0xfd, 0x2e, 0x4c, 0x2b, 0x55, 0xcd, 0xa7, 0x5e, 0x3b, 0x33, 0xe3, 0x30,
	0x75, 0x7d, 0x75, 0x4f, 0x3c, 0x1a, 0x46, 0x9d, 0x31, 0x7e, 0x03, 0xf7, 0x9f, 0x8e, 0x9d, 0x50,
	0xd5, 0xda, 0xa7, 0xba, 0x5e, 0xc6, 0x2b, 0x98, 0x65, 0x42, 0x89, 0xfa, 0x4f, 0x57, 0x89, 0x5f,
	0xc3, 0x5d, 0x8e, 0x9d, 0x36, 0xf4, 0xa5, 0xd7, 0x24, 0x6e, 0xe4, 0x0f, 0x70, 0xff, 0x15, 0x5d,
	0x00, 0xf1, 0xdf, 0x47, 0xdd, 0x4e, 0x86, 0x5f, 0x6a, 0xf5, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xbf,
	0x75, 0x1a, 0x65, 0xc3, 0x02, 0x00, 0x00,
}
