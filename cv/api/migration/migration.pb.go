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
// source: go.chromium.org/luci/cv/api/migration/migration.proto

package migrationpb

import (
	v1 "go.chromium.org/luci/cv/api/bigquery/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type ReportRunsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Runs []*Run `protobuf:"bytes,1,rep,name=runs,proto3" json:"runs,omitempty"`
}

func (x *ReportRunsRequest) Reset() {
	*x = ReportRunsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_api_migration_migration_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportRunsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportRunsRequest) ProtoMessage() {}

func (x *ReportRunsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_api_migration_migration_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportRunsRequest.ProtoReflect.Descriptor instead.
func (*ReportRunsRequest) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_api_migration_migration_proto_rawDescGZIP(), []int{0}
}

func (x *ReportRunsRequest) GetRuns() []*Run {
	if x != nil {
		return x.Runs
	}
	return nil
}

type ReportFinishedRunRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Run *Run `protobuf:"bytes,1,opt,name=run,proto3" json:"run,omitempty"`
}

func (x *ReportFinishedRunRequest) Reset() {
	*x = ReportFinishedRunRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_api_migration_migration_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportFinishedRunRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportFinishedRunRequest) ProtoMessage() {}

func (x *ReportFinishedRunRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_api_migration_migration_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportFinishedRunRequest.ProtoReflect.Descriptor instead.
func (*ReportFinishedRunRequest) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_api_migration_migration_proto_rawDescGZIP(), []int{1}
}

func (x *ReportFinishedRunRequest) GetRun() *Run {
	if x != nil {
		return x.Run
	}
	return nil
}

type ReportUsedNetrcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GerritHost  string `protobuf:"bytes,1,opt,name=gerrit_host,json=gerritHost,proto3" json:"gerrit_host,omitempty"`
	AccessToken string `protobuf:"bytes,2,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
}

func (x *ReportUsedNetrcRequest) Reset() {
	*x = ReportUsedNetrcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_api_migration_migration_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportUsedNetrcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportUsedNetrcRequest) ProtoMessage() {}

func (x *ReportUsedNetrcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_api_migration_migration_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportUsedNetrcRequest.ProtoReflect.Descriptor instead.
func (*ReportUsedNetrcRequest) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_api_migration_migration_proto_rawDescGZIP(), []int{2}
}

func (x *ReportUsedNetrcRequest) GetGerritHost() string {
	if x != nil {
		return x.GerritHost
	}
	return ""
}

func (x *ReportUsedNetrcRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

type FetchActiveRunsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LuciProject string `protobuf:"bytes,1,opt,name=luci_project,json=luciProject,proto3" json:"luci_project,omitempty"`
}

func (x *FetchActiveRunsRequest) Reset() {
	*x = FetchActiveRunsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_api_migration_migration_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchActiveRunsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchActiveRunsRequest) ProtoMessage() {}

func (x *FetchActiveRunsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_api_migration_migration_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchActiveRunsRequest.ProtoReflect.Descriptor instead.
func (*FetchActiveRunsRequest) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_api_migration_migration_proto_rawDescGZIP(), []int{3}
}

func (x *FetchActiveRunsRequest) GetLuciProject() string {
	if x != nil {
		return x.LuciProject
	}
	return ""
}

type FetchActiveRunsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Runs []*Run `protobuf:"bytes,1,rep,name=runs,proto3" json:"runs,omitempty"`
}

func (x *FetchActiveRunsResponse) Reset() {
	*x = FetchActiveRunsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_api_migration_migration_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchActiveRunsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchActiveRunsResponse) ProtoMessage() {}

func (x *FetchActiveRunsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_api_migration_migration_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchActiveRunsResponse.ProtoReflect.Descriptor instead.
func (*FetchActiveRunsResponse) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_api_migration_migration_proto_rawDescGZIP(), []int{4}
}

func (x *FetchActiveRunsResponse) GetRuns() []*Run {
	if x != nil {
		return x.Runs
	}
	return nil
}

type Run struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Attempt *v1.Attempt `protobuf:"bytes,1,opt,name=attempt,proto3" json:"attempt,omitempty"`
	// CV's run ID.
	Id  string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Cls []*RunCL `protobuf:"bytes,3,rep,name=cls,proto3" json:"cls,omitempty"`
}

func (x *Run) Reset() {
	*x = Run{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_api_migration_migration_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Run) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Run) ProtoMessage() {}

func (x *Run) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_api_migration_migration_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Run.ProtoReflect.Descriptor instead.
func (*Run) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_api_migration_migration_proto_rawDescGZIP(), []int{5}
}

func (x *Run) GetAttempt() *v1.Attempt {
	if x != nil {
		return x.Attempt
	}
	return nil
}

func (x *Run) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Run) GetCls() []*RunCL {
	if x != nil {
		return x.Cls
	}
	return nil
}

type RunCL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// CV's CLID. Used to identify Deps.
	Id int64            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Gc *v1.GerritChange `protobuf:"bytes,2,opt,name=gc,proto3" json:"gc,omitempty"`
	// Hint to CQDaemon to avoid needless re-fetching.
	UpdatedTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=updated_time,json=updatedTime,proto3" json:"updated_time,omitempty"`
	Trigger     *RunCL_Trigger         `protobuf:"bytes,4,opt,name=trigger,proto3" json:"trigger,omitempty"`
	Deps        []*RunCL_Dep           `protobuf:"bytes,5,rep,name=deps,proto3" json:"deps,omitempty"`
}

