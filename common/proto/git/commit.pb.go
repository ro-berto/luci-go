// Copyright 2017 The LUCI Authors.
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
// source: go.chromium.org/luci/common/proto/git/commit.proto

package git

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

type Commit_TreeDiff_ChangeType int32

const (
	Commit_TreeDiff_ADD    Commit_TreeDiff_ChangeType = 0
	Commit_TreeDiff_COPY   Commit_TreeDiff_ChangeType = 1
	Commit_TreeDiff_DELETE Commit_TreeDiff_ChangeType = 2
	Commit_TreeDiff_MODIFY Commit_TreeDiff_ChangeType = 3
	Commit_TreeDiff_RENAME Commit_TreeDiff_ChangeType = 4
)

// Enum value maps for Commit_TreeDiff_ChangeType.
var (
	Commit_TreeDiff_ChangeType_name = map[int32]string{
		0: "ADD",
		1: "COPY",
		2: "DELETE",
		3: "MODIFY",
		4: "RENAME",
	}
	Commit_TreeDiff_ChangeType_value = map[string]int32{
		"ADD":    0,
		"COPY":   1,
		"DELETE": 2,
		"MODIFY": 3,
		"RENAME": 4,
	}
)

func (x Commit_TreeDiff_ChangeType) Enum() *Commit_TreeDiff_ChangeType {
	p := new(Commit_TreeDiff_ChangeType)
	*p = x
	return p
}

func (x Commit_TreeDiff_ChangeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Commit_TreeDiff_ChangeType) Descriptor() protoreflect.EnumDescriptor {
	return file_go_chromium_org_luci_common_proto_git_commit_proto_enumTypes[0].Descriptor()
}

func (Commit_TreeDiff_ChangeType) Type() protoreflect.EnumType {
	return &file_go_chromium_org_luci_common_proto_git_commit_proto_enumTypes[0]
}

func (x Commit_TreeDiff_ChangeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Commit_TreeDiff_ChangeType.Descriptor instead.
func (Commit_TreeDiff_ChangeType) EnumDescriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_common_proto_git_commit_proto_rawDescGZIP(), []int{0, 1, 0}
}

// Commit is a single parsed commit as represented in a git log or git show
// expression.
type Commit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The hex sha1 of the commit.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The hex sha1 of the tree for this commit.
	Tree string `protobuf:"bytes,2,opt,name=tree,proto3" json:"tree,omitempty"`
	// The hex sha1's of each of this commits' parents.
	Parents   []string     `protobuf:"bytes,3,rep,name=parents,proto3" json:"parents,omitempty"`
	Author    *Commit_User `protobuf:"bytes,4,opt,name=author,proto3" json:"author,omitempty"`
	Committer *Commit_User `protobuf:"bytes,5,opt,name=committer,proto3" json:"committer,omitempty"`
	// This is the entire unaltered message body.
	Message  string             `protobuf:"bytes,6,opt,name=message,proto3" json:"message,omitempty"`
	TreeDiff []*Commit_TreeDiff `protobuf:"bytes,7,rep,name=tree_diff,json=treeDiff,proto3" json:"tree_diff,omitempty"`
}

func (x *Commit) Reset() {
	*x = Commit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_common_proto_git_commit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Commit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Commit) ProtoMessage() {}

func (x *Commit) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_common_proto_git_commit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Commit.ProtoReflect.Descriptor instead.
func (*Commit) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_common_proto_git_commit_proto_rawDescGZIP(), []int{0}
}

func (x *Commit) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Commit) GetTree() string {
	if x != nil {
		return x.Tree
	}
	return ""
}

func (x *Commit) GetParents() []string {
	if x != nil {
		return x.Parents
	}
	return nil
}

func (x *Commit) GetAuthor() *Commit_User {
	if x != nil {
		return x.Author
	}
	return nil
}

func (x *Commit) GetCommitter() *Commit_User {
	if x != nil {
		return x.Committer
	}
	return nil
}

func (x *Commit) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Commit) GetTreeDiff() []*Commit_TreeDiff {
	if x != nil {
		return x.TreeDiff
	}
	return nil
}

// File is a single file as represented in a git tree.
type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID is sha1 hash of the file contents
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Path is the path to file, without leading "/".
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	// Mode is file mode, e.g. 0100644 (octal, often shows up 33188 in decimal).
	Mode uint32 `protobuf:"varint,3,opt,name=mode,proto3" json:"mode,omitempty"` // TODO: add type, perhaps as a enum if needed.
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_common_proto_git_commit_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_common_proto_git_commit_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_common_proto_git_commit_proto_rawDescGZIP(), []int{1}
}

