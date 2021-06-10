// Copyright 2015 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.1
// source: go.chromium.org/luci/logdog/api/config/svcconfig/transport.proto

package svcconfig

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

// Transport is the transport configuration.
type Transport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Type is the transport configuration that is being used.
	//
	// Types that are assignable to Type:
	//	*Transport_Pubsub
	Type isTransport_Type `protobuf_oneof:"Type"`
}

func (x *Transport) Reset() {
	*x = Transport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_logdog_api_config_svcconfig_transport_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transport) ProtoMessage() {}

func (x *Transport) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_logdog_api_config_svcconfig_transport_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transport.ProtoReflect.Descriptor instead.
func (*Transport) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_logdog_api_config_svcconfig_transport_proto_rawDescGZIP(), []int{0}
}

func (m *Transport) GetType() isTransport_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *Transport) GetPubsub() *Transport_PubSub {
	if x, ok := x.GetType().(*Transport_Pubsub); ok {
		return x.Pubsub
	}
	return nil
}

type isTransport_Type interface {
	isTransport_Type()
}

type Transport_Pubsub struct {
	Pubsub *Transport_PubSub `protobuf:"bytes,1,opt,name=pubsub,proto3,oneof"`
}

func (*Transport_Pubsub) isTransport_Type() {}

// PubSub is a transport configuration for Google Cloud Pub/Sub.
type Transport_PubSub struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the authentication group for administrators.
	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	// The name of the authentication group for administrators.
	Topic string `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	// The name of the authentication group for administrators.
	Subscription string `protobuf:"bytes,3,opt,name=subscription,proto3" json:"subscription,omitempty"`
}

func (x *Transport_PubSub) Reset() {
	*x = Transport_PubSub{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_logdog_api_config_svcconfig_transport_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transport_PubSub) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transport_PubSub) ProtoMessage() {}

func (x *Transport_PubSub) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_logdog_api_config_svcconfig_transport_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transport_PubSub.ProtoReflect.Descriptor instead.
func (*Transport_PubSub) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_logdog_api_config_svcconfig_transport_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Transport_PubSub) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *Transport_PubSub) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *Transport_PubSub) GetSubscription() string {
	if x != nil {
		return x.Subscription
	}
	return ""
}

var File_go_chromium_org_luci_logdog_api_config_svcconfig_transport_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_logdog_api_config_svcconfig_transport_proto_rawDesc = []byte{
	0x0a, 0x40, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x6c, 0x6f, 0x67, 0x64, 0x6f, 0x67, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x73, 0x76, 0x63, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x09, 0x73, 0x76, 0x63, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0xa8, 0x01,
	0x0a, 0x09, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x35, 0x0a, 0x06, 0x70,
	0x75, 0x62, 0x73, 0x75, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x76,
	0x63, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x2e, 0x50, 0x75, 0x62, 0x53, 0x75, 0x62, 0x48, 0x00, 0x52, 0x06, 0x70, 0x75, 0x62, 0x73,
	0x75, 0x62, 0x1a, 0x5c, 0x0a, 0x06, 0x50, 0x75, 0x62, 0x53, 0x75, 0x62, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x22, 0x0a, 0x0c,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x06, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x6f, 0x2e, 0x63,
	0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69,
	0x2f, 0x6c, 0x6f, 0x67, 0x64, 0x6f, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2f, 0x73, 0x76, 0x63, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_logdog_api_config_svcconfig_transport_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_logdog_api_config_svcconfig_transport_proto_rawDescData = file_go_chromium_org_luci_logdog_api_config_svcconfig_transport_proto_rawDesc
)

func file_go_chromium_org_luci_logdog_api_config_svcconfig_transport_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_logdog_api_config_svcconfig_transport_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_logdog_api_config_svcconfig_transport_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_logdog_api_config_svcconfig_transport_proto_rawDescData)
	})
	return file_go_chromium_org_luci_logdog_api_config_svcconfig_transport_proto_rawDescData
}

var file_go_chromium_org_luci_logdog_api_config_svcconfig_transport_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_go_chromium_org_luci_logdog_api_config_svcconfig_transport_proto_goTypes = []interface{}{
	(*Transport)(nil),        // 0: svcconfig.Transport
	(*Transport_PubSub)(nil), // 1: svcconfig.Transport.PubSub
}
var file_go_chromium_org_luci_logdog_api_config_svcconfig_transport_proto_depIdxs = []int32{
	1, // 0: svcconfig.Transport.pubsub:type_name -> svcconfig.Transport.PubSub
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_logdog_api_config_svcconfig_transport_proto_init() }
func file_go_chromium_org_luci_logdog_api_config_svcconfig_transport_proto_init() {
	if File_go_chromium_org_luci_logdog_api_config_svcconfig_transport_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_logdog_api_config_svcconfig_transport_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transport); i {
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
		file_go_chromium_org_luci_logdog_api_config_svcconfig_transport_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transport_PubSub); i {
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
	file_go_chromium_org_luci_logdog_api_config_svcconfig_transport_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Transport_Pubsub)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_go_chromium_org_luci_logdog_api_config_svcconfig_transport_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_logdog_api_config_svcconfig_transport_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_logdog_api_config_svcconfig_transport_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_logdog_api_config_svcconfig_transport_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_logdog_api_config_svcconfig_transport_proto = out.File
	file_go_chromium_org_luci_logdog_api_config_svcconfig_transport_proto_rawDesc = nil
	file_go_chromium_org_luci_logdog_api_config_svcconfig_transport_proto_goTypes = nil
	file_go_chromium_org_luci_logdog_api_config_svcconfig_transport_proto_depIdxs = nil
}
