// Copyright 2020 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Messages that describe internal representation of LUCI Realms.
//
// They are produced and distributed by the LUCI Auth service based on a high
// level representation fetched via LUCI Config from realms.cfg project config
// files. See realms_config.proto in the auth service source code for schema.
//
// Comments in this file is the authoritative documentation of how LUCI services
// should interpret realms when making authorization decisions.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: go.chromium.org/luci/server/auth/service/protocol/components/auth/proto/realms.proto

package protocol

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

// Realms is a complete definition of all known permissions and realms in a LUCI
// deployment.
//
// It is generated and distributed across all LUCI services (as part of the
// AuthDB) by the LUCI Auth service.
//
// Note that this is a denormalized internal representation of realms which is
// derived from the high level user-facing representation supplied via multiple
// realms.cfg config files in various config sets. See comments for Realm
// message for details.
//
// The internal representation doesn't have a notion of roles or realm
// inheritance. These concepts are handled by the LUCI Auth service and
// individual downstream services generally **must not care** how it works.
// Instead they should follow the rules outlined in comments in this file (or
// equivalently just use the Realms API exposed by the LUCI auth libraries).
//
// Next ID: 5.
type Realms struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// API version is incremented whenever the semantic meaning of Realms message
	// changes in some backward incompatible way (e.g. some message grows a new
	// field that *must* be checked by services). LUCI services must reject Realms
	// messages that have API versions they don't recognize. It is a precaution
	// against misinterpreting the realms configuration.
	//
	// The current version is 1.
	ApiVersion int64 `protobuf:"varint,1,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	// List of all possible permissions in alphabetical order.
	//
	// Acts as a universal set of permissions in Binding messages.
	//
	// Services may also use this field to check that permissions they are about
	// to use are actually known to the LUCI auth system. This is useful for
	// debugging when adding or removing permissions.
	//
	// See Permission message for more details.
	Permissions []*Permission `protobuf:"bytes,2,rep,name=permissions,proto3" json:"permissions,omitempty"`
	// List of all conditions referenced by bindings.
	//
	// Bindings reference them by their zero-based index in this list.
	Conditions []*Condition `protobuf:"bytes,4,rep,name=conditions,proto3" json:"conditions,omitempty"`
	// List of all registered realms in alphabetical order.
	//
	// See Realm message for more details.
	Realms []*Realm `protobuf:"bytes,3,rep,name=realms,proto3" json:"realms,omitempty"`
}

func (x *Realms) Reset() {
	*x = Realms{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_realms_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Realms) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Realms) ProtoMessage() {}

func (x *Realms) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_realms_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Realms.ProtoReflect.Descriptor instead.
func (*Realms) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_realms_proto_rawDescGZIP(), []int{0}
}

func (x *Realms) GetApiVersion() int64 {
	if x != nil {
		return x.ApiVersion
	}
	return 0
}

func (x *Realms) GetPermissions() []*Permission {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *Realms) GetConditions() []*Condition {
	if x != nil {
		return x.Conditions
	}
	return nil
}

func (x *Realms) GetRealms() []*Realm {
	if x != nil {
		return x.Realms
	}
	return nil
}

// Permission is a symbol that has form "<service>.<subject>.<verb>", which
// describes some elementary action ("<verb>") that can be done to some category
// of resources ("<subject>"), managed by some particular kind of LUCI service
// ("<service>").
//
// Within each individual realm (see Realm message), a principal (such as an end
// user or a service account) can have zero or more permissions that describe
// what this user can actually do to resources belonging to the realm. See Realm
// message for the definition of what "belonging to the realm" means.
//
// Examples of permissions:
//   * buildbucket.build.create
//   * swarming.pool.listBots
//   * swarming.task.cancel
//
// Note that permission names are composed of generic terms, not some specific
// IDs of service deployments or resources. Generally, using a concrete
// permission name in the service's source code as a constant should look
// natural.
//
// A permission can be marked as "internal". Internal permissions are not
// allowed to appear in custom roles in user-defined project realms.cfg files.
// They can be used in internal realms (defined in realms.cfg in the LUCI Auth
// service config set, see comments for Realm message) and they are added to
// some predefined roles by the LUCI Auth service itself. They are used to setup
// ACLs for internal interactions between LUCI components.
//
// Each individual LUCI service should document what permissions it checks and
// when. It becomes a part of service's public API. Usually services should
// check only permissions of resources they own (e.g. "<service>.<subject>.*"),
// but in exceptional cases they may also check permissions intended for other
// services. This is primarily useful for services that somehow "proxy" access
// to resources.
//
// Field `permissions` in Realms message describes all permissions known to the
// LUCI Auth service. The LUCI Auth service guarantees that all permissions
// mentioned in all realms (in `realms` field) are among `permissions` set.
//
// If a LUCI service checks a permission that is no longer (or not yet) listed
// in the `permissions` set, the check should succeed with "no permission"
// result, and produce a warning in service's logs.
type Permission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`          // "<service>.<subject>.<verb>"
	Internal bool   `protobuf:"varint,2,opt,name=internal,proto3" json:"internal,omitempty"` // internal permissions cannot be used in project realms
}

