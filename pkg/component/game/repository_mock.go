// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package game is a generated GoMock package.
package game

import (
	model "github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockRepository is a mock of Repository interface
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// CreateRoom mocks base method
func (m *MockRepository) CreateRoom() *model.Room {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRoom")
	ret0, _ := ret[0].(*model.Room)
	return ret0
}

// CreateRoom indicates an expected call of CreateRoom
func (mr *MockRepositoryMockRecorder) CreateRoom() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRoom", reflect.TypeOf((*MockRepository)(nil).CreateRoom))
}

// GetAllRooms mocks base method
func (m *MockRepository) GetAllRooms() []*model.Room {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllRooms")
	ret0, _ := ret[0].([]*model.Room)
	return ret0
}

// GetAllRooms indicates an expected call of GetAllRooms
func (mr *MockRepositoryMockRecorder) GetAllRooms() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllRooms", reflect.TypeOf((*MockRepository)(nil).GetAllRooms))
}

// GetRoomByID mocks base method
func (m *MockRepository) GetRoomByID(id int) *model.Room {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoomByID", id)
	ret0, _ := ret[0].(*model.Room)
	return ret0
}

// GetRoomByID indicates an expected call of GetRoomByID
func (mr *MockRepositoryMockRecorder) GetRoomByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoomByID", reflect.TypeOf((*MockRepository)(nil).GetRoomByID), id)
}

// DeleteRoom mocks base method
func (m *MockRepository) DeleteRoom(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRoom", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRoom indicates an expected call of DeleteRoom
func (mr *MockRepositoryMockRecorder) DeleteRoom(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRoom", reflect.TypeOf((*MockRepository)(nil).DeleteRoom), id)
}

// AddPlayer mocks base method
func (m *MockRepository) AddPlayer(roomID int, player model.Player) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPlayer", roomID, player)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddPlayer indicates an expected call of AddPlayer
func (mr *MockRepositoryMockRecorder) AddPlayer(roomID, player interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPlayer", reflect.TypeOf((*MockRepository)(nil).AddPlayer), roomID, player)
}

// DeletePlayer mocks base method
func (m *MockRepository) DeletePlayer(roomID, userID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePlayer", roomID, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePlayer indicates an expected call of DeletePlayer
func (mr *MockRepositoryMockRecorder) DeletePlayer(roomID, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePlayer", reflect.TypeOf((*MockRepository)(nil).DeletePlayer), roomID, userID)
}

// SetPlayerDirection mocks base method
func (m *MockRepository) SetPlayerDirection(roomID, userID, direction int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPlayerDirection", roomID, userID, direction)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetPlayerDirection indicates an expected call of SetPlayerDirection
func (mr *MockRepositoryMockRecorder) SetPlayerDirection(roomID, userID, direction interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPlayerDirection", reflect.TypeOf((*MockRepository)(nil).SetPlayerDirection), roomID, userID, direction)
}

// SetPlayerSpeed mocks base method
func (m *MockRepository) SetPlayerSpeed(roomID, userID, speed int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPlayerSpeed", roomID, userID, speed)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetPlayerSpeed indicates an expected call of SetPlayerSpeed
func (mr *MockRepositoryMockRecorder) SetPlayerSpeed(roomID, userID, speed interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPlayerSpeed", reflect.TypeOf((*MockRepository)(nil).SetPlayerSpeed), roomID, userID, speed)
}

// SetPlayerPosition mocks base method
func (m *MockRepository) SetPlayerPosition(roomID, userID int, position model.Position) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPlayerPosition", roomID, userID, position)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetPlayerPosition indicates an expected call of SetPlayerPosition
func (mr *MockRepositoryMockRecorder) SetPlayerPosition(roomID, userID, position interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPlayerPosition", reflect.TypeOf((*MockRepository)(nil).SetPlayerPosition), roomID, userID, position)
}

// SetPlayerSize mocks base method
func (m *MockRepository) SetPlayerSize(roomID, userID, size int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPlayerSize", roomID, userID, size)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetPlayerSize indicates an expected call of SetPlayerSize
func (mr *MockRepositoryMockRecorder) SetPlayerSize(roomID, userID, size interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPlayerSize", reflect.TypeOf((*MockRepository)(nil).SetPlayerSize), roomID, userID, size)
}

// AddFood mocks base method
func (m *MockRepository) AddFood(roomID int, food []model.Food) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFood", roomID, food)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddFood indicates an expected call of AddFood
func (mr *MockRepositoryMockRecorder) AddFood(roomID, food interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFood", reflect.TypeOf((*MockRepository)(nil).AddFood), roomID, food)
}

// DeleteFood mocks base method
func (m *MockRepository) DeleteFood(roomID int, foodIDs []int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFood", roomID, foodIDs)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFood indicates an expected call of DeleteFood
func (mr *MockRepositoryMockRecorder) DeleteFood(roomID, foodIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFood", reflect.TypeOf((*MockRepository)(nil).DeleteFood), roomID, foodIDs)
}

// GetFoodInRange mocks base method
func (m *MockRepository) GetFoodInRange(roomID int, topLeftPoint, bottomRightPoint model.Position) ([]int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFoodInRange", roomID, topLeftPoint, bottomRightPoint)
	ret0, _ := ret[0].([]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFoodInRange indicates an expected call of GetFoodInRange
func (mr *MockRepositoryMockRecorder) GetFoodInRange(roomID, topLeftPoint, bottomRightPoint interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFoodInRange", reflect.TypeOf((*MockRepository)(nil).GetFoodInRange), roomID, topLeftPoint, bottomRightPoint)
}
