// Copyright 2017 The LUCI Authors.
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
// source: go.chromium.org/luci/scheduler/appengine/internal/tq.proto

package internal

import (
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

// ReadProjectConfigTask is used to import jobs of some project.
//
// Queue: "read-project-config".
type ReadProjectConfigTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *ReadProjectConfigTask) Reset() {
	*x = ReadProjectConfigTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadProjectConfigTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadProjectConfigTask) ProtoMessage() {}

func (x *ReadProjectConfigTask) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadProjectConfigTask.ProtoReflect.Descriptor instead.
func (*ReadProjectConfigTask) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_rawDescGZIP(), []int{0}
}

func (x *ReadProjectConfigTask) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

// LaunchInvocationTask is used to start running (or retry a lunch of) a single
// invocation.
//
// It is enqueued non-transactionally, but with the deduplication key.
//
// Queue: "launches".
type LaunchInvocationTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobId string `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	InvId int64  `protobuf:"varint,2,opt,name=inv_id,json=invId,proto3" json:"inv_id,omitempty"`
}

func (x *LaunchInvocationTask) Reset() {
	*x = LaunchInvocationTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LaunchInvocationTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LaunchInvocationTask) ProtoMessage() {}

func (x *LaunchInvocationTask) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LaunchInvocationTask.ProtoReflect.Descriptor instead.
func (*LaunchInvocationTask) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_rawDescGZIP(), []int{1}
}

func (x *LaunchInvocationTask) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *LaunchInvocationTask) GetInvId() int64 {
	if x != nil {
		return x.InvId
	}
	return 0
}

// LaunchInvocationsBatchTask is used to kick off several invocations at once.
//
// It is enqueued transactionally. It fans out into many LaunchInvocationTask.
//
// Queue: "batches".
type LaunchInvocationsBatchTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tasks []*LaunchInvocationTask `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
}

func (x *LaunchInvocationsBatchTask) Reset() {
	*x = LaunchInvocationsBatchTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LaunchInvocationsBatchTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LaunchInvocationsBatchTask) ProtoMessage() {}

func (x *LaunchInvocationsBatchTask) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LaunchInvocationsBatchTask.ProtoReflect.Descriptor instead.
func (*LaunchInvocationsBatchTask) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_rawDescGZIP(), []int{2}
}

func (x *LaunchInvocationsBatchTask) GetTasks() []*LaunchInvocationTask {
	if x != nil {
		return x.Tasks
	}
	return nil
}

// TriageJobStateTask looks at the state of the job and decided what to do next.
//
// Enqueued non-transactionally. It is throttled to run approximately once per
// second. It looks at pending triggers and recently finished invocations and
// launches new invocations (or schedules timers to do it later).
//
// Queue: "triages".
type TriageJobStateTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobId string `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
}

func (x *TriageJobStateTask) Reset() {
	*x = TriageJobStateTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TriageJobStateTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TriageJobStateTask) ProtoMessage() {}

func (x *TriageJobStateTask) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TriageJobStateTask.ProtoReflect.Descriptor instead.
func (*TriageJobStateTask) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_rawDescGZIP(), []int{3}
}

func (x *TriageJobStateTask) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

// KickTriageTask can be used to transactionally initiate a new triage.
//
// We can't transactionally enqueue TriageJobStateTask, since its throttling
// mechanism uses memcache and named tasks, which are not available inside
// transactions. So instead transactions can enqueue KickTriageTask, which in
// turn will enqueue TriageJobStateTask (with throttling).
//
// Queue: "triages".
type KickTriageTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobId string `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
}

func (x *KickTriageTask) Reset() {
	*x = KickTriageTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KickTriageTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KickTriageTask) ProtoMessage() {}

func (x *KickTriageTask) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KickTriageTask.ProtoReflect.Descriptor instead.
func (*KickTriageTask) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_rawDescGZIP(), []int{4}
}

func (x *KickTriageTask) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

