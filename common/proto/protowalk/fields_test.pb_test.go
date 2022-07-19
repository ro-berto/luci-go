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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: go.chromium.org/luci/common/proto/protowalk/fields_test.proto

package protowalk

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CustomExt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MustEqual string `protobuf:"bytes,1,opt,name=must_equal,json=mustEqual,proto3" json:"must_equal,omitempty"`
}

func (x *CustomExt) Reset() {
	*x = CustomExt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_common_proto_protowalk_fields_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomExt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomExt) ProtoMessage() {}

func (x *CustomExt) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_common_proto_protowalk_fields_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomExt.ProtoReflect.Descriptor instead.
func (*CustomExt) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_common_proto_protowalk_fields_test_proto_rawDescGZIP(), []int{0}
}

func (x *CustomExt) GetMustEqual() string {
	if x != nil {
		return x.MustEqual
	}
	return ""
}

type Inner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Regular string `protobuf:"bytes,1,opt,name=regular,proto3" json:"regular,omitempty"`
	// Deprecated: Do not use.
	Deprecated  string                     `protobuf:"bytes,2,opt,name=deprecated,proto3" json:"deprecated,omitempty"`
	Output      string                     `protobuf:"bytes,3,opt,name=output,proto3" json:"output,omitempty"`
	Req         string                     `protobuf:"bytes,4,opt,name=req,proto3" json:"req,omitempty"`
	Custom      string                     `protobuf:"bytes,5,opt,name=custom,proto3" json:"custom,omitempty"`
	SingleEmbed *Inner_Embedded            `protobuf:"bytes,6,opt,name=single_embed,json=singleEmbed,proto3" json:"single_embed,omitempty"`
	MultiEmbed  []*Inner_Embedded          `protobuf:"bytes,7,rep,name=multi_embed,json=multiEmbed,proto3" json:"multi_embed,omitempty"`
	MapEmbed    map[string]*Inner_Embedded `protobuf:"bytes,8,rep,name=map_embed,json=mapEmbed,proto3" json:"map_embed,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Inner) Reset() {
	*x = Inner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_common_proto_protowalk_fields_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Inner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Inner) ProtoMessage() {}

func (x *Inner) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_common_proto_protowalk_fields_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Inner.ProtoReflect.Descriptor instead.
func (*Inner) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_common_proto_protowalk_fields_test_proto_rawDescGZIP(), []int{1}
}

func (x *Inner) GetRegular() string {
	if x != nil {
		return x.Regular
	}
	return ""
}

// Deprecated: Do not use.
func (x *Inner) GetDeprecated() string {
	if x != nil {
		return x.Deprecated
	}
	return ""
}

func (x *Inner) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

func (x *Inner) GetReq() string {
	if x != nil {
		return x.Req
	}
	return ""
}

func (x *Inner) GetCustom() string {
	if x != nil {
		return x.Custom
	}
	return ""
}

func (x *Inner) GetSingleEmbed() *Inner_Embedded {
	if x != nil {
		return x.SingleEmbed
	}
	return nil
}

func (x *Inner) GetMultiEmbed() []*Inner_Embedded {
	if x != nil {
		return x.MultiEmbed
	}
	return nil
}

func (x *Inner) GetMapEmbed() map[string]*Inner_Embedded {
	if x != nil {
		return x.MapEmbed
	}
	return nil
}

type Outer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Regular string `protobuf:"bytes,1,opt,name=regular,proto3" json:"regular,omitempty"`
	// Deprecated: Do not use.
	Deprecated  string            `protobuf:"bytes,2,opt,name=deprecated,proto3" json:"deprecated,omitempty"`
	Output      string            `protobuf:"bytes,3,opt,name=output,proto3" json:"output,omitempty"`
	Req         string            `protobuf:"bytes,4,opt,name=req,proto3" json:"req,omitempty"`
	Custom      string            `protobuf:"bytes,5,opt,name=custom,proto3" json:"custom,omitempty"`
	SingleInner *Inner            `protobuf:"bytes,6,opt,name=single_inner,json=singleInner,proto3" json:"single_inner,omitempty"`
	MultiInner  []*Inner          `protobuf:"bytes,7,rep,name=multi_inner,json=multiInner,proto3" json:"multi_inner,omitempty"`
	MapInner    map[string]*Inner `protobuf:"bytes,8,rep,name=map_inner,json=mapInner,proto3" json:"map_inner,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Deprecated: Do not use.
	MultiDeprecated []*Inner         `protobuf:"bytes,9,rep,name=multi_deprecated,json=multiDeprecated,proto3" json:"multi_deprecated,omitempty"`
	IntMapInner     map[int32]*Inner `protobuf:"bytes,10,rep,name=int_map_inner,json=intMapInner,proto3" json:"int_map_inner,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	OutputInner     *Inner           `protobuf:"bytes,11,opt,name=output_inner,json=outputInner,proto3" json:"output_inner,omitempty"`
}

func (x *Outer) Reset() {
	*x = Outer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_common_proto_protowalk_fields_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Outer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Outer) ProtoMessage() {}

func (x *Outer) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_common_proto_protowalk_fields_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Outer.ProtoReflect.Descriptor instead.
func (*Outer) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_common_proto_protowalk_fields_test_proto_rawDescGZIP(), []int{2}
}

func (x *Outer) GetRegular() string {
	if x != nil {
		return x.Regular
	}
	return ""
}

// Deprecated: Do not use.
func (x *Outer) GetDeprecated() string {
	if x != nil {
		return x.Deprecated
	}
	return ""
}

func (x *Outer) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

func (x *Outer) GetReq() string {
	if x != nil {
		return x.Req
	}
	return ""
}

func (x *Outer) GetCustom() string {
	if x != nil {
		return x.Custom
	}
	return ""
}

func (x *Outer) GetSingleInner() *Inner {
	if x != nil {
		return x.SingleInner
	}
	return nil
}

func (x *Outer) GetMultiInner() []*Inner {
	if x != nil {
		return x.MultiInner
	}
	return nil
}

func (x *Outer) GetMapInner() map[string]*Inner {
	if x != nil {
		return x.MapInner
	}
	return nil
}

// Deprecated: Do not use.
func (x *Outer) GetMultiDeprecated() []*Inner {
	if x != nil {
		return x.MultiDeprecated
	}
	return nil
}

func (x *Outer) GetIntMapInner() map[int32]*Inner {
	if x != nil {
		return x.IntMapInner
	}
	return nil
}

func (x *Outer) GetOutputInner() *Inner {
	if x != nil {
		return x.OutputInner
	}
	return nil
}

type Inner_Embedded struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Regular string `protobuf:"bytes,1,opt,name=regular,proto3" json:"regular,omitempty"`
	// Deprecated: Do not use.
	Deprecated string `protobuf:"bytes,2,opt,name=deprecated,proto3" json:"deprecated,omitempty"`
	Output     string `protobuf:"bytes,3,opt,name=output,proto3" json:"output,omitempty"`
	Req        string `protobuf:"bytes,4,opt,name=req,proto3" json:"req,omitempty"`
	Custom     string `protobuf:"bytes,5,opt,name=custom,proto3" json:"custom,omitempty"`
}

func (x *Inner_Embedded) Reset() {
	*x = Inner_Embedded{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_common_proto_protowalk_fields_test_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Inner_Embedded) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Inner_Embedded) ProtoMessage() {}

func (x *Inner_Embedded) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_common_proto_protowalk_fields_test_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Inner_Embedded.ProtoReflect.Descriptor instead.
func (*Inner_Embedded) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_common_proto_protowalk_fields_test_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Inner_Embedded) GetRegular() string {
	if x != nil {
		return x.Regular
	}
	return ""
}

// Deprecated: Do not use.
func (x *Inner_Embedded) GetDeprecated() string {
	if x != nil {
		return x.Deprecated
	}
	return ""
}

func (x *Inner_Embedded) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

func (x *Inner_Embedded) GetReq() string {
	if x != nil {
		return x.Req
	}
	return ""
}

func (x *Inner_Embedded) GetCustom() string {
	if x != nil {
		return x.Custom
	}
	return ""
}

var file_go_chromium_org_luci_common_proto_protowalk_fields_test_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*CustomExt)(nil),
		Field:         50002,
		Name:          "protowalk.custom",
		Tag:           "bytes,50002,opt,name=custom",
		Filename:      "go.chromium.org/luci/common/proto/protowalk/fields_test.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional protowalk.CustomExt custom = 50002;
	E_Custom = &file_go_chromium_org_luci_common_proto_protowalk_fields_test_proto_extTypes[0]
)

var File_go_chromium_org_luci_common_proto_protowalk_fields_test_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_common_proto_protowalk_fields_test_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x77, 0x61, 0x6c, 0x6b, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x77, 0x61, 0x6c, 0x6b, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68,
	0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2a, 0x0a,
	0x09, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x78, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x75,
	0x73, 0x74, 0x5f, 0x65, 0x71, 0x75, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6d, 0x75, 0x73, 0x74, 0x45, 0x71, 0x75, 0x61, 0x6c, 0x22, 0xd1, 0x04, 0x0a, 0x05, 0x49, 0x6e,
	0x6e, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x12, 0x22, 0x0a,
	0x0a, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65,
	0x64, 0x12, 0x1b, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x15,
	0x0a, 0x03, 0x72, 0x65, 0x71, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x03, 0x72, 0x65, 0x71, 0x12, 0x23, 0x0a, 0x06, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0x92, 0xb5, 0x18, 0x07, 0x0a, 0x05, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x52, 0x06, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x12, 0x3c, 0x0a, 0x0c, 0x73, 0x69,
	0x6e, 0x67, 0x6c, 0x65, 0x5f, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x77, 0x61, 0x6c, 0x6b, 0x2e, 0x49, 0x6e, 0x6e,
	0x65, 0x72, 0x2e, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x52, 0x0b, 0x73, 0x69, 0x6e,
	0x67, 0x6c, 0x65, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x12, 0x3a, 0x0a, 0x0b, 0x6d, 0x75, 0x6c, 0x74,
	0x69, 0x5f, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x77, 0x61, 0x6c, 0x6b, 0x2e, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x2e,
	0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x52, 0x0a, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x45,
	0x6d, 0x62, 0x65, 0x64, 0x12, 0x3b, 0x0a, 0x09, 0x6d, 0x61, 0x70, 0x5f, 0x65, 0x6d, 0x62, 0x65,
	0x64, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x77,
	0x61, 0x6c, 0x6b, 0x2e, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x4d, 0x61, 0x70, 0x45, 0x6d, 0x62,
	0x65, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x61, 0x70, 0x45, 0x6d, 0x62, 0x65,
	0x64, 0x1a, 0xa1, 0x01, 0x0a, 0x08, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x72, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x72, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x12, 0x22, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x72,
	0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01,
	0x52, 0x0a, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x06,
	0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x15, 0x0a, 0x03, 0x72, 0x65, 0x71,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x03, 0x72, 0x65, 0x71,
	0x12, 0x23, 0x0a, 0x06, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0b, 0x92, 0xb5, 0x18, 0x07, 0x0a, 0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x06, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x1a, 0x56, 0x0a, 0x0d, 0x4d, 0x61, 0x70, 0x45, 0x6d, 0x62, 0x65,
	0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x77,
	0x61, 0x6c, 0x6b, 0x2e, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64,
	0x65, 0x64, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xa6, 0x05,
	0x0a, 0x05, 0x4f, 0x75, 0x74, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x67, 0x75, 0x6c,
	0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x67, 0x75, 0x6c, 0x61,
	0x72, 0x12, 0x22, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x72, 0x65,
	0x63, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x12, 0x15, 0x0a, 0x03, 0x72, 0x65, 0x71, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x02, 0x52, 0x03, 0x72, 0x65, 0x71, 0x12, 0x23, 0x0a, 0x06, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0x92, 0xb5, 0x18, 0x07, 0x0a,
	0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x06, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x12, 0x33,
	0x0a, 0x0c, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x77, 0x61, 0x6c, 0x6b,
	0x2e, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x0b, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x49, 0x6e,
	0x6e, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x0b, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x69, 0x6e, 0x6e,
	0x65, 0x72, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x77, 0x61, 0x6c, 0x6b, 0x2e, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x0a, 0x6d, 0x75, 0x6c, 0x74,
	0x69, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x3b, 0x0a, 0x09, 0x6d, 0x61, 0x70, 0x5f, 0x69, 0x6e,
	0x6e, 0x65, 0x72, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x77, 0x61, 0x6c, 0x6b, 0x2e, 0x4f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x4d, 0x61, 0x70, 0x49,
	0x6e, 0x6e, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x61, 0x70, 0x49, 0x6e,
	0x6e, 0x65, 0x72, 0x12, 0x3f, 0x0a, 0x10, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x64, 0x65, 0x70,
	0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x77, 0x61, 0x6c, 0x6b, 0x2e, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x42,
	0x02, 0x18, 0x01, 0x52, 0x0f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x44, 0x65, 0x70, 0x72, 0x65, 0x63,
	0x61, 0x74, 0x65, 0x64, 0x12, 0x45, 0x0a, 0x0d, 0x69, 0x6e, 0x74, 0x5f, 0x6d, 0x61, 0x70, 0x5f,
	0x69, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x77, 0x61, 0x6c, 0x6b, 0x2e, 0x4f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x49, 0x6e,
	0x74, 0x4d, 0x61, 0x70, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b,
	0x69, 0x6e, 0x74, 0x4d, 0x61, 0x70, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x0c, 0x6f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x77, 0x61, 0x6c, 0x6b, 0x2e, 0x49, 0x6e,
	0x6e, 0x65, 0x72, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0b, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x49, 0x6e, 0x6e, 0x65, 0x72, 0x1a, 0x4d, 0x0a, 0x0d, 0x4d, 0x61, 0x70, 0x49, 0x6e, 0x6e, 0x65,
	0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x26, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x77,
	0x61, 0x6c, 0x6b, 0x2e, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x1a, 0x50, 0x0a, 0x10, 0x49, 0x6e, 0x74, 0x4d, 0x61, 0x70, 0x49, 0x6e,
	0x6e, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x26, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x77, 0x61, 0x6c, 0x6b, 0x2e, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x3a, 0x4d, 0x0a, 0x06, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0xd2, 0x86, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x77,
	0x61, 0x6c, 0x6b, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x78, 0x74, 0x52, 0x06, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f,
	0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x77, 0x61, 0x6c, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_common_proto_protowalk_fields_test_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_common_proto_protowalk_fields_test_proto_rawDescData = file_go_chromium_org_luci_common_proto_protowalk_fields_test_proto_rawDesc
)

func file_go_chromium_org_luci_common_proto_protowalk_fields_test_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_common_proto_protowalk_fields_test_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_common_proto_protowalk_fields_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_common_proto_protowalk_fields_test_proto_rawDescData)
	})
	return file_go_chromium_org_luci_common_proto_protowalk_fields_test_proto_rawDescData
}

var file_go_chromium_org_luci_common_proto_protowalk_fields_test_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_go_chromium_org_luci_common_proto_protowalk_fields_test_proto_goTypes = []interface{}{
	(*CustomExt)(nil),                 // 0: protowalk.CustomExt
	(*Inner)(nil),                     // 1: protowalk.Inner
	(*Outer)(nil),                     // 2: protowalk.Outer
	(*Inner_Embedded)(nil),            // 3: protowalk.Inner.Embedded
	nil,                               // 4: protowalk.Inner.MapEmbedEntry
	nil,                               // 5: protowalk.Outer.MapInnerEntry
	nil,                               // 6: protowalk.Outer.IntMapInnerEntry
	(*descriptorpb.FieldOptions)(nil), // 7: google.protobuf.FieldOptions
}
var file_go_chromium_org_luci_common_proto_protowalk_fields_test_proto_depIdxs = []int32{
	3,  // 0: protowalk.Inner.single_embed:type_name -> protowalk.Inner.Embedded
	3,  // 1: protowalk.Inner.multi_embed:type_name -> protowalk.Inner.Embedded
	4,  // 2: protowalk.Inner.map_embed:type_name -> protowalk.Inner.MapEmbedEntry
	1,  // 3: protowalk.Outer.single_inner:type_name -> protowalk.Inner
	1,  // 4: protowalk.Outer.multi_inner:type_name -> protowalk.Inner
	5,  // 5: protowalk.Outer.map_inner:type_name -> protowalk.Outer.MapInnerEntry
	1,  // 6: protowalk.Outer.multi_deprecated:type_name -> protowalk.Inner
	6,  // 7: protowalk.Outer.int_map_inner:type_name -> protowalk.Outer.IntMapInnerEntry
	1,  // 8: protowalk.Outer.output_inner:type_name -> protowalk.Inner
	3,  // 9: protowalk.Inner.MapEmbedEntry.value:type_name -> protowalk.Inner.Embedded
	1,  // 10: protowalk.Outer.MapInnerEntry.value:type_name -> protowalk.Inner
	1,  // 11: protowalk.Outer.IntMapInnerEntry.value:type_name -> protowalk.Inner
	7,  // 12: protowalk.custom:extendee -> google.protobuf.FieldOptions
	0,  // 13: protowalk.custom:type_name -> protowalk.CustomExt
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	13, // [13:14] is the sub-list for extension type_name
	12, // [12:13] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_common_proto_protowalk_fields_test_proto_init() }
func file_go_chromium_org_luci_common_proto_protowalk_fields_test_proto_init() {
	if File_go_chromium_org_luci_common_proto_protowalk_fields_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_common_proto_protowalk_fields_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomExt); i {
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
		file_go_chromium_org_luci_common_proto_protowalk_fields_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Inner); i {
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
		file_go_chromium_org_luci_common_proto_protowalk_fields_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Outer); i {
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
		file_go_chromium_org_luci_common_proto_protowalk_fields_test_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Inner_Embedded); i {
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
			RawDescriptor: file_go_chromium_org_luci_common_proto_protowalk_fields_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_common_proto_protowalk_fields_test_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_common_proto_protowalk_fields_test_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_common_proto_protowalk_fields_test_proto_msgTypes,
		ExtensionInfos:    file_go_chromium_org_luci_common_proto_protowalk_fields_test_proto_extTypes,
	}.Build()
	File_go_chromium_org_luci_common_proto_protowalk_fields_test_proto = out.File
	file_go_chromium_org_luci_common_proto_protowalk_fields_test_proto_rawDesc = nil
	file_go_chromium_org_luci_common_proto_protowalk_fields_test_proto_goTypes = nil
	file_go_chromium_org_luci_common_proto_protowalk_fields_test_proto_depIdxs = nil
}
