package middleware

import (
	"log"
	"net/http"
)

func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Method", "GET DELETE PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		log.Println("Cors middleware run......")

		next.ServeHTTP(w, r)
	})
}
