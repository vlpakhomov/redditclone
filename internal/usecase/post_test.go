package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/internal/entity"
	"gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/internal/mock"
)

func TestGetPostsSuccess(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	serv := mock.NewMockIService(ctrl)
	genID := mock.NewMockIGeneratorID(ctrl)
	u := NewUseCase(serv, genID)

	// GetPostsSuccess
	{
		postsReturn := []entity.PostExtend{
			{
				ID: "postID",
			},
		}
		serv.EXPECT().GetPosts(ctx).Return(postsReturn, nil)

		posts, errGetPosts := u.GetPosts(ctx)
		require.Nilf(t, errGetPosts, "[usecase.test.GetPostsSuccess]: u.GetPosts(ctx) failed: %v", errGetPosts)

		if !reflect.DeepEqual(postsReturn, posts) {
			t.Errorf("[usecase.test.GetPostsSuccess]: results not match, want %v, have %v", postsReturn, posts)
			return
		}
	}
}

func TestGetPostsFailure(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	serv := mock.NewMockIService(ctrl)
	genID := mock.NewMockIGeneratorID(ctrl)
	u := NewUseCase(serv, genID)

	// ErrServGetPosts
	{
		postsReturn := []entity.PostExtend{}
		serv.EXPECT().GetPosts(ctx).Return(postsReturn, errors.New("serv.GetPosts failed"))

		posts, errGetPosts := u.GetPosts(ctx)
		require.NotNil(t, errGetPosts, "[usecase.test.GetPostsFailure(ErrServGetPosts)]: expected error from u.AddComment(ctx, postID, comment)")

		if !reflect.DeepEqual(postsReturn, posts) {
			t.Errorf("[usecase.test.GetPostsFailure(ErrServGetPosts)]: results not match, want %v, have %v", postsReturn, posts)
			return
		}
	}
}

func TestAddPostSuccess(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	serv := mock.NewMockIService(ctrl)
	genID := mock.NewMockIGeneratorID(ctrl)
	u := NewUseCase(serv, genID)

	// AddPostSuccess
	{
		postIDReturn := "commentID"
		genID.EXPECT().Generate(ctx).Return(postIDReturn, nil)

		post := entity.Post{}
		postExtend := entity.NewPostExtend(post, postIDReturn)
		serv.EXPECT().AddPost(ctx, *postExtend).Return(nil)

		postResult, errAddPost := u.AddPost(ctx, post)
		require.Nilf(t, errAddPost, "[usecase.test.AddPostSuccess]: u.AddPost(ctx, post) failed: %v", errAddPost)

		if !reflect.DeepEqual(*postExtend, postResult) {
			t.Errorf("[usecase.test.AddPostSuccess]: results not match, want %v, have %v", postExtend, postResult)
			return
		}
	}
}

func TestAddPostFailure(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	serv := mock.NewMockIService(ctrl)
	genID := mock.NewMockIGeneratorID(ctrl)
	u := NewUseCase(serv, genID)

	// ErrConstructObject
	{
		postIDReturn := ""
		genID.EXPECT().Generate(ctx).Return(postIDReturn, errors.New("genID.Generate"))

		post := entity.Post{}
		postReturn := entity.PostExtend{}

		postResult, errAddPost := u.AddPost(ctx, post)
		require.NotNil(t, errAddPost, "[usecase.test.AddPostFailure(ErrConstructObject)]: expected error from u.AddPost(ctx, post)")

		if !reflect.DeepEqual(postReturn, postResult) {
			t.Errorf("[usecase.test.AddPostFailure(ErrConstructObject)]: results not match, want %v, have %v", postReturn, postResult)
			return
		}
	}

	// ErrServAddPost
	{
		postIDReturn := "commentID"
		genID.EXPECT().Generate(ctx).Return(postIDReturn, nil)

		post := entity.Post{}
		postExtend := entity.NewPostExtend(post, postIDReturn)
		serv.EXPECT().AddPost(ctx, *postExtend).Return(errors.New("serv.AddPost failed"))

		postReturn := entity.PostExtend{}

		postResult, errAddPost := u.AddPost(ctx, post)
		require.NotNil(t, errAddPost, "[usecase.test.AddPostFailure(ErrServAddPost)]: expected error from u.AddPost(ctx, post)")

		if !reflect.DeepEqual(postReturn, postResult) {
			t.Errorf("[usecase.test.AddPostFailure(ErrServAddPost)]: results not match, want %v, have %v", postReturn, postResult)
			return
		}
	}
}

