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
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: go.chromium.org/luci/cv/internal/run/eventpb/events.proto

package eventpb

import (
	changelist "go.chromium.org/luci/cv/internal/changelist"
	tryjob "go.chromium.org/luci/cv/internal/tryjob"
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

// Event is a container for all kinds of events a Run Manager can receive.
type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Event:
	//
	//	*Event_Start
	//	*Event_Cancel
	//	*Event_Poke
	//	*Event_NewConfig
	//	*Event_ClsUpdated
	//	*Event_ReadyForSubmission
	//	*Event_ClsSubmitted
	//	*Event_SubmissionCompleted
	//	*Event_LongOpCompleted
	//	*Event_TryjobsUpdated
	Event isEvent_Event `protobuf_oneof:"event"`
	// Instructs Run Manager that this event can only be processed after
	// this timestamp.
	ProcessAfter *timestamppb.Timestamp `protobuf:"bytes,20,opt,name=process_after,json=processAfter,proto3" json:"process_after,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_internal_run_eventpb_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_internal_run_eventpb_events_proto_msgTypes[0]
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
	return file_go_chromium_org_luci_cv_internal_run_eventpb_events_proto_rawDescGZIP(), []int{0}
}

func (m *Event) GetEvent() isEvent_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (x *Event) GetStart() *Start {
	if x, ok := x.GetEvent().(*Event_Start); ok {
		return x.Start
	}
	return nil
}

func (x *Event) GetCancel() *Cancel {
	if x, ok := x.GetEvent().(*Event_Cancel); ok {
		return x.Cancel
	}
	return nil
}

func (x *Event) GetPoke() *Poke {
	if x, ok := x.GetEvent().(*Event_Poke); ok {
		return x.Poke
	}
	return nil
}

func (x *Event) GetNewConfig() *NewConfig {
	if x, ok := x.GetEvent().(*Event_NewConfig); ok {
		return x.NewConfig
	}
	return nil
}

func (x *Event) GetClsUpdated() *changelist.CLUpdatedEvents {
	if x, ok := x.GetEvent().(*Event_ClsUpdated); ok {
		return x.ClsUpdated
	}
	return nil
}

func (x *Event) GetReadyForSubmission() *ReadyForSubmission {
	if x, ok := x.GetEvent().(*Event_ReadyForSubmission); ok {
		return x.ReadyForSubmission
	}
	return nil
}

func (x *Event) GetClsSubmitted() *CLsSubmitted {
	if x, ok := x.GetEvent().(*Event_ClsSubmitted); ok {
		return x.ClsSubmitted
	}
	return nil
}

func (x *Event) GetSubmissionCompleted() *SubmissionCompleted {
	if x, ok := x.GetEvent().(*Event_SubmissionCompleted); ok {
		return x.SubmissionCompleted
	}
	return nil
}

func (x *Event) GetLongOpCompleted() *LongOpCompleted {
	if x, ok := x.GetEvent().(*Event_LongOpCompleted); ok {
		return x.LongOpCompleted
	}
	return nil
}

func (x *Event) GetTryjobsUpdated() *tryjob.TryjobUpdatedEvents {
	if x, ok := x.GetEvent().(*Event_TryjobsUpdated); ok {
		return x.TryjobsUpdated
	}
	return nil
}

func (x *Event) GetProcessAfter() *timestamppb.Timestamp {
	if x != nil {
		return x.ProcessAfter
	}
	return nil
}

type isEvent_Event interface {
	isEvent_Event()
}

type Event_Start struct {
	// On Start event, Run Manager will start the Run.
	//
	// The Run entity must already exist.
	Start *Start `protobuf:"bytes,1,opt,name=start,proto3,oneof"`
}

type Event_Cancel struct {
	// On Cancel event, Run Manager will cancel the Run.
	Cancel *Cancel `protobuf:"bytes,2,opt,name=cancel,proto3,oneof"`
}

type Event_Poke struct {
	// On Poke event, Run Manager will check the state of the Run and perform
	// any action if necessary.
	//
	// Sent periodically by Project Manager.
	Poke *Poke `protobuf:"bytes,3,opt,name=poke,proto3,oneof"`
}

type Event_NewConfig struct {
	// On NewConfig event, Run Manager will update config the Run for the
	// given RunID.
	//
	// Sent by Project Manager, which guarantees these events are sent in order
	// of config updates. See also its `eversion` field.
	NewConfig *NewConfig `protobuf:"bytes,4,opt,name=new_config,json=newConfig,proto3,oneof"`
}

type Event_ClsUpdated struct {
	// On cls_updated event, Run Manager will decide whether to cancel the Run
	// based on the current CL state.
	ClsUpdated *changelist.CLUpdatedEvents `protobuf:"bytes,9,opt,name=cls_updated,json=clsUpdated,proto3,oneof"`
}

type Event_ReadyForSubmission struct {
	// On ReadyForSubmission event, Run Manager will try to submit CLs in
	// this Run.
	ReadyForSubmission *ReadyForSubmission `protobuf:"bytes,6,opt,name=ready_for_submission,json=readyForSubmission,proto3,oneof"`
}

type Event_ClsSubmitted struct {
	// On CLSubmitted event, Run Manager will mark this CL as submitted.
	ClsSubmitted *CLsSubmitted `protobuf:"bytes,12,opt,name=cls_submitted,json=clsSubmitted,proto3,oneof"`
}

type Event_SubmissionCompleted struct {
	// On SubmissionCompleted event, Run Manager will look at the submission
	// result and decide whether to retry this submission or succeed or fail
	// the Run.
	SubmissionCompleted *SubmissionCompleted `protobuf:"bytes,8,opt,name=submission_completed,json=submissionCompleted,proto3,oneof"`
}

type Event_LongOpCompleted struct {
	// On LongOpCompleted event is sent by Run Manager to itself to mark
	// completion of a long operation.
	LongOpCompleted *LongOpCompleted `protobuf:"bytes,10,opt,name=long_op_completed,json=longOpCompleted,proto3,oneof"`
}

type Event_TryjobsUpdated struct {
	// On tryjobs_updated event, Run Manager decides the next step for those
	// tryjobs (e.g. retry) and this Run (e.g. fail the Run for critical
	// tryjob failure).
	TryjobsUpdated *tryjob.TryjobUpdatedEvents `protobuf:"bytes,11,opt,name=tryjobs_updated,json=tryjobsUpdated,proto3,oneof"`
}

func (*Event_Start) isEvent_Event() {}

func (*Event_Cancel) isEvent_Event() {}

func (*Event_Poke) isEvent_Event() {}

func (*Event_NewConfig) isEvent_Event() {}

func (*Event_ClsUpdated) isEvent_Event() {}

func (*Event_ReadyForSubmission) isEvent_Event() {}

func (*Event_ClsSubmitted) isEvent_Event() {}

func (*Event_SubmissionCompleted) isEvent_Event() {}

func (*Event_LongOpCompleted) isEvent_Event() {}

func (*Event_TryjobsUpdated) isEvent_Event() {}

type Start struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Start) Reset() {
	*x = Start{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_internal_run_eventpb_events_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Start) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Start) ProtoMessage() {}

func (x *Start) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_internal_run_eventpb_events_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Start.ProtoReflect.Descriptor instead.
func (*Start) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_internal_run_eventpb_events_proto_rawDescGZIP(), []int{1}
}

type Cancel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// reason for cancelling the run.
	Reason string `protobuf:"bytes,1,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (x *Cancel) Reset() {
	*x = Cancel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_internal_run_eventpb_events_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cancel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cancel) ProtoMessage() {}

