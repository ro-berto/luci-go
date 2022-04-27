// Copyright 2018 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: go.chromium.org/luci/machine-db/api/crimson/v1/nics.proto

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

// A network interface in the database.
type NIC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of this NIC. With machine, uniquely identifies this NIC.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The machine this NIC belongs to. With name, uniquely identifies this NIC.
	Machine string `protobuf:"bytes,2,opt,name=machine,proto3" json:"machine,omitempty"`
	// The MAC address associated with this NIC.
	MacAddress string `protobuf:"bytes,3,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
	// The switch this NIC is connected to.
	Switch string `protobuf:"bytes,4,opt,name=switch,proto3" json:"switch,omitempty"`
	// The switchport this NIC is connected to.
	Switchport int32 `protobuf:"varint,5,opt,name=switchport,proto3" json:"switchport,omitempty"`
	// The name of this NIC on the network.
	Hostname string `protobuf:"bytes,6,opt,name=hostname,proto3" json:"hostname,omitempty"`
	// The IPv4 address associated with this NIC.
	Ipv4 string `protobuf:"bytes,7,opt,name=ipv4,proto3" json:"ipv4,omitempty"`
}

func (x *NIC) Reset() {
	*x = NIC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_machine_db_api_crimson_v1_nics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NIC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NIC) ProtoMessage() {}

func (x *NIC) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_machine_db_api_crimson_v1_nics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NIC.ProtoReflect.Descriptor instead.
func (*NIC) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_machine_db_api_crimson_v1_nics_proto_rawDescGZIP(), []int{0}
}

func (x *NIC) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NIC) GetMachine() string {
	if x != nil {
		return x.Machine
	}
	return ""
}

func (x *NIC) GetMacAddress() string {
	if x != nil {
		return x.MacAddress
	}
	return ""
}

func (x *NIC) GetSwitch() string {
	if x != nil {
		return x.Switch
	}
	return ""
}

func (x *NIC) GetSwitchport() int32 {
	if x != nil {
		return x.Switchport
	}
	return 0
}

func (x *NIC) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *NIC) GetIpv4() string {
	if x != nil {
		return x.Ipv4
	}
	return ""
}

// A request to create a new NIC in the database.
type CreateNICRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The NIC to create in the database.
	Nic *NIC `protobuf:"bytes,1,opt,name=nic,proto3" json:"nic,omitempty"`
}

func (x *CreateNICRequest) Reset() {
	*x = CreateNICRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_machine_db_api_crimson_v1_nics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNICRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNICRequest) ProtoMessage() {}

func (x *CreateNICRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_machine_db_api_crimson_v1_nics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNICRequest.ProtoReflect.Descriptor instead.
func (*CreateNICRequest) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_machine_db_api_crimson_v1_nics_proto_rawDescGZIP(), []int{1}
}

func (x *CreateNICRequest) GetNic() *NIC {
	if x != nil {
		return x.Nic
	}
	return nil
}

// A request to delete a NIC from the database.
type DeleteNICRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the NIC to delete.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The machine the NIC belongs to.
	Machine string `protobuf:"bytes,2,opt,name=machine,proto3" json:"machine,omitempty"`
}

func (x *DeleteNICRequest) Reset() {
	*x = DeleteNICRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_machine_db_api_crimson_v1_nics_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNICRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNICRequest) ProtoMessage() {}

func (x *DeleteNICRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_machine_db_api_crimson_v1_nics_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNICRequest.ProtoReflect.Descriptor instead.
func (*DeleteNICRequest) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_machine_db_api_crimson_v1_nics_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteNICRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeleteNICRequest) GetMachine() string {
	if x != nil {
		return x.Machine
	}
	return ""
}

// A request to list NICs in the database.
type ListNICsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The names of NICs to filter retrieved NICs on.
	Names []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
	// The machines to filter retrieved NICs on.
	Machines []string `protobuf:"bytes,2,rep,name=machines,proto3" json:"machines,omitempty"`
	// The MAC addresses to filter retrieved NICs on.
	MacAddresses []string `protobuf:"bytes,3,rep,name=mac_addresses,json=macAddresses,proto3" json:"mac_addresses,omitempty"`
	// The switches to filter retrieved NICs on.
	Switches []string `protobuf:"bytes,4,rep,name=switches,proto3" json:"switches,omitempty"`
	// The hostnames of NICs to filter retrieved NICs on.
	Hostnames []string `protobuf:"bytes,5,rep,name=hostnames,proto3" json:"hostnames,omitempty"`
	// The IPv4 addresses of NICs to filter retrieved NICs on.
	Ipv4S []string `protobuf:"bytes,6,rep,name=ipv4s,proto3" json:"ipv4s,omitempty"`
}

func (x *ListNICsRequest) Reset() {
	*x = ListNICsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_machine_db_api_crimson_v1_nics_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNICsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNICsRequest) ProtoMessage() {}

func (x *ListNICsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_machine_db_api_crimson_v1_nics_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNICsRequest.ProtoReflect.Descriptor instead.
func (*ListNICsRequest) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_machine_db_api_crimson_v1_nics_proto_rawDescGZIP(), []int{3}
}

func (x *ListNICsRequest) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

func (x *ListNICsRequest) GetMachines() []string {
	if x != nil {
		return x.Machines
	}
	return nil
}

func (x *ListNICsRequest) GetMacAddresses() []string {
	if x != nil {
		return x.MacAddresses
	}
	return nil
}

func (x *ListNICsRequest) GetSwitches() []string {
	if x != nil {
		return x.Switches
	}
	return nil
}

func (x *ListNICsRequest) GetHostnames() []string {
	if x != nil {
		return x.Hostnames
	}
	return nil
}

func (x *ListNICsRequest) GetIpv4S() []string {
	if x != nil {
		return x.Ipv4S
	}
	return nil
}

// A response containing a list of NICs in the database.
type ListNICsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The NICs matching this request.
	Nics []*NIC `protobuf:"bytes,1,rep,name=nics,proto3" json:"nics,omitempty"` // TODO(smut): Support page tokens.
}

func (x *ListNICsResponse) Reset() {
	*x = ListNICsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_machine_db_api_crimson_v1_nics_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNICsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNICsResponse) ProtoMessage() {}

func (x *ListNICsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_machine_db_api_crimson_v1_nics_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNICsResponse.ProtoReflect.Descriptor instead.
func (*ListNICsResponse) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_machine_db_api_crimson_v1_nics_proto_rawDescGZIP(), []int{4}
}

func (x *ListNICsResponse) GetNics() []*NIC {
	if x != nil {
		return x.Nics
	}
	return nil
}

// A request to update a NIC in the database.
type UpdateNICRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The NIC to update in the database.
	Nic *NIC `protobuf:"bytes,1,opt,name=nic,proto3" json:"nic,omitempty"`
	// The fields to update in the NIC.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateNICRequest) Reset() {
	*x = UpdateNICRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_machine_db_api_crimson_v1_nics_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNICRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNICRequest) ProtoMessage() {}

func (x *UpdateNICRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_machine_db_api_crimson_v1_nics_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNICRequest.ProtoReflect.Descriptor instead.
func (*UpdateNICRequest) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_machine_db_api_crimson_v1_nics_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateNICRequest) GetNic() *NIC {
	if x != nil {
		return x.Nic
	}
	return nil
}

func (x *UpdateNICRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

var File_go_chromium_org_luci_machine_db_api_crimson_v1_nics_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_machine_db_api_crimson_v1_nics_proto_rawDesc = []byte{
	0x0a, 0x39, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2d, 0x64,
	0x62, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x72, 0x69, 0x6d, 0x73, 0x6f, 0x6e, 0x2f, 0x76, 0x31,
	0x2f, 0x6e, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x72, 0x69,
	0x6d, 0x73, 0x6f, 0x6e, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbc, 0x01, 0x0a, 0x03, 0x4e, 0x49, 0x43, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x6d, 0x61, 0x63, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x6d, 0x61, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x77, 0x69, 0x74, 0x63, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x70,
	0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x77, 0x69, 0x74, 0x63,
	0x68, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x70, 0x76, 0x34, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x69, 0x70, 0x76, 0x34, 0x22, 0x32, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e,
	0x49, 0x43, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x03, 0x6e, 0x69, 0x63,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x72, 0x69, 0x6d, 0x73, 0x6f, 0x6e,
	0x2e, 0x4e, 0x49, 0x43, 0x52, 0x03, 0x6e, 0x69, 0x63, 0x22, 0x40, 0x0a, 0x10, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4e, 0x49, 0x43, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x22, 0xb8, 0x01, 0x0a, 0x0f,
	0x4c, 0x69, 0x73, 0x74, 0x4e, 0x49, 0x43, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x73, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x61, 0x63, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x61, 0x63, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68,
	0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68,
	0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x69, 0x70, 0x76, 0x34, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x05, 0x69, 0x70, 0x76, 0x34, 0x73, 0x22, 0x34, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x49,
	0x43, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x04, 0x6e, 0x69,
	0x63, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x72, 0x69, 0x6d, 0x73,
	0x6f, 0x6e, 0x2e, 0x4e, 0x49, 0x43, 0x52, 0x04, 0x6e, 0x69, 0x63, 0x73, 0x22, 0x6f, 0x0a, 0x10,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x49, 0x43, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1e, 0x0a, 0x03, 0x6e, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x63, 0x72, 0x69, 0x6d, 0x73, 0x6f, 0x6e, 0x2e, 0x4e, 0x49, 0x43, 0x52, 0x03, 0x6e, 0x69, 0x63,
	0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73,
	0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x42, 0x38, 0x5a,
	0x36, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2d, 0x64, 0x62,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x72, 0x69, 0x6d, 0x73, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b,
	0x63, 0x72, 0x69, 0x6d, 0x73, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_machine_db_api_crimson_v1_nics_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_machine_db_api_crimson_v1_nics_proto_rawDescData = file_go_chromium_org_luci_machine_db_api_crimson_v1_nics_proto_rawDesc
)

func file_go_chromium_org_luci_machine_db_api_crimson_v1_nics_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_machine_db_api_crimson_v1_nics_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_machine_db_api_crimson_v1_nics_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_machine_db_api_crimson_v1_nics_proto_rawDescData)
	})
	return file_go_chromium_org_luci_machine_db_api_crimson_v1_nics_proto_rawDescData
}

var file_go_chromium_org_luci_machine_db_api_crimson_v1_nics_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_go_chromium_org_luci_machine_db_api_crimson_v1_nics_proto_goTypes = []interface{}{
	(*NIC)(nil),                   // 0: crimson.NIC
	(*CreateNICRequest)(nil),      // 1: crimson.CreateNICRequest
	(*DeleteNICRequest)(nil),      // 2: crimson.DeleteNICRequest
	(*ListNICsRequest)(nil),       // 3: crimson.ListNICsRequest
	(*ListNICsResponse)(nil),      // 4: crimson.ListNICsResponse
	(*UpdateNICRequest)(nil),      // 5: crimson.UpdateNICRequest
	(*fieldmaskpb.FieldMask)(nil), // 6: google.protobuf.FieldMask
}
var file_go_chromium_org_luci_machine_db_api_crimson_v1_nics_proto_depIdxs = []int32{
	0, // 0: crimson.CreateNICRequest.nic:type_name -> crimson.NIC
	0, // 1: crimson.ListNICsResponse.nics:type_name -> crimson.NIC
	0, // 2: crimson.UpdateNICRequest.nic:type_name -> crimson.NIC
	6, // 3: crimson.UpdateNICRequest.update_mask:type_name -> google.protobuf.FieldMask
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_machine_db_api_crimson_v1_nics_proto_init() }
func file_go_chromium_org_luci_machine_db_api_crimson_v1_nics_proto_init() {
	if File_go_chromium_org_luci_machine_db_api_crimson_v1_nics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_machine_db_api_crimson_v1_nics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NIC); i {
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
		file_go_chromium_org_luci_machine_db_api_crimson_v1_nics_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNICRequest); i {
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
		file_go_chromium_org_luci_machine_db_api_crimson_v1_nics_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNICRequest); i {
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
		file_go_chromium_org_luci_machine_db_api_crimson_v1_nics_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNICsRequest); i {
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
		file_go_chromium_org_luci_machine_db_api_crimson_v1_nics_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNICsResponse); i {
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
		file_go_chromium_org_luci_machine_db_api_crimson_v1_nics_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNICRequest); i {
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
			RawDescriptor: file_go_chromium_org_luci_machine_db_api_crimson_v1_nics_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_machine_db_api_crimson_v1_nics_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_machine_db_api_crimson_v1_nics_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_machine_db_api_crimson_v1_nics_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_machine_db_api_crimson_v1_nics_proto = out.File
	file_go_chromium_org_luci_machine_db_api_crimson_v1_nics_proto_rawDesc = nil
	file_go_chromium_org_luci_machine_db_api_crimson_v1_nics_proto_goTypes = nil
	file_go_chromium_org_luci_machine_db_api_crimson_v1_nics_proto_depIdxs = nil
}
