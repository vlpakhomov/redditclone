package service

import (
	"context"
	"fmt"
	"sort"

	"gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/internal/entity"
	"gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/internal/interfaces"
)

func (s *service) AddPost(ctx context.Context, post entity.PostExtend) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	errAddPost := s.postsRepo.Add(ctx, post)
	if errAddPost != nil {
		return fmt.Errorf("[service.AddPost]->%w", errAddPost)
	}

	return nil
}

func (s *service) GetPost(ctx context.Context, postID string) (entity.PostExtend, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	post, err := s.postsRepo.Get(ctx, postID)
	if err != nil {
		return entity.PostExtend{}, fmt.Errorf("[service.GetPost]->%w", err)
	}

	return post, nil
}

func (s *service) GetPosts(ctx context.Context) ([]entity.PostExtend, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	posts, errGetAll := s.postsRepo.GetAll(ctx)
	if errGetAll != nil {
		return []entity.PostExtend{}, fmt.Errorf("[service.GetPosts]->%w", errGetAll)
	}

	return s.SortPostsByTime(posts), nil
}

func (s *service) GetPostsWithCategory(ctx context.Context, category string) ([]entity.PostExtend, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	posts, errGetPosts := s.postsRepo.GetWhereCategory(ctx, category)
	if errGetPosts != nil {
		return []entity.PostExtend{}, fmt.Errorf("[service.GetWhereCategory]->%w", errGetPosts)
	}
	return s.SortPostsByTime(posts), nil
}

func (s *service) GetPostsWithUser(ctx context.Context, username string) ([]entity.PostExtend, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	posts, errGetPosts := s.postsRepo.GetWhereUsername(ctx, username)
	if errGetPosts != nil {
		return []entity.PostExtend{}, fmt.Errorf("[service.GetPostsWithUser]->%w", errGetPosts)
	}

	return s.SortPostsByTime(posts), nil
}

func (s *service) UpvotePost(ctx context.Context, userID string, postID string) (entity.PostExtend, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	post, err := s.postsRepo.Get(ctx, postID)
	if err != nil {
		return entity.PostExtend{}, fmt.Errorf("[service.UnvotePost]->%w", err)
	}

	err = post.Upvote(userID)
	if err != nil {
		return entity.PostExtend{}, fmt.Errorf("[service.UnvotePost]->%w", err)
	}

	errUpdate := s.postsRepo.Update(ctx, postID, post)
	if errUpdate != nil {
		return entity.PostExtend{}, fmt.Errorf("[service.UnvotePost]->%w", errUpdate)
	}

	return post, nil
}

func (s *service) DownvotePost(ctx context.Context, userID string, postID string) (entity.PostExtend, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	post, errGet := s.postsRepo.Get(ctx, postID)
	if errGet != nil {
		return entity.PostExtend{}, fmt.Errorf("[service.DownvotePost]->%w", errGet)
	}

	errDownvote := post.Downvote(userID)
	if errDownvote != nil {
		return entity.PostExtend{}, fmt.Errorf("[service.DownvotePost]->%w", errDownvote)
	}

	errUpdate := s.postsRepo.Update(ctx, postID, post)
	if errUpdate != nil {
		return entity.PostExtend{}, fmt.Errorf("[service.DownvotePost]->%w", errUpdate)
	}

	return post, nil
}

func (s *service) UnvotePost(ctx context.Context, userID string, postID string) (entity.PostExtend, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	post, errGet := s.postsRepo.Get(ctx, postID)
	if errGet != nil {
		return entity.PostExtend{}, fmt.Errorf("[service.UnvotePost]->%w", errGet)
	}

	errUnvote := post.Unvote(userID)
	if errUnvote != nil {
		return entity.PostExtend{}, fmt.Errorf("[service.UnvotePost]->%w", errUnvote)
	}

	errUpdate := s.postsRepo.Update(ctx, postID, post)
	if errUpdate != nil {
		return entity.PostExtend{}, fmt.Errorf("[service.UnvotePost]->%w", errUpdate)
	}

	return post, nil
}

func (s *service) DeletePost(ctx context.Context, username string, postID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	post, errGet := s.postsRepo.Get(ctx, postID)
	if errGet != nil {
		return fmt.Errorf("[service.DeletePost]->%w", errGet)
	}

	if post.Author.Username != username {
		return fmt.Errorf("[service.DeletePost]: %w", interfaces.ErrAccessDenied)
	}

	errDelete := s.postsRepo.Delete(ctx, postID)
	if errDelete != nil {
		return fmt.Errorf("[service.DeletePost]->%w", errDelete)
	}

	return nil
}

func (s *service) SortPostsByTime(posts []entity.PostExtend) []entity.PostExtend {
	sort.SliceStable(posts, func(i, j int) bool {
		return posts[i].CreatedTime.Before(posts[j].CreatedTime)
	})
	return posts
}
