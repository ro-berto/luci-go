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
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: go.chromium.org/luci/cv/internal/gerrit/poller/storage.proto

package poller

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// QueryState represents execution of a single Gerrit query over time.
//
// Exactly one of (or_projects, common_project_prefix) must be specified.
// Not using oneof to avoid wrapping or_projects in a message as oneof doesn't
// support repeated fields.
type QueryState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Host is Gerrit host.
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	// Enumerated Gerrit projects, e.g. ["infra/infra", "infra/luci/luci-go"].
	OrProjects []string `protobuf:"bytes,2,rep,name=or_projects,json=orProjects,proto3" json:"or_projects,omitempty"`
	// Common Gerrit project prefix, e.g. "chromiumos/".
	CommonProjectPrefix string `protobuf:"bytes,3,opt,name=common_project_prefix,json=commonProjectPrefix,proto3" json:"common_project_prefix,omitempty"`
	// When the last full poll was started.
	LastFullTime *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=last_full_time,json=lastFullTime,proto3" json:"last_full_time,omitempty"`
	// When the last incremental poll was started.
	LastIncrTime *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=last_incr_time,json=lastIncrTime,proto3" json:"last_incr_time,omitempty"`
	// Changes are changes which were last observed by the query execution.
	//
	// These are not CL IDs, but Gerrit change numbers.
	//
	// The full poll resets these.
	// The incremental poll adds newly discovered CLs.
	//
	// Sorted.
	Changes []int64 `protobuf:"varint,13,rep,packed,name=changes,proto3" json:"changes,omitempty"`
}

func (x *QueryState) Reset() {
	*x = QueryState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_internal_gerrit_poller_storage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryState) ProtoMessage() {}

func (x *QueryState) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_internal_gerrit_poller_storage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryState.ProtoReflect.Descriptor instead.
func (*QueryState) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_internal_gerrit_poller_storage_proto_rawDescGZIP(), []int{0}
}

func (x *QueryState) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *QueryState) GetOrProjects() []string {
	if x != nil {
		return x.OrProjects
	}
	return nil
}

func (x *QueryState) GetCommonProjectPrefix() string {
	if x != nil {
		return x.CommonProjectPrefix
	}
	return ""
}

func (x *QueryState) GetLastFullTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastFullTime
	}
	return nil
}

func (x *QueryState) GetLastIncrTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastIncrTime
	}
	return nil
}

func (x *QueryState) GetChanges() []int64 {
	if x != nil {
		return x.Changes
	}
	return nil
}

// QueryStates exists to reference several QueryStates as a single property in a
// Datastore entity.
type QueryStates struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	States []*QueryState `protobuf:"bytes,1,rep,name=states,proto3" json:"states,omitempty"`
}

func (x *QueryStates) Reset() {
	*x = QueryStates{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_internal_gerrit_poller_storage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryStates) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryStates) ProtoMessage() {}

func (x *QueryStates) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_internal_gerrit_poller_storage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryStates.ProtoReflect.Descriptor instead.
func (*QueryStates) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_internal_gerrit_poller_storage_proto_rawDescGZIP(), []int{1}
}

func (x *QueryStates) GetStates() []*QueryState {
	if x != nil {
		return x.States
	}
	return nil
}

var File_go_chromium_org_luci_cv_internal_gerrit_poller_storage_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_cv_internal_gerrit_poller_storage_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x76, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x67, 0x65, 0x72, 0x72, 0x69, 0x74, 0x2f, 0x70, 0x6f, 0x6c, 0x6c, 0x65, 0x72,
	0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19,
	0x63, 0x76, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x67, 0x65, 0x72, 0x72,
	0x69, 0x74, 0x2e, 0x70, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x02, 0x0a, 0x0a, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x6f, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0a, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x32,
	0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x12, 0x40, 0x0a, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x66, 0x75, 0x6c, 0x6c, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x46, 0x75, 0x6c, 0x6c,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x69, 0x6e, 0x63,
	0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x49, 0x6e,
	0x63, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x03, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73,
	0x22, 0x4c, 0x0a, 0x0b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x12,
	0x3d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x25, 0x2e, 0x63, 0x76, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x67, 0x65,
	0x72, 0x72, 0x69, 0x74, 0x2e, 0x70, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x42, 0x37,
	0x5a, 0x35, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x76, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x67, 0x65, 0x72, 0x72, 0x69, 0x74, 0x2f, 0x70, 0x6f, 0x6c, 0x6c, 0x65, 0x72,
	0x3b, 0x70, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_cv_internal_gerrit_poller_storage_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_cv_internal_gerrit_poller_storage_proto_rawDescData = file_go_chromium_org_luci_cv_internal_gerrit_poller_storage_proto_rawDesc
)

func file_go_chromium_org_luci_cv_internal_gerrit_poller_storage_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_cv_internal_gerrit_poller_storage_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_cv_internal_gerrit_poller_storage_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_cv_internal_gerrit_poller_storage_proto_rawDescData)
	})
	return file_go_chromium_org_luci_cv_internal_gerrit_poller_storage_proto_rawDescData
}

var file_go_chromium_org_luci_cv_internal_gerrit_poller_storage_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_go_chromium_org_luci_cv_internal_gerrit_poller_storage_proto_goTypes = []interface{}{
	(*QueryState)(nil),            // 0: cv.internal.gerrit.poller.QueryState
	(*QueryStates)(nil),           // 1: cv.internal.gerrit.poller.QueryStates
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_go_chromium_org_luci_cv_internal_gerrit_poller_storage_proto_depIdxs = []int32{
	2, // 0: cv.internal.gerrit.poller.QueryState.last_full_time:type_name -> google.protobuf.Timestamp
	2, // 1: cv.internal.gerrit.poller.QueryState.last_incr_time:type_name -> google.protobuf.Timestamp
	0, // 2: cv.internal.gerrit.poller.QueryStates.states:type_name -> cv.internal.gerrit.poller.QueryState
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_cv_internal_gerrit_poller_storage_proto_init() }
func file_go_chromium_org_luci_cv_internal_gerrit_poller_storage_proto_init() {
	if File_go_chromium_org_luci_cv_internal_gerrit_poller_storage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_cv_internal_gerrit_poller_storage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryState); i {
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
		file_go_chromium_org_luci_cv_internal_gerrit_poller_storage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryStates); i {
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
			RawDescriptor: file_go_chromium_org_luci_cv_internal_gerrit_poller_storage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_cv_internal_gerrit_poller_storage_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_cv_internal_gerrit_poller_storage_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_cv_internal_gerrit_poller_storage_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_cv_internal_gerrit_poller_storage_proto = out.File
	file_go_chromium_org_luci_cv_internal_gerrit_poller_storage_proto_rawDesc = nil
	file_go_chromium_org_luci_cv_internal_gerrit_poller_storage_proto_goTypes = nil
	file_go_chromium_org_luci_cv_internal_gerrit_poller_storage_proto_depIdxs = nil
}