func (x *RunCL) Reset() {
	*x = RunCL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_api_migration_migration_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunCL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunCL) ProtoMessage() {}

func (x *RunCL) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_api_migration_migration_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunCL.ProtoReflect.Descriptor instead.
func (*RunCL) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_api_migration_migration_proto_rawDescGZIP(), []int{6}
}

func (x *RunCL) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RunCL) GetGc() *v1.GerritChange {
	if x != nil {
		return x.Gc
	}
	return nil
}

func (x *RunCL) GetUpdatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedTime
	}
	return nil
}

func (x *RunCL) GetTrigger() *RunCL_Trigger {
	if x != nil {
		return x.Trigger
	}
	return nil
}

func (x *RunCL) GetDeps() []*RunCL_Dep {
	if x != nil {
		return x.Deps
	}
	return nil
}

type RunCL_Trigger struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	// Gerrit account ID.
	AccountId int64 `protobuf:"varint,3,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	// User email, if known.
	Email string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *RunCL_Trigger) Reset() {
	*x = RunCL_Trigger{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_api_migration_migration_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunCL_Trigger) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunCL_Trigger) ProtoMessage() {}

func (x *RunCL_Trigger) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_api_migration_migration_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunCL_Trigger.ProtoReflect.Descriptor instead.
func (*RunCL_Trigger) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_api_migration_migration_proto_rawDescGZIP(), []int{6, 0}
}

func (x *RunCL_Trigger) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *RunCL_Trigger) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *RunCL_Trigger) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type RunCL_Dep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// CV's CLID. Guaranteed to match one of the RunCL in the same Run.
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// True means Dep is the immediate git parent and must be submitted first.
	Hard bool `protobuf:"varint,2,opt,name=hard,proto3" json:"hard,omitempty"`
}

func (x *RunCL_Dep) Reset() {
	*x = RunCL_Dep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_api_migration_migration_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunCL_Dep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunCL_Dep) ProtoMessage() {}

func (x *RunCL_Dep) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_api_migration_migration_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunCL_Dep.ProtoReflect.Descriptor instead.
func (*RunCL_Dep) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_api_migration_migration_proto_rawDescGZIP(), []int{6, 1}
}

func (x *RunCL_Dep) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RunCL_Dep) GetHard() bool {
	if x != nil {
		return x.Hard
	}
	return false
}

var File_go_chromium_org_luci_cv_api_migration_migration_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_cv_api_migration_migration_proto_rawDesc = []byte{
	0x0a, 0x35, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x76, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x69,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x35, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x76, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x69,
	0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x37, 0x0a, 0x11, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x52, 0x75, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x04,
	0x72, 0x75, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x69, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x75, 0x6e, 0x52, 0x04, 0x72, 0x75, 0x6e, 0x73,
	0x22, 0x3c, 0x0a, 0x18, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68,
	0x65, 0x64, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x03,
	0x72, 0x75, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x69, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x75, 0x6e, 0x52, 0x03, 0x72, 0x75, 0x6e, 0x22, 0x5c,
	0x0a, 0x16, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x55, 0x73, 0x65, 0x64, 0x4e, 0x65, 0x74, 0x72,
	0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x67, 0x65, 0x72, 0x72,
	0x69, 0x74, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x67,
	0x65, 0x72, 0x72, 0x69, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x3b, 0x0a, 0x16,
	0x46, 0x65, 0x74, 0x63, 0x68, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x75, 0x6e, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x75, 0x63, 0x69, 0x5f, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x75,
	0x63, 0x69, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x3d, 0x0a, 0x17, 0x46, 0x65, 0x74,
	0x63, 0x68, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x75, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x72, 0x75, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52,
	0x75, 0x6e, 0x52, 0x04, 0x72, 0x75, 0x6e, 0x73, 0x22, 0x66, 0x0a, 0x03, 0x52, 0x75, 0x6e, 0x12,
	0x2b, 0x0a, 0x07, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x41, 0x74, 0x74, 0x65,
	0x6d, 0x70, 0x74, 0x52, 0x07, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x03,
	0x63, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x69, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x75, 0x6e, 0x43, 0x4c, 0x52, 0x03, 0x63, 0x6c, 0x73,
	0x22, 0xf7, 0x02, 0x0a, 0x05, 0x52, 0x75, 0x6e, 0x43, 0x4c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x02, 0x67, 0x63,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x2e, 0x47, 0x65, 0x72, 0x72, 0x69, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x02,
	0x67, 0x63, 0x12, 0x3d, 0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x32, 0x0a, 0x07, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52,
	0x75, 0x6e, 0x43, 0x4c, 0x2e, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x07, 0x74, 0x72,
	0x69, 0x67, 0x67, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x65, 0x70, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x52, 0x75, 0x6e, 0x43, 0x4c, 0x2e, 0x44, 0x65, 0x70, 0x52, 0x04, 0x64, 0x65, 0x70, 0x73, 0x1a,
	0x6e, 0x0a, 0x07, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x1a,
	0x29, 0x0a, 0x03, 0x44, 0x65, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x72, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x68, 0x61, 0x72, 0x64, 0x32, 0xc9, 0x02, 0x0a, 0x09, 0x4d,
	0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x42, 0x0a, 0x0a, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x52, 0x75, 0x6e, 0x73, 0x12, 0x1c, 0x2e, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x75, 0x6e, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x50, 0x0a, 0x11,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x52, 0x75,
	0x6e, 0x12, 0x23, 0x2e, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x52, 0x75, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x4c,
	0x0a, 0x0f, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x55, 0x73, 0x65, 0x64, 0x4e, 0x65, 0x74, 0x72,
	0x63, 0x12, 0x21, 0x2e, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x55, 0x73, 0x65, 0x64, 0x4e, 0x65, 0x74, 0x72, 0x63, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x58, 0x0a, 0x0f,
	0x46, 0x65, 0x74, 0x63, 0x68, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x75, 0x6e, 0x73, 0x12,
	0x21, 0x2e, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x46, 0x65, 0x74, 0x63,
	0x68, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x75, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x22, 0x2e, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x46,
	0x65, 0x74, 0x63, 0x68, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x75, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72,
	0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63,
	0x76, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3b,
	0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_cv_api_migration_migration_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_cv_api_migration_migration_proto_rawDescData = file_go_chromium_org_luci_cv_api_migration_migration_proto_rawDesc
)

func file_go_chromium_org_luci_cv_api_migration_migration_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_cv_api_migration_migration_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_cv_api_migration_migration_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_cv_api_migration_migration_proto_rawDescData)
	})
	return file_go_chromium_org_luci_cv_api_migration_migration_proto_rawDescData
}

var file_go_chromium_org_luci_cv_api_migration_migration_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_go_chromium_org_luci_cv_api_migration_migration_proto_goTypes = []interface{}{
	(*ReportRunsRequest)(nil),        // 0: migration.ReportRunsRequest
	(*ReportFinishedRunRequest)(nil), // 1: migration.ReportFinishedRunRequest
	(*ReportUsedNetrcRequest)(nil),   // 2: migration.ReportUsedNetrcRequest
	(*FetchActiveRunsRequest)(nil),   // 3: migration.FetchActiveRunsRequest
	(*FetchActiveRunsResponse)(nil),  // 4: migration.FetchActiveRunsResponse
	(*Run)(nil),                      // 5: migration.Run
	(*RunCL)(nil),                    // 6: migration.RunCL
	(*RunCL_Trigger)(nil),            // 7: migration.RunCL.Trigger
	(*RunCL_Dep)(nil),                // 8: migration.RunCL.Dep
	(*v1.Attempt)(nil),               // 9: bigquery.Attempt
	(*v1.GerritChange)(nil),          // 10: bigquery.GerritChange
	(*timestamppb.Timestamp)(nil),    // 11: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),            // 12: google.protobuf.Empty
}
var file_go_chromium_org_luci_cv_api_migration_migration_proto_depIdxs = []int32{
	5,  // 0: migration.ReportRunsRequest.runs:type_name -> migration.Run
	5,  // 1: migration.ReportFinishedRunRequest.run:type_name -> migration.Run
	5,  // 2: migration.FetchActiveRunsResponse.runs:type_name -> migration.Run
	9,  // 3: migration.Run.attempt:type_name -> bigquery.Attempt
	6,  // 4: migration.Run.cls:type_name -> migration.RunCL
	10, // 5: migration.RunCL.gc:type_name -> bigquery.GerritChange
	11, // 6: migration.RunCL.updated_time:type_name -> google.protobuf.Timestamp
	7,  // 7: migration.RunCL.trigger:type_name -> migration.RunCL.Trigger
	8,  // 8: migration.RunCL.deps:type_name -> migration.RunCL.Dep
	11, // 9: migration.RunCL.Trigger.time:type_name -> google.protobuf.Timestamp
	0,  // 10: migration.Migration.ReportRuns:input_type -> migration.ReportRunsRequest
	1,  // 11: migration.Migration.ReportFinishedRun:input_type -> migration.ReportFinishedRunRequest
	2,  // 12: migration.Migration.ReportUsedNetrc:input_type -> migration.ReportUsedNetrcRequest
	3,  // 13: migration.Migration.FetchActiveRuns:input_type -> migration.FetchActiveRunsRequest
	12, // 14: migration.Migration.ReportRuns:output_type -> google.protobuf.Empty
	12, // 15: migration.Migration.ReportFinishedRun:output_type -> google.protobuf.Empty
	12, // 16: migration.Migration.ReportUsedNetrc:output_type -> google.protobuf.Empty
	4,  // 17: migration.Migration.FetchActiveRuns:output_type -> migration.FetchActiveRunsResponse
	14, // [14:18] is the sub-list for method output_type
	10, // [10:14] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_cv_api_migration_migration_proto_init() }
func file_go_chromium_org_luci_cv_api_migration_migration_proto_init() {
	if File_go_chromium_org_luci_cv_api_migration_migration_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_cv_api_migration_migration_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportRunsRequest); i {
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
		file_go_chromium_org_luci_cv_api_migration_migration_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportFinishedRunRequest); i {
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
		file_go_chromium_org_luci_cv_api_migration_migration_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportUsedNetrcRequest); i {
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
		file_go_chromium_org_luci_cv_api_migration_migration_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchActiveRunsRequest); i {
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
		file_go_chromium_org_luci_cv_api_migration_migration_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchActiveRunsResponse); i {
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
		file_go_chromium_org_luci_cv_api_migration_migration_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Run); i {
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
		file_go_chromium_org_luci_cv_api_migration_migration_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunCL); i {
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
		file_go_chromium_org_luci_cv_api_migration_migration_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunCL_Trigger); i {
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
		file_go_chromium_org_luci_cv_api_migration_migration_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunCL_Dep); i {
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
			RawDescriptor: file_go_chromium_org_luci_cv_api_migration_migration_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_go_chromium_org_luci_cv_api_migration_migration_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_cv_api_migration_migration_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_cv_api_migration_migration_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_cv_api_migration_migration_proto = out.File
	file_go_chromium_org_luci_cv_api_migration_migration_proto_rawDesc = nil
	file_go_chromium_org_luci_cv_api_migration_migration_proto_goTypes = nil
	file_go_chromium_org_luci_cv_api_migration_migration_proto_depIdxs = nil
}
