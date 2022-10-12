// Copyright 2015 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: go.chromium.org/luci/cipd/client/cipd/internal/messages/messages.proto

package messages

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

// BlobWithSHA256 is a wrapper around a binary blob with SHA256 hash to verify
// its integrity.
type BlobWithSHA256 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blob   []byte `protobuf:"bytes,1,opt,name=blob,proto3" json:"blob,omitempty"`
	Sha256 []byte `protobuf:"bytes,3,opt,name=sha256,proto3" json:"sha256,omitempty"`
}

func (x *BlobWithSHA256) Reset() {
	*x = BlobWithSHA256{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cipd_client_cipd_internal_messages_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlobWithSHA256) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlobWithSHA256) ProtoMessage() {}

func (x *BlobWithSHA256) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cipd_client_cipd_internal_messages_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlobWithSHA256.ProtoReflect.Descriptor instead.
func (*BlobWithSHA256) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cipd_client_cipd_internal_messages_messages_proto_rawDescGZIP(), []int{0}
}

func (x *BlobWithSHA256) GetBlob() []byte {
	if x != nil {
		return x.Blob
	}
	return nil
}

func (x *BlobWithSHA256) GetSha256() []byte {
	if x != nil {
		return x.Sha256
	}
	return nil
}

// TagCache stores a mapping (service, package name, tag) -> instance ID to
// speed up subsequent ResolveVersion calls when tags are used.
//
// It also contains a '(service, instance_id, file_name) -> encoded ObjectRef'
// mapping which is used for client self-update purposes. file_name is
// case-sensitive and must always use POSIX-style slashes.
//
// A service is specified by its hostname. We make it part of the key since
// same tags may point to different instances on different services.
type TagCache struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Capped list of entries, most recently resolved is last.
	Entries     []*TagCache_Entry     `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
	FileEntries []*TagCache_FileEntry `protobuf:"bytes,2,rep,name=file_entries,json=fileEntries,proto3" json:"file_entries,omitempty"`
}

func (x *TagCache) Reset() {
	*x = TagCache{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cipd_client_cipd_internal_messages_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagCache) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagCache) ProtoMessage() {}

func (x *TagCache) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cipd_client_cipd_internal_messages_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagCache.ProtoReflect.Descriptor instead.
func (*TagCache) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cipd_client_cipd_internal_messages_messages_proto_rawDescGZIP(), []int{1}
}

func (x *TagCache) GetEntries() []*TagCache_Entry {
	if x != nil {
		return x.Entries
	}
	return nil
}

func (x *TagCache) GetFileEntries() []*TagCache_FileEntry {
	if x != nil {
		return x.FileEntries
	}
	return nil
}

// InstanceCache stores a list of instances and their last access time.
//
// This cache does not depend on a service being used, since an instance's ID is
// derived only from its contents (regardless from where it was downloaded).
type InstanceCache struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Entries is a map of {instance id -> information about instance}.
	Entries map[string]*InstanceCache_Entry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// LastSynced is timestamp when we synchronized Entries with actual
	// instance files.
	LastSynced *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=last_synced,json=lastSynced,proto3" json:"last_synced,omitempty"`
}

func (x *InstanceCache) Reset() {
	*x = InstanceCache{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cipd_client_cipd_internal_messages_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstanceCache) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstanceCache) ProtoMessage() {}

func (x *InstanceCache) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cipd_client_cipd_internal_messages_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstanceCache.ProtoReflect.Descriptor instead.
func (*InstanceCache) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cipd_client_cipd_internal_messages_messages_proto_rawDescGZIP(), []int{2}
}

func (x *InstanceCache) GetEntries() map[string]*InstanceCache_Entry {
	if x != nil {
		return x.Entries
	}
	return nil
}

func (x *InstanceCache) GetLastSynced() *timestamppb.Timestamp {
	if x != nil {
		return x.LastSynced
	}
	return nil
}

type TagCache_Entry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service    string `protobuf:"bytes,4,opt,name=service,proto3" json:"service,omitempty"`                         // e.g. 'chrome-infra-packages.appspot.com'
	Package    string `protobuf:"bytes,1,opt,name=package,proto3" json:"package,omitempty"`                         // name of a tagged CIPD package
	Tag        string `protobuf:"bytes,2,opt,name=tag,proto3" json:"tag,omitempty"`                                 // the tag, e.g. 'k:v'
	InstanceId string `protobuf:"bytes,3,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"` // the instance ID it resolves to
}

