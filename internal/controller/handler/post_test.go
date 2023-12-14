package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
	"gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/internal/controller/auth"
	"gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/internal/controller/middleware"
	"gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/internal/entity"
	"gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/internal/interfaces"
	"gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/internal/mock"
)

func TestGetPostsSuccess(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	u := mock.NewMockIUseCase(ctrl)
	h := NewHandler(u, nil, nil)

	// GetPostsSuccess
	{

		req := httptest.NewRequest("GET", "/api/posts/", nil).WithContext(ctx)

		ctxReq := req.Context()

		postsResult := []entity.PostExtend{
			{
				ID: "postID1",
			},
			{
				ID: "postID2",
			},
			{
				ID: "postID3",
			},
		}

		u.EXPECT().GetPosts(ctxReq).Return(postsResult, nil)

		w := httptest.NewRecorder()

		h.GetPosts(w, req)

		resp := w.Result()
		bodyResp, errIoReadAll := io.ReadAll(resp.Body)
		require.Nilf(t, errIoReadAll, "[handler.test.GetPostsSuccess]: io.ReadAll(resp.Body) failed: %v", errIoReadAll)

		postsReturn := []entity.PostExtend{}

		errUnmarshalResp := json.Unmarshal(bodyResp, &postsReturn)
		require.Nilf(t, errUnmarshalResp, "[handler.test.GetPostsSuccess]: json.Unmarshal(bodyResp, &postsReturn)failed: %v", errUnmarshalResp)

		if !reflect.DeepEqual(postsReturn, postsResult) {
			t.Errorf("[handler.test.GetPostsSuccess]: results not match, want %v, have %v", postsReturn, postsResult)
			return
		}
	}
}

func TestGetPostsFailure(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	u := mock.NewMockIUseCase(ctrl)
	h := NewHandler(u, nil, nil)

	// GetPostsFailure
	{
		req := httptest.NewRequest("GET", "/api/posts/", nil).WithContext(ctx)

		ctxReq := req.Context()

		postsResult := []entity.PostExtend{}

		u.EXPECT().GetPosts(ctxReq).Return(postsResult, errors.New("usecase.GetPosts failed"))

		codeResult := http.StatusInternalServerError
		w := httptest.NewRecorder()

		h.GetPosts(w, req)

		codeResp := w.Code

		require.Equalf(t, codeResult, codeResp, "[handler.test.GetPostsFailure]: results not match, want %v, have %v", codeResp, codeResult)
	}

	// ErrWriteResp
	{
		req := httptest.NewRequest("GET", "/api/posts/", nil).WithContext(ctx)

		ctxReq := req.Context()

		postsResult := []entity.PostExtend{
			{
				ID: "postID1",
			},
			{
				ID: "postID2",
			},
			{
				ID: "postID3",
			},
		}

		u.EXPECT().GetPosts(ctxReq).Return(postsResult, nil)

		codeResult := http.StatusOK
		d := &dummyResponseWriter{}

		h.GetPosts(d, req)

		codeResp := d.Code

		require.Equalf(t, codeResult, codeResp, "[handler.test.ErrWriteResp]: results not match, want %v, have %v", codeResp, codeResult)
	}
}

