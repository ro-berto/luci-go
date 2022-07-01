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
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: go.chromium.org/luci/deploy/api/modelpb/deployment.proto

package modelpb

import (
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

// Types of events to notify on.
type DeploymentConfig_Notification_Event int32

const (
	DeploymentConfig_Notification_EVENT_UNSPECIFIED DeploymentConfig_Notification_Event = 0
	// An actuation is starting.
	//
	// Always emitted regardless of any prior history.
	DeploymentConfig_Notification_ACTUATION_STARTING DeploymentConfig_Notification_Event = 1
	// An actuation has finished successfully.
	//
	// Always emitted regardless of any prior history.
	DeploymentConfig_Notification_ACTUATION_SUCCEEDED DeploymentConfig_Notification_Event = 2
	// An actuation failed, perhaps after several retries.
	//
	// First few failures (below `consecutive_failures` threshold) will *not*
	// result in an event. Every consecutive failure after that will result
	// in an event.
	DeploymentConfig_Notification_ACTUATION_FAILED DeploymentConfig_Notification_Event = 3
	// An actuation succeeded after a series of failures.
	//
	// Emitted if the actuation succeeded after >= `consecutive_failures`
	// consecutive failures. Overrides ACTUATION_SUCCEEDED if emitted by the
	// same state transition.
	DeploymentConfig_Notification_ACTUATION_FIXED DeploymentConfig_Notification_Event = 4
)

// Enum value maps for DeploymentConfig_Notification_Event.
var (
	DeploymentConfig_Notification_Event_name = map[int32]string{
		0: "EVENT_UNSPECIFIED",
		1: "ACTUATION_STARTING",
		2: "ACTUATION_SUCCEEDED",
		3: "ACTUATION_FAILED",
		4: "ACTUATION_FIXED",
	}
	DeploymentConfig_Notification_Event_value = map[string]int32{
		"EVENT_UNSPECIFIED":   0,
		"ACTUATION_STARTING":  1,
		"ACTUATION_SUCCEEDED": 2,
		"ACTUATION_FAILED":    3,
		"ACTUATION_FIXED":     4,
	}
)

func (x DeploymentConfig_Notification_Event) Enum() *DeploymentConfig_Notification_Event {
	p := new(DeploymentConfig_Notification_Event)
	*p = x
	return p
}

func (x DeploymentConfig_Notification_Event) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeploymentConfig_Notification_Event) Descriptor() protoreflect.EnumDescriptor {
	return file_go_chromium_org_luci_deploy_api_modelpb_deployment_proto_enumTypes[0].Descriptor()
}

func (DeploymentConfig_Notification_Event) Type() protoreflect.EnumType {
	return &file_go_chromium_org_luci_deploy_api_modelpb_deployment_proto_enumTypes[0]
}

func (x DeploymentConfig_Notification_Event) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeploymentConfig_Notification_Event.Descriptor instead.
func (DeploymentConfig_Notification_Event) EnumDescriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_deploy_api_modelpb_deployment_proto_rawDescGZIP(), []int{2, 0, 0}
}

// Identifier of a deployment: a reference to its config location.
//
// A deployment is located in some directory of a git repository on `HEAD` ref.
//
// The directory path has two segments: the path to the root of the
// configuration tree, and the path to the particular configuration within this
// tree.
type DeploymentID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Hostname of the GoB server with the IaC repository, e.g. "chrome-internal".
	RepoHost string `protobuf:"bytes,1,opt,name=repo_host,json=repoHost,proto3" json:"repo_host,omitempty"`
	// Name of the IaC repository relative to the host, e.g. "infradata/gae".
	RepoName string `protobuf:"bytes,2,opt,name=repo_name,json=repoName,proto3" json:"repo_name,omitempty"`
	// Path to the root of the IaC config tree in the repository, e.g. ".".
	RepoPath string `protobuf:"bytes,3,opt,name=repo_path,json=repoPath,proto3" json:"repo_path,omitempty"`
	// Path to the directory inside the IaC configs, e.g. "apps/luci-deploy/prod".
	ConfigPath string `protobuf:"bytes,4,opt,name=config_path,json=configPath,proto3" json:"config_path,omitempty"`
}

