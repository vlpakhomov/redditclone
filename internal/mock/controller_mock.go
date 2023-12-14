// Code generated by MockGen. DO NOT EDIT.
// Source: controller.go

// Package mock is a generated GoMock package.
package mock

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	auth "gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/internal/controller/auth"
	entity "gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/internal/entity"
)

// MockIHandler is a mock of IHandler interface.
type MockIHandler struct {
	ctrl     *gomock.Controller
	recorder *MockIHandlerMockRecorder
}

// MockIHandlerMockRecorder is the mock recorder for MockIHandler.
type MockIHandlerMockRecorder struct {
	mock *MockIHandler
}

// NewMockIHandler creates a new mock instance.
func NewMockIHandler(ctrl *gomock.Controller) *MockIHandler {
	mock := &MockIHandler{ctrl: ctrl}
	mock.recorder = &MockIHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIHandler) EXPECT() *MockIHandlerMockRecorder {
	return m.recorder
}

// AddComment mocks base method.
func (m *MockIHandler) AddComment(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddComment", w, r)
}

// AddComment indicates an expected call of AddComment.
func (mr *MockIHandlerMockRecorder) AddComment(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddComment", reflect.TypeOf((*MockIHandler)(nil).AddComment), w, r)
}

// AddPost mocks base method.
func (m *MockIHandler) AddPost(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddPost", w, r)
}

// AddPost indicates an expected call of AddPost.
func (mr *MockIHandlerMockRecorder) AddPost(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPost", reflect.TypeOf((*MockIHandler)(nil).AddPost), w, r)
}

// DeleteComment mocks base method.
func (m *MockIHandler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteComment", w, r)
}

// DeleteComment indicates an expected call of DeleteComment.
func (mr *MockIHandlerMockRecorder) DeleteComment(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteComment", reflect.TypeOf((*MockIHandler)(nil).DeleteComment), w, r)
}

// DeletePost mocks base method.
func (m *MockIHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeletePost", w, r)
}

// DeletePost indicates an expected call of DeletePost.
func (mr *MockIHandlerMockRecorder) DeletePost(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePost", reflect.TypeOf((*MockIHandler)(nil).DeletePost), w, r)
}

// Downvote mocks base method.
func (m *MockIHandler) Downvote(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Downvote", w, r)
}

// Downvote indicates an expected call of Downvote.
func (mr *MockIHandlerMockRecorder) Downvote(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Downvote", reflect.TypeOf((*MockIHandler)(nil).Downvote), w, r)
}

// GetPost mocks base method.
func (m *MockIHandler) GetPost(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetPost", w, r)
}

// GetPost indicates an expected call of GetPost.
func (mr *MockIHandlerMockRecorder) GetPost(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPost", reflect.TypeOf((*MockIHandler)(nil).GetPost), w, r)
}

// GetPosts mocks base method.
func (m *MockIHandler) GetPosts(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetPosts", w, r)
}

// GetPosts indicates an expected call of GetPosts.
func (mr *MockIHandlerMockRecorder) GetPosts(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPosts", reflect.TypeOf((*MockIHandler)(nil).GetPosts), w, r)
}

// GetPostsWithCategory mocks base method.
func (m *MockIHandler) GetPostsWithCategory(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetPostsWithCategory", w, r)
}

// GetPostsWithCategory indicates an expected call of GetPostsWithCategory.
func (mr *MockIHandlerMockRecorder) GetPostsWithCategory(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostsWithCategory", reflect.TypeOf((*MockIHandler)(nil).GetPostsWithCategory), w, r)
}

// GetPostsWithUser mocks base method.
func (m *MockIHandler) GetPostsWithUser(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetPostsWithUser", w, r)
}

// GetPostsWithUser indicates an expected call of GetPostsWithUser.
func (mr *MockIHandlerMockRecorder) GetPostsWithUser(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostsWithUser", reflect.TypeOf((*MockIHandler)(nil).GetPostsWithUser), w, r)
}

