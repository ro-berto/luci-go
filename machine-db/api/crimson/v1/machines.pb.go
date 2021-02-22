// Copyright 2017 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.1
// source: go.chromium.org/luci/machine-db/api/crimson/v1/machines.proto

package crimson

import (
	v1 "go.chromium.org/luci/machine-db/api/common/v1"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
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

// A machine in the database.
type Machine struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of this machine. Uniquely identifies this machine.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The type of platform this machine is.
	Platform string `protobuf:"bytes,2,opt,name=platform,proto3" json:"platform,omitempty"`
	// The rack this machine belongs to.
	Rack string `protobuf:"bytes,3,opt,name=rack,proto3" json:"rack,omitempty"`
	// A description of this machine.
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// The asset tag associated with this machine.
	AssetTag string `protobuf:"bytes,5,opt,name=asset_tag,json=assetTag,proto3" json:"asset_tag,omitempty"`
	// The service tag associated with this machine.
	ServiceTag string `protobuf:"bytes,6,opt,name=service_tag,json=serviceTag,proto3" json:"service_tag,omitempty"`
	// The deployment ticket associated with this machine.
	DeploymentTicket string `protobuf:"bytes,7,opt,name=deployment_ticket,json=deploymentTicket,proto3" json:"deployment_ticket,omitempty"`
	// The state of this machine.
	State v1.State `protobuf:"varint,8,opt,name=state,proto3,enum=common.State" json:"state,omitempty"`
	// The datacenter this machine belongs to.
	// When creating a machine, omit this field. It will be inferred from the rack.
	Datacenter string `protobuf:"bytes,9,opt,name=datacenter,proto3" json:"datacenter,omitempty"`
	// The DRAC password associated with this machine.
	DracPassword string `protobuf:"bytes,10,opt,name=drac_password,json=dracPassword,proto3" json:"drac_password,omitempty"`
}

func (x *Machine) Reset() {
	*x = Machine{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_machine_db_api_crimson_v1_machines_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Machine) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Machine) ProtoMessage() {}

func (x *Machine) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_machine_db_api_crimson_v1_machines_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Machine.ProtoReflect.Descriptor instead.
func (*Machine) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_machine_db_api_crimson_v1_machines_proto_rawDescGZIP(), []int{0}
}

func (x *Machine) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Machine) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *Machine) GetRack() string {
	if x != nil {
		return x.Rack
	}
	return ""
}

func (x *Machine) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Machine) GetAssetTag() string {
	if x != nil {
		return x.AssetTag
	}
	return ""
}

func (x *Machine) GetServiceTag() string {
	if x != nil {
		return x.ServiceTag
	}
	return ""
}

func (x *Machine) GetDeploymentTicket() string {
	if x != nil {
		return x.DeploymentTicket
	}
	return ""
}

func (x *Machine) GetState() v1.State {
	if x != nil {
		return x.State
	}
	return v1.State_STATE_UNSPECIFIED
}

func (x *Machine) GetDatacenter() string {
	if x != nil {
		return x.Datacenter
	}
	return ""
}

func (x *Machine) GetDracPassword() string {
	if x != nil {
		return x.DracPassword
	}
	return ""
}

// A request to create a new machine in the database.
type CreateMachineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The machine to create in the database.
	Machine *Machine `protobuf:"bytes,1,opt,name=machine,proto3" json:"machine,omitempty"`
}

func (x *CreateMachineRequest) Reset() {
	*x = CreateMachineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_machine_db_api_crimson_v1_machines_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMachineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMachineRequest) ProtoMessage() {}

func (x *CreateMachineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_machine_db_api_crimson_v1_machines_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMachineRequest.ProtoReflect.Descriptor instead.
func (*CreateMachineRequest) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_machine_db_api_crimson_v1_machines_proto_rawDescGZIP(), []int{1}
}

func (x *CreateMachineRequest) GetMachine() *Machine {
	if x != nil {
		return x.Machine
	}
	return nil
}

// A request to delete a machine from the database.
type DeleteMachineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the machine to delete.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteMachineRequest) Reset() {
	*x = DeleteMachineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_machine_db_api_crimson_v1_machines_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMachineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMachineRequest) ProtoMessage() {}

func (x *DeleteMachineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_machine_db_api_crimson_v1_machines_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMachineRequest.ProtoReflect.Descriptor instead.
func (*DeleteMachineRequest) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_machine_db_api_crimson_v1_machines_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteMachineRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// A request to list machines in the database.
type ListMachinesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The names of machines to get.
	Names []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
	// The platforms to filter retrieved machines on.
	Platforms []string `protobuf:"bytes,2,rep,name=platforms,proto3" json:"platforms,omitempty"`
	// The racks to filter retrieved machines on.
	Racks []string `protobuf:"bytes,3,rep,name=racks,proto3" json:"racks,omitempty"`
	// The states to filter retrieved machines on.
	States []v1.State `protobuf:"varint,4,rep,packed,name=states,proto3,enum=common.State" json:"states,omitempty"`
	// The datacenters to filter retrieved machines on.
	Datacenters []string `protobuf:"bytes,5,rep,name=datacenters,proto3" json:"datacenters,omitempty"`
}

