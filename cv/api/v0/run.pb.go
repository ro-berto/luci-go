// Copyright 2021 The LUCI Authors.
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
// source: go.chromium.org/luci/cv/api/v0/run.proto

package cvpb

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

// Status describes the status of a CV Run.
type Run_Status int32

const (
	// Unspecified status.
	Run_STATUS_UNSPECIFIED Run_Status = 0
	// Run is pending to start.
	//
	// It is either because Run Manager hasn't processed the StartEvent yet or
	// the RunOwner has exhausted all the quota and waiting for new quota to
	// be available.
	Run_PENDING Run_Status = 1
	// Run is running.
	Run_RUNNING Run_Status = 2
	// Run is waiting for submission.
	//
	// Run is in this status if one of the following scenario is true:
	//  1. Tree is closed at the time Run attempts to submit.
	//  2. There is another Run in the same LUCI Project that is currently
	//     submitting.
	//  3. The submission is rate limited according to the submit option in
	//     Project Config.
	//
	// This status is cancellable.
	Run_WAITING_FOR_SUBMISSION Run_Status = 4
	// Run is submitting.
	//
	// A Run can't be cancelled while submitting. A Run may transition from
	// this status to either `WAITING_FOR_SUBMISSION` status or a non-cancelled
	// terminal status.
	Run_SUBMITTING Run_Status = 5
	// ENDED_MASK can be used as a bitmask to check if a Run has ended.
	// This MUST NOT be used as the status of a Run.
	Run_ENDED_MASK Run_Status = 64
	// Run ends successfully.
	Run_SUCCEEDED Run_Status = 65
	// Run ends unsuccessfully.
	Run_FAILED Run_Status = 66
	// Run is cancelled.
	Run_CANCELLED Run_Status = 67
)

// Enum value maps for Run_Status.
var (
	Run_Status_name = map[int32]string{
		0:  "STATUS_UNSPECIFIED",
		1:  "PENDING",
		2:  "RUNNING",
		4:  "WAITING_FOR_SUBMISSION",
		5:  "SUBMITTING",
		64: "ENDED_MASK",
		65: "SUCCEEDED",
		66: "FAILED",
		67: "CANCELLED",
	}
	Run_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED":     0,
		"PENDING":                1,
		"RUNNING":                2,
		"WAITING_FOR_SUBMISSION": 4,
		"SUBMITTING":             5,
		"ENDED_MASK":             64,
		"SUCCEEDED":              65,
		"FAILED":                 66,
		"CANCELLED":              67,
	}
)

func (x Run_Status) Enum() *Run_Status {
	p := new(Run_Status)
	*p = x
	return p
}

func (x Run_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Run_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_go_chromium_org_luci_cv_api_v0_run_proto_enumTypes[0].Descriptor()
}

func (Run_Status) Type() protoreflect.EnumType {
	return &file_go_chromium_org_luci_cv_api_v0_run_proto_enumTypes[0]
}

func (x Run_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Run_Status.Descriptor instead.
func (Run_Status) EnumDescriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_api_v0_run_proto_rawDescGZIP(), []int{0, 0}
}

