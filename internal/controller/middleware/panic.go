package middleware

import (
	"log"
	"net/http"
)

func Panic(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if errRecover := recover(); errRecover != nil {
				log.Printf("||FATAL|| [middleware.Panic]: occured panic: %v", errRecover)
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()
		h.ServeHTTP(w, r)
	})
}
