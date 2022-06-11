package middlewares

import (
	"net/http"

	"github.com/henbk/go-twitter-api/jwt"
)

func ValidateJsonWebToken(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		_, _, _, err := jwt.ValidateToken(token)

		if err != nil {
			http.Error(rw, "Error processing the token! "+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(rw, r)
	}
}
