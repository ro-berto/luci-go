// Copyright 2021 The LUCI Authors.
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
// source: go.chromium.org/luci/cv/internal/changelist/task.proto

package changelist

import (
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

type UpdateCLTask_Requester int32

const (
	UpdateCLTask_REQUESTER_CLASS_UNSPECIFIED UpdateCLTask_Requester = 0
	UpdateCLTask_INCR_POLL_MATCHED           UpdateCLTask_Requester = 1
	UpdateCLTask_FULL_POLL_MATCHED           UpdateCLTask_Requester = 2
	UpdateCLTask_FULL_POLL_UNMATCHED         UpdateCLTask_Requester = 3
	UpdateCLTask_PUBSUB_POLL                 UpdateCLTask_Requester = 4
	UpdateCLTask_CL_PURGER                   UpdateCLTask_Requester = 5
	UpdateCLTask_RPC_ADMIN                   UpdateCLTask_Requester = 6
	UpdateCLTask_RUN_POKE                    UpdateCLTask_Requester = 7
	UpdateCLTask_RUN_REMOVAL                 UpdateCLTask_Requester = 8
	UpdateCLTask_CANCEL_CL_TRIGGER           UpdateCLTask_Requester = 9
	UpdateCLTask_UPDATE_CONFIG               UpdateCLTask_Requester = 10
)

// Enum value maps for UpdateCLTask_Requester.
var (
	UpdateCLTask_Requester_name = map[int32]string{
		0:  "REQUESTER_CLASS_UNSPECIFIED",
		1:  "INCR_POLL_MATCHED",
		2:  "FULL_POLL_MATCHED",
		3:  "FULL_POLL_UNMATCHED",
		4:  "PUBSUB_POLL",
		5:  "CL_PURGER",
		6:  "RPC_ADMIN",
		7:  "RUN_POKE",
		8:  "RUN_REMOVAL",
		9:  "CANCEL_CL_TRIGGER",
		10: "UPDATE_CONFIG",
	}
	UpdateCLTask_Requester_value = map[string]int32{
		"REQUESTER_CLASS_UNSPECIFIED": 0,
		"INCR_POLL_MATCHED":           1,
		"FULL_POLL_MATCHED":           2,
		"FULL_POLL_UNMATCHED":         3,
		"PUBSUB_POLL":                 4,
		"CL_PURGER":                   5,
		"RPC_ADMIN":                   6,
		"RUN_POKE":                    7,
		"RUN_REMOVAL":                 8,
		"CANCEL_CL_TRIGGER":           9,
		"UPDATE_CONFIG":               10,
	}
)

func (x UpdateCLTask_Requester) Enum() *UpdateCLTask_Requester {
	p := new(UpdateCLTask_Requester)
	*p = x
	return p
}

func (x UpdateCLTask_Requester) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UpdateCLTask_Requester) Descriptor() protoreflect.EnumDescriptor {
	return file_go_chromium_org_luci_cv_internal_changelist_task_proto_enumTypes[0].Descriptor()
}

func (UpdateCLTask_Requester) Type() protoreflect.EnumType {
	return &file_go_chromium_org_luci_cv_internal_changelist_task_proto_enumTypes[0]
}

func (x UpdateCLTask_Requester) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UpdateCLTask_Requester.Descriptor instead.
func (UpdateCLTask_Requester) EnumDescriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_internal_changelist_task_proto_rawDescGZIP(), []int{0, 0}
}

// UpdateCLTask is for updating a single CL.
//
// Queue: "update-cl".
type UpdateCLTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LuciProject string `protobuf:"bytes,1,opt,name=luci_project,json=luciProject,proto3" json:"luci_project,omitempty"`
	// At least one of internal or external ID must be given.
	Id         int64  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"` // internal CLID
	ExternalId string `protobuf:"bytes,3,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
	// TODO(crbug.com/1358208): remove updated_hint.
	UpdatedHint *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=updated_hint,json=updatedHint,proto3" json:"updated_hint,omitempty"`
	// Requester identifies various scenarios that enqueue UpdateCLTask(s).
	//
	// This is used to track UpdateCLTask(s) by the requester for monitoring
	// purposes.
	Requester UpdateCLTask_Requester `protobuf:"varint,5,opt,name=requester,proto3,enum=cv.internal.changelist.UpdateCLTask_Requester" json:"requester,omitempty"`
	// True if the UpdateCLTask was enqueued to resolve a dependency.
	IsForDep bool `protobuf:"varint,6,opt,name=is_for_dep,json=isForDep,proto3" json:"is_for_dep,omitempty"`
	// Hint provides various hints for the snapshot to be fetched.
	Hint *UpdateCLTask_Hint `protobuf:"bytes,7,opt,name=hint,proto3" json:"hint,omitempty"`
}

