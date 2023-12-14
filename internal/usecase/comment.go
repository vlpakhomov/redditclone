package usecase

import (
	"context"
	"fmt"

	"gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/internal/entity"
	"gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/internal/interfaces"
)

func (u *usecase) AddComment(ctx context.Context, postID string, comment entity.Comment) (entity.PostExtend, error) {
	id, errGenID := u.genID.Generate(ctx)
	if errGenID != nil {
		return entity.PostExtend{}, fmt.Errorf("[usecase.AddComment] entity.NewCommentExtend(comment) failed: %w", interfaces.ErrConstructObject)
	}
	commentExtend := entity.NewCommentExtend(comment, id)

	post, errAddComment := u.service.AddComment(ctx, postID, *commentExtend)
	if errAddComment != nil {
		return entity.PostExtend{}, fmt.Errorf("[usecase.AddComment]->%w", errAddComment)
	}

	return post, nil
}

func (u *usecase) DeleteComment(ctx context.Context, username string, postID string, commentID string) (entity.PostExtend, error) {
	post, err := u.service.DeleteComment(ctx, username, postID, commentID)
	if err != nil {
		return entity.PostExtend{}, fmt.Errorf("[usecase.DeleteComment]->%w", err)
	}

	return post, nil
}
