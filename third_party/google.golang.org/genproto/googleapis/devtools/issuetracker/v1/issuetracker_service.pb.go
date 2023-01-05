// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package issuetracker

import (
	context "context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// View on Issue. Pass this enum to rpcs that returns an Issue message to
// control which subsets of fields to get.
type IssueView int32

const (
	// Unspecified. It defaults to BASIC.
	IssueView_ISSUE_VIEW_UNSPECIFIED IssueView = 0
	// Basic fields.
	IssueView_BASIC IssueView = 1
	// Include all fields.
	IssueView_FULL IssueView = 2
)

// Enum value maps for IssueView.
var (
	IssueView_name = map[int32]string{
		0: "ISSUE_VIEW_UNSPECIFIED",
		1: "BASIC",
		2: "FULL",
	}
	IssueView_value = map[string]int32{
		"ISSUE_VIEW_UNSPECIFIED": 0,
		"BASIC":                  1,
		"FULL":                   2,
	}
)

func (x IssueView) Enum() *IssueView {
	p := new(IssueView)
	*p = x
	return p
}

func (x IssueView) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IssueView) Descriptor() protoreflect.EnumDescriptor {
	return file_google_devtools_issuetracker_v1_issuetracker_service_proto_enumTypes[0].Descriptor()
}

func (IssueView) Type() protoreflect.EnumType {
	return &file_google_devtools_issuetracker_v1_issuetracker_service_proto_enumTypes[0]
}

func (x IssueView) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IssueView.Descriptor instead.
func (IssueView) EnumDescriptor() ([]byte, []int) {
	return file_google_devtools_issuetracker_v1_issuetracker_service_proto_rawDescGZIP(), []int{0}
}

// Issue comment view that specifies which fields are returned. Note: view names
// are prefixed to avoid name collision with other view enums.
type IssueCommentView int32

const (
	// The default / unset value. The API will default to the BASIC view.
	IssueCommentView_ISSUE_COMMENT_VIEW_UNSPECIFIED IssueCommentView = 0
	// Include everything except for the formatted_comment field.
	IssueCommentView_ISSUE_COMMENT_VIEW_BASIC IssueCommentView = 1
	// Include everything.
	IssueCommentView_ISSUE_COMMENT_VIEW_FULL IssueCommentView = 2
)

// Enum value maps for IssueCommentView.
var (
	IssueCommentView_name = map[int32]string{
		0: "ISSUE_COMMENT_VIEW_UNSPECIFIED",
		1: "ISSUE_COMMENT_VIEW_BASIC",
		2: "ISSUE_COMMENT_VIEW_FULL",
	}
	IssueCommentView_value = map[string]int32{
		"ISSUE_COMMENT_VIEW_UNSPECIFIED": 0,
		"ISSUE_COMMENT_VIEW_BASIC":       1,
		"ISSUE_COMMENT_VIEW_FULL":        2,
	}
)

func (x IssueCommentView) Enum() *IssueCommentView {
	p := new(IssueCommentView)
	*p = x
	return p
}

func (x IssueCommentView) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IssueCommentView) Descriptor() protoreflect.EnumDescriptor {
	return file_google_devtools_issuetracker_v1_issuetracker_service_proto_enumTypes[1].Descriptor()
}

func (IssueCommentView) Type() protoreflect.EnumType {
	return &file_google_devtools_issuetracker_v1_issuetracker_service_proto_enumTypes[1]
}

func (x IssueCommentView) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IssueCommentView.Descriptor instead.
func (IssueCommentView) EnumDescriptor() ([]byte, []int) {
	return file_google_devtools_issuetracker_v1_issuetracker_service_proto_rawDescGZIP(), []int{1}
}

// Request object for IssueTracker.GetComponent.
type GetComponentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Numeric ID of the component to look up.
	ComponentId int64 `protobuf:"varint,1,opt,name=component_id,json=componentId,proto3" json:"component_id,omitempty"`
}

func (x *GetComponentRequest) Reset() {
	*x = GetComponentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetComponentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetComponentRequest) ProtoMessage() {}

func (x *GetComponentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetComponentRequest.ProtoReflect.Descriptor instead.
func (*GetComponentRequest) Descriptor() ([]byte, []int) {
	return file_google_devtools_issuetracker_v1_issuetracker_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetComponentRequest) GetComponentId() int64 {
	if x != nil {
		return x.ComponentId
	}
	return 0
}

// Request object for IssueTracker.ListIssues.
type ListIssuesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Query language for issues requests is defined at:
	// https://developers.google.com/issue-tracker/concepts/search-query-language
	// *Required.*
	Query string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	// Order parameter.  Order is ascending ("asc") by default, but can
	// be defined with the sort field followed by a space and "asc" or
	// "desc".
	//
	// Examples:
	//     "issue_id",
	//     "modified_time asc",
	//
	// Secondary sorts may be specified in comma-separated format.
	//
	// Examples:
	//    "priority asc, created_time desc"
	//    "custom_field:1234, modified_time desc"
	//
	// Valid sort fields are:
	//   * archived
	//   * assignee
	//   * cc_count
	//   * component_path
	//   * created_time
	//   * custom_field:<id>
	//   * deletion_time
	//   * duplicate_count
	//   * found_in_versions
	//   * in_prod
	//   * issue_id
	//   * last_modifier
	//   * modified_time
	//   * priority
	//   * reporter
	//   * severity
	//   * status
	//   * targeted_to_versions
	//   * title
	//   * type
	//   * verified_in_versions
	//   * verified_time
	//   * verifier
	//   * vote_count
	OrderBy string `protobuf:"bytes,2,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	// Default page_size = 25.
	// Maximum page_size = 500.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Pagination token. Optional.
	PageToken string `protobuf:"bytes,4,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// The view of the issue to return.
	View IssueView `protobuf:"varint,5,opt,name=view,proto3,enum=google.devtools.issuetracker.v1.IssueView" json:"view,omitempty"`
}

func (x *ListIssuesRequest) Reset() {
	*x = ListIssuesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListIssuesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListIssuesRequest) ProtoMessage() {}

func (x *ListIssuesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListIssuesRequest.ProtoReflect.Descriptor instead.
func (*ListIssuesRequest) Descriptor() ([]byte, []int) {
	return file_google_devtools_issuetracker_v1_issuetracker_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListIssuesRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *ListIssuesRequest) GetOrderBy() string {
	if x != nil {
		return x.OrderBy
	}
	return ""
}

func (x *ListIssuesRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListIssuesRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListIssuesRequest) GetView() IssueView {
	if x != nil {
		return x.View
	}
	return IssueView_ISSUE_VIEW_UNSPECIFIED
}

// Response object for IssueTracker.ListIssues.
type ListIssuesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The current page of issues.
	Issues []*Issue `protobuf:"bytes,1,rep,name=issues,proto3" json:"issues,omitempty"`
	// Pagination token for next page of results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	// Total number of results. This is an approximation.
	TotalSize int32 `protobuf:"varint,3,opt,name=total_size,json=totalSize,proto3" json:"total_size,omitempty"`
}

func (x *ListIssuesResponse) Reset() {
	*x = ListIssuesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListIssuesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListIssuesResponse) ProtoMessage() {}

