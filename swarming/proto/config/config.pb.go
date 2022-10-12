// Copyright 2016 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: go.chromium.org/luci/swarming/proto/config/config.proto

package configpb

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

// Schema for settings.cfg service config file in luci-config.
type SettingsCfg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id to inject into pages if applicable.
	GoogleAnalytics string `protobuf:"bytes,1,opt,name=google_analytics,json=googleAnalytics,proto3" json:"google_analytics,omitempty"`
	// The number of seconds an old task can be deduped from.
	// Default is one week: 7*24*60*60 = 604800
	ReusableTaskAgeSecs int32 `protobuf:"varint,2,opt,name=reusable_task_age_secs,json=reusableTaskAgeSecs,proto3" json:"reusable_task_age_secs,omitempty"`
	// The amount of time that has to pass before a machine is considered dead.
	// Default is 600 (10 minutes).
	BotDeathTimeoutSecs int32 `protobuf:"varint,3,opt,name=bot_death_timeout_secs,json=botDeathTimeoutSecs,proto3" json:"bot_death_timeout_secs,omitempty"`
	// Enable ts_mon based monitoring.
	EnableTsMonitoring bool `protobuf:"varint,4,opt,name=enable_ts_monitoring,json=enableTsMonitoring,proto3" json:"enable_ts_monitoring,omitempty"`
	// (deprecated, see pools.proto) Configuration for swarming-cipd integration.
	Cipd *CipdSettings `protobuf:"bytes,6,opt,name=cipd,proto3" json:"cipd,omitempty"`
	// Emergency setting to disable bot task reaping. When set, all bots are
	// always put to sleep and are never granted task.
	ForceBotsToSleepAndNotRunTask bool `protobuf:"varint,8,opt,name=force_bots_to_sleep_and_not_run_task,json=forceBotsToSleepAndNotRunTask,proto3" json:"force_bots_to_sleep_and_not_run_task,omitempty"`
	// oauth client id for the ui. This is created in the developer's console
	// under Credentials.
	UiClientId string `protobuf:"bytes,9,opt,name=ui_client_id,json=uiClientId,proto3" json:"ui_client_id,omitempty"`
	// A url to a task display server (e.g. milo).  This should have a %s where
	// a task id can go.
	DisplayServerUrlTemplate string `protobuf:"bytes,11,opt,name=display_server_url_template,json=displayServerUrlTemplate,proto3" json:"display_server_url_template,omitempty"`
	// Sets a maximum sleep time in seconds for bots that limits the exponental
	// backoff. If missing, the task scheduler will provide the default maximum
	// (usually 60s, but see bot_code/task_scheduler.py for details).
	MaxBotSleepTime int32 `protobuf:"varint,12,opt,name=max_bot_sleep_time,json=maxBotSleepTime,proto3" json:"max_bot_sleep_time,omitempty"`
	// Names of the authorization groups used by components/auth.
	Auth *AuthSettings `protobuf:"bytes,13,opt,name=auth,proto3" json:"auth,omitempty"`
	// Sets the default gRPC proxy for the bot's Isolate server calls.
	BotIsolateGrpcProxy string `protobuf:"bytes,14,opt,name=bot_isolate_grpc_proxy,json=botIsolateGrpcProxy,proto3" json:"bot_isolate_grpc_proxy,omitempty"`
	// Sets the default gRPC proxy for the bot's Swarming server calls.
	BotSwarmingGrpcProxy string `protobuf:"bytes,15,opt,name=bot_swarming_grpc_proxy,json=botSwarmingGrpcProxy,proto3" json:"bot_swarming_grpc_proxy,omitempty"`
	// Any extra urls that should be added to frame-src, e.g. anything that
	// will be linked to from the display server.
	// This originally added things to child-src, which was deprecated:
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy/child-src
	ExtraChildSrcCspUrl []string `protobuf:"bytes,16,rep,name=extra_child_src_csp_url,json=extraChildSrcCspUrl,proto3" json:"extra_child_src_csp_url,omitempty"`
	// Whether tasks should be run in FIFO or LIFO order.
	UseLifo bool `protobuf:"varint,17,opt,name=use_lifo,json=useLifo,proto3" json:"use_lifo,omitempty"`
	// Whether swarming should batch notifications to the external scheduler.
	EnableBatchEsNotifications bool `protobuf:"varint,18,opt,name=enable_batch_es_notifications,json=enableBatchEsNotifications,proto3" json:"enable_batch_es_notifications,omitempty"`
	// Configuration for Swarming-ResultDB integration.
	Resultdb *ResultDBSettings `protobuf:"bytes,19,opt,name=resultdb,proto3" json:"resultdb,omitempty"`
	// Configuration for RBE-CAS integration.
	Cas *CASSettings `protobuf:"bytes,20,opt,name=cas,proto3" json:"cas,omitempty"`
}

