// Copyright 2017 The LUCI Authors.
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
// source: go.chromium.org/luci/scheduler/appengine/task/gitiles/pb/messages.proto

package pb

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

// Child is the last part and its sha1 tip.
type Child struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Suffix string `protobuf:"bytes,1,opt,name=suffix,proto3" json:"suffix,omitempty"`
	Sha1   []byte `protobuf:"bytes,2,opt,name=sha1,proto3" json:"sha1,omitempty"`
}

func (x *Child) Reset() {
	*x = Child{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_scheduler_appengine_task_gitiles_pb_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Child) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Child) ProtoMessage() {}

func (x *Child) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_scheduler_appengine_task_gitiles_pb_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Child.ProtoReflect.Descriptor instead.
func (*Child) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_scheduler_appengine_task_gitiles_pb_messages_proto_rawDescGZIP(), []int{0}
}

func (x *Child) GetSuffix() string {
	if x != nil {
		return x.Suffix
	}
	return ""
}

func (x *Child) GetSha1() []byte {
	if x != nil {
		return x.Sha1
	}
	return nil
}

// RefSpace is a bunch of children which share the same ref namespace (prefix).
type RefSpace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Prefix   string   `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	Children []*Child `protobuf:"bytes,2,rep,name=children,proto3" json:"children,omitempty"`
}

func (x *RefSpace) Reset() {
	*x = RefSpace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_scheduler_appengine_task_gitiles_pb_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefSpace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefSpace) ProtoMessage() {}

func (x *RefSpace) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_scheduler_appengine_task_gitiles_pb_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefSpace.ProtoReflect.Descriptor instead.
func (*RefSpace) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_scheduler_appengine_task_gitiles_pb_messages_proto_rawDescGZIP(), []int{1}
}

func (x *RefSpace) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

func (x *RefSpace) GetChildren() []*Child {
	if x != nil {
		return x.Children
	}
	return nil
}

// RepositoryState stores tips of all watched refs in a repo.
type RepositoryState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Spaces []*RefSpace `protobuf:"bytes,1,rep,name=spaces,proto3" json:"spaces,omitempty"`
}

func (x *RepositoryState) Reset() {
	*x = RepositoryState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_scheduler_appengine_task_gitiles_pb_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepositoryState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepositoryState) ProtoMessage() {}

func (x *RepositoryState) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_scheduler_appengine_task_gitiles_pb_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepositoryState.ProtoReflect.Descriptor instead.
func (*RepositoryState) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_scheduler_appengine_task_gitiles_pb_messages_proto_rawDescGZIP(), []int{2}
}

func (x *RepositoryState) GetSpaces() []*RefSpace {
	if x != nil {
		return x.Spaces
	}
	return nil
}

// DebugState is returned as part of GetDebugJobState RPC response.
type DebugState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Known   []*DebugState_Ref `protobuf:"bytes,1,rep,name=known,proto3" json:"known,omitempty"`     // tips as stored in the datastore
	Current []*DebugState_Ref `protobuf:"bytes,2,rep,name=current,proto3" json:"current,omitempty"` // tips as returned by gitiles right now
}

func (x *DebugState) Reset() {
	*x = DebugState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_scheduler_appengine_task_gitiles_pb_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DebugState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DebugState) ProtoMessage() {}

func (x *DebugState) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_scheduler_appengine_task_gitiles_pb_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DebugState.ProtoReflect.Descriptor instead.
func (*DebugState) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_scheduler_appengine_task_gitiles_pb_messages_proto_rawDescGZIP(), []int{3}
}

func (x *DebugState) GetKnown() []*DebugState_Ref {
	if x != nil {
		return x.Known
	}
	return nil
}

func (x *DebugState) GetCurrent() []*DebugState_Ref {
	if x != nil {
		return x.Current
	}
	return nil
}

type DebugState_Ref struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ref    string `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`
	Commit string `protobuf:"bytes,2,opt,name=commit,proto3" json:"commit,omitempty"`
}

func (x *DebugState_Ref) Reset() {
	*x = DebugState_Ref{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_scheduler_appengine_task_gitiles_pb_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DebugState_Ref) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DebugState_Ref) ProtoMessage() {}

func (x *DebugState_Ref) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_scheduler_appengine_task_gitiles_pb_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DebugState_Ref.ProtoReflect.Descriptor instead.
func (*DebugState_Ref) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_scheduler_appengine_task_gitiles_pb_messages_proto_rawDescGZIP(), []int{3, 0}
}

func (x *DebugState_Ref) GetRef() string {
	if x != nil {
		return x.Ref
	}
	return ""
}

func (x *DebugState_Ref) GetCommit() string {
	if x != nil {
		return x.Commit
	}
	return ""
}

var File_go_chromium_org_luci_scheduler_appengine_task_gitiles_pb_messages_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_scheduler_appengine_task_gitiles_pb_messages_proto_rawDesc = []byte{
	0x0a, 0x47, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72,
	0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f,
	0x67, 0x69, 0x74, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x67, 0x69, 0x74, 0x69, 0x6c,
	0x65, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0x33, 0x0a, 0x05, 0x43,
	0x68, 0x69, 0x6c, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x68, 0x61, 0x31, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x73, 0x68, 0x61, 0x31,
	0x22, 0x57, 0x0a, 0x08, 0x52, 0x65, 0x66, 0x53, 0x70, 0x61, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x12, 0x33, 0x0a, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x69, 0x74, 0x69, 0x6c, 0x65, 0x73,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x52,
	0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x22, 0x45, 0x0a, 0x0f, 0x52, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x32, 0x0a, 0x06,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x69, 0x74, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x52, 0x65, 0x66, 0x53, 0x70, 0x61, 0x63, 0x65, 0x52, 0x06, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73,
	0x22, 0xb1, 0x01, 0x0a, 0x0a, 0x44, 0x65, 0x62, 0x75, 0x67, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x36, 0x0a, 0x05, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x67, 0x69, 0x74, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x2e, 0x44, 0x65, 0x62, 0x75, 0x67, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x52, 0x65, 0x66,
	0x52, 0x05, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x12, 0x3a, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x69, 0x74, 0x69, 0x6c,
	0x65, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x62, 0x75,
	0x67, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x52, 0x65, 0x66, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x1a, 0x2f, 0x0a, 0x03, 0x52, 0x65, 0x66, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65,
	0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72, 0x65, 0x66, 0x12, 0x16, 0x0a, 0x06,
	0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f,
	0x6d, 0x6d, 0x69, 0x74, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d,
	0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x67, 0x69, 0x74, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_scheduler_appengine_task_gitiles_pb_messages_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_scheduler_appengine_task_gitiles_pb_messages_proto_rawDescData = file_go_chromium_org_luci_scheduler_appengine_task_gitiles_pb_messages_proto_rawDesc
)

func file_go_chromium_org_luci_scheduler_appengine_task_gitiles_pb_messages_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_scheduler_appengine_task_gitiles_pb_messages_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_scheduler_appengine_task_gitiles_pb_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_scheduler_appengine_task_gitiles_pb_messages_proto_rawDescData)
	})
	return file_go_chromium_org_luci_scheduler_appengine_task_gitiles_pb_messages_proto_rawDescData
}

var file_go_chromium_org_luci_scheduler_appengine_task_gitiles_pb_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_go_chromium_org_luci_scheduler_appengine_task_gitiles_pb_messages_proto_goTypes = []interface{}{
	(*Child)(nil),           // 0: gitiles.messages.Child
	(*RefSpace)(nil),        // 1: gitiles.messages.RefSpace
	(*RepositoryState)(nil), // 2: gitiles.messages.RepositoryState
	(*DebugState)(nil),      // 3: gitiles.messages.DebugState
	(*DebugState_Ref)(nil),  // 4: gitiles.messages.DebugState.Ref
}
var file_go_chromium_org_luci_scheduler_appengine_task_gitiles_pb_messages_proto_depIdxs = []int32{
	0, // 0: gitiles.messages.RefSpace.children:type_name -> gitiles.messages.Child
	1, // 1: gitiles.messages.RepositoryState.spaces:type_name -> gitiles.messages.RefSpace
	4, // 2: gitiles.messages.DebugState.known:type_name -> gitiles.messages.DebugState.Ref
	4, // 3: gitiles.messages.DebugState.current:type_name -> gitiles.messages.DebugState.Ref
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_scheduler_appengine_task_gitiles_pb_messages_proto_init() }
func file_go_chromium_org_luci_scheduler_appengine_task_gitiles_pb_messages_proto_init() {
	if File_go_chromium_org_luci_scheduler_appengine_task_gitiles_pb_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_scheduler_appengine_task_gitiles_pb_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Child); i {
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
		file_go_chromium_org_luci_scheduler_appengine_task_gitiles_pb_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefSpace); i {
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
		file_go_chromium_org_luci_scheduler_appengine_task_gitiles_pb_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepositoryState); i {
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
		file_go_chromium_org_luci_scheduler_appengine_task_gitiles_pb_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DebugState); i {
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
		file_go_chromium_org_luci_scheduler_appengine_task_gitiles_pb_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DebugState_Ref); i {
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
			RawDescriptor: file_go_chromium_org_luci_scheduler_appengine_task_gitiles_pb_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_scheduler_appengine_task_gitiles_pb_messages_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_scheduler_appengine_task_gitiles_pb_messages_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_scheduler_appengine_task_gitiles_pb_messages_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_scheduler_appengine_task_gitiles_pb_messages_proto = out.File
	file_go_chromium_org_luci_scheduler_appengine_task_gitiles_pb_messages_proto_rawDesc = nil
	file_go_chromium_org_luci_scheduler_appengine_task_gitiles_pb_messages_proto_goTypes = nil
	file_go_chromium_org_luci_scheduler_appengine_task_gitiles_pb_messages_proto_depIdxs = nil
}
