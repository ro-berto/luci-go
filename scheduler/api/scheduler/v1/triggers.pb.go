// Copyright 2016 The LUCI Authors.
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
// source: go.chromium.org/luci/scheduler/api/scheduler/v1/triggers.proto

package scheduler

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Trigger can be emitted by triggering tasks (such as Gitiles tasks) or through
// API and consumed by triggered tasks (such as Buildbucket tasks).
type Trigger struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique identifier of the trigger.
	//
	// It is used to deduplicate and hence provide idempotency for adding
	// a trigger. Each job has an internal buffer with IDs of recent triggers it
	// received. Triggers that have already been seen are silently skipped. The
	// buffer is periodically cleaned, so old IDs can be potentially reused,
	// though you should not rely on that.
	//
	// Must be provided by whoever emits the trigger. Can be anything at all, as
	// long as it is unique.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Optional user friendly name for this trigger that shows up in Scheduler UI.
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	// Optional HTTP link to display in Scheduler UI.
	Url string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	// Actual trigger data. Its type defines how the trigger will be processed
	// by the Scheduler, see corresponding protos.
	//
	// Types that are assignable to Payload:
	//	*Trigger_Cron
	//	*Trigger_Webui
	//	*Trigger_Noop
	//	*Trigger_Gitiles
	//	*Trigger_Buildbucket
	Payload isTrigger_Payload `protobuf_oneof:"payload"`
}

func (x *Trigger) Reset() {
	*x = Trigger{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_scheduler_api_scheduler_v1_triggers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Trigger) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trigger) ProtoMessage() {}

func (x *Trigger) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_scheduler_api_scheduler_v1_triggers_proto_msgTypes[0]
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
	return file_go_chromium_org_luci_scheduler_api_scheduler_v1_triggers_proto_rawDescGZIP(), []int{0}
}

func (x *Trigger) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Trigger) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Trigger) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (m *Trigger) GetPayload() isTrigger_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *Trigger) GetCron() *CronTrigger {
	if x, ok := x.GetPayload().(*Trigger_Cron); ok {
		return x.Cron
	}
	return nil
}

func (x *Trigger) GetWebui() *WebUITrigger {
	if x, ok := x.GetPayload().(*Trigger_Webui); ok {
		return x.Webui
	}
	return nil
}

func (x *Trigger) GetNoop() *NoopTrigger {
	if x, ok := x.GetPayload().(*Trigger_Noop); ok {
		return x.Noop
	}
	return nil
}

func (x *Trigger) GetGitiles() *GitilesTrigger {
	if x, ok := x.GetPayload().(*Trigger_Gitiles); ok {
		return x.Gitiles
	}
	return nil
}

func (x *Trigger) GetBuildbucket() *BuildbucketTrigger {
	if x, ok := x.GetPayload().(*Trigger_Buildbucket); ok {
		return x.Buildbucket
	}
	return nil
}

type isTrigger_Payload interface {
	isTrigger_Payload()
}

type Trigger_Cron struct {
	Cron *CronTrigger `protobuf:"bytes,40,opt,name=cron,proto3,oneof"`
}

type Trigger_Webui struct {
	Webui *WebUITrigger `protobuf:"bytes,41,opt,name=webui,proto3,oneof"`
}

type Trigger_Noop struct {
	Noop *NoopTrigger `protobuf:"bytes,50,opt,name=noop,proto3,oneof"`
}

type Trigger_Gitiles struct {
	Gitiles *GitilesTrigger `protobuf:"bytes,51,opt,name=gitiles,proto3,oneof"`
}

type Trigger_Buildbucket struct {
	Buildbucket *BuildbucketTrigger `protobuf:"bytes,52,opt,name=buildbucket,proto3,oneof"`
}

func (*Trigger_Cron) isTrigger_Payload() {}

func (*Trigger_Webui) isTrigger_Payload() {}

func (*Trigger_Noop) isTrigger_Payload() {}

func (*Trigger_Gitiles) isTrigger_Payload() {}

func (*Trigger_Buildbucket) isTrigger_Payload() {}

