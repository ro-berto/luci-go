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
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.0
// source: go.chromium.org/luci/cv/api/config/legacy/tricium.proto

package tricium

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

// All supported data types.
type DataType int32

const (
	DataType_NONE             DataType = 0
	DataType_GIT_FILE_DETAILS DataType = 1
	DataType_RESULTS          DataType = 2
)

// Enum value maps for DataType.
var (
	DataType_name = map[int32]string{
		0: "NONE",
		1: "GIT_FILE_DETAILS",
		2: "RESULTS",
	}
	DataType_value = map[string]int32{
		"NONE":             0,
		"GIT_FILE_DETAILS": 1,
		"RESULTS":          2,
	}
)

func (x DataType) Enum() *DataType {
	p := new(DataType)
	*p = x
	return p
}

func (x DataType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DataType) Descriptor() protoreflect.EnumDescriptor {
	return file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_enumTypes[0].Descriptor()
}

func (DataType) Type() protoreflect.EnumType {
	return &file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_enumTypes[0]
}

func (x DataType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DataType.Descriptor instead.
func (DataType) EnumDescriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_rawDescGZIP(), []int{0}
}

type Platform int32

const (
	Platform_UNSPECIFIED Platform = 0
	Platform_LINUX       Platform = 1
)

// Enum value maps for Platform.
var (
	Platform_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "LINUX",
	}
	Platform_value = map[string]int32{
		"UNSPECIFIED": 0,
		"LINUX":       1,
	}
)

func (x Platform) Enum() *Platform {
	p := new(Platform)
	*p = x
	return p
}

func (x Platform) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Platform) Descriptor() protoreflect.EnumDescriptor {
	return file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_enumTypes[1].Descriptor()
}

func (Platform) Type() protoreflect.EnumType {
	return &file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_enumTypes[1]
}

func (x Platform) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Platform.Descriptor instead.
func (Platform) EnumDescriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_rawDescGZIP(), []int{1}
}

type Function_Type int32

const (
	Function_NONE     Function_Type = 0
	Function_ANALYZER Function_Type = 1
)

// Enum value maps for Function_Type.
var (
	Function_Type_name = map[int32]string{
		0: "NONE",
		1: "ANALYZER",
	}
	Function_Type_value = map[string]int32{
		"NONE":     0,
		"ANALYZER": 1,
	}
)

func (x Function_Type) Enum() *Function_Type {
	p := new(Function_Type)
	*p = x
	return p
}

func (x Function_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Function_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_enumTypes[2].Descriptor()
}

func (Function_Type) Type() protoreflect.EnumType {
	return &file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_enumTypes[2]
}

func (x Function_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Function_Type.Descriptor instead.
func (Function_Type) EnumDescriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_rawDescGZIP(), []int{1, 0}
}

// Tricium project configuration.
//
// Specifies details needed to connect a project to Tricium.
// Adds project-specific functions and selects shared function
// implementations.
type ProjectConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Analyzer definitions.
	//
	// Each analyzer generally corresponds to one builder.
	Functions []*Function `protobuf:"bytes,1,rep,name=functions,proto3" json:"functions,omitempty"`
	// Selection of function implementations to run for this project.
	//
	// An analyzer is only enabled if there is a selections entry. Generally all
	// defined functions are listed as selections. Note that the function
	// (analyzer) name must match.
	Selections []*Selection `protobuf:"bytes,2,rep,name=selections,proto3" json:"selections,omitempty"`
	// Repositories, including Gerrit details.
	Repos []*RepoDetails `protobuf:"bytes,3,rep,name=repos,proto3" json:"repos,omitempty"`
	// General service account for this project.
	ServiceAccount string `protobuf:"bytes,4,opt,name=service_account,json=serviceAccount,proto3" json:"service_account,omitempty"`
}

func (x *ProjectConfig) Reset() {
	*x = ProjectConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectConfig) ProtoMessage() {}

func (x *ProjectConfig) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectConfig.ProtoReflect.Descriptor instead.
func (*ProjectConfig) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_rawDescGZIP(), []int{0}
}

func (x *ProjectConfig) GetFunctions() []*Function {
	if x != nil {
		return x.Functions
	}
	return nil
}

func (x *ProjectConfig) GetSelections() []*Selection {
	if x != nil {
		return x.Selections
	}
	return nil
}