func TestGetPostsWithCategorySuccess(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	serv := mock.NewMockIService(ctrl)
	genID := mock.NewMockIGeneratorID(ctrl)
	u := NewUseCase(serv, genID)

	// GetPostsWithCategorySuccess
	{
		category := "category"
		postsReturn := []entity.PostExtend{
			{
				Post: entity.Post{
					Category: category,
				},
				ID: "postID",
			},
		}
		serv.EXPECT().GetPostsWithCategory(ctx, category).Return(postsReturn, nil)

		posts, errGetPosts := u.GetPostsWithCategory(ctx, category)
		require.Nilf(t, errGetPosts, "[usecase.test.GetPostsWithCategorySuccess]:  u.GetPostsWithCategory(ctx, category) failed: %v", errGetPosts)

		if !reflect.DeepEqual(postsReturn, posts) {
			t.Errorf("[usecase.test.GetPostsWithCategorySuccess]: results not match, want %v, have %v", postsReturn, posts)
			return
		}
	}
}

func TestGetPostsWithCategoryFailure(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	serv := mock.NewMockIService(ctrl)
	genID := mock.NewMockIGeneratorID(ctrl)
	u := NewUseCase(serv, genID)

	// ErrServGetPostsWithCategory
	{
		category := "category"
		postsReturn := []entity.PostExtend{}
		serv.EXPECT().GetPostsWithCategory(ctx, category).Return(postsReturn, errors.New("serv.GetPostsWithCategory failed"))

		posts, errGetPostsWithCategory := u.GetPostsWithCategory(ctx, category)
		require.NotNil(t, errGetPostsWithCategory, "[usecase.test.GetPostsWithCategoryFailure(ErrServGetPostsWithCategory)]: expected error from u.GetPostsWithCategory(ctx, category)")

		if !reflect.DeepEqual(postsReturn, posts) {
			t.Errorf("[usecase.test.GetPostsWithCategoryFailure(ErrServGetPostsWithCategory)]: results not match, want %v, have %v", postsReturn, posts)
			return
		}
	}
}

func TestGetPostsWithUserSuccess(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	serv := mock.NewMockIService(ctrl)
	genID := mock.NewMockIGeneratorID(ctrl)
	u := NewUseCase(serv, genID)

	// GetPostsWithUserSuccess
	{
		username := "username"
		postsReturn := []entity.PostExtend{
			{
				Post: entity.Post{
					Author: entity.Author{
						Username: username,
					},
				},
				ID: "postID",
			},
		}
		serv.EXPECT().GetPostsWithUser(ctx, username).Return(postsReturn, nil)

		posts, errGetPostsWithUser := u.GetPostsWithUser(ctx, username)
		require.Nilf(t, errGetPostsWithUser, "[usecase.test.GetPostsWithUserSuccess]:  u.GetPostsWithUser(ctx, username) failed: %v", errGetPostsWithUser)

		if !reflect.DeepEqual(postsReturn, posts) {
			t.Errorf("[usecase.test.GetPostsWithUserSuccess]: results not match, want %v, have %v", postsReturn, posts)
			return
		}
	}
}

func TestGetPostsWithUserFailure(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	serv := mock.NewMockIService(ctrl)
	genID := mock.NewMockIGeneratorID(ctrl)
	u := NewUseCase(serv, genID)

	// ErrServGetPostsWithUser
	{
		username := "username"
		postsReturn := []entity.PostExtend{}
		serv.EXPECT().GetPostsWithUser(ctx, username).Return(postsReturn, errors.New("serv.GetPostsWithUser failed"))

		posts, errGetPostsWithUser := u.GetPostsWithUser(ctx, username)
		require.NotNil(t, errGetPostsWithUser, "[usecase.test.GetPostsWithUserFailure(ErrServGetPostsWithUser)]: expected error from u.GetPostsWithUser(ctx, username)")

		if !reflect.DeepEqual(postsReturn, posts) {
			t.Errorf("[usecase.test.usecase.test.GetPostsWithUserFailure(ErrServGetPostsWithUser)]: results not match, want %v, have %v", postsReturn, posts)
			return
		}
	}
}