func (x *File) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *File) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *File) GetMode() uint32 {
	if x != nil {
		return x.Mode
	}
	return 0
}

// User represents the (name, email, timestamp) Commit header for author and/or
// commtter.
type Commit_User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Time  *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *Commit_User) Reset() {
	*x = Commit_User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_common_proto_git_commit_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Commit_User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Commit_User) ProtoMessage() {}

func (x *Commit_User) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_common_proto_git_commit_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Commit_User.ProtoReflect.Descriptor instead.
func (*Commit_User) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_common_proto_git_commit_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Commit_User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Commit_User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Commit_User) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

// Each TreeDiff represents a single file that's changed between this commit
// and the "previous" commit, where "previous" depends on the context of how
// this Commit object was produced (i.e. the specific `git log` invocation, or
// similar command).
//
// Note that these are an artifact of the `git log` expression, not of the
// commit itself (since git log has different ways that it could sort the
// commits in the log, and thus different ways it could calculate these
// diffs). In particular, you should avoid caching the TreeDiff data using
// only the Commit.id as the key.
//
// The old_* fields correspond to the matching file in the previous commit (in
// the case of COPY/DELETE/MODIFY/RENAME), telling its blob hash, file mode
// and path name.
//
// The new_* fields correspond to the matching file in this commit (in the
// case of ADD/COPY/MODIFY/RENAME), telling its blob hash, file mode and path
// name.
type Commit_TreeDiff struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// How this file changed.
	Type    Commit_TreeDiff_ChangeType `protobuf:"varint,1,opt,name=type,proto3,enum=git.Commit_TreeDiff_ChangeType" json:"type,omitempty"`
	OldId   string                     `protobuf:"bytes,2,opt,name=old_id,json=oldId,proto3" json:"old_id,omitempty"`
	OldMode uint32                     `protobuf:"varint,3,opt,name=old_mode,json=oldMode,proto3" json:"old_mode,omitempty"`
	OldPath string                     `protobuf:"bytes,4,opt,name=old_path,json=oldPath,proto3" json:"old_path,omitempty"`
	NewId   string                     `protobuf:"bytes,5,opt,name=new_id,json=newId,proto3" json:"new_id,omitempty"`
	NewMode uint32                     `protobuf:"varint,6,opt,name=new_mode,json=newMode,proto3" json:"new_mode,omitempty"`
	NewPath string                     `protobuf:"bytes,7,opt,name=new_path,json=newPath,proto3" json:"new_path,omitempty"`
}

func (x *Commit_TreeDiff) Reset() {
	*x = Commit_TreeDiff{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_common_proto_git_commit_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Commit_TreeDiff) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Commit_TreeDiff) ProtoMessage() {}

func (x *Commit_TreeDiff) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_common_proto_git_commit_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Commit_TreeDiff.ProtoReflect.Descriptor instead.
func (*Commit_TreeDiff) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_common_proto_git_commit_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Commit_TreeDiff) GetType() Commit_TreeDiff_ChangeType {
	if x != nil {
		return x.Type
	}
	return Commit_TreeDiff_ADD
}

func (x *Commit_TreeDiff) GetOldId() string {
	if x != nil {
		return x.OldId
	}
	return ""
}

func (x *Commit_TreeDiff) GetOldMode() uint32 {
	if x != nil {
		return x.OldMode
	}
	return 0
}

func (x *Commit_TreeDiff) GetOldPath() string {
	if x != nil {
		return x.OldPath
	}
	return ""
}

func (x *Commit_TreeDiff) GetNewId() string {
	if x != nil {
		return x.NewId
	}
	return ""
}

func (x *Commit_TreeDiff) GetNewMode() uint32 {
	if x != nil {
		return x.NewMode
	}
	return 0
}

func (x *Commit_TreeDiff) GetNewPath() string {
	if x != nil {
		return x.NewPath
	}
	return ""
}