// Run includes the high-level information about a CV Run.
//
// WARNING: this message is EXPERIMENTAL
// - The field definitions can change w/o notice.
// - No backward compatibility guaranteed.
// - Please contact CV maintainers at luci-eng@ before using this message.
type Run struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the Run.
	//
	// The format of an ID is "projects/$luci-project/runs/$id", where
	// - luci-project is the name of the LUCI project the Run belongs to
	// - id is an opaque key unique in the LUCI project.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Status of the Run.
	Status Run_Status `protobuf:"varint,2,opt,name=status,proto3,enum=cv.v0.Run_Status" json:"status,omitempty"`
	// eversion is the entity version, which is monotonically increasing.
	Eversion int64 `protobuf:"varint,3,opt,name=eversion,proto3" json:"eversion,omitempty"`
	// Mode dictates the behavior of the Run.
	//
	// The possible values include
	// - FULL_RUN
	// - DRY_RUN
	// - QUICK_DRY_RUN
	//
	// If the mode is FULL_RUN, the Run triggers TryJobs and then submits the CL
	// if they succeeded.
	// If the mode is DRY_RUN, the Run trigger TryJobs w/o submission.
	// If the mode is QUICK_DRY_RUN, the Run triggers a different, usually
	// smaller/faster, set of TryJobs.
	Mode string `protobuf:"bytes,4,opt,name=mode,proto3" json:"mode,omitempty"`
	// Time when the Run was created.
	//
	// This is the timestamp of the vote, on a Gerrit CL, that triggered the Run.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The time when the Run was started.
	StartTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// The time when the Run was last updated.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// The time when the Run was ended.
	EndTime *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// Owner of the Run.
	//
	// For a single-CL Run, this is the preferred email of the owner of
	// the Gerrit CL (whoever authenticated to Gerrit to upload the first
	// patchset of the CL). Note that Gerrit CL owner may differ from author and
	// committer encoded in the Git commit. Also, depending on Gerrit
	// configuration, later patchsets could be uploaded by different accounts to
	// that of the CL owner.
	//
	// For a multi-CL Run, this is the owner of the Gerrit CL which has the latest
	// triggering timestamp (e.g. latest CQ+2 vote).
	Owner string `protobuf:"bytes,9,opt,name=owner,proto3" json:"owner,omitempty"`
	// The Gerrit changes involved in this Run.
	Cls []*GerritChange `protobuf:"bytes,10,rep,name=cls,proto3" json:"cls,omitempty"`
	// The tryjobs of the Run.
	//
	// Note that this data is a snapshot at the time Run has ended. Therefore,
	// some tryjobs may not have ended yet. If you need fresh data, query the
	// corresponding backend system using the returned ID.
	//
	// It may also contain tryjobs that are no longer required.
	// TODO(yiwzhang): Re-evalute whether the statement above is true after
	// tryjobs are handled by LUCI CV instead of CQDaemon.
	Tryjobs []*Tryjob `protobuf:"bytes,11,rep,name=tryjobs,proto3" json:"tryjobs,omitempty"`
	// The state of Run Submission.
	//
	// Unset if Submission hasn't started.
	Submission *Run_Submission `protobuf:"bytes,12,opt,name=submission,proto3" json:"submission,omitempty"`
}

func (x *Run) Reset() {
	*x = Run{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_api_v0_run_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Run) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Run) ProtoMessage() {}

func (x *Run) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_api_v0_run_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Run.ProtoReflect.Descriptor instead.
func (*Run) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_api_v0_run_proto_rawDescGZIP(), []int{0}
}

func (x *Run) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Run) GetStatus() Run_Status {
	if x != nil {
		return x.Status
	}
	return Run_STATUS_UNSPECIFIED
}

func (x *Run) GetEversion() int64 {
	if x != nil {
		return x.Eversion
	}
	return 0
}

func (x *Run) GetMode() string {
	if x != nil {
		return x.Mode
	}
	return ""
}

func (x *Run) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Run) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *Run) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Run) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *Run) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *Run) GetCls() []*GerritChange {
	if x != nil {
		return x.Cls
	}
	return nil
}

func (x *Run) GetTryjobs() []*Tryjob {
	if x != nil {
		return x.Tryjobs
	}
	return nil
}

func (x *Run) GetSubmission() *Run_Submission {
	if x != nil {
		return x.Submission
	}
	return nil
}

// A Gerrit patchset.
type GerritChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Gerrit hostname, e.g. "chromium-review.googlesource.com".
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	// Change number, e.g. 12345.
	Change int64 `protobuf:"varint,2,opt,name=change,proto3" json:"change,omitempty"`
	// Patch set number, e.g. 1.
	Patchset int32 `protobuf:"varint,3,opt,name=patchset,proto3" json:"patchset,omitempty"`
}

func (x *GerritChange) Reset() {
	*x = GerritChange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_api_v0_run_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GerritChange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GerritChange) ProtoMessage() {}