func (x *ListIssuesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListIssuesResponse.ProtoReflect.Descriptor instead.
func (*ListIssuesResponse) Descriptor() ([]byte, []int) {
	return file_google_devtools_issuetracker_v1_issuetracker_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListIssuesResponse) GetIssues() []*Issue {
	if x != nil {
		return x.Issues
	}
	return nil
}

func (x *ListIssuesResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

func (x *ListIssuesResponse) GetTotalSize() int32 {
	if x != nil {
		return x.TotalSize
	}
	return 0
}

// Request object for IssueTracker.BatchGetIssues.
type BatchGetIssuesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Numeric ID of the issues to look up.
	IssueIds []int64 `protobuf:"varint,1,rep,packed,name=issue_ids,json=issueIds,proto3" json:"issue_ids,omitempty"`
	// The view of the issue to return.
	View IssueView `protobuf:"varint,2,opt,name=view,proto3,enum=google.devtools.issuetracker.v1.IssueView" json:"view,omitempty"`
}

func (x *BatchGetIssuesRequest) Reset() {
	*x = BatchGetIssuesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchGetIssuesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchGetIssuesRequest) ProtoMessage() {}

func (x *BatchGetIssuesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchGetIssuesRequest.ProtoReflect.Descriptor instead.
func (*BatchGetIssuesRequest) Descriptor() ([]byte, []int) {
	return file_google_devtools_issuetracker_v1_issuetracker_service_proto_rawDescGZIP(), []int{3}
}

func (x *BatchGetIssuesRequest) GetIssueIds() []int64 {
	if x != nil {
		return x.IssueIds
	}
	return nil
}

func (x *BatchGetIssuesRequest) GetView() IssueView {
	if x != nil {
		return x.View
	}
	return IssueView_ISSUE_VIEW_UNSPECIFIED
}

// Response object for IssueTracker.BatchGetIssues.
type BatchGetIssuesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The requested issues.
	Issues []*Issue `protobuf:"bytes,1,rep,name=issues,proto3" json:"issues,omitempty"`
}

func (x *BatchGetIssuesResponse) Reset() {
	*x = BatchGetIssuesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchGetIssuesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchGetIssuesResponse) ProtoMessage() {}

func (x *BatchGetIssuesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchGetIssuesResponse.ProtoReflect.Descriptor instead.
func (*BatchGetIssuesResponse) Descriptor() ([]byte, []int) {
	return file_google_devtools_issuetracker_v1_issuetracker_service_proto_rawDescGZIP(), []int{4}
}

func (x *BatchGetIssuesResponse) GetIssues() []*Issue {
	if x != nil {
		return x.Issues
	}
	return nil
}

// Request object for IssueTracker.GetIssue.
type GetIssueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Numeric ID of the issue to look up.
	IssueId int64 `protobuf:"varint,1,opt,name=issue_id,json=issueId,proto3" json:"issue_id,omitempty"`
	// The view of the issue to return.
	View IssueView `protobuf:"varint,2,opt,name=view,proto3,enum=google.devtools.issuetracker.v1.IssueView" json:"view,omitempty"`
}

func (x *GetIssueRequest) Reset() {
	*x = GetIssueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIssueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIssueRequest) ProtoMessage() {}

func (x *GetIssueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIssueRequest.ProtoReflect.Descriptor instead.
func (*GetIssueRequest) Descriptor() ([]byte, []int) {
	return file_google_devtools_issuetracker_v1_issuetracker_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetIssueRequest) GetIssueId() int64 {
	if x != nil {
		return x.IssueId
	}
	return 0
}

func (x *GetIssueRequest) GetView() IssueView {
	if x != nil {
		return x.View
	}
	return IssueView_ISSUE_VIEW_UNSPECIFIED
}

// Request object for IssueTracker.CreateIssue.
type CreateIssueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required fields in the Issue.issue_state:
	//   * title
	//   * component_id
	// Status requirements:
	//  * When an issue is NEW, issue_id must be absent.
	//  * When an issue is not NEW, an assignee must be provided.
	//  * A canonical_issue_id may not be set if the state is not DUPLICATE.
	//  * When an issue is in the VERIFIED state, it must have a verifier.
	// If these conditions are not met, the response will be an error.
	// Additional fields:
	//  * A provided created_time will be ignored, and a value from the
	//    actual time of creation/save will be determined by the API
	//    implementation.
	//  * user will be set to the id of the requesting user.  A provided
	//    value will be ignored.
	//  * reporter, if not set or if set to a blank email address, will be set
	//    to the id of the requesting user.
	//  * issue_comment can be set with initial comment.
	//  * attachments can be set with attachment metadata. The
	//    attachment_data_ref field in the response will contain the attachment
	//    resource name to pass to ByteStream.Write to upload the actual
	//    attachment data.
	Issue *Issue `protobuf:"bytes,1,opt,name=issue,proto3" json:"issue,omitempty"`
}

func (x *CreateIssueRequest) Reset() {
	*x = CreateIssueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateIssueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateIssueRequest) ProtoMessage() {}

func (x *CreateIssueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateIssueRequest.ProtoReflect.Descriptor instead.
func (*CreateIssueRequest) Descriptor() ([]byte, []int) {
	return file_google_devtools_issuetracker_v1_issuetracker_service_proto_rawDescGZIP(), []int{6}
}

func (x *CreateIssueRequest) GetIssue() *Issue {
	if x != nil {
		return x.Issue
	}
	return nil
}

// Request object for IssueTracker.ModifyIssue.
type ModifyIssueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Numeric ID of the issue to update.
	IssueId int64 `protobuf:"varint,1,opt,name=issue_id,json=issueId,proto3" json:"issue_id,omitempty"`
	// Field mask indicating what fields are set in the add IssueState.
	// Note that if any fields need to be set, this mask is required. If it is
	// not provided, the add IssueState will be ignored.
	AddMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=add_mask,json=addMask,proto3" json:"add_mask,omitempty"`
	// Fields present that are single value fields will replace the
	// current value. Fields that are collections will append the
	// values provided. Note that if any fields need to be set, add_mask above is
	// required. If it is not provided, the add IssueState will be ignored.
	Add *IssueState `protobuf:"bytes,3,opt,name=add,proto3" json:"add,omitempty"`
	// Field mask indicating what fields are set in the remove IssueState.
	// Note that if any fields need to be cleared, this mask is required. If it
	// is not provided, the remove IssueState will be ignored.
	RemoveMask *fieldmaskpb.FieldMask `protobuf:"bytes,4,opt,name=remove_mask,json=removeMask,proto3" json:"remove_mask,omitempty"`
	// Fields present that are single value fields with an empty value
	// will clear the current value. Fields that are collections will
	// remove the values provided. Note that if any fields need to be cleared,
	// remove_mask above is required. If it is not provided, the remove IssueState
	// will be ignored.
	Remove *IssueState `protobuf:"bytes,5,opt,name=remove,proto3" json:"remove,omitempty"`
	// A comment and/or attachments to append to this issue.
	IssueComment *IssueComment `protobuf:"bytes,6,opt,name=issue_comment,json=issueComment,proto3" json:"issue_comment,omitempty"`
}

func (x *ModifyIssueRequest) Reset() {
	*x = ModifyIssueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifyIssueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifyIssueRequest) ProtoMessage() {}

func (x *ModifyIssueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifyIssueRequest.ProtoReflect.Descriptor instead.
func (*ModifyIssueRequest) Descriptor() ([]byte, []int) {
	return file_google_devtools_issuetracker_v1_issuetracker_service_proto_rawDescGZIP(), []int{7}
}

func (x *ModifyIssueRequest) GetIssueId() int64 {
	if x != nil {
		return x.IssueId
	}
	return 0
}

func (x *ModifyIssueRequest) GetAddMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.AddMask
	}
	return nil
}

func (x *ModifyIssueRequest) GetAdd() *IssueState {
	if x != nil {
		return x.Add
	}
	return nil
}

func (x *ModifyIssueRequest) GetRemoveMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.RemoveMask
	}
	return nil
}

func (x *ModifyIssueRequest) GetRemove() *IssueState {
	if x != nil {
		return x.Remove
	}
	return nil
}

func (x *ModifyIssueRequest) GetIssueComment() *IssueComment {
	if x != nil {
		return x.IssueComment
	}
	return nil
}

// Request object for IssueTracker.CreateIssueRelationship.
type CreateIssueRelationshipRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Numeric ID of the source issue.
	IssueId int64 `protobuf:"varint,1,opt,name=issue_id,json=issueId,proto3" json:"issue_id,omitempty"`
	// Type of the relationship to create.
	RelationshipType IssueRelationshipType `protobuf:"varint,2,opt,name=relationship_type,json=relationshipType,proto3,enum=google.devtools.issuetracker.v1.IssueRelationshipType" json:"relationship_type,omitempty"`
	// Issue relationship to create. The target issue is inserted after
	// the issue specified by previous_target_issue_id. To insert to the top,
	// leave previous_target_issue_id as 0.
	IssueRelationship *IssueRelationship `protobuf:"bytes,3,opt,name=issue_relationship,json=issueRelationship,proto3" json:"issue_relationship,omitempty"`
}

func (x *CreateIssueRelationshipRequest) Reset() {
	*x = CreateIssueRelationshipRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateIssueRelationshipRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateIssueRelationshipRequest) ProtoMessage() {}

func (x *CreateIssueRelationshipRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateIssueRelationshipRequest.ProtoReflect.Descriptor instead.
func (*CreateIssueRelationshipRequest) Descriptor() ([]byte, []int) {
	return file_google_devtools_issuetracker_v1_issuetracker_service_proto_rawDescGZIP(), []int{8}
}

func (x *CreateIssueRelationshipRequest) GetIssueId() int64 {
	if x != nil {
		return x.IssueId
	}
	return 0
}

func (x *CreateIssueRelationshipRequest) GetRelationshipType() IssueRelationshipType {
	if x != nil {
		return x.RelationshipType
	}
	return IssueRelationshipType_ISSUE_RELATIONSHIP_TYPE_UNSPECIFIED
}

func (x *CreateIssueRelationshipRequest) GetIssueRelationship() *IssueRelationship {
	if x != nil {
		return x.IssueRelationship
	}
	return nil
}

// Request object for IssueTracker.ListIssueRelationships.
type ListIssueRelationshipsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Numeric ID of the source issue.
	IssueId int64 `protobuf:"varint,1,opt,name=issue_id,json=issueId,proto3" json:"issue_id,omitempty"`
	// Type of the relationship to list.
	RelationshipType IssueRelationshipType `protobuf:"varint,2,opt,name=relationship_type,json=relationshipType,proto3,enum=google.devtools.issuetracker.v1.IssueRelationshipType" json:"relationship_type,omitempty"`
	// The view of the issue to return.
	View IssueView `protobuf:"varint,3,opt,name=view,proto3,enum=google.devtools.issuetracker.v1.IssueView" json:"view,omitempty"`
	// Use specific atoms described here to refine results. By default, result
	// will not include issue relationships where target issue has been archived.
	// Query atoms:
	//   include_archived:[true|false]
	Query string `protobuf:"bytes,4,opt,name=query,proto3" json:"query,omitempty"`
}

func (x *ListIssueRelationshipsRequest) Reset() {
	*x = ListIssueRelationshipsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListIssueRelationshipsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListIssueRelationshipsRequest) ProtoMessage() {}

