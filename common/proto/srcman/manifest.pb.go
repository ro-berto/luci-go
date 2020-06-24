// Copyright 2017 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.1
// source: go.chromium.org/luci/common/proto/srcman/manifest.proto

package srcman

import (
	proto "github.com/golang/protobuf/proto"
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

// A Manifest attempts to make an accurate accounting of source/data directories
// during the execution of a LUCI task.
//
// These directories are primarily in the form of e.g. git checkouts of
// source, but also include things like isolated hashes and CIPD package
// deployments. In the future, other deployment forms may be supported (like
// other SCMs).
//
// The purpose of this manifest is so that other parts of the LUCI stack (e.g.
// Milo) can work with the descriptions of this deployed data as a first-class
// citizen. Initially this Manifest will be used to allow Milo to display diffs
// between jobs, but it will also be useful for tools and humans to get a
// record of exactly what data went into this LUCI task.
//
// Source Manifests can be emitted from recipes using the
// 'recipe_engine/source_manifest' module.
type Manifest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Version will increment on backwards-incompatible changes only. Backwards
	// compatible changes will not alter this version number.
	//
	// Currently, the only valid version number is 0.
	Version int32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// Map of local file system directory path (with forward slashes) to
	// a Directory message containing one or more deployments.
	//
	// The local path is relative to some job-specific root. This should be used
	// for informational/display/organization purposes. In particular, think VERY
	// CAREFULLY before you configure remote services/recipes to look for
	// particular filesystem layouts here. For example, if you want to look for
	// "the version of chromium/src checked out by the job", prefer to look for
	// a Directory which checks out "chromium/src", as opposed to assuming this
	// checkout lives in a top-level folder called "src". The reason for this is
	// that jobs SHOULD reserve the right to do their checkouts in any way they
	// please.
	//
	// If you feel like you need to make some service configuration which uses one
	// of these local filesystem paths as a key, please consult with the Chrome
	// Infrastructure team to see if there's a better alternative.
	//
	// Ex.
	//   "": {...}  // root directory
	//   "src/third_party/something": {...}
	Directories map[string]*Manifest_Directory `protobuf:"bytes,2,rep,name=directories,proto3" json:"directories,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Manifest) Reset() {
	*x = Manifest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_common_proto_srcman_manifest_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Manifest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Manifest) ProtoMessage() {}

func (x *Manifest) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_common_proto_srcman_manifest_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Manifest.ProtoReflect.Descriptor instead.
func (*Manifest) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_common_proto_srcman_manifest_proto_rawDescGZIP(), []int{0}
}

func (x *Manifest) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Manifest) GetDirectories() map[string]*Manifest_Directory {
	if x != nil {
		return x.Directories
	}
	return nil
}

// Links to an externally stored Manifest proto.
type ManifestLink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The fully qualified url of the Manifest proto. It's expected that this is
	// a binary logdog stream consisting of exactly one Manifest proto. For now
	// this will always be the `logdog` uri scheme, though it's feasible to put
	// other uri schemes here later.
	//
	// Ex.
	//   logdog://logs.chromium.org/infra/build/12345/+/some/path
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// The hash of the Manifest's raw binary form (i.e. the bytes at the end of
	// `url`, without any interpretation or decoding). Milo will use this as an
	// optimization; Manifests will be interned once into Milo's datastore.
	// Future hashes which match will not be loaded from the url, but will be
	// assumed to be identical. If the sha256 doesn't match the data at the URL,
	// Milo may render this build with the wrong manifest.
	//
	// This is the raw sha256, so it must be exactly 32 bytes.
	Sha256 []byte `protobuf:"bytes,2,opt,name=sha256,proto3" json:"sha256,omitempty"`
}

func (x *ManifestLink) Reset() {
	*x = ManifestLink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_common_proto_srcman_manifest_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManifestLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManifestLink) ProtoMessage() {}

func (x *ManifestLink) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_common_proto_srcman_manifest_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManifestLink.ProtoReflect.Descriptor instead.
func (*ManifestLink) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_common_proto_srcman_manifest_proto_rawDescGZIP(), []int{1}
}

func (x *ManifestLink) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ManifestLink) GetSha256() []byte {
	if x != nil {
		return x.Sha256
	}
	return nil
}

type Manifest_GitCheckout struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The canonicalized URL of the original repo that is considered the “source
	// of truth” for the source code.
	//
	// Ex.
	//   https://chromium.googlesource.com/chromium/tools/build
	//   https://chromium.googlesource.com/infra/luci/recipes-py
	RepoUrl string `protobuf:"bytes,1,opt,name=repo_url,json=repoUrl,proto3" json:"repo_url,omitempty"`
	// If different from repo_url, this can be the URL of the repo that the source
	// was actually fetched from (i.e. a mirror).
	//
	// If this is empty, it's presumed to be equal to repo_url.
	//
	// Ex.
	//   https://github.com/luci/recipes-py
	FetchUrl string `protobuf:"bytes,2,opt,name=fetch_url,json=fetchUrl,proto3" json:"fetch_url,omitempty"`
	// The fully resolved revision (commit hash) of the source.
	//
	// This must always be a revision on the hosted repo (not any locally
	// generated commit).
	//
	// Ex.
	//   3617b0eea7ec74b8e731a23fed2f4070cbc284c4
	Revision string `protobuf:"bytes,3,opt,name=revision,proto3" json:"revision,omitempty"`
	// The ref that the task used to resolve/fetch the revision of the source
	// (if any).
	//
	// This must always be a ref on the hosted repo (not any local alias
	// like 'refs/remotes/...').
	//
	// This must always be an absolute ref (i.e. starts with 'refs/'). An
	// example of a non-absolute ref would be 'master'.
	//
	// Ex.
	//   refs/heads/master
	FetchRef string `protobuf:"bytes,4,opt,name=fetch_ref,json=fetchRef,proto3" json:"fetch_ref,omitempty"`
	// If the checkout had a CL associated with it (i.e. a gerrit commit), this
	// is the fully resolved revision (commit hash) of the CL. If there was no
	// CL, this is empty. Typically the checkout application (e.g. bot_update)
	// rebases this revision on top of the `revision` fetched above.
	//
	// If specified, this must always be a revision on the hosted repo (not any
	// locally generated commit).
	//
	// Ex.
	//   6b0b5c12443cfb93305f8d9e21f8d762c8dad9f0
	PatchRevision string `protobuf:"bytes,5,opt,name=patch_revision,json=patchRevision,proto3" json:"patch_revision,omitempty"`
	// If the checkout had a CL associated with it, this is the ref that the
	// task used to fetch patch_revision. If `patch_revision` is supplied, this
	// field is required. If there was no CL, this is empty.
	//
	// If specified, this must always be a ref on the hosted repo (not any local
	// alias like 'refs/remotes/...').
	//
	// This must always be an absolute ref (i.e. starts with 'refs/').
	//
	// Ex.
	//   refs/changes/04/511804/4
	PatchFetchRef string `protobuf:"bytes,6,opt,name=patch_fetch_ref,json=patchFetchRef,proto3" json:"patch_fetch_ref,omitempty"`
}

func (x *Manifest_GitCheckout) Reset() {
	*x = Manifest_GitCheckout{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_common_proto_srcman_manifest_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Manifest_GitCheckout) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Manifest_GitCheckout) ProtoMessage() {}

func (x *Manifest_GitCheckout) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_common_proto_srcman_manifest_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Manifest_GitCheckout.ProtoReflect.Descriptor instead.
func (*Manifest_GitCheckout) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_common_proto_srcman_manifest_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Manifest_GitCheckout) GetRepoUrl() string {
	if x != nil {
		return x.RepoUrl
	}
	return ""
}

func (x *Manifest_GitCheckout) GetFetchUrl() string {
	if x != nil {
		return x.FetchUrl
	}
	return ""
}

func (x *Manifest_GitCheckout) GetRevision() string {
	if x != nil {
		return x.Revision
	}
	return ""
}

func (x *Manifest_GitCheckout) GetFetchRef() string {
	if x != nil {
		return x.FetchRef
	}
	return ""
}

func (x *Manifest_GitCheckout) GetPatchRevision() string {
	if x != nil {
		return x.PatchRevision
	}
	return ""
}

func (x *Manifest_GitCheckout) GetPatchFetchRef() string {
	if x != nil {
		return x.PatchFetchRef
	}
	return ""
}

type Manifest_CIPDPackage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The package pattern that was given to the CIPD client (if known).
	//
	// Ex.
	//   infra/tools/luci/led/${platform}
	PackagePattern string `protobuf:"bytes,1,opt,name=package_pattern,json=packagePattern,proto3" json:"package_pattern,omitempty"`
	// The fully resolved instance ID of the deployed package.
	//
	// Ex.
	//   0cfafb3a705bd8f05f86c6444ff500397fbb711c
	InstanceId string `protobuf:"bytes,2,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	// The unresolved version ID of the deployed package (if known).
	//
	// Ex.
	//   git_revision:aaf3a2cfccc227b5141caa1b6b3502c9907d7420
	//   latest
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *Manifest_CIPDPackage) Reset() {
	*x = Manifest_CIPDPackage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_common_proto_srcman_manifest_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Manifest_CIPDPackage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Manifest_CIPDPackage) ProtoMessage() {}

