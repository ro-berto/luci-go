// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/gce/api/instances/v1/service.proto

package instances

import prpc "go.chromium.org/luci/grpc/prpc"

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// A disk associated with a GCE instance.
type Disk struct {
	// The image associated with this disk.
	// https://cloud.google.com/compute/docs/reference/rest/v1/images/list.
	Image                string   `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Disk) Reset()         { *m = Disk{} }
func (m *Disk) String() string { return proto.CompactTextString(m) }
func (*Disk) ProtoMessage()    {}
func (*Disk) Descriptor() ([]byte, []int) {
	return fileDescriptor_85758a446589b372, []int{0}
}

func (m *Disk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Disk.Unmarshal(m, b)
}
func (m *Disk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Disk.Marshal(b, m, deterministic)
}
func (m *Disk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Disk.Merge(m, src)
}
func (m *Disk) XXX_Size() int {
	return xxx_messageInfo_Disk.Size(m)
}
func (m *Disk) XXX_DiscardUnknown() {
	xxx_messageInfo_Disk.DiscardUnknown(m)
}

var xxx_messageInfo_Disk proto.InternalMessageInfo

func (m *Disk) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

// A GCE instance configured to exist.
// The instance actually exists iff the created timestamp is set.
type Instance struct {
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
	Created *timestamp.Timestamp `protobuf:"bytes,5,opt,name=created,proto3" json:"created,omitempty"`
	// The lifetime of this instance.
	// At the end of its lifetime, the instance is deleted.
	Lifetime int64 `protobuf:"varint,6,opt,name=lifetime,proto3" json:"lifetime,omitempty"`
	// The hostname of the Swarming server this instance should connect to.
	Swarming string `protobuf:"bytes,7,opt,name=swarming,proto3" json:"swarming,omitempty"`
	// The timestamp when this instance connected to Swarming.
	Connected *timestamp.Timestamp `protobuf:"bytes,8,opt,name=connected,proto3" json:"connected,omitempty"`
	// The timeout of this instance.
	// If no Swarming bot has connected by the timeout, the instance is deleted.
	Timeout int64 `protobuf:"varint,9,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// Whether this instance is drained or not.
	// A drained instance will be safely deleted regardless of lifetime.
	Drained bool `protobuf:"varint,10,opt,name=drained,proto3" json:"drained,omitempty"`
	// The config revision associated with this instance.
	ConfigRevision string `protobuf:"bytes,11,opt,name=config_revision,json=configRevision,proto3" json:"config_revision,omitempty"`
	// The disks associated with this instance.
	Disks                []*Disk  `protobuf:"bytes,12,rep,name=disks,proto3" json:"disks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Instance) Reset()         { *m = Instance{} }
func (m *Instance) String() string { return proto.CompactTextString(m) }
func (*Instance) ProtoMessage()    {}
func (*Instance) Descriptor() ([]byte, []int) {
	return fileDescriptor_85758a446589b372, []int{1}
}

func (m *Instance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Instance.Unmarshal(m, b)
}
func (m *Instance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Instance.Marshal(b, m, deterministic)
}
func (m *Instance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Instance.Merge(m, src)
}
func (m *Instance) XXX_Size() int {
	return xxx_messageInfo_Instance.Size(m)
}
func (m *Instance) XXX_DiscardUnknown() {
	xxx_messageInfo_Instance.DiscardUnknown(m)
}

var xxx_messageInfo_Instance proto.InternalMessageInfo

func (m *Instance) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Instance) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *Instance) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

func (m *Instance) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *Instance) GetCreated() *timestamp.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *Instance) GetLifetime() int64 {
	if m != nil {
		return m.Lifetime
	}
	return 0
}

func (m *Instance) GetSwarming() string {
	if m != nil {
		return m.Swarming
	}
	return ""
}

func (m *Instance) GetConnected() *timestamp.Timestamp {
	if m != nil {
		return m.Connected
	}
	return nil
}

func (m *Instance) GetTimeout() int64 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *Instance) GetDrained() bool {
	if m != nil {
		return m.Drained
	}
	return false
}

func (m *Instance) GetConfigRevision() string {
	if m != nil {
		return m.ConfigRevision
	}
	return ""
}

func (m *Instance) GetDisks() []*Disk {
	if m != nil {
		return m.Disks
	}
	return nil
}

// A request to delete an instance.
type DeleteRequest struct {
	// The ID of the instance to delete.
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_85758a446589b372, []int{2}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// A request to get an existing instance.
type GetRequest struct {
	// The ID of the instance to get.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The hostname of the instance to get.
	Hostname             string   `protobuf:"bytes,2,opt,name=hostname,proto3" json:"hostname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_85758a446589b372, []int{3}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetRequest) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

// A request to list existing instances.
type ListRequest struct {
	// The prefix to list instances for.
	Prefix string `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	// The value of next_page_token received in a ListResponse. Used to get the
	// next page of instances. If empty, gets the first page.
	PageToken            string   `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_85758a446589b372, []int{4}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *ListRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// A response to a request to list instances.
type ListResponse struct {
	// The prefix the instances are for.
	Prefix string `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	// The instances matching the prefix.
	Instances []*Instance `protobuf:"bytes,2,rep,name=instances,proto3" json:"instances,omitempty"`
	// The value to use as the page_token in a ListRequest to get the next page of
	// instances. If empty, there are no more instances.
	NextPageToken        string   `protobuf:"bytes,3,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_85758a446589b372, []int{5}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *ListResponse) GetInstances() []*Instance {
	if m != nil {
		return m.Instances
	}
	return nil
}

func (m *ListResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

func init() {
	proto.RegisterType((*Disk)(nil), "instances.Disk")
	proto.RegisterType((*Instance)(nil), "instances.Instance")
	proto.RegisterType((*DeleteRequest)(nil), "instances.DeleteRequest")
	proto.RegisterType((*GetRequest)(nil), "instances.GetRequest")
	proto.RegisterType((*ListRequest)(nil), "instances.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "instances.ListResponse")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/gce/api/instances/v1/service.proto", fileDescriptor_85758a446589b372)
}

var fileDescriptor_85758a446589b372 = []byte{
	// 522 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x95, 0x93, 0x34, 0x89, 0x27, 0x6d, 0x23, 0x2d, 0x10, 0x56, 0x06, 0xd4, 0xc8, 0x12, 0x90,
	0x93, 0xad, 0x04, 0xa4, 0x56, 0x9c, 0x83, 0xaa, 0x4a, 0x1c, 0x90, 0xd5, 0x7b, 0xe4, 0xda, 0x13,
	0x77, 0x49, 0xbc, 0x6b, 0x76, 0x37, 0xa1, 0x70, 0xe3, 0x1f, 0xf0, 0x5f, 0xf8, 0x83, 0x68, 0xbd,
	0xb6, 0x63, 0xb5, 0x14, 0x6e, 0x7e, 0xef, 0xcd, 0xce, 0xbc, 0xf9, 0x30, 0x9c, 0x67, 0x22, 0x48,
	0x6e, 0xa5, 0xc8, 0xd9, 0x2e, 0x0f, 0x84, 0xcc, 0xc2, 0xed, 0x2e, 0x61, 0x61, 0x96, 0x60, 0x18,
	0x17, 0x2c, 0x64, 0x5c, 0xe9, 0x98, 0x27, 0xa8, 0xc2, 0xfd, 0x3c, 0x54, 0x28, 0xf7, 0x2c, 0xc1,
	0xa0, 0x90, 0x42, 0x0b, 0xe2, 0x36, 0x9a, 0xf7, 0x22, 0x13, 0x22, 0xdb, 0x62, 0x58, 0x0a, 0x37,
	0xbb, 0x75, 0x88, 0x79, 0xa1, 0xbf, 0xdb, 0x38, 0xef, 0xec, 0xbe, 0xa8, 0x59, 0x8e, 0x4a, 0xc7,
	0x79, 0x61, 0x03, 0xfc, 0x97, 0xd0, 0x5b, 0x32, 0xb5, 0x21, 0x4f, 0xe1, 0x88, 0xe5, 0x71, 0x86,
	0xd4, 0x99, 0x3a, 0x33, 0x37, 0xb2, 0xc0, 0xff, 0xd5, 0x85, 0xe1, 0x55, 0x55, 0x89, 0x9c, 0x42,
	0x87, 0xa5, 0x95, 0xde, 0x61, 0x29, 0xf1, 0x60, 0x78, 0x2b, 0x94, 0xe6, 0x71, 0x8e, 0xb4, 0x53,
	0xb2, 0x0d, 0x26, 0x04, 0x7a, 0x3f, 0x04, 0x47, 0xda, 0x2d, 0xf9, 0xf2, 0x9b, 0x50, 0x18, 0x14,
	0x52, 0x7c, 0xc1, 0x44, 0xd3, 0x5e, 0x49, 0xd7, 0x90, 0xbc, 0x87, 0x41, 0x22, 0x31, 0xd6, 0x98,
	0xd2, 0xa3, 0xa9, 0x33, 0x1b, 0x2d, 0xbc, 0xc0, 0xfa, 0x0e, 0x6a, 0xdf, 0xc1, 0x75, 0xed, 0x3b,
	0xaa, 0x43, 0x4d, 0xfd, 0x2d, 0x5b, 0xa3, 0xe9, 0x88, 0xf6, 0xa7, 0xce, 0xac, 0x1b, 0x35, 0xd8,
	0x68, 0xea, 0x5b, 0x2c, 0x73, 0xc6, 0x33, 0x3a, 0xb0, 0xde, 0x6a, 0x4c, 0x2e, 0xc0, 0x4d, 0x04,
	0xe7, 0x98, 0x98, 0x7a, 0xc3, 0xff, 0xd6, 0x3b, 0x04, 0x9b, 0x0e, 0x4c, 0x76, 0xb1, 0xd3, 0xd4,
	0x2d, 0x0b, 0xd6, 0xd0, 0x28, 0xa9, 0x8c, 0x19, 0xc7, 0x94, 0xc2, 0xd4, 0x99, 0x0d, 0xa3, 0x1a,
	0x92, 0xb7, 0x30, 0x4e, 0x04, 0x5f, 0xb3, 0x6c, 0x25, 0x71, 0xcf, 0x14, 0x13, 0x9c, 0x8e, 0x4a,
	0x43, 0xa7, 0x96, 0x8e, 0x2a, 0x96, 0xbc, 0x86, 0xa3, 0x94, 0xa9, 0x8d, 0xa2, 0xc7, 0xd3, 0xee,
	0x6c, 0xb4, 0x18, 0x07, 0xcd, 0x8a, 0x03, 0xb3, 0xa1, 0xc8, 0xaa, 0xfe, 0x19, 0x9c, 0x2c, 0x71,
	0x8b, 0x1a, 0x23, 0xfc, 0xba, 0x43, 0xa5, 0xef, 0xaf, 0xc5, 0xbf, 0x00, 0xb8, 0x44, 0xfd, 0x88,
	0xfa, 0xaf, 0xa5, 0xf9, 0x4b, 0x18, 0x7d, 0x62, 0xaa, 0x79, 0x3a, 0x81, 0x7e, 0x21, 0x71, 0xcd,
	0xee, 0xaa, 0xe7, 0x15, 0x22, 0xaf, 0x00, 0x8a, 0x38, 0xc3, 0x95, 0x16, 0x1b, 0xe4, 0x55, 0x12,
	0xd7, 0x30, 0xd7, 0x86, 0xf0, 0x7f, 0x3a, 0x70, 0x6c, 0xd3, 0xa8, 0x42, 0x70, 0x85, 0x8f, 0xe6,
	0x99, 0xc3, 0xe1, 0x8a, 0x69, 0xa7, 0x6c, 0xfa, 0x49, 0xab, 0xe9, 0xfa, 0xee, 0xa2, 0x43, 0x14,
	0x79, 0x03, 0x63, 0x8e, 0x77, 0x7a, 0xd5, 0xaa, 0x6f, 0x2f, 0xec, 0xc4, 0xd0, 0x9f, 0x6b, 0x0f,
	0x8b, 0xdf, 0x0e, 0xb8, 0x57, 0xcd, 0xab, 0x0f, 0xd0, 0xb7, 0x23, 0x23, 0xb4, 0x3d, 0xd4, 0xf6,
	0x14, 0xbd, 0xc9, 0x83, 0x0b, 0xf8, 0x68, 0x7e, 0x23, 0x32, 0x87, 0xee, 0x25, 0x6a, 0xf2, 0xac,
	0xf5, 0xf0, 0x30, 0x5d, 0xef, 0x6f, 0x7e, 0xc9, 0x39, 0xf4, 0x4c, 0xff, 0x64, 0xd2, 0x12, 0x5b,
	0x73, 0xf5, 0x9e, 0x3f, 0xe0, 0xed, 0xa0, 0x6e, 0xfa, 0x65, 0xed, 0x77, 0x7f, 0x02, 0x00, 0x00,
	0xff, 0xff, 0xe7, 0x6b, 0xeb, 0x7e, 0x16, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// InstancesClient is the client API for Instances service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InstancesClient interface {
	// Delete deletes an instance asynchronously.
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*empty.Empty, error)
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

func (c *instancesPRPCClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
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
	cc *grpc.ClientConn
}

func NewInstancesClient(cc *grpc.ClientConn) InstancesClient {
	return &instancesClient{cc}
}

func (c *instancesClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
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
	Delete(context.Context, *DeleteRequest) (*empty.Empty, error)
	// Get returns an existing instance.
	Get(context.Context, *GetRequest) (*Instance, error)
	// List returns existing instances.
	List(context.Context, *ListRequest) (*ListResponse, error)
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
