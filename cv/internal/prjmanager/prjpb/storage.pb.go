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
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.0
// source: go.chromium.org/luci/cv/internal/prjmanager/prjpb/storage.proto

package prjpb

import (
	changelist "go.chromium.org/luci/cv/internal/changelist"
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

type Status int32

const (
	Status_STATUS_UNSPECIFIED Status = 0
	Status_STARTED            Status = 1
	Status_STOPPING           Status = 2
	Status_STOPPED            Status = 3
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "STARTED",
		2: "STOPPING",
		3: "STOPPED",
	}
	Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"STARTED":            1,
		"STOPPING":           2,
		"STOPPED":            3,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_storage_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_storage_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_storage_proto_rawDescGZIP(), []int{0}
}

type PCL_Status int32

const (
	PCL_PCL_STATUS_UNSPECIFIED PCL_Status = 0
	// OK means CL metadata below is correct and CL is watched by this project.
	//
	// Value 0 is chosen such that it's not serialized, since this is the most
	// common state.
	PCL_OK PCL_Status = 0
	// UNKNOWN means Datastore CL entity doesn't have the info yet.
	PCL_UNKNOWN PCL_Status = 1
	// UNWATCHED means CL isn't watched by this LUCI project.
	PCL_UNWATCHED PCL_Status = 2
	// DELETED means CL's Datastore entity got deleted.
	//
	// This is used to temporary mark a PCL before deleting it entirely from
	// PState to avoid dangling references from components.
	PCL_DELETED PCL_Status = 3
)

// Enum value maps for PCL_Status.
var (
	PCL_Status_name = map[int32]string{
		0: "PCL_STATUS_UNSPECIFIED",
		// Duplicate value: 0: "OK",
		1: "UNKNOWN",
		2: "UNWATCHED",
		3: "DELETED",
	}
	PCL_Status_value = map[string]int32{
		"PCL_STATUS_UNSPECIFIED": 0,
		"OK":                     0,
		"UNKNOWN":                1,
		"UNWATCHED":              2,
		"DELETED":                3,
	}
)

func (x PCL_Status) Enum() *PCL_Status {
	p := new(PCL_Status)
	*p = x
	return p
}

func (x PCL_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PCL_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_storage_proto_enumTypes[1].Descriptor()
}

func (PCL_Status) Type() protoreflect.EnumType {
	return &file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_storage_proto_enumTypes[1]
}

func (x PCL_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PCL_Status.Descriptor instead.
func (PCL_Status) EnumDescriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_storage_proto_rawDescGZIP(), []int{1, 0}
}

// PState is the PM state of a specific LUCI project.
//
// Semantically, it's a collection of CLs somehow grouped into components (see
// Component message below), each of which may have several active (a.k.a.
// Incomplete) Runs valid at a specific project's config version.
//
// Most CLs are watched by the LUCI project, but to assist with error reporting,
// it also tracks unwatched CLs if they are dependencies of some actually
// watched CLs.
type PState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of LUCI project.
	LuciProject string `protobuf:"bytes,1,opt,name=luci_project,json=luciProject,proto3" json:"luci_project,omitempty"`
	// Status of the Project.
	Status Status `protobuf:"varint,2,opt,name=status,proto3,enum=cv.prjmanager.prjpb.Status" json:"status,omitempty"`
	// Config hash pins specific project config version.
	ConfigHash string `protobuf:"bytes,3,opt,name=config_hash,json=configHash,proto3" json:"config_hash,omitempty"`
	// Config group names intern the names referenced in PCL entities to reduce
	// memory and at-rest footprint.
	//
	// See also https://en.wikipedia.org/wiki/String_interning.
	ConfigGroupNames []string `protobuf:"bytes,4,rep,name=config_group_names,json=configGroupNames,proto3" json:"config_group_names,omitempty"`
	// PCLs are currently tracked CLs.
	//
	// Includes deps which are of not yet known kind (because CL doesn't yet have
	// a snapshot) or unwatched.
	//
	// Sorted by CL ID.
	Pcls []*PCL `protobuf:"bytes,11,rep,name=pcls,proto3" json:"pcls,omitempty"`
	// Components are a partition of CLs in the list above.
	//
	// An active CL (watched or used to be watched and still member of a Run) may
	// belong to at most 1 component, while unwatched dep may be referenced by
	// several.
	Components []*Component `protobuf:"bytes,12,rep,name=components,proto3" json:"components,omitempty"`
	// PurgingCLs are CLs currently being purged.
	//
	// They are tracked in PState to avoid creating Runs with such CLs.
	//
	// A CL being purged does not necessarily have a corresponding PCL.
	// A PurgingCL is kept in PState until purging process stops, regardless of
	// successful or failed.
	//
	// See more in PurgingCL doc.
	//
	// Sorted by CL ID.
	PurgingCls []*PurgingCL `protobuf:"bytes,13,rep,name=purging_cls,json=purgingCls,proto3" json:"purging_cls,omitempty"`
	// If true, components partition must be redone as soon as possible.
	RepartitionRequired bool `protobuf:"varint,21,opt,name=repartition_required,json=repartitionRequired,proto3" json:"repartition_required,omitempty"`
	// PRuns which can't yet be added to any component but should be. Sorted by
	// Run ID.
	//
	// In response to OnRunCreated event, PM may append to this list new Runs if
	// either:
	//   * not all Run's CLs are already known to PM;
	//   * Run's CLs are currently partitioned into different components.
	//
	// Thus,
	//   * CLs referenced by these PRuns may not be tracked;
	//   * If this field is not empty, re-partioning may be required.
	CreatedPruns []*PRun `protobuf:"bytes,22,rep,name=created_pruns,json=createdPruns,proto3" json:"created_pruns,omitempty"`
	// If set, establishes when components should be re-evaluated.
	NextEvalTime *timestamppb.Timestamp `protobuf:"bytes,23,opt,name=next_eval_time,json=nextEvalTime,proto3" json:"next_eval_time,omitempty"`
}

