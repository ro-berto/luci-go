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
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: go.chromium.org/luci/server/tq/internal/tqpb/tasks.proto

package tqpb

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

// SweepTask is used to distribute sweeping work items and options when doing
// distributed sweeps.
//
// All fields are required.
type SweepTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// DB is an identifier of the database used for reminders.
	//
	// It must be registered in the process that does the sweeping.
	Db string `protobuf:"bytes,1,opt,name=db,proto3" json:"db,omitempty"`
	// Partition specifies the range of keys to scan as [Low..High).
	//
	// It is a string of the form "<hex-low>_<hex-high>".
	Partition string `protobuf:"bytes,2,opt,name=partition,proto3" json:"partition,omitempty"`
	// Identifies a registered Lessor implementation to grab leases through.
	LessorId string `protobuf:"bytes,3,opt,name=lessor_id,json=lessorId,proto3" json:"lessor_id,omitempty"`
	// Identifier of a lease section ID to grab leases on sub-partitions through.
	LeaseSectionId string `protobuf:"bytes,4,opt,name=lease_section_id,json=leaseSectionId,proto3" json:"lease_section_id,omitempty"`
	// Total number of shards used when this task was generated.
	//
	// Used FYI only in logs.
	ShardCount int32 `protobuf:"varint,5,opt,name=shard_count,json=shardCount,proto3" json:"shard_count,omitempty"`
	// Shard number in the range of [0 .. Shards).
	//
	// Used FYI only in logs.
	ShardIndex int32 `protobuf:"varint,6,opt,name=shard_index,json=shardIndex,proto3" json:"shard_index,omitempty"`
	// Level counts recursion level for monitoring/debugging purposes.
	//
	// The root sweeper triggers tasks at level=0. If there is a big backlog,
	// level=0 task will offload some work to level=1 tasks. level > 1 should not
	// normally happen and indicates either a bug or a very overloaded system.
	//
	// level > 2 won't be executed at all.
	Level int32 `protobuf:"varint,7,opt,name=level,proto3" json:"level,omitempty"`
	// Length of the reminder keys in the partition.
	//
	// Used to figure out the upper bound of the scan. Usually 16.
	KeySpaceBytes int32 `protobuf:"varint,8,opt,name=key_space_bytes,json=keySpaceBytes,proto3" json:"key_space_bytes,omitempty"`
	// Caps maximum number of reminders to process.
	//
	// Usually in hundreds.
	TasksPerScan int32 `protobuf:"varint,9,opt,name=tasks_per_scan,json=tasksPerScan,proto3" json:"tasks_per_scan,omitempty"`
	// Caps the number of follow-up scans.
	//
	// Usually 16.
	SecondaryScanShards int32 `protobuf:"varint,10,opt,name=secondary_scan_shards,json=secondaryScanShards,proto3" json:"secondary_scan_shards,omitempty"`
}

func (x *SweepTask) Reset() {
	*x = SweepTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_server_tq_internal_tqpb_tasks_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SweepTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SweepTask) ProtoMessage() {}

func (x *SweepTask) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_server_tq_internal_tqpb_tasks_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SweepTask.ProtoReflect.Descriptor instead.
func (*SweepTask) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_server_tq_internal_tqpb_tasks_proto_rawDescGZIP(), []int{0}
}

func (x *SweepTask) GetDb() string {
	if x != nil {
		return x.Db
	}
	return ""
}

func (x *SweepTask) GetPartition() string {
	if x != nil {
		return x.Partition
	}
	return ""
}

func (x *SweepTask) GetLessorId() string {
	if x != nil {
		return x.LessorId
	}
	return ""
}

func (x *SweepTask) GetLeaseSectionId() string {
	if x != nil {
		return x.LeaseSectionId
	}
	return ""
}

func (x *SweepTask) GetShardCount() int32 {
	if x != nil {
		return x.ShardCount
	}
	return 0
}

func (x *SweepTask) GetShardIndex() int32 {
	if x != nil {
		return x.ShardIndex
	}
	return 0
}

