// Copyright 2017 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.1
// source: go.chromium.org/luci/vpython/api/vpython/spec.proto

package vpython

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

// Spec is a "vpython" environment specification.
type Spec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Python version to use. This should be of the form:
	// "Major[.Minor[.Patch]]"
	//
	// If specified,
	// - The Major version will be enforced absolutely. Python 3 will not be
	//   preferred over Python 2 because '3' is greater than '2'.
	// - The remaining versions, if specified, will be regarded as *minimum*
	//   versions. In other words, if "2.7.4" is specified and the system has
	//   "2.7.12", that will suffice. Similarly, "2.6" would accept a "2.7"
	//   interpreter.
	//
	// If empty, the default Python interpreter ("python") will be used.
	PythonVersion string          `protobuf:"bytes,1,opt,name=python_version,json=pythonVersion,proto3" json:"python_version,omitempty"`
	Wheel         []*Spec_Package `protobuf:"bytes,2,rep,name=wheel,proto3" json:"wheel,omitempty"`
	// The VirtualEnv package.
	//
	// This should be left empty to use the `vpython` default package
	// (recommended).
	Virtualenv *Spec_Package `protobuf:"bytes,3,opt,name=virtualenv,proto3" json:"virtualenv,omitempty"`
	// Specification-provided PEP425 verification tags.
	//
	// By default, verification will be performed against a default set of
	// environment parameters. However, a given specification may offer its own
	// set of PEP425 tags representing the systems that it wants to be verified
	// against.
	VerifyPep425Tag []*PEP425Tag `protobuf:"bytes,4,rep,name=verify_pep425_tag,json=verifyPep425Tag,proto3" json:"verify_pep425_tag,omitempty"`
}

func (x *Spec) Reset() {
	*x = Spec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_vpython_api_vpython_spec_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Spec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Spec) ProtoMessage() {}

func (x *Spec) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_vpython_api_vpython_spec_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Spec.ProtoReflect.Descriptor instead.
func (*Spec) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_vpython_api_vpython_spec_proto_rawDescGZIP(), []int{0}
}

func (x *Spec) GetPythonVersion() string {
	if x != nil {
		return x.PythonVersion
	}
	return ""
}

func (x *Spec) GetWheel() []*Spec_Package {
	if x != nil {
		return x.Wheel
	}
	return nil
}

func (x *Spec) GetVirtualenv() *Spec_Package {
	if x != nil {
		return x.Virtualenv
	}
	return nil
}

func (x *Spec) GetVerifyPep425Tag() []*PEP425Tag {
	if x != nil {
		return x.VerifyPep425Tag
	}
	return nil
}

// A definition for a remote package. The type of package depends on the
// configured package resolver.
type Spec_Package struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the package.
	//
	// - For CIPD, this is the package name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The package version.
	//
	// - For CIPD, this will be any recognized CIPD version (i.e., ID, tag, or
	//   ref).
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// Optional PEP425 tags to determine whether this package is included on the
	// target system. If no match tags are specified, this package will always
	// be included. If match tags are specified, the package will be included if
	// any system PEP425 tags match at least one of the match tags.
	//
	// A match will succeed if any system PEP425 tag field matches the
	// corresponding field in the PEP425 tag. If the match tag omits a field
	// (partial), that field will not be considered. For example, if a match
	// tag specifies just an ABI field, any system PEP425 tag with that ABI will
	// be considered a successful match, regardless of other field values.
	MatchTag []*PEP425Tag `protobuf:"bytes,3,rep,name=match_tag,json=matchTag,proto3" json:"match_tag,omitempty"`
	// Optional PEP425 tags to determine whether this package is NOT included on
	// the target system. This has the opposite behavior as "match_tag": if any
	// host tags match any tags in this list, the package will not be installed
	// on this host.
	//
	// A "not_match_tag" overrides a "match_tag", so if a host has tags that
	// match entries in both, the package will be not considered a match.
	NotMatchTag []*PEP425Tag `protobuf:"bytes,4,rep,name=not_match_tag,json=notMatchTag,proto3" json:"not_match_tag,omitempty"`
}

func (x *Spec_Package) Reset() {
	*x = Spec_Package{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_vpython_api_vpython_spec_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Spec_Package) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Spec_Package) ProtoMessage() {}

func (x *Spec_Package) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_vpython_api_vpython_spec_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Spec_Package.ProtoReflect.Descriptor instead.
func (*Spec_Package) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_vpython_api_vpython_spec_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Spec_Package) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Spec_Package) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Spec_Package) GetMatchTag() []*PEP425Tag {
	if x != nil {
		return x.MatchTag
	}
	return nil
}

func (x *Spec_Package) GetNotMatchTag() []*PEP425Tag {
	if x != nil {
		return x.NotMatchTag
	}
	return nil
}