func (x *TagCache_Entry) Reset() {
	*x = TagCache_Entry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cipd_client_cipd_internal_messages_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagCache_Entry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagCache_Entry) ProtoMessage() {}

func (x *TagCache_Entry) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cipd_client_cipd_internal_messages_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagCache_Entry.ProtoReflect.Descriptor instead.
func (*TagCache_Entry) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cipd_client_cipd_internal_messages_messages_proto_rawDescGZIP(), []int{1, 0}
}

func (x *TagCache_Entry) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *TagCache_Entry) GetPackage() string {
	if x != nil {
		return x.Package
	}
	return ""
}

func (x *TagCache_Entry) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *TagCache_Entry) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

type TagCache_FileEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service    string `protobuf:"bytes,5,opt,name=service,proto3" json:"service,omitempty"`                         // e.g. 'chrome-infra-packages.appspot.com'
	Package    string `protobuf:"bytes,1,opt,name=package,proto3" json:"package,omitempty"`                         // name of a CIPD package containing the file
	InstanceId string `protobuf:"bytes,2,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"` // identifier of the CIPD package instance
	FileName   string `protobuf:"bytes,3,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`       // file name inside the package, POSIX-style slashes
	ObjectRef  string `protobuf:"bytes,4,opt,name=object_ref,json=objectRef,proto3" json:"object_ref,omitempty"`    // file's ObjectRef as encoded by ObjectRefToInstanceID (for legacy reasons)
}

func (x *TagCache_FileEntry) Reset() {
	*x = TagCache_FileEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cipd_client_cipd_internal_messages_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagCache_FileEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagCache_FileEntry) ProtoMessage() {}

func (x *TagCache_FileEntry) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cipd_client_cipd_internal_messages_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagCache_FileEntry.ProtoReflect.Descriptor instead.
func (*TagCache_FileEntry) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cipd_client_cipd_internal_messages_messages_proto_rawDescGZIP(), []int{1, 1}
}

func (x *TagCache_FileEntry) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *TagCache_FileEntry) GetPackage() string {
	if x != nil {
		return x.Package
	}
	return ""
}

func (x *TagCache_FileEntry) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

func (x *TagCache_FileEntry) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *TagCache_FileEntry) GetObjectRef() string {
	if x != nil {
		return x.ObjectRef
	}
	return ""
}

// Entry stores info about an instance.
type InstanceCache_Entry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// LastAccess is last time this instance was retrieved from or put to the
	// cache.
	LastAccess *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=last_access,json=lastAccess,proto3" json:"last_access,omitempty"`
}

func (x *InstanceCache_Entry) Reset() {
	*x = InstanceCache_Entry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cipd_client_cipd_internal_messages_messages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstanceCache_Entry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstanceCache_Entry) ProtoMessage() {}

func (x *InstanceCache_Entry) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cipd_client_cipd_internal_messages_messages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstanceCache_Entry.ProtoReflect.Descriptor instead.
func (*InstanceCache_Entry) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cipd_client_cipd_internal_messages_messages_proto_rawDescGZIP(), []int{2, 0}
}

func (x *InstanceCache_Entry) GetLastAccess() *timestamppb.Timestamp {
	if x != nil {
		return x.LastAccess
	}
	return nil
}

var File_go_chromium_org_luci_cipd_client_cipd_internal_messages_messages_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_cipd_client_cipd_internal_messages_messages_proto_rawDesc = []byte{
	0x0a, 0x46, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x69, 0x70, 0x64, 0x2f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2f, 0x63, 0x69, 0x70, 0x64, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a, 0x0e, 0x42, 0x6c, 0x6f, 0x62, 0x57, 0x69, 0x74, 0x68, 0x53,
	0x48, 0x41, 0x32, 0x35, 0x36, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6c, 0x6f, 0x62, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6c, 0x6f, 0x62, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x68, 0x61,
	0x32, 0x35, 0x36, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x73, 0x68, 0x61, 0x32, 0x35,
	0x36, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x22, 0x8e, 0x03, 0x0a, 0x08, 0x54, 0x61, 0x67, 0x43,
	0x61, 0x63, 0x68, 0x65, 0x12, 0x32, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x54, 0x61, 0x67, 0x43, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x3f, 0x0a, 0x0c, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x54, 0x61, 0x67, 0x43, 0x61, 0x63,
	0x68, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x66, 0x69,
	0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x1a, 0x6e, 0x0a, 0x05, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x1a, 0x9c, 0x01, 0x0a, 0x09, 0x46, 0x69,
	0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x22, 0xad, 0x02, 0x0a, 0x0d, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65, 0x12, 0x3e, 0x0a, 0x07, 0x65, 0x6e,
	0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x43,
	0x61, 0x63, 0x68, 0x65, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x3b, 0x0a, 0x0b, 0x6c, 0x61,
	0x73, 0x74, 0x5f, 0x73, 0x79, 0x6e, 0x63, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x6c, 0x61, 0x73,
	0x74, 0x53, 0x79, 0x6e, 0x63, 0x65, 0x64, 0x1a, 0x44, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x3b, 0x0a, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x1a, 0x59, 0x0a,
	0x0c, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x33, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x6f, 0x2e, 0x63,
	0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69,
	0x2f, 0x63, 0x69, 0x70, 0x64, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x69, 0x70,
	0x64, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_cipd_client_cipd_internal_messages_messages_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_cipd_client_cipd_internal_messages_messages_proto_rawDescData = file_go_chromium_org_luci_cipd_client_cipd_internal_messages_messages_proto_rawDesc
)

