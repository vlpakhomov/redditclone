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

func TestAddPostSuccess(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	postRepo := mock.NewMockIPostRepository(ctrl)
	s := NewService(postRepo, nil)

	// AddPostSuccess
	{
		newPost := entity.PostExtend{}
		postRepo.EXPECT().Add(ctx, newPost).Return(nil)

		errAddPost := s.AddPost(ctx, newPost)
		require.Nilf(t, errAddPost, "[service.test.AddPostSuccess]: s.AddPost(ctx, newPost) failed: %v", errAddPost)
	}
}

func TestAddPostFailure(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	postRepo := mock.NewMockIPostRepository(ctrl)
	s := NewService(postRepo, nil)

	// ErrRepoAdd
	{
		newPost := entity.PostExtend{}
		postRepo.EXPECT().Add(ctx, newPost).Return(errors.New("postRepo.Add failed"))

		errAddPost := s.AddPost(ctx, newPost)
		require.NotNil(t, errAddPost, "[service.test.AddPostFailure(ErrRepoAdd)]: expected error from s.AddPost(ctx, newPost)")
	}
}

func TestGetPostSuccess(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	postRepo := mock.NewMockIPostRepository(ctrl)
	s := NewService(postRepo, nil)

	// GetPostSuccess
	{
		postID := "postID"
		postReturn := entity.PostExtend{
			ID: "id",
		}
		postRepo.EXPECT().Get(ctx, postID).Return(postReturn, nil)

		post, errGetPost := s.GetPost(ctx, postID)
		require.Nilf(t, errGetPost, "[service.test.GetPostSuccess]: s.GetPost(ctx, postID) failed: %v", errGetPost)

		if !reflect.DeepEqual(postReturn, post) {
			t.Errorf("[service.test.GetPostSuccess]: results not match, want %v, have %v", postReturn, post)
			return
		}
	}
}

func TestGetPostFailure(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	postRepo := mock.NewMockIPostRepository(ctrl)
	s := NewService(postRepo, nil)

	// ErrRepoGet
	{
		postID := "postID"
		postReturn := entity.PostExtend{}
		postRepo.EXPECT().Get(ctx, postID).Return(postReturn, errors.New("post.Repo.Get failed"))

		post, errGetPost := s.GetPost(ctx, postID)
		require.NotNil(t, errGetPost, "[service.test.GetPostFailure(ErrRepoGet)]: s.GetPosts(ctx) failed: %v", errGetPost)

		if !reflect.DeepEqual(postReturn, post) {
			t.Errorf("[service.test.GetPostFailure(ErrRepoGet)]: results not match, want %v, have %v", postReturn, post)
			return
		}
	}
}

func TestGetPostsSuccess(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	postRepo := mock.NewMockIPostRepository(ctrl)
	s := NewService(postRepo, nil)

	// GetPostsSuccess
	{
		postsReturn := []entity.PostExtend{{
			ID: "id",
		}}
		postRepo.EXPECT().GetAll(ctx).Return(postsReturn, nil)

		posts, errGetPosts := s.GetPosts(ctx)
		require.Nilf(t, errGetPosts, "[service.test.GetPostsSuccess]: s.GetPosts(ctx) failed: %v", errGetPosts)

		if !reflect.DeepEqual(postsReturn, posts) {
			t.Errorf("[service.test.GetPostsSuccess]: results not match, want %v, have %v", postsReturn, posts)
			return
		}
	}
}

func TestGetPostsFailure(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	postRepo := mock.NewMockIPostRepository(ctrl)
	s := NewService(postRepo, nil)

	// ErrRepoGetAll
	{
		postsReturn := []entity.PostExtend{}
		postRepo.EXPECT().GetAll(ctx).Return(postsReturn, errors.New("postRepo.GetAll failed"))

		posts, errGetPosts := s.GetPosts(ctx)
		require.NotNil(t, errGetPosts, "[service.test.GetPostsFailure(ErrRepoGetAll)]: expected error from s.GetPosts(ctx)")

		if !reflect.DeepEqual(postsReturn, posts) {
			t.Errorf("[service.test.GetPostsFailure(ErrRepoGetAll)]: results not match, want %v, have %v", postsReturn, posts)
			return
		}
	}
}

