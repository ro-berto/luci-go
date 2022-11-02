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
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: go.chromium.org/luci/cv/settings/listener/settings.proto

package listenerpb

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type Settings_GerritSubscription_MessageFormat int32

const (
	Settings_GerritSubscription_MESSAGE_FORMAT_UNSPECIFIED Settings_GerritSubscription_MessageFormat = 0
	Settings_GerritSubscription_JSON                       Settings_GerritSubscription_MessageFormat = 1
	Settings_GerritSubscription_PROTO_BINARY               Settings_GerritSubscription_MessageFormat = 2
)

// Enum value maps for Settings_GerritSubscription_MessageFormat.
var (
	Settings_GerritSubscription_MessageFormat_name = map[int32]string{
		0: "MESSAGE_FORMAT_UNSPECIFIED",
		1: "JSON",
		2: "PROTO_BINARY",
	}
	Settings_GerritSubscription_MessageFormat_value = map[string]int32{
		"MESSAGE_FORMAT_UNSPECIFIED": 0,
		"JSON":                       1,
		"PROTO_BINARY":               2,
	}
)

func (x Settings_GerritSubscription_MessageFormat) Enum() *Settings_GerritSubscription_MessageFormat {
	p := new(Settings_GerritSubscription_MessageFormat)
	*p = x
	return p
}

func (x Settings_GerritSubscription_MessageFormat) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Settings_GerritSubscription_MessageFormat) Descriptor() protoreflect.EnumDescriptor {
	return file_go_chromium_org_luci_cv_settings_listener_settings_proto_enumTypes[0].Descriptor()
}

func (Settings_GerritSubscription_MessageFormat) Type() protoreflect.EnumType {
	return &file_go_chromium_org_luci_cv_settings_listener_settings_proto_enumTypes[0]
}

func (x Settings_GerritSubscription_MessageFormat) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Settings_GerritSubscription_MessageFormat.Descriptor instead.
func (Settings_GerritSubscription_MessageFormat) EnumDescriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_settings_listener_settings_proto_rawDescGZIP(), []int{0, 1, 0}
}

// Settings defines fields for configuring listener settings.
type Settings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Subscriptions for the Gerrit hosts that have enabled Gerrit Pub/Sub.
	//
	// To enable Gerrit pub/sub for a given LUCI project, the subscription of
	// all the Gerrit hosts listed in the project config must be added in this
	// field. If not, the config validation will fail.
	GerritSubscriptions []*Settings_GerritSubscription `protobuf:"bytes,1,rep,name=gerrit_subscriptions,json=gerritSubscriptions,proto3" json:"gerrit_subscriptions,omitempty"`
	// If a LUCI Project matches any of the regexps, CV will use the pubsub
	// listener to find changes in the Gerrit hosts listed in the project config.
	//
	// If not, CV will use the incremental poller to find changes in the Gerrit
	// hosts.
	//
	// Deprecated: Do not use.
	EnabledProjectRegexps []string `protobuf:"bytes,2,rep,name=enabled_project_regexps,json=enabledProjectRegexps,proto3" json:"enabled_project_regexps,omitempty"`
	// If a LUCI Project matches any of the regexps, CV will not use the pubsub
	// listener to find changes in the Gerrit hosts listed in the project config.
	//
	// Instead, CV will use the incremental poller to find changes from the Gerrit
	// hosts.
	DisabledProjectRegexps []string `protobuf:"bytes,3,rep,name=disabled_project_regexps,json=disabledProjectRegexps,proto3" json:"disabled_project_regexps,omitempty"`
}

func (x *Settings) Reset() {
	*x = Settings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_settings_listener_settings_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Settings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Settings) ProtoMessage() {}

func (x *Settings) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_settings_listener_settings_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Settings.ProtoReflect.Descriptor instead.
func (*Settings) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_settings_listener_settings_proto_rawDescGZIP(), []int{0}
}

func (x *Settings) GetGerritSubscriptions() []*Settings_GerritSubscription {
	if x != nil {
		return x.GerritSubscriptions
	}
	return nil
}

// Deprecated: Do not use.
func (x *Settings) GetEnabledProjectRegexps() []string {
	if x != nil {
		return x.EnabledProjectRegexps
	}
	return nil
}

func (x *Settings) GetDisabledProjectRegexps() []string {
	if x != nil {
		return x.DisabledProjectRegexps
	}
	return nil
}