// CronTrigger is used internally by Scheduler to implement periodic jobs.
//
// It is emitted by the cron state machines whenever it decides the scheduler
// should launch the invocation.
//
// Note: such triggers can't be scheduled through external Scheduler API. They
// may appear in the API responses though.
type CronTrigger struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Generation int64 `protobuf:"varint,1,opt,name=generation,proto3" json:"generation,omitempty"` // monotonically increasing number
}

func (x *CronTrigger) Reset() {
	*x = CronTrigger{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_scheduler_api_scheduler_v1_triggers_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CronTrigger) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CronTrigger) ProtoMessage() {}

func (x *CronTrigger) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_scheduler_api_scheduler_v1_triggers_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CronTrigger.ProtoReflect.Descriptor instead.
func (*CronTrigger) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_scheduler_api_scheduler_v1_triggers_proto_rawDescGZIP(), []int{1}
}

func (x *CronTrigger) GetGeneration() int64 {
	if x != nil {
		return x.Generation
	}
	return 0
}

// WebUITrigger is emitted whenever users click "Trigger" button in UI.
//
// Note: such triggers can't be scheduled through external Scheduler API (to
// avoid confusion). They may appear in the API responses though.
type WebUITrigger struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WebUITrigger) Reset() {
	*x = WebUITrigger{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_scheduler_api_scheduler_v1_triggers_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebUITrigger) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebUITrigger) ProtoMessage() {}

func (x *WebUITrigger) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_scheduler_api_scheduler_v1_triggers_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebUITrigger.ProtoReflect.Descriptor instead.
func (*WebUITrigger) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_scheduler_api_scheduler_v1_triggers_proto_rawDescGZIP(), []int{2}
}

// NoopTrigger is used by Scheduler integration tests to represent test
// triggers.
type NoopTrigger struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *NoopTrigger) Reset() {
	*x = NoopTrigger{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_scheduler_api_scheduler_v1_triggers_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoopTrigger) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoopTrigger) ProtoMessage() {}

func (x *NoopTrigger) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_scheduler_api_scheduler_v1_triggers_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoopTrigger.ProtoReflect.Descriptor instead.
func (*NoopTrigger) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_scheduler_api_scheduler_v1_triggers_proto_rawDescGZIP(), []int{3}
}

func (x *NoopTrigger) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

// GitilesTrigger is emitted by sources that watch Gitiles and consumed by
// Buildbucket tasks.
//
// Such triggers are emitted whenever the repository state changes or via
// EmitTriggers API.
type GitilesTrigger struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Repo     string `protobuf:"bytes,1,opt,name=repo,proto3" json:"repo,omitempty"`         // an URL of a repo that changed
	Ref      string `protobuf:"bytes,2,opt,name=ref,proto3" json:"ref,omitempty"`           // a ref that changed, in full, e.g "refs/heads/master"
	Revision string `protobuf:"bytes,3,opt,name=revision,proto3" json:"revision,omitempty"` // a revision (SHA1 in hex) pointed to by the ref
	// Properties and tags to add in addition to ones generated from the commit.
	Properties *structpb.Struct `protobuf:"bytes,4,opt,name=properties,proto3" json:"properties,omitempty"`
	Tags       []string         `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *GitilesTrigger) Reset() {
	*x = GitilesTrigger{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_scheduler_api_scheduler_v1_triggers_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GitilesTrigger) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GitilesTrigger) ProtoMessage() {}

func (x *GitilesTrigger) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_scheduler_api_scheduler_v1_triggers_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GitilesTrigger.ProtoReflect.Descriptor instead.
func (*GitilesTrigger) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_scheduler_api_scheduler_v1_triggers_proto_rawDescGZIP(), []int{4}
}

func (x *GitilesTrigger) GetRepo() string {
	if x != nil {
		return x.Repo
	}
	return ""
}

func (x *GitilesTrigger) GetRef() string {
	if x != nil {
		return x.Ref
	}
	return ""
}

func (x *GitilesTrigger) GetRevision() string {
	if x != nil {
		return x.Revision
	}
	return ""
}

func (x *GitilesTrigger) GetProperties() *structpb.Struct {
	if x != nil {
		return x.Properties
	}
	return nil
}

func (x *GitilesTrigger) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

// BuildbucketTrigger is emitted by sources that request a general build and
// consumed by Buildbucket tasks.
//
// The trigger contains information to pass to the new Buildbucket build.
//
// Note: what builds to trigger is specified separately, either in the job
// configuration (when one job triggers another) or via the API request
// parameters (when triggering through public API).
type BuildbucketTrigger struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Properties *structpb.Struct `protobuf:"bytes,1,opt,name=properties,proto3" json:"properties,omitempty"`
	Tags       []string         `protobuf:"bytes,2,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *BuildbucketTrigger) Reset() {
	*x = BuildbucketTrigger{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_scheduler_api_scheduler_v1_triggers_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildbucketTrigger) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildbucketTrigger) ProtoMessage() {}

