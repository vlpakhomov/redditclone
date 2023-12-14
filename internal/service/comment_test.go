package service

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/internal/entity"
	"gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/internal/interfaces"
	"gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/internal/mock"
)

func TestAddCommentSuccess(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	postRepo := mock.NewMockIPostRepository(ctrl)
	s := NewService(postRepo, nil)

	// AddCommentSuccess
	{
		postID := "postID"
		comment := entity.CommentExtend{}
		postReturn := entity.PostExtend{
			ID: "id",
		}
		postRepo.EXPECT().AddComment(ctx, postID, comment).Return(postReturn, nil)

		post, errAddComment := s.AddComment(ctx, postID, comment)
		require.Nilf(t, errAddComment, "[service.test.AddCommentSuccess]: s.AddComment(ctx, user) failed: %v", errAddComment)

		if !reflect.DeepEqual(postReturn, post) {
			t.Errorf("[service.test.AddCommentSuccess]: results not match, want %v, have %v", postReturn, post)
			return
		}
	}
}

func TestAddCommentFailure(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	postRepo := mock.NewMockIPostRepository(ctrl)
	s := NewService(postRepo, nil)

	// ErrAddComment
	{
		postID := "postID"
		comment := entity.CommentExtend{}
		postReturn := entity.PostExtend{}
		postRepo.EXPECT().AddComment(ctx, postID, comment).Return(postReturn, errors.New("postRepo.AddComment failed"))

		post, errAddComment := s.AddComment(ctx, postID, comment)
		require.NotNil(t, errAddComment, "[service.test.AddCommentFailure(ErrAddComment)]: expected error from s.AddComment(ctx, user)")

		if !reflect.DeepEqual(postReturn, post) {
			t.Errorf("[service.test.AddCommentFailure(ErrAddComment)]: results not match, want %v, have %v", postReturn, post)
			return
		}
	}
}

func TestDeleteCommentSuccess(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	postRepo := mock.NewMockIPostRepository(ctrl)
	s := NewService(postRepo, nil)

	// DeleteCommentSuccess
	{
		username := "username"
		postID := "postID"
		commentID := "commentID"
		commentReturn := entity.CommentExtend{
			Comment: entity.Comment{
				Author: entity.Author{
					Username: username,
				},
			},
		}
		postRepo.EXPECT().GetComment(ctx, postID, commentID).Return(commentReturn, nil)

		postReturn := entity.PostExtend{
			ID: "id",
		}
		postRepo.EXPECT().DeleteComment(ctx, postID, commentID).Return(postReturn, nil)

		post, errDeleteComment := s.DeleteComment(ctx, username, postID, commentID)
		require.Nilf(t, errDeleteComment, "[service.test.DeleteCommentSuccess]: s.DeleteComment(ctx, username, postID, commentID) failed: %v", errDeleteComment)

		if !reflect.DeepEqual(postReturn, post) {
			t.Errorf("[service.test.DeleteCommentSuccess]: results not match, want %v, have %v", postReturn, post)
			return
		}
	}
}

func TestDeleteCommentFailure(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	postRepo := mock.NewMockIPostRepository(ctrl)
	s := NewService(postRepo, nil)

	// ErrRepoGetComment
	{
		username := "username"
		postID := "postID"
		commentID := "commentID"
		commentReturn := entity.CommentExtend{
			Comment: entity.Comment{
				Author: entity.Author{
					Username: username,
				},
			},
		}
		postRepo.EXPECT().GetComment(ctx, postID, commentID).Return(commentReturn, errors.New("postRepo.GetComment failed"))

		postReturn := entity.PostExtend{}

		post, errDeleteComment := s.DeleteComment(ctx, username, postID, commentID)
		require.NotNil(t, errDeleteComment, "[service.test.DeleteCommentFailure(ErrRepoGetComment)]: expected error from s.DeleteComment(ctx, username, postID, commentID)")

		if !reflect.DeepEqual(postReturn, post) {
			t.Errorf("[service.test.DeleteCommentFailure(ErrRepoGetComment)]: results not match, want %v, have %v", postReturn, post)
			return
		}

	}

	// ErrAccessDeined
	{
		username := "username"
		postID := "postID"
		commentID := "commentID"
		commentReturn := entity.CommentExtend{}
		postRepo.EXPECT().GetComment(ctx, postID, commentID).Return(commentReturn, nil)

		postReturn := entity.PostExtend{}

		post, errDeleteComment := s.DeleteComment(ctx, username, postID, commentID)
		require.NotNil(t, errDeleteComment, "[service.test.DeleteCommentFailure(ErrAccessDeined)]: expected error from s.DeleteComment(ctx, username, postID, commentID)")

		if !reflect.DeepEqual(postReturn, post) {
			t.Errorf("[service.test.DeleteCommentFailure(ErrAccessDeined)]: results not match, want %v, have %v", postReturn, post)
			return
		}

		require.Truef(t, errors.Is(errDeleteComment, interfaces.ErrAccessDenied), "[service.test.DeleteCommentFailure(ErrAccessDeined)]: error not match, want %v, have %v", interfaces.ErrAccessDenied, errDeleteComment)
	}

	// ErrRepoDeleteComment
	{
		username := "username"
		postID := "postID"
		commentID := "commentID"
		commentReturn := entity.CommentExtend{
			Comment: entity.Comment{
				Author: entity.Author{
					Username: username,
				},
			},
		}
		postRepo.EXPECT().GetComment(ctx, postID, commentID).Return(commentReturn, nil)

		postReturn := entity.PostExtend{}
		postRepo.EXPECT().DeleteComment(ctx, postID, commentID).Return(postReturn, errors.New("postRepo.DeleteComment failed"))

		post, errDeleteComment := s.DeleteComment(ctx, username, postID, commentID)
		require.NotNil(t, errDeleteComment, "[service.test.DeleteCommentFailure(ErrRepoDeleteComment)]: expected error from s.DeleteComment(ctx, username, postID, commentID)")

		if !reflect.DeepEqual(postReturn, post) {
			t.Errorf("[service.test.DeleteCommentFailure(ErrRepoDeleteComment)]: results not match, want %v, have %v", postReturn, post)
			return
		}
	}
}
