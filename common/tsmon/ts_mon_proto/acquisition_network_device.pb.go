// Copyright 2016 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.0
// source: go.chromium.org/luci/common/tsmon/ts_mon_proto/acquisition_network_device.proto

package ts_mon_proto

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

type NetworkDevice_TypeId int32

const (
	NetworkDevice_MESSAGE_TYPE_ID NetworkDevice_TypeId = 34049749
)

// Enum value maps for NetworkDevice_TypeId.
var (
	NetworkDevice_TypeId_name = map[int32]string{
		34049749: "MESSAGE_TYPE_ID",
	}
	NetworkDevice_TypeId_value = map[string]int32{
		"MESSAGE_TYPE_ID": 34049749,
	}
)

func (x NetworkDevice_TypeId) Enum() *NetworkDevice_TypeId {
	p := new(NetworkDevice_TypeId)
	*p = x
	return p
}

func (x NetworkDevice_TypeId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NetworkDevice_TypeId) Descriptor() protoreflect.EnumDescriptor {
	return file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_network_device_proto_enumTypes[0].Descriptor()
}

func (NetworkDevice_TypeId) Type() protoreflect.EnumType {
	return &file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_network_device_proto_enumTypes[0]
}

func (x NetworkDevice_TypeId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *NetworkDevice_TypeId) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = NetworkDevice_TypeId(num)
	return nil
}

// Deprecated: Use NetworkDevice_TypeId.Descriptor instead.
func (NetworkDevice_TypeId) EnumDescriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_network_device_proto_rawDescGZIP(), []int{0, 0}
}

type NetworkDevice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProxyEnvironment *string `protobuf:"bytes,5,opt,name=proxy_environment,json=proxyEnvironment" json:"proxy_environment,omitempty"`
	AcquisitionName  *string `protobuf:"bytes,10,opt,name=acquisition_name,json=acquisitionName" json:"acquisition_name,omitempty"`
	Pop              *string `protobuf:"bytes,30,opt,name=pop" json:"pop,omitempty"`
	Alertable        *bool   `protobuf:"varint,101,opt,name=alertable" json:"alertable,omitempty"`
	Realm            *string `protobuf:"bytes,102,opt,name=realm" json:"realm,omitempty"`
	Asn              *int64  `protobuf:"varint,103,opt,name=asn" json:"asn,omitempty"`
	Metro            *string `protobuf:"bytes,104,opt,name=metro" json:"metro,omitempty"`
	Role             *string `protobuf:"bytes,105,opt,name=role" json:"role,omitempty"`
	Hostname         *string `protobuf:"bytes,106,opt,name=hostname" json:"hostname,omitempty"`
	Vendor           *string `protobuf:"bytes,70,opt,name=vendor" json:"vendor,omitempty"`
	Hostgroup        *string `protobuf:"bytes,108,opt,name=hostgroup" json:"hostgroup,omitempty"`
	ProxyZone        *string `protobuf:"bytes,100,opt,name=proxy_zone,json=proxyZone" json:"proxy_zone,omitempty"`
}

func (x *NetworkDevice) Reset() {
	*x = NetworkDevice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_network_device_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkDevice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkDevice) ProtoMessage() {}

func (x *NetworkDevice) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_network_device_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkDevice.ProtoReflect.Descriptor instead.
func (*NetworkDevice) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_network_device_proto_rawDescGZIP(), []int{0}
}

func (x *NetworkDevice) GetProxyEnvironment() string {
	if x != nil && x.ProxyEnvironment != nil {
		return *x.ProxyEnvironment
	}
	return ""
}

func (x *NetworkDevice) GetAcquisitionName() string {
	if x != nil && x.AcquisitionName != nil {
		return *x.AcquisitionName
	}
	return ""
}

func (x *NetworkDevice) GetPop() string {
	if x != nil && x.Pop != nil {
		return *x.Pop
	}
	return ""
}

func (x *NetworkDevice) GetAlertable() bool {
	if x != nil && x.Alertable != nil {
		return *x.Alertable
	}
	return false
}

