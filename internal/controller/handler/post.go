package handler

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/internal/controller/auth"
	"gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/internal/controller/middleware"
	"gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/internal/entity"
	"gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/internal/interfaces"
)

func (h *Handler) GetPosts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	posts, errGetPosts := h.useCase.GetPosts(ctx)
	if errGetPosts != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("||ERROR|| [Handler.GetPosts]: h.useCase.GetPosts(ctx) failed: %v\n", errGetPosts)
		return
	}

	resp, errPostsMarshal := json.Marshal(posts)
	if errPostsMarshal != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("||ERROR|| [Handler.GetPosts]: json.Marshal(posts) failed: %v\n", errPostsMarshal)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(resp); err != nil {
		log.Printf("||ERROR|| [Handler.GetPosts]: w.Write(resp) failed: %v\n", err)
		return
	}
}

func (h *Handler) AddPost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	rawSession := r.Context().Value(auth.SessKey)
	if rawSession == nil {
		w.WriteHeader(http.StatusUnauthorized)
		log.Println("||ERROR|| [Handler.AddPost]: unauthorized")
		return
	}
	session, ok := rawSession.(auth.Session)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("||ERROR|| [Handler.AddPost]: type assertion rawSession.(auth.Session) failed")
		return
	}
	username := session.Username
	id := session.ID

	rawCallTime := r.Context().Value(middleware.CallTimeKey)
	if rawCallTime == nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("[||ERROR|| Handler.AddPost]: r.Context().Value(middleware.CallTimeKey) failed")
		return
	}
	callTime, ok := rawCallTime.(time.Time)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("||ERROR|| [Handler.AddPost]: type assertion rawCallTime.(time.Time) failed")
		return
	}

	body, errBodyReadAll := io.ReadAll(r.Body)
	if errBodyReadAll != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("||ERROR|| [Handler.AddPost]: io.ReadAll(req.Body) failed: %v\n", errBodyReadAll)
		return
	}

	objReq := struct {
		Category string
		Text     string
		URL      string
		Title    string
		Type     string
	}{}

	errReqUnmarshal := json.Unmarshal(body, &objReq)
	if errReqUnmarshal != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("||ERROR|| [Handler.AddPost]: json.Unmarshal(body, &objReq) failed: %v\n", errReqUnmarshal)
		return
	}

	author := entity.Author{
		Username: username,
		ID:       id,
	}
	post := entity.Post{
		Score:            0,
		Views:            1,
		Type:             objReq.Type,
		Title:            objReq.Title,
		Author:           author,
		Category:         objReq.Category,
		Text:             objReq.Text,
		URL:              objReq.URL,
		Votes:            []entity.Vote{},
		Comments:         []entity.CommentExtend{},
		CreatedTime:      callTime,
		UpvotePercentage: 0,
	}
	postExtend, errAddPost := h.useCase.AddPost(ctx, post)
	if errAddPost != nil {
		log.Printf("||ERROR|| [Handler.AddPost]->%v", errAddPost)

		objResp := struct {
			Message string `json:"message"`
		}{}
		switch {
		case errors.Is(errAddPost, interfaces.ErrConstructObject):
			objResp.Message = errConstructObjectMsg
		default:
			objResp.Message = errUnknownErrorMsg
		}

		resp, errRespMarshal := json.Marshal(objResp)
		if errRespMarshal != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("||ERROR|| [Handler.AddPost]: json.Marshal(objResp) failed: %v\n", errRespMarshal)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		if _, errRespWrite := w.Write(resp); errRespWrite != nil {
			log.Printf("||ERROR|| [Handler.AddPost]: w.Write(resp) failed: %v\n", errRespWrite)
			return
		}

		return
	}

	resp, errPostMarshal := json.Marshal(postExtend)
	if errPostMarshal != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("||ERROR|| [Handler.AddPost]: json.Marshal(postExtend) failed: %v\n", errPostMarshal)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, errRespWrite := w.Write(resp); errRespWrite != nil {
		log.Printf("||ERROR|| [Handler.AddPost]: w.Write(resp) failed: %v\n", errRespWrite)
		return
	}

}

func (h *Handler) GetPostsWithCategory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	category := mux.Vars(r)["categoryName"]

	posts, errGetPosts := h.useCase.GetPostsWithCategory(ctx, category)
	if errGetPosts != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("||ERROR|| [Handler.GetPostsWithCategory]: h.useCase.GetPostsWithCategory(ctx, category) failed: %v\n", errGetPosts)
		return
	}

	resp, errPostsMarshal := json.Marshal(posts)
	if errPostsMarshal != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("||ERROR|| [Handler.GetPostsWithCategory]: json.Marshal(posts) failed: %v\n", errPostsMarshal)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, errRespWrite := w.Write(resp); errRespWrite != nil {
		log.Printf("||ERROR|| [Handler.GetPostsWithCategory]: w.Write(resp) failed: %v\n", errRespWrite)
		return
	}

}