func TestAddPostSuccess(t *testing.T) {
	ctx := context.Background()
	username := "username"
	userID := "id"
	session := auth.Session{
		Username: username,
		ID:       userID,
	}
	callTime := time.Now().Round(0)
	ctxWithValue := context.WithValue(ctx, auth.SessKey, session)
	ctxWithValue = context.WithValue(ctxWithValue, middleware.CallTimeKey, callTime)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	u := mock.NewMockIUseCase(ctrl)
	h := NewHandler(u, nil, nil)

	// AddTextPostSuccess
	{
		category := "category"
		text := "text"
		title := "title"
		kind := "type"
		objReq := struct {
			Category string
			Text     string
			Title    string
			Type     string
		}{
			Category: category,
			Text:     text,
			Title:    title,
			Type:     kind,
		}
		bodyReq, errMarshalReq := json.Marshal(objReq)
		require.Nilf(t, errMarshalReq, "[handler.test.AddTextPostSuccess]: json.Marshal(objReq) failed: %v", errMarshalReq)

		reader := bytes.NewReader(bodyReq)

		req := httptest.NewRequest("POST", "/api/post", reader).WithContext(ctxWithValue)

		ctxReq := req.Context()

		postID := "postID"
		post := entity.Post{
			Score:    0,
			Views:    1,
			Category: category,
			Text:     text,
			Title:    title,
			Type:     kind,
			Author: entity.Author{
				Username: username,
				ID:       userID,
			},
			Votes:            []entity.Vote{},
			Comments:         []entity.CommentExtend{},
			CreatedTime:      callTime,
			UpvotePercentage: 0,
		}
		postReturn := entity.PostExtend{
			Post: post,
			ID:   postID,
		}
		u.EXPECT().AddPost(ctxReq, post).Return(postReturn, nil)

		w := httptest.NewRecorder()

		h.AddPost(w, req)

		resp := w.Result()
		bodyResp, errIoReadAll := io.ReadAll(resp.Body)
		require.Nilf(t, errIoReadAll, "[handler.test.AddTextPostSuccess]: io.ReadAll(resp.Body) failed: %v", errIoReadAll)

		postResult := entity.PostExtend{}

		errUnmarshalResp := json.Unmarshal(bodyResp, &postResult)
		require.Nilf(t, errUnmarshalResp, "[handler.test.AddTextPostSuccess]: json.Unmarshal(bodyResp, &objResp)failed: %v", errUnmarshalResp)

		if !reflect.DeepEqual(postReturn, postResult) {
			t.Errorf("[handler.test.AddTextPostSuccess]: results not match, want %v, have %v", postReturn, postResult)
			return
		}
	}

	// AddLinkPostSuccess
	{
		category := "category"
		url := "url"
		title := "title"
		kind := "link"
		objReq := struct {
			Category string
			URL      string
			Title    string
			Type     string
		}{
			Category: category,
			URL:      url,
			Title:    title,
			Type:     kind,
		}
		bodyReq, errMarshalReq := json.Marshal(objReq)
		require.Nilf(t, errMarshalReq, "[handler.test.AddTextPostSuccess]: json.Marshal(objReq) failed: %v", errMarshalReq)

		reader := bytes.NewReader(bodyReq)

		req := httptest.NewRequest("POST", "/api/post", reader).WithContext(ctxWithValue)

		ctxReq := req.Context()

		postID := "postID"
		post := entity.Post{
			Score:    0,
			Views:    1,
			Category: category,
			URL:      url,
			Title:    title,
			Type:     kind,
			Author: entity.Author{
				Username: username,
				ID:       userID,
			},
			Votes:            []entity.Vote{},
			Comments:         []entity.CommentExtend{},
			CreatedTime:      callTime,
			UpvotePercentage: 0,
		}
		postReturn := entity.PostExtend{
			Post: post,
			ID:   postID,
		}
		u.EXPECT().AddPost(ctxReq, post).Return(postReturn, nil)

		w := httptest.NewRecorder()

		h.AddPost(w, req)

		resp := w.Result()
		bodyResp, errIoReadAll := io.ReadAll(resp.Body)
		require.Nilf(t, errIoReadAll, "[handler.test.AddTextPostSuccess]: io.ReadAll(resp.Body) failed: %v", errIoReadAll)

		postResult := entity.PostExtend{}

		errUnmarshalResp := json.Unmarshal(bodyResp, &postResult)
		require.Nilf(t, errUnmarshalResp, "[handler.test.AddTextPostSuccess]: json.Unmarshal(bodyResp, &objResp)failed: %v", errUnmarshalResp)

		if !reflect.DeepEqual(postReturn, postResult) {
			t.Errorf("[handler.test.AddTextPostSuccess]: results not match, want %v, have %v", postReturn, postResult)
			return
		}
	}
}

