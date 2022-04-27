// Copyright 2015 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: go.chromium.org/luci/grpc/prpc/talk/buildbot/proto/buildbot.proto

package buildbot

import prpc "go.chromium.org/luci/grpc/prpc"

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BuildState int32

const (
	BuildState_UNSET     BuildState = 0
	BuildState_PENDING   BuildState = 1
	BuildState_RUNNING   BuildState = 2
	BuildState_SUCCESS   BuildState = 3
	BuildState_FAILURE   BuildState = 4
	BuildState_EXCEPTION BuildState = 5
)

// Enum value maps for BuildState.
var (
	BuildState_name = map[int32]string{
		0: "UNSET",
		1: "PENDING",
		2: "RUNNING",
		3: "SUCCESS",
		4: "FAILURE",
		5: "EXCEPTION",
	}
	BuildState_value = map[string]int32{
		"UNSET":     0,
		"PENDING":   1,
		"RUNNING":   2,
		"SUCCESS":   3,
		"FAILURE":   4,
		"EXCEPTION": 5,
	}
)

func (x BuildState) Enum() *BuildState {
	p := new(BuildState)
	*p = x
	return p
}

func (x BuildState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BuildState) Descriptor() protoreflect.EnumDescriptor {
	return file_go_chromium_org_luci_grpc_prpc_talk_buildbot_proto_buildbot_proto_enumTypes[0].Descriptor()
}

func (BuildState) Type() protoreflect.EnumType {
	return &file_go_chromium_org_luci_grpc_prpc_talk_buildbot_proto_buildbot_proto_enumTypes[0]
}

func (x BuildState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BuildState.Descriptor instead.
func (BuildState) EnumDescriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_grpc_prpc_talk_buildbot_proto_buildbot_proto_rawDescGZIP(), []int{0}
}

// SearchReqeust specifies a search criteria.
type SearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Master filters by master name, e.g. "master.XXX".
	Master string `protobuf:"bytes,1,opt,name=master,proto3" json:"master,omitempty"`
	// State filters by build state.
	State BuildState `protobuf:"varint,2,opt,name=state,proto3,enum=buildbot.BuildState" json:"state,omitempty"`
	// Builder filters by builder name.
	Builder string `protobuf:"bytes,3,opt,name=builder,proto3" json:"builder,omitempty"`
}

func (x *SearchRequest) Reset() {
	*x = SearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_grpc_prpc_talk_buildbot_proto_buildbot_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRequest) ProtoMessage() {}

func (x *SearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_grpc_prpc_talk_buildbot_proto_buildbot_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRequest.ProtoReflect.Descriptor instead.
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_grpc_prpc_talk_buildbot_proto_buildbot_proto_rawDescGZIP(), []int{0}
}

func (x *SearchRequest) GetMaster() string {
	if x != nil {
		return x.Master
	}
	return ""
}

func (x *SearchRequest) GetState() BuildState {
	if x != nil {
		return x.State
	}
	return BuildState_UNSET
}

func (x *SearchRequest) GetBuilder() string {
	if x != nil {
		return x.Builder
	}
	return ""
}

type SearchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Builds []*Build `protobuf:"bytes,1,rep,name=builds,proto3" json:"builds,omitempty"`
}

func (x *SearchResponse) Reset() {
	*x = SearchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_grpc_prpc_talk_buildbot_proto_buildbot_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResponse) ProtoMessage() {}

func (x *SearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_grpc_prpc_talk_buildbot_proto_buildbot_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResponse.ProtoReflect.Descriptor instead.
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_grpc_prpc_talk_buildbot_proto_buildbot_proto_rawDescGZIP(), []int{1}
}

func (x *SearchResponse) GetBuilds() []*Build {
	if x != nil {
		return x.Builds
	}
	return nil
}

type Build struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Master  string     `protobuf:"bytes,1,opt,name=master,proto3" json:"master,omitempty"`
	Builder string     `protobuf:"bytes,2,opt,name=builder,proto3" json:"builder,omitempty"`
	Number  int32      `protobuf:"varint,3,opt,name=number,proto3" json:"number,omitempty"`
	State   BuildState `protobuf:"varint,4,opt,name=state,proto3,enum=buildbot.BuildState" json:"state,omitempty"`
}

func (x *Build) Reset() {
	*x = Build{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_grpc_prpc_talk_buildbot_proto_buildbot_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Build) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Build) ProtoMessage() {}

func (x *Build) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_grpc_prpc_talk_buildbot_proto_buildbot_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Build.ProtoReflect.Descriptor instead.
func (*Build) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_grpc_prpc_talk_buildbot_proto_buildbot_proto_rawDescGZIP(), []int{2}
}

func (x *Build) GetMaster() string {
	if x != nil {
		return x.Master
	}
	return ""
}

