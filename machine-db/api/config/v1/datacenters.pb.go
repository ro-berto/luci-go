// Copyright 2017 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.1
// source: go.chromium.org/luci/machine-db/api/config/v1/datacenters.proto

package config

import (
	v1 "go.chromium.org/luci/machine-db/api/common/v1"
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

// A switch to store in the database.
type Switch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of this switch. Must be globally unique.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A description of this switch.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// The number of ports this switch has.
	Ports int32 `protobuf:"varint,3,opt,name=ports,proto3" json:"ports,omitempty"`
	// The state of this switch.
	State v1.State `protobuf:"varint,4,opt,name=state,proto3,enum=common.State" json:"state,omitempty"`
}

func (x *Switch) Reset() {
	*x = Switch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_machine_db_api_config_v1_datacenters_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Switch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Switch) ProtoMessage() {}

func (x *Switch) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_machine_db_api_config_v1_datacenters_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Switch.ProtoReflect.Descriptor instead.
func (*Switch) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_machine_db_api_config_v1_datacenters_proto_rawDescGZIP(), []int{0}
}

func (x *Switch) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Switch) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Switch) GetPorts() int32 {
	if x != nil {
		return x.Ports
	}
	return 0
}

func (x *Switch) GetState() v1.State {
	if x != nil {
		return x.State
	}
	return v1.State_STATE_UNSPECIFIED
}