func TestAddPostFailure(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	u := mock.NewMockIUseCase(ctrl)
	h := NewHandler(u, nil, nil)

	// ErrUnauthorized
	{
		ctx := context.Background()

		req := httptest.NewRequest("POST", "/api/post", nil).WithContext(ctx)

		codeResult := http.StatusUnauthorized
		w := httptest.NewRecorder()

		h.AddPost(w, req)

		codeResp := w.Code

		require.Equalf(t, codeResult, codeResp, "[handler.test.ErrUnauthorized]: results not match, want %v, have %v", codeResult, codeResp)

	}

	// ErrTypeAssertionSession
	{
		ctx := context.Background()
		session := struct {
		}{}
		ctxWithValue := context.WithValue(ctx, auth.SessKey, session)

		req := httptest.NewRequest("POST", "/api/post", nil).WithContext(ctxWithValue)

		codeResult := http.StatusInternalServerError
		w := httptest.NewRecorder()

		h.AddPost(w, req)

		codeResp := w.Code

		require.Equalf(t, codeResult, codeResp, "[handler.test.ErrTypeAssertionSession]: results not match, want %v, have %v", codeResult, codeResp)

	}

	// ErrCallTimeNotInContext
	{
		ctx := context.Background()
		username := "username"
		userID := "id"
		session := auth.Session{
			Username: username,
			ID:       userID,
		}
		ctxWithValue := context.WithValue(ctx, auth.SessKey, session)

		req := httptest.NewRequest("POST", "/api/post", nil).WithContext(ctxWithValue)

		codeResult := http.StatusInternalServerError
		w := httptest.NewRecorder()

		h.AddPost(w, req)

		codeResp := w.Code

		require.Equalf(t, codeResult, codeResp, "[handler.test.ErrCallTimeNotInContext]: results not match, want %v, have %v", codeResult, codeResp)
	}

	// ErrTypeAssertionCallTime
	{
		ctx := context.Background()
		username := "username"
		userID := "id"
		session := auth.Session{
			Username: username,
			ID:       userID,
		}
		callTime := 0
		ctxWithValue := context.WithValue(ctx, auth.SessKey, session)
		ctxWithValue = context.WithValue(ctxWithValue, middleware.CallTimeKey, callTime)

		req := httptest.NewRequest("POST", "/api/post", nil).WithContext(ctxWithValue)

		codeResult := http.StatusInternalServerError
		w := httptest.NewRecorder()

		h.AddPost(w, req)

		codeResp := w.Code

		require.Equalf(t, codeResult, codeResp, "[handler.test.ErrTypeAssertionCallTime]: results not match, want %v, have %v", codeResult, codeResp)
	}

	// ErrReadAllBody
	{
		ctx := context.Background()
		username := "username"
		userID := "id"
		session := auth.Session{
			Username: username,
			ID:       userID,
		}
		callTime := time.Now().Round(0)
		ctxWithValue := context.WithValue(ctx, auth.SessKey, session)
		ctxWithValue = context.WithValue(ctxWithValue, middleware.CallTimeKey, callTime)

		b := &bodyReader{}
		req := httptest.NewRequest("POST", "/api/post", b).WithContext(ctxWithValue)

		codeResult := http.StatusInternalServerError
		w := httptest.NewRecorder()

		h.AddPost(w, req)

		codeResp := w.Code

		require.Equalf(t, codeResult, codeResp, "[handler.test.ErrReadAllBody]: results not match, want %v, have %v", codeResult, codeResp)
	}

	// ErrUnmarshalReqBody
	{
		ctx := context.Background()
		username := "username"
		userID := "id"
		session := auth.Session{
			Username: username,
			ID:       userID,
		}
		callTime := time.Now().Round(0)
		ctxWithValue := context.WithValue(ctx, auth.SessKey, session)
		ctxWithValue = context.WithValue(ctxWithValue, middleware.CallTimeKey, callTime)

		category := false
		url := "url"
		title := "title"
		kind := "link"
		objReq := struct {
			Category bool
			URL      string
			Title    string
			Type     string
		}{
			Category: category,
			URL:      url,
			Title:    title,
			Type:     kind,
		}
		bodyReq, errMarshalReq := json.Marshal(objReq)
		require.Nilf(t, errMarshalReq, "[handler.test.ErrUnmarshalReqBody]: json.Marshal(objReq) failed: %v", errMarshalReq)

		reader := bytes.NewReader(bodyReq)

		req := httptest.NewRequest("POST", "/api/post", reader).WithContext(ctxWithValue)

		codeResult := http.StatusInternalServerError
		w := httptest.NewRecorder()

		h.AddPost(w, req)

		codeResp := w.Code

		require.Equalf(t, codeResult, codeResp, "[handler.test.ErrUnmarshalReqBody]: results not match, want %v, have %v", codeResult, codeResp)
	}

	// ErrAddPostConstructObjectMsg
	{
		ctx := context.Background()
		username := "username"
		userID := "id"
		session := auth.Session{
			Username: username,
			ID:       userID,
		}
		callTime := time.Now().Round(0)
		ctxWithValue := context.WithValue(ctx, auth.SessKey, session)
		ctxWithValue = context.WithValue(ctxWithValue, middleware.CallTimeKey, callTime)

		category := "category"
		url := "url"
		title := "title"
		kind := "link"
		objReq := struct {
			Category string
			URL      string
			Title    string
			Type     string
		}{
			Category: category,
			URL:      url,
			Title:    title,
			Type:     kind,
		}
		bodyReq, errMarshalReq := json.Marshal(objReq)
		require.Nilf(t, errMarshalReq, "[handler.test.ErrAddPostConstructObjectMsg]: json.Marshal(objReq) failed: %v", errMarshalReq)

		reader := bytes.NewReader(bodyReq)

		req := httptest.NewRequest("POST", "/api/post", reader).WithContext(ctxWithValue)

		ctxReq := req.Context()

		author := entity.Author{
			Username: username,
			ID:       userID,
		}
		post := entity.Post{
			Score:            0,
			Views:            1,
			Type:             objReq.Type,
			Title:            objReq.Title,
			Author:           author,
			Category:         objReq.Category,
			URL:              objReq.URL,
			Votes:            []entity.Vote{},
			Comments:         []entity.CommentExtend{},
			CreatedTime:      callTime,
			UpvotePercentage: 0,
		}
		postReturn := entity.PostExtend{}
		errAddPostReturn := interfaces.ErrConstructObject
		u.EXPECT().AddPost(ctxReq, post).Return(postReturn, errAddPostReturn)

		w := httptest.NewRecorder()

		objResult := struct {
			Message string `json:"message"`
		}{
			Message: interfaces.ErrConstructObject.Error(),
		}

		h.AddPost(w, req)

		resp := w.Result()
		bodyResp, errIoReadAll := io.ReadAll(resp.Body)
		require.Nilf(t, errIoReadAll, "[handler.test.ErrAddPostConstructObjectMsg]: io.ReadAll(resp.Body) failed: %v", errIoReadAll)

		objResp := struct {
			Message string `json:"message"`
		}{}

		errUnmarshalResp := json.Unmarshal(bodyResp, &objResp)
		require.Nilf(t, errUnmarshalResp, "[handler.test.ErrAddPostConstructObjectMsg]: json.Unmarshal(bodyResp, &objResp)failed: %v", errUnmarshalResp)

		if !reflect.DeepEqual(objResp, objResult) {
			t.Errorf("[handler.test.ErrAddPostConstructObjectMsg]: results not match, want %v, have %v", objResp, objResult)
			return
		}

	}

	// ErrAddPostUnknownErrorMsg
	{
		ctx := context.Background()
		username := "username"
		userID := "id"
		session := auth.Session{
			Username: username,
			ID:       userID,
		}
		callTime := time.Now().Round(0)
		ctxWithValue := context.WithValue(ctx, auth.SessKey, session)
		ctxWithValue = context.WithValue(ctxWithValue, middleware.CallTimeKey, callTime)

		category := "category"
		url := "url"
		title := "title"
		kind := "link"
		objReq := struct {
			Category string
			URL      string
			Title    string
			Type     string
		}{
			Category: category,
			URL:      url,
			Title:    title,
			Type:     kind,
		}
		bodyReq, errMarshalReq := json.Marshal(objReq)
		require.Nilf(t, errMarshalReq, "[handler.test.ErrAddPostUnknownErrorMsg]: json.Marshal(objReq) failed: %v", errMarshalReq)

		reader := bytes.NewReader(bodyReq)

		req := httptest.NewRequest("POST", "/api/post", reader).WithContext(ctxWithValue)

		ctxReq := req.Context()

		post := entity.Post{
			Score: 0,
			Views: 1,
			Type:  objReq.Type,
			Title: objReq.Title,
			Author: entity.Author{
				Username: username,
				ID:       userID,
			},
			Category:         objReq.Category,
			URL:              objReq.URL,
			Votes:            []entity.Vote{},
			Comments:         []entity.CommentExtend{},
			CreatedTime:      callTime,
			UpvotePercentage: 0,
		}
		postReturn := entity.PostExtend{}
		errAddPostReturn := errors.New("usecase.AddPost failed")
		u.EXPECT().AddPost(ctxReq, post).Return(postReturn, errAddPostReturn)

		w := httptest.NewRecorder()

		objResult := struct {
			Message string `json:"message"`
		}{
			Message: interfaces.ErrUnknownErrorMsg.Error(),
		}

		h.AddPost(w, req)

		resp := w.Result()
		bodyResp, errIoReadAll := io.ReadAll(resp.Body)
		require.Nilf(t, errIoReadAll, "[handler.test.ErrAddPostUnknownErrorMsg]: io.ReadAll(resp.Body) failed: %v", errIoReadAll)

		objResp := struct {
			Message string `json:"message"`
		}{}

		errUnmarshalResp := json.Unmarshal(bodyResp, &objResp)
		require.Nilf(t, errUnmarshalResp, "[handler.test.ErrAddPostUnknownErrorMsg]: json.Unmarshal(bodyResp, &objResp)failed: %v", errUnmarshalResp)

		if !reflect.DeepEqual(objResp, objResult) {
			t.Errorf("[handler.test.ErrAddPostUnknownErrorMsg]: results not match, want %v, have %v", objResp, objResult)
			return
		}
	}

	// ErrAddPostWriteResp
	{
		ctx := context.Background()
		username := "username"
		userID := "id"
		session := auth.Session{
			Username: username,
			ID:       userID,
		}
		callTime := time.Now().Round(0)
		ctxWithValue := context.WithValue(ctx, auth.SessKey, session)
		ctxWithValue = context.WithValue(ctxWithValue, middleware.CallTimeKey, callTime)

		category := "category"
		url := "url"
		title := "title"
		kind := "link"
		objReq := struct {
			Category string
			URL      string
			Title    string
			Type     string
		}{
			Category: category,
			URL:      url,
			Title:    title,
			Type:     kind,
		}
		bodyReq, errMarshalReq := json.Marshal(objReq)
		require.Nilf(t, errMarshalReq, "[handler.test.ErrAddPostWriteResp]: json.Marshal(objReq) failed: %v", errMarshalReq)

		reader := bytes.NewReader(bodyReq)

		req := httptest.NewRequest("POST", "/api/post", reader).WithContext(ctxWithValue)

		ctxReq := req.Context()

		post := entity.Post{
			Score: 0,
			Views: 1,
			Type:  objReq.Type,
			Title: objReq.Title,
			Author: entity.Author{
				Username: username,
				ID:       userID,
			},
			Category:         objReq.Category,
			URL:              objReq.URL,
			Votes:            []entity.Vote{},
			Comments:         []entity.CommentExtend{},
			CreatedTime:      callTime,
			UpvotePercentage: 0,
		}
		postReturn := entity.PostExtend{}
		errAddPostReturn := errors.New("usecase.AddPost failed")
		u.EXPECT().AddPost(ctxReq, post).Return(postReturn, errAddPostReturn)

		d := &dummyResponseWriter{}

		codeResult := http.StatusInternalServerError
		h.AddPost(d, req)

		codeResp := d.Code

		require.Equalf(t, codeResult, codeResp, "[handler.test.ErrAddPostWriteResp]: results not match, want %v, have %v", codeResult, codeResp)
	}

	// ErrWriteResp
	{
		ctx := context.Background()
		username := "username"
		userID := "id"
		session := auth.Session{
			Username: username,
			ID:       userID,
		}
		callTime := time.Now().Round(0)
		ctxWithValue := context.WithValue(ctx, auth.SessKey, session)
		ctxWithValue = context.WithValue(ctxWithValue, middleware.CallTimeKey, callTime)

		category := "category"
		url := "url"
		title := "title"
		kind := "link"
		objReq := struct {
			Category string
			URL      string
			Title    string
			Type     string
		}{
			Category: category,
			URL:      url,
			Title:    title,
			Type:     kind,
		}
		bodyReq, errMarshalReq := json.Marshal(objReq)
		require.Nilf(t, errMarshalReq, "[handler.test.ErrWriteResp]: json.Marshal(objReq) failed: %v", errMarshalReq)

		reader := bytes.NewReader(bodyReq)

		req := httptest.NewRequest("POST", "/api/post", reader).WithContext(ctxWithValue)

		ctxReq := req.Context()

		postID := "postID"
		post := entity.Post{
			Score:    0,
			Views:    1,
			Category: category,
			URL:      url,
			Title:    title,
			Type:     kind,
			Author: entity.Author{
				Username: username,
				ID:       userID,
			},
			Votes:            []entity.Vote{},
			Comments:         []entity.CommentExtend{},
			CreatedTime:      callTime,
			UpvotePercentage: 0,
		}
		postReturn := entity.PostExtend{
			Post: post,
			ID:   postID,
		}
		u.EXPECT().AddPost(ctxReq, post).Return(postReturn, nil)

		d := &dummyResponseWriter{}

		codeResult := http.StatusOK
		h.AddPost(d, req)

		codeResp := d.Code

		require.Equalf(t, codeResult, codeResp, "[handler.test.ErrWriteResp]: results not match, want %v, have %v", codeResult, codeResp)
	}

}

