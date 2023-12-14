package usecase

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

	serv := mock.NewMockIService(ctrl)
	genID := mock.NewMockIGeneratorID(ctrl)
	u := NewUseCase(serv, genID)

	// AddCommentSuccess
	{
		commentIDReturn := "commentID"
		genID.EXPECT().Generate(ctx).Return(commentIDReturn, nil)

		postID := "postID"
		comment := entity.Comment{}
		commentExtend := entity.NewCommentExtend(comment, commentIDReturn)
		postReturn := entity.PostExtend{
			Post: entity.Post{
				Comments: []entity.CommentExtend{
					*commentExtend,
				},
			},
			ID: postID,
		}
		serv.EXPECT().AddComment(ctx, postID, *commentExtend).Return(postReturn, nil)

		post, errAddComment := u.AddComment(ctx, postID, comment)
		require.Nilf(t, errAddComment, "[usecase.test.AddCommentSuccess]: u.AddComment(ctx, postID, comment) failed: %v", errAddComment)

		if !reflect.DeepEqual(postReturn, post) {
			t.Errorf("[usecase.test.AddCommentSuccess]: results not match, want %v, have %v", postReturn, post)
			return
		}

	}
}

func TestAddCommentFailure(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	serv := mock.NewMockIService(ctrl)
	genID := mock.NewMockIGeneratorID(ctrl)
	u := NewUseCase(serv, genID)

	// ErrConstructObject
	{
		idReturn := ""
		genID.EXPECT().Generate(ctx).Return(idReturn, errors.New("genID.Generate failed"))

		postID := "postID"
		comment := entity.Comment{}

		postReturn := entity.PostExtend{}

		post, errAddComment := u.AddComment(ctx, postID, comment)
		require.NotNil(t, errAddComment, "[usecase.test.AddCommentFailure(ErrConstructObject)]: expected error from u.AddComment(ctx, postID, comment)")

		if !reflect.DeepEqual(postReturn, post) {
			t.Errorf("[usecase.test.AddCommentFailure(ErrConstructObject)]: results not match, want %v, have %v", postReturn, post)
			return
		}

		require.Truef(t, errors.Is(errAddComment, interfaces.ErrConstructObject), "[usecase.test.AddCommentFailure(ErrConstructObject)]: error not match, want %v, have %v", interfaces.ErrConstructObject, errAddComment)
	}

	// ErrServAddComment
	{
		idReturn := "commentID"
		genID.EXPECT().Generate(ctx).Return(idReturn, nil)

		postID := "postID"
		comment := entity.Comment{}
		commentExtend := entity.NewCommentExtend(comment, idReturn)
		postReturn := entity.PostExtend{}
		serv.EXPECT().AddComment(ctx, postID, *commentExtend).Return(postReturn, errors.New("service.AddComment failed"))

		post, errAddComment := u.AddComment(ctx, postID, comment)
		require.NotNil(t, errAddComment, "[usecase.test.AddCommentFailure(ErrServAddComment)]: expected error from u.AddComment(ctx, postID, comment)")

		if !reflect.DeepEqual(postReturn, post) {
			t.Errorf("[usecase.test.AddCommentFailure(ErrServAddComment)]: results not match, want %v, have %v", postReturn, post)
			return
		}
	}
}

func TestDeleteCommentSuccess(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	serv := mock.NewMockIService(ctrl)
	genID := mock.NewMockIGeneratorID(ctrl)
	u := NewUseCase(serv, genID)

	// DeleteCommentSuccess
	{
		username := "username"
		postID := "postID"
		commentID := "commentID"
		postReturn := entity.PostExtend{
			Post: entity.Post{
				Author: entity.Author{
					Username: username,
				},
			},
			ID: postID,
		}

		serv.EXPECT().DeleteComment(ctx, username, postID, commentID).Return(postReturn, nil)

		post, errDeleteComment := u.DeleteComment(ctx, username, postID, commentID)
		require.Nilf(t, errDeleteComment, "[usecase.test.DeleteCommentSuccess]: u.DeleteComment(ctx, username, postID, commentID) failed: %v", errDeleteComment)

		if !reflect.DeepEqual(postReturn, post) {
			t.Errorf("[usecase.test.DeleteCommentSuccess]: results not match, want %v, have %v", postReturn, post)
			return
		}
	}
}

func TestDeleteCommentFailure(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	serv := mock.NewMockIService(ctrl)
	genID := mock.NewMockIGeneratorID(ctrl)
	u := NewUseCase(serv, genID)

	// ErrServDeleteComment
	{
		username := "username"
		postID := "postID"
		commentID := "commentID"
		postReturn := entity.PostExtend{}

		serv.EXPECT().DeleteComment(ctx, username, postID, commentID).Return(postReturn, errors.New("serv.DeleteComment failed"))

		post, errDeleteComment := u.DeleteComment(ctx, username, postID, commentID)
		require.NotNil(t, errDeleteComment, "[usecase.test.DeleteCommentFailure(ErrServDeleteComment)]: expected error from u.AddComment(ctx, postID, comment)")

		if !reflect.DeepEqual(postReturn, post) {
			t.Errorf("[usecase.test.DeleteCommentFailure(ErrServDeleteComment)]: results not match, want %v, have %v", postReturn, post)
			return
		}
	}
}
