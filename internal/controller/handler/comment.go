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

func (h *Handler) AddComment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	rawSession := r.Context().Value(auth.SessKey)
	if rawSession == nil {
		w.WriteHeader(http.StatusUnauthorized)
		log.Println("[||ERROR|| Handler.AddComment]: unauthorized")
		return
	}
	session, ok := rawSession.(auth.Session)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("||ERROR|| [Handler.AddComment]: type assertion rawSession.(auth.Session) failed")
		return
	}
	username := session.Username
	id := session.ID

	rawCallTime := r.Context().Value(middleware.CallTimeKey)
	if rawSession == nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("[||ERROR|| Handler.AddComment]: r.Context().Value(middleware.CallTimeKey) failed")
		return
	}
	callTime, ok := rawCallTime.(time.Time)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("||ERROR|| [Handler.AddComment]: type assertion rawCallTime.(time.Time) failed")
		return
	}

	postID := mux.Vars(r)["postID"]

	body, errBodyReadAll := io.ReadAll(r.Body)
	if errBodyReadAll != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("||ERROR|| [Handler.AddComment]: io.ReadAll(req.Body) failed: %v\n", errBodyReadAll)
		return
	}

	objReq := struct {
		Comment string `json:"comment"`
	}{}

	errReqUnmarshal := json.Unmarshal(body, &objReq)
	if errReqUnmarshal != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("||ERROR|| [Handler.AddComment]: json.Unmarshal(req.Body, &objReq) failed: %v\n", errReqUnmarshal)
		return
	}

	author := entity.Author{
		Username: username,
		ID:       id,
	}
	comment := entity.Comment{
		Author:      author,
		CreatedTime: callTime,
		Body:        objReq.Comment,
	}

	post, errAddComment := h.useCase.AddComment(ctx, postID, comment)
	if errAddComment != nil {
		log.Printf("||ERROR|| [Handler.AddComment]->%v", errAddComment)

		objResp := struct {
			Message string `json:"message"`
		}{}
		switch {
		case errors.Is(errAddComment, interfaces.ErrConstructObject):
			objResp.Message = errConstructObjectMsg

		default:
			objResp.Message = errUnknownErrorMsg
		}
		resp, errRespMarshal := json.Marshal(objResp)
		if errRespMarshal != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("||ERROR|| [Handler.DeleteComment]: json.Marshal(objResp) failed: %v\n", errRespMarshal)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		if _, errRespWrite := w.Write(resp); errRespWrite != nil {
			log.Printf("||ERROR|| [Handler.DeleteComment]: w.Write(resp) failed: %v\n", errRespWrite)
			return
		}

		return
	}

	resp, errPostMarshal := json.Marshal(post)
	if errPostMarshal != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("||ERROR|| [Handler.AddComment]: json.Marshal(post) failed: %v\n", errPostMarshal)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, errRespWrite := w.Write(resp); errRespWrite != nil {
		log.Printf("||ERROR|| [Handler.AddComment]: w.Write(resp) failed: %v\n", errRespWrite)
		return
	}

}

func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	rawSession := r.Context().Value(auth.SessKey)
	if rawSession == nil {
		w.WriteHeader(http.StatusUnauthorized)
		log.Println("||ERROR|| [Handler.DeleteComment]: unauthorized")
		return
	}
	session, ok := rawSession.(auth.Session)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("||ERROR|| [Handler.DeleteComment]: type assertion rawSession.(auth.Session) failed")
		return
	}
	username := session.Username

	postID := mux.Vars(r)["postID"]
	commentID := mux.Vars(r)["commentID"]

	post, errDeleteComment := h.useCase.DeleteComment(ctx, username, postID, commentID)
	if errDeleteComment != nil {
		log.Printf("||ERROR|| [Handler.DeleteComment]->%v", errDeleteComment)

		switch {
		case errors.Is(errDeleteComment, interfaces.ErrAccessDenied):
			w.WriteHeader(http.StatusUnauthorized)
			return
		default:
			objResp := struct {
				Message string `json:"message"`
			}{
				Message: errUnknownErrorMsg,
			}

			resp, errRespMarshal := json.Marshal(objResp)
			if errRespMarshal != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Printf("||ERROR|| [Handler.DeleteComment]: json.Marshal(objResp) failed: %v\n", errRespMarshal)
				return
			}

			w.WriteHeader(http.StatusInternalServerError)
			if _, errRespWrite := w.Write(resp); errRespWrite != nil {
				log.Printf("||ERROR|| [Handler.DeleteComment]: w.Write(resp) failed: %v\n", errRespWrite)
				return
			}

			return
		}
	}

	resp, errPostMarshal := json.Marshal(post)
	if errPostMarshal != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("||ERROR|| [Handler.DeleteComment]: json.Marshal(post) failed: %v\n", errPostMarshal)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, errRespWrite := w.Write(resp); errRespWrite != nil {
		log.Printf("||ERROR|| [Handler.DeleteComment]: w.Write(resp) failed: %v\n", errRespWrite)
		return
	}

}