func TestGetPostsWithCategorySuccess(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	postRepo := mock.NewMockIPostRepository(ctrl)
	s := NewService(postRepo, nil)

	// GetPostsWithCategorySuccess
	{
		category := "category"
		postsReturn := []entity.PostExtend{{
			Post: entity.Post{
				Category: category,
			},
		}}
		postRepo.EXPECT().GetWhereCategory(ctx, category).Return(postsReturn, nil)

		posts, errGetPostsWithCategory := s.GetPostsWithCategory(ctx, category)
		require.Nilf(t, errGetPostsWithCategory, "[service.test.GetPostsWithCategorySuccess]: s.GetPostsWithCategory(ctx, category) failed: %v", errGetPostsWithCategory)

		if !reflect.DeepEqual(postsReturn, posts) {
			t.Errorf("[service.test.GetPostsWithCategorySuccess]: results not match, want %v, have %v", postsReturn, posts)
			return
		}
	}
}

func TestGetPostsWithCategoryFailure(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	postRepo := mock.NewMockIPostRepository(ctrl)
	s := NewService(postRepo, nil)

	// ErrRepoGetWhereCategory
	{
		category := "category"
		postsReturn := []entity.PostExtend{}
		postRepo.EXPECT().GetWhereCategory(ctx, category).Return(postsReturn, errors.New("postRepo.GetWhereCategory failed"))

		posts, errGetPostsWithCategory := s.GetPostsWithCategory(ctx, category)
		require.NotNil(t, errGetPostsWithCategory, "[service.test.GetPostsWithCategoryFailure(ErrRepoGetWhereCategory)]: expected error from s.GetPostsWithCategory(ctx, category)")

		if !reflect.DeepEqual(postsReturn, posts) {
			t.Errorf("[service.test.GetPostsWithCategoryFailure(ErrRepoGetWhereCategory)]: results not match, want %v, have %v", postsReturn, posts)
			return
		}
	}
}

func TestGetPostsWithUserSuccess(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	postRepo := mock.NewMockIPostRepository(ctrl)
	s := NewService(postRepo, nil)

	// GetPostsWithUserSuccess
	{
		username := "username"
		postsReturn := []entity.PostExtend{{
			Post: entity.Post{
				Author: entity.Author{
					Username: username,
				},
			}},
		}
		postRepo.EXPECT().GetWhereUsername(ctx, username).Return(postsReturn, nil)

		posts, errGetPostsWithUser := s.GetPostsWithUser(ctx, username)
		require.Nilf(t, errGetPostsWithUser, "[service.test.GetPostsWithUserSuccess]: s.GetPostsWithUser(ctx, username) failed: %v", errGetPostsWithUser)

		if !reflect.DeepEqual(postsReturn, posts) {
			t.Errorf("[service.test.GetPostsWithUserSuccess]: results not match, want %v, have %v", postsReturn, posts)
			return
		}
	}
}

func TestGetPostsWithUserFailure(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	postRepo := mock.NewMockIPostRepository(ctrl)
	s := NewService(postRepo, nil)

	// GetPostsWithUserFailure
	{
		username := "username"
		postsReturn := []entity.PostExtend{}
		postRepo.EXPECT().GetWhereUsername(ctx, username).Return(postsReturn, errors.New("postRepo.GetWhereUsername failed"))

		posts, errGetPostsWithUser := s.GetPostsWithUser(ctx, username)
		require.NotNil(t, errGetPostsWithUser, "[service.test.GetPostsWithUserFailure]: expected error from s.GetPostsWithUser(ctx, username)")

		if !reflect.DeepEqual(postsReturn, posts) {
			t.Errorf("[service.test.GetPostsWithUserFailure]: results not match, want %v, have %v", postsReturn, posts)
			return
		}
	}
}