func file_go_chromium_org_luci_cipd_client_cipd_internal_messages_messages_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_cipd_client_cipd_internal_messages_messages_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_cipd_client_cipd_internal_messages_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_cipd_client_cipd_internal_messages_messages_proto_rawDescData)
	})
	return file_go_chromium_org_luci_cipd_client_cipd_internal_messages_messages_proto_rawDescData
}

var file_go_chromium_org_luci_cipd_client_cipd_internal_messages_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_go_chromium_org_luci_cipd_client_cipd_internal_messages_messages_proto_goTypes = []interface{}{
	(*BlobWithSHA256)(nil),        // 0: messages.BlobWithSHA256
	(*TagCache)(nil),              // 1: messages.TagCache
	(*InstanceCache)(nil),         // 2: messages.InstanceCache
	(*TagCache_Entry)(nil),        // 3: messages.TagCache.Entry
	(*TagCache_FileEntry)(nil),    // 4: messages.TagCache.FileEntry
	(*InstanceCache_Entry)(nil),   // 5: messages.InstanceCache.Entry
	nil,                           // 6: messages.InstanceCache.EntriesEntry
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_go_chromium_org_luci_cipd_client_cipd_internal_messages_messages_proto_depIdxs = []int32{
	3, // 0: messages.TagCache.entries:type_name -> messages.TagCache.Entry
	4, // 1: messages.TagCache.file_entries:type_name -> messages.TagCache.FileEntry
	6, // 2: messages.InstanceCache.entries:type_name -> messages.InstanceCache.EntriesEntry
	7, // 3: messages.InstanceCache.last_synced:type_name -> google.protobuf.Timestamp
	7, // 4: messages.InstanceCache.Entry.last_access:type_name -> google.protobuf.Timestamp
	5, // 5: messages.InstanceCache.EntriesEntry.value:type_name -> messages.InstanceCache.Entry
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_cipd_client_cipd_internal_messages_messages_proto_init() }
func file_go_chromium_org_luci_cipd_client_cipd_internal_messages_messages_proto_init() {
	if File_go_chromium_org_luci_cipd_client_cipd_internal_messages_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_cipd_client_cipd_internal_messages_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlobWithSHA256); i {
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
		file_go_chromium_org_luci_cipd_client_cipd_internal_messages_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagCache); i {
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
		file_go_chromium_org_luci_cipd_client_cipd_internal_messages_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstanceCache); i {
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
		file_go_chromium_org_luci_cipd_client_cipd_internal_messages_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagCache_Entry); i {
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
		file_go_chromium_org_luci_cipd_client_cipd_internal_messages_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagCache_FileEntry); i {
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
		file_go_chromium_org_luci_cipd_client_cipd_internal_messages_messages_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstanceCache_Entry); i {
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
			RawDescriptor: file_go_chromium_org_luci_cipd_client_cipd_internal_messages_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_cipd_client_cipd_internal_messages_messages_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_cipd_client_cipd_internal_messages_messages_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_cipd_client_cipd_internal_messages_messages_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_cipd_client_cipd_internal_messages_messages_proto = out.File
	file_go_chromium_org_luci_cipd_client_cipd_internal_messages_messages_proto_rawDesc = nil
	file_go_chromium_org_luci_cipd_client_cipd_internal_messages_messages_proto_goTypes = nil
	file_go_chromium_org_luci_cipd_client_cipd_internal_messages_messages_proto_depIdxs = nil
}