func (x *Permission) Reset() {
	*x = Permission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_realms_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Permission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Permission) ProtoMessage() {}

func (x *Permission) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_realms_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Permission.ProtoReflect.Descriptor instead.
func (*Permission) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_realms_proto_rawDescGZIP(), []int{1}
}

func (x *Permission) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Permission) GetInternal() bool {
	if x != nil {
		return x.Internal
	}
	return false
}

// Condition defines a predicate that takes a set of `attribute = value` pairs
// with a context of a particular permission check and returns True if the
// binding guarded by this condition should be respected.
//
// When a service checks a permission, it should pass to the authorization
// library a string-valued dictionary of attributes that describe the context
// of the permission check. It may contain things like the name of the resource
// being accessed, or parameters of the incoming RPC request that triggered
// the check.
//
// A list of available attributes and meaning of their values depends on
// the permission being checked and the service should document it in its API
// documentation.
type Condition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Op:
	//	*Condition_Restrict
	Op isCondition_Op `protobuf_oneof:"op"`
}

func (x *Condition) Reset() {
	*x = Condition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_realms_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Condition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Condition) ProtoMessage() {}

func (x *Condition) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_realms_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Condition.ProtoReflect.Descriptor instead.
func (*Condition) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_realms_proto_rawDescGZIP(), []int{2}
}

func (m *Condition) GetOp() isCondition_Op {
	if m != nil {
		return m.Op
	}
	return nil
}

func (x *Condition) GetRestrict() *Condition_AttributeRestriction {
	if x, ok := x.GetOp().(*Condition_Restrict); ok {
		return x.Restrict
	}
	return nil
}

type isCondition_Op interface {
	isCondition_Op()
}

type Condition_Restrict struct {
	Restrict *Condition_AttributeRestriction `protobuf:"bytes,1,opt,name=restrict,proto3,oneof"`
}

func (*Condition_Restrict) isCondition_Op() {}

// Realm is a named collection of (<principal>, <permission>) pairs.
//
// Realms are primarily defined in realms.cfg project config files. Such realms
// are called project realms. They are controlled by respective **project**
// owners and used to define ACLs for resources owned by these projects.
//
// There's a special set of realms (called internal realms or, sometimes, global
// realms) that are defined in realms.cfg in the LUCI Auth service config set.
// They are controlled by LUCI **deployment** owners and used to define ACLs for
// resources that are associated with LUCI deployment or LUCI services (and do
// not belong to any particular LUCI project). They are also allowed to use
// internal roles and permissions to define administrative-level ACLs (i.e. ACLs
// that transcend project boundaries).
//
// A full realm name has form "<project>:<realm>", where:
//   * "<project>" is a name of the LUCI project that defined the realm or
//     literal "@internal" for internal realms.
//   * "<realm>" is a name of the realm from a realms.cfg config file. This name
//     is also known as a project-scoped name, since it makes sense only within
//     a scope of some concrete LUCI project.
//
// A LUCI resource can point to exactly one realm by referring to its full
// "<project>:<realm>" name. Such reference can either be calculated on the fly
// from other resource's properties, or be stored alongside the resource's data.
// We say that such resource "belongs to the realm" or "lives in the realm" or
// is just "in the realm". We also say that such resource belongs to the project
// "<project>". The corresponding Realm message then describes who can do what
// to the resource.
//
// The logic of how resources get assigned to realms is a part of the public API
// of the service that owns resources. Some services may use a static realm
// assignment via project configuration files, others may do it dynamically by
// accepting a realm when a resource is created via an RPC.
//
// There are two special realms (both optional) that a project can have:
// "<project>:@root" and "<project>:@legacy".
//
// The root realm should be used as a fallback when an existing resource points
// to a realm that doesn't exist. Without the root realm, such resources become
// effectively inaccessible and this may be undesirable. The root realm usually
// contains only administrative-level bindings.
//
// The legacy realm should be used for legacy resources created before the
// realms mechanism was introduced in case the service can't figure out a more
// appropriate realm based on resource's properties. The service must clearly
// document when and how it uses the legacy realm (if it uses it at all).
//
// The actual list of (<principal>, <permission>) pairs is defined via a list of
// bindings, where each binding basically says "all these principals have all
// these permissions". In other words, each binding defines some subset of
// permissions and the overall realm permissions is a union of all such subsets.
// Subsets defined by bindings may potentially intersect or be empty.
//
// The LUCI Auth service constructs bindings by interpreting realms.cfg files
// using some set of rules. Individual LUCI services **must not care** about
// what these rules really are. They should use only the end result (in the form
// of bindings) provided in the Realm message. This allows to decouple the
// high-level user-facing language for defining permissions from the
// implementation of each individual LUCI service that checks permissions.
//
// A realm can also carry some small amount of data (usually auth related) that
// LUCI services use when dealing with this realm. It should be something that
// all (or at least more than one) LUCI services use. Configuration specific to
// a single service should be in this service's project config instead.
type Realm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the realm as "<project>:<realm>" string, where:
	//   "<project>" matches `^([a-z0-9\-_]{1,100}|@internal)$`.
	//   "<realm>" matches `^([a-z0-9_\.\-/]{1,400}|@root|@legacy)$`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A list of bindings in lexicographical order of their `permissions` fields.
	Bindings []*Binding `protobuf:"bytes,2,rep,name=bindings,proto3" json:"bindings,omitempty"`
	// Associated data extracted from the realms.cfg project config.
	Data *RealmData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Realm) Reset() {
	*x = Realm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_realms_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Realm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Realm) ProtoMessage() {}

