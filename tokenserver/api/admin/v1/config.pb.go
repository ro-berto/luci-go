// Copyright 2016 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: go.chromium.org/luci/tokenserver/api/admin/v1/config.proto

package admin

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

// TokenServerConfig is read from tokenserver.cfg in luci-config.
type TokenServerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of CAs we trust.
	CertificateAuthority []*CertificateAuthorityConfig `protobuf:"bytes,1,rep,name=certificate_authority,json=certificateAuthority,proto3" json:"certificate_authority,omitempty"`
}

func (x *TokenServerConfig) Reset() {
	*x = TokenServerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_tokenserver_api_admin_v1_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenServerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenServerConfig) ProtoMessage() {}

func (x *TokenServerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_tokenserver_api_admin_v1_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenServerConfig.ProtoReflect.Descriptor instead.
func (*TokenServerConfig) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_tokenserver_api_admin_v1_config_proto_rawDescGZIP(), []int{0}
}

func (x *TokenServerConfig) GetCertificateAuthority() []*CertificateAuthorityConfig {
	if x != nil {
		return x.CertificateAuthority
	}
	return nil
}

// CertificateAuthorityConfig defines a single CA we trust.
//
// Such CA issues certificates for nodes that use The Token Service. Each node
// has a private key and certificate with Common Name set to the FQDN of this
// node, e.g. "CN=slave43-c1.c.chromecompute.google.com.internal".
//
// The Token Server uses this CN to derive an identity string for a machine. It
// splits FQDN into a hostname ("slave43-c1") and a domain name
// ("c.chromecompute.google.com.internal"), searches for a domain name in
// "known_domains" set, and, if it is present, uses parameters described there
// for generating a token that contains machine's FQDN and certificate serial
// number (among other things, see MachineTokenBody in machine_token.proto).
type CertificateAuthorityConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UniqueId    int64    `protobuf:"varint,6,opt,name=unique_id,json=uniqueId,proto3" json:"unique_id,omitempty"`         // ID of this CA, will be embedded into tokens.
	Cn          string   `protobuf:"bytes,1,opt,name=cn,proto3" json:"cn,omitempty"`                                      // CA Common Name, must match Subject CN in the cert
	CertPath    string   `protobuf:"bytes,2,opt,name=cert_path,json=certPath,proto3" json:"cert_path,omitempty"`          // path to the root certificate file in luci-config
	CrlUrl      string   `protobuf:"bytes,3,opt,name=crl_url,json=crlUrl,proto3" json:"crl_url,omitempty"`                // where to fetch CRL from
	UseOauth    bool     `protobuf:"varint,4,opt,name=use_oauth,json=useOauth,proto3" json:"use_oauth,omitempty"`         // true to send Authorization header when fetching CRL
	OauthScopes []string `protobuf:"bytes,7,rep,name=oauth_scopes,json=oauthScopes,proto3" json:"oauth_scopes,omitempty"` // OAuth scopes to use when fetching CRL
	// KnownDomains describes parameters to use for each particular domain.
	KnownDomains []*DomainConfig `protobuf:"bytes,5,rep,name=known_domains,json=knownDomains,proto3" json:"known_domains,omitempty"`
}

func (x *CertificateAuthorityConfig) Reset() {
	*x = CertificateAuthorityConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_tokenserver_api_admin_v1_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CertificateAuthorityConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CertificateAuthorityConfig) ProtoMessage() {}

func (x *CertificateAuthorityConfig) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_tokenserver_api_admin_v1_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CertificateAuthorityConfig.ProtoReflect.Descriptor instead.
func (*CertificateAuthorityConfig) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_tokenserver_api_admin_v1_config_proto_rawDescGZIP(), []int{1}
}

func (x *CertificateAuthorityConfig) GetUniqueId() int64 {
	if x != nil {
		return x.UniqueId
	}
	return 0
}

func (x *CertificateAuthorityConfig) GetCn() string {
	if x != nil {
		return x.Cn
	}
	return ""
}

func (x *CertificateAuthorityConfig) GetCertPath() string {
	if x != nil {
		return x.CertPath
	}
	return ""
}

func (x *CertificateAuthorityConfig) GetCrlUrl() string {
	if x != nil {
		return x.CrlUrl
	}
	return ""
}