func (x *PState) Reset() {
	*x = PState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_storage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PState) ProtoMessage() {}

func (x *PState) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_storage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PState.ProtoReflect.Descriptor instead.
func (*PState) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_storage_proto_rawDescGZIP(), []int{0}
}

func (x *PState) GetLuciProject() string {
	if x != nil {
		return x.LuciProject
	}
	return ""
}

func (x *PState) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_STATUS_UNSPECIFIED
}

func (x *PState) GetConfigHash() string {
	if x != nil {
		return x.ConfigHash
	}
	return ""
}

func (x *PState) GetConfigGroupNames() []string {
	if x != nil {
		return x.ConfigGroupNames
	}
	return nil
}

func (x *PState) GetPcls() []*PCL {
	if x != nil {
		return x.Pcls
	}
	return nil
}

func (x *PState) GetComponents() []*Component {
	if x != nil {
		return x.Components
	}
	return nil
}

func (x *PState) GetPurgingCls() []*PurgingCL {
	if x != nil {
		return x.PurgingCls
	}
	return nil
}

func (x *PState) GetRepartitionRequired() bool {
	if x != nil {
		return x.RepartitionRequired
	}
	return false
}

func (x *PState) GetCreatedPruns() []*PRun {
	if x != nil {
		return x.CreatedPruns
	}
	return nil
}

func (x *PState) GetNextEvalTime() *timestamppb.Timestamp {
	if x != nil {
		return x.NextEvalTime
	}
	return nil
}