func (x *UpdateCLTask) Reset() {
	*x = UpdateCLTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_internal_changelist_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCLTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCLTask) ProtoMessage() {}

func (x *UpdateCLTask) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_internal_changelist_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCLTask.ProtoReflect.Descriptor instead.
func (*UpdateCLTask) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_internal_changelist_task_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateCLTask) GetLuciProject() string {
	if x != nil {
		return x.LuciProject
	}
	return ""
}

func (x *UpdateCLTask) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateCLTask) GetExternalId() string {
	if x != nil {
		return x.ExternalId
	}
	return ""
}

func (x *UpdateCLTask) GetUpdatedHint() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedHint
	}
	return nil
}

func (x *UpdateCLTask) GetRequester() UpdateCLTask_Requester {
	if x != nil {
		return x.Requester
	}
	return UpdateCLTask_REQUESTER_CLASS_UNSPECIFIED
}

func (x *UpdateCLTask) GetIsForDep() bool {
	if x != nil {
		return x.IsForDep
	}
	return false
}

func (x *UpdateCLTask) GetHint() *UpdateCLTask_Hint {
	if x != nil {
		return x.Hint
	}
	return nil
}

// BatchUpdateCLTask is for updating many CLs.
//
// When executed, it just enqueues its tasks as individual UpdateCLTask TQ
// for independent execution.
//
// Queue: "update-cl".
type BatchUpdateCLTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tasks []*UpdateCLTask `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
}

func (x *BatchUpdateCLTask) Reset() {
	*x = BatchUpdateCLTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_internal_changelist_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchUpdateCLTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchUpdateCLTask) ProtoMessage() {}

func (x *BatchUpdateCLTask) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_internal_changelist_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchUpdateCLTask.ProtoReflect.Descriptor instead.
func (*BatchUpdateCLTask) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_internal_changelist_task_proto_rawDescGZIP(), []int{1}
}

func (x *BatchUpdateCLTask) GetTasks() []*UpdateCLTask {
	if x != nil {
		return x.Tasks
	}
	return nil
}

// BatchOnCLUpdatedTask notifies many Projects and Runs about updated CLs.
//
// Queue: "notify-on-cl-updated".
type BatchOnCLUpdatedTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Projects map[string]*CLUpdatedEvents `protobuf:"bytes,1,rep,name=projects,proto3" json:"projects,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Runs     map[string]*CLUpdatedEvents `protobuf:"bytes,2,rep,name=runs,proto3" json:"runs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *BatchOnCLUpdatedTask) Reset() {
	*x = BatchOnCLUpdatedTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_internal_changelist_task_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchOnCLUpdatedTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchOnCLUpdatedTask) ProtoMessage() {}

func (x *BatchOnCLUpdatedTask) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_internal_changelist_task_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchOnCLUpdatedTask.ProtoReflect.Descriptor instead.
func (*BatchOnCLUpdatedTask) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_internal_changelist_task_proto_rawDescGZIP(), []int{2}
}

func (x *BatchOnCLUpdatedTask) GetProjects() map[string]*CLUpdatedEvents {
	if x != nil {
		return x.Projects
	}
	return nil
}

func (x *BatchOnCLUpdatedTask) GetRuns() map[string]*CLUpdatedEvents {
	if x != nil {
		return x.Runs
	}
	return nil
}

type UpdateCLTask_Hint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The external update time of the Snapshot to fetch.
	ExternalUpdateTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=external_update_time,json=externalUpdateTime,proto3" json:"external_update_time,omitempty"`
}

func (x *UpdateCLTask_Hint) Reset() {
	*x = UpdateCLTask_Hint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_internal_changelist_task_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCLTask_Hint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCLTask_Hint) ProtoMessage() {}

func (x *UpdateCLTask_Hint) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_internal_changelist_task_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCLTask_Hint.ProtoReflect.Descriptor instead.
func (*UpdateCLTask_Hint) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_internal_changelist_task_proto_rawDescGZIP(), []int{0, 0}
}

func (x *UpdateCLTask_Hint) GetExternalUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ExternalUpdateTime
	}
	return nil
}