type Settings_ReceiveSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The number of goroutines that Listener will spawn for the subscription.
	//
	// 10, if unset.
	//
	// This doesn't limit the number of buffered messages that are waiting to
	// be processed or are being processed.
	//
	// Use max_outstanding_messages to limit he number of buffered messages.
	NumGoroutines uint64 `protobuf:"varint,1,opt,name=num_goroutines,json=numGoroutines,proto3" json:"num_goroutines,omitempty"`
	// The maximum number of unacknowledged but not yet expired messages.
	//
	// 1000, if unset.
	// If < 0, there will be no limit.
	MaxOutstandingMessages int64 `protobuf:"varint,2,opt,name=max_outstanding_messages,json=maxOutstandingMessages,proto3" json:"max_outstanding_messages,omitempty"`
}

func (x *Settings_ReceiveSettings) Reset() {
	*x = Settings_ReceiveSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_settings_listener_settings_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Settings_ReceiveSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Settings_ReceiveSettings) ProtoMessage() {}

func (x *Settings_ReceiveSettings) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_settings_listener_settings_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Settings_ReceiveSettings.ProtoReflect.Descriptor instead.
func (*Settings_ReceiveSettings) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_settings_listener_settings_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Settings_ReceiveSettings) GetNumGoroutines() uint64 {
	if x != nil {
		return x.NumGoroutines
	}
	return 0
}

func (x *Settings_ReceiveSettings) GetMaxOutstandingMessages() int64 {
	if x != nil {
		return x.MaxOutstandingMessages
	}
	return 0
}

type Settings_GerritSubscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Gerrit host w/o scheme.
	// For example, chromium-review.googlesource.com
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	// The subscription ID of the host. If unset, `host` is the subscription ID.
	//
	// Note that this is subscription ID, not subscription name.
	// Subscription name is the full path of a subscription in the format of
	// projects/$project/subscription/$sub_id.
	SubscriptionId string `protobuf:"bytes,2,opt,name=subscription_id,json=subscriptionId,proto3" json:"subscription_id,omitempty"`
	// Configuration for the pubsub receive function.
	ReceiveSettings *Settings_ReceiveSettings `protobuf:"bytes,3,opt,name=receive_settings,json=receiveSettings,proto3" json:"receive_settings,omitempty"`
	// The format of the pubsub payload.
	MessageFormat Settings_GerritSubscription_MessageFormat `protobuf:"varint,4,opt,name=message_format,json=messageFormat,proto3,enum=listener.Settings_GerritSubscription_MessageFormat" json:"message_format,omitempty"`
}

func (x *Settings_GerritSubscription) Reset() {
	*x = Settings_GerritSubscription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_settings_listener_settings_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Settings_GerritSubscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Settings_GerritSubscription) ProtoMessage() {}

func (x *Settings_GerritSubscription) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_settings_listener_settings_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Settings_GerritSubscription.ProtoReflect.Descriptor instead.
func (*Settings_GerritSubscription) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_settings_listener_settings_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Settings_GerritSubscription) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *Settings_GerritSubscription) GetSubscriptionId() string {
	if x != nil {
		return x.SubscriptionId
	}
	return ""
}

func (x *Settings_GerritSubscription) GetReceiveSettings() *Settings_ReceiveSettings {
	if x != nil {
		return x.ReceiveSettings
	}
	return nil
}

func (x *Settings_GerritSubscription) GetMessageFormat() Settings_GerritSubscription_MessageFormat {
	if x != nil {
		return x.MessageFormat
	}
	return Settings_GerritSubscription_MESSAGE_FORMAT_UNSPECIFIED
}