func (x *ProjectConfig) GetRepos() []*RepoDetails {
	if x != nil {
		return x.Repos
	}
	return nil
}

func (x *ProjectConfig) GetServiceAccount() string {
	if x != nil {
		return x.ServiceAccount
	}
	return ""
}

// Tricium analyzer definition.
type Function struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of this function. Required.
	//
	// Should always be ANALYZER.
	Type Function_Type `protobuf:"varint,1,opt,name=type,proto3,enum=cv.legacy.tricium.Function_Type" json:"type,omitempty"`
	// The name of the analyzer. Required.
	//
	// The name must be unique among Tricium functions within a Tricium instance.
	// The name is expected to be CamelCase; no spaces, underscores or dashes are
	// allowed.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Data needed by this analyzer. Required.
	//
	// Should always be GIT_FILE_DETAILS.
	Needs DataType `protobuf:"varint,3,opt,name=needs,proto3,enum=cv.legacy.tricium.DataType" json:"needs,omitempty"`
	// Data provided by this analyzer. Required.
	//
	// Should always be RESULTS.
	Provides DataType `protobuf:"varint,4,opt,name=provides,proto3,enum=cv.legacy.tricium.DataType" json:"provides,omitempty"`
	// Path filters for this analyzer.
	//
	// Defined as a glob. The path filters only apply to the last part of the
	// path.
	PathFilters []string `protobuf:"bytes,5,rep,name=path_filters,json=pathFilters,proto3" json:"path_filters,omitempty"` // Default: "*"
	// Function implementations.
	//
	// Originally the idea was that an analyzer may run on many different
	// platforms and the comments from different platforms may be merged.
	//
	// This was not done in practice, so the number of impls should always be one.
	Impls []*Impl `protobuf:"bytes,6,rep,name=impls,proto3" json:"impls,omitempty"`
}

func (x *Function) Reset() {
	*x = Function{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Function) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Function) ProtoMessage() {}

func (x *Function) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Function.ProtoReflect.Descriptor instead.
func (*Function) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_rawDescGZIP(), []int{1}
}

func (x *Function) GetType() Function_Type {
	if x != nil {
		return x.Type
	}
	return Function_NONE
}

func (x *Function) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Function) GetNeeds() DataType {
	if x != nil {
		return x.Needs
	}
	return DataType_NONE
}

func (x *Function) GetProvides() DataType {
	if x != nil {
		return x.Provides
	}
	return DataType_NONE
}

func (x *Function) GetPathFilters() []string {
	if x != nil {
		return x.PathFilters
	}
	return nil
}

func (x *Function) GetImpls() []*Impl {
	if x != nil {
		return x.Impls
	}
	return nil
}

// Analyzer implementation.
type Impl struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProvidesForPlatform Platform `protobuf:"varint,1,opt,name=provides_for_platform,json=providesForPlatform,proto3,enum=cv.legacy.tricium.Platform" json:"provides_for_platform,omitempty"`
	// The platform to run this implementation on.
	//
	// This particular value of this field isn't significant, because
	// the platform is determined by the builder.
	RuntimePlatform Platform `protobuf:"varint,2,opt,name=runtime_platform,json=runtimePlatform,proto3,enum=cv.legacy.tricium.Platform" json:"runtime_platform,omitempty"`
	// Recipe for recipe-based implementation.
	Recipe *Recipe `protobuf:"bytes,3,opt,name=recipe,proto3" json:"recipe,omitempty"`
}

func (x *Impl) Reset() {
	*x = Impl{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Impl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Impl) ProtoMessage() {}

func (x *Impl) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Impl.ProtoReflect.Descriptor instead.
func (*Impl) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_rawDescGZIP(), []int{2}
}

func (x *Impl) GetProvidesForPlatform() Platform {
	if x != nil {
		return x.ProvidesForPlatform
	}
	return Platform_UNSPECIFIED
}

func (x *Impl) GetRuntimePlatform() Platform {
	if x != nil {
		return x.RuntimePlatform
	}
	return Platform_UNSPECIFIED
}

func (x *Impl) GetRecipe() *Recipe {
	if x != nil {
		return x.Recipe
	}
	return nil
}

