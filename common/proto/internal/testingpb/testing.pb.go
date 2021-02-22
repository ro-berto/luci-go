// Copyright 2018 The Swarming Authors. All rights reserved.
// Use of this source code is governed by the Apache v2.0 license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.1
// source: go.chromium.org/luci/common/proto/internal/testingpb/testing.proto

package testingpb

import (
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Some struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	I int64 `protobuf:"varint,1,opt,name=i,proto3" json:"i,omitempty"`
}

func (x *Some) Reset() {
	*x = Some{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_common_proto_internal_testingpb_testing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Some) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Some) ProtoMessage() {}

func (x *Some) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_common_proto_internal_testingpb_testing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Some.ProtoReflect.Descriptor instead.
func (*Some) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_common_proto_internal_testingpb_testing_proto_rawDescGZIP(), []int{0}
}

func (x *Some) GetI() int64 {
	if x != nil {
		return x.I
	}
	return 0
}

type Simple struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Some   *Some                 `protobuf:"bytes,2,opt,name=some,proto3" json:"some,omitempty"`
	Fields *field_mask.FieldMask `protobuf:"bytes,100,opt,name=fields,proto3" json:"fields,omitempty"`
}

func (x *Simple) Reset() {
	*x = Simple{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_common_proto_internal_testingpb_testing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Simple) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Simple) ProtoMessage() {}

func (x *Simple) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_common_proto_internal_testingpb_testing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Simple.ProtoReflect.Descriptor instead.
func (*Simple) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_common_proto_internal_testingpb_testing_proto_rawDescGZIP(), []int{1}
}

func (x *Simple) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Simple) GetSome() *Some {
	if x != nil {
		return x.Some
	}
	return nil
}

func (x *Simple) GetFields() *field_mask.FieldMask {
	if x != nil {
		return x.Fields
	}
	return nil
}

type Props struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Properties *structpb.Struct      `protobuf:"bytes,6,opt,name=properties,proto3" json:"properties,omitempty"`
	Fields     *field_mask.FieldMask `protobuf:"bytes,100,opt,name=fields,proto3" json:"fields,omitempty"`
}

func (x *Props) Reset() {
	*x = Props{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_common_proto_internal_testingpb_testing_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Props) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Props) ProtoMessage() {}

func (x *Props) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_common_proto_internal_testingpb_testing_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Props.ProtoReflect.Descriptor instead.
func (*Props) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_common_proto_internal_testingpb_testing_proto_rawDescGZIP(), []int{2}
}

func (x *Props) GetProperties() *structpb.Struct {
	if x != nil {
		return x.Properties
	}
	return nil
}

func (x *Props) GetFields() *field_mask.FieldMask {
	if x != nil {
		return x.Fields
	}
	return nil
}

type WithInner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msgs []*WithInner_Inner `protobuf:"bytes,1,rep,name=msgs,proto3" json:"msgs,omitempty"`
}

func (x *WithInner) Reset() {
	*x = WithInner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_common_proto_internal_testingpb_testing_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WithInner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WithInner) ProtoMessage() {}

func (x *WithInner) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_common_proto_internal_testingpb_testing_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WithInner.ProtoReflect.Descriptor instead.
func (*WithInner) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_common_proto_internal_testingpb_testing_proto_rawDescGZIP(), []int{3}
}

func (x *WithInner) GetMsgs() []*WithInner_Inner {
	if x != nil {
		return x.Msgs
	}
	return nil
}

type Full struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num            int32            `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
	Nums           []int32          `protobuf:"varint,2,rep,packed,name=nums,proto3" json:"nums,omitempty"`
	Str            string           `protobuf:"bytes,3,opt,name=str,proto3" json:"str,omitempty"`
	Strs           []string         `protobuf:"bytes,4,rep,name=strs,proto3" json:"strs,omitempty"`
	Msg            *Full            `protobuf:"bytes,5,opt,name=msg,proto3" json:"msg,omitempty"`
	Msgs           []*Full          `protobuf:"bytes,6,rep,name=msgs,proto3" json:"msgs,omitempty"`
	MapStrNum      map[string]int32 `protobuf:"bytes,7,rep,name=map_str_num,json=mapStrNum,proto3" json:"map_str_num,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	MapNumStr      map[int32]string `protobuf:"bytes,8,rep,name=map_num_str,json=mapNumStr,proto3" json:"map_num_str,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	MapBoolStr     map[bool]string  `protobuf:"bytes,9,rep,name=map_bool_str,json=mapBoolStr,proto3" json:"map_bool_str,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	MapStrMsg      map[string]*Full `protobuf:"bytes,10,rep,name=map_str_msg,json=mapStrMsg,proto3" json:"map_str_msg,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	JsonName       string           `protobuf:"bytes,11,opt,name=json_name,json=jsonName,proto3" json:"json_name,omitempty"`
	JsonNameOption string           `protobuf:"bytes,12,opt,name=json_name_option,json=another_json_name,proto3" json:"json_name_option,omitempty"`
}

