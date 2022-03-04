package middlewares

import (
	"net/http"

	"github.com/henbk/go-twitter-api/db"
)

func DatabaseConnectionCheck(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if !db.GetConnectionStatus() {
			http.Error(rw, "Lost connection to database", 500)
			return
		}

		next.ServeHTTP(rw, r)
	}
}
