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
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: go.chromium.org/luci/cv/internal/run/eventpb/longop.proto

package eventpb

import (
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

type LongOpCompleted_Status int32

const (
	LongOpCompleted_LONG_OP_STATUS_UNSPECIFIED LongOpCompleted_Status = 0
	// SUCCEEDED means the long operation succeeded.
	LongOpCompleted_SUCCEEDED LongOpCompleted_Status = 1
	// FAILED means the long operation experienced a failure.
	LongOpCompleted_FAILED LongOpCompleted_Status = 2
	// Cancelled is true if the LongOp detected that it was cancelled by the Run
	// Manager and thus stopped its working before completion.
	LongOpCompleted_CANCELLED LongOpCompleted_Status = 3
	// Expired means the long operation expired.
	//
	// If this is the case, the `result` field may be unset: this should be rare
	// but may happen if Run Manager detects expiry before a LongOpCompleted
	// event is sent by the long operation handling task.
	LongOpCompleted_EXPIRED LongOpCompleted_Status = 4
)

// Enum value maps for LongOpCompleted_Status.
var (
	LongOpCompleted_Status_name = map[int32]string{
		0: "LONG_OP_STATUS_UNSPECIFIED",
		1: "SUCCEEDED",
		2: "FAILED",
		3: "CANCELLED",
		4: "EXPIRED",
	}
	LongOpCompleted_Status_value = map[string]int32{
		"LONG_OP_STATUS_UNSPECIFIED": 0,
		"SUCCEEDED":                  1,
		"FAILED":                     2,
		"CANCELLED":                  3,
		"EXPIRED":                    4,
	}
)

func (x LongOpCompleted_Status) Enum() *LongOpCompleted_Status {
	p := new(LongOpCompleted_Status)
	*p = x
	return p
}

func (x LongOpCompleted_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LongOpCompleted_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_go_chromium_org_luci_cv_internal_run_eventpb_longop_proto_enumTypes[0].Descriptor()
}

func (LongOpCompleted_Status) Type() protoreflect.EnumType {
	return &file_go_chromium_org_luci_cv_internal_run_eventpb_longop_proto_enumTypes[0]
}

func (x LongOpCompleted_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LongOpCompleted_Status.Descriptor instead.
func (LongOpCompleted_Status) EnumDescriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_internal_run_eventpb_longop_proto_rawDescGZIP(), []int{0, 0}
}

type LongOpCompleted struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Long Operation ID.
	OperationId string `protobuf:"bytes,1,opt,name=operation_id,json=operationId,proto3" json:"operation_id,omitempty"`
	// Status of the long operation.
	Status LongOpCompleted_Status `protobuf:"varint,2,opt,name=status,proto3,enum=cv.internal.run.eventpb.LongOpCompleted_Status" json:"status,omitempty"`
	// Types that are assignable to Result:
	//
	//	*LongOpCompleted_PostStartMessage_
	//	*LongOpCompleted_CancelTriggers_
	//	*LongOpCompleted_ExecuteTryjobs
	Result isLongOpCompleted_Result `protobuf_oneof:"result"`
}

func (x *LongOpCompleted) Reset() {
	*x = LongOpCompleted{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_internal_run_eventpb_longop_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LongOpCompleted) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LongOpCompleted) ProtoMessage() {}

func (x *LongOpCompleted) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_internal_run_eventpb_longop_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LongOpCompleted.ProtoReflect.Descriptor instead.
func (*LongOpCompleted) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_internal_run_eventpb_longop_proto_rawDescGZIP(), []int{0}
}

func (x *LongOpCompleted) GetOperationId() string {
	if x != nil {
		return x.OperationId
	}
	return ""
}

func (x *LongOpCompleted) GetStatus() LongOpCompleted_Status {
	if x != nil {
		return x.Status
	}
	return LongOpCompleted_LONG_OP_STATUS_UNSPECIFIED
}

func (m *LongOpCompleted) GetResult() isLongOpCompleted_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (x *LongOpCompleted) GetPostStartMessage() *LongOpCompleted_PostStartMessage {
	if x, ok := x.GetResult().(*LongOpCompleted_PostStartMessage_); ok {
		return x.PostStartMessage
	}
	return nil
}

