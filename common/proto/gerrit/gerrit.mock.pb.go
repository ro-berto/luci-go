// Code generated by MockGen. DO NOT EDIT.
// Source: gerrit.pb.go

// Package gerrit is a generated GoMock package.
package gerrit

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

// MockGerritClient is a mock of GerritClient interface.
type MockGerritClient struct {
	ctrl     *gomock.Controller
	recorder *MockGerritClientMockRecorder
}

// MockGerritClientMockRecorder is the mock recorder for MockGerritClient.
type MockGerritClientMockRecorder struct {
	mock *MockGerritClient
}

// NewMockGerritClient creates a new mock instance.
func NewMockGerritClient(ctrl *gomock.Controller) *MockGerritClient {
	mock := &MockGerritClient{ctrl: ctrl}
	mock.recorder = &MockGerritClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGerritClient) EXPECT() *MockGerritClientMockRecorder {
	return m.recorder
}

// ListProjects mocks base method.
func (m *MockGerritClient) ListProjects(ctx context.Context, in *ListProjectsRequest, opts ...grpc.CallOption) (*ListProjectsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListProjects", varargs...)
	ret0, _ := ret[0].(*ListProjectsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProjects indicates an expected call of ListProjects.
func (mr *MockGerritClientMockRecorder) ListProjects(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProjects", reflect.TypeOf((*MockGerritClient)(nil).ListProjects), varargs...)
}

// GetRefInfo mocks base method.
func (m *MockGerritClient) GetRefInfo(ctx context.Context, in *RefInfoRequest, opts ...grpc.CallOption) (*RefInfo, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRefInfo", varargs...)
	ret0, _ := ret[0].(*RefInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRefInfo indicates an expected call of GetRefInfo.
func (mr *MockGerritClientMockRecorder) GetRefInfo(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRefInfo", reflect.TypeOf((*MockGerritClient)(nil).GetRefInfo), varargs...)
}

// ListFileOwners mocks base method.
func (m *MockGerritClient) ListFileOwners(ctx context.Context, in *ListFileOwnersRequest, opts ...grpc.CallOption) (*ListOwnersResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListFileOwners", varargs...)
	ret0, _ := ret[0].(*ListOwnersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFileOwners indicates an expected call of ListFileOwners.
func (mr *MockGerritClientMockRecorder) ListFileOwners(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFileOwners", reflect.TypeOf((*MockGerritClient)(nil).ListFileOwners), varargs...)
}

// ListChanges mocks base method.
func (m *MockGerritClient) ListChanges(ctx context.Context, in *ListChangesRequest, opts ...grpc.CallOption) (*ListChangesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListChanges", varargs...)
	ret0, _ := ret[0].(*ListChangesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListChanges indicates an expected call of ListChanges.
func (mr *MockGerritClientMockRecorder) ListChanges(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListChanges", reflect.TypeOf((*MockGerritClient)(nil).ListChanges), varargs...)
}

// GetChange mocks base method.
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

// GetChange indicates an expected call of GetChange.
func (mr *MockGerritClientMockRecorder) GetChange(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChange", reflect.TypeOf((*MockGerritClient)(nil).GetChange), varargs...)
}

// GetMergeable mocks base method.
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

// GetMergeable indicates an expected call of GetMergeable.
func (mr *MockGerritClientMockRecorder) GetMergeable(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMergeable", reflect.TypeOf((*MockGerritClient)(nil).GetMergeable), varargs...)
}

// ListFiles mocks base method.
func (m *MockGerritClient) ListFiles(ctx context.Context, in *ListFilesRequest, opts ...grpc.CallOption) (*ListFilesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListFiles", varargs...)
	ret0, _ := ret[0].(*ListFilesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFiles indicates an expected call of ListFiles.
func (mr *MockGerritClientMockRecorder) ListFiles(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFiles", reflect.TypeOf((*MockGerritClient)(nil).ListFiles), varargs...)
}

// CreateChange mocks base method.
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

// CreateChange indicates an expected call of CreateChange.
func (mr *MockGerritClientMockRecorder) CreateChange(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateChange", reflect.TypeOf((*MockGerritClient)(nil).CreateChange), varargs...)
}

// ChangeEditFileContent mocks base method.
func (m *MockGerritClient) ChangeEditFileContent(ctx context.Context, in *ChangeEditFileContentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ChangeEditFileContent", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeEditFileContent indicates an expected call of ChangeEditFileContent.
func (mr *MockGerritClientMockRecorder) ChangeEditFileContent(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeEditFileContent", reflect.TypeOf((*MockGerritClient)(nil).ChangeEditFileContent), varargs...)
}

// DeleteEditFileContent mocks base method.
func (m *MockGerritClient) DeleteEditFileContent(ctx context.Context, in *DeleteEditFileContentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteEditFileContent", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteEditFileContent indicates an expected call of DeleteEditFileContent.
func (mr *MockGerritClientMockRecorder) DeleteEditFileContent(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEditFileContent", reflect.TypeOf((*MockGerritClient)(nil).DeleteEditFileContent), varargs...)
}

// ChangeEditPublish mocks base method.
func (m *MockGerritClient) ChangeEditPublish(ctx context.Context, in *ChangeEditPublishRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ChangeEditPublish", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeEditPublish indicates an expected call of ChangeEditPublish.
func (mr *MockGerritClientMockRecorder) ChangeEditPublish(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeEditPublish", reflect.TypeOf((*MockGerritClient)(nil).ChangeEditPublish), varargs...)
}

// AddReviewer mocks base method.
func (m *MockGerritClient) AddReviewer(ctx context.Context, in *AddReviewerRequest, opts ...grpc.CallOption) (*AddReviewerResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddReviewer", varargs...)
	ret0, _ := ret[0].(*AddReviewerResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddReviewer indicates an expected call of AddReviewer.
func (mr *MockGerritClientMockRecorder) AddReviewer(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddReviewer", reflect.TypeOf((*MockGerritClient)(nil).AddReviewer), varargs...)
}

// SetReview mocks base method.
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

// SetReview indicates an expected call of SetReview.
func (mr *MockGerritClientMockRecorder) SetReview(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetReview", reflect.TypeOf((*MockGerritClient)(nil).SetReview), varargs...)
}

// AddToAttentionSet mocks base method.
func (m *MockGerritClient) AddToAttentionSet(ctx context.Context, in *AttentionSetRequest, opts ...grpc.CallOption) (*AccountInfo, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddToAttentionSet", varargs...)
	ret0, _ := ret[0].(*AccountInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddToAttentionSet indicates an expected call of AddToAttentionSet.
func (mr *MockGerritClientMockRecorder) AddToAttentionSet(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddToAttentionSet", reflect.TypeOf((*MockGerritClient)(nil).AddToAttentionSet), varargs...)
}

// SubmitChange mocks base method.
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

// SubmitChange indicates an expected call of SubmitChange.
func (mr *MockGerritClientMockRecorder) SubmitChange(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitChange", reflect.TypeOf((*MockGerritClient)(nil).SubmitChange), varargs...)
}

// AbandonChange mocks base method.
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

// AbandonChange indicates an expected call of AbandonChange.
func (mr *MockGerritClientMockRecorder) AbandonChange(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AbandonChange", reflect.TypeOf((*MockGerritClient)(nil).AbandonChange), varargs...)
}

// MockGerritServer is a mock of GerritServer interface.
type MockGerritServer struct {
	ctrl     *gomock.Controller
	recorder *MockGerritServerMockRecorder
}

// MockGerritServerMockRecorder is the mock recorder for MockGerritServer.
type MockGerritServerMockRecorder struct {
	mock *MockGerritServer
}

// NewMockGerritServer creates a new mock instance.
func NewMockGerritServer(ctrl *gomock.Controller) *MockGerritServer {
	mock := &MockGerritServer{ctrl: ctrl}
	mock.recorder = &MockGerritServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGerritServer) EXPECT() *MockGerritServerMockRecorder {
	return m.recorder
}

// ListProjects mocks base method.
func (m *MockGerritServer) ListProjects(arg0 context.Context, arg1 *ListProjectsRequest) (*ListProjectsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProjects", arg0, arg1)
	ret0, _ := ret[0].(*ListProjectsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProjects indicates an expected call of ListProjects.
func (mr *MockGerritServerMockRecorder) ListProjects(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProjects", reflect.TypeOf((*MockGerritServer)(nil).ListProjects), arg0, arg1)
}

// GetRefInfo mocks base method.
func (m *MockGerritServer) GetRefInfo(arg0 context.Context, arg1 *RefInfoRequest) (*RefInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRefInfo", arg0, arg1)
	ret0, _ := ret[0].(*RefInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRefInfo indicates an expected call of GetRefInfo.
func (mr *MockGerritServerMockRecorder) GetRefInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRefInfo", reflect.TypeOf((*MockGerritServer)(nil).GetRefInfo), arg0, arg1)
}

// ListFileOwners mocks base method.
func (m *MockGerritServer) ListFileOwners(arg0 context.Context, arg1 *ListFileOwnersRequest) (*ListOwnersResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFileOwners", arg0, arg1)
	ret0, _ := ret[0].(*ListOwnersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFileOwners indicates an expected call of ListFileOwners.
func (mr *MockGerritServerMockRecorder) ListFileOwners(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFileOwners", reflect.TypeOf((*MockGerritServer)(nil).ListFileOwners), arg0, arg1)
}

// ListChanges mocks base method.
func (m *MockGerritServer) ListChanges(arg0 context.Context, arg1 *ListChangesRequest) (*ListChangesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListChanges", arg0, arg1)
	ret0, _ := ret[0].(*ListChangesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListChanges indicates an expected call of ListChanges.
func (mr *MockGerritServerMockRecorder) ListChanges(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListChanges", reflect.TypeOf((*MockGerritServer)(nil).ListChanges), arg0, arg1)
}

// GetChange mocks base method.
func (m *MockGerritServer) GetChange(arg0 context.Context, arg1 *GetChangeRequest) (*ChangeInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChange", arg0, arg1)
	ret0, _ := ret[0].(*ChangeInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChange indicates an expected call of GetChange.
func (mr *MockGerritServerMockRecorder) GetChange(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChange", reflect.TypeOf((*MockGerritServer)(nil).GetChange), arg0, arg1)
}

// GetMergeable mocks base method.
func (m *MockGerritServer) GetMergeable(arg0 context.Context, arg1 *GetMergeableRequest) (*MergeableInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMergeable", arg0, arg1)
	ret0, _ := ret[0].(*MergeableInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMergeable indicates an expected call of GetMergeable.
func (mr *MockGerritServerMockRecorder) GetMergeable(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMergeable", reflect.TypeOf((*MockGerritServer)(nil).GetMergeable), arg0, arg1)
}

// ListFiles mocks base method.
func (m *MockGerritServer) ListFiles(arg0 context.Context, arg1 *ListFilesRequest) (*ListFilesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFiles", arg0, arg1)
	ret0, _ := ret[0].(*ListFilesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFiles indicates an expected call of ListFiles.
func (mr *MockGerritServerMockRecorder) ListFiles(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFiles", reflect.TypeOf((*MockGerritServer)(nil).ListFiles), arg0, arg1)
}

// CreateChange mocks base method.
func (m *MockGerritServer) CreateChange(arg0 context.Context, arg1 *CreateChangeRequest) (*ChangeInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateChange", arg0, arg1)
	ret0, _ := ret[0].(*ChangeInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateChange indicates an expected call of CreateChange.
func (mr *MockGerritServerMockRecorder) CreateChange(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateChange", reflect.TypeOf((*MockGerritServer)(nil).CreateChange), arg0, arg1)
}

// ChangeEditFileContent mocks base method.
func (m *MockGerritServer) ChangeEditFileContent(arg0 context.Context, arg1 *ChangeEditFileContentRequest) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeEditFileContent", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeEditFileContent indicates an expected call of ChangeEditFileContent.
func (mr *MockGerritServerMockRecorder) ChangeEditFileContent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeEditFileContent", reflect.TypeOf((*MockGerritServer)(nil).ChangeEditFileContent), arg0, arg1)
}

// DeleteEditFileContent mocks base method.
func (m *MockGerritServer) DeleteEditFileContent(arg0 context.Context, arg1 *DeleteEditFileContentRequest) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEditFileContent", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteEditFileContent indicates an expected call of DeleteEditFileContent.
func (mr *MockGerritServerMockRecorder) DeleteEditFileContent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEditFileContent", reflect.TypeOf((*MockGerritServer)(nil).DeleteEditFileContent), arg0, arg1)
}

// ChangeEditPublish mocks base method.
func (m *MockGerritServer) ChangeEditPublish(arg0 context.Context, arg1 *ChangeEditPublishRequest) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeEditPublish", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeEditPublish indicates an expected call of ChangeEditPublish.
func (mr *MockGerritServerMockRecorder) ChangeEditPublish(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeEditPublish", reflect.TypeOf((*MockGerritServer)(nil).ChangeEditPublish), arg0, arg1)
}

// AddReviewer mocks base method.
func (m *MockGerritServer) AddReviewer(arg0 context.Context, arg1 *AddReviewerRequest) (*AddReviewerResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddReviewer", arg0, arg1)
	ret0, _ := ret[0].(*AddReviewerResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddReviewer indicates an expected call of AddReviewer.
func (mr *MockGerritServerMockRecorder) AddReviewer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddReviewer", reflect.TypeOf((*MockGerritServer)(nil).AddReviewer), arg0, arg1)
}

// SetReview mocks base method.
func (m *MockGerritServer) SetReview(arg0 context.Context, arg1 *SetReviewRequest) (*ReviewResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetReview", arg0, arg1)
	ret0, _ := ret[0].(*ReviewResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetReview indicates an expected call of SetReview.
func (mr *MockGerritServerMockRecorder) SetReview(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetReview", reflect.TypeOf((*MockGerritServer)(nil).SetReview), arg0, arg1)
}

// AddToAttentionSet mocks base method.
func (m *MockGerritServer) AddToAttentionSet(arg0 context.Context, arg1 *AttentionSetRequest) (*AccountInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddToAttentionSet", arg0, arg1)
	ret0, _ := ret[0].(*AccountInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddToAttentionSet indicates an expected call of AddToAttentionSet.
func (mr *MockGerritServerMockRecorder) AddToAttentionSet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddToAttentionSet", reflect.TypeOf((*MockGerritServer)(nil).AddToAttentionSet), arg0, arg1)
}

// SubmitChange mocks base method.
func (m *MockGerritServer) SubmitChange(arg0 context.Context, arg1 *SubmitChangeRequest) (*ChangeInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubmitChange", arg0, arg1)
	ret0, _ := ret[0].(*ChangeInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubmitChange indicates an expected call of SubmitChange.
func (mr *MockGerritServerMockRecorder) SubmitChange(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitChange", reflect.TypeOf((*MockGerritServer)(nil).SubmitChange), arg0, arg1)
}

// AbandonChange mocks base method.
func (m *MockGerritServer) AbandonChange(arg0 context.Context, arg1 *AbandonChangeRequest) (*ChangeInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AbandonChange", arg0, arg1)
	ret0, _ := ret[0].(*ChangeInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AbandonChange indicates an expected call of AbandonChange.
func (mr *MockGerritServerMockRecorder) AbandonChange(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AbandonChange", reflect.TypeOf((*MockGerritServer)(nil).AbandonChange), arg0, arg1)
}
