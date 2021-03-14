// Code generated by MockGen. DO NOT EDIT.
// Source: app/domain/repository/post_repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/toshiykst/golang-rest-api/app/domain/model"
)

// MockPostRepository is a mock of PostRepository interface.
type MockPostRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPostRepositoryMockRecorder
}

// MockPostRepositoryMockRecorder is the mock recorder for MockPostRepository.
type MockPostRepositoryMockRecorder struct {
	mock *MockPostRepository
}

// NewMockPostRepository creates a new mock instance.
func NewMockPostRepository(ctrl *gomock.Controller) *MockPostRepository {
	mock := &MockPostRepository{ctrl: ctrl}
	mock.recorder = &MockPostRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPostRepository) EXPECT() *MockPostRepositoryMockRecorder {
	return m.recorder
}

// CreatePost mocks base method.
func (m *MockPostRepository) CreatePost(arg0 *model.Post) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePost", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatePost indicates an expected call of CreatePost.
func (mr *MockPostRepositoryMockRecorder) CreatePost(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePost", reflect.TypeOf((*MockPostRepository)(nil).CreatePost), arg0)
}

// DeletePost mocks base method.
func (m *MockPostRepository) DeletePost(arg0 *model.Post) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePost", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePost indicates an expected call of DeletePost.
func (mr *MockPostRepositoryMockRecorder) DeletePost(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePost", reflect.TypeOf((*MockPostRepository)(nil).DeletePost), arg0)
}

// FindPost mocks base method.
func (m *MockPostRepository) FindPost(id int) (model.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindPost", id)
	ret0, _ := ret[0].(model.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindPost indicates an expected call of FindPost.
func (mr *MockPostRepositoryMockRecorder) FindPost(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindPost", reflect.TypeOf((*MockPostRepository)(nil).FindPost), id)
}

// FindPosts mocks base method.
func (m *MockPostRepository) FindPosts() (model.Posts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindPosts")
	ret0, _ := ret[0].(model.Posts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindPosts indicates an expected call of FindPosts.
func (mr *MockPostRepositoryMockRecorder) FindPosts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindPosts", reflect.TypeOf((*MockPostRepository)(nil).FindPosts))
}

// UpdatePost mocks base method.
func (m *MockPostRepository) UpdatePost(arg0 *model.Post) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePost", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePost indicates an expected call of UpdatePost.
func (mr *MockPostRepositoryMockRecorder) UpdatePost(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePost", reflect.TypeOf((*MockPostRepository)(nil).UpdatePost), arg0)
}