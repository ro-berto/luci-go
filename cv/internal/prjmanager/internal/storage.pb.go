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
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.1
// source: go.chromium.org/luci/cv/internal/prjmanager/internal/storage.proto

package internal

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
	return file_go_chromium_org_luci_cv_internal_prjmanager_internal_storage_proto_enumTypes[0].Descriptor()
}

func (PCL_Status) Type() protoreflect.EnumType {
	return &file_go_chromium_org_luci_cv_internal_prjmanager_internal_storage_proto_enumTypes[0]
}

func (x PCL_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PCL_Status.Descriptor instead.
func (PCL_Status) EnumDescriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_internal_prjmanager_internal_storage_proto_rawDescGZIP(), []int{1, 0}
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
	// If true, components partition must be redone as soon as possible.
	DirtyComponents bool `protobuf:"varint,21,opt,name=dirty_components,json=dirtyComponents,proto3" json:"dirty_components,omitempty"`
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
}

func (x *PState) Reset() {
	*x = PState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_internal_prjmanager_internal_storage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PState) ProtoMessage() {}

func (x *PState) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_internal_prjmanager_internal_storage_proto_msgTypes[0]
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
	return file_go_chromium_org_luci_cv_internal_prjmanager_internal_storage_proto_rawDescGZIP(), []int{0}
}

func (x *PState) GetLuciProject() string {
	if x != nil {
		return x.LuciProject
	}
	return ""
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

func (x *PState) GetDirtyComponents() bool {
	if x != nil {
		return x.DirtyComponents
	}
	return false
}

func (x *PState) GetCreatedPruns() []*PRun {
	if x != nil {
		return x.CreatedPruns
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
	Status   PCL_Status `protobuf:"varint,3,opt,name=status,proto3,enum=cv.prjmanager.internal.PCL_Status" json:"status,omitempty"`
	// Indexes in PState.config_group_names identifying ConfigGroup which watches
	// this CL.
	//
	// Normally, contains exactly 1 index.
	// May have > 1 index, which means 2+ non-fallback config groups watch this
	// CL, which is not allowed and will be signalled to CV users.
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
}

func (x *PCL) Reset() {
	*x = PCL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_internal_prjmanager_internal_storage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PCL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PCL) ProtoMessage() {}

func (x *PCL) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_internal_prjmanager_internal_storage_proto_msgTypes[1]
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
	return file_go_chromium_org_luci_cv_internal_prjmanager_internal_storage_proto_rawDescGZIP(), []int{1}
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
		mi := &file_go_chromium_org_luci_cv_internal_prjmanager_internal_storage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PRun) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PRun) ProtoMessage() {}

func (x *PRun) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_internal_prjmanager_internal_storage_proto_msgTypes[2]
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
	return file_go_chromium_org_luci_cv_internal_prjmanager_internal_storage_proto_rawDescGZIP(), []int{2}
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
	// If true, this component must be re-evaluated as soon as possible.
	Dirty bool `protobuf:"varint,11,opt,name=dirty,proto3" json:"dirty,omitempty"`
}

func (x *Component) Reset() {
	*x = Component{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_internal_prjmanager_internal_storage_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Component) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Component) ProtoMessage() {}

func (x *Component) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_internal_prjmanager_internal_storage_proto_msgTypes[3]
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
	return file_go_chromium_org_luci_cv_internal_prjmanager_internal_storage_proto_rawDescGZIP(), []int{3}
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

func (x *Component) GetDirty() bool {
	if x != nil {
		return x.Dirty
	}
	return false
}

var File_go_chromium_org_luci_cv_internal_prjmanager_internal_storage_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_cv_internal_prjmanager_internal_storage_proto_rawDesc = []byte{
	0x0a, 0x42, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x76, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6a, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x63, 0x76, 0x2e, 0x70, 0x72, 0x6a, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x39, 0x67,
	0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c,
	0x75, 0x63, 0x69, 0x2f, 0x63, 0x76, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72,
	0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63,
	0x76, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x72, 0x75, 0x6e, 0x2f, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdc, 0x02, 0x0a,
	0x06, 0x50, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x75, 0x63, 0x69, 0x5f,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c,
	0x75, 0x63, 0x69, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x61, 0x73, 0x68, 0x12, 0x2c, 0x0a, 0x12, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x2f, 0x0a, 0x04, 0x70, 0x63, 0x6c,
	0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x76, 0x2e, 0x70, 0x72, 0x6a,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2e, 0x50, 0x43, 0x4c, 0x52, 0x04, 0x70, 0x63, 0x6c, 0x73, 0x12, 0x41, 0x0a, 0x0a, 0x63, 0x6f,
	0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x63, 0x76, 0x2e, 0x70, 0x72, 0x6a, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x29, 0x0a,
	0x10, 0x64, 0x69, 0x72, 0x74, 0x79, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x73, 0x18, 0x15, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x64, 0x69, 0x72, 0x74, 0x79, 0x43, 0x6f,
	0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x41, 0x0a, 0x0d, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x70, 0x72, 0x75, 0x6e, 0x73, 0x18, 0x16, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x63, 0x76, 0x2e, 0x70, 0x72, 0x6a, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x50, 0x52, 0x75, 0x6e, 0x52, 0x0c, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x50, 0x72, 0x75, 0x6e, 0x73, 0x22, 0xd1, 0x02, 0x0a, 0x03,
	0x50, 0x43, 0x4c, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6c, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x63, 0x6c, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x65, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x3a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x63, 0x76, 0x2e, 0x70, 0x72, 0x6a, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x50, 0x43, 0x4c,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x30, 0x0a, 0x14, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x05, 0x52, 0x12, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x65,
	0x73, 0x12, 0x26, 0x0a, 0x04, 0x64, 0x65, 0x70, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x63, 0x76, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x2e,
	0x44, 0x65, 0x70, 0x52, 0x04, 0x64, 0x65, 0x70, 0x73, 0x12, 0x29, 0x0a, 0x07, 0x74, 0x72, 0x69,
	0x67, 0x67, 0x65, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x76, 0x2e,
	0x72, 0x75, 0x6e, 0x2e, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x07, 0x74, 0x72, 0x69,
	0x67, 0x67, 0x65, 0x72, 0x22, 0x59, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a,
	0x0a, 0x16, 0x50, 0x43, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b,
	0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12,
	0x0d, 0x0a, 0x09, 0x55, 0x4e, 0x57, 0x41, 0x54, 0x43, 0x48, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0b,
	0x0a, 0x07, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x03, 0x1a, 0x02, 0x10, 0x01, 0x22,
	0x2c, 0x0a, 0x04, 0x50, 0x52, 0x75, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x69, 0x64, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6c, 0x69, 0x64, 0x73, 0x22, 0xac, 0x01,
	0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6c, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6c, 0x69, 0x64,
	0x73, 0x12, 0x3f, 0x0a, 0x0d, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x70, 0x72, 0x75, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x76, 0x2e, 0x70, 0x72, 0x6a, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x50, 0x52, 0x75, 0x6e, 0x52,
	0x05, 0x70, 0x72, 0x75, 0x6e, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x69, 0x72, 0x74, 0x79, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x64, 0x69, 0x72, 0x74, 0x79, 0x42, 0x3f, 0x5a, 0x3d,
	0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x76, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x70, 0x72, 0x6a, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x3b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_cv_internal_prjmanager_internal_storage_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_cv_internal_prjmanager_internal_storage_proto_rawDescData = file_go_chromium_org_luci_cv_internal_prjmanager_internal_storage_proto_rawDesc
)

