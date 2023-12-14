package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/internal/entity"
	"gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/internal/mock"
)

func TestSignUpSuccess(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	u := mock.NewMockIUseCase(ctrl)
	a := mock.NewMockIAuthManager(ctrl)
	h := NewHandler(u, nil, a)

	// SignUpSuccess
	{

		password := "password"
		username := "username"
		objReq := struct {
			Password string
			Username string
		}{
			Password: password,
			Username: username,
		}
		bodyReq, errMarshalReq := json.Marshal(objReq)
		require.Nilf(t, errMarshalReq, "[handler.test.SignUpSuccess]: json.Marshal(objReq) failed: %v", errMarshalReq)

		reader := bytes.NewReader(bodyReq)

		req := httptest.NewRequest("POST", "/api/register", reader).WithContext(ctx)

		ctxReq := req.Context()

		user := entity.User{
			Username: username,
			Password: password,
		}
		userID := "userID"
		userReturn := entity.UserExtend{
			User: user,
			ID:   userID,
		}
		u.EXPECT().SignUp(ctxReq, user).Return(userReturn, nil)

		tokenReturn := "token"

		a.EXPECT().CreateSession(userReturn).Return(tokenReturn, nil)

		w := httptest.NewRecorder()

		h.SignUp(w, req)

		resp := w.Result()
		bodyResp, errIoReadAll := io.ReadAll(resp.Body)
		require.Nilf(t, errIoReadAll, "[handler.test.SignUpSuccess]: io.ReadAll(resp.Body) failed: %v", errIoReadAll)

		objResp := struct {
			Token string
		}{}

		errUnmarshalResp := json.Unmarshal(bodyResp, &objResp)
		require.Nilf(t, errUnmarshalResp, "[handler.test.SignUpSuccess]: json.Unmarshal(bodyResp, &tokenResult)failed: %v", errUnmarshalResp)

		tokenResult := objResp.Token
		if !reflect.DeepEqual(tokenReturn, tokenResult) {
			t.Errorf("[handler.test.SignUpSuccess]: results not match, want %v, have %v", tokenReturn, tokenResult)
			return
		}
	}

}

func TestLoginSuccess(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	u := mock.NewMockIUseCase(ctrl)
	a := mock.NewMockIAuthManager(ctrl)
	h := NewHandler(u, nil, a)

	// LoginSuccess
	{

		password := "password"
		username := "username"
		objReq := struct {
			Password string
			Username string
		}{
			Password: password,
			Username: username,
		}
		bodyReq, errMarshalReq := json.Marshal(objReq)
		require.Nilf(t, errMarshalReq, "[handler.test.LoginSuccess]: json.Marshal(objReq) failed: %v", errMarshalReq)

		reader := bytes.NewReader(bodyReq)

		req := httptest.NewRequest("POST", "/api/login", reader).WithContext(ctx)

		ctxReq := req.Context()

		userID := "userID"
		userReturn := entity.UserExtend{
			User: entity.User{
				Username: username,
				Password: password,
			},
			ID: userID,
		}
		u.EXPECT().Login(ctxReq, username, password).Return(userReturn, nil)

		tokenReturn := "token"

		a.EXPECT().CreateSession(userReturn).Return(tokenReturn, nil)

		w := httptest.NewRecorder()

		h.Login(w, req)

		resp := w.Result()
		bodyResp, errIoReadAll := io.ReadAll(resp.Body)
		require.Nilf(t, errIoReadAll, "[handler.test.LoginSuccess]: io.ReadAll(resp.Body) failed: %v", errIoReadAll)

		objResp := struct {
			Token string
		}{}

		errUnmarshalResp := json.Unmarshal(bodyResp, &objResp)
		require.Nilf(t, errUnmarshalResp, "[handler.test.LoginSuccess]: json.Unmarshal(bodyResp, &tokenResult)failed: %v", errUnmarshalResp)

		tokenResult := objResp.Token
		if !reflect.DeepEqual(tokenReturn, tokenResult) {
			t.Errorf("[handler.test.LoginSuccess]: results not match, want %v, have %v", tokenReturn, tokenResult)
			return
		}
	}

}
