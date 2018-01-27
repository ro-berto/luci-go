// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/machine-db/api/crimson/v1/crimson.proto

/*
Package crimson is a generated protocol buffer package.

It is generated from these files:
	go.chromium.org/luci/machine-db/api/crimson/v1/crimson.proto
	go.chromium.org/luci/machine-db/api/crimson/v1/datacenters.proto
	go.chromium.org/luci/machine-db/api/crimson/v1/machines.proto
	go.chromium.org/luci/machine-db/api/crimson/v1/nics.proto
	go.chromium.org/luci/machine-db/api/crimson/v1/oses.proto
	go.chromium.org/luci/machine-db/api/crimson/v1/physical_hosts.proto
	go.chromium.org/luci/machine-db/api/crimson/v1/platforms.proto
	go.chromium.org/luci/machine-db/api/crimson/v1/racks.proto
	go.chromium.org/luci/machine-db/api/crimson/v1/switches.proto
	go.chromium.org/luci/machine-db/api/crimson/v1/vlans.proto
	go.chromium.org/luci/machine-db/api/crimson/v1/vms.proto

It has these top-level messages:
	ListDatacentersRequest
	Datacenter
	ListDatacentersResponse
	Machine
	CreateMachineRequest
	DeleteMachineRequest
	ListMachinesRequest
	ListMachinesResponse
	NIC
	CreateNICRequest
	ListNICsRequest
	ListNICsResponse
	ListOSesRequest
	OS
	ListOSesResponse
	PhysicalHost
	CreatePhysicalHostRequest
	ListPhysicalHostsRequest
	ListPhysicalHostsResponse
	ListPlatformsRequest
	Platform
	ListPlatformsResponse
	Rack
	ListRacksRequest
	ListRacksResponse
	Switch
	ListSwitchesRequest
	ListSwitchesResponse
	ListVLANsRequest
	VLAN
	ListVLANsResponse
	VM
	CreateVMRequest
	ListVMsRequest
	ListVMsResponse
*/
package crimson

import prpc "go.chromium.org/luci/grpc/prpc"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Crimson service

