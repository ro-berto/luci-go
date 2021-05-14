// Copyright 2016 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.0
// source: go.chromium.org/luci/dm/api/service/v1/graph_query.proto

package dm

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

type GraphQuery_Search_Domain int32

const (
	GraphQuery_Search_QUEST   GraphQuery_Search_Domain = 0
	GraphQuery_Search_ATTEMPT GraphQuery_Search_Domain = 1
)

// Enum value maps for GraphQuery_Search_Domain.
var (
	GraphQuery_Search_Domain_name = map[int32]string{
		0: "QUEST",
		1: "ATTEMPT",
	}
	GraphQuery_Search_Domain_value = map[string]int32{
		"QUEST":   0,
		"ATTEMPT": 1,
	}
)

func (x GraphQuery_Search_Domain) Enum() *GraphQuery_Search_Domain {
	p := new(GraphQuery_Search_Domain)
	*p = x
	return p
}

func (x GraphQuery_Search_Domain) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GraphQuery_Search_Domain) Descriptor() protoreflect.EnumDescriptor {
	return file_go_chromium_org_luci_dm_api_service_v1_graph_query_proto_enumTypes[0].Descriptor()
}

func (GraphQuery_Search_Domain) Type() protoreflect.EnumType {
	return &file_go_chromium_org_luci_dm_api_service_v1_graph_query_proto_enumTypes[0]
}

func (x GraphQuery_Search_Domain) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GraphQuery_Search_Domain.Descriptor instead.
func (GraphQuery_Search_Domain) EnumDescriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_dm_api_service_v1_graph_query_proto_rawDescGZIP(), []int{0, 1, 0}
}

// GraphQuery represents a single query into the state of DM's dependency graph.
// It's a required parameter for WalkGraphReq.
type GraphQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// AttemptList allows you to list one or more specific attempts as the result
	// of the query. If a quest contains the attempt number 0, or is empty, it
	// means 'all attempts for this quest'.
	AttemptList *AttemptList `protobuf:"bytes,1,opt,name=attempt_list,json=attemptList,proto3" json:"attempt_list,omitempty"`
	// attempt_range allows you to list a range of attempts in a single quest.
	// low must be > 0, and high must be > low. The range is [low, high). High may
	// be higher than the highest attempt, and low may be lower than the lowest
	// attempt (but not 0).
	AttemptRange []*GraphQuery_AttemptRange `protobuf:"bytes,2,rep,name=attempt_range,json=attemptRange,proto3" json:"attempt_range,omitempty"`
	Search       []*GraphQuery_Search       `protobuf:"bytes,3,rep,name=search,proto3" json:"search,omitempty"`
}

func (x *GraphQuery) Reset() {
	*x = GraphQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_dm_api_service_v1_graph_query_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GraphQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GraphQuery) ProtoMessage() {}

func (x *GraphQuery) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_dm_api_service_v1_graph_query_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GraphQuery.ProtoReflect.Descriptor instead.
func (*GraphQuery) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_dm_api_service_v1_graph_query_proto_rawDescGZIP(), []int{0}
}

func (x *GraphQuery) GetAttemptList() *AttemptList {
	if x != nil {
		return x.AttemptList
	}
	return nil
}

func (x *GraphQuery) GetAttemptRange() []*GraphQuery_AttemptRange {
	if x != nil {
		return x.AttemptRange
	}
	return nil
}

func (x *GraphQuery) GetSearch() []*GraphQuery_Search {
	if x != nil {
		return x.Search
	}
	return nil
}

type GraphQuery_AttemptRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Quest string `protobuf:"bytes,1,opt,name=quest,proto3" json:"quest,omitempty"`
	Low   uint32 `protobuf:"varint,2,opt,name=low,proto3" json:"low,omitempty"`
	High  uint32 `protobuf:"varint,3,opt,name=high,proto3" json:"high,omitempty"`
}

func (x *GraphQuery_AttemptRange) Reset() {
	*x = GraphQuery_AttemptRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_dm_api_service_v1_graph_query_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GraphQuery_AttemptRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GraphQuery_AttemptRange) ProtoMessage() {}

