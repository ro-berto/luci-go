// Copyright 2018 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.1
// source: go.chromium.org/luci/gce/api/tasks/v1/tasks.proto

package tasks

import (
	v1 "go.chromium.org/luci/gce/api/config/v1"
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

// A task to count the VMs in a config.
type CountVMs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the config whose VMs to count.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CountVMs) Reset() {
	*x = CountVMs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountVMs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountVMs) ProtoMessage() {}

func (x *CountVMs) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountVMs.ProtoReflect.Descriptor instead.
func (*CountVMs) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto_rawDescGZIP(), []int{0}
}

func (x *CountVMs) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// A task to create a GCE instance from a VM.
type CreateInstance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the VM to create a GCE instance from.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateInstance) Reset() {
	*x = CreateInstance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateInstance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateInstance) ProtoMessage() {}

func (x *CreateInstance) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateInstance.ProtoReflect.Descriptor instead.
func (*CreateInstance) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto_rawDescGZIP(), []int{1}
}

func (x *CreateInstance) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// A task to delete a Swarming bot associated with a VM.
type DeleteBot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the VM to delete a Swarming bot for.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The hostname of the Swarming bot to delete.
	Hostname string `protobuf:"bytes,2,opt,name=hostname,proto3" json:"hostname,omitempty"`
}

func (x *DeleteBot) Reset() {
	*x = DeleteBot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBot) ProtoMessage() {}

func (x *DeleteBot) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBot.ProtoReflect.Descriptor instead.
func (*DeleteBot) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteBot) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeleteBot) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

// A task to destroy a GCE instance created from a VM.
type DestroyInstance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the VM to destroy a GCE instance for.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The URL of the GCE instance to destroy.
	Url string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *DestroyInstance) Reset() {
	*x = DestroyInstance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DestroyInstance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DestroyInstance) ProtoMessage() {}

func (x *DestroyInstance) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DestroyInstance.ProtoReflect.Descriptor instead.
func (*DestroyInstance) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto_rawDescGZIP(), []int{3}
}

func (x *DestroyInstance) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DestroyInstance) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

