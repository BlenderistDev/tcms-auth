// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/telegram/telegram_grpc.pb.go

// Package mock_telegram is a generated GoMock package.
package mock_telegram

import (
	context "context"
	reflect "reflect"
	telegram "tcms-auth/pkg/telegram"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// MockTelegramClient is a mock of TelegramClient interface.
type MockTelegramClient struct {
	ctrl     *gomock.Controller
	recorder *MockTelegramClientMockRecorder
}

// MockTelegramClientMockRecorder is the mock recorder for MockTelegramClient.
type MockTelegramClientMockRecorder struct {
	mock *MockTelegramClient
}

// NewMockTelegramClient creates a new mock instance.
func NewMockTelegramClient(ctrl *gomock.Controller) *MockTelegramClient {
	mock := &MockTelegramClient{ctrl: ctrl}
	mock.recorder = &MockTelegramClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTelegramClient) EXPECT() *MockTelegramClientMockRecorder {
	return m.recorder
}

// GetDialogs mocks base method.
func (m *MockTelegramClient) GetDialogs(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*telegram.DialogsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDialogs", varargs...)
	ret0, _ := ret[0].(*telegram.DialogsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDialogs indicates an expected call of GetDialogs.
func (mr *MockTelegramClientMockRecorder) GetDialogs(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDialogs", reflect.TypeOf((*MockTelegramClient)(nil).GetDialogs), varargs...)
}

// GetUser mocks base method.
func (m *MockTelegramClient) GetUser(ctx context.Context, in *telegram.GetUserRequest, opts ...grpc.CallOption) (*telegram.UserResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUser", varargs...)
	ret0, _ := ret[0].(*telegram.UserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockTelegramClientMockRecorder) GetUser(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockTelegramClient)(nil).GetUser), varargs...)
}

// Login mocks base method.
func (m *MockTelegramClient) Login(ctx context.Context, in *telegram.LoginMessage, opts ...grpc.CallOption) (*telegram.Result, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Login", varargs...)
	ret0, _ := ret[0].(*telegram.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockTelegramClientMockRecorder) Login(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockTelegramClient)(nil).Login), varargs...)
}