func file_go_chromium_org_luci_cv_internal_prjmanager_internal_storage_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_cv_internal_prjmanager_internal_storage_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_cv_internal_prjmanager_internal_storage_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_cv_internal_prjmanager_internal_storage_proto_rawDescData)
	})
	return file_go_chromium_org_luci_cv_internal_prjmanager_internal_storage_proto_rawDescData
}

var file_go_chromium_org_luci_cv_internal_prjmanager_internal_storage_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_go_chromium_org_luci_cv_internal_prjmanager_internal_storage_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_go_chromium_org_luci_cv_internal_prjmanager_internal_storage_proto_goTypes = []interface{}{
	(PCL_Status)(0),               // 0: cv.prjmanager.internal.PCL.Status
	(*PState)(nil),                // 1: cv.prjmanager.internal.PState
	(*PCL)(nil),                   // 2: cv.prjmanager.internal.PCL
	(*PRun)(nil),                  // 3: cv.prjmanager.internal.PRun
	(*Component)(nil),             // 4: cv.prjmanager.internal.Component
	(*changelist.Dep)(nil),        // 5: cv.changelist.Dep
	(*run.Trigger)(nil),           // 6: cv.run.Trigger
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_go_chromium_org_luci_cv_internal_prjmanager_internal_storage_proto_depIdxs = []int32{
	2, // 0: cv.prjmanager.internal.PState.pcls:type_name -> cv.prjmanager.internal.PCL
	4, // 1: cv.prjmanager.internal.PState.components:type_name -> cv.prjmanager.internal.Component
	3, // 2: cv.prjmanager.internal.PState.created_pruns:type_name -> cv.prjmanager.internal.PRun
	0, // 3: cv.prjmanager.internal.PCL.status:type_name -> cv.prjmanager.internal.PCL.Status
	5, // 4: cv.prjmanager.internal.PCL.deps:type_name -> cv.changelist.Dep
	6, // 5: cv.prjmanager.internal.PCL.trigger:type_name -> cv.run.Trigger
	7, // 6: cv.prjmanager.internal.Component.decision_time:type_name -> google.protobuf.Timestamp
	3, // 7: cv.prjmanager.internal.Component.pruns:type_name -> cv.prjmanager.internal.PRun
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_cv_internal_prjmanager_internal_storage_proto_init() }
func file_go_chromium_org_luci_cv_internal_prjmanager_internal_storage_proto_init() {
	if File_go_chromium_org_luci_cv_internal_prjmanager_internal_storage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_cv_internal_prjmanager_internal_storage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_go_chromium_org_luci_cv_internal_prjmanager_internal_storage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_go_chromium_org_luci_cv_internal_prjmanager_internal_storage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_go_chromium_org_luci_cv_internal_prjmanager_internal_storage_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_go_chromium_org_luci_cv_internal_prjmanager_internal_storage_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_cv_internal_prjmanager_internal_storage_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_cv_internal_prjmanager_internal_storage_proto_depIdxs,
		EnumInfos:         file_go_chromium_org_luci_cv_internal_prjmanager_internal_storage_proto_enumTypes,
		MessageInfos:      file_go_chromium_org_luci_cv_internal_prjmanager_internal_storage_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_cv_internal_prjmanager_internal_storage_proto = out.File
	file_go_chromium_org_luci_cv_internal_prjmanager_internal_storage_proto_rawDesc = nil
	file_go_chromium_org_luci_cv_internal_prjmanager_internal_storage_proto_goTypes = nil
	file_go_chromium_org_luci_cv_internal_prjmanager_internal_storage_proto_depIdxs = nil
}
