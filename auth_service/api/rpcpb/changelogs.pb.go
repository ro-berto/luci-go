// Copyright 2022 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: go.chromium.org/luci/auth_service/api/rpcpb/changelogs.proto

package rpcpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// ListChangeLogsRequest is a request to get a list of change logs, which can
// be filtered by auth_db_rev and/or target.
type ListChangeLogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// AuthDB revision that the change log was made.
	AuthDbRev int64 `protobuf:"varint,1,opt,name=auth_db_rev,json=authDbRev,proto3" json:"auth_db_rev,omitempty"`
	// Entity that was changed in the change log.
	Target string `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	// The value of next_page_token received in a ListChangeLogsResponse. Used
	// to get the next page of change logs. If empty, gets the first page.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// The maximum number of change logs to include in the response.
	PageSize int32 `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *ListChangeLogsRequest) Reset() {
	*x = ListChangeLogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_auth_service_api_rpcpb_changelogs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListChangeLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListChangeLogsRequest) ProtoMessage() {}

func (x *ListChangeLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_auth_service_api_rpcpb_changelogs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListChangeLogsRequest.ProtoReflect.Descriptor instead.
func (*ListChangeLogsRequest) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_auth_service_api_rpcpb_changelogs_proto_rawDescGZIP(), []int{0}
}

func (x *ListChangeLogsRequest) GetAuthDbRev() int64 {
	if x != nil {
		return x.AuthDbRev
	}
	return 0
}

func (x *ListChangeLogsRequest) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *ListChangeLogsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListChangeLogsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

// ListChangeLogsResponse contains a list of change logs that matched the query.
type ListChangeLogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of change logs.
	Changes []*AuthDBChange `protobuf:"bytes,1,rep,name=changes,proto3" json:"changes,omitempty"`
	// The value to use as the page_token in a ListChangeLogsRequest to get the
	// next page of change logs. If empty, there are no more change logs.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListChangeLogsResponse) Reset() {
	*x = ListChangeLogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_auth_service_api_rpcpb_changelogs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListChangeLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListChangeLogsResponse) ProtoMessage() {}

func (x *ListChangeLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_auth_service_api_rpcpb_changelogs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListChangeLogsResponse.ProtoReflect.Descriptor instead.
func (*ListChangeLogsResponse) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_auth_service_api_rpcpb_changelogs_proto_rawDescGZIP(), []int{1}
}

func (x *ListChangeLogsResponse) GetChanges() []*AuthDBChange {
	if x != nil {
		return x.Changes
	}
	return nil
}

