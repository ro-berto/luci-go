// Copyright 2018 The LUCI Authors.
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

// Schemas for service configs.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: go.chromium.org/luci/buildbucket/proto/service_config.proto

package buildbucketpb

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

// Schema of settings.cfg file, a service config.
type SettingsCfg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Swarmbucket settings.
	Swarming   *SwarmingSettings   `protobuf:"bytes,1,opt,name=swarming,proto3" json:"swarming,omitempty"`
	Logdog     *LogDogSettings     `protobuf:"bytes,2,opt,name=logdog,proto3" json:"logdog,omitempty"`
	Resultdb   *ResultDBSettings   `protobuf:"bytes,4,opt,name=resultdb,proto3" json:"resultdb,omitempty"`
	Experiment *ExperimentSettings `protobuf:"bytes,5,opt,name=experiment,proto3" json:"experiment,omitempty"`
	// List of Gerrit hosts to force git authentication for.
	//
	// By default public hosts are accessed anonymously, and the anonymous access
	// has very low quota. Context needs to know all such hostnames in advance to
	// be able to force authenticated access to them.
	KnownPublicGerritHosts []string `protobuf:"bytes,3,rep,name=known_public_gerrit_hosts,json=knownPublicGerritHosts,proto3" json:"known_public_gerrit_hosts,omitempty"`
}

func (x *SettingsCfg) Reset() {
	*x = SettingsCfg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_buildbucket_proto_service_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SettingsCfg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SettingsCfg) ProtoMessage() {}

func (x *SettingsCfg) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_buildbucket_proto_service_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SettingsCfg.ProtoReflect.Descriptor instead.
func (*SettingsCfg) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_buildbucket_proto_service_config_proto_rawDescGZIP(), []int{0}
}

func (x *SettingsCfg) GetSwarming() *SwarmingSettings {
	if x != nil {
		return x.Swarming
	}
	return nil
}

func (x *SettingsCfg) GetLogdog() *LogDogSettings {
	if x != nil {
		return x.Logdog
	}
	return nil
}

func (x *SettingsCfg) GetResultdb() *ResultDBSettings {
	if x != nil {
		return x.Resultdb
	}
	return nil
}

func (x *SettingsCfg) GetExperiment() *ExperimentSettings {
	if x != nil {
		return x.Experiment
	}
	return nil
}

func (x *SettingsCfg) GetKnownPublicGerritHosts() []string {
	if x != nil {
		return x.KnownPublicGerritHosts
	}
	return nil
}

// Swarmbucket settings.
type SwarmingSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Swarmbucket build URLs will point to this Milo instance.
	MiloHostname string `protobuf:"bytes,2,opt,name=milo_hostname,json=miloHostname,proto3" json:"milo_hostname,omitempty"`
	// These caches are available to all builders implicitly.
	// A builder may override a cache specified here.
	GlobalCaches []*Builder_CacheEntry `protobuf:"bytes,4,rep,name=global_caches,json=globalCaches,proto3" json:"global_caches,omitempty"`
	// Packages available to the user executable in $PATH.
	// Installed in "{TASK_RUN_DIR}/cipd_bin_packages".
	// "{TASK_RUN_DIR}/cipd_bin_packages" and
	// "{TASK_RUN_DIR}/cipd_bin_packages/bin" are prepended to $PATH.
	UserPackages []*SwarmingSettings_Package `protobuf:"bytes,5,rep,name=user_packages,json=userPackages,proto3" json:"user_packages,omitempty"`
	// Package of buildbucket agent,
	// https://chromium.googlesource.com/infra/luci/luci-go/+/HEAD/buildbucket/cmd/bbagent
	// used to run LUCI executables.
	BbagentPackage *SwarmingSettings_Package `protobuf:"bytes,8,opt,name=bbagent_package,json=bbagentPackage,proto3" json:"bbagent_package,omitempty"`
	// CIPD package of kitchen binary. DEPRECATED. TODO(nodir): remove.
	KitchenPackage *SwarmingSettings_Package `protobuf:"bytes,7,opt,name=kitchen_package,json=kitchenPackage,proto3" json:"kitchen_package,omitempty"`
}

func (x *SwarmingSettings) Reset() {
	*x = SwarmingSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_buildbucket_proto_service_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SwarmingSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SwarmingSettings) ProtoMessage() {}