// A KVM to store in the database.
type KVM struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of this KVM on the network. Must be globally unique.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A description of this KVM.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// The type of platform this KVM is.
	Platform string `protobuf:"bytes,3,opt,name=platform,proto3" json:"platform,omitempty"`
	// The rack this KVM belongs to.
	// Must be the name of a rack in the same datacenter.
	Rack string `protobuf:"bytes,4,opt,name=rack,proto3" json:"rack,omitempty"`
	// The MAC address associated with this KVM.
	MacAddress string `protobuf:"bytes,5,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
	// The IPv4 address associated with this KVM.
	Ipv4 string `protobuf:"bytes,6,opt,name=ipv4,proto3" json:"ipv4,omitempty"`
	// The state of this KVM.
	State v1.State `protobuf:"varint,7,opt,name=state,proto3,enum=common.State" json:"state,omitempty"`
}

func (x *KVM) Reset() {
	*x = KVM{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_machine_db_api_config_v1_datacenters_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KVM) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KVM) ProtoMessage() {}

func (x *KVM) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_machine_db_api_config_v1_datacenters_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KVM.ProtoReflect.Descriptor instead.
func (*KVM) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_machine_db_api_config_v1_datacenters_proto_rawDescGZIP(), []int{1}
}

func (x *KVM) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *KVM) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *KVM) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *KVM) GetRack() string {
	if x != nil {
		return x.Rack
	}
	return ""
}

func (x *KVM) GetMacAddress() string {
	if x != nil {
		return x.MacAddress
	}
	return ""
}

func (x *KVM) GetIpv4() string {
	if x != nil {
		return x.Ipv4
	}
	return ""
}

func (x *KVM) GetState() v1.State {
	if x != nil {
		return x.State
	}
	return v1.State_STATE_UNSPECIFIED
}

// A rack to store in the database.
type Rack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of this rack. Must be globally unique.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A description of this rack.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// The switches belonging to this rack.
	Switch []*Switch `protobuf:"bytes,3,rep,name=switch,proto3" json:"switch,omitempty"`
	// The state of this rack.
	State v1.State `protobuf:"varint,4,opt,name=state,proto3,enum=common.State" json:"state,omitempty"`
	// The KVM serving this rack.
	Kvm string `protobuf:"bytes,5,opt,name=kvm,proto3" json:"kvm,omitempty"`
}

func (x *Rack) Reset() {
	*x = Rack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_machine_db_api_config_v1_datacenters_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rack) ProtoMessage() {}

func (x *Rack) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_machine_db_api_config_v1_datacenters_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rack.ProtoReflect.Descriptor instead.
func (*Rack) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_machine_db_api_config_v1_datacenters_proto_rawDescGZIP(), []int{2}
}

func (x *Rack) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Rack) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Rack) GetSwitch() []*Switch {
	if x != nil {
		return x.Switch
	}
	return nil
}

func (x *Rack) GetState() v1.State {
	if x != nil {
		return x.State
	}
	return v1.State_STATE_UNSPECIFIED
}

func (x *Rack) GetKvm() string {
	if x != nil {
		return x.Kvm
	}
	return ""
}

// A datacenter to store in the database.
type Datacenter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of this datacenter. Must be globally unique.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A description of this datacenter.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// The racks belonging to this datacenter.
	Rack []*Rack `protobuf:"bytes,3,rep,name=rack,proto3" json:"rack,omitempty"`
	// The state of this datacenter.
	State v1.State `protobuf:"varint,4,opt,name=state,proto3,enum=common.State" json:"state,omitempty"`
	// The KVMs belonging to this datacenter.
	Kvm []*KVM `protobuf:"bytes,5,rep,name=kvm,proto3" json:"kvm,omitempty"`
}

func (x *Datacenter) Reset() {
	*x = Datacenter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_machine_db_api_config_v1_datacenters_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Datacenter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Datacenter) ProtoMessage() {}

func (x *Datacenter) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_machine_db_api_config_v1_datacenters_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Datacenter.ProtoReflect.Descriptor instead.
func (*Datacenter) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_machine_db_api_config_v1_datacenters_proto_rawDescGZIP(), []int{3}
}

func (x *Datacenter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Datacenter) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Datacenter) GetRack() []*Rack {
	if x != nil {
		return x.Rack
	}
	return nil
}

func (x *Datacenter) GetState() v1.State {
	if x != nil {
		return x.State
	}
	return v1.State_STATE_UNSPECIFIED
}

func (x *Datacenter) GetKvm() []*KVM {
	if x != nil {
		return x.Kvm
	}
	return nil
}

// A list of datacenter config files.
type Datacenters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of names of datacenter config files.
	Datacenter []string `protobuf:"bytes,1,rep,name=datacenter,proto3" json:"datacenter,omitempty"`
}

func (x *Datacenters) Reset() {
	*x = Datacenters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_machine_db_api_config_v1_datacenters_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Datacenters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Datacenters) ProtoMessage() {}

func (x *Datacenters) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_machine_db_api_config_v1_datacenters_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Datacenters.ProtoReflect.Descriptor instead.
func (*Datacenters) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_machine_db_api_config_v1_datacenters_proto_rawDescGZIP(), []int{4}
}

func (x *Datacenters) GetDatacenter() []string {
	if x != nil {
		return x.Datacenter
	}
	return nil
}

var File_go_chromium_org_luci_machine_db_api_config_v1_datacenters_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_machine_db_api_config_v1_datacenters_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2d, 0x64,
	0x62, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x3a, 0x67, 0x6f, 0x2e, 0x63, 0x68,
	0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f,
	0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2d, 0x64, 0x62, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x79, 0x0a, 0x06, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x22, 0xc5, 0x01, 0x0a, 0x03, 0x4b, 0x56, 0x4d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61,
	0x63, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x61, 0x63, 0x6b, 0x12, 0x1f,
	0x0a, 0x0b, 0x6d, 0x61, 0x63, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x61, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x69, 0x70, 0x76, 0x34, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69,
	0x70, 0x76, 0x34, 0x12, 0x23, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x9b, 0x01, 0x0a, 0x04, 0x52, 0x61, 0x63,
	0x6b, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x06, 0x73, 0x77, 0x69, 0x74, 0x63,
	0x68, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x52, 0x06, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x12,
	0x23, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x76, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x76, 0x6d, 0x22, 0xa8, 0x01, 0x0a, 0x0a, 0x44, 0x61, 0x74, 0x61, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x04, 0x72,
	0x61, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x52, 0x61, 0x63, 0x6b, 0x52, 0x04, 0x72, 0x61, 0x63, 0x6b, 0x12, 0x23, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x1d, 0x0a, 0x03, 0x6b, 0x76, 0x6d, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4b, 0x56, 0x4d, 0x52, 0x03, 0x6b, 0x76,
	0x6d, 0x22, 0x2d, 0x0a, 0x0b, 0x44, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73,
	0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x42, 0x36, 0x5a, 0x34, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e,
	0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x2d, 0x64, 0x62, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76,
	0x31, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_machine_db_api_config_v1_datacenters_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_machine_db_api_config_v1_datacenters_proto_rawDescData = file_go_chromium_org_luci_machine_db_api_config_v1_datacenters_proto_rawDesc
)

func file_go_chromium_org_luci_machine_db_api_config_v1_datacenters_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_machine_db_api_config_v1_datacenters_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_machine_db_api_config_v1_datacenters_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_machine_db_api_config_v1_datacenters_proto_rawDescData)
	})
	return file_go_chromium_org_luci_machine_db_api_config_v1_datacenters_proto_rawDescData
}

var file_go_chromium_org_luci_machine_db_api_config_v1_datacenters_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_go_chromium_org_luci_machine_db_api_config_v1_datacenters_proto_goTypes = []interface{}{
	(*Switch)(nil),      // 0: config.Switch
	(*KVM)(nil),         // 1: config.KVM
	(*Rack)(nil),        // 2: config.Rack
	(*Datacenter)(nil),  // 3: config.Datacenter
	(*Datacenters)(nil), // 4: config.Datacenters
	(v1.State)(0),       // 5: common.State
}
var file_go_chromium_org_luci_machine_db_api_config_v1_datacenters_proto_depIdxs = []int32{
	5, // 0: config.Switch.state:type_name -> common.State
	5, // 1: config.KVM.state:type_name -> common.State
	0, // 2: config.Rack.switch:type_name -> config.Switch
	5, // 3: config.Rack.state:type_name -> common.State
	2, // 4: config.Datacenter.rack:type_name -> config.Rack
	5, // 5: config.Datacenter.state:type_name -> common.State
	1, // 6: config.Datacenter.kvm:type_name -> config.KVM
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_machine_db_api_config_v1_datacenters_proto_init() }
func file_go_chromium_org_luci_machine_db_api_config_v1_datacenters_proto_init() {
	if File_go_chromium_org_luci_machine_db_api_config_v1_datacenters_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_machine_db_api_config_v1_datacenters_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Switch); i {
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
		file_go_chromium_org_luci_machine_db_api_config_v1_datacenters_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KVM); i {
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
		file_go_chromium_org_luci_machine_db_api_config_v1_datacenters_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rack); i {
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
		file_go_chromium_org_luci_machine_db_api_config_v1_datacenters_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Datacenter); i {
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
		file_go_chromium_org_luci_machine_db_api_config_v1_datacenters_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Datacenters); i {
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
			RawDescriptor: file_go_chromium_org_luci_machine_db_api_config_v1_datacenters_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_machine_db_api_config_v1_datacenters_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_machine_db_api_config_v1_datacenters_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_machine_db_api_config_v1_datacenters_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_machine_db_api_config_v1_datacenters_proto = out.File
	file_go_chromium_org_luci_machine_db_api_config_v1_datacenters_proto_rawDesc = nil
	file_go_chromium_org_luci_machine_db_api_config_v1_datacenters_proto_goTypes = nil
	file_go_chromium_org_luci_machine_db_api_config_v1_datacenters_proto_depIdxs = nil
}