func (x *LongOpCompleted) GetCancelTriggers() *LongOpCompleted_CancelTriggers {
	if x, ok := x.GetResult().(*LongOpCompleted_CancelTriggers_); ok {
		return x.CancelTriggers
	}
	return nil
}

func (x *LongOpCompleted) GetExecuteTryjobs() *tryjob.ExecuteTryjobsResult {
	if x, ok := x.GetResult().(*LongOpCompleted_ExecuteTryjobs); ok {
		return x.ExecuteTryjobs
	}
	return nil
}

type isLongOpCompleted_Result interface {
	isLongOpCompleted_Result()
}

type LongOpCompleted_PostStartMessage_ struct {
	PostStartMessage *LongOpCompleted_PostStartMessage `protobuf:"bytes,3,opt,name=post_start_message,json=postStartMessage,proto3,oneof"`
}

type LongOpCompleted_CancelTriggers_ struct {
	CancelTriggers *LongOpCompleted_CancelTriggers `protobuf:"bytes,4,opt,name=cancel_triggers,json=cancelTriggers,proto3,oneof"`
}

type LongOpCompleted_ExecuteTryjobs struct {
	ExecuteTryjobs *tryjob.ExecuteTryjobsResult `protobuf:"bytes,5,opt,name=execute_tryjobs,json=executeTryjobs,proto3,oneof"`
}

func (*LongOpCompleted_PostStartMessage_) isLongOpCompleted_Result() {}

func (*LongOpCompleted_CancelTriggers_) isLongOpCompleted_Result() {}

func (*LongOpCompleted_ExecuteTryjobs) isLongOpCompleted_Result() {}

type LongOpCompleted_PostStartMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// CLIDs on which the message was posted.
	Posted []int64 `protobuf:"varint,1,rep,packed,name=posted,proto3" json:"posted,omitempty"`
	// Maps CLID to the permanent error.
	PermanentErrors map[int64]string `protobuf:"bytes,2,rep,name=permanent_errors,json=permanentErrors,proto3" json:"permanent_errors,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Approximate time when CV became certain that the message was posted on
	// the last CL.
	Time *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *LongOpCompleted_PostStartMessage) Reset() {
	*x = LongOpCompleted_PostStartMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_internal_run_eventpb_longop_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LongOpCompleted_PostStartMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LongOpCompleted_PostStartMessage) ProtoMessage() {}

func (x *LongOpCompleted_PostStartMessage) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_internal_run_eventpb_longop_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LongOpCompleted_PostStartMessage.ProtoReflect.Descriptor instead.
func (*LongOpCompleted_PostStartMessage) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_internal_run_eventpb_longop_proto_rawDescGZIP(), []int{0, 0}
}

func (x *LongOpCompleted_PostStartMessage) GetPosted() []int64 {
	if x != nil {
		return x.Posted
	}
	return nil
}

func (x *LongOpCompleted_PostStartMessage) GetPermanentErrors() map[int64]string {
	if x != nil {
		return x.PermanentErrors
	}
	return nil
}

func (x *LongOpCompleted_PostStartMessage) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

type LongOpCompleted_CancelTriggers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The top-level long ops status will be SUCCEEDED iff all triggers are
	// cancelled successfully, in other word, all results have `cancelled_at`
	// set.
	Results []*LongOpCompleted_CancelTriggers_Result `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *LongOpCompleted_CancelTriggers) Reset() {
	*x = LongOpCompleted_CancelTriggers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_internal_run_eventpb_longop_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LongOpCompleted_CancelTriggers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LongOpCompleted_CancelTriggers) ProtoMessage() {}

func (x *LongOpCompleted_CancelTriggers) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_internal_run_eventpb_longop_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LongOpCompleted_CancelTriggers.ProtoReflect.Descriptor instead.
func (*LongOpCompleted_CancelTriggers) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_internal_run_eventpb_longop_proto_rawDescGZIP(), []int{0, 1}
}

func (x *LongOpCompleted_CancelTriggers) GetResults() []*LongOpCompleted_CancelTriggers_Result {
	if x != nil {
		return x.Results
	}
	return nil
}