func (x *NetworkDevice) GetRealm() string {
	if x != nil && x.Realm != nil {
		return *x.Realm
	}
	return ""
}

func (x *NetworkDevice) GetAsn() int64 {
	if x != nil && x.Asn != nil {
		return *x.Asn
	}
	return 0
}

func (x *NetworkDevice) GetMetro() string {
	if x != nil && x.Metro != nil {
		return *x.Metro
	}
	return ""
}

func (x *NetworkDevice) GetRole() string {
	if x != nil && x.Role != nil {
		return *x.Role
	}
	return ""
}

func (x *NetworkDevice) GetHostname() string {
	if x != nil && x.Hostname != nil {
		return *x.Hostname
	}
	return ""
}

func (x *NetworkDevice) GetVendor() string {
	if x != nil && x.Vendor != nil {
		return *x.Vendor
	}
	return ""
}

func (x *NetworkDevice) GetHostgroup() string {
	if x != nil && x.Hostgroup != nil {
		return *x.Hostgroup
	}
	return ""
}

func (x *NetworkDevice) GetProxyZone() string {
	if x != nil && x.ProxyZone != nil {
		return *x.ProxyZone
	}
	return ""
}

var File_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_network_device_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_network_device_proto_rawDesc = []byte{
	0x0a, 0x4f, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x73,
	0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x73, 0x5f, 0x6d, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x61, 0x63, 0x71, 0x75, 0x69, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x74, 0x73, 0x5f, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xfc, 0x02, 0x0a, 0x0d, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x2b, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x65, 0x6e, 0x76, 0x69, 0x72,
	0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x29,
	0x0a, 0x10, 0x61, 0x63, 0x71, 0x75, 0x69, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x63, 0x71, 0x75, 0x69, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x6f, 0x70,
	0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x6f, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09,
	0x61, 0x6c, 0x65, 0x72, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x61,
	0x6c, 0x6d, 0x18, 0x66, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x12,
	0x10, 0x0a, 0x03, 0x61, 0x73, 0x6e, 0x18, 0x67, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x61, 0x73,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x65, 0x74, 0x72, 0x6f, 0x18, 0x68, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6d, 0x65, 0x74, 0x72, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18,
	0x69, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x68,
	0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x6a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68,
	0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f,
	0x72, 0x18, 0x46, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12,
	0x1c, 0x0a, 0x09, 0x68, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x6c, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x68, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x64, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5a, 0x6f, 0x6e, 0x65, 0x22, 0x20, 0x0a, 0x06,
	0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x0f, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47,
	0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x44, 0x10, 0xd5, 0x9d, 0x9e, 0x10, 0x42, 0x30,
	0x5a, 0x2e, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x73,
	0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x73, 0x5f, 0x6d, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_network_device_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_network_device_proto_rawDescData = file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_network_device_proto_rawDesc
)

func file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_network_device_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_network_device_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_network_device_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_network_device_proto_rawDescData)
	})
	return file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_network_device_proto_rawDescData
}

var file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_network_device_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_network_device_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_network_device_proto_goTypes = []interface{}{
	(NetworkDevice_TypeId)(0), // 0: ts_mon.proto.NetworkDevice.TypeId
	(*NetworkDevice)(nil),     // 1: ts_mon.proto.NetworkDevice
}
var file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_network_device_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_network_device_proto_init()
}
func file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_network_device_proto_init() {
	if File_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_network_device_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_network_device_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkDevice); i {
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
			RawDescriptor: file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_network_device_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_network_device_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_network_device_proto_depIdxs,
		EnumInfos:         file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_network_device_proto_enumTypes,
		MessageInfos:      file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_network_device_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_network_device_proto = out.File
	file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_network_device_proto_rawDesc = nil
	file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_network_device_proto_goTypes = nil
	file_go_chromium_org_luci_common_tsmon_ts_mon_proto_acquisition_network_device_proto_depIdxs = nil
}
