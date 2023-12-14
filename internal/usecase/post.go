package usecase

import (
	"context"
	"fmt"

	"gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/internal/entity"
	"gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/internal/interfaces"
)

func (u *usecase) GetPosts(ctx context.Context) ([]entity.PostExtend, error) {
	posts, errGetPosts := u.service.GetPosts(ctx)
	if errGetPosts != nil {
		return []entity.PostExtend{}, fmt.Errorf("[usecase.GetPosts]->%w", errGetPosts)
	}

	return posts, nil
}

func (u *usecase) AddPost(ctx context.Context, post entity.Post) (entity.PostExtend, error) {
	id, errGenID := u.genID.Generate(ctx)
	if errGenID != nil {
		return entity.PostExtend{}, fmt.Errorf("[usecase.AddPost]: %w", interfaces.ErrConstructObject)
	}
	postExtend := entity.NewPostExtend(post, id)

	errAddPost := u.service.AddPost(ctx, *postExtend)
	if errAddPost != nil {
		return entity.PostExtend{}, fmt.Errorf("[usecase.AddPost]->%w", errAddPost)
	}

	return *postExtend, nil
}

func (u *usecase) GetPostsWithCategory(ctx context.Context, category string) ([]entity.PostExtend, error) {
	posts, errGetPosts := u.service.GetPostsWithCategory(ctx, category)
	if errGetPosts != nil {
		return []entity.PostExtend{}, fmt.Errorf("[usecase.GetPostsWithCategory]->%w", errGetPosts)
	}

	return posts, nil
}

func (u *usecase) GetPostsWithUser(ctx context.Context, username string) ([]entity.PostExtend, error) {
	posts, errGetPosts := u.service.GetPostsWithUser(ctx, username)
	if errGetPosts != nil {
		return []entity.PostExtend{}, fmt.Errorf("[usecase.GetPostsWithUser]->%w", errGetPosts)
	}

	return posts, nil
}

func (u *usecase) GetPost(ctx context.Context, postID string) (entity.PostExtend, error) {
	post, err := u.service.GetPost(ctx, postID)
	if err != nil {
		return entity.PostExtend{}, fmt.Errorf("[usecase.GetPost]->%w", err)
	}

	return post, nil
}

func (u *usecase) DeletePost(ctx context.Context, username string, postID string) error {
	err := u.service.DeletePost(ctx, username, postID)
	if err != nil {
		return fmt.Errorf("[usecase.DeletePost]->%w", err)
	}

	return nil
}

func (u *usecase) Upvote(ctx context.Context, userID string, postID string) (entity.PostExtend, error) {
	post, err := u.service.UpvotePost(ctx, userID, postID)
	if err != nil {
		return entity.PostExtend{}, fmt.Errorf("[usecase.Upvote]->%w", err)
	}

	return post, nil
}

func (u *usecase) Downvote(ctx context.Context, userID string, postID string) (entity.PostExtend, error) {
	post, err := u.service.DownvotePost(ctx, userID, postID)
	if err != nil {
		return entity.PostExtend{}, fmt.Errorf("[usecase.Downvote]->%w", err)
	}

	return post, nil
}

func (u *usecase) Unvote(ctx context.Context, userID string, postID string) (entity.PostExtend, error) {
	post, err := u.service.UnvotePost(ctx, userID, postID)
	if err != nil {
		return entity.PostExtend{}, fmt.Errorf("[usecase.Unvote]->%w", err)
	}

	return post, nil
}
