// Copyright 2019 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: go.chromium.org/luci/gce/api/instances/v1/service.proto

package instances

import prpc "go.chromium.org/luci/grpc/prpc"

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

// A disk associated with a GCE instance.
type Disk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The image associated with this disk.
	// https://cloud.google.com/compute/docs/reference/rest/v1/images/list.
	Image string `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *Disk) Reset() {
	*x = Disk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_gce_api_instances_v1_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Disk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Disk) ProtoMessage() {}

func (x *Disk) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_gce_api_instances_v1_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Disk.ProtoReflect.Descriptor instead.
func (*Disk) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_gce_api_instances_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *Disk) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

// A network interface associated with a GCE instance.
type NetworkInterface struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The internal IP address associated with this network interface.
	InternalIp string `protobuf:"bytes,1,opt,name=internal_ip,json=internalIp,proto3" json:"internal_ip,omitempty"`
	// The external IP addresses associated with this network interface.
	ExternalIps []string `protobuf:"bytes,2,rep,name=external_ips,json=externalIps,proto3" json:"external_ips,omitempty"`
}

func (x *NetworkInterface) Reset() {
	*x = NetworkInterface{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_gce_api_instances_v1_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkInterface) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkInterface) ProtoMessage() {}

func (x *NetworkInterface) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_gce_api_instances_v1_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkInterface.ProtoReflect.Descriptor instead.
func (*NetworkInterface) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_gce_api_instances_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *NetworkInterface) GetInternalIp() string {
	if x != nil {
		return x.InternalIp
	}
	return ""
}

func (x *NetworkInterface) GetExternalIps() []string {
	if x != nil {
		return x.ExternalIps
	}
	return nil
}

// A GCE instance configured to exist.
// The instance actually exists iff the created timestamp is set.
type Instance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the VM this instance was created from.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The hostname associated with this instance.
	Hostname string `protobuf:"bytes,2,opt,name=hostname,proto3" json:"hostname,omitempty"`
	// The zone associated with this instance.
	// https://cloud.google.com/compute/docs/reference/rest/v1/zones/list.
	Zone string `protobuf:"bytes,3,opt,name=zone,proto3" json:"zone,omitempty"`
	// The GCP project associated with this instance.
	Project string `protobuf:"bytes,4,opt,name=project,proto3" json:"project,omitempty"`
	// The timestamp when this instance was created.
	Created *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created,proto3" json:"created,omitempty"`
	// The lifetime of this instance.
	// At the end of its lifetime, the instance is deleted.
	Lifetime int64 `protobuf:"varint,6,opt,name=lifetime,proto3" json:"lifetime,omitempty"`
	// The hostname of the Swarming server this instance should connect to.
	Swarming string `protobuf:"bytes,7,opt,name=swarming,proto3" json:"swarming,omitempty"`
	// The timestamp when this instance connected to Swarming.
	Connected *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=connected,proto3" json:"connected,omitempty"`
	// The timeout of this instance.
	// If no Swarming bot has connected by the timeout, the instance is deleted.
	Timeout int64 `protobuf:"varint,9,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// Whether this instance is drained or not.
	// A drained instance will be safely deleted regardless of lifetime.
	Drained bool `protobuf:"varint,10,opt,name=drained,proto3" json:"drained,omitempty"`
	// The config revision associated with this instance.
	ConfigRevision string `protobuf:"bytes,11,opt,name=config_revision,json=configRevision,proto3" json:"config_revision,omitempty"`
	// The disks associated with this instance.
	Disks []*Disk `protobuf:"bytes,12,rep,name=disks,proto3" json:"disks,omitempty"`
	// The network interfaces associated with this instance.
	NetworkInterfaces []*NetworkInterface `protobuf:"bytes,13,rep,name=network_interfaces,json=networkInterfaces,proto3" json:"network_interfaces,omitempty"`
	// The prefix associated with this instance.
	Prefix string `protobuf:"bytes,14,opt,name=prefix,proto3" json:"prefix,omitempty"`
}

func (x *Instance) Reset() {
	*x = Instance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_gce_api_instances_v1_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Instance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Instance) ProtoMessage() {}

func (x *Instance) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_gce_api_instances_v1_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Instance.ProtoReflect.Descriptor instead.
func (*Instance) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_gce_api_instances_v1_service_proto_rawDescGZIP(), []int{2}
}

func (x *Instance) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Instance) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *Instance) GetZone() string {
	if x != nil {
		return x.Zone
	}
	return ""
}

func (x *Instance) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *Instance) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *Instance) GetLifetime() int64 {
	if x != nil {
		return x.Lifetime
	}
	return 0
}

func (x *Instance) GetSwarming() string {
	if x != nil {
		return x.Swarming
	}
	return ""
}

func (x *Instance) GetConnected() *timestamppb.Timestamp {
	if x != nil {
		return x.Connected
	}
	return nil
}

func (x *Instance) GetTimeout() int64 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

func (x *Instance) GetDrained() bool {
	if x != nil {
		return x.Drained
	}
	return false
}

func (x *Instance) GetConfigRevision() string {
	if x != nil {
		return x.ConfigRevision
	}
	return ""
}

func (x *Instance) GetDisks() []*Disk {
	if x != nil {
		return x.Disks
	}
	return nil
}

func (x *Instance) GetNetworkInterfaces() []*NetworkInterface {
	if x != nil {
		return x.NetworkInterfaces
	}
	return nil
}

func (x *Instance) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

// A request to delete an instance.
type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the instance to delete.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The hostname of the instance to delete.
	Hostname string `protobuf:"bytes,2,opt,name=hostname,proto3" json:"hostname,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_gce_api_instances_v1_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_gce_api_instances_v1_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_gce_api_instances_v1_service_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeleteRequest) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