// PCL is a tracked CL.
type PCL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clid     int64      `protobuf:"varint,1,opt,name=clid,proto3" json:"clid,omitempty"`
	Eversion int64      `protobuf:"varint,2,opt,name=eversion,proto3" json:"eversion,omitempty"`
	Status   PCL_Status `protobuf:"varint,3,opt,name=status,proto3,enum=cv.prjmanager.prjpb.PCL_Status" json:"status,omitempty"`
	// Indexes in PState.config_group_names identifying ConfigGroup which watches
	// this CL.
	//
	// Normally, contains exactly 1 index.
	// May have > 1 index, which means 2+ non-fallback config groups watch this
	// CL, which is not allowed and will be signalled to CV users.
	// TODO(tandrii): move >1 index case to be tracked via `errors` field.
	ConfigGroupIndexes []int32 `protobuf:"varint,4,rep,packed,name=config_group_indexes,json=configGroupIndexes,proto3" json:"config_group_indexes,omitempty"`
	// Deps refers to CLs in PState.PCLs which are dependencies of the PCL.
	Deps []*changelist.Dep `protobuf:"bytes,11,rep,name=deps,proto3" json:"deps,omitempty"`
	// Trigger is CQDaemon-compatible record of who/when triggered CQ on this CL.
	//
	// It may be nil, if CL is not triggered but nevertheless tracked as either:
	//  * a dependency of another CL.
	//  * previously triggered member of an incomplete Run, which is probably
	//    being finalized right now by its Run Manager.
	//
	// TODO(tandrii): don't store potentially long user's email,
	// which isn't necessary for PM decision making. It should be (re)-computed
	// based on CL snapshots at the time only when actual Run is actually being
	// created.
	Trigger *run.Trigger `protobuf:"bytes,12,opt,name=trigger,proto3" json:"trigger,omitempty"`
	// Submitted means CV isn't going to work on a CL, but CL is still tracked as
	// a dep of another CL or as a member of an incomplete Run (though the other
	// Run will probably finish soon).
	Submitted bool `protobuf:"varint,13,opt,name=submitted,proto3" json:"submitted,omitempty"`
	// If true, CL's owner has no known email.
	// TODO(tandrii): deprecate in favor of `errors`.
	OwnerLacksEmail bool `protobuf:"varint,14,opt,name=owner_lacks_email,json=ownerLacksEmail,proto3" json:"owner_lacks_email,omitempty"`
	// If set, describes one or more problems with CL.
	Errors []*changelist.CLError `protobuf:"bytes,15,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (x *PCL) Reset() {
	*x = PCL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_storage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PCL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PCL) ProtoMessage() {}

func (x *PCL) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_storage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PCL.ProtoReflect.Descriptor instead.
func (*PCL) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_storage_proto_rawDescGZIP(), []int{1}
}

func (x *PCL) GetClid() int64 {
	if x != nil {
		return x.Clid
	}
	return 0
}

func (x *PCL) GetEversion() int64 {
	if x != nil {
		return x.Eversion
	}
	return 0
}

func (x *PCL) GetStatus() PCL_Status {
	if x != nil {
		return x.Status
	}
	return PCL_PCL_STATUS_UNSPECIFIED
}

func (x *PCL) GetConfigGroupIndexes() []int32 {
	if x != nil {
		return x.ConfigGroupIndexes
	}
	return nil
}

func (x *PCL) GetDeps() []*changelist.Dep {
	if x != nil {
		return x.Deps
	}
	return nil
}

func (x *PCL) GetTrigger() *run.Trigger {
	if x != nil {
		return x.Trigger
	}
	return nil
}

func (x *PCL) GetSubmitted() bool {
	if x != nil {
		return x.Submitted
	}
	return false
}

func (x *PCL) GetOwnerLacksEmail() bool {
	if x != nil {
		return x.OwnerLacksEmail
	}
	return false
}

func (x *PCL) GetErrors() []*changelist.CLError {
	if x != nil {
		return x.Errors
	}
	return nil
}

// PRun is an incomplete Run on which CV is currently working.
//
// It is referenced by at most 1 component.
type PRun struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// CV's Run ID.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// IDs of CLs involved. Sorted.
	//
	// Actual Run may orders its CLs in a different way.
	Clids []int64 `protobuf:"varint,2,rep,packed,name=clids,proto3" json:"clids,omitempty"`
}

func (x *PRun) Reset() {
	*x = PRun{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_storage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PRun) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PRun) ProtoMessage() {}

func (x *PRun) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_storage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PRun.ProtoReflect.Descriptor instead.
func (*PRun) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_storage_proto_rawDescGZIP(), []int{2}
}

func (x *PRun) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PRun) GetClids() []int64 {
	if x != nil {
		return x.Clids
	}
	return nil
}

// Component is a set of CLs related to each other.
type Component struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// CL IDs of the tracked CLs in this component. Sorted.
	//
	// Each referenced CL must be in PState.PCLs list.
	// Each referenced CL may have deps not in this list if they are either
	// PCL.Status.UNKNOWN or PCL.Status.UNWATCHED.
	//
	// A referenced CL is normally watched by this LUCI project. In rare cases,
	// referenced CL is no longer watched by this LUCI project but is still kept
	// in a component becaues the CL is still a member of an incomplete Run in
	// this component. In this case, the CL's deps are no longer tracked.
	Clids []int64 `protobuf:"varint,1,rep,packed,name=clids,proto3" json:"clids,omitempty"`
	// Decision time is the earliest time when this component should be
	// re-evaluated.
	//
	// Can be set to far future meaning no need for re-evaluation without an
	// external event (e.g., CLUpdated or RunFinished).
	DecisionTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=decision_time,json=decisionTime,proto3" json:"decision_time,omitempty"`
	// Incomplete Runs working on CLs from this component.
	//
	// Sorted by Run's ID.
	Pruns []*PRun `protobuf:"bytes,3,rep,name=pruns,proto3" json:"pruns,omitempty"`
	// If true, this component must be triaged as soon as possible.
	TriageRequired bool `protobuf:"varint,11,opt,name=triage_required,json=triageRequired,proto3" json:"triage_required,omitempty"`
}

