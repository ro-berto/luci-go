// Copyright 2022 The LUCI Authors.
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
// source: go.chromium.org/luci/deploy/api/modelpb/artifact.proto

package modelpb

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

// Artifact kind.
type ArtifactID_Kind int32

const (
	ArtifactID_KIND_UNSPECIFIED ArtifactID_Kind = 0
	ArtifactID_GAE_TARBALL      ArtifactID_Kind = 1 // a tarball with GAE code built by cloudbuildhelper
	ArtifactID_DOCKER_IMAGE     ArtifactID_Kind = 2 // a docker image
)

// Enum value maps for ArtifactID_Kind.
var (
	ArtifactID_Kind_name = map[int32]string{
		0: "KIND_UNSPECIFIED",
		1: "GAE_TARBALL",
		2: "DOCKER_IMAGE",
	}
	ArtifactID_Kind_value = map[string]int32{
		"KIND_UNSPECIFIED": 0,
		"GAE_TARBALL":      1,
		"DOCKER_IMAGE":     2,
	}
)

func (x ArtifactID_Kind) Enum() *ArtifactID_Kind {
	p := new(ArtifactID_Kind)
	*p = x
	return p
}

func (x ArtifactID_Kind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ArtifactID_Kind) Descriptor() protoreflect.EnumDescriptor {
	return file_go_chromium_org_luci_deploy_api_modelpb_artifact_proto_enumTypes[0].Descriptor()
}

func (ArtifactID_Kind) Type() protoreflect.EnumType {
	return &file_go_chromium_org_luci_deploy_api_modelpb_artifact_proto_enumTypes[0]
}

func (x ArtifactID_Kind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ArtifactID_Kind.Descriptor instead.
func (ArtifactID_Kind) EnumDescriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_deploy_api_modelpb_artifact_proto_rawDescGZIP(), []int{0, 0}
}

// Identifier of an artifact version.
type ArtifactID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind ArtifactID_Kind `protobuf:"varint,1,opt,name=kind,proto3,enum=deploy.model.ArtifactID_Kind" json:"kind,omitempty"`
	// Name of the artifact, e.g. a tarball path or container image name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Its version label, e.g. "47712-fe5d339".
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *ArtifactID) Reset() {
	*x = ArtifactID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_deploy_api_modelpb_artifact_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArtifactID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArtifactID) ProtoMessage() {}

func (x *ArtifactID) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_deploy_api_modelpb_artifact_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArtifactID.ProtoReflect.Descriptor instead.
func (*ArtifactID) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_deploy_api_modelpb_artifact_proto_rawDescGZIP(), []int{0}
}

func (x *ArtifactID) GetKind() ArtifactID_Kind {
	if x != nil {
		return x.Kind
	}
	return ArtifactID_KIND_UNSPECIFIED
}

func (x *ArtifactID) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ArtifactID) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

// Metadata about an artifact version.
//
// Immutable.
type Artifact struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Its full ID (including version).
	Id *ArtifactID `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Artifact URL in the storage, the format depends on the artifact kind.
	Location string `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	// Its digest as "<algo>:<hex>".
	Digest string `protobuf:"bytes,3,opt,name=digest,proto3" json:"digest,omitempty"`
	// When it was published to the IaC repository the first time.
	Published *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=published,proto3" json:"published,omitempty"`
	// Reference to the source code the artifact was built from, for change logs.
	Sources []*ArtifactSource `protobuf:"bytes,5,rep,name=sources,proto3" json:"sources,omitempty"`
	// Links to logs and UI pages related to the artifact.
	Links *ArtifactLinks `protobuf:"bytes,6,opt,name=links,proto3" json:"links,omitempty"`
}

func (x *Artifact) Reset() {
	*x = Artifact{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_deploy_api_modelpb_artifact_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Artifact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Artifact) ProtoMessage() {}

func (x *Artifact) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_deploy_api_modelpb_artifact_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Artifact.ProtoReflect.Descriptor instead.
func (*Artifact) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_deploy_api_modelpb_artifact_proto_rawDescGZIP(), []int{1}
}

func (x *Artifact) GetId() *ArtifactID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Artifact) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *Artifact) GetDigest() string {
	if x != nil {
		return x.Digest
	}
	return ""
}