func (x *ListChangeLogsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

// AuthDBChange refers to a change log entry.
type AuthDBChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Fields common across all change types.
	ChangeType     string                 `protobuf:"bytes,1,opt,name=change_type,json=changeType,proto3" json:"change_type,omitempty"`
	Target         string                 `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	AuthDbRev      int64                  `protobuf:"varint,3,opt,name=auth_db_rev,json=authDbRev,proto3" json:"auth_db_rev,omitempty"`
	Who            string                 `protobuf:"bytes,4,opt,name=who,proto3" json:"who,omitempty"`
	When           *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=when,proto3" json:"when,omitempty"`
	Comment        string                 `protobuf:"bytes,6,opt,name=comment,proto3" json:"comment,omitempty"`
	AppVersion     string                 `protobuf:"bytes,7,opt,name=app_version,json=appVersion,proto3" json:"app_version,omitempty"`
	Description    string                 `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	OldDescription string                 `protobuf:"bytes,9,opt,name=old_description,json=oldDescription,proto3" json:"old_description,omitempty"`
	// Fields specific to AuthDBGroupChange.
	Owners    string   `protobuf:"bytes,10,opt,name=owners,proto3" json:"owners,omitempty"`
	OldOwners string   `protobuf:"bytes,11,opt,name=old_owners,json=oldOwners,proto3" json:"old_owners,omitempty"`
	Members   []string `protobuf:"bytes,12,rep,name=members,proto3" json:"members,omitempty"`
	Globs     []string `protobuf:"bytes,13,rep,name=globs,proto3" json:"globs,omitempty"`
	Nested    []string `protobuf:"bytes,14,rep,name=nested,proto3" json:"nested,omitempty"`
	// Fields specific to AuthDBIPAllowlistChange.
	Subnets []string `protobuf:"bytes,15,rep,name=subnets,proto3" json:"subnets,omitempty"`
	// Fields specific to AuthDBIPAllowlistAssignmentChange.
	Identity    string `protobuf:"bytes,16,opt,name=identity,proto3" json:"identity,omitempty"`
	IpAllowList string `protobuf:"bytes,17,opt,name=ip_allow_list,json=ipAllowList,proto3" json:"ip_allow_list,omitempty"`
	// Fields specific to AuthDBConfigChange.
	OauthClientId            string   `protobuf:"bytes,18,opt,name=oauth_client_id,json=oauthClientId,proto3" json:"oauth_client_id,omitempty"`
	OauthClientSecret        string   `protobuf:"bytes,19,opt,name=oauth_client_secret,json=oauthClientSecret,proto3" json:"oauth_client_secret,omitempty"`
	OauthAdditionalClientIds []string `protobuf:"bytes,20,rep,name=oauth_additional_client_ids,json=oauthAdditionalClientIds,proto3" json:"oauth_additional_client_ids,omitempty"`
	TokenServerUrlOld        string   `protobuf:"bytes,21,opt,name=token_server_url_old,json=tokenServerUrlOld,proto3" json:"token_server_url_old,omitempty"`
	TokenServerUrlNew        string   `protobuf:"bytes,22,opt,name=token_server_url_new,json=tokenServerUrlNew,proto3" json:"token_server_url_new,omitempty"`
	SecurityConfigOld        string   `protobuf:"bytes,23,opt,name=security_config_old,json=securityConfigOld,proto3" json:"security_config_old,omitempty"`
	SecurityConfigNew        string   `protobuf:"bytes,24,opt,name=security_config_new,json=securityConfigNew,proto3" json:"security_config_new,omitempty"`
	// Fields specific to AuthRealmsGlobalsChange.
	PermissionsAdded   []string `protobuf:"bytes,25,rep,name=permissions_added,json=permissionsAdded,proto3" json:"permissions_added,omitempty"`
	PermissionsChanged []string `protobuf:"bytes,26,rep,name=permissions_changed,json=permissionsChanged,proto3" json:"permissions_changed,omitempty"`
	PermissionsRemoved []string `protobuf:"bytes,27,rep,name=permissions_removed,json=permissionsRemoved,proto3" json:"permissions_removed,omitempty"`
	// Fields specific to AuthProjectRealmsChange.
	ConfigRevOld string `protobuf:"bytes,28,opt,name=config_rev_old,json=configRevOld,proto3" json:"config_rev_old,omitempty"`
	ConfigRevNew string `protobuf:"bytes,29,opt,name=config_rev_new,json=configRevNew,proto3" json:"config_rev_new,omitempty"`
	PermsRevOld  string `protobuf:"bytes,30,opt,name=perms_rev_old,json=permsRevOld,proto3" json:"perms_rev_old,omitempty"`
	PermsRevNew  string `protobuf:"bytes,31,opt,name=perms_rev_new,json=permsRevNew,proto3" json:"perms_rev_new,omitempty"`
}

func (x *AuthDBChange) Reset() {
	*x = AuthDBChange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_auth_service_api_rpcpb_changelogs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthDBChange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthDBChange) ProtoMessage() {}

func (x *AuthDBChange) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_auth_service_api_rpcpb_changelogs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthDBChange.ProtoReflect.Descriptor instead.
func (*AuthDBChange) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_auth_service_api_rpcpb_changelogs_proto_rawDescGZIP(), []int{2}
}

func (x *AuthDBChange) GetChangeType() string {
	if x != nil {
		return x.ChangeType
	}
	return ""
}

func (x *AuthDBChange) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *AuthDBChange) GetAuthDbRev() int64 {
	if x != nil {
		return x.AuthDbRev
	}
	return 0
}

func (x *AuthDBChange) GetWho() string {
	if x != nil {
		return x.Who
	}
	return ""
}

func (x *AuthDBChange) GetWhen() *timestamppb.Timestamp {
	if x != nil {
		return x.When
	}
	return nil
}

func (x *AuthDBChange) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *AuthDBChange) GetAppVersion() string {
	if x != nil {
		return x.AppVersion
	}
	return ""
}

func (x *AuthDBChange) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AuthDBChange) GetOldDescription() string {
	if x != nil {
		return x.OldDescription
	}
	return ""
}

func (x *AuthDBChange) GetOwners() string {
	if x != nil {
		return x.Owners
	}
	return ""
}

func (x *AuthDBChange) GetOldOwners() string {
	if x != nil {
		return x.OldOwners
	}
	return ""
}

func (x *AuthDBChange) GetMembers() []string {
	if x != nil {
		return x.Members
	}
	return nil
}

func (x *AuthDBChange) GetGlobs() []string {
	if x != nil {
		return x.Globs
	}
	return nil
}

func (x *AuthDBChange) GetNested() []string {
	if x != nil {
		return x.Nested
	}
	return nil
}

func (x *AuthDBChange) GetSubnets() []string {
	if x != nil {
		return x.Subnets
	}
	return nil
}

func (x *AuthDBChange) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *AuthDBChange) GetIpAllowList() string {
	if x != nil {
		return x.IpAllowList
	}
	return ""
}

func (x *AuthDBChange) GetOauthClientId() string {
	if x != nil {
		return x.OauthClientId
	}
	return ""
}

func (x *AuthDBChange) GetOauthClientSecret() string {
	if x != nil {
		return x.OauthClientSecret
	}
	return ""
}

func (x *AuthDBChange) GetOauthAdditionalClientIds() []string {
	if x != nil {
		return x.OauthAdditionalClientIds
	}
	return nil
}

func (x *AuthDBChange) GetTokenServerUrlOld() string {
	if x != nil {
		return x.TokenServerUrlOld
	}
	return ""
}

func (x *AuthDBChange) GetTokenServerUrlNew() string {
	if x != nil {
		return x.TokenServerUrlNew
	}
	return ""
}

func (x *AuthDBChange) GetSecurityConfigOld() string {
	if x != nil {
		return x.SecurityConfigOld
	}
	return ""
}

func (x *AuthDBChange) GetSecurityConfigNew() string {
	if x != nil {
		return x.SecurityConfigNew
	}
	return ""
}

func (x *AuthDBChange) GetPermissionsAdded() []string {
	if x != nil {
		return x.PermissionsAdded
	}
	return nil
}

func (x *AuthDBChange) GetPermissionsChanged() []string {
	if x != nil {
		return x.PermissionsChanged
	}
	return nil
}

func (x *AuthDBChange) GetPermissionsRemoved() []string {
	if x != nil {
		return x.PermissionsRemoved
	}
	return nil
}

func (x *AuthDBChange) GetConfigRevOld() string {
	if x != nil {
		return x.ConfigRevOld
	}
	return ""
}

func (x *AuthDBChange) GetConfigRevNew() string {
	if x != nil {
		return x.ConfigRevNew
	}
	return ""
}

func (x *AuthDBChange) GetPermsRevOld() string {
	if x != nil {
		return x.PermsRevOld
	}
	return ""
}

func (x *AuthDBChange) GetPermsRevNew() string {
	if x != nil {
		return x.PermsRevNew
	}
	return ""
}

var File_go_chromium_org_luci_auth_service_api_rpcpb_changelogs_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_auth_service_api_rpcpb_changelogs_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x70, 0x63, 0x70, 0x62, 0x2f, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x01,
	0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4c, 0x6f, 0x67, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x5f,
	0x64, 0x62, 0x5f, 0x72, 0x65, 0x76, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x75,
	0x74, 0x68, 0x44, 0x62, 0x52, 0x65, 0x76, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x76, 0x0a, 0x16, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x44, 0x42, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e,
	0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0x84, 0x09, 0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x44, 0x42, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x1e, 0x0a,
	0x0b, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x64, 0x62, 0x5f, 0x72, 0x65, 0x76, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x61, 0x75, 0x74, 0x68, 0x44, 0x62, 0x52, 0x65, 0x76, 0x12, 0x10, 0x0a,
	0x03, 0x77, 0x68, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x77, 0x68, 0x6f, 0x12,
	0x2e, 0x0a, 0x04, 0x77, 0x68, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x77, 0x68, 0x65, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x70, 0x70,
	0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x61, 0x70, 0x70, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x0f,
	0x6f, 0x6c, 0x64, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x6c, 0x64, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x1d, 0x0a,
	0x0a, 0x6f, 0x6c, 0x64, 0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6f, 0x6c, 0x64, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x6c, 0x6f, 0x62, 0x73, 0x18,
	0x0d, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x67, 0x6c, 0x6f, 0x62, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x65,
	0x73, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x73, 0x18,
	0x0f, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x22, 0x0a, 0x0d, 0x69, 0x70,
	0x5f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x69, 0x70, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x26,
	0x0a, 0x0f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x13, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x5f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x13, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x11, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x3d, 0x0a, 0x1b, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x5f,
	0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x14, 0x20, 0x03, 0x28, 0x09, 0x52, 0x18, 0x6f, 0x61, 0x75,
	0x74, 0x68, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x73, 0x12, 0x2f, 0x0a, 0x14, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x5f, 0x6f, 0x6c, 0x64, 0x18, 0x15, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x11, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x55, 0x72, 0x6c, 0x4f, 0x6c, 0x64, 0x12, 0x2f, 0x0a, 0x14, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x5f, 0x6e, 0x65, 0x77, 0x18, 0x16,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x55, 0x72, 0x6c, 0x4e, 0x65, 0x77, 0x12, 0x2e, 0x0a, 0x13, 0x73, 0x65, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x6f, 0x6c, 0x64, 0x18, 0x17,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x4f, 0x6c, 0x64, 0x12, 0x2e, 0x0a, 0x13, 0x73, 0x65, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x6e, 0x65, 0x77, 0x18, 0x18,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x4e, 0x65, 0x77, 0x12, 0x2b, 0x0a, 0x11, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x61, 0x64, 0x64, 0x65, 0x64, 0x18, 0x19, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x10, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x41,
	0x64, 0x64, 0x65, 0x64, 0x12, 0x2f, 0x0a, 0x13, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x18, 0x1a, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x12, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x64, 0x12, 0x2f, 0x0a, 0x13, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x18, 0x1b, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x12, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x5f, 0x72, 0x65, 0x76, 0x5f, 0x6f, 0x6c, 0x64, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x76, 0x4f, 0x6c, 0x64, 0x12, 0x24, 0x0a, 0x0e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x72, 0x65, 0x76, 0x5f, 0x6e, 0x65, 0x77, 0x18, 0x1d,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x76, 0x4e,
	0x65, 0x77, 0x12, 0x22, 0x0a, 0x0d, 0x70, 0x65, 0x72, 0x6d, 0x73, 0x5f, 0x72, 0x65, 0x76, 0x5f,
	0x6f, 0x6c, 0x64, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x73,
	0x52, 0x65, 0x76, 0x4f, 0x6c, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x70, 0x65, 0x72, 0x6d, 0x73, 0x5f,
	0x72, 0x65, 0x76, 0x5f, 0x6e, 0x65, 0x77, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x65, 0x72, 0x6d, 0x73, 0x52, 0x65, 0x76, 0x4e, 0x65, 0x77, 0x32, 0x69, 0x0a, 0x0a, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x5b, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x23, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x24, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f,
	0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72,
	0x70, 0x63, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_auth_service_api_rpcpb_changelogs_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_auth_service_api_rpcpb_changelogs_proto_rawDescData = file_go_chromium_org_luci_auth_service_api_rpcpb_changelogs_proto_rawDesc
)

func file_go_chromium_org_luci_auth_service_api_rpcpb_changelogs_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_auth_service_api_rpcpb_changelogs_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_auth_service_api_rpcpb_changelogs_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_auth_service_api_rpcpb_changelogs_proto_rawDescData)
	})
	return file_go_chromium_org_luci_auth_service_api_rpcpb_changelogs_proto_rawDescData
}

var file_go_chromium_org_luci_auth_service_api_rpcpb_changelogs_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_go_chromium_org_luci_auth_service_api_rpcpb_changelogs_proto_goTypes = []interface{}{
	(*ListChangeLogsRequest)(nil),  // 0: auth.service.ListChangeLogsRequest
	(*ListChangeLogsResponse)(nil), // 1: auth.service.ListChangeLogsResponse
	(*AuthDBChange)(nil),           // 2: auth.service.AuthDBChange
	(*timestamppb.Timestamp)(nil),  // 3: google.protobuf.Timestamp
}
var file_go_chromium_org_luci_auth_service_api_rpcpb_changelogs_proto_depIdxs = []int32{
	2, // 0: auth.service.ListChangeLogsResponse.changes:type_name -> auth.service.AuthDBChange
	3, // 1: auth.service.AuthDBChange.when:type_name -> google.protobuf.Timestamp
	0, // 2: auth.service.ChangeLogs.ListChangeLogs:input_type -> auth.service.ListChangeLogsRequest
	1, // 3: auth.service.ChangeLogs.ListChangeLogs:output_type -> auth.service.ListChangeLogsResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_auth_service_api_rpcpb_changelogs_proto_init() }
func file_go_chromium_org_luci_auth_service_api_rpcpb_changelogs_proto_init() {
	if File_go_chromium_org_luci_auth_service_api_rpcpb_changelogs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_auth_service_api_rpcpb_changelogs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListChangeLogsRequest); i {
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
		file_go_chromium_org_luci_auth_service_api_rpcpb_changelogs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListChangeLogsResponse); i {
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
		file_go_chromium_org_luci_auth_service_api_rpcpb_changelogs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthDBChange); i {
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
			RawDescriptor: file_go_chromium_org_luci_auth_service_api_rpcpb_changelogs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_go_chromium_org_luci_auth_service_api_rpcpb_changelogs_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_auth_service_api_rpcpb_changelogs_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_auth_service_api_rpcpb_changelogs_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_auth_service_api_rpcpb_changelogs_proto = out.File
	file_go_chromium_org_luci_auth_service_api_rpcpb_changelogs_proto_rawDesc = nil
	file_go_chromium_org_luci_auth_service_api_rpcpb_changelogs_proto_goTypes = nil
	file_go_chromium_org_luci_auth_service_api_rpcpb_changelogs_proto_depIdxs = nil
}