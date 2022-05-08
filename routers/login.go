package routers

import (
	"encoding/json"
	"net/http"

	"github.com/henbk/go-twitter-api/db"
	"github.com/henbk/go-twitter-api/jwt"
	"github.com/henbk/go-twitter-api/models"
	"github.com/henbk/go-twitter-api/responses"
)

func Login(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("content-type", "application/json")
	// rw.Header().Set("content-type", "application/json")

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(rw, "Incorrect user password/email.", http.StatusBadRequest)
		return
	}

	if user.Email == "" {
		http.Error(rw, "The user email is required to login.", http.StatusBadRequest)
	}

	if user.Password == "" {
		http.Error(rw, "The user password is required to login.", http.StatusBadRequest)
	}

	mongoUser, userExistsInMongo, sucessfulLogin := db.ValidateLogin(user.Email, user.Password)

	if !userExistsInMongo {
		http.Error(rw, "The user does not exist in the database.", http.StatusBadRequest)
		return
	}

	if !sucessfulLogin {
		http.Error(rw, "Incorrect user password/email.", http.StatusBadRequest)
		return
	}

	jwtKey, err := jwt.GenerateJsonWebToken(mongoUser)

	if err != nil {
		http.Error(rw, "Error while generating authentication token.", http.StatusBadRequest)
		return
	}

	response := &responses.LoginResponse{
		Token: jwtKey,
	}

	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(response)

}
