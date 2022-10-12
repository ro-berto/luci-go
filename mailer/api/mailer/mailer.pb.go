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
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: go.chromium.org/luci/mailer/api/mailer/mailer.proto

package mailer

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// SendMailRequest is passed to SendMail RPC and contains the email to send.
//
// Addresses may be of any form permitted by RFC 822. At least one of To, Cc,
// or Bcc must be non-empty.
//
// At least one of TextBody or HtmlBody must be non-empty.
type SendMailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A unique identifier for this request to guarantee idempotency.
	//
	// Restricted to 36 ASCII characters. A random UUID is recommended. This
	// request is only idempotent if a `request_id` is provided.
	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// Sender is put into "From" email header field.
	//
	// The server will validate this field and reject requests that use disallowed
	// values (based on the authenticated identity of the caller).
	Sender string `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	// ReplyTo is put into "Reply-To" email header field.
	ReplyTo string `protobuf:"bytes,3,opt,name=reply_to,json=replyTo,proto3" json:"reply_to,omitempty"`
	// To is put into "To" email header field.
	To []string `protobuf:"bytes,4,rep,name=to,proto3" json:"to,omitempty"`
	// Cc is put into "Cc" email header field.
	Cc []string `protobuf:"bytes,5,rep,name=cc,proto3" json:"cc,omitempty"`
	// Bcc is put into "Bcc" email header field.
	Bcc []string `protobuf:"bytes,6,rep,name=bcc,proto3" json:"bcc,omitempty"`
	// Subject is put into "Subject" email header field.
	Subject string `protobuf:"bytes,7,opt,name=subject,proto3" json:"subject,omitempty"`
	// TextBody contains a plaintext body of the email message.
	TextBody string `protobuf:"bytes,8,opt,name=text_body,json=textBody,proto3" json:"text_body,omitempty"`
	// HtmlBody contains an HTML body of the email message.
	HtmlBody string `protobuf:"bytes,9,opt,name=html_body,json=htmlBody,proto3" json:"html_body,omitempty"`
}

func (x *SendMailRequest) Reset() {
	*x = SendMailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_mailer_api_mailer_mailer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMailRequest) ProtoMessage() {}

func (x *SendMailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_mailer_api_mailer_mailer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMailRequest.ProtoReflect.Descriptor instead.
func (*SendMailRequest) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_mailer_api_mailer_mailer_proto_rawDescGZIP(), []int{0}
}

func (x *SendMailRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *SendMailRequest) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *SendMailRequest) GetReplyTo() string {
	if x != nil {
		return x.ReplyTo
	}
	return ""
}

func (x *SendMailRequest) GetTo() []string {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *SendMailRequest) GetCc() []string {
	if x != nil {
		return x.Cc
	}
	return nil
}

func (x *SendMailRequest) GetBcc() []string {
	if x != nil {
		return x.Bcc
	}
	return nil
}

func (x *SendMailRequest) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *SendMailRequest) GetTextBody() string {
	if x != nil {
		return x.TextBody
	}
	return ""
}

func (x *SendMailRequest) GetHtmlBody() string {
	if x != nil {
		return x.HtmlBody
	}
	return ""
}

// SendMailResponse is returned by SendMail RPC.
type SendMailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// MessageId as an opaque identifier of the enqueued email.
	MessageId string `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
}

func (x *SendMailResponse) Reset() {
	*x = SendMailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_mailer_api_mailer_mailer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMailResponse) ProtoMessage() {}

func (x *SendMailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_mailer_api_mailer_mailer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMailResponse.ProtoReflect.Descriptor instead.
func (*SendMailResponse) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_mailer_api_mailer_mailer_proto_rawDescGZIP(), []int{1}
}

func (x *SendMailResponse) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

var File_go_chromium_org_luci_mailer_api_mailer_mailer_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_mailer_api_mailer_mailer_proto_rawDesc = []byte{
	0x0a, 0x33, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6c, 0x75, 0x63, 0x69, 0x2e, 0x6d, 0x61, 0x69, 0x6c,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf3, 0x01, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64, 0x4d,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x06,
	0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x5f,
	0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x54,
	0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x02, 0x74,
	0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x63, 0x63, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x02, 0x63,
	0x63, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x63, 0x63, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03,
	0x62, 0x63, 0x63, 0x12, 0x1d, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x78, 0x74, 0x42, 0x6f, 0x64, 0x79, 0x12,
	0x1b, 0x0a, 0x09, 0x68, 0x74, 0x6d, 0x6c, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x68, 0x74, 0x6d, 0x6c, 0x42, 0x6f, 0x64, 0x79, 0x22, 0x31, 0x0a, 0x10,
	0x53, 0x65, 0x6e, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x32,
	0x57, 0x0a, 0x06, 0x4d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x12, 0x4d, 0x0a, 0x08, 0x53, 0x65, 0x6e,
	0x64, 0x4d, 0x61, 0x69, 0x6c, 0x12, 0x1f, 0x2e, 0x6c, 0x75, 0x63, 0x69, 0x2e, 0x6d, 0x61, 0x69,
	0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6c, 0x75, 0x63, 0x69, 0x2e, 0x6d, 0x61,
	0x69, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x6f, 0x2e, 0x63,
	0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69,
	0x2f, 0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x61, 0x69, 0x6c,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_mailer_api_mailer_mailer_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_mailer_api_mailer_mailer_proto_rawDescData = file_go_chromium_org_luci_mailer_api_mailer_mailer_proto_rawDesc
)

func file_go_chromium_org_luci_mailer_api_mailer_mailer_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_mailer_api_mailer_mailer_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_mailer_api_mailer_mailer_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_mailer_api_mailer_mailer_proto_rawDescData)
	})
	return file_go_chromium_org_luci_mailer_api_mailer_mailer_proto_rawDescData
}

var file_go_chromium_org_luci_mailer_api_mailer_mailer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_go_chromium_org_luci_mailer_api_mailer_mailer_proto_goTypes = []interface{}{
	(*SendMailRequest)(nil),  // 0: luci.mailer.v1.SendMailRequest
	(*SendMailResponse)(nil), // 1: luci.mailer.v1.SendMailResponse
}
var file_go_chromium_org_luci_mailer_api_mailer_mailer_proto_depIdxs = []int32{
	0, // 0: luci.mailer.v1.Mailer.SendMail:input_type -> luci.mailer.v1.SendMailRequest
	1, // 1: luci.mailer.v1.Mailer.SendMail:output_type -> luci.mailer.v1.SendMailResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_mailer_api_mailer_mailer_proto_init() }
func file_go_chromium_org_luci_mailer_api_mailer_mailer_proto_init() {
	if File_go_chromium_org_luci_mailer_api_mailer_mailer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_mailer_api_mailer_mailer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMailRequest); i {
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
		file_go_chromium_org_luci_mailer_api_mailer_mailer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMailResponse); i {
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
			RawDescriptor: file_go_chromium_org_luci_mailer_api_mailer_mailer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_go_chromium_org_luci_mailer_api_mailer_mailer_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_mailer_api_mailer_mailer_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_mailer_api_mailer_mailer_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_mailer_api_mailer_mailer_proto = out.File
	file_go_chromium_org_luci_mailer_api_mailer_mailer_proto_rawDesc = nil
	file_go_chromium_org_luci_mailer_api_mailer_mailer_proto_goTypes = nil
	file_go_chromium_org_luci_mailer_api_mailer_mailer_proto_depIdxs = nil
}
