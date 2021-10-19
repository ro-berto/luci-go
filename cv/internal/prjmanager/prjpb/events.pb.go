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
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: go.chromium.org/luci/cv/internal/prjmanager/prjpb/events.proto

package prjpb

import (
	changelist "go.chromium.org/luci/cv/internal/changelist"
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

// Event is a container for all kinds of events a project manager can receive.
type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Event:
	//	*Event_NewConfig
	//	*Event_Poke
	//	*Event_ClsUpdated
	//	*Event_RunCreated
	//	*Event_RunFinished
	//	*Event_PurgeCompleted
	Event isEvent_Event `protobuf_oneof:"event"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_events_proto_rawDescGZIP(), []int{0}
}

func (m *Event) GetEvent() isEvent_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (x *Event) GetNewConfig() *NewConfig {
	if x, ok := x.GetEvent().(*Event_NewConfig); ok {
		return x.NewConfig
	}
	return nil
}

func (x *Event) GetPoke() *Poke {
	if x, ok := x.GetEvent().(*Event_Poke); ok {
		return x.Poke
	}
	return nil
}

func (x *Event) GetClsUpdated() *changelist.CLUpdatedEvents {
	if x, ok := x.GetEvent().(*Event_ClsUpdated); ok {
		return x.ClsUpdated
	}
	return nil
}

func (x *Event) GetRunCreated() *RunCreated {
	if x, ok := x.GetEvent().(*Event_RunCreated); ok {
		return x.RunCreated
	}
	return nil
}

func (x *Event) GetRunFinished() *RunFinished {
	if x, ok := x.GetEvent().(*Event_RunFinished); ok {
		return x.RunFinished
	}
	return nil
}

func (x *Event) GetPurgeCompleted() *PurgeCompleted {
	if x, ok := x.GetEvent().(*Event_PurgeCompleted); ok {
		return x.PurgeCompleted
	}
	return nil
}

type isEvent_Event interface {
	isEvent_Event()
}

type Event_NewConfig struct {
	NewConfig *NewConfig `protobuf:"bytes,1,opt,name=new_config,json=newConfig,proto3,oneof"`
}

type Event_Poke struct {
	Poke *Poke `protobuf:"bytes,2,opt,name=poke,proto3,oneof"`
}

type Event_ClsUpdated struct {
	ClsUpdated *changelist.CLUpdatedEvents `protobuf:"bytes,7,opt,name=cls_updated,json=clsUpdated,proto3,oneof"`
}

type Event_RunCreated struct {
	RunCreated *RunCreated `protobuf:"bytes,4,opt,name=run_created,json=runCreated,proto3,oneof"`
}

type Event_RunFinished struct {
	RunFinished *RunFinished `protobuf:"bytes,5,opt,name=run_finished,json=runFinished,proto3,oneof"`
}

type Event_PurgeCompleted struct {
	PurgeCompleted *PurgeCompleted `protobuf:"bytes,6,opt,name=purge_completed,json=purgeCompleted,proto3,oneof"`
}

func (*Event_NewConfig) isEvent_Event() {}

func (*Event_Poke) isEvent_Event() {}

func (*Event_ClsUpdated) isEvent_Event() {}

func (*Event_RunCreated) isEvent_Event() {}

func (*Event_RunFinished) isEvent_Event() {}

func (*Event_PurgeCompleted) isEvent_Event() {}

// NewConfig is sent to PM by Project Config updater upon saving newest config
// in datastore.
type NewConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NewConfig) Reset() {
	*x = NewConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_events_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewConfig) ProtoMessage() {}

func (x *NewConfig) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_events_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewConfig.ProtoReflect.Descriptor instead.
func (*NewConfig) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_events_proto_rawDescGZIP(), []int{1}
}

// Poke is sent to PM by Project Config updater.
type Poke struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Poke) Reset() {
	*x = Poke{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_events_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Poke) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Poke) ProtoMessage() {}

func (x *Poke) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_events_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Poke.ProtoReflect.Descriptor instead.
func (*Poke) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_events_proto_rawDescGZIP(), []int{2}
}

// RunCreated is sent to PM by either itself or API-based Run creation.
type RunCreated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RunId string `protobuf:"bytes,1,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
}

