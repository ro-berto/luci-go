// Copyright 2021 The LUCI Authors.
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
// source: go.chromium.org/luci/cipd/api/cipd/v1/access_log.proto

package api

import (
	_ "go.chromium.org/luci/common/bq/pb"
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

// AccessLogEntry defines a schema for `access` BigQuery table with access logs.
//
// This is a best effort log populated using in-memory buffers. Some entries may
// be dropped if a process crashes before it flushes the buffer.
//
// Field types must be compatible with BigQuery Storage Write API, see
// https://cloud.google.com/bigquery/docs/write-api#data_type_conversions
type AccessLogEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method           string   `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`                                                 // the RPC method, e.g. "/cipd.Repository/ListPrefix"
	Timestamp        int64    `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`                                          // microseconds since epoch
	Package          string   `protobuf:"bytes,3,opt,name=package,proto3" json:"package,omitempty"`                                               // CIPD package or package prefix (if known)
	Instance         string   `protobuf:"bytes,4,opt,name=instance,proto3" json:"instance,omitempty"`                                             // CIPD instance ID (if known)
	Version          string   `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`                                               // CIPD package version (if known)
	Tags             []string `protobuf:"bytes,6,rep,name=tags,proto3" json:"tags,omitempty"`                                                     // the requested tags (if known)
	Metadata         []string `protobuf:"bytes,7,rep,name=metadata,proto3" json:"metadata,omitempty"`                                             // the requested metadata keys (if known)
	Flags            []string `protobuf:"bytes,8,rep,name=flags,proto3" json:"flags,omitempty"`                                                   // encoding of boolean flags in the request
	CallIdentity     string   `protobuf:"bytes,9,opt,name=call_identity,json=callIdentity,proto3" json:"call_identity,omitempty"`                 // identity used to authorize the call
	PeerIdentity     string   `protobuf:"bytes,10,opt,name=peer_identity,json=peerIdentity,proto3" json:"peer_identity,omitempty"`                // identity of a service that made the RPC
	PeerIp           string   `protobuf:"bytes,11,opt,name=peer_ip,json=peerIp,proto3" json:"peer_ip,omitempty"`                                  // IP address of the caller
	UserAgent        string   `protobuf:"bytes,12,opt,name=user_agent,json=userAgent,proto3" json:"user_agent,omitempty"`                         // value of "User-Agent" header
	ServiceVersion   string   `protobuf:"bytes,13,opt,name=service_version,json=serviceVersion,proto3" json:"service_version,omitempty"`          // GAE app and version that handled the request
	ProcessId        string   `protobuf:"bytes,14,opt,name=process_id,json=processId,proto3" json:"process_id,omitempty"`                         // identifier of the concrete backend process
	RequestId        string   `protobuf:"bytes,15,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`                         // Cloud Trace ID of the request
	AuthDbRev        int64    `protobuf:"varint,16,opt,name=auth_db_rev,json=authDbRev,proto3" json:"auth_db_rev,omitempty"`                      // AuthDB revision used to authorize the call
	ResponseCode     string   `protobuf:"bytes,17,opt,name=response_code,json=responseCode,proto3" json:"response_code,omitempty"`                // canonical gRPC response code (e.g. "OK")
	ResponseErr      string   `protobuf:"bytes,18,opt,name=response_err,json=responseErr,proto3" json:"response_err,omitempty"`                   // the response error message if any
	ResponseTimeUsec int64    `protobuf:"varint,19,opt,name=response_time_usec,json=responseTimeUsec,proto3" json:"response_time_usec,omitempty"` // time spent handling the request in microseconds
}

func (x *AccessLogEntry) Reset() {
	*x = AccessLogEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cipd_api_cipd_v1_access_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessLogEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessLogEntry) ProtoMessage() {}

func (x *AccessLogEntry) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cipd_api_cipd_v1_access_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessLogEntry.ProtoReflect.Descriptor instead.
func (*AccessLogEntry) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cipd_api_cipd_v1_access_log_proto_rawDescGZIP(), []int{0}
}

func (x *AccessLogEntry) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *AccessLogEntry) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *AccessLogEntry) GetPackage() string {
	if x != nil {
		return x.Package
	}
	return ""
}

func (x *AccessLogEntry) GetInstance() string {
	if x != nil {
		return x.Instance
	}
	return ""
}

func (x *AccessLogEntry) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *AccessLogEntry) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *AccessLogEntry) GetMetadata() []string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *AccessLogEntry) GetFlags() []string {
	if x != nil {
		return x.Flags
	}
	return nil
}

func (x *AccessLogEntry) GetCallIdentity() string {
	if x != nil {
		return x.CallIdentity
	}
	return ""
}