func (x *SwarmingSettings) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_buildbucket_proto_service_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SwarmingSettings.ProtoReflect.Descriptor instead.
func (*SwarmingSettings) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_buildbucket_proto_service_config_proto_rawDescGZIP(), []int{1}
}

func (x *SwarmingSettings) GetMiloHostname() string {
	if x != nil {
		return x.MiloHostname
	}
	return ""
}

func (x *SwarmingSettings) GetGlobalCaches() []*Builder_CacheEntry {
	if x != nil {
		return x.GlobalCaches
	}
	return nil
}

func (x *SwarmingSettings) GetUserPackages() []*SwarmingSettings_Package {
	if x != nil {
		return x.UserPackages
	}
	return nil
}

func (x *SwarmingSettings) GetBbagentPackage() *SwarmingSettings_Package {
	if x != nil {
		return x.BbagentPackage
	}
	return nil
}

func (x *SwarmingSettings) GetKitchenPackage() *SwarmingSettings_Package {
	if x != nil {
		return x.KitchenPackage
	}
	return nil
}

type LogDogSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Hostname of the LogDog instance to use, e.g. "logs.chromium.org".
	Hostname string `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
}

func (x *LogDogSettings) Reset() {
	*x = LogDogSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_buildbucket_proto_service_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogDogSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogDogSettings) ProtoMessage() {}

func (x *LogDogSettings) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_buildbucket_proto_service_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogDogSettings.ProtoReflect.Descriptor instead.
func (*LogDogSettings) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_buildbucket_proto_service_config_proto_rawDescGZIP(), []int{2}
}

func (x *LogDogSettings) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

// ExperimentSettings controls all well-known global experiment values.
type ExperimentSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Experiments []*ExperimentSettings_Experiment `protobuf:"bytes,1,rep,name=experiments,proto3" json:"experiments,omitempty"`
}

func (x *ExperimentSettings) Reset() {
	*x = ExperimentSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_buildbucket_proto_service_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExperimentSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExperimentSettings) ProtoMessage() {}

func (x *ExperimentSettings) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_buildbucket_proto_service_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExperimentSettings.ProtoReflect.Descriptor instead.
func (*ExperimentSettings) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_buildbucket_proto_service_config_proto_rawDescGZIP(), []int{3}
}

func (x *ExperimentSettings) GetExperiments() []*ExperimentSettings_Experiment {
	if x != nil {
		return x.Experiments
	}
	return nil
}

// A predicate for a builder.
type BuilderPredicate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// OR-connected list of regular expressions for a string
	// "{project}/{bucket}/{builder}".
	// Each regex is wrapped in ^ and $ automatically.
	// Examples:
	//
	//   # All builders in "chromium" project
	//   regex: "chromium/.+"
	//   # A specific builder.
	//   regex: "infra/ci/infra-continuous-trusty-64"
	//
	// Defaults to [".*"].
	Regex []string `protobuf:"bytes,1,rep,name=regex,proto3" json:"regex,omitempty"`
	// Like regex field, but negation. Negation always wins.
	RegexExclude []string `protobuf:"bytes,2,rep,name=regex_exclude,json=regexExclude,proto3" json:"regex_exclude,omitempty"`
}

func (x *BuilderPredicate) Reset() {
	*x = BuilderPredicate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_buildbucket_proto_service_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuilderPredicate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuilderPredicate) ProtoMessage() {}

func (x *BuilderPredicate) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_buildbucket_proto_service_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuilderPredicate.ProtoReflect.Descriptor instead.
func (*BuilderPredicate) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_buildbucket_proto_service_config_proto_rawDescGZIP(), []int{4}
}

func (x *BuilderPredicate) GetRegex() []string {
	if x != nil {
		return x.Regex
	}
	return nil
}

func (x *BuilderPredicate) GetRegexExclude() []string {
	if x != nil {
		return x.RegexExclude
	}
	return nil
}

type ResultDBSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Hostname of the ResultDB instance to use, e.g. "results.api.cr.dev".
	Hostname string `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
}

func (x *ResultDBSettings) Reset() {
	*x = ResultDBSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_buildbucket_proto_service_config_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultDBSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultDBSettings) ProtoMessage() {}

func (x *ResultDBSettings) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_buildbucket_proto_service_config_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultDBSettings.ProtoReflect.Descriptor instead.
func (*ResultDBSettings) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_buildbucket_proto_service_config_proto_rawDescGZIP(), []int{5}
}