func (h *Handler) GetPostsWithUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	username := mux.Vars(r)["username"]

	posts, errGetPosts := h.useCase.GetPostsWithUser(ctx, username)
	if errGetPosts != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("||ERROR|| [Handler.GetPostsWithUser]: h.useCase.GetPostsWithUser(ctx, username) failed: %v\n", errGetPosts)
		return
	}

	resp, errPostMarshal := json.Marshal(posts)
	if errPostMarshal != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("||ERROR|| [Handler.GetPostsWithUser]: json.Marshal(posts) failed: %v\n", errPostMarshal)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, errRespWrite := w.Write(resp); errRespWrite != nil {
		log.Printf("||ERROR|| [Handler.GetPostsWithUser]: w.Write(resp) failed: %v\n", errRespWrite)
		return
	}
}

func (h *Handler) GetPost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	postID := mux.Vars(r)["postID"]

	post, errGetPosts := h.useCase.GetPost(ctx, postID)
	if errGetPosts != nil {
		log.Printf("||ERROR|| [Handler.GetPost]->%v", errGetPosts)

		objResp := struct {
			Message string `json:"message"`
		}{}
		switch {
		case errors.Is(errGetPosts, interfaces.ErrPostNotExists):
			objResp.Message = interfaces.ErrPostNotExists.Error()
		default:
			objResp.Message = interfaces.ErrUnknownErrorMsg.Error()
		}

		resp, errRespMarshal := json.Marshal(objResp)
		if errRespMarshal != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("||ERROR|| [Handler.GetPost]: json.Marshal(objResp) failed: %v\n", errRespMarshal)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		if _, errRespWrite := w.Write(resp); errRespWrite != nil {
			log.Printf("||ERROR|| [Handler.GetPost]: w.Write(resp) failed: %v\n", errRespWrite)
			return
		}

		return
	}

	resp, errPostMarshal := json.Marshal(post)
	if errPostMarshal != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("||ERROR|| [Handler.GetPost]: json.Marshal(posts) failed: %v\n", errPostMarshal)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, errRespWrite := w.Write(resp); errRespWrite != nil {
		log.Printf("||ERROR|| [Handler.GetPost]: w.Write(resp) failed: %v\n", errRespWrite)
		return
	}
}

func (h *Handler) DeletePost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	rawSession := r.Context().Value(auth.SessKey)
	if rawSession == nil {
		w.WriteHeader(http.StatusUnauthorized)
		log.Println("||ERROR|| [Handler.DeletePost]: unauthorized")
		return
	}
	session, ok := rawSession.(auth.Session)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("||ERROR|| [Handler.DeletePost]: type assertion rawSession.(auth.Session) failed")
		return
	}
	username := session.Username

	postID := mux.Vars(r)["postID"]

	errDeletePost := h.useCase.DeletePost(ctx, username, postID)
	if errDeletePost != nil {
		log.Printf("||ERROR|| [Handler.DeletePost]->%v", errDeletePost)

		objResp := struct {
			Message string `json:"message"`
		}{}
		switch {
		case errors.Is(errDeletePost, interfaces.ErrAccessDenied):
			w.WriteHeader(http.StatusUnauthorized)
			return
		case errors.Is(errDeletePost, interfaces.ErrPostNotExists):
			objResp.Message = errPostNotFoundMsg

		default:
			objResp.Message = errUnknownErrorMsg

		}
		resp, errRespMarshal := json.Marshal(objResp)
		if errRespMarshal != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("||ERROR|| [Handler.DeletePost]: json.Marshal(objResp) failed: %v\n", errRespMarshal)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		if _, errRespWrite := w.Write(resp); errRespWrite != nil {
			log.Printf("||ERROR|| [Handler.DeletePost]: w.Write(resp) failed: %v\n", errRespWrite)
			return
		}

		return
	}

	objResponse := struct {
		Message string
	}{
		Message: successMsg,
	}

	resp, errRespMarshal := json.Marshal(objResponse)
	if errRespMarshal != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("||ERROR|| [Handler.DeletePost]: json.Marshal(objResponse) failed: %v\n", errRespMarshal)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, errRespWrite := w.Write(resp); errRespWrite != nil {
		log.Printf("||ERROR|| [Handler.DeletePost]: w.Write(resp) failed: %v\n", errRespWrite)
		return
	}
}