func (x *Cancel) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_internal_run_eventpb_events_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cancel.ProtoReflect.Descriptor instead.
func (*Cancel) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_internal_run_eventpb_events_proto_rawDescGZIP(), []int{2}
}

func (x *Cancel) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

type Poke struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Poke) Reset() {
	*x = Poke{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_internal_run_eventpb_events_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Poke) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Poke) ProtoMessage() {}

func (x *Poke) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_internal_run_eventpb_events_proto_msgTypes[3]
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
	return file_go_chromium_org_luci_cv_internal_run_eventpb_events_proto_rawDescGZIP(), []int{3}
}

type NewConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Hash identifying project config version to update to.
	Hash string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	// Eversion of the project config version identify by hash.
	//
	// Provided for identifying the latest NewConfig message
	// if there are more than one outstanding NewConfig event.
	Eversion int64 `protobuf:"varint,2,opt,name=eversion,proto3" json:"eversion,omitempty"`
}

func (x *NewConfig) Reset() {
	*x = NewConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_internal_run_eventpb_events_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewConfig) ProtoMessage() {}

func (x *NewConfig) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_internal_run_eventpb_events_proto_msgTypes[4]
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
	return file_go_chromium_org_luci_cv_internal_run_eventpb_events_proto_rawDescGZIP(), []int{4}
}