// MuteChat mocks base method.
func (m *MockTelegramClient) MuteChat(ctx context.Context, in *telegram.MuteChatRequest, opts ...grpc.CallOption) (*telegram.Result, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MuteChat", varargs...)
	ret0, _ := ret[0].(*telegram.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MuteChat indicates an expected call of MuteChat.
func (mr *MockTelegramClientMockRecorder) MuteChat(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MuteChat", reflect.TypeOf((*MockTelegramClient)(nil).MuteChat), varargs...)
}

// MuteUser mocks base method.
func (m *MockTelegramClient) MuteUser(ctx context.Context, in *telegram.MuteUserRequest, opts ...grpc.CallOption) (*telegram.Result, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MuteUser", varargs...)
	ret0, _ := ret[0].(*telegram.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MuteUser indicates an expected call of MuteUser.
func (mr *MockTelegramClientMockRecorder) MuteUser(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MuteUser", reflect.TypeOf((*MockTelegramClient)(nil).MuteUser), varargs...)
}

// Send mocks base method.
func (m *MockTelegramClient) Send(ctx context.Context, in *telegram.SendMessageRequest, opts ...grpc.CallOption) (*telegram.Result, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Send", varargs...)
	ret0, _ := ret[0].(*telegram.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Send indicates an expected call of Send.
func (mr *MockTelegramClientMockRecorder) Send(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockTelegramClient)(nil).Send), varargs...)
}

// Sign mocks base method.
func (m *MockTelegramClient) Sign(ctx context.Context, in *telegram.SignMessage, opts ...grpc.CallOption) (*telegram.SignResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Sign", varargs...)
	ret0, _ := ret[0].(*telegram.SignResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sign indicates an expected call of Sign.
func (mr *MockTelegramClientMockRecorder) Sign(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sign", reflect.TypeOf((*MockTelegramClient)(nil).Sign), varargs...)
}

// MockTelegramServer is a mock of TelegramServer interface.
type MockTelegramServer struct {
	ctrl     *gomock.Controller
	recorder *MockTelegramServerMockRecorder
}

// MockTelegramServerMockRecorder is the mock recorder for MockTelegramServer.
type MockTelegramServerMockRecorder struct {
	mock *MockTelegramServer
}

// NewMockTelegramServer creates a new mock instance.
func NewMockTelegramServer(ctrl *gomock.Controller) *MockTelegramServer {
	mock := &MockTelegramServer{ctrl: ctrl}
	mock.recorder = &MockTelegramServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTelegramServer) EXPECT() *MockTelegramServerMockRecorder {
	return m.recorder
}

// GetDialogs mocks base method.
func (m *MockTelegramServer) GetDialogs(arg0 context.Context, arg1 *emptypb.Empty) (*telegram.DialogsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDialogs", arg0, arg1)
	ret0, _ := ret[0].(*telegram.DialogsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDialogs indicates an expected call of GetDialogs.
func (mr *MockTelegramServerMockRecorder) GetDialogs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDialogs", reflect.TypeOf((*MockTelegramServer)(nil).GetDialogs), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockTelegramServer) GetUser(arg0 context.Context, arg1 *telegram.GetUserRequest) (*telegram.UserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(*telegram.UserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockTelegramServerMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockTelegramServer)(nil).GetUser), arg0, arg1)
}

// Login mocks base method.
func (m *MockTelegramServer) Login(arg0 context.Context, arg1 *telegram.LoginMessage) (*telegram.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", arg0, arg1)
	ret0, _ := ret[0].(*telegram.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockTelegramServerMockRecorder) Login(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockTelegramServer)(nil).Login), arg0, arg1)
}

// MuteChat mocks base method.
func (m *MockTelegramServer) MuteChat(arg0 context.Context, arg1 *telegram.MuteChatRequest) (*telegram.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MuteChat", arg0, arg1)
	ret0, _ := ret[0].(*telegram.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MuteChat indicates an expected call of MuteChat.
func (mr *MockTelegramServerMockRecorder) MuteChat(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MuteChat", reflect.TypeOf((*MockTelegramServer)(nil).MuteChat), arg0, arg1)
}

// MuteUser mocks base method.
func (m *MockTelegramServer) MuteUser(arg0 context.Context, arg1 *telegram.MuteUserRequest) (*telegram.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MuteUser", arg0, arg1)
	ret0, _ := ret[0].(*telegram.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MuteUser indicates an expected call of MuteUser.
func (mr *MockTelegramServerMockRecorder) MuteUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MuteUser", reflect.TypeOf((*MockTelegramServer)(nil).MuteUser), arg0, arg1)
}

// Send mocks base method.
func (m *MockTelegramServer) Send(arg0 context.Context, arg1 *telegram.SendMessageRequest) (*telegram.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0, arg1)
	ret0, _ := ret[0].(*telegram.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Send indicates an expected call of Send.
func (mr *MockTelegramServerMockRecorder) Send(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockTelegramServer)(nil).Send), arg0, arg1)
}

// Sign mocks base method.
func (m *MockTelegramServer) Sign(arg0 context.Context, arg1 *telegram.SignMessage) (*telegram.SignResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sign", arg0, arg1)
	ret0, _ := ret[0].(*telegram.SignResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sign indicates an expected call of Sign.
func (mr *MockTelegramServerMockRecorder) Sign(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sign", reflect.TypeOf((*MockTelegramServer)(nil).Sign), arg0, arg1)
}

// mustEmbedUnimplementedTelegramServer mocks base method.
func (m *MockTelegramServer) mustEmbedUnimplementedTelegramServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedTelegramServer")
}

// mustEmbedUnimplementedTelegramServer indicates an expected call of mustEmbedUnimplementedTelegramServer.
func (mr *MockTelegramServerMockRecorder) mustEmbedUnimplementedTelegramServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedTelegramServer", reflect.TypeOf((*MockTelegramServer)(nil).mustEmbedUnimplementedTelegramServer))
}

// MockUnsafeTelegramServer is a mock of UnsafeTelegramServer interface.
type MockUnsafeTelegramServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeTelegramServerMockRecorder
}

// MockUnsafeTelegramServerMockRecorder is the mock recorder for MockUnsafeTelegramServer.
type MockUnsafeTelegramServerMockRecorder struct {
	mock *MockUnsafeTelegramServer
}

// NewMockUnsafeTelegramServer creates a new mock instance.
func NewMockUnsafeTelegramServer(ctrl *gomock.Controller) *MockUnsafeTelegramServer {
	mock := &MockUnsafeTelegramServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeTelegramServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeTelegramServer) EXPECT() *MockUnsafeTelegramServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedTelegramServer mocks base method.
func (m *MockUnsafeTelegramServer) mustEmbedUnimplementedTelegramServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedTelegramServer")
}

// mustEmbedUnimplementedTelegramServer indicates an expected call of mustEmbedUnimplementedTelegramServer.
func (mr *MockUnsafeTelegramServerMockRecorder) mustEmbedUnimplementedTelegramServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedTelegramServer", reflect.TypeOf((*MockUnsafeTelegramServer)(nil).mustEmbedUnimplementedTelegramServer))
}