func (x *DeploymentID) Reset() {
	*x = DeploymentID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_deploy_api_modelpb_deployment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeploymentID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeploymentID) ProtoMessage() {}

func (x *DeploymentID) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_deploy_api_modelpb_deployment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeploymentID.ProtoReflect.Descriptor instead.
func (*DeploymentID) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_deploy_api_modelpb_deployment_proto_rawDescGZIP(), []int{0}
}

func (x *DeploymentID) GetRepoHost() string {
	if x != nil {
		return x.RepoHost
	}
	return ""
}

func (x *DeploymentID) GetRepoName() string {
	if x != nil {
		return x.RepoName
	}
	return ""
}

func (x *DeploymentID) GetRepoPath() string {
	if x != nil {
		return x.RepoPath
	}
	return ""
}

func (x *DeploymentID) GetConfigPath() string {
	if x != nil {
		return x.ConfigPath
	}
	return ""
}

// Deployment as defined in the IaC repo.
type Deployment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The stable identifier of the deployment based on the config location.
	Id *DeploymentID `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The revision of the HEAD of IaC repository with the deployment.
	//
	// Can be obtained with `git rev-parse HEAD`. Moves frequently.
	RepoRev string `protobuf:"bytes,2,opt,name=repo_rev,json=repoRev,proto3" json:"repo_rev,omitempty"`
	// The revision of the deployment config directory.
	//
	// Can be obtained `git log --pretty=tformat:"%H" -n1 <config_path>`. If it
	// changes, the deployment most like is changing too.
	ConfigRev string `protobuf:"bytes,3,opt,name=config_rev,json=configRev,proto3" json:"config_rev,omitempty"`
	// Configuration for the deployment system itself, as defined in the IaC repo.
	//
	// Applies to all assets associated with this deployment.
	Config *DeploymentConfig `protobuf:"bytes,4,opt,name=config,proto3" json:"config,omitempty"`
	// Details of the commit matching `config_rev`.
	LatestCommit *CommitDetails `protobuf:"bytes,5,opt,name=latest_commit,json=latestCommit,proto3" json:"latest_commit,omitempty"`
}

func (x *Deployment) Reset() {
	*x = Deployment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_deploy_api_modelpb_deployment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Deployment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Deployment) ProtoMessage() {}

func (x *Deployment) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_deploy_api_modelpb_deployment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Deployment.ProtoReflect.Descriptor instead.
func (*Deployment) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_deploy_api_modelpb_deployment_proto_rawDescGZIP(), []int{1}
}

func (x *Deployment) GetId() *DeploymentID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Deployment) GetRepoRev() string {
	if x != nil {
		return x.RepoRev
	}
	return ""
}

func (x *Deployment) GetConfigRev() string {
	if x != nil {
		return x.ConfigRev
	}
	return ""
}

func (x *Deployment) GetConfig() *DeploymentConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *Deployment) GetLatestCommit() *CommitDetails {
	if x != nil {
		return x.LatestCommit
	}
	return nil
}

// Deployment configuration, as defined in the IaC repo.
type DeploymentConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// How long the actuation can run before it is considered crashed.
	//
	// Default is 20 min.
	ActuationTimeout *durationpb.Duration             `protobuf:"bytes,1,opt,name=actuation_timeout,json=actuationTimeout,proto3" json:"actuation_timeout,omitempty"`
	Notifications    []*DeploymentConfig_Notification `protobuf:"bytes,2,rep,name=notifications,proto3" json:"notifications,omitempty"`
}

func (x *DeploymentConfig) Reset() {
	*x = DeploymentConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_deploy_api_modelpb_deployment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeploymentConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeploymentConfig) ProtoMessage() {}

func (x *DeploymentConfig) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_deploy_api_modelpb_deployment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeploymentConfig.ProtoReflect.Descriptor instead.
func (*DeploymentConfig) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_deploy_api_modelpb_deployment_proto_rawDescGZIP(), []int{2}
}

func (x *DeploymentConfig) GetActuationTimeout() *durationpb.Duration {
	if x != nil {
		return x.ActuationTimeout
	}
	return nil
}

func (x *DeploymentConfig) GetNotifications() []*DeploymentConfig_Notification {
	if x != nil {
		return x.Notifications
	}
	return nil
}

// Details of an IaC repo commit, to show in the UI.
type CommitDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Author name, as git understands it.
	AuthorName string `protobuf:"bytes,1,opt,name=author_name,json=authorName,proto3" json:"author_name,omitempty"`
	// Author email, as git understands it.
	AuthorEmail string `protobuf:"bytes,2,opt,name=author_email,json=authorEmail,proto3" json:"author_email,omitempty"`
	// Full commit message, including subject and footers.
	CommitMessage string `protobuf:"bytes,3,opt,name=commit_message,json=commitMessage,proto3" json:"commit_message,omitempty"`
}

func (x *CommitDetails) Reset() {
	*x = CommitDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_deploy_api_modelpb_deployment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitDetails) ProtoMessage() {}

func (x *CommitDetails) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_deploy_api_modelpb_deployment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitDetails.ProtoReflect.Descriptor instead.
func (*CommitDetails) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_deploy_api_modelpb_deployment_proto_rawDescGZIP(), []int{3}
}

func (x *CommitDetails) GetAuthorName() string {
	if x != nil {
		return x.AuthorName
	}
	return ""
}

func (x *CommitDetails) GetAuthorEmail() string {
	if x != nil {
		return x.AuthorEmail
	}
	return ""
}

func (x *CommitDetails) GetCommitMessage() string {
	if x != nil {
		return x.CommitMessage
	}
	return ""
}

// Who to notify on noteworthy events.
//
// Various asset state transitions emit 0 or more events per transition. For
// each event kind `notifications` define a set of destinations to send it
// to. During a state transition, all emitted events are bucketed by their
// destination, then redundant events are trimmed (e.g. if a single
// destination is getting ACTUATION_FIXED and ACTUATION_SUCCEEDED events, only
// ACTUATION_FIXED will be retained, since it makes little sense to send
// two success notifications to the same destination at the same time).
type DeploymentConfig_Notification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events []DeploymentConfig_Notification_Event `protobuf:"varint,1,rep,packed,name=events,proto3,enum=deploy.model.DeploymentConfig_Notification_Event" json:"events,omitempty"`
	// List of emails to send notifications to.
	Emails     []string                                   `protobuf:"bytes,2,rep,name=emails,proto3" json:"emails,omitempty"`
	ChatSpaces []*DeploymentConfig_Notification_ChatSpace `protobuf:"bytes,3,rep,name=chat_spaces,json=chatSpaces,proto3" json:"chat_spaces,omitempty"`
	// Consecutive failures threshold for ACTUATION_FAILED and ACTUATION_FIXED.
	//
	// First few failures (below the threshold) will *not* result in
	// a notification. This is useful to avoid spamming on flaky failures
	// resolved by automatic retries.
	ConsecutiveFailures int32 `protobuf:"varint,4,opt,name=consecutive_failures,json=consecutiveFailures,proto3" json:"consecutive_failures,omitempty"`
}

func (x *DeploymentConfig_Notification) Reset() {
	*x = DeploymentConfig_Notification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_deploy_api_modelpb_deployment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeploymentConfig_Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeploymentConfig_Notification) ProtoMessage() {}

func (x *DeploymentConfig_Notification) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_deploy_api_modelpb_deployment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeploymentConfig_Notification.ProtoReflect.Descriptor instead.
func (*DeploymentConfig_Notification) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_deploy_api_modelpb_deployment_proto_rawDescGZIP(), []int{2, 0}
}

func (x *DeploymentConfig_Notification) GetEvents() []DeploymentConfig_Notification_Event {
	if x != nil {
		return x.Events
	}
	return nil
}

func (x *DeploymentConfig_Notification) GetEmails() []string {
	if x != nil {
		return x.Emails
	}
	return nil
}

func (x *DeploymentConfig_Notification) GetChatSpaces() []*DeploymentConfig_Notification_ChatSpace {
	if x != nil {
		return x.ChatSpaces
	}
	return nil
}

func (x *DeploymentConfig_Notification) GetConsecutiveFailures() int32 {
	if x != nil {
		return x.ConsecutiveFailures
	}
	return 0
}

// List of Google Chat spaces to send notifications to.
type DeploymentConfig_Notification_ChatSpace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A preregistered web hook URL to POST notifications too.
	//
	// See https://developers.google.com/chat/how-tos/webhooks#create_a_webhook
	WebhookUrl string `protobuf:"bytes,1,opt,name=webhook_url,json=webhookUrl,proto3" json:"webhook_url,omitempty"`
}

func (x *DeploymentConfig_Notification_ChatSpace) Reset() {
	*x = DeploymentConfig_Notification_ChatSpace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_deploy_api_modelpb_deployment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeploymentConfig_Notification_ChatSpace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeploymentConfig_Notification_ChatSpace) ProtoMessage() {}

func (x *DeploymentConfig_Notification_ChatSpace) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_deploy_api_modelpb_deployment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeploymentConfig_Notification_ChatSpace.ProtoReflect.Descriptor instead.
func (*DeploymentConfig_Notification_ChatSpace) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_deploy_api_modelpb_deployment_proto_rawDescGZIP(), []int{2, 0, 0}
}

func (x *DeploymentConfig_Notification_ChatSpace) GetWebhookUrl() string {
	if x != nil {
		return x.WebhookUrl
	}
	return ""
}

var File_go_chromium_org_luci_deploy_api_modelpb_deployment_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_deploy_api_modelpb_deployment_proto_rawDesc = []byte{
	0x0a, 0x38, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x70, 0x62, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x64, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x01, 0x0a, 0x0c, 0x44, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x70,
	0x6f, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65,
	0x70, 0x6f, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x70, 0x6f, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x70, 0x6f, 0x5f, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x50, 0x61, 0x74, 0x68,
	0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x61, 0x74,
	0x68, 0x22, 0xec, 0x01, 0x0a, 0x0a, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x2a, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x72, 0x65, 0x70, 0x6f, 0x5f, 0x72, 0x65, 0x76, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x72, 0x65, 0x70, 0x6f, 0x52, 0x65, 0x76, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x5f, 0x72, 0x65, 0x76, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x65, 0x76, 0x12, 0x36, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x40,
	0x0a, 0x0d, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x52, 0x0c, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x22, 0xd6, 0x04, 0x0a, 0x10, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x46, 0x0a, 0x11, 0x61, 0x63, 0x74, 0x75, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x61, 0x63, 0x74,
	0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x51, 0x0a,
	0x0d, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0d, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x1a, 0xa6, 0x03, 0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x49, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0e, 0x32, 0x31, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x73, 0x12, 0x56, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x64, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x53, 0x70, 0x61, 0x63, 0x65,
	0x52, 0x0a, 0x63, 0x68, 0x61, 0x74, 0x53, 0x70, 0x61, 0x63, 0x65, 0x73, 0x12, 0x31, 0x0a, 0x14,
	0x63, 0x6f, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x66, 0x61, 0x69, 0x6c,
	0x75, 0x72, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x13, 0x63, 0x6f, 0x6e, 0x73,
	0x65, 0x63, 0x75, 0x74, 0x69, 0x76, 0x65, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x73, 0x1a,
	0x2c, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x74, 0x53, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x55, 0x72, 0x6c, 0x22, 0x7a, 0x0a,
	0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x15, 0x0a, 0x11, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x16, 0x0a,
	0x12, 0x41, 0x43, 0x54, 0x55, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54,
	0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x41, 0x43, 0x54, 0x55, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x45, 0x44, 0x45, 0x44, 0x10, 0x02, 0x12, 0x14,
	0x0a, 0x10, 0x41, 0x43, 0x54, 0x55, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x41, 0x49, 0x4c,
	0x45, 0x44, 0x10, 0x03, 0x12, 0x13, 0x0a, 0x0f, 0x41, 0x43, 0x54, 0x55, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x46, 0x49, 0x58, 0x45, 0x44, 0x10, 0x04, 0x22, 0x7a, 0x0a, 0x0d, 0x43, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x25,
	0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f,
	0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_deploy_api_modelpb_deployment_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_deploy_api_modelpb_deployment_proto_rawDescData = file_go_chromium_org_luci_deploy_api_modelpb_deployment_proto_rawDesc
)

func file_go_chromium_org_luci_deploy_api_modelpb_deployment_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_deploy_api_modelpb_deployment_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_deploy_api_modelpb_deployment_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_deploy_api_modelpb_deployment_proto_rawDescData)
	})
	return file_go_chromium_org_luci_deploy_api_modelpb_deployment_proto_rawDescData
}

var file_go_chromium_org_luci_deploy_api_modelpb_deployment_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_go_chromium_org_luci_deploy_api_modelpb_deployment_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_go_chromium_org_luci_deploy_api_modelpb_deployment_proto_goTypes = []interface{}{
	(DeploymentConfig_Notification_Event)(0),        // 0: deploy.model.DeploymentConfig.Notification.Event
	(*DeploymentID)(nil),                            // 1: deploy.model.DeploymentID
	(*Deployment)(nil),                              // 2: deploy.model.Deployment
	(*DeploymentConfig)(nil),                        // 3: deploy.model.DeploymentConfig
	(*CommitDetails)(nil),                           // 4: deploy.model.CommitDetails
	(*DeploymentConfig_Notification)(nil),           // 5: deploy.model.DeploymentConfig.Notification
	(*DeploymentConfig_Notification_ChatSpace)(nil), // 6: deploy.model.DeploymentConfig.Notification.ChatSpace
	(*durationpb.Duration)(nil),                     // 7: google.protobuf.Duration
}
var file_go_chromium_org_luci_deploy_api_modelpb_deployment_proto_depIdxs = []int32{
	1, // 0: deploy.model.Deployment.id:type_name -> deploy.model.DeploymentID
	3, // 1: deploy.model.Deployment.config:type_name -> deploy.model.DeploymentConfig
	4, // 2: deploy.model.Deployment.latest_commit:type_name -> deploy.model.CommitDetails
	7, // 3: deploy.model.DeploymentConfig.actuation_timeout:type_name -> google.protobuf.Duration
	5, // 4: deploy.model.DeploymentConfig.notifications:type_name -> deploy.model.DeploymentConfig.Notification
	0, // 5: deploy.model.DeploymentConfig.Notification.events:type_name -> deploy.model.DeploymentConfig.Notification.Event
	6, // 6: deploy.model.DeploymentConfig.Notification.chat_spaces:type_name -> deploy.model.DeploymentConfig.Notification.ChatSpace
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_deploy_api_modelpb_deployment_proto_init() }
func file_go_chromium_org_luci_deploy_api_modelpb_deployment_proto_init() {
	if File_go_chromium_org_luci_deploy_api_modelpb_deployment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_deploy_api_modelpb_deployment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeploymentID); i {
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
		file_go_chromium_org_luci_deploy_api_modelpb_deployment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Deployment); i {
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
		file_go_chromium_org_luci_deploy_api_modelpb_deployment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeploymentConfig); i {
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
		file_go_chromium_org_luci_deploy_api_modelpb_deployment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitDetails); i {
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
		file_go_chromium_org_luci_deploy_api_modelpb_deployment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeploymentConfig_Notification); i {
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
		file_go_chromium_org_luci_deploy_api_modelpb_deployment_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeploymentConfig_Notification_ChatSpace); i {
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
			RawDescriptor: file_go_chromium_org_luci_deploy_api_modelpb_deployment_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_deploy_api_modelpb_deployment_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_deploy_api_modelpb_deployment_proto_depIdxs,
		EnumInfos:         file_go_chromium_org_luci_deploy_api_modelpb_deployment_proto_enumTypes,
		MessageInfos:      file_go_chromium_org_luci_deploy_api_modelpb_deployment_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_deploy_api_modelpb_deployment_proto = out.File
	file_go_chromium_org_luci_deploy_api_modelpb_deployment_proto_rawDesc = nil
	file_go_chromium_org_luci_deploy_api_modelpb_deployment_proto_goTypes = nil
	file_go_chromium_org_luci_deploy_api_modelpb_deployment_proto_depIdxs = nil
}
