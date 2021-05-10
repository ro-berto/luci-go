// Copyright 2020 The LUCI Authors.
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
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.1
// source: go.chromium.org/luci/resultdb/proto/v1/artifact.proto

package resultpb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// A file produced during a build/test, typically a test artifact.
// The parent resource is either a TestResult or an Invocation.
//
// An invocation-level artifact might be related to tests, or it might not, for
// example it may be used to store build step logs when streaming support is
// added.
type Artifact struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Can be used to refer to this artifact.
	// Format:
	// - For invocation-level artifacts:
	//   "invocations/{INVOCATION_ID}/artifacts/{ARTIFACT_ID}".
	// - For test-result-level artifacts:
	//   "invocations/{INVOCATION_ID}/tests/{URL_ESCAPED_TEST_ID}/results/{RESULT_ID}/artifacts/{ARTIFACT_ID}".
	// where URL_ESCAPED_TEST_ID is the test_id escaped with
	// https://golang.org/pkg/net/url/#PathEscape (see also https://aip.dev/122),
	// and ARTIFACT_ID is documented below.
	// Examples: "screenshot.png", "traces/a.txt".
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A local identifier of the artifact, unique within the parent resource.
	// MAY have slashes, but MUST NOT start with a slash.
	// SHOULD not use backslashes.
	// Regex: ^[[:word:]]([[:print:]]{0,254}[[:word:]])?$
	ArtifactId string `protobuf:"bytes,2,opt,name=artifact_id,json=artifactId,proto3" json:"artifact_id,omitempty"`
	// A signed short-lived URL to fetch the contents of the artifact.
	// See also fetch_url_expiration.
	FetchUrl string `protobuf:"bytes,3,opt,name=fetch_url,json=fetchUrl,proto3" json:"fetch_url,omitempty"`
	// When fetch_url expires. If expired, re-request this Artifact.
	FetchUrlExpiration *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=fetch_url_expiration,json=fetchUrlExpiration,proto3" json:"fetch_url_expiration,omitempty"`
	// Media type of the artifact.
	// Logs are typically "text/plain" and screenshots are typically "image/png".
	// Optional.
	ContentType string `protobuf:"bytes,5,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	// Size of the file.
	// Can be used in UI to decide between displaying the artifact inline or only
	// showing a link if it is too large.
	SizeBytes int64 `protobuf:"varint,6,opt,name=size_bytes,json=sizeBytes,proto3" json:"size_bytes,omitempty"`
	// Contents of the artifact.
	// This is INPUT_ONLY, and taken by BatchCreateArtifacts().
	// All getter RPCs, such as ListArtifacts(), do not populate values into
	// the field in the response.
	Contents []byte `protobuf:"bytes,7,opt,name=contents,proto3" json:"contents,omitempty"`
}

func (x *Artifact) Reset() {
	*x = Artifact{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_resultdb_proto_v1_artifact_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Artifact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Artifact) ProtoMessage() {}

func (x *Artifact) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_resultdb_proto_v1_artifact_proto_msgTypes[0]
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
	return file_go_chromium_org_luci_resultdb_proto_v1_artifact_proto_rawDescGZIP(), []int{0}
}

func (x *Artifact) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Artifact) GetArtifactId() string {
	if x != nil {
		return x.ArtifactId
	}
	return ""
}

func (x *Artifact) GetFetchUrl() string {
	if x != nil {
		return x.FetchUrl
	}
	return ""
}

func (x *Artifact) GetFetchUrlExpiration() *timestamppb.Timestamp {
	if x != nil {
		return x.FetchUrlExpiration
	}
	return nil
}

func (x *Artifact) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *Artifact) GetSizeBytes() int64 {
	if x != nil {
		return x.SizeBytes
	}
	return 0
}

func (x *Artifact) GetContents() []byte {
	if x != nil {
		return x.Contents
	}
	return nil
}

var File_go_chromium_org_luci_resultdb_proto_v1_artifact_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_resultdb_proto_v1_artifact_proto_rawDesc = []byte{
	0x0a, 0x35, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x64, 0x62, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x6c, 0x75, 0x63, 0x69, 0x2e, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61,
	0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x02, 0x0a, 0x08,
	0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x66, 0x65, 0x74, 0x63, 0x68, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x66, 0x65, 0x74, 0x63, 0x68, 0x55, 0x72, 0x6c, 0x12, 0x4c, 0x0a, 0x14, 0x66, 0x65,
	0x74, 0x63, 0x68, 0x5f, 0x75, 0x72, 0x6c, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x12, 0x66, 0x65, 0x74, 0x63, 0x68, 0x55, 0x72, 0x6c, 0x45, 0x78,
	0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x69, 0x7a, 0x65, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x73, 0x69, 0x7a, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x08, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x03, 0xe0, 0x41,
	0x04, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x31, 0x5a, 0x2f, 0x67,
	0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c,
	0x75, 0x63, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x64, 0x62, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x76, 0x31, 0x3b, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_resultdb_proto_v1_artifact_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_resultdb_proto_v1_artifact_proto_rawDescData = file_go_chromium_org_luci_resultdb_proto_v1_artifact_proto_rawDesc
)

func file_go_chromium_org_luci_resultdb_proto_v1_artifact_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_resultdb_proto_v1_artifact_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_resultdb_proto_v1_artifact_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_resultdb_proto_v1_artifact_proto_rawDescData)
	})
	return file_go_chromium_org_luci_resultdb_proto_v1_artifact_proto_rawDescData
}

var file_go_chromium_org_luci_resultdb_proto_v1_artifact_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_go_chromium_org_luci_resultdb_proto_v1_artifact_proto_goTypes = []interface{}{
	(*Artifact)(nil),              // 0: luci.resultdb.v1.Artifact
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_go_chromium_org_luci_resultdb_proto_v1_artifact_proto_depIdxs = []int32{
	1, // 0: luci.resultdb.v1.Artifact.fetch_url_expiration:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_resultdb_proto_v1_artifact_proto_init() }
func file_go_chromium_org_luci_resultdb_proto_v1_artifact_proto_init() {
	if File_go_chromium_org_luci_resultdb_proto_v1_artifact_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_resultdb_proto_v1_artifact_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_go_chromium_org_luci_resultdb_proto_v1_artifact_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_resultdb_proto_v1_artifact_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_resultdb_proto_v1_artifact_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_resultdb_proto_v1_artifact_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_resultdb_proto_v1_artifact_proto = out.File
	file_go_chromium_org_luci_resultdb_proto_v1_artifact_proto_rawDesc = nil
	file_go_chromium_org_luci_resultdb_proto_v1_artifact_proto_goTypes = nil
	file_go_chromium_org_luci_resultdb_proto_v1_artifact_proto_depIdxs = nil
}