func TestUpvotePostSuccess(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	postRepo := mock.NewMockIPostRepository(ctrl)
	s := NewService(postRepo, nil)

	// UpvotePostSuccess
	{
		userID := "userID"
		postID := "postID"
		postReturn := entity.PostExtend{
			ID: "id",
		}
		postRepo.EXPECT().Get(ctx, postID).Return(postReturn, nil)

		newPost := entity.PostExtend{
			ID: "id",
		}
		errUpvote := newPost.Upvote(userID)
		require.Nilf(t, errUpvote, "[service.test.UpvotePostSuccess]: newPost.Upvote(userID) failed: %v", errUpvote)

		postRepo.EXPECT().Update(ctx, postID, newPost).Return(nil)

		post, errUpvotePost := s.UpvotePost(ctx, userID, postID)
		require.Nilf(t, errUpvote, "[service.test.UpvotePostSuccess]: s.UpvotePost(ctx, userID, postID) failed: %v", errUpvotePost)

		if !reflect.DeepEqual(newPost, post) {
			t.Errorf("[service.test.UpvotePostSuccess]: results not match, want %v, have %v", newPost, post)
			return
		}
	}
}

func TestUpvotePostFailure(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	postRepo := mock.NewMockIPostRepository(ctrl)
	s := NewService(postRepo, nil)

	// ErrRepoGet
	{
		userID := "userID"
		postID := "postID"
		postReturn := entity.PostExtend{}
		postRepo.EXPECT().Get(ctx, postID).Return(postReturn, errors.New("postRepo.Get failed"))

		post, errUpvotePost := s.UpvotePost(ctx, userID, postID)
		require.NotNil(t, errUpvotePost, "[service.test.UpvotePostFailure(ErrRepoGet)]: expect error from s.UpvotePost(ctx, userID, postID)")

		if !reflect.DeepEqual(postReturn, post) {
			t.Errorf("[service.test.UpvotePostFailure(ErrRepoGet)]: results not match, want %v, have %v", postReturn, post)
			return
		}
	}

	// ErrUpvote
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
			ID: "id",
		}
		postRepo.EXPECT().Get(ctx, postID).Return(postReturn, nil)

		newPost := entity.PostExtend{
			Post: entity.Post{
				Votes: []entity.Vote{
					{
						UserID: userID,
						Vote:   1,
					},
				},
			},
			ID: "id",
		}
		errUpvote := newPost.Upvote(userID)
		require.NotNil(t, errUpvote, "[service.test.UpvotePostFailure(ErrUpvote)]: expected error newPost.Upvote(userID)")

		post, errUpvotePost := s.UpvotePost(ctx, userID, postID)
		require.NotNil(t, errUpvotePost, "[service.test.UpvotePostFailure(ErrUpvote)]: expect error from s.UpvotePost(ctx, userID, postID)")

		postReturn = entity.PostExtend{}
		if !reflect.DeepEqual(postReturn, post) {
			t.Errorf("[service.test.UpvotePostFailure(ErrUpvote)]: results not match, want %v, have %v", postReturn, post)
			return
		}
	}

	// ErrRepoUpdate
	{
		userID := "userID"
		postID := "postID"
		postReturn := entity.PostExtend{
			ID: "id",
		}
		postRepo.EXPECT().Get(ctx, postID).Return(postReturn, nil)

		newPost := entity.PostExtend{
			ID: "id",
		}
		errUpvote := newPost.Upvote(userID)
		require.Nilf(t, errUpvote, "[service.test.UpvotePostFailure(ErrRepoUpdate)]: newPost.Upvote(userID) failed: %v", errUpvote)

		postRepo.EXPECT().Update(ctx, postID, newPost).Return(errors.New("postRepo.Update failed"))

		post, errUpvotePost := s.UpvotePost(ctx, userID, postID)
		require.NotNil(t, errUpvotePost, "[service.test.UpvotePostFailure(ErrRepoUpdate)]: expect error from s.UpvotePost(ctx, userID, postID)")

		postReturn = entity.PostExtend{}
		if !reflect.DeepEqual(postReturn, post) {
			t.Errorf("[service.test.UpvotePostFailure(ErrRepoUpdate)]: results not match, want %v, have %v", postReturn, post)
			return
		}
	}
}