func (x *Artifact) GetPublished() *timestamppb.Timestamp {
	if x != nil {
		return x.Published
	}
	return nil
}

func (x *Artifact) GetSources() []*ArtifactSource {
	if x != nil {
		return x.Sources
	}
	return nil
}

func (x *Artifact) GetLinks() *ArtifactLinks {
	if x != nil {
		return x.Links
	}
	return nil
}

// Reference to the source code to calculate change logs.
//
// It is not a full manifest, just "interesting" directories worthy of inclusion
// in the change log.
type ArtifactSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Full git repository URL.
	Repository string `protobuf:"bytes,1,opt,name=repository,proto3" json:"repository,omitempty"`
	// Git revision.
	Revision string `protobuf:"bytes,2,opt,name=revision,proto3" json:"revision,omitempty"`
	// List of the directories inside this repo with sources to include.
	Sources []string `protobuf:"bytes,3,rep,name=sources,proto3" json:"sources,omitempty"`
}

func (x *ArtifactSource) Reset() {
	*x = ArtifactSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_deploy_api_modelpb_artifact_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArtifactSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArtifactSource) ProtoMessage() {}

func (x *ArtifactSource) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_deploy_api_modelpb_artifact_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArtifactSource.ProtoReflect.Descriptor instead.
func (*ArtifactSource) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_deploy_api_modelpb_artifact_proto_rawDescGZIP(), []int{2}
}

func (x *ArtifactSource) GetRepository() string {
	if x != nil {
		return x.Repository
	}
	return ""
}

func (x *ArtifactSource) GetRevision() string {
	if x != nil {
		return x.Revision
	}
	return ""
}

func (x *ArtifactSource) GetSources() []string {
	if x != nil {
		return x.Sources
	}
	return nil
}

// Links to human-readable logs and UI pages related to the artifact.
//
// All are optional.
type ArtifactLinks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Link to where the artifact is defined (e.g. its build configs).
	Definition string `protobuf:"bytes,1,opt,name=definition,proto3" json:"definition,omitempty"`
	// Link to view the artifact via its storage UI (e.g. GCR Cloud Console link).
	View string `protobuf:"bytes,2,opt,name=view,proto3" json:"view,omitempty"`
	// Link to the buildbucket build that produced this artifact.
	Buildbucket string `protobuf:"bytes,3,opt,name=buildbucket,proto3" json:"buildbucket,omitempty"`
	// Link to the cloudbuild build that produced this artifact.
	Cloudbuild string `protobuf:"bytes,4,opt,name=cloudbuild,proto3" json:"cloudbuild,omitempty"`
}

func (x *ArtifactLinks) Reset() {
	*x = ArtifactLinks{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_deploy_api_modelpb_artifact_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArtifactLinks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArtifactLinks) ProtoMessage() {}

func (x *ArtifactLinks) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_deploy_api_modelpb_artifact_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArtifactLinks.ProtoReflect.Descriptor instead.
func (*ArtifactLinks) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_deploy_api_modelpb_artifact_proto_rawDescGZIP(), []int{3}
}

func (x *ArtifactLinks) GetDefinition() string {
	if x != nil {
		return x.Definition
	}
	return ""
}

func (x *ArtifactLinks) GetView() string {
	if x != nil {
		return x.View
	}
	return ""
}

func (x *ArtifactLinks) GetBuildbucket() string {
	if x != nil {
		return x.Buildbucket
	}
	return ""
}

func (x *ArtifactLinks) GetCloudbuild() string {
	if x != nil {
		return x.Cloudbuild
	}
	return ""
}