func (x *Manifest_CIPDPackage) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_common_proto_srcman_manifest_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Manifest_CIPDPackage.ProtoReflect.Descriptor instead.
func (*Manifest_CIPDPackage) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_common_proto_srcman_manifest_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Manifest_CIPDPackage) GetPackagePattern() string {
	if x != nil {
		return x.PackagePattern
	}
	return ""
}

func (x *Manifest_CIPDPackage) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

func (x *Manifest_CIPDPackage) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type Manifest_Isolated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The namespace of the isolated document.
	//
	// Ex.
	//   default-gzip
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// The hash of the isolated document.
	//
	// Ex.
	//   62a7df62ea122380afb306bb4d9cdac1bc7e9a96
	Hash string `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *Manifest_Isolated) Reset() {
	*x = Manifest_Isolated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_common_proto_srcman_manifest_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Manifest_Isolated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Manifest_Isolated) ProtoMessage() {}

func (x *Manifest_Isolated) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_common_proto_srcman_manifest_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Manifest_Isolated.ProtoReflect.Descriptor instead.
func (*Manifest_Isolated) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_common_proto_srcman_manifest_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Manifest_Isolated) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Manifest_Isolated) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

// A Directory contains one or more descriptions of deployed artifacts. Note
// that due to the practical nature of jobs on bots, it may be the case that
// a given directory contains e.g. a git checkout and multiple cipd packages.
type Manifest_Directory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GitCheckout *Manifest_GitCheckout `protobuf:"bytes,1,opt,name=git_checkout,json=gitCheckout,proto3" json:"git_checkout,omitempty"`
	// The canonicalized hostname of the CIPD server which hosts the CIPD
	// packages (if any). If no CIPD packages are in this Directory, this must
	// be blank.
	//
	// Ex.
	//   chrome-infra-packages.appspot.com
	CipdServerHost string `protobuf:"bytes,2,opt,name=cipd_server_host,json=cipdServerHost,proto3" json:"cipd_server_host,omitempty"`
	// Maps CIPD package name to CIPDPackage.
	//
	// Ex.
	//   "some/package/name": {...}
	//   "other/package": {...}
	CipdPackage map[string]*Manifest_CIPDPackage `protobuf:"bytes,4,rep,name=cipd_package,json=cipdPackage,proto3" json:"cipd_package,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The canonicalized hostname of the isolated server which hosts the
	// isolated. If no Isolated objects are in this Directory, this must be
	// blank.
	//
	// Ex.
	//   isolateserver.appspot.com
	IsolatedServerHost string `protobuf:"bytes,5,opt,name=isolated_server_host,json=isolatedServerHost,proto3" json:"isolated_server_host,omitempty"`
	// A list of all isolateds which have been installed in this directory.
	Isolated []*Manifest_Isolated `protobuf:"bytes,6,rep,name=isolated,proto3" json:"isolated,omitempty"`
}

