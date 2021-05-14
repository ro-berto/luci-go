// Copyright 2017 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Messages for the task queue.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.0
// source: go.chromium.org/luci/luci_notify/internal/tq.proto

package internal

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

// EmailTask represents a single email notification to be dispatched.
type EmailTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Recipients is a list of email addresses to send the email to.
	// TODO(nodir): make it non-repeated.
	Recipients []string `protobuf:"bytes,1,rep,name=recipients,proto3" json:"recipients,omitempty"`
	// Subject is the subject line of the email to be sent.
	Subject string `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	// DEPRECATED. See body_gzip.
	Body string `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	// Gzipped, HTML-formatted string containing the body of the email
	// to be sent.
	BodyGzip []byte `protobuf:"bytes,4,opt,name=body_gzip,json=bodyGzip,proto3" json:"body_gzip,omitempty"`
}

func (x *EmailTask) Reset() {
	*x = EmailTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_luci_notify_internal_tq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailTask) ProtoMessage() {}

func (x *EmailTask) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_luci_notify_internal_tq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailTask.ProtoReflect.Descriptor instead.
func (*EmailTask) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_luci_notify_internal_tq_proto_rawDescGZIP(), []int{0}
}

func (x *EmailTask) GetRecipients() []string {
	if x != nil {
		return x.Recipients
	}
	return nil
}

func (x *EmailTask) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *EmailTask) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *EmailTask) GetBodyGzip() []byte {
	if x != nil {
		return x.BodyGzip
	}
	return nil
}

var File_go_chromium_org_luci_luci_notify_internal_tq_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_luci_notify_internal_tq_proto_rawDesc = []byte{
	0x0a, 0x32, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x5f, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x71, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x22, 0x76,
	0x0a, 0x09, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x1e, 0x0a, 0x0a, 0x72,
	0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0a, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x6f, 0x64,
	0x79, 0x5f, 0x67, 0x7a, 0x69, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x62, 0x6f,
	0x64, 0x79, 0x47, 0x7a, 0x69, 0x70, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72,
	0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x6c,
	0x75, 0x63, 0x69, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_luci_notify_internal_tq_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_luci_notify_internal_tq_proto_rawDescData = file_go_chromium_org_luci_luci_notify_internal_tq_proto_rawDesc
)

func file_go_chromium_org_luci_luci_notify_internal_tq_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_luci_notify_internal_tq_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_luci_notify_internal_tq_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_luci_notify_internal_tq_proto_rawDescData)
	})
	return file_go_chromium_org_luci_luci_notify_internal_tq_proto_rawDescData
}

var file_go_chromium_org_luci_luci_notify_internal_tq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_go_chromium_org_luci_luci_notify_internal_tq_proto_goTypes = []interface{}{
	(*EmailTask)(nil), // 0: internal.EmailTask
}
var file_go_chromium_org_luci_luci_notify_internal_tq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_luci_notify_internal_tq_proto_init() }
func file_go_chromium_org_luci_luci_notify_internal_tq_proto_init() {
	if File_go_chromium_org_luci_luci_notify_internal_tq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_luci_notify_internal_tq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailTask); i {
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
			RawDescriptor: file_go_chromium_org_luci_luci_notify_internal_tq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_luci_notify_internal_tq_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_luci_notify_internal_tq_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_luci_notify_internal_tq_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_luci_notify_internal_tq_proto = out.File
	file_go_chromium_org_luci_luci_notify_internal_tq_proto_rawDesc = nil
	file_go_chromium_org_luci_luci_notify_internal_tq_proto_goTypes = nil
	file_go_chromium_org_luci_luci_notify_internal_tq_proto_depIdxs = nil
}
