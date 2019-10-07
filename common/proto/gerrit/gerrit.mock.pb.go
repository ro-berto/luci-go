// Code generated by MockGen. DO NOT EDIT.
// Source: gerrit.pb.go

// Package gerrit is a generated GoMock package.
package gerrit

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockGerritClient is a mock of GerritClient interface
type MockGerritClient struct {
	ctrl     *gomock.Controller
	recorder *MockGerritClientMockRecorder
}

// MockGerritClientMockRecorder is the mock recorder for MockGerritClient
type MockGerritClientMockRecorder struct {
	mock *MockGerritClient
}

// NewMockGerritClient creates a new mock instance
func NewMockGerritClient(ctrl *gomock.Controller) *MockGerritClient {
	mock := &MockGerritClient{ctrl: ctrl}
	mock.recorder = &MockGerritClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGerritClient) EXPECT() *MockGerritClientMockRecorder {
	return m.recorder
}

// GetChange mocks base method
func (m *MockGerritClient) GetChange(ctx context.Context, in *GetChangeRequest, opts ...grpc.CallOption) (*ChangeInfo, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetChange", varargs...)
	ret0, _ := ret[0].(*ChangeInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChange indicates an expected call of GetChange
func (mr *MockGerritClientMockRecorder) GetChange(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChange", reflect.TypeOf((*MockGerritClient)(nil).GetChange), varargs...)
}

// CreateChange mocks base method
func (m *MockGerritClient) CreateChange(ctx context.Context, in *CreateChangeRequest, opts ...grpc.CallOption) (*ChangeInfo, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateChange", varargs...)
	ret0, _ := ret[0].(*ChangeInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateChange indicates an expected call of CreateChange
func (mr *MockGerritClientMockRecorder) CreateChange(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateChange", reflect.TypeOf((*MockGerritClient)(nil).CreateChange), varargs...)
}

// ChangeEditFileContent mocks base method
func (m *MockGerritClient) ChangeEditFileContent(ctx context.Context, in *ChangeEditFileContentRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ChangeEditFileContent", varargs...)
	ret0, _ := ret[0].(*empty.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeEditFileContent indicates an expected call of ChangeEditFileContent
func (mr *MockGerritClientMockRecorder) ChangeEditFileContent(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeEditFileContent", reflect.TypeOf((*MockGerritClient)(nil).ChangeEditFileContent), varargs...)
}

// DeleteEditFileContent mocks base method
func (m *MockGerritClient) DeleteEditFileContent(ctx context.Context, in *DeleteEditFileContentRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteEditFileContent", varargs...)
	ret0, _ := ret[0].(*empty.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteEditFileContent indicates an expected call of DeleteEditFileContent
func (mr *MockGerritClientMockRecorder) DeleteEditFileContent(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEditFileContent", reflect.TypeOf((*MockGerritClient)(nil).DeleteEditFileContent), varargs...)
}

// ChangeEditPublish mocks base method
func (m *MockGerritClient) ChangeEditPublish(ctx context.Context, in *ChangeEditPublishRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ChangeEditPublish", varargs...)
	ret0, _ := ret[0].(*empty.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeEditPublish indicates an expected call of ChangeEditPublish
func (mr *MockGerritClientMockRecorder) ChangeEditPublish(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeEditPublish", reflect.TypeOf((*MockGerritClient)(nil).ChangeEditPublish), varargs...)
}

// SetReview mocks base method
func (m *MockGerritClient) SetReview(ctx context.Context, in *SetReviewRequest, opts ...grpc.CallOption) (*ReviewResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetReview", varargs...)
	ret0, _ := ret[0].(*ReviewResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetReview indicates an expected call of SetReview
func (mr *MockGerritClientMockRecorder) SetReview(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetReview", reflect.TypeOf((*MockGerritClient)(nil).SetReview), varargs...)
}

// SubmitChange mocks base method
func (m *MockGerritClient) SubmitChange(ctx context.Context, in *SubmitChangeRequest, opts ...grpc.CallOption) (*ChangeInfo, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SubmitChange", varargs...)
	ret0, _ := ret[0].(*ChangeInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubmitChange indicates an expected call of SubmitChange
func (mr *MockGerritClientMockRecorder) SubmitChange(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitChange", reflect.TypeOf((*MockGerritClient)(nil).SubmitChange), varargs...)
}

// AbandonChange mocks base method
func (m *MockGerritClient) AbandonChange(ctx context.Context, in *AbandonChangeRequest, opts ...grpc.CallOption) (*ChangeInfo, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AbandonChange", varargs...)
	ret0, _ := ret[0].(*ChangeInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AbandonChange indicates an expected call of AbandonChange
func (mr *MockGerritClientMockRecorder) AbandonChange(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AbandonChange", reflect.TypeOf((*MockGerritClient)(nil).AbandonChange), varargs...)
}

// GetMergeable mocks base method
func (m *MockGerritClient) GetMergeable(ctx context.Context, in *GetMergeableRequest, opts ...grpc.CallOption) (*MergeableInfo, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMergeable", varargs...)
	ret0, _ := ret[0].(*MergeableInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMergeable indicates an expected call of GetMergeable
func (mr *MockGerritClientMockRecorder) GetMergeable(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMergeable", reflect.TypeOf((*MockGerritClient)(nil).GetMergeable), varargs...)
}

// MockGerritServer is a mock of GerritServer interface
type MockGerritServer struct {
	ctrl     *gomock.Controller
	recorder *MockGerritServerMockRecorder
}

// MockGerritServerMockRecorder is the mock recorder for MockGerritServer
type MockGerritServerMockRecorder struct {
	mock *MockGerritServer
}

// NewMockGerritServer creates a new mock instance
func NewMockGerritServer(ctrl *gomock.Controller) *MockGerritServer {
	mock := &MockGerritServer{ctrl: ctrl}
	mock.recorder = &MockGerritServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGerritServer) EXPECT() *MockGerritServerMockRecorder {
	return m.recorder
}

// GetChange mocks base method
func (m *MockGerritServer) GetChange(arg0 context.Context, arg1 *GetChangeRequest) (*ChangeInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChange", arg0, arg1)
	ret0, _ := ret[0].(*ChangeInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChange indicates an expected call of GetChange
func (mr *MockGerritServerMockRecorder) GetChange(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChange", reflect.TypeOf((*MockGerritServer)(nil).GetChange), arg0, arg1)
}

// CreateChange mocks base method
func (m *MockGerritServer) CreateChange(arg0 context.Context, arg1 *CreateChangeRequest) (*ChangeInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateChange", arg0, arg1)
	ret0, _ := ret[0].(*ChangeInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateChange indicates an expected call of CreateChange
func (mr *MockGerritServerMockRecorder) CreateChange(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateChange", reflect.TypeOf((*MockGerritServer)(nil).CreateChange), arg0, arg1)
}

// ChangeEditFileContent mocks base method
func (m *MockGerritServer) ChangeEditFileContent(arg0 context.Context, arg1 *ChangeEditFileContentRequest) (*empty.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeEditFileContent", arg0, arg1)
	ret0, _ := ret[0].(*empty.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeEditFileContent indicates an expected call of ChangeEditFileContent
func (mr *MockGerritServerMockRecorder) ChangeEditFileContent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeEditFileContent", reflect.TypeOf((*MockGerritServer)(nil).ChangeEditFileContent), arg0, arg1)
}

// DeleteEditFileContent mocks base method
func (m *MockGerritServer) DeleteEditFileContent(arg0 context.Context, arg1 *DeleteEditFileContentRequest) (*empty.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEditFileContent", arg0, arg1)
	ret0, _ := ret[0].(*empty.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteEditFileContent indicates an expected call of DeleteEditFileContent
func (mr *MockGerritServerMockRecorder) DeleteEditFileContent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEditFileContent", reflect.TypeOf((*MockGerritServer)(nil).DeleteEditFileContent), arg0, arg1)
}

// ChangeEditPublish mocks base method
func (m *MockGerritServer) ChangeEditPublish(arg0 context.Context, arg1 *ChangeEditPublishRequest) (*empty.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeEditPublish", arg0, arg1)
	ret0, _ := ret[0].(*empty.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeEditPublish indicates an expected call of ChangeEditPublish
func (mr *MockGerritServerMockRecorder) ChangeEditPublish(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeEditPublish", reflect.TypeOf((*MockGerritServer)(nil).ChangeEditPublish), arg0, arg1)
}

// SetReview mocks base method
func (m *MockGerritServer) SetReview(arg0 context.Context, arg1 *SetReviewRequest) (*ReviewResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetReview", arg0, arg1)
	ret0, _ := ret[0].(*ReviewResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetReview indicates an expected call of SetReview
func (mr *MockGerritServerMockRecorder) SetReview(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetReview", reflect.TypeOf((*MockGerritServer)(nil).SetReview), arg0, arg1)
}

// SubmitChange mocks base method
func (m *MockGerritServer) SubmitChange(arg0 context.Context, arg1 *SubmitChangeRequest) (*ChangeInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubmitChange", arg0, arg1)
	ret0, _ := ret[0].(*ChangeInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubmitChange indicates an expected call of SubmitChange
func (mr *MockGerritServerMockRecorder) SubmitChange(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitChange", reflect.TypeOf((*MockGerritServer)(nil).SubmitChange), arg0, arg1)
}

// AbandonChange mocks base method
func (m *MockGerritServer) AbandonChange(arg0 context.Context, arg1 *AbandonChangeRequest) (*ChangeInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AbandonChange", arg0, arg1)
	ret0, _ := ret[0].(*ChangeInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AbandonChange indicates an expected call of AbandonChange
func (mr *MockGerritServerMockRecorder) AbandonChange(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AbandonChange", reflect.TypeOf((*MockGerritServer)(nil).AbandonChange), arg0, arg1)
}

// GetMergeable mocks base method
func (m *MockGerritServer) GetMergeable(arg0 context.Context, arg1 *GetMergeableRequest) (*MergeableInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMergeable", arg0, arg1)
	ret0, _ := ret[0].(*MergeableInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMergeable indicates an expected call of GetMergeable
func (mr *MockGerritServerMockRecorder) GetMergeable(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMergeable", reflect.TypeOf((*MockGerritServer)(nil).GetMergeable), arg0, arg1)
}
