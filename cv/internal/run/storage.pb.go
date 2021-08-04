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
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.0
// source: go.chromium.org/luci/cv/internal/run/storage.proto

package run

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

// Trigger describes who/how CV was triggered on a specific CL.
type Trigger struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	// Mode is string value of run.Mode.
	Mode string `protobuf:"bytes,2,opt,name=mode,proto3" json:"mode,omitempty"`
	// Additional label is recorded in case applicable ConfigGroup had additional
	// modes, e.g. for QUICK_DRY_RUN mode.
	AdditionalLabel string `protobuf:"bytes,5,opt,name=additional_label,json=additionalLabel,proto3" json:"additional_label,omitempty"`
	// Triggering user email if known.
	//
	// Gerrit doesn't guarantee that every user has set their preferred email,
	// but LUCI ACLs are based entirely on user emails. Thus, Runs with the email
	// unset are quickly aborted by CQDaemon.
	//
	// TODO(tandrii): once CQDaemon is deleted, require email to start a Run,
	// and remove Gerrit-specific gerrit_account_id.
	Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	// Gerrit account ID. Always known.
	GerritAccountId int64 `protobuf:"varint,4,opt,name=gerrit_account_id,json=gerritAccountId,proto3" json:"gerrit_account_id,omitempty"`
}

func (x *Trigger) Reset() {
	*x = Trigger{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_internal_run_storage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Trigger) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trigger) ProtoMessage() {}

func (x *Trigger) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_internal_run_storage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trigger.ProtoReflect.Descriptor instead.
func (*Trigger) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_internal_run_storage_proto_rawDescGZIP(), []int{0}
}

func (x *Trigger) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *Trigger) GetMode() string {
	if x != nil {
		return x.Mode
	}
	return ""
}

func (x *Trigger) GetAdditionalLabel() string {
	if x != nil {
		return x.AdditionalLabel
	}
	return ""
}

func (x *Trigger) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Trigger) GetGerritAccountId() int64 {
	if x != nil {
		return x.GerritAccountId
	}
	return 0
}

// Submission describes the current state of Run submission.
type Submission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The deadline of this submission.
	//
	// If the deadline is not set or has already expired, a RunManager task
	// can claim the exclusive privilege by setting the deadline to a future
	// timestamp (generally, end of task deadline).
	Deadline *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=deadline,proto3" json:"deadline,omitempty"`
	// ID of the task that executes this submission.
	TaskId string `protobuf:"bytes,2,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	// IDs of all CLs that should be submitted in this submission.
	//
	// Must be ordered in submission order.
	Cls []int64 `protobuf:"varint,3,rep,packed,name=cls,proto3" json:"cls,omitempty"`
	// IDs of all CLs that have been submitted successfully already.
	SubmittedCls []int64 `protobuf:"varint,4,rep,packed,name=submitted_cls,json=submittedCls,proto3" json:"submitted_cls,omitempty"`
	// IDs of all CLs that fails to submit if any.
	//
	// CLs that are neither in this list nor in the `submitted_cls` should be
	// treated as if CV has never attempted to submit them.
	//
	// This could be empty even when the entire submission fails, which would be
	// typically caused by faulty infrastructure (e.g. Task Queue not executing
	// a Run Manager task before the whole submission timeout is reached).
	FailedCls []int64 `protobuf:"varint,5,rep,packed,name=failed_cls,json=failedCls,proto3" json:"failed_cls,omitempty"`
	// If True, Tree is currently in open state.
	TreeOpen bool `protobuf:"varint,10,opt,name=tree_open,json=treeOpen,proto3" json:"tree_open,omitempty"`
	// The timestamp when the Tree status was last fetched.
	LastTreeCheckTime *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=last_tree_check_time,json=lastTreeCheckTime,proto3" json:"last_tree_check_time,omitempty"`
}

func (x *Submission) Reset() {
	*x = Submission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_internal_run_storage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Submission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Submission) ProtoMessage() {}

func (x *Submission) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_internal_run_storage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Submission.ProtoReflect.Descriptor instead.
func (*Submission) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_internal_run_storage_proto_rawDescGZIP(), []int{1}
}

func (x *Submission) GetDeadline() *timestamppb.Timestamp {
	if x != nil {
		return x.Deadline
	}
	return nil
}

func (x *Submission) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *Submission) GetCls() []int64 {
	if x != nil {
		return x.Cls
	}
	return nil
}

