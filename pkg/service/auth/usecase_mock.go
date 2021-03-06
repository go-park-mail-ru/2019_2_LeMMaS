// Code generated by MockGen. DO NOT EDIT.
// Source: usecase.go

// Package auth is a generated GoMock package.
package auth

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAuthUsecase is a mock of AuthUsecase interface
type MockAuthUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockAuthUsecaseMockRecorder
}

// MockAuthUsecaseMockRecorder is the mock recorder for MockAuthUsecase
type MockAuthUsecaseMockRecorder struct {
	mock *MockAuthUsecase
}

// NewMockAuthUsecase creates a new mock instance
func NewMockAuthUsecase(ctrl *gomock.Controller) *MockAuthUsecase {
	mock := &MockAuthUsecase{ctrl: ctrl}
	mock.recorder = &MockAuthUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthUsecase) EXPECT() *MockAuthUsecaseMockRecorder {
	return m.recorder
}

// Login mocks base method
func (m *MockAuthUsecase) Login(email, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", email, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login
func (mr *MockAuthUsecaseMockRecorder) Login(email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockAuthUsecase)(nil).Login), email, password)
}

// Logout mocks base method
func (m *MockAuthUsecase) Logout(session string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logout", session)
	ret0, _ := ret[0].(error)
	return ret0
}

// Logout indicates an expected call of Logout
func (mr *MockAuthUsecaseMockRecorder) Logout(session interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logout", reflect.TypeOf((*MockAuthUsecase)(nil).Logout), session)
}

// Register mocks base method
func (m *MockAuthUsecase) Register(email, password, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", email, password, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register
func (mr *MockAuthUsecaseMockRecorder) Register(email, password, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockAuthUsecase)(nil).Register), email, password, name)
}

// GetUser mocks base method
func (m *MockAuthUsecase) GetUser(session string) (int, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", session)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser
func (mr *MockAuthUsecaseMockRecorder) GetUser(session interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockAuthUsecase)(nil).GetUser), session)
}

// GetPasswordHash mocks base method
func (m *MockAuthUsecase) GetPasswordHash(password string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPasswordHash", password)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetPasswordHash indicates an expected call of GetPasswordHash
func (mr *MockAuthUsecaseMockRecorder) GetPasswordHash(password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPasswordHash", reflect.TypeOf((*MockAuthUsecase)(nil).GetPasswordHash), password)
}