func (x *SettingsCfg) Reset() {
	*x = SettingsCfg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_swarming_proto_config_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SettingsCfg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SettingsCfg) ProtoMessage() {}

func (x *SettingsCfg) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_swarming_proto_config_config_proto_msgTypes[0]
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
	return file_go_chromium_org_luci_swarming_proto_config_config_proto_rawDescGZIP(), []int{0}
}

func (x *SettingsCfg) GetGoogleAnalytics() string {
	if x != nil {
		return x.GoogleAnalytics
	}
	return ""
}

func (x *SettingsCfg) GetReusableTaskAgeSecs() int32 {
	if x != nil {
		return x.ReusableTaskAgeSecs
	}
	return 0
}

func (x *SettingsCfg) GetBotDeathTimeoutSecs() int32 {
	if x != nil {
		return x.BotDeathTimeoutSecs
	}
	return 0
}

func (x *SettingsCfg) GetEnableTsMonitoring() bool {
	if x != nil {
		return x.EnableTsMonitoring
	}
	return false
}

func (x *SettingsCfg) GetCipd() *CipdSettings {
	if x != nil {
		return x.Cipd
	}
	return nil
}

func (x *SettingsCfg) GetForceBotsToSleepAndNotRunTask() bool {
	if x != nil {
		return x.ForceBotsToSleepAndNotRunTask
	}
	return false
}

func (x *SettingsCfg) GetUiClientId() string {
	if x != nil {
		return x.UiClientId
	}
	return ""
}

func (x *SettingsCfg) GetDisplayServerUrlTemplate() string {
	if x != nil {
		return x.DisplayServerUrlTemplate
	}
	return ""
}

func (x *SettingsCfg) GetMaxBotSleepTime() int32 {
	if x != nil {
		return x.MaxBotSleepTime
	}
	return 0
}

func (x *SettingsCfg) GetAuth() *AuthSettings {
	if x != nil {
		return x.Auth
	}
	return nil
}

func (x *SettingsCfg) GetBotIsolateGrpcProxy() string {
	if x != nil {
		return x.BotIsolateGrpcProxy
	}
	return ""
}

func (x *SettingsCfg) GetBotSwarmingGrpcProxy() string {
	if x != nil {
		return x.BotSwarmingGrpcProxy
	}
	return ""
}

func (x *SettingsCfg) GetExtraChildSrcCspUrl() []string {
	if x != nil {
		return x.ExtraChildSrcCspUrl
	}
	return nil
}

func (x *SettingsCfg) GetUseLifo() bool {
	if x != nil {
		return x.UseLifo
	}
	return false
}

func (x *SettingsCfg) GetEnableBatchEsNotifications() bool {
	if x != nil {
		return x.EnableBatchEsNotifications
	}
	return false
}

func (x *SettingsCfg) GetResultdb() *ResultDBSettings {
	if x != nil {
		return x.Resultdb
	}
	return nil
}

func (x *SettingsCfg) GetCas() *CASSettings {
	if x != nil {
		return x.Cas
	}
	return nil
}

// A CIPD package.
type CipdPackage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A template of a full CIPD package name, e.g.
	// "infra/tools/cipd/${platform}"
	// See also cipd.ALL_PARAMS.
	PackageName string `protobuf:"bytes,1,opt,name=package_name,json=packageName,proto3" json:"package_name,omitempty"`
	// Valid package version for all packages matched by package name.
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *CipdPackage) Reset() {
	*x = CipdPackage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_swarming_proto_config_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CipdPackage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CipdPackage) ProtoMessage() {}

func (x *CipdPackage) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_swarming_proto_config_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CipdPackage.ProtoReflect.Descriptor instead.
func (*CipdPackage) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_swarming_proto_config_config_proto_rawDescGZIP(), []int{1}
}

func (x *CipdPackage) GetPackageName() string {
	if x != nil {
		return x.PackageName
	}
	return ""
}

func (x *CipdPackage) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

