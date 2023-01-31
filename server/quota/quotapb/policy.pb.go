// Copyright 2022 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: go.chromium.org/luci/server/quota/quotapb/policy.proto

package quotapb

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Policy_Options int32

const (
	Policy_NO_OPTIONS Policy_Options = 0
	// Indicates that this Policy covers a resource type which represents an
	// absolute quantity (e.g. number of builds in flight, current amount of
	// storage used, etc.). Accounts flagged with this option cannot be manually
	// manipulated via the Admin API, even with `quota.accounts.write`
	// permission. Applications which need to expose 'reset' functionality for
	// these should expose their own endpoints for this (or, ideally, don't
	// allow these Accounts to get out of sync with reality in the first place
	// :))
	Policy_ABSOLUTE_RESOURCE Policy_Options = 1
)

// Enum value maps for Policy_Options.
var (
	Policy_Options_name = map[int32]string{
		0: "NO_OPTIONS",
		1: "ABSOLUTE_RESOURCE",
	}
	Policy_Options_value = map[string]int32{
		"NO_OPTIONS":        0,
		"ABSOLUTE_RESOURCE": 1,
	}
)

func (x Policy_Options) Enum() *Policy_Options {
	p := new(Policy_Options)
	*p = x
	return p
}

func (x Policy_Options) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Policy_Options) Descriptor() protoreflect.EnumDescriptor {
	return file_go_chromium_org_luci_server_quota_quotapb_policy_proto_enumTypes[0].Descriptor()
}

func (Policy_Options) Type() protoreflect.EnumType {
	return &file_go_chromium_org_luci_server_quota_quotapb_policy_proto_enumTypes[0]
}

func (x Policy_Options) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Policy_Options.Descriptor instead.
func (Policy_Options) EnumDescriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_server_quota_quotapb_policy_proto_rawDescGZIP(), []int{0, 0}
}

