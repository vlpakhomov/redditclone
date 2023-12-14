// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/internal/entity"
)

// MockIUserRepository is a mock of IUserRepository interface.
type MockIUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIUserRepositoryMockRecorder
}

// MockIUserRepositoryMockRecorder is the mock recorder for MockIUserRepository.
type MockIUserRepositoryMockRecorder struct {
	mock *MockIUserRepository
}

// NewMockIUserRepository creates a new mock instance.
func NewMockIUserRepository(ctrl *gomock.Controller) *MockIUserRepository {
	mock := &MockIUserRepository{ctrl: ctrl}
	mock.recorder = &MockIUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserRepository) EXPECT() *MockIUserRepositoryMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockIUserRepository) Add(ctx context.Context, user entity.UserExtend) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockIUserRepositoryMockRecorder) Add(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockIUserRepository)(nil).Add), ctx, user)
}

// Contains mocks base method.
func (m *MockIUserRepository) Contains(ctx context.Context, username string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Contains", ctx, username)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Contains indicates an expected call of Contains.
func (mr *MockIUserRepositoryMockRecorder) Contains(ctx, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Contains", reflect.TypeOf((*MockIUserRepository)(nil).Contains), ctx, username)
}

// Get mocks base method.
func (m *MockIUserRepository) Get(ctx context.Context, username string) (entity.UserExtend, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, username)
	ret0, _ := ret[0].(entity.UserExtend)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockIUserRepositoryMockRecorder) Get(ctx, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIUserRepository)(nil).Get), ctx, username)
}

// MockIPostRepository is a mock of IPostRepository interface.
type MockIPostRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIPostRepositoryMockRecorder
}

// MockIPostRepositoryMockRecorder is the mock recorder for MockIPostRepository.
type MockIPostRepositoryMockRecorder struct {
	mock *MockIPostRepository
}

// NewMockIPostRepository creates a new mock instance.
func NewMockIPostRepository(ctrl *gomock.Controller) *MockIPostRepository {
	mock := &MockIPostRepository{ctrl: ctrl}
	mock.recorder = &MockIPostRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIPostRepository) EXPECT() *MockIPostRepositoryMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockIPostRepository) Add(ctx context.Context, post entity.PostExtend) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", ctx, post)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockIPostRepositoryMockRecorder) Add(ctx, post interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockIPostRepository)(nil).Add), ctx, post)
}

// AddComment mocks base method.
func (m *MockIPostRepository) AddComment(ctx context.Context, postID string, comment entity.CommentExtend) (entity.PostExtend, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddComment", ctx, postID, comment)
	ret0, _ := ret[0].(entity.PostExtend)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddComment indicates an expected call of AddComment.
func (mr *MockIPostRepositoryMockRecorder) AddComment(ctx, postID, comment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddComment", reflect.TypeOf((*MockIPostRepository)(nil).AddComment), ctx, postID, comment)
}

// Delete mocks base method.
func (m *MockIPostRepository) Delete(ctx context.Context, postID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, postID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIPostRepositoryMockRecorder) Delete(ctx, postID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIPostRepository)(nil).Delete), ctx, postID)
}

// DeleteComment mocks base method.
func (m *MockIPostRepository) DeleteComment(ctx context.Context, postID, commentID string) (entity.PostExtend, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteComment", ctx, postID, commentID)
	ret0, _ := ret[0].(entity.PostExtend)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteComment indicates an expected call of DeleteComment.
func (mr *MockIPostRepositoryMockRecorder) DeleteComment(ctx, postID, commentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteComment", reflect.TypeOf((*MockIPostRepository)(nil).DeleteComment), ctx, postID, commentID)
}

// Get mocks base method.
func (m *MockIPostRepository) Get(ctx context.Context, postID string) (entity.PostExtend, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, postID)
	ret0, _ := ret[0].(entity.PostExtend)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockIPostRepositoryMockRecorder) Get(ctx, postID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIPostRepository)(nil).Get), ctx, postID)
}

// GetAll mocks base method.
func (m *MockIPostRepository) GetAll(ctx context.Context) ([]entity.PostExtend, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx)
	ret0, _ := ret[0].([]entity.PostExtend)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockIPostRepositoryMockRecorder) GetAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockIPostRepository)(nil).GetAll), ctx)
}

// GetComment mocks base method.
func (m *MockIPostRepository) GetComment(ctx context.Context, postID, commentID string) (entity.CommentExtend, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetComment", ctx, postID, commentID)
	ret0, _ := ret[0].(entity.CommentExtend)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetComment indicates an expected call of GetComment.
func (mr *MockIPostRepositoryMockRecorder) GetComment(ctx, postID, commentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComment", reflect.TypeOf((*MockIPostRepository)(nil).GetComment), ctx, postID, commentID)
}

// GetWhereCategory mocks base method.
func (m *MockIPostRepository) GetWhereCategory(ctx context.Context, category string) ([]entity.PostExtend, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWhereCategory", ctx, category)
	ret0, _ := ret[0].([]entity.PostExtend)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWhereCategory indicates an expected call of GetWhereCategory.
func (mr *MockIPostRepositoryMockRecorder) GetWhereCategory(ctx, category interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWhereCategory", reflect.TypeOf((*MockIPostRepository)(nil).GetWhereCategory), ctx, category)
}

// GetWhereUsername mocks base method.
func (m *MockIPostRepository) GetWhereUsername(ctx context.Context, username string) ([]entity.PostExtend, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWhereUsername", ctx, username)
	ret0, _ := ret[0].([]entity.PostExtend)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWhereUsername indicates an expected call of GetWhereUsername.
func (mr *MockIPostRepositoryMockRecorder) GetWhereUsername(ctx, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWhereUsername", reflect.TypeOf((*MockIPostRepository)(nil).GetWhereUsername), ctx, username)
}

// Update mocks base method.
func (m *MockIPostRepository) Update(ctx context.Context, postID string, newPost entity.PostExtend) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, postID, newPost)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockIPostRepositoryMockRecorder) Update(ctx, postID, newPost interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIPostRepository)(nil).Update), ctx, postID, newPost)
}