func (x *ResultDBSettings) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

// CIPD package. Does not specify installation path.
type SwarmingSettings_Package struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// CIPD package name, e.g. "infra/python/cpython/${platform}"
	PackageName string `protobuf:"bytes,1,opt,name=package_name,json=packageName,proto3" json:"package_name,omitempty"`
	// CIPD instance version, e.g. "version:2.7.15.chromium14".
	// Used for non-canary builds.
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// CIPD instance version for canary builds.
	// Defaults to version.
	VersionCanary string `protobuf:"bytes,3,opt,name=version_canary,json=versionCanary,proto3" json:"version_canary,omitempty"`
	// Include in builders matching the predicate.
	Builders *BuilderPredicate `protobuf:"bytes,4,opt,name=builders,proto3" json:"builders,omitempty"`
	// Subdirectory to install the package into, relative to the installation
	// root. Useful if installing two packages at the same root would conflict.
	Subdir string `protobuf:"bytes,5,opt,name=subdir,proto3" json:"subdir,omitempty"`
}

func (x *SwarmingSettings_Package) Reset() {
	*x = SwarmingSettings_Package{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_buildbucket_proto_service_config_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SwarmingSettings_Package) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SwarmingSettings_Package) ProtoMessage() {}

func (x *SwarmingSettings_Package) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_buildbucket_proto_service_config_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SwarmingSettings_Package.ProtoReflect.Descriptor instead.
func (*SwarmingSettings_Package) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_buildbucket_proto_service_config_proto_rawDescGZIP(), []int{1, 0}
}

func (x *SwarmingSettings_Package) GetPackageName() string {
	if x != nil {
		return x.PackageName
	}
	return ""
}

func (x *SwarmingSettings_Package) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *SwarmingSettings_Package) GetVersionCanary() string {
	if x != nil {
		return x.VersionCanary
	}
	return ""
}

func (x *SwarmingSettings_Package) GetBuilders() *BuilderPredicate {
	if x != nil {
		return x.Builders
	}
	return nil
}

func (x *SwarmingSettings_Package) GetSubdir() string {
	if x != nil {
		return x.Subdir
	}
	return ""
}