func (x *GraphQuery_AttemptRange) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_dm_api_service_v1_graph_query_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GraphQuery_AttemptRange.ProtoReflect.Descriptor instead.
func (*GraphQuery_AttemptRange) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_dm_api_service_v1_graph_query_proto_rawDescGZIP(), []int{0, 0}
}

func (x *GraphQuery_AttemptRange) GetQuest() string {
	if x != nil {
		return x.Quest
	}
	return ""
}

func (x *GraphQuery_AttemptRange) GetLow() uint32 {
	if x != nil {
		return x.Low
	}
	return 0
}

func (x *GraphQuery_AttemptRange) GetHigh() uint32 {
	if x != nil {
		return x.High
	}
	return 0
}

// A Search allows you to query objects whose properties match all of the
// provided filters. Filters take the form of a dot-delimited path. For
// example, say that we had the following objects:
//
//   Quest(id=deadbeef):
//     created = <timestamp>  #sort
//     descriptor.distributor_config_name = "foo"
//     descriptor.json_payload = {
//       "key": "value",
//       "multi": ["some", 10, "values", true],
//       "sub": [{"msg": 11}, {"msg": 12}],
//     }
//
//   Attempt(id=deadbeef|1):
//     created = <timestamp>  #sort
//     attempt_type = Finished
//     finished.expiration = <timestamp>
//     finished.json_result = {
//       "rslt": "yes",
//       "ok": true,
//     }
//
// Then you could query (in pseudo-proto):
//   domain: Attempt
//   approx_filters: {
//     "attempt_type": ["Finished"],
//     "$quest.descriptor.json_payload.multi": [true, 10],
//     "$quest.descriptor.json_payload.sub.msg": [11, 10],
//     "finished.json_result.ok": [true],
//   }
//
// Or:
//
//   domain: Attempt
//   exact_filters: {
//     "$quest.descriptor.json_payload.multi[1]": [10],
//     "$quest.descriptor.json_payload.sub[0].msg": [11],
//   }
//
// Literal '.' and '[' characters may be escaped with a backslash.
type GraphQuery_Search struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Domain indicates which class of objects your query applies to. The fields
	// available to query are defined by the `data` field in the corresponding
	// GraphData message.
	//
	// Additionally `Attempt` has a special field $quest whose subfields are
	// queriable in the exact same way that a search in a Quest domain works.
	Domain GraphQuery_Search_Domain `protobuf:"varint,1,opt,name=domain,proto3,enum=dm.GraphQuery_Search_Domain" json:"domain,omitempty"`
	// Start and End are optional restrictions on the first sort property. For
	// now, these are just restrictions on the 'created' timestamp for either
	// the Quest or Attempt, depending on the SearchDomain.
	Start *PropertyValue `protobuf:"bytes,3,opt,name=start,proto3" json:"start,omitempty"`
	End   *PropertyValue `protobuf:"bytes,4,opt,name=end,proto3" json:"end,omitempty"`
	// ApproxFilters allows you to filter on 'approximate' fields. Approximate
	// fields are the json path to the value, without any array subscripts. For
	// example, if your document looked like:
	//
	//   {
	//     "some": ["list", {"of": ["data", "and", "stuff"]}],
	//   }
	//
	// Then the following approximate filters would match:
	//   "some" = ["list"]
	//   "some.of" = ["data"]
	//   "some.of" = ["and"]
	//   "some.of" = ["stuff"]
	//   "some.of" = ["stuff", "and"]
	//   "some.of" = ["stuff", "and", "data"]
	//
	// This is useful for filtering documents where the order of parameters
	// in a list or sublist isn't known, or doesn't matter.
	ApproxFilters map[string]*MultiPropertyValue `protobuf:"bytes,5,rep,name=approx_filters,json=approxFilters,proto3" json:"approx_filters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// ExactFilters allows you to filter on 'exact' fields. Exact fields are the
	// json path to the value, including array subscripts. For example if your
	// document looked like:
	//
	//   {
	//     "some": ["list", {"of": ["data", "and", "stuff"]}],
	//   }
	//
	// Then the following exact filters would match:
	//   "some[0]" = "list"
	//   "some[1].of[0]" = "data"
	//   "some[1].of[1]" = "and"
	//   "some[1].of[2]" = "stuff"
	//
	// This is useful for filtering documents where the order of parameters
	// in a list or sublist matters.
	ExactFilters map[string]*PropertyValue `protobuf:"bytes,6,rep,name=exact_filters,json=exactFilters,proto3" json:"exact_filters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GraphQuery_Search) Reset() {
	*x = GraphQuery_Search{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_dm_api_service_v1_graph_query_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GraphQuery_Search) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GraphQuery_Search) ProtoMessage() {}