func (x *NewConfig) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *NewConfig) GetEversion() int64 {
	if x != nil {
		return x.Eversion
	}
	return 0
}

type ReadyForSubmission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReadyForSubmission) Reset() {
	*x = ReadyForSubmission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_internal_run_eventpb_events_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadyForSubmission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadyForSubmission) ProtoMessage() {}

func (x *ReadyForSubmission) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_internal_run_eventpb_events_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadyForSubmission.ProtoReflect.Descriptor instead.
func (*ReadyForSubmission) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_internal_run_eventpb_events_proto_rawDescGZIP(), []int{5}
}

type CLsSubmitted struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Internal ID of CLs that were submitted successfully.
	Clids []int64 `protobuf:"varint,1,rep,packed,name=clids,proto3" json:"clids,omitempty"`
}

func (x *CLsSubmitted) Reset() {
	*x = CLsSubmitted{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_internal_run_eventpb_events_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CLsSubmitted) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CLsSubmitted) ProtoMessage() {}

func (x *CLsSubmitted) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_internal_run_eventpb_events_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CLsSubmitted.ProtoReflect.Descriptor instead.
func (*CLsSubmitted) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_internal_run_eventpb_events_proto_rawDescGZIP(), []int{6}
}

func (x *CLsSubmitted) GetClids() []int64 {
	if x != nil {
		return x.Clids
	}
	return nil
}

