// Copyright 2022 The LUCI Authors.
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

// vim: noexpandtab

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: go.chromium.org/luci/common/proto/reflectutil/reflectutil_test.proto

package reflectutil

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

type TestMapMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BoolMap   map[bool]string   `protobuf:"bytes,1,rep,name=bool_map,json=boolMap,proto3" json:"bool_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Int32Map  map[int32]string  `protobuf:"bytes,2,rep,name=int32_map,json=int32Map,proto3" json:"int32_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Int64Map  map[int64]string  `protobuf:"bytes,3,rep,name=int64_map,json=int64Map,proto3" json:"int64_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Uint32Map map[uint32]string `protobuf:"bytes,4,rep,name=uint32_map,json=uint32Map,proto3" json:"uint32_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Uint64Map map[uint64]string `protobuf:"bytes,5,rep,name=uint64_map,json=uint64Map,proto3" json:"uint64_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	StringMap map[string]string `protobuf:"bytes,6,rep,name=string_map,json=stringMap,proto3" json:"string_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *TestMapMessage) Reset() {
	*x = TestMapMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_common_proto_reflectutil_reflectutil_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestMapMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestMapMessage) ProtoMessage() {}

func (x *TestMapMessage) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_common_proto_reflectutil_reflectutil_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestMapMessage.ProtoReflect.Descriptor instead.
func (*TestMapMessage) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_common_proto_reflectutil_reflectutil_test_proto_rawDescGZIP(), []int{0}
}

func (x *TestMapMessage) GetBoolMap() map[bool]string {
	if x != nil {
		return x.BoolMap
	}
	return nil
}

func (x *TestMapMessage) GetInt32Map() map[int32]string {
	if x != nil {
		return x.Int32Map
	}
	return nil
}

func (x *TestMapMessage) GetInt64Map() map[int64]string {
	if x != nil {
		return x.Int64Map
	}
	return nil
}

func (x *TestMapMessage) GetUint32Map() map[uint32]string {
	if x != nil {
		return x.Uint32Map
	}
	return nil
}

func (x *TestMapMessage) GetUint64Map() map[uint64]string {
	if x != nil {
		return x.Uint64Map
	}
	return nil
}

func (x *TestMapMessage) GetStringMap() map[string]string {
	if x != nil {
		return x.StringMap
	}
	return nil
}

type TestPathMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SingleInner *TestPathMessage_Inner            `protobuf:"bytes,1,opt,name=single_inner,json=singleInner,proto3" json:"single_inner,omitempty"`
	MapInner    map[string]*TestPathMessage_Inner `protobuf:"bytes,2,rep,name=map_inner,json=mapInner,proto3" json:"map_inner,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	IntMapInner map[int32]*TestPathMessage_Inner  `protobuf:"bytes,3,rep,name=int_map_inner,json=intMapInner,proto3" json:"int_map_inner,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	MultiInner  []*TestPathMessage_Inner          `protobuf:"bytes,4,rep,name=multi_inner,json=multiInner,proto3" json:"multi_inner,omitempty"`
}

func (x *TestPathMessage) Reset() {
	*x = TestPathMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_common_proto_reflectutil_reflectutil_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestPathMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestPathMessage) ProtoMessage() {}

func (x *TestPathMessage) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_common_proto_reflectutil_reflectutil_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestPathMessage.ProtoReflect.Descriptor instead.
func (*TestPathMessage) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_common_proto_reflectutil_reflectutil_test_proto_rawDescGZIP(), []int{1}
}

func (x *TestPathMessage) GetSingleInner() *TestPathMessage_Inner {
	if x != nil {
		return x.SingleInner
	}
	return nil
}

func (x *TestPathMessage) GetMapInner() map[string]*TestPathMessage_Inner {
	if x != nil {
		return x.MapInner
	}
	return nil
}

func (x *TestPathMessage) GetIntMapInner() map[int32]*TestPathMessage_Inner {
	if x != nil {
		return x.IntMapInner
	}
	return nil
}

func (x *TestPathMessage) GetMultiInner() []*TestPathMessage_Inner {
	if x != nil {
		return x.MultiInner
	}
	return nil
}

type TestShallowCopyMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field         string                                   `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	RepeatedField []string                                 `protobuf:"bytes,2,rep,name=repeated_field,json=repeatedField,proto3" json:"repeated_field,omitempty"`
	MappedField   map[string]string                        `protobuf:"bytes,3,rep,name=mapped_field,json=mappedField,proto3" json:"mapped_field,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	InnerMsg      *TestShallowCopyMessage_Inner            `protobuf:"bytes,4,opt,name=inner_msg,json=innerMsg,proto3" json:"inner_msg,omitempty"`
	RepeatedMsg   []*TestShallowCopyMessage_Inner          `protobuf:"bytes,5,rep,name=repeated_msg,json=repeatedMsg,proto3" json:"repeated_msg,omitempty"`
	MappedMsg     map[string]*TestShallowCopyMessage_Inner `protobuf:"bytes,6,rep,name=mapped_msg,json=mappedMsg,proto3" json:"mapped_msg,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *TestShallowCopyMessage) Reset() {
	*x = TestShallowCopyMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_common_proto_reflectutil_reflectutil_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestShallowCopyMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestShallowCopyMessage) ProtoMessage() {}

func (x *TestShallowCopyMessage) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_common_proto_reflectutil_reflectutil_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestShallowCopyMessage.ProtoReflect.Descriptor instead.
func (*TestShallowCopyMessage) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_common_proto_reflectutil_reflectutil_test_proto_rawDescGZIP(), []int{2}
}

func (x *TestShallowCopyMessage) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *TestShallowCopyMessage) GetRepeatedField() []string {
	if x != nil {
		return x.RepeatedField
	}
	return nil
}

func (x *TestShallowCopyMessage) GetMappedField() map[string]string {
	if x != nil {
		return x.MappedField
	}
	return nil
}

func (x *TestShallowCopyMessage) GetInnerMsg() *TestShallowCopyMessage_Inner {
	if x != nil {
		return x.InnerMsg
	}
	return nil
}

func (x *TestShallowCopyMessage) GetRepeatedMsg() []*TestShallowCopyMessage_Inner {
	if x != nil {
		return x.RepeatedMsg
	}
	return nil
}

func (x *TestShallowCopyMessage) GetMappedMsg() map[string]*TestShallowCopyMessage_Inner {
	if x != nil {
		return x.MappedMsg
	}
	return nil
}

type TestPathMessage_Inner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Str string `protobuf:"bytes,1,opt,name=str,proto3" json:"str,omitempty"`
}

func (x *TestPathMessage_Inner) Reset() {
	*x = TestPathMessage_Inner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_common_proto_reflectutil_reflectutil_test_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestPathMessage_Inner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestPathMessage_Inner) ProtoMessage() {}

func (x *TestPathMessage_Inner) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_common_proto_reflectutil_reflectutil_test_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestPathMessage_Inner.ProtoReflect.Descriptor instead.
func (*TestPathMessage_Inner) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_common_proto_reflectutil_reflectutil_test_proto_rawDescGZIP(), []int{1, 0}
}

func (x *TestPathMessage_Inner) GetStr() string {
	if x != nil {
		return x.Str
	}
	return ""
}

type TestShallowCopyMessage_Inner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field string `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
}

func (x *TestShallowCopyMessage_Inner) Reset() {
	*x = TestShallowCopyMessage_Inner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_common_proto_reflectutil_reflectutil_test_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestShallowCopyMessage_Inner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestShallowCopyMessage_Inner) ProtoMessage() {}

func (x *TestShallowCopyMessage_Inner) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_common_proto_reflectutil_reflectutil_test_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestShallowCopyMessage_Inner.ProtoReflect.Descriptor instead.
func (*TestShallowCopyMessage_Inner) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_common_proto_reflectutil_reflectutil_test_proto_rawDescGZIP(), []int{2, 0}
}

func (x *TestShallowCopyMessage_Inner) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