func (x *GraphQuery_Search) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_dm_api_service_v1_graph_query_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GraphQuery_Search.ProtoReflect.Descriptor instead.
func (*GraphQuery_Search) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_dm_api_service_v1_graph_query_proto_rawDescGZIP(), []int{0, 1}
}

func (x *GraphQuery_Search) GetDomain() GraphQuery_Search_Domain {
	if x != nil {
		return x.Domain
	}
	return GraphQuery_Search_QUEST
}

func (x *GraphQuery_Search) GetStart() *PropertyValue {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *GraphQuery_Search) GetEnd() *PropertyValue {
	if x != nil {
		return x.End
	}
	return nil
}

func (x *GraphQuery_Search) GetApproxFilters() map[string]*MultiPropertyValue {
	if x != nil {
		return x.ApproxFilters
	}
	return nil
}

func (x *GraphQuery_Search) GetExactFilters() map[string]*PropertyValue {
	if x != nil {
		return x.ExactFilters
	}
	return nil
}

var File_go_chromium_org_luci_dm_api_service_v1_graph_query_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_dm_api_service_v1_graph_query_proto_rawDesc = []byte{
	0x0a, 0x38, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x64, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x5f, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x64, 0x6d, 0x1a, 0x32,
	0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x6c, 0x75, 0x63, 0x69, 0x2f, 0x64, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x8a, 0x06, 0x0a, 0x0a, 0x47, 0x72, 0x61, 0x70, 0x68, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x12, 0x32, 0x0a, 0x0c, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x64, 0x6d, 0x2e, 0x41, 0x74, 0x74,
	0x65, 0x6d, 0x70, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x0b, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x0d, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74,
	0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x64,
	0x6d, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x41, 0x74, 0x74,
	0x65, 0x6d, 0x70, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x0c, 0x61, 0x74, 0x74, 0x65, 0x6d,
	0x70, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x64, 0x6d, 0x2e, 0x47, 0x72, 0x61,
	0x70, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x06,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x1a, 0x4a, 0x0a, 0x0c, 0x41, 0x74, 0x74, 0x65, 0x6d, 0x70,
	0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x6c, 0x6f, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6c, 0x6f, 0x77, 0x12, 0x12,
	0x0a, 0x04, 0x68, 0x69, 0x67, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x68, 0x69,
	0x67, 0x68, 0x1a, 0x8a, 0x04, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x34, 0x0a,
	0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e,
	0x64, 0x6d, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x06, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x12, 0x27, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x64, 0x6d, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x23, 0x0a, 0x03,
	0x65, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x64, 0x6d, 0x2e, 0x50,
	0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x03, 0x65, 0x6e,
	0x64, 0x12, 0x4f, 0x0a, 0x0e, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x78, 0x5f, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x64, 0x6d, 0x2e, 0x47,
	0x72, 0x61, 0x70, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x2e, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x78, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x0d, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x78, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x12, 0x4c, 0x0a, 0x0d, 0x65, 0x78, 0x61, 0x63, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x64, 0x6d, 0x2e, 0x47,
	0x72, 0x61, 0x70, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x2e, 0x45, 0x78, 0x61, 0x63, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0c, 0x65, 0x78, 0x61, 0x63, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x1a, 0x58, 0x0a, 0x12, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x78, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x64, 0x6d, 0x2e, 0x4d, 0x75, 0x6c,
	0x74, 0x69, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x52, 0x0a, 0x11, 0x45, 0x78,
	0x61, 0x63, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x27, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x64, 0x6d, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x20,
	0x0a, 0x06, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x09, 0x0a, 0x05, 0x51, 0x55, 0x45, 0x53,
	0x54, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x54, 0x54, 0x45, 0x4d, 0x50, 0x54, 0x10, 0x01,
	0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x52, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x62, 0x79, 0x42,
	0x2b, 0x5a, 0x29, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x64, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x64, 0x6d, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_dm_api_service_v1_graph_query_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_dm_api_service_v1_graph_query_proto_rawDescData = file_go_chromium_org_luci_dm_api_service_v1_graph_query_proto_rawDesc
)