func (x *GerritChange) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_api_v0_run_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GerritChange.ProtoReflect.Descriptor instead.
func (*GerritChange) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_api_v0_run_proto_rawDescGZIP(), []int{1}
}

func (x *GerritChange) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *GerritChange) GetChange() int64 {
	if x != nil {
		return x.Change
	}
	return 0
}

func (x *GerritChange) GetPatchset() int32 {
	if x != nil {
		return x.Patchset
	}
	return 0
}

// Submission represents the state of a Run Submission.
type Run_Submission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Indexes of CLs in Run.CL IDs that have been submitted succeessfully.
	SubmittedClIndexes []int32 `protobuf:"varint,2,rep,packed,name=submitted_cl_indexes,json=submittedClIndexes,proto3" json:"submitted_cl_indexes,omitempty"`
	// Indexes of CLs in Run.CL IDs that failed to be submitted.
	//
	// CLs that are neither in this list nor in the `submitted_cl_indexes`
	// should be treated as if CV has never attempted to submit them.
	FailedClIndexes []int32 `protobuf:"varint,3,rep,packed,name=failed_cl_indexes,json=failedClIndexes,proto3" json:"failed_cl_indexes,omitempty"`
}

func (x *Run_Submission) Reset() {
	*x = Run_Submission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_api_v0_run_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Run_Submission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Run_Submission) ProtoMessage() {}

func (x *Run_Submission) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_api_v0_run_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Run_Submission.ProtoReflect.Descriptor instead.
func (*Run_Submission) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_api_v0_run_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Run_Submission) GetSubmittedClIndexes() []int32 {
	if x != nil {
		return x.SubmittedClIndexes
	}
	return nil
}

func (x *Run_Submission) GetFailedClIndexes() []int32 {
	if x != nil {
		return x.FailedClIndexes
	}
	return nil
}

var File_go_chromium_org_luci_cv_api_v0_run_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_cv_api_v0_run_proto_rawDesc = []byte{
	0x0a, 0x28, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x76, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x30,
	0x2f, 0x72, 0x75, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x63, 0x76, 0x2e, 0x76,
	0x30, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2b, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e,
	0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x76, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x30, 0x2f, 0x74, 0x72, 0x79, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x88, 0x06, 0x0a, 0x03, 0x52, 0x75, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x29, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x63, 0x76, 0x2e, 0x76, 0x30, 0x2e,
	0x52, 0x75, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x65, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x6f,
	0x64, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x03, 0x63, 0x6c, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x63, 0x76, 0x2e, 0x76, 0x30, 0x2e, 0x47, 0x65, 0x72, 0x72, 0x69, 0x74,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x03, 0x63, 0x6c, 0x73, 0x12, 0x27, 0x0a, 0x07, 0x74,
	0x72, 0x79, 0x6a, 0x6f, 0x62, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63,
	0x76, 0x2e, 0x76, 0x30, 0x2e, 0x54, 0x72, 0x79, 0x6a, 0x6f, 0x62, 0x52, 0x07, 0x74, 0x72, 0x79,
	0x6a, 0x6f, 0x62, 0x73, 0x12, 0x35, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x76, 0x2e, 0x76, 0x30,
	0x2e, 0x52, 0x75, 0x6e, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x0a, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x6a, 0x0a, 0x0a, 0x53,
	0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x14, 0x73, 0x75, 0x62,
	0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x6c, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x12, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x74,
	0x65, 0x64, 0x43, 0x6c, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x66,
	0x61, 0x69, 0x6c, 0x65, 0x64, 0x5f, 0x63, 0x6c, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0f, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x43, 0x6c,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x22, 0xa0, 0x01, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45,
	0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x55, 0x4e, 0x4e, 0x49,
	0x4e, 0x47, 0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16, 0x57, 0x41, 0x49, 0x54, 0x49, 0x4e, 0x47, 0x5f,
	0x46, 0x4f, 0x52, 0x5f, 0x53, 0x55, 0x42, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x04,
	0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x55, 0x42, 0x4d, 0x49, 0x54, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x05,
	0x12, 0x0e, 0x0a, 0x0a, 0x45, 0x4e, 0x44, 0x45, 0x44, 0x5f, 0x4d, 0x41, 0x53, 0x4b, 0x10, 0x40,
	0x12, 0x0d, 0x0a, 0x09, 0x53, 0x55, 0x43, 0x43, 0x45, 0x45, 0x44, 0x45, 0x44, 0x10, 0x41, 0x12,
	0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x42, 0x12, 0x0d, 0x0a, 0x09, 0x43,
	0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x43, 0x22, 0x56, 0x0a, 0x0c, 0x47, 0x65,
	0x72, 0x72, 0x69, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x74, 0x63, 0x68, 0x73,
	0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x74, 0x63, 0x68, 0x73,
	0x65, 0x74, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75,
	0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x76, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x30, 0x3b, 0x63, 0x76, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_go_chromium_org_luci_cv_api_v0_run_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_cv_api_v0_run_proto_rawDescData = file_go_chromium_org_luci_cv_api_v0_run_proto_rawDesc
)