// Specification of a recipe for a recipe-based analyzer.
type Recipe struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Project ID, e.g. "chromium".
	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	// Bucket name, e.g. "try".
	Bucket string `protobuf:"bytes,2,opt,name=bucket,proto3" json:"bucket,omitempty"`
	// Builder name, e.g. "linux-rel".
	Builder string `protobuf:"bytes,3,opt,name=builder,proto3" json:"builder,omitempty"`
}

func (x *Recipe) Reset() {
	*x = Recipe{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Recipe) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Recipe) ProtoMessage() {}

func (x *Recipe) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Recipe.ProtoReflect.Descriptor instead.
func (*Recipe) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_rawDescGZIP(), []int{3}
}

func (x *Recipe) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *Recipe) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *Recipe) GetBuilder() string {
	if x != nil {
		return x.Builder
	}
	return ""
}

// Selection of function implementations to run for a project.
type Selection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of function to run.
	Function string `protobuf:"bytes,1,opt,name=function,proto3" json:"function,omitempty"`
	// Name of platform to retrieve results from.
	Platform Platform `protobuf:"varint,2,opt,name=platform,proto3,enum=cv.legacy.tricium.Platform" json:"platform,omitempty"`
}

func (x *Selection) Reset() {
	*x = Selection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Selection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Selection) ProtoMessage() {}

func (x *Selection) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Selection.ProtoReflect.Descriptor instead.
func (*Selection) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_rawDescGZIP(), []int{4}
}

func (x *Selection) GetFunction() string {
	if x != nil {
		return x.Function
	}
	return ""
}

func (x *Selection) GetPlatform() Platform {
	if x != nil {
		return x.Platform
	}
	return Platform_UNSPECIFIED
}

// Repository details for one repository.
type RepoDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GerritProject *RepoDetails_GerritProject `protobuf:"bytes,1,opt,name=gerrit_project,json=gerritProject,proto3" json:"gerrit_project,omitempty"`
	// Whitelisted groups.
	//
	// The owner of a change will be checked for membership of a whitelisted
	// group. Absence of this field means all groups are whitelisted.
	//
	// Group names must be known to the Chrome infra auth service,
	// https://chrome-infra-auth.appspot.com. Contact a Chromium trooper
	// if you need to add or modify a group: g.co/bugatrooper.
	WhitelistedGroup []string `protobuf:"bytes,7,rep,name=whitelisted_group,json=whitelistedGroup,proto3" json:"whitelisted_group,omitempty"`
}

func (x *RepoDetails) Reset() {
	*x = RepoDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepoDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepoDetails) ProtoMessage() {}

func (x *RepoDetails) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepoDetails.ProtoReflect.Descriptor instead.
func (*RepoDetails) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_rawDescGZIP(), []int{5}
}

func (x *RepoDetails) GetGerritProject() *RepoDetails_GerritProject {
	if x != nil {
		return x.GerritProject
	}
	return nil
}

func (x *RepoDetails) GetWhitelistedGroup() []string {
	if x != nil {
		return x.WhitelistedGroup
	}
	return nil
}