var File_go_chromium_org_luci_cv_settings_listener_settings_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_cv_settings_listener_settings_proto_rawDesc = []byte{
	0x0a, 0x38, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x76, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6c, 0x69, 0x73, 0x74,
	0x65, 0x6e, 0x65, 0x72, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb1, 0x05,
	0x0a, 0x08, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x58, 0x0a, 0x14, 0x67, 0x65,
	0x72, 0x72, 0x69, 0x74, 0x5f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x65,
	0x6e, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x47, 0x65, 0x72,
	0x72, 0x69, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x13, 0x67, 0x65, 0x72, 0x72, 0x69, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3a, 0x0a, 0x17, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x5f,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x67, 0x65, 0x78, 0x70, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x15, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x67, 0x65, 0x78, 0x70, 0x73,
	0x12, 0x38, 0x0a, 0x18, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x5f, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x67, 0x65, 0x78, 0x70, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x16, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x67, 0x65, 0x78, 0x70, 0x73, 0x1a, 0x72, 0x0a, 0x0f, 0x52, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x25, 0x0a,
	0x0e, 0x6e, 0x75, 0x6d, 0x5f, 0x67, 0x6f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x6e, 0x75, 0x6d, 0x47, 0x6f, 0x72, 0x6f, 0x75, 0x74,
	0x69, 0x6e, 0x65, 0x73, 0x12, 0x38, 0x0a, 0x18, 0x6d, 0x61, 0x78, 0x5f, 0x6f, 0x75, 0x74, 0x73,
	0x74, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x16, 0x6d, 0x61, 0x78, 0x4f, 0x75, 0x74, 0x73, 0x74,
	0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x1a, 0xe0,
	0x02, 0x0a, 0x12, 0x47, 0x65, 0x72, 0x72, 0x69, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x0b, 0xfa, 0x42, 0x08, 0x72, 0x06, 0x10, 0x01, 0xba, 0x01, 0x01, 0x2f,
	0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x4d, 0x0a, 0x10, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6c, 0x69, 0x73, 0x74,
	0x65, 0x6e, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x52, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x0f, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x64,
	0x0a, 0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x33, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65,
	0x72, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x47, 0x65, 0x72, 0x72, 0x69,
	0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x42, 0x08, 0xfa, 0x42, 0x05,
	0x82, 0x01, 0x02, 0x20, 0x00, 0x52, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x46, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x22, 0x4b, 0x0a, 0x0d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x46,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x1e, 0x0a, 0x1a, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45,
	0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4a, 0x53, 0x4f, 0x4e, 0x10, 0x01, 0x12,
	0x10, 0x0a, 0x0c, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x5f, 0x42, 0x49, 0x4e, 0x41, 0x52, 0x59, 0x10,
	0x02, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d,
	0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x76, 0x2f, 0x73, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x3b, 0x6c,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_go_chromium_org_luci_cv_settings_listener_settings_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_cv_settings_listener_settings_proto_rawDescData = file_go_chromium_org_luci_cv_settings_listener_settings_proto_rawDesc
)

func file_go_chromium_org_luci_cv_settings_listener_settings_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_cv_settings_listener_settings_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_cv_settings_listener_settings_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_cv_settings_listener_settings_proto_rawDescData)
	})
	return file_go_chromium_org_luci_cv_settings_listener_settings_proto_rawDescData
}

var file_go_chromium_org_luci_cv_settings_listener_settings_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_go_chromium_org_luci_cv_settings_listener_settings_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_go_chromium_org_luci_cv_settings_listener_settings_proto_goTypes = []interface{}{
	(Settings_GerritSubscription_MessageFormat)(0), // 0: listener.Settings.GerritSubscription.MessageFormat
	(*Settings)(nil),                    // 1: listener.Settings
	(*Settings_ReceiveSettings)(nil),    // 2: listener.Settings.ReceiveSettings
	(*Settings_GerritSubscription)(nil), // 3: listener.Settings.GerritSubscription
}
var file_go_chromium_org_luci_cv_settings_listener_settings_proto_depIdxs = []int32{
	3, // 0: listener.Settings.gerrit_subscriptions:type_name -> listener.Settings.GerritSubscription
	2, // 1: listener.Settings.GerritSubscription.receive_settings:type_name -> listener.Settings.ReceiveSettings
	0, // 2: listener.Settings.GerritSubscription.message_format:type_name -> listener.Settings.GerritSubscription.MessageFormat
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_cv_settings_listener_settings_proto_init() }
func file_go_chromium_org_luci_cv_settings_listener_settings_proto_init() {
	if File_go_chromium_org_luci_cv_settings_listener_settings_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_cv_settings_listener_settings_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Settings); i {
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
		file_go_chromium_org_luci_cv_settings_listener_settings_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Settings_ReceiveSettings); i {
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
		file_go_chromium_org_luci_cv_settings_listener_settings_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Settings_GerritSubscription); i {
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
			RawDescriptor: file_go_chromium_org_luci_cv_settings_listener_settings_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_cv_settings_listener_settings_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_cv_settings_listener_settings_proto_depIdxs,
		EnumInfos:         file_go_chromium_org_luci_cv_settings_listener_settings_proto_enumTypes,
		MessageInfos:      file_go_chromium_org_luci_cv_settings_listener_settings_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_cv_settings_listener_settings_proto = out.File
	file_go_chromium_org_luci_cv_settings_listener_settings_proto_rawDesc = nil
	file_go_chromium_org_luci_cv_settings_listener_settings_proto_goTypes = nil
	file_go_chromium_org_luci_cv_settings_listener_settings_proto_depIdxs = nil
}