// Index mocks base method.
func (m *MockIHandler) Index(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Index", w, r)
}

// Index indicates an expected call of Index.
func (mr *MockIHandlerMockRecorder) Index(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Index", reflect.TypeOf((*MockIHandler)(nil).Index), w, r)
}

// Login mocks base method.
func (m *MockIHandler) Login(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Login", w, r)
}

// Login indicates an expected call of Login.
func (mr *MockIHandlerMockRecorder) Login(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockIHandler)(nil).Login), w, r)
}

// SignUp mocks base method.
func (m *MockIHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SignUp", w, r)
}

// SignUp indicates an expected call of SignUp.
func (mr *MockIHandlerMockRecorder) SignUp(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUp", reflect.TypeOf((*MockIHandler)(nil).SignUp), w, r)
}

// Unvote mocks base method.
func (m *MockIHandler) Unvote(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Unvote", w, r)
}

// Unvote indicates an expected call of Unvote.
func (mr *MockIHandlerMockRecorder) Unvote(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unvote", reflect.TypeOf((*MockIHandler)(nil).Unvote), w, r)
}

// Upvote mocks base method.
func (m *MockIHandler) Upvote(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Upvote", w, r)
}

// Upvote indicates an expected call of Upvote.
func (mr *MockIHandlerMockRecorder) Upvote(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upvote", reflect.TypeOf((*MockIHandler)(nil).Upvote), w, r)
}

// MockIAuthManager is a mock of IAuthManager interface.
type MockIAuthManager struct {
	ctrl     *gomock.Controller
	recorder *MockIAuthManagerMockRecorder
}

// MockIAuthManagerMockRecorder is the mock recorder for MockIAuthManager.
type MockIAuthManagerMockRecorder struct {
	mock *MockIAuthManager
}

// NewMockIAuthManager creates a new mock instance.
func NewMockIAuthManager(ctrl *gomock.Controller) *MockIAuthManager {
	mock := &MockIAuthManager{ctrl: ctrl}
	mock.recorder = &MockIAuthManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIAuthManager) EXPECT() *MockIAuthManagerMockRecorder {
	return m.recorder
}

// CreateSession mocks base method.
func (m *MockIAuthManager) CreateSession(user entity.UserExtend) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", user)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSession indicates an expected call of CreateSession.
func (mr *MockIAuthManagerMockRecorder) CreateSession(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockIAuthManager)(nil).CreateSession), user)
}

// CreateToken mocks base method.
func (m *MockIAuthManager) CreateToken(user entity.UserExtend) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateToken", user)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateToken indicates an expected call of CreateToken.
func (mr *MockIAuthManagerMockRecorder) CreateToken(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateToken", reflect.TypeOf((*MockIAuthManager)(nil).CreateToken), user)
}

// DeleteSession mocks base method.
func (m *MockIAuthManager) DeleteSession(sid auth.SessionID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSession", sid)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSession indicates an expected call of DeleteSession.
func (mr *MockIAuthManagerMockRecorder) DeleteSession(sid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSession", reflect.TypeOf((*MockIAuthManager)(nil).DeleteSession), sid)
}

// GetSession mocks base method.
func (m *MockIAuthManager) GetSession(session auth.Session) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSession", session)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetSession indicates an expected call of GetSession.
func (mr *MockIAuthManagerMockRecorder) GetSession(session interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSession", reflect.TypeOf((*MockIAuthManager)(nil).GetSession), session)
}

// ParseToken mocks base method.
func (m *MockIAuthManager) ParseToken(accessToken string) (*auth.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseToken", accessToken)
	ret0, _ := ret[0].(*auth.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseToken indicates an expected call of ParseToken.
func (mr *MockIAuthManagerMockRecorder) ParseToken(accessToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseToken", reflect.TypeOf((*MockIAuthManager)(nil).ParseToken), accessToken)
}