var File_go_chromium_org_luci_cv_internal_changelist_task_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_cv_internal_changelist_task_proto_rawDesc = []byte{
	0x0a, 0x36, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x76, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x74, 0x61,
	0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x63, 0x76, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x69, 0x73, 0x74,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x39, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x76, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x90, 0x05, 0x0a,
	0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x4c, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x21, 0x0a,
	0x0c, 0x6c, 0x75, 0x63, 0x69, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x75, 0x63, 0x69, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49,
	0x64, 0x12, 0x3d, 0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x68, 0x69, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x48, 0x69, 0x6e, 0x74,
	0x12, 0x4c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x63, 0x76, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x4c, 0x54, 0x61, 0x73, 0x6b, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1c,
	0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x66, 0x6f, 0x72, 0x5f, 0x64, 0x65, 0x70, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x46, 0x6f, 0x72, 0x44, 0x65, 0x70, 0x12, 0x3d, 0x0a, 0x04,
	0x68, 0x69, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63, 0x76, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c,
	0x69, 0x73, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x4c, 0x54, 0x61, 0x73, 0x6b,
	0x2e, 0x48, 0x69, 0x6e, 0x74, 0x52, 0x04, 0x68, 0x69, 0x6e, 0x74, 0x1a, 0x54, 0x0a, 0x04, 0x48,
	0x69, 0x6e, 0x74, 0x12, 0x4c, 0x0a, 0x14, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x12, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x22, 0xeb, 0x01, 0x0a, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x72, 0x12,
	0x1f, 0x0a, 0x1b, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x43, 0x4c, 0x41,
	0x53, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x15, 0x0a, 0x11, 0x49, 0x4e, 0x43, 0x52, 0x5f, 0x50, 0x4f, 0x4c, 0x4c, 0x5f, 0x4d, 0x41,
	0x54, 0x43, 0x48, 0x45, 0x44, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x46, 0x55, 0x4c, 0x4c, 0x5f,
	0x50, 0x4f, 0x4c, 0x4c, 0x5f, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x45, 0x44, 0x10, 0x02, 0x12, 0x17,
	0x0a, 0x13, 0x46, 0x55, 0x4c, 0x4c, 0x5f, 0x50, 0x4f, 0x4c, 0x4c, 0x5f, 0x55, 0x4e, 0x4d, 0x41,
	0x54, 0x43, 0x48, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x50, 0x55, 0x42, 0x53, 0x55,
	0x42, 0x5f, 0x50, 0x4f, 0x4c, 0x4c, 0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4c, 0x5f, 0x50,
	0x55, 0x52, 0x47, 0x45, 0x52, 0x10, 0x05, 0x12, 0x0d, 0x0a, 0x09, 0x52, 0x50, 0x43, 0x5f, 0x41,
	0x44, 0x4d, 0x49, 0x4e, 0x10, 0x06, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x55, 0x4e, 0x5f, 0x50, 0x4f,
	0x4b, 0x45, 0x10, 0x07, 0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x55, 0x4e, 0x5f, 0x52, 0x45, 0x4d, 0x4f,
	0x56, 0x41, 0x4c, 0x10, 0x08, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x5f,
	0x43, 0x4c, 0x5f, 0x54, 0x52, 0x49, 0x47, 0x47, 0x45, 0x52, 0x10, 0x09, 0x12, 0x11, 0x0a, 0x0d,
	0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x10, 0x0a, 0x22,
	0x4f, 0x0a, 0x11, 0x42, 0x61, 0x74, 0x63, 0x68, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x4c,
	0x54, 0x61, 0x73, 0x6b, 0x12, 0x3a, 0x0a, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x76, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x4c, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73,
	0x22, 0x82, 0x03, 0x0a, 0x14, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x6e, 0x43, 0x4c, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x56, 0x0a, 0x08, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x63, 0x76,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x6c, 0x69, 0x73, 0x74, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x6e, 0x43, 0x4c, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x12, 0x4a, 0x0a, 0x04, 0x72, 0x75, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x36, 0x2e, 0x63, 0x76, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x6e,
	0x43, 0x4c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x2e, 0x52, 0x75,
	0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x72, 0x75, 0x6e, 0x73, 0x1a, 0x64, 0x0a,
	0x0d, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x3d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x63, 0x76, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x43, 0x4c, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x1a, 0x60, 0x0a, 0x09, 0x52, 0x75, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x3d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x27, 0x2e, 0x63, 0x76, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x43, 0x4c, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f,
	0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x76,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x6c, 0x69, 0x73, 0x74, 0x3b, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_cv_internal_changelist_task_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_cv_internal_changelist_task_proto_rawDescData = file_go_chromium_org_luci_cv_internal_changelist_task_proto_rawDesc
)