// A request to get an existing instance.
type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the instance to get.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The hostname of the instance to get.
	Hostname string `protobuf:"bytes,2,opt,name=hostname,proto3" json:"hostname,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_gce_api_instances_v1_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_gce_api_instances_v1_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_gce_api_instances_v1_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetRequest) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

// A request to list existing instances.
type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The prefix to list instances for.
	Prefix string `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	// The value of next_page_token received in a ListResponse. Used to get the
	// next page of instances. If empty, gets the first page.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// The maximum number of results to include in the response.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// A filter to apply when listing instances. Currently the only supported
	// filter is "disks.image=<image>" where <image> is the name of the image to
	// filter for.
	Filter string `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_gce_api_instances_v1_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_gce_api_instances_v1_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_gce_api_instances_v1_service_proto_rawDescGZIP(), []int{5}
}

func (x *ListRequest) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

func (x *ListRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

// A response to a request to list instances.
type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The prefix the instances are for.
	Prefix string `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	// The instances matching the prefix.
	Instances []*Instance `protobuf:"bytes,2,rep,name=instances,proto3" json:"instances,omitempty"`
	// The value to use as the page_token in a ListRequest to get the next page of
	// instances. If empty, there are no more instances.
	NextPageToken string `protobuf:"bytes,3,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_gce_api_instances_v1_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_gce_api_instances_v1_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_gce_api_instances_v1_service_proto_rawDescGZIP(), []int{6}
}

func (x *ListResponse) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

func (x *ListResponse) GetInstances() []*Instance {
	if x != nil {
		return x.Instances
	}
	return nil
}

func (x *ListResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_go_chromium_org_luci_gce_api_instances_v1_service_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_gce_api_instances_v1_service_proto_rawDesc = []byte{
	0x0a, 0x37, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x67, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x1c, 0x0a, 0x04, 0x44, 0x69, 0x73, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x22, 0x56, 0x0a, 0x10, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x5f, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x49, 0x70, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x5f, 0x69, 0x70, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x70, 0x73, 0x22, 0xf4, 0x03, 0x0a, 0x08, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x34, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x77, 0x61, 0x72, 0x6d, 0x69, 0x6e, 0x67, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x77, 0x61, 0x72, 0x6d, 0x69, 0x6e, 0x67, 0x12, 0x38, 0x0a,
	0x09, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x64, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x05, 0x64, 0x69, 0x73, 0x6b, 0x73, 0x18, 0x0c, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2e,
	0x44, 0x69, 0x73, 0x6b, 0x52, 0x05, 0x64, 0x69, 0x73, 0x6b, 0x73, 0x12, 0x4a, 0x0a, 0x12, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x73, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x52, 0x11, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x22,
	0x3b, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x38, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f,
	0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f,
	0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x79, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x22, 0x81, 0x01, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x31, 0x0a, 0x09, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x52, 0x09, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x26, 0x0a,
	0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xb3, 0x01, 0x0a, 0x09, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x73, 0x12, 0x3a, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x18, 0x2e,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12,
	0x31, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x15, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x12, 0x37, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x17, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x35, 0x5a, 0x33, 0x67,
	0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c,
	0x75, 0x63, 0x69, 0x2f, 0x67, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_gce_api_instances_v1_service_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_gce_api_instances_v1_service_proto_rawDescData = file_go_chromium_org_luci_gce_api_instances_v1_service_proto_rawDesc
)

func file_go_chromium_org_luci_gce_api_instances_v1_service_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_gce_api_instances_v1_service_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_gce_api_instances_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_gce_api_instances_v1_service_proto_rawDescData)
	})
	return file_go_chromium_org_luci_gce_api_instances_v1_service_proto_rawDescData
}

var file_go_chromium_org_luci_gce_api_instances_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_go_chromium_org_luci_gce_api_instances_v1_service_proto_goTypes = []interface{}{
	(*Disk)(nil),                  // 0: instances.Disk
	(*NetworkInterface)(nil),      // 1: instances.NetworkInterface
	(*Instance)(nil),              // 2: instances.Instance
	(*DeleteRequest)(nil),         // 3: instances.DeleteRequest
	(*GetRequest)(nil),            // 4: instances.GetRequest
	(*ListRequest)(nil),           // 5: instances.ListRequest
	(*ListResponse)(nil),          // 6: instances.ListResponse
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 8: google.protobuf.Empty
}
var file_go_chromium_org_luci_gce_api_instances_v1_service_proto_depIdxs = []int32{
	7, // 0: instances.Instance.created:type_name -> google.protobuf.Timestamp
	7, // 1: instances.Instance.connected:type_name -> google.protobuf.Timestamp
	0, // 2: instances.Instance.disks:type_name -> instances.Disk
	1, // 3: instances.Instance.network_interfaces:type_name -> instances.NetworkInterface
	2, // 4: instances.ListResponse.instances:type_name -> instances.Instance
	3, // 5: instances.Instances.Delete:input_type -> instances.DeleteRequest
	4, // 6: instances.Instances.Get:input_type -> instances.GetRequest
	5, // 7: instances.Instances.List:input_type -> instances.ListRequest
	8, // 8: instances.Instances.Delete:output_type -> google.protobuf.Empty
	2, // 9: instances.Instances.Get:output_type -> instances.Instance
	6, // 10: instances.Instances.List:output_type -> instances.ListResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_gce_api_instances_v1_service_proto_init() }
func file_go_chromium_org_luci_gce_api_instances_v1_service_proto_init() {
	if File_go_chromium_org_luci_gce_api_instances_v1_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_gce_api_instances_v1_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Disk); i {
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
		file_go_chromium_org_luci_gce_api_instances_v1_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkInterface); i {
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
		file_go_chromium_org_luci_gce_api_instances_v1_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Instance); i {
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
		file_go_chromium_org_luci_gce_api_instances_v1_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
		file_go_chromium_org_luci_gce_api_instances_v1_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_go_chromium_org_luci_gce_api_instances_v1_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_go_chromium_org_luci_gce_api_instances_v1_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
			RawDescriptor: file_go_chromium_org_luci_gce_api_instances_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_go_chromium_org_luci_gce_api_instances_v1_service_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_gce_api_instances_v1_service_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_gce_api_instances_v1_service_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_gce_api_instances_v1_service_proto = out.File
	file_go_chromium_org_luci_gce_api_instances_v1_service_proto_rawDesc = nil
	file_go_chromium_org_luci_gce_api_instances_v1_service_proto_goTypes = nil
	file_go_chromium_org_luci_gce_api_instances_v1_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// InstancesClient is the client API for Instances service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InstancesClient interface {
	// Delete deletes an instance asynchronously.
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Get returns an existing instance.
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Instance, error)
	// List returns existing instances.
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
}
type instancesPRPCClient struct {
	client *prpc.Client
}

func NewInstancesPRPCClient(client *prpc.Client) InstancesClient {
	return &instancesPRPCClient{client}
}

func (c *instancesPRPCClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.client.Call(ctx, "instances.Instances", "Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instancesPRPCClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Instance, error) {
	out := new(Instance)
	err := c.client.Call(ctx, "instances.Instances", "Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instancesPRPCClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.client.Call(ctx, "instances.Instances", "List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type instancesClient struct {
	cc grpc.ClientConnInterface
}

func NewInstancesClient(cc grpc.ClientConnInterface) InstancesClient {
	return &instancesClient{cc}
}

func (c *instancesClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/instances.Instances/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instancesClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Instance, error) {
	out := new(Instance)
	err := c.cc.Invoke(ctx, "/instances.Instances/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instancesClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/instances.Instances/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InstancesServer is the server API for Instances service.
type InstancesServer interface {
	// Delete deletes an instance asynchronously.
	Delete(context.Context, *DeleteRequest) (*emptypb.Empty, error)
	// Get returns an existing instance.
	Get(context.Context, *GetRequest) (*Instance, error)
	// List returns existing instances.
	List(context.Context, *ListRequest) (*ListResponse, error)
}

// UnimplementedInstancesServer can be embedded to have forward compatible implementations.
type UnimplementedInstancesServer struct {
}

func (*UnimplementedInstancesServer) Delete(context.Context, *DeleteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedInstancesServer) Get(context.Context, *GetRequest) (*Instance, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedInstancesServer) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterInstancesServer(s prpc.Registrar, srv InstancesServer) {
	s.RegisterService(&_Instances_serviceDesc, srv)
}

func _Instances_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstancesServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/instances.Instances/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstancesServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Instances_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstancesServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/instances.Instances/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstancesServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Instances_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstancesServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/instances.Instances/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstancesServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Instances_serviceDesc = grpc.ServiceDesc{
	ServiceName: "instances.Instances",
	HandlerType: (*InstancesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Delete",
			Handler:    _Instances_Delete_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Instances_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Instances_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go.chromium.org/luci/gce/api/instances/v1/service.proto",
}
