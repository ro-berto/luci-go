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
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.0
// source: go.chromium.org/luci/cipd/api/cipd/v1/cas.proto

package api

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

// Supported hashing algorithms used by the content-addressable storage.
//
// Literal names are important, since they are used to construct Google Storage
// paths internally.
type HashAlgo int32

const (
	HashAlgo_HASH_ALGO_UNSPECIFIED HashAlgo = 0
	HashAlgo_SHA1                  HashAlgo = 1
	HashAlgo_SHA256                HashAlgo = 2
)

// Enum value maps for HashAlgo.
var (
	HashAlgo_name = map[int32]string{
		0: "HASH_ALGO_UNSPECIFIED",
		1: "SHA1",
		2: "SHA256",
	}
	HashAlgo_value = map[string]int32{
		"HASH_ALGO_UNSPECIFIED": 0,
		"SHA1":                  1,
		"SHA256":                2,
	}
)

func (x HashAlgo) Enum() *HashAlgo {
	p := new(HashAlgo)
	*p = x
	return p
}

func (x HashAlgo) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HashAlgo) Descriptor() protoreflect.EnumDescriptor {
	return file_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto_enumTypes[0].Descriptor()
}

func (HashAlgo) Type() protoreflect.EnumType {
	return &file_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto_enumTypes[0]
}

func (x HashAlgo) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HashAlgo.Descriptor instead.
func (HashAlgo) EnumDescriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto_rawDescGZIP(), []int{0}
}

type UploadStatus int32

const (
	UploadStatus_UPLOAD_STATUS_UNSPECIFIED UploadStatus = 0
	UploadStatus_UPLOADING                 UploadStatus = 1 // the data is being uploaded now
	UploadStatus_VERIFYING                 UploadStatus = 2 // the object's hash is being verified now
	UploadStatus_PUBLISHED                 UploadStatus = 3 // the object has been published in the CAS
	UploadStatus_ERRORED                   UploadStatus = 4 // there were fatal errors during the finalization
	UploadStatus_CANCELED                  UploadStatus = 5 // the operation was canceled via CancelUpload
)

// Enum value maps for UploadStatus.
var (
	UploadStatus_name = map[int32]string{
		0: "UPLOAD_STATUS_UNSPECIFIED",
		1: "UPLOADING",
		2: "VERIFYING",
		3: "PUBLISHED",
		4: "ERRORED",
		5: "CANCELED",
	}
	UploadStatus_value = map[string]int32{
		"UPLOAD_STATUS_UNSPECIFIED": 0,
		"UPLOADING":                 1,
		"VERIFYING":                 2,
		"PUBLISHED":                 3,
		"ERRORED":                   4,
		"CANCELED":                  5,
	}
)

func (x UploadStatus) Enum() *UploadStatus {
	p := new(UploadStatus)
	*p = x
	return p
}

func (x UploadStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UploadStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto_enumTypes[1].Descriptor()
}

func (UploadStatus) Type() protoreflect.EnumType {
	return &file_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto_enumTypes[1]
}

func (x UploadStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UploadStatus.Descriptor instead.
func (UploadStatus) EnumDescriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto_rawDescGZIP(), []int{1}
}

// A reference to an object in the content-addressable storage.
type ObjectRef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HashAlgo  HashAlgo `protobuf:"varint,1,opt,name=hash_algo,json=hashAlgo,proto3,enum=cipd.HashAlgo" json:"hash_algo,omitempty"` // an algorithm applied to the body to get the name
	HexDigest string   `protobuf:"bytes,2,opt,name=hex_digest,json=hexDigest,proto3" json:"hex_digest,omitempty"`                  // the name as lowercase hex string, e.g 'abcdef...'
}

func (x *ObjectRef) Reset() {
	*x = ObjectRef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectRef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectRef) ProtoMessage() {}