var File_go_chromium_org_luci_common_proto_reflectutil_reflectutil_test_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_common_proto_reflectutil_reflectutil_test_proto_rawDesc = []byte{
	0x0a, 0x44, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x74, 0x69, 0x6c, 0x2f,
	0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x74, 0x69, 0x6c, 0x5f, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x75,
	0x74, 0x69, 0x6c, 0x22, 0xb6, 0x06, 0x0a, 0x0e, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x61, 0x70, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x43, 0x0a, 0x08, 0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x6d,
	0x61, 0x70, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x72, 0x65, 0x66, 0x6c, 0x65,
	0x63, 0x74, 0x75, 0x74, 0x69, 0x6c, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x61, 0x70, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x07, 0x62, 0x6f, 0x6f, 0x6c, 0x4d, 0x61, 0x70, 0x12, 0x46, 0x0a, 0x09, 0x69,
	0x6e, 0x74, 0x33, 0x32, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29,
	0x2e, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x74, 0x69, 0x6c, 0x2e, 0x54, 0x65, 0x73,
	0x74, 0x4d, 0x61, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x49, 0x6e, 0x74, 0x33,
	0x32, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x33, 0x32,
	0x4d, 0x61, 0x70, 0x12, 0x46, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x6d, 0x61, 0x70,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74,
	0x75, 0x74, 0x69, 0x6c, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x61, 0x70, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x4d, 0x61, 0x70, 0x12, 0x49, 0x0a, 0x0a, 0x75,
	0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2a, 0x2e, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x74, 0x69, 0x6c, 0x2e, 0x54, 0x65,
	0x73, 0x74, 0x4d, 0x61, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x55, 0x69, 0x6e,
	0x74, 0x33, 0x32, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x75, 0x69, 0x6e,
	0x74, 0x33, 0x32, 0x4d, 0x61, 0x70, 0x12, 0x49, 0x0a, 0x0a, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34,
	0x5f, 0x6d, 0x61, 0x70, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x72, 0x65, 0x66,
	0x6c, 0x65, 0x63, 0x74, 0x75, 0x74, 0x69, 0x6c, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x61, 0x70,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x55, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x4d, 0x61,
	0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x4d, 0x61,
	0x70, 0x12, 0x49, 0x0a, 0x0a, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x61, 0x70, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x75,
	0x74, 0x69, 0x6c, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x61, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x09, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x70, 0x1a, 0x3a, 0x0a, 0x0c,
	0x42, 0x6f, 0x6f, 0x6c, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3b, 0x0a, 0x0d, 0x49, 0x6e, 0x74, 0x33,
	0x32, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3b, 0x0a, 0x0d, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x4d, 0x61,
	0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x1a, 0x3c, 0x0a, 0x0e, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x4d, 0x61, 0x70, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x1a, 0x3c, 0x0a, 0x0e, 0x55, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3c,
	0x0a, 0x0e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x99, 0x04, 0x0a,
	0x0f, 0x54, 0x65, 0x73, 0x74, 0x50, 0x61, 0x74, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x45, 0x0a, 0x0c, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x6e, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74,
	0x75, 0x74, 0x69, 0x6c, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x50, 0x61, 0x74, 0x68, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x0b, 0x73, 0x69, 0x6e, 0x67,
	0x6c, 0x65, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x47, 0x0a, 0x09, 0x6d, 0x61, 0x70, 0x5f, 0x69,
	0x6e, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x72, 0x65, 0x66,
	0x6c, 0x65, 0x63, 0x74, 0x75, 0x74, 0x69, 0x6c, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x50, 0x61, 0x74,
	0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x61, 0x70, 0x49, 0x6e, 0x6e, 0x65,
	0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x61, 0x70, 0x49, 0x6e, 0x6e, 0x65, 0x72,
	0x12, 0x51, 0x0a, 0x0d, 0x69, 0x6e, 0x74, 0x5f, 0x6d, 0x61, 0x70, 0x5f, 0x69, 0x6e, 0x6e, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63,
	0x74, 0x75, 0x74, 0x69, 0x6c, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x50, 0x61, 0x74, 0x68, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x49, 0x6e, 0x74, 0x4d, 0x61, 0x70, 0x49, 0x6e, 0x6e, 0x65,
	0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x69, 0x6e, 0x74, 0x4d, 0x61, 0x70, 0x49, 0x6e,
	0x6e, 0x65, 0x72, 0x12, 0x43, 0x0a, 0x0b, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x69, 0x6e, 0x6e,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x72, 0x65, 0x66, 0x6c, 0x65,
	0x63, 0x74, 0x75, 0x74, 0x69, 0x6c, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x50, 0x61, 0x74, 0x68, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x0a, 0x6d, 0x75,
	0x6c, 0x74, 0x69, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x1a, 0x19, 0x0a, 0x05, 0x49, 0x6e, 0x6e, 0x65,
	0x72, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x74, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x73, 0x74, 0x72, 0x1a, 0x5f, 0x0a, 0x0d, 0x4d, 0x61, 0x70, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x38, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x75,
	0x74, 0x69, 0x6c, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x50, 0x61, 0x74, 0x68, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x1a, 0x62, 0x0a, 0x10, 0x49, 0x6e, 0x74, 0x4d, 0x61, 0x70, 0x49, 0x6e,
	0x6e, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x38, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x72, 0x65, 0x66, 0x6c,
	0x65, 0x63, 0x74, 0x75, 0x74, 0x69, 0x6c, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x50, 0x61, 0x74, 0x68,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xdf, 0x04, 0x0a, 0x16, 0x54, 0x65, 0x73,
	0x74, 0x53, 0x68, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x70, 0x79, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x70,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0d, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x12, 0x57, 0x0a, 0x0c, 0x6d, 0x61, 0x70, 0x70, 0x65, 0x64, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74,
	0x75, 0x74, 0x69, 0x6c, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x53, 0x68, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x43, 0x6f, 0x70, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x61, 0x70, 0x70,
	0x65, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x6d, 0x61,
	0x70, 0x70, 0x65, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x46, 0x0a, 0x09, 0x69, 0x6e, 0x6e,
	0x65, 0x72, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x72,
	0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x74, 0x69, 0x6c, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x53,
	0x68, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x70, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x08, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x4d, 0x73,
	0x67, 0x12, 0x4c, 0x0a, 0x0c, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x6d, 0x73,
	0x67, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63,
	0x74, 0x75, 0x74, 0x69, 0x6c, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x53, 0x68, 0x61, 0x6c, 0x6c, 0x6f,
	0x77, 0x43, 0x6f, 0x70, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x49, 0x6e, 0x6e,
	0x65, 0x72, 0x52, 0x0b, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x4d, 0x73, 0x67, 0x12,
	0x51, 0x0a, 0x0a, 0x6d, 0x61, 0x70, 0x70, 0x65, 0x64, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x74, 0x69,
	0x6c, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x53, 0x68, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x70,
	0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x61, 0x70, 0x70, 0x65, 0x64, 0x4d,
	0x73, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x6d, 0x61, 0x70, 0x70, 0x65, 0x64, 0x4d,
	0x73, 0x67, 0x1a, 0x1d, 0x0a, 0x05, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x1a, 0x3e, 0x0a, 0x10, 0x4d, 0x61, 0x70, 0x70, 0x65, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x67, 0x0a, 0x0e, 0x4d, 0x61, 0x70, 0x70, 0x65, 0x64, 0x4d, 0x73, 0x67, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x74,
	0x69, 0x6c, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x53, 0x68, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x43, 0x6f,
	0x70, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x6f,
	0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75,
	0x63, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x74, 0x69, 0x6c, 0x3b, 0x72, 0x65, 0x66, 0x6c,
	0x65, 0x63, 0x74, 0x75, 0x74, 0x69, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_common_proto_reflectutil_reflectutil_test_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_common_proto_reflectutil_reflectutil_test_proto_rawDescData = file_go_chromium_org_luci_common_proto_reflectutil_reflectutil_test_proto_rawDesc
)