func (x *Manifest_Directory) Reset() {
	*x = Manifest_Directory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_common_proto_srcman_manifest_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Manifest_Directory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Manifest_Directory) ProtoMessage() {}

func (x *Manifest_Directory) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_common_proto_srcman_manifest_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Manifest_Directory.ProtoReflect.Descriptor instead.
func (*Manifest_Directory) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_common_proto_srcman_manifest_proto_rawDescGZIP(), []int{0, 3}
}

func (x *Manifest_Directory) GetGitCheckout() *Manifest_GitCheckout {
	if x != nil {
		return x.GitCheckout
	}
	return nil
}

func (x *Manifest_Directory) GetCipdServerHost() string {
	if x != nil {
		return x.CipdServerHost
	}
	return ""
}

func (x *Manifest_Directory) GetCipdPackage() map[string]*Manifest_CIPDPackage {
	if x != nil {
		return x.CipdPackage
	}
	return nil
}

func (x *Manifest_Directory) GetIsolatedServerHost() string {
	if x != nil {
		return x.IsolatedServerHost
	}
	return ""
}

func (x *Manifest_Directory) GetIsolated() []*Manifest_Isolated {
	if x != nil {
		return x.Isolated
	}
	return nil
}

var File_go_chromium_org_luci_common_proto_srcman_manifest_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_common_proto_srcman_manifest_proto_rawDesc = []byte{
	0x0a, 0x37, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x72, 0x63, 0x6d, 0x61, 0x6e, 0x2f, 0x6d, 0x61, 0x6e, 0x69, 0x66,
	0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x72, 0x63, 0x6d, 0x61,
	0x6e, 0x22, 0xd6, 0x07, 0x0a, 0x08, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x43, 0x0a, 0x0b, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x73, 0x72, 0x63, 0x6d, 0x61, 0x6e, 0x2e, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x2e,
	0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0b, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x1a, 0xcd, 0x01,
	0x0a, 0x0b, 0x47, 0x69, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x72, 0x65, 0x70, 0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x72, 0x65, 0x70, 0x6f, 0x55, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x65, 0x74, 0x63,
	0x68, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x65, 0x74,
	0x63, 0x68, 0x55, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x65, 0x74, 0x63, 0x68, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x66, 0x12, 0x25,
	0x0a, 0x0e, 0x70, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x76,
	0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0f, 0x70, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x66,
	0x65, 0x74, 0x63, 0x68, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x70, 0x61, 0x74, 0x63, 0x68, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x66, 0x1a, 0x71, 0x0a,
	0x0b, 0x43, 0x49, 0x50, 0x44, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x27, 0x0a, 0x0f,
	0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x50, 0x61,
	0x74, 0x74, 0x65, 0x72, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x1a, 0x3c, 0x0a, 0x08, 0x49, 0x73, 0x6f, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61,
	0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x1a, 0x8d,
	0x03, 0x0a, 0x09, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x3f, 0x0a, 0x0c,
	0x67, 0x69, 0x74, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x72, 0x63, 0x6d, 0x61, 0x6e, 0x2e, 0x4d, 0x61, 0x6e, 0x69,
	0x66, 0x65, 0x73, 0x74, 0x2e, 0x47, 0x69, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74,
	0x52, 0x0b, 0x67, 0x69, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x12, 0x28, 0x0a,
	0x10, 0x63, 0x69, 0x70, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x68, 0x6f, 0x73,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x69, 0x70, 0x64, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x4e, 0x0a, 0x0c, 0x63, 0x69, 0x70, 0x64, 0x5f,
	0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e,
	0x73, 0x72, 0x63, 0x6d, 0x61, 0x6e, 0x2e, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x2e,
	0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x43, 0x69, 0x70, 0x64, 0x50, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x63, 0x69, 0x70, 0x64,
	0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x69, 0x73, 0x6f, 0x6c, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x69, 0x73, 0x6f, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x08, 0x69, 0x73, 0x6f,
	0x6c, 0x61, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x72,
	0x63, 0x6d, 0x61, 0x6e, 0x2e, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x2e, 0x49, 0x73,
	0x6f, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x52, 0x08, 0x69, 0x73, 0x6f, 0x6c, 0x61, 0x74, 0x65, 0x64,
	0x1a, 0x5c, 0x0a, 0x10, 0x43, 0x69, 0x70, 0x64, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x32, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x72, 0x63, 0x6d, 0x61, 0x6e, 0x2e, 0x4d,
	0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x49, 0x50, 0x44, 0x50, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x5a,
	0x0a, 0x10, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x30, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x72, 0x63, 0x6d, 0x61, 0x6e, 0x2e, 0x4d, 0x61, 0x6e,
	0x69, 0x66, 0x65, 0x73, 0x74, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x38, 0x0a, 0x0c, 0x4d, 0x61,
	0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x68, 0x61, 0x32, 0x35, 0x36, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x73, 0x68,
	0x61, 0x32, 0x35, 0x36, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d,
	0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x72, 0x63, 0x6d, 0x61, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_common_proto_srcman_manifest_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_common_proto_srcman_manifest_proto_rawDescData = file_go_chromium_org_luci_common_proto_srcman_manifest_proto_rawDesc
)

