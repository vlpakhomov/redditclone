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

func TestSignUpSuccess(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	serv := mock.NewMockIService(ctrl)
	genID := mock.NewMockIGeneratorID(ctrl)
	u := NewUseCase(serv, genID)

	// SignUpSuccess
	{
		username := "username"
		existsReturn := false
		serv.EXPECT().ContainsUser(ctx, username).Return(existsReturn, nil)

		userIDReturn := "userID"
		genID.EXPECT().Generate(ctx).Return(userIDReturn, nil)

		user := entity.User{
			Username: username,
		}
		userExtend := entity.NewUserExtend(user, userIDReturn)

		serv.EXPECT().AddUser(ctx, *userExtend).Return(nil)

		userResult, errSignUp := u.SignUp(ctx, user)
		require.Nilf(t, errSignUp, "[usecase.test.SignUpSuccess]: u.SignUp(ctx, user) failed: %v", errSignUp)

		if !reflect.DeepEqual(*userExtend, userResult) {
			t.Errorf("[usecase.test.SignUpSuccess]: results not match, want %v, have %v", userExtend, userResult)
			return
		}
	}
}

func TestSignUpFailure(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	serv := mock.NewMockIService(ctrl)
	genID := mock.NewMockIGeneratorID(ctrl)
	u := NewUseCase(serv, genID)

	// ErrServContainsUser
	{
		username := "username"
		existsReturn := false
		serv.EXPECT().ContainsUser(ctx, username).Return(existsReturn, errors.New("serv.ContainsUser failed"))

		user := entity.User{
			Username: username,
		}
		userReturn := entity.UserExtend{}

		userResult, errSignUp := u.SignUp(ctx, user)
		require.NotNil(t, errSignUp, "[usecase.test.SignUpFailure(ErrServContainsUser)]: expected error from u.SignUp(ctx, user)")

		if !reflect.DeepEqual(userReturn, userResult) {
			t.Errorf("[usecase.test.SignUpFailure(ErrServContainsUser)]: results not match, want %v, have %v", userReturn, userResult)
			return
		}
	}

	// ErrUserExists
	{
		username := "username"
		existsReturn := true
		serv.EXPECT().ContainsUser(ctx, username).Return(existsReturn, nil)

		user := entity.User{
			Username: username,
		}

		userReturn := entity.UserExtend{}

		userResult, errSignUp := u.SignUp(ctx, user)
		require.NotNil(t, errSignUp, "[usecase.test.SignUpFailure(ErrUserExists)]: expected error from u.SignUp(ctx, user)", errSignUp)

		if !reflect.DeepEqual(userReturn, userResult) {
			t.Errorf("[usecase.test.SignUpFailure(ErrUserExists)]: results not match, want %v, have %v", userReturn, userResult)
			return
		}

	}

	// ErrConstructObject
	{

		username := "username"
		existsReturn := false
		serv.EXPECT().ContainsUser(ctx, username).Return(existsReturn, nil)

		userIDReturn := ""
		genID.EXPECT().Generate(ctx).Return(userIDReturn, errors.New("genID.Generate failed"))

		user := entity.User{
			Username: username,
		}
		userReturn := entity.UserExtend{}

		userResult, errSignUp := u.SignUp(ctx, user)
		require.NotNil(t, errSignUp, "[usecase.test.SignUpFailure(ErrConstructObject)]: expected error from u.SignUp(ctx, user)")

		if !reflect.DeepEqual(userReturn, userResult) {
			t.Errorf("[usecase.test.SignUpFailure(ErrConstructObject)]: results not match, want %v, have %v", userReturn, userResult)
			return
		}

		require.Truef(t, errors.Is(errSignUp, interfaces.ErrConstructObject), "[usecase.test.SignUpFailure(ErrConstructObject)]: results not match, want %v, have %v", interfaces.ErrConstructObject, errSignUp)

	}

	// ErrServAddUser
	{
		username := "username"
		existsReturn := false
		serv.EXPECT().ContainsUser(ctx, username).Return(existsReturn, nil)

		userIDReturn := "userID"
		genID.EXPECT().Generate(ctx).Return(userIDReturn, nil)

		user := entity.User{
			Username: username,
		}
		userExtend := entity.NewUserExtend(user, userIDReturn)

		serv.EXPECT().AddUser(ctx, *userExtend).Return(errors.New("serv.AddUser failed"))

		userReturn := entity.UserExtend{}

		userResult, errSignUp := u.SignUp(ctx, user)
		require.NotNil(t, errSignUp, "[usecase.test.SignUpFailure(ErrConstructObject)]: expected error from u.SignUp(ctx, user)")

		if !reflect.DeepEqual(userReturn, userResult) {
			t.Errorf("[usecase.test.SignUpSuccess]: results not match, want %v, have %v", userReturn, userResult)
			return
		}
	}
}

func TestLoginSuccess(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	serv := mock.NewMockIService(ctrl)
	genID := mock.NewMockIGeneratorID(ctrl)
	u := NewUseCase(serv, genID)

	// LoginSuccess
	{
		username := "username"
		password := "password"
		userReturn := entity.UserExtend{
			User: entity.User{
				Username: username,
				Password: password,
			},
		}
		serv.EXPECT().GetUser(ctx, username).Return(userReturn, nil)

		userResult, errLogin := u.Login(ctx, username, password)
		require.Nilf(t, errLogin, "[usecase.test.LoginSuccess]: u.SignUp(ctx, user) failed: %v", errLogin)

		if !reflect.DeepEqual(userReturn, userResult) {
			t.Errorf("[usecase.test.LoginSuccess]: results not match, want %v, have %v", userReturn, userResult)
			return
		}
	}
}

func TestLoginFailure(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	serv := mock.NewMockIService(ctrl)
	genID := mock.NewMockIGeneratorID(ctrl)
	u := NewUseCase(serv, genID)

	// ErrServGetUser
	{
		username := "username"
		password := "password"
		userReturn := entity.UserExtend{}
		serv.EXPECT().GetUser(ctx, username).Return(userReturn, errors.New("serv.GetUser failed"))

		userResult, errLogin := u.Login(ctx, username, password)
		require.NotNil(t, errLogin, "[usecase.test.LoginFailure(ErrServGetUser)]: expected error from u.Login(ctx, username, password)")

		if !reflect.DeepEqual(userReturn, userResult) {
			t.Errorf("[usecase.test.LoginFailure(ErrServGetUser)]: results not match, want %v, have %v", userReturn, userResult)
			return
		}
	}

	// ErrInvalidPassword
	{
		username := "username"
		password := "password"
		userReturn := entity.UserExtend{
			User: entity.User{
				Username: username,
				Password: password + ".",
			},
		}
		serv.EXPECT().GetUser(ctx, username).Return(userReturn, nil)

		userResult, errLogin := u.Login(ctx, username, password)
		require.NotNil(t, errLogin, "[usecase.test.LoginFailure(ErrInvalidPassword)]: expected error from u.Login(ctx, username, password)")

		userReturn = entity.UserExtend{}
		if !reflect.DeepEqual(userReturn, userResult) {
			t.Errorf("[usecase.test.LoginFailure(ErrInvalidPassword)]: results not match, want %v, have %v", userReturn, userResult)
			return
		}

		require.Truef(t, errors.Is(errLogin, interfaces.ErrInvalidPassword), "[usecase.test.LoginFailure(ErrInvalidPassword)]: results not match, want %v, have %v", interfaces.ErrInvalidPassword, errLogin)
	}
}
