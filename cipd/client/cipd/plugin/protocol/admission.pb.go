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
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: go.chromium.org/luci/cipd/client/cipd/plugin/protocol/admission.proto

package protocol

import (
	v1 "go.chromium.org/luci/cipd/api/cipd/v1"
	status "google.golang.org/genproto/googleapis/rpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ListAdmissionsRequest carries arguments for ListAdmissions RPC.
type ListAdmissionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProtocolVersion int32  `protobuf:"varint,1,opt,name=protocol_version,json=protocolVersion,proto3" json:"protocol_version,omitempty"` // currently should be 1
	PluginVersion   string `protobuf:"bytes,2,opt,name=plugin_version,json=pluginVersion,proto3" json:"plugin_version,omitempty"`        // arbitrary string for logs
}

func (x *ListAdmissionsRequest) Reset() {
	*x = ListAdmissionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cipd_client_cipd_plugin_protocol_admission_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAdmissionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAdmissionsRequest) ProtoMessage() {}

func (x *ListAdmissionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cipd_client_cipd_plugin_protocol_admission_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAdmissionsRequest.ProtoReflect.Descriptor instead.
func (*ListAdmissionsRequest) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cipd_client_cipd_plugin_protocol_admission_proto_rawDescGZIP(), []int{0}
}

func (x *ListAdmissionsRequest) GetProtocolVersion() int32 {
	if x != nil {
		return x.ProtocolVersion
	}
	return 0
}

func (x *ListAdmissionsRequest) GetPluginVersion() string {
	if x != nil {
		return x.PluginVersion
	}
	return ""
}

// Admission is sent by CIPD client when it attempts to install a package.
//
// The plugin may allow or forbid this via ResolveAdmission RPC.
type Admission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdmissionId string        `protobuf:"bytes,1,opt,name=admission_id,json=admissionId,proto3" json:"admission_id,omitempty"` // an opaque ID identifies this request
	ServiceUrl  string        `protobuf:"bytes,2,opt,name=service_url,json=serviceUrl,proto3" json:"service_url,omitempty"`    // https:// address of the CIPD backend
	Package     string        `protobuf:"bytes,3,opt,name=package,proto3" json:"package,omitempty"`                            // a package being installed
	Instance    *v1.ObjectRef `protobuf:"bytes,4,opt,name=instance,proto3" json:"instance,omitempty"`                          // a concrete package instance being installed
}

func (x *Admission) Reset() {
	*x = Admission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cipd_client_cipd_plugin_protocol_admission_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Admission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Admission) ProtoMessage() {}

func (x *Admission) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cipd_client_cipd_plugin_protocol_admission_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Admission.ProtoReflect.Descriptor instead.
func (*Admission) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cipd_client_cipd_plugin_protocol_admission_proto_rawDescGZIP(), []int{1}
}

func (x *Admission) GetAdmissionId() string {
	if x != nil {
		return x.AdmissionId
	}
	return ""
}

func (x *Admission) GetServiceUrl() string {
	if x != nil {
		return x.ServiceUrl
	}
	return ""
}

func (x *Admission) GetPackage() string {
	if x != nil {
		return x.Package
	}
	return ""
}

func (x *Admission) GetInstance() *v1.ObjectRef {
	if x != nil {
		return x.Instance
	}
	return nil
}

// ResolveAdmissionRequest carries a judgment on some Admission.
type ResolveAdmissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdmissionId string         `protobuf:"bytes,1,opt,name=admission_id,json=admissionId,proto3" json:"admission_id,omitempty"` // same as in the corresponding Admission
	Status      *status.Status `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`                              // not OK if the deployment is denied
}

func (x *ResolveAdmissionRequest) Reset() {
	*x = ResolveAdmissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cipd_client_cipd_plugin_protocol_admission_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResolveAdmissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolveAdmissionRequest) ProtoMessage() {}

func (x *ResolveAdmissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cipd_client_cipd_plugin_protocol_admission_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolveAdmissionRequest.ProtoReflect.Descriptor instead.
func (*ResolveAdmissionRequest) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cipd_client_cipd_plugin_protocol_admission_proto_rawDescGZIP(), []int{2}
}

func (x *ResolveAdmissionRequest) GetAdmissionId() string {
	if x != nil {
		return x.AdmissionId
	}
	return ""
}

func (x *ResolveAdmissionRequest) GetStatus() *status.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

