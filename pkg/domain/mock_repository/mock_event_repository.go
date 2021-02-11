// Code generated by MockGen. DO NOT EDIT.
// Source: domain/repository/event_repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	model "github.com/THEToilet/events-server/pkg/domain/model"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockEventRepository is a mock of EventRepository interface
type MockEventRepository struct {
	ctrl     *gomock.Controller
	recorder *MockEventRepositoryMockRecorder
}

// MockEventRepositoryMockRecorder is the mock recorder for MockEventRepository
type MockEventRepositoryMockRecorder struct {
	mock *MockEventRepository
}

// NewMockEventRepository creates a new mock instance
func NewMockEventRepository(ctrl *gomock.Controller) *MockEventRepository {
	mock := &MockEventRepository{ctrl: ctrl}
	mock.recorder = &MockEventRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEventRepository) EXPECT() *MockEventRepositoryMockRecorder {
	return m.recorder
}

// Find mocks base method
func (m *MockEventRepository) Find(id string) (*model.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*model.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockEventRepositoryMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockEventRepository)(nil).Find), id)
}

// FindAll mocks base method
func (m *MockEventRepository) FindAll() ([]*model.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]*model.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll
func (mr *MockEventRepositoryMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockEventRepository)(nil).FindAll))
}

// Save mocks base method
func (m *MockEventRepository) Save(id string) (*model.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", id)
	ret0, _ := ret[0].(*model.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save
func (mr *MockEventRepositoryMockRecorder) Save(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockEventRepository)(nil).Save), id)
}

// Delete mocks base method
func (m *MockEventRepository) Delete(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockEventRepositoryMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockEventRepository)(nil).Delete), id)
}

// Update mocks base method
func (m *MockEventRepository) Update(id string) (*model.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", id)
	ret0, _ := ret[0].(*model.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockEventRepositoryMockRecorder) Update(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockEventRepository)(nil).Update), id)
}