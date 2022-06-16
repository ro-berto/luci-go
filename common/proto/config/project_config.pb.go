// Copyright 2014 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Schemas for some of config files in projects/<project_id> config sets.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: go.chromium.org/luci/common/proto/config/project_config.proto

package config

import (
	_ "go.chromium.org/luci/common/proto"
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

// Schema for project.cfg
type ProjectCfg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Full name of the project.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A list of identities that have read-only access to the project.
	//
	// An element is one of:
	// * "group:<group>", where group is defined on auth server.
	// * "<email>"
	// * "<kind>:<value>" (for non-email identities)
	//
	// If not specified, only admins and trusted services have access.
	// Talk to admins to determine the group name appropriate for your project.
	Access []string `protobuf:"bytes,2,rep,name=access,proto3" json:"access,omitempty"`
	// Populated when the config is generated by `lucicfg`.
	//
	// Contains information about how the config was generated.
	Lucicfg *GeneratorMetadata `protobuf:"bytes,3,opt,name=lucicfg,proto3" json:"lucicfg,omitempty"`
}

func (x *ProjectCfg) Reset() {
	*x = ProjectCfg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_common_proto_config_project_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectCfg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectCfg) ProtoMessage() {}

func (x *ProjectCfg) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_common_proto_config_project_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectCfg.ProtoReflect.Descriptor instead.
func (*ProjectCfg) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_common_proto_config_project_config_proto_rawDescGZIP(), []int{0}
}

func (x *ProjectCfg) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProjectCfg) GetAccess() []string {
	if x != nil {
		return x.Access
	}
	return nil
}

func (x *ProjectCfg) GetLucicfg() *GeneratorMetadata {
	if x != nil {
		return x.Lucicfg
	}
	return nil
}