func (x *Build) GetBuilder() string {
	if x != nil {
		return x.Builder
	}
	return ""
}

func (x *Build) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *Build) GetState() BuildState {
	if x != nil {
		return x.State
	}
	return BuildState_UNSET
}

// ScheduleRequest defines builds to schedule.
type ScheduleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Master is a "master.XXX" string that defines where to schedule builds.
	Master string `protobuf:"bytes,1,opt,name=master,proto3" json:"master,omitempty"`
	// Builds is a list of builds to schedule.
	Builds []*ScheduleRequest_BuildDef `protobuf:"bytes,2,rep,name=builds,proto3" json:"builds,omitempty"`
}

func (x *ScheduleRequest) Reset() {
	*x = ScheduleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_grpc_prpc_talk_buildbot_proto_buildbot_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScheduleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleRequest) ProtoMessage() {}

func (x *ScheduleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_grpc_prpc_talk_buildbot_proto_buildbot_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleRequest.ProtoReflect.Descriptor instead.
func (*ScheduleRequest) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_grpc_prpc_talk_buildbot_proto_buildbot_proto_rawDescGZIP(), []int{3}
}

func (x *ScheduleRequest) GetMaster() string {
	if x != nil {
		return x.Master
	}
	return ""
}

func (x *ScheduleRequest) GetBuilds() []*ScheduleRequest_BuildDef {
	if x != nil {
		return x.Builds
	}
	return nil
}

// HelloReply contains a greeting.
type ScheduleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Builds []*Build `protobuf:"bytes,1,rep,name=builds,proto3" json:"builds,omitempty"`
}

func (x *ScheduleResponse) Reset() {
	*x = ScheduleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_grpc_prpc_talk_buildbot_proto_buildbot_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScheduleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleResponse) ProtoMessage() {}

func (x *ScheduleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_grpc_prpc_talk_buildbot_proto_buildbot_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleResponse.ProtoReflect.Descriptor instead.
func (*ScheduleResponse) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_grpc_prpc_talk_buildbot_proto_buildbot_proto_rawDescGZIP(), []int{4}
}

func (x *ScheduleResponse) GetBuilds() []*Build {
	if x != nil {
		return x.Builds
	}
	return nil
}

// Build is a build to schedule.
type ScheduleRequest_BuildDef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Builder defines the build script.
	Builder string `protobuf:"bytes,1,opt,name=builder,proto3" json:"builder,omitempty"`
	// Branch defines what to fetch.
	Branch string `protobuf:"bytes,2,opt,name=branch,proto3" json:"branch,omitempty"`
	// Revision is a commit hash to checkout
	Revision string `protobuf:"bytes,3,opt,name=revision,proto3" json:"revision,omitempty"`
	// Properties are "key:value" pairs.
	Properties []string `protobuf:"bytes,4,rep,name=properties,proto3" json:"properties,omitempty"`
	// Blamelist is a list of user email addressed to blame if this build
	// fails.
	Blamelist []string `protobuf:"bytes,5,rep,name=blamelist,proto3" json:"blamelist,omitempty"`
}

func (x *ScheduleRequest_BuildDef) Reset() {
	*x = ScheduleRequest_BuildDef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_grpc_prpc_talk_buildbot_proto_buildbot_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScheduleRequest_BuildDef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleRequest_BuildDef) ProtoMessage() {}

func (x *ScheduleRequest_BuildDef) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_grpc_prpc_talk_buildbot_proto_buildbot_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleRequest_BuildDef.ProtoReflect.Descriptor instead.
func (*ScheduleRequest_BuildDef) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_grpc_prpc_talk_buildbot_proto_buildbot_proto_rawDescGZIP(), []int{3, 0}
}

func (x *ScheduleRequest_BuildDef) GetBuilder() string {
	if x != nil {
		return x.Builder
	}
	return ""
}

func (x *ScheduleRequest_BuildDef) GetBranch() string {
	if x != nil {
		return x.Branch
	}
	return ""
}

func (x *ScheduleRequest_BuildDef) GetRevision() string {
	if x != nil {
		return x.Revision
	}
	return ""
}

func (x *ScheduleRequest_BuildDef) GetProperties() []string {
	if x != nil {
		return x.Properties
	}
	return nil
}

func (x *ScheduleRequest_BuildDef) GetBlamelist() []string {
	if x != nil {
		return x.Blamelist
	}
	return nil
}