var File_go_chromium_org_luci_deploy_api_modelpb_artifact_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_deploy_api_modelpb_artifact_proto_rawDesc = []byte{
	0x0a, 0x36, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x70, 0x62, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xae, 0x01, 0x0a, 0x0a, 0x41, 0x72, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x74, 0x49, 0x44, 0x12, 0x31, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x49, 0x44, 0x2e, 0x4b,
	0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x3f, 0x0a, 0x04, 0x4b, 0x69, 0x6e, 0x64, 0x12,
	0x14, 0x0a, 0x10, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x47, 0x41, 0x45, 0x5f, 0x54, 0x41, 0x52,
	0x42, 0x41, 0x4c, 0x4c, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x44, 0x4f, 0x43, 0x4b, 0x45, 0x52,
	0x5f, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x10, 0x02, 0x22, 0x8d, 0x02, 0x0a, 0x08, 0x41, 0x72, 0x74,
	0x69, 0x66, 0x61, 0x63, 0x74, 0x12, 0x28, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x64,
	0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x69, 0x67,
	0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x12, 0x36, 0x0a,
	0x07, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x72,
	0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x07, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x31, 0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x4c, 0x69, 0x6e, 0x6b,
	0x73, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x22, 0x66, 0x0a, 0x0e, 0x41, 0x72, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x22, 0x85, 0x01, 0x0a, 0x0d, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x4c, 0x69, 0x6e,
	0x6b, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x76, 0x69, 0x65, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x76, 0x69, 0x65, 0x77, 0x12, 0x20, 0x0a, 0x0b, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x6f, 0x2e, 0x63,
	0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69,
	0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_deploy_api_modelpb_artifact_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_deploy_api_modelpb_artifact_proto_rawDescData = file_go_chromium_org_luci_deploy_api_modelpb_artifact_proto_rawDesc
)

func file_go_chromium_org_luci_deploy_api_modelpb_artifact_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_deploy_api_modelpb_artifact_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_deploy_api_modelpb_artifact_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_deploy_api_modelpb_artifact_proto_rawDescData)
	})
	return file_go_chromium_org_luci_deploy_api_modelpb_artifact_proto_rawDescData
}

var file_go_chromium_org_luci_deploy_api_modelpb_artifact_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_go_chromium_org_luci_deploy_api_modelpb_artifact_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_go_chromium_org_luci_deploy_api_modelpb_artifact_proto_goTypes = []interface{}{
	(ArtifactID_Kind)(0),          // 0: deploy.model.ArtifactID.Kind
	(*ArtifactID)(nil),            // 1: deploy.model.ArtifactID
	(*Artifact)(nil),              // 2: deploy.model.Artifact
	(*ArtifactSource)(nil),        // 3: deploy.model.ArtifactSource
	(*ArtifactLinks)(nil),         // 4: deploy.model.ArtifactLinks
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_go_chromium_org_luci_deploy_api_modelpb_artifact_proto_depIdxs = []int32{
	0, // 0: deploy.model.ArtifactID.kind:type_name -> deploy.model.ArtifactID.Kind
	1, // 1: deploy.model.Artifact.id:type_name -> deploy.model.ArtifactID
	5, // 2: deploy.model.Artifact.published:type_name -> google.protobuf.Timestamp
	3, // 3: deploy.model.Artifact.sources:type_name -> deploy.model.ArtifactSource
	4, // 4: deploy.model.Artifact.links:type_name -> deploy.model.ArtifactLinks
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_deploy_api_modelpb_artifact_proto_init() }
func file_go_chromium_org_luci_deploy_api_modelpb_artifact_proto_init() {
	if File_go_chromium_org_luci_deploy_api_modelpb_artifact_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_deploy_api_modelpb_artifact_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArtifactID); i {
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
		file_go_chromium_org_luci_deploy_api_modelpb_artifact_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Artifact); i {
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
		file_go_chromium_org_luci_deploy_api_modelpb_artifact_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArtifactSource); i {
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
		file_go_chromium_org_luci_deploy_api_modelpb_artifact_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArtifactLinks); i {
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
			RawDescriptor: file_go_chromium_org_luci_deploy_api_modelpb_artifact_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_deploy_api_modelpb_artifact_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_deploy_api_modelpb_artifact_proto_depIdxs,
		EnumInfos:         file_go_chromium_org_luci_deploy_api_modelpb_artifact_proto_enumTypes,
		MessageInfos:      file_go_chromium_org_luci_deploy_api_modelpb_artifact_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_deploy_api_modelpb_artifact_proto = out.File
	file_go_chromium_org_luci_deploy_api_modelpb_artifact_proto_rawDesc = nil
	file_go_chromium_org_luci_deploy_api_modelpb_artifact_proto_goTypes = nil
	file_go_chromium_org_luci_deploy_api_modelpb_artifact_proto_depIdxs = nil
}