func file_go_chromium_org_luci_cv_internal_changelist_task_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_cv_internal_changelist_task_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_cv_internal_changelist_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_cv_internal_changelist_task_proto_rawDescData)
	})
	return file_go_chromium_org_luci_cv_internal_changelist_task_proto_rawDescData
}

var file_go_chromium_org_luci_cv_internal_changelist_task_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_go_chromium_org_luci_cv_internal_changelist_task_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_go_chromium_org_luci_cv_internal_changelist_task_proto_goTypes = []interface{}{
	(UpdateCLTask_Requester)(0),   // 0: cv.internal.changelist.UpdateCLTask.Requester
	(*UpdateCLTask)(nil),          // 1: cv.internal.changelist.UpdateCLTask
	(*BatchUpdateCLTask)(nil),     // 2: cv.internal.changelist.BatchUpdateCLTask
	(*BatchOnCLUpdatedTask)(nil),  // 3: cv.internal.changelist.BatchOnCLUpdatedTask
	(*UpdateCLTask_Hint)(nil),     // 4: cv.internal.changelist.UpdateCLTask.Hint
	nil,                           // 5: cv.internal.changelist.BatchOnCLUpdatedTask.ProjectsEntry
	nil,                           // 6: cv.internal.changelist.BatchOnCLUpdatedTask.RunsEntry
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
	(*CLUpdatedEvents)(nil),       // 8: cv.internal.changelist.CLUpdatedEvents
}
var file_go_chromium_org_luci_cv_internal_changelist_task_proto_depIdxs = []int32{
	7, // 0: cv.internal.changelist.UpdateCLTask.updated_hint:type_name -> google.protobuf.Timestamp
	0, // 1: cv.internal.changelist.UpdateCLTask.requester:type_name -> cv.internal.changelist.UpdateCLTask.Requester
	4, // 2: cv.internal.changelist.UpdateCLTask.hint:type_name -> cv.internal.changelist.UpdateCLTask.Hint
	1, // 3: cv.internal.changelist.BatchUpdateCLTask.tasks:type_name -> cv.internal.changelist.UpdateCLTask
	5, // 4: cv.internal.changelist.BatchOnCLUpdatedTask.projects:type_name -> cv.internal.changelist.BatchOnCLUpdatedTask.ProjectsEntry
	6, // 5: cv.internal.changelist.BatchOnCLUpdatedTask.runs:type_name -> cv.internal.changelist.BatchOnCLUpdatedTask.RunsEntry
	7, // 6: cv.internal.changelist.UpdateCLTask.Hint.external_update_time:type_name -> google.protobuf.Timestamp
	8, // 7: cv.internal.changelist.BatchOnCLUpdatedTask.ProjectsEntry.value:type_name -> cv.internal.changelist.CLUpdatedEvents
	8, // 8: cv.internal.changelist.BatchOnCLUpdatedTask.RunsEntry.value:type_name -> cv.internal.changelist.CLUpdatedEvents
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_cv_internal_changelist_task_proto_init() }
func file_go_chromium_org_luci_cv_internal_changelist_task_proto_init() {
	if File_go_chromium_org_luci_cv_internal_changelist_task_proto != nil {
		return
	}
	file_go_chromium_org_luci_cv_internal_changelist_storage_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_cv_internal_changelist_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCLTask); i {
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
		file_go_chromium_org_luci_cv_internal_changelist_task_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchUpdateCLTask); i {
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
		file_go_chromium_org_luci_cv_internal_changelist_task_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchOnCLUpdatedTask); i {
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
		file_go_chromium_org_luci_cv_internal_changelist_task_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCLTask_Hint); i {
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
			RawDescriptor: file_go_chromium_org_luci_cv_internal_changelist_task_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_cv_internal_changelist_task_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_cv_internal_changelist_task_proto_depIdxs,
		EnumInfos:         file_go_chromium_org_luci_cv_internal_changelist_task_proto_enumTypes,
		MessageInfos:      file_go_chromium_org_luci_cv_internal_changelist_task_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_cv_internal_changelist_task_proto = out.File
	file_go_chromium_org_luci_cv_internal_changelist_task_proto_rawDesc = nil
	file_go_chromium_org_luci_cv_internal_changelist_task_proto_goTypes = nil
	file_go_chromium_org_luci_cv_internal_changelist_task_proto_depIdxs = nil
}
