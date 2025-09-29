package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/enghasib/laundry_service/utils"
)

func (m *Middlewares) Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Authentication middlewares call.....")
		// header
		AuthenticationHeader := r.Header.Get("Authorization")
		if AuthenticationHeader == "" {
			http.Error(w, "Unauthorized:", http.StatusUnauthorized)
			return
		}

		//split header and grep the token
		headerArr := strings.Split(AuthenticationHeader, " ")
		if len(headerArr) != 2 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		jwt_token := headerArr[1]

		// verify token
		isVerified, err := utils.Verify(jwt_token, m.cnf.JwtSecretKey)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if !isVerified {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
