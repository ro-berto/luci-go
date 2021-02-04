// Copyright 2018 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.1
// source: go.chromium.org/luci/buildbucket/appengine/tasks/defs/tasks.proto

package taskdefs

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

// A task to cancel a Swarming task.
type CancelSwarmingTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Hostname (e.g. "chromium-swarm.appspot.com") where the task should be
	// cancelled.
	Hostname string `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	// Swarming task ID to cancel.
	TaskId string `protobuf:"bytes,2,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	// LUCI realm the task belongs to or "" if realms is disabled.
	Realm string `protobuf:"bytes,3,opt,name=realm,proto3" json:"realm,omitempty"`
}

func (x *CancelSwarmingTask) Reset() {
	*x = CancelSwarmingTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_buildbucket_appengine_tasks_defs_tasks_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelSwarmingTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelSwarmingTask) ProtoMessage() {}

func (x *CancelSwarmingTask) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_buildbucket_appengine_tasks_defs_tasks_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelSwarmingTask.ProtoReflect.Descriptor instead.
func (*CancelSwarmingTask) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_buildbucket_appengine_tasks_defs_tasks_proto_rawDescGZIP(), []int{0}
}

func (x *CancelSwarmingTask) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *CancelSwarmingTask) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *CancelSwarmingTask) GetRealm() string {
	if x != nil {
		return x.Realm
	}
	return ""
}

// A task to create a Swarming task.
type CreateSwarmingTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of a build in the datastore. See model.Build.
	BuildId int64 `protobuf:"varint,1,opt,name=build_id,json=buildId,proto3" json:"build_id,omitempty"`
}

func (x *CreateSwarmingTask) Reset() {
	*x = CreateSwarmingTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_buildbucket_appengine_tasks_defs_tasks_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSwarmingTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSwarmingTask) ProtoMessage() {}

func (x *CreateSwarmingTask) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_buildbucket_appengine_tasks_defs_tasks_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSwarmingTask.ProtoReflect.Descriptor instead.
func (*CreateSwarmingTask) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_buildbucket_appengine_tasks_defs_tasks_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSwarmingTask) GetBuildId() int64 {
	if x != nil {
		return x.BuildId
	}
	return 0
}

// A task to export a build to BigQuery.
type ExportBigQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of a build in the datastore. See model.Build.
	BuildId int64 `protobuf:"varint,1,opt,name=build_id,json=buildId,proto3" json:"build_id,omitempty"`
}

func (x *ExportBigQuery) Reset() {
	*x = ExportBigQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_buildbucket_appengine_tasks_defs_tasks_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportBigQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportBigQuery) ProtoMessage() {}

func (x *ExportBigQuery) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_buildbucket_appengine_tasks_defs_tasks_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportBigQuery.ProtoReflect.Descriptor instead.
func (*ExportBigQuery) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_buildbucket_appengine_tasks_defs_tasks_proto_rawDescGZIP(), []int{2}
}

func (x *ExportBigQuery) GetBuildId() int64 {
	if x != nil {
		return x.BuildId
	}
	return 0
}

// A task to finalize an invocation in ResultDB.
type FinalizeResultDB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of a build in the datastore. See model.Build.
	BuildId int64 `protobuf:"varint,1,opt,name=build_id,json=buildId,proto3" json:"build_id,omitempty"`
}

func (x *FinalizeResultDB) Reset() {
	*x = FinalizeResultDB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_buildbucket_appengine_tasks_defs_tasks_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinalizeResultDB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinalizeResultDB) ProtoMessage() {}

func (x *FinalizeResultDB) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_buildbucket_appengine_tasks_defs_tasks_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinalizeResultDB.ProtoReflect.Descriptor instead.
func (*FinalizeResultDB) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_buildbucket_appengine_tasks_defs_tasks_proto_rawDescGZIP(), []int{3}
}

func (x *FinalizeResultDB) GetBuildId() int64 {
	if x != nil {
		return x.BuildId
	}
	return 0
}

// A task to send a build notification on Pub/Sub.
type NotifyPubSub struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of a build in the datastore. See model.Build.
	BuildId int64 `protobuf:"varint,1,opt,name=build_id,json=buildId,proto3" json:"build_id,omitempty"`
	// Whether to use the global or build-specific callback Pub/Sub topic.
	Callback bool `protobuf:"varint,2,opt,name=callback,proto3" json:"callback,omitempty"`
}

func (x *NotifyPubSub) Reset() {
	*x = NotifyPubSub{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_buildbucket_appengine_tasks_defs_tasks_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyPubSub) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyPubSub) ProtoMessage() {}

func (x *NotifyPubSub) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_buildbucket_appengine_tasks_defs_tasks_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyPubSub.ProtoReflect.Descriptor instead.
func (*NotifyPubSub) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_buildbucket_appengine_tasks_defs_tasks_proto_rawDescGZIP(), []int{4}
}

