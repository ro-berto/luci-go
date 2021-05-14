// Copyright 2018 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.0
// source: go.chromium.org/luci/machine-db/api/crimson/v1/dracs.proto

package crimson

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A DRAC in the database.
type DRAC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of this DRAC on the network. Uniquely identifies this DRAC.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The machine this DRAC belongs to. Uniquely identifies this DRAC.
	Machine string `protobuf:"bytes,2,opt,name=machine,proto3" json:"machine,omitempty"`
	// The IPv4 address associated with this DRAC.
	Ipv4 string `protobuf:"bytes,3,opt,name=ipv4,proto3" json:"ipv4,omitempty"`
	// The VLAN this DRAC belongs to.
	// When creating a DRAC, omit this field. It will be inferred from the IPv4 address.
	Vlan int64 `protobuf:"varint,4,opt,name=vlan,proto3" json:"vlan,omitempty"`
	// The MAC address associated with this DRAC.
	MacAddress string `protobuf:"bytes,5,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
	// The switch this DRAC is connected to.
	Switch string `protobuf:"bytes,6,opt,name=switch,proto3" json:"switch,omitempty"`
	// The switchport this DRAC is connected to.
	Switchport int32 `protobuf:"varint,7,opt,name=switchport,proto3" json:"switchport,omitempty"`
}

func (x *DRAC) Reset() {
	*x = DRAC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_machine_db_api_crimson_v1_dracs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DRAC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DRAC) ProtoMessage() {}

func (x *DRAC) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_machine_db_api_crimson_v1_dracs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DRAC.ProtoReflect.Descriptor instead.
func (*DRAC) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_machine_db_api_crimson_v1_dracs_proto_rawDescGZIP(), []int{0}
}

func (x *DRAC) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DRAC) GetMachine() string {
	if x != nil {
		return x.Machine
	}
	return ""
}

func (x *DRAC) GetIpv4() string {
	if x != nil {
		return x.Ipv4
	}
	return ""
}

func (x *DRAC) GetVlan() int64 {
	if x != nil {
		return x.Vlan
	}
	return 0
}

func (x *DRAC) GetMacAddress() string {
	if x != nil {
		return x.MacAddress
	}
	return ""
}

func (x *DRAC) GetSwitch() string {
	if x != nil {
		return x.Switch
	}
	return ""
}

func (x *DRAC) GetSwitchport() int32 {
	if x != nil {
		return x.Switchport
	}
	return 0
}

// A request to create a new DRAC in the database.
type CreateDRACRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The DRAC to create in the database.
	Drac *DRAC `protobuf:"bytes,1,opt,name=drac,proto3" json:"drac,omitempty"`
}

func (x *CreateDRACRequest) Reset() {
	*x = CreateDRACRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_machine_db_api_crimson_v1_dracs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDRACRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDRACRequest) ProtoMessage() {}

func (x *CreateDRACRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_machine_db_api_crimson_v1_dracs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDRACRequest.ProtoReflect.Descriptor instead.
func (*CreateDRACRequest) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_machine_db_api_crimson_v1_dracs_proto_rawDescGZIP(), []int{1}
}

func (x *CreateDRACRequest) GetDrac() *DRAC {
	if x != nil {
		return x.Drac
	}
	return nil
}

// A request to list DRACs in the database.
type ListDRACsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The names of DRACs to get.
	Names []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
	// The machines to filter retrieved DRACs on.
	Machines []string `protobuf:"bytes,2,rep,name=machines,proto3" json:"machines,omitempty"`
	// The IPv4 addresses to filter retrieved DRACs on.
	Ipv4S []string `protobuf:"bytes,3,rep,name=ipv4s,proto3" json:"ipv4s,omitempty"`
	// The VLANs to filter retrieved DRACs on.
	Vlans []int64 `protobuf:"varint,4,rep,packed,name=vlans,proto3" json:"vlans,omitempty"`
	// The MAC addresses to filter retrieved DRACs on.
	MacAddresses []string `protobuf:"bytes,5,rep,name=mac_addresses,json=macAddresses,proto3" json:"mac_addresses,omitempty"`
	// The switches to filter retrieved DRACs on.
	Switches []string `protobuf:"bytes,6,rep,name=switches,proto3" json:"switches,omitempty"`
}

func (x *ListDRACsRequest) Reset() {
	*x = ListDRACsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_machine_db_api_crimson_v1_dracs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDRACsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDRACsRequest) ProtoMessage() {}

func (x *ListDRACsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_machine_db_api_crimson_v1_dracs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDRACsRequest.ProtoReflect.Descriptor instead.
func (*ListDRACsRequest) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_machine_db_api_crimson_v1_dracs_proto_rawDescGZIP(), []int{2}
}

func (x *ListDRACsRequest) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

func (x *ListDRACsRequest) GetMachines() []string {
	if x != nil {
		return x.Machines
	}
	return nil
}

func (x *ListDRACsRequest) GetIpv4S() []string {
	if x != nil {
		return x.Ipv4S
	}
	return nil
}

func (x *ListDRACsRequest) GetVlans() []int64 {
	if x != nil {
		return x.Vlans
	}
	return nil
}

func (x *ListDRACsRequest) GetMacAddresses() []string {
	if x != nil {
		return x.MacAddresses
	}
	return nil
}

func (x *ListDRACsRequest) GetSwitches() []string {
	if x != nil {
		return x.Switches
	}
	return nil
}

// A response containing a list of DRACs in the database.
type ListDRACsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The DRACs matching this request.
	Dracs []*DRAC `protobuf:"bytes,1,rep,name=dracs,proto3" json:"dracs,omitempty"` // TODO(smut): Support page tokens.
}

func (x *ListDRACsResponse) Reset() {
	*x = ListDRACsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_machine_db_api_crimson_v1_dracs_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDRACsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDRACsResponse) ProtoMessage() {}

func (x *ListDRACsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_machine_db_api_crimson_v1_dracs_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDRACsResponse.ProtoReflect.Descriptor instead.
func (*ListDRACsResponse) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_machine_db_api_crimson_v1_dracs_proto_rawDescGZIP(), []int{3}
}

func (x *ListDRACsResponse) GetDracs() []*DRAC {
	if x != nil {
		return x.Dracs
	}
	return nil
}

// A request to update a DRAC in the database.
type UpdateDRACRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The DRAC to update in the database.
	Drac *DRAC `protobuf:"bytes,1,opt,name=drac,proto3" json:"drac,omitempty"`
	// The fields to update in the DRAC.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateDRACRequest) Reset() {
	*x = UpdateDRACRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_machine_db_api_crimson_v1_dracs_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDRACRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDRACRequest) ProtoMessage() {}

func (x *UpdateDRACRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_machine_db_api_crimson_v1_dracs_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDRACRequest.ProtoReflect.Descriptor instead.
func (*UpdateDRACRequest) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_machine_db_api_crimson_v1_dracs_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateDRACRequest) GetDrac() *DRAC {
	if x != nil {
		return x.Drac
	}
	return nil
}

func (x *UpdateDRACRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

var File_go_chromium_org_luci_machine_db_api_crimson_v1_dracs_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_machine_db_api_crimson_v1_dracs_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2d, 0x64,
	0x62, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x72, 0x69, 0x6d, 0x73, 0x6f, 0x6e, 0x2f, 0x76, 0x31,
	0x2f, 0x64, 0x72, 0x61, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x72,
	0x69, 0x6d, 0x73, 0x6f, 0x6e, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb5, 0x01, 0x0a, 0x04, 0x44, 0x52, 0x41, 0x43,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x69, 0x70, 0x76, 0x34, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x70,
	0x76, 0x34, 0x12, 0x12, 0x0a, 0x04, 0x76, 0x6c, 0x61, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x76, 0x6c, 0x61, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x63, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x61, 0x63,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x77, 0x69, 0x74, 0x63,
	0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x12,
	0x1e, 0x0a, 0x0a, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x70, 0x6f, 0x72, 0x74, 0x22,
	0x36, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x52, 0x41, 0x43, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x04, 0x64, 0x72, 0x61, 0x63, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x72, 0x69, 0x6d, 0x73, 0x6f, 0x6e, 0x2e, 0x44, 0x52, 0x41,
	0x43, 0x52, 0x04, 0x64, 0x72, 0x61, 0x63, 0x22, 0xb1, 0x01, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74,
	0x44, 0x52, 0x41, 0x43, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x69, 0x70, 0x76, 0x34, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x69,
	0x70, 0x76, 0x34, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x6c, 0x61, 0x6e, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x03, 0x52, 0x05, 0x76, 0x6c, 0x61, 0x6e, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x61,
	0x63, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0c, 0x6d, 0x61, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x08, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x65, 0x73, 0x22, 0x38, 0x0a, 0x11, 0x4c,
	0x69, 0x73, 0x74, 0x44, 0x52, 0x41, 0x43, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x23, 0x0a, 0x05, 0x64, 0x72, 0x61, 0x63, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x63, 0x72, 0x69, 0x6d, 0x73, 0x6f, 0x6e, 0x2e, 0x44, 0x52, 0x41, 0x43, 0x52, 0x05,
	0x64, 0x72, 0x61, 0x63, 0x73, 0x22, 0x73, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44,
	0x52, 0x41, 0x43, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x04, 0x64, 0x72,
	0x61, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x72, 0x69, 0x6d, 0x73,
	0x6f, 0x6e, 0x2e, 0x44, 0x52, 0x41, 0x43, 0x52, 0x04, 0x64, 0x72, 0x61, 0x63, 0x12, 0x3b, 0x0a,
	0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x6f,
	0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75,
	0x63, 0x69, 0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2d, 0x64, 0x62, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x72, 0x69, 0x6d, 0x73, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x72, 0x69,
	0x6d, 0x73, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_machine_db_api_crimson_v1_dracs_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_machine_db_api_crimson_v1_dracs_proto_rawDescData = file_go_chromium_org_luci_machine_db_api_crimson_v1_dracs_proto_rawDesc
)

func file_go_chromium_org_luci_machine_db_api_crimson_v1_dracs_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_machine_db_api_crimson_v1_dracs_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_machine_db_api_crimson_v1_dracs_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_machine_db_api_crimson_v1_dracs_proto_rawDescData)
	})
	return file_go_chromium_org_luci_machine_db_api_crimson_v1_dracs_proto_rawDescData
}

var file_go_chromium_org_luci_machine_db_api_crimson_v1_dracs_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_go_chromium_org_luci_machine_db_api_crimson_v1_dracs_proto_goTypes = []interface{}{
	(*DRAC)(nil),                  // 0: crimson.DRAC
	(*CreateDRACRequest)(nil),     // 1: crimson.CreateDRACRequest
	(*ListDRACsRequest)(nil),      // 2: crimson.ListDRACsRequest
	(*ListDRACsResponse)(nil),     // 3: crimson.ListDRACsResponse
	(*UpdateDRACRequest)(nil),     // 4: crimson.UpdateDRACRequest
	(*fieldmaskpb.FieldMask)(nil), // 5: google.protobuf.FieldMask
}
var file_go_chromium_org_luci_machine_db_api_crimson_v1_dracs_proto_depIdxs = []int32{
	0, // 0: crimson.CreateDRACRequest.drac:type_name -> crimson.DRAC
	0, // 1: crimson.ListDRACsResponse.dracs:type_name -> crimson.DRAC
	0, // 2: crimson.UpdateDRACRequest.drac:type_name -> crimson.DRAC
	5, // 3: crimson.UpdateDRACRequest.update_mask:type_name -> google.protobuf.FieldMask
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_machine_db_api_crimson_v1_dracs_proto_init() }
func file_go_chromium_org_luci_machine_db_api_crimson_v1_dracs_proto_init() {
	if File_go_chromium_org_luci_machine_db_api_crimson_v1_dracs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_machine_db_api_crimson_v1_dracs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DRAC); i {
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
		file_go_chromium_org_luci_machine_db_api_crimson_v1_dracs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDRACRequest); i {
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
		file_go_chromium_org_luci_machine_db_api_crimson_v1_dracs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDRACsRequest); i {
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
		file_go_chromium_org_luci_machine_db_api_crimson_v1_dracs_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDRACsResponse); i {
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
		file_go_chromium_org_luci_machine_db_api_crimson_v1_dracs_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDRACRequest); i {
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
			RawDescriptor: file_go_chromium_org_luci_machine_db_api_crimson_v1_dracs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_machine_db_api_crimson_v1_dracs_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_machine_db_api_crimson_v1_dracs_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_machine_db_api_crimson_v1_dracs_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_machine_db_api_crimson_v1_dracs_proto = out.File
	file_go_chromium_org_luci_machine_db_api_crimson_v1_dracs_proto_rawDesc = nil
	file_go_chromium_org_luci_machine_db_api_crimson_v1_dracs_proto_goTypes = nil
	file_go_chromium_org_luci_machine_db_api_crimson_v1_dracs_proto_depIdxs = nil
}
