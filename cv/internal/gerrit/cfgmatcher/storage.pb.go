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
// 	protoc        v3.12.1
// source: go.chromium.org/luci/cv/internal/gerrit/cfgmatcher/storage.proto

package cfgmatcher

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

// Groups keeps config groups of a single LUCI project relevant to a specific
// Git repository (aka Gerrit project).
//
// For now, this message is just a wrapper for a list of groups as they appear
// in LUCI Project's CV config file. In the future, this can optimized into
// treap-like structure based on known ref prefix if there are 100+ ref specs
// for the same repo.
type Groups struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Groups []*Group `protobuf:"bytes,1,rep,name=groups,proto3" json:"groups,omitempty"`
}

func (x *Groups) Reset() {
	*x = Groups{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_internal_gerrit_cfgmatcher_storage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Groups) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Groups) ProtoMessage() {}

func (x *Groups) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_internal_gerrit_cfgmatcher_storage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Groups.ProtoReflect.Descriptor instead.
func (*Groups) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_internal_gerrit_cfgmatcher_storage_proto_rawDescGZIP(), []int{0}
}

func (x *Groups) GetGroups() []*Group {
	if x != nil {
		return x.Groups
	}
	return nil
}

// Group represents one config group applied to just 1 Git repository.
//
// For full documentation, see ConfigGroup of api/config/v2/cq.proto.
type Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ConfigGroupID, as stored in ConfigGroup datastore entity.
	//
	// Used by gobmap.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Index of the ConfigGroup names interned in MatcherState.
	//
	// Used by MatcherState.
	Index int32 `protobuf:"varint,3,opt,name=index,proto3" json:"index,omitempty"`
	// If set, this ConfigGroup will be selected if no other ConfigGroup matches
	// refspec. At most 1 group will have this set (this is validated before
	// config is injected).
	Fallback bool `protobuf:"varint,2,opt,name=fallback,proto3" json:"fallback,omitempty"`
	// Regular expression that a CL's target ref must match. Required.
	//
	// It's constructed from `ref_regexp`s of CV config.
	Include string `protobuf:"bytes,13,opt,name=include,proto3" json:"include,omitempty"`
	// Regular expression that a CL's target ref must NOT match. Required.
	//
	// It's constructed from `ref_regexp_exclude`s of CV config.
	Exclude string `protobuf:"bytes,14,opt,name=exclude,proto3" json:"exclude,omitempty"`
}

func (x *Group) Reset() {
	*x = Group{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_internal_gerrit_cfgmatcher_storage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_internal_gerrit_cfgmatcher_storage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Group.ProtoReflect.Descriptor instead.
func (*Group) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_internal_gerrit_cfgmatcher_storage_proto_rawDescGZIP(), []int{1}
}

func (x *Group) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Group) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Group) GetFallback() bool {
	if x != nil {
		return x.Fallback
	}
	return false
}

func (x *Group) GetInclude() string {
	if x != nil {
		return x.Include
	}
	return ""
}

func (x *Group) GetExclude() string {
	if x != nil {
		return x.Exclude
	}
	return ""
}