func (x *Submission) GetSubmittedCls() []int64 {
	if x != nil {
		return x.SubmittedCls
	}
	return nil
}

func (x *Submission) GetFailedCls() []int64 {
	if x != nil {
		return x.FailedCls
	}
	return nil
}

func (x *Submission) GetTreeOpen() bool {
	if x != nil {
		return x.TreeOpen
	}
	return false
}

func (x *Submission) GetLastTreeCheckTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastTreeCheckTime
	}
	return nil
}

// Options are Run-specific additions on top of LUCI project config.
type Options struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If true, submitting the Run isn't blocked on open tree.
	//
	// If false (default), respects project configuration.
	SkipTreeChecks bool `protobuf:"varint,1,opt,name=skip_tree_checks,json=skipTreeChecks,proto3" json:"skip_tree_checks,omitempty"`
	// If true, `builders.equivalent_to{...}` sections are ignored when triggering
	// tryjobs.
	//
	// If false (default), respects project configuration.
	SkipEquivalentBuilders bool `protobuf:"varint,2,opt,name=skip_equivalent_builders,json=skipEquivalentBuilders,proto3" json:"skip_equivalent_builders,omitempty"`
	// If true, no longer useful tryjobs won't be cancelled.
	//
	// If false (default), respects project configuration.
	AvoidCancellingTryjobs bool `protobuf:"varint,3,opt,name=avoid_cancelling_tryjobs,json=avoidCancellingTryjobs,proto3" json:"avoid_cancelling_tryjobs,omitempty"`
	// If true, no tryjobs will be triggered except "presubmit" regardless of
	// project configuration.
	//
	// "presubmit" builders are legacy which are currently configured with
	// "disable_reuse: true" in project config. To skip triggering them,
	// skip_presubmit must be set to true.
	// TODO(https://crbug.com/950074): ignore.
	//
	// If false (default), respects project configuration.
	SkipTryjobs bool `protobuf:"varint,4,opt,name=skip_tryjobs,json=skipTryjobs,proto3" json:"skip_tryjobs,omitempty"`
	// Deprecated per https://crbug.com/950074.
	// See skip_tryjobs doc.
	SkipPresubmit bool `protobuf:"varint,5,opt,name=skip_presubmit,json=skipPresubmit,proto3" json:"skip_presubmit,omitempty"`
}

func (x *Options) Reset() {
	*x = Options{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_internal_run_storage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Options) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Options) ProtoMessage() {}

func (x *Options) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_internal_run_storage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Options.ProtoReflect.Descriptor instead.
func (*Options) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_internal_run_storage_proto_rawDescGZIP(), []int{2}
}

func (x *Options) GetSkipTreeChecks() bool {
	if x != nil {
		return x.SkipTreeChecks
	}
	return false
}

func (x *Options) GetSkipEquivalentBuilders() bool {
	if x != nil {
		return x.SkipEquivalentBuilders
	}
	return false
}

func (x *Options) GetAvoidCancellingTryjobs() bool {
	if x != nil {
		return x.AvoidCancellingTryjobs
	}
	return false
}

func (x *Options) GetSkipTryjobs() bool {
	if x != nil {
		return x.SkipTryjobs
	}
	return false
}

func (x *Options) GetSkipPresubmit() bool {
	if x != nil {
		return x.SkipPresubmit
	}
	return false
}