var File_go_chromium_org_luci_cv_internal_run_eventpb_events_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_cv_internal_run_eventpb_events_proto_rawDesc = []byte{
	0x0a, 0x39, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x76, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x72, 0x75, 0x6e, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x63, 0x76, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x39, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69,
	0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x76, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x69,
	0x73, 0x74, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x35, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x76, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x74, 0x72, 0x79, 0x6a, 0x6f, 0x62, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x39, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f,
	0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x76,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x72, 0x75, 0x6e, 0x2f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x6c, 0x6f, 0x6e, 0x67, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x3d, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e,
	0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x76, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x72, 0x75, 0x6e, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x70, 0x62,
	0x2f, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xda, 0x06, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x36, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x76, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x48, 0x00, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x12, 0x39, 0x0a, 0x06, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x76, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x43, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x48, 0x00, 0x52, 0x06, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x12, 0x33,
	0x0a, 0x04, 0x70, 0x6f, 0x6b, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63,
	0x76, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x50, 0x6f, 0x6b, 0x65, 0x48, 0x00, 0x52, 0x04, 0x70,
	0x6f, 0x6b, 0x65, 0x12, 0x43, 0x0a, 0x0a, 0x6e, 0x65, 0x77, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x76, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x70,
	0x62, 0x2e, 0x4e, 0x65, 0x77, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x09, 0x6e,
	0x65, 0x77, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x4a, 0x0a, 0x0b, 0x63, 0x6c, 0x73, 0x5f,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x63, 0x76, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x43, 0x4c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x48, 0x00, 0x52, 0x0a, 0x63, 0x6c, 0x73, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x12, 0x5f, 0x0a, 0x14, 0x72, 0x65, 0x61, 0x64, 0x79, 0x5f, 0x66, 0x6f,
	0x72, 0x5f, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x63, 0x76, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x61,
	0x64, 0x79, 0x46, 0x6f, 0x72, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x48,
	0x00, 0x52, 0x12, 0x72, 0x65, 0x61, 0x64, 0x79, 0x46, 0x6f, 0x72, 0x53, 0x75, 0x62, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x4c, 0x0a, 0x0d, 0x63, 0x6c, 0x73, 0x5f, 0x73, 0x75, 0x62,
	0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63,
	0x76, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x43, 0x4c, 0x73, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74,
	0x74, 0x65, 0x64, 0x48, 0x00, 0x52, 0x0c, 0x63, 0x6c, 0x73, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74,
	0x74, 0x65, 0x64, 0x12, 0x61, 0x0a, 0x14, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2c, 0x2e, 0x63, 0x76, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e,
	0x72, 0x75, 0x6e, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x53, 0x75, 0x62, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x48,
	0x00, 0x52, 0x13, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x56, 0x0a, 0x11, 0x6c, 0x6f, 0x6e, 0x67, 0x5f, 0x6f,
	0x70, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x28, 0x2e, 0x63, 0x76, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e,
	0x72, 0x75, 0x6e, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x6e, 0x67,
	0x4f, 0x70, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x48, 0x00, 0x52, 0x0f, 0x6c,
	0x6f, 0x6e, 0x67, 0x4f, 0x70, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x52,
	0x0a, 0x0f, 0x74, 0x72, 0x79, 0x6a, 0x6f, 0x62, 0x73, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x76, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x74, 0x72, 0x79, 0x6a, 0x6f, 0x62, 0x2e, 0x54, 0x72, 0x79,
	0x6a, 0x6f, 0x62, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x48, 0x00, 0x52, 0x0e, 0x74, 0x72, 0x79, 0x6a, 0x6f, 0x62, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x12, 0x3f, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x61, 0x66,
	0x74, 0x65, 0x72, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x41, 0x66,
	0x74, 0x65, 0x72, 0x42, 0x07, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4a, 0x04, 0x08, 0x07,
	0x10, 0x08, 0x4a, 0x04, 0x08, 0x1e, 0x10, 0x1f, 0x4a, 0x04, 0x08, 0x20, 0x10, 0x21, 0x22, 0x07,
	0x0a, 0x05, 0x53, 0x74, 0x61, 0x72, 0x74, 0x22, 0x20, 0x0a, 0x06, 0x43, 0x61, 0x6e, 0x63, 0x65,
	0x6c, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0x06, 0x0a, 0x04, 0x50, 0x6f, 0x6b,
	0x65, 0x22, 0x3b, 0x0a, 0x09, 0x4e, 0x65, 0x77, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12,
	0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61,
	0x73, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x65, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x14,
	0x0a, 0x12, 0x52, 0x65, 0x61, 0x64, 0x79, 0x46, 0x6f, 0x72, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x22, 0x24, 0x0a, 0x0c, 0x43, 0x4c, 0x73, 0x53, 0x75, 0x62, 0x6d, 0x69,
	0x74, 0x74, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6c, 0x69, 0x64, 0x73, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x6f,
	0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75,
	0x63, 0x69, 0x2f, 0x63, 0x76, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x72,
	0x75, 0x6e, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x3b, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_cv_internal_run_eventpb_events_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_cv_internal_run_eventpb_events_proto_rawDescData = file_go_chromium_org_luci_cv_internal_run_eventpb_events_proto_rawDesc
)

func file_go_chromium_org_luci_cv_internal_run_eventpb_events_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_cv_internal_run_eventpb_events_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_cv_internal_run_eventpb_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_cv_internal_run_eventpb_events_proto_rawDescData)
	})
	return file_go_chromium_org_luci_cv_internal_run_eventpb_events_proto_rawDescData
}

var file_go_chromium_org_luci_cv_internal_run_eventpb_events_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_go_chromium_org_luci_cv_internal_run_eventpb_events_proto_goTypes = []interface{}{
	(*Event)(nil),                      // 0: cv.internal.run.eventpb.Event
	(*Start)(nil),                      // 1: cv.internal.run.eventpb.Start
	(*Cancel)(nil),                     // 2: cv.internal.run.eventpb.Cancel
	(*Poke)(nil),                       // 3: cv.internal.run.eventpb.Poke
	(*NewConfig)(nil),                  // 4: cv.internal.run.eventpb.NewConfig
	(*ReadyForSubmission)(nil),         // 5: cv.internal.run.eventpb.ReadyForSubmission
	(*CLsSubmitted)(nil),               // 6: cv.internal.run.eventpb.CLsSubmitted
	(*changelist.CLUpdatedEvents)(nil), // 7: cv.internal.changelist.CLUpdatedEvents
	(*SubmissionCompleted)(nil),        // 8: cv.internal.run.eventpb.SubmissionCompleted
	(*LongOpCompleted)(nil),            // 9: cv.internal.run.eventpb.LongOpCompleted
	(*tryjob.TryjobUpdatedEvents)(nil), // 10: cv.internal.tryjob.TryjobUpdatedEvents
	(*timestamppb.Timestamp)(nil),      // 11: google.protobuf.Timestamp
}
var file_go_chromium_org_luci_cv_internal_run_eventpb_events_proto_depIdxs = []int32{
	1,  // 0: cv.internal.run.eventpb.Event.start:type_name -> cv.internal.run.eventpb.Start
	2,  // 1: cv.internal.run.eventpb.Event.cancel:type_name -> cv.internal.run.eventpb.Cancel
	3,  // 2: cv.internal.run.eventpb.Event.poke:type_name -> cv.internal.run.eventpb.Poke
	4,  // 3: cv.internal.run.eventpb.Event.new_config:type_name -> cv.internal.run.eventpb.NewConfig
	7,  // 4: cv.internal.run.eventpb.Event.cls_updated:type_name -> cv.internal.changelist.CLUpdatedEvents
	5,  // 5: cv.internal.run.eventpb.Event.ready_for_submission:type_name -> cv.internal.run.eventpb.ReadyForSubmission
	6,  // 6: cv.internal.run.eventpb.Event.cls_submitted:type_name -> cv.internal.run.eventpb.CLsSubmitted
	8,  // 7: cv.internal.run.eventpb.Event.submission_completed:type_name -> cv.internal.run.eventpb.SubmissionCompleted
	9,  // 8: cv.internal.run.eventpb.Event.long_op_completed:type_name -> cv.internal.run.eventpb.LongOpCompleted
	10, // 9: cv.internal.run.eventpb.Event.tryjobs_updated:type_name -> cv.internal.tryjob.TryjobUpdatedEvents
	11, // 10: cv.internal.run.eventpb.Event.process_after:type_name -> google.protobuf.Timestamp
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_cv_internal_run_eventpb_events_proto_init() }
func file_go_chromium_org_luci_cv_internal_run_eventpb_events_proto_init() {
	if File_go_chromium_org_luci_cv_internal_run_eventpb_events_proto != nil {
		return
	}
	file_go_chromium_org_luci_cv_internal_run_eventpb_longop_proto_init()
	file_go_chromium_org_luci_cv_internal_run_eventpb_submission_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_cv_internal_run_eventpb_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_go_chromium_org_luci_cv_internal_run_eventpb_events_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Start); i {
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
		file_go_chromium_org_luci_cv_internal_run_eventpb_events_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cancel); i {
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
		file_go_chromium_org_luci_cv_internal_run_eventpb_events_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_go_chromium_org_luci_cv_internal_run_eventpb_events_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_go_chromium_org_luci_cv_internal_run_eventpb_events_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadyForSubmission); i {
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
		file_go_chromium_org_luci_cv_internal_run_eventpb_events_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CLsSubmitted); i {
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
	file_go_chromium_org_luci_cv_internal_run_eventpb_events_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Event_Start)(nil),
		(*Event_Cancel)(nil),
		(*Event_Poke)(nil),
		(*Event_NewConfig)(nil),
		(*Event_ClsUpdated)(nil),
		(*Event_ReadyForSubmission)(nil),
		(*Event_ClsSubmitted)(nil),
		(*Event_SubmissionCompleted)(nil),
		(*Event_LongOpCompleted)(nil),
		(*Event_TryjobsUpdated)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_go_chromium_org_luci_cv_internal_run_eventpb_events_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_cv_internal_run_eventpb_events_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_cv_internal_run_eventpb_events_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_cv_internal_run_eventpb_events_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_cv_internal_run_eventpb_events_proto = out.File
	file_go_chromium_org_luci_cv_internal_run_eventpb_events_proto_rawDesc = nil
	file_go_chromium_org_luci_cv_internal_run_eventpb_events_proto_goTypes = nil
	file_go_chromium_org_luci_cv_internal_run_eventpb_events_proto_depIdxs = nil
}