// MatcherState is serializable state of a matcher for a single LUCI project at
// specific config hash (version).
type MatcherState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConfigHash string `protobuf:"bytes,1,opt,name=config_hash,json=configHash,proto3" json:"config_hash,omitempty"`
	// Interned config group names.
	//
	// Combine with config_hash to obtain config.ConfigGroupID.
	ConfigGroupNames []string `protobuf:"bytes,2,rep,name=config_group_names,json=configGroupNames,proto3" json:"config_group_names,omitempty"`
	// Maps Gerrit hosts to watched projects.
	Hosts map[string]*MatcherState_Projects `protobuf:"bytes,3,rep,name=hosts,proto3" json:"hosts,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MatcherState) Reset() {
	*x = MatcherState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_internal_gerrit_cfgmatcher_storage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatcherState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatcherState) ProtoMessage() {}

func (x *MatcherState) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_internal_gerrit_cfgmatcher_storage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatcherState.ProtoReflect.Descriptor instead.
func (*MatcherState) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_internal_gerrit_cfgmatcher_storage_proto_rawDescGZIP(), []int{2}
}

func (x *MatcherState) GetConfigHash() string {
	if x != nil {
		return x.ConfigHash
	}
	return ""
}

func (x *MatcherState) GetConfigGroupNames() []string {
	if x != nil {
		return x.ConfigGroupNames
	}
	return nil
}

func (x *MatcherState) GetHosts() map[string]*MatcherState_Projects {
	if x != nil {
		return x.Hosts
	}
	return nil
}

type MatcherState_Projects struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Maps Gerrit project (aka Gerrit repo) to one or more config groups.
	Projects map[string]*Groups `protobuf:"bytes,1,rep,name=projects,proto3" json:"projects,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MatcherState_Projects) Reset() {
	*x = MatcherState_Projects{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_internal_gerrit_cfgmatcher_storage_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatcherState_Projects) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatcherState_Projects) ProtoMessage() {}

func (x *MatcherState_Projects) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_internal_gerrit_cfgmatcher_storage_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatcherState_Projects.ProtoReflect.Descriptor instead.
func (*MatcherState_Projects) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_internal_gerrit_cfgmatcher_storage_proto_rawDescGZIP(), []int{2, 1}
}

func (x *MatcherState_Projects) GetProjects() map[string]*Groups {
	if x != nil {
		return x.Projects
	}
	return nil
}

var File_go_chromium_org_luci_cv_internal_gerrit_cfgmatcher_storage_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_cv_internal_gerrit_cfgmatcher_storage_proto_rawDesc = []byte{
	0x0a, 0x40, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x76, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x67, 0x65, 0x72, 0x72, 0x69, 0x74, 0x2f, 0x63, 0x66, 0x67, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x14, 0x63, 0x76, 0x2e, 0x67, 0x65, 0x72, 0x72, 0x69, 0x74, 0x2e, 0x63, 0x66,
	0x67, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x22, 0x3d, 0x0a, 0x06, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x12, 0x33, 0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x76, 0x2e, 0x67, 0x65, 0x72, 0x72, 0x69, 0x74, 0x2e, 0x63,
	0x66, 0x67, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x22, 0x7d, 0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x61, 0x6c, 0x6c, 0x62, 0x61,
	0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x66, 0x61, 0x6c, 0x6c, 0x62, 0x61,
	0x63, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65,
	0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x22, 0xc8, 0x03, 0x0a, 0x0c, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x48, 0x61, 0x73, 0x68, 0x12, 0x2c, 0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x43, 0x0a, 0x05, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x76, 0x2e, 0x67, 0x65, 0x72, 0x72, 0x69,
	0x74, 0x2e, 0x63, 0x66, 0x67, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x1a, 0x65, 0x0a, 0x0a, 0x48,
	0x6f, 0x73, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x41, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x63, 0x76, 0x2e,
	0x67, 0x65, 0x72, 0x72, 0x69, 0x74, 0x2e, 0x63, 0x66, 0x67, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x1a, 0xbc, 0x01, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12,
	0x55, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x39, 0x2e, 0x63, 0x76, 0x2e, 0x67, 0x65, 0x72, 0x72, 0x69, 0x74, 0x2e, 0x63, 0x66,
	0x67, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x1a, 0x59, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x32, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x76, 0x2e, 0x67, 0x65,
	0x72, 0x72, 0x69, 0x74, 0x2e, 0x63, 0x66, 0x67, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d,
	0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x76, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x72, 0x72, 0x69, 0x74, 0x2f, 0x63, 0x66, 0x67,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x3b, 0x63, 0x66, 0x67, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_cv_internal_gerrit_cfgmatcher_storage_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_cv_internal_gerrit_cfgmatcher_storage_proto_rawDescData = file_go_chromium_org_luci_cv_internal_gerrit_cfgmatcher_storage_proto_rawDesc
)

