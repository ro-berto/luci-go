// Copyright 2021 The LUCI Authors.
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
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: go.chromium.org/luci/cv/api/v1/pubsub.proto

package cvpb

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

// PubSubRun includes the high-level information about the CV Run sent via
// PubSub.
//
// This includes a subset of the fields defined in Run message.
// Use "runs.GetRun" rpc to retrieve the full field set of Runs.
type PubSubRun struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unqiue ID of the Run.
	//
	// The format of an ID is "projects/$luci-project/runs/$id", where
	// - luci-project is the name of the LUCI project the Run belongs to
	// - id is an opaque key unique in the LUCI project.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// status of the Run.
	Status Run_Status `protobuf:"varint,2,opt,name=status,proto3,enum=cv.v1.Run_Status" json:"status,omitempty"`
	// eversion is the entity version, which is monotonically increasing.
	Eversion int64 `protobuf:"varint,3,opt,name=eversion,proto3" json:"eversion,omitempty"`
}

func (x *PubSubRun) Reset() {
	*x = PubSubRun{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_api_v1_pubsub_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PubSubRun) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PubSubRun) ProtoMessage() {}

func (x *PubSubRun) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_api_v1_pubsub_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PubSubRun.ProtoReflect.Descriptor instead.
func (*PubSubRun) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_api_v1_pubsub_proto_rawDescGZIP(), []int{0}
}

func (x *PubSubRun) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PubSubRun) GetStatus() Run_Status {
	if x != nil {
		return x.Status
	}
	return Run_STATUS_UNSPECIFIED
}

func (x *PubSubRun) GetEversion() int64 {
	if x != nil {
		return x.Eversion
	}
	return 0
}

var File_go_chromium_org_luci_cv_api_v1_pubsub_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_cv_api_v1_pubsub_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x76, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x63,
	0x76, 0x2e, 0x76, 0x31, 0x1a, 0x28, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75,
	0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x76, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x75, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x62,
	0x0a, 0x09, 0x50, 0x75, 0x62, 0x53, 0x75, 0x62, 0x52, 0x75, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x29, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x63, 0x76,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x65, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75,
	0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x76, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x76, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_go_chromium_org_luci_cv_api_v1_pubsub_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_cv_api_v1_pubsub_proto_rawDescData = file_go_chromium_org_luci_cv_api_v1_pubsub_proto_rawDesc
)

func file_go_chromium_org_luci_cv_api_v1_pubsub_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_cv_api_v1_pubsub_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_cv_api_v1_pubsub_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_cv_api_v1_pubsub_proto_rawDescData)
	})
	return file_go_chromium_org_luci_cv_api_v1_pubsub_proto_rawDescData
}

var file_go_chromium_org_luci_cv_api_v1_pubsub_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_go_chromium_org_luci_cv_api_v1_pubsub_proto_goTypes = []interface{}{
	(*PubSubRun)(nil), // 0: cv.v1.PubSubRun
	(Run_Status)(0),   // 1: cv.v1.Run.Status
}
var file_go_chromium_org_luci_cv_api_v1_pubsub_proto_depIdxs = []int32{
	1, // 0: cv.v1.PubSubRun.status:type_name -> cv.v1.Run.Status
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_cv_api_v1_pubsub_proto_init() }
func file_go_chromium_org_luci_cv_api_v1_pubsub_proto_init() {
	if File_go_chromium_org_luci_cv_api_v1_pubsub_proto != nil {
		return
	}
	file_go_chromium_org_luci_cv_api_v1_run_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_cv_api_v1_pubsub_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PubSubRun); i {
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
			RawDescriptor: file_go_chromium_org_luci_cv_api_v1_pubsub_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_cv_api_v1_pubsub_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_cv_api_v1_pubsub_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_cv_api_v1_pubsub_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_cv_api_v1_pubsub_proto = out.File
	file_go_chromium_org_luci_cv_api_v1_pubsub_proto_rawDesc = nil
	file_go_chromium_org_luci_cv_api_v1_pubsub_proto_goTypes = nil
	file_go_chromium_org_luci_cv_api_v1_pubsub_proto_depIdxs = nil
}