func (x *Full) Reset() {
	*x = Full{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_common_proto_internal_testingpb_testing_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Full) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Full) ProtoMessage() {}

func (x *Full) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_common_proto_internal_testingpb_testing_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Full.ProtoReflect.Descriptor instead.
func (*Full) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_common_proto_internal_testingpb_testing_proto_rawDescGZIP(), []int{4}
}

func (x *Full) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *Full) GetNums() []int32 {
	if x != nil {
		return x.Nums
	}
	return nil
}

func (x *Full) GetStr() string {
	if x != nil {
		return x.Str
	}
	return ""
}

func (x *Full) GetStrs() []string {
	if x != nil {
		return x.Strs
	}
	return nil
}

func (x *Full) GetMsg() *Full {
	if x != nil {
		return x.Msg
	}
	return nil
}

func (x *Full) GetMsgs() []*Full {
	if x != nil {
		return x.Msgs
	}
	return nil
}

func (x *Full) GetMapStrNum() map[string]int32 {
	if x != nil {
		return x.MapStrNum
	}
	return nil
}

func (x *Full) GetMapNumStr() map[int32]string {
	if x != nil {
		return x.MapNumStr
	}
	return nil
}

func (x *Full) GetMapBoolStr() map[bool]string {
	if x != nil {
		return x.MapBoolStr
	}
	return nil
}

func (x *Full) GetMapStrMsg() map[string]*Full {
	if x != nil {
		return x.MapStrMsg
	}
	return nil
}

func (x *Full) GetJsonName() string {
	if x != nil {
		return x.JsonName
	}
	return ""
}

func (x *Full) GetJsonNameOption() string {
	if x != nil {
		return x.JsonNameOption
	}
	return ""
}

type WithInner_Inner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Msg:
	//	*WithInner_Inner_Simple
	//	*WithInner_Inner_Props
	Msg isWithInner_Inner_Msg `protobuf_oneof:"msg"`
}

func (x *WithInner_Inner) Reset() {
	*x = WithInner_Inner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_common_proto_internal_testingpb_testing_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WithInner_Inner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WithInner_Inner) ProtoMessage() {}

func (x *WithInner_Inner) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_common_proto_internal_testingpb_testing_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WithInner_Inner.ProtoReflect.Descriptor instead.
func (*WithInner_Inner) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_common_proto_internal_testingpb_testing_proto_rawDescGZIP(), []int{3, 0}
}

func (m *WithInner_Inner) GetMsg() isWithInner_Inner_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (x *WithInner_Inner) GetSimple() *Simple {
	if x, ok := x.GetMsg().(*WithInner_Inner_Simple); ok {
		return x.Simple
	}
	return nil
}

func (x *WithInner_Inner) GetProps() *Props {
	if x, ok := x.GetMsg().(*WithInner_Inner_Props); ok {
		return x.Props
	}
	return nil
}

type isWithInner_Inner_Msg interface {
	isWithInner_Inner_Msg()
}

type WithInner_Inner_Simple struct {
	Simple *Simple `protobuf:"bytes,1,opt,name=simple,proto3,oneof"`
}

type WithInner_Inner_Props struct {
	Props *Props `protobuf:"bytes,2,opt,name=props,proto3,oneof"`
}

func (*WithInner_Inner_Simple) isWithInner_Inner_Msg() {}

func (*WithInner_Inner_Props) isWithInner_Inner_Msg() {}

