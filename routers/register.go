package routers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/henbk/go-twitter-api/db"
	"github.com/henbk/go-twitter-api/models"
)

func Register(rw http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		fmt.Printf("Json object was not decoded succesfully %v", err)
		http.Error(rw, "Json object was not decoded succesfully", 400)
		return
	}

	// Validations

	if len(user.Email) == 0 {
		http.Error(rw, "User email is required", 400)
		return
	}

	if len(user.Password) < 6 {
		http.Error(rw, "The user password must be at least 6 characters long", 400)
		return
	}

	_, userExists, _ := db.CheckUserExists(user.Email)

	if userExists {
		http.Error(rw, "There is an already created user with this emai", 400)
		return
	}

	// Registration

	_, status, err := db.InsertUser(user)

	if err != nil {
		http.Error(rw, "An error ocurred while inserting the new user into the database. "+err.Error(), 500)
		return
	}

	if !status {
		http.Error(rw, "An error ocurred while inserting the new user into the database.", 500)
		return
	}

	rw.WriteHeader(http.StatusCreated)

}
