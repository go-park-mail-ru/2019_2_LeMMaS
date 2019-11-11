// Code generated by MockGen. DO NOT EDIT.
// Source: usecase.go

// Package user is a generated GoMock package.
package user

import (
	model "github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"
	gomock "github.com/golang/mock/gomock"
	io "io"
	reflect "reflect"
)

// MockUsecase is a mock of Usecase interface
type MockUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockUsecaseMockRecorder
}

// MockUsecaseMockRecorder is the mock recorder for MockUsecase
type MockUsecaseMockRecorder struct {
	mock *MockUsecase
}

// NewMockUsecase creates a new mock instance
func NewMockUsecase(ctrl *gomock.Controller) *MockUsecase {
	mock := &MockUsecase{ctrl: ctrl}
	mock.recorder = &MockUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUsecase) EXPECT() *MockUsecaseMockRecorder {
	return m.recorder
}

// GetAllUsers mocks base method
func (m *MockUsecase) GetAllUsers() ([]model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUsers")
	ret0, _ := ret[0].([]model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUsers indicates an expected call of GetAllUsers
func (mr *MockUsecaseMockRecorder) GetAllUsers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsers", reflect.TypeOf((*MockUsecase)(nil).GetAllUsers))
}

// GetUserBySessionID mocks base method
func (m *MockUsecase) GetUserBySessionID(sessionID string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserBySessionID", sessionID)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserBySessionID indicates an expected call of GetUserBySessionID
func (mr *MockUsecaseMockRecorder) GetUserBySessionID(sessionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserBySessionID", reflect.TypeOf((*MockUsecase)(nil).GetUserBySessionID), sessionID)
}

// UpdateUser mocks base method
func (m *MockUsecase) UpdateUser(id int, password, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", id, password, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser
func (mr *MockUsecaseMockRecorder) UpdateUser(id, password, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUsecase)(nil).UpdateUser), id, password, name)
}

// UpdateUserAvatar mocks base method
func (m *MockUsecase) UpdateUserAvatar(user *model.User, avatarFile io.Reader, avatarPath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserAvatar", user, avatarFile, avatarPath)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserAvatar indicates an expected call of UpdateUserAvatar
func (mr *MockUsecaseMockRecorder) UpdateUserAvatar(user, avatarFile, avatarPath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserAvatar", reflect.TypeOf((*MockUsecase)(nil).UpdateUserAvatar), user, avatarFile, avatarPath)
}

// GetAvatarUrlByName mocks base method
func (m *MockUsecase) GetAvatarUrlByName(name string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAvatarUrlByName", name)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetAvatarUrlByName indicates an expected call of GetAvatarUrlByName
func (mr *MockUsecaseMockRecorder) GetAvatarUrlByName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAvatarUrlByName", reflect.TypeOf((*MockUsecase)(nil).GetAvatarUrlByName), name)
}

// Register mocks base method
func (m *MockUsecase) Register(email, password, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", email, password, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register
func (mr *MockUsecaseMockRecorder) Register(email, password, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockUsecase)(nil).Register), email, password, name)
}

// Login mocks base method
func (m *MockUsecase) Login(email, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", email, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login
func (mr *MockUsecaseMockRecorder) Login(email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockUsecase)(nil).Login), email, password)
}

// Logout mocks base method
func (m *MockUsecase) Logout(sessionID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logout", sessionID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Logout indicates an expected call of Logout
func (mr *MockUsecaseMockRecorder) Logout(sessionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logout", reflect.TypeOf((*MockUsecase)(nil).Logout), sessionID)
}