type LongOpCompleted_CancelTriggers_Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the CL that this long op is trying to cancel its trigger.
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// The external id of the CL.
	ExternalId string `protobuf:"bytes,2,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
	// Types that are assignable to Detail:
	//
	//	*LongOpCompleted_CancelTriggers_Result_SuccessInfo
	//	*LongOpCompleted_CancelTriggers_Result_FailureInfo
	Detail isLongOpCompleted_CancelTriggers_Result_Detail `protobuf_oneof:"detail"`
}

func (x *LongOpCompleted_CancelTriggers_Result) Reset() {
	*x = LongOpCompleted_CancelTriggers_Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_internal_run_eventpb_longop_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LongOpCompleted_CancelTriggers_Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LongOpCompleted_CancelTriggers_Result) ProtoMessage() {}

func (x *LongOpCompleted_CancelTriggers_Result) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_internal_run_eventpb_longop_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LongOpCompleted_CancelTriggers_Result.ProtoReflect.Descriptor instead.
func (*LongOpCompleted_CancelTriggers_Result) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_internal_run_eventpb_longop_proto_rawDescGZIP(), []int{0, 1, 0}
}

func (x *LongOpCompleted_CancelTriggers_Result) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LongOpCompleted_CancelTriggers_Result) GetExternalId() string {
	if x != nil {
		return x.ExternalId
	}
	return ""
}

func (m *LongOpCompleted_CancelTriggers_Result) GetDetail() isLongOpCompleted_CancelTriggers_Result_Detail {
	if m != nil {
		return m.Detail
	}
	return nil
}

func (x *LongOpCompleted_CancelTriggers_Result) GetSuccessInfo() *LongOpCompleted_CancelTriggers_Result_Success {
	if x, ok := x.GetDetail().(*LongOpCompleted_CancelTriggers_Result_SuccessInfo); ok {
		return x.SuccessInfo
	}
	return nil
}

func (x *LongOpCompleted_CancelTriggers_Result) GetFailureInfo() *LongOpCompleted_CancelTriggers_Result_Failure {
	if x, ok := x.GetDetail().(*LongOpCompleted_CancelTriggers_Result_FailureInfo); ok {
		return x.FailureInfo
	}
	return nil
}

type isLongOpCompleted_CancelTriggers_Result_Detail interface {
	isLongOpCompleted_CancelTriggers_Result_Detail()
}

type LongOpCompleted_CancelTriggers_Result_SuccessInfo struct {
	SuccessInfo *LongOpCompleted_CancelTriggers_Result_Success `protobuf:"bytes,3,opt,name=success_info,json=successInfo,proto3,oneof"`
}

type LongOpCompleted_CancelTriggers_Result_FailureInfo struct {
	FailureInfo *LongOpCompleted_CancelTriggers_Result_Failure `protobuf:"bytes,4,opt,name=failure_info,json=failureInfo,proto3,oneof"`
}

func (*LongOpCompleted_CancelTriggers_Result_SuccessInfo) isLongOpCompleted_CancelTriggers_Result_Detail() {
}

func (*LongOpCompleted_CancelTriggers_Result_FailureInfo) isLongOpCompleted_CancelTriggers_Result_Detail() {
}

type LongOpCompleted_CancelTriggers_Result_Success struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The timestamp when trigger is successfully cancelled from LUCI CV
	// PoV.
	//
	// It is possible by the time LUCI CV is trying to cancel the trigger,
	// it has already been removed by the user. Therefore, this timestamp
	// is the   time when LUCI CV observes that trigger is removed.
	CancelledAt *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=cancelled_at,json=cancelledAt,proto3" json:"cancelled_at,omitempty"`
}

func (x *LongOpCompleted_CancelTriggers_Result_Success) Reset() {
	*x = LongOpCompleted_CancelTriggers_Result_Success{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_internal_run_eventpb_longop_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LongOpCompleted_CancelTriggers_Result_Success) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LongOpCompleted_CancelTriggers_Result_Success) ProtoMessage() {}

func (x *LongOpCompleted_CancelTriggers_Result_Success) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_internal_run_eventpb_longop_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LongOpCompleted_CancelTriggers_Result_Success.ProtoReflect.Descriptor instead.
func (*LongOpCompleted_CancelTriggers_Result_Success) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_internal_run_eventpb_longop_proto_rawDescGZIP(), []int{0, 1, 0, 0}
}