func TestGetPostsWithCategorySuccess(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	u := mock.NewMockIUseCase(ctrl)
	h := NewHandler(u, nil, nil)

	// GetPostsWithCategorySuccess
	{

		categoryName := "categoryName"
		req := httptest.NewRequest("GET", "/api/posts/"+categoryName, nil).WithContext(ctx)
		vars := map[string]string{
			"categoryName": categoryName,
		}

		req = mux.SetURLVars(req, vars)

		ctxReq := req.Context()

		postsResult := []entity.PostExtend{
			{
				Post: entity.Post{
					Category: categoryName,
				},
				ID: "postID1",
			},
			{
				Post: entity.Post{
					Category: categoryName,
				},
				ID: "postID2",
			},
			{
				Post: entity.Post{
					Category: categoryName,
				},
				ID: "postID3",
			},
		}

		u.EXPECT().GetPostsWithCategory(ctxReq, categoryName).Return(postsResult, nil)

		w := httptest.NewRecorder()

		h.GetPostsWithCategory(w, req)

		resp := w.Result()
		bodyResp, errIoReadAll := io.ReadAll(resp.Body)
		require.Nilf(t, errIoReadAll, "[handler.test.GetPostsWithCategorySuccess]: io.ReadAll(resp.Body) failed: %v", errIoReadAll)

		postsReturn := []entity.PostExtend{}

		errUnmarshalResp := json.Unmarshal(bodyResp, &postsReturn)
		require.Nilf(t, errUnmarshalResp, "[handler.test.GetPostsWithCategorySuccess]: json.Unmarshal(bodyResp, &postsReturn)failed: %v", errUnmarshalResp)

		if !reflect.DeepEqual(postsReturn, postsResult) {
			t.Errorf("[handler.test.GetPostsWithCategorySuccess]: results not match, want %v, have %v", postsReturn, postsResult)
			return
		}
	}
}