var File_go_chromium_org_luci_common_proto_internal_testingpb_testing_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_common_proto_internal_testingpb_testing_proto_rawDesc = []byte{
	0x0a, 0x42, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61,
	0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x14, 0x0a, 0x04, 0x53, 0x6f, 0x6d, 0x65, 0x12, 0x0c,
	0x0a, 0x01, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x69, 0x22, 0x78, 0x0a, 0x06,
	0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2a, 0x0a, 0x04, 0x73, 0x6f, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x6f, 0x6d, 0x65, 0x52, 0x04, 0x73, 0x6f,
	0x6d, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x64, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x06,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x22, 0x74, 0x0a, 0x05, 0x50, 0x72, 0x6f, 0x70, 0x73, 0x12,
	0x37, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0a, 0x70, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x32, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x4d, 0x61, 0x73, 0x6b, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x22, 0xb7, 0x01, 0x0a,
	0x09, 0x57, 0x69, 0x74, 0x68, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x35, 0x0a, 0x04, 0x6d, 0x73,
	0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x57, 0x69, 0x74, 0x68,
	0x49, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x04, 0x6d, 0x73, 0x67,
	0x73, 0x1a, 0x73, 0x0a, 0x05, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x06, 0x73, 0x69,
	0x6d, 0x70, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x69,
	0x6d, 0x70, 0x6c, 0x65, 0x48, 0x00, 0x52, 0x06, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x2f,
	0x0a, 0x05, 0x70, 0x72, 0x6f, 0x70, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x50, 0x72, 0x6f, 0x70, 0x73, 0x48, 0x00, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x70, 0x73, 0x42,
	0x05, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x22, 0xa2, 0x06, 0x0a, 0x04, 0x46, 0x75, 0x6c, 0x6c, 0x12,
	0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75,
	0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x75, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52,
	0x04, 0x6e, 0x75, 0x6d, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x74, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x73, 0x74, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x74, 0x72, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x73, 0x74, 0x72, 0x73, 0x12, 0x28, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x46, 0x75, 0x6c, 0x6c,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x2a, 0x0a, 0x04, 0x6d, 0x73, 0x67, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x46, 0x75, 0x6c, 0x6c, 0x52, 0x04, 0x6d, 0x73, 0x67,
	0x73, 0x12, 0x45, 0x0a, 0x0b, 0x6d, 0x61, 0x70, 0x5f, 0x73, 0x74, 0x72, 0x5f, 0x6e, 0x75, 0x6d,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x46, 0x75, 0x6c, 0x6c, 0x2e, 0x4d,
	0x61, 0x70, 0x53, 0x74, 0x72, 0x4e, 0x75, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x6d,
	0x61, 0x70, 0x53, 0x74, 0x72, 0x4e, 0x75, 0x6d, 0x12, 0x45, 0x0a, 0x0b, 0x6d, 0x61, 0x70, 0x5f,
	0x6e, 0x75, 0x6d, 0x5f, 0x73, 0x74, 0x72, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x46, 0x75, 0x6c, 0x6c, 0x2e, 0x4d, 0x61, 0x70, 0x4e, 0x75, 0x6d, 0x53, 0x74, 0x72, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x6d, 0x61, 0x70, 0x4e, 0x75, 0x6d, 0x53, 0x74, 0x72, 0x12,
	0x48, 0x0a, 0x0c, 0x6d, 0x61, 0x70, 0x5f, 0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x73, 0x74, 0x72, 0x18,
	0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x46, 0x75, 0x6c, 0x6c, 0x2e, 0x4d, 0x61,
	0x70, 0x42, 0x6f, 0x6f, 0x6c, 0x53, 0x74, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x6d,
	0x61, 0x70, 0x42, 0x6f, 0x6f, 0x6c, 0x53, 0x74, 0x72, 0x12, 0x45, 0x0a, 0x0b, 0x6d, 0x61, 0x70,
	0x5f, 0x73, 0x74, 0x72, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x46, 0x75, 0x6c, 0x6c, 0x2e, 0x4d, 0x61, 0x70, 0x53, 0x74, 0x72, 0x4d, 0x73, 0x67,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x6d, 0x61, 0x70, 0x53, 0x74, 0x72, 0x4d, 0x73, 0x67,
	0x12, 0x1b, 0x0a, 0x09, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6a, 0x73, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a,
	0x10, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x61, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72,
	0x5f, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x3c, 0x0a, 0x0e, 0x4d, 0x61,
	0x70, 0x53, 0x74, 0x72, 0x4e, 0x75, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3c, 0x0a, 0x0e, 0x4d, 0x61, 0x70, 0x4e,
	0x75, 0x6d, 0x53, 0x74, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3d, 0x0a, 0x0f, 0x4d, 0x61, 0x70, 0x42, 0x6f, 0x6f,
	0x6c, 0x53, 0x74, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x54, 0x0a, 0x0e, 0x4d, 0x61, 0x70, 0x53, 0x74, 0x72, 0x4d,
	0x73, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x46, 0x75, 0x6c, 0x6c,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x40, 0x5a, 0x3e, 0x67,
	0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c,
	0x75, 0x63, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x70, 0x62, 0x3b, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_common_proto_internal_testingpb_testing_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_common_proto_internal_testingpb_testing_proto_rawDescData = file_go_chromium_org_luci_common_proto_internal_testingpb_testing_proto_rawDesc
)

