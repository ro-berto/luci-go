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
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: go.chromium.org/luci/resultdb/proto/config/project_config.proto

package resultpb

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

// ProjectConfig is the project-specific configuration data for Luci ResultDB.
type ProjectConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Per realm allow list to control GCS artifacts that could be uploaded to
	// ResultDB for the associated realm.
	// Since ResultDB returns GCS artifacts through signed urls, the allow list
	// is needed to prevent potential exploit where user could gain access to
	// artifacts in GCS locations they don't have access to by feigning the
	// uploaded artifact GCS path.
	RealmGcsAllowlist []*RealmGcsAllowList `protobuf:"bytes,1,rep,name=realm_gcs_allowlist,json=realmGcsAllowlist,proto3" json:"realm_gcs_allowlist,omitempty"`
}

func (x *ProjectConfig) Reset() {
	*x = ProjectConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_resultdb_proto_config_project_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectConfig) ProtoMessage() {}

func (x *ProjectConfig) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_resultdb_proto_config_project_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectConfig.ProtoReflect.Descriptor instead.
func (*ProjectConfig) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_resultdb_proto_config_project_config_proto_rawDescGZIP(), []int{0}
}

func (x *ProjectConfig) GetRealmGcsAllowlist() []*RealmGcsAllowList {
	if x != nil {
		return x.RealmGcsAllowlist
	}
	return nil
}

// Capture the per realm GCS artifact allow list.
type RealmGcsAllowList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Realm name. e.g. testplatform, cq, etc.
	// Note that the realm is implicitly scoped to the project the config is
	// defined.
	Realm string `protobuf:"bytes,1,opt,name=realm,proto3" json:"realm,omitempty"`
	// Allowed GCS bucket prefixes associated with the realm.
	// Each bucket should be an entry along with allowed object prefixes for the
	// bucket.
	// There should not be multiple entries for the same bucket.
	GcsBucketPrefixes []*GcsBucketPrefixes `protobuf:"bytes,2,rep,name=gcs_bucket_prefixes,json=gcsBucketPrefixes,proto3" json:"gcs_bucket_prefixes,omitempty"`
}

func (x *RealmGcsAllowList) Reset() {
	*x = RealmGcsAllowList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_resultdb_proto_config_project_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RealmGcsAllowList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RealmGcsAllowList) ProtoMessage() {}

func (x *RealmGcsAllowList) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_resultdb_proto_config_project_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RealmGcsAllowList.ProtoReflect.Descriptor instead.
func (*RealmGcsAllowList) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_resultdb_proto_config_project_config_proto_rawDescGZIP(), []int{1}
}

func (x *RealmGcsAllowList) GetRealm() string {
	if x != nil {
		return x.Realm
	}
	return ""
}

func (x *RealmGcsAllowList) GetGcsBucketPrefixes() []*GcsBucketPrefixes {
	if x != nil {
		return x.GcsBucketPrefixes
	}
	return nil
}

// Capture a GCS bucket along with the allowed object prefixes.
type GcsBucketPrefixes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// GCS bucket name.
	// e.g. chromeos-test-logs
	Bucket string `protobuf:"bytes,1,opt,name=bucket,proto3" json:"bucket,omitempty"`
	// Object prefixes that are allowed for the bucket.
	// Wildcard can be used to indicate everything e.g. '*'.
	AllowedPrefixes []string `protobuf:"bytes,2,rep,name=allowed_prefixes,json=allowedPrefixes,proto3" json:"allowed_prefixes,omitempty"`
}

func (x *GcsBucketPrefixes) Reset() {
	*x = GcsBucketPrefixes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_resultdb_proto_config_project_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GcsBucketPrefixes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcsBucketPrefixes) ProtoMessage() {}

func (x *GcsBucketPrefixes) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_resultdb_proto_config_project_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcsBucketPrefixes.ProtoReflect.Descriptor instead.
func (*GcsBucketPrefixes) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_resultdb_proto_config_project_config_proto_rawDescGZIP(), []int{2}
}

func (x *GcsBucketPrefixes) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *GcsBucketPrefixes) GetAllowedPrefixes() []string {
	if x != nil {
		return x.AllowedPrefixes
	}
	return nil
}