var File_go_chromium_org_luci_grpc_prpc_talk_buildbot_proto_buildbot_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_grpc_prpc_talk_buildbot_proto_buildbot_proto_rawDesc = []byte{
	0x0a, 0x41, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x70, 0x63,
	0x2f, 0x74, 0x61, 0x6c, 0x6b, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x6f, 0x74, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x6f, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x6f, 0x74, 0x22, 0x6d, 0x0a,
	0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x6f, 0x74,
	0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x22, 0x39, 0x0a, 0x0e,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27,
	0x0a, 0x06, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x6f, 0x74, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52,
	0x06, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x22, 0x7d, 0x0a, 0x05, 0x42, 0x75, 0x69, 0x6c, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x62, 0x6f, 0x74, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0xfe, 0x01, 0x0a, 0x0f, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x61,
	0x73, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x61, 0x73, 0x74,
	0x65, 0x72, 0x12, 0x3a, 0x0a, 0x06, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x6f, 0x74, 0x2e, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x42, 0x75,
	0x69, 0x6c, 0x64, 0x44, 0x65, 0x66, 0x52, 0x06, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x1a, 0x96,
	0x01, 0x0a, 0x08, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x44, 0x65, 0x66, 0x12, 0x18, 0x0a, 0x07, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x1a, 0x0a,
	0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x70,
	0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x6c, 0x61,
	0x6d, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x62, 0x6c,
	0x61, 0x6d, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x3b, 0x0a, 0x10, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x62, 0x6f, 0x74, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x06, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x73, 0x2a, 0x5a, 0x0a, 0x0a, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x55, 0x4e, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x0b, 0x0a,
	0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x55,
	0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45,
	0x53, 0x53, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x10,
	0x04, 0x12, 0x0d, 0x0a, 0x09, 0x45, 0x58, 0x43, 0x45, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x05,
	0x32, 0x8e, 0x01, 0x0a, 0x08, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x6f, 0x74, 0x12, 0x3d, 0x0a,
	0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x17, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62,
	0x6f, 0x74, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x6f, 0x74, 0x2e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x08,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x19, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x62, 0x6f, 0x74, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x6f, 0x74, 0x2e, 0x53,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d,
	0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70,
	0x72, 0x70, 0x63, 0x2f, 0x74, 0x61, 0x6c, 0x6b, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x6f,
	0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x6f, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_grpc_prpc_talk_buildbot_proto_buildbot_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_grpc_prpc_talk_buildbot_proto_buildbot_proto_rawDescData = file_go_chromium_org_luci_grpc_prpc_talk_buildbot_proto_buildbot_proto_rawDesc
)

func file_go_chromium_org_luci_grpc_prpc_talk_buildbot_proto_buildbot_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_grpc_prpc_talk_buildbot_proto_buildbot_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_grpc_prpc_talk_buildbot_proto_buildbot_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_grpc_prpc_talk_buildbot_proto_buildbot_proto_rawDescData)
	})
	return file_go_chromium_org_luci_grpc_prpc_talk_buildbot_proto_buildbot_proto_rawDescData
}