func file_go_chromium_org_luci_dm_api_service_v1_graph_query_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_dm_api_service_v1_graph_query_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_dm_api_service_v1_graph_query_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_dm_api_service_v1_graph_query_proto_rawDescData)
	})
	return file_go_chromium_org_luci_dm_api_service_v1_graph_query_proto_rawDescData
}

var file_go_chromium_org_luci_dm_api_service_v1_graph_query_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_go_chromium_org_luci_dm_api_service_v1_graph_query_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_go_chromium_org_luci_dm_api_service_v1_graph_query_proto_goTypes = []interface{}{
	(GraphQuery_Search_Domain)(0),   // 0: dm.GraphQuery.Search.Domain
	(*GraphQuery)(nil),              // 1: dm.GraphQuery
	(*GraphQuery_AttemptRange)(nil), // 2: dm.GraphQuery.AttemptRange
	(*GraphQuery_Search)(nil),       // 3: dm.GraphQuery.Search
	nil,                             // 4: dm.GraphQuery.Search.ApproxFiltersEntry
	nil,                             // 5: dm.GraphQuery.Search.ExactFiltersEntry
	(*AttemptList)(nil),             // 6: dm.AttemptList
	(*PropertyValue)(nil),           // 7: dm.PropertyValue
	(*MultiPropertyValue)(nil),      // 8: dm.MultiPropertyValue
}
var file_go_chromium_org_luci_dm_api_service_v1_graph_query_proto_depIdxs = []int32{
	6,  // 0: dm.GraphQuery.attempt_list:type_name -> dm.AttemptList
	2,  // 1: dm.GraphQuery.attempt_range:type_name -> dm.GraphQuery.AttemptRange
	3,  // 2: dm.GraphQuery.search:type_name -> dm.GraphQuery.Search
	0,  // 3: dm.GraphQuery.Search.domain:type_name -> dm.GraphQuery.Search.Domain
	7,  // 4: dm.GraphQuery.Search.start:type_name -> dm.PropertyValue
	7,  // 5: dm.GraphQuery.Search.end:type_name -> dm.PropertyValue
	4,  // 6: dm.GraphQuery.Search.approx_filters:type_name -> dm.GraphQuery.Search.ApproxFiltersEntry
	5,  // 7: dm.GraphQuery.Search.exact_filters:type_name -> dm.GraphQuery.Search.ExactFiltersEntry
	8,  // 8: dm.GraphQuery.Search.ApproxFiltersEntry.value:type_name -> dm.MultiPropertyValue
	7,  // 9: dm.GraphQuery.Search.ExactFiltersEntry.value:type_name -> dm.PropertyValue
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_dm_api_service_v1_graph_query_proto_init() }
func file_go_chromium_org_luci_dm_api_service_v1_graph_query_proto_init() {
	if File_go_chromium_org_luci_dm_api_service_v1_graph_query_proto != nil {
		return
	}
	file_go_chromium_org_luci_dm_api_service_v1_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_dm_api_service_v1_graph_query_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GraphQuery); i {
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
		file_go_chromium_org_luci_dm_api_service_v1_graph_query_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GraphQuery_AttemptRange); i {
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
		file_go_chromium_org_luci_dm_api_service_v1_graph_query_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GraphQuery_Search); i {
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
			RawDescriptor: file_go_chromium_org_luci_dm_api_service_v1_graph_query_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_dm_api_service_v1_graph_query_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_dm_api_service_v1_graph_query_proto_depIdxs,
		EnumInfos:         file_go_chromium_org_luci_dm_api_service_v1_graph_query_proto_enumTypes,
		MessageInfos:      file_go_chromium_org_luci_dm_api_service_v1_graph_query_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_dm_api_service_v1_graph_query_proto = out.File
	file_go_chromium_org_luci_dm_api_service_v1_graph_query_proto_rawDesc = nil
	file_go_chromium_org_luci_dm_api_service_v1_graph_query_proto_goTypes = nil
	file_go_chromium_org_luci_dm_api_service_v1_graph_query_proto_depIdxs = nil
}