func (x *ListMachinesRequest) Reset() {
	*x = ListMachinesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_machine_db_api_crimson_v1_machines_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMachinesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMachinesRequest) ProtoMessage() {}

func (x *ListMachinesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_machine_db_api_crimson_v1_machines_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMachinesRequest.ProtoReflect.Descriptor instead.
func (*ListMachinesRequest) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_machine_db_api_crimson_v1_machines_proto_rawDescGZIP(), []int{3}
}

func (x *ListMachinesRequest) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

func (x *ListMachinesRequest) GetPlatforms() []string {
	if x != nil {
		return x.Platforms
	}
	return nil
}

func (x *ListMachinesRequest) GetRacks() []string {
	if x != nil {
		return x.Racks
	}
	return nil
}

func (x *ListMachinesRequest) GetStates() []v1.State {
	if x != nil {
		return x.States
	}
	return nil
}

func (x *ListMachinesRequest) GetDatacenters() []string {
	if x != nil {
		return x.Datacenters
	}
	return nil
}

// A response containing a list of machines in the database.
type ListMachinesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The machines matching this request.
	Machines []*Machine `protobuf:"bytes,1,rep,name=machines,proto3" json:"machines,omitempty"` // TODO(smut): Support page tokens.
}

func (x *ListMachinesResponse) Reset() {
	*x = ListMachinesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_machine_db_api_crimson_v1_machines_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMachinesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMachinesResponse) ProtoMessage() {}

func (x *ListMachinesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_machine_db_api_crimson_v1_machines_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMachinesResponse.ProtoReflect.Descriptor instead.
func (*ListMachinesResponse) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_machine_db_api_crimson_v1_machines_proto_rawDescGZIP(), []int{4}
}

func (x *ListMachinesResponse) GetMachines() []*Machine {
	if x != nil {
		return x.Machines
	}
	return nil
}

// A request to rename a machine in the database.
type RenameMachineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the machine to rename.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The new name to give this machine.
	NewName string `protobuf:"bytes,2,opt,name=new_name,json=newName,proto3" json:"new_name,omitempty"`
}

func (x *RenameMachineRequest) Reset() {
	*x = RenameMachineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_machine_db_api_crimson_v1_machines_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RenameMachineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenameMachineRequest) ProtoMessage() {}

func (x *RenameMachineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_machine_db_api_crimson_v1_machines_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenameMachineRequest.ProtoReflect.Descriptor instead.
func (*RenameMachineRequest) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_machine_db_api_crimson_v1_machines_proto_rawDescGZIP(), []int{5}
}

func (x *RenameMachineRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RenameMachineRequest) GetNewName() string {
	if x != nil {
		return x.NewName
	}
	return ""
}

// A request to update a machine in the database.
type UpdateMachineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The machine to update in the database.
	Machine *Machine `protobuf:"bytes,1,opt,name=machine,proto3" json:"machine,omitempty"`
	// The fields to update in the machine.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateMachineRequest) Reset() {
	*x = UpdateMachineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_machine_db_api_crimson_v1_machines_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMachineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMachineRequest) ProtoMessage() {}

func (x *UpdateMachineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_machine_db_api_crimson_v1_machines_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMachineRequest.ProtoReflect.Descriptor instead.
func (*UpdateMachineRequest) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_machine_db_api_crimson_v1_machines_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateMachineRequest) GetMachine() *Machine {
	if x != nil {
		return x.Machine
	}
	return nil
}

func (x *UpdateMachineRequest) GetUpdateMask() *field_mask.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

