// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.0
// source: go.chromium.org/luci/client/flagpb/unmarshal_test.proto

package flagpb

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

type E int32

const (
	E_V0 E = 0
	E_V1 E = 1
)

// Enum value maps for E.
var (
	E_name = map[int32]string{
		0: "V0",
		1: "V1",
	}
	E_value = map[string]int32{
		"V0": 0,
		"V1": 1,
	}
)

func (x E) Enum() *E {
	p := new(E)
	*p = x
	return p
}

func (x E) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (E) Descriptor() protoreflect.EnumDescriptor {
	return file_go_chromium_org_luci_client_flagpb_unmarshal_test_proto_enumTypes[0].Descriptor()
}

func (E) Type() protoreflect.EnumType {
	return &file_go_chromium_org_luci_client_flagpb_unmarshal_test_proto_enumTypes[0]
}

func (x E) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use E.Descriptor instead.
func (E) EnumDescriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_client_flagpb_unmarshal_test_proto_rawDescGZIP(), []int{0}
}

type M1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	S  string  `protobuf:"bytes,1,opt,name=s,proto3" json:"s,omitempty"`
	I  int32   `protobuf:"varint,2,opt,name=i,proto3" json:"i,omitempty"`
	Ri []int32 `protobuf:"varint,3,rep,packed,name=ri,proto3" json:"ri,omitempty"`
	B  bool    `protobuf:"varint,4,opt,name=b,proto3" json:"b,omitempty"`
	Rb []bool  `protobuf:"varint,6,rep,packed,name=rb,proto3" json:"rb,omitempty"`
	Bb []byte  `protobuf:"bytes,5,opt,name=bb,proto3" json:"bb,omitempty"`
}

func (x *M1) Reset() {
	*x = M1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_client_flagpb_unmarshal_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *M1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*M1) ProtoMessage() {}

func (x *M1) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_client_flagpb_unmarshal_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use M1.ProtoReflect.Descriptor instead.
func (*M1) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_client_flagpb_unmarshal_test_proto_rawDescGZIP(), []int{0}
}

func (x *M1) GetS() string {
	if x != nil {
		return x.S
	}
	return ""
}

func (x *M1) GetI() int32 {
	if x != nil {
		return x.I
	}
	return 0
}

func (x *M1) GetRi() []int32 {
	if x != nil {
		return x.Ri
	}
	return nil
}

func (x *M1) GetB() bool {
	if x != nil {
		return x.B
	}
	return false
}

func (x *M1) GetRb() []bool {
	if x != nil {
		return x.Rb
	}
	return nil
}

func (x *M1) GetBb() []byte {
	if x != nil {
		return x.Bb
	}
	return nil
}

type M2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	M1 *M1 `protobuf:"bytes,1,opt,name=m1,proto3" json:"m1,omitempty"`
	E  E   `protobuf:"varint,2,opt,name=e,proto3,enum=flagpb.E" json:"e,omitempty"`
}

func (x *M2) Reset() {
	*x = M2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_client_flagpb_unmarshal_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *M2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*M2) ProtoMessage() {}

func (x *M2) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_client_flagpb_unmarshal_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use M2.ProtoReflect.Descriptor instead.
func (*M2) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_client_flagpb_unmarshal_test_proto_rawDescGZIP(), []int{1}
}

func (x *M2) GetM1() *M1 {
	if x != nil {
		return x.M1
	}
	return nil
}

func (x *M2) GetE() E {
	if x != nil {
		return x.E
	}
	return E_V0
}

type M3 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	M1 []*M1  `protobuf:"bytes,1,rep,name=m1,proto3" json:"m1,omitempty"`
	M2 *M2    `protobuf:"bytes,2,opt,name=m2,proto3" json:"m2,omitempty"`
	B  bool   `protobuf:"varint,3,opt,name=b,proto3" json:"b,omitempty"`
	S  string `protobuf:"bytes,4,opt,name=s,proto3" json:"s,omitempty"`
	Bt []byte `protobuf:"bytes,5,opt,name=bt,proto3" json:"bt,omitempty"`
}

func (x *M3) Reset() {
	*x = M3{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_client_flagpb_unmarshal_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *M3) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*M3) ProtoMessage() {}

func (x *M3) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_client_flagpb_unmarshal_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use M3.ProtoReflect.Descriptor instead.
func (*M3) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_client_flagpb_unmarshal_test_proto_rawDescGZIP(), []int{2}
}

func (x *M3) GetM1() []*M1 {
	if x != nil {
		return x.M1
	}
	return nil
}