// InvocationFinishedTask is emitted by the invocation when it finishes.
//
// It is enqueued transactionally.
//
// Queue: "completions".
type InvocationFinishedTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobId    string              `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	InvId    int64               `protobuf:"varint,2,opt,name=inv_id,json=invId,proto3" json:"inv_id,omitempty"`
	Triggers *FanOutTriggersTask `protobuf:"bytes,3,opt,name=triggers,proto3" json:"triggers,omitempty"`
}

func (x *InvocationFinishedTask) Reset() {
	*x = InvocationFinishedTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvocationFinishedTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvocationFinishedTask) ProtoMessage() {}

func (x *InvocationFinishedTask) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvocationFinishedTask.ProtoReflect.Descriptor instead.
func (*InvocationFinishedTask) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_rawDescGZIP(), []int{5}
}

func (x *InvocationFinishedTask) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *InvocationFinishedTask) GetInvId() int64 {
	if x != nil {
		return x.InvId
	}
	return 0
}

func (x *InvocationFinishedTask) GetTriggers() *FanOutTriggersTask {
	if x != nil {
		return x.Triggers
	}
	return nil
}

// FanOutTriggersTask is a batch task that emits a bunch of triggers.
//
// It is enqueued transactionally. It fans out into many EnqueueTriggersTask,
// one per job ID.
//
// Queue: "triggers".
type FanOutTriggersTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobIds   []string   `protobuf:"bytes,1,rep,name=job_ids,json=jobIds,proto3" json:"job_ids,omitempty"`
	Triggers []*Trigger `protobuf:"bytes,2,rep,name=triggers,proto3" json:"triggers,omitempty"`
}

func (x *FanOutTriggersTask) Reset() {
	*x = FanOutTriggersTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FanOutTriggersTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FanOutTriggersTask) ProtoMessage() {}

func (x *FanOutTriggersTask) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FanOutTriggersTask.ProtoReflect.Descriptor instead.
func (*FanOutTriggersTask) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_rawDescGZIP(), []int{6}
}

func (x *FanOutTriggersTask) GetJobIds() []string {
	if x != nil {
		return x.JobIds
	}
	return nil
}

func (x *FanOutTriggersTask) GetTriggers() []*Trigger {
	if x != nil {
		return x.Triggers
	}
	return nil
}

// EnqueueTriggersTask adds given triggers to a job's pending triggers set.
//
// Enqueued non-transactionally (from FanOutTriggersTask) and transactionally
// (when emitting single trigger from a cron).
//
// Queue: "triggers".
type EnqueueTriggersTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobId    string     `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	Triggers []*Trigger `protobuf:"bytes,2,rep,name=triggers,proto3" json:"triggers,omitempty"`
}

func (x *EnqueueTriggersTask) Reset() {
	*x = EnqueueTriggersTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnqueueTriggersTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnqueueTriggersTask) ProtoMessage() {}

func (x *EnqueueTriggersTask) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnqueueTriggersTask.ProtoReflect.Descriptor instead.
func (*EnqueueTriggersTask) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_rawDescGZIP(), []int{7}
}

func (x *EnqueueTriggersTask) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *EnqueueTriggersTask) GetTriggers() []*Trigger {
	if x != nil {
		return x.Triggers
	}
	return nil
}

// ScheduleTimersTask adds a bunch of delayed invocation calls.
//
// It is enqueued transactionally. Results in a bunch of TimerTask calls.
//
// Queue: "timers".
type ScheduleTimersTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobId  string   `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	InvId  int64    `protobuf:"varint,2,opt,name=inv_id,json=invId,proto3" json:"inv_id,omitempty"`
	Timers []*Timer `protobuf:"bytes,3,rep,name=timers,proto3" json:"timers,omitempty"`
}

func (x *ScheduleTimersTask) Reset() {
	*x = ScheduleTimersTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScheduleTimersTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleTimersTask) ProtoMessage() {}

func (x *ScheduleTimersTask) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleTimersTask.ProtoReflect.Descriptor instead.
func (*ScheduleTimersTask) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_rawDescGZIP(), []int{8}
}

func (x *ScheduleTimersTask) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *ScheduleTimersTask) GetInvId() int64 {
	if x != nil {
		return x.InvId
	}
	return 0
}

func (x *ScheduleTimersTask) GetTimers() []*Timer {
	if x != nil {
		return x.Timers
	}
	return nil
}

// TimerTask corresponds to delayed calls added through AddTimer controller API.
//
// Enqueued either transactionally or not. Deduplicated based on invocation's
// PendingTimers set: any timers not in the set are silently skipped.
//
// Queue: "timers".
type TimerTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobId string `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	InvId int64  `protobuf:"varint,2,opt,name=inv_id,json=invId,proto3" json:"inv_id,omitempty"`
	Timer *Timer `protobuf:"bytes,3,opt,name=timer,proto3" json:"timer,omitempty"`
}

