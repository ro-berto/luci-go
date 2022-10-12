// Copyright (c) 2016 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: go.chromium.org/luci/milo/api/config/settings.proto

package config

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Settings represents the format for the global (service) config for Milo.
type Settings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Buildbucket *Settings_Buildbucket `protobuf:"bytes,2,opt,name=buildbucket,proto3" json:"buildbucket,omitempty"`
	Swarming    *Settings_Swarming    `protobuf:"bytes,3,opt,name=swarming,proto3" json:"swarming,omitempty"`
	// source_acls instructs Milo to provide Git/Gerrit data
	// (e.g., blamelist) to some of its users on entire subdomains or individual
	// repositories (Gerrit "projects").
	//
	// Multiple records are allowed, but each host and project must appear only in
	// one record.
	//
	// See go/milo-git-acls for design rationales.
	SourceAcls   []*Settings_SourceAcls `protobuf:"bytes,4,rep,name=source_acls,json=sourceAcls,proto3" json:"source_acls,omitempty"`
	Resultdb     *Settings_ResultDB     `protobuf:"bytes,5,opt,name=resultdb,proto3" json:"resultdb,omitempty"`
	LuciAnalysis *Settings_LuciAnalysis `protobuf:"bytes,7,opt,name=luci_analysis,json=luciAnalysis,proto3" json:"luci_analysis,omitempty"`
}

func (x *Settings) Reset() {
	*x = Settings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_milo_api_config_settings_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Settings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Settings) ProtoMessage() {}

func (x *Settings) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_milo_api_config_settings_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Settings.ProtoReflect.Descriptor instead.
func (*Settings) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_milo_api_config_settings_proto_rawDescGZIP(), []int{0}
}

func (x *Settings) GetBuildbucket() *Settings_Buildbucket {
	if x != nil {
		return x.Buildbucket
	}
	return nil
}

func (x *Settings) GetSwarming() *Settings_Swarming {
	if x != nil {
		return x.Swarming
	}
	return nil
}

func (x *Settings) GetSourceAcls() []*Settings_SourceAcls {
	if x != nil {
		return x.SourceAcls
	}
	return nil
}

func (x *Settings) GetResultdb() *Settings_ResultDB {
	if x != nil {
		return x.Resultdb
	}
	return nil
}

func (x *Settings) GetLuciAnalysis() *Settings_LuciAnalysis {
	if x != nil {
		return x.LuciAnalysis
	}
	return nil
}