func (x *CertificateAuthorityConfig) GetUseOauth() bool {
	if x != nil {
		return x.UseOauth
	}
	return false
}

func (x *CertificateAuthorityConfig) GetOauthScopes() []string {
	if x != nil {
		return x.OauthScopes
	}
	return nil
}

func (x *CertificateAuthorityConfig) GetKnownDomains() []*DomainConfig {
	if x != nil {
		return x.KnownDomains
	}
	return nil
}

// DomainConfig is used inside CertificateAuthorityConfig.
type DomainConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Domain is domain names of hosts this config applies to.
	//
	// Machines that reside in a subdomain of given domain are also considered
	// part of it, e.g. both FQDNs "host.example.com" and "host.abc.example.com"
	// match domain "example.com".
	Domain []string `protobuf:"bytes,1,rep,name=domain,proto3" json:"domain,omitempty"`
	// MachineTokenLifetime is how long generated machine tokens live, in seconds.
	//
	// If 0, machine tokens are not allowed.
	MachineTokenLifetime int64 `protobuf:"varint,5,opt,name=machine_token_lifetime,json=machineTokenLifetime,proto3" json:"machine_token_lifetime,omitempty"`
}

func (x *DomainConfig) Reset() {
	*x = DomainConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_tokenserver_api_admin_v1_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DomainConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DomainConfig) ProtoMessage() {}

func (x *DomainConfig) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_tokenserver_api_admin_v1_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DomainConfig.ProtoReflect.Descriptor instead.
func (*DomainConfig) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_tokenserver_api_admin_v1_config_proto_rawDescGZIP(), []int{2}
}

func (x *DomainConfig) GetDomain() []string {
	if x != nil {
		return x.Domain
	}
	return nil
}

func (x *DomainConfig) GetMachineTokenLifetime() int64 {
	if x != nil {
		return x.MachineTokenLifetime
	}
	return 0
}

