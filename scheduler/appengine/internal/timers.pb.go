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
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.0
// source: go.chromium.org/luci/scheduler/appengine/internal/timers.proto

package internal

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

// Timer can be emitted by any invocation if it wants to be poked later.
//
// Timers are scoped to single invocation and owned by it, so we don't include
// invocation reference here. It is always available from the context of calls.
type Timer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique in time identifier of this timer, auto-generated.
	//
	// It is used to deduplicate and hence provide idempotency for adding
	// timers.
	//
	// Set by the engine, can't be overridden.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Timestamp when the timer was created.
	//
	// Set by the engine, can't be overridden.
	Created *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created,proto3" json:"created,omitempty"`
	// Target time when this timer activates.
	//
	// Should be provided by whoever emits the timer.
	Eta *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=eta,proto3" json:"eta,omitempty"`
	// User friendly name for this timer that shows up in UI.
	//
	// Can be provided by whoever emits the timer. Doesn't have to be unique.
	Title string `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	// Arbitrary optional payload passed verbatim to the invocation.
	Payload []byte `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *Timer) Reset() {
	*x = Timer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_scheduler_appengine_internal_timers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Timer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Timer) ProtoMessage() {}

func (x *Timer) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_scheduler_appengine_internal_timers_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Timer.ProtoReflect.Descriptor instead.
func (*Timer) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_scheduler_appengine_internal_timers_proto_rawDescGZIP(), []int{0}
}

func (x *Timer) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Timer) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *Timer) GetEta() *timestamppb.Timestamp {
	if x != nil {
		return x.Eta
	}
	return nil
}

func (x *Timer) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Timer) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

// TimerList is what we store in datastore entities.
type TimerList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timers []*Timer `protobuf:"bytes,1,rep,name=timers,proto3" json:"timers,omitempty"`
}

func (x *TimerList) Reset() {
	*x = TimerList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_scheduler_appengine_internal_timers_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimerList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimerList) ProtoMessage() {}

func (x *TimerList) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_scheduler_appengine_internal_timers_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimerList.ProtoReflect.Descriptor instead.
func (*TimerList) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_scheduler_appengine_internal_timers_proto_rawDescGZIP(), []int{1}
}

func (x *TimerList) GetTimers() []*Timer {
	if x != nil {
		return x.Timers
	}
	return nil
}

var File_go_chromium_org_luci_scheduler_appengine_internal_timers_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_scheduler_appengine_internal_timers_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72,
	0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x72,
	0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xab, 0x01, 0x0a, 0x05, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x34, 0x0a, 0x07,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x12, 0x2c, 0x0a, 0x03, 0x65, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x03, 0x65, 0x74, 0x61,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x22, 0x3b, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2e, 0x0a,
	0x06, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x73, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x72, 0x52, 0x06, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x73, 0x42, 0x33, 0x5a,
	0x31, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2f,
	0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_scheduler_appengine_internal_timers_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_scheduler_appengine_internal_timers_proto_rawDescData = file_go_chromium_org_luci_scheduler_appengine_internal_timers_proto_rawDesc
)

func file_go_chromium_org_luci_scheduler_appengine_internal_timers_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_scheduler_appengine_internal_timers_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_scheduler_appengine_internal_timers_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_scheduler_appengine_internal_timers_proto_rawDescData)
	})
	return file_go_chromium_org_luci_scheduler_appengine_internal_timers_proto_rawDescData
}

var file_go_chromium_org_luci_scheduler_appengine_internal_timers_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_go_chromium_org_luci_scheduler_appengine_internal_timers_proto_goTypes = []interface{}{
	(*Timer)(nil),                 // 0: internal.timers.Timer
	(*TimerList)(nil),             // 1: internal.timers.TimerList
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_go_chromium_org_luci_scheduler_appengine_internal_timers_proto_depIdxs = []int32{
	2, // 0: internal.timers.Timer.created:type_name -> google.protobuf.Timestamp
	2, // 1: internal.timers.Timer.eta:type_name -> google.protobuf.Timestamp
	0, // 2: internal.timers.TimerList.timers:type_name -> internal.timers.Timer
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_scheduler_appengine_internal_timers_proto_init() }
func file_go_chromium_org_luci_scheduler_appengine_internal_timers_proto_init() {
	if File_go_chromium_org_luci_scheduler_appengine_internal_timers_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_scheduler_appengine_internal_timers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Timer); i {
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
		file_go_chromium_org_luci_scheduler_appengine_internal_timers_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimerList); i {
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
			RawDescriptor: file_go_chromium_org_luci_scheduler_appengine_internal_timers_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_scheduler_appengine_internal_timers_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_scheduler_appengine_internal_timers_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_scheduler_appengine_internal_timers_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_scheduler_appengine_internal_timers_proto = out.File
	file_go_chromium_org_luci_scheduler_appengine_internal_timers_proto_rawDesc = nil
	file_go_chromium_org_luci_scheduler_appengine_internal_timers_proto_goTypes = nil
	file_go_chromium_org_luci_scheduler_appengine_internal_timers_proto_depIdxs = nil
}
