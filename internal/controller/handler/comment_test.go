package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
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
	"gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/internal/mock"
)

func TestAddCommentSuccess(t *testing.T) {
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

	// AddCommentSuccess
	{
		bodyComment := "comment"
		objReq := struct {
			Comment string
		}{
			Comment: bodyComment,
		}
		bodyReq, errMarshalReq := json.Marshal(objReq)
		require.Nilf(t, errMarshalReq, "[handler.test.AddCommentSuccess]: json.Marshal(objReq) failed: %v", errMarshalReq)

		reader := bytes.NewReader(bodyReq)

		req := httptest.NewRequest("POST", "/api/post/postID", reader).WithContext(ctxWithValue)
		vars := map[string]string{
			"postID": "postID",
		}
		req = mux.SetURLVars(req, vars)

		ctxReq := req.Context()

		postID := "postID"
		comment := entity.Comment{
			Author: entity.Author{
				Username: username,
				ID:       userID,
			},
			Body:        bodyComment,
			CreatedTime: callTime,
		}
		commentID := "commentID"
		postReturn := entity.PostExtend{
			Post: entity.Post{
				Comments: []entity.CommentExtend{
					{
						Comment: entity.Comment{
							Author: entity.Author{
								Username: username,
								ID:       userID,
							},
							Body:        bodyComment,
							CreatedTime: callTime,
						},
						ID: commentID,
					},
				},
			},
		}
		u.EXPECT().AddComment(ctxReq, postID, comment).Return(postReturn, nil)

		w := httptest.NewRecorder()

		h.AddComment(w, req)

		resp := w.Result()
		bodyResp, errIoReadAll := io.ReadAll(resp.Body)
		require.Nilf(t, errIoReadAll, "[handler.test.AddCommentSuccess]: io.ReadAll(resp.Body) failed: %v", errIoReadAll)

		postResult := entity.PostExtend{}

		errUnmarshalResp := json.Unmarshal(bodyResp, &postResult)
		require.Nilf(t, errUnmarshalResp, "[handler.test.AddCommentSuccess]: json.Unmarshal(bodyResp, &objResp)failed: %v", errUnmarshalResp)

		require.Equal(t, postReturn.Comments[0].CreatedTime, postResult.Comments[0].CreatedTime, "fail")
		if !reflect.DeepEqual(postReturn, postResult) {
			t.Errorf("[handler.test.AddCommentSuccess]: results not match, want %v\n, have %v", postReturn, postResult)
			return
		}
	}
}

func TestAddCommentFailure(t *testing.T) {
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

	// AddCommentSuccess
	{
		bodyComment := "comment"
		objReq := struct {
			Comment string
		}{
			Comment: bodyComment,
		}
		bodyReq, errMarshalReq := json.Marshal(objReq)
		require.Nilf(t, errMarshalReq, "[handler.test.AddCommentSuccess]: json.Marshal(objReq) failed: %v", errMarshalReq)

		reader := bytes.NewReader(bodyReq)

		req := httptest.NewRequest("POST", "/api/post/postID", reader).WithContext(ctxWithValue)
		vars := map[string]string{
			"postID": "postID",
		}
		req = mux.SetURLVars(req, vars)

		ctxReq := req.Context()

		postID := "postID"
		comment := entity.Comment{
			Author: entity.Author{
				Username: username,
				ID:       userID,
			},
			Body:        bodyComment,
			CreatedTime: callTime,
		}
		commentID := "commentID"
		postReturn := entity.PostExtend{
			Post: entity.Post{
				Comments: []entity.CommentExtend{
					{
						Comment: entity.Comment{
							Author: entity.Author{
								Username: username,
								ID:       userID,
							},
							Body:        bodyComment,
							CreatedTime: callTime,
						},
						ID: commentID,
					},
				},
			},
		}
		u.EXPECT().AddComment(ctxReq, postID, comment).Return(postReturn, nil)

		w := httptest.NewRecorder()

		h.AddComment(w, req)

		resp := w.Result()
		bodyResp, errIoReadAll := io.ReadAll(resp.Body)
		require.Nilf(t, errIoReadAll, "[handler.test.AddCommentSuccess]: io.ReadAll(resp.Body) failed: %v", errIoReadAll)

		postResult := entity.PostExtend{}

		errUnmarshalResp := json.Unmarshal(bodyResp, &postResult)
		require.Nilf(t, errUnmarshalResp, "[handler.test.AddCommentSuccess]: json.Unmarshal(bodyResp, &objResp)failed: %v", errUnmarshalResp)

		if !reflect.DeepEqual(postReturn, postResult) {
			t.Errorf("[handler.test.AddCommentSuccess]: results not match, want %v, have %v", postReturn, postResult)
			return
		}
	}
}

func TestDeleteCommentSuccess(t *testing.T) {
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

	// DeleteCommentSuccess
	{
		bodyComment := "comment"
		objReq := struct {
			Comment string
		}{
			Comment: bodyComment,
		}
		bodyReq, errMarshalReq := json.Marshal(objReq)
		require.Nilf(t, errMarshalReq, "[handler.test.AddCommentSuccess]: json.Marshal(objReq) failed: %v", errMarshalReq)

		reader := bytes.NewReader(bodyReq)

		postID := "postID"
		commentID := "commentID"
		req := httptest.NewRequest("DELETE", "/api/post/postID/commentID", reader).WithContext(ctxWithValue)
		vars := map[string]string{
			"postID":    postID,
			"commentID": commentID,
		}
		req = mux.SetURLVars(req, vars)

		ctxReq := req.Context()

		postReturn := entity.PostExtend{
			ID: postID,
		}
		u.EXPECT().DeleteComment(ctxReq, username, postID, commentID).Return(postReturn, nil)

		w := httptest.NewRecorder()

		h.DeleteComment(w, req)

		resp := w.Result()
		bodyResp, errIoReadAll := io.ReadAll(resp.Body)
		require.Nilf(t, errIoReadAll, "[handler.test.DeleteCommentSuccess]: io.ReadAll(resp.Body) failed: %v", errIoReadAll)

		postResult := entity.PostExtend{}

		errUnmarshalResp := json.Unmarshal(bodyResp, &postResult)
		require.Nilf(t, errUnmarshalResp, "[handler.test.DeleteCommentSuccess]: json.Unmarshal(bodyResp, &postResult)failed: %v", errUnmarshalResp)

		if !reflect.DeepEqual(postReturn, postResult) {
			t.Errorf("[handler.test.DeleteCommentSuccess]: results not match, want %v, have %v", postReturn, postResult)
			return
		}
	}
}
