// Copyright (c) 2015 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.1
// source: go.chromium.org/luci/common/proto/srcman/manifest_diff.proto

package srcman

import (
	proto "github.com/golang/protobuf/proto"
	git "go.chromium.org/luci/common/proto/git"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Stat indicates how a given item has changed.
type ManifestDiff_Stat int32

const (
	// These two items are identical
	ManifestDiff_EQUAL ManifestDiff_Stat = 0
	// The item was added in `new` compared to `old`
	ManifestDiff_ADDED ManifestDiff_Stat = 1
	// The item was removed in `new` compared to `old`
	ManifestDiff_REMOVED ManifestDiff_Stat = 2
	// The item is in both, but is incomparable (e.g. repo_url changed from
	// `old` to `new`).
	ManifestDiff_MODIFIED ManifestDiff_Stat = 4
	// The item is in both, and is directly comparable (e.g. different
	// revisions of the same repo_url). This only applies to the revision fields
	// of SCM messages.
	//
	// This is 0x8 | MODIFIED, so that users who don't care about DIFF v.
	// MODIFIED can check `Status & MODIFIED`.
	ManifestDiff_DIFF ManifestDiff_Stat = 12
)

// Enum value maps for ManifestDiff_Stat.
var (
	ManifestDiff_Stat_name = map[int32]string{
		0:  "EQUAL",
		1:  "ADDED",
		2:  "REMOVED",
		4:  "MODIFIED",
		12: "DIFF",
	}
	ManifestDiff_Stat_value = map[string]int32{
		"EQUAL":    0,
		"ADDED":    1,
		"REMOVED":  2,
		"MODIFIED": 4,
		"DIFF":     12,
	}
)

func (x ManifestDiff_Stat) Enum() *ManifestDiff_Stat {
	p := new(ManifestDiff_Stat)
	*p = x
	return p
}

func (x ManifestDiff_Stat) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ManifestDiff_Stat) Descriptor() protoreflect.EnumDescriptor {
	return file_go_chromium_org_luci_common_proto_srcman_manifest_diff_proto_enumTypes[0].Descriptor()
}

func (ManifestDiff_Stat) Type() protoreflect.EnumType {
	return &file_go_chromium_org_luci_common_proto_srcman_manifest_diff_proto_enumTypes[0]
}

func (x ManifestDiff_Stat) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ManifestDiff_Stat.Descriptor instead.
func (ManifestDiff_Stat) EnumDescriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_common_proto_srcman_manifest_diff_proto_rawDescGZIP(), []int{0, 0}
}

