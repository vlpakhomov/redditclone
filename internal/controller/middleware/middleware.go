package middleware

import (
	"context"
	"net/http"
	"time"
)

type callTimeKey string

const (
	CallTimeKey callTimeKey = "callTime"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callTime := time.Now()
		ctx := context.WithValue(r.Context(), CallTimeKey, callTime)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