// Specifies a Gerrit project and its corresponding git repo.
type RepoDetails_GerritProject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Gerrit host to connect to.
	//
	// Value must not include the schema part; it will be assumed to be "https".
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	// Gerrit project name.
	Project string `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
	// Full URL for the corresponding git repo.
	GitUrl string `protobuf:"bytes,3,opt,name=git_url,json=gitUrl,proto3" json:"git_url,omitempty"`
}

func (x *RepoDetails_GerritProject) Reset() {
	*x = RepoDetails_GerritProject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepoDetails_GerritProject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepoDetails_GerritProject) ProtoMessage() {}

func (x *RepoDetails_GerritProject) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepoDetails_GerritProject.ProtoReflect.Descriptor instead.
func (*RepoDetails_GerritProject) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_rawDescGZIP(), []int{5, 0}
}

func (x *RepoDetails_GerritProject) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *RepoDetails_GerritProject) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *RepoDetails_GerritProject) GetGitUrl() string {
	if x != nil {
		return x.GitUrl
	}
	return ""
}

var File_go_chromium_org_luci_cv_api_config_legacy_tricium_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_rawDesc = []byte{
	0x0a, 0x37, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x76, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x2f, 0x74, 0x72, 0x69, 0x63,
	0x69, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x63, 0x76, 0x2e, 0x6c, 0x65,
	0x67, 0x61, 0x63, 0x79, 0x2e, 0x74, 0x72, 0x69, 0x63, 0x69, 0x75, 0x6d, 0x22, 0xe7, 0x01, 0x0a,
	0x0d, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x39,
	0x0a, 0x09, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x76, 0x2e, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x2e, 0x74, 0x72,
	0x69, 0x63, 0x69, 0x75, 0x6d, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09,
	0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3c, 0x0a, 0x0a, 0x73, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x63, 0x76, 0x2e, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x2e, 0x74, 0x72, 0x69, 0x63, 0x69, 0x75,
	0x6d, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x73, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x34, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6f, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x76, 0x2e, 0x6c, 0x65, 0x67, 0x61,
	0x63, 0x79, 0x2e, 0x74, 0x72, 0x69, 0x63, 0x69, 0x75, 0x6d, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x05, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x12, 0x27, 0x0a,
	0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xb2, 0x02, 0x0a, 0x08, 0x46, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x20, 0x2e, 0x63, 0x76, 0x2e, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x2e, 0x74, 0x72,
	0x69, 0x63, 0x69, 0x75, 0x6d, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x31, 0x0a,
	0x05, 0x6e, 0x65, 0x65, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x63,
	0x76, 0x2e, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x2e, 0x74, 0x72, 0x69, 0x63, 0x69, 0x75, 0x6d,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x6e, 0x65, 0x65, 0x64, 0x73,
	0x12, 0x37, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x63, 0x76, 0x2e, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x2e, 0x74,
	0x72, 0x69, 0x63, 0x69, 0x75, 0x6d, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x74,
	0x68, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0b, 0x70, 0x61, 0x74, 0x68, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x2d, 0x0a, 0x05,
	0x69, 0x6d, 0x70, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x76,
	0x2e, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x2e, 0x74, 0x72, 0x69, 0x63, 0x69, 0x75, 0x6d, 0x2e,
	0x49, 0x6d, 0x70, 0x6c, 0x52, 0x05, 0x69, 0x6d, 0x70, 0x6c, 0x73, 0x22, 0x1e, 0x0a, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0c, 0x0a,
	0x08, 0x41, 0x4e, 0x41, 0x4c, 0x59, 0x5a, 0x45, 0x52, 0x10, 0x01, 0x22, 0xd2, 0x01, 0x0a, 0x04,
	0x49, 0x6d, 0x70, 0x6c, 0x12, 0x4f, 0x0a, 0x15, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x73,
	0x5f, 0x66, 0x6f, 0x72, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x63, 0x76, 0x2e, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x2e,
	0x74, 0x72, 0x69, 0x63, 0x69, 0x75, 0x6d, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x52, 0x13, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x73, 0x46, 0x6f, 0x72, 0x50, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x46, 0x0a, 0x10, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1b, 0x2e, 0x63, 0x76, 0x2e, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x2e, 0x74, 0x72, 0x69, 0x63,
	0x69, 0x75, 0x6d, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x0f, 0x72, 0x75,
	0x6e, 0x74, 0x69, 0x6d, 0x65, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x31, 0x0a,
	0x06, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x63, 0x76, 0x2e, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x2e, 0x74, 0x72, 0x69, 0x63, 0x69, 0x75,
	0x6d, 0x2e, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x52, 0x06, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65,
	0x22, 0x54, 0x0a, 0x06, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x22, 0x60, 0x0a, 0x09, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x37, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1b, 0x2e, 0x63, 0x76, 0x2e, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x2e, 0x74, 0x72,
	0x69, 0x63, 0x69, 0x75, 0x6d, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x08,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x22, 0xe7, 0x01, 0x0a, 0x0b, 0x52, 0x65, 0x70,
	0x6f, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x53, 0x0a, 0x0e, 0x67, 0x65, 0x72, 0x72,
	0x69, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2c, 0x2e, 0x63, 0x76, 0x2e, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x2e, 0x74, 0x72, 0x69,
	0x63, 0x69, 0x75, 0x6d, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x2e, 0x47, 0x65, 0x72, 0x72, 0x69, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x0d,
	0x67, 0x65, 0x72, 0x72, 0x69, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x2b, 0x0a,
	0x11, 0x77, 0x68, 0x69, 0x74, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x77, 0x68, 0x69, 0x74, 0x65, 0x6c,
	0x69, 0x73, 0x74, 0x65, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x1a, 0x56, 0x0a, 0x0d, 0x47, 0x65,
	0x72, 0x72, 0x69, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68,
	0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x69, 0x74,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x69, 0x74, 0x55,
	0x72, 0x6c, 0x2a, 0x37, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08,
	0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x47, 0x49, 0x54, 0x5f,
	0x46, 0x49, 0x4c, 0x45, 0x5f, 0x44, 0x45, 0x54, 0x41, 0x49, 0x4c, 0x53, 0x10, 0x01, 0x12, 0x0b,
	0x0a, 0x07, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x53, 0x10, 0x02, 0x2a, 0x26, 0x0a, 0x08, 0x50,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x49, 0x4e, 0x55,
	0x58, 0x10, 0x01, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69,
	0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x76, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79,
	0x3b, 0x74, 0x72, 0x69, 0x63, 0x69, 0x75, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_rawDescData = file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_rawDesc
)

func file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_rawDescData)
	})
	return file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_rawDescData
}

var file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_goTypes = []interface{}{
	(DataType)(0),                     // 0: cv.legacy.tricium.DataType
	(Platform)(0),                     // 1: cv.legacy.tricium.Platform
	(Function_Type)(0),                // 2: cv.legacy.tricium.Function.Type
	(*ProjectConfig)(nil),             // 3: cv.legacy.tricium.ProjectConfig
	(*Function)(nil),                  // 4: cv.legacy.tricium.Function
	(*Impl)(nil),                      // 5: cv.legacy.tricium.Impl
	(*Recipe)(nil),                    // 6: cv.legacy.tricium.Recipe
	(*Selection)(nil),                 // 7: cv.legacy.tricium.Selection
	(*RepoDetails)(nil),               // 8: cv.legacy.tricium.RepoDetails
	(*RepoDetails_GerritProject)(nil), // 9: cv.legacy.tricium.RepoDetails.GerritProject
}
var file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_depIdxs = []int32{
	4,  // 0: cv.legacy.tricium.ProjectConfig.functions:type_name -> cv.legacy.tricium.Function
	7,  // 1: cv.legacy.tricium.ProjectConfig.selections:type_name -> cv.legacy.tricium.Selection
	8,  // 2: cv.legacy.tricium.ProjectConfig.repos:type_name -> cv.legacy.tricium.RepoDetails
	2,  // 3: cv.legacy.tricium.Function.type:type_name -> cv.legacy.tricium.Function.Type
	0,  // 4: cv.legacy.tricium.Function.needs:type_name -> cv.legacy.tricium.DataType
	0,  // 5: cv.legacy.tricium.Function.provides:type_name -> cv.legacy.tricium.DataType
	5,  // 6: cv.legacy.tricium.Function.impls:type_name -> cv.legacy.tricium.Impl
	1,  // 7: cv.legacy.tricium.Impl.provides_for_platform:type_name -> cv.legacy.tricium.Platform
	1,  // 8: cv.legacy.tricium.Impl.runtime_platform:type_name -> cv.legacy.tricium.Platform
	6,  // 9: cv.legacy.tricium.Impl.recipe:type_name -> cv.legacy.tricium.Recipe
	1,  // 10: cv.legacy.tricium.Selection.platform:type_name -> cv.legacy.tricium.Platform
	9,  // 11: cv.legacy.tricium.RepoDetails.gerrit_project:type_name -> cv.legacy.tricium.RepoDetails.GerritProject
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_init() }
func file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_init() {
	if File_go_chromium_org_luci_cv_api_config_legacy_tricium_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectConfig); i {
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
		file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Function); i {
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
		file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Impl); i {
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
		file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Recipe); i {
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
		file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Selection); i {
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
		file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepoDetails); i {
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
		file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepoDetails_GerritProject); i {
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
			RawDescriptor: file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_depIdxs,
		EnumInfos:         file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_enumTypes,
		MessageInfos:      file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_cv_api_config_legacy_tricium_proto = out.File
	file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_rawDesc = nil
	file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_goTypes = nil
	file_go_chromium_org_luci_cv_api_config_legacy_tricium_proto_depIdxs = nil
}
