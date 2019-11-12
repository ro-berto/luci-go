// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/resultdb/proto/rpc/v1/resultdb.proto

package rpcpb

import prpc "go.chromium.org/luci/grpc/prpc"

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// A request message for GetInvocation RPC.
type GetInvocationRequest struct {
	// The name of the invocation to request, see Invocation.name.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetInvocationRequest) Reset()         { *m = GetInvocationRequest{} }
func (m *GetInvocationRequest) String() string { return proto.CompactTextString(m) }
func (*GetInvocationRequest) ProtoMessage()    {}
func (*GetInvocationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b48732830a4fcfbd, []int{0}
}

func (m *GetInvocationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetInvocationRequest.Unmarshal(m, b)
}
func (m *GetInvocationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetInvocationRequest.Marshal(b, m, deterministic)
}
func (m *GetInvocationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetInvocationRequest.Merge(m, src)
}
func (m *GetInvocationRequest) XXX_Size() int {
	return xxx_messageInfo_GetInvocationRequest.Size(m)
}
func (m *GetInvocationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetInvocationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetInvocationRequest proto.InternalMessageInfo

func (m *GetInvocationRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// A request message for GetTestResult RPC.
type GetTestResultRequest struct {
	// The name of the test result to request, see TestResult.name.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTestResultRequest) Reset()         { *m = GetTestResultRequest{} }
func (m *GetTestResultRequest) String() string { return proto.CompactTextString(m) }
func (*GetTestResultRequest) ProtoMessage()    {}
func (*GetTestResultRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b48732830a4fcfbd, []int{1}
}

func (m *GetTestResultRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTestResultRequest.Unmarshal(m, b)
}
func (m *GetTestResultRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTestResultRequest.Marshal(b, m, deterministic)
}
func (m *GetTestResultRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTestResultRequest.Merge(m, src)
}
func (m *GetTestResultRequest) XXX_Size() int {
	return xxx_messageInfo_GetTestResultRequest.Size(m)
}
func (m *GetTestResultRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTestResultRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTestResultRequest proto.InternalMessageInfo

func (m *GetTestResultRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// A request message for ListTestResults RPC.
type ListTestResultsRequest struct {
	// Name of the invocation, e.g. "invocations/{id}".
	Invocation string `protobuf:"bytes,1,opt,name=invocation,proto3" json:"invocation,omitempty"`
	// The maximum number of test results to return.
	//
	// The service may return fewer than this value.
	// If unspecified, at most 100 test results will be returned.
	// The maximum value is 1000; values above 1000 will be coerced to 1000.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// A page token, received from a previous `ListTestResults` call.
	// Provide this to retrieve the subsequent page.
	//
	// When paginating, all other parameters provided to `ListTestResults` MUST
	// match the call that provided the page token.
	//
	// Does NOT guarantee returning all test results from chained reads until
	// cursor exhaustion unless the invocation was finalized at the time of first
	// ListTestResults request.
	PageToken            string   `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTestResultsRequest) Reset()         { *m = ListTestResultsRequest{} }
func (m *ListTestResultsRequest) String() string { return proto.CompactTextString(m) }
func (*ListTestResultsRequest) ProtoMessage()    {}
func (*ListTestResultsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b48732830a4fcfbd, []int{2}
}

func (m *ListTestResultsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTestResultsRequest.Unmarshal(m, b)
}
func (m *ListTestResultsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTestResultsRequest.Marshal(b, m, deterministic)
}
func (m *ListTestResultsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTestResultsRequest.Merge(m, src)
}
func (m *ListTestResultsRequest) XXX_Size() int {
	return xxx_messageInfo_ListTestResultsRequest.Size(m)
}
func (m *ListTestResultsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTestResultsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListTestResultsRequest proto.InternalMessageInfo

func (m *ListTestResultsRequest) GetInvocation() string {
	if m != nil {
		return m.Invocation
	}
	return ""
}

func (m *ListTestResultsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListTestResultsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// A response message for ListTestResults RPC.
type ListTestResultsResponse struct {
	// The test results from the specified invocation.
	TestResults []*TestResult `protobuf:"bytes,1,rep,name=test_results,json=testResults,proto3" json:"test_results,omitempty"`
	// A token, which can be sent as `page_token` to retrieve the next page.
	// If this field is omitted, there were no subsequent pages at the time of
	// request.
	// If the invocation is not finalized, more results may appear later.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTestResultsResponse) Reset()         { *m = ListTestResultsResponse{} }
func (m *ListTestResultsResponse) String() string { return proto.CompactTextString(m) }
func (*ListTestResultsResponse) ProtoMessage()    {}
func (*ListTestResultsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b48732830a4fcfbd, []int{3}
}

func (m *ListTestResultsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTestResultsResponse.Unmarshal(m, b)
}
func (m *ListTestResultsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTestResultsResponse.Marshal(b, m, deterministic)
}
func (m *ListTestResultsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTestResultsResponse.Merge(m, src)
}
func (m *ListTestResultsResponse) XXX_Size() int {
	return xxx_messageInfo_ListTestResultsResponse.Size(m)
}
func (m *ListTestResultsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTestResultsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListTestResultsResponse proto.InternalMessageInfo

func (m *ListTestResultsResponse) GetTestResults() []*TestResult {
	if m != nil {
		return m.TestResults
	}
	return nil
}

func (m *ListTestResultsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// A request message for GetTestExoneration RPC.
type GetTestExonerationRequest struct {
	// The name of the test exoneration to request, see TestExoneration.name.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTestExonerationRequest) Reset()         { *m = GetTestExonerationRequest{} }
func (m *GetTestExonerationRequest) String() string { return proto.CompactTextString(m) }
func (*GetTestExonerationRequest) ProtoMessage()    {}
func (*GetTestExonerationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b48732830a4fcfbd, []int{4}
}

func (m *GetTestExonerationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTestExonerationRequest.Unmarshal(m, b)
}
func (m *GetTestExonerationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTestExonerationRequest.Marshal(b, m, deterministic)
}
func (m *GetTestExonerationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTestExonerationRequest.Merge(m, src)
}
func (m *GetTestExonerationRequest) XXX_Size() int {
	return xxx_messageInfo_GetTestExonerationRequest.Size(m)
}
func (m *GetTestExonerationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTestExonerationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTestExonerationRequest proto.InternalMessageInfo

func (m *GetTestExonerationRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// A request message for ListTestExonerations RPC.
type ListTestExonerationsRequest struct {
	// Name of the invocation, e.g. "invocations/{id}".
	Invocation string `protobuf:"bytes,1,opt,name=invocation,proto3" json:"invocation,omitempty"`
	// The maximum number of test exonerations to return.
	//
	// The service may return fewer than this value.
	// If unspecified, at most 100 test exonerations will be returned.
	// The maximum value is 1000; values above 1000 will be coerced to 1000.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// A page token, received from a previous `ListTestExonerations` call.
	// Provide this to retrieve the subsequent page.
	//
	// When paginating, all other parameters provided to `ListTestExonerations`
	// MUST match the call that provided the page token.
	PageToken            string   `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTestExonerationsRequest) Reset()         { *m = ListTestExonerationsRequest{} }
func (m *ListTestExonerationsRequest) String() string { return proto.CompactTextString(m) }
func (*ListTestExonerationsRequest) ProtoMessage()    {}
func (*ListTestExonerationsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b48732830a4fcfbd, []int{5}
}

func (m *ListTestExonerationsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTestExonerationsRequest.Unmarshal(m, b)
}
func (m *ListTestExonerationsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTestExonerationsRequest.Marshal(b, m, deterministic)
}
func (m *ListTestExonerationsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTestExonerationsRequest.Merge(m, src)
}
func (m *ListTestExonerationsRequest) XXX_Size() int {
	return xxx_messageInfo_ListTestExonerationsRequest.Size(m)
}
func (m *ListTestExonerationsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTestExonerationsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListTestExonerationsRequest proto.InternalMessageInfo

func (m *ListTestExonerationsRequest) GetInvocation() string {
	if m != nil {
		return m.Invocation
	}
	return ""
}

func (m *ListTestExonerationsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListTestExonerationsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// A response message for ListTestExonerations RPC.
type ListTestExonerationsResponse struct {
	// The test exonerations from the specified invocation.
	TestExonerations []*TestExoneration `protobuf:"bytes,1,rep,name=test_exonerations,json=testExonerations,proto3" json:"test_exonerations,omitempty"`
	// A token, which can be sent as `page_token` to retrieve the next page.
	// If this field is omitted, there were no subsequent pages at the time of
	// request.
	// If the invocation is not finalized, more results may appear later.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTestExonerationsResponse) Reset()         { *m = ListTestExonerationsResponse{} }
func (m *ListTestExonerationsResponse) String() string { return proto.CompactTextString(m) }
func (*ListTestExonerationsResponse) ProtoMessage()    {}
func (*ListTestExonerationsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b48732830a4fcfbd, []int{6}
}

func (m *ListTestExonerationsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTestExonerationsResponse.Unmarshal(m, b)
}
func (m *ListTestExonerationsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTestExonerationsResponse.Marshal(b, m, deterministic)
}
func (m *ListTestExonerationsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTestExonerationsResponse.Merge(m, src)
}
func (m *ListTestExonerationsResponse) XXX_Size() int {
	return xxx_messageInfo_ListTestExonerationsResponse.Size(m)
}
func (m *ListTestExonerationsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTestExonerationsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListTestExonerationsResponse proto.InternalMessageInfo

func (m *ListTestExonerationsResponse) GetTestExonerations() []*TestExoneration {
	if m != nil {
		return m.TestExonerations
	}
	return nil
}

func (m *ListTestExonerationsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// A request message for QueryTestResults RPC.
type QueryTestResultsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryTestResultsRequest) Reset()         { *m = QueryTestResultsRequest{} }
func (m *QueryTestResultsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryTestResultsRequest) ProtoMessage()    {}
func (*QueryTestResultsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b48732830a4fcfbd, []int{7}
}

func (m *QueryTestResultsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryTestResultsRequest.Unmarshal(m, b)
}
func (m *QueryTestResultsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryTestResultsRequest.Marshal(b, m, deterministic)
}
func (m *QueryTestResultsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryTestResultsRequest.Merge(m, src)
}
func (m *QueryTestResultsRequest) XXX_Size() int {
	return xxx_messageInfo_QueryTestResultsRequest.Size(m)
}
func (m *QueryTestResultsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryTestResultsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryTestResultsRequest proto.InternalMessageInfo

// A response message for QueryTestResults RPC.
type QueryTestResultsResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryTestResultsResponse) Reset()         { *m = QueryTestResultsResponse{} }
func (m *QueryTestResultsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryTestResultsResponse) ProtoMessage()    {}
func (*QueryTestResultsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b48732830a4fcfbd, []int{8}
}

func (m *QueryTestResultsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryTestResultsResponse.Unmarshal(m, b)
}
func (m *QueryTestResultsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryTestResultsResponse.Marshal(b, m, deterministic)
}
func (m *QueryTestResultsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryTestResultsResponse.Merge(m, src)
}
func (m *QueryTestResultsResponse) XXX_Size() int {
	return xxx_messageInfo_QueryTestResultsResponse.Size(m)
}
func (m *QueryTestResultsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryTestResultsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryTestResultsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GetInvocationRequest)(nil), "luci.resultdb.rpc.v1.GetInvocationRequest")
	proto.RegisterType((*GetTestResultRequest)(nil), "luci.resultdb.rpc.v1.GetTestResultRequest")
	proto.RegisterType((*ListTestResultsRequest)(nil), "luci.resultdb.rpc.v1.ListTestResultsRequest")
	proto.RegisterType((*ListTestResultsResponse)(nil), "luci.resultdb.rpc.v1.ListTestResultsResponse")
	proto.RegisterType((*GetTestExonerationRequest)(nil), "luci.resultdb.rpc.v1.GetTestExonerationRequest")
	proto.RegisterType((*ListTestExonerationsRequest)(nil), "luci.resultdb.rpc.v1.ListTestExonerationsRequest")
	proto.RegisterType((*ListTestExonerationsResponse)(nil), "luci.resultdb.rpc.v1.ListTestExonerationsResponse")
	proto.RegisterType((*QueryTestResultsRequest)(nil), "luci.resultdb.rpc.v1.QueryTestResultsRequest")
	proto.RegisterType((*QueryTestResultsResponse)(nil), "luci.resultdb.rpc.v1.QueryTestResultsResponse")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/resultdb/proto/rpc/v1/resultdb.proto", fileDescriptor_b48732830a4fcfbd)
}

var fileDescriptor_b48732830a4fcfbd = []byte{
	// 522 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x8d, 0x13, 0x8a, 0xda, 0x29, 0x55, 0xcb, 0x2a, 0x22, 0xa9, 0x0b, 0x22, 0x32, 0x02, 0x45,
	0x08, 0xbc, 0x24, 0x9c, 0x50, 0xb9, 0x10, 0x40, 0x08, 0x89, 0x03, 0x98, 0x9e, 0xb8, 0x58, 0x8e,
	0x3b, 0xb8, 0x2b, 0x12, 0xaf, 0xd9, 0x5d, 0x47, 0x6d, 0x0f, 0x70, 0xe3, 0x47, 0x70, 0xe6, 0x87,
	0xf1, 0x53, 0x90, 0x3f, 0xe2, 0x6c, 0x63, 0x27, 0x75, 0x2e, 0x5c, 0xdf, 0xce, 0x9b, 0x79, 0xf3,
	0xfc, 0xc6, 0xf0, 0x22, 0xe0, 0xb6, 0x7f, 0x26, 0xf8, 0x94, 0xc5, 0x53, 0x9b, 0x8b, 0x80, 0x4e,
	0x62, 0x9f, 0x51, 0x81, 0x32, 0x9e, 0xa8, 0xd3, 0x31, 0x8d, 0x04, 0x57, 0x9c, 0x8a, 0xc8, 0xa7,
	0xb3, 0x41, 0x81, 0xda, 0x29, 0x4a, 0xda, 0x49, 0xa9, 0x5d, 0x80, 0x22, 0xf2, 0xed, 0xd9, 0xc0,
	0xbc, 0x1f, 0x70, 0x1e, 0x4c, 0x90, 0x7a, 0x11, 0xa3, 0x5f, 0x19, 0x4e, 0x4e, 0xdd, 0x31, 0x9e,
	0x79, 0x33, 0xc6, 0x45, 0x46, 0x33, 0x8f, 0x37, 0x98, 0xc8, 0xc2, 0x19, 0xf7, 0x3d, 0xc5, 0x78,
	0x98, 0x93, 0x5f, 0x6e, 0x40, 0x56, 0x28, 0x95, 0x9b, 0x3d, 0x65, 0x6c, 0x8b, 0x42, 0xfb, 0x1d,
	0xaa, 0xf7, 0x45, 0x53, 0x07, 0xbf, 0xc7, 0x28, 0x15, 0xe9, 0xc0, 0x8d, 0xd0, 0x9b, 0x62, 0xd7,
	0xe8, 0x19, 0xfd, 0x9d, 0x51, 0xeb, 0xef, 0xab, 0xa6, 0x93, 0x02, 0x39, 0xe1, 0x04, 0xa5, 0x72,
	0xd2, 0x3e, 0xd7, 0x12, 0x2e, 0xe0, 0xce, 0x07, 0x26, 0x35, 0x86, 0x9c, 0x53, 0x1e, 0x00, 0x2c,
	0xb6, 0xd1, 0x89, 0x1a, 0x4c, 0x8e, 0x60, 0x27, 0xf2, 0x02, 0x74, 0x25, 0xbb, 0xc4, 0x6e, 0xb3,
	0x67, 0xf4, 0xb7, 0x9c, 0xed, 0x04, 0xf8, 0xcc, 0x2e, 0x91, 0xdc, 0x03, 0x48, 0x1f, 0x15, 0xff,
	0x86, 0x61, 0xb7, 0x95, 0x74, 0x70, 0xd2, 0xf2, 0x93, 0x04, 0xb0, 0x7e, 0x19, 0xd0, 0x29, 0xcd,
	0x96, 0x11, 0x0f, 0x25, 0x92, 0xd7, 0x70, 0x4b, 0x73, 0x43, 0x76, 0x8d, 0x5e, 0xab, 0xbf, 0x3b,
	0xec, 0xd9, 0x55, 0x5f, 0xd0, 0xd6, 0xd6, 0xdd, 0x55, 0x8b, 0x66, 0xe4, 0x11, 0xec, 0x87, 0x78,
	0xae, 0x5c, 0x4d, 0x44, 0x33, 0x15, 0xb1, 0x97, 0xc0, 0x1f, 0x0b, 0x21, 0x14, 0x0e, 0x73, 0xd3,
	0xde, 0x9e, 0xf3, 0x10, 0xc5, 0x15, 0xab, 0x89, 0xee, 0x5c, 0x6e, 0xda, 0x0f, 0x38, 0x9a, 0x0b,
	0xd7, 0x18, 0xff, 0xcf, 0xb9, 0xdf, 0x06, 0xdc, 0xad, 0x16, 0x90, 0xdb, 0xe7, 0xc0, 0xed, 0xd4,
	0x3e, 0xd4, 0x1e, 0x73, 0x0f, 0x1f, 0xae, 0xf6, 0x50, 0xdf, 0xfe, 0x40, 0x2d, 0xf5, 0xae, 0xed,
	0xe6, 0x21, 0x74, 0x3e, 0xc5, 0x28, 0x2e, 0xca, 0x91, 0xb2, 0x4c, 0xe8, 0x96, 0x9f, 0x32, 0xc9,
	0xc3, 0x3f, 0x5b, 0xb0, 0x9d, 0x61, 0x6f, 0x46, 0xc4, 0x85, 0xbd, 0x2b, 0xb9, 0x27, 0x8f, 0xab,
	0x55, 0x57, 0x1d, 0x87, 0xb9, 0x22, 0x25, 0x8b, 0x42, 0xab, 0x91, 0x0f, 0x58, 0xe8, 0x58, 0x33,
	0xa0, 0x74, 0x4c, 0xe6, 0xb5, 0x31, 0xb4, 0x1a, 0x24, 0x82, 0xfd, 0xa5, 0x6c, 0x93, 0x27, 0xd5,
	0xb4, 0xea, 0xf3, 0x33, 0x9f, 0xd6, 0xac, 0xce, 0xec, 0xb3, 0x1a, 0x24, 0x04, 0x52, 0x4e, 0x31,
	0xa1, 0x6b, 0xf7, 0x2a, 0xe7, 0xdd, 0xac, 0x97, 0x0f, 0xab, 0x41, 0x7e, 0x42, 0xbb, 0x2a, 0x83,
	0x64, 0xb0, 0x5e, 0x78, 0xc5, 0xc1, 0x98, 0xc3, 0x4d, 0x28, 0xc5, 0xc2, 0x12, 0x0e, 0x96, 0xd3,
	0x44, 0x56, 0xb8, 0xb6, 0x22, 0x90, 0xa6, 0x5d, 0xb7, 0x7c, 0x3e, 0x74, 0x34, 0xfc, 0xf2, 0xac,
	0xfe, 0x1f, 0xfd, 0x58, 0x44, 0x7e, 0x34, 0x1e, 0xdf, 0x4c, 0xb1, 0xe7, 0xff, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x14, 0xb1, 0x0d, 0xd7, 0xbb, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ResultDBClient is the client API for ResultDB service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ResultDBClient interface {
	// Retrieves an invocation.
	GetInvocation(ctx context.Context, in *GetInvocationRequest, opts ...grpc.CallOption) (*Invocation, error)
	// Retrieves a test result.
	GetTestResult(ctx context.Context, in *GetTestResultRequest, opts ...grpc.CallOption) (*TestResult, error)
	// Retrieves test results for a parent invocation.
	//
	// Note: response does not contain test results of included invocations.
	// Use QueryTestResults instead.
	ListTestResults(ctx context.Context, in *ListTestResultsRequest, opts ...grpc.CallOption) (*ListTestResultsResponse, error)
	// Retrieves a test exoneration.
	GetTestExoneration(ctx context.Context, in *GetTestExonerationRequest, opts ...grpc.CallOption) (*TestExoneration, error)
	// Retrieves test exonerations for a parent invocation.
	// Note: response does not contain test results of included invocations.
	ListTestExonerations(ctx context.Context, in *ListTestExonerationsRequest, opts ...grpc.CallOption) (*ListTestExonerationsResponse, error)
	// Retrieves test results from an invocation.
	// Supports invocation inclusions.
	// Supports filtering based on "TestResult.expected", exonerations, test
	// path and variant.
	QueryTestResults(ctx context.Context, in *QueryTestResultsRequest, opts ...grpc.CallOption) (*QueryTestResultsResponse, error)
}
type resultDBPRPCClient struct {
	client *prpc.Client
}

func NewResultDBPRPCClient(client *prpc.Client) ResultDBClient {
	return &resultDBPRPCClient{client}
}

func (c *resultDBPRPCClient) GetInvocation(ctx context.Context, in *GetInvocationRequest, opts ...grpc.CallOption) (*Invocation, error) {
	out := new(Invocation)
	err := c.client.Call(ctx, "luci.resultdb.rpc.v1.ResultDB", "GetInvocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resultDBPRPCClient) GetTestResult(ctx context.Context, in *GetTestResultRequest, opts ...grpc.CallOption) (*TestResult, error) {
	out := new(TestResult)
	err := c.client.Call(ctx, "luci.resultdb.rpc.v1.ResultDB", "GetTestResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resultDBPRPCClient) ListTestResults(ctx context.Context, in *ListTestResultsRequest, opts ...grpc.CallOption) (*ListTestResultsResponse, error) {
	out := new(ListTestResultsResponse)
	err := c.client.Call(ctx, "luci.resultdb.rpc.v1.ResultDB", "ListTestResults", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resultDBPRPCClient) GetTestExoneration(ctx context.Context, in *GetTestExonerationRequest, opts ...grpc.CallOption) (*TestExoneration, error) {
	out := new(TestExoneration)
	err := c.client.Call(ctx, "luci.resultdb.rpc.v1.ResultDB", "GetTestExoneration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resultDBPRPCClient) ListTestExonerations(ctx context.Context, in *ListTestExonerationsRequest, opts ...grpc.CallOption) (*ListTestExonerationsResponse, error) {
	out := new(ListTestExonerationsResponse)
	err := c.client.Call(ctx, "luci.resultdb.rpc.v1.ResultDB", "ListTestExonerations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resultDBPRPCClient) QueryTestResults(ctx context.Context, in *QueryTestResultsRequest, opts ...grpc.CallOption) (*QueryTestResultsResponse, error) {
	out := new(QueryTestResultsResponse)
	err := c.client.Call(ctx, "luci.resultdb.rpc.v1.ResultDB", "QueryTestResults", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type resultDBClient struct {
	cc *grpc.ClientConn
}

func NewResultDBClient(cc *grpc.ClientConn) ResultDBClient {
	return &resultDBClient{cc}
}

func (c *resultDBClient) GetInvocation(ctx context.Context, in *GetInvocationRequest, opts ...grpc.CallOption) (*Invocation, error) {
	out := new(Invocation)
	err := c.cc.Invoke(ctx, "/luci.resultdb.rpc.v1.ResultDB/GetInvocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resultDBClient) GetTestResult(ctx context.Context, in *GetTestResultRequest, opts ...grpc.CallOption) (*TestResult, error) {
	out := new(TestResult)
	err := c.cc.Invoke(ctx, "/luci.resultdb.rpc.v1.ResultDB/GetTestResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resultDBClient) ListTestResults(ctx context.Context, in *ListTestResultsRequest, opts ...grpc.CallOption) (*ListTestResultsResponse, error) {
	out := new(ListTestResultsResponse)
	err := c.cc.Invoke(ctx, "/luci.resultdb.rpc.v1.ResultDB/ListTestResults", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resultDBClient) GetTestExoneration(ctx context.Context, in *GetTestExonerationRequest, opts ...grpc.CallOption) (*TestExoneration, error) {
	out := new(TestExoneration)
	err := c.cc.Invoke(ctx, "/luci.resultdb.rpc.v1.ResultDB/GetTestExoneration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resultDBClient) ListTestExonerations(ctx context.Context, in *ListTestExonerationsRequest, opts ...grpc.CallOption) (*ListTestExonerationsResponse, error) {
	out := new(ListTestExonerationsResponse)
	err := c.cc.Invoke(ctx, "/luci.resultdb.rpc.v1.ResultDB/ListTestExonerations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resultDBClient) QueryTestResults(ctx context.Context, in *QueryTestResultsRequest, opts ...grpc.CallOption) (*QueryTestResultsResponse, error) {
	out := new(QueryTestResultsResponse)
	err := c.cc.Invoke(ctx, "/luci.resultdb.rpc.v1.ResultDB/QueryTestResults", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResultDBServer is the server API for ResultDB service.
type ResultDBServer interface {
	// Retrieves an invocation.
	GetInvocation(context.Context, *GetInvocationRequest) (*Invocation, error)
	// Retrieves a test result.
	GetTestResult(context.Context, *GetTestResultRequest) (*TestResult, error)
	// Retrieves test results for a parent invocation.
	//
	// Note: response does not contain test results of included invocations.
	// Use QueryTestResults instead.
	ListTestResults(context.Context, *ListTestResultsRequest) (*ListTestResultsResponse, error)
	// Retrieves a test exoneration.
	GetTestExoneration(context.Context, *GetTestExonerationRequest) (*TestExoneration, error)
	// Retrieves test exonerations for a parent invocation.
	// Note: response does not contain test results of included invocations.
	ListTestExonerations(context.Context, *ListTestExonerationsRequest) (*ListTestExonerationsResponse, error)
	// Retrieves test results from an invocation.
	// Supports invocation inclusions.
	// Supports filtering based on "TestResult.expected", exonerations, test
	// path and variant.
	QueryTestResults(context.Context, *QueryTestResultsRequest) (*QueryTestResultsResponse, error)
}

// UnimplementedResultDBServer can be embedded to have forward compatible implementations.
type UnimplementedResultDBServer struct {
}

func (*UnimplementedResultDBServer) GetInvocation(ctx context.Context, req *GetInvocationRequest) (*Invocation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInvocation not implemented")
}
func (*UnimplementedResultDBServer) GetTestResult(ctx context.Context, req *GetTestResultRequest) (*TestResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTestResult not implemented")
}
func (*UnimplementedResultDBServer) ListTestResults(ctx context.Context, req *ListTestResultsRequest) (*ListTestResultsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTestResults not implemented")
}
func (*UnimplementedResultDBServer) GetTestExoneration(ctx context.Context, req *GetTestExonerationRequest) (*TestExoneration, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTestExoneration not implemented")
}
func (*UnimplementedResultDBServer) ListTestExonerations(ctx context.Context, req *ListTestExonerationsRequest) (*ListTestExonerationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTestExonerations not implemented")
}
func (*UnimplementedResultDBServer) QueryTestResults(ctx context.Context, req *QueryTestResultsRequest) (*QueryTestResultsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryTestResults not implemented")
}

func RegisterResultDBServer(s prpc.Registrar, srv ResultDBServer) {
	s.RegisterService(&_ResultDB_serviceDesc, srv)
}

func _ResultDB_GetInvocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInvocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResultDBServer).GetInvocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/luci.resultdb.rpc.v1.ResultDB/GetInvocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResultDBServer).GetInvocation(ctx, req.(*GetInvocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResultDB_GetTestResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTestResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResultDBServer).GetTestResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/luci.resultdb.rpc.v1.ResultDB/GetTestResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResultDBServer).GetTestResult(ctx, req.(*GetTestResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResultDB_ListTestResults_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTestResultsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResultDBServer).ListTestResults(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/luci.resultdb.rpc.v1.ResultDB/ListTestResults",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResultDBServer).ListTestResults(ctx, req.(*ListTestResultsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResultDB_GetTestExoneration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTestExonerationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResultDBServer).GetTestExoneration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/luci.resultdb.rpc.v1.ResultDB/GetTestExoneration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResultDBServer).GetTestExoneration(ctx, req.(*GetTestExonerationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResultDB_ListTestExonerations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTestExonerationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResultDBServer).ListTestExonerations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/luci.resultdb.rpc.v1.ResultDB/ListTestExonerations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResultDBServer).ListTestExonerations(ctx, req.(*ListTestExonerationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResultDB_QueryTestResults_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTestResultsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResultDBServer).QueryTestResults(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/luci.resultdb.rpc.v1.ResultDB/QueryTestResults",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResultDBServer).QueryTestResults(ctx, req.(*QueryTestResultsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ResultDB_serviceDesc = grpc.ServiceDesc{
	ServiceName: "luci.resultdb.rpc.v1.ResultDB",
	HandlerType: (*ResultDBServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetInvocation",
			Handler:    _ResultDB_GetInvocation_Handler,
		},
		{
			MethodName: "GetTestResult",
			Handler:    _ResultDB_GetTestResult_Handler,
		},
		{
			MethodName: "ListTestResults",
			Handler:    _ResultDB_ListTestResults_Handler,
		},
		{
			MethodName: "GetTestExoneration",
			Handler:    _ResultDB_GetTestExoneration_Handler,
		},
		{
			MethodName: "ListTestExonerations",
			Handler:    _ResultDB_ListTestExonerations_Handler,
		},
		{
			MethodName: "QueryTestResults",
			Handler:    _ResultDB_QueryTestResults_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go.chromium.org/luci/resultdb/proto/rpc/v1/resultdb.proto",
}