func TestDownvotePostSucess(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	postRepo := mock.NewMockIPostRepository(ctrl)
	s := NewService(postRepo, nil)

	// DownvotePostSuccess
	{
		userID := "userID"
		postID := "postID"
		postReturn := entity.PostExtend{
			ID: "id",
		}
		postRepo.EXPECT().Get(ctx, postID).Return(postReturn, nil)

		newPost := entity.PostExtend{
			ID: "id",
		}
		errUpvote := newPost.Downvote(userID)
		require.Nilf(t, errUpvote, "[service.test.DownvotePostSuccess]: newPost.Downvote(userID) failed: %v", errUpvote)

		postRepo.EXPECT().Update(ctx, postID, newPost).Return(nil)

		post, errUpvotePost := s.DownvotePost(ctx, userID, postID)
		require.Nilf(t, errUpvote, "[service.test.DownvotePostSuccess]: s.DownvotePost(ctx, userID, postID) failed: %v", errUpvotePost)

		if !reflect.DeepEqual(newPost, post) {
			t.Errorf("[service.test.DownvotePostSuccess]: results not match, want %v, have %v", newPost, post)
			return
		}
	}
}

func TestDownvotePostFailure(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	postRepo := mock.NewMockIPostRepository(ctrl)
	s := NewService(postRepo, nil)

	// ErrRepoGet
	{
		userID := "userID"
		postID := "postID"
		postReturn := entity.PostExtend{}
		postRepo.EXPECT().Get(ctx, postID).Return(postReturn, errors.New("postRepo.Get failed"))

		post, errDownvotePost := s.DownvotePost(ctx, userID, postID)
		require.NotNil(t, errDownvotePost, "[service.test.DownvotePostFailure(ErrRepoGet)]: expect error from s.DownvotePost(ctx, userID, postID)")

		if !reflect.DeepEqual(postReturn, post) {
			t.Errorf("[service.test.DownvotePostFailure(ErrRepoGet)]: results not match, want %v, have %v", postReturn, post)
			return
		}
	}

	// ErrDownvote
	{
		userID := "userID"
		postID := "postID"
		postReturn := entity.PostExtend{
			Post: entity.Post{
				Votes: []entity.Vote{
					{
						UserID: userID,
						Vote:   -1,
					},
				},
			},
			ID: "id",
		}
		postRepo.EXPECT().Get(ctx, postID).Return(postReturn, nil)

		newPost := entity.PostExtend{
			Post: entity.Post{
				Votes: []entity.Vote{
					{
						UserID: userID,
						Vote:   -1,
					},
				},
			},
			ID: "id",
		}
		errDownvote := newPost.Downvote(userID)
		require.NotNil(t, errDownvote, "[service.test.DownvotePostFailure(ErrDownvote)]: expected error newPost.Downvote(userID)")

		post, errDownvotePost := s.DownvotePost(ctx, userID, postID)
		require.NotNil(t, errDownvotePost, "[service.test.DownvotePostFailure(ErrDownvote)]: expect error from s.DownvotePost(ctx, userID, postID)")

		postReturn = entity.PostExtend{}
		if !reflect.DeepEqual(postReturn, post) {
			t.Errorf("[service.test.DownvotePostFailure(ErrDownvote)]: results not match, want %v, have %v", postReturn, post)
			return
		}
	}

	// ErrRepoUpdate
	{
		userID := "userID"
		postID := "postID"
		postReturn := entity.PostExtend{
			ID: "id",
		}
		postRepo.EXPECT().Get(ctx, postID).Return(postReturn, nil)

		newPost := entity.PostExtend{
			ID: "id",
		}
		errDownvote := newPost.Downvote(userID)
		require.Nilf(t, errDownvote, "[service.test.DownvotePostFailure(ErrRepoUpdate)]: newPost.Downvote(userID) failed: %v", errDownvote)

		postRepo.EXPECT().Update(ctx, postID, newPost).Return(errors.New("postRepo.Update failed"))

		post, errDownvotePost := s.DownvotePost(ctx, userID, postID)
		require.NotNil(t, errDownvotePost, "[service.test.DownvotePostFailure(ErrRepoUpdate)]: expect error from s.DownvotePost(ctx, userID, postID)")

		postReturn = entity.PostExtend{}
		if !reflect.DeepEqual(postReturn, post) {
			t.Errorf("[service.test.DownvotePostFailure(ErrRepoUpdate)]: results not match, want %v, have %v", postReturn, post)
			return
		}
	}
}