func (x *BuildbucketTrigger) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_scheduler_api_scheduler_v1_triggers_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildbucketTrigger.ProtoReflect.Descriptor instead.
func (*BuildbucketTrigger) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_scheduler_api_scheduler_v1_triggers_proto_rawDescGZIP(), []int{5}
}

func (x *BuildbucketTrigger) GetProperties() *structpb.Struct {
	if x != nil {
		return x.Properties
	}
	return nil
}

func (x *BuildbucketTrigger) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

var File_go_chromium_org_luci_scheduler_api_scheduler_v1_triggers_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_scheduler_api_scheduler_v1_triggers_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x09, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd3, 0x02, 0x0a, 0x07, 0x54, 0x72,
	0x69, 0x67, 0x67, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x2c, 0x0a,
	0x04, 0x63, 0x72, 0x6f, 0x6e, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x6f, 0x6e, 0x54, 0x72, 0x69, 0x67,
	0x67, 0x65, 0x72, 0x48, 0x00, 0x52, 0x04, 0x63, 0x72, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x05, 0x77,
	0x65, 0x62, 0x75, 0x69, 0x18, 0x29, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x57, 0x65, 0x62, 0x55, 0x49, 0x54, 0x72, 0x69, 0x67,
	0x67, 0x65, 0x72, 0x48, 0x00, 0x52, 0x05, 0x77, 0x65, 0x62, 0x75, 0x69, 0x12, 0x2c, 0x0a, 0x04,
	0x6e, 0x6f, 0x6f, 0x70, 0x18, 0x32, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x6f, 0x70, 0x54, 0x72, 0x69, 0x67, 0x67,
	0x65, 0x72, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x6f, 0x6f, 0x70, 0x12, 0x35, 0x0a, 0x07, 0x67, 0x69,
	0x74, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x33, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x47, 0x69, 0x74, 0x69, 0x6c, 0x65, 0x73, 0x54,
	0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x48, 0x00, 0x52, 0x07, 0x67, 0x69, 0x74, 0x69, 0x6c, 0x65,
	0x73, 0x12, 0x41, 0x0a, 0x0b, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x18, 0x34, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x72, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x54, 0x72,
	0x69, 0x67, 0x67, 0x65, 0x72, 0x48, 0x00, 0x52, 0x0b, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22,
	0x2d, 0x0a, 0x0b, 0x43, 0x72, 0x6f, 0x6e, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x12, 0x1e,
	0x0a, 0x0a, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x0e,
	0x0a, 0x0c, 0x57, 0x65, 0x62, 0x55, 0x49, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x22, 0x21,
	0x0a, 0x0b, 0x4e, 0x6f, 0x6f, 0x70, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x9f, 0x01, 0x0a, 0x0e, 0x47, 0x69, 0x74, 0x69, 0x6c, 0x65, 0x73, 0x54, 0x72, 0x69,
	0x67, 0x67, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x66, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72, 0x65, 0x66, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x37, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x69, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x22, 0x61, 0x0a, 0x12, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x12, 0x37, 0x0a, 0x0a, 0x70, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69,
	0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72,
	0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_scheduler_api_scheduler_v1_triggers_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_scheduler_api_scheduler_v1_triggers_proto_rawDescData = file_go_chromium_org_luci_scheduler_api_scheduler_v1_triggers_proto_rawDesc
)