var File_go_chromium_org_luci_common_proto_git_commit_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_common_proto_git_commit_proto_rawDesc = []byte{
	0x0a, 0x32, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x69, 0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x67, 0x69, 0x74, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf0, 0x04, 0x0a, 0x06, 0x43,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x72, 0x65, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x72, 0x65, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x28, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x69, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x2e, 0x0a,
	0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x67, 0x69, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x31, 0x0a, 0x09, 0x74, 0x72, 0x65, 0x65, 0x5f,
	0x64, 0x69, 0x66, 0x66, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x69, 0x74,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x2e, 0x54, 0x72, 0x65, 0x65, 0x44, 0x69, 0x66, 0x66,
	0x52, 0x08, 0x74, 0x72, 0x65, 0x65, 0x44, 0x69, 0x66, 0x66, 0x1a, 0x60, 0x0a, 0x04, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x2e, 0x0a, 0x04,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x1a, 0x9e, 0x02, 0x0a,
	0x08, 0x54, 0x72, 0x65, 0x65, 0x44, 0x69, 0x66, 0x66, 0x12, 0x33, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x67, 0x69, 0x74, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x69, 0x74, 0x2e, 0x54, 0x72, 0x65, 0x65, 0x44, 0x69, 0x66, 0x66, 0x2e, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x15,
	0x0a, 0x06, 0x6f, 0x6c, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6f, 0x6c, 0x64, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x6c, 0x64, 0x5f, 0x6d, 0x6f, 0x64,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6f, 0x6c, 0x64, 0x4d, 0x6f, 0x64, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x6f, 0x6c, 0x64, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6f, 0x6c, 0x64, 0x50, 0x61, 0x74, 0x68, 0x12, 0x15, 0x0a, 0x06, 0x6e,
	0x65, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x65, 0x77,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x65, 0x77, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6e, 0x65, 0x77, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x6e, 0x65, 0x77, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6e, 0x65, 0x77, 0x50, 0x61, 0x74, 0x68, 0x22, 0x43, 0x0a, 0x0a, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x44, 0x44, 0x10, 0x00, 0x12,
	0x08, 0x0a, 0x04, 0x43, 0x4f, 0x50, 0x59, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c,
	0x45, 0x54, 0x45, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x4f, 0x44, 0x49, 0x46, 0x59, 0x10,
	0x03, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x45, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x04, 0x22, 0x3e, 0x0a,
	0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x64,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x42, 0x27, 0x5a,
	0x25, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x69, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_common_proto_git_commit_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_common_proto_git_commit_proto_rawDescData = file_go_chromium_org_luci_common_proto_git_commit_proto_rawDesc
)

func file_go_chromium_org_luci_common_proto_git_commit_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_common_proto_git_commit_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_common_proto_git_commit_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_common_proto_git_commit_proto_rawDescData)
	})
	return file_go_chromium_org_luci_common_proto_git_commit_proto_rawDescData
}

var file_go_chromium_org_luci_common_proto_git_commit_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_go_chromium_org_luci_common_proto_git_commit_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_go_chromium_org_luci_common_proto_git_commit_proto_goTypes = []interface{}{
	(Commit_TreeDiff_ChangeType)(0), // 0: git.Commit.TreeDiff.ChangeType
	(*Commit)(nil),                  // 1: git.Commit
	(*File)(nil),                    // 2: git.File
	(*Commit_User)(nil),             // 3: git.Commit.User
	(*Commit_TreeDiff)(nil),         // 4: git.Commit.TreeDiff
	(*timestamppb.Timestamp)(nil),   // 5: google.protobuf.Timestamp
}
var file_go_chromium_org_luci_common_proto_git_commit_proto_depIdxs = []int32{
	3, // 0: git.Commit.author:type_name -> git.Commit.User
	3, // 1: git.Commit.committer:type_name -> git.Commit.User
	4, // 2: git.Commit.tree_diff:type_name -> git.Commit.TreeDiff
	5, // 3: git.Commit.User.time:type_name -> google.protobuf.Timestamp
	0, // 4: git.Commit.TreeDiff.type:type_name -> git.Commit.TreeDiff.ChangeType
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_common_proto_git_commit_proto_init() }
func file_go_chromium_org_luci_common_proto_git_commit_proto_init() {
	if File_go_chromium_org_luci_common_proto_git_commit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_common_proto_git_commit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Commit); i {
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
		file_go_chromium_org_luci_common_proto_git_commit_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File); i {
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
		file_go_chromium_org_luci_common_proto_git_commit_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Commit_User); i {
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
		file_go_chromium_org_luci_common_proto_git_commit_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Commit_TreeDiff); i {
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
			RawDescriptor: file_go_chromium_org_luci_common_proto_git_commit_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_common_proto_git_commit_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_common_proto_git_commit_proto_depIdxs,
		EnumInfos:         file_go_chromium_org_luci_common_proto_git_commit_proto_enumTypes,
		MessageInfos:      file_go_chromium_org_luci_common_proto_git_commit_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_common_proto_git_commit_proto = out.File
	file_go_chromium_org_luci_common_proto_git_commit_proto_rawDesc = nil
	file_go_chromium_org_luci_common_proto_git_commit_proto_goTypes = nil
	file_go_chromium_org_luci_common_proto_git_commit_proto_depIdxs = nil
}