var File_go_chromium_org_luci_vpython_api_vpython_spec_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_vpython_api_vpython_spec_proto_rawDesc = []byte{
	0x0a, 0x33, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x76, 0x70, 0x79, 0x74, 0x68, 0x6f, 0x6e, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x70, 0x79, 0x74, 0x68, 0x6f, 0x6e, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x76, 0x70, 0x79, 0x74, 0x68, 0x6f, 0x6e, 0x1a, 0x35,
	0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x6c, 0x75, 0x63, 0x69, 0x2f, 0x76, 0x70, 0x79, 0x74, 0x68, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x70, 0x79, 0x74, 0x68, 0x6f, 0x6e, 0x2f, 0x70, 0x65, 0x70, 0x34, 0x32, 0x35, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf4, 0x02, 0x0a, 0x04, 0x53, 0x70, 0x65, 0x63, 0x12, 0x25,
	0x0a, 0x0e, 0x70, 0x79, 0x74, 0x68, 0x6f, 0x6e, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x79, 0x74, 0x68, 0x6f, 0x6e, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x05, 0x77, 0x68, 0x65, 0x65, 0x6c, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x76, 0x70, 0x79, 0x74, 0x68, 0x6f, 0x6e, 0x2e, 0x53,
	0x70, 0x65, 0x63, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x05, 0x77, 0x68, 0x65,
	0x65, 0x6c, 0x12, 0x35, 0x0a, 0x0a, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x65, 0x6e, 0x76,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x76, 0x70, 0x79, 0x74, 0x68, 0x6f, 0x6e,
	0x2e, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x0a, 0x76,
	0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x65, 0x6e, 0x76, 0x12, 0x3e, 0x0a, 0x11, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x5f, 0x70, 0x65, 0x70, 0x34, 0x32, 0x35, 0x5f, 0x74, 0x61, 0x67, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x76, 0x70, 0x79, 0x74, 0x68, 0x6f, 0x6e, 0x2e, 0x50,
	0x45, 0x50, 0x34, 0x32, 0x35, 0x54, 0x61, 0x67, 0x52, 0x0f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x50, 0x65, 0x70, 0x34, 0x32, 0x35, 0x54, 0x61, 0x67, 0x1a, 0xa0, 0x01, 0x0a, 0x07, 0x50, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x09, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x74, 0x61, 0x67,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x76, 0x70, 0x79, 0x74, 0x68, 0x6f, 0x6e,
	0x2e, 0x50, 0x45, 0x50, 0x34, 0x32, 0x35, 0x54, 0x61, 0x67, 0x52, 0x08, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x54, 0x61, 0x67, 0x12, 0x36, 0x0a, 0x0d, 0x6e, 0x6f, 0x74, 0x5f, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x5f, 0x74, 0x61, 0x67, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x76, 0x70,
	0x79, 0x74, 0x68, 0x6f, 0x6e, 0x2e, 0x50, 0x45, 0x50, 0x34, 0x32, 0x35, 0x54, 0x61, 0x67, 0x52,
	0x0b, 0x6e, 0x6f, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x61, 0x67, 0x42, 0x2a, 0x5a, 0x28,
	0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x6c, 0x75, 0x63, 0x69, 0x2f, 0x76, 0x70, 0x79, 0x74, 0x68, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x70, 0x79, 0x74, 0x68, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_vpython_api_vpython_spec_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_vpython_api_vpython_spec_proto_rawDescData = file_go_chromium_org_luci_vpython_api_vpython_spec_proto_rawDesc
)

func file_go_chromium_org_luci_vpython_api_vpython_spec_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_vpython_api_vpython_spec_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_vpython_api_vpython_spec_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_vpython_api_vpython_spec_proto_rawDescData)
	})
	return file_go_chromium_org_luci_vpython_api_vpython_spec_proto_rawDescData
}

var file_go_chromium_org_luci_vpython_api_vpython_spec_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_go_chromium_org_luci_vpython_api_vpython_spec_proto_goTypes = []interface{}{
	(*Spec)(nil),         // 0: vpython.Spec
	(*Spec_Package)(nil), // 1: vpython.Spec.Package
	(*PEP425Tag)(nil),    // 2: vpython.PEP425Tag
}
var file_go_chromium_org_luci_vpython_api_vpython_spec_proto_depIdxs = []int32{
	1, // 0: vpython.Spec.wheel:type_name -> vpython.Spec.Package
	1, // 1: vpython.Spec.virtualenv:type_name -> vpython.Spec.Package
	2, // 2: vpython.Spec.verify_pep425_tag:type_name -> vpython.PEP425Tag
	2, // 3: vpython.Spec.Package.match_tag:type_name -> vpython.PEP425Tag
	2, // 4: vpython.Spec.Package.not_match_tag:type_name -> vpython.PEP425Tag
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_vpython_api_vpython_spec_proto_init() }
func file_go_chromium_org_luci_vpython_api_vpython_spec_proto_init() {
	if File_go_chromium_org_luci_vpython_api_vpython_spec_proto != nil {
		return
	}
	file_go_chromium_org_luci_vpython_api_vpython_pep425_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_vpython_api_vpython_spec_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Spec); i {
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
		file_go_chromium_org_luci_vpython_api_vpython_spec_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Spec_Package); i {
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
			RawDescriptor: file_go_chromium_org_luci_vpython_api_vpython_spec_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_vpython_api_vpython_spec_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_vpython_api_vpython_spec_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_vpython_api_vpython_spec_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_vpython_api_vpython_spec_proto = out.File
	file_go_chromium_org_luci_vpython_api_vpython_spec_proto_rawDesc = nil
	file_go_chromium_org_luci_vpython_api_vpython_spec_proto_goTypes = nil
	file_go_chromium_org_luci_vpython_api_vpython_spec_proto_depIdxs = nil
}