// ManifestDiff holds basic difference information between two source manifests.
type ManifestDiff struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The older of the two manifests.
	Old *Manifest `protobuf:"bytes,1,opt,name=old,proto3" json:"old,omitempty"`
	// The newer of the two manifests.
	New *Manifest `protobuf:"bytes,2,opt,name=new,proto3" json:"new,omitempty"`
	// Indicates if there is some overall difference between old and new.
	Overall     ManifestDiff_Stat                  `protobuf:"varint,3,opt,name=overall,proto3,enum=srcman.ManifestDiff_Stat" json:"overall,omitempty"`
	Directories map[string]*ManifestDiff_Directory `protobuf:"bytes,4,rep,name=directories,proto3" json:"directories,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ManifestDiff) Reset() {
	*x = ManifestDiff{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_common_proto_srcman_manifest_diff_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManifestDiff) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManifestDiff) ProtoMessage() {}

func (x *ManifestDiff) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_common_proto_srcman_manifest_diff_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManifestDiff.ProtoReflect.Descriptor instead.
func (*ManifestDiff) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_common_proto_srcman_manifest_diff_proto_rawDescGZIP(), []int{0}
}

func (x *ManifestDiff) GetOld() *Manifest {
	if x != nil {
		return x.Old
	}
	return nil
}

func (x *ManifestDiff) GetNew() *Manifest {
	if x != nil {
		return x.New
	}
	return nil
}

func (x *ManifestDiff) GetOverall() ManifestDiff_Stat {
	if x != nil {
		return x.Overall
	}
	return ManifestDiff_EQUAL
}

func (x *ManifestDiff) GetDirectories() map[string]*ManifestDiff_Directory {
	if x != nil {
		return x.Directories
	}
	return nil
}

type ManifestDiff_GitCheckout struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Indicates if there is some overall difference between old and new.
	Overall ManifestDiff_Stat `protobuf:"varint,1,opt,name=overall,proto3,enum=srcman.ManifestDiff_Stat" json:"overall,omitempty"`
	// Indicates the status for the `revision` field.
	//
	// If this is DIFF, it is sensible to compute
	//   `git log repo_url old.revision new.revision`
	Revision ManifestDiff_Stat `protobuf:"varint,2,opt,name=revision,proto3,enum=srcman.ManifestDiff_Stat" json:"revision,omitempty"`
	// Indicates the status for the `patch_revision` field. It evaluates
	// the patch_fetch_ref values to ensure that old and new are different
	// patches from the same CL.
	//
	// If this is DIFF, it is sensible to compute
	//   `git log repo_url old.patch_revision new.patch_revision`
	PatchRevision ManifestDiff_Stat `protobuf:"varint,3,opt,name=patch_revision,json=patchRevision,proto3,enum=srcman.ManifestDiff_Stat" json:"patch_revision,omitempty"`
	// The URL that should be used for RPCs. It may differ from the url in old
	// or new if the service computing this ManifestDiff knows of e.g. a repo
	// URL migration.
	RepoUrl string `protobuf:"bytes,4,opt,name=repo_url,json=repoUrl,proto3" json:"repo_url,omitempty"`
	// If revision==DIFF, this may be populated with git history occurring
	// between the two base revisions.
	History []*git.Commit `protobuf:"bytes,5,rep,name=history,proto3" json:"history,omitempty"`
}

func (x *ManifestDiff_GitCheckout) Reset() {
	*x = ManifestDiff_GitCheckout{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_common_proto_srcman_manifest_diff_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManifestDiff_GitCheckout) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManifestDiff_GitCheckout) ProtoMessage() {}

func (x *ManifestDiff_GitCheckout) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_common_proto_srcman_manifest_diff_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManifestDiff_GitCheckout.ProtoReflect.Descriptor instead.
func (*ManifestDiff_GitCheckout) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_common_proto_srcman_manifest_diff_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ManifestDiff_GitCheckout) GetOverall() ManifestDiff_Stat {
	if x != nil {
		return x.Overall
	}
	return ManifestDiff_EQUAL
}

func (x *ManifestDiff_GitCheckout) GetRevision() ManifestDiff_Stat {
	if x != nil {
		return x.Revision
	}
	return ManifestDiff_EQUAL
}

func (x *ManifestDiff_GitCheckout) GetPatchRevision() ManifestDiff_Stat {
	if x != nil {
		return x.PatchRevision
	}
	return ManifestDiff_EQUAL
}

func (x *ManifestDiff_GitCheckout) GetRepoUrl() string {
	if x != nil {
		return x.RepoUrl
	}
	return ""
}

func (x *ManifestDiff_GitCheckout) GetHistory() []*git.Commit {
	if x != nil {
		return x.History
	}
	return nil
}

type ManifestDiff_Directory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This is the overall status for this Directory.
	Overall        ManifestDiff_Stat         `protobuf:"varint,1,opt,name=overall,proto3,enum=srcman.ManifestDiff_Stat" json:"overall,omitempty"`
	GitCheckout    *ManifestDiff_GitCheckout `protobuf:"bytes,2,opt,name=git_checkout,json=gitCheckout,proto3" json:"git_checkout,omitempty"`
	CipdServerHost ManifestDiff_Stat         `protobuf:"varint,3,opt,name=cipd_server_host,json=cipdServerHost,proto3,enum=srcman.ManifestDiff_Stat" json:"cipd_server_host,omitempty"`
	// Note: this will only ever be MODIFIED, because we cannot (yet) determine
	// if two versions of a cipd package are diffable. We may later implement
	// DIFF detection (i.e. if both packages use `version:<sha1>` tags).
	CipdPackage        map[string]ManifestDiff_Stat `protobuf:"bytes,4,rep,name=cipd_package,json=cipdPackage,proto3" json:"cipd_package,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=srcman.ManifestDiff_Stat"`
	IsolatedServerHost ManifestDiff_Stat            `protobuf:"varint,5,opt,name=isolated_server_host,json=isolatedServerHost,proto3,enum=srcman.ManifestDiff_Stat" json:"isolated_server_host,omitempty"`
	// This merely indicates if the list of isolated hashes was the same or not;
	// there's not a good way to register the two lists.
	//
	// Since order-of-application for isolateds matters, this will indicate
	// MODIFIED if the order of isolated hashes changes.
	Isolated ManifestDiff_Stat `protobuf:"varint,6,opt,name=isolated,proto3,enum=srcman.ManifestDiff_Stat" json:"isolated,omitempty"`
}

func (x *ManifestDiff_Directory) Reset() {
	*x = ManifestDiff_Directory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_common_proto_srcman_manifest_diff_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManifestDiff_Directory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManifestDiff_Directory) ProtoMessage() {}

func (x *ManifestDiff_Directory) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_common_proto_srcman_manifest_diff_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManifestDiff_Directory.ProtoReflect.Descriptor instead.
func (*ManifestDiff_Directory) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_common_proto_srcman_manifest_diff_proto_rawDescGZIP(), []int{0, 1}
}

func (x *ManifestDiff_Directory) GetOverall() ManifestDiff_Stat {
	if x != nil {
		return x.Overall
	}
	return ManifestDiff_EQUAL
}

func (x *ManifestDiff_Directory) GetGitCheckout() *ManifestDiff_GitCheckout {
	if x != nil {
		return x.GitCheckout
	}
	return nil
}

func (x *ManifestDiff_Directory) GetCipdServerHost() ManifestDiff_Stat {
	if x != nil {
		return x.CipdServerHost
	}
	return ManifestDiff_EQUAL
}

func (x *ManifestDiff_Directory) GetCipdPackage() map[string]ManifestDiff_Stat {
	if x != nil {
		return x.CipdPackage
	}
	return nil
}

func (x *ManifestDiff_Directory) GetIsolatedServerHost() ManifestDiff_Stat {
	if x != nil {
		return x.IsolatedServerHost
	}
	return ManifestDiff_EQUAL
}

func (x *ManifestDiff_Directory) GetIsolated() ManifestDiff_Stat {
	if x != nil {
		return x.Isolated
	}
	return ManifestDiff_EQUAL
}

var File_go_chromium_org_luci_common_proto_srcman_manifest_diff_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_common_proto_srcman_manifest_diff_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x72, 0x63, 0x6d, 0x61, 0x6e, 0x2f, 0x6d, 0x61, 0x6e, 0x69, 0x66,
	0x65, 0x73, 0x74, 0x5f, 0x64, 0x69, 0x66, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x73, 0x72, 0x63, 0x6d, 0x61, 0x6e, 0x1a, 0x32, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d,
	0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x69, 0x74, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x37, 0x67, 0x6f, 0x2e, 0x63,
	0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x72,
	0x63, 0x6d, 0x61, 0x6e, 0x2f, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xf7, 0x08, 0x0a, 0x0c, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74,
	0x44, 0x69, 0x66, 0x66, 0x12, 0x22, 0x0a, 0x03, 0x6f, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x73, 0x72, 0x63, 0x6d, 0x61, 0x6e, 0x2e, 0x4d, 0x61, 0x6e, 0x69, 0x66,
	0x65, 0x73, 0x74, 0x52, 0x03, 0x6f, 0x6c, 0x64, 0x12, 0x22, 0x0a, 0x03, 0x6e, 0x65, 0x77, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x72, 0x63, 0x6d, 0x61, 0x6e, 0x2e, 0x4d,
	0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x52, 0x03, 0x6e, 0x65, 0x77, 0x12, 0x33, 0x0a, 0x07,
	0x6f, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e,
	0x73, 0x72, 0x63, 0x6d, 0x61, 0x6e, 0x2e, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x44,
	0x69, 0x66, 0x66, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x52, 0x07, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x6c,
	0x6c, 0x12, 0x47, 0x0a, 0x0b, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x72, 0x63, 0x6d, 0x61, 0x6e, 0x2e,
	0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x44, 0x69, 0x66, 0x66, 0x2e, 0x44, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x1a, 0xfd, 0x01, 0x0a, 0x0b, 0x47,
	0x69, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x12, 0x33, 0x0a, 0x07, 0x6f, 0x76,
	0x65, 0x72, 0x61, 0x6c, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x73, 0x72,
	0x63, 0x6d, 0x61, 0x6e, 0x2e, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x44, 0x69, 0x66,
	0x66, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x52, 0x07, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x6c, 0x12,
	0x35, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x19, 0x2e, 0x73, 0x72, 0x63, 0x6d, 0x61, 0x6e, 0x2e, 0x4d, 0x61, 0x6e, 0x69, 0x66,
	0x65, 0x73, 0x74, 0x44, 0x69, 0x66, 0x66, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x52, 0x08, 0x72, 0x65,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x0e, 0x70, 0x61, 0x74, 0x63, 0x68, 0x5f,
	0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19,
	0x2e, 0x73, 0x72, 0x63, 0x6d, 0x61, 0x6e, 0x2e, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74,
	0x44, 0x69, 0x66, 0x66, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x52, 0x0d, 0x70, 0x61, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x65, 0x70, 0x6f,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x70, 0x6f,
	0x55, 0x72, 0x6c, 0x12, 0x25, 0x0a, 0x07, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x67, 0x69, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x52, 0x07, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x1a, 0xfd, 0x03, 0x0a, 0x09, 0x44,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x33, 0x0a, 0x07, 0x6f, 0x76, 0x65, 0x72,
	0x61, 0x6c, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x73, 0x72, 0x63, 0x6d,
	0x61, 0x6e, 0x2e, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x44, 0x69, 0x66, 0x66, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x52, 0x07, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x6c, 0x12, 0x43, 0x0a,
	0x0c, 0x67, 0x69, 0x74, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x72, 0x63, 0x6d, 0x61, 0x6e, 0x2e, 0x4d, 0x61, 0x6e,
	0x69, 0x66, 0x65, 0x73, 0x74, 0x44, 0x69, 0x66, 0x66, 0x2e, 0x47, 0x69, 0x74, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x0b, 0x67, 0x69, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f,
	0x75, 0x74, 0x12, 0x43, 0x0a, 0x10, 0x63, 0x69, 0x70, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x73,
	0x72, 0x63, 0x6d, 0x61, 0x6e, 0x2e, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x44, 0x69,
	0x66, 0x66, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x52, 0x0e, 0x63, 0x69, 0x70, 0x64, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x52, 0x0a, 0x0c, 0x63, 0x69, 0x70, 0x64, 0x5f,
	0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e,
	0x73, 0x72, 0x63, 0x6d, 0x61, 0x6e, 0x2e, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x44,
	0x69, 0x66, 0x66, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x43, 0x69,
	0x70, 0x64, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b,
	0x63, 0x69, 0x70, 0x64, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x4b, 0x0a, 0x14, 0x69,
	0x73, 0x6f, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x68,
	0x6f, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x73, 0x72, 0x63, 0x6d,
	0x61, 0x6e, 0x2e, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x44, 0x69, 0x66, 0x66, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x52, 0x12, 0x69, 0x73, 0x6f, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x08, 0x69, 0x73, 0x6f, 0x6c,
	0x61, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x73, 0x72, 0x63,
	0x6d, 0x61, 0x6e, 0x2e, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x44, 0x69, 0x66, 0x66,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x52, 0x08, 0x69, 0x73, 0x6f, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x1a,
	0x59, 0x0a, 0x10, 0x43, 0x69, 0x70, 0x64, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x73, 0x72, 0x63, 0x6d, 0x61, 0x6e, 0x2e, 0x4d, 0x61,
	0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x44, 0x69, 0x66, 0x66, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x5e, 0x0a, 0x10, 0x44, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x34, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x73, 0x72, 0x63, 0x6d, 0x61, 0x6e, 0x2e, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73,
	0x74, 0x44, 0x69, 0x66, 0x66, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x41, 0x0a, 0x04, 0x53, 0x74,
	0x61, 0x74, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x51, 0x55, 0x41, 0x4c, 0x10, 0x00, 0x12, 0x09, 0x0a,
	0x05, 0x41, 0x44, 0x44, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x45, 0x4d, 0x4f,
	0x56, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x4d, 0x4f, 0x44, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x04, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x49, 0x46, 0x46, 0x10, 0x0c, 0x42, 0x2a, 0x5a,
	0x28, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x73, 0x72, 0x63, 0x6d, 0x61, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_go_chromium_org_luci_common_proto_srcman_manifest_diff_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_common_proto_srcman_manifest_diff_proto_rawDescData = file_go_chromium_org_luci_common_proto_srcman_manifest_diff_proto_rawDesc
)

func file_go_chromium_org_luci_common_proto_srcman_manifest_diff_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_common_proto_srcman_manifest_diff_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_common_proto_srcman_manifest_diff_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_common_proto_srcman_manifest_diff_proto_rawDescData)
	})
	return file_go_chromium_org_luci_common_proto_srcman_manifest_diff_proto_rawDescData
}

var file_go_chromium_org_luci_common_proto_srcman_manifest_diff_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_go_chromium_org_luci_common_proto_srcman_manifest_diff_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_go_chromium_org_luci_common_proto_srcman_manifest_diff_proto_goTypes = []interface{}{
	(ManifestDiff_Stat)(0),           // 0: srcman.ManifestDiff.Stat
	(*ManifestDiff)(nil),             // 1: srcman.ManifestDiff
	(*ManifestDiff_GitCheckout)(nil), // 2: srcman.ManifestDiff.GitCheckout
	(*ManifestDiff_Directory)(nil),   // 3: srcman.ManifestDiff.Directory
	nil,                              // 4: srcman.ManifestDiff.DirectoriesEntry
	nil,                              // 5: srcman.ManifestDiff.Directory.CipdPackageEntry
	(*Manifest)(nil),                 // 6: srcman.Manifest
	(*git.Commit)(nil),               // 7: git.Commit
}
var file_go_chromium_org_luci_common_proto_srcman_manifest_diff_proto_depIdxs = []int32{
	6,  // 0: srcman.ManifestDiff.old:type_name -> srcman.Manifest
	6,  // 1: srcman.ManifestDiff.new:type_name -> srcman.Manifest
	0,  // 2: srcman.ManifestDiff.overall:type_name -> srcman.ManifestDiff.Stat
	4,  // 3: srcman.ManifestDiff.directories:type_name -> srcman.ManifestDiff.DirectoriesEntry
	0,  // 4: srcman.ManifestDiff.GitCheckout.overall:type_name -> srcman.ManifestDiff.Stat
	0,  // 5: srcman.ManifestDiff.GitCheckout.revision:type_name -> srcman.ManifestDiff.Stat
	0,  // 6: srcman.ManifestDiff.GitCheckout.patch_revision:type_name -> srcman.ManifestDiff.Stat
	7,  // 7: srcman.ManifestDiff.GitCheckout.history:type_name -> git.Commit
	0,  // 8: srcman.ManifestDiff.Directory.overall:type_name -> srcman.ManifestDiff.Stat
	2,  // 9: srcman.ManifestDiff.Directory.git_checkout:type_name -> srcman.ManifestDiff.GitCheckout
	0,  // 10: srcman.ManifestDiff.Directory.cipd_server_host:type_name -> srcman.ManifestDiff.Stat
	5,  // 11: srcman.ManifestDiff.Directory.cipd_package:type_name -> srcman.ManifestDiff.Directory.CipdPackageEntry
	0,  // 12: srcman.ManifestDiff.Directory.isolated_server_host:type_name -> srcman.ManifestDiff.Stat
	0,  // 13: srcman.ManifestDiff.Directory.isolated:type_name -> srcman.ManifestDiff.Stat
	3,  // 14: srcman.ManifestDiff.DirectoriesEntry.value:type_name -> srcman.ManifestDiff.Directory
	0,  // 15: srcman.ManifestDiff.Directory.CipdPackageEntry.value:type_name -> srcman.ManifestDiff.Stat
	16, // [16:16] is the sub-list for method output_type
	16, // [16:16] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_common_proto_srcman_manifest_diff_proto_init() }
func file_go_chromium_org_luci_common_proto_srcman_manifest_diff_proto_init() {
	if File_go_chromium_org_luci_common_proto_srcman_manifest_diff_proto != nil {
		return
	}
	file_go_chromium_org_luci_common_proto_srcman_manifest_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_common_proto_srcman_manifest_diff_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManifestDiff); i {
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
		file_go_chromium_org_luci_common_proto_srcman_manifest_diff_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManifestDiff_GitCheckout); i {
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
		file_go_chromium_org_luci_common_proto_srcman_manifest_diff_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManifestDiff_Directory); i {
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
			RawDescriptor: file_go_chromium_org_luci_common_proto_srcman_manifest_diff_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_common_proto_srcman_manifest_diff_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_common_proto_srcman_manifest_diff_proto_depIdxs,
		EnumInfos:         file_go_chromium_org_luci_common_proto_srcman_manifest_diff_proto_enumTypes,
		MessageInfos:      file_go_chromium_org_luci_common_proto_srcman_manifest_diff_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_common_proto_srcman_manifest_diff_proto = out.File
	file_go_chromium_org_luci_common_proto_srcman_manifest_diff_proto_rawDesc = nil
	file_go_chromium_org_luci_common_proto_srcman_manifest_diff_proto_goTypes = nil
	file_go_chromium_org_luci_common_proto_srcman_manifest_diff_proto_depIdxs = nil
}
