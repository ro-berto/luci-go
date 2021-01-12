// Copyright 2017 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.1
// source: go.chromium.org/luci/tokenserver/api/oauth_token_grant.proto

package tokenserver

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

// OAuthTokenGrantBody contains the internal guts of an oauth token grant.
//
// It gets serialized, signed and stuffed into OAuthTokenGrantEnvelope, which
// then also gets serialized to get the final blob with the grant. This blob is
// then base64-encoded and returned to the caller of MintOAuthTokenGrant.
type OAuthTokenGrantBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifier of this token as generated by the token server.
	//
	// Used for logging and tracking purposes.
	//
	// TODO(vadimsh): It may later be used for revocation purposes.
	TokenId int64 `protobuf:"varint,1,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
	// Service account email the end user wants to act as.
	ServiceAccount string `protobuf:"bytes,2,opt,name=service_account,json=serviceAccount,proto3" json:"service_account,omitempty"`
	// Who can pass this token to MintOAuthTokenViaGrant to get an OAuth token.
	//
	// A string of the form "user:<email>". On Swarming, this is Swarming's own
	// service account name.
	Proxy string `protobuf:"bytes,3,opt,name=proxy,proto3" json:"proxy,omitempty"`
	// An end user that wants to act as the service account (perhaps indirectly).
	//
	// A string of the form "user:<email>". On Swarming, this is an identity of
	// a user that posted the task.
	//
	// Used by MintOAuthTokenViaGrant to recheck that the access is still allowed.
	EndUser string `protobuf:"bytes,4,opt,name=end_user,json=endUser,proto3" json:"end_user,omitempty"`
	// When the token was generated (and when it becomes valid).
	IssuedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=issued_at,json=issuedAt,proto3" json:"issued_at,omitempty"`
	// How long the token is considered valid (in seconds).
	//
	// It may become invalid sooner if the token server policy changes and the
	// new policy doesn't allow this token.
	ValidityDuration int64 `protobuf:"varint,6,opt,name=validity_duration,json=validityDuration,proto3" json:"validity_duration,omitempty"`
}

func (x *OAuthTokenGrantBody) Reset() {
	*x = OAuthTokenGrantBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_tokenserver_api_oauth_token_grant_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OAuthTokenGrantBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OAuthTokenGrantBody) ProtoMessage() {}

func (x *OAuthTokenGrantBody) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_tokenserver_api_oauth_token_grant_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OAuthTokenGrantBody.ProtoReflect.Descriptor instead.
func (*OAuthTokenGrantBody) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_tokenserver_api_oauth_token_grant_proto_rawDescGZIP(), []int{0}
}

func (x *OAuthTokenGrantBody) GetTokenId() int64 {
	if x != nil {
		return x.TokenId
	}
	return 0
}

func (x *OAuthTokenGrantBody) GetServiceAccount() string {
	if x != nil {
		return x.ServiceAccount
	}
	return ""
}

func (x *OAuthTokenGrantBody) GetProxy() string {
	if x != nil {
		return x.Proxy
	}
	return ""
}

func (x *OAuthTokenGrantBody) GetEndUser() string {
	if x != nil {
		return x.EndUser
	}
	return ""
}

func (x *OAuthTokenGrantBody) GetIssuedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.IssuedAt
	}
	return nil
}

func (x *OAuthTokenGrantBody) GetValidityDuration() int64 {
	if x != nil {
		return x.ValidityDuration
	}
	return 0
}

// OAuthTokenGrantEnvelope is what is actually being serialized and send to
// the callers of MintOAuthTokenGrant (after being encoded using base64 standard
// raw encoding).
type OAuthTokenGrantEnvelope struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenBody      []byte `protobuf:"bytes,1,opt,name=token_body,json=tokenBody,proto3" json:"token_body,omitempty"`                  // serialized OAuthTokenGrantBody
	KeyId          string `protobuf:"bytes,2,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`                              // id of a token server private key used for signing
	Pkcs1Sha256Sig []byte `protobuf:"bytes,3,opt,name=pkcs1_sha256_sig,json=pkcs1Sha256Sig,proto3" json:"pkcs1_sha256_sig,omitempty"` // signature of 'token_body'
}

func (x *OAuthTokenGrantEnvelope) Reset() {
	*x = OAuthTokenGrantEnvelope{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_tokenserver_api_oauth_token_grant_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OAuthTokenGrantEnvelope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OAuthTokenGrantEnvelope) ProtoMessage() {}

func (x *OAuthTokenGrantEnvelope) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_tokenserver_api_oauth_token_grant_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OAuthTokenGrantEnvelope.ProtoReflect.Descriptor instead.
func (*OAuthTokenGrantEnvelope) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_tokenserver_api_oauth_token_grant_proto_rawDescGZIP(), []int{1}
}

