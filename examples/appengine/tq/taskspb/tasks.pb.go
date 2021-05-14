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
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.0
// source: go.chromium.org/luci/examples/appengine/tq/taskspb/tasks.proto

package taskspb

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

type CountDownTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int64 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *CountDownTask) Reset() {
	*x = CountDownTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_examples_appengine_tq_taskspb_tasks_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountDownTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountDownTask) ProtoMessage() {}

func (x *CountDownTask) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_examples_appengine_tq_taskspb_tasks_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountDownTask.ProtoReflect.Descriptor instead.
func (*CountDownTask) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_examples_appengine_tq_taskspb_tasks_proto_rawDescGZIP(), []int{0}
}

func (x *CountDownTask) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

var File_go_chromium_org_luci_examples_appengine_tq_taskspb_tasks_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_examples_appengine_tq_taskspb_tasks_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f,
	0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x74, 0x71, 0x2f, 0x74, 0x61, 0x73,
	0x6b, 0x73, 0x70, 0x62, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x22, 0x6c, 0x75, 0x63, 0x69, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e,
	0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x74, 0x71, 0x2e, 0x74, 0x61, 0x73,
	0x6b, 0x73, 0x70, 0x62, 0x22, 0x27, 0x0a, 0x0d, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x6f, 0x77,
	0x6e, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x34, 0x5a,
	0x32, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x61,
	0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x74, 0x71, 0x2f, 0x74, 0x61, 0x73, 0x6b,
	0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_examples_appengine_tq_taskspb_tasks_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_examples_appengine_tq_taskspb_tasks_proto_rawDescData = file_go_chromium_org_luci_examples_appengine_tq_taskspb_tasks_proto_rawDesc
)

func file_go_chromium_org_luci_examples_appengine_tq_taskspb_tasks_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_examples_appengine_tq_taskspb_tasks_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_examples_appengine_tq_taskspb_tasks_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_examples_appengine_tq_taskspb_tasks_proto_rawDescData)
	})
	return file_go_chromium_org_luci_examples_appengine_tq_taskspb_tasks_proto_rawDescData
}

var file_go_chromium_org_luci_examples_appengine_tq_taskspb_tasks_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_go_chromium_org_luci_examples_appengine_tq_taskspb_tasks_proto_goTypes = []interface{}{
	(*CountDownTask)(nil), // 0: luci.examples.appengine.tq.taskspb.CountDownTask
}
var file_go_chromium_org_luci_examples_appengine_tq_taskspb_tasks_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_examples_appengine_tq_taskspb_tasks_proto_init() }
func file_go_chromium_org_luci_examples_appengine_tq_taskspb_tasks_proto_init() {
	if File_go_chromium_org_luci_examples_appengine_tq_taskspb_tasks_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_examples_appengine_tq_taskspb_tasks_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountDownTask); i {
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
			RawDescriptor: file_go_chromium_org_luci_examples_appengine_tq_taskspb_tasks_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_examples_appengine_tq_taskspb_tasks_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_examples_appengine_tq_taskspb_tasks_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_examples_appengine_tq_taskspb_tasks_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_examples_appengine_tq_taskspb_tasks_proto = out.File
	file_go_chromium_org_luci_examples_appengine_tq_taskspb_tasks_proto_rawDesc = nil
	file_go_chromium_org_luci_examples_appengine_tq_taskspb_tasks_proto_goTypes = nil
	file_go_chromium_org_luci_examples_appengine_tq_taskspb_tasks_proto_depIdxs = nil
}