func (h *Handler) Upvote(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	val := ctx.Value(auth.SessKey)
	if val == nil {
		w.WriteHeader(http.StatusUnauthorized)
		log.Println("||ERROR|| [Handler.Upvote]: unauthorized")
		return
	}
	session, ok := val.(auth.Session)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(" ||ERROR|| [Handler.Upvote]: type assertion rawSession.(auth.Session) failed")
		return
	}

	postID := mux.Vars(r)["postID"]

	post, errUpvote := h.useCase.Upvote(ctx, session.ID, postID)

	if errUpvote != nil {
		log.Printf("||ERROR|| [Handler.Upvote]->%v", errUpvote)

		objResp := struct {
			Message string `json:"message"`
		}{}
		switch {
		case errors.Is(errUpvote, interfaces.ErrPostNotExists):
			objResp.Message = errPostNotFoundMsg
		case errors.Is(errUpvote, entity.ErrAlreadyUpvote):
			objResp.Message = errAlreadyUpvoteMsg
		default:
			objResp.Message = errUnknownErrorMsg
		}

		resp, errRespMarshal := json.Marshal(objResp)
		if errRespMarshal != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("||ERROR|| [Handler.Upvote]: json.Marshal(objResp) failed: %v\n", errRespMarshal)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		if _, errRespWrite := w.Write(resp); errRespWrite != nil {
			log.Printf("||ERROR|| [Handler.Upvote]: w.Write(resp) failed: %v\n", errRespWrite)
			return
		}

		return
	}

	resp, errPostMarshal := json.Marshal(post)
	if errPostMarshal != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("||ERROR|| [Handler.Upvote]: json.Marshal(post) failed: %v\n", errPostMarshal)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, errRespWrite := w.Write(resp); errRespWrite != nil {
		log.Printf("||ERROR|| [Handler.Upvote]: w.Write(resp) failed: %v\n", errRespWrite)
		return
	}
}

func (h *Handler) Downvote(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	val := ctx.Value(auth.SessKey)
	if val == nil {
		w.WriteHeader(http.StatusUnauthorized)
		log.Println("||ERROR|| [Handler.Downvote]: unauthorized")
		return
	}
	session, ok := val.(auth.Session)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("||ERROR|| [Handler.Downvote]: type assertion rawSession.(auth.Session) failed")
		return
	}

	postID := mux.Vars(r)["postID"]

	post, errDownvote := h.useCase.Downvote(ctx, session.ID, postID)

	if errDownvote != nil {
		log.Printf("||ERROR|| [Handler.Downvote]->%v", errDownvote)

		objResp := struct {
			Message string `json:"message"`
		}{}
		switch {
		case errors.Is(errDownvote, interfaces.ErrPostNotExists):
			objResp.Message = errPostNotFoundMsg
		case errors.Is(errDownvote, entity.ErrAlreadyDownvote):
			objResp.Message = errAlreadyDownvoteMsg
		default:
			objResp.Message = errUnknownErrorMsg
		}

		resp, errRespMarshal := json.Marshal(objResp)
		if errRespMarshal != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("||ERROR|| [Handler.Downvote]: json.Marshal(objResp) failed: %v\n", errRespMarshal)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		if _, errRespWrite := w.Write(resp); errRespWrite != nil {
			log.Printf("||ERROR|| [Handler.Downvote]: w.Write(resp) failed: %v\n", errRespWrite)
			return
		}

		return
	}

	resp, errPostMarshal := json.Marshal(post)
	if errPostMarshal != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("||ERROR|| [Handler.Downvote]: json.Marshal(post) failed: %v\n", errPostMarshal)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, errRespWrite := w.Write(resp); errRespWrite != nil {
		log.Printf("||ERROR|| [Handler.Downvote]: w.Write(resp) failed: %v\n", errRespWrite)
		return
	}

}

func (h *Handler) Unvote(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	val := ctx.Value(auth.SessKey)
	if val == nil {
		w.WriteHeader(http.StatusUnauthorized)
		log.Println("||ERROR|| [Handler.Unvote]: unauthorized")
		return
	}
	session, ok := val.(auth.Session)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("||ERROR|| [Handler.Unvote]: type assertion rawSession.(auth.Session) failed")
		return
	}

	postID := mux.Vars(r)["postID"]

	post, errUnvote := h.useCase.Unvote(ctx, session.ID, postID)

	if errUnvote != nil {
		log.Printf("||ERROR|| [Handler.Unvote]->%v", errUnvote)

		objResp := struct {
			Message string `json:"message"`
		}{}
		switch {
		case errors.Is(errUnvote, interfaces.ErrPostNotExists):
			objResp.Message = errPostNotFoundMsg
		case errors.Is(errUnvote, entity.ErrAlreadyUnvote):
			objResp.Message = errAlreadyUnvoteMsg
		default:
			objResp.Message = errUnknownErrorMsg
		}

		resp, errRespMarshal := json.Marshal(objResp)
		if errRespMarshal != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("||ERROR|| [Handler.Unvote]: json.Marshal(objResp) failed: %v\n", errRespMarshal)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		if _, errRespWrite := w.Write(resp); errRespWrite != nil {
			log.Printf("||ERROR|| [Handler.Unvote]: w.Write(resp) failed: %v\n", errRespWrite)
			return
		}

		return
	}

	resp, errPostMarshal := json.Marshal(post)
	if errPostMarshal != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("||ERROR|| [Handler.Unvote]: json.Marshal(post) failed: %v\n", errPostMarshal)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, errRespWrite := w.Write(resp); errRespWrite != nil {
		log.Printf("||ERROR|| [Handler.Unvote]: w.Write(resp) failed: %v\n", errRespWrite)
		return
	}
}