var File_go_chromium_org_luci_resultdb_proto_config_project_config_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_resultdb_proto_config_project_config_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x64, 0x62, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x14, 0x6c, 0x75, 0x63, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x64, 0x62,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x68, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x57, 0x0a, 0x13, 0x72, 0x65, 0x61, 0x6c,
	0x6d, 0x5f, 0x67, 0x63, 0x73, 0x5f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6c, 0x75, 0x63, 0x69, 0x2e, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x64, 0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x52, 0x65, 0x61,
	0x6c, 0x6d, 0x47, 0x63, 0x73, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x11,
	0x72, 0x65, 0x61, 0x6c, 0x6d, 0x47, 0x63, 0x73, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x6c, 0x69, 0x73,
	0x74, 0x22, 0x82, 0x01, 0x0a, 0x11, 0x52, 0x65, 0x61, 0x6c, 0x6d, 0x47, 0x63, 0x73, 0x41, 0x6c,
	0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x61, 0x6c, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x12, 0x57, 0x0a,
	0x13, 0x67, 0x63, 0x73, 0x5f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x70, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6c, 0x75, 0x63,
	0x69, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x64, 0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x47, 0x63, 0x73, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x50, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x65, 0x73, 0x52, 0x11, 0x67, 0x63, 0x73, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x50, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x65, 0x73, 0x22, 0x56, 0x0a, 0x11, 0x47, 0x63, 0x73, 0x42, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x62,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x70,
	0x72, 0x65, 0x66, 0x69, 0x78, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x61,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x65, 0x73, 0x42, 0x35,
	0x5a, 0x33, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x64, 0x62, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x3b, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_resultdb_proto_config_project_config_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_resultdb_proto_config_project_config_proto_rawDescData = file_go_chromium_org_luci_resultdb_proto_config_project_config_proto_rawDesc
)

func file_go_chromium_org_luci_resultdb_proto_config_project_config_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_resultdb_proto_config_project_config_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_resultdb_proto_config_project_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_resultdb_proto_config_project_config_proto_rawDescData)
	})
	return file_go_chromium_org_luci_resultdb_proto_config_project_config_proto_rawDescData
}

var file_go_chromium_org_luci_resultdb_proto_config_project_config_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_go_chromium_org_luci_resultdb_proto_config_project_config_proto_goTypes = []interface{}{
	(*ProjectConfig)(nil),     // 0: luci.resultdb.config.ProjectConfig
	(*RealmGcsAllowList)(nil), // 1: luci.resultdb.config.RealmGcsAllowList
	(*GcsBucketPrefixes)(nil), // 2: luci.resultdb.config.GcsBucketPrefixes
}
var file_go_chromium_org_luci_resultdb_proto_config_project_config_proto_depIdxs = []int32{
	1, // 0: luci.resultdb.config.ProjectConfig.realm_gcs_allowlist:type_name -> luci.resultdb.config.RealmGcsAllowList
	2, // 1: luci.resultdb.config.RealmGcsAllowList.gcs_bucket_prefixes:type_name -> luci.resultdb.config.GcsBucketPrefixes
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_resultdb_proto_config_project_config_proto_init() }
func file_go_chromium_org_luci_resultdb_proto_config_project_config_proto_init() {
	if File_go_chromium_org_luci_resultdb_proto_config_project_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_resultdb_proto_config_project_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectConfig); i {
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
		file_go_chromium_org_luci_resultdb_proto_config_project_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RealmGcsAllowList); i {
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
		file_go_chromium_org_luci_resultdb_proto_config_project_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GcsBucketPrefixes); i {
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
			RawDescriptor: file_go_chromium_org_luci_resultdb_proto_config_project_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_resultdb_proto_config_project_config_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_resultdb_proto_config_project_config_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_resultdb_proto_config_project_config_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_resultdb_proto_config_project_config_proto = out.File
	file_go_chromium_org_luci_resultdb_proto_config_project_config_proto_rawDesc = nil
	file_go_chromium_org_luci_resultdb_proto_config_project_config_proto_goTypes = nil
	file_go_chromium_org_luci_resultdb_proto_config_project_config_proto_depIdxs = nil
}
