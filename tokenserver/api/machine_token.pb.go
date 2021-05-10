// Copyright 2016 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.1
// source: go.chromium.org/luci/tokenserver/api/machine_token.proto

package tokenserver

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

// The kinds of machine tokens the token server can mint.
//
// Passed to MintMachineToken and InspectMachineToken.
//
// Reserved: 1.
type MachineTokenType int32

const (
	MachineTokenType_UNKNOWN_TYPE       MachineTokenType = 0 // used if the field is not initialized
	MachineTokenType_LUCI_MACHINE_TOKEN MachineTokenType = 2 // matches serialized MachineTokenEnvelope
)

// Enum value maps for MachineTokenType.
var (
	MachineTokenType_name = map[int32]string{
		0: "UNKNOWN_TYPE",
		2: "LUCI_MACHINE_TOKEN",
	}
	MachineTokenType_value = map[string]int32{
		"UNKNOWN_TYPE":       0,
		"LUCI_MACHINE_TOKEN": 2,
	}
)

func (x MachineTokenType) Enum() *MachineTokenType {
	p := new(MachineTokenType)
	*p = x
	return p
}

func (x MachineTokenType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MachineTokenType) Descriptor() protoreflect.EnumDescriptor {
	return file_go_chromium_org_luci_tokenserver_api_machine_token_proto_enumTypes[0].Descriptor()
}

func (MachineTokenType) Type() protoreflect.EnumType {
	return &file_go_chromium_org_luci_tokenserver_api_machine_token_proto_enumTypes[0]
}

func (x MachineTokenType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MachineTokenType.Descriptor instead.
func (MachineTokenType) EnumDescriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_tokenserver_api_machine_token_proto_rawDescGZIP(), []int{0}
}

// MachineTokenBody describes internal structure of the machine token.
//
// The token will be put in HTTP headers and its body shouldn't be too large.
// For that reason we use unix timestamps instead of google.protobuf.Timestamp
// (no need for microsecond precision), and assume certificate serial numbers
// are smallish uint64 integers (not random blobs).
type MachineTokenBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Machine identity this token conveys (machine FQDN).
	//
	// It is extracted from a Common Name of a certificate used as a basis for
	// the token.
	MachineFqdn string `protobuf:"bytes,1,opt,name=machine_fqdn,json=machineFqdn,proto3" json:"machine_fqdn,omitempty"`
	// Service account email that signed this token.
	//
	// When verifying the token backends will check that the issuer is in
	// "auth-token-servers" group.
	IssuedBy string `protobuf:"bytes,2,opt,name=issued_by,json=issuedBy,proto3" json:"issued_by,omitempty"`
	// Unix timestamp in seconds when this token was issued. Required.
	IssuedAt uint64 `protobuf:"varint,3,opt,name=issued_at,json=issuedAt,proto3" json:"issued_at,omitempty"`
	// Number of seconds the token is considered valid.
	//
	// Usually 3600. Set by the token server. Required.
	Lifetime uint64 `protobuf:"varint,4,opt,name=lifetime,proto3" json:"lifetime,omitempty"`
	// Id of a CA that issued machine certificate used to make this token.
	//
	// These IDs are defined in token server config (via unique_id field).
	CaId int64 `protobuf:"varint,5,opt,name=ca_id,json=caId,proto3" json:"ca_id,omitempty"`
	// Serial number of the machine certificate used to make this token.
	//
	// ca_id and cert_sn together uniquely identify the certificate, and can be
	// used to check for certificate revocation (by asking token server whether
	// the given certificate is in CRL). Revocation checks are optional, most
	// callers can rely on expiration checks only.
	CertSn uint64 `protobuf:"varint,6,opt,name=cert_sn,json=certSn,proto3" json:"cert_sn,omitempty"`
}

func (x *MachineTokenBody) Reset() {
	*x = MachineTokenBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_tokenserver_api_machine_token_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MachineTokenBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MachineTokenBody) ProtoMessage() {}

func (x *MachineTokenBody) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_tokenserver_api_machine_token_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MachineTokenBody.ProtoReflect.Descriptor instead.
func (*MachineTokenBody) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_tokenserver_api_machine_token_proto_rawDescGZIP(), []int{0}
}

func (x *MachineTokenBody) GetMachineFqdn() string {
	if x != nil {
		return x.MachineFqdn
	}
	return ""
}

func (x *MachineTokenBody) GetIssuedBy() string {
	if x != nil {
		return x.IssuedBy
	}
	return ""
}

func (x *MachineTokenBody) GetIssuedAt() uint64 {
	if x != nil {
		return x.IssuedAt
	}
	return 0
}

func (x *MachineTokenBody) GetLifetime() uint64 {
	if x != nil {
		return x.Lifetime
	}
	return 0
}

func (x *MachineTokenBody) GetCaId() int64 {
	if x != nil {
		return x.CaId
	}
	return 0
}

func (x *MachineTokenBody) GetCertSn() uint64 {
	if x != nil {
		return x.CertSn
	}
	return 0
}

