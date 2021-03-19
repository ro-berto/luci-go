// Copyright 2019 The LUCI Authors.
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
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.1
// source: go.chromium.org/luci/cv/api/recipe/v1/cq.proto

package recipe

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

type Output_Retry int32

const (
	Output_OUTPUT_RETRY_UNSPECIFIED Output_Retry = 0
	// Default. Allow CQ to retry the build.
	//
	// Does NOT force CQ to retry this build, since it depends on other factors,
	// such as the applicable project's CQ config.
	Output_OUTPUT_RETRY_ALLOWED Output_Retry = 1
	// Denies retries regardless of other factors.
	//
	// This is equivalent to setting legacy top-level `"do_not_retry": true`
	// output property.
	// TODO(tandrii): deprecate and remove the legacy property.
	Output_OUTPUT_RETRY_DENIED Output_Retry = 2
)

// Enum value maps for Output_Retry.
var (
	Output_Retry_name = map[int32]string{
		0: "OUTPUT_RETRY_UNSPECIFIED",
		1: "OUTPUT_RETRY_ALLOWED",
		2: "OUTPUT_RETRY_DENIED",
	}
	Output_Retry_value = map[string]int32{
		"OUTPUT_RETRY_UNSPECIFIED": 0,
		"OUTPUT_RETRY_ALLOWED":     1,
		"OUTPUT_RETRY_DENIED":      2,
	}
)

func (x Output_Retry) Enum() *Output_Retry {
	p := new(Output_Retry)
	*p = x
	return p
}

func (x Output_Retry) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Output_Retry) Descriptor() protoreflect.EnumDescriptor {
	return file_go_chromium_org_luci_cv_api_recipe_v1_cq_proto_enumTypes[0].Descriptor()
}

func (Output_Retry) Type() protoreflect.EnumType {
	return &file_go_chromium_org_luci_cv_api_recipe_v1_cq_proto_enumTypes[0]
}

func (x Output_Retry) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Output_Retry.Descriptor instead.
func (Output_Retry) EnumDescriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_api_recipe_v1_cq_proto_rawDescGZIP(), []int{1, 0}
}

// Input provides CQ metadata for CQ-triggered tryjob.
type Input struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If true, CQ is active for the current build. CQ is considered "active" for
	// a build if CQ triggered the build, either directly or indirectly.
	Active bool `protobuf:"varint,1,opt,name=active,proto3" json:"active,omitempty"`
	// If false, CQ would try to submit CL(s) if all other checks pass.
	// If true, CQ won't try to submit.
	//
	// DEPRECATED: Use run_mode instead.
	DryRun bool `protobuf:"varint,2,opt,name=dry_run,json=dryRun,proto3" json:"dry_run,omitempty"`
	// If true, CQ will not take this build into account while deciding whether
	// CL is good or not. See also `experiment_percentage` of CQ's config file.
	Experimental bool `protobuf:"varint,3,opt,name=experimental,proto3" json:"experimental,omitempty"`
	// If true, CQ triggered this build directly, otherwise typically indicates a
	// child build triggered by a CQ triggered one (possibly indirectly).
	//
	// Can be spoofed. *DO NOT USE FOR SECURITY CHECKS.*
	//
	// One possible use is to distinguish which builds must be cancelled manually,
	// and which (top_level=True) CQ would cancel itself.
	TopLevel bool `protobuf:"varint,4,opt,name=top_level,json=topLevel,proto3" json:"top_level,omitempty"`
	// The mode of the CQ Run that triggers this Tryjob.
	RunMode string `protobuf:"bytes,5,opt,name=run_mode,json=runMode,proto3" json:"run_mode,omitempty"`
}

func (x *Input) Reset() {
	*x = Input{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_api_recipe_v1_cq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Input) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Input) ProtoMessage() {}

func (x *Input) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_api_recipe_v1_cq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Input.ProtoReflect.Descriptor instead.
func (*Input) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_api_recipe_v1_cq_proto_rawDescGZIP(), []int{0}
}

func (x *Input) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *Input) GetDryRun() bool {
	if x != nil {
		return x.DryRun
	}
	return false
}

func (x *Input) GetExperimental() bool {
	if x != nil {
		return x.Experimental
	}
	return false
}

func (x *Input) GetTopLevel() bool {
	if x != nil {
		return x.TopLevel
	}
	return false
}

func (x *Input) GetRunMode() string {
	if x != nil {
		return x.RunMode
	}
	return ""
}

