package usecase

import (
	"context"
	"fmt"

	"gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/internal/entity"
	"gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/internal/interfaces"
)

func (u *usecase) SignUp(ctx context.Context, user entity.User) (entity.UserExtend, error) {
	u.mu.Lock()
	defer u.mu.Unlock()

	exists, errContains := u.service.ContainsUser(ctx, user.Username)
	if errContains != nil {
		return entity.UserExtend{}, fmt.Errorf("[usecase.SignUp]->%w", errContains)
	}
	if exists {
		return entity.UserExtend{}, fmt.Errorf("[usecase.SignUp]: %w", interfaces.ErrUserExists)
	}

	id, errGenID := u.genID.Generate(ctx)
	if errGenID != nil {
		return entity.UserExtend{}, fmt.Errorf("[usecase.SignUp]: %w", interfaces.ErrConstructObject)
	}

	userExtend := entity.NewUserExtend(user, id)

	errAddUser := u.service.AddUser(ctx, *userExtend)
	if errAddUser != nil {
		return entity.UserExtend{}, fmt.Errorf("[usecase.SignUp]: %w", errAddUser)
	}

	return *userExtend, nil
}

func (u *usecase) Login(ctx context.Context, username string, password string) (entity.UserExtend, error) {
	user, err := u.service.GetUser(ctx, username)
	if err != nil {
		return entity.UserExtend{}, fmt.Errorf("[usecase.Login]->%w", err)
	}

	if user.Password != password {
		return entity.UserExtend{}, fmt.Errorf("[usecase.Login]: %w", interfaces.ErrInvalidPassword)
	}

	return user, nil

}