func (x *ListIssueRelationshipsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListIssueRelationshipsRequest.ProtoReflect.Descriptor instead.
func (*ListIssueRelationshipsRequest) Descriptor() ([]byte, []int) {
	return file_google_devtools_issuetracker_v1_issuetracker_service_proto_rawDescGZIP(), []int{9}
}

func (x *ListIssueRelationshipsRequest) GetIssueId() int64 {
	if x != nil {
		return x.IssueId
	}
	return 0
}

func (x *ListIssueRelationshipsRequest) GetRelationshipType() IssueRelationshipType {
	if x != nil {
		return x.RelationshipType
	}
	return IssueRelationshipType_ISSUE_RELATIONSHIP_TYPE_UNSPECIFIED
}

func (x *ListIssueRelationshipsRequest) GetView() IssueView {
	if x != nil {
		return x.View
	}
	return IssueView_ISSUE_VIEW_UNSPECIFIED
}

func (x *ListIssueRelationshipsRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

// Response object for IssueTracker.ListIssueRelationships.
type ListIssueRelationshipsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of issue relationships. All target issues are included regardless of
	// the caller's issue view access. target_issue_id is always set. target_issue
	// is set only if the caller has issue VIEW access to the target issue.
	IssueRelationships []*IssueRelationship `protobuf:"bytes,1,rep,name=issue_relationships,json=issueRelationships,proto3" json:"issue_relationships,omitempty"`
}

func (x *ListIssueRelationshipsResponse) Reset() {
	*x = ListIssueRelationshipsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListIssueRelationshipsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListIssueRelationshipsResponse) ProtoMessage() {}

func (x *ListIssueRelationshipsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListIssueRelationshipsResponse.ProtoReflect.Descriptor instead.
func (*ListIssueRelationshipsResponse) Descriptor() ([]byte, []int) {
	return file_google_devtools_issuetracker_v1_issuetracker_service_proto_rawDescGZIP(), []int{10}
}

func (x *ListIssueRelationshipsResponse) GetIssueRelationships() []*IssueRelationship {
	if x != nil {
		return x.IssueRelationships
	}
	return nil
}

// Request object for IssueTracker.ListIssueUpdates.
type ListIssueUpdatesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Numeric ID of the issue.
	IssueId int64 `protobuf:"varint,1,opt,name=issue_id,json=issueId,proto3" json:"issue_id,omitempty"`
	// Updates are sorted by version. Can be specified as either
	// ASC or DESC. Default is DESC.
	SortBy string `protobuf:"bytes,2,opt,name=sort_by,json=sortBy,proto3" json:"sort_by,omitempty"`
	// If 0, the full set of IssueUpdates will be returned. Default 0.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Pagination token. Optional.
	PageToken string `protobuf:"bytes,4,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// Issue comment view. When unspecified, it will default to BASIC.
	IssueCommentView IssueCommentView `protobuf:"varint,5,opt,name=issue_comment_view,json=issueCommentView,proto3,enum=google.devtools.issuetracker.v1.IssueCommentView" json:"issue_comment_view,omitempty"`
}

func (x *ListIssueUpdatesRequest) Reset() {
	*x = ListIssueUpdatesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListIssueUpdatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListIssueUpdatesRequest) ProtoMessage() {}

func (x *ListIssueUpdatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListIssueUpdatesRequest.ProtoReflect.Descriptor instead.
func (*ListIssueUpdatesRequest) Descriptor() ([]byte, []int) {
	return file_google_devtools_issuetracker_v1_issuetracker_service_proto_rawDescGZIP(), []int{11}
}

func (x *ListIssueUpdatesRequest) GetIssueId() int64 {
	if x != nil {
		return x.IssueId
	}
	return 0
}

func (x *ListIssueUpdatesRequest) GetSortBy() string {
	if x != nil {
		return x.SortBy
	}
	return ""
}

func (x *ListIssueUpdatesRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListIssueUpdatesRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListIssueUpdatesRequest) GetIssueCommentView() IssueCommentView {
	if x != nil {
		return x.IssueCommentView
	}
	return IssueCommentView_ISSUE_COMMENT_VIEW_UNSPECIFIED
}

// Response object for IssueTracker.ListIssueUpdates.
type ListIssueUpdatesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The current page of IssueUpdates.
	IssueUpdates []*IssueUpdate `protobuf:"bytes,1,rep,name=issue_updates,json=issueUpdates,proto3" json:"issue_updates,omitempty"`
	// Pagination token for next page of results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	// Total number of results.
	TotalSize int32 `protobuf:"varint,3,opt,name=total_size,json=totalSize,proto3" json:"total_size,omitempty"`
}

func (x *ListIssueUpdatesResponse) Reset() {
	*x = ListIssueUpdatesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListIssueUpdatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListIssueUpdatesResponse) ProtoMessage() {}

func (x *ListIssueUpdatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListIssueUpdatesResponse.ProtoReflect.Descriptor instead.
func (*ListIssueUpdatesResponse) Descriptor() ([]byte, []int) {
	return file_google_devtools_issuetracker_v1_issuetracker_service_proto_rawDescGZIP(), []int{12}
}

func (x *ListIssueUpdatesResponse) GetIssueUpdates() []*IssueUpdate {
	if x != nil {
		return x.IssueUpdates
	}
	return nil
}

func (x *ListIssueUpdatesResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

func (x *ListIssueUpdatesResponse) GetTotalSize() int32 {
	if x != nil {
		return x.TotalSize
	}
	return 0
}

// Request object for IssueTracker.ListIssueComments.
type ListIssueCommentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Numeric ID of the issue.
	IssueId int64 `protobuf:"varint,1,opt,name=issue_id,json=issueId,proto3" json:"issue_id,omitempty"`
	// Comments are sorted by created_time. Can be specified as either
	// ASC or DESC. Default is DESC.
	SortBy string `protobuf:"bytes,2,opt,name=sort_by,json=sortBy,proto3" json:"sort_by,omitempty"`
	// The number of results to return. Default 25.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Pagination token. Optional.
	PageToken string `protobuf:"bytes,4,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// Issue comment view. When unspecified, it will default to BASIC.
	IssueCommentView IssueCommentView `protobuf:"varint,5,opt,name=issue_comment_view,json=issueCommentView,proto3,enum=google.devtools.issuetracker.v1.IssueCommentView" json:"issue_comment_view,omitempty"`
}

func (x *ListIssueCommentsRequest) Reset() {
	*x = ListIssueCommentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListIssueCommentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListIssueCommentsRequest) ProtoMessage() {}

func (x *ListIssueCommentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListIssueCommentsRequest.ProtoReflect.Descriptor instead.
func (*ListIssueCommentsRequest) Descriptor() ([]byte, []int) {
	return file_google_devtools_issuetracker_v1_issuetracker_service_proto_rawDescGZIP(), []int{13}
}

func (x *ListIssueCommentsRequest) GetIssueId() int64 {
	if x != nil {
		return x.IssueId
	}
	return 0
}

func (x *ListIssueCommentsRequest) GetSortBy() string {
	if x != nil {
		return x.SortBy
	}
	return ""
}

func (x *ListIssueCommentsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListIssueCommentsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListIssueCommentsRequest) GetIssueCommentView() IssueCommentView {
	if x != nil {
		return x.IssueCommentView
	}
	return IssueCommentView_ISSUE_COMMENT_VIEW_UNSPECIFIED
}

// Response object for IssueTracker.ListIssueComments.
type ListIssueCommentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The current page of IssueComments.
	IssueComments []*IssueComment `protobuf:"bytes,1,rep,name=issue_comments,json=issueComments,proto3" json:"issue_comments,omitempty"`
	// Pagination token for next page of results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	// Total number of results.
	TotalSize int32 `protobuf:"varint,3,opt,name=total_size,json=totalSize,proto3" json:"total_size,omitempty"`
}

func (x *ListIssueCommentsResponse) Reset() {
	*x = ListIssueCommentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListIssueCommentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListIssueCommentsResponse) ProtoMessage() {}

func (x *ListIssueCommentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListIssueCommentsResponse.ProtoReflect.Descriptor instead.
func (*ListIssueCommentsResponse) Descriptor() ([]byte, []int) {
	return file_google_devtools_issuetracker_v1_issuetracker_service_proto_rawDescGZIP(), []int{14}
}

func (x *ListIssueCommentsResponse) GetIssueComments() []*IssueComment {
	if x != nil {
		return x.IssueComments
	}
	return nil
}

func (x *ListIssueCommentsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

func (x *ListIssueCommentsResponse) GetTotalSize() int32 {
	if x != nil {
		return x.TotalSize
	}
	return 0
}

// Request object for IssueTracker.CreateIssueComment. Only new text comment
// will be created.
type CreateIssueCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Numeric ID of the issue.
	IssueId int64 `protobuf:"varint,1,opt,name=issue_id,json=issueId,proto3" json:"issue_id,omitempty"`
	// Contains the comment text
	Comment *IssueComment `protobuf:"bytes,2,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *CreateIssueCommentRequest) Reset() {
	*x = CreateIssueCommentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateIssueCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateIssueCommentRequest) ProtoMessage() {}

func (x *CreateIssueCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateIssueCommentRequest.ProtoReflect.Descriptor instead.
func (*CreateIssueCommentRequest) Descriptor() ([]byte, []int) {
	return file_google_devtools_issuetracker_v1_issuetracker_service_proto_rawDescGZIP(), []int{15}
}

func (x *CreateIssueCommentRequest) GetIssueId() int64 {
	if x != nil {
		return x.IssueId
	}
	return 0
}

func (x *CreateIssueCommentRequest) GetComment() *IssueComment {
	if x != nil {
		return x.Comment
	}
	return nil
}

// Request object for IssueTracker.ListAttachments.
type ListAttachmentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Numeric ID of the issue.
	IssueId int64 `protobuf:"varint,1,opt,name=issue_id,json=issueId,proto3" json:"issue_id,omitempty"`
}

func (x *ListAttachmentsRequest) Reset() {
	*x = ListAttachmentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[16]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAttachmentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAttachmentsRequest) ProtoMessage() {}

func (x *ListAttachmentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[16]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAttachmentsRequest.ProtoReflect.Descriptor instead.
func (*ListAttachmentsRequest) Descriptor() ([]byte, []int) {
	return file_google_devtools_issuetracker_v1_issuetracker_service_proto_rawDescGZIP(), []int{16}
}

func (x *ListAttachmentsRequest) GetIssueId() int64 {
	if x != nil {
		return x.IssueId
	}
	return 0
}

// Response object for IssueTracker.ListAttachments.
type ListAttachmentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Metadata of attachments on the specified issue.
	Attachments []*Attachment `protobuf:"bytes,1,rep,name=attachments,proto3" json:"attachments,omitempty"`
}

func (x *ListAttachmentsResponse) Reset() {
	*x = ListAttachmentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[17]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAttachmentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAttachmentsResponse) ProtoMessage() {}

func (x *ListAttachmentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[17]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAttachmentsResponse.ProtoReflect.Descriptor instead.
func (*ListAttachmentsResponse) Descriptor() ([]byte, []int) {
	return file_google_devtools_issuetracker_v1_issuetracker_service_proto_rawDescGZIP(), []int{17}
}

func (x *ListAttachmentsResponse) GetAttachments() []*Attachment {
	if x != nil {
		return x.Attachments
	}
	return nil
}

// Request object for IssueTracker.CreateHotlistEntry
// Used to add an issue to a hotlist
type CreateHotlistEntryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Numeric ID of the hotlist.
	HotlistId int64 `protobuf:"varint,1,opt,name=hotlist_id,json=hotlistId,proto3" json:"hotlist_id,omitempty"`
	// HotlistEntry position will not be taken into account. Issues are always
	// prepended to the hotlist and have position = 1. If the issue is already
	// present on the hotlist it will be moved to position = 1.
	HotlistEntry *HotlistEntry `protobuf:"bytes,2,opt,name=hotlist_entry,json=hotlistEntry,proto3" json:"hotlist_entry,omitempty"`
}

func (x *CreateHotlistEntryRequest) Reset() {
	*x = CreateHotlistEntryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[18]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateHotlistEntryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHotlistEntryRequest) ProtoMessage() {}

func (x *CreateHotlistEntryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[18]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHotlistEntryRequest.ProtoReflect.Descriptor instead.
func (*CreateHotlistEntryRequest) Descriptor() ([]byte, []int) {
	return file_google_devtools_issuetracker_v1_issuetracker_service_proto_rawDescGZIP(), []int{18}
}

func (x *CreateHotlistEntryRequest) GetHotlistId() int64 {
	if x != nil {
		return x.HotlistId
	}
	return 0
}

func (x *CreateHotlistEntryRequest) GetHotlistEntry() *HotlistEntry {
	if x != nil {
		return x.HotlistEntry
	}
	return nil
}

// Request object for IssueTracker.DeleteHotlistEntry
// Used to delete an issue from a hotlist
type DeleteHotlistEntryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Numeric ID of the hotlist.
	HotlistId int64 `protobuf:"varint,1,opt,name=hotlist_id,json=hotlistId,proto3" json:"hotlist_id,omitempty"`
	// Numeric ID of the issue to remove from hotlist
	IssueId int64 `protobuf:"varint,2,opt,name=issue_id,json=issueId,proto3" json:"issue_id,omitempty"`
}

func (x *DeleteHotlistEntryRequest) Reset() {
	*x = DeleteHotlistEntryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[19]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteHotlistEntryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteHotlistEntryRequest) ProtoMessage() {}

func (x *DeleteHotlistEntryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[19]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteHotlistEntryRequest.ProtoReflect.Descriptor instead.
func (*DeleteHotlistEntryRequest) Descriptor() ([]byte, []int) {
	return file_google_devtools_issuetracker_v1_issuetracker_service_proto_rawDescGZIP(), []int{19}
}

func (x *DeleteHotlistEntryRequest) GetHotlistId() int64 {
	if x != nil {
		return x.HotlistId
	}
	return 0
}

func (x *DeleteHotlistEntryRequest) GetIssueId() int64 {
	if x != nil {
		return x.IssueId
	}
	return 0
}

var File_google_devtools_issuetracker_v1_issuetracker_service_proto protoreflect.FileDescriptor

var file_google_devtools_issuetracker_v1_issuetracker_service_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c,
	0x73, 0x2f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x69, 0x73,
	0x73, 0x75, 0x65, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x64, 0x65,
	0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x74, 0x72, 0x61, 0x63,
	0x6b, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x74, 0x72, 0x61, 0x63,
	0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61,
	0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x38, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x22, 0xc0, 0x01, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x19,
	0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x3e, 0x0a, 0x04, 0x76, 0x69, 0x65, 0x77, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76,
	0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x65, 0x74, 0x72, 0x61, 0x63, 0x6b,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x56, 0x69, 0x65, 0x77, 0x52,
	0x04, 0x76, 0x69, 0x65, 0x77, 0x22, 0x9b, 0x01, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x73,
	0x73, 0x75, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x06,
	0x69, 0x73, 0x73, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x69,
	0x73, 0x73, 0x75, 0x65, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49,
	0x73, 0x73, 0x75, 0x65, 0x52, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0f,
	0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53,
	0x69, 0x7a, 0x65, 0x22, 0x74, 0x0a, 0x15, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x49,
	0x73, 0x73, 0x75, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x69, 0x73, 0x73, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52,
	0x08, 0x69, 0x73, 0x73, 0x75, 0x65, 0x49, 0x64, 0x73, 0x12, 0x3e, 0x0a, 0x04, 0x76, 0x69, 0x65,
	0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x65, 0x74,
	0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x56,
	0x69, 0x65, 0x77, 0x52, 0x04, 0x76, 0x69, 0x65, 0x77, 0x22, 0x58, 0x0a, 0x16, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x47, 0x65, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76,
	0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x65, 0x74, 0x72, 0x61, 0x63, 0x6b,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x06, 0x69, 0x73, 0x73,
	0x75, 0x65, 0x73, 0x22, 0x6c, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x69, 0x73, 0x73, 0x75, 0x65, 0x49,
	0x64, 0x12, 0x3e, 0x0a, 0x04, 0x76, 0x69, 0x65, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c,
	0x73, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x65, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x56, 0x69, 0x65, 0x77, 0x52, 0x04, 0x76, 0x69, 0x65,
	0x77, 0x22, 0x52, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x73, 0x73, 0x75, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x05, 0x69, 0x73, 0x73, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x65, 0x74, 0x72,
	0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x05,
	0x69, 0x73, 0x73, 0x75, 0x65, 0x22, 0xfb, 0x02, 0x0a, 0x12, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79,
	0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08,
	0x69, 0x73, 0x73, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x69, 0x73, 0x73, 0x75, 0x65, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x08, 0x61, 0x64, 0x64, 0x5f, 0x6d,
	0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x07, 0x61, 0x64, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x3d,
	0x0a, 0x03, 0x61, 0x64, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x69, 0x73,
	0x73, 0x75, 0x65, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x73,
	0x73, 0x75, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x03, 0x61, 0x64, 0x64, 0x12, 0x3b, 0x0a,
	0x0b, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a,
	0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x43, 0x0a, 0x06, 0x72, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x69, 0x73, 0x73,
	0x75, 0x65, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x73, 0x73,
	0x75, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12,
	0x52, 0x0a, 0x0d, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x65, 0x74, 0x72,
	0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0c, 0x69, 0x73, 0x73, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x22, 0x83, 0x02, 0x0a, 0x1e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x73,
	0x73, 0x75, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x69, 0x73, 0x73, 0x75, 0x65, 0x49,
	0x64, 0x12, 0x63, 0x0a, 0x11, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69,
	0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x36, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x69,
	0x73, 0x73, 0x75, 0x65, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49,
	0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x10, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68,
	0x69, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x61, 0x0a, 0x12, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f,
	0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x32, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74,
	0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x65, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x52, 0x11, 0x69, 0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x22, 0xf5, 0x01, 0x0a, 0x1d, 0x4c, 0x69,
	0x73, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x68, 0x69, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x69,
	0x73, 0x73, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x69,
	0x73, 0x73, 0x75, 0x65, 0x49, 0x64, 0x12, 0x63, 0x0a, 0x11, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x36, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f,
	0x6f, 0x6c, 0x73, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x65, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x68, 0x69, 0x70, 0x54, 0x79, 0x70, 0x65, 0x52, 0x10, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3e, 0x0a, 0x04, 0x76,
	0x69, 0x65, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x69, 0x73, 0x73, 0x75,
	0x65, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x73, 0x73, 0x75,
	0x65, 0x56, 0x69, 0x65, 0x77, 0x52, 0x04, 0x76, 0x69, 0x65, 0x77, 0x12, 0x14, 0x0a, 0x05, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x22, 0x85, 0x01, 0x0a, 0x1e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x63, 0x0a, 0x13, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x32, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f,
	0x6f, 0x6c, 0x73, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x65, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x68, 0x69, 0x70, 0x52, 0x12, 0x69, 0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x73, 0x22, 0xea, 0x01, 0x0a, 0x17, 0x4c, 0x69,
	0x73, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x69, 0x73, 0x73, 0x75, 0x65, 0x49, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x62, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x5f, 0x0a, 0x12, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x31, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f,
	0x6f, 0x6c, 0x73, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x65, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x56, 0x69, 0x65, 0x77, 0x52, 0x10, 0x69, 0x73, 0x73, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x56, 0x69, 0x65, 0x77, 0x22, 0xb4, 0x01, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x49,
	0x73, 0x73, 0x75, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x0d, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x69, 0x73, 0x73,
	0x75, 0x65, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x73, 0x73,
	0x75, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x0c, 0x69, 0x73, 0x73, 0x75, 0x65, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d,
	0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x69, 0x7a, 0x65, 0x22, 0xeb, 0x01,
	0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73,
	0x73, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x69, 0x73,
	0x73, 0x75, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x62, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x5f, 0x0a, 0x12, 0x69, 0x73,
	0x73, 0x75, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x69, 0x65, 0x77,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x31, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x65, 0x74, 0x72,
	0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x69, 0x65, 0x77, 0x52, 0x10, 0x69, 0x73, 0x73, 0x75, 0x65,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x69, 0x65, 0x77, 0x22, 0xb8, 0x01, 0x0a, 0x19,
	0x4c, 0x69, 0x73, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a, 0x0e, 0x69, 0x73, 0x73,
	0x75, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f,
	0x6f, 0x6c, 0x73, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x65, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x0d, 0x69, 0x73, 0x73, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61,
	0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x7f, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x49, 0x73, 0x73, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x69, 0x73, 0x73, 0x75, 0x65, 0x49, 0x64, 0x12, 0x47,
	0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c,
	0x73, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x65, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x33, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x41,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x69, 0x73, 0x73, 0x75, 0x65, 0x49, 0x64, 0x22, 0x68, 0x0a, 0x17,
	0x4c, 0x69, 0x73, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0b, 0x61, 0x74, 0x74, 0x61, 0x63,
	0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x69,
	0x73, 0x73, 0x75, 0x65, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x61, 0x74, 0x74, 0x61, 0x63,
	0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x8e, 0x01, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x48, 0x6f, 0x74, 0x6c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x6f, 0x74, 0x6c, 0x69, 0x73, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x68, 0x6f, 0x74, 0x6c, 0x69, 0x73,
	0x74, 0x49, 0x64, 0x12, 0x52, 0x0a, 0x0d, 0x68, 0x6f, 0x74, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x65,
	0x6e, 0x74, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x69, 0x73, 0x73,
	0x75, 0x65, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x6f, 0x74,
	0x6c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x68, 0x6f, 0x74, 0x6c, 0x69,
	0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x55, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x48, 0x6f, 0x74, 0x6c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x6f, 0x74, 0x6c, 0x69, 0x73, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x68, 0x6f, 0x74, 0x6c, 0x69, 0x73,
	0x74, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x69, 0x73, 0x73, 0x75, 0x65, 0x49, 0x64, 0x2a, 0x3c,
	0x0a, 0x09, 0x49, 0x73, 0x73, 0x75, 0x65, 0x56, 0x69, 0x65, 0x77, 0x12, 0x1a, 0x0a, 0x16, 0x49,
	0x53, 0x53, 0x55, 0x45, 0x5f, 0x56, 0x49, 0x45, 0x57, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x42, 0x41, 0x53, 0x49, 0x43,
	0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x55, 0x4c, 0x4c, 0x10, 0x02, 0x2a, 0x71, 0x0a, 0x10,
	0x49, 0x73, 0x73, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x69, 0x65, 0x77,
	0x12, 0x22, 0x0a, 0x1e, 0x49, 0x53, 0x53, 0x55, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x45, 0x4e,
	0x54, 0x5f, 0x56, 0x49, 0x45, 0x57, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18, 0x49, 0x53, 0x53, 0x55, 0x45, 0x5f, 0x43, 0x4f,
	0x4d, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x56, 0x49, 0x45, 0x57, 0x5f, 0x42, 0x41, 0x53, 0x49, 0x43,
	0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x49, 0x53, 0x53, 0x55, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x4d,
	0x45, 0x4e, 0x54, 0x5f, 0x56, 0x49, 0x45, 0x57, 0x5f, 0x46, 0x55, 0x4c, 0x4c, 0x10, 0x02, 0x32,
	0xfd, 0x12, 0x0a, 0x0c, 0x49, 0x73, 0x73, 0x75, 0x65, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72,
	0x12, 0x99, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x12, 0x34, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f,
	0x6f, 0x6c, 0x73, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x65, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x65, 0x74,
	0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x12, 0x1f, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x3d, 0x2a, 0x7d, 0x12, 0x89, 0x01, 0x0a,
	0x0a, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x73, 0x12, 0x32, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x69, 0x73,
	0x73, 0x75, 0x65, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x33, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c,
	0x73, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x65, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x76,
	0x31, 0x2f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x73, 0x12, 0x9e, 0x01, 0x0a, 0x0e, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x47, 0x65, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x73, 0x12, 0x36, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x69, 0x73,
	0x73, 0x75, 0x65, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76,
	0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x65, 0x74, 0x72, 0x61, 0x63, 0x6b,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x49, 0x73,
	0x73, 0x75, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x73,
	0x3a, 0x62, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x12, 0x85, 0x01, 0x0a, 0x08, 0x47, 0x65,
	0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x12, 0x30, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x65, 0x74, 0x72,
	0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x73, 0x73, 0x75,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x65,
	0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65,
	0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x73,
	0x73, 0x75, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x3d, 0x2a,
	0x7d, 0x12, 0x85, 0x01, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x73, 0x73, 0x75,
	0x65, 0x12, 0x33, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f,
	0x6f, 0x6c, 0x73, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x65, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x65, 0x74, 0x72,
	0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x22, 0x19,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x22, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x73, 0x73, 0x75,
	0x65, 0x73, 0x3a, 0x05, 0x69, 0x73, 0x73, 0x75, 0x65, 0x12, 0x95, 0x01, 0x0a, 0x0b, 0x4d, 0x6f,
	0x64, 0x69, 0x66, 0x79, 0x49, 0x73, 0x73, 0x75, 0x65, 0x12, 0x33, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x69, 0x73, 0x73, 0x75,
	0x65, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x64, 0x69,
	0x66, 0x79, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73,
	0x2e, 0x69, 0x73, 0x73, 0x75, 0x65, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x22, 0x1e,
	0x2f, 0x76, 0x31, 0x2f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x73, 0x73, 0x75,
	0x65, 0x5f, 0x69, 0x64, 0x3d, 0x2a, 0x7d, 0x3a, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x3a, 0x01,
	0x2a, 0x12, 0xd1, 0x01, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x73, 0x73, 0x75,
	0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x12, 0x3f, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e,
	0x69, 0x73, 0x73, 0x75, 0x65, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73,
	0x2e, 0x69, 0x73, 0x73, 0x75, 0x65, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68,
	0x69, 0x70, 0x22, 0x41, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3b, 0x22, 0x25, 0x2f, 0x76, 0x31, 0x2f,
	0x69, 0x73, 0x73, 0x75, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f, 0x69, 0x64,
	0x3d, 0x2a, 0x7d, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70,
	0x73, 0x3a, 0x12, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x68, 0x69, 0x70, 0x12, 0xc8, 0x01, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x73,
	0x73, 0x75, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x73,
	0x12, 0x3e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f,
	0x6c, 0x73, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x65, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x3f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f,
	0x6c, 0x73, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x65, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x2d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x12, 0x25, 0x2f, 0x76, 0x31, 0x2f, 0x69,
	0x73, 0x73, 0x75, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x3d,
	0x2a, 0x7d, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x73,
	0x12, 0xb5, 0x01, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x73, 0x12, 0x38, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64,
	0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x65, 0x74, 0x72, 0x61,
	0x63, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x73, 0x73, 0x75,
	0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x39, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c,
	0x73, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x65, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2c, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x26, 0x12, 0x24, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x73, 0x2f, 0x7b,
	0x69, 0x73, 0x73, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x3d, 0x2a, 0x7d, 0x2f, 0x69, 0x73, 0x73, 0x75,
	0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x12, 0xb2, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x49, 0x73, 0x73, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x3a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c,
	0x73, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x65, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x73, 0x73, 0x75, 0x65, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x69, 0x73,
	0x73, 0x75, 0x65, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x73,
	0x73, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x31, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x2b, 0x22, 0x20, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x73, 0x2f, 0x7b,
	0x69, 0x73, 0x73, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x3d, 0x2a, 0x7d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x3a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0xb4, 0x01,
	0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x39, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76,
	0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x65, 0x74, 0x72, 0x61, 0x63, 0x6b,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73,
	0x2e, 0x69, 0x73, 0x73, 0x75, 0x65, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x22, 0x12, 0x20, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x73, 0x2f, 0x7b,
	0x69, 0x73, 0x73, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x3d, 0x2a, 0x7d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0xb1, 0x01, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x37, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x65,
	0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x38, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f,
	0x6f, 0x6c, 0x73, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x65, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2b, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x25, 0x12, 0x23, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x73, 0x2f,
	0x7b, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x3d, 0x2a, 0x7d, 0x2f, 0x61, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0xaf, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x48, 0x6f, 0x74, 0x6c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x3a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c,
	0x73, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x65, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x74, 0x6c, 0x69, 0x73, 0x74, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x69, 0x73,
	0x73, 0x75, 0x65, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x6f,
	0x74, 0x6c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x2e, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x28, 0x1a, 0x23, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x6f, 0x74, 0x6c, 0x69, 0x73, 0x74, 0x73,
	0x2f, 0x7b, 0x68, 0x6f, 0x74, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x3d, 0x2a, 0x7d, 0x2f,
	0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0xa2, 0x01, 0x0a, 0x12, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x6f, 0x74, 0x6c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x3a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f,
	0x6f, 0x6c, 0x73, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x65, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x6f, 0x74, 0x6c, 0x69, 0x73,
	0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x38, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x32, 0x2a, 0x30, 0x2f,
	0x76, 0x31, 0x2f, 0x68, 0x6f, 0x74, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x2f, 0x7b, 0x68, 0x6f, 0x74,
	0x6c, 0x69, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x3d, 0x2a, 0x7d, 0x2f, 0x65, 0x6e, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x2f, 0x7b, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x3d, 0x2a, 0x7d, 0x1a,
	0x2e, 0xca, 0x41, 0x2b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x2d,
	0x69, 0x73, 0x73, 0x75, 0x65, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2d, 0x63, 0x32, 0x70,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x42,
	0x8e, 0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64,
	0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x65, 0x74, 0x72, 0x61,
	0x63, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x18, 0x49, 0x73, 0x73, 0x75, 0x65, 0x54, 0x72,
	0x61, 0x63, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x64, 0x65, 0x76, 0x74, 0x6f,
	0x6f, 0x6c, 0x73, 0x2f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x3b, 0x69, 0x73, 0x73, 0x75, 0x65, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_devtools_issuetracker_v1_issuetracker_service_proto_rawDescOnce sync.Once
	file_google_devtools_issuetracker_v1_issuetracker_service_proto_rawDescData = file_google_devtools_issuetracker_v1_issuetracker_service_proto_rawDesc
)

func file_google_devtools_issuetracker_v1_issuetracker_service_proto_rawDescGZIP() []byte {
	file_google_devtools_issuetracker_v1_issuetracker_service_proto_rawDescOnce.Do(func() {
		file_google_devtools_issuetracker_v1_issuetracker_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_devtools_issuetracker_v1_issuetracker_service_proto_rawDescData)
	})
	return file_google_devtools_issuetracker_v1_issuetracker_service_proto_rawDescData
}

var file_google_devtools_issuetracker_v1_issuetracker_service_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes = make([]protoimpl.MessageInfo, 20)
var file_google_devtools_issuetracker_v1_issuetracker_service_proto_goTypes = []interface{}{
	(IssueView)(0),                         // 0: google.devtools.issuetracker.v1.IssueView
	(IssueCommentView)(0),                  // 1: google.devtools.issuetracker.v1.IssueCommentView
	(*GetComponentRequest)(nil),            // 2: google.devtools.issuetracker.v1.GetComponentRequest
	(*ListIssuesRequest)(nil),              // 3: google.devtools.issuetracker.v1.ListIssuesRequest
	(*ListIssuesResponse)(nil),             // 4: google.devtools.issuetracker.v1.ListIssuesResponse
	(*BatchGetIssuesRequest)(nil),          // 5: google.devtools.issuetracker.v1.BatchGetIssuesRequest
	(*BatchGetIssuesResponse)(nil),         // 6: google.devtools.issuetracker.v1.BatchGetIssuesResponse
	(*GetIssueRequest)(nil),                // 7: google.devtools.issuetracker.v1.GetIssueRequest
	(*CreateIssueRequest)(nil),             // 8: google.devtools.issuetracker.v1.CreateIssueRequest
	(*ModifyIssueRequest)(nil),             // 9: google.devtools.issuetracker.v1.ModifyIssueRequest
	(*CreateIssueRelationshipRequest)(nil), // 10: google.devtools.issuetracker.v1.CreateIssueRelationshipRequest
	(*ListIssueRelationshipsRequest)(nil),  // 11: google.devtools.issuetracker.v1.ListIssueRelationshipsRequest
	(*ListIssueRelationshipsResponse)(nil), // 12: google.devtools.issuetracker.v1.ListIssueRelationshipsResponse
	(*ListIssueUpdatesRequest)(nil),        // 13: google.devtools.issuetracker.v1.ListIssueUpdatesRequest
	(*ListIssueUpdatesResponse)(nil),       // 14: google.devtools.issuetracker.v1.ListIssueUpdatesResponse
	(*ListIssueCommentsRequest)(nil),       // 15: google.devtools.issuetracker.v1.ListIssueCommentsRequest
	(*ListIssueCommentsResponse)(nil),      // 16: google.devtools.issuetracker.v1.ListIssueCommentsResponse
	(*CreateIssueCommentRequest)(nil),      // 17: google.devtools.issuetracker.v1.CreateIssueCommentRequest
	(*ListAttachmentsRequest)(nil),         // 18: google.devtools.issuetracker.v1.ListAttachmentsRequest
	(*ListAttachmentsResponse)(nil),        // 19: google.devtools.issuetracker.v1.ListAttachmentsResponse
	(*CreateHotlistEntryRequest)(nil),      // 20: google.devtools.issuetracker.v1.CreateHotlistEntryRequest
	(*DeleteHotlistEntryRequest)(nil),      // 21: google.devtools.issuetracker.v1.DeleteHotlistEntryRequest
	(*Issue)(nil),                          // 22: google.devtools.issuetracker.v1.Issue
	(*fieldmaskpb.FieldMask)(nil),          // 23: google.protobuf.FieldMask
	(*IssueState)(nil),                     // 24: google.devtools.issuetracker.v1.IssueState
	(*IssueComment)(nil),                   // 25: google.devtools.issuetracker.v1.IssueComment
	(IssueRelationshipType)(0),             // 26: google.devtools.issuetracker.v1.IssueRelationshipType
	(*IssueRelationship)(nil),              // 27: google.devtools.issuetracker.v1.IssueRelationship
	(*IssueUpdate)(nil),                    // 28: google.devtools.issuetracker.v1.IssueUpdate
	(*Attachment)(nil),                     // 29: google.devtools.issuetracker.v1.Attachment
	(*HotlistEntry)(nil),                   // 30: google.devtools.issuetracker.v1.HotlistEntry
	(*Component)(nil),                      // 31: google.devtools.issuetracker.v1.Component
	(*emptypb.Empty)(nil),                  // 32: google.protobuf.Empty
}
var file_google_devtools_issuetracker_v1_issuetracker_service_proto_depIdxs = []int32{
	0,  // 0: google.devtools.issuetracker.v1.ListIssuesRequest.view:type_name -> google.devtools.issuetracker.v1.IssueView
	22, // 1: google.devtools.issuetracker.v1.ListIssuesResponse.issues:type_name -> google.devtools.issuetracker.v1.Issue
	0,  // 2: google.devtools.issuetracker.v1.BatchGetIssuesRequest.view:type_name -> google.devtools.issuetracker.v1.IssueView
	22, // 3: google.devtools.issuetracker.v1.BatchGetIssuesResponse.issues:type_name -> google.devtools.issuetracker.v1.Issue
	0,  // 4: google.devtools.issuetracker.v1.GetIssueRequest.view:type_name -> google.devtools.issuetracker.v1.IssueView
	22, // 5: google.devtools.issuetracker.v1.CreateIssueRequest.issue:type_name -> google.devtools.issuetracker.v1.Issue
	23, // 6: google.devtools.issuetracker.v1.ModifyIssueRequest.add_mask:type_name -> google.protobuf.FieldMask
	24, // 7: google.devtools.issuetracker.v1.ModifyIssueRequest.add:type_name -> google.devtools.issuetracker.v1.IssueState
	23, // 8: google.devtools.issuetracker.v1.ModifyIssueRequest.remove_mask:type_name -> google.protobuf.FieldMask
	24, // 9: google.devtools.issuetracker.v1.ModifyIssueRequest.remove:type_name -> google.devtools.issuetracker.v1.IssueState
	25, // 10: google.devtools.issuetracker.v1.ModifyIssueRequest.issue_comment:type_name -> google.devtools.issuetracker.v1.IssueComment
	26, // 11: google.devtools.issuetracker.v1.CreateIssueRelationshipRequest.relationship_type:type_name -> google.devtools.issuetracker.v1.IssueRelationshipType
	27, // 12: google.devtools.issuetracker.v1.CreateIssueRelationshipRequest.issue_relationship:type_name -> google.devtools.issuetracker.v1.IssueRelationship
	26, // 13: google.devtools.issuetracker.v1.ListIssueRelationshipsRequest.relationship_type:type_name -> google.devtools.issuetracker.v1.IssueRelationshipType
	0,  // 14: google.devtools.issuetracker.v1.ListIssueRelationshipsRequest.view:type_name -> google.devtools.issuetracker.v1.IssueView
	27, // 15: google.devtools.issuetracker.v1.ListIssueRelationshipsResponse.issue_relationships:type_name -> google.devtools.issuetracker.v1.IssueRelationship
	1,  // 16: google.devtools.issuetracker.v1.ListIssueUpdatesRequest.issue_comment_view:type_name -> google.devtools.issuetracker.v1.IssueCommentView
	28, // 17: google.devtools.issuetracker.v1.ListIssueUpdatesResponse.issue_updates:type_name -> google.devtools.issuetracker.v1.IssueUpdate
	1,  // 18: google.devtools.issuetracker.v1.ListIssueCommentsRequest.issue_comment_view:type_name -> google.devtools.issuetracker.v1.IssueCommentView
	25, // 19: google.devtools.issuetracker.v1.ListIssueCommentsResponse.issue_comments:type_name -> google.devtools.issuetracker.v1.IssueComment
	25, // 20: google.devtools.issuetracker.v1.CreateIssueCommentRequest.comment:type_name -> google.devtools.issuetracker.v1.IssueComment
	29, // 21: google.devtools.issuetracker.v1.ListAttachmentsResponse.attachments:type_name -> google.devtools.issuetracker.v1.Attachment
	30, // 22: google.devtools.issuetracker.v1.CreateHotlistEntryRequest.hotlist_entry:type_name -> google.devtools.issuetracker.v1.HotlistEntry
	2,  // 23: google.devtools.issuetracker.v1.IssueTracker.GetComponent:input_type -> google.devtools.issuetracker.v1.GetComponentRequest
	3,  // 24: google.devtools.issuetracker.v1.IssueTracker.ListIssues:input_type -> google.devtools.issuetracker.v1.ListIssuesRequest
	5,  // 25: google.devtools.issuetracker.v1.IssueTracker.BatchGetIssues:input_type -> google.devtools.issuetracker.v1.BatchGetIssuesRequest
	7,  // 26: google.devtools.issuetracker.v1.IssueTracker.GetIssue:input_type -> google.devtools.issuetracker.v1.GetIssueRequest
	8,  // 27: google.devtools.issuetracker.v1.IssueTracker.CreateIssue:input_type -> google.devtools.issuetracker.v1.CreateIssueRequest
	9,  // 28: google.devtools.issuetracker.v1.IssueTracker.ModifyIssue:input_type -> google.devtools.issuetracker.v1.ModifyIssueRequest
	10, // 29: google.devtools.issuetracker.v1.IssueTracker.CreateIssueRelationship:input_type -> google.devtools.issuetracker.v1.CreateIssueRelationshipRequest
	11, // 30: google.devtools.issuetracker.v1.IssueTracker.ListIssueRelationships:input_type -> google.devtools.issuetracker.v1.ListIssueRelationshipsRequest
	13, // 31: google.devtools.issuetracker.v1.IssueTracker.ListIssueUpdates:input_type -> google.devtools.issuetracker.v1.ListIssueUpdatesRequest
	17, // 32: google.devtools.issuetracker.v1.IssueTracker.CreateIssueComment:input_type -> google.devtools.issuetracker.v1.CreateIssueCommentRequest
	15, // 33: google.devtools.issuetracker.v1.IssueTracker.ListIssueComments:input_type -> google.devtools.issuetracker.v1.ListIssueCommentsRequest
	18, // 34: google.devtools.issuetracker.v1.IssueTracker.ListAttachments:input_type -> google.devtools.issuetracker.v1.ListAttachmentsRequest
	20, // 35: google.devtools.issuetracker.v1.IssueTracker.CreateHotlistEntry:input_type -> google.devtools.issuetracker.v1.CreateHotlistEntryRequest
	21, // 36: google.devtools.issuetracker.v1.IssueTracker.DeleteHotlistEntry:input_type -> google.devtools.issuetracker.v1.DeleteHotlistEntryRequest
	31, // 37: google.devtools.issuetracker.v1.IssueTracker.GetComponent:output_type -> google.devtools.issuetracker.v1.Component
	4,  // 38: google.devtools.issuetracker.v1.IssueTracker.ListIssues:output_type -> google.devtools.issuetracker.v1.ListIssuesResponse
	6,  // 39: google.devtools.issuetracker.v1.IssueTracker.BatchGetIssues:output_type -> google.devtools.issuetracker.v1.BatchGetIssuesResponse
	22, // 40: google.devtools.issuetracker.v1.IssueTracker.GetIssue:output_type -> google.devtools.issuetracker.v1.Issue
	22, // 41: google.devtools.issuetracker.v1.IssueTracker.CreateIssue:output_type -> google.devtools.issuetracker.v1.Issue
	22, // 42: google.devtools.issuetracker.v1.IssueTracker.ModifyIssue:output_type -> google.devtools.issuetracker.v1.Issue
	27, // 43: google.devtools.issuetracker.v1.IssueTracker.CreateIssueRelationship:output_type -> google.devtools.issuetracker.v1.IssueRelationship
	12, // 44: google.devtools.issuetracker.v1.IssueTracker.ListIssueRelationships:output_type -> google.devtools.issuetracker.v1.ListIssueRelationshipsResponse
	14, // 45: google.devtools.issuetracker.v1.IssueTracker.ListIssueUpdates:output_type -> google.devtools.issuetracker.v1.ListIssueUpdatesResponse
	25, // 46: google.devtools.issuetracker.v1.IssueTracker.CreateIssueComment:output_type -> google.devtools.issuetracker.v1.IssueComment
	16, // 47: google.devtools.issuetracker.v1.IssueTracker.ListIssueComments:output_type -> google.devtools.issuetracker.v1.ListIssueCommentsResponse
	19, // 48: google.devtools.issuetracker.v1.IssueTracker.ListAttachments:output_type -> google.devtools.issuetracker.v1.ListAttachmentsResponse
	30, // 49: google.devtools.issuetracker.v1.IssueTracker.CreateHotlistEntry:output_type -> google.devtools.issuetracker.v1.HotlistEntry
	32, // 50: google.devtools.issuetracker.v1.IssueTracker.DeleteHotlistEntry:output_type -> google.protobuf.Empty
	37, // [37:51] is the sub-list for method output_type
	23, // [23:37] is the sub-list for method input_type
	23, // [23:23] is the sub-list for extension type_name
	23, // [23:23] is the sub-list for extension extendee
	0,  // [0:23] is the sub-list for field type_name
}

func init() { file_google_devtools_issuetracker_v1_issuetracker_service_proto_init() }
func file_google_devtools_issuetracker_v1_issuetracker_service_proto_init() {
	if File_google_devtools_issuetracker_v1_issuetracker_service_proto != nil {
		return
	}
	file_google_devtools_issuetracker_v1_issuetracker_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetComponentRequest); i {
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
		file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListIssuesRequest); i {
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
		file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListIssuesResponse); i {
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
		file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchGetIssuesRequest); i {
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
		file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchGetIssuesResponse); i {
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
		file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIssueRequest); i {
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
		file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateIssueRequest); i {
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
		file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModifyIssueRequest); i {
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
		file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateIssueRelationshipRequest); i {
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
		file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListIssueRelationshipsRequest); i {
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
		file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListIssueRelationshipsResponse); i {
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
		file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListIssueUpdatesRequest); i {
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
		file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListIssueUpdatesResponse); i {
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
		file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListIssueCommentsRequest); i {
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
		file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListIssueCommentsResponse); i {
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
		file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateIssueCommentRequest); i {
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
		file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[16].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAttachmentsRequest); i {
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
		file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[17].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAttachmentsResponse); i {
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
		file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[18].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateHotlistEntryRequest); i {
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
		file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes[19].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteHotlistEntryRequest); i {
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
			RawDescriptor: file_google_devtools_issuetracker_v1_issuetracker_service_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   20,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_devtools_issuetracker_v1_issuetracker_service_proto_goTypes,
		DependencyIndexes: file_google_devtools_issuetracker_v1_issuetracker_service_proto_depIdxs,
		EnumInfos:         file_google_devtools_issuetracker_v1_issuetracker_service_proto_enumTypes,
		MessageInfos:      file_google_devtools_issuetracker_v1_issuetracker_service_proto_msgTypes,
	}.Build()
	File_google_devtools_issuetracker_v1_issuetracker_service_proto = out.File
	file_google_devtools_issuetracker_v1_issuetracker_service_proto_rawDesc = nil
	file_google_devtools_issuetracker_v1_issuetracker_service_proto_goTypes = nil
	file_google_devtools_issuetracker_v1_issuetracker_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// IssueTrackerClient is the client API for IssueTracker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IssueTrackerClient interface {
	// Gets a component, including its parent hierarchy info.
	GetComponent(ctx context.Context, in *GetComponentRequest, opts ...grpc.CallOption) (*Component, error)
	// Searches issues, and returns issues with their current state.
	ListIssues(ctx context.Context, in *ListIssuesRequest, opts ...grpc.CallOption) (*ListIssuesResponse, error)
	// Gets multiple issues with their current state by their ID. Non-existing
	// issues, or issues that the caller does not have access to, are silently
	// ignored. Note: The maximum number of issues that can be retrieved in one call is
	// limited to 100.
	BatchGetIssues(ctx context.Context, in *BatchGetIssuesRequest, opts ...grpc.CallOption) (*BatchGetIssuesResponse, error)
	// Gets an issue with its current state.
	GetIssue(ctx context.Context, in *GetIssueRequest, opts ...grpc.CallOption) (*Issue, error)
	// Creates a new issue within a component, and returns the new object.
	CreateIssue(ctx context.Context, in *CreateIssueRequest, opts ...grpc.CallOption) (*Issue, error)
	// Updates an issue based on add and remove IssueState.  Returns the
	// modified issue.
	ModifyIssue(ctx context.Context, in *ModifyIssueRequest, opts ...grpc.CallOption) (*Issue, error)
	// Creates a new issue relationship.
	// Requires issue EDIT on the source issue and issue VIEW on the target issue.
	// For relationship_type = CHILD, requires issue EDIT on the source issue and
	// issue VIEW on the target issue.
	CreateIssueRelationship(ctx context.Context, in *CreateIssueRelationshipRequest, opts ...grpc.CallOption) (*IssueRelationship, error)
	// Lists issue relationships under an issue of a type.
	// Requires issue VIEW on the source issue. All target issues are included
	// regardless of the caller's issue view access. target_issue_id is always
	// set. target_issue is set only if the caller has issue VIEW access to the
	// target issue.
	ListIssueRelationships(ctx context.Context, in *ListIssueRelationshipsRequest, opts ...grpc.CallOption) (*ListIssueRelationshipsResponse, error)
	// Fetch a collection of IssueUpdate objects representing the change
	// history of an issue, ordered by IssueUpdate.version.
	ListIssueUpdates(ctx context.Context, in *ListIssueUpdatesRequest, opts ...grpc.CallOption) (*ListIssueUpdatesResponse, error)
	// Creates a new issue comment in an issue
	CreateIssueComment(ctx context.Context, in *CreateIssueCommentRequest, opts ...grpc.CallOption) (*IssueComment, error)
	// Fetches a list of IssueComment objects.
	ListIssueComments(ctx context.Context, in *ListIssueCommentsRequest, opts ...grpc.CallOption) (*ListIssueCommentsResponse, error)
	// List attachments that belong to an issue. Only returns attachment metadata.
	ListAttachments(ctx context.Context, in *ListAttachmentsRequest, opts ...grpc.CallOption) (*ListAttachmentsResponse, error)
	// Adds an issue to a hotlist by creating a HotlistEntry. Returns the created
	// HotlistEntry.
	// Requires hotlist APPEND and issue VIEW permission
	CreateHotlistEntry(ctx context.Context, in *CreateHotlistEntryRequest, opts ...grpc.CallOption) (*HotlistEntry, error)
	// Removes an issue from a hotlist by deleting hotlistEntry. Removing an issue
	// from a hotlist it does not belong to will do nothing and return.
	// Requires hotlist APPEND and issue VIEW permission
	DeleteHotlistEntry(ctx context.Context, in *DeleteHotlistEntryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type issueTrackerClient struct {
	cc grpc.ClientConnInterface
}

func NewIssueTrackerClient(cc grpc.ClientConnInterface) IssueTrackerClient {
	return &issueTrackerClient{cc}
}

func (c *issueTrackerClient) GetComponent(ctx context.Context, in *GetComponentRequest, opts ...grpc.CallOption) (*Component, error) {
	out := new(Component)
	err := c.cc.Invoke(ctx, "/google.devtools.issuetracker.v1.IssueTracker/GetComponent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *issueTrackerClient) ListIssues(ctx context.Context, in *ListIssuesRequest, opts ...grpc.CallOption) (*ListIssuesResponse, error) {
	out := new(ListIssuesResponse)
	err := c.cc.Invoke(ctx, "/google.devtools.issuetracker.v1.IssueTracker/ListIssues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *issueTrackerClient) BatchGetIssues(ctx context.Context, in *BatchGetIssuesRequest, opts ...grpc.CallOption) (*BatchGetIssuesResponse, error) {
	out := new(BatchGetIssuesResponse)
	err := c.cc.Invoke(ctx, "/google.devtools.issuetracker.v1.IssueTracker/BatchGetIssues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *issueTrackerClient) GetIssue(ctx context.Context, in *GetIssueRequest, opts ...grpc.CallOption) (*Issue, error) {
	out := new(Issue)
	err := c.cc.Invoke(ctx, "/google.devtools.issuetracker.v1.IssueTracker/GetIssue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *issueTrackerClient) CreateIssue(ctx context.Context, in *CreateIssueRequest, opts ...grpc.CallOption) (*Issue, error) {
	out := new(Issue)
	err := c.cc.Invoke(ctx, "/google.devtools.issuetracker.v1.IssueTracker/CreateIssue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *issueTrackerClient) ModifyIssue(ctx context.Context, in *ModifyIssueRequest, opts ...grpc.CallOption) (*Issue, error) {
	out := new(Issue)
	err := c.cc.Invoke(ctx, "/google.devtools.issuetracker.v1.IssueTracker/ModifyIssue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *issueTrackerClient) CreateIssueRelationship(ctx context.Context, in *CreateIssueRelationshipRequest, opts ...grpc.CallOption) (*IssueRelationship, error) {
	out := new(IssueRelationship)
	err := c.cc.Invoke(ctx, "/google.devtools.issuetracker.v1.IssueTracker/CreateIssueRelationship", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *issueTrackerClient) ListIssueRelationships(ctx context.Context, in *ListIssueRelationshipsRequest, opts ...grpc.CallOption) (*ListIssueRelationshipsResponse, error) {
	out := new(ListIssueRelationshipsResponse)
	err := c.cc.Invoke(ctx, "/google.devtools.issuetracker.v1.IssueTracker/ListIssueRelationships", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *issueTrackerClient) ListIssueUpdates(ctx context.Context, in *ListIssueUpdatesRequest, opts ...grpc.CallOption) (*ListIssueUpdatesResponse, error) {
	out := new(ListIssueUpdatesResponse)
	err := c.cc.Invoke(ctx, "/google.devtools.issuetracker.v1.IssueTracker/ListIssueUpdates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *issueTrackerClient) CreateIssueComment(ctx context.Context, in *CreateIssueCommentRequest, opts ...grpc.CallOption) (*IssueComment, error) {
	out := new(IssueComment)
	err := c.cc.Invoke(ctx, "/google.devtools.issuetracker.v1.IssueTracker/CreateIssueComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *issueTrackerClient) ListIssueComments(ctx context.Context, in *ListIssueCommentsRequest, opts ...grpc.CallOption) (*ListIssueCommentsResponse, error) {
	out := new(ListIssueCommentsResponse)
	err := c.cc.Invoke(ctx, "/google.devtools.issuetracker.v1.IssueTracker/ListIssueComments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *issueTrackerClient) ListAttachments(ctx context.Context, in *ListAttachmentsRequest, opts ...grpc.CallOption) (*ListAttachmentsResponse, error) {
	out := new(ListAttachmentsResponse)
	err := c.cc.Invoke(ctx, "/google.devtools.issuetracker.v1.IssueTracker/ListAttachments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *issueTrackerClient) CreateHotlistEntry(ctx context.Context, in *CreateHotlistEntryRequest, opts ...grpc.CallOption) (*HotlistEntry, error) {
	out := new(HotlistEntry)
	err := c.cc.Invoke(ctx, "/google.devtools.issuetracker.v1.IssueTracker/CreateHotlistEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *issueTrackerClient) DeleteHotlistEntry(ctx context.Context, in *DeleteHotlistEntryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/google.devtools.issuetracker.v1.IssueTracker/DeleteHotlistEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IssueTrackerServer is the server API for IssueTracker service.
type IssueTrackerServer interface {
	// Gets a component, including its parent hierarchy info.
	GetComponent(context.Context, *GetComponentRequest) (*Component, error)
	// Searches issues, and returns issues with their current state.
	ListIssues(context.Context, *ListIssuesRequest) (*ListIssuesResponse, error)
	// Gets multiple issues with their current state by their ID. Non-existing
	// issues, or issues that the caller does not have access to, are silently
	// ignored. Note: The maximum number of issues that can be retrieved in one call is
	// limited to 100.
	BatchGetIssues(context.Context, *BatchGetIssuesRequest) (*BatchGetIssuesResponse, error)
	// Gets an issue with its current state.
	GetIssue(context.Context, *GetIssueRequest) (*Issue, error)
	// Creates a new issue within a component, and returns the new object.
	CreateIssue(context.Context, *CreateIssueRequest) (*Issue, error)
	// Updates an issue based on add and remove IssueState.  Returns the
	// modified issue.
	ModifyIssue(context.Context, *ModifyIssueRequest) (*Issue, error)
	// Creates a new issue relationship.
	// Requires issue EDIT on the source issue and issue VIEW on the target issue.
	// For relationship_type = CHILD, requires issue EDIT on the source issue and
	// issue VIEW on the target issue.
	CreateIssueRelationship(context.Context, *CreateIssueRelationshipRequest) (*IssueRelationship, error)
	// Lists issue relationships under an issue of a type.
	// Requires issue VIEW on the source issue. All target issues are included
	// regardless of the caller's issue view access. target_issue_id is always
	// set. target_issue is set only if the caller has issue VIEW access to the
	// target issue.
	ListIssueRelationships(context.Context, *ListIssueRelationshipsRequest) (*ListIssueRelationshipsResponse, error)
	// Fetch a collection of IssueUpdate objects representing the change
	// history of an issue, ordered by IssueUpdate.version.
	ListIssueUpdates(context.Context, *ListIssueUpdatesRequest) (*ListIssueUpdatesResponse, error)
	// Creates a new issue comment in an issue
	CreateIssueComment(context.Context, *CreateIssueCommentRequest) (*IssueComment, error)
	// Fetches a list of IssueComment objects.
	ListIssueComments(context.Context, *ListIssueCommentsRequest) (*ListIssueCommentsResponse, error)
	// List attachments that belong to an issue. Only returns attachment metadata.
	ListAttachments(context.Context, *ListAttachmentsRequest) (*ListAttachmentsResponse, error)
	// Adds an issue to a hotlist by creating a HotlistEntry. Returns the created
	// HotlistEntry.
	// Requires hotlist APPEND and issue VIEW permission
	CreateHotlistEntry(context.Context, *CreateHotlistEntryRequest) (*HotlistEntry, error)
	// Removes an issue from a hotlist by deleting hotlistEntry. Removing an issue
	// from a hotlist it does not belong to will do nothing and return.
	// Requires hotlist APPEND and issue VIEW permission
	DeleteHotlistEntry(context.Context, *DeleteHotlistEntryRequest) (*emptypb.Empty, error)
}

// UnimplementedIssueTrackerServer can be embedded to have forward compatible implementations.
type UnimplementedIssueTrackerServer struct {
}

func (*UnimplementedIssueTrackerServer) GetComponent(context.Context, *GetComponentRequest) (*Component, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComponent not implemented")
}
func (*UnimplementedIssueTrackerServer) ListIssues(context.Context, *ListIssuesRequest) (*ListIssuesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListIssues not implemented")
}
func (*UnimplementedIssueTrackerServer) BatchGetIssues(context.Context, *BatchGetIssuesRequest) (*BatchGetIssuesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchGetIssues not implemented")
}
func (*UnimplementedIssueTrackerServer) GetIssue(context.Context, *GetIssueRequest) (*Issue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIssue not implemented")
}
func (*UnimplementedIssueTrackerServer) CreateIssue(context.Context, *CreateIssueRequest) (*Issue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateIssue not implemented")
}
func (*UnimplementedIssueTrackerServer) ModifyIssue(context.Context, *ModifyIssueRequest) (*Issue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyIssue not implemented")
}
func (*UnimplementedIssueTrackerServer) CreateIssueRelationship(context.Context, *CreateIssueRelationshipRequest) (*IssueRelationship, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateIssueRelationship not implemented")
}
func (*UnimplementedIssueTrackerServer) ListIssueRelationships(context.Context, *ListIssueRelationshipsRequest) (*ListIssueRelationshipsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListIssueRelationships not implemented")
}
func (*UnimplementedIssueTrackerServer) ListIssueUpdates(context.Context, *ListIssueUpdatesRequest) (*ListIssueUpdatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListIssueUpdates not implemented")
}
func (*UnimplementedIssueTrackerServer) CreateIssueComment(context.Context, *CreateIssueCommentRequest) (*IssueComment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateIssueComment not implemented")
}
func (*UnimplementedIssueTrackerServer) ListIssueComments(context.Context, *ListIssueCommentsRequest) (*ListIssueCommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListIssueComments not implemented")
}
func (*UnimplementedIssueTrackerServer) ListAttachments(context.Context, *ListAttachmentsRequest) (*ListAttachmentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAttachments not implemented")
}
func (*UnimplementedIssueTrackerServer) CreateHotlistEntry(context.Context, *CreateHotlistEntryRequest) (*HotlistEntry, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHotlistEntry not implemented")
}
func (*UnimplementedIssueTrackerServer) DeleteHotlistEntry(context.Context, *DeleteHotlistEntryRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteHotlistEntry not implemented")
}

func RegisterIssueTrackerServer(s *grpc.Server, srv IssueTrackerServer) {
	s.RegisterService(&_IssueTracker_serviceDesc, srv)
}

func _IssueTracker_GetComponent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetComponentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IssueTrackerServer).GetComponent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.issuetracker.v1.IssueTracker/GetComponent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IssueTrackerServer).GetComponent(ctx, req.(*GetComponentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IssueTracker_ListIssues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListIssuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IssueTrackerServer).ListIssues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.issuetracker.v1.IssueTracker/ListIssues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IssueTrackerServer).ListIssues(ctx, req.(*ListIssuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IssueTracker_BatchGetIssues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchGetIssuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IssueTrackerServer).BatchGetIssues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.issuetracker.v1.IssueTracker/BatchGetIssues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IssueTrackerServer).BatchGetIssues(ctx, req.(*BatchGetIssuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IssueTracker_GetIssue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIssueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IssueTrackerServer).GetIssue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.issuetracker.v1.IssueTracker/GetIssue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IssueTrackerServer).GetIssue(ctx, req.(*GetIssueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IssueTracker_CreateIssue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateIssueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IssueTrackerServer).CreateIssue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.issuetracker.v1.IssueTracker/CreateIssue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IssueTrackerServer).CreateIssue(ctx, req.(*CreateIssueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IssueTracker_ModifyIssue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyIssueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IssueTrackerServer).ModifyIssue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.issuetracker.v1.IssueTracker/ModifyIssue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IssueTrackerServer).ModifyIssue(ctx, req.(*ModifyIssueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IssueTracker_CreateIssueRelationship_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateIssueRelationshipRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IssueTrackerServer).CreateIssueRelationship(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.issuetracker.v1.IssueTracker/CreateIssueRelationship",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IssueTrackerServer).CreateIssueRelationship(ctx, req.(*CreateIssueRelationshipRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IssueTracker_ListIssueRelationships_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListIssueRelationshipsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IssueTrackerServer).ListIssueRelationships(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.issuetracker.v1.IssueTracker/ListIssueRelationships",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IssueTrackerServer).ListIssueRelationships(ctx, req.(*ListIssueRelationshipsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IssueTracker_ListIssueUpdates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListIssueUpdatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IssueTrackerServer).ListIssueUpdates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.issuetracker.v1.IssueTracker/ListIssueUpdates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IssueTrackerServer).ListIssueUpdates(ctx, req.(*ListIssueUpdatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IssueTracker_CreateIssueComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateIssueCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IssueTrackerServer).CreateIssueComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.issuetracker.v1.IssueTracker/CreateIssueComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IssueTrackerServer).CreateIssueComment(ctx, req.(*CreateIssueCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IssueTracker_ListIssueComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListIssueCommentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IssueTrackerServer).ListIssueComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.issuetracker.v1.IssueTracker/ListIssueComments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IssueTrackerServer).ListIssueComments(ctx, req.(*ListIssueCommentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IssueTracker_ListAttachments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAttachmentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IssueTrackerServer).ListAttachments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.issuetracker.v1.IssueTracker/ListAttachments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IssueTrackerServer).ListAttachments(ctx, req.(*ListAttachmentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IssueTracker_CreateHotlistEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateHotlistEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IssueTrackerServer).CreateHotlistEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.issuetracker.v1.IssueTracker/CreateHotlistEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IssueTrackerServer).CreateHotlistEntry(ctx, req.(*CreateHotlistEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IssueTracker_DeleteHotlistEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteHotlistEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IssueTrackerServer).DeleteHotlistEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.issuetracker.v1.IssueTracker/DeleteHotlistEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IssueTrackerServer).DeleteHotlistEntry(ctx, req.(*DeleteHotlistEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _IssueTracker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.devtools.issuetracker.v1.IssueTracker",
	HandlerType: (*IssueTrackerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetComponent",
			Handler:    _IssueTracker_GetComponent_Handler,
		},
		{
			MethodName: "ListIssues",
			Handler:    _IssueTracker_ListIssues_Handler,
		},
		{
			MethodName: "BatchGetIssues",
			Handler:    _IssueTracker_BatchGetIssues_Handler,
		},
		{
			MethodName: "GetIssue",
			Handler:    _IssueTracker_GetIssue_Handler,
		},
		{
			MethodName: "CreateIssue",
			Handler:    _IssueTracker_CreateIssue_Handler,
		},
		{
			MethodName: "ModifyIssue",
			Handler:    _IssueTracker_ModifyIssue_Handler,
		},
		{
			MethodName: "CreateIssueRelationship",
			Handler:    _IssueTracker_CreateIssueRelationship_Handler,
		},
		{
			MethodName: "ListIssueRelationships",
			Handler:    _IssueTracker_ListIssueRelationships_Handler,
		},
		{
			MethodName: "ListIssueUpdates",
			Handler:    _IssueTracker_ListIssueUpdates_Handler,
		},
		{
			MethodName: "CreateIssueComment",
			Handler:    _IssueTracker_CreateIssueComment_Handler,
		},
		{
			MethodName: "ListIssueComments",
			Handler:    _IssueTracker_ListIssueComments_Handler,
		},
		{
			MethodName: "ListAttachments",
			Handler:    _IssueTracker_ListAttachments_Handler,
		},
		{
			MethodName: "CreateHotlistEntry",
			Handler:    _IssueTracker_CreateHotlistEntry_Handler,
		},
		{
			MethodName: "DeleteHotlistEntry",
			Handler:    _IssueTracker_DeleteHotlistEntry_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/devtools/issuetracker/v1/issuetracker_service.proto",
}
