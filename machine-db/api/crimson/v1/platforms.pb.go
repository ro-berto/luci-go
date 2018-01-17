// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/machine-db/api/crimson/v1/platforms.proto

package crimson

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// ListPlatformsRequest is a request to retrieve platforms.
type ListPlatformsRequest struct {
	// The names of platforms to retrieve.
	Names []string `protobuf:"bytes,1,rep,name=names" json:"names,omitempty"`
}

func (m *ListPlatformsRequest) Reset()                    { *m = ListPlatformsRequest{} }
func (m *ListPlatformsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListPlatformsRequest) ProtoMessage()               {}
func (*ListPlatformsRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *ListPlatformsRequest) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

// Platform describes a platform.
type Platform struct {
	// The name of this platform. Uniquely identifies this platform.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// A description of this platform.
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
}

func (m *Platform) Reset()                    { *m = Platform{} }
func (m *Platform) String() string            { return proto.CompactTextString(m) }
func (*Platform) ProtoMessage()               {}
func (*Platform) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *Platform) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Platform) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// ListPlatformsResponse is a response to a request to retrieve platforms.
type ListPlatformsResponse struct {
	// The platforms matching the request.
	Platforms []*Platform `protobuf:"bytes,1,rep,name=platforms" json:"platforms,omitempty"`
}

func (m *ListPlatformsResponse) Reset()                    { *m = ListPlatformsResponse{} }
func (m *ListPlatformsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListPlatformsResponse) ProtoMessage()               {}
func (*ListPlatformsResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

func (m *ListPlatformsResponse) GetPlatforms() []*Platform {
	if m != nil {
		return m.Platforms
	}
	return nil
}

func init() {
	proto.RegisterType((*ListPlatformsRequest)(nil), "crimson.ListPlatformsRequest")
	proto.RegisterType((*Platform)(nil), "crimson.Platform")
	proto.RegisterType((*ListPlatformsResponse)(nil), "crimson.ListPlatformsResponse")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/machine-db/api/crimson/v1/platforms.proto", fileDescriptor5)
}

var fileDescriptor5 = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0xbd, 0x4e, 0xc4, 0x30,
	0x10, 0x84, 0x15, 0xfe, 0xb3, 0x57, 0x61, 0x1d, 0x52, 0xca, 0x28, 0xd5, 0x15, 0x60, 0x0b, 0xe8,
	0x11, 0x25, 0x05, 0x05, 0xf2, 0x1b, 0x38, 0x8e, 0x49, 0x56, 0x8a, 0xbd, 0xc6, 0xeb, 0xf0, 0xfc,
	0x08, 0x93, 0x00, 0xba, 0xce, 0x9e, 0xef, 0x93, 0x66, 0x16, 0x9e, 0x46, 0x92, 0x76, 0x4a, 0xe4,
	0x71, 0xf1, 0x92, 0xd2, 0xa8, 0xe6, 0xc5, 0xa2, 0xf2, 0xc6, 0x4e, 0x18, 0xdc, 0xdd, 0xd0, 0x2b,
	0x13, 0x51, 0xd9, 0x84, 0x9e, 0x29, 0xa8, 0xcf, 0x7b, 0x15, 0x67, 0x93, 0xdf, 0x29, 0x79, 0x96,
	0x31, 0x51, 0x26, 0x71, 0xb9, 0xb2, 0xee, 0x16, 0xf6, 0xaf, 0xc8, 0xf9, 0x6d, 0xe3, 0xda, 0x7d,
	0x2c, 0x8e, 0xb3, 0xd8, 0xc3, 0x79, 0x30, 0xde, 0x71, 0x53, 0xb5, 0xa7, 0x87, 0x5a, 0xff, 0x7c,
	0xba, 0x67, 0xb8, 0xda, 0x4c, 0x21, 0xe0, 0xec, 0x3b, 0x6c, 0xaa, 0xb6, 0x3a, 0xd4, 0xba, 0xbc,
	0x45, 0x0b, 0xbb, 0xc1, 0xb1, 0x4d, 0x18, 0x33, 0x52, 0x68, 0x4e, 0x0a, 0xfa, 0x1f, 0x75, 0x2f,
	0x70, 0x73, 0xd4, 0xc7, 0x91, 0x02, 0x3b, 0xa1, 0xa0, 0xfe, 0x1d, 0x59, 0x4a, 0x77, 0x0f, 0xd7,
	0x72, 0x5d, 0x29, 0x37, 0x5d, 0xff, 0x39, 0xfd, 0x45, 0xb9, 0xe4, 0xf1, 0x2b, 0x00, 0x00, 0xff,
	0xff, 0x7a, 0x71, 0xed, 0x70, 0x0b, 0x01, 0x00, 0x00,
}
