// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package api is a generated GoMock package.
package api

import (
	gomock "github.com/golang/mock/gomock"
	io "io"
	reflect "reflect"
)

// MockFileRepository is a mock of FileRepository interface
type MockFileRepository struct {
	ctrl     *gomock.Controller
	recorder *MockFileRepositoryMockRecorder
}

// MockFileRepositoryMockRecorder is the mock recorder for MockFileRepository
type MockFileRepositoryMockRecorder struct {
	mock *MockFileRepository
}

// NewMockFileRepository creates a new mock instance
func NewMockFileRepository(ctrl *gomock.Controller) *MockFileRepository {
	mock := &MockFileRepository{ctrl: ctrl}
	mock.recorder = &MockFileRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFileRepository) EXPECT() *MockFileRepositoryMockRecorder {
	return m.recorder
}

// Store mocks base method
func (m *MockFileRepository) Store(file io.Reader) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", file)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Store indicates an expected call of Store
func (mr *MockFileRepositoryMockRecorder) Store(file interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockFileRepository)(nil).Store), file)
}