var File_go_chromium_org_luci_machine_db_api_crimson_v1_machines_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_machine_db_api_crimson_v1_machines_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2d, 0x64,
	0x62, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x72, 0x69, 0x6d, 0x73, 0x6f, 0x6e, 0x2f, 0x76, 0x31,
	0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x63, 0x72, 0x69, 0x6d, 0x73, 0x6f, 0x6e, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3a, 0x67, 0x6f, 0x2e, 0x63,
	0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69,
	0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2d, 0x64, 0x62, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc4, 0x02, 0x0a, 0x07, 0x4d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x72, 0x61, 0x63, 0x6b, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x5f, 0x74, 0x61, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x54, 0x61, 0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x74, 0x61, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x54, 0x61, 0x67, 0x12, 0x2b, 0x0a, 0x11, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x12, 0x23, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x61,
	0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x72, 0x61, 0x63,
	0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x64, 0x72, 0x61, 0x63, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x42, 0x0a,
	0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x07, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x72, 0x69, 0x6d, 0x73, 0x6f, 0x6e,
	0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x07, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x65, 0x22, 0x2a, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xa8, 0x01,
	0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x61, 0x63,
	0x6b, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x72, 0x61, 0x63, 0x6b, 0x73, 0x12,
	0x25, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0e, 0x32,
	0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x61, 0x74,
	0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x22, 0x44, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74,
	0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2c, 0x0a, 0x08, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x72, 0x69, 0x6d, 0x73, 0x6f, 0x6e, 0x2e, 0x4d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x52, 0x08, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x22, 0x45,
	0x0a, 0x14, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x65,
	0x77, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65,
	0x77, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x7f, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a,
	0x07, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x63, 0x72, 0x69, 0x6d, 0x73, 0x6f, 0x6e, 0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x52, 0x07, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_machine_db_api_crimson_v1_machines_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_machine_db_api_crimson_v1_machines_proto_rawDescData = file_go_chromium_org_luci_machine_db_api_crimson_v1_machines_proto_rawDesc
)

func file_go_chromium_org_luci_machine_db_api_crimson_v1_machines_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_machine_db_api_crimson_v1_machines_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_machine_db_api_crimson_v1_machines_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_machine_db_api_crimson_v1_machines_proto_rawDescData)
	})
	return file_go_chromium_org_luci_machine_db_api_crimson_v1_machines_proto_rawDescData
}

var file_go_chromium_org_luci_machine_db_api_crimson_v1_machines_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_go_chromium_org_luci_machine_db_api_crimson_v1_machines_proto_goTypes = []interface{}{
	(*Machine)(nil),              // 0: crimson.Machine
	(*CreateMachineRequest)(nil), // 1: crimson.CreateMachineRequest
	(*DeleteMachineRequest)(nil), // 2: crimson.DeleteMachineRequest
	(*ListMachinesRequest)(nil),  // 3: crimson.ListMachinesRequest
	(*ListMachinesResponse)(nil), // 4: crimson.ListMachinesResponse
	(*RenameMachineRequest)(nil), // 5: crimson.RenameMachineRequest
	(*UpdateMachineRequest)(nil), // 6: crimson.UpdateMachineRequest
	(v1.State)(0),                // 7: common.State
	(*field_mask.FieldMask)(nil), // 8: google.protobuf.FieldMask
}
var file_go_chromium_org_luci_machine_db_api_crimson_v1_machines_proto_depIdxs = []int32{
	7, // 0: crimson.Machine.state:type_name -> common.State
	0, // 1: crimson.CreateMachineRequest.machine:type_name -> crimson.Machine
	7, // 2: crimson.ListMachinesRequest.states:type_name -> common.State
	0, // 3: crimson.ListMachinesResponse.machines:type_name -> crimson.Machine
	0, // 4: crimson.UpdateMachineRequest.machine:type_name -> crimson.Machine
	8, // 5: crimson.UpdateMachineRequest.update_mask:type_name -> google.protobuf.FieldMask
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_machine_db_api_crimson_v1_machines_proto_init() }
func file_go_chromium_org_luci_machine_db_api_crimson_v1_machines_proto_init() {
	if File_go_chromium_org_luci_machine_db_api_crimson_v1_machines_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_machine_db_api_crimson_v1_machines_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Machine); i {
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
		file_go_chromium_org_luci_machine_db_api_crimson_v1_machines_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMachineRequest); i {
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
		file_go_chromium_org_luci_machine_db_api_crimson_v1_machines_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMachineRequest); i {
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
		file_go_chromium_org_luci_machine_db_api_crimson_v1_machines_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMachinesRequest); i {
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
		file_go_chromium_org_luci_machine_db_api_crimson_v1_machines_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMachinesResponse); i {
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
		file_go_chromium_org_luci_machine_db_api_crimson_v1_machines_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RenameMachineRequest); i {
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
		file_go_chromium_org_luci_machine_db_api_crimson_v1_machines_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMachineRequest); i {
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
			RawDescriptor: file_go_chromium_org_luci_machine_db_api_crimson_v1_machines_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_machine_db_api_crimson_v1_machines_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_machine_db_api_crimson_v1_machines_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_machine_db_api_crimson_v1_machines_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_machine_db_api_crimson_v1_machines_proto = out.File
	file_go_chromium_org_luci_machine_db_api_crimson_v1_machines_proto_rawDesc = nil
	file_go_chromium_org_luci_machine_db_api_crimson_v1_machines_proto_goTypes = nil
	file_go_chromium_org_luci_machine_db_api_crimson_v1_machines_proto_depIdxs = nil
}
