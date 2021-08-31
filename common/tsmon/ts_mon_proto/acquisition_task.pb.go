// Copyright 2015 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: go.chromium.org/luci/common/tsmon/ts_mon_proto/acquisition_task.proto

package ts_mon_proto

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

type Task_TypeId int32

const (
	Task_MESSAGE_TYPE_ID Task_TypeId = 34049749
)

// Enum value maps for Task_TypeId.
var (
	Task_TypeId_name = map[int32]string{
		34049749: "MESSAGE_TYPE_ID",
	}
	Task_TypeId_value = map[string]int32{
		"MESSAGE_TYPE_ID": 34049749,
	}
)

func (x Task_TypeId) Enum() *Task_TypeId {
	p := new(Task_TypeId)
	*p = x
	return p
}

func (x Task_TypeId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Task_TypeId) Descriptor() protoreflect.EnumDescriptor {
	return file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_task_proto_enumTypes[0].Descriptor()
}

func (Task_TypeId) Type() protoreflect.EnumType {
	return &file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_task_proto_enumTypes[0]
}

func (x Task_TypeId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Task_TypeId) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Task_TypeId(num)
	return nil
}

// Deprecated: Use Task_TypeId.Descriptor instead.
func (Task_TypeId) EnumDescriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_task_proto_rawDescGZIP(), []int{0, 0}
}

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProxyEnvironment *string `protobuf:"bytes,5,opt,name=proxy_environment,json=proxyEnvironment" json:"proxy_environment,omitempty"`
	AcquisitionName  *string `protobuf:"bytes,10,opt,name=acquisition_name,json=acquisitionName" json:"acquisition_name,omitempty"`
	ServiceName      *string `protobuf:"bytes,20,opt,name=service_name,json=serviceName" json:"service_name,omitempty"`
	JobName          *string `protobuf:"bytes,30,opt,name=job_name,json=jobName" json:"job_name,omitempty"`
	DataCenter       *string `protobuf:"bytes,40,opt,name=data_center,json=dataCenter" json:"data_center,omitempty"`
	HostName         *string `protobuf:"bytes,50,opt,name=host_name,json=hostName" json:"host_name,omitempty"`
	TaskNum          *int32  `protobuf:"varint,60,opt,name=task_num,json=taskNum" json:"task_num,omitempty"`
	ProxyZone        *string `protobuf:"bytes,70,opt,name=proxy_zone,json=proxyZone" json:"proxy_zone,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_task_proto_rawDescGZIP(), []int{0}
}

func (x *Task) GetProxyEnvironment() string {
	if x != nil && x.ProxyEnvironment != nil {
		return *x.ProxyEnvironment
	}
	return ""
}

func (x *Task) GetAcquisitionName() string {
	if x != nil && x.AcquisitionName != nil {
		return *x.AcquisitionName
	}
	return ""
}

func (x *Task) GetServiceName() string {
	if x != nil && x.ServiceName != nil {
		return *x.ServiceName
	}
	return ""
}

func (x *Task) GetJobName() string {
	if x != nil && x.JobName != nil {
		return *x.JobName
	}
	return ""
}

func (x *Task) GetDataCenter() string {
	if x != nil && x.DataCenter != nil {
		return *x.DataCenter
	}
	return ""
}

func (x *Task) GetHostName() string {
	if x != nil && x.HostName != nil {
		return *x.HostName
	}
	return ""
}

func (x *Task) GetTaskNum() int32 {
	if x != nil && x.TaskNum != nil {
		return *x.TaskNum
	}
	return 0
}

func (x *Task) GetProxyZone() string {
	if x != nil && x.ProxyZone != nil {
		return *x.ProxyZone
	}
	return ""
}

var File_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_task_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_task_proto_rawDesc = []byte{
	0x0a, 0x45, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x73,
	0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x73, 0x5f, 0x6d, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x61, 0x63, 0x71, 0x75, 0x69, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x61, 0x73,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x74, 0x73, 0x5f, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb6, 0x02, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x2b,
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x61,
	0x63, 0x71, 0x75, 0x69, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x63, 0x71, 0x75, 0x69, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6a, 0x6f, 0x62,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6a, 0x6f, 0x62,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x43,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x3c,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x4e, 0x75, 0x6d, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x46, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5a, 0x6f, 0x6e, 0x65, 0x22, 0x20, 0x0a, 0x06,
	0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x0f, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47,
	0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x44, 0x10, 0xd5, 0x9d, 0x9e, 0x10, 0x42, 0x30,
	0x5a, 0x2e, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x73,
	0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x73, 0x5f, 0x6d, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_task_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_task_proto_rawDescData = file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_task_proto_rawDesc
)

func file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_task_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_task_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_task_proto_rawDescData)
	})
	return file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_task_proto_rawDescData
}

var file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_task_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_task_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_task_proto_goTypes = []interface{}{
	(Task_TypeId)(0), // 0: ts_mon.proto.Task.TypeId
	(*Task)(nil),     // 1: ts_mon.proto.Task
}
var file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_task_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_task_proto_init() }
func file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_task_proto_init() {
	if File_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_task_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
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
			RawDescriptor: file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_task_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_task_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_task_proto_depIdxs,
		EnumInfos:         file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_task_proto_enumTypes,
		MessageInfos:      file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_task_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_task_proto = out.File
	file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_task_proto_rawDesc = nil
	file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_task_proto_goTypes = nil
	file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_task_proto_depIdxs = nil
}