func (x *Realm) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_realms_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Realm.ProtoReflect.Descriptor instead.
func (*Realm) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_realms_proto_rawDescGZIP(), []int{3}
}

func (x *Realm) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Realm) GetBindings() []*Binding {
	if x != nil {
		return x.Bindings
	}
	return nil
}

func (x *Realm) GetData() *RealmData {
	if x != nil {
		return x.Data
	}
	return nil
}

// Binding assigns all specified permissions to all specified principals.
type Binding struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Permissions in increasing order of their indexes.
	//
	// This set is a subset of `permissions` in the Realms message. Each element
	// is an index of a permission in the `permissions` list in the Realms
	// message.
	//
	// These indexes are not stable across different copies of Realms message.
	// They must not be stored or processed in isolation from the containing
	// Realms message.
	Permissions []uint32 `protobuf:"varint,1,rep,packed,name=permissions,proto3" json:"permissions,omitempty"`
	// A set of principals to grant all above permissions to.
	//
	// Each entry can either be an identity string (like "user:<email>") or a
	// LUCI group reference "group:<name>".
	//
	// Ordered alphabetically.
	Principals []string `protobuf:"bytes,2,rep,name=principals,proto3" json:"principals,omitempty"`
	// Conditions in increasing order of their indexes.
	//
	// Each element is an index of a condition in the `conditions` list in the
	// Realms message. These conditions are ANDed together. See the public API for
	// details of the semantics.
	//
	// These indexes are not stable across different copies of Realms message.
	// They must not be stored or processed in isolation from the containing
	// Realms message.
	Conditions []uint32 `protobuf:"varint,3,rep,packed,name=conditions,proto3" json:"conditions,omitempty"`
}

func (x *Binding) Reset() {
	*x = Binding{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_realms_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Binding) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Binding) ProtoMessage() {}

func (x *Binding) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_realms_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Binding.ProtoReflect.Descriptor instead.
func (*Binding) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_realms_proto_rawDescGZIP(), []int{4}
}

func (x *Binding) GetPermissions() []uint32 {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *Binding) GetPrincipals() []string {
	if x != nil {
		return x.Principals
	}
	return nil
}

func (x *Binding) GetConditions() []uint32 {
	if x != nil {
		return x.Conditions
	}
	return nil
}

// RealmData is semi-arbitrary non-ACL data extracted from the realms.cfg
// project config and attached to a realm.
type RealmData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Used only during Realms migration to gradually roll out the enforcement.
	EnforceInService []string `protobuf:"bytes,1,rep,name=enforce_in_service,json=enforceInService,proto3" json:"enforce_in_service,omitempty"`
}

func (x *RealmData) Reset() {
	*x = RealmData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_realms_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RealmData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RealmData) ProtoMessage() {}

func (x *RealmData) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_realms_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RealmData.ProtoReflect.Descriptor instead.
func (*RealmData) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_realms_proto_rawDescGZIP(), []int{5}
}

func (x *RealmData) GetEnforceInService() []string {
	if x != nil {
		return x.EnforceInService
	}
	return nil
}

// AttributeRestriction checks that the attributes set (as a set of key-value
// pairs) contains a particular (attribute, value) pair. Checked values are
// given as a list. The check succeeds if ANY of (attribute, value) pairs are
// present: `any((attribute, value) in attrs for value in values)`.
type Condition_AttributeRestriction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Attribute string   `protobuf:"bytes,1,opt,name=attribute,proto3" json:"attribute,omitempty"`
	Values    []string `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"` // sorted alphabetically, no dups
}