func (x *NotifyPubSub) GetBuildId() int64 {
	if x != nil {
		return x.BuildId
	}
	return 0
}

func (x *NotifyPubSub) GetCallback() bool {
	if x != nil {
		return x.Callback
	}
	return false
}

var File_go_chromium_org_luci_buildbucket_appengine_tasks_defs_tasks_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_buildbucket_appengine_tasks_defs_tasks_proto_rawDesc = []byte{
	0x0a, 0x41, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x74, 0x61, 0x73,
	0x6b, 0x73, 0x2f, 0x64, 0x65, 0x66, 0x73, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x64, 0x65, 0x66, 0x73, 0x22, 0x5f, 0x0a,
	0x12, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x53, 0x77, 0x61, 0x72, 0x6d, 0x69, 0x6e, 0x67, 0x54,
	0x61, 0x73, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x61, 0x6c,
	0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x22, 0x2f,
	0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x77, 0x61, 0x72, 0x6d, 0x69, 0x6e, 0x67,
	0x54, 0x61, 0x73, 0x6b, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x64, 0x22,
	0x2b, 0x0a, 0x0e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x42, 0x69, 0x67, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x64, 0x22, 0x2d, 0x0a, 0x10,
	0x46, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x44, 0x42,
	0x12, 0x19, 0x0a, 0x08, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x64, 0x22, 0x45, 0x0a, 0x0c, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x50, 0x75, 0x62, 0x53, 0x75, 0x62, 0x12, 0x19, 0x0a, 0x08, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61,
	0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61,
	0x63, 0x6b, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75,
	0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2f, 0x64, 0x65, 0x66, 0x73, 0x3b, 0x74, 0x61, 0x73, 0x6b,
	0x64, 0x65, 0x66, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_buildbucket_appengine_tasks_defs_tasks_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_buildbucket_appengine_tasks_defs_tasks_proto_rawDescData = file_go_chromium_org_luci_buildbucket_appengine_tasks_defs_tasks_proto_rawDesc
)

func file_go_chromium_org_luci_buildbucket_appengine_tasks_defs_tasks_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_buildbucket_appengine_tasks_defs_tasks_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_buildbucket_appengine_tasks_defs_tasks_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_buildbucket_appengine_tasks_defs_tasks_proto_rawDescData)
	})
	return file_go_chromium_org_luci_buildbucket_appengine_tasks_defs_tasks_proto_rawDescData
}

var file_go_chromium_org_luci_buildbucket_appengine_tasks_defs_tasks_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_go_chromium_org_luci_buildbucket_appengine_tasks_defs_tasks_proto_goTypes = []interface{}{
	(*CancelSwarmingTask)(nil), // 0: taskdefs.CancelSwarmingTask
	(*CreateSwarmingTask)(nil), // 1: taskdefs.CreateSwarmingTask
	(*ExportBigQuery)(nil),     // 2: taskdefs.ExportBigQuery
	(*FinalizeResultDB)(nil),   // 3: taskdefs.FinalizeResultDB
	(*NotifyPubSub)(nil),       // 4: taskdefs.NotifyPubSub
}
var file_go_chromium_org_luci_buildbucket_appengine_tasks_defs_tasks_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_buildbucket_appengine_tasks_defs_tasks_proto_init() }
func file_go_chromium_org_luci_buildbucket_appengine_tasks_defs_tasks_proto_init() {
	if File_go_chromium_org_luci_buildbucket_appengine_tasks_defs_tasks_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_buildbucket_appengine_tasks_defs_tasks_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelSwarmingTask); i {
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
		file_go_chromium_org_luci_buildbucket_appengine_tasks_defs_tasks_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSwarmingTask); i {
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
		file_go_chromium_org_luci_buildbucket_appengine_tasks_defs_tasks_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExportBigQuery); i {
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
		file_go_chromium_org_luci_buildbucket_appengine_tasks_defs_tasks_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FinalizeResultDB); i {
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
		file_go_chromium_org_luci_buildbucket_appengine_tasks_defs_tasks_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyPubSub); i {
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
			RawDescriptor: file_go_chromium_org_luci_buildbucket_appengine_tasks_defs_tasks_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_buildbucket_appengine_tasks_defs_tasks_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_buildbucket_appengine_tasks_defs_tasks_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_buildbucket_appengine_tasks_defs_tasks_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_buildbucket_appengine_tasks_defs_tasks_proto = out.File
	file_go_chromium_org_luci_buildbucket_appengine_tasks_defs_tasks_proto_rawDesc = nil
	file_go_chromium_org_luci_buildbucket_appengine_tasks_defs_tasks_proto_goTypes = nil
	file_go_chromium_org_luci_buildbucket_appengine_tasks_defs_tasks_proto_depIdxs = nil
}