var File_go_chromium_org_luci_cipd_client_cipd_plugin_protocol_admission_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_cipd_client_cipd_plugin_protocol_admission_proto_rawDesc = []byte{
	0x0a, 0x45, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x69, 0x70, 0x64, 0x2f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2f, 0x63, 0x69, 0x70, 0x64, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63, 0x69, 0x70, 0x64, 0x2e, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x67, 0x6f, 0x2e, 0x63,
	0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69,
	0x2f, 0x63, 0x69, 0x70, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x69, 0x70, 0x64, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x69, 0x0a, 0x15, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x25, 0x0a, 0x0e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x96, 0x01, 0x0a, 0x09, 0x41, 0x64, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x64, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x64, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x69, 0x70, 0x64, 0x2e, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x66, 0x52, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x22,
	0x68, 0x0a, 0x17, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x64,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x61, 0x64, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x2a, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xae, 0x01, 0x0a, 0x0a, 0x41, 0x64,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x4e, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74,
	0x41, 0x64, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x22, 0x2e, 0x63, 0x69, 0x70,
	0x64, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x64, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x63, 0x69, 0x70, 0x64, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x41, 0x64, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x30, 0x01, 0x12, 0x50, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x6f,
	0x6c, 0x76, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x2e, 0x63,
	0x69, 0x70, 0x64, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c,
	0x76, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x6f,
	0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75,
	0x63, 0x69, 0x2f, 0x63, 0x69, 0x70, 0x64, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x63,
	0x69, 0x70, 0x64, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_cipd_client_cipd_plugin_protocol_admission_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_cipd_client_cipd_plugin_protocol_admission_proto_rawDescData = file_go_chromium_org_luci_cipd_client_cipd_plugin_protocol_admission_proto_rawDesc
)

func file_go_chromium_org_luci_cipd_client_cipd_plugin_protocol_admission_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_cipd_client_cipd_plugin_protocol_admission_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_cipd_client_cipd_plugin_protocol_admission_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_cipd_client_cipd_plugin_protocol_admission_proto_rawDescData)
	})
	return file_go_chromium_org_luci_cipd_client_cipd_plugin_protocol_admission_proto_rawDescData
}

var file_go_chromium_org_luci_cipd_client_cipd_plugin_protocol_admission_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_go_chromium_org_luci_cipd_client_cipd_plugin_protocol_admission_proto_goTypes = []interface{}{
	(*ListAdmissionsRequest)(nil),   // 0: cipd.plugin.ListAdmissionsRequest
	(*Admission)(nil),               // 1: cipd.plugin.Admission
	(*ResolveAdmissionRequest)(nil), // 2: cipd.plugin.ResolveAdmissionRequest
	(*v1.ObjectRef)(nil),            // 3: cipd.ObjectRef
	(*status.Status)(nil),           // 4: google.rpc.Status
	(*emptypb.Empty)(nil),           // 5: google.protobuf.Empty
}
var file_go_chromium_org_luci_cipd_client_cipd_plugin_protocol_admission_proto_depIdxs = []int32{
	3, // 0: cipd.plugin.Admission.instance:type_name -> cipd.ObjectRef
	4, // 1: cipd.plugin.ResolveAdmissionRequest.status:type_name -> google.rpc.Status
	0, // 2: cipd.plugin.Admissions.ListAdmissions:input_type -> cipd.plugin.ListAdmissionsRequest
	2, // 3: cipd.plugin.Admissions.ResolveAdmission:input_type -> cipd.plugin.ResolveAdmissionRequest
	1, // 4: cipd.plugin.Admissions.ListAdmissions:output_type -> cipd.plugin.Admission
	5, // 5: cipd.plugin.Admissions.ResolveAdmission:output_type -> google.protobuf.Empty
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_cipd_client_cipd_plugin_protocol_admission_proto_init() }
func file_go_chromium_org_luci_cipd_client_cipd_plugin_protocol_admission_proto_init() {
	if File_go_chromium_org_luci_cipd_client_cipd_plugin_protocol_admission_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_cipd_client_cipd_plugin_protocol_admission_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAdmissionsRequest); i {
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
		file_go_chromium_org_luci_cipd_client_cipd_plugin_protocol_admission_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Admission); i {
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
		file_go_chromium_org_luci_cipd_client_cipd_plugin_protocol_admission_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResolveAdmissionRequest); i {
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
			RawDescriptor: file_go_chromium_org_luci_cipd_client_cipd_plugin_protocol_admission_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_go_chromium_org_luci_cipd_client_cipd_plugin_protocol_admission_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_cipd_client_cipd_plugin_protocol_admission_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_cipd_client_cipd_plugin_protocol_admission_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_cipd_client_cipd_plugin_protocol_admission_proto = out.File
	file_go_chromium_org_luci_cipd_client_cipd_plugin_protocol_admission_proto_rawDesc = nil
	file_go_chromium_org_luci_cipd_client_cipd_plugin_protocol_admission_proto_goTypes = nil
	file_go_chromium_org_luci_cipd_client_cipd_plugin_protocol_admission_proto_depIdxs = nil
}