// A Policy represents a single quota policy.
//
// A single Policy will typically be used to govern many Accounts.
//
// Policies are always loaded into the database within a PolicyConfig.
type Policy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The number of resources to assign to an Account when first accessed under
	// this Policy.
	//
	// Must be <= `limit`.
	Default uint64 `protobuf:"varint,1,opt,name=default,proto3" json:"default,omitempty"`
	// The maximum balance of Accounts managed under this Policy.
	//
	// Operations with a positive delta will be capped to this value, unless they
	// specify `POLICY_HARD_LIMIT`, in which case exceeding this limit will be an
	// error.
	//
	// If this policy has a positive refill, accounts with this policy will
	// gradually fill to this limit over time (but will never refill past it).
	//
	// NOTE: When assigning a new Policy to an existing Account it's possible for
	// an Account balance to exceed this value.
	//
	// For example, say an Account had a balance of 100 under a Policy with
	// a limit of 100, but then you set a new Policy with a limit of 50. In this
	// case, the Account balance remains at 100. However, the Account would not
	// gain any additional refill under the new Policy until it was brought below
	// 50 (the new limit).
	//
	// This is done because applications using the quota library may not have full
	// consistency with their Policy choice (e.g. they may choose a Policy based
	// on group membership, which is volatile, or some application nodes may have
	// gotten configuration to use a new Policy version while others haven't).
	Limit  uint64         `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Refill *Policy_Refill `protobuf:"bytes,3,opt,name=refill,proto3" json:"refill,omitempty"`
	// Bitwise-OR of Options values.
	Options int32 `protobuf:"varint,4,opt,name=options,proto3" json:"options,omitempty"`
	// The amount of time that Accounts created with this Policy should persist
	// after being written. Each Op on the Account refreshes the timeout.
	//
	// This could be used to create temporary quota Accounts based on e.g. IP
	// address which automatically garbage collect after a certain time.
	//
	// A value of 0 means an 'infinite' Account lifetime (the default).
	// It's recommended to pick some very large value for this rather than 0, to
	// allow Redis to prune old Accounts when it needs to do garbage collection.
	Lifetime *durationpb.Duration `protobuf:"bytes,5,opt,name=lifetime,proto3" json:"lifetime,omitempty"`
}

func (x *Policy) Reset() {
	*x = Policy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_server_quota_quotapb_policy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Policy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Policy) ProtoMessage() {}

func (x *Policy) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_server_quota_quotapb_policy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Policy.ProtoReflect.Descriptor instead.
func (*Policy) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_server_quota_quotapb_policy_proto_rawDescGZIP(), []int{0}
}

func (x *Policy) GetDefault() uint64 {
	if x != nil {
		return x.Default
	}
	return 0
}

func (x *Policy) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *Policy) GetRefill() *Policy_Refill {
	if x != nil {
		return x.Refill
	}
	return nil
}

func (x *Policy) GetOptions() int32 {
	if x != nil {
		return x.Options
	}
	return 0
}

func (x *Policy) GetLifetime() *durationpb.Duration {
	if x != nil {
		return x.Lifetime
	}
	return nil
}

// Refill describes how Accounts under this Policy refill (or drain) over
// time.
//
// The Refill process mimics a cron, starting at UTC midnight + offset, waking
// up every `interval` seconds to add `units` to the Account balance (up to
// `limit`). This refill operation only happens when an Account is actually
// interacted with, however.
type Policy_Refill struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The number of units to add to the Account banance every `interval`.
	//
	// The refill process is discrete; From T0..T0+interval, none of
	// these units will appear in the Account balance. At T0+interval, all
	// the units will be added.
	//
	// Note that it's permitted to have a negative refill `units` to have
	// Account balances drain back to 0 over time.
	//
	// It's not permitted for the units to be 0 (just omit the Refill message
	// entirely in that case).
	Units int64 `protobuf:"varint,1,opt,name=units,proto3" json:"units,omitempty"`
	// The number of seconds between refill events, synchronized to UTC midnight
	// + `offset`.
	//
	// If this is 0 and `units` is positive, the Account will be treated as if
	// it always has `limit` quota.
	//
	// It is an error for this to be 0 with negative `units`. To achieve this,
	// just make a Policy with a limit of 0 and no Refill.
	//
	// Refill events occur synchronized to "midnight" in UTC. So if you set this
	// to 60, then each minute-after-UTC-midnight, the Account will gain
	// `units`. This synchronization makes quota Account refill more
	// predictable.
	//
	// The offset from UTC is currently configed on the Policy (i.e. to support
	// policies which are synched with different time zones), but this
	// presumably could instead be configured on a per-Account basis, if it were
	// needed.
	//
	// This MUST evenly divide 24h (86400). For example, an interval of 71 is
	// NOT OK because it would divide the day into 1216.9 intervals, meaning
	// that the refresh 'cycle' could not correctly reset at midnight every day.
	// An interval of 72 IS ok though, because it evenly divides the day into
	// 1200 refresh periods.
	Interval uint32 `protobuf:"varint,2,opt,name=interval,proto3" json:"interval,omitempty"`
	// An offset from UTC midnight. This will be used to establish when the
	// associated Accounts 'start their day', and can be used to implement
	// a rudimentary timezone alignment for quota Accounts.
	Offset uint32 `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *Policy_Refill) Reset() {
	*x = Policy_Refill{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_server_quota_quotapb_policy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Policy_Refill) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Policy_Refill) ProtoMessage() {}

func (x *Policy_Refill) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_server_quota_quotapb_policy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Policy_Refill.ProtoReflect.Descriptor instead.
func (*Policy_Refill) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_server_quota_quotapb_policy_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Policy_Refill) GetUnits() int64 {
	if x != nil {
		return x.Units
	}
	return 0
}

func (x *Policy_Refill) GetInterval() uint32 {
	if x != nil {
		return x.Interval
	}
	return 0
}

func (x *Policy_Refill) GetOffset() uint32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