func (x *Component) Reset() {
	*x = Component{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_storage_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Component) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Component) ProtoMessage() {}

func (x *Component) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_storage_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Component.ProtoReflect.Descriptor instead.
func (*Component) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_storage_proto_rawDescGZIP(), []int{3}
}

func (x *Component) GetClids() []int64 {
	if x != nil {
		return x.Clids
	}
	return nil
}

func (x *Component) GetDecisionTime() *timestamppb.Timestamp {
	if x != nil {
		return x.DecisionTime
	}
	return nil
}

func (x *Component) GetPruns() []*PRun {
	if x != nil {
		return x.Pruns
	}
	return nil
}

func (x *Component) GetTriageRequired() bool {
	if x != nil {
		return x.TriageRequired
	}
	return false
}

// PurgingCL represents purging of a CL due to some problem.
//
// The purging process is initiated during PM state mutation while atomically
// adding a TQ task to perform the actual purge.
//
// Purging itself constitutes removing whatever triggered CV on a CL as well as
// posting the reason for purging to the user.
//
// Individual CLs are purged independently, even if CLs are related.
//
// Upon TQ task completion, the task handler notifies PM back via an
// PurgeCompleted event. For fail-safe reasons, there is a deadline to
// perform the purge. PM keeps the PurgingCL in PState until either deadline is
// reached OR PurgeCompleted event is received.
type PurgingCL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// CL ID which is being purged.
	Clid int64 `protobuf:"varint,1,opt,name=clid,proto3" json:"clid,omitempty"`
	// Operation ID is a unique within a project identifier of a purge operation
	// to use in PurgeCompleted events.
	OperationId string `protobuf:"bytes,2,opt,name=operation_id,json=operationId,proto3" json:"operation_id,omitempty"`
	// Deadline is obeyed by the purging TQ task.
	//
	// TQ task SHOULD not modify a CL (e.g. via Gerrit RPCs) beyond this point.
	// This is merely best effort, as an RPC to external system initiated before
	// this deadline may still complete after it.
	//
	// If PM doesn't receive PurgeCompleted event before this deadline + some grace
	// period, PM will consider purge operation expired and it'll be removed from
	// PState.
	Deadline *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=deadline,proto3" json:"deadline,omitempty"`
}

func (x *PurgingCL) Reset() {
	*x = PurgingCL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_storage_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PurgingCL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurgingCL) ProtoMessage() {}

func (x *PurgingCL) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_storage_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PurgingCL.ProtoReflect.Descriptor instead.
func (*PurgingCL) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_storage_proto_rawDescGZIP(), []int{4}
}

func (x *PurgingCL) GetClid() int64 {
	if x != nil {
		return x.Clid
	}
	return 0
}

func (x *PurgingCL) GetOperationId() string {
	if x != nil {
		return x.OperationId
	}
	return ""
}

func (x *PurgingCL) GetDeadline() *timestamppb.Timestamp {
	if x != nil {
		return x.Deadline
	}
	return nil
}