// DelegationPermissions is read from delegation.cfg in luci-config.
type DelegationPermissions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Rules specify what calls to MintDelegationToken are allowed.
	//
	// Rules are evaluated independently. One and only one rule should match the
	// request to allow the operation. If none rules or more than one rule match,
	// the request will be denied.
	//
	// See DelegationRule comments for more details.
	Rules []*DelegationRule `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
}

func (x *DelegationPermissions) Reset() {
	*x = DelegationPermissions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_tokenserver_api_admin_v1_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelegationPermissions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelegationPermissions) ProtoMessage() {}

func (x *DelegationPermissions) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_tokenserver_api_admin_v1_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelegationPermissions.ProtoReflect.Descriptor instead.
func (*DelegationPermissions) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_tokenserver_api_admin_v1_config_proto_rawDescGZIP(), []int{3}
}

func (x *DelegationPermissions) GetRules() []*DelegationRule {
	if x != nil {
		return x.Rules
	}
	return nil
}

// DelegationRule describes a single allowed case of using delegation tokens.
//
// An incoming MintDelegationTokenRequest is basically a tuple of:
//   - 'requestor_id' - an identity of whoever makes the request.
//   - 'delegated_identity' - an identity to delegate.
//   - 'audience' - a set of identities that will be able to use the token.
//   - 'services' - a set of services that should accept the token.
//
// A request matches a rule iff:
//   - 'requestor_id' is in 'requestor' set.
//   - 'delegated_identity' is in 'allowed_to_impersonate' set.
//   - 'audience' is a subset of 'allowed_audience' set.
//   - 'services' is a subset of 'target_service' set.
//
// The presence of a matching rule permits to mint the token. The rule also
// provides an upper bound on allowed validity_duration, and the rule's name
// is logged in the audit trail.
type DelegationRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A descriptive name of this rule, for the audit log.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Email of developers that own this rule, to know who to contact.
	Owner []string `protobuf:"bytes,2,rep,name=owner,proto3" json:"owner,omitempty"`
	// A set of callers to which this rule applies.
	//
	// Matched against verified credentials of a caller of MintDelegationToken.
	//
	// Each element is either:
	//  * An identity string ("user:<email>").
	//  * A group reference ("group:<name>").
	//
	// The groups specified here are expanded when MintDelegationTokenRequest is
	// evaluated.
	Requestor []string `protobuf:"bytes,3,rep,name=requestor,proto3" json:"requestor,omitempty"`
	// Identities that are allowed to be delegated/impersonated by the requestor.
	//
	// Matched against 'delegated_identity' field of MintDelegationTokenRequest.
	//
	// Each element is either:
	//  * An identity string ("user:<email>").
	//  * A group reference ("group:<name>").
	//  * A special identifier "REQUESTOR" that is substituted by the requestor
	//    identity when evaluating the rule.
	//
	// "REQUESTOR" allows one to generate tokens that delegate their own identity
	// to some target audience.
	//
	// The groups specified here are expanded when MintDelegationTokenRequest is
	// evaluated.
	AllowedToImpersonate []string `protobuf:"bytes,4,rep,name=allowed_to_impersonate,json=allowedToImpersonate,proto3" json:"allowed_to_impersonate,omitempty"`
	// A set of identities that should be able to use the new token.
	//
	// Matched against 'audience' field of MintDelegationTokenRequest.
	//
	// Each element is either:
	//  * An identity string ("user:<email>").
	//  * A group reference ("group:<name>").
	//  * A special identifier "REQUESTOR" that is substituted by the requestor
	//    identity when evaluating the rule.
	//  * A special token "*" that means "any bearer can use the new token,
	//    including anonymous".
	//
	// "REQUESTOR" is typically used here for rules that allow requestors to
	// impersonate someone else. The corresponding tokens have the requestor as
	// the only allowed audience.
	//
	// The groups specified here are NOT expanded when MintDelegationTokenRequest
	// is evaluated. To match the rule, MintDelegationTokenRequest must specify
	// subset of 'allowed_audience' groups explicitly in 'audience' field.
	AllowedAudience []string `protobuf:"bytes,5,rep,name=allowed_audience,json=allowedAudience,proto3" json:"allowed_audience,omitempty"`
	// A set of services that should be able to accept the new token.
	//
	// Matched against 'services' field of MintDelegationTokenRequest.
	//
	// Each element is either:
	//  * A service identity string ("service:<id>").
	//  * A special token "*" that mean "any LUCI service should accept the
	//    token".
	TargetService []string `protobuf:"bytes,6,rep,name=target_service,json=targetService,proto3" json:"target_service,omitempty"`
	// Maximum allowed validity duration (sec) of minted delegation tokens.
	//
	// Default is 12 hours.
	MaxValidityDuration int64 `protobuf:"varint,7,opt,name=max_validity_duration,json=maxValidityDuration,proto3" json:"max_validity_duration,omitempty"`
}

func (x *DelegationRule) Reset() {
	*x = DelegationRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_tokenserver_api_admin_v1_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelegationRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelegationRule) ProtoMessage() {}

func (x *DelegationRule) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_tokenserver_api_admin_v1_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelegationRule.ProtoReflect.Descriptor instead.
func (*DelegationRule) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_tokenserver_api_admin_v1_config_proto_rawDescGZIP(), []int{4}
}

func (x *DelegationRule) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DelegationRule) GetOwner() []string {
	if x != nil {
		return x.Owner
	}
	return nil
}

func (x *DelegationRule) GetRequestor() []string {
	if x != nil {
		return x.Requestor
	}
	return nil
}

func (x *DelegationRule) GetAllowedToImpersonate() []string {
	if x != nil {
		return x.AllowedToImpersonate
	}
	return nil
}

func (x *DelegationRule) GetAllowedAudience() []string {
	if x != nil {
		return x.AllowedAudience
	}
	return nil
}

func (x *DelegationRule) GetTargetService() []string {
	if x != nil {
		return x.TargetService
	}
	return nil
}

func (x *DelegationRule) GetMaxValidityDuration() int64 {
	if x != nil {
		return x.MaxValidityDuration
	}
	return 0
}

// ServiceAccountsProjectMapping defines what service accounts belong to what
// LUCI projects.
//
// Used by MintServiceAccountToken RPC as a final authorization step, after
// checking that the usage of the service account is allowed by Realms ACLs.
//
// This is a stop gap solution until the Token Server learns to use
// project-scoped accounts when calling Cloud IAM. Once this happens, we can
// move information contained in ServiceAccountsProjectMapping into Cloud IAM
// permissions.
//
// This message is stored as project_owned_accounts.cfg in luci-config.
type ServiceAccountsProjectMapping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Each entry maps a bunch of service accounts to one or more projects.
	Mapping []*ServiceAccountsProjectMapping_Mapping `protobuf:"bytes,1,rep,name=mapping,proto3" json:"mapping,omitempty"`
}

func (x *ServiceAccountsProjectMapping) Reset() {
	*x = ServiceAccountsProjectMapping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_tokenserver_api_admin_v1_config_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceAccountsProjectMapping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceAccountsProjectMapping) ProtoMessage() {}

func (x *ServiceAccountsProjectMapping) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_tokenserver_api_admin_v1_config_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceAccountsProjectMapping.ProtoReflect.Descriptor instead.
func (*ServiceAccountsProjectMapping) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_tokenserver_api_admin_v1_config_proto_rawDescGZIP(), []int{5}
}

func (x *ServiceAccountsProjectMapping) GetMapping() []*ServiceAccountsProjectMapping_Mapping {
	if x != nil {
		return x.Mapping
	}
	return nil
}

type ServiceAccountsProjectMapping_Mapping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Names of LUCI projects.
	Project []string `protobuf:"bytes,1,rep,name=project,proto3" json:"project,omitempty"`
	// Emails of service accounts allowed to be used by all these projects.
	ServiceAccount []string `protobuf:"bytes,2,rep,name=service_account,json=serviceAccount,proto3" json:"service_account,omitempty"`
}

func (x *ServiceAccountsProjectMapping_Mapping) Reset() {
	*x = ServiceAccountsProjectMapping_Mapping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_tokenserver_api_admin_v1_config_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceAccountsProjectMapping_Mapping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceAccountsProjectMapping_Mapping) ProtoMessage() {}

func (x *ServiceAccountsProjectMapping_Mapping) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_tokenserver_api_admin_v1_config_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceAccountsProjectMapping_Mapping.ProtoReflect.Descriptor instead.
func (*ServiceAccountsProjectMapping_Mapping) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_tokenserver_api_admin_v1_config_proto_rawDescGZIP(), []int{5, 0}
}

func (x *ServiceAccountsProjectMapping_Mapping) GetProject() []string {
	if x != nil {
		return x.Project
	}
	return nil
}

func (x *ServiceAccountsProjectMapping_Mapping) GetServiceAccount() []string {
	if x != nil {
		return x.ServiceAccount
	}
	return nil
}

var File_go_chromium_org_luci_tokenserver_api_admin_v1_config_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_tokenserver_api_admin_v1_config_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x22,
	0x77, 0x0a, 0x11, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x62, 0x0a, 0x15, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x14, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x22, 0x85, 0x02, 0x0a, 0x1a, 0x43, 0x65, 0x72,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x6e, 0x69, 0x71, 0x75,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x75, 0x6e, 0x69, 0x71,
	0x75, 0x65, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x63, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x63, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x65, 0x72, 0x74, 0x5f, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x65, 0x72, 0x74, 0x50, 0x61, 0x74,
	0x68, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x72, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x63, 0x72, 0x6c, 0x55, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73,
	0x65, 0x5f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x4f, 0x61, 0x75, 0x74, 0x68, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x61, 0x75, 0x74, 0x68,
	0x5f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x6f,
	0x61, 0x75, 0x74, 0x68, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x12, 0x44, 0x0a, 0x0d, 0x6b, 0x6e,
	0x6f, 0x77, 0x6e, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x0c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73,
	0x22, 0x74, 0x0a, 0x0c, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x34, 0x0a, 0x16, 0x6d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x6c, 0x69, 0x66, 0x65, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x14, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x4a, 0x04,
	0x08, 0x02, 0x10, 0x03, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x04, 0x4a, 0x04, 0x08, 0x04, 0x10, 0x05,
	0x4a, 0x04, 0x08, 0x06, 0x10, 0x07, 0x22, 0x50, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x37, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x75, 0x6c,
	0x65, 0x52, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x22, 0x94, 0x02, 0x0a, 0x0e, 0x44, 0x65, 0x6c,
	0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x6f, 0x72, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x6f, 0x72, 0x12, 0x34, 0x0a, 0x16, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x74,
	0x6f, 0x5f, 0x69, 0x6d, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x14, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x54, 0x6f, 0x49, 0x6d,
	0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x61, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x64, 0x5f, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x41, 0x75, 0x64, 0x69,
	0x65, 0x6e, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x15, 0x6d,
	0x61, 0x78, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x69, 0x74, 0x79, 0x5f, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x13, 0x6d, 0x61, 0x78, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x69, 0x74, 0x79, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0xc1, 0x01, 0x0a, 0x1d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x12, 0x52, 0x0a, 0x07, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x38, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x61, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x2e, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x6d, 0x61,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x1a, 0x4c, 0x0a, 0x07, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69,
	0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_tokenserver_api_admin_v1_config_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_tokenserver_api_admin_v1_config_proto_rawDescData = file_go_chromium_org_luci_tokenserver_api_admin_v1_config_proto_rawDesc
)

func file_go_chromium_org_luci_tokenserver_api_admin_v1_config_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_tokenserver_api_admin_v1_config_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_tokenserver_api_admin_v1_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_tokenserver_api_admin_v1_config_proto_rawDescData)
	})
	return file_go_chromium_org_luci_tokenserver_api_admin_v1_config_proto_rawDescData
}

var file_go_chromium_org_luci_tokenserver_api_admin_v1_config_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_go_chromium_org_luci_tokenserver_api_admin_v1_config_proto_goTypes = []interface{}{
	(*TokenServerConfig)(nil),                     // 0: tokenserver.admin.TokenServerConfig
	(*CertificateAuthorityConfig)(nil),            // 1: tokenserver.admin.CertificateAuthorityConfig
	(*DomainConfig)(nil),                          // 2: tokenserver.admin.DomainConfig
	(*DelegationPermissions)(nil),                 // 3: tokenserver.admin.DelegationPermissions
	(*DelegationRule)(nil),                        // 4: tokenserver.admin.DelegationRule
	(*ServiceAccountsProjectMapping)(nil),         // 5: tokenserver.admin.ServiceAccountsProjectMapping
	(*ServiceAccountsProjectMapping_Mapping)(nil), // 6: tokenserver.admin.ServiceAccountsProjectMapping.Mapping
}
var file_go_chromium_org_luci_tokenserver_api_admin_v1_config_proto_depIdxs = []int32{
	1, // 0: tokenserver.admin.TokenServerConfig.certificate_authority:type_name -> tokenserver.admin.CertificateAuthorityConfig
	2, // 1: tokenserver.admin.CertificateAuthorityConfig.known_domains:type_name -> tokenserver.admin.DomainConfig
	4, // 2: tokenserver.admin.DelegationPermissions.rules:type_name -> tokenserver.admin.DelegationRule
	6, // 3: tokenserver.admin.ServiceAccountsProjectMapping.mapping:type_name -> tokenserver.admin.ServiceAccountsProjectMapping.Mapping
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_tokenserver_api_admin_v1_config_proto_init() }
func file_go_chromium_org_luci_tokenserver_api_admin_v1_config_proto_init() {
	if File_go_chromium_org_luci_tokenserver_api_admin_v1_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_tokenserver_api_admin_v1_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenServerConfig); i {
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
		file_go_chromium_org_luci_tokenserver_api_admin_v1_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CertificateAuthorityConfig); i {
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
		file_go_chromium_org_luci_tokenserver_api_admin_v1_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DomainConfig); i {
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
		file_go_chromium_org_luci_tokenserver_api_admin_v1_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelegationPermissions); i {
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
		file_go_chromium_org_luci_tokenserver_api_admin_v1_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelegationRule); i {
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
		file_go_chromium_org_luci_tokenserver_api_admin_v1_config_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceAccountsProjectMapping); i {
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
		file_go_chromium_org_luci_tokenserver_api_admin_v1_config_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceAccountsProjectMapping_Mapping); i {
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
			RawDescriptor: file_go_chromium_org_luci_tokenserver_api_admin_v1_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_tokenserver_api_admin_v1_config_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_tokenserver_api_admin_v1_config_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_tokenserver_api_admin_v1_config_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_tokenserver_api_admin_v1_config_proto = out.File
	file_go_chromium_org_luci_tokenserver_api_admin_v1_config_proto_rawDesc = nil
	file_go_chromium_org_luci_tokenserver_api_admin_v1_config_proto_goTypes = nil
	file_go_chromium_org_luci_tokenserver_api_admin_v1_config_proto_depIdxs = nil
}