func file_go_chromium_org_luci_scheduler_api_scheduler_v1_triggers_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_scheduler_api_scheduler_v1_triggers_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_scheduler_api_scheduler_v1_triggers_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_scheduler_api_scheduler_v1_triggers_proto_rawDescData)
	})
	return file_go_chromium_org_luci_scheduler_api_scheduler_v1_triggers_proto_rawDescData
}

var file_go_chromium_org_luci_scheduler_api_scheduler_v1_triggers_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_go_chromium_org_luci_scheduler_api_scheduler_v1_triggers_proto_goTypes = []interface{}{
	(*Trigger)(nil),            // 0: scheduler.Trigger
	(*CronTrigger)(nil),        // 1: scheduler.CronTrigger
	(*WebUITrigger)(nil),       // 2: scheduler.WebUITrigger
	(*NoopTrigger)(nil),        // 3: scheduler.NoopTrigger
	(*GitilesTrigger)(nil),     // 4: scheduler.GitilesTrigger
	(*BuildbucketTrigger)(nil), // 5: scheduler.BuildbucketTrigger
	(*structpb.Struct)(nil),    // 6: google.protobuf.Struct
}
var file_go_chromium_org_luci_scheduler_api_scheduler_v1_triggers_proto_depIdxs = []int32{
	1, // 0: scheduler.Trigger.cron:type_name -> scheduler.CronTrigger
	2, // 1: scheduler.Trigger.webui:type_name -> scheduler.WebUITrigger
	3, // 2: scheduler.Trigger.noop:type_name -> scheduler.NoopTrigger
	4, // 3: scheduler.Trigger.gitiles:type_name -> scheduler.GitilesTrigger
	5, // 4: scheduler.Trigger.buildbucket:type_name -> scheduler.BuildbucketTrigger
	6, // 5: scheduler.GitilesTrigger.properties:type_name -> google.protobuf.Struct
	6, // 6: scheduler.BuildbucketTrigger.properties:type_name -> google.protobuf.Struct
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_scheduler_api_scheduler_v1_triggers_proto_init() }
func file_go_chromium_org_luci_scheduler_api_scheduler_v1_triggers_proto_init() {
	if File_go_chromium_org_luci_scheduler_api_scheduler_v1_triggers_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_scheduler_api_scheduler_v1_triggers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_go_chromium_org_luci_scheduler_api_scheduler_v1_triggers_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CronTrigger); i {
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
		file_go_chromium_org_luci_scheduler_api_scheduler_v1_triggers_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebUITrigger); i {
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
		file_go_chromium_org_luci_scheduler_api_scheduler_v1_triggers_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoopTrigger); i {
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
		file_go_chromium_org_luci_scheduler_api_scheduler_v1_triggers_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GitilesTrigger); i {
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
		file_go_chromium_org_luci_scheduler_api_scheduler_v1_triggers_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildbucketTrigger); i {
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
	file_go_chromium_org_luci_scheduler_api_scheduler_v1_triggers_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Trigger_Cron)(nil),
		(*Trigger_Webui)(nil),
		(*Trigger_Noop)(nil),
		(*Trigger_Gitiles)(nil),
		(*Trigger_Buildbucket)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_go_chromium_org_luci_scheduler_api_scheduler_v1_triggers_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_scheduler_api_scheduler_v1_triggers_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_scheduler_api_scheduler_v1_triggers_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_scheduler_api_scheduler_v1_triggers_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_scheduler_api_scheduler_v1_triggers_proto = out.File
	file_go_chromium_org_luci_scheduler_api_scheduler_v1_triggers_proto_rawDesc = nil
	file_go_chromium_org_luci_scheduler_api_scheduler_v1_triggers_proto_goTypes = nil
	file_go_chromium_org_luci_scheduler_api_scheduler_v1_triggers_proto_depIdxs = nil
}