func (x *M3) GetM2() *M2 {
	if x != nil {
		return x.M2
	}
	return nil
}

func (x *M3) GetB() bool {
	if x != nil {
		return x.B
	}
	return false
}

func (x *M3) GetS() string {
	if x != nil {
		return x.S
	}
	return ""
}

func (x *M3) GetBt() []byte {
	if x != nil {
		return x.Bt
	}
	return nil
}

type MapContainer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ss  map[string]string `protobuf:"bytes,1,rep,name=ss,proto3" json:"ss,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Ii  map[int32]int32   `protobuf:"bytes,2,rep,name=ii,proto3" json:"ii,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Sm1 map[string]*M1    `protobuf:"bytes,3,rep,name=sm1,proto3" json:"sm1,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MapContainer) Reset() {
	*x = MapContainer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_client_flagpb_unmarshal_test_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapContainer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapContainer) ProtoMessage() {}

func (x *MapContainer) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_client_flagpb_unmarshal_test_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapContainer.ProtoReflect.Descriptor instead.
func (*MapContainer) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_client_flagpb_unmarshal_test_proto_rawDescGZIP(), []int{3}
}

func (x *MapContainer) GetSs() map[string]string {
	if x != nil {
		return x.Ss
	}
	return nil
}

func (x *MapContainer) GetIi() map[int32]int32 {
	if x != nil {
		return x.Ii
	}
	return nil
}

func (x *MapContainer) GetSm1() map[string]*M1 {
	if x != nil {
		return x.Sm1
	}
	return nil
}

var File_go_chromium_org_luci_client_flagpb_unmarshal_test_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_client_flagpb_unmarshal_test_proto_rawDesc = []byte{
	0x0a, 0x37, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x66, 0x6c,
	0x61, 0x67, 0x70, 0x62, 0x2f, 0x75, 0x6e, 0x6d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x5f, 0x74,
	0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x66, 0x6c, 0x61, 0x67, 0x70,
	0x62, 0x22, 0x5e, 0x0a, 0x02, 0x4d, 0x31, 0x12, 0x0c, 0x0a, 0x01, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x01, 0x73, 0x12, 0x0c, 0x0a, 0x01, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x01, 0x69, 0x12, 0x0e, 0x0a, 0x02, 0x72, 0x69, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52,
	0x02, 0x72, 0x69, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x01,
	0x62, 0x12, 0x0e, 0x0a, 0x02, 0x72, 0x62, 0x18, 0x06, 0x20, 0x03, 0x28, 0x08, 0x52, 0x02, 0x72,
	0x62, 0x12, 0x0e, 0x0a, 0x02, 0x62, 0x62, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x62,
	0x62, 0x22, 0x39, 0x0a, 0x02, 0x4d, 0x32, 0x12, 0x1a, 0x0a, 0x02, 0x6d, 0x31, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x66, 0x6c, 0x61, 0x67, 0x70, 0x62, 0x2e, 0x4d, 0x31, 0x52,
	0x02, 0x6d, 0x31, 0x12, 0x17, 0x0a, 0x01, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09,
	0x2e, 0x66, 0x6c, 0x61, 0x67, 0x70, 0x62, 0x2e, 0x45, 0x52, 0x01, 0x65, 0x22, 0x68, 0x0a, 0x02,
	0x4d, 0x33, 0x12, 0x1a, 0x0a, 0x02, 0x6d, 0x31, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x66, 0x6c, 0x61, 0x67, 0x70, 0x62, 0x2e, 0x4d, 0x31, 0x52, 0x02, 0x6d, 0x31, 0x12, 0x1a,
	0x0a, 0x02, 0x6d, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x66, 0x6c, 0x61,
	0x67, 0x70, 0x62, 0x2e, 0x4d, 0x32, 0x52, 0x02, 0x6d, 0x32, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x01, 0x62, 0x12, 0x0c, 0x0a, 0x01, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x01, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x62, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x02, 0x62, 0x74, 0x22, 0xcd, 0x02, 0x0a, 0x0c, 0x4d, 0x61, 0x70, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x02, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x66, 0x6c, 0x61, 0x67, 0x70, 0x62, 0x2e, 0x4d, 0x61, 0x70,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x53, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x02, 0x73, 0x73, 0x12, 0x2c, 0x0a, 0x02, 0x69, 0x69, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x66, 0x6c, 0x61, 0x67, 0x70, 0x62, 0x2e, 0x4d, 0x61, 0x70, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x49, 0x69, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x02, 0x69, 0x69, 0x12, 0x2f, 0x0a, 0x03, 0x73, 0x6d, 0x31, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x66, 0x6c, 0x61, 0x67, 0x70, 0x62, 0x2e, 0x4d, 0x61, 0x70, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x53, 0x6d, 0x31, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x03, 0x73, 0x6d, 0x31, 0x1a, 0x35, 0x0a, 0x07, 0x53, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x35, 0x0a, 0x07, 0x49,
	0x69, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x1a, 0x42, 0x0a, 0x08, 0x53, 0x6d, 0x31, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x20, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x66, 0x6c, 0x61, 0x67, 0x70, 0x62, 0x2e, 0x4d, 0x31, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x13, 0x0a, 0x01, 0x45, 0x12, 0x06, 0x0a, 0x02, 0x56,
	0x30, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x56, 0x31, 0x10, 0x01, 0x42, 0x24, 0x5a, 0x22, 0x67,
	0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c,
	0x75, 0x63, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x66, 0x6c, 0x61, 0x67, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_client_flagpb_unmarshal_test_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_client_flagpb_unmarshal_test_proto_rawDescData = file_go_chromium_org_luci_client_flagpb_unmarshal_test_proto_rawDesc
)

