// Code generated by MockGen. DO NOT EDIT.
// Source: app/usecase/post_usecase.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/toshiykst/golang-rest-api/app/domain/model"
)

// MockPostUsecase is a mock of PostUsecase interface.
type MockPostUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockPostUsecaseMockRecorder
}

// MockPostUsecaseMockRecorder is the mock recorder for MockPostUsecase.
type MockPostUsecaseMockRecorder struct {
	mock *MockPostUsecase
}

// NewMockPostUsecase creates a new mock instance.
func NewMockPostUsecase(ctrl *gomock.Controller) *MockPostUsecase {
	mock := &MockPostUsecase{ctrl: ctrl}
	mock.recorder = &MockPostUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPostUsecase) EXPECT() *MockPostUsecaseMockRecorder {
	return m.recorder
}

// CreatePost mocks base method.
func (m *MockPostUsecase) CreatePost(arg0 *model.Post) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePost", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatePost indicates an expected call of CreatePost.
func (mr *MockPostUsecaseMockRecorder) CreatePost(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePost", reflect.TypeOf((*MockPostUsecase)(nil).CreatePost), arg0)
}

// DeletePost mocks base method.
func (m *MockPostUsecase) DeletePost(arg0 *model.Post) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePost", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePost indicates an expected call of DeletePost.
func (mr *MockPostUsecaseMockRecorder) DeletePost(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePost", reflect.TypeOf((*MockPostUsecase)(nil).DeletePost), arg0)
}

// GetPost mocks base method.
func (m *MockPostUsecase) GetPost(id int) (model.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPost", id)
	ret0, _ := ret[0].(model.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPost indicates an expected call of GetPost.
func (mr *MockPostUsecaseMockRecorder) GetPost(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPost", reflect.TypeOf((*MockPostUsecase)(nil).GetPost), id)
}

// GetPosts mocks base method.
func (m *MockPostUsecase) GetPosts() (model.Posts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPosts")
	ret0, _ := ret[0].(model.Posts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPosts indicates an expected call of GetPosts.
func (mr *MockPostUsecaseMockRecorder) GetPosts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPosts", reflect.TypeOf((*MockPostUsecase)(nil).GetPosts))
}

// UpdatePost mocks base method.
func (m *MockPostUsecase) UpdatePost(arg0 *model.Post) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePost", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePost indicates an expected call of UpdatePost.
func (mr *MockPostUsecaseMockRecorder) UpdatePost(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePost", reflect.TypeOf((*MockPostUsecase)(nil).UpdatePost), arg0)
}