// Code generated by MockGen. DO NOT EDIT.
// Source: usecase.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/internal/entity"
)

// MockIUseCase is a mock of IUseCase interface.
type MockIUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockIUseCaseMockRecorder
}

// MockIUseCaseMockRecorder is the mock recorder for MockIUseCase.
type MockIUseCaseMockRecorder struct {
	mock *MockIUseCase
}

// NewMockIUseCase creates a new mock instance.
func NewMockIUseCase(ctrl *gomock.Controller) *MockIUseCase {
	mock := &MockIUseCase{ctrl: ctrl}
	mock.recorder = &MockIUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUseCase) EXPECT() *MockIUseCaseMockRecorder {
	return m.recorder
}

// AddComment mocks base method.
func (m *MockIUseCase) AddComment(ctx context.Context, postID string, comment entity.Comment) (entity.PostExtend, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddComment", ctx, postID, comment)
	ret0, _ := ret[0].(entity.PostExtend)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddComment indicates an expected call of AddComment.
func (mr *MockIUseCaseMockRecorder) AddComment(ctx, postID, comment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddComment", reflect.TypeOf((*MockIUseCase)(nil).AddComment), ctx, postID, comment)
}

// AddPost mocks base method.
func (m *MockIUseCase) AddPost(ctx context.Context, post entity.Post) (entity.PostExtend, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPost", ctx, post)
	ret0, _ := ret[0].(entity.PostExtend)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddPost indicates an expected call of AddPost.
func (mr *MockIUseCaseMockRecorder) AddPost(ctx, post interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPost", reflect.TypeOf((*MockIUseCase)(nil).AddPost), ctx, post)
}

// DeleteComment mocks base method.
func (m *MockIUseCase) DeleteComment(ctx context.Context, username, postID, commentID string) (entity.PostExtend, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteComment", ctx, username, postID, commentID)
	ret0, _ := ret[0].(entity.PostExtend)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteComment indicates an expected call of DeleteComment.
func (mr *MockIUseCaseMockRecorder) DeleteComment(ctx, username, postID, commentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteComment", reflect.TypeOf((*MockIUseCase)(nil).DeleteComment), ctx, username, postID, commentID)
}

// DeletePost mocks base method.
func (m *MockIUseCase) DeletePost(ctx context.Context, username, postID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePost", ctx, username, postID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePost indicates an expected call of DeletePost.
func (mr *MockIUseCaseMockRecorder) DeletePost(ctx, username, postID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePost", reflect.TypeOf((*MockIUseCase)(nil).DeletePost), ctx, username, postID)
}

// Downvote mocks base method.
func (m *MockIUseCase) Downvote(ctx context.Context, userID, postID string) (entity.PostExtend, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Downvote", ctx, userID, postID)
	ret0, _ := ret[0].(entity.PostExtend)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Downvote indicates an expected call of Downvote.
func (mr *MockIUseCaseMockRecorder) Downvote(ctx, userID, postID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Downvote", reflect.TypeOf((*MockIUseCase)(nil).Downvote), ctx, userID, postID)
}

// GetPost mocks base method.
func (m *MockIUseCase) GetPost(ctx context.Context, postID string) (entity.PostExtend, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPost", ctx, postID)
	ret0, _ := ret[0].(entity.PostExtend)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPost indicates an expected call of GetPost.
func (mr *MockIUseCaseMockRecorder) GetPost(ctx, postID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPost", reflect.TypeOf((*MockIUseCase)(nil).GetPost), ctx, postID)
}

// GetPosts mocks base method.
func (m *MockIUseCase) GetPosts(ctx context.Context) ([]entity.PostExtend, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPosts", ctx)
	ret0, _ := ret[0].([]entity.PostExtend)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPosts indicates an expected call of GetPosts.
func (mr *MockIUseCaseMockRecorder) GetPosts(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPosts", reflect.TypeOf((*MockIUseCase)(nil).GetPosts), ctx)
}

// GetPostsWithCategory mocks base method.
func (m *MockIUseCase) GetPostsWithCategory(ctx context.Context, category string) ([]entity.PostExtend, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostsWithCategory", ctx, category)
	ret0, _ := ret[0].([]entity.PostExtend)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostsWithCategory indicates an expected call of GetPostsWithCategory.
func (mr *MockIUseCaseMockRecorder) GetPostsWithCategory(ctx, category interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostsWithCategory", reflect.TypeOf((*MockIUseCase)(nil).GetPostsWithCategory), ctx, category)
}

// GetPostsWithUser mocks base method.
func (m *MockIUseCase) GetPostsWithUser(ctx context.Context, username string) ([]entity.PostExtend, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostsWithUser", ctx, username)
	ret0, _ := ret[0].([]entity.PostExtend)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostsWithUser indicates an expected call of GetPostsWithUser.
func (mr *MockIUseCaseMockRecorder) GetPostsWithUser(ctx, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostsWithUser", reflect.TypeOf((*MockIUseCase)(nil).GetPostsWithUser), ctx, username)
}

// Login mocks base method.
func (m *MockIUseCase) Login(ctx context.Context, username, password string) (entity.UserExtend, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx, username, password)
	ret0, _ := ret[0].(entity.UserExtend)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockIUseCaseMockRecorder) Login(ctx, username, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockIUseCase)(nil).Login), ctx, username, password)
}

// SignUp mocks base method.
func (m *MockIUseCase) SignUp(ctx context.Context, user entity.User) (entity.UserExtend, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignUp", ctx, user)
	ret0, _ := ret[0].(entity.UserExtend)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignUp indicates an expected call of SignUp.
func (mr *MockIUseCaseMockRecorder) SignUp(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUp", reflect.TypeOf((*MockIUseCase)(nil).SignUp), ctx, user)
}

// Unvote mocks base method.
func (m *MockIUseCase) Unvote(ctx context.Context, userID, postID string) (entity.PostExtend, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unvote", ctx, userID, postID)
	ret0, _ := ret[0].(entity.PostExtend)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unvote indicates an expected call of Unvote.
func (mr *MockIUseCaseMockRecorder) Unvote(ctx, userID, postID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unvote", reflect.TypeOf((*MockIUseCase)(nil).Unvote), ctx, userID, postID)
}

// Upvote mocks base method.
func (m *MockIUseCase) Upvote(ctx context.Context, userID, postID string) (entity.PostExtend, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upvote", ctx, userID, postID)
	ret0, _ := ret[0].(entity.PostExtend)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Upvote indicates an expected call of Upvote.
func (mr *MockIUseCaseMockRecorder) Upvote(ctx, userID, postID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upvote", reflect.TypeOf((*MockIUseCase)(nil).Upvote), ctx, userID, postID)
}

// MockIGeneratorID is a mock of IGeneratorID interface.
type MockIGeneratorID struct {
	ctrl     *gomock.Controller
	recorder *MockIGeneratorIDMockRecorder
}

// MockIGeneratorIDMockRecorder is the mock recorder for MockIGeneratorID.
type MockIGeneratorIDMockRecorder struct {
	mock *MockIGeneratorID
}

// NewMockIGeneratorID creates a new mock instance.
func NewMockIGeneratorID(ctrl *gomock.Controller) *MockIGeneratorID {
	mock := &MockIGeneratorID{ctrl: ctrl}
	mock.recorder = &MockIGeneratorIDMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIGeneratorID) EXPECT() *MockIGeneratorIDMockRecorder {
	return m.recorder
}

// Generate mocks base method.
func (m *MockIGeneratorID) Generate(ctx context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generate", ctx)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Generate indicates an expected call of Generate.
func (mr *MockIGeneratorIDMockRecorder) Generate(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generate", reflect.TypeOf((*MockIGeneratorID)(nil).Generate), ctx)
}