func (x *AccessLogEntry) GetPeerIdentity() string {
	if x != nil {
		return x.PeerIdentity
	}
	return ""
}

func (x *AccessLogEntry) GetPeerIp() string {
	if x != nil {
		return x.PeerIp
	}
	return ""
}

func (x *AccessLogEntry) GetUserAgent() string {
	if x != nil {
		return x.UserAgent
	}
	return ""
}

func (x *AccessLogEntry) GetServiceVersion() string {
	if x != nil {
		return x.ServiceVersion
	}
	return ""
}

func (x *AccessLogEntry) GetProcessId() string {
	if x != nil {
		return x.ProcessId
	}
	return ""
}

func (x *AccessLogEntry) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *AccessLogEntry) GetAuthDbRev() int64 {
	if x != nil {
		return x.AuthDbRev
	}
	return 0
}

func (x *AccessLogEntry) GetResponseCode() string {
	if x != nil {
		return x.ResponseCode
	}
	return ""
}

func (x *AccessLogEntry) GetResponseErr() string {
	if x != nil {
		return x.ResponseErr
	}
	return ""
}

func (x *AccessLogEntry) GetResponseTimeUsec() int64 {
	if x != nil {
		return x.ResponseTimeUsec
	}
	return 0
}

var File_go_chromium_org_luci_cipd_api_cipd_v1_access_log_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_cipd_api_cipd_v1_access_log_proto_rawDesc = []byte{
	0x0a, 0x36, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x69, 0x70, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x69, 0x70, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6c,
	0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x69, 0x70, 0x64, 0x1a, 0x2f,
	0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x71, 0x2f, 0x70,
	0x62, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xec, 0x04, 0x0a, 0x0e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x2d, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0f, 0xe2,
	0xbc, 0x24, 0x0b, 0x0a, 0x09, 0x54, 0x49, 0x4d, 0x45, 0x53, 0x54, 0x41, 0x4d, 0x50, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67,
	0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6c, 0x61,
	0x67, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x12,
	0x23, 0x0a, 0x0d, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x61, 0x6c, 0x6c, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x65, 0x65,
	0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x65, 0x65,
	0x72, 0x5f, 0x69, 0x70, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x65, 0x65, 0x72,
	0x49, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x68,
	0x5f, 0x64, 0x62, 0x5f, 0x72, 0x65, 0x76, 0x18, 0x10, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61,
	0x75, 0x74, 0x68, 0x44, 0x62, 0x52, 0x65, 0x76, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x65, 0x72, 0x72, 0x18, 0x12, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x45, 0x72, 0x72,
	0x12, 0x2c, 0x0a, 0x12, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x5f, 0x75, 0x73, 0x65, 0x63, 0x18, 0x13, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x73, 0x65, 0x63, 0x42, 0x2b,
	0x5a, 0x29, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x69, 0x70, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x69, 0x70, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_cipd_api_cipd_v1_access_log_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_cipd_api_cipd_v1_access_log_proto_rawDescData = file_go_chromium_org_luci_cipd_api_cipd_v1_access_log_proto_rawDesc
)

func file_go_chromium_org_luci_cipd_api_cipd_v1_access_log_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_cipd_api_cipd_v1_access_log_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_cipd_api_cipd_v1_access_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_cipd_api_cipd_v1_access_log_proto_rawDescData)
	})
	return file_go_chromium_org_luci_cipd_api_cipd_v1_access_log_proto_rawDescData
}

var file_go_chromium_org_luci_cipd_api_cipd_v1_access_log_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_go_chromium_org_luci_cipd_api_cipd_v1_access_log_proto_goTypes = []interface{}{
	(*AccessLogEntry)(nil), // 0: cipd.AccessLogEntry
}
var file_go_chromium_org_luci_cipd_api_cipd_v1_access_log_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_cipd_api_cipd_v1_access_log_proto_init() }
func file_go_chromium_org_luci_cipd_api_cipd_v1_access_log_proto_init() {
	if File_go_chromium_org_luci_cipd_api_cipd_v1_access_log_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_cipd_api_cipd_v1_access_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessLogEntry); i {
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
			RawDescriptor: file_go_chromium_org_luci_cipd_api_cipd_v1_access_log_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_cipd_api_cipd_v1_access_log_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_cipd_api_cipd_v1_access_log_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_cipd_api_cipd_v1_access_log_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_cipd_api_cipd_v1_access_log_proto = out.File
	file_go_chromium_org_luci_cipd_api_cipd_v1_access_log_proto_rawDesc = nil
	file_go_chromium_org_luci_cipd_api_cipd_v1_access_log_proto_goTypes = nil
	file_go_chromium_org_luci_cipd_api_cipd_v1_access_log_proto_depIdxs = nil
}