// Output provides build-specific instructions back to CQ.
//
// Unless stated otherwise, each Output message field can be set even on builds
// not triggered directly or indirectly by CQ itself. For example, `git cl try`
// or Gerrit UI can be used to trigger a build directly, which can then instruct
// CQ not to retry it.
//
// CQ periodically checks the Output of still running builds, too,
// and may act on the Output even before a build is completed.
type Output struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Buildbucket build IDs which this build has triggered for CQ to wait on.
	//
	// Required when using triggered_by builders in project's CQ config.
	// This is useful to allow the triggering builder to finish without waiting
	// for its child builds, which can be efficiently done by CQ.
	//
	// This is equivalent to setting legacy top-level "triggered_build_ids" output
	// property.
	// TODO(tandrii): deprecate and remove the legacy property.
	TriggeredBuildIds []int64 `protobuf:"varint,1,rep,packed,name=triggered_build_ids,json=triggeredBuildIds,proto3" json:"triggered_build_ids,omitempty"`
	// Retry controls whether this build can be retried by CQ.
	Retry Output_Retry `protobuf:"varint,2,opt,name=retry,proto3,enum=cq.recipe.Output_Retry" json:"retry,omitempty"`
	// Reuse restricts potential reuse of this build by a later CQ run.
	//
	// NOTE: even if reuse is not restricted here, reuse is still subject to other
	// restrictions in applicable project's CQ config.
	//
	// If empty (default), reuse is *allowed*.
	//
	// If specified, the order matters: the first matching Reuse message wins.
	// If specified and no Reuse match the run, reuse is *not allowed*.
	// If any individual Reuse block is invalid, reuse is *not allowed*.
	//
	// Examples:
	//
	//  1. To prohibit reuse only for Full runs, do:
	//     {mode_regexp: "fullrun" deny: true}
	//     {mode_regexp: ".+"      deny: false}
	//
	//  2. To prohibit reuse for everything except Dry Runs, do:
	//     {mode_regexp: "dryrun"}
	Reuse []*Output_Reuse `protobuf:"bytes,3,rep,name=reuse,proto3" json:"reuse,omitempty"`
}

func (x *Output) Reset() {
	*x = Output{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_api_recipe_v1_cq_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Output) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Output) ProtoMessage() {}

func (x *Output) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_api_recipe_v1_cq_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Output.ProtoReflect.Descriptor instead.
func (*Output) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_api_recipe_v1_cq_proto_rawDescGZIP(), []int{1}
}

func (x *Output) GetTriggeredBuildIds() []int64 {
	if x != nil {
		return x.TriggeredBuildIds
	}
	return nil
}

func (x *Output) GetRetry() Output_Retry {
	if x != nil {
		return x.Retry
	}
	return Output_OUTPUT_RETRY_UNSPECIFIED
}

func (x *Output) GetReuse() []*Output_Reuse {
	if x != nil {
		return x.Reuse
	}
	return nil
}

type Output_Reuse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Regular expression for modes of Runs for which this Reuse block applies.
	// Required.
	//
	// Implicitly wrapped with (?i)$...^  (= complete case-insensitive match).
	//
	// For example,
	//   ".+" will match all modes of Runs,
	//   "dryrun" and "fullrun" will match only Dry and Full runs, respectively.
	ModeRegexp string `protobuf:"bytes,1,opt,name=mode_regexp,json=modeRegexp,proto3" json:"mode_regexp,omitempty"`
	// If deny is true, then reuse of this build in the future Runs of the
	// matched mode is not allowed.
	//
	// If false, then reuse is allowed. It's useful to stop the matching in case
	// of several Reuse messages.
	Deny bool `protobuf:"varint,2,opt,name=deny,proto3" json:"deny,omitempty"` // TODO(crbug/753103): add reuse duration or deadline.
}

func (x *Output_Reuse) Reset() {
	*x = Output_Reuse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_api_recipe_v1_cq_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Output_Reuse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Output_Reuse) ProtoMessage() {}

func (x *Output_Reuse) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_api_recipe_v1_cq_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Output_Reuse.ProtoReflect.Descriptor instead.
func (*Output_Reuse) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_api_recipe_v1_cq_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Output_Reuse) GetModeRegexp() string {
	if x != nil {
		return x.ModeRegexp
	}
	return ""
}

func (x *Output_Reuse) GetDeny() bool {
	if x != nil {
		return x.Deny
	}
	return false
}

