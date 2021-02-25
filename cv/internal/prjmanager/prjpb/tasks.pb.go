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
// source: go.chromium.org/luci/cv/internal/prjmanager/prjpb/tasks.proto

package prjpb

import (
	run "go.chromium.org/luci/cv/internal/run"
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

// PokePMTask sends a signal to ProjectManager to process events.
//
// Always used with de-duplication and thus can't be created from a transaction.
//
// Queue: "manage-project".
type PokePMTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LuciProject string                 `protobuf:"bytes,1,opt,name=luci_project,json=luciProject,proto3" json:"luci_project,omitempty"`
	Eta         *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=eta,proto3" json:"eta,omitempty"`
}

func (x *PokePMTask) Reset() {
	*x = PokePMTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_tasks_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PokePMTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PokePMTask) ProtoMessage() {}

func (x *PokePMTask) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_tasks_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PokePMTask.ProtoReflect.Descriptor instead.
func (*PokePMTask) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_tasks_proto_rawDescGZIP(), []int{0}
}

func (x *PokePMTask) GetLuciProject() string {
	if x != nil {
		return x.LuciProject
	}
	return ""
}

func (x *PokePMTask) GetEta() *timestamppb.Timestamp {
	if x != nil {
		return x.Eta
	}
	return nil
}

// KickPokePMTask starts a task to actually enqueue PokePMTask.
//
// It exists in order to poke ProjectManager from a transaction.
//
// Queue: "manage-project".
type KickPokePMTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LuciProject string                 `protobuf:"bytes,1,opt,name=luci_project,json=luciProject,proto3" json:"luci_project,omitempty"`
	Eta         *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=eta,proto3" json:"eta,omitempty"`
}

func (x *KickPokePMTask) Reset() {
	*x = KickPokePMTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_tasks_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KickPokePMTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KickPokePMTask) ProtoMessage() {}

func (x *KickPokePMTask) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_tasks_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KickPokePMTask.ProtoReflect.Descriptor instead.
func (*KickPokePMTask) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_tasks_proto_rawDescGZIP(), []int{1}
}

func (x *KickPokePMTask) GetLuciProject() string {
	if x != nil {
		return x.LuciProject
	}
	return ""
}

func (x *KickPokePMTask) GetEta() *timestamppb.Timestamp {
	if x != nil {
		return x.Eta
	}
	return nil
}