type CrimsonClient interface {
	// Lists datacenters in the database.
	ListDatacenters(ctx context.Context, in *ListDatacentersRequest, opts ...grpc.CallOption) (*ListDatacentersResponse, error)
	// Lists operating systems in the database.
	ListOSes(ctx context.Context, in *ListOSesRequest, opts ...grpc.CallOption) (*ListOSesResponse, error)
	// Lists platforms in the database.
	ListPlatforms(ctx context.Context, in *ListPlatformsRequest, opts ...grpc.CallOption) (*ListPlatformsResponse, error)
	// Lists racks in the database.
	ListRacks(ctx context.Context, in *ListRacksRequest, opts ...grpc.CallOption) (*ListRacksResponse, error)
	// Lists switches in the database.
	ListSwitches(ctx context.Context, in *ListSwitchesRequest, opts ...grpc.CallOption) (*ListSwitchesResponse, error)
	// Lists VLANs in the database.
	ListVLANs(ctx context.Context, in *ListVLANsRequest, opts ...grpc.CallOption) (*ListVLANsResponse, error)
	// Creates a new machine in the database.
	CreateMachine(ctx context.Context, in *CreateMachineRequest, opts ...grpc.CallOption) (*Machine, error)
	// Deletes a machine from the database.
	DeleteMachine(ctx context.Context, in *DeleteMachineRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	// Lists machines in the database.
	ListMachines(ctx context.Context, in *ListMachinesRequest, opts ...grpc.CallOption) (*ListMachinesResponse, error)
	// Creates a new NIC in the database.
	CreateNIC(ctx context.Context, in *CreateNICRequest, opts ...grpc.CallOption) (*NIC, error)
	// Lists NICs in the database.
	ListNICs(ctx context.Context, in *ListNICsRequest, opts ...grpc.CallOption) (*ListNICsResponse, error)
	// Creates a new physical host in the database.
	CreatePhysicalHost(ctx context.Context, in *CreatePhysicalHostRequest, opts ...grpc.CallOption) (*PhysicalHost, error)
	// Lists physical hosts in the database.
	ListPhysicalHosts(ctx context.Context, in *ListPhysicalHostsRequest, opts ...grpc.CallOption) (*ListPhysicalHostsResponse, error)
	// Creates a new VM in the database.
	CreateVM(ctx context.Context, in *CreateVMRequest, opts ...grpc.CallOption) (*VM, error)
	// Lists VMs in the database.
	ListVMs(ctx context.Context, in *ListVMsRequest, opts ...grpc.CallOption) (*ListVMsResponse, error)
}
type crimsonPRPCClient struct {
	client *prpc.Client
}

func NewCrimsonPRPCClient(client *prpc.Client) CrimsonClient {
	return &crimsonPRPCClient{client}
}

func (c *crimsonPRPCClient) ListDatacenters(ctx context.Context, in *ListDatacentersRequest, opts ...grpc.CallOption) (*ListDatacentersResponse, error) {
	out := new(ListDatacentersResponse)
	err := c.client.Call(ctx, "crimson.Crimson", "ListDatacenters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crimsonPRPCClient) ListOSes(ctx context.Context, in *ListOSesRequest, opts ...grpc.CallOption) (*ListOSesResponse, error) {
	out := new(ListOSesResponse)
	err := c.client.Call(ctx, "crimson.Crimson", "ListOSes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crimsonPRPCClient) ListPlatforms(ctx context.Context, in *ListPlatformsRequest, opts ...grpc.CallOption) (*ListPlatformsResponse, error) {
	out := new(ListPlatformsResponse)
	err := c.client.Call(ctx, "crimson.Crimson", "ListPlatforms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crimsonPRPCClient) ListRacks(ctx context.Context, in *ListRacksRequest, opts ...grpc.CallOption) (*ListRacksResponse, error) {
	out := new(ListRacksResponse)
	err := c.client.Call(ctx, "crimson.Crimson", "ListRacks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crimsonPRPCClient) ListSwitches(ctx context.Context, in *ListSwitchesRequest, opts ...grpc.CallOption) (*ListSwitchesResponse, error) {
	out := new(ListSwitchesResponse)
	err := c.client.Call(ctx, "crimson.Crimson", "ListSwitches", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crimsonPRPCClient) ListVLANs(ctx context.Context, in *ListVLANsRequest, opts ...grpc.CallOption) (*ListVLANsResponse, error) {
	out := new(ListVLANsResponse)
	err := c.client.Call(ctx, "crimson.Crimson", "ListVLANs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crimsonPRPCClient) CreateMachine(ctx context.Context, in *CreateMachineRequest, opts ...grpc.CallOption) (*Machine, error) {
	out := new(Machine)
	err := c.client.Call(ctx, "crimson.Crimson", "CreateMachine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crimsonPRPCClient) DeleteMachine(ctx context.Context, in *DeleteMachineRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := c.client.Call(ctx, "crimson.Crimson", "DeleteMachine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crimsonPRPCClient) ListMachines(ctx context.Context, in *ListMachinesRequest, opts ...grpc.CallOption) (*ListMachinesResponse, error) {
	out := new(ListMachinesResponse)
	err := c.client.Call(ctx, "crimson.Crimson", "ListMachines", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crimsonPRPCClient) CreateNIC(ctx context.Context, in *CreateNICRequest, opts ...grpc.CallOption) (*NIC, error) {
	out := new(NIC)
	err := c.client.Call(ctx, "crimson.Crimson", "CreateNIC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crimsonPRPCClient) ListNICs(ctx context.Context, in *ListNICsRequest, opts ...grpc.CallOption) (*ListNICsResponse, error) {
	out := new(ListNICsResponse)
	err := c.client.Call(ctx, "crimson.Crimson", "ListNICs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crimsonPRPCClient) CreatePhysicalHost(ctx context.Context, in *CreatePhysicalHostRequest, opts ...grpc.CallOption) (*PhysicalHost, error) {
	out := new(PhysicalHost)
	err := c.client.Call(ctx, "crimson.Crimson", "CreatePhysicalHost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crimsonPRPCClient) ListPhysicalHosts(ctx context.Context, in *ListPhysicalHostsRequest, opts ...grpc.CallOption) (*ListPhysicalHostsResponse, error) {
	out := new(ListPhysicalHostsResponse)
	err := c.client.Call(ctx, "crimson.Crimson", "ListPhysicalHosts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crimsonPRPCClient) CreateVM(ctx context.Context, in *CreateVMRequest, opts ...grpc.CallOption) (*VM, error) {
	out := new(VM)
	err := c.client.Call(ctx, "crimson.Crimson", "CreateVM", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crimsonPRPCClient) ListVMs(ctx context.Context, in *ListVMsRequest, opts ...grpc.CallOption) (*ListVMsResponse, error) {
	out := new(ListVMsResponse)
	err := c.client.Call(ctx, "crimson.Crimson", "ListVMs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type crimsonClient struct {
	cc *grpc.ClientConn
}

func NewCrimsonClient(cc *grpc.ClientConn) CrimsonClient {
	return &crimsonClient{cc}
}

func (c *crimsonClient) ListDatacenters(ctx context.Context, in *ListDatacentersRequest, opts ...grpc.CallOption) (*ListDatacentersResponse, error) {
	out := new(ListDatacentersResponse)
	err := grpc.Invoke(ctx, "/crimson.Crimson/ListDatacenters", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crimsonClient) ListOSes(ctx context.Context, in *ListOSesRequest, opts ...grpc.CallOption) (*ListOSesResponse, error) {
	out := new(ListOSesResponse)
	err := grpc.Invoke(ctx, "/crimson.Crimson/ListOSes", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crimsonClient) ListPlatforms(ctx context.Context, in *ListPlatformsRequest, opts ...grpc.CallOption) (*ListPlatformsResponse, error) {
	out := new(ListPlatformsResponse)
	err := grpc.Invoke(ctx, "/crimson.Crimson/ListPlatforms", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crimsonClient) ListRacks(ctx context.Context, in *ListRacksRequest, opts ...grpc.CallOption) (*ListRacksResponse, error) {
	out := new(ListRacksResponse)
	err := grpc.Invoke(ctx, "/crimson.Crimson/ListRacks", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crimsonClient) ListSwitches(ctx context.Context, in *ListSwitchesRequest, opts ...grpc.CallOption) (*ListSwitchesResponse, error) {
	out := new(ListSwitchesResponse)
	err := grpc.Invoke(ctx, "/crimson.Crimson/ListSwitches", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crimsonClient) ListVLANs(ctx context.Context, in *ListVLANsRequest, opts ...grpc.CallOption) (*ListVLANsResponse, error) {
	out := new(ListVLANsResponse)
	err := grpc.Invoke(ctx, "/crimson.Crimson/ListVLANs", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crimsonClient) CreateMachine(ctx context.Context, in *CreateMachineRequest, opts ...grpc.CallOption) (*Machine, error) {
	out := new(Machine)
	err := grpc.Invoke(ctx, "/crimson.Crimson/CreateMachine", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crimsonClient) DeleteMachine(ctx context.Context, in *DeleteMachineRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/crimson.Crimson/DeleteMachine", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crimsonClient) ListMachines(ctx context.Context, in *ListMachinesRequest, opts ...grpc.CallOption) (*ListMachinesResponse, error) {
	out := new(ListMachinesResponse)
	err := grpc.Invoke(ctx, "/crimson.Crimson/ListMachines", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crimsonClient) CreateNIC(ctx context.Context, in *CreateNICRequest, opts ...grpc.CallOption) (*NIC, error) {
	out := new(NIC)
	err := grpc.Invoke(ctx, "/crimson.Crimson/CreateNIC", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crimsonClient) ListNICs(ctx context.Context, in *ListNICsRequest, opts ...grpc.CallOption) (*ListNICsResponse, error) {
	out := new(ListNICsResponse)
	err := grpc.Invoke(ctx, "/crimson.Crimson/ListNICs", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crimsonClient) CreatePhysicalHost(ctx context.Context, in *CreatePhysicalHostRequest, opts ...grpc.CallOption) (*PhysicalHost, error) {
	out := new(PhysicalHost)
	err := grpc.Invoke(ctx, "/crimson.Crimson/CreatePhysicalHost", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crimsonClient) ListPhysicalHosts(ctx context.Context, in *ListPhysicalHostsRequest, opts ...grpc.CallOption) (*ListPhysicalHostsResponse, error) {
	out := new(ListPhysicalHostsResponse)
	err := grpc.Invoke(ctx, "/crimson.Crimson/ListPhysicalHosts", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crimsonClient) CreateVM(ctx context.Context, in *CreateVMRequest, opts ...grpc.CallOption) (*VM, error) {
	out := new(VM)
	err := grpc.Invoke(ctx, "/crimson.Crimson/CreateVM", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crimsonClient) ListVMs(ctx context.Context, in *ListVMsRequest, opts ...grpc.CallOption) (*ListVMsResponse, error) {
	out := new(ListVMsResponse)
	err := grpc.Invoke(ctx, "/crimson.Crimson/ListVMs", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Crimson service

type CrimsonServer interface {
	// Lists datacenters in the database.
	ListDatacenters(context.Context, *ListDatacentersRequest) (*ListDatacentersResponse, error)
	// Lists operating systems in the database.
	ListOSes(context.Context, *ListOSesRequest) (*ListOSesResponse, error)
	// Lists platforms in the database.
	ListPlatforms(context.Context, *ListPlatformsRequest) (*ListPlatformsResponse, error)
	// Lists racks in the database.
	ListRacks(context.Context, *ListRacksRequest) (*ListRacksResponse, error)
	// Lists switches in the database.
	ListSwitches(context.Context, *ListSwitchesRequest) (*ListSwitchesResponse, error)
	// Lists VLANs in the database.
	ListVLANs(context.Context, *ListVLANsRequest) (*ListVLANsResponse, error)
	// Creates a new machine in the database.
	CreateMachine(context.Context, *CreateMachineRequest) (*Machine, error)
	// Deletes a machine from the database.
	DeleteMachine(context.Context, *DeleteMachineRequest) (*google_protobuf.Empty, error)
	// Lists machines in the database.
	ListMachines(context.Context, *ListMachinesRequest) (*ListMachinesResponse, error)
	// Creates a new NIC in the database.
	CreateNIC(context.Context, *CreateNICRequest) (*NIC, error)
	// Lists NICs in the database.
	ListNICs(context.Context, *ListNICsRequest) (*ListNICsResponse, error)
	// Creates a new physical host in the database.
	CreatePhysicalHost(context.Context, *CreatePhysicalHostRequest) (*PhysicalHost, error)
	// Lists physical hosts in the database.
	ListPhysicalHosts(context.Context, *ListPhysicalHostsRequest) (*ListPhysicalHostsResponse, error)
	// Creates a new VM in the database.
	CreateVM(context.Context, *CreateVMRequest) (*VM, error)
	// Lists VMs in the database.
	ListVMs(context.Context, *ListVMsRequest) (*ListVMsResponse, error)
}

func RegisterCrimsonServer(s prpc.Registrar, srv CrimsonServer) {
	s.RegisterService(&_Crimson_serviceDesc, srv)
}

func _Crimson_ListDatacenters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDatacentersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrimsonServer).ListDatacenters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crimson.Crimson/ListDatacenters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrimsonServer).ListDatacenters(ctx, req.(*ListDatacentersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Crimson_ListOSes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOSesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrimsonServer).ListOSes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crimson.Crimson/ListOSes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrimsonServer).ListOSes(ctx, req.(*ListOSesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Crimson_ListPlatforms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPlatformsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrimsonServer).ListPlatforms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crimson.Crimson/ListPlatforms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrimsonServer).ListPlatforms(ctx, req.(*ListPlatformsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Crimson_ListRacks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRacksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrimsonServer).ListRacks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crimson.Crimson/ListRacks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrimsonServer).ListRacks(ctx, req.(*ListRacksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Crimson_ListSwitches_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSwitchesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrimsonServer).ListSwitches(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crimson.Crimson/ListSwitches",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrimsonServer).ListSwitches(ctx, req.(*ListSwitchesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Crimson_ListVLANs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListVLANsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrimsonServer).ListVLANs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crimson.Crimson/ListVLANs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrimsonServer).ListVLANs(ctx, req.(*ListVLANsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Crimson_CreateMachine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMachineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrimsonServer).CreateMachine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crimson.Crimson/CreateMachine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrimsonServer).CreateMachine(ctx, req.(*CreateMachineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Crimson_DeleteMachine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMachineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrimsonServer).DeleteMachine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crimson.Crimson/DeleteMachine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrimsonServer).DeleteMachine(ctx, req.(*DeleteMachineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Crimson_ListMachines_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMachinesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrimsonServer).ListMachines(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crimson.Crimson/ListMachines",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrimsonServer).ListMachines(ctx, req.(*ListMachinesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Crimson_CreateNIC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNICRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrimsonServer).CreateNIC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crimson.Crimson/CreateNIC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrimsonServer).CreateNIC(ctx, req.(*CreateNICRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Crimson_ListNICs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNICsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrimsonServer).ListNICs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crimson.Crimson/ListNICs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrimsonServer).ListNICs(ctx, req.(*ListNICsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Crimson_CreatePhysicalHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePhysicalHostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrimsonServer).CreatePhysicalHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crimson.Crimson/CreatePhysicalHost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrimsonServer).CreatePhysicalHost(ctx, req.(*CreatePhysicalHostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Crimson_ListPhysicalHosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPhysicalHostsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrimsonServer).ListPhysicalHosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crimson.Crimson/ListPhysicalHosts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrimsonServer).ListPhysicalHosts(ctx, req.(*ListPhysicalHostsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Crimson_CreateVM_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVMRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrimsonServer).CreateVM(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crimson.Crimson/CreateVM",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrimsonServer).CreateVM(ctx, req.(*CreateVMRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Crimson_ListVMs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListVMsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrimsonServer).ListVMs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crimson.Crimson/ListVMs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrimsonServer).ListVMs(ctx, req.(*ListVMsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Crimson_serviceDesc = grpc.ServiceDesc{
	ServiceName: "crimson.Crimson",
	HandlerType: (*CrimsonServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListDatacenters",
			Handler:    _Crimson_ListDatacenters_Handler,
		},
		{
			MethodName: "ListOSes",
			Handler:    _Crimson_ListOSes_Handler,
		},
		{
			MethodName: "ListPlatforms",
			Handler:    _Crimson_ListPlatforms_Handler,
		},
		{
			MethodName: "ListRacks",
			Handler:    _Crimson_ListRacks_Handler,
		},
		{
			MethodName: "ListSwitches",
			Handler:    _Crimson_ListSwitches_Handler,
		},
		{
			MethodName: "ListVLANs",
			Handler:    _Crimson_ListVLANs_Handler,
		},
		{
			MethodName: "CreateMachine",
			Handler:    _Crimson_CreateMachine_Handler,
		},
		{
			MethodName: "DeleteMachine",
			Handler:    _Crimson_DeleteMachine_Handler,
		},
		{
			MethodName: "ListMachines",
			Handler:    _Crimson_ListMachines_Handler,
		},
		{
			MethodName: "CreateNIC",
			Handler:    _Crimson_CreateNIC_Handler,
		},
		{
			MethodName: "ListNICs",
			Handler:    _Crimson_ListNICs_Handler,
		},
		{
			MethodName: "CreatePhysicalHost",
			Handler:    _Crimson_CreatePhysicalHost_Handler,
		},
		{
			MethodName: "ListPhysicalHosts",
			Handler:    _Crimson_ListPhysicalHosts_Handler,
		},
		{
			MethodName: "CreateVM",
			Handler:    _Crimson_CreateVM_Handler,
		},
		{
			MethodName: "ListVMs",
			Handler:    _Crimson_ListVMs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go.chromium.org/luci/machine-db/api/crimson/v1/crimson.proto",
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/machine-db/api/crimson/v1/crimson.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 512 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xdf, 0x6f, 0xd3, 0x30,
	0x10, 0xc7, 0xdf, 0x68, 0x67, 0x56, 0x01, 0x96, 0x80, 0x51, 0x18, 0x3f, 0xf6, 0x4e, 0xa2, 0x01,
	0x0f, 0x80, 0x06, 0x0c, 0x32, 0x10, 0x13, 0x4b, 0x37, 0x6d, 0xa8, 0x0f, 0xbc, 0x20, 0xd7, 0xf3,
	0x1a, 0x8b, 0x24, 0x0e, 0xb6, 0x3b, 0xb4, 0x3f, 0x8b, 0xff, 0x10, 0x39, 0xf6, 0xb9, 0x71, 0x48,
	0x27, 0xa5, 0x6f, 0xed, 0xf7, 0x7b, 0xf7, 0xb9, 0xcb, 0xdd, 0xc9, 0x68, 0x6f, 0x2e, 0x22, 0x9a,
	0x49, 0x51, 0xf0, 0x45, 0x11, 0x09, 0x39, 0x8f, 0xf3, 0x05, 0xe5, 0x71, 0x41, 0x68, 0xc6, 0x4b,
	0xf6, 0xfc, 0x7c, 0x16, 0x93, 0x8a, 0xc7, 0x54, 0xf2, 0x42, 0x89, 0x32, 0xbe, 0xdc, 0x85, 0x9f,
	0x51, 0x25, 0x85, 0x16, 0x78, 0xe0, 0xfe, 0x8e, 0x1f, 0xce, 0x85, 0x98, 0xe7, 0x2c, 0xae, 0xe5,
	0xd9, 0xe2, 0x22, 0x66, 0x45, 0xa5, 0xaf, 0x6c, 0xd4, 0x78, 0xbf, 0x67, 0x8d, 0x73, 0xa2, 0x09,
	0x65, 0xa5, 0x66, 0x52, 0x39, 0xc2, 0xbb, 0x9e, 0x04, 0xe7, 0x40, 0xfa, 0x9b, 0x9e, 0xe9, 0x25,
	0xa7, 0xeb, 0xa6, 0x0a, 0xe5, 0xab, 0x26, 0x3d, 0x53, 0xab, 0xec, 0x4a, 0x71, 0x4a, 0xf2, 0x9f,
	0x99, 0x50, 0x1a, 0x20, 0xef, 0xfb, 0x42, 0x72, 0xa2, 0x2f, 0x84, 0x2c, 0x20, 0xff, 0x6d, 0xcf,
	0x7c, 0x49, 0xe8, 0xaf, 0x75, 0xa7, 0xae, 0xfe, 0x70, 0x4d, 0x33, 0xb6, 0x6e, 0xe9, 0xcb, 0x9c,
	0x94, 0x90, 0xfb, 0xba, 0x6f, 0x2e, 0x7c, 0xf0, 0x8b, 0xbf, 0x43, 0x34, 0x48, 0xac, 0x81, 0xbf,
	0xa3, 0x5b, 0x47, 0x5c, 0xe9, 0x83, 0xe5, 0x3d, 0xe1, 0x27, 0x11, 0x5c, 0x70, 0xcb, 0x39, 0x65,
	0xbf, 0x17, 0x4c, 0xe9, 0xf1, 0xd3, 0xd5, 0x01, 0xaa, 0x12, 0xa5, 0x62, 0xf8, 0x03, 0x1a, 0x1a,
	0xeb, 0xf8, 0x8c, 0x29, 0xbc, 0x15, 0x44, 0x1b, 0x09, 0x38, 0x0f, 0x3a, 0x1c, 0x07, 0x98, 0xa0,
	0x91, 0xd1, 0x4e, 0x60, 0x55, 0x78, 0x3b, 0x88, 0xf5, 0x3a, 0xa0, 0x1e, 0xaf, 0xb2, 0x1d, 0xef,
	0x13, 0xda, 0x30, 0xc6, 0xa9, 0x59, 0x1d, 0x0e, 0xeb, 0xd6, 0x1a, 0x70, 0xc6, 0x5d, 0x96, 0x63,
	0x7c, 0x43, 0x9b, 0x46, 0x3c, 0x73, 0x2b, 0xc4, 0x8f, 0x82, 0x58, 0x90, 0x81, 0xb4, 0xbd, 0xc2,
	0x0d, 0x1b, 0x9a, 0x1e, 0x7d, 0x9c, 0xb4, 0x1b, 0xaa, 0xb5, 0xee, 0x86, 0x9c, 0xe5, 0x18, 0xfb,
	0x68, 0x94, 0x48, 0x46, 0x34, 0x4b, 0xed, 0xd2, 0x1b, 0x43, 0x0a, 0x74, 0x60, 0xdd, 0xf6, 0x36,
	0x24, 0x7c, 0x41, 0xa3, 0x03, 0x96, 0xb3, 0x2e, 0x42, 0xa0, 0x03, 0xe1, 0x5e, 0x64, 0x1f, 0xb1,
	0x08, 0x1e, 0xb1, 0xe8, 0xb3, 0x79, 0xc4, 0x60, 0x34, 0x2e, 0xba, 0x3d, 0x1a, 0x90, 0xbb, 0x47,
	0xb3, 0x74, 0xdd, 0x67, 0xbd, 0x42, 0x1b, 0xb6, 0xfd, 0xc9, 0x61, 0xd2, 0x18, 0x8d, 0xd7, 0x00,
	0xb3, 0xe9, 0x2d, 0x13, 0xe8, 0x4e, 0x6e, 0x72, 0x98, 0xb4, 0x4f, 0xce, 0x48, 0xdd, 0x27, 0x67,
	0x1d, 0x57, 0xf6, 0x18, 0x61, 0x5b, 0xe2, 0xc4, 0x3d, 0x32, 0x5f, 0x85, 0xd2, 0x78, 0xa7, 0x55,
	0xbf, 0x69, 0x02, 0xf4, 0xae, 0x8f, 0x09, 0x52, 0x7f, 0xa0, 0x3b, 0xf5, 0x31, 0x36, 0x34, 0x85,
	0x9f, 0x85, 0x87, 0xda, 0xf4, 0x00, 0xb7, 0x73, 0x5d, 0x88, 0x6b, 0x76, 0x17, 0x0d, 0x6d, 0x3f,
	0xd3, 0xb4, 0xf1, 0xb5, 0x20, 0x01, 0xe9, 0xa6, 0x77, 0xa6, 0x29, 0xde, 0x43, 0x83, 0xfa, 0x84,
	0x52, 0x85, 0xef, 0x87, 0x47, 0x95, 0xfa, 0xd2, 0x5b, 0xff, 0x1b, 0xb6, 0xe0, 0xec, 0x46, 0xbd,
	0xf1, 0x97, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x66, 0xaa, 0x85, 0x6e, 0x0d, 0x07, 0x00, 0x00,
}
