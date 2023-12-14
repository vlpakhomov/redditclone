package interfaces

import (
	"context"
	"errors"

	"gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/internal/entity"
)

var (
	ErrAccessDenied = errors.New("access denied")
)

//go:generate mockgen -source=service.go -destination=../mock/service_mock.go -package=mock IService
type IService interface {
	AddUser(ctx context.Context, user entity.UserExtend) error
	GetUser(ctx context.Context, username string) (entity.UserExtend, error)
	ContainsUser(ctx context.Context, username string) (bool, error)

	AddPost(ctx context.Context, post entity.PostExtend) error
	GetPost(ctx context.Context, postID string) (entity.PostExtend, error)
	GetPosts(ctx context.Context) ([]entity.PostExtend, error)
	GetPostsWithCategory(ctx context.Context, category string) ([]entity.PostExtend, error)
	GetPostsWithUser(ctx context.Context, username string) ([]entity.PostExtend, error)
	UpvotePost(ctx context.Context, userID string, postID string) (entity.PostExtend, error)
	DownvotePost(ctx context.Context, userID string, postID string) (entity.PostExtend, error)
	UnvotePost(ctx context.Context, userID string, postID string) (entity.PostExtend, error)
	DeletePost(ctx context.Context, username string, postID string) error

	AddComment(ctx context.Context, postID string, comment entity.CommentExtend) (entity.PostExtend, error)
	DeleteComment(ctx context.Context, username string, postID string, commentID string) (entity.PostExtend, error)

	SortPostsByTime(posts []entity.PostExtend) []entity.PostExtend
}