// MachineTokenEnvelope is what is actually being serialized and represented
// as a machine token (after being encoded using base64 standard raw encoding).
//
// Resulting token (including base64 encoding) is usually ~500 bytes long.
type MachineTokenEnvelope struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenBody []byte `protobuf:"bytes,1,opt,name=token_body,json=tokenBody,proto3" json:"token_body,omitempty"` // serialized MachineTokenBody
	KeyId     string `protobuf:"bytes,2,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`             // id of a token server private key used for signing
	RsaSha256 []byte `protobuf:"bytes,3,opt,name=rsa_sha256,json=rsaSha256,proto3" json:"rsa_sha256,omitempty"` // signature of 'token_body'
}

func (x *MachineTokenEnvelope) Reset() {
	*x = MachineTokenEnvelope{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_tokenserver_api_machine_token_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MachineTokenEnvelope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MachineTokenEnvelope) ProtoMessage() {}

func (x *MachineTokenEnvelope) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_tokenserver_api_machine_token_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MachineTokenEnvelope.ProtoReflect.Descriptor instead.
func (*MachineTokenEnvelope) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_tokenserver_api_machine_token_proto_rawDescGZIP(), []int{1}
}

func (x *MachineTokenEnvelope) GetTokenBody() []byte {
	if x != nil {
		return x.TokenBody
	}
	return nil
}

func (x *MachineTokenEnvelope) GetKeyId() string {
	if x != nil {
		return x.KeyId
	}
	return ""
}

func (x *MachineTokenEnvelope) GetRsaSha256() []byte {
	if x != nil {
		return x.RsaSha256
	}
	return nil
}

var File_go_chromium_org_luci_tokenserver_api_machine_token_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_tokenserver_api_machine_token_proto_rawDesc = []byte{
	0x0a, 0x38, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0xb9, 0x01, 0x0a, 0x10, 0x4d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x21, 0x0a, 0x0c,
	0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f, 0x66, 0x71, 0x64, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x46, 0x71, 0x64, 0x6e, 0x12,
	0x1b, 0x0a, 0x09, 0x69, 0x73, 0x73, 0x75, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x69, 0x73, 0x73, 0x75, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1b, 0x0a, 0x09,
	0x69, 0x73, 0x73, 0x75, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x08, 0x69, 0x73, 0x73, 0x75, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x69, 0x66,
	0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x6c, 0x69, 0x66,
	0x65, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x13, 0x0a, 0x05, 0x63, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x61, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x65,
	0x72, 0x74, 0x5f, 0x73, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x63, 0x65, 0x72,
	0x74, 0x53, 0x6e, 0x22, 0x6b, 0x0a, 0x14, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x15, 0x0a, 0x06, 0x6b, 0x65,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6b, 0x65, 0x79, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x73, 0x61, 0x5f, 0x73, 0x68, 0x61, 0x32, 0x35, 0x36, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x72, 0x73, 0x61, 0x53, 0x68, 0x61, 0x32, 0x35, 0x36,
	0x2a, 0x3c, 0x0a, 0x10, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x4c, 0x55, 0x43, 0x49, 0x5f, 0x4d,
	0x41, 0x43, 0x48, 0x49, 0x4e, 0x45, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x10, 0x02, 0x42, 0x32,
	0x5a, 0x30, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_tokenserver_api_machine_token_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_tokenserver_api_machine_token_proto_rawDescData = file_go_chromium_org_luci_tokenserver_api_machine_token_proto_rawDesc
)

func file_go_chromium_org_luci_tokenserver_api_machine_token_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_tokenserver_api_machine_token_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_tokenserver_api_machine_token_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_tokenserver_api_machine_token_proto_rawDescData)
	})
	return file_go_chromium_org_luci_tokenserver_api_machine_token_proto_rawDescData
}

var file_go_chromium_org_luci_tokenserver_api_machine_token_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_go_chromium_org_luci_tokenserver_api_machine_token_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_go_chromium_org_luci_tokenserver_api_machine_token_proto_goTypes = []interface{}{
	(MachineTokenType)(0),        // 0: tokenserver.MachineTokenType
	(*MachineTokenBody)(nil),     // 1: tokenserver.MachineTokenBody
	(*MachineTokenEnvelope)(nil), // 2: tokenserver.MachineTokenEnvelope
}
var file_go_chromium_org_luci_tokenserver_api_machine_token_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_tokenserver_api_machine_token_proto_init() }
func file_go_chromium_org_luci_tokenserver_api_machine_token_proto_init() {
	if File_go_chromium_org_luci_tokenserver_api_machine_token_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_tokenserver_api_machine_token_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MachineTokenBody); i {
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
		file_go_chromium_org_luci_tokenserver_api_machine_token_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MachineTokenEnvelope); i {
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
			RawDescriptor: file_go_chromium_org_luci_tokenserver_api_machine_token_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_tokenserver_api_machine_token_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_tokenserver_api_machine_token_proto_depIdxs,
		EnumInfos:         file_go_chromium_org_luci_tokenserver_api_machine_token_proto_enumTypes,
		MessageInfos:      file_go_chromium_org_luci_tokenserver_api_machine_token_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_tokenserver_api_machine_token_proto = out.File
	file_go_chromium_org_luci_tokenserver_api_machine_token_proto_rawDesc = nil
	file_go_chromium_org_luci_tokenserver_api_machine_token_proto_goTypes = nil
	file_go_chromium_org_luci_tokenserver_api_machine_token_proto_depIdxs = nil
}