func (x *TimerTask) Reset() {
	*x = TimerTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimerTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimerTask) ProtoMessage() {}

func (x *TimerTask) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimerTask.ProtoReflect.Descriptor instead.
func (*TimerTask) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_rawDescGZIP(), []int{9}
}

func (x *TimerTask) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *TimerTask) GetInvId() int64 {
	if x != nil {
		return x.InvId
	}
	return 0
}

func (x *TimerTask) GetTimer() *Timer {
	if x != nil {
		return x.Timer
	}
	return nil
}

// CronTickTask is scheduled based on the job's cron schedule.
//
// It is enqueued transactionally when the job changes state (e.g. the job
// appears for the first time or its schedule changes) or from previous cron
// ticks.
//
// Queue: "crons".
type CronTickTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobId     string `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	TickNonce int64  `protobuf:"varint,2,opt,name=tick_nonce,json=tickNonce,proto3" json:"tick_nonce,omitempty"` // used to skip no longer interesting ticks
}

func (x *CronTickTask) Reset() {
	*x = CronTickTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CronTickTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CronTickTask) ProtoMessage() {}

func (x *CronTickTask) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CronTickTask.ProtoReflect.Descriptor instead.
func (*CronTickTask) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_rawDescGZIP(), []int{10}
}

func (x *CronTickTask) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *CronTickTask) GetTickNonce() int64 {
	if x != nil {
		return x.TickNonce
	}
	return 0
}

var File_go_chromium_org_luci_scheduler_appengine_internal_tq_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72,
	0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x74, 0x71, 0x1a, 0x3e, 0x67, 0x6f, 0x2e, 0x63, 0x68,
	0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x40, 0x67, 0x6f, 0x2e, 0x63, 0x68,
	0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x72, 0x69,
	0x67, 0x67, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x36, 0x0a, 0x15, 0x52,
	0x65, 0x61, 0x64, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x54, 0x61, 0x73, 0x6b, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x64, 0x22, 0x44, 0x0a, 0x14, 0x4c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x49, 0x6e, 0x76,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x15, 0x0a, 0x06, 0x6a,
	0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62,
	0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x6e, 0x76, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x69, 0x6e, 0x76, 0x49, 0x64, 0x22, 0x55, 0x0a, 0x1a, 0x4c, 0x61, 0x75,
	0x6e, 0x63, 0x68, 0x49, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x37, 0x0a, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2e, 0x74, 0x71, 0x2e, 0x4c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x49, 0x6e, 0x76, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73,
	0x22, 0x2b, 0x0a, 0x12, 0x54, 0x72, 0x69, 0x61, 0x67, 0x65, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x22, 0x27, 0x0a,
	0x0e, 0x4b, 0x69, 0x63, 0x6b, 0x54, 0x72, 0x69, 0x61, 0x67, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12,
	0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x22, 0x83, 0x01, 0x0a, 0x16, 0x49, 0x6e, 0x76, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x54, 0x61, 0x73,
	0x6b, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x6e, 0x76, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x69, 0x6e, 0x76, 0x49, 0x64, 0x12,
	0x3b, 0x0a, 0x08, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x74, 0x71, 0x2e,
	0x46, 0x61, 0x6e, 0x4f, 0x75, 0x74, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x73, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x08, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x73, 0x22, 0x65, 0x0a, 0x12,
	0x46, 0x61, 0x6e, 0x4f, 0x75, 0x74, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x73, 0x54, 0x61,
	0x73, 0x6b, 0x12, 0x17, 0x0a, 0x07, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x06, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x73, 0x12, 0x36, 0x0a, 0x08, 0x74,
	0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72,
	0x73, 0x2e, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x08, 0x74, 0x72, 0x69, 0x67, 0x67,
	0x65, 0x72, 0x73, 0x22, 0x64, 0x0a, 0x13, 0x45, 0x6e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x54, 0x72,
	0x69, 0x67, 0x67, 0x65, 0x72, 0x73, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f,
	0x62, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49,
	0x64, 0x12, 0x36, 0x0a, 0x08, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x74,
	0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x73, 0x2e, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52,
	0x08, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x73, 0x22, 0x72, 0x0a, 0x12, 0x53, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x73, 0x54, 0x61, 0x73, 0x6b, 0x12,
	0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x6e, 0x76, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x69, 0x6e, 0x76, 0x49, 0x64, 0x12, 0x2e, 0x0a,
	0x06, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x73, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x72, 0x52, 0x06, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x73, 0x22, 0x67, 0x0a,
	0x09, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f,
	0x62, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49,
	0x64, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x6e, 0x76, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x69, 0x6e, 0x76, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x05, 0x74, 0x69, 0x6d, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x73, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x52,
	0x05, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x22, 0x44, 0x0a, 0x0c, 0x43, 0x72, 0x6f, 0x6e, 0x54, 0x69,
	0x63, 0x6b, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x74, 0x69, 0x63, 0x6b, 0x5f, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x74, 0x69, 0x63, 0x6b, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x42, 0x33, 0x5a, 0x31,
	0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x6c, 0x75, 0x63, 0x69, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2f, 0x61,
	0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_rawDescData = file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_rawDesc
)

func file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_rawDescData)
	})
	return file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_rawDescData
}

var file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_goTypes = []interface{}{
	(*ReadProjectConfigTask)(nil),      // 0: internal.tq.ReadProjectConfigTask
	(*LaunchInvocationTask)(nil),       // 1: internal.tq.LaunchInvocationTask
	(*LaunchInvocationsBatchTask)(nil), // 2: internal.tq.LaunchInvocationsBatchTask
	(*TriageJobStateTask)(nil),         // 3: internal.tq.TriageJobStateTask
	(*KickTriageTask)(nil),             // 4: internal.tq.KickTriageTask
	(*InvocationFinishedTask)(nil),     // 5: internal.tq.InvocationFinishedTask
	(*FanOutTriggersTask)(nil),         // 6: internal.tq.FanOutTriggersTask
	(*EnqueueTriggersTask)(nil),        // 7: internal.tq.EnqueueTriggersTask
	(*ScheduleTimersTask)(nil),         // 8: internal.tq.ScheduleTimersTask
	(*TimerTask)(nil),                  // 9: internal.tq.TimerTask
	(*CronTickTask)(nil),               // 10: internal.tq.CronTickTask
	(*Trigger)(nil),                    // 11: internal.triggers.Trigger
	(*Timer)(nil),                      // 12: internal.timers.Timer
}
var file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_depIdxs = []int32{
	1,  // 0: internal.tq.LaunchInvocationsBatchTask.tasks:type_name -> internal.tq.LaunchInvocationTask
	6,  // 1: internal.tq.InvocationFinishedTask.triggers:type_name -> internal.tq.FanOutTriggersTask
	11, // 2: internal.tq.FanOutTriggersTask.triggers:type_name -> internal.triggers.Trigger
	11, // 3: internal.tq.EnqueueTriggersTask.triggers:type_name -> internal.triggers.Trigger
	12, // 4: internal.tq.ScheduleTimersTask.timers:type_name -> internal.timers.Timer
	12, // 5: internal.tq.TimerTask.timer:type_name -> internal.timers.Timer
	6,  // [6:6] is the sub-list for method output_type
	6,  // [6:6] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_init() }
func file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_init() {
	if File_go_chromium_org_luci_scheduler_appengine_internal_tq_proto != nil {
		return
	}
	file_go_chromium_org_luci_scheduler_appengine_internal_timers_proto_init()
	file_go_chromium_org_luci_scheduler_appengine_internal_triggers_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadProjectConfigTask); i {
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
		file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LaunchInvocationTask); i {
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
		file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LaunchInvocationsBatchTask); i {
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
		file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TriageJobStateTask); i {
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
		file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KickTriageTask); i {
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
		file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvocationFinishedTask); i {
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
		file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FanOutTriggersTask); i {
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
		file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnqueueTriggersTask); i {
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
		file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScheduleTimersTask); i {
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
		file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimerTask); i {
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
		file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CronTickTask); i {
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
			RawDescriptor: file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_scheduler_appengine_internal_tq_proto = out.File
	file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_rawDesc = nil
	file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_goTypes = nil
	file_go_chromium_org_luci_scheduler_appengine_internal_tq_proto_depIdxs = nil
}