type Settings_Buildbucket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name is the user friendly name of the Buildbucket instance we're pointing to.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// host is the hostname of the buildbucket instance we're pointing to (sans scheme).
	Host string `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	// project is the name of the Google Cloud project that the pubsub topic
	// belongs to.
	//
	// Deprecated: this is no longer used. The buildbucket subscription should
	// be configured via GCP console.
	// TODO(crbug/1255983): set up subscription configuration via terraform.
	Project string `protobuf:"bytes,3,opt,name=project,proto3" json:"project,omitempty"`
}

func (x *Settings_Buildbucket) Reset() {
	*x = Settings_Buildbucket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_milo_api_config_settings_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Settings_Buildbucket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Settings_Buildbucket) ProtoMessage() {}

func (x *Settings_Buildbucket) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_milo_api_config_settings_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Settings_Buildbucket.ProtoReflect.Descriptor instead.
func (*Settings_Buildbucket) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_milo_api_config_settings_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Settings_Buildbucket) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Settings_Buildbucket) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *Settings_Buildbucket) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

type Settings_Swarming struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// default_host is the hostname of the swarming host Milo defaults to, if
	// none is specified.  Default host is implicitly an allowed host.
	DefaultHost string `protobuf:"bytes,1,opt,name=default_host,json=defaultHost,proto3" json:"default_host,omitempty"`
	// allowed_hosts is a list of hostnames of swarming instances that Milo is
	// allowed to talk to.  This is specified here for security reasons,
	// because Milo will hand out its oauth2 token to a swarming host.
	AllowedHosts []string `protobuf:"bytes,2,rep,name=allowed_hosts,json=allowedHosts,proto3" json:"allowed_hosts,omitempty"`
}

func (x *Settings_Swarming) Reset() {
	*x = Settings_Swarming{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_milo_api_config_settings_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Settings_Swarming) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Settings_Swarming) ProtoMessage() {}

func (x *Settings_Swarming) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_milo_api_config_settings_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Settings_Swarming.ProtoReflect.Descriptor instead.
func (*Settings_Swarming) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_milo_api_config_settings_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Settings_Swarming) GetDefaultHost() string {
	if x != nil {
		return x.DefaultHost
	}
	return ""
}

func (x *Settings_Swarming) GetAllowedHosts() []string {
	if x != nil {
		return x.AllowedHosts
	}
	return nil
}

// SourceAcls grants read access on a set of Git/Gerrit hosts or projects.
type Settings_SourceAcls struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// host grants read access on all project at this host.
	//
	// For more granularity, use the project field instead.
	//
	// For *.googlesource.com domains, host should not be a Gerrit host,
	// i.e.  it shouldn't be <subdomain>-review.googlesource.com.
	Hosts []string `protobuf:"bytes,1,rep,name=hosts,proto3" json:"hosts,omitempty"`
	// project is a URL to a Git repository.
	//
	// Read access is granted on both git data and Gerrit CLs of this project.
	//
	// For *.googlesource.com Git repositories:
	//   URL Path should not start with '/a/' (forced authentication).
	//   URL Path should not end with '.git' (redundant).
	Projects []string `protobuf:"bytes,2,rep,name=projects,proto3" json:"projects,omitempty"`
	// readers are allowed to read git/gerrit data from targets.
	//
	// Three types of identity strings are supported:
	//  * Emails.                   For example: "someuser@example.com"
	//  * Chrome-infra-auth Groups. For example: "group:committers"
	//  * Auth service identities.  For example: "kind:name"
	//
	// Required.
	Readers []string `protobuf:"bytes,3,rep,name=readers,proto3" json:"readers,omitempty"`
}

func (x *Settings_SourceAcls) Reset() {
	*x = Settings_SourceAcls{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_milo_api_config_settings_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Settings_SourceAcls) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Settings_SourceAcls) ProtoMessage() {}

func (x *Settings_SourceAcls) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_milo_api_config_settings_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Settings_SourceAcls.ProtoReflect.Descriptor instead.
func (*Settings_SourceAcls) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_milo_api_config_settings_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Settings_SourceAcls) GetHosts() []string {
	if x != nil {
		return x.Hosts
	}
	return nil
}

func (x *Settings_SourceAcls) GetProjects() []string {
	if x != nil {
		return x.Projects
	}
	return nil
}

func (x *Settings_SourceAcls) GetReaders() []string {
	if x != nil {
		return x.Readers
	}
	return nil
}

type Settings_ResultDB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// host is the hostname of the ResultDB instance we're pointing to (sans scheme).
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
}

func (x *Settings_ResultDB) Reset() {
	*x = Settings_ResultDB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_milo_api_config_settings_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Settings_ResultDB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Settings_ResultDB) ProtoMessage() {}

func (x *Settings_ResultDB) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_milo_api_config_settings_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Settings_ResultDB.ProtoReflect.Descriptor instead.
func (*Settings_ResultDB) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_milo_api_config_settings_proto_rawDescGZIP(), []int{0, 3}
}

func (x *Settings_ResultDB) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

type Settings_LuciAnalysis struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// host is the hostname of the LUCI Analysis instance we're pointing to (sans scheme).
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
}

func (x *Settings_LuciAnalysis) Reset() {
	*x = Settings_LuciAnalysis{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_milo_api_config_settings_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Settings_LuciAnalysis) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Settings_LuciAnalysis) ProtoMessage() {}

func (x *Settings_LuciAnalysis) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_milo_api_config_settings_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Settings_LuciAnalysis.ProtoReflect.Descriptor instead.
func (*Settings_LuciAnalysis) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_milo_api_config_settings_proto_rawDescGZIP(), []int{0, 4}
}

func (x *Settings_LuciAnalysis) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

var File_go_chromium_org_luci_milo_api_config_settings_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_milo_api_config_settings_proto_rawDesc = []byte{
	0x0a, 0x33, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x6d, 0x69, 0x6c, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x69, 0x6c, 0x6f, 0x22, 0xff, 0x04, 0x0a, 0x08,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x3c, 0x0a, 0x0b, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x6d, 0x69, 0x6c, 0x6f, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x42, 0x75,
	0x69, 0x6c, 0x64, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x0b, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x33, 0x0a, 0x08, 0x73, 0x77, 0x61, 0x72, 0x6d, 0x69,
	0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x69, 0x6c, 0x6f, 0x2e,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x53, 0x77, 0x61, 0x72, 0x6d, 0x69, 0x6e,
	0x67, 0x52, 0x08, 0x73, 0x77, 0x61, 0x72, 0x6d, 0x69, 0x6e, 0x67, 0x12, 0x3a, 0x0a, 0x0b, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x6d, 0x69, 0x6c, 0x6f, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x63, 0x6c, 0x73, 0x52, 0x0a, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x41, 0x63, 0x6c, 0x73, 0x12, 0x33, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x64, 0x62, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x69, 0x6c, 0x6f,
	0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x44, 0x42, 0x52, 0x08, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x64, 0x62, 0x12, 0x40, 0x0a, 0x0d,
	0x6c, 0x75, 0x63, 0x69, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6d, 0x69, 0x6c, 0x6f, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x2e, 0x4c, 0x75, 0x63, 0x69, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73,
	0x52, 0x0c, 0x6c, 0x75, 0x63, 0x69, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x1a, 0x4f,
	0x0a, 0x0b, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x1a,
	0x52, 0x0a, 0x08, 0x53, 0x77, 0x61, 0x72, 0x6d, 0x69, 0x6e, 0x67, 0x12, 0x21, 0x0a, 0x0c, 0x64,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x23,
	0x0a, 0x0d, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x48, 0x6f,
	0x73, 0x74, 0x73, 0x1a, 0x58, 0x0a, 0x0a, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x63, 0x6c,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x05, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x1a, 0x1e, 0x0a,
	0x08, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x44, 0x42, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x1a, 0x22, 0x0a,
	0x0c, 0x4c, 0x75, 0x63, 0x69, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73,
	0x74, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x4a, 0x04, 0x08, 0x06, 0x10, 0x07, 0x42, 0x26, 0x5a,
	0x24, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x6d, 0x69, 0x6c, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_milo_api_config_settings_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_milo_api_config_settings_proto_rawDescData = file_go_chromium_org_luci_milo_api_config_settings_proto_rawDesc
)

func file_go_chromium_org_luci_milo_api_config_settings_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_milo_api_config_settings_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_milo_api_config_settings_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_milo_api_config_settings_proto_rawDescData)
	})
	return file_go_chromium_org_luci_milo_api_config_settings_proto_rawDescData
}

var file_go_chromium_org_luci_milo_api_config_settings_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_go_chromium_org_luci_milo_api_config_settings_proto_goTypes = []interface{}{
	(*Settings)(nil),              // 0: milo.Settings
	(*Settings_Buildbucket)(nil),  // 1: milo.Settings.Buildbucket
	(*Settings_Swarming)(nil),     // 2: milo.Settings.Swarming
	(*Settings_SourceAcls)(nil),   // 3: milo.Settings.SourceAcls
	(*Settings_ResultDB)(nil),     // 4: milo.Settings.ResultDB
	(*Settings_LuciAnalysis)(nil), // 5: milo.Settings.LuciAnalysis
}
var file_go_chromium_org_luci_milo_api_config_settings_proto_depIdxs = []int32{
	1, // 0: milo.Settings.buildbucket:type_name -> milo.Settings.Buildbucket
	2, // 1: milo.Settings.swarming:type_name -> milo.Settings.Swarming
	3, // 2: milo.Settings.source_acls:type_name -> milo.Settings.SourceAcls
	4, // 3: milo.Settings.resultdb:type_name -> milo.Settings.ResultDB
	5, // 4: milo.Settings.luci_analysis:type_name -> milo.Settings.LuciAnalysis
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_milo_api_config_settings_proto_init() }
func file_go_chromium_org_luci_milo_api_config_settings_proto_init() {
	if File_go_chromium_org_luci_milo_api_config_settings_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_milo_api_config_settings_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Settings); i {
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
		file_go_chromium_org_luci_milo_api_config_settings_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Settings_Buildbucket); i {
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
		file_go_chromium_org_luci_milo_api_config_settings_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Settings_Swarming); i {
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
		file_go_chromium_org_luci_milo_api_config_settings_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Settings_SourceAcls); i {
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
		file_go_chromium_org_luci_milo_api_config_settings_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Settings_ResultDB); i {
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
		file_go_chromium_org_luci_milo_api_config_settings_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Settings_LuciAnalysis); i {
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
			RawDescriptor: file_go_chromium_org_luci_milo_api_config_settings_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_milo_api_config_settings_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_milo_api_config_settings_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_milo_api_config_settings_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_milo_api_config_settings_proto = out.File
	file_go_chromium_org_luci_milo_api_config_settings_proto_rawDesc = nil
	file_go_chromium_org_luci_milo_api_config_settings_proto_goTypes = nil
	file_go_chromium_org_luci_milo_api_config_settings_proto_depIdxs = nil
}