func (x *SweepTask) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *SweepTask) GetKeySpaceBytes() int32 {
	if x != nil {
		return x.KeySpaceBytes
	}
	return 0
}

func (x *SweepTask) GetTasksPerScan() int32 {
	if x != nil {
		return x.TasksPerScan
	}
	return 0
}

func (x *SweepTask) GetSecondaryScanShards() int32 {
	if x != nil {
		return x.SecondaryScanShards
	}
	return 0
}

var File_go_chromium_org_luci_server_tq_internal_tqpb_tasks_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_server_tq_internal_tqpb_tasks_proto_rawDesc = []byte{
	0x0a, 0x38, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x74, 0x71,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x71, 0x70, 0x62, 0x2f, 0x74,
	0x61, 0x73, 0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x6c, 0x75, 0x63, 0x69,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x74, 0x71, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x22, 0xda, 0x02, 0x0a, 0x09, 0x53, 0x77, 0x65, 0x65, 0x70, 0x54, 0x61, 0x73,
	0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x64, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x64,
	0x62, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1b, 0x0a, 0x09, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x53, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x68, 0x61, 0x72, 0x64, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x68, 0x61,
	0x72, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x68, 0x61, 0x72, 0x64,
	0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x68,
	0x61, 0x72, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x26,
	0x0a, 0x0f, 0x6b, 0x65, 0x79, 0x5f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x62, 0x79, 0x74, 0x65,
	0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x6b, 0x65, 0x79, 0x53, 0x70, 0x61, 0x63,
	0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x0e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x5f,
	0x70, 0x65, 0x72, 0x5f, 0x73, 0x63, 0x61, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c,
	0x74, 0x61, 0x73, 0x6b, 0x73, 0x50, 0x65, 0x72, 0x53, 0x63, 0x61, 0x6e, 0x12, 0x32, 0x0a, 0x15,
	0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x5f, 0x73, 0x63, 0x61, 0x6e, 0x5f, 0x73,
	0x68, 0x61, 0x72, 0x64, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x13, 0x73, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x53, 0x63, 0x61, 0x6e, 0x53, 0x68, 0x61, 0x72, 0x64, 0x73,
	0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e,
	0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f,
	0x74, 0x71, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x71, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_server_tq_internal_tqpb_tasks_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_server_tq_internal_tqpb_tasks_proto_rawDescData = file_go_chromium_org_luci_server_tq_internal_tqpb_tasks_proto_rawDesc
)

func file_go_chromium_org_luci_server_tq_internal_tqpb_tasks_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_server_tq_internal_tqpb_tasks_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_server_tq_internal_tqpb_tasks_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_server_tq_internal_tqpb_tasks_proto_rawDescData)
	})
	return file_go_chromium_org_luci_server_tq_internal_tqpb_tasks_proto_rawDescData
}

var file_go_chromium_org_luci_server_tq_internal_tqpb_tasks_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_go_chromium_org_luci_server_tq_internal_tqpb_tasks_proto_goTypes = []interface{}{
	(*SweepTask)(nil), // 0: luci.server.tq.internal.SweepTask
}
var file_go_chromium_org_luci_server_tq_internal_tqpb_tasks_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_server_tq_internal_tqpb_tasks_proto_init() }
func file_go_chromium_org_luci_server_tq_internal_tqpb_tasks_proto_init() {
	if File_go_chromium_org_luci_server_tq_internal_tqpb_tasks_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_server_tq_internal_tqpb_tasks_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SweepTask); i {
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
			RawDescriptor: file_go_chromium_org_luci_server_tq_internal_tqpb_tasks_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_server_tq_internal_tqpb_tasks_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_server_tq_internal_tqpb_tasks_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_server_tq_internal_tqpb_tasks_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_server_tq_internal_tqpb_tasks_proto = out.File
	file_go_chromium_org_luci_server_tq_internal_tqpb_tasks_proto_rawDesc = nil
	file_go_chromium_org_luci_server_tq_internal_tqpb_tasks_proto_goTypes = nil
	file_go_chromium_org_luci_server_tq_internal_tqpb_tasks_proto_depIdxs = nil
}
