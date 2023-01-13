// Copyright 2018 The Swarming Authors. All rights reserved.
// Use of this source code is governed by the Apache v2.0 license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: go.chromium.org/luci/buildbucket/proto/notification.proto

package buildbucketpb

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

// A notification about a build.
// Deprecated: this is no longer in use anymore.
type Notification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// When this notification was created.
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Cloud Project ID of the Buildbucket instance that sent this notification,
	// e.g. "cr-buildbucket".
	// Useful if a service listens to both prod and dev instances of buildbucket.
	AppId string `protobuf:"bytes,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	// Buildbucket build ID.
	// Use GetBuild rpc to load the contents.
	BuildId int64 `protobuf:"varint,3,opt,name=build_id,json=buildId,proto3" json:"build_id,omitempty"`
	// User-defined opaque blob specified in NotificationConfig.user_data.
	UserData []byte `protobuf:"bytes,4,opt,name=user_data,json=userData,proto3" json:"user_data,omitempty"`
}

func (x *Notification) Reset() {
	*x = Notification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_buildbucket_proto_notification_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notification) ProtoMessage() {}

func (x *Notification) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_buildbucket_proto_notification_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notification.ProtoReflect.Descriptor instead.
func (*Notification) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_buildbucket_proto_notification_proto_rawDescGZIP(), []int{0}
}

func (x *Notification) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Notification) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *Notification) GetBuildId() int64 {
	if x != nil {
		return x.BuildId
	}
	return 0
}

func (x *Notification) GetUserData() []byte {
	if x != nil {
		return x.UserData
	}
	return nil
}

// Configuration of notifications.
type NotificationConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Target Cloud PubSub topic.
	// Usually has format "projects/{cloud project}/topics/{topic name}".
	//
	// The PubSub message data schema is:
	//
	//	 {
	//	  'build': ${BuildMessage},
	//	  'user_data': ${NotificationConfig.user_data}
	//	  'hostname': 'cr-buildbucket.appspot.com',
	//	}
	//
	// where the BuildMessage is
	// https://chromium.googlesource.com/infra/infra.git/+/b3204748243a9e4bf815a7024e921be46e3e1747/appengine/cr-buildbucket/legacy/api_common.py#94
	//
	// Note: The above data schema is in a legacy format and will be changed soon.
	// So for new users who want to use this, please contact LUCI owners first.
	//
	// <buildbucket-app-id>@appspot.gserviceaccount.com must have
	// "pubsub.topics.publish" permissions on the topic, where
	// <buildbucket-app-id> is usually "cr-buildbucket."
	PubsubTopic string `protobuf:"bytes,1,opt,name=pubsub_topic,json=pubsubTopic,proto3" json:"pubsub_topic,omitempty"`
	// Will be available in Notification.user_data.
	// Max length: 4096.
	UserData []byte `protobuf:"bytes,2,opt,name=user_data,json=userData,proto3" json:"user_data,omitempty"`
}

func (x *NotificationConfig) Reset() {
	*x = NotificationConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_buildbucket_proto_notification_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationConfig) ProtoMessage() {}

func (x *NotificationConfig) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_buildbucket_proto_notification_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationConfig.ProtoReflect.Descriptor instead.
func (*NotificationConfig) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_buildbucket_proto_notification_proto_rawDescGZIP(), []int{1}
}

func (x *NotificationConfig) GetPubsubTopic() string {
	if x != nil {
		return x.PubsubTopic
	}
	return ""
}

func (x *NotificationConfig) GetUserData() []byte {
	if x != nil {
		return x.UserData
	}
	return nil
}

// BuildsV2PubSub is the "builds_v2" pubsub topic message data schema.
// Attributes of this pubsub message:
// - "project"
// - "bucket"
// - "builder"
// - "is_completed" (The value is either "true" or "false" in string.)
type BuildsV2PubSub struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Contains all field except large fields
	Build *Build `protobuf:"bytes,1,opt,name=build,proto3" json:"build,omitempty"`
	// A Compressed bytes in proto binary format of buildbucket.v2.Build where
	// it only contains the large build fields - build.input.properties,
	// build.output.properties and build.steps.
	BuildLargeFields []byte `protobuf:"bytes,2,opt,name=build_large_fields,json=buildLargeFields,proto3" json:"build_large_fields,omitempty"`
	// The compression method the above `build_large_fields` uses. By default, it
	// is ZLIB as this is the most common one and is the built-in lib in many
	// programming languages.
	Compression Compression `protobuf:"varint,3,opt,name=compression,proto3,enum=buildbucket.v2.Compression" json:"compression,omitempty"`
}

func (x *BuildsV2PubSub) Reset() {
	*x = BuildsV2PubSub{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_buildbucket_proto_notification_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildsV2PubSub) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildsV2PubSub) ProtoMessage() {}

func (x *BuildsV2PubSub) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_buildbucket_proto_notification_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildsV2PubSub.ProtoReflect.Descriptor instead.
func (*BuildsV2PubSub) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_buildbucket_proto_notification_proto_rawDescGZIP(), []int{2}
}

func (x *BuildsV2PubSub) GetBuild() *Build {
	if x != nil {
		return x.Build
	}
	return nil
}

func (x *BuildsV2PubSub) GetBuildLargeFields() []byte {
	if x != nil {
		return x.BuildLargeFields
	}
	return nil
}

func (x *BuildsV2PubSub) GetCompression() Compression {
	if x != nil {
		return x.Compression
	}
	return Compression_ZLIB
}

var File_go_chromium_org_luci_buildbucket_proto_notification_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_buildbucket_proto_notification_proto_rawDesc = []byte{
	0x0a, 0x39, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x32, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x67, 0x6f,
	0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75,
	0x63, 0x69, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x33, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x01, 0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x22,
	0x54, 0x0a, 0x12, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x5f,
	0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x75, 0x62,
	0x73, 0x75, 0x62, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x44, 0x61, 0x74, 0x61, 0x22, 0xaa, 0x01, 0x0a, 0x0e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x73,
	0x56, 0x32, 0x50, 0x75, 0x62, 0x53, 0x75, 0x62, 0x12, 0x2b, 0x0a, 0x05, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x05,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x6c,
	0x61, 0x72, 0x67, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x10, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x4c, 0x61, 0x72, 0x67, 0x65, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x12, 0x3d, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75,
	0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_buildbucket_proto_notification_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_buildbucket_proto_notification_proto_rawDescData = file_go_chromium_org_luci_buildbucket_proto_notification_proto_rawDesc
)

func file_go_chromium_org_luci_buildbucket_proto_notification_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_buildbucket_proto_notification_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_buildbucket_proto_notification_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_buildbucket_proto_notification_proto_rawDescData)
	})
	return file_go_chromium_org_luci_buildbucket_proto_notification_proto_rawDescData
}

var file_go_chromium_org_luci_buildbucket_proto_notification_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_go_chromium_org_luci_buildbucket_proto_notification_proto_goTypes = []interface{}{
	(*Notification)(nil),          // 0: buildbucket.v2.Notification
	(*NotificationConfig)(nil),    // 1: buildbucket.v2.NotificationConfig
	(*BuildsV2PubSub)(nil),        // 2: buildbucket.v2.BuildsV2PubSub
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*Build)(nil),                 // 4: buildbucket.v2.Build
	(Compression)(0),              // 5: buildbucket.v2.Compression
}
var file_go_chromium_org_luci_buildbucket_proto_notification_proto_depIdxs = []int32{
	3, // 0: buildbucket.v2.Notification.timestamp:type_name -> google.protobuf.Timestamp
	4, // 1: buildbucket.v2.BuildsV2PubSub.build:type_name -> buildbucket.v2.Build
	5, // 2: buildbucket.v2.BuildsV2PubSub.compression:type_name -> buildbucket.v2.Compression
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_buildbucket_proto_notification_proto_init() }
func file_go_chromium_org_luci_buildbucket_proto_notification_proto_init() {
	if File_go_chromium_org_luci_buildbucket_proto_notification_proto != nil {
		return
	}
	file_go_chromium_org_luci_buildbucket_proto_build_proto_init()
	file_go_chromium_org_luci_buildbucket_proto_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_buildbucket_proto_notification_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Notification); i {
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
		file_go_chromium_org_luci_buildbucket_proto_notification_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotificationConfig); i {
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
		file_go_chromium_org_luci_buildbucket_proto_notification_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildsV2PubSub); i {
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
			RawDescriptor: file_go_chromium_org_luci_buildbucket_proto_notification_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_buildbucket_proto_notification_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_buildbucket_proto_notification_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_buildbucket_proto_notification_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_buildbucket_proto_notification_proto = out.File
	file_go_chromium_org_luci_buildbucket_proto_notification_proto_rawDesc = nil
	file_go_chromium_org_luci_buildbucket_proto_notification_proto_goTypes = nil
	file_go_chromium_org_luci_buildbucket_proto_notification_proto_depIdxs = nil
}
