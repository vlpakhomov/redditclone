package middleware

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/internal/controller/auth"
	"gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/internal/interfaces"
)

func Auth(a interfaces.IAuthManager, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rawAccessToken := r.Header.Get(auth.AuthHeader)

		if rawAccessToken != "" {
			accesToken := strings.TrimPrefix(rawAccessToken, auth.TokenPrefix)
			session, errParseToken := a.ParseToken(accesToken)
			if errParseToken != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Printf("[middleware.Auth]: a.ParseToken(accesToken) failed: %v\n", errParseToken)
			}
			errGetSession := a.GetSession(*session)
			if errGetSession != nil {
				log.Printf("[middleware.Auth]: a.ParseToken(accesToken) failed: %v\n", errGetSession)

				switch {
				case errors.Is(errGetSession, auth.ErrInactiveToken):
					w.WriteHeader(http.StatusUnauthorized)
					return
				case errors.Is(errGetSession, auth.ErrSessionUnmarshal):
					w.WriteHeader(http.StatusInternalServerError)
					return
				default:
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
			}

			ctx := auth.ContextWithSession(r.Context(), *session)

			h.ServeHTTP(w, r.WithContext(ctx))

		} else {
			h.ServeHTTP(w, r)
		}
	})
}