// PurgeCLTask starts a task to purge a CL.
//
// Queue: "purge-cls".
type PurgeCLTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LuciProject string              `protobuf:"bytes,1,opt,name=luci_project,json=luciProject,proto3" json:"luci_project,omitempty"`
	PurgingCl   *PurgingCL          `protobuf:"bytes,2,opt,name=purging_cl,json=purgingCl,proto3" json:"purging_cl,omitempty"`
	Trigger     *run.Trigger        `protobuf:"bytes,3,opt,name=trigger,proto3" json:"trigger,omitempty"`
	Reason      *PurgeCLTask_Reason `protobuf:"bytes,4,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (x *PurgeCLTask) Reset() {
	*x = PurgeCLTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_tasks_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PurgeCLTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurgeCLTask) ProtoMessage() {}

func (x *PurgeCLTask) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_tasks_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PurgeCLTask.ProtoReflect.Descriptor instead.
func (*PurgeCLTask) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_tasks_proto_rawDescGZIP(), []int{2}
}

func (x *PurgeCLTask) GetLuciProject() string {
	if x != nil {
		return x.LuciProject
	}
	return ""
}

func (x *PurgeCLTask) GetPurgingCl() *PurgingCL {
	if x != nil {
		return x.PurgingCl
	}
	return nil
}

func (x *PurgeCLTask) GetTrigger() *run.Trigger {
	if x != nil {
		return x.Trigger
	}
	return nil
}

func (x *PurgeCLTask) GetReason() *PurgeCLTask_Reason {
	if x != nil {
		return x.Reason
	}
	return nil
}

type PurgeCLTask_Reason struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PurgeCLTask_Reason) Reset() {
	*x = PurgeCLTask_Reason{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_tasks_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PurgeCLTask_Reason) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurgeCLTask_Reason) ProtoMessage() {}

func (x *PurgeCLTask_Reason) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_tasks_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PurgeCLTask_Reason.ProtoReflect.Descriptor instead.
func (*PurgeCLTask_Reason) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_tasks_proto_rawDescGZIP(), []int{2, 0}
}

var File_go_chromium_org_luci_cv_internal_prjmanager_prjpb_tasks_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_tasks_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x76, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6a, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x70, 0x72,
	0x6a, 0x70, 0x62, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x13, 0x63, 0x76, 0x2e, 0x70, 0x72, 0x6a, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6a, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69,
	0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x76, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x72, 0x75, 0x6e, 0x2f, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3f, 0x67, 0x6f, 0x2e, 0x63, 0x68,
	0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f,
	0x63, 0x76, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6a, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6a, 0x70, 0x62, 0x2f, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5d, 0x0a, 0x0a, 0x50, 0x6f,
	0x6b, 0x65, 0x50, 0x4d, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x75, 0x63, 0x69,
	0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x6c, 0x75, 0x63, 0x69, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x2c, 0x0a, 0x03, 0x65,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x03, 0x65, 0x74, 0x61, 0x22, 0x61, 0x0a, 0x0e, 0x4b, 0x69, 0x63,
	0x6b, 0x50, 0x6f, 0x6b, 0x65, 0x50, 0x4d, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x21, 0x0a, 0x0c, 0x6c,
	0x75, 0x63, 0x69, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x6c, 0x75, 0x63, 0x69, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x2c,
	0x0a, 0x03, 0x65, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x03, 0x65, 0x74, 0x61, 0x22, 0xe5, 0x01, 0x0a,
	0x0b, 0x50, 0x75, 0x72, 0x67, 0x65, 0x43, 0x4c, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x21, 0x0a, 0x0c,
	0x6c, 0x75, 0x63, 0x69, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x6c, 0x75, 0x63, 0x69, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x3d, 0x0a, 0x0a, 0x70, 0x75, 0x72, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x76, 0x2e, 0x70, 0x72, 0x6a, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6a, 0x70, 0x62, 0x2e, 0x50, 0x75, 0x72, 0x67, 0x69, 0x6e,
	0x67, 0x43, 0x4c, 0x52, 0x09, 0x70, 0x75, 0x72, 0x67, 0x69, 0x6e, 0x67, 0x43, 0x6c, 0x12, 0x29,
	0x0a, 0x07, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x63, 0x76, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72,
	0x52, 0x07, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x12, 0x3f, 0x0a, 0x06, 0x72, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x76, 0x2e, 0x70,
	0x72, 0x6a, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6a, 0x70, 0x62, 0x2e,
	0x50, 0x75, 0x72, 0x67, 0x65, 0x43, 0x4c, 0x54, 0x61, 0x73, 0x6b, 0x2e, 0x52, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x1a, 0x08, 0x0a, 0x06, 0x52, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d,
	0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x76, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6a, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6a, 0x70, 0x62, 0x3b, 0x70, 0x72, 0x6a, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_tasks_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_tasks_proto_rawDescData = file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_tasks_proto_rawDesc
)

func file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_tasks_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_tasks_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_tasks_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_tasks_proto_rawDescData)
	})
	return file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_tasks_proto_rawDescData
}

var file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_tasks_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_tasks_proto_goTypes = []interface{}{
	(*PokePMTask)(nil),            // 0: cv.prjmanager.prjpb.PokePMTask
	(*KickPokePMTask)(nil),        // 1: cv.prjmanager.prjpb.KickPokePMTask
	(*PurgeCLTask)(nil),           // 2: cv.prjmanager.prjpb.PurgeCLTask
	(*PurgeCLTask_Reason)(nil),    // 3: cv.prjmanager.prjpb.PurgeCLTask.Reason
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
	(*PurgingCL)(nil),             // 5: cv.prjmanager.prjpb.PurgingCL
	(*run.Trigger)(nil),           // 6: cv.run.Trigger
}
var file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_tasks_proto_depIdxs = []int32{
	4, // 0: cv.prjmanager.prjpb.PokePMTask.eta:type_name -> google.protobuf.Timestamp
	4, // 1: cv.prjmanager.prjpb.KickPokePMTask.eta:type_name -> google.protobuf.Timestamp
	5, // 2: cv.prjmanager.prjpb.PurgeCLTask.purging_cl:type_name -> cv.prjmanager.prjpb.PurgingCL
	6, // 3: cv.prjmanager.prjpb.PurgeCLTask.trigger:type_name -> cv.run.Trigger
	3, // 4: cv.prjmanager.prjpb.PurgeCLTask.reason:type_name -> cv.prjmanager.prjpb.PurgeCLTask.Reason
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_tasks_proto_init() }
func file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_tasks_proto_init() {
	if File_go_chromium_org_luci_cv_internal_prjmanager_prjpb_tasks_proto != nil {
		return
	}
	file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_storage_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_tasks_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PokePMTask); i {
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
		file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_tasks_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KickPokePMTask); i {
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
		file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_tasks_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PurgeCLTask); i {
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
		file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_tasks_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PurgeCLTask_Reason); i {
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
			RawDescriptor: file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_tasks_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_tasks_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_tasks_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_tasks_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_cv_internal_prjmanager_prjpb_tasks_proto = out.File
	file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_tasks_proto_rawDesc = nil
	file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_tasks_proto_goTypes = nil
	file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_tasks_proto_depIdxs = nil
}
