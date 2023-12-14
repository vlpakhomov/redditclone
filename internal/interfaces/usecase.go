package interfaces

import (
	"context"
	"errors"

	"gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/internal/entity"
)

var (
	ErrUserExists      = errors.New("user already exists")
	ErrInvalidPassword = errors.New("invalID password")
	ErrConstructObject = errors.New("construct object error")
)

//go:generate mockgen -source=usecase.go -destination=../mock/usecase_mock.go -package=mock IUseCase
type IUseCase interface {
	SignUp(ctx context.Context, user entity.User) (entity.UserExtend, error)
	Login(ctx context.Context, username string, password string) (entity.UserExtend, error)

	GetPosts(ctx context.Context) ([]entity.PostExtend, error)
	AddPost(ctx context.Context, post entity.Post) (entity.PostExtend, error)
	GetPostsWithCategory(ctx context.Context, category string) ([]entity.PostExtend, error)
	GetPostsWithUser(ctx context.Context, username string) ([]entity.PostExtend, error)
	GetPost(ctx context.Context, postID string) (entity.PostExtend, error)
	DeletePost(ctx context.Context, username string, postID string) error

	AddComment(ctx context.Context, postID string, comment entity.Comment) (entity.PostExtend, error)
	DeleteComment(ctx context.Context, username string, postID string, commentID string) (entity.PostExtend, error)

	Upvote(ctx context.Context, userID string, postID string) (entity.PostExtend, error)
	Downvote(ctx context.Context, userID string, postID string) (entity.PostExtend, error)
	Unvote(ctx context.Context, userID string, postID string) (entity.PostExtend, error)
}

//go:generate mockgen -source=usecase.go -destination=../mock/usecase_mock.go -package=mock IGeneratorID
type IGeneratorID interface {
	Generate(ctx context.Context) (string, error)
}