func file_go_chromium_org_luci_client_flagpb_unmarshal_test_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_client_flagpb_unmarshal_test_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_client_flagpb_unmarshal_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_client_flagpb_unmarshal_test_proto_rawDescData)
	})
	return file_go_chromium_org_luci_client_flagpb_unmarshal_test_proto_rawDescData
}

var file_go_chromium_org_luci_client_flagpb_unmarshal_test_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_go_chromium_org_luci_client_flagpb_unmarshal_test_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_go_chromium_org_luci_client_flagpb_unmarshal_test_proto_goTypes = []interface{}{
	(E)(0),               // 0: flagpb.E
	(*M1)(nil),           // 1: flagpb.M1
	(*M2)(nil),           // 2: flagpb.M2
	(*M3)(nil),           // 3: flagpb.M3
	(*MapContainer)(nil), // 4: flagpb.MapContainer
	nil,                  // 5: flagpb.MapContainer.SsEntry
	nil,                  // 6: flagpb.MapContainer.IiEntry
	nil,                  // 7: flagpb.MapContainer.Sm1Entry
}
var file_go_chromium_org_luci_client_flagpb_unmarshal_test_proto_depIdxs = []int32{
	1, // 0: flagpb.M2.m1:type_name -> flagpb.M1
	0, // 1: flagpb.M2.e:type_name -> flagpb.E
	1, // 2: flagpb.M3.m1:type_name -> flagpb.M1
	2, // 3: flagpb.M3.m2:type_name -> flagpb.M2
	5, // 4: flagpb.MapContainer.ss:type_name -> flagpb.MapContainer.SsEntry
	6, // 5: flagpb.MapContainer.ii:type_name -> flagpb.MapContainer.IiEntry
	7, // 6: flagpb.MapContainer.sm1:type_name -> flagpb.MapContainer.Sm1Entry
	1, // 7: flagpb.MapContainer.Sm1Entry.value:type_name -> flagpb.M1
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_client_flagpb_unmarshal_test_proto_init() }
func file_go_chromium_org_luci_client_flagpb_unmarshal_test_proto_init() {
	if File_go_chromium_org_luci_client_flagpb_unmarshal_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_client_flagpb_unmarshal_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*M1); i {
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
		file_go_chromium_org_luci_client_flagpb_unmarshal_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*M2); i {
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
		file_go_chromium_org_luci_client_flagpb_unmarshal_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*M3); i {
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
		file_go_chromium_org_luci_client_flagpb_unmarshal_test_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MapContainer); i {
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
			RawDescriptor: file_go_chromium_org_luci_client_flagpb_unmarshal_test_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_client_flagpb_unmarshal_test_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_client_flagpb_unmarshal_test_proto_depIdxs,
		EnumInfos:         file_go_chromium_org_luci_client_flagpb_unmarshal_test_proto_enumTypes,
		MessageInfos:      file_go_chromium_org_luci_client_flagpb_unmarshal_test_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_client_flagpb_unmarshal_test_proto = out.File
	file_go_chromium_org_luci_client_flagpb_unmarshal_test_proto_rawDesc = nil
	file_go_chromium_org_luci_client_flagpb_unmarshal_test_proto_goTypes = nil
	file_go_chromium_org_luci_client_flagpb_unmarshal_test_proto_depIdxs = nil
}