// A task to create a particular VM.
type CreateVM struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the VM to create.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The attributes of the VM.
	Attributes *v1.VM `protobuf:"bytes,2,opt,name=attributes,proto3" json:"attributes,omitempty"`
	// The ID of the config this VM belongs to.
	Config string `protobuf:"bytes,3,opt,name=config,proto3" json:"config,omitempty"`
	// The timestamp when this task was created.
	Created *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created,proto3" json:"created,omitempty"`
	// The index of the VM to create.
	Index int32 `protobuf:"varint,5,opt,name=index,proto3" json:"index,omitempty"`
	// The lifetime of the VM in seconds.
	Lifetime int64 `protobuf:"varint,6,opt,name=lifetime,proto3" json:"lifetime,omitempty"`
	// The prefix to use when naming this VM.
	Prefix string `protobuf:"bytes,7,opt,name=prefix,proto3" json:"prefix,omitempty"`
	// The config revision this VM is created from.
	Revision string `protobuf:"bytes,8,opt,name=revision,proto3" json:"revision,omitempty"`
	// The hostname of the Swarming server this VM connects to.
	Swarming string `protobuf:"bytes,9,opt,name=swarming,proto3" json:"swarming,omitempty"`
	// The timeout of the VM in seconds.
	Timeout int64 `protobuf:"varint,10,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *CreateVM) Reset() {
	*x = CreateVM{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateVM) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVM) ProtoMessage() {}

func (x *CreateVM) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVM.ProtoReflect.Descriptor instead.
func (*CreateVM) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto_rawDescGZIP(), []int{4}
}

func (x *CreateVM) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateVM) GetAttributes() *v1.VM {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *CreateVM) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

func (x *CreateVM) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *CreateVM) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *CreateVM) GetLifetime() int64 {
	if x != nil {
		return x.Lifetime
	}
	return 0
}

func (x *CreateVM) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

func (x *CreateVM) GetRevision() string {
	if x != nil {
		return x.Revision
	}
	return ""
}

func (x *CreateVM) GetSwarming() string {
	if x != nil {
		return x.Swarming
	}
	return ""
}

func (x *CreateVM) GetTimeout() int64 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

// A task to expand a config.
type ExpandConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the config to expand.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ExpandConfig) Reset() {
	*x = ExpandConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpandConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpandConfig) ProtoMessage() {}

func (x *ExpandConfig) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpandConfig.ProtoReflect.Descriptor instead.
func (*ExpandConfig) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto_rawDescGZIP(), []int{5}
}

func (x *ExpandConfig) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// A task to manage a Swarming bot associated with a VM.
type ManageBot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the VM to manage a Swarming bot for.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ManageBot) Reset() {
	*x = ManageBot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManageBot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManageBot) ProtoMessage() {}

func (x *ManageBot) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManageBot.ProtoReflect.Descriptor instead.
func (*ManageBot) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto_rawDescGZIP(), []int{6}
}

func (x *ManageBot) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// A task to report GCE quota utilization.
type ReportQuota struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the project to report quota utilization for.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ReportQuota) Reset() {
	*x = ReportQuota{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportQuota) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportQuota) ProtoMessage() {}

func (x *ReportQuota) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportQuota.ProtoReflect.Descriptor instead.
func (*ReportQuota) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto_rawDescGZIP(), []int{7}
}

func (x *ReportQuota) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// A task to terminate a Swarming bot associated with a VM.
type TerminateBot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the VM to terminate a Swarming bot for.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The hostname of the Swarming bot to terminate.
	Hostname string `protobuf:"bytes,2,opt,name=hostname,proto3" json:"hostname,omitempty"`
}

func (x *TerminateBot) Reset() {
	*x = TerminateBot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TerminateBot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TerminateBot) ProtoMessage() {}

func (x *TerminateBot) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TerminateBot.ProtoReflect.Descriptor instead.
func (*TerminateBot) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto_rawDescGZIP(), []int{8}
}

func (x *TerminateBot) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TerminateBot) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

var File_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto_rawDesc = []byte{
	0x0a, 0x31, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x67, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74,
	0x61, 0x73, 0x6b, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x33, 0x67, 0x6f, 0x2e,
	0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63,
	0x69, 0x2f, 0x67, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x56, 0x4d, 0x73, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x20, 0x0a, 0x0e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x37,
	0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x68,
	0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68,
	0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x33, 0x0a, 0x0f, 0x44, 0x65, 0x73, 0x74, 0x72,
	0x6f, 0x79, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0xb0, 0x02, 0x0a,
	0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x4d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2a, 0x0a, 0x0a, 0x61, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x56, 0x4d, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x34, 0x0a,
	0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x69, 0x66,
	0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6c, 0x69, 0x66,
	0x65, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x1a, 0x0a,
	0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x77, 0x61,
	0x72, 0x6d, 0x69, 0x6e, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x77, 0x61,
	0x72, 0x6d, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22,
	0x1e, 0x0a, 0x0c, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x1b, 0x0a, 0x09, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x42, 0x6f, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1d, 0x0a, 0x0b,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3a, 0x0a, 0x0c, 0x54,
	0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x68,
	0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68,
	0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x6f, 0x2e, 0x63, 0x68,
	0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f,
	0x67, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2f, 0x76, 0x31,
	0x3b, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto_rawDescData = file_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto_rawDesc
)

func file_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto_rawDescData)
	})
	return file_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto_rawDescData
}

var file_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto_goTypes = []interface{}{
	(*CountVMs)(nil),              // 0: tasks.CountVMs
	(*CreateInstance)(nil),        // 1: tasks.CreateInstance
	(*DeleteBot)(nil),             // 2: tasks.DeleteBot
	(*DestroyInstance)(nil),       // 3: tasks.DestroyInstance
	(*CreateVM)(nil),              // 4: tasks.CreateVM
	(*ExpandConfig)(nil),          // 5: tasks.ExpandConfig
	(*ManageBot)(nil),             // 6: tasks.ManageBot
	(*ReportQuota)(nil),           // 7: tasks.ReportQuota
	(*TerminateBot)(nil),          // 8: tasks.TerminateBot
	(*v1.VM)(nil),                 // 9: config.VM
	(*timestamppb.Timestamp)(nil), // 10: google.protobuf.Timestamp
}
var file_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto_depIdxs = []int32{
	9,  // 0: tasks.CreateVM.attributes:type_name -> config.VM
	10, // 1: tasks.CreateVM.created:type_name -> google.protobuf.Timestamp
	2,  // [2:2] is the sub-list for method output_type
	2,  // [2:2] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto_init() }
func file_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto_init() {
	if File_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountVMs); i {
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
		file_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateInstance); i {
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
		file_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBot); i {
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
		file_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DestroyInstance); i {
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
		file_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateVM); i {
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
		file_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpandConfig); i {
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
		file_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManageBot); i {
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
		file_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportQuota); i {
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
		file_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TerminateBot); i {
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
			RawDescriptor: file_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto = out.File
	file_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto_rawDesc = nil
	file_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto_goTypes = nil
	file_go_chromium_org_luci_gce_api_tasks_v1_tasks_proto_depIdxs = nil
}
