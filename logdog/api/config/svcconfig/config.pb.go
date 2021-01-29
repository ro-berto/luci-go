// Copyright 2015 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.1
// source: go.chromium.org/luci/logdog/api/config/svcconfig/config.proto

package svcconfig

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Config is the overall instance configuration.
type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Configuration for the Butler's log transport.
	Transport *Transport `protobuf:"bytes,10,opt,name=transport,proto3" json:"transport,omitempty"`
	// Configuration for intermediate Storage.
	Storage *Storage `protobuf:"bytes,11,opt,name=storage,proto3" json:"storage,omitempty"`
	// Coordinator is the coordinator service configuration.
	Coordinator *Coordinator `protobuf:"bytes,20,opt,name=coordinator,proto3" json:"coordinator,omitempty"`
	// Collector is the collector fleet configuration.
	Collector *Collector `protobuf:"bytes,21,opt,name=collector,proto3" json:"collector,omitempty"`
	// Archivist microservice configuration.
	Archivist *Archivist `protobuf:"bytes,22,opt,name=archivist,proto3" json:"archivist,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_logdog_api_config_svcconfig_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_logdog_api_config_svcconfig_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_logdog_api_config_svcconfig_config_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetTransport() *Transport {
	if x != nil {
		return x.Transport
	}
	return nil
}

func (x *Config) GetStorage() *Storage {
	if x != nil {
		return x.Storage
	}
	return nil
}

func (x *Config) GetCoordinator() *Coordinator {
	if x != nil {
		return x.Coordinator
	}
	return nil
}

func (x *Config) GetCollector() *Collector {
	if x != nil {
		return x.Collector
	}
	return nil
}

func (x *Config) GetArchivist() *Archivist {
	if x != nil {
		return x.Archivist
	}
	return nil
}

// Coordinator is the Coordinator service configuration.
type Coordinator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the authentication group for administrators.
	AdminAuthGroup string `protobuf:"bytes,10,opt,name=admin_auth_group,json=adminAuthGroup,proto3" json:"admin_auth_group,omitempty"`
	// The name of the authentication group for backend services.
	ServiceAuthGroup string `protobuf:"bytes,11,opt,name=service_auth_group,json=serviceAuthGroup,proto3" json:"service_auth_group,omitempty"`
	// A list of origin URLs that are allowed to perform CORS RPC calls.
	RpcAllowOrigins []string `protobuf:"bytes,20,rep,name=rpc_allow_origins,json=rpcAllowOrigins,proto3" json:"rpc_allow_origins,omitempty"`
	// The maximum amount of time after a prefix has been registered when log
	// streams may also be registered under that prefix.
	//
	// After the expiration period has passed, new log stream registration will
	// fail.
	//
	// Project configurations or stream prefix regitrations may override this by
	// providing >= 0 values for prefix expiration. The smallest configured
	// expiration will be applied.
	PrefixExpiration *durationpb.Duration `protobuf:"bytes,21,opt,name=prefix_expiration,json=prefixExpiration,proto3" json:"prefix_expiration,omitempty"`
	// The full path of the archival Pub/Sub topic.
	//
	// The Coordinator must have permission to publish to this topic.
	ArchiveTopic string `protobuf:"bytes,30,opt,name=archive_topic,json=archiveTopic,proto3" json:"archive_topic,omitempty"`
	// The amount of time after an archive request has been dispatched before it
	// should be executed.
	//
	// Since terminal messages can arrive out of order, the archival request may
	// be kicked off before all of the log stream data has been loaded into
	// intermediate storage. If this happens, the Archivist will retry archival
	// later autometically.
	//
	// This parameter is an optimization to stop the archivist from wasting its
	// time until the log stream has a reasonable expectation of being available.
	ArchiveSettleDelay *durationpb.Duration `protobuf:"bytes,31,opt,name=archive_settle_delay,json=archiveSettleDelay,proto3" json:"archive_settle_delay,omitempty"`
	// The amount of time before a log stream is candidate for archival regardless
	// of whether or not it's been terminated or complete.
	//
	// This is a failsafe designed to ensure that log streams with missing records
	// or no terminal record (e.g., Butler crashed) are eventually archived.
	//
	// This should be fairly large (days) to avoid prematurely archiving
	// long-running streams, but should be considerably smaller than the
	// intermediate storage data retention period.
	//
	// If a project's "max_stream_age" is smaller than this value, it will be used
	// on that project's streams.
	ArchiveDelayMax *durationpb.Duration `protobuf:"bytes,32,opt,name=archive_delay_max,json=archiveDelayMax,proto3" json:"archive_delay_max,omitempty"`
}

func (x *Coordinator) Reset() {
	*x = Coordinator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_logdog_api_config_svcconfig_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Coordinator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Coordinator) ProtoMessage() {}

func (x *Coordinator) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_logdog_api_config_svcconfig_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Coordinator.ProtoReflect.Descriptor instead.
func (*Coordinator) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_logdog_api_config_svcconfig_config_proto_rawDescGZIP(), []int{1}
}

func (x *Coordinator) GetAdminAuthGroup() string {
	if x != nil {
		return x.AdminAuthGroup
	}
	return ""
}

func (x *Coordinator) GetServiceAuthGroup() string {
	if x != nil {
		return x.ServiceAuthGroup
	}
	return ""
}

func (x *Coordinator) GetRpcAllowOrigins() []string {
	if x != nil {
		return x.RpcAllowOrigins
	}
	return nil
}

func (x *Coordinator) GetPrefixExpiration() *durationpb.Duration {
	if x != nil {
		return x.PrefixExpiration
	}
	return nil
}

func (x *Coordinator) GetArchiveTopic() string {
	if x != nil {
		return x.ArchiveTopic
	}
	return ""
}

func (x *Coordinator) GetArchiveSettleDelay() *durationpb.Duration {
	if x != nil {
		return x.ArchiveSettleDelay
	}
	return nil
}

func (x *Coordinator) GetArchiveDelayMax() *durationpb.Duration {
	if x != nil {
		return x.ArchiveDelayMax
	}
	return nil
}

// Collector is the set of configuration parameters for Collector instances.
type Collector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The maximum number of concurrent transport messages to process. If <= 0,
	// a default will be chosen based on the transport.
	MaxConcurrentMessages int32 `protobuf:"varint,1,opt,name=max_concurrent_messages,json=maxConcurrentMessages,proto3" json:"max_concurrent_messages,omitempty"`
	// The maximum number of concurrent workers to process each ingested message.
	// If <= 0, collector.DefaultMaxMessageWorkers will be used.
	MaxMessageWorkers int32 `protobuf:"varint,2,opt,name=max_message_workers,json=maxMessageWorkers,proto3" json:"max_message_workers,omitempty"`
	// The maximum number of log stream states to cache locally. If <= 0, a
	// default will be used.
	StateCacheSize int32 `protobuf:"varint,3,opt,name=state_cache_size,json=stateCacheSize,proto3" json:"state_cache_size,omitempty"`
	// The maximum amount of time that cached stream state is valid. If <= 0, a
	// default will be used.
	StateCacheExpiration *durationpb.Duration `protobuf:"bytes,4,opt,name=state_cache_expiration,json=stateCacheExpiration,proto3" json:"state_cache_expiration,omitempty"`
}

func (x *Collector) Reset() {
	*x = Collector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_logdog_api_config_svcconfig_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Collector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Collector) ProtoMessage() {}

func (x *Collector) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_logdog_api_config_svcconfig_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Collector.ProtoReflect.Descriptor instead.
func (*Collector) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_logdog_api_config_svcconfig_config_proto_rawDescGZIP(), []int{2}
}

func (x *Collector) GetMaxConcurrentMessages() int32 {
	if x != nil {
		return x.MaxConcurrentMessages
	}
	return 0
}

func (x *Collector) GetMaxMessageWorkers() int32 {
	if x != nil {
		return x.MaxMessageWorkers
	}
	return 0
}

func (x *Collector) GetStateCacheSize() int32 {
	if x != nil {
		return x.StateCacheSize
	}
	return 0
}

func (x *Collector) GetStateCacheExpiration() *durationpb.Duration {
	if x != nil {
		return x.StateCacheExpiration
	}
	return nil
}

// Configuration for the Archivist microservice.
type Archivist struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the archival Pub/Sub subscription.
	//
	// This should be connected to "archive_topic", and the Archivist must have
	// permission to consume from this subscription.
	Subscription string `protobuf:"bytes,1,opt,name=subscription,proto3" json:"subscription,omitempty"`
	// The number of tasks to run at a time. If blank, the archivist will choose a
	// default value.
	Tasks int32 `protobuf:"varint,2,opt,name=tasks,proto3" json:"tasks,omitempty"`
	// The name of the staging storage bucket. All projects will share the same
	// staging bucket. Logs for a project will be staged under:
	//
	// gs://<gs_staging_bucket>/<app-id>/<project-name>/...
	GsStagingBucket string `protobuf:"bytes,3,opt,name=gs_staging_bucket,json=gsStagingBucket,proto3" json:"gs_staging_bucket,omitempty"`
	// Service-wide index configuration. This is used if per-project configuration
	// is not specified.
	ArchiveIndexConfig *ArchiveIndexConfig `protobuf:"bytes,10,opt,name=archive_index_config,json=archiveIndexConfig,proto3" json:"archive_index_config,omitempty"`
}

func (x *Archivist) Reset() {
	*x = Archivist{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_logdog_api_config_svcconfig_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Archivist) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Archivist) ProtoMessage() {}

func (x *Archivist) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_logdog_api_config_svcconfig_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Archivist.ProtoReflect.Descriptor instead.
func (*Archivist) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_logdog_api_config_svcconfig_config_proto_rawDescGZIP(), []int{3}
}

func (x *Archivist) GetSubscription() string {
	if x != nil {
		return x.Subscription
	}
	return ""
}

func (x *Archivist) GetTasks() int32 {
	if x != nil {
		return x.Tasks
	}
	return 0
}

func (x *Archivist) GetGsStagingBucket() string {
	if x != nil {
		return x.GsStagingBucket
	}
	return ""
}

func (x *Archivist) GetArchiveIndexConfig() *ArchiveIndexConfig {
	if x != nil {
		return x.ArchiveIndexConfig
	}
	return nil
}

var File_go_chromium_org_luci_logdog_api_config_svcconfig_config_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_logdog_api_config_svcconfig_config_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x6c, 0x6f, 0x67, 0x64, 0x6f, 0x67, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x73, 0x76, 0x63, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x09, 0x73, 0x76, 0x63, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x3f, 0x67, 0x6f, 0x2e, 0x63,
	0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69,
	0x2f, 0x6c, 0x6f, 0x67, 0x64, 0x6f, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2f, 0x73, 0x76, 0x63, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x61, 0x72, 0x63,
	0x68, 0x69, 0x76, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3e, 0x67, 0x6f, 0x2e,
	0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63,
	0x69, 0x2f, 0x6c, 0x6f, 0x67, 0x64, 0x6f, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2f, 0x73, 0x76, 0x63, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x40, 0x67, 0x6f, 0x2e,
	0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63,
	0x69, 0x2f, 0x6c, 0x6f, 0x67, 0x64, 0x6f, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2f, 0x73, 0x76, 0x63, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c, 0x02,
	0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x32, 0x0a, 0x09, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x76,
	0x63, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x52, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x2c, 0x0a, 0x07,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x73, 0x76, 0x63, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x38, 0x0a, 0x0b, 0x63, 0x6f,
	0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x73, 0x76, 0x63, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6f, 0x72,
	0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x0b, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e,
	0x61, 0x74, 0x6f, 0x72, 0x12, 0x32, 0x0a, 0x09, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x76, 0x63, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x09, 0x63,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x32, 0x0a, 0x09, 0x61, 0x72, 0x63, 0x68,
	0x69, 0x76, 0x69, 0x73, 0x74, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x76,
	0x63, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x69, 0x73,
	0x74, 0x52, 0x09, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x69, 0x73, 0x74, 0x22, 0x92, 0x03, 0x0a,
	0x0b, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x28, 0x0a, 0x10,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x41, 0x75, 0x74,
	0x68, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x2c, 0x0a, 0x12, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x75, 0x74, 0x68, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x12, 0x2a, 0x0a, 0x11, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x6c, 0x6c, 0x6f,
	0x77, 0x5f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x73, 0x18, 0x14, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0f, 0x72, 0x70, 0x63, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x73,
	0x12, 0x46, 0x0a, 0x11, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x45, 0x78,
	0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x72, 0x63, 0x68,
	0x69, 0x76, 0x65, 0x5f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x4b, 0x0a,
	0x14, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x5f,
	0x64, 0x65, 0x6c, 0x61, 0x79, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x12, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x53,
	0x65, 0x74, 0x74, 0x6c, 0x65, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x12, 0x45, 0x0a, 0x11, 0x61, 0x72,
	0x63, 0x68, 0x69, 0x76, 0x65, 0x5f, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x5f, 0x6d, 0x61, 0x78, 0x18,
	0x20, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x4d, 0x61,
	0x78, 0x22, 0xee, 0x01, 0x0a, 0x09, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12,
	0x36, 0x0a, 0x17, 0x6d, 0x61, 0x78, 0x5f, 0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x15, 0x6d, 0x61, 0x78, 0x43, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x2e, 0x0a, 0x13, 0x6d, 0x61, 0x78, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x6d, 0x61, 0x78, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x12, 0x28, 0x0a, 0x10, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x4f, 0x0a, 0x16, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65,
	0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x14, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65, 0x45, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0xdc, 0x01, 0x0a, 0x09, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x69, 0x73, 0x74,
	0x12, 0x22, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x67, 0x73,
	0x5f, 0x73, 0x74, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x67, 0x73, 0x53, 0x74, 0x61, 0x67, 0x69, 0x6e, 0x67,
	0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x4f, 0x0a, 0x14, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76,
	0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x76, 0x63, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x12, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4a, 0x04, 0x08, 0x0d, 0x10, 0x0e, 0x52, 0x12, 0x72,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x6c, 0x6c, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x73, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d,
	0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x6c, 0x6f, 0x67, 0x64, 0x6f, 0x67,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x73, 0x76, 0x63, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_logdog_api_config_svcconfig_config_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_logdog_api_config_svcconfig_config_proto_rawDescData = file_go_chromium_org_luci_logdog_api_config_svcconfig_config_proto_rawDesc
)

func file_go_chromium_org_luci_logdog_api_config_svcconfig_config_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_logdog_api_config_svcconfig_config_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_logdog_api_config_svcconfig_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_logdog_api_config_svcconfig_config_proto_rawDescData)
	})
	return file_go_chromium_org_luci_logdog_api_config_svcconfig_config_proto_rawDescData
}

var file_go_chromium_org_luci_logdog_api_config_svcconfig_config_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_go_chromium_org_luci_logdog_api_config_svcconfig_config_proto_goTypes = []interface{}{
	(*Config)(nil),              // 0: svcconfig.Config
	(*Coordinator)(nil),         // 1: svcconfig.Coordinator
	(*Collector)(nil),           // 2: svcconfig.Collector
	(*Archivist)(nil),           // 3: svcconfig.Archivist
	(*Transport)(nil),           // 4: svcconfig.Transport
	(*Storage)(nil),             // 5: svcconfig.Storage
	(*durationpb.Duration)(nil), // 6: google.protobuf.Duration
	(*ArchiveIndexConfig)(nil),  // 7: svcconfig.ArchiveIndexConfig
}
var file_go_chromium_org_luci_logdog_api_config_svcconfig_config_proto_depIdxs = []int32{
	4,  // 0: svcconfig.Config.transport:type_name -> svcconfig.Transport
	5,  // 1: svcconfig.Config.storage:type_name -> svcconfig.Storage
	1,  // 2: svcconfig.Config.coordinator:type_name -> svcconfig.Coordinator
	2,  // 3: svcconfig.Config.collector:type_name -> svcconfig.Collector
	3,  // 4: svcconfig.Config.archivist:type_name -> svcconfig.Archivist
	6,  // 5: svcconfig.Coordinator.prefix_expiration:type_name -> google.protobuf.Duration
	6,  // 6: svcconfig.Coordinator.archive_settle_delay:type_name -> google.protobuf.Duration
	6,  // 7: svcconfig.Coordinator.archive_delay_max:type_name -> google.protobuf.Duration
	6,  // 8: svcconfig.Collector.state_cache_expiration:type_name -> google.protobuf.Duration
	7,  // 9: svcconfig.Archivist.archive_index_config:type_name -> svcconfig.ArchiveIndexConfig
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_logdog_api_config_svcconfig_config_proto_init() }
func file_go_chromium_org_luci_logdog_api_config_svcconfig_config_proto_init() {
	if File_go_chromium_org_luci_logdog_api_config_svcconfig_config_proto != nil {
		return
	}
	file_go_chromium_org_luci_logdog_api_config_svcconfig_archival_proto_init()
	file_go_chromium_org_luci_logdog_api_config_svcconfig_storage_proto_init()
	file_go_chromium_org_luci_logdog_api_config_svcconfig_transport_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_logdog_api_config_svcconfig_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
		file_go_chromium_org_luci_logdog_api_config_svcconfig_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Coordinator); i {
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
		file_go_chromium_org_luci_logdog_api_config_svcconfig_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Collector); i {
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
		file_go_chromium_org_luci_logdog_api_config_svcconfig_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Archivist); i {
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
			RawDescriptor: file_go_chromium_org_luci_logdog_api_config_svcconfig_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_logdog_api_config_svcconfig_config_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_logdog_api_config_svcconfig_config_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_logdog_api_config_svcconfig_config_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_logdog_api_config_svcconfig_config_proto = out.File
	file_go_chromium_org_luci_logdog_api_config_svcconfig_config_proto_rawDesc = nil
	file_go_chromium_org_luci_logdog_api_config_svcconfig_config_proto_goTypes = nil
	file_go_chromium_org_luci_logdog_api_config_svcconfig_config_proto_depIdxs = nil
}
