package handler

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/internal/entity"
	"gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/internal/interfaces"
)

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	body, errBodyReadAll := io.ReadAll(r.Body)
	if errBodyReadAll != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("||ERROR|| [Handler.SignUp]: io.ReadAll(req.Body) failed: %v\n", errBodyReadAll)
		return
	}

	objReq := entity.User{}

	errReqUnmarshal := json.Unmarshal(body, &objReq)
	if errReqUnmarshal != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("||ERROR|| [Handler.SignUp]: json.Unmarshal(req.Body, &objReq) failed: %v\n", errReqUnmarshal)
		return
	}

	userExtend, errSignUp := h.useCase.SignUp(ctx, objReq)
	if errSignUp != nil {
		log.Printf("||ERROR|| [Handler.SignUp]->%v\n", errSignUp)

		objResp := struct {
			Errors []handlerError `json:"errors"`
		}{}
		objRespInvalid := struct {
			Message string `json:"message"`
		}{}
		statusCode := 0
		switch {
		case errors.Is(errSignUp, interfaces.ErrUserExists):
			objResp.Errors = append(objResp.Errors, handlerError{
				Location: "body",
				Param:    "username",
				Value:    objReq.Username,
				Msg:      errUserAlreadyExists,
			})
			statusCode = http.StatusUnprocessableEntity
		case errors.Is(errSignUp, interfaces.ErrConstructObject):
			objRespInvalid.Message = errConstructObjectMsg
			statusCode = http.StatusInternalServerError
		default:
			objRespInvalid.Message = errUnknownErrorMsg
			statusCode = http.StatusInternalServerError
		}

		var resp []byte
		if len(objResp.Errors) > 0 {
			res, errRespMarshal := json.Marshal(objResp)
			if errRespMarshal != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Printf("||ERROR|| [Handler.SignUp]: json.Marshal(objResp) failed: %v\n", errRespMarshal)
				return
			}
			resp = res
		} else {
			res, errRespMarshal := json.Marshal(objRespInvalid)
			if errRespMarshal != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Printf("||ERROR|| [Handler.SignUp]: json.Marshal(objResp) failed: %v\n", errRespMarshal)
				return
			}
			resp = res
		}

		w.WriteHeader(statusCode)
		if _, errRespWrite := w.Write(resp); errRespWrite != nil {
			log.Printf("||ERROR|| [Handler.SignUp]: w.Write(resp) failed: %v\n", errRespWrite)
			return
		}

		return
	}

	accessToken, errCreateSession := h.Auth.CreateSession(userExtend)
	if errCreateSession != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("||ERROR|| [Handler.SignUp]: h.Auth.CreateToken(userExtend) failed: %v\n", errCreateSession)
		return
	}

	objResp := struct {
		Token string `json:"token"`
	}{
		Token: accessToken,
	}

	resp, errRespMarshal := json.Marshal(objResp)
	if errRespMarshal != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("||ERROR|| [Handler.SignUp]: json.Marshal(objResp) failed: %v\n", errRespMarshal)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, errRespWrite := w.Write(resp); errRespWrite != nil {
		log.Printf("||ERROR|| [Handler.SignUp]: w.Write(resp) failed: %v\n", errRespWrite)
		return
	}
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	body, errBodyReadAll := io.ReadAll(r.Body)
	if errBodyReadAll != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("||ERROR|| [Handler.Login]: io.ReadAll(req.Body) failed: %v\n", errBodyReadAll)
		return
	}

	objReq := entity.User{}

	errReqUnmarshal := json.Unmarshal(body, &objReq)
	if errReqUnmarshal != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("||ERROR|| [Handler.Login]: json.Unmarshal(req.Body, &objReq) failed: %v\n", errReqUnmarshal)
		return
	}

	userExtend, errLogin := h.useCase.Login(ctx, objReq.Username, objReq.Password)
	if errLogin != nil {
		log.Printf("||ERROR|| [Handler.Login]->%v\n", errLogin)

		objResp := struct {
			Message string `json:"message"`
		}{}
		statusCode := 0
		switch {
		case errors.Is(errLogin, interfaces.ErrUserNotExists):
			objResp.Message = errUserNotExistsMsg
			statusCode = http.StatusBadRequest
		case errors.Is(errLogin, interfaces.ErrInvalidPassword):
			objResp.Message = errInvalidPasswordMsg
			statusCode = http.StatusUnauthorized
		default:
			objResp.Message = errUnknownErrorMsg
			statusCode = http.StatusInternalServerError
		}

		resp, errRespMarshal := json.Marshal(objResp)
		if errRespMarshal != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("||ERROR|| [Handler.Login]: json.Marshal(objResp) failed: %v\n", errRespMarshal)
			return
		}

		w.WriteHeader(statusCode)
		if _, errRespWrite := w.Write(resp); errRespWrite != nil {
			log.Printf("||ERROR|| [Handler.Login]: w.Write(resp) failed: %v\n", errRespWrite)
			return
		}

		return
	}

	accessToken, errCreateToken := h.Auth.CreateSession(userExtend)
	if errCreateToken != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("||ERROR|| [Handler.Login]: h.Auth.CreateToken(userExtend) failed: %v\n", errCreateToken)
		return
	}

	objResp := struct {
		Token string `json:"token"`
	}{
		Token: accessToken,
	}

	resp, errRespMarshal := json.Marshal(objResp)
	if errRespMarshal != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("||ERROR|| [Handler.Login]: json.Marshal(objResp) failed: %v\n", errRespMarshal)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, errRespWrite := w.Write(resp); errRespWrite != nil {
		log.Printf("||ERROR|| [Handler.Login]: w.Write(resp) failed: %v\n", errRespWrite)
		return
	}
}