func TestGetPostsWithUserSuccess(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	u := mock.NewMockIUseCase(ctrl)
	h := NewHandler(u, nil, nil)

	// GetPostsWithUserSuccess
	{

		username := "username"
		req := httptest.NewRequest("GET", "/api/user/"+username, nil).WithContext(ctx)
		vars := map[string]string{
			"username": username,
		}

		req = mux.SetURLVars(req, vars)

		ctxReq := req.Context()

		postsResult := []entity.PostExtend{
			{
				Post: entity.Post{
					Author: entity.Author{
						Username: username,
					},
				},
				ID: "postID1",
			},
			{
				Post: entity.Post{
					Author: entity.Author{
						Username: username,
					},
				},
				ID: "postID2",
			},
			{
				Post: entity.Post{
					Author: entity.Author{
						Username: username,
					},
				},
				ID: "postID3",
			},
		}

		u.EXPECT().GetPostsWithUser(ctxReq, username).Return(postsResult, nil)

		w := httptest.NewRecorder()

		h.GetPostsWithUser(w, req)

		resp := w.Result()
		bodyResp, errIoReadAll := io.ReadAll(resp.Body)
		require.Nilf(t, errIoReadAll, "[handler.test.GetPostsWithUserSuccess]: io.ReadAll(resp.Body) failed: %v", errIoReadAll)

		postsReturn := []entity.PostExtend{}

		errUnmarshalResp := json.Unmarshal(bodyResp, &postsReturn)
		require.Nilf(t, errUnmarshalResp, "[handler.test.GetPostsWithUserSuccess]: json.Unmarshal(bodyResp, &postsReturn)failed: %v", errUnmarshalResp)

		if !reflect.DeepEqual(postsReturn, postsResult) {
			t.Errorf("[handler.test.GetPostsWithUserSuccess]: results not match, want %v, have %v", postsReturn, postsResult)
			return
		}
	}
}

