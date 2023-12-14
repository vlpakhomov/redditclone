package service

import (
	"context"
	"fmt"

	"gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/internal/entity"
	"gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/internal/interfaces"
)

func (s *service) AddComment(ctx context.Context, postID string, comment entity.CommentExtend) (entity.PostExtend, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	post, errAddComment := s.postsRepo.AddComment(ctx, postID, comment)
	if errAddComment != nil {
		return entity.PostExtend{}, fmt.Errorf("[service.AddComment]->%w", errAddComment)
	}
	return post, nil
}

func (s *service) DeleteComment(ctx context.Context, username string, postID string, commentID string) (entity.PostExtend, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	comment, errGetComment := s.postsRepo.GetComment(ctx, postID, commentID)
	if errGetComment != nil {
		return entity.PostExtend{}, fmt.Errorf("[service.DeleteComment]->%w", errGetComment)
	}
	if comment.Author.Username != username {
		return entity.PostExtend{}, fmt.Errorf("[service.DeleteComment]: %w", interfaces.ErrAccessDenied)
	}

	post, errDeleteComment := s.postsRepo.DeleteComment(ctx, postID, commentID)
	if errDeleteComment != nil {
		return entity.PostExtend{}, fmt.Errorf("[service.DeleteComment]->%w", errDeleteComment)
	}
	return post, nil
}