func (x *ObjectRef) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectRef.ProtoReflect.Descriptor instead.
func (*ObjectRef) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto_rawDescGZIP(), []int{0}
}

func (x *ObjectRef) GetHashAlgo() HashAlgo {
	if x != nil {
		return x.HashAlgo
	}
	return HashAlgo_HASH_ALGO_UNSPECIFIED
}

func (x *ObjectRef) GetHexDigest() string {
	if x != nil {
		return x.HexDigest
	}
	return ""
}

type GetObjectURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A reference to the object the client wants to fetch.
	Object *ObjectRef `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
	// If present, the returned URL will be served with Content-Disposition header
	// that includes the given filename. It makes browsers save the file under the
	// given name.
	DownloadFilename string `protobuf:"bytes,2,opt,name=download_filename,json=downloadFilename,proto3" json:"download_filename,omitempty"`
}

func (x *GetObjectURLRequest) Reset() {
	*x = GetObjectURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetObjectURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetObjectURLRequest) ProtoMessage() {}

func (x *GetObjectURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetObjectURLRequest.ProtoReflect.Descriptor instead.
func (*GetObjectURLRequest) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto_rawDescGZIP(), []int{1}
}

func (x *GetObjectURLRequest) GetObject() *ObjectRef {
	if x != nil {
		return x.Object
	}
	return nil
}

func (x *GetObjectURLRequest) GetDownloadFilename() string {
	if x != nil {
		return x.DownloadFilename
	}
	return ""
}

type ObjectURL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A signed HTTPS URL to the object's body.
	//
	// Fetching it doesn't require authentication. Expires after some unspecified
	// short amount of time. It is expected that callers will use it immediately.
	//
	// The URL isn't guaranteed to have any particular internal structure. Do not
	// attempt to parse it.
	SignedUrl string `protobuf:"bytes,1,opt,name=signed_url,json=signedUrl,proto3" json:"signed_url,omitempty"`
}

func (x *ObjectURL) Reset() {
	*x = ObjectURL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectURL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectURL) ProtoMessage() {}

func (x *ObjectURL) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectURL.ProtoReflect.Descriptor instead.
func (*ObjectURL) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto_rawDescGZIP(), []int{2}
}

func (x *ObjectURL) GetSignedUrl() string {
	if x != nil {
		return x.SignedUrl
	}
	return ""
}

type BeginUploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A reference to the object the client wants to put in the storage, if known.
	//
	// If such object already exists, RPC will finish with ALREADY_EXISTS status
	// right away.
	//
	// If this field is missing (in case the client doesn't know the hash yet),
	// the client MUST supply hash_algo field, to let the backend know what
	// hashing algorithm it should use for calculating object's hash.
	//
	// The calculated hash will be returned back to the client as part of
	// UploadOperation ('object' field) by FinishUpload call.
	Object *ObjectRef `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
	// An algorithm to use to derive object's name during uploads when the final
	// hash of the object is not yet known.
	//
	// Optional if 'object' is present.
	//
	// If both 'object' and 'hash_algo' are present, 'object.hash_algo' MUST match
	// 'hash_algo'.
	HashAlgo HashAlgo `protobuf:"varint,2,opt,name=hash_algo,json=hashAlgo,proto3,enum=cipd.HashAlgo" json:"hash_algo,omitempty"`
}

func (x *BeginUploadRequest) Reset() {
	*x = BeginUploadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BeginUploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BeginUploadRequest) ProtoMessage() {}

func (x *BeginUploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BeginUploadRequest.ProtoReflect.Descriptor instead.
func (*BeginUploadRequest) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto_rawDescGZIP(), []int{3}
}

func (x *BeginUploadRequest) GetObject() *ObjectRef {
	if x != nil {
		return x.Object
	}
	return nil
}

func (x *BeginUploadRequest) GetHashAlgo() HashAlgo {
	if x != nil {
		return x.HashAlgo
	}
	return HashAlgo_HASH_ALGO_UNSPECIFIED
}

type FinishUploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An identifier of an upload operation returned by BeginUpload RPC.
	UploadOperationId string `protobuf:"bytes,1,opt,name=upload_operation_id,json=uploadOperationId,proto3" json:"upload_operation_id,omitempty"`
	// If set, instructs Storage to skip the hash verification and just assume the
	// uploaded object has the given hash.
	//
	// This is used internally by the service as an optimization for cases when
	// it trusts the uploaded data (for example, when it upload it itself).
	//
	// External callers are denied usage of this field. Attempt to use it results
	// in PERMISSION_DENIED.
	ForceHash *ObjectRef `protobuf:"bytes,2,opt,name=force_hash,json=forceHash,proto3" json:"force_hash,omitempty"`
}

func (x *FinishUploadRequest) Reset() {
	*x = FinishUploadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinishUploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinishUploadRequest) ProtoMessage() {}

func (x *FinishUploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinishUploadRequest.ProtoReflect.Descriptor instead.
func (*FinishUploadRequest) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto_rawDescGZIP(), []int{4}
}

func (x *FinishUploadRequest) GetUploadOperationId() string {
	if x != nil {
		return x.UploadOperationId
	}
	return ""
}

func (x *FinishUploadRequest) GetForceHash() *ObjectRef {
	if x != nil {
		return x.ForceHash
	}
	return nil
}

type CancelUploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An identifier of an upload operation returned by BeginUpload RPC.
	UploadOperationId string `protobuf:"bytes,1,opt,name=upload_operation_id,json=uploadOperationId,proto3" json:"upload_operation_id,omitempty"`
}

func (x *CancelUploadRequest) Reset() {
	*x = CancelUploadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelUploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelUploadRequest) ProtoMessage() {}

func (x *CancelUploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelUploadRequest.ProtoReflect.Descriptor instead.
func (*CancelUploadRequest) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto_rawDescGZIP(), []int{5}
}

func (x *CancelUploadRequest) GetUploadOperationId() string {
	if x != nil {
		return x.UploadOperationId
	}
	return ""
}

type UploadOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An opaque string that identifies this upload operation.
	//
	// It acts as a temporary authorization token for FinishUpload RPC. Treat it
	// as a secret.
	OperationId string `protobuf:"bytes,1,opt,name=operation_id,json=operationId,proto3" json:"operation_id,omitempty"`
	// URL the client should use in Google Storage Resumable Upload protocol to
	// upload the object's body.
	//
	// No authentication is required to upload data to this URL, so it also should
	// be treated as a secret.
	UploadUrl string `protobuf:"bytes,2,opt,name=upload_url,json=uploadUrl,proto3" json:"upload_url,omitempty"`
	// Status of the upload operation.
	Status UploadStatus `protobuf:"varint,3,opt,name=status,proto3,enum=cipd.UploadStatus" json:"status,omitempty"`
	// For PUBLISHED status, the reference to the published object.
	//
	// This is in particular useful for uploads when the hash of the object is not
	// known until the upload is finished.
	Object *ObjectRef `protobuf:"bytes,4,opt,name=object,proto3" json:"object,omitempty"`
	// For ERRORED status, a human readable error message.
	ErrorMessage string `protobuf:"bytes,5,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
}

func (x *UploadOperation) Reset() {
	*x = UploadOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadOperation) ProtoMessage() {}

func (x *UploadOperation) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadOperation.ProtoReflect.Descriptor instead.
func (*UploadOperation) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto_rawDescGZIP(), []int{6}
}

func (x *UploadOperation) GetOperationId() string {
	if x != nil {
		return x.OperationId
	}
	return ""
}

func (x *UploadOperation) GetUploadUrl() string {
	if x != nil {
		return x.UploadUrl
	}
	return ""
}

func (x *UploadOperation) GetStatus() UploadStatus {
	if x != nil {
		return x.Status
	}
	return UploadStatus_UPLOAD_STATUS_UNSPECIFIED
}

func (x *UploadOperation) GetObject() *ObjectRef {
	if x != nil {
		return x.Object
	}
	return nil
}

func (x *UploadOperation) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

var File_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x69, 0x70, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x69, 0x70, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x63, 0x69, 0x70, 0x64, 0x22, 0x57, 0x0a, 0x09, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x52, 0x65, 0x66, 0x12, 0x2b, 0x0a, 0x09, 0x68, 0x61, 0x73, 0x68, 0x5f, 0x61, 0x6c, 0x67,
	0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x63, 0x69, 0x70, 0x64, 0x2e, 0x48,
	0x61, 0x73, 0x68, 0x41, 0x6c, 0x67, 0x6f, 0x52, 0x08, 0x68, 0x61, 0x73, 0x68, 0x41, 0x6c, 0x67,
	0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x65, 0x78, 0x5f, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x68, 0x65, 0x78, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74,
	0x22, 0x6b, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x55, 0x52, 0x4c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x69, 0x70, 0x64, 0x2e, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x2b, 0x0a, 0x11, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x66, 0x69, 0x6c,
	0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x64, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2a, 0x0a,
	0x09, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x55, 0x52, 0x4c, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x69,
	0x67, 0x6e, 0x65, 0x64, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x55, 0x72, 0x6c, 0x22, 0x6a, 0x0a, 0x12, 0x42, 0x65, 0x67,
	0x69, 0x6e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x27, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x63, 0x69, 0x70, 0x64, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66,
	0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x2b, 0x0a, 0x09, 0x68, 0x61, 0x73, 0x68,
	0x5f, 0x61, 0x6c, 0x67, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x63, 0x69,
	0x70, 0x64, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x41, 0x6c, 0x67, 0x6f, 0x52, 0x08, 0x68, 0x61, 0x73,
	0x68, 0x41, 0x6c, 0x67, 0x6f, 0x22, 0x75, 0x0a, 0x13, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x13,
	0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x75, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x0a,
	0x66, 0x6f, 0x72, 0x63, 0x65, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x63, 0x69, 0x70, 0x64, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x66, 0x52, 0x09, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x48, 0x61, 0x73, 0x68, 0x22, 0x45, 0x0a, 0x13,
	0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x13, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x11, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x22, 0xcd, 0x01, 0x0a, 0x0f, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x72, 0x6c, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x63, 0x69, 0x70, 0x64,
	0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x27, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x69, 0x70, 0x64, 0x2e, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x23,
	0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2a, 0x3b, 0x0a, 0x08, 0x48, 0x61, 0x73, 0x68, 0x41, 0x6c, 0x67, 0x6f, 0x12,
	0x19, 0x0a, 0x15, 0x48, 0x41, 0x53, 0x48, 0x5f, 0x41, 0x4c, 0x47, 0x4f, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x48,
	0x41, 0x31, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x48, 0x41, 0x32, 0x35, 0x36, 0x10, 0x02,
	0x2a, 0x75, 0x0a, 0x0c, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1d, 0x0a, 0x19, 0x55, 0x50, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x0d, 0x0a, 0x09, 0x55, 0x50, 0x4c, 0x4f, 0x41, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0d,
	0x0a, 0x09, 0x56, 0x45, 0x52, 0x49, 0x46, 0x59, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0d, 0x0a,
	0x09, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x53, 0x48, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x45, 0x44, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x41, 0x4e,
	0x43, 0x45, 0x4c, 0x45, 0x44, 0x10, 0x05, 0x32, 0x89, 0x02, 0x0a, 0x07, 0x53, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x12, 0x3a, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x55, 0x52, 0x4c, 0x12, 0x19, 0x2e, 0x63, 0x69, 0x70, 0x64, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f,
	0x2e, 0x63, 0x69, 0x70, 0x64, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x55, 0x52, 0x4c, 0x12,
	0x3e, 0x0a, 0x0b, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x18,
	0x2e, 0x63, 0x69, 0x70, 0x64, 0x2e, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x63, 0x69, 0x70, 0x64, 0x2e,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x40, 0x0a, 0x0c, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12,
	0x19, 0x2e, 0x63, 0x69, 0x70, 0x64, 0x2e, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x63, 0x69, 0x70,
	0x64, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x40, 0x0a, 0x0c, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x12, 0x19, 0x2e, 0x63, 0x69, 0x70, 0x64, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x63,
	0x69, 0x70, 0x64, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69,
	0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x69, 0x70, 0x64,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x69, 0x70, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x70, 0x69,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto_rawDescData = file_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto_rawDesc
)

func file_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto_rawDescData)
	})
	return file_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto_rawDescData
}

var file_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto_goTypes = []interface{}{
	(HashAlgo)(0),               // 0: cipd.HashAlgo
	(UploadStatus)(0),           // 1: cipd.UploadStatus
	(*ObjectRef)(nil),           // 2: cipd.ObjectRef
	(*GetObjectURLRequest)(nil), // 3: cipd.GetObjectURLRequest
	(*ObjectURL)(nil),           // 4: cipd.ObjectURL
	(*BeginUploadRequest)(nil),  // 5: cipd.BeginUploadRequest
	(*FinishUploadRequest)(nil), // 6: cipd.FinishUploadRequest
	(*CancelUploadRequest)(nil), // 7: cipd.CancelUploadRequest
	(*UploadOperation)(nil),     // 8: cipd.UploadOperation
}
var file_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto_depIdxs = []int32{
	0,  // 0: cipd.ObjectRef.hash_algo:type_name -> cipd.HashAlgo
	2,  // 1: cipd.GetObjectURLRequest.object:type_name -> cipd.ObjectRef
	2,  // 2: cipd.BeginUploadRequest.object:type_name -> cipd.ObjectRef
	0,  // 3: cipd.BeginUploadRequest.hash_algo:type_name -> cipd.HashAlgo
	2,  // 4: cipd.FinishUploadRequest.force_hash:type_name -> cipd.ObjectRef
	1,  // 5: cipd.UploadOperation.status:type_name -> cipd.UploadStatus
	2,  // 6: cipd.UploadOperation.object:type_name -> cipd.ObjectRef
	3,  // 7: cipd.Storage.GetObjectURL:input_type -> cipd.GetObjectURLRequest
	5,  // 8: cipd.Storage.BeginUpload:input_type -> cipd.BeginUploadRequest
	6,  // 9: cipd.Storage.FinishUpload:input_type -> cipd.FinishUploadRequest
	7,  // 10: cipd.Storage.CancelUpload:input_type -> cipd.CancelUploadRequest
	4,  // 11: cipd.Storage.GetObjectURL:output_type -> cipd.ObjectURL
	8,  // 12: cipd.Storage.BeginUpload:output_type -> cipd.UploadOperation
	8,  // 13: cipd.Storage.FinishUpload:output_type -> cipd.UploadOperation
	8,  // 14: cipd.Storage.CancelUpload:output_type -> cipd.UploadOperation
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto_init() }
func file_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto_init() {
	if File_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectRef); i {
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
		file_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetObjectURLRequest); i {
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
		file_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectURL); i {
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
		file_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BeginUploadRequest); i {
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
		file_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FinishUploadRequest); i {
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
		file_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelUploadRequest); i {
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
		file_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadOperation); i {
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
			RawDescriptor: file_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto_depIdxs,
		EnumInfos:         file_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto_enumTypes,
		MessageInfos:      file_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto = out.File
	file_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto_rawDesc = nil
	file_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto_goTypes = nil
	file_go_chromium_org_luci_cipd_api_cipd_v1_cas_proto_depIdxs = nil
}
