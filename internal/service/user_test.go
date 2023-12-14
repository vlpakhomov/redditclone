package service

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

func TestAddUserSuccess(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mock.NewMockIUserRepository(ctrl)
	s := NewService(nil, userRepo)

	// AddUserSuccess
	{
		user := entity.UserExtend{}
		userRepo.EXPECT().Add(ctx, user).Return(nil)

		errAddUser := s.AddUser(ctx, user)
		require.Nilf(t, errAddUser, "[service.test.AddUserSuccess]: s.AddUser(ctx, user) failed: %v", errAddUser)
	}
}

func TestAddUserFailure(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mock.NewMockIUserRepository(ctrl)
	s := NewService(nil, userRepo)

	// ErrAddUser
	{
		user := entity.UserExtend{
			User: entity.User{
				Username: "username",
				Password: "password",
			},
			ID: "id",
		}
		userRepo.EXPECT().Add(ctx, user).Return(errors.New("userRepo.Add failed"))

		errAddUser := s.AddUser(ctx, user)
		require.NotNil(t, errAddUser, "[service.test.AddUserFailure(ErrAddUser)]: expected error from s.AddUser(ctx, user)")
	}
}

func TestGetUserSuccess(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mock.NewMockIUserRepository(ctrl)
	s := NewService(nil, userRepo)

	// GetUserSuccess
	{
		username := "username"
		userReturn := entity.UserExtend{
			ID: "id",
		}
		userRepo.EXPECT().Get(ctx, username).Return(userReturn, nil)

		user, errGetUser := s.GetUser(ctx, username)
		require.Nilf(t, errGetUser, "[service.test.GetUserSuccess]: s.GetUser(ctx, username) failed: %v", errGetUser)

		if !reflect.DeepEqual(userReturn, user) {
			t.Errorf("[service.test.GetCommentSuccess]: results not match, want %v, have %v", userReturn, user)
			return
		}
	}
}

func TestGetUserFailure(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mock.NewMockIUserRepository(ctrl)
	s := NewService(nil, userRepo)

	// ErrRepoGet
	{
		username := "username"
		userReturn := entity.UserExtend{}
		userRepo.EXPECT().Get(ctx, username).Return(userReturn, errors.New("userRepo.Get failed"))

		user, errGetUser := s.GetUser(ctx, username)
		require.NotNil(t, errGetUser, "[service.test.GetUserFailure(ErrRepoGet)]: expected error from s.GetUser(ctx, username)")

		if !reflect.DeepEqual(userReturn, user) {
			t.Errorf("[service.test.GetUserFailure(ErrRepoGet)]: results not match, want %v, have %v", userReturn, user)
			return
		}
	}
}

func TestContainsUserSuccess(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mock.NewMockIUserRepository(ctrl)
	s := NewService(nil, userRepo)

	// ContainsUserSuccess
	{
		username := "username"
		existsReturn := true
		userRepo.EXPECT().Contains(ctx, username).Return(existsReturn, nil)

		exists, errContainsUser := s.ContainsUser(ctx, username)
		require.Nilf(t, errContainsUser, "[service.test.ContainsUserSuccess(ContainsUserSuccess)]: s.ContainsUser(ctx, username) failed: %v", errContainsUser)

		if !reflect.DeepEqual(existsReturn, exists) {
			t.Errorf("[service.test.ContainsUserSuccess(ContainsUserSuccess)]: results not match, want %v, have %v", existsReturn, exists)
			return
		}
	}
}

func TestContainsUserFailure(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mock.NewMockIUserRepository(ctrl)
	s := NewService(nil, userRepo)

	// ErrRepoContains
	{
		username := "username"
		existsReturn := false
		userRepo.EXPECT().Contains(ctx, username).Return(existsReturn, errors.New("userRepo.Contains failed"))

		exists, errContainsUser := s.ContainsUser(ctx, username)
		require.NotNil(t, errContainsUser, "[service.test.ContainsUserFailure(ErrRepoContains)]: expected error from s.ContainsUser(ctx, username)")

		if !reflect.DeepEqual(existsReturn, exists) {
			t.Errorf("[service.test.ContainsUserFailure(ErrRepoContains)]: results not match, want %v, have %v", existsReturn, exists)
			return
		}
	}
}