func TestGetPostSuccess(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	u := mock.NewMockIUseCase(ctrl)
	h := NewHandler(u, nil, nil)

	// GetPostSuccess
	{

		postID := "postID"
		req := httptest.NewRequest("GET", "/api/post/"+postID, nil).WithContext(ctx)
		vars := map[string]string{
			"postID": postID,
		}

		req = mux.SetURLVars(req, vars)

		ctxReq := req.Context()

		postResult := entity.PostExtend{
			ID: postID,
		}

		u.EXPECT().GetPost(ctxReq, postID).Return(postResult, nil)

		w := httptest.NewRecorder()

		h.GetPost(w, req)

		resp := w.Result()
		bodyResp, errIoReadAll := io.ReadAll(resp.Body)
		require.Nilf(t, errIoReadAll, "[handler.test.GetPostSuccess]: io.ReadAll(resp.Body) failed: %v", errIoReadAll)

		postReturn := entity.PostExtend{}

		errUnmarshalResp := json.Unmarshal(bodyResp, &postReturn)
		require.Nilf(t, errUnmarshalResp, "[handler.test.GetPostSuccess]: json.Unmarshal(bodyResp, &postReturn)failed: %v", errUnmarshalResp)

		if !reflect.DeepEqual(postReturn, postResult) {
			t.Errorf("[handler.test.GetPostSuccess]: results not match, want %v, have %v", postReturn, postResult)
			return
		}
	}
}

