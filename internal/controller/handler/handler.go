package handler

import (
	"html/template"
	"log"
	"net/http"

	"gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/internal/interfaces"
)

const (
	errConstructObjectMsg = "construct object error"
	errUnknownErrorMsg    = "unknown error"
	errPostNotFoundMsg    = "post not found"
	errAlreadyUpvoteMsg   = "already upvote"
	errAlreadyDownvoteMsg = "already downvote"
	errAlreadyUnvoteMsg   = "already unvote"
	errInvalidPasswordMsg = "invalid password"
	errUserNotExistsMsg   = "user doesn't exists"
	errUserAlreadyExists  = "user already exists"
	successMsg            = "success"
)

type handlerError struct {
	Location string `json:"location"`
	Param    string `json:"param"`
	Value    string `json:"value"`
	Msg      string `json:"msg"`
}

type Handler struct {
	useCase  interfaces.IUseCase
	template *template.Template
	Auth     interfaces.IAuthManager
}

func NewHandler(useCase interfaces.IUseCase, template *template.Template, auth interfaces.IAuthManager) *Handler {
	return &Handler{
		useCase:  useCase,
		template: template,
		Auth:     auth,
	}
}

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	err := h.template.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("||ERROR|| h.template.ExecuteTemplate(w, name, nil) failed: %v\n", err)
	}
}
