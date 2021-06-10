// Copyright 2016 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.1
// source: go.chromium.org/luci/logdog/api/endpoints/coordinator/services/v1/tasks.proto

package logdog

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

// ArchiveTask is a task queue task description for the archival of a single
// log stream.
type ArchiveTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the project that this stream is bound to.
	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	// The hash ID of the log stream to archive.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// The TaskQueue Name of this task, used in the taskqueue pipeline.
	// This is required for ACKing.
	TaskName string `protobuf:"bytes,7,opt,name=task_name,json=taskName,proto3" json:"task_name,omitempty"`
	// The realm name (within the project) that the prefix of the stream is
	// associated with.
	Realm string `protobuf:"bytes,8,opt,name=realm,proto3" json:"realm,omitempty"`
	// TODO(hinoka): Remove this after crbug.com/923557
	// The archival key of the log stream. If this key doesn't match the key in
	// the log stream state, the request is superfluous and should be deleted.
	Key []byte `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *ArchiveTask) Reset() {
	*x = ArchiveTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_logdog_api_endpoints_coordinator_services_v1_tasks_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArchiveTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArchiveTask) ProtoMessage() {}

func (x *ArchiveTask) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_logdog_api_endpoints_coordinator_services_v1_tasks_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArchiveTask.ProtoReflect.Descriptor instead.
func (*ArchiveTask) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_logdog_api_endpoints_coordinator_services_v1_tasks_proto_rawDescGZIP(), []int{0}
}

func (x *ArchiveTask) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *ArchiveTask) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ArchiveTask) GetTaskName() string {
	if x != nil {
		return x.TaskName
	}
	return ""
}

func (x *ArchiveTask) GetRealm() string {
	if x != nil {
		return x.Realm
	}
	return ""
}

func (x *ArchiveTask) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

var File_go_chromium_org_luci_logdog_api_endpoints_coordinator_services_v1_tasks_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_logdog_api_endpoints_coordinator_services_v1_tasks_proto_rawDesc = []byte{
	0x0a, 0x4d, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x6c, 0x6f, 0x67, 0x64, 0x6f, 0x67, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2f, 0x63, 0x6f, 0x6f, 0x72,
	0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x6c, 0x6f, 0x67, 0x64, 0x6f, 0x67, 0x22, 0x7c, 0x0a, 0x0b, 0x41, 0x72, 0x63, 0x68, 0x69,
	0x76, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65,
	0x61, 0x6c, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x42, 0x4a, 0x5a, 0x48, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f,
	0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x6c, 0x6f,
	0x67, 0x64, 0x6f, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x73, 0x2f, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x6c, 0x6f, 0x67, 0x64, 0x6f,
	0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_logdog_api_endpoints_coordinator_services_v1_tasks_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_logdog_api_endpoints_coordinator_services_v1_tasks_proto_rawDescData = file_go_chromium_org_luci_logdog_api_endpoints_coordinator_services_v1_tasks_proto_rawDesc
)

func file_go_chromium_org_luci_logdog_api_endpoints_coordinator_services_v1_tasks_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_logdog_api_endpoints_coordinator_services_v1_tasks_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_logdog_api_endpoints_coordinator_services_v1_tasks_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_logdog_api_endpoints_coordinator_services_v1_tasks_proto_rawDescData)
	})
	return file_go_chromium_org_luci_logdog_api_endpoints_coordinator_services_v1_tasks_proto_rawDescData
}

var file_go_chromium_org_luci_logdog_api_endpoints_coordinator_services_v1_tasks_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_go_chromium_org_luci_logdog_api_endpoints_coordinator_services_v1_tasks_proto_goTypes = []interface{}{
	(*ArchiveTask)(nil), // 0: logdog.ArchiveTask
}
var file_go_chromium_org_luci_logdog_api_endpoints_coordinator_services_v1_tasks_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_go_chromium_org_luci_logdog_api_endpoints_coordinator_services_v1_tasks_proto_init()
}
func file_go_chromium_org_luci_logdog_api_endpoints_coordinator_services_v1_tasks_proto_init() {
	if File_go_chromium_org_luci_logdog_api_endpoints_coordinator_services_v1_tasks_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_logdog_api_endpoints_coordinator_services_v1_tasks_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArchiveTask); i {
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
			RawDescriptor: file_go_chromium_org_luci_logdog_api_endpoints_coordinator_services_v1_tasks_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_logdog_api_endpoints_coordinator_services_v1_tasks_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_logdog_api_endpoints_coordinator_services_v1_tasks_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_logdog_api_endpoints_coordinator_services_v1_tasks_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_logdog_api_endpoints_coordinator_services_v1_tasks_proto = out.File
	file_go_chromium_org_luci_logdog_api_endpoints_coordinator_services_v1_tasks_proto_rawDesc = nil
	file_go_chromium_org_luci_logdog_api_endpoints_coordinator_services_v1_tasks_proto_goTypes = nil
	file_go_chromium_org_luci_logdog_api_endpoints_coordinator_services_v1_tasks_proto_depIdxs = nil
}