func (x *RunCreated) Reset() {
	*x = RunCreated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_events_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunCreated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunCreated) ProtoMessage() {}

func (x *RunCreated) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_events_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunCreated.ProtoReflect.Descriptor instead.
func (*RunCreated) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_events_proto_rawDescGZIP(), []int{3}
}

func (x *RunCreated) GetRunId() string {
	if x != nil {
		return x.RunId
	}
	return ""
}

// RunFinished is sent to PM by Run Manager after or atomically with changing Run's
// status to a final status.
type RunFinished struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RunId string `protobuf:"bytes,1,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
}

func (x *RunFinished) Reset() {
	*x = RunFinished{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_events_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunFinished) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunFinished) ProtoMessage() {}

func (x *RunFinished) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_events_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunFinished.ProtoReflect.Descriptor instead.
func (*RunFinished) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_events_proto_rawDescGZIP(), []int{4}
}

func (x *RunFinished) GetRunId() string {
	if x != nil {
		return x.RunId
	}
	return ""
}

// PurgingCompleted is sent to PM by TQ task purging a CL.
//
// See storage.proto:PurgingCL doc.
//
// There is no status of the purge because it's the CL state that matters,
// hence success or failure will reach PM via CLUpdatedEvent.
type PurgeCompleted struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Operation ID suffices to identify a purge.
	OperationId string `protobuf:"bytes,1,opt,name=operation_id,json=operationId,proto3" json:"operation_id,omitempty"`
}

func (x *PurgeCompleted) Reset() {
	*x = PurgeCompleted{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_events_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PurgeCompleted) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurgeCompleted) ProtoMessage() {}

func (x *PurgeCompleted) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_events_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PurgeCompleted.ProtoReflect.Descriptor instead.
func (*PurgeCompleted) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_events_proto_rawDescGZIP(), []int{5}
}

func (x *PurgeCompleted) GetOperationId() string {
	if x != nil {
		return x.OperationId
	}
	return ""
}

var File_go_chromium_org_luci_cv_internal_prjmanager_prjpb_events_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_events_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x76, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6a, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x70, 0x72,
	0x6a, 0x70, 0x62, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1c, 0x63, 0x76, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72,
	0x6a, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6a, 0x70, 0x62, 0x1a, 0x39,
	0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x76, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdc, 0x03, 0x0a, 0x05, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x48, 0x0a, 0x0a, 0x6e, 0x65, 0x77, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x76, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6a, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6a, 0x70, 0x62, 0x2e, 0x4e, 0x65, 0x77, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x48, 0x00, 0x52, 0x09, 0x6e, 0x65, 0x77, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x38, 0x0a,
	0x04, 0x70, 0x6f, 0x6b, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x76,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6a, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6a, 0x70, 0x62, 0x2e, 0x50, 0x6f, 0x6b, 0x65, 0x48,
	0x00, 0x52, 0x04, 0x70, 0x6f, 0x6b, 0x65, 0x12, 0x4a, 0x0a, 0x0b, 0x63, 0x6c, 0x73, 0x5f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63,
	0x76, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x43, 0x4c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x48, 0x00, 0x52, 0x0a, 0x63, 0x6c, 0x73, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x12, 0x4b, 0x0a, 0x0b, 0x72, 0x75, 0x6e, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x63, 0x76, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6a, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6a, 0x70, 0x62, 0x2e, 0x52, 0x75, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x48, 0x00, 0x52, 0x0a, 0x72, 0x75, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x12, 0x4e, 0x0a, 0x0c, 0x72, 0x75, 0x6e, 0x5f, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63, 0x76, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6a, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6a, 0x70, 0x62, 0x2e, 0x52, 0x75, 0x6e, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65,
	0x64, 0x48, 0x00, 0x52, 0x0b, 0x72, 0x75, 0x6e, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64,
	0x12, 0x57, 0x0a, 0x0f, 0x70, 0x75, 0x72, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x63, 0x76, 0x2e, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6a, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6a, 0x70, 0x62, 0x2e, 0x50, 0x75, 0x72, 0x67, 0x65, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x48, 0x00, 0x52, 0x0e, 0x70, 0x75, 0x72, 0x67, 0x65,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x04, 0x22, 0x0b, 0x0a, 0x09, 0x4e, 0x65, 0x77, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x06, 0x0a, 0x04, 0x50, 0x6f, 0x6b, 0x65, 0x22, 0x23, 0x0a,
	0x0a, 0x52, 0x75, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x72,
	0x75, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x75, 0x6e,
	0x49, 0x64, 0x22, 0x24, 0x0a, 0x0b, 0x52, 0x75, 0x6e, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65,
	0x64, 0x12, 0x15, 0x0a, 0x06, 0x72, 0x75, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x72, 0x75, 0x6e, 0x49, 0x64, 0x22, 0x33, 0x0a, 0x0e, 0x50, 0x75, 0x72, 0x67,
	0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x42, 0x39, 0x5a,
	0x37, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x76, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x70, 0x72, 0x6a, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6a,
	0x70, 0x62, 0x3b, 0x70, 0x72, 0x6a, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_events_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_events_proto_rawDescData = file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_events_proto_rawDesc
)

func file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_events_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_events_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_events_proto_rawDescData)
	})
	return file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_events_proto_rawDescData
}

var file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_events_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_events_proto_goTypes = []interface{}{
	(*Event)(nil),                      // 0: cv.internal.prjmanager.prjpb.Event
	(*NewConfig)(nil),                  // 1: cv.internal.prjmanager.prjpb.NewConfig
	(*Poke)(nil),                       // 2: cv.internal.prjmanager.prjpb.Poke
	(*RunCreated)(nil),                 // 3: cv.internal.prjmanager.prjpb.RunCreated
	(*RunFinished)(nil),                // 4: cv.internal.prjmanager.prjpb.RunFinished
	(*PurgeCompleted)(nil),             // 5: cv.internal.prjmanager.prjpb.PurgeCompleted
	(*changelist.CLUpdatedEvents)(nil), // 6: cv.internal.changelist.CLUpdatedEvents
}
var file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_events_proto_depIdxs = []int32{
	1, // 0: cv.internal.prjmanager.prjpb.Event.new_config:type_name -> cv.internal.prjmanager.prjpb.NewConfig
	2, // 1: cv.internal.prjmanager.prjpb.Event.poke:type_name -> cv.internal.prjmanager.prjpb.Poke
	6, // 2: cv.internal.prjmanager.prjpb.Event.cls_updated:type_name -> cv.internal.changelist.CLUpdatedEvents
	3, // 3: cv.internal.prjmanager.prjpb.Event.run_created:type_name -> cv.internal.prjmanager.prjpb.RunCreated
	4, // 4: cv.internal.prjmanager.prjpb.Event.run_finished:type_name -> cv.internal.prjmanager.prjpb.RunFinished
	5, // 5: cv.internal.prjmanager.prjpb.Event.purge_completed:type_name -> cv.internal.prjmanager.prjpb.PurgeCompleted
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_events_proto_init() }
func file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_events_proto_init() {
	if File_go_chromium_org_luci_cv_internal_prjmanager_prjpb_events_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_events_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewConfig); i {
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
		file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_events_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Poke); i {
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
		file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_events_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunCreated); i {
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
		file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_events_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunFinished); i {
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
		file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_events_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PurgeCompleted); i {
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
	file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_events_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Event_NewConfig)(nil),
		(*Event_Poke)(nil),
		(*Event_ClsUpdated)(nil),
		(*Event_RunCreated)(nil),
		(*Event_RunFinished)(nil),
		(*Event_PurgeCompleted)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_events_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_events_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_events_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_events_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_cv_internal_prjmanager_prjpb_events_proto = out.File
	file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_events_proto_rawDesc = nil
	file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_events_proto_goTypes = nil
	file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_events_proto_depIdxs = nil
}