var File_go_chromium_org_luci_server_quota_quotapb_policy_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_server_quota_quotapb_policy_proto_rawDesc = []byte{
	0x0a, 0x36, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x71, 0x75,
	0x6f, 0x74, 0x61, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x70, 0x62, 0x2f, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x29, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72,
	0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x75, 0x63, 0x69, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x71, 0x75, 0x6f, 0x74,
	0x61, 0x70, 0x62, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb5, 0x03, 0x0a,
	0x06, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x28, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x0e, 0xfa, 0x42, 0x0b, 0x32, 0x09, 0x18,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x0f, 0x52, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x12, 0x24, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x42, 0x0e, 0xfa, 0x42, 0x0b, 0x32, 0x09, 0x18, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x0f,
	0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x50, 0x0a, 0x06, 0x72, 0x65, 0x66, 0x69, 0x6c,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72,
	0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x75, 0x63, 0x69, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x71, 0x75, 0x6f, 0x74,
	0x61, 0x70, 0x62, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x52, 0x65, 0x66, 0x69, 0x6c,
	0x6c, 0x52, 0x06, 0x72, 0x65, 0x66, 0x69, 0x6c, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x35, 0x0a, 0x08, 0x6c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x08, 0x6c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x1a, 0x85, 0x01, 0x0a, 0x06, 0x52,
	0x65, 0x66, 0x69, 0x6c, 0x6c, 0x12, 0x31, 0x0a, 0x05, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x42, 0x1b, 0xfa, 0x42, 0x18, 0x22, 0x16, 0x18, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0x0f, 0x28, 0x81, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0xf0, 0xff, 0x01, 0x38,
	0x00, 0x52, 0x05, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x12, 0x25, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x2a,
	0x04, 0x18, 0x80, 0xa3, 0x05, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12,
	0x21, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x42,
	0x09, 0xfa, 0x42, 0x06, 0x2a, 0x04, 0x18, 0x80, 0xa3, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x22, 0x30, 0x0a, 0x07, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x0e, 0x0a,
	0x0a, 0x4e, 0x4f, 0x5f, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x10, 0x00, 0x12, 0x15, 0x0a,
	0x11, 0x41, 0x42, 0x53, 0x4f, 0x4c, 0x55, 0x54, 0x45, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52,
	0x43, 0x45, 0x10, 0x01, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d,
	0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_server_quota_quotapb_policy_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_server_quota_quotapb_policy_proto_rawDescData = file_go_chromium_org_luci_server_quota_quotapb_policy_proto_rawDesc
)

func file_go_chromium_org_luci_server_quota_quotapb_policy_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_server_quota_quotapb_policy_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_server_quota_quotapb_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_server_quota_quotapb_policy_proto_rawDescData)
	})
	return file_go_chromium_org_luci_server_quota_quotapb_policy_proto_rawDescData
}

var file_go_chromium_org_luci_server_quota_quotapb_policy_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_go_chromium_org_luci_server_quota_quotapb_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_go_chromium_org_luci_server_quota_quotapb_policy_proto_goTypes = []interface{}{
	(Policy_Options)(0),         // 0: go.chromium.org.luci.server.quota.quotapb.Policy.Options
	(*Policy)(nil),              // 1: go.chromium.org.luci.server.quota.quotapb.Policy
	(*Policy_Refill)(nil),       // 2: go.chromium.org.luci.server.quota.quotapb.Policy.Refill
	(*durationpb.Duration)(nil), // 3: google.protobuf.Duration
}
var file_go_chromium_org_luci_server_quota_quotapb_policy_proto_depIdxs = []int32{
	2, // 0: go.chromium.org.luci.server.quota.quotapb.Policy.refill:type_name -> go.chromium.org.luci.server.quota.quotapb.Policy.Refill
	3, // 1: go.chromium.org.luci.server.quota.quotapb.Policy.lifetime:type_name -> google.protobuf.Duration
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_server_quota_quotapb_policy_proto_init() }
func file_go_chromium_org_luci_server_quota_quotapb_policy_proto_init() {
	if File_go_chromium_org_luci_server_quota_quotapb_policy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_server_quota_quotapb_policy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Policy); i {
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
		file_go_chromium_org_luci_server_quota_quotapb_policy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Policy_Refill); i {
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
			RawDescriptor: file_go_chromium_org_luci_server_quota_quotapb_policy_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_server_quota_quotapb_policy_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_server_quota_quotapb_policy_proto_depIdxs,
		EnumInfos:         file_go_chromium_org_luci_server_quota_quotapb_policy_proto_enumTypes,
		MessageInfos:      file_go_chromium_org_luci_server_quota_quotapb_policy_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_server_quota_quotapb_policy_proto = out.File
	file_go_chromium_org_luci_server_quota_quotapb_policy_proto_rawDesc = nil
	file_go_chromium_org_luci_server_quota_quotapb_policy_proto_goTypes = nil
	file_go_chromium_org_luci_server_quota_quotapb_policy_proto_depIdxs = nil
}