func (x *OAuthTokenGrantEnvelope) GetTokenBody() []byte {
	if x != nil {
		return x.TokenBody
	}
	return nil
}

func (x *OAuthTokenGrantEnvelope) GetKeyId() string {
	if x != nil {
		return x.KeyId
	}
	return ""
}

func (x *OAuthTokenGrantEnvelope) GetPkcs1Sha256Sig() []byte {
	if x != nil {
		return x.Pkcs1Sha256Sig
	}
	return nil
}

var File_go_chromium_org_luci_tokenserver_api_oauth_token_grant_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_tokenserver_api_oauth_token_grant_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x5f, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf0, 0x01, 0x0a,
	0x13, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x47, 0x72, 0x61, 0x6e, 0x74,
	0x42, 0x6f, 0x64, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x64, 0x12,
	0x27, 0x0a, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x19,
	0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x65, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x37, 0x0a, 0x09, 0x69, 0x73, 0x73,
	0x75, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x69, 0x73, 0x73, 0x75, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x2b, 0x0a, 0x11, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x69, 0x74, 0x79, 0x5f, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x69, 0x74, 0x79, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x79, 0x0a, 0x17, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x47, 0x72, 0x61,
	0x6e, 0x74, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x15, 0x0a, 0x06, 0x6b, 0x65, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6b, 0x65, 0x79, 0x49, 0x64,
	0x12, 0x28, 0x0a, 0x10, 0x70, 0x6b, 0x63, 0x73, 0x31, 0x5f, 0x73, 0x68, 0x61, 0x32, 0x35, 0x36,
	0x5f, 0x73, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x70, 0x6b, 0x63, 0x73,
	0x31, 0x53, 0x68, 0x61, 0x32, 0x35, 0x36, 0x53, 0x69, 0x67, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x6f,
	0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75,
	0x63, 0x69, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61,
	0x70, 0x69, 0x3b, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_tokenserver_api_oauth_token_grant_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_tokenserver_api_oauth_token_grant_proto_rawDescData = file_go_chromium_org_luci_tokenserver_api_oauth_token_grant_proto_rawDesc
)

func file_go_chromium_org_luci_tokenserver_api_oauth_token_grant_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_tokenserver_api_oauth_token_grant_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_tokenserver_api_oauth_token_grant_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_tokenserver_api_oauth_token_grant_proto_rawDescData)
	})
	return file_go_chromium_org_luci_tokenserver_api_oauth_token_grant_proto_rawDescData
}

var file_go_chromium_org_luci_tokenserver_api_oauth_token_grant_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_go_chromium_org_luci_tokenserver_api_oauth_token_grant_proto_goTypes = []interface{}{
	(*OAuthTokenGrantBody)(nil),     // 0: tokenserver.OAuthTokenGrantBody
	(*OAuthTokenGrantEnvelope)(nil), // 1: tokenserver.OAuthTokenGrantEnvelope
	(*timestamppb.Timestamp)(nil),   // 2: google.protobuf.Timestamp
}
var file_go_chromium_org_luci_tokenserver_api_oauth_token_grant_proto_depIdxs = []int32{
	2, // 0: tokenserver.OAuthTokenGrantBody.issued_at:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_tokenserver_api_oauth_token_grant_proto_init() }
func file_go_chromium_org_luci_tokenserver_api_oauth_token_grant_proto_init() {
	if File_go_chromium_org_luci_tokenserver_api_oauth_token_grant_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_tokenserver_api_oauth_token_grant_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OAuthTokenGrantBody); i {
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
		file_go_chromium_org_luci_tokenserver_api_oauth_token_grant_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OAuthTokenGrantEnvelope); i {
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
			RawDescriptor: file_go_chromium_org_luci_tokenserver_api_oauth_token_grant_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_tokenserver_api_oauth_token_grant_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_tokenserver_api_oauth_token_grant_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_tokenserver_api_oauth_token_grant_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_tokenserver_api_oauth_token_grant_proto = out.File
	file_go_chromium_org_luci_tokenserver_api_oauth_token_grant_proto_rawDesc = nil
	file_go_chromium_org_luci_tokenserver_api_oauth_token_grant_proto_goTypes = nil
	file_go_chromium_org_luci_tokenserver_api_oauth_token_grant_proto_depIdxs = nil
}
