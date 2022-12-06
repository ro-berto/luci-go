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
// source: go.chromium.org/luci/provenance/api/speepb/v1/network_proxy.proto

package speepb

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

// NetworkActivityLog records a network request made by the build.
type NetworkActivityLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HttpRequests    []*NetworkActivityLog_HttpRequest    `protobuf:"bytes,1,rep,name=http_requests,json=httpRequests,proto3" json:"http_requests,omitempty"`
	RawRequests     []*NetworkActivityLog_RawRequest     `protobuf:"bytes,2,rep,name=raw_requests,json=rawRequests,proto3" json:"raw_requests,omitempty"`
	PackageInstalls []*NetworkActivityLog_PackageInstall `protobuf:"bytes,3,rep,name=package_installs,json=packageInstalls,proto3" json:"package_installs,omitempty"`
}

func (x *NetworkActivityLog) Reset() {
	*x = NetworkActivityLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_provenance_api_speepb_v1_network_proxy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkActivityLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkActivityLog) ProtoMessage() {}

func (x *NetworkActivityLog) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_provenance_api_speepb_v1_network_proxy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkActivityLog.ProtoReflect.Descriptor instead.
func (*NetworkActivityLog) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_provenance_api_speepb_v1_network_proxy_proto_rawDescGZIP(), []int{0}
}

func (x *NetworkActivityLog) GetHttpRequests() []*NetworkActivityLog_HttpRequest {
	if x != nil {
		return x.HttpRequests
	}
	return nil
}

func (x *NetworkActivityLog) GetRawRequests() []*NetworkActivityLog_RawRequest {
	if x != nil {
		return x.RawRequests
	}
	return nil
}

func (x *NetworkActivityLog) GetPackageInstalls() []*NetworkActivityLog_PackageInstall {
	if x != nil {
		return x.PackageInstalls
	}
	return nil
}

type NetworkActivityLog_HttpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method *string `protobuf:"bytes,1,opt,name=method,proto3,oneof" json:"method,omitempty"`
	Scheme *string `protobuf:"bytes,2,opt,name=scheme,proto3,oneof" json:"scheme,omitempty"`
	Host   *string `protobuf:"bytes,3,opt,name=host,proto3,oneof" json:"host,omitempty"`
	Path   *string `protobuf:"bytes,4,opt,name=path,proto3,oneof" json:"path,omitempty"`
}

func (x *NetworkActivityLog_HttpRequest) Reset() {
	*x = NetworkActivityLog_HttpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_provenance_api_speepb_v1_network_proxy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkActivityLog_HttpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkActivityLog_HttpRequest) ProtoMessage() {}

func (x *NetworkActivityLog_HttpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_provenance_api_speepb_v1_network_proxy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkActivityLog_HttpRequest.ProtoReflect.Descriptor instead.
func (*NetworkActivityLog_HttpRequest) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_provenance_api_speepb_v1_network_proxy_proto_rawDescGZIP(), []int{0, 0}
}

func (x *NetworkActivityLog_HttpRequest) GetMethod() string {
	if x != nil && x.Method != nil {
		return *x.Method
	}
	return ""
}

func (x *NetworkActivityLog_HttpRequest) GetScheme() string {
	if x != nil && x.Scheme != nil {
		return *x.Scheme
	}
	return ""
}

func (x *NetworkActivityLog_HttpRequest) GetHost() string {
	if x != nil && x.Host != nil {
		return *x.Host
	}
	return ""
}

func (x *NetworkActivityLog_HttpRequest) GetPath() string {
	if x != nil && x.Path != nil {
		return *x.Path
	}
	return ""
}

type NetworkActivityLog_RawRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Protocol *string `protobuf:"bytes,1,opt,name=protocol,proto3,oneof" json:"protocol,omitempty"`
	Address  *string `protobuf:"bytes,2,opt,name=address,proto3,oneof" json:"address,omitempty"`
}

func (x *NetworkActivityLog_RawRequest) Reset() {
	*x = NetworkActivityLog_RawRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_provenance_api_speepb_v1_network_proxy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkActivityLog_RawRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkActivityLog_RawRequest) ProtoMessage() {}