var File_go_chromium_org_luci_cv_internal_prjmanager_prjpb_storage_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_storage_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x76, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6a, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x70, 0x72,
	0x6a, 0x70, 0x62, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x13, 0x63, 0x76, 0x2e, 0x70, 0x72, 0x6a, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6a, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x39, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f,
	0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x76,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x6c, 0x69, 0x73, 0x74, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x32, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e,
	0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x76, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x72, 0x75, 0x6e, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x04, 0x0a, 0x06, 0x50, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x75, 0x63, 0x69, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x75, 0x63, 0x69, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x33, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x63, 0x76, 0x2e, 0x70, 0x72, 0x6a, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6a, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x61, 0x73, 0x68, 0x12, 0x2c, 0x0a, 0x12, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x04, 0x70, 0x63, 0x6c, 0x73,
	0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x76, 0x2e, 0x70, 0x72, 0x6a, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6a, 0x70, 0x62, 0x2e, 0x50, 0x43, 0x4c,
	0x52, 0x04, 0x70, 0x63, 0x6c, 0x73, 0x12, 0x3e, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x76, 0x2e,
	0x70, 0x72, 0x6a, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6a, 0x70, 0x62,
	0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x3f, 0x0a, 0x0b, 0x70, 0x75, 0x72, 0x67, 0x69, 0x6e,
	0x67, 0x5f, 0x63, 0x6c, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x76,
	0x2e, 0x70, 0x72, 0x6a, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6a, 0x70,
	0x62, 0x2e, 0x50, 0x75, 0x72, 0x67, 0x69, 0x6e, 0x67, 0x43, 0x4c, 0x52, 0x0a, 0x70, 0x75, 0x72,
	0x67, 0x69, 0x6e, 0x67, 0x43, 0x6c, 0x73, 0x12, 0x31, 0x0a, 0x14, 0x72, 0x65, 0x70, 0x61, 0x72,
	0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18,
	0x15, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x72, 0x65, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x3e, 0x0a, 0x0d, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x70, 0x72, 0x75, 0x6e, 0x73, 0x18, 0x16, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x63, 0x76, 0x2e, 0x70, 0x72, 0x6a, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6a, 0x70, 0x62, 0x2e, 0x50, 0x52, 0x75, 0x6e, 0x52, 0x0c, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x50, 0x72, 0x75, 0x6e, 0x73, 0x12, 0x40, 0x0a, 0x0e, 0x6e, 0x65,
	0x78, 0x74, 0x5f, 0x65, 0x76, 0x61, 0x6c, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x17, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c,
	0x6e, 0x65, 0x78, 0x74, 0x45, 0x76, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xc8, 0x03, 0x0a,
	0x03, 0x50, 0x43, 0x4c, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6c, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x63, 0x6c, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x65, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x37, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x63, 0x76, 0x2e, 0x70, 0x72, 0x6a, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6a, 0x70, 0x62, 0x2e, 0x50, 0x43, 0x4c, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x30, 0x0a,
	0x14, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x05, 0x52, 0x12, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x12,
	0x26, 0x0a, 0x04, 0x64, 0x65, 0x70, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x63, 0x76, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x44, 0x65,
	0x70, 0x52, 0x04, 0x64, 0x65, 0x70, 0x73, 0x12, 0x29, 0x0a, 0x07, 0x74, 0x72, 0x69, 0x67, 0x67,
	0x65, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x76, 0x2e, 0x72, 0x75,
	0x6e, 0x2e, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x07, 0x74, 0x72, 0x69, 0x67, 0x67,
	0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64,
	0x12, 0x2a, 0x0a, 0x11, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x6c, 0x61, 0x63, 0x6b, 0x73, 0x5f,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x4c, 0x61, 0x63, 0x6b, 0x73, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x2e, 0x0a, 0x06,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63,
	0x76, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x43, 0x4c, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0x59, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x16, 0x50, 0x43, 0x4c, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x4e, 0x57, 0x41, 0x54,
	0x43, 0x48, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45,
	0x44, 0x10, 0x03, 0x1a, 0x02, 0x10, 0x01, 0x22, 0x2c, 0x0a, 0x04, 0x50, 0x52, 0x75, 0x6e, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6c, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x03, 0x52, 0x05,
	0x63, 0x6c, 0x69, 0x64, 0x73, 0x22, 0xbc, 0x01, 0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x03, 0x52, 0x05, 0x63, 0x6c, 0x69, 0x64, 0x73, 0x12, 0x3f, 0x0a, 0x0d, 0x64, 0x65, 0x63,
	0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x64, 0x65,
	0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x05, 0x70, 0x72,
	0x75, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x76, 0x2e, 0x70,
	0x72, 0x6a, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6a, 0x70, 0x62, 0x2e,
	0x50, 0x52, 0x75, 0x6e, 0x52, 0x05, 0x70, 0x72, 0x75, 0x6e, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x74,
	0x72, 0x69, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x74, 0x72, 0x69, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x64, 0x22, 0x7a, 0x0a, 0x09, 0x50, 0x75, 0x72, 0x67, 0x69, 0x6e, 0x67, 0x43,
	0x4c, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6c, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x63, 0x6c, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x08, 0x64, 0x65, 0x61, 0x64,
	0x6c, 0x69, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65,
	0x2a, 0x48, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x54, 0x41, 0x52, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12,
	0x0c, 0x0a, 0x08, 0x53, 0x54, 0x4f, 0x50, 0x50, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0b, 0x0a,
	0x07, 0x53, 0x54, 0x4f, 0x50, 0x50, 0x45, 0x44, 0x10, 0x03, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x6f,
	0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75,
	0x63, 0x69, 0x2f, 0x63, 0x76, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70,
	0x72, 0x6a, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6a, 0x70, 0x62, 0x3b,
	0x70, 0x72, 0x6a, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_storage_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_storage_proto_rawDescData = file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_storage_proto_rawDesc
)

