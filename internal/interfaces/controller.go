package interfaces

import (
	"errors"
	"net/http"

	"gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/internal/controller/auth"
	"gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/internal/entity"
)

var (
	ErrUnknownErrorMsg = errors.New("unknown error")
	ErrSuccessMsg      = errors.New("success")
)

type IHandler interface {
	SignUp(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)

	GetPosts(w http.ResponseWriter, r *http.Request)
	AddPost(w http.ResponseWriter, r *http.Request)
	GetPostsWithCategory(w http.ResponseWriter, r *http.Request)
	GetPostsWithUser(w http.ResponseWriter, r *http.Request)
	GetPost(w http.ResponseWriter, r *http.Request)
	DeletePost(w http.ResponseWriter, r *http.Request)

	AddComment(w http.ResponseWriter, r *http.Request)
	DeleteComment(w http.ResponseWriter, r *http.Request)

	Upvote(w http.ResponseWriter, r *http.Request)
	Downvote(w http.ResponseWriter, r *http.Request)
	Unvote(w http.ResponseWriter, r *http.Request)

	Index(w http.ResponseWriter, r *http.Request)
}

//go:generate mockgen -source=controller.go -destination=../mock/controller_mock.go -package=mock IAuthManager
type IAuthManager interface {
	CreateToken(user entity.UserExtend) (string, error)
	ParseToken(accessToken string) (*auth.Session, error)
	CreateSession(user entity.UserExtend) (string, error)
	GetSession(session auth.Session) error
	DeleteSession(sid auth.SessionID) error
}