func file_go_chromium_org_luci_common_proto_internal_testingpb_testing_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_common_proto_internal_testingpb_testing_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_common_proto_internal_testingpb_testing_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_common_proto_internal_testingpb_testing_proto_rawDescData)
	})
	return file_go_chromium_org_luci_common_proto_internal_testingpb_testing_proto_rawDescData
}

var file_go_chromium_org_luci_common_proto_internal_testingpb_testing_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_go_chromium_org_luci_common_proto_internal_testingpb_testing_proto_goTypes = []interface{}{
	(*Some)(nil),                 // 0: internal.testing.Some
	(*Simple)(nil),               // 1: internal.testing.Simple
	(*Props)(nil),                // 2: internal.testing.Props
	(*WithInner)(nil),            // 3: internal.testing.WithInner
	(*Full)(nil),                 // 4: internal.testing.Full
	(*WithInner_Inner)(nil),      // 5: internal.testing.WithInner.Inner
	nil,                          // 6: internal.testing.Full.MapStrNumEntry
	nil,                          // 7: internal.testing.Full.MapNumStrEntry
	nil,                          // 8: internal.testing.Full.MapBoolStrEntry
	nil,                          // 9: internal.testing.Full.MapStrMsgEntry
	(*field_mask.FieldMask)(nil), // 10: google.protobuf.FieldMask
	(*structpb.Struct)(nil),      // 11: google.protobuf.Struct
}
var file_go_chromium_org_luci_common_proto_internal_testingpb_testing_proto_depIdxs = []int32{
	0,  // 0: internal.testing.Simple.some:type_name -> internal.testing.Some
	10, // 1: internal.testing.Simple.fields:type_name -> google.protobuf.FieldMask
	11, // 2: internal.testing.Props.properties:type_name -> google.protobuf.Struct
	10, // 3: internal.testing.Props.fields:type_name -> google.protobuf.FieldMask
	5,  // 4: internal.testing.WithInner.msgs:type_name -> internal.testing.WithInner.Inner
	4,  // 5: internal.testing.Full.msg:type_name -> internal.testing.Full
	4,  // 6: internal.testing.Full.msgs:type_name -> internal.testing.Full
	6,  // 7: internal.testing.Full.map_str_num:type_name -> internal.testing.Full.MapStrNumEntry
	7,  // 8: internal.testing.Full.map_num_str:type_name -> internal.testing.Full.MapNumStrEntry
	8,  // 9: internal.testing.Full.map_bool_str:type_name -> internal.testing.Full.MapBoolStrEntry
	9,  // 10: internal.testing.Full.map_str_msg:type_name -> internal.testing.Full.MapStrMsgEntry
	1,  // 11: internal.testing.WithInner.Inner.simple:type_name -> internal.testing.Simple
	2,  // 12: internal.testing.WithInner.Inner.props:type_name -> internal.testing.Props
	4,  // 13: internal.testing.Full.MapStrMsgEntry.value:type_name -> internal.testing.Full
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_common_proto_internal_testingpb_testing_proto_init() }
func file_go_chromium_org_luci_common_proto_internal_testingpb_testing_proto_init() {
	if File_go_chromium_org_luci_common_proto_internal_testingpb_testing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_common_proto_internal_testingpb_testing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Some); i {
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
		file_go_chromium_org_luci_common_proto_internal_testingpb_testing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Simple); i {
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
		file_go_chromium_org_luci_common_proto_internal_testingpb_testing_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Props); i {
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
		file_go_chromium_org_luci_common_proto_internal_testingpb_testing_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WithInner); i {
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
		file_go_chromium_org_luci_common_proto_internal_testingpb_testing_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Full); i {
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
		file_go_chromium_org_luci_common_proto_internal_testingpb_testing_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WithInner_Inner); i {
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
	file_go_chromium_org_luci_common_proto_internal_testingpb_testing_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*WithInner_Inner_Simple)(nil),
		(*WithInner_Inner_Props)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_go_chromium_org_luci_common_proto_internal_testingpb_testing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_common_proto_internal_testingpb_testing_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_common_proto_internal_testingpb_testing_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_common_proto_internal_testingpb_testing_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_common_proto_internal_testingpb_testing_proto = out.File
	file_go_chromium_org_luci_common_proto_internal_testingpb_testing_proto_rawDesc = nil
	file_go_chromium_org_luci_common_proto_internal_testingpb_testing_proto_goTypes = nil
	file_go_chromium_org_luci_common_proto_internal_testingpb_testing_proto_depIdxs = nil
}
