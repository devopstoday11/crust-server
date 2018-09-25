// Code generated by MockGen. DO NOT EDIT.
// Source: sam/service/organisation.go

// Package service is a generated GoMock package.
package service

import (
	context "context"
	types "github.com/crusttech/crust/sam/types"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockOrganisationService is a mock of OrganisationService interface
type MockOrganisationService struct {
	ctrl     *gomock.Controller
	recorder *MockOrganisationServiceMockRecorder
}

// MockOrganisationServiceMockRecorder is the mock recorder for MockOrganisationService
type MockOrganisationServiceMockRecorder struct {
	mock *MockOrganisationService
}

// NewMockOrganisationService creates a new mock instance
func NewMockOrganisationService(ctrl *gomock.Controller) *MockOrganisationService {
	mock := &MockOrganisationService{ctrl: ctrl}
	mock.recorder = &MockOrganisationServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOrganisationService) EXPECT() *MockOrganisationServiceMockRecorder {
	return m.recorder
}

// With mocks base method
func (m *MockOrganisationService) With(ctx context.Context) OrganisationService {
	ret := m.ctrl.Call(m, "With", ctx)
	ret0, _ := ret[0].(OrganisationService)
	return ret0
}

// With indicates an expected call of With
func (mr *MockOrganisationServiceMockRecorder) With(ctx interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "With", reflect.TypeOf((*MockOrganisationService)(nil).With), ctx)
}

// FindByID mocks base method
func (m *MockOrganisationService) FindByID(organisationID uint64) (*types.Organisation, error) {
	ret := m.ctrl.Call(m, "FindByID", organisationID)
	ret0, _ := ret[0].(*types.Organisation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID
func (mr *MockOrganisationServiceMockRecorder) FindByID(organisationID interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockOrganisationService)(nil).FindByID), organisationID)
}

// Find mocks base method
func (m *MockOrganisationService) Find(filter *types.OrganisationFilter) ([]*types.Organisation, error) {
	ret := m.ctrl.Call(m, "Find", filter)
	ret0, _ := ret[0].([]*types.Organisation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockOrganisationServiceMockRecorder) Find(filter interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockOrganisationService)(nil).Find), filter)
}

// Create mocks base method
func (m *MockOrganisationService) Create(organisation *types.Organisation) (*types.Organisation, error) {
	ret := m.ctrl.Call(m, "Create", organisation)
	ret0, _ := ret[0].(*types.Organisation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockOrganisationServiceMockRecorder) Create(organisation interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockOrganisationService)(nil).Create), organisation)
}

// Update mocks base method
func (m *MockOrganisationService) Update(organisation *types.Organisation) (*types.Organisation, error) {
	ret := m.ctrl.Call(m, "Update", organisation)
	ret0, _ := ret[0].(*types.Organisation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockOrganisationServiceMockRecorder) Update(organisation interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockOrganisationService)(nil).Update), organisation)
}

// Archive mocks base method
func (m *MockOrganisationService) Archive(ID uint64) error {
	ret := m.ctrl.Call(m, "Archive", ID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Archive indicates an expected call of Archive
func (mr *MockOrganisationServiceMockRecorder) Archive(ID interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Archive", reflect.TypeOf((*MockOrganisationService)(nil).Archive), ID)
}

// Unarchive mocks base method
func (m *MockOrganisationService) Unarchive(ID uint64) error {
	ret := m.ctrl.Call(m, "Unarchive", ID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unarchive indicates an expected call of Unarchive
func (mr *MockOrganisationServiceMockRecorder) Unarchive(ID interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unarchive", reflect.TypeOf((*MockOrganisationService)(nil).Unarchive), ID)
}

// Delete mocks base method
func (m *MockOrganisationService) Delete(ID uint64) error {
	ret := m.ctrl.Call(m, "Delete", ID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockOrganisationServiceMockRecorder) Delete(ID interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockOrganisationService)(nil).Delete), ID)
}