func (x *LongOpCompleted_CancelTriggers_Result_Success) GetCancelledAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CancelledAt
	}
	return nil
}

type LongOpCompleted_CancelTriggers_Result_Failure struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The message when CV failed to cancel the trigger.
	FailureMessage string `protobuf:"bytes,1,opt,name=failure_message,json=failureMessage,proto3" json:"failure_message,omitempty"`
}

func (x *LongOpCompleted_CancelTriggers_Result_Failure) Reset() {
	*x = LongOpCompleted_CancelTriggers_Result_Failure{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_internal_run_eventpb_longop_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LongOpCompleted_CancelTriggers_Result_Failure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LongOpCompleted_CancelTriggers_Result_Failure) ProtoMessage() {}

func (x *LongOpCompleted_CancelTriggers_Result_Failure) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_internal_run_eventpb_longop_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LongOpCompleted_CancelTriggers_Result_Failure.ProtoReflect.Descriptor instead.
func (*LongOpCompleted_CancelTriggers_Result_Failure) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_internal_run_eventpb_longop_proto_rawDescGZIP(), []int{0, 1, 0, 1}
}

func (x *LongOpCompleted_CancelTriggers_Result_Failure) GetFailureMessage() string {
	if x != nil {
		return x.FailureMessage
	}
	return ""
}

var File_go_chromium_org_luci_cv_internal_run_eventpb_longop_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_cv_internal_run_eventpb_longop_proto_rawDesc = []byte{
	0x0a, 0x39, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x76, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x72, 0x75, 0x6e, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x6c,
	0x6f, 0x6e, 0x67, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x63, 0x76, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69,
	0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x76, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x72, 0x79, 0x6a, 0x6f, 0x62, 0x2f, 0x74,
	0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb3, 0x0a, 0x0a, 0x0f, 0x4c, 0x6f,
	0x6e, 0x67, 0x4f, 0x70, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x21, 0x0a,
	0x0c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x47, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x2f, 0x2e, 0x63, 0x76, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x72,
	0x75, 0x6e, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x6e, 0x67, 0x4f,
	0x70, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x69, 0x0a, 0x12, 0x70, 0x6f, 0x73,
	0x74, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x63, 0x76, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e,
	0x4c, 0x6f, 0x6e, 0x67, 0x4f, 0x70, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x2e,
	0x50, 0x6f, 0x73, 0x74, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x48, 0x00, 0x52, 0x10, 0x70, 0x6f, 0x73, 0x74, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x62, 0x0a, 0x0f, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x5f, 0x74,
	0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e,
	0x63, 0x76, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x72, 0x75, 0x6e, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x6e, 0x67, 0x4f, 0x70, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x54, 0x72,
	0x69, 0x67, 0x67, 0x65, 0x72, 0x73, 0x48, 0x00, 0x52, 0x0e, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x73, 0x12, 0x53, 0x0a, 0x0f, 0x65, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x65, 0x5f, 0x74, 0x72, 0x79, 0x6a, 0x6f, 0x62, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x28, 0x2e, 0x63, 0x76, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e,
	0x74, 0x72, 0x79, 0x6a, 0x6f, 0x62, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x54, 0x72,
	0x79, 0x6a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48, 0x00, 0x52, 0x0e, 0x65,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x54, 0x72, 0x79, 0x6a, 0x6f, 0x62, 0x73, 0x1a, 0x99, 0x02,
	0x0a, 0x10, 0x50, 0x6f, 0x73, 0x74, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x03, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x64, 0x12, 0x79, 0x0a, 0x10, 0x70, 0x65,
	0x72, 0x6d, 0x61, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x4e, 0x2e, 0x63, 0x76, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x4c,
	0x6f, 0x6e, 0x67, 0x4f, 0x70, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x2e, 0x50,
	0x6f, 0x73, 0x74, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x50, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x65, 0x6e, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0f, 0x70, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x65, 0x6e, 0x74, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x1a, 0x42, 0x0a, 0x14, 0x50, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x65,
	0x6e, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x88, 0x04, 0x0a, 0x0e, 0x43, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x73, 0x12, 0x58, 0x0a, 0x07,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3e, 0x2e,
	0x63, 0x76, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x72, 0x75, 0x6e, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x6e, 0x67, 0x4f, 0x70, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x54, 0x72,
	0x69, 0x67, 0x67, 0x65, 0x72, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x1a, 0x9b, 0x03, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x49, 0x64, 0x12, 0x6b, 0x0a, 0x0c, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x46, 0x2e, 0x63, 0x76, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x6e, 0x67, 0x4f, 0x70, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72,
	0x73, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x48, 0x00, 0x52, 0x0b, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x6b, 0x0a, 0x0c, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x46, 0x2e, 0x63, 0x76, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e,
	0x4c, 0x6f, 0x6e, 0x67, 0x4f, 0x70, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x2e,
	0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x73, 0x2e, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x48, 0x00, 0x52,
	0x0b, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x48, 0x0a, 0x07,
	0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x3d, 0x0a, 0x0c, 0x63, 0x61, 0x6e, 0x63, 0x65,
	0x6c, 0x6c, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x63, 0x61, 0x6e, 0x63, 0x65,
	0x6c, 0x6c, 0x65, 0x64, 0x41, 0x74, 0x1a, 0x32, 0x0a, 0x07, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72,
	0x65, 0x12, 0x27, 0x0a, 0x0f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x66, 0x61, 0x69, 0x6c,
	0x75, 0x72, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x22, 0x5f, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e,
	0x0a, 0x1a, 0x4c, 0x4f, 0x4e, 0x47, 0x5f, 0x4f, 0x50, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0d,
	0x0a, 0x09, 0x53, 0x55, 0x43, 0x43, 0x45, 0x45, 0x44, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0a, 0x0a,
	0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x41, 0x4e,
	0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x58, 0x50, 0x49,
	0x52, 0x45, 0x44, 0x10, 0x04, 0x42, 0x08, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42,
	0x36, 0x5a, 0x34, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x76, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x72, 0x75, 0x6e, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x3b,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_cv_internal_run_eventpb_longop_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_cv_internal_run_eventpb_longop_proto_rawDescData = file_go_chromium_org_luci_cv_internal_run_eventpb_longop_proto_rawDesc
)