// Settings for Swarming-CIPD integration.
type CipdSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// URL of the default CIPD server to use if it is not specified in a task.
	// Must start with "https://" or "http://",
	// e.g. "https://chrome-infra-packages.appspot.com".
	DefaultServer string `protobuf:"bytes,1,opt,name=default_server,json=defaultServer,proto3" json:"default_server,omitempty"`
	// Package of the default CIPD client to use if it is not specified in a
	// task.
	DefaultClientPackage *CipdPackage `protobuf:"bytes,2,opt,name=default_client_package,json=defaultClientPackage,proto3" json:"default_client_package,omitempty"`
}

func (x *CipdSettings) Reset() {
	*x = CipdSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_swarming_proto_config_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CipdSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CipdSettings) ProtoMessage() {}

func (x *CipdSettings) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_swarming_proto_config_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CipdSettings.ProtoReflect.Descriptor instead.
func (*CipdSettings) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_swarming_proto_config_config_proto_rawDescGZIP(), []int{2}
}

func (x *CipdSettings) GetDefaultServer() string {
	if x != nil {
		return x.DefaultServer
	}
	return ""
}

func (x *CipdSettings) GetDefaultClientPackage() *CipdPackage {
	if x != nil {
		return x.DefaultClientPackage
	}
	return nil
}

type AuthSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Members of this group have full administrative access.
	//
	// Grants:
	// - config view and edit
	// - delete any bot
	// - all of bot_bootstrap_group membership
	// - all of privileged_users_group membership
	AdminsGroup string `protobuf:"bytes,1,opt,name=admins_group,json=adminsGroup,proto3" json:"admins_group,omitempty"`
	// Members of this group can fetch swarming bot code and bootstrap bots.
	//
	// Grants:
	// - bot create: create a token to anonymously fetch the bot code.
	BotBootstrapGroup string `protobuf:"bytes,2,opt,name=bot_bootstrap_group,json=botBootstrapGroup,proto3" json:"bot_bootstrap_group,omitempty"`
	// Members of this group can schedule tasks and see everyone else's tasks.
	//
	// Grants:
	// - cancel any task
	// - edit (terminate) any bot
	// - all of view_all_bots_group membership
	// - all of view_all_tasks_group membership
	PrivilegedUsersGroup string `protobuf:"bytes,3,opt,name=privileged_users_group,json=privilegedUsersGroup,proto3" json:"privileged_users_group,omitempty"`
	// Members of this group can schedule tasks and see only their own tasks.
	//
	// Grants:
	// - create a task
	// - view and edit own task
	UsersGroup string `protobuf:"bytes,4,opt,name=users_group,json=usersGroup,proto3" json:"users_group,omitempty"`
	// Members of this group can view all bots. This is a read-only group.
	//
	// Grants:
	// - view all bots
	ViewAllBotsGroup string `protobuf:"bytes,5,opt,name=view_all_bots_group,json=viewAllBotsGroup,proto3" json:"view_all_bots_group,omitempty"`
	// Members of this group can view all tasks. This is a read-only group.
	//
	// Grants:
	// - view all tasks
	ViewAllTasksGroup string `protobuf:"bytes,6,opt,name=view_all_tasks_group,json=viewAllTasksGroup,proto3" json:"view_all_tasks_group,omitempty"`
	// List of Realm permissions enforced by default.
	// This field will be deprecated after migration.
	EnforcedRealmPermissions []RealmPermission `protobuf:"varint,7,rep,packed,name=enforced_realm_permissions,json=enforcedRealmPermissions,proto3,enum=swarming.config.RealmPermission" json:"enforced_realm_permissions,omitempty"`
}

func (x *AuthSettings) Reset() {
	*x = AuthSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_swarming_proto_config_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthSettings) ProtoMessage() {}

func (x *AuthSettings) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_swarming_proto_config_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthSettings.ProtoReflect.Descriptor instead.
func (*AuthSettings) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_swarming_proto_config_config_proto_rawDescGZIP(), []int{3}
}

func (x *AuthSettings) GetAdminsGroup() string {
	if x != nil {
		return x.AdminsGroup
	}
	return ""
}

func (x *AuthSettings) GetBotBootstrapGroup() string {
	if x != nil {
		return x.BotBootstrapGroup
	}
	return ""
}

func (x *AuthSettings) GetPrivilegedUsersGroup() string {
	if x != nil {
		return x.PrivilegedUsersGroup
	}
	return ""
}

func (x *AuthSettings) GetUsersGroup() string {
	if x != nil {
		return x.UsersGroup
	}
	return ""
}