func file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_storage_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_storage_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_storage_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_storage_proto_rawDescData)
	})
	return file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_storage_proto_rawDescData
}

var file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_storage_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_storage_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_storage_proto_goTypes = []interface{}{
	(Status)(0),                   // 0: cv.prjmanager.prjpb.Status
	(PCL_Status)(0),               // 1: cv.prjmanager.prjpb.PCL.Status
	(*PState)(nil),                // 2: cv.prjmanager.prjpb.PState
	(*PCL)(nil),                   // 3: cv.prjmanager.prjpb.PCL
	(*PRun)(nil),                  // 4: cv.prjmanager.prjpb.PRun
	(*Component)(nil),             // 5: cv.prjmanager.prjpb.Component
	(*PurgingCL)(nil),             // 6: cv.prjmanager.prjpb.PurgingCL
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
	(*changelist.Dep)(nil),        // 8: cv.changelist.Dep
	(*run.Trigger)(nil),           // 9: cv.run.Trigger
	(*changelist.CLError)(nil),    // 10: cv.changelist.CLError
}
var file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_storage_proto_depIdxs = []int32{
	0,  // 0: cv.prjmanager.prjpb.PState.status:type_name -> cv.prjmanager.prjpb.Status
	3,  // 1: cv.prjmanager.prjpb.PState.pcls:type_name -> cv.prjmanager.prjpb.PCL
	5,  // 2: cv.prjmanager.prjpb.PState.components:type_name -> cv.prjmanager.prjpb.Component
	6,  // 3: cv.prjmanager.prjpb.PState.purging_cls:type_name -> cv.prjmanager.prjpb.PurgingCL
	4,  // 4: cv.prjmanager.prjpb.PState.created_pruns:type_name -> cv.prjmanager.prjpb.PRun
	7,  // 5: cv.prjmanager.prjpb.PState.next_eval_time:type_name -> google.protobuf.Timestamp
	1,  // 6: cv.prjmanager.prjpb.PCL.status:type_name -> cv.prjmanager.prjpb.PCL.Status
	8,  // 7: cv.prjmanager.prjpb.PCL.deps:type_name -> cv.changelist.Dep
	9,  // 8: cv.prjmanager.prjpb.PCL.trigger:type_name -> cv.run.Trigger
	10, // 9: cv.prjmanager.prjpb.PCL.errors:type_name -> cv.changelist.CLError
	7,  // 10: cv.prjmanager.prjpb.Component.decision_time:type_name -> google.protobuf.Timestamp
	4,  // 11: cv.prjmanager.prjpb.Component.pruns:type_name -> cv.prjmanager.prjpb.PRun
	7,  // 12: cv.prjmanager.prjpb.PurgingCL.deadline:type_name -> google.protobuf.Timestamp
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_storage_proto_init() }
func file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_storage_proto_init() {
	if File_go_chromium_org_luci_cv_internal_prjmanager_prjpb_storage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_storage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PState); i {
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
		file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_storage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PCL); i {
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
		file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_storage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PRun); i {
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
		file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_storage_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Component); i {
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
		file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_storage_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PurgingCL); i {
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
			RawDescriptor: file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_storage_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_storage_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_storage_proto_depIdxs,
		EnumInfos:         file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_storage_proto_enumTypes,
		MessageInfos:      file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_storage_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_cv_internal_prjmanager_prjpb_storage_proto = out.File
	file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_storage_proto_rawDesc = nil
	file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_storage_proto_goTypes = nil
	file_go_chromium_org_luci_cv_internal_prjmanager_prjpb_storage_proto_depIdxs = nil
}