type ExperimentSettings_Experiment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the global experiment.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The value (% chance, 0 - 100) of the global experiment.
	Value int32 `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	// Allows temporary exclusion of builders from the experiment.
	// Each line here should have a corresponding bug to remove the exclusion.
	Builders *BuilderPredicate `protobuf:"bytes,3,opt,name=builders,proto3" json:"builders,omitempty"`
	// If this is true it means that the experiment has no effect, and is safe
	// to stop setting in user configs. Additionally, Buildbucket will stop
	// setting this experiment negatively on Builds.
	//
	// When removing a global experiment, set this to true rather than removing
	// the experiment entirely, because Buildbucket still needs this to permit
	// (and ignore) user configs which still mention reserved experiments (e.g.
	// if we have "luci.something" which someone explicitly specifies, and we
	// ramp it to 100% and remove it from the global spec, Buildbucket will
	// start complaining that users are using a reserved experiment name, rather
	// than just ignoring it).
	//
	// If inactive experiments appear in user configurations, it may cause
	// warnings to be printed e.g. at config validation time and/or on the LUCI
	// UI, etc.
	Inactive bool `protobuf:"varint,4,opt,name=inactive,proto3" json:"inactive,omitempty"`
}

func (x *ExperimentSettings_Experiment) Reset() {
	*x = ExperimentSettings_Experiment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_buildbucket_proto_service_config_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExperimentSettings_Experiment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExperimentSettings_Experiment) ProtoMessage() {}

func (x *ExperimentSettings_Experiment) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_buildbucket_proto_service_config_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExperimentSettings_Experiment.ProtoReflect.Descriptor instead.
func (*ExperimentSettings_Experiment) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_buildbucket_proto_service_config_proto_rawDescGZIP(), []int{3, 0}
}

func (x *ExperimentSettings_Experiment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ExperimentSettings_Experiment) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *ExperimentSettings_Experiment) GetBuilders() *BuilderPredicate {
	if x != nil {
		return x.Builders
	}
	return nil
}

func (x *ExperimentSettings_Experiment) GetInactive() bool {
	if x != nil {
		return x.Inactive
	}
	return false
}

var File_go_chromium_org_luci_buildbucket_proto_service_config_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_buildbucket_proto_service_config_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x1a, 0x3b, 0x67, 0x6f, 0x2e, 0x63,
	0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69,
	0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb4, 0x02, 0x0a, 0x0b, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x43, 0x66, 0x67, 0x12, 0x39, 0x0a, 0x08, 0x73, 0x77, 0x61, 0x72, 0x6d,
	0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x53, 0x77, 0x61, 0x72, 0x6d, 0x69, 0x6e, 0x67,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x08, 0x73, 0x77, 0x61, 0x72, 0x6d, 0x69,
	0x6e, 0x67, 0x12, 0x33, 0x0a, 0x06, 0x6c, 0x6f, 0x67, 0x64, 0x6f, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x2e, 0x4c, 0x6f, 0x67, 0x44, 0x6f, 0x67, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52,
	0x06, 0x6c, 0x6f, 0x67, 0x64, 0x6f, 0x67, 0x12, 0x39, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x64, 0x62, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x44, 0x42,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x08, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x64, 0x62, 0x12, 0x3f, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x39, 0x0a, 0x19, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x5f, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x5f, 0x67, 0x65, 0x72, 0x72, 0x69, 0x74, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x16, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x47, 0x65, 0x72, 0x72, 0x69, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x22, 0xb8,
	0x04, 0x0a, 0x10, 0x53, 0x77, 0x61, 0x72, 0x6d, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x69, 0x6c, 0x6f, 0x5f, 0x68, 0x6f, 0x73, 0x74,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x69, 0x6c, 0x6f,
	0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x44, 0x0a, 0x0d, 0x67, 0x6c, 0x6f, 0x62,
	0x61, 0x6c, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x42, 0x75,
	0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0c, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x43, 0x61, 0x63, 0x68, 0x65, 0x73, 0x12, 0x4a,
	0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x2e, 0x53, 0x77, 0x61, 0x72, 0x6d, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x0c, 0x75, 0x73,
	0x65, 0x72, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x12, 0x4e, 0x0a, 0x0f, 0x62, 0x62,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x2e, 0x53, 0x77, 0x61, 0x72, 0x6d, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x0e, 0x62, 0x62, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x4e, 0x0a, 0x0f, 0x6b, 0x69,
	0x74, 0x63, 0x68, 0x65, 0x6e, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x2e, 0x53, 0x77, 0x61, 0x72, 0x6d, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x0e, 0x6b, 0x69, 0x74, 0x63,
	0x68, 0x65, 0x6e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x1a, 0xc0, 0x01, 0x0a, 0x07, 0x50,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x63,
	0x61, 0x6e, 0x61, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x12, 0x39, 0x0a, 0x08, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64,
	0x65, 0x72, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x08, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x65, 0x72, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x75, 0x62, 0x64, 0x69, 0x72, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x75, 0x62, 0x64, 0x69, 0x72, 0x4a, 0x04, 0x08,
	0x01, 0x10, 0x02, 0x4a, 0x04, 0x08, 0x06, 0x10, 0x07, 0x22, 0x2c, 0x0a, 0x0e, 0x4c, 0x6f, 0x67,
	0x44, 0x6f, 0x67, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x68,
	0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68,
	0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xf2, 0x01, 0x0a, 0x12, 0x45, 0x78, 0x70, 0x65,
	0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x4c,
	0x0a, 0x0b, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x0b, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x8d, 0x01, 0x0a,
	0x0a, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x39, 0x0a, 0x08, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x50, 0x72, 0x65,
	0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x08, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x69, 0x6e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x4d, 0x0a, 0x10,
	0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x67, 0x65, 0x78, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x05, 0x72, 0x65, 0x67, 0x65, 0x78, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x67, 0x65, 0x78, 0x5f,
	0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x72,
	0x65, 0x67, 0x65, 0x78, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x22, 0x2e, 0x0a, 0x10, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x44, 0x42, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x36, 0x5a, 0x34, 0x67,
	0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c,
	0x75, 0x63, 0x69, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_buildbucket_proto_service_config_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_buildbucket_proto_service_config_proto_rawDescData = file_go_chromium_org_luci_buildbucket_proto_service_config_proto_rawDesc
)

func file_go_chromium_org_luci_buildbucket_proto_service_config_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_buildbucket_proto_service_config_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_buildbucket_proto_service_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_buildbucket_proto_service_config_proto_rawDescData)
	})
	return file_go_chromium_org_luci_buildbucket_proto_service_config_proto_rawDescData
}

var file_go_chromium_org_luci_buildbucket_proto_service_config_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_go_chromium_org_luci_buildbucket_proto_service_config_proto_goTypes = []interface{}{
	(*SettingsCfg)(nil),                   // 0: buildbucket.SettingsCfg
	(*SwarmingSettings)(nil),              // 1: buildbucket.SwarmingSettings
	(*LogDogSettings)(nil),                // 2: buildbucket.LogDogSettings
	(*ExperimentSettings)(nil),            // 3: buildbucket.ExperimentSettings
	(*BuilderPredicate)(nil),              // 4: buildbucket.BuilderPredicate
	(*ResultDBSettings)(nil),              // 5: buildbucket.ResultDBSettings
	(*SwarmingSettings_Package)(nil),      // 6: buildbucket.SwarmingSettings.Package
	(*ExperimentSettings_Experiment)(nil), // 7: buildbucket.ExperimentSettings.Experiment
	(*Builder_CacheEntry)(nil),            // 8: buildbucket.Builder.CacheEntry
}
var file_go_chromium_org_luci_buildbucket_proto_service_config_proto_depIdxs = []int32{
	1,  // 0: buildbucket.SettingsCfg.swarming:type_name -> buildbucket.SwarmingSettings
	2,  // 1: buildbucket.SettingsCfg.logdog:type_name -> buildbucket.LogDogSettings
	5,  // 2: buildbucket.SettingsCfg.resultdb:type_name -> buildbucket.ResultDBSettings
	3,  // 3: buildbucket.SettingsCfg.experiment:type_name -> buildbucket.ExperimentSettings
	8,  // 4: buildbucket.SwarmingSettings.global_caches:type_name -> buildbucket.Builder.CacheEntry
	6,  // 5: buildbucket.SwarmingSettings.user_packages:type_name -> buildbucket.SwarmingSettings.Package
	6,  // 6: buildbucket.SwarmingSettings.bbagent_package:type_name -> buildbucket.SwarmingSettings.Package
	6,  // 7: buildbucket.SwarmingSettings.kitchen_package:type_name -> buildbucket.SwarmingSettings.Package
	7,  // 8: buildbucket.ExperimentSettings.experiments:type_name -> buildbucket.ExperimentSettings.Experiment
	4,  // 9: buildbucket.SwarmingSettings.Package.builders:type_name -> buildbucket.BuilderPredicate
	4,  // 10: buildbucket.ExperimentSettings.Experiment.builders:type_name -> buildbucket.BuilderPredicate
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_buildbucket_proto_service_config_proto_init() }
func file_go_chromium_org_luci_buildbucket_proto_service_config_proto_init() {
	if File_go_chromium_org_luci_buildbucket_proto_service_config_proto != nil {
		return
	}
	file_go_chromium_org_luci_buildbucket_proto_project_config_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_buildbucket_proto_service_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SettingsCfg); i {
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
		file_go_chromium_org_luci_buildbucket_proto_service_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SwarmingSettings); i {
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
		file_go_chromium_org_luci_buildbucket_proto_service_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogDogSettings); i {
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
		file_go_chromium_org_luci_buildbucket_proto_service_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExperimentSettings); i {
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
		file_go_chromium_org_luci_buildbucket_proto_service_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuilderPredicate); i {
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
		file_go_chromium_org_luci_buildbucket_proto_service_config_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultDBSettings); i {
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
		file_go_chromium_org_luci_buildbucket_proto_service_config_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SwarmingSettings_Package); i {
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
		file_go_chromium_org_luci_buildbucket_proto_service_config_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExperimentSettings_Experiment); i {
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
			RawDescriptor: file_go_chromium_org_luci_buildbucket_proto_service_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_buildbucket_proto_service_config_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_buildbucket_proto_service_config_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_buildbucket_proto_service_config_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_buildbucket_proto_service_config_proto = out.File
	file_go_chromium_org_luci_buildbucket_proto_service_config_proto_rawDesc = nil
	file_go_chromium_org_luci_buildbucket_proto_service_config_proto_goTypes = nil
	file_go_chromium_org_luci_buildbucket_proto_service_config_proto_depIdxs = nil
}