func TestGetPostFailure(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	u := mock.NewMockIUseCase(ctrl)
	h := NewHandler(u, nil, nil)

	// ErrUnknownErrorMsg
	{

		postID := "postID"
		req := httptest.NewRequest("GET", "/api/post/"+postID, nil).WithContext(ctx)
		vars := map[string]string{
			"postID": postID,
		}

		req = mux.SetURLVars(req, vars)

		ctxReq := req.Context()

		postResult := entity.PostExtend{}

		u.EXPECT().GetPost(ctxReq, postID).Return(postResult, errors.New("usecase.GetPost failed"))

		w := httptest.NewRecorder()

		objResult := struct {
			Message string `json:"message"`
		}{
			Message: interfaces.ErrUnknownErrorMsg.Error(),
		}

		h.GetPost(w, req)

		resp := w.Result()
		bodyResp, errIoReadAll := io.ReadAll(resp.Body)
		require.Nilf(t, errIoReadAll, "[handler.test.GetPostFailure]: io.ReadAll(resp.Body) failed: %v", errIoReadAll)

		objResp := struct {
			Message string `json:"message"`
		}{}

		errUnmarshalResp := json.Unmarshal(bodyResp, &objResp)
		require.Nilf(t, errUnmarshalResp, "[handler.test.GetPostFailure]: json.Unmarshal(bodyResp, &postReturn)failed: %v", errUnmarshalResp)

		if !reflect.DeepEqual(objResp, objResult) {
			t.Errorf("[handler.test.GetPostFailure]: results not match, want %v, have %v", objResult, objResp)
			return
		}
	}

	// ErrPostNotFoundMsg
	{
		postID := "postID"
		req := httptest.NewRequest("GET", "/api/post/"+postID, nil).WithContext(ctx)
		vars := map[string]string{
			"postID": postID,
		}

		req = mux.SetURLVars(req, vars)

		ctxReq := req.Context()

		postResult := entity.PostExtend{}
		errResult := fmt.Errorf("[usecase.GetPost]->%w", interfaces.ErrPostNotExists)
		u.EXPECT().GetPost(ctxReq, postID).Return(postResult, errResult)

		w := httptest.NewRecorder()

		objResult := struct {
			Message string `json:"message"`
		}{
			Message: interfaces.ErrPostNotExists.Error(),
		}

		h.GetPost(w, req)

		resp := w.Result()
		bodyResp, errIoReadAll := io.ReadAll(resp.Body)
		require.Nilf(t, errIoReadAll, "[handler.test.ErrPostNotFoundMsg]: io.ReadAll(resp.Body) failed: %v", errIoReadAll)

		objResp := struct {
			Message string `json:"message"`
		}{}

		errUnmarshalResp := json.Unmarshal(bodyResp, &objResp)
		require.Nilf(t, errUnmarshalResp, "[handler.test.ErrPostNotFoundMsg]: json.Unmarshal(bodyResp, &postReturn)failed: %v", errUnmarshalResp)

		if !reflect.DeepEqual(objResp, objResult) {
			t.Errorf("[handler.test.ErrPostNotFoundMsg]: results not match, want %v, have %v", objResult, objResp)
			return
		}
	}

	// ErrWriteResp
	{
		postID := "postID"
		req := httptest.NewRequest("GET", "/api/post/"+postID, nil).WithContext(ctx)
		vars := map[string]string{
			"postID": postID,
		}

		req = mux.SetURLVars(req, vars)

		ctxReq := req.Context()

		postResult := entity.PostExtend{
			ID: postID,
		}
		u.EXPECT().GetPost(ctxReq, postID).Return(postResult, nil)

		d := &dummyResponseWriter{}

		codeResult := http.StatusOK
		h.GetPost(d, req)

		codeResp := d.Code
		require.Equalf(t, codeResult, codeResp, "[handler.test.ErrWriteResp]: results not match, want %v, have %v", codeResult, codeResp)
	}
}