func file_go_chromium_org_luci_common_proto_reflectutil_reflectutil_test_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_common_proto_reflectutil_reflectutil_test_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_common_proto_reflectutil_reflectutil_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_common_proto_reflectutil_reflectutil_test_proto_rawDescData)
	})
	return file_go_chromium_org_luci_common_proto_reflectutil_reflectutil_test_proto_rawDescData
}

var file_go_chromium_org_luci_common_proto_reflectutil_reflectutil_test_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_go_chromium_org_luci_common_proto_reflectutil_reflectutil_test_proto_goTypes = []interface{}{
	(*TestMapMessage)(nil),               // 0: reflectutil.TestMapMessage
	(*TestPathMessage)(nil),              // 1: reflectutil.TestPathMessage
	(*TestShallowCopyMessage)(nil),       // 2: reflectutil.TestShallowCopyMessage
	nil,                                  // 3: reflectutil.TestMapMessage.BoolMapEntry
	nil,                                  // 4: reflectutil.TestMapMessage.Int32MapEntry
	nil,                                  // 5: reflectutil.TestMapMessage.Int64MapEntry
	nil,                                  // 6: reflectutil.TestMapMessage.Uint32MapEntry
	nil,                                  // 7: reflectutil.TestMapMessage.Uint64MapEntry
	nil,                                  // 8: reflectutil.TestMapMessage.StringMapEntry
	(*TestPathMessage_Inner)(nil),        // 9: reflectutil.TestPathMessage.Inner
	nil,                                  // 10: reflectutil.TestPathMessage.MapInnerEntry
	nil,                                  // 11: reflectutil.TestPathMessage.IntMapInnerEntry
	(*TestShallowCopyMessage_Inner)(nil), // 12: reflectutil.TestShallowCopyMessage.Inner
	nil,                                  // 13: reflectutil.TestShallowCopyMessage.MappedFieldEntry
	nil,                                  // 14: reflectutil.TestShallowCopyMessage.MappedMsgEntry
}
var file_go_chromium_org_luci_common_proto_reflectutil_reflectutil_test_proto_depIdxs = []int32{
	3,  // 0: reflectutil.TestMapMessage.bool_map:type_name -> reflectutil.TestMapMessage.BoolMapEntry
	4,  // 1: reflectutil.TestMapMessage.int32_map:type_name -> reflectutil.TestMapMessage.Int32MapEntry
	5,  // 2: reflectutil.TestMapMessage.int64_map:type_name -> reflectutil.TestMapMessage.Int64MapEntry
	6,  // 3: reflectutil.TestMapMessage.uint32_map:type_name -> reflectutil.TestMapMessage.Uint32MapEntry
	7,  // 4: reflectutil.TestMapMessage.uint64_map:type_name -> reflectutil.TestMapMessage.Uint64MapEntry
	8,  // 5: reflectutil.TestMapMessage.string_map:type_name -> reflectutil.TestMapMessage.StringMapEntry
	9,  // 6: reflectutil.TestPathMessage.single_inner:type_name -> reflectutil.TestPathMessage.Inner
	10, // 7: reflectutil.TestPathMessage.map_inner:type_name -> reflectutil.TestPathMessage.MapInnerEntry
	11, // 8: reflectutil.TestPathMessage.int_map_inner:type_name -> reflectutil.TestPathMessage.IntMapInnerEntry
	9,  // 9: reflectutil.TestPathMessage.multi_inner:type_name -> reflectutil.TestPathMessage.Inner
	13, // 10: reflectutil.TestShallowCopyMessage.mapped_field:type_name -> reflectutil.TestShallowCopyMessage.MappedFieldEntry
	12, // 11: reflectutil.TestShallowCopyMessage.inner_msg:type_name -> reflectutil.TestShallowCopyMessage.Inner
	12, // 12: reflectutil.TestShallowCopyMessage.repeated_msg:type_name -> reflectutil.TestShallowCopyMessage.Inner
	14, // 13: reflectutil.TestShallowCopyMessage.mapped_msg:type_name -> reflectutil.TestShallowCopyMessage.MappedMsgEntry
	9,  // 14: reflectutil.TestPathMessage.MapInnerEntry.value:type_name -> reflectutil.TestPathMessage.Inner
	9,  // 15: reflectutil.TestPathMessage.IntMapInnerEntry.value:type_name -> reflectutil.TestPathMessage.Inner
	12, // 16: reflectutil.TestShallowCopyMessage.MappedMsgEntry.value:type_name -> reflectutil.TestShallowCopyMessage.Inner
	17, // [17:17] is the sub-list for method output_type
	17, // [17:17] is the sub-list for method input_type
	17, // [17:17] is the sub-list for extension type_name
	17, // [17:17] is the sub-list for extension extendee
	0,  // [0:17] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_common_proto_reflectutil_reflectutil_test_proto_init() }
func file_go_chromium_org_luci_common_proto_reflectutil_reflectutil_test_proto_init() {
	if File_go_chromium_org_luci_common_proto_reflectutil_reflectutil_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_common_proto_reflectutil_reflectutil_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestMapMessage); i {
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
		file_go_chromium_org_luci_common_proto_reflectutil_reflectutil_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestPathMessage); i {
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
		file_go_chromium_org_luci_common_proto_reflectutil_reflectutil_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestShallowCopyMessage); i {
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
		file_go_chromium_org_luci_common_proto_reflectutil_reflectutil_test_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestPathMessage_Inner); i {
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
		file_go_chromium_org_luci_common_proto_reflectutil_reflectutil_test_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestShallowCopyMessage_Inner); i {
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
			RawDescriptor: file_go_chromium_org_luci_common_proto_reflectutil_reflectutil_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_common_proto_reflectutil_reflectutil_test_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_common_proto_reflectutil_reflectutil_test_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_common_proto_reflectutil_reflectutil_test_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_common_proto_reflectutil_reflectutil_test_proto = out.File
	file_go_chromium_org_luci_common_proto_reflectutil_reflectutil_test_proto_rawDesc = nil
	file_go_chromium_org_luci_common_proto_reflectutil_reflectutil_test_proto_goTypes = nil
	file_go_chromium_org_luci_common_proto_reflectutil_reflectutil_test_proto_depIdxs = nil
}