func (x *NetworkActivityLog_RawRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_provenance_api_speepb_v1_network_proxy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkActivityLog_RawRequest.ProtoReflect.Descriptor instead.
func (*NetworkActivityLog_RawRequest) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_provenance_api_speepb_v1_network_proxy_proto_rawDescGZIP(), []int{0, 1}
}

func (x *NetworkActivityLog_RawRequest) GetProtocol() string {
	if x != nil && x.Protocol != nil {
		return *x.Protocol
	}
	return ""
}

func (x *NetworkActivityLog_RawRequest) GetAddress() string {
	if x != nil && x.Address != nil {
		return *x.Address
	}
	return ""
}

type NetworkActivityLog_PackageInstall struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ecosystem   *string `protobuf:"bytes,1,opt,name=ecosystem,proto3,oneof" json:"ecosystem,omitempty"`
	Name        *string `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Version     *string `protobuf:"bytes,3,opt,name=version,proto3,oneof" json:"version,omitempty"`
	DownloadUri *string `protobuf:"bytes,4,opt,name=download_uri,json=downloadUri,proto3,oneof" json:"download_uri,omitempty"`
}

func (x *NetworkActivityLog_PackageInstall) Reset() {
	*x = NetworkActivityLog_PackageInstall{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_provenance_api_speepb_v1_network_proxy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkActivityLog_PackageInstall) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkActivityLog_PackageInstall) ProtoMessage() {}

func (x *NetworkActivityLog_PackageInstall) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_provenance_api_speepb_v1_network_proxy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkActivityLog_PackageInstall.ProtoReflect.Descriptor instead.
func (*NetworkActivityLog_PackageInstall) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_provenance_api_speepb_v1_network_proxy_proto_rawDescGZIP(), []int{0, 2}
}

func (x *NetworkActivityLog_PackageInstall) GetEcosystem() string {
	if x != nil && x.Ecosystem != nil {
		return *x.Ecosystem
	}
	return ""
}

func (x *NetworkActivityLog_PackageInstall) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *NetworkActivityLog_PackageInstall) GetVersion() string {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return ""
}

func (x *NetworkActivityLog_PackageInstall) GetDownloadUri() string {
	if x != nil && x.DownloadUri != nil {
		return *x.DownloadUri
	}
	return ""
}

var File_go_chromium_org_luci_provenance_api_speepb_v1_network_proxy_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_provenance_api_speepb_v1_network_proxy_proto_rawDesc = []byte{
	0x0a, 0x41, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x6e, 0x61, 0x6e, 0x63,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x70, 0x65, 0x65, 0x70, 0x62, 0x2f, 0x76, 0x31, 0x2f,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x70, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x22, 0xd9, 0x05, 0x0a,
	0x12, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x4c, 0x6f, 0x67, 0x12, 0x4c, 0x0a, 0x0d, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x73, 0x70, 0x65,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x4c, 0x6f, 0x67, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x0c, 0x68, 0x74, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x73, 0x12, 0x49, 0x0a, 0x0c, 0x72, 0x61, 0x77, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73, 0x70, 0x65, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x4c, 0x6f, 0x67, 0x2e, 0x52, 0x61, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x0b, 0x72, 0x61, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x55, 0x0a, 0x10,
	0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x73, 0x70, 0x65, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x4c, 0x6f, 0x67, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6c, 0x6c, 0x52, 0x0f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6c, 0x6c, 0x73, 0x1a, 0xa1, 0x01, 0x0a, 0x0b, 0x48, 0x74, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x88, 0x01, 0x01,
	0x12, 0x1b, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x01, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a,
	0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x04, 0x68,
	0x6f, 0x73, 0x74, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x88, 0x01, 0x01, 0x42,
	0x09, 0x0a, 0x07, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x42, 0x07,
	0x0a, 0x05, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x1a, 0x65, 0x0a, 0x0a, 0x52, 0x61, 0x77, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x1a, 0xc7,
	0x01, 0x0a, 0x0e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c,
	0x6c, 0x12, 0x21, 0x0a, 0x09, 0x65, 0x63, 0x6f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x65, 0x63, 0x6f, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x26, 0x0a, 0x0c,
	0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x03, 0x52, 0x0b, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x72,
	0x69, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x65, 0x63, 0x6f, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x64, 0x6f, 0x77, 0x6e,
	0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x75, 0x72, 0x69, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x6f, 0x2e, 0x63,
	0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69,
	0x2f, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x73, 0x70, 0x65, 0x65, 0x70, 0x62, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x70, 0x65, 0x65, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_provenance_api_speepb_v1_network_proxy_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_provenance_api_speepb_v1_network_proxy_proto_rawDescData = file_go_chromium_org_luci_provenance_api_speepb_v1_network_proxy_proto_rawDesc
)

func file_go_chromium_org_luci_provenance_api_speepb_v1_network_proxy_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_provenance_api_speepb_v1_network_proxy_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_provenance_api_speepb_v1_network_proxy_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_provenance_api_speepb_v1_network_proxy_proto_rawDescData)
	})
	return file_go_chromium_org_luci_provenance_api_speepb_v1_network_proxy_proto_rawDescData
}

var file_go_chromium_org_luci_provenance_api_speepb_v1_network_proxy_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_go_chromium_org_luci_provenance_api_speepb_v1_network_proxy_proto_goTypes = []interface{}{
	(*NetworkActivityLog)(nil),                // 0: spee.v1.NetworkActivityLog
	(*NetworkActivityLog_HttpRequest)(nil),    // 1: spee.v1.NetworkActivityLog.HttpRequest
	(*NetworkActivityLog_RawRequest)(nil),     // 2: spee.v1.NetworkActivityLog.RawRequest
	(*NetworkActivityLog_PackageInstall)(nil), // 3: spee.v1.NetworkActivityLog.PackageInstall
}
var file_go_chromium_org_luci_provenance_api_speepb_v1_network_proxy_proto_depIdxs = []int32{
	1, // 0: spee.v1.NetworkActivityLog.http_requests:type_name -> spee.v1.NetworkActivityLog.HttpRequest
	2, // 1: spee.v1.NetworkActivityLog.raw_requests:type_name -> spee.v1.NetworkActivityLog.RawRequest
	3, // 2: spee.v1.NetworkActivityLog.package_installs:type_name -> spee.v1.NetworkActivityLog.PackageInstall
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_provenance_api_speepb_v1_network_proxy_proto_init() }
func file_go_chromium_org_luci_provenance_api_speepb_v1_network_proxy_proto_init() {
	if File_go_chromium_org_luci_provenance_api_speepb_v1_network_proxy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_provenance_api_speepb_v1_network_proxy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkActivityLog); i {
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
		file_go_chromium_org_luci_provenance_api_speepb_v1_network_proxy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkActivityLog_HttpRequest); i {
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
		file_go_chromium_org_luci_provenance_api_speepb_v1_network_proxy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkActivityLog_RawRequest); i {
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
		file_go_chromium_org_luci_provenance_api_speepb_v1_network_proxy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkActivityLog_PackageInstall); i {
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
	file_go_chromium_org_luci_provenance_api_speepb_v1_network_proxy_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_go_chromium_org_luci_provenance_api_speepb_v1_network_proxy_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_go_chromium_org_luci_provenance_api_speepb_v1_network_proxy_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_go_chromium_org_luci_provenance_api_speepb_v1_network_proxy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_provenance_api_speepb_v1_network_proxy_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_provenance_api_speepb_v1_network_proxy_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_provenance_api_speepb_v1_network_proxy_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_provenance_api_speepb_v1_network_proxy_proto = out.File
	file_go_chromium_org_luci_provenance_api_speepb_v1_network_proxy_proto_rawDesc = nil
	file_go_chromium_org_luci_provenance_api_speepb_v1_network_proxy_proto_goTypes = nil
	file_go_chromium_org_luci_provenance_api_speepb_v1_network_proxy_proto_depIdxs = nil
}