func TestDeletePostSuccess(t *testing.T) {
	ctx := context.Background()
	username := "username"
	userID := "id"
	session := auth.Session{
		Username: username,
		ID:       userID,
	}
	callTime := time.Now().Round(0)
	ctxWithValue := context.WithValue(ctx, auth.SessKey, session)
	ctxWithValue = context.WithValue(ctxWithValue, middleware.CallTimeKey, callTime)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	u := mock.NewMockIUseCase(ctrl)
	h := NewHandler(u, nil, nil)

	// DeletePostSuccess
	{

		postID := "postID"
		req := httptest.NewRequest("DELETE", "/api/post/"+postID, nil).WithContext(ctxWithValue)
		vars := map[string]string{
			"postID": postID,
		}
		req = mux.SetURLVars(req, vars)

		ctxReq := req.Context()

		u.EXPECT().DeletePost(ctxReq, username, postID).Return(nil)

		w := httptest.NewRecorder()

		objResult := struct {
			Message string
		}{
			Message: "success",
		}

		h.DeletePost(w, req)

		resp := w.Result()
		bodyResp, errIoReadAll := io.ReadAll(resp.Body)
		require.Nilf(t, errIoReadAll, "[handler.test.DeletePostSuccess]: io.ReadAll(resp.Body) failed: %v", errIoReadAll)

		objResp := struct {
			Message string
		}{}

		errUnmarshalResp := json.Unmarshal(bodyResp, &objResp)
		require.Nilf(t, errUnmarshalResp, "[handler.test.DeletePostSuccess]: json.Unmarshal(bodyResp, &objResp)failed: %v", errUnmarshalResp)

		if !reflect.DeepEqual(objResp, objResult) {
			t.Errorf("[handler.test.DeletePostSuccess]: results not match, want %v, have %v", objResp, objResult)
			return
		}
	}
}