func file_go_chromium_org_luci_cv_internal_gerrit_cfgmatcher_storage_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_cv_internal_gerrit_cfgmatcher_storage_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_cv_internal_gerrit_cfgmatcher_storage_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_cv_internal_gerrit_cfgmatcher_storage_proto_rawDescData)
	})
	return file_go_chromium_org_luci_cv_internal_gerrit_cfgmatcher_storage_proto_rawDescData
}

var file_go_chromium_org_luci_cv_internal_gerrit_cfgmatcher_storage_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_go_chromium_org_luci_cv_internal_gerrit_cfgmatcher_storage_proto_goTypes = []interface{}{
	(*Groups)(nil),                // 0: cv.gerrit.cfgmatcher.Groups
	(*Group)(nil),                 // 1: cv.gerrit.cfgmatcher.Group
	(*MatcherState)(nil),          // 2: cv.gerrit.cfgmatcher.MatcherState
	nil,                           // 3: cv.gerrit.cfgmatcher.MatcherState.HostsEntry
	(*MatcherState_Projects)(nil), // 4: cv.gerrit.cfgmatcher.MatcherState.Projects
	nil,                           // 5: cv.gerrit.cfgmatcher.MatcherState.Projects.ProjectsEntry
}
var file_go_chromium_org_luci_cv_internal_gerrit_cfgmatcher_storage_proto_depIdxs = []int32{
	1, // 0: cv.gerrit.cfgmatcher.Groups.groups:type_name -> cv.gerrit.cfgmatcher.Group
	3, // 1: cv.gerrit.cfgmatcher.MatcherState.hosts:type_name -> cv.gerrit.cfgmatcher.MatcherState.HostsEntry
	4, // 2: cv.gerrit.cfgmatcher.MatcherState.HostsEntry.value:type_name -> cv.gerrit.cfgmatcher.MatcherState.Projects
	5, // 3: cv.gerrit.cfgmatcher.MatcherState.Projects.projects:type_name -> cv.gerrit.cfgmatcher.MatcherState.Projects.ProjectsEntry
	0, // 4: cv.gerrit.cfgmatcher.MatcherState.Projects.ProjectsEntry.value:type_name -> cv.gerrit.cfgmatcher.Groups
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_cv_internal_gerrit_cfgmatcher_storage_proto_init() }
func file_go_chromium_org_luci_cv_internal_gerrit_cfgmatcher_storage_proto_init() {
	if File_go_chromium_org_luci_cv_internal_gerrit_cfgmatcher_storage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_cv_internal_gerrit_cfgmatcher_storage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Groups); i {
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
		file_go_chromium_org_luci_cv_internal_gerrit_cfgmatcher_storage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Group); i {
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
		file_go_chromium_org_luci_cv_internal_gerrit_cfgmatcher_storage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatcherState); i {
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
		file_go_chromium_org_luci_cv_internal_gerrit_cfgmatcher_storage_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatcherState_Projects); i {
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
			RawDescriptor: file_go_chromium_org_luci_cv_internal_gerrit_cfgmatcher_storage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_cv_internal_gerrit_cfgmatcher_storage_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_cv_internal_gerrit_cfgmatcher_storage_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_cv_internal_gerrit_cfgmatcher_storage_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_cv_internal_gerrit_cfgmatcher_storage_proto = out.File
	file_go_chromium_org_luci_cv_internal_gerrit_cfgmatcher_storage_proto_rawDesc = nil
	file_go_chromium_org_luci_cv_internal_gerrit_cfgmatcher_storage_proto_goTypes = nil
	file_go_chromium_org_luci_cv_internal_gerrit_cfgmatcher_storage_proto_depIdxs = nil
}
