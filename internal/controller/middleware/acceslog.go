package middleware

import (
	"log"
	"net/http"
	"time"
)

func AccessLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("||INFO|| request log: method - %s remote_addr - %s url - %s time %s\n", r.Method, r.RemoteAddr, r.URL.Path, time.Since(start).String())
	})
}
