// Code generated by MockGen. DO NOT EDIT.
// Source: sam/service/attachment.go

// Package service is a generated GoMock package.
package service

import (
	context "context"
	types "github.com/crusttech/crust/sam/types"
	gomock "github.com/golang/mock/gomock"
	io "io"
	reflect "reflect"
)

// MockAttachmentService is a mock of AttachmentService interface
type MockAttachmentService struct {
	ctrl     *gomock.Controller
	recorder *MockAttachmentServiceMockRecorder
}

// MockAttachmentServiceMockRecorder is the mock recorder for MockAttachmentService
type MockAttachmentServiceMockRecorder struct {
	mock *MockAttachmentService
}

// NewMockAttachmentService creates a new mock instance
func NewMockAttachmentService(ctrl *gomock.Controller) *MockAttachmentService {
	mock := &MockAttachmentService{ctrl: ctrl}
	mock.recorder = &MockAttachmentServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAttachmentService) EXPECT() *MockAttachmentServiceMockRecorder {
	return m.recorder
}

// With mocks base method
func (m *MockAttachmentService) With(ctx context.Context) AttachmentService {
	ret := m.ctrl.Call(m, "With", ctx)
	ret0, _ := ret[0].(AttachmentService)
	return ret0
}

// With indicates an expected call of With
func (mr *MockAttachmentServiceMockRecorder) With(ctx interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "With", reflect.TypeOf((*MockAttachmentService)(nil).With), ctx)
}

// FindByID mocks base method
func (m *MockAttachmentService) FindByID(id uint64) (*types.Attachment, error) {
	ret := m.ctrl.Call(m, "FindByID", id)
	ret0, _ := ret[0].(*types.Attachment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID
func (mr *MockAttachmentServiceMockRecorder) FindByID(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockAttachmentService)(nil).FindByID), id)
}

// Create mocks base method
func (m *MockAttachmentService) Create(channelId uint64, name string, size int64, fh io.ReadSeeker) (*types.Attachment, error) {
	ret := m.ctrl.Call(m, "Create", channelId, name, size, fh)
	ret0, _ := ret[0].(*types.Attachment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockAttachmentServiceMockRecorder) Create(channelId, name, size, fh interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAttachmentService)(nil).Create), channelId, name, size, fh)
}

// LoadFromMessages mocks base method
func (m *MockAttachmentService) LoadFromMessages(mm types.MessageSet) error {
	ret := m.ctrl.Call(m, "LoadFromMessages", mm)
	ret0, _ := ret[0].(error)
	return ret0
}

// LoadFromMessages indicates an expected call of LoadFromMessages
func (mr *MockAttachmentServiceMockRecorder) LoadFromMessages(mm interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadFromMessages", reflect.TypeOf((*MockAttachmentService)(nil).LoadFromMessages), mm)
}

// OpenOriginal mocks base method
func (m *MockAttachmentService) OpenOriginal(att *types.Attachment) (io.ReadSeeker, error) {
	ret := m.ctrl.Call(m, "OpenOriginal", att)
	ret0, _ := ret[0].(io.ReadSeeker)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenOriginal indicates an expected call of OpenOriginal
func (mr *MockAttachmentServiceMockRecorder) OpenOriginal(att interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenOriginal", reflect.TypeOf((*MockAttachmentService)(nil).OpenOriginal), att)
}

// OpenPreview mocks base method
func (m *MockAttachmentService) OpenPreview(att *types.Attachment) (io.ReadSeeker, error) {
	ret := m.ctrl.Call(m, "OpenPreview", att)
	ret0, _ := ret[0].(io.ReadSeeker)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenPreview indicates an expected call of OpenPreview
func (mr *MockAttachmentServiceMockRecorder) OpenPreview(att interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenPreview", reflect.TypeOf((*MockAttachmentService)(nil).OpenPreview), att)
}