func TestUnvotePostSuccess(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	postRepo := mock.NewMockIPostRepository(ctrl)
	s := NewService(postRepo, nil)

	// UnvotePostSuccess
	{
		userID := "userID"
		postID := "postID"
		postReturn := entity.PostExtend{
			Post: entity.Post{
				Score: 1,
				Votes: []entity.Vote{
					{
						UserID: userID,
						Vote:   1,
					},
				},
				UpvotePercentage: 100,
			},
			ID: postID,
		}
		postRepo.EXPECT().Get(ctx, postID).Return(postReturn, nil)

		newPost := entity.PostExtend{
			Post: entity.Post{
				Score: 1,
				Votes: []entity.Vote{
					{
						UserID: userID,
						Vote:   1,
					},
				},
				UpvotePercentage: 100,
			},
			ID: postID,
		}
		errUnvote := newPost.Unvote(userID)
		require.Nilf(t, errUnvote, "[service.test.UnvotePostSucess]: newPost.Unvote(userID) failed: %v", errUnvote)

		postRepo.EXPECT().Update(ctx, postID, newPost).Return(nil)

		post, errUnvotePost := s.UnvotePost(ctx, userID, postID)
		require.Nilf(t, errUnvotePost, "[service.test.UnvotePostSucess]: s.UnvotePost(ctx, userID, postID) failed: %v", errUnvotePost)

		if !reflect.DeepEqual(newPost, post) {
			t.Errorf("[service.test.UnvotePostSucess]: results not match, want %v, have %v", newPost, post)
			return
		}
	}
}

func TestUnvotePostFailure(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	postRepo := mock.NewMockIPostRepository(ctrl)
	s := NewService(postRepo, nil)

	// ErrRepoGet
	{
		userID := "userID"
		postID := "postID"
		postReturn := entity.PostExtend{}
		postRepo.EXPECT().Get(ctx, postID).Return(postReturn, errors.New("postRepo.Get failed"))

		post, errUnvotePost := s.UnvotePost(ctx, userID, postID)
		require.NotNil(t, errUnvotePost, "[service.test.UnvotePostFailure(ErrRepoGet)]: expect error from s.DownvotePost(ctx, userID, postID)")

		if !reflect.DeepEqual(postReturn, post) {
			t.Errorf("[service.test.UnvotePostFailure(ErrRepoGet)]: results not match, want %v, have %v", postReturn, post)
			return
		}
	}

	// ErrUnvote
	{
		userID := "userID"
		postID := "postID"
		postReturn := entity.PostExtend{
			ID: "id",
		}
		postRepo.EXPECT().Get(ctx, postID).Return(postReturn, nil)

		newPost := postReturn
		errUnvote := newPost.Unvote(userID)
		require.NotNil(t, errUnvote, "[service.test.UnvotePostFailure(ErrUnvote)]: expected error newPost.Unvote(userID)")

		post, errUnvotePost := s.UnvotePost(ctx, userID, postID)
		require.NotNil(t, errUnvotePost, "[service.test.UnvotePostFailure(ErrUnvote)]: expect error from s.UnvotePost(ctx, userID, postID)")

		postReturn = entity.PostExtend{}
		if !reflect.DeepEqual(postReturn, post) {
			t.Errorf("[service.test.UnvotePostFailure(ErrUnvote)]: results not match, want %v, have %v", postReturn, post)
			return
		}
	}

	// ErrRepoUpdate
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
			ID: "id",
		}
		postRepo.EXPECT().Get(ctx, postID).Return(postReturn, nil)

		newPost := entity.PostExtend{
			Post: entity.Post{
				Votes: []entity.Vote{
					{
						UserID: userID,
						Vote:   1,
					},
				},
			},
			ID: "id",
		}
		errUnvote := newPost.Unvote(userID)
		require.Nilf(t, errUnvote, "[service.test.UnvotePostFailure(ErrRepoUpdate)]: newPost.Unvote(userID) failed: %v", errUnvote)

		postRepo.EXPECT().Update(ctx, postID, newPost).Return(errors.New("postRepo.Update failed"))

		post, errUnvotePost := s.UnvotePost(ctx, userID, postID)
		require.NotNil(t, errUnvotePost, "[service.test.UnvotePostFailure(ErrRepoUpdate)]: expect error from s.UnvotePost(ctx, userID, postID)")

		postReturn = entity.PostExtend{}
		if !reflect.DeepEqual(postReturn, post) {
			t.Errorf("[service.test.UnvotePostFailure(ErrRepoUpdate)]: results not match, want %v, have %v", postReturn, post)
			return
		}
	}
}

func TestDeletePostSucess(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	postRepo := mock.NewMockIPostRepository(ctrl)
	s := NewService(postRepo, nil)

	// DeletePostSucces
	{
		username := "username"
		postID := "postID"
		postReturn := entity.PostExtend{
			Post: entity.Post{
				Author: entity.Author{
					Username: username,
				},
			},
		}
		postRepo.EXPECT().Get(ctx, postID).Return(postReturn, nil)

		postRepo.EXPECT().Delete(ctx, postID).Return(nil)

		errDeletePost := s.DeletePost(ctx, username, postID)
		require.Nilf(t, errDeletePost, "[service.test.DeletePostSucess]: s.DeletePost(ctx, username, postID) failed: %v", errDeletePost)
	}
}

func TestDeletePostFailure(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	postRepo := mock.NewMockIPostRepository(ctrl)
	s := NewService(postRepo, nil)

	// ErrRepoGet
	{
		username := "username"
		postID := "postID"
		postReturn := entity.PostExtend{}
		postRepo.EXPECT().Get(ctx, postID).Return(postReturn, errors.New("postRepo.Get failed"))

		errDeletePost := s.DeletePost(ctx, username, postID)
		require.NotNil(t, errDeletePost, "[service.test.DeletePostFailure(ErrRepoGet)]: expect error from s.DeletePost(ctx, username, postID)")
	}

	// ErrAccessDenied
	{
		username := "username"
		postID := "postID"
		postReturn := entity.PostExtend{
			Post: entity.Post{
				Author: entity.Author{
					Username: username + ".",
				},
			},
		}
		postRepo.EXPECT().Get(ctx, postID).Return(postReturn, nil)

		errDeletePost := s.DeletePost(ctx, username, postID)
		require.NotNil(t, errDeletePost, "[service.test.DeletePostFailure(ErrAccessDenied)]: expected error from s.DeletePost(ctx, username, postID)")

		require.Truef(t, errors.Is(errDeletePost, interfaces.ErrAccessDenied), "[service.test.DeletePostFailure]: error not match, want %v, have %v", interfaces.ErrAccessDenied, errDeletePost)
	}

	// ErrDelete
	{
		username := "username"
		postID := "postID"
		postReturn := entity.PostExtend{
			Post: entity.Post{
				Author: entity.Author{
					Username: username,
				},
			},
		}
		postRepo.EXPECT().Get(ctx, postID).Return(postReturn, nil)

		postRepo.EXPECT().Delete(ctx, postID).Return(errors.New("postRepo.Delete failed"))

		errDeletePost := s.DeletePost(ctx, username, postID)
		require.NotNil(t, errDeletePost, "[service.test.DeletePostFailure(ErrDelete)]: expected error from s.DeletePost(ctx, username, postID)")
	}
}
