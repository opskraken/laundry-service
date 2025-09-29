package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		duration := time.Since(time.Now())

		log.Println(r.URL.Path, "->", r.Method, "->", r.Proto, "-> time", duration)

		next.ServeHTTP(w, r)
	})
}