func file_go_chromium_org_luci_cv_internal_run_eventpb_longop_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_cv_internal_run_eventpb_longop_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_cv_internal_run_eventpb_longop_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_cv_internal_run_eventpb_longop_proto_rawDescData)
	})
	return file_go_chromium_org_luci_cv_internal_run_eventpb_longop_proto_rawDescData
}

var file_go_chromium_org_luci_cv_internal_run_eventpb_longop_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_go_chromium_org_luci_cv_internal_run_eventpb_longop_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_go_chromium_org_luci_cv_internal_run_eventpb_longop_proto_goTypes = []interface{}{
	(LongOpCompleted_Status)(0),                           // 0: cv.internal.run.eventpb.LongOpCompleted.Status
	(*LongOpCompleted)(nil),                               // 1: cv.internal.run.eventpb.LongOpCompleted
	(*LongOpCompleted_PostStartMessage)(nil),              // 2: cv.internal.run.eventpb.LongOpCompleted.PostStartMessage
	(*LongOpCompleted_CancelTriggers)(nil),                // 3: cv.internal.run.eventpb.LongOpCompleted.CancelTriggers
	nil,                                                   // 4: cv.internal.run.eventpb.LongOpCompleted.PostStartMessage.PermanentErrorsEntry
	(*LongOpCompleted_CancelTriggers_Result)(nil),         // 5: cv.internal.run.eventpb.LongOpCompleted.CancelTriggers.Result
	(*LongOpCompleted_CancelTriggers_Result_Success)(nil), // 6: cv.internal.run.eventpb.LongOpCompleted.CancelTriggers.Result.Success
	(*LongOpCompleted_CancelTriggers_Result_Failure)(nil), // 7: cv.internal.run.eventpb.LongOpCompleted.CancelTriggers.Result.Failure
	(*tryjob.ExecuteTryjobsResult)(nil),                   // 8: cv.internal.tryjob.ExecuteTryjobsResult
	(*timestamppb.Timestamp)(nil),                         // 9: google.protobuf.Timestamp
}
var file_go_chromium_org_luci_cv_internal_run_eventpb_longop_proto_depIdxs = []int32{
	0,  // 0: cv.internal.run.eventpb.LongOpCompleted.status:type_name -> cv.internal.run.eventpb.LongOpCompleted.Status
	2,  // 1: cv.internal.run.eventpb.LongOpCompleted.post_start_message:type_name -> cv.internal.run.eventpb.LongOpCompleted.PostStartMessage
	3,  // 2: cv.internal.run.eventpb.LongOpCompleted.cancel_triggers:type_name -> cv.internal.run.eventpb.LongOpCompleted.CancelTriggers
	8,  // 3: cv.internal.run.eventpb.LongOpCompleted.execute_tryjobs:type_name -> cv.internal.tryjob.ExecuteTryjobsResult
	4,  // 4: cv.internal.run.eventpb.LongOpCompleted.PostStartMessage.permanent_errors:type_name -> cv.internal.run.eventpb.LongOpCompleted.PostStartMessage.PermanentErrorsEntry
	9,  // 5: cv.internal.run.eventpb.LongOpCompleted.PostStartMessage.time:type_name -> google.protobuf.Timestamp
	5,  // 6: cv.internal.run.eventpb.LongOpCompleted.CancelTriggers.results:type_name -> cv.internal.run.eventpb.LongOpCompleted.CancelTriggers.Result
	6,  // 7: cv.internal.run.eventpb.LongOpCompleted.CancelTriggers.Result.success_info:type_name -> cv.internal.run.eventpb.LongOpCompleted.CancelTriggers.Result.Success
	7,  // 8: cv.internal.run.eventpb.LongOpCompleted.CancelTriggers.Result.failure_info:type_name -> cv.internal.run.eventpb.LongOpCompleted.CancelTriggers.Result.Failure
	9,  // 9: cv.internal.run.eventpb.LongOpCompleted.CancelTriggers.Result.Success.cancelled_at:type_name -> google.protobuf.Timestamp
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_cv_internal_run_eventpb_longop_proto_init() }
func file_go_chromium_org_luci_cv_internal_run_eventpb_longop_proto_init() {
	if File_go_chromium_org_luci_cv_internal_run_eventpb_longop_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_cv_internal_run_eventpb_longop_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LongOpCompleted); i {
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
		file_go_chromium_org_luci_cv_internal_run_eventpb_longop_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LongOpCompleted_PostStartMessage); i {
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
		file_go_chromium_org_luci_cv_internal_run_eventpb_longop_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LongOpCompleted_CancelTriggers); i {
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
		file_go_chromium_org_luci_cv_internal_run_eventpb_longop_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LongOpCompleted_CancelTriggers_Result); i {
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
		file_go_chromium_org_luci_cv_internal_run_eventpb_longop_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LongOpCompleted_CancelTriggers_Result_Success); i {
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
		file_go_chromium_org_luci_cv_internal_run_eventpb_longop_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LongOpCompleted_CancelTriggers_Result_Failure); i {
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
	file_go_chromium_org_luci_cv_internal_run_eventpb_longop_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*LongOpCompleted_PostStartMessage_)(nil),
		(*LongOpCompleted_CancelTriggers_)(nil),
		(*LongOpCompleted_ExecuteTryjobs)(nil),
	}
	file_go_chromium_org_luci_cv_internal_run_eventpb_longop_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*LongOpCompleted_CancelTriggers_Result_SuccessInfo)(nil),
		(*LongOpCompleted_CancelTriggers_Result_FailureInfo)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_go_chromium_org_luci_cv_internal_run_eventpb_longop_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_cv_internal_run_eventpb_longop_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_cv_internal_run_eventpb_longop_proto_depIdxs,
		EnumInfos:         file_go_chromium_org_luci_cv_internal_run_eventpb_longop_proto_enumTypes,
		MessageInfos:      file_go_chromium_org_luci_cv_internal_run_eventpb_longop_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_cv_internal_run_eventpb_longop_proto = out.File
	file_go_chromium_org_luci_cv_internal_run_eventpb_longop_proto_rawDesc = nil
	file_go_chromium_org_luci_cv_internal_run_eventpb_longop_proto_goTypes = nil
	file_go_chromium_org_luci_cv_internal_run_eventpb_longop_proto_depIdxs = nil
}