func file_go_chromium_org_luci_common_proto_srcman_manifest_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_common_proto_srcman_manifest_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_common_proto_srcman_manifest_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_common_proto_srcman_manifest_proto_rawDescData)
	})
	return file_go_chromium_org_luci_common_proto_srcman_manifest_proto_rawDescData
}

var file_go_chromium_org_luci_common_proto_srcman_manifest_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_go_chromium_org_luci_common_proto_srcman_manifest_proto_goTypes = []interface{}{
	(*Manifest)(nil),             // 0: srcman.Manifest
	(*ManifestLink)(nil),         // 1: srcman.ManifestLink
	(*Manifest_GitCheckout)(nil), // 2: srcman.Manifest.GitCheckout
	(*Manifest_CIPDPackage)(nil), // 3: srcman.Manifest.CIPDPackage
	(*Manifest_Isolated)(nil),    // 4: srcman.Manifest.Isolated
	(*Manifest_Directory)(nil),   // 5: srcman.Manifest.Directory
	nil,                          // 6: srcman.Manifest.DirectoriesEntry
	nil,                          // 7: srcman.Manifest.Directory.CipdPackageEntry
}
var file_go_chromium_org_luci_common_proto_srcman_manifest_proto_depIdxs = []int32{
	6, // 0: srcman.Manifest.directories:type_name -> srcman.Manifest.DirectoriesEntry
	2, // 1: srcman.Manifest.Directory.git_checkout:type_name -> srcman.Manifest.GitCheckout
	7, // 2: srcman.Manifest.Directory.cipd_package:type_name -> srcman.Manifest.Directory.CipdPackageEntry
	4, // 3: srcman.Manifest.Directory.isolated:type_name -> srcman.Manifest.Isolated
	5, // 4: srcman.Manifest.DirectoriesEntry.value:type_name -> srcman.Manifest.Directory
	3, // 5: srcman.Manifest.Directory.CipdPackageEntry.value:type_name -> srcman.Manifest.CIPDPackage
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_common_proto_srcman_manifest_proto_init() }
func file_go_chromium_org_luci_common_proto_srcman_manifest_proto_init() {
	if File_go_chromium_org_luci_common_proto_srcman_manifest_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_common_proto_srcman_manifest_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Manifest); i {
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
		file_go_chromium_org_luci_common_proto_srcman_manifest_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManifestLink); i {
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
		file_go_chromium_org_luci_common_proto_srcman_manifest_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Manifest_GitCheckout); i {
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
		file_go_chromium_org_luci_common_proto_srcman_manifest_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Manifest_CIPDPackage); i {
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
		file_go_chromium_org_luci_common_proto_srcman_manifest_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Manifest_Isolated); i {
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
		file_go_chromium_org_luci_common_proto_srcman_manifest_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Manifest_Directory); i {
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
			RawDescriptor: file_go_chromium_org_luci_common_proto_srcman_manifest_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_common_proto_srcman_manifest_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_common_proto_srcman_manifest_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_common_proto_srcman_manifest_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_common_proto_srcman_manifest_proto = out.File
	file_go_chromium_org_luci_common_proto_srcman_manifest_proto_rawDesc = nil
	file_go_chromium_org_luci_common_proto_srcman_manifest_proto_goTypes = nil
	file_go_chromium_org_luci_common_proto_srcman_manifest_proto_depIdxs = nil
}