var File_go_chromium_org_luci_cv_api_recipe_v1_cq_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_cv_api_recipe_v1_cq_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x76, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65,
	0x63, 0x69, 0x70, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x09, 0x63, 0x71, 0x2e, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x22, 0x94, 0x01, 0x0a, 0x05,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x17, 0x0a,
	0x07, 0x64, 0x72, 0x79, 0x5f, 0x72, 0x75, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x64, 0x72, 0x79, 0x52, 0x75, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69,
	0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x65, 0x78,
	0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x6f,
	0x70, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x74,
	0x6f, 0x70, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x75, 0x6e, 0x5f, 0x6d,
	0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x4d, 0x6f,
	0x64, 0x65, 0x22, 0xae, 0x02, 0x0a, 0x06, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x2e, 0x0a,
	0x13, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x65, 0x64, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x11, 0x74, 0x72, 0x69, 0x67,
	0x67, 0x65, 0x72, 0x65, 0x64, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x64, 0x73, 0x12, 0x2d, 0x0a,
	0x05, 0x72, 0x65, 0x74, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x63,
	0x71, 0x2e, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e,
	0x52, 0x65, 0x74, 0x72, 0x79, 0x52, 0x05, 0x72, 0x65, 0x74, 0x72, 0x79, 0x12, 0x2d, 0x0a, 0x05,
	0x72, 0x65, 0x75, 0x73, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x71,
	0x2e, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e, 0x52,
	0x65, 0x75, 0x73, 0x65, 0x52, 0x05, 0x72, 0x65, 0x75, 0x73, 0x65, 0x1a, 0x3c, 0x0a, 0x05, 0x52,
	0x65, 0x75, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x6f, 0x64, 0x65, 0x5f, 0x72, 0x65, 0x67,
	0x65, 0x78, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x67, 0x65, 0x78, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x6e, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x04, 0x64, 0x65, 0x6e, 0x79, 0x22, 0x58, 0x0a, 0x05, 0x52, 0x65, 0x74,
	0x72, 0x79, 0x12, 0x1c, 0x0a, 0x18, 0x4f, 0x55, 0x54, 0x50, 0x55, 0x54, 0x5f, 0x52, 0x45, 0x54,
	0x52, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x18, 0x0a, 0x14, 0x4f, 0x55, 0x54, 0x50, 0x55, 0x54, 0x5f, 0x52, 0x45, 0x54, 0x52, 0x59,
	0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x44, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x4f, 0x55,
	0x54, 0x50, 0x55, 0x54, 0x5f, 0x52, 0x45, 0x54, 0x52, 0x59, 0x5f, 0x44, 0x45, 0x4e, 0x49, 0x45,
	0x44, 0x10, 0x02, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69,
	0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x76, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x72, 0x65, 0x63,
	0x69, 0x70, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_cv_api_recipe_v1_cq_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_cv_api_recipe_v1_cq_proto_rawDescData = file_go_chromium_org_luci_cv_api_recipe_v1_cq_proto_rawDesc
)

func file_go_chromium_org_luci_cv_api_recipe_v1_cq_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_cv_api_recipe_v1_cq_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_cv_api_recipe_v1_cq_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_cv_api_recipe_v1_cq_proto_rawDescData)
	})
	return file_go_chromium_org_luci_cv_api_recipe_v1_cq_proto_rawDescData
}

var file_go_chromium_org_luci_cv_api_recipe_v1_cq_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_go_chromium_org_luci_cv_api_recipe_v1_cq_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_go_chromium_org_luci_cv_api_recipe_v1_cq_proto_goTypes = []interface{}{
	(Output_Retry)(0),    // 0: cq.recipe.Output.Retry
	(*Input)(nil),        // 1: cq.recipe.Input
	(*Output)(nil),       // 2: cq.recipe.Output
	(*Output_Reuse)(nil), // 3: cq.recipe.Output.Reuse
}
var file_go_chromium_org_luci_cv_api_recipe_v1_cq_proto_depIdxs = []int32{
	0, // 0: cq.recipe.Output.retry:type_name -> cq.recipe.Output.Retry
	3, // 1: cq.recipe.Output.reuse:type_name -> cq.recipe.Output.Reuse
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_cv_api_recipe_v1_cq_proto_init() }
func file_go_chromium_org_luci_cv_api_recipe_v1_cq_proto_init() {
	if File_go_chromium_org_luci_cv_api_recipe_v1_cq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_cv_api_recipe_v1_cq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Input); i {
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
		file_go_chromium_org_luci_cv_api_recipe_v1_cq_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Output); i {
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
		file_go_chromium_org_luci_cv_api_recipe_v1_cq_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Output_Reuse); i {
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
			RawDescriptor: file_go_chromium_org_luci_cv_api_recipe_v1_cq_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_cv_api_recipe_v1_cq_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_cv_api_recipe_v1_cq_proto_depIdxs,
		EnumInfos:         file_go_chromium_org_luci_cv_api_recipe_v1_cq_proto_enumTypes,
		MessageInfos:      file_go_chromium_org_luci_cv_api_recipe_v1_cq_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_cv_api_recipe_v1_cq_proto = out.File
	file_go_chromium_org_luci_cv_api_recipe_v1_cq_proto_rawDesc = nil
	file_go_chromium_org_luci_cv_api_recipe_v1_cq_proto_goTypes = nil
	file_go_chromium_org_luci_cv_api_recipe_v1_cq_proto_depIdxs = nil
}