func (x *Condition_AttributeRestriction) Reset() {
	*x = Condition_AttributeRestriction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_realms_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Condition_AttributeRestriction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Condition_AttributeRestriction) ProtoMessage() {}

func (x *Condition_AttributeRestriction) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_realms_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Condition_AttributeRestriction.ProtoReflect.Descriptor instead.
func (*Condition_AttributeRestriction) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_realms_proto_rawDescGZIP(), []int{2, 0}
}

func (x *Condition_AttributeRestriction) GetAttribute() string {
	if x != nil {
		return x.Attribute
	}
	return ""
}

func (x *Condition_AttributeRestriction) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

var File_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_realms_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_realms_proto_rawDesc = []byte{
	0x0a, 0x54, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x73, 0x22, 0xe9,
	0x01, 0x0a, 0x06, 0x52, 0x65, 0x61, 0x6c, 0x6d, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x70, 0x69,
	0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x44, 0x0a, 0x0b, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x73, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x41, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x73, 0x2e, 0x43, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x35, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x73, 0x2e, 0x52, 0x65, 0x61,
	0x6c, 0x6d, 0x52, 0x06, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x73, 0x22, 0x3c, 0x0a, 0x0a, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x22, 0xb5, 0x01, 0x0a, 0x09, 0x43, 0x6f, 0x6e,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x54, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69,
	0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x72, 0x65, 0x61, 0x6c, 0x6d,
	0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x48, 0x00, 0x52, 0x08, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x1a, 0x4c, 0x0a, 0x14,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x42, 0x04, 0x0a, 0x02, 0x6f, 0x70,
	0x22, 0x8f, 0x01, 0x0a, 0x05, 0x52, 0x65, 0x61, 0x6c, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3b,
	0x0a, 0x08, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x73, 0x2e, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x52, 0x08, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x35, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x72, 0x65, 0x61, 0x6c,
	0x6d, 0x73, 0x2e, 0x52, 0x65, 0x61, 0x6c, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x6b, 0x0a, 0x07, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x20, 0x0a,
	0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0d, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x73, 0x12,
	0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0d, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22,
	0x39, 0x0a, 0x09, 0x52, 0x65, 0x61, 0x6c, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2c, 0x0a, 0x12,
	0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63,
	0x65, 0x49, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x6f,
	0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75,
	0x63, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_realms_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_realms_proto_rawDescData = file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_realms_proto_rawDesc
)

func file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_realms_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_realms_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_realms_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_realms_proto_rawDescData)
	})
	return file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_realms_proto_rawDescData
}

var file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_realms_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_realms_proto_goTypes = []interface{}{
	(*Realms)(nil),                         // 0: components.auth.realms.Realms
	(*Permission)(nil),                     // 1: components.auth.realms.Permission
	(*Condition)(nil),                      // 2: components.auth.realms.Condition
	(*Realm)(nil),                          // 3: components.auth.realms.Realm
	(*Binding)(nil),                        // 4: components.auth.realms.Binding
	(*RealmData)(nil),                      // 5: components.auth.realms.RealmData
	(*Condition_AttributeRestriction)(nil), // 6: components.auth.realms.Condition.AttributeRestriction
}
var file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_realms_proto_depIdxs = []int32{
	1, // 0: components.auth.realms.Realms.permissions:type_name -> components.auth.realms.Permission
	2, // 1: components.auth.realms.Realms.conditions:type_name -> components.auth.realms.Condition
	3, // 2: components.auth.realms.Realms.realms:type_name -> components.auth.realms.Realm
	6, // 3: components.auth.realms.Condition.restrict:type_name -> components.auth.realms.Condition.AttributeRestriction
	4, // 4: components.auth.realms.Realm.bindings:type_name -> components.auth.realms.Binding
	5, // 5: components.auth.realms.Realm.data:type_name -> components.auth.realms.RealmData
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() {
	file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_realms_proto_init()
}
func file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_realms_proto_init() {
	if File_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_realms_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_realms_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Realms); i {
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
		file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_realms_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Permission); i {
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
		file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_realms_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Condition); i {
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
		file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_realms_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Realm); i {
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
		file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_realms_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Binding); i {
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
		file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_realms_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RealmData); i {
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
		file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_realms_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Condition_AttributeRestriction); i {
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
	file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_realms_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Condition_Restrict)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_realms_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_realms_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_realms_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_realms_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_realms_proto = out.File
	file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_realms_proto_rawDesc = nil
	file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_realms_proto_goTypes = nil
	file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_realms_proto_depIdxs = nil
}