func (x *AuthSettings) GetViewAllBotsGroup() string {
	if x != nil {
		return x.ViewAllBotsGroup
	}
	return ""
}

func (x *AuthSettings) GetViewAllTasksGroup() string {
	if x != nil {
		return x.ViewAllTasksGroup
	}
	return ""
}

func (x *AuthSettings) GetEnforcedRealmPermissions() []RealmPermission {
	if x != nil {
		return x.EnforcedRealmPermissions
	}
	return nil
}

// Settings for Swarming-ResultDB integration.
type ResultDBSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// URL of the ResultDB server to use.
	// Must start with "https://" or "http://",
	// e.g. "https://results.api.cr.dev".
	Server string `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
}

func (x *ResultDBSettings) Reset() {
	*x = ResultDBSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_swarming_proto_config_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultDBSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultDBSettings) ProtoMessage() {}

func (x *ResultDBSettings) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_swarming_proto_config_config_proto_msgTypes[4]
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
	return file_go_chromium_org_luci_swarming_proto_config_config_proto_rawDescGZIP(), []int{4}
}

func (x *ResultDBSettings) GetServer() string {
	if x != nil {
		return x.Server
	}
	return ""
}

// Settings for CAS integration.
type CASSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Host of the CAS viewer.
	// Must start with "https://" or "http://",
	// e.g. "https://cas-viewer.appspot.com".
	ViewerServer string `protobuf:"bytes,1,opt,name=viewer_server,json=viewerServer,proto3" json:"viewer_server,omitempty"`
}

func (x *CASSettings) Reset() {
	*x = CASSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_swarming_proto_config_config_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CASSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CASSettings) ProtoMessage() {}

func (x *CASSettings) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_swarming_proto_config_config_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CASSettings.ProtoReflect.Descriptor instead.
func (*CASSettings) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_swarming_proto_config_config_proto_rawDescGZIP(), []int{5}
}

func (x *CASSettings) GetViewerServer() string {
	if x != nil {
		return x.ViewerServer
	}
	return ""
}

var File_go_chromium_org_luci_swarming_proto_config_config_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_swarming_proto_config_config_proto_rawDesc = []byte{
	0x0a, 0x37, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x73, 0x77, 0x61, 0x72, 0x6d, 0x69, 0x6e, 0x67, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x73, 0x77, 0x61, 0x72, 0x6d,
	0x69, 0x6e, 0x67, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x2f, 0x67, 0x6f, 0x2e, 0x63,
	0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x37, 0x67, 0x6f, 0x2e,
	0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63,
	0x69, 0x2f, 0x73, 0x77, 0x61, 0x72, 0x6d, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x96, 0x07, 0x0a, 0x0b, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x43, 0x66, 0x67, 0x12, 0x29, 0x0a, 0x10, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x61,
	0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x12,
	0x33, 0x0a, 0x16, 0x72, 0x65, 0x75, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x74, 0x61, 0x73, 0x6b,
	0x5f, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x13, 0x72, 0x65, 0x75, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x41, 0x67, 0x65,
	0x53, 0x65, 0x63, 0x73, 0x12, 0x33, 0x0a, 0x16, 0x62, 0x6f, 0x74, 0x5f, 0x64, 0x65, 0x61, 0x74,
	0x68, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x13, 0x62, 0x6f, 0x74, 0x44, 0x65, 0x61, 0x74, 0x68, 0x54, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x53, 0x65, 0x63, 0x73, 0x12, 0x30, 0x0a, 0x14, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x5f, 0x74, 0x73, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e,
	0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x54,
	0x73, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x31, 0x0a, 0x04, 0x63,
	0x69, 0x70, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x77, 0x61, 0x72,
	0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x69, 0x70, 0x64,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x04, 0x63, 0x69, 0x70, 0x64, 0x12, 0x4b,
	0x0a, 0x24, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x5f, 0x62, 0x6f, 0x74, 0x73, 0x5f, 0x74, 0x6f, 0x5f,
	0x73, 0x6c, 0x65, 0x65, 0x70, 0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x72, 0x75,
	0x6e, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1d, 0x66, 0x6f,
	0x72, 0x63, 0x65, 0x42, 0x6f, 0x74, 0x73, 0x54, 0x6f, 0x53, 0x6c, 0x65, 0x65, 0x70, 0x41, 0x6e,
	0x64, 0x4e, 0x6f, 0x74, 0x52, 0x75, 0x6e, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x20, 0x0a, 0x0c, 0x75,
	0x69, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x75, 0x69, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x3d, 0x0a,
	0x1b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f,
	0x75, 0x72, 0x6c, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x18, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x55, 0x72, 0x6c, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x0a, 0x12,
	0x6d, 0x61, 0x78, 0x5f, 0x62, 0x6f, 0x74, 0x5f, 0x73, 0x6c, 0x65, 0x65, 0x70, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x6d, 0x61, 0x78, 0x42, 0x6f, 0x74,
	0x53, 0x6c, 0x65, 0x65, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x61, 0x75, 0x74,
	0x68, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x77, 0x61, 0x72, 0x6d, 0x69,
	0x6e, 0x67, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x12, 0x33, 0x0a, 0x16,
	0x62, 0x6f, 0x74, 0x5f, 0x69, 0x73, 0x6f, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x67, 0x72, 0x70, 0x63,
	0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x62, 0x6f,
	0x74, 0x49, 0x73, 0x6f, 0x6c, 0x61, 0x74, 0x65, 0x47, 0x72, 0x70, 0x63, 0x50, 0x72, 0x6f, 0x78,
	0x79, 0x12, 0x35, 0x0a, 0x17, 0x62, 0x6f, 0x74, 0x5f, 0x73, 0x77, 0x61, 0x72, 0x6d, 0x69, 0x6e,
	0x67, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x14, 0x62, 0x6f, 0x74, 0x53, 0x77, 0x61, 0x72, 0x6d, 0x69, 0x6e, 0x67, 0x47,
	0x72, 0x70, 0x63, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x34, 0x0a, 0x17, 0x65, 0x78, 0x74, 0x72,
	0x61, 0x5f, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x5f, 0x73, 0x72, 0x63, 0x5f, 0x63, 0x73, 0x70, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x10, 0x20, 0x03, 0x28, 0x09, 0x52, 0x13, 0x65, 0x78, 0x74, 0x72, 0x61,
	0x43, 0x68, 0x69, 0x6c, 0x64, 0x53, 0x72, 0x63, 0x43, 0x73, 0x70, 0x55, 0x72, 0x6c, 0x12, 0x19,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x5f, 0x6c, 0x69, 0x66, 0x6f, 0x18, 0x11, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x75, 0x73, 0x65, 0x4c, 0x69, 0x66, 0x6f, 0x12, 0x41, 0x0a, 0x1d, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x5f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x65, 0x73, 0x5f, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x12, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x1a, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x45, 0x73, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3d, 0x0a, 0x08,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x64, 0x62, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x73, 0x77, 0x61, 0x72, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x44, 0x42, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x52, 0x08, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x64, 0x62, 0x12, 0x2e, 0x0a, 0x03, 0x63,
	0x61, 0x73, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x77, 0x61, 0x72, 0x6d,
	0x69, 0x6e, 0x67, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x41, 0x53, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x03, 0x63, 0x61, 0x73, 0x4a, 0x04, 0x08, 0x07, 0x10,
	0x08, 0x4a, 0x04, 0x08, 0x0a, 0x10, 0x0b, 0x4a, 0x04, 0x08, 0x05, 0x10, 0x06, 0x22, 0x4a, 0x0a,
	0x0b, 0x43, 0x69, 0x70, 0x64, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x89, 0x01, 0x0a, 0x0c, 0x43, 0x69,
	0x70, 0x64, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x12, 0x52, 0x0a, 0x16, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x77, 0x61, 0x72, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x43, 0x69, 0x70, 0x64, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52,
	0x14, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x22, 0xf8, 0x02, 0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x73,
	0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x73, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x2e, 0x0a, 0x13, 0x62, 0x6f, 0x74,
	0x5f, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x62, 0x6f, 0x74, 0x42, 0x6f, 0x6f, 0x74, 0x73,
	0x74, 0x72, 0x61, 0x70, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x34, 0x0a, 0x16, 0x70, 0x72, 0x69,
	0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x70, 0x72, 0x69, 0x76, 0x69,
	0x6c, 0x65, 0x67, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x73, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12,
	0x1f, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x73, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x12, 0x2d, 0x0a, 0x13, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x61, 0x6c, 0x6c, 0x5f, 0x62, 0x6f, 0x74,
	0x73, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x76,
	0x69, 0x65, 0x77, 0x41, 0x6c, 0x6c, 0x42, 0x6f, 0x74, 0x73, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12,
	0x2f, 0x0a, 0x14, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x61, 0x6c, 0x6c, 0x5f, 0x74, 0x61, 0x73, 0x6b,
	0x73, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x76,
	0x69, 0x65, 0x77, 0x41, 0x6c, 0x6c, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x12, 0x5e, 0x0a, 0x1a, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x61,
	0x6c, 0x6d, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x73, 0x77, 0x61, 0x72, 0x6d, 0x69, 0x6e, 0x67, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x52, 0x65, 0x61, 0x6c, 0x6d, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x18, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x64,
	0x52, 0x65, 0x61, 0x6c, 0x6d, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x22, 0x2a, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x44, 0x42, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x32, 0x0a, 0x0b,
	0x43, 0x41, 0x53, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x76,
	0x69, 0x65, 0x77, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x76, 0x69, 0x65, 0x77, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x42, 0x81, 0x01, 0x5a, 0x33, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d,
	0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x73, 0x77, 0x61, 0x72, 0x6d, 0x69,
	0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x3b,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x70, 0x62, 0xa2, 0xfe, 0x23, 0x48, 0x0a, 0x46, 0x68, 0x74,
	0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2d, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x70, 0x6f, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x73,
	0x77, 0x61, 0x72, 0x6d, 0x69, 0x6e, 0x67, 0x3a, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x2e, 0x63, 0x66, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_swarming_proto_config_config_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_swarming_proto_config_config_proto_rawDescData = file_go_chromium_org_luci_swarming_proto_config_config_proto_rawDesc
)

func file_go_chromium_org_luci_swarming_proto_config_config_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_swarming_proto_config_config_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_swarming_proto_config_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_swarming_proto_config_config_proto_rawDescData)
	})
	return file_go_chromium_org_luci_swarming_proto_config_config_proto_rawDescData
}

var file_go_chromium_org_luci_swarming_proto_config_config_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_go_chromium_org_luci_swarming_proto_config_config_proto_goTypes = []interface{}{
	(*SettingsCfg)(nil),      // 0: swarming.config.SettingsCfg
	(*CipdPackage)(nil),      // 1: swarming.config.CipdPackage
	(*CipdSettings)(nil),     // 2: swarming.config.CipdSettings
	(*AuthSettings)(nil),     // 3: swarming.config.AuthSettings
	(*ResultDBSettings)(nil), // 4: swarming.config.ResultDBSettings
	(*CASSettings)(nil),      // 5: swarming.config.CASSettings
	(RealmPermission)(0),     // 6: swarming.config.RealmPermission
}
var file_go_chromium_org_luci_swarming_proto_config_config_proto_depIdxs = []int32{
	2, // 0: swarming.config.SettingsCfg.cipd:type_name -> swarming.config.CipdSettings
	3, // 1: swarming.config.SettingsCfg.auth:type_name -> swarming.config.AuthSettings
	4, // 2: swarming.config.SettingsCfg.resultdb:type_name -> swarming.config.ResultDBSettings
	5, // 3: swarming.config.SettingsCfg.cas:type_name -> swarming.config.CASSettings
	1, // 4: swarming.config.CipdSettings.default_client_package:type_name -> swarming.config.CipdPackage
	6, // 5: swarming.config.AuthSettings.enforced_realm_permissions:type_name -> swarming.config.RealmPermission
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_swarming_proto_config_config_proto_init() }
func file_go_chromium_org_luci_swarming_proto_config_config_proto_init() {
	if File_go_chromium_org_luci_swarming_proto_config_config_proto != nil {
		return
	}
	file_go_chromium_org_luci_swarming_proto_config_realms_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_swarming_proto_config_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_go_chromium_org_luci_swarming_proto_config_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CipdPackage); i {
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
		file_go_chromium_org_luci_swarming_proto_config_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CipdSettings); i {
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
		file_go_chromium_org_luci_swarming_proto_config_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthSettings); i {
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
		file_go_chromium_org_luci_swarming_proto_config_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_go_chromium_org_luci_swarming_proto_config_config_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CASSettings); i {
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
			RawDescriptor: file_go_chromium_org_luci_swarming_proto_config_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_swarming_proto_config_config_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_swarming_proto_config_config_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_swarming_proto_config_config_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_swarming_proto_config_config_proto = out.File
	file_go_chromium_org_luci_swarming_proto_config_config_proto_rawDesc = nil
	file_go_chromium_org_luci_swarming_proto_config_config_proto_goTypes = nil
	file_go_chromium_org_luci_swarming_proto_config_config_proto_depIdxs = nil
}
