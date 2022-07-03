package routers

import (
	"encoding/json"
	"net/http"

	"github.com/henbk/go-twitter-api/db"
)

func Profile(rw http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("id")
	// TODO: change to receive "@username" to instead of mongo db id to be closer to real twitter behavior

	if userID == "" {
		http.Error(rw, `the "id" parameter is mandatory`, http.StatusBadRequest)
		return
	}

	user, err := db.GetUserByID(userID)
	user.Password = "" // avoid sending the user encrypted password to the client

	if err != nil {
		http.Error(rw, "error while trying to find the user with the provided ID", http.StatusNotFound)
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)

	json.NewEncoder(rw).Encode(user)

}
