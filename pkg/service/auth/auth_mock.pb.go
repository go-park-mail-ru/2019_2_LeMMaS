// Code generated by MockGen. DO NOT EDIT.
// Source: auth.pb.go

// Package auth is a generated GoMock package.
package auth

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockAuthClient is a mock of AuthClient interface
type MockAuthClient struct {
	ctrl     *gomock.Controller
	recorder *MockAuthClientMockRecorder
}

// MockAuthClientMockRecorder is the mock recorder for MockAuthClient
type MockAuthClientMockRecorder struct {
	mock *MockAuthClient
}

// NewMockAuthClient creates a new mock instance
func NewMockAuthClient(ctrl *gomock.Controller) *MockAuthClient {
	mock := &MockAuthClient{ctrl: ctrl}
	mock.recorder = &MockAuthClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthClient) EXPECT() *MockAuthClientMockRecorder {
	return m.recorder
}

// Login mocks base method
func (m *MockAuthClient) Login(ctx context.Context, in *LoginParams, opts ...grpc.CallOption) (*LoginResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Login", varargs...)
	ret0, _ := ret[0].(*LoginResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login
func (mr *MockAuthClientMockRecorder) Login(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockAuthClient)(nil).Login), varargs...)
}

// Logout mocks base method
func (m *MockAuthClient) Logout(ctx context.Context, in *LogoutParams, opts ...grpc.CallOption) (*LogoutResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Logout", varargs...)
	ret0, _ := ret[0].(*LogoutResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Logout indicates an expected call of Logout
func (mr *MockAuthClientMockRecorder) Logout(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logout", reflect.TypeOf((*MockAuthClient)(nil).Logout), varargs...)
}

// Register mocks base method
func (m *MockAuthClient) Register(ctx context.Context, in *RegisterParams, opts ...grpc.CallOption) (*RegisterResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Register", varargs...)
	ret0, _ := ret[0].(*RegisterResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register
func (mr *MockAuthClientMockRecorder) Register(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockAuthClient)(nil).Register), varargs...)
}

// GetUser mocks base method
func (m *MockAuthClient) GetUser(ctx context.Context, in *GetUserParams, opts ...grpc.CallOption) (*GetUserResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUser", varargs...)
	ret0, _ := ret[0].(*GetUserResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser
func (mr *MockAuthClientMockRecorder) GetUser(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockAuthClient)(nil).GetUser), varargs...)
}

// GetPasswordHash mocks base method
func (m *MockAuthClient) GetPasswordHash(ctx context.Context, in *GetPasswordHashParams, opts ...grpc.CallOption) (*GetPasswordHashResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPasswordHash", varargs...)
	ret0, _ := ret[0].(*GetPasswordHashResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPasswordHash indicates an expected call of GetPasswordHash
func (mr *MockAuthClientMockRecorder) GetPasswordHash(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPasswordHash", reflect.TypeOf((*MockAuthClient)(nil).GetPasswordHash), varargs...)
}

// MockAuthServer is a mock of AuthServer interface
type MockAuthServer struct {
	ctrl     *gomock.Controller
	recorder *MockAuthServerMockRecorder
}

// MockAuthServerMockRecorder is the mock recorder for MockAuthServer
type MockAuthServerMockRecorder struct {
	mock *MockAuthServer
}

// NewMockAuthServer creates a new mock instance
func NewMockAuthServer(ctrl *gomock.Controller) *MockAuthServer {
	mock := &MockAuthServer{ctrl: ctrl}
	mock.recorder = &MockAuthServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthServer) EXPECT() *MockAuthServerMockRecorder {
	return m.recorder
}

// Login mocks base method
func (m *MockAuthServer) Login(arg0 context.Context, arg1 *LoginParams) (*LoginResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", arg0, arg1)
	ret0, _ := ret[0].(*LoginResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login
func (mr *MockAuthServerMockRecorder) Login(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockAuthServer)(nil).Login), arg0, arg1)
}

// Logout mocks base method
func (m *MockAuthServer) Logout(arg0 context.Context, arg1 *LogoutParams) (*LogoutResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logout", arg0, arg1)
	ret0, _ := ret[0].(*LogoutResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Logout indicates an expected call of Logout
func (mr *MockAuthServerMockRecorder) Logout(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logout", reflect.TypeOf((*MockAuthServer)(nil).Logout), arg0, arg1)
}

// Register mocks base method
func (m *MockAuthServer) Register(arg0 context.Context, arg1 *RegisterParams) (*RegisterResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", arg0, arg1)
	ret0, _ := ret[0].(*RegisterResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register
func (mr *MockAuthServerMockRecorder) Register(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockAuthServer)(nil).Register), arg0, arg1)
}

// GetUser mocks base method
func (m *MockAuthServer) GetUser(arg0 context.Context, arg1 *GetUserParams) (*GetUserResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(*GetUserResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser
func (mr *MockAuthServerMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockAuthServer)(nil).GetUser), arg0, arg1)
}

// GetPasswordHash mocks base method
func (m *MockAuthServer) GetPasswordHash(arg0 context.Context, arg1 *GetPasswordHashParams) (*GetPasswordHashResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPasswordHash", arg0, arg1)
	ret0, _ := ret[0].(*GetPasswordHashResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPasswordHash indicates an expected call of GetPasswordHash
func (mr *MockAuthServerMockRecorder) GetPasswordHash(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPasswordHash", reflect.TypeOf((*MockAuthServer)(nil).GetPasswordHash), arg0, arg1)
}