// GeneratorMetadata is produced by lucicfg to describe how it generated the
// config.
type GeneratorMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Version of lucicfg binary used to generate the config.
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// Path to the main package relative to project.cfg.
	PackageDir string `protobuf:"bytes,5,opt,name=package_dir,json=packageDir,proto3" json:"package_dir,omitempty"`
	// Directory with generated LUCI project configs relative to the main package.
	ConfigDir string `protobuf:"bytes,2,opt,name=config_dir,json=configDir,proto3" json:"config_dir,omitempty"`
	// Name of the entry point Starlark file at the root of the main package.
	EntryPoint string `protobuf:"bytes,3,opt,name=entry_point,json=entryPoint,proto3" json:"entry_point,omitempty"`
	// Set of vars passed to lucicfg as "-var ...".
	Vars map[string]string `protobuf:"bytes,4,rep,name=vars,proto3" json:"vars,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Set of enabled lucicfg experiments.
	Experiments []string `protobuf:"bytes,6,rep,name=experiments,proto3" json:"experiments,omitempty"`
}

func (x *GeneratorMetadata) Reset() {
	*x = GeneratorMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_common_proto_config_project_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeneratorMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeneratorMetadata) ProtoMessage() {}

func (x *GeneratorMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_common_proto_config_project_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeneratorMetadata.ProtoReflect.Descriptor instead.
func (*GeneratorMetadata) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_common_proto_config_project_config_proto_rawDescGZIP(), []int{1}
}

func (x *GeneratorMetadata) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *GeneratorMetadata) GetPackageDir() string {
	if x != nil {
		return x.PackageDir
	}
	return ""
}

func (x *GeneratorMetadata) GetConfigDir() string {
	if x != nil {
		return x.ConfigDir
	}
	return ""
}

func (x *GeneratorMetadata) GetEntryPoint() string {
	if x != nil {
		return x.EntryPoint
	}
	return ""
}

func (x *GeneratorMetadata) GetVars() map[string]string {
	if x != nil {
		return x.Vars
	}
	return nil
}

func (x *GeneratorMetadata) GetExperiments() []string {
	if x != nil {
		return x.Experiments
	}
	return nil
}

// DEPRECATED per crbug/924803. DO NOT USE in your project.
// Contact luci-team@ if you have a need for this.
//
// Schema of refs.cfg.
type RefsCfg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of refs that have configuration files and need to be fetched into
	// luci-config. Refs are accessible through get_refs() API endpoint.
	// A CI service can read all refs of all projects and build them.
	Refs []*RefsCfg_Ref `protobuf:"bytes,1,rep,name=refs,proto3" json:"refs,omitempty"`
}

func (x *RefsCfg) Reset() {
	*x = RefsCfg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_common_proto_config_project_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefsCfg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefsCfg) ProtoMessage() {}

func (x *RefsCfg) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_common_proto_config_project_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefsCfg.ProtoReflect.Descriptor instead.
func (*RefsCfg) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_common_proto_config_project_config_proto_rawDescGZIP(), []int{2}
}

func (x *RefsCfg) GetRefs() []*RefsCfg_Ref {
	if x != nil {
		return x.Refs
	}
	return nil
}

type RefsCfg_Ref struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the ref. Must start with "refs/".
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Path to config directory for the ref. Defaults to "infra/config".
	ConfigPath string `protobuf:"bytes,3,opt,name=config_path,json=configPath,proto3" json:"config_path,omitempty"`
}

func (x *RefsCfg_Ref) Reset() {
	*x = RefsCfg_Ref{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_common_proto_config_project_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefsCfg_Ref) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefsCfg_Ref) ProtoMessage() {}

func (x *RefsCfg_Ref) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_common_proto_config_project_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefsCfg_Ref.ProtoReflect.Descriptor instead.
func (*RefsCfg_Ref) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_common_proto_config_project_config_proto_rawDescGZIP(), []int{2, 0}
}

func (x *RefsCfg_Ref) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RefsCfg_Ref) GetConfigPath() string {
	if x != nil {
		return x.ConfigPath
	}
	return ""
}

var File_go_chromium_org_luci_common_proto_config_project_config_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_common_proto_config_project_config_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x2f, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f,
	0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x73, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x43, 0x66, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x39, 0x0a, 0x07, 0x6c, 0x75, 0x63, 0x69, 0x63, 0x66, 0x67, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x04,
	0xb0, 0xfe, 0x23, 0x01, 0x52, 0x07, 0x6c, 0x75, 0x63, 0x69, 0x63, 0x66, 0x67, 0x22, 0xa2, 0x02,
	0x0a, 0x11, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a,
	0x0b, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x64, 0x69, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x44, 0x69, 0x72, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x64, 0x69, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x44, 0x69, 0x72, 0x12, 0x1f, 0x0a,
	0x0b, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x37,
	0x0a, 0x04, 0x76, 0x61, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x56, 0x61, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x04, 0x76, 0x61, 0x72, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x65, 0x72,
	0x69, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x78,
	0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x37, 0x0a, 0x09, 0x56, 0x61, 0x72,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x6e, 0x0a, 0x07, 0x52, 0x65, 0x66, 0x73, 0x43, 0x66, 0x67, 0x12, 0x27, 0x0a,
	0x04, 0x72, 0x65, 0x66, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x52, 0x65, 0x66, 0x73, 0x43, 0x66, 0x67, 0x2e, 0x52, 0x65, 0x66,
	0x52, 0x04, 0x72, 0x65, 0x66, 0x73, 0x1a, 0x3a, 0x0a, 0x03, 0x52, 0x65, 0x66, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x61,
	0x74, 0x68, 0x42, 0x6c, 0x5a, 0x28, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75,
	0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0xa2, 0xfe,
	0x23, 0x3e, 0x0a, 0x3c, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x6c, 0x75, 0x63, 0x69,
	0x2d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x70, 0x6f, 0x74, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x3a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x63, 0x66, 0x67,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_common_proto_config_project_config_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_common_proto_config_project_config_proto_rawDescData = file_go_chromium_org_luci_common_proto_config_project_config_proto_rawDesc
)

func file_go_chromium_org_luci_common_proto_config_project_config_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_common_proto_config_project_config_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_common_proto_config_project_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_common_proto_config_project_config_proto_rawDescData)
	})
	return file_go_chromium_org_luci_common_proto_config_project_config_proto_rawDescData
}

var file_go_chromium_org_luci_common_proto_config_project_config_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_go_chromium_org_luci_common_proto_config_project_config_proto_goTypes = []interface{}{
	(*ProjectCfg)(nil),        // 0: config.ProjectCfg
	(*GeneratorMetadata)(nil), // 1: config.GeneratorMetadata
	(*RefsCfg)(nil),           // 2: config.RefsCfg
	nil,                       // 3: config.GeneratorMetadata.VarsEntry
	(*RefsCfg_Ref)(nil),       // 4: config.RefsCfg.Ref
}
var file_go_chromium_org_luci_common_proto_config_project_config_proto_depIdxs = []int32{
	1, // 0: config.ProjectCfg.lucicfg:type_name -> config.GeneratorMetadata
	3, // 1: config.GeneratorMetadata.vars:type_name -> config.GeneratorMetadata.VarsEntry
	4, // 2: config.RefsCfg.refs:type_name -> config.RefsCfg.Ref
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_common_proto_config_project_config_proto_init() }
func file_go_chromium_org_luci_common_proto_config_project_config_proto_init() {
	if File_go_chromium_org_luci_common_proto_config_project_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_common_proto_config_project_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectCfg); i {
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
		file_go_chromium_org_luci_common_proto_config_project_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeneratorMetadata); i {
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
		file_go_chromium_org_luci_common_proto_config_project_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefsCfg); i {
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
		file_go_chromium_org_luci_common_proto_config_project_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefsCfg_Ref); i {
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
			RawDescriptor: file_go_chromium_org_luci_common_proto_config_project_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_common_proto_config_project_config_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_common_proto_config_project_config_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_common_proto_config_project_config_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_common_proto_config_project_config_proto = out.File
	file_go_chromium_org_luci_common_proto_config_project_config_proto_rawDesc = nil
	file_go_chromium_org_luci_common_proto_config_project_config_proto_goTypes = nil
	file_go_chromium_org_luci_common_proto_config_project_config_proto_depIdxs = nil
}
