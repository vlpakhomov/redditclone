package interfaces

import (
	"context"
	"errors"

	"gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/internal/entity"
)

var (
	ErrPostNotExists    = errors.New("post doesn't exists")
	ErrUserNotExists    = errors.New("user doesn't exists")
	ErrCommentNotExists = errors.New("comment doesn't exists")
)

//go:generate mockgen -source=repository.go -destination=../mock/repository_mock.go -package=mock IUserRepository
type IUserRepository interface {
	Add(ctx context.Context, user entity.UserExtend) error
	Get(ctx context.Context, username string) (entity.UserExtend, error)
	Contains(ctx context.Context, username string) (bool, error)
}

//go:generate mockgen -source=repository.go -destination=../mock/repository_mock.go -package=mock IPostRepository
type IPostRepository interface {
	Add(ctx context.Context, post entity.PostExtend) error
	Get(ctx context.Context, postID string) (entity.PostExtend, error)
	GetWhereCategory(ctx context.Context, category string) ([]entity.PostExtend, error)
	GetWhereUsername(ctx context.Context, username string) ([]entity.PostExtend, error)
	GetAll(ctx context.Context) ([]entity.PostExtend, error)
	Update(ctx context.Context, postID string, newPost entity.PostExtend) error
	Delete(ctx context.Context, postID string) error

	AddComment(ctx context.Context, postID string, comment entity.CommentExtend) (entity.PostExtend, error)
	GetComment(ctx context.Context, postID string, commentID string) (entity.CommentExtend, error)
	DeleteComment(ctx context.Context, postID string, commentID string) (entity.PostExtend, error)
}