var File_go_chromium_org_luci_cv_internal_run_storage_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_cv_internal_run_storage_proto_rawDesc = []byte{
	0x0a, 0x32, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x76, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x72, 0x75, 0x6e, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x76, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2e, 0x72, 0x75, 0x6e, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xba, 0x01, 0x0a, 0x07, 0x54, 0x72, 0x69, 0x67, 0x67,
	0x65, 0x72, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x2a, 0x0a, 0x11, 0x67, 0x65, 0x72, 0x72, 0x69,
	0x74, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0f, 0x67, 0x65, 0x72, 0x72, 0x69, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x64, 0x22, 0x9d, 0x02, 0x0a, 0x0a, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61,
	0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73,
	0x6b, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x03,
	0x52, 0x03, 0x63, 0x6c, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x74,
	0x65, 0x64, 0x5f, 0x63, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0c, 0x73, 0x75,
	0x62, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x43, 0x6c, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x61,
	0x69, 0x6c, 0x65, 0x64, 0x5f, 0x63, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x03, 0x52, 0x09,
	0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x43, 0x6c, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x72, 0x65,
	0x65, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x74, 0x72,
	0x65, 0x65, 0x4f, 0x70, 0x65, 0x6e, 0x12, 0x4b, 0x0a, 0x14, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x74,
	0x72, 0x65, 0x65, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x11, 0x6c, 0x61, 0x73, 0x74, 0x54, 0x72, 0x65, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54,
	0x69, 0x6d, 0x65, 0x22, 0xf1, 0x01, 0x0a, 0x07, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x28, 0x0a, 0x10, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x74, 0x72, 0x65, 0x65, 0x5f, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x73, 0x6b, 0x69, 0x70, 0x54,
	0x72, 0x65, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x12, 0x38, 0x0a, 0x18, 0x73, 0x6b, 0x69,
	0x70, 0x5f, 0x65, 0x71, 0x75, 0x69, 0x76, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x5f, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x73, 0x6b, 0x69,
	0x70, 0x45, 0x71, 0x75, 0x69, 0x76, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x42, 0x75, 0x69, 0x6c, 0x64,
	0x65, 0x72, 0x73, 0x12, 0x38, 0x0a, 0x18, 0x61, 0x76, 0x6f, 0x69, 0x64, 0x5f, 0x63, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x72, 0x79, 0x6a, 0x6f, 0x62, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x61, 0x76, 0x6f, 0x69, 0x64, 0x43, 0x61, 0x6e, 0x63,
	0x65, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x54, 0x72, 0x79, 0x6a, 0x6f, 0x62, 0x73, 0x12, 0x21, 0x0a,
	0x0c, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x74, 0x72, 0x79, 0x6a, 0x6f, 0x62, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0b, 0x73, 0x6b, 0x69, 0x70, 0x54, 0x72, 0x79, 0x6a, 0x6f, 0x62, 0x73,
	0x12, 0x25, 0x0a, 0x0e, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x70, 0x72, 0x65, 0x73, 0x75, 0x62, 0x6d,
	0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x73, 0x6b, 0x69, 0x70, 0x50, 0x72,
	0x65, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x6f, 0x2e, 0x63, 0x68,
	0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f,
	0x63, 0x76, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x72, 0x75, 0x6e, 0x3b,
	0x72, 0x75, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_cv_internal_run_storage_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_cv_internal_run_storage_proto_rawDescData = file_go_chromium_org_luci_cv_internal_run_storage_proto_rawDesc
)

func file_go_chromium_org_luci_cv_internal_run_storage_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_cv_internal_run_storage_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_cv_internal_run_storage_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_cv_internal_run_storage_proto_rawDescData)
	})
	return file_go_chromium_org_luci_cv_internal_run_storage_proto_rawDescData
}

var file_go_chromium_org_luci_cv_internal_run_storage_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_go_chromium_org_luci_cv_internal_run_storage_proto_goTypes = []interface{}{
	(*Trigger)(nil),               // 0: cv.internal.run.Trigger
	(*Submission)(nil),            // 1: cv.internal.run.Submission
	(*Options)(nil),               // 2: cv.internal.run.Options
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_go_chromium_org_luci_cv_internal_run_storage_proto_depIdxs = []int32{
	3, // 0: cv.internal.run.Trigger.time:type_name -> google.protobuf.Timestamp
	3, // 1: cv.internal.run.Submission.deadline:type_name -> google.protobuf.Timestamp
	3, // 2: cv.internal.run.Submission.last_tree_check_time:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_cv_internal_run_storage_proto_init() }
func file_go_chromium_org_luci_cv_internal_run_storage_proto_init() {
	if File_go_chromium_org_luci_cv_internal_run_storage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_cv_internal_run_storage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Trigger); i {
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
		file_go_chromium_org_luci_cv_internal_run_storage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Submission); i {
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
		file_go_chromium_org_luci_cv_internal_run_storage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Options); i {
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
			RawDescriptor: file_go_chromium_org_luci_cv_internal_run_storage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_cv_internal_run_storage_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_cv_internal_run_storage_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_cv_internal_run_storage_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_cv_internal_run_storage_proto = out.File
	file_go_chromium_org_luci_cv_internal_run_storage_proto_rawDesc = nil
	file_go_chromium_org_luci_cv_internal_run_storage_proto_goTypes = nil
	file_go_chromium_org_luci_cv_internal_run_storage_proto_depIdxs = nil
}