func file_go_chromium_org_luci_cv_api_v0_run_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_cv_api_v0_run_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_cv_api_v0_run_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_cv_api_v0_run_proto_rawDescData)
	})
	return file_go_chromium_org_luci_cv_api_v0_run_proto_rawDescData
}

var file_go_chromium_org_luci_cv_api_v0_run_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_go_chromium_org_luci_cv_api_v0_run_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_go_chromium_org_luci_cv_api_v0_run_proto_goTypes = []interface{}{
	(Run_Status)(0),               // 0: cv.v0.Run.Status
	(*Run)(nil),                   // 1: cv.v0.Run
	(*GerritChange)(nil),          // 2: cv.v0.GerritChange
	(*Run_Submission)(nil),        // 3: cv.v0.Run.Submission
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
	(*Tryjob)(nil),                // 5: cv.v0.Tryjob
}
var file_go_chromium_org_luci_cv_api_v0_run_proto_depIdxs = []int32{
	0, // 0: cv.v0.Run.status:type_name -> cv.v0.Run.Status
	4, // 1: cv.v0.Run.create_time:type_name -> google.protobuf.Timestamp
	4, // 2: cv.v0.Run.start_time:type_name -> google.protobuf.Timestamp
	4, // 3: cv.v0.Run.update_time:type_name -> google.protobuf.Timestamp
	4, // 4: cv.v0.Run.end_time:type_name -> google.protobuf.Timestamp
	2, // 5: cv.v0.Run.cls:type_name -> cv.v0.GerritChange
	5, // 6: cv.v0.Run.tryjobs:type_name -> cv.v0.Tryjob
	3, // 7: cv.v0.Run.submission:type_name -> cv.v0.Run.Submission
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_cv_api_v0_run_proto_init() }
func file_go_chromium_org_luci_cv_api_v0_run_proto_init() {
	if File_go_chromium_org_luci_cv_api_v0_run_proto != nil {
		return
	}
	file_go_chromium_org_luci_cv_api_v0_tryjob_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_cv_api_v0_run_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Run); i {
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
		file_go_chromium_org_luci_cv_api_v0_run_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GerritChange); i {
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
		file_go_chromium_org_luci_cv_api_v0_run_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Run_Submission); i {
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
			RawDescriptor: file_go_chromium_org_luci_cv_api_v0_run_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_cv_api_v0_run_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_cv_api_v0_run_proto_depIdxs,
		EnumInfos:         file_go_chromium_org_luci_cv_api_v0_run_proto_enumTypes,
		MessageInfos:      file_go_chromium_org_luci_cv_api_v0_run_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_cv_api_v0_run_proto = out.File
	file_go_chromium_org_luci_cv_api_v0_run_proto_rawDesc = nil
	file_go_chromium_org_luci_cv_api_v0_run_proto_goTypes = nil
	file_go_chromium_org_luci_cv_api_v0_run_proto_depIdxs = nil
}