var file_go_chromium_org_luci_grpc_prpc_talk_buildbot_proto_buildbot_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_go_chromium_org_luci_grpc_prpc_talk_buildbot_proto_buildbot_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_go_chromium_org_luci_grpc_prpc_talk_buildbot_proto_buildbot_proto_goTypes = []interface{}{
	(BuildState)(0),                  // 0: buildbot.BuildState
	(*SearchRequest)(nil),            // 1: buildbot.SearchRequest
	(*SearchResponse)(nil),           // 2: buildbot.SearchResponse
	(*Build)(nil),                    // 3: buildbot.Build
	(*ScheduleRequest)(nil),          // 4: buildbot.ScheduleRequest
	(*ScheduleResponse)(nil),         // 5: buildbot.ScheduleResponse
	(*ScheduleRequest_BuildDef)(nil), // 6: buildbot.ScheduleRequest.BuildDef
}
var file_go_chromium_org_luci_grpc_prpc_talk_buildbot_proto_buildbot_proto_depIdxs = []int32{
	0, // 0: buildbot.SearchRequest.state:type_name -> buildbot.BuildState
	3, // 1: buildbot.SearchResponse.builds:type_name -> buildbot.Build
	0, // 2: buildbot.Build.state:type_name -> buildbot.BuildState
	6, // 3: buildbot.ScheduleRequest.builds:type_name -> buildbot.ScheduleRequest.BuildDef
	3, // 4: buildbot.ScheduleResponse.builds:type_name -> buildbot.Build
	1, // 5: buildbot.Buildbot.Search:input_type -> buildbot.SearchRequest
	4, // 6: buildbot.Buildbot.Schedule:input_type -> buildbot.ScheduleRequest
	2, // 7: buildbot.Buildbot.Search:output_type -> buildbot.SearchResponse
	5, // 8: buildbot.Buildbot.Schedule:output_type -> buildbot.ScheduleResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_grpc_prpc_talk_buildbot_proto_buildbot_proto_init() }
func file_go_chromium_org_luci_grpc_prpc_talk_buildbot_proto_buildbot_proto_init() {
	if File_go_chromium_org_luci_grpc_prpc_talk_buildbot_proto_buildbot_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_grpc_prpc_talk_buildbot_proto_buildbot_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchRequest); i {
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
		file_go_chromium_org_luci_grpc_prpc_talk_buildbot_proto_buildbot_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchResponse); i {
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
		file_go_chromium_org_luci_grpc_prpc_talk_buildbot_proto_buildbot_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Build); i {
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
		file_go_chromium_org_luci_grpc_prpc_talk_buildbot_proto_buildbot_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScheduleRequest); i {
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
		file_go_chromium_org_luci_grpc_prpc_talk_buildbot_proto_buildbot_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScheduleResponse); i {
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
		file_go_chromium_org_luci_grpc_prpc_talk_buildbot_proto_buildbot_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScheduleRequest_BuildDef); i {
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
			RawDescriptor: file_go_chromium_org_luci_grpc_prpc_talk_buildbot_proto_buildbot_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_go_chromium_org_luci_grpc_prpc_talk_buildbot_proto_buildbot_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_grpc_prpc_talk_buildbot_proto_buildbot_proto_depIdxs,
		EnumInfos:         file_go_chromium_org_luci_grpc_prpc_talk_buildbot_proto_buildbot_proto_enumTypes,
		MessageInfos:      file_go_chromium_org_luci_grpc_prpc_talk_buildbot_proto_buildbot_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_grpc_prpc_talk_buildbot_proto_buildbot_proto = out.File
	file_go_chromium_org_luci_grpc_prpc_talk_buildbot_proto_buildbot_proto_rawDesc = nil
	file_go_chromium_org_luci_grpc_prpc_talk_buildbot_proto_buildbot_proto_goTypes = nil
	file_go_chromium_org_luci_grpc_prpc_talk_buildbot_proto_buildbot_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BuildbotClient is the client API for Buildbot service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BuildbotClient interface {
	// Search returns builds matching a criteria.
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
	// Schedule puts new builds to a queue.
	Schedule(ctx context.Context, in *ScheduleRequest, opts ...grpc.CallOption) (*ScheduleResponse, error)
}
type buildbotPRPCClient struct {
	client *prpc.Client
}

func NewBuildbotPRPCClient(client *prpc.Client) BuildbotClient {
	return &buildbotPRPCClient{client}
}

func (c *buildbotPRPCClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.client.Call(ctx, "buildbot.Buildbot", "Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildbotPRPCClient) Schedule(ctx context.Context, in *ScheduleRequest, opts ...grpc.CallOption) (*ScheduleResponse, error) {
	out := new(ScheduleResponse)
	err := c.client.Call(ctx, "buildbot.Buildbot", "Schedule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type buildbotClient struct {
	cc grpc.ClientConnInterface
}

func NewBuildbotClient(cc grpc.ClientConnInterface) BuildbotClient {
	return &buildbotClient{cc}
}

func (c *buildbotClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/buildbot.Buildbot/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildbotClient) Schedule(ctx context.Context, in *ScheduleRequest, opts ...grpc.CallOption) (*ScheduleResponse, error) {
	out := new(ScheduleResponse)
	err := c.cc.Invoke(ctx, "/buildbot.Buildbot/Schedule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BuildbotServer is the server API for Buildbot service.
type BuildbotServer interface {
	// Search returns builds matching a criteria.
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
	// Schedule puts new builds to a queue.
	Schedule(context.Context, *ScheduleRequest) (*ScheduleResponse, error)
}

// UnimplementedBuildbotServer can be embedded to have forward compatible implementations.
type UnimplementedBuildbotServer struct {
}

func (*UnimplementedBuildbotServer) Search(context.Context, *SearchRequest) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (*UnimplementedBuildbotServer) Schedule(context.Context, *ScheduleRequest) (*ScheduleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Schedule not implemented")
}

func RegisterBuildbotServer(s prpc.Registrar, srv BuildbotServer) {
	s.RegisterService(&_Buildbot_serviceDesc, srv)
}

func _Buildbot_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildbotServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buildbot.Buildbot/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildbotServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Buildbot_Schedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildbotServer).Schedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buildbot.Buildbot/Schedule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildbotServer).Schedule(ctx, req.(*ScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Buildbot_serviceDesc = grpc.ServiceDesc{
	ServiceName: "buildbot.Buildbot",
	HandlerType: (*BuildbotServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _Buildbot_Search_Handler,
		},
		{
			MethodName: "Schedule",
			Handler:    _Buildbot_Schedule_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go.chromium.org/luci/grpc/prpc/talk/buildbot/proto/buildbot.proto",
}