func TestGetPostSucess(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	serv := mock.NewMockIService(ctrl)
	genID := mock.NewMockIGeneratorID(ctrl)
	u := NewUseCase(serv, genID)

	// GetPostSuccess
	{
		postID := "postID"
		postReturn := entity.PostExtend{
			ID: postID,
		}
		serv.EXPECT().GetPost(ctx, postID).Return(postReturn, nil)

		posts, errGetPost := u.GetPost(ctx, postID)
		require.Nilf(t, errGetPost, "[usecase.test.GetPostSucess]:  u.GetPost(ctx, postID) failed: %v", errGetPost)

		if !reflect.DeepEqual(postReturn, posts) {
			t.Errorf("[usecase.test.GetPostSucess]: results not match, want %v, have %v", postReturn, posts)
			return
		}
	}
}

func TestGetPostFailure(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	serv := mock.NewMockIService(ctrl)
	genID := mock.NewMockIGeneratorID(ctrl)
	u := NewUseCase(serv, genID)

	// ErrServGetPost
	{
		postID := "postID"
		postReturn := entity.PostExtend{}
		serv.EXPECT().GetPost(ctx, postID).Return(postReturn, errors.New("serv.GetPost failed"))

		posts, errGetPost := u.GetPost(ctx, postID)
		require.NotNil(t, errGetPost, "[usecase.test.GetPostFailure(ErrServGetPost)]: expected error from u.GetPost(ctx, postID)")

		if !reflect.DeepEqual(postReturn, posts) {
			t.Errorf("[usecase.test.GetPostFailure(ErrServGetPost)]: results not match, want %v, have %v", postReturn, posts)
			return
		}
	}
}

func TestDeletePostSuccess(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	serv := mock.NewMockIService(ctrl)
	genID := mock.NewMockIGeneratorID(ctrl)
	u := NewUseCase(serv, genID)

	// DeletePostSuccess
	{
		username := "username"
		postID := "postID"
		serv.EXPECT().DeletePost(ctx, username, postID).Return(nil)

		errDeletePost := u.DeletePost(ctx, username, postID)
		require.Nilf(t, errDeletePost, "[usecase.test.DeletePostSuccess]:  u.DeletePost(ctx, username, postID) failed: %v", errDeletePost)

	}
}

func TestDeletePostFailure(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	serv := mock.NewMockIService(ctrl)
	genID := mock.NewMockIGeneratorID(ctrl)
	u := NewUseCase(serv, genID)

	// ErrServDeletePost
	{
		username := "username"
		postID := "postID"
		serv.EXPECT().DeletePost(ctx, username, postID).Return(errors.New("serv.DeletePost failed"))

		errDeletePost := u.DeletePost(ctx, username, postID)
		require.NotNil(t, errDeletePost, "[usecase.test.DeletePostFailure(ErrServDeletePost)]: expected error from u.DeletePost(ctx, username, postID)")

	}
}

func TestUpvotePostSuccess(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	serv := mock.NewMockIService(ctrl)
	genID := mock.NewMockIGeneratorID(ctrl)
	u := NewUseCase(serv, genID)

	// UpvotePostSuccess
	{
		userID := "userID"
		postID := "postID"
		postReturn := entity.PostExtend{
			Post: entity.Post{
				Votes: []entity.Vote{
					{
						UserID: userID,
						Vote:   1,
					},
				},
			},
			ID: postID,
		}
		serv.EXPECT().UpvotePost(ctx, userID, postID).Return(postReturn, nil)

		post, errUpvotePost := u.Upvote(ctx, userID, postID)
		require.Nilf(t, errUpvotePost, "[usecase.test.UpvotePostSuccess]: u.Upvote(ctx, userID, postID) failed: %v", errUpvotePost)

		if !reflect.DeepEqual(postReturn, post) {
			t.Errorf("[usecase.test.UpvotePostSuccess]: results not match, want %v, have %v", postReturn, post)
			return
		}

	}
}

func TestUpvotePostFailure(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	serv := mock.NewMockIService(ctrl)
	genID := mock.NewMockIGeneratorID(ctrl)
	u := NewUseCase(serv, genID)

	// ErrServUpvotePost
	{
		userID := "userID"
		postID := "postID"
		postReturn := entity.PostExtend{}
		serv.EXPECT().UpvotePost(ctx, userID, postID).Return(postReturn, errors.New("serv.UpvotePost failed"))

		post, errUpvotePost := u.Upvote(ctx, userID, postID)
		require.NotNil(t, errUpvotePost, "[usecase.test.UpvotePostFailure(ErrServUpvotePost)]: expected error from u.Upvote(ctx, userID, postID)")

		if !reflect.DeepEqual(postReturn, post) {
			t.Errorf("[usecase.test.UpvotePostFailure(ErrServUpvotePost)]: results not match, want %v, have %v", postReturn, post)
			return
		}

	}
}

func TestDownvotePostSuccess(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	serv := mock.NewMockIService(ctrl)
	genID := mock.NewMockIGeneratorID(ctrl)
	u := NewUseCase(serv, genID)

	// UpvotePostSuccess
	{
		userID := "userID"
		postID := "postID"
		postReturn := entity.PostExtend{
			ID: postID,
		}
		serv.EXPECT().DownvotePost(ctx, userID, postID).Return(postReturn, nil)

		post, errDownvotePost := u.Downvote(ctx, userID, postID)
		require.Nilf(t, errDownvotePost, "[usecase.test.DownvotePostSuccess]: u.Downvote(ctx, userID, postID) failed: %v", errDownvotePost)

		if !reflect.DeepEqual(postReturn, post) {
			t.Errorf("[usecase.test.DownvotePostSuccess]: results not match, want %v, have %v", postReturn, post)
			return
		}

	}
}

func TestDownvotePostFailure(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	serv := mock.NewMockIService(ctrl)
	genID := mock.NewMockIGeneratorID(ctrl)
	u := NewUseCase(serv, genID)

	// ErrServDownovtePost
	{
		userID := "userID"
		postID := "postID"
		postReturn := entity.PostExtend{}
		serv.EXPECT().DownvotePost(ctx, userID, postID).Return(postReturn, errors.New("serv.Downvote failed"))

		post, errDownvotePost := u.Downvote(ctx, userID, postID)
		require.NotNil(t, errDownvotePost, "[usecase.test.DownvotePostFailure(ErrServDownovtePost)]: exp u.Upvote(ctx, userID, postID) failed: %v", errDownvotePost)

		if !reflect.DeepEqual(postReturn, post) {
			t.Errorf("[usecase.test.DownvotePostFailure(ErrServDownovtePost)]: results not match, want %v, have %v", postReturn, post)
			return
		}

	}
}

func TestUnvotePostSuccess(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	serv := mock.NewMockIService(ctrl)
	genID := mock.NewMockIGeneratorID(ctrl)
	u := NewUseCase(serv, genID)

	// UnvotePostSuccess
	{
		userID := "userID"
		postID := "postID"
		postReturn := entity.PostExtend{
			ID: postID,
		}
		serv.EXPECT().UnvotePost(ctx, userID, postID).Return(postReturn, nil)

		post, errUnvotePost := u.Unvote(ctx, userID, postID)
		require.Nilf(t, errUnvotePost, "[usecase.test.UnvotePostSuccess]: u.Unvote(ctx, userID, postID) failed: %v", errUnvotePost)

		if !reflect.DeepEqual(postReturn, post) {
			t.Errorf("[usecase.test.UnvotePostSuccess]: results not match, want %v, have %v", postReturn, post)
			return
		}

	}
}

func TestUnvotePostFailure(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	serv := mock.NewMockIService(ctrl)
	genID := mock.NewMockIGeneratorID(ctrl)
	u := NewUseCase(serv, genID)

	// ErrServUnvotePost
	{
		userID := "userID"
		postID := "postID"
		postReturn := entity.PostExtend{}
		serv.EXPECT().UnvotePost(ctx, userID, postID).Return(postReturn, errors.New("serv.UnvotePost failed"))

		post, errUnvotePost := u.Unvote(ctx, userID, postID)
		require.NotNil(t, errUnvotePost, "[usecase.test.UnvotePostFailure(ErrServUnvotePost)]: expected error from u.Unvote(ctx, userID, postID)")

		if !reflect.DeepEqual(postReturn, post) {
			t.Errorf("[usecase.test.UnvotePostFailure(ErrServUnvotePost)]: results not match, want %v, have %v", postReturn, post)
			return
		}

	}
}
