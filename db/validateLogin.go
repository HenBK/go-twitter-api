package db

import (
	"github.com/henbk/go-twitter-api/models"
	"golang.org/x/crypto/bcrypt"
)

func ValidateLogin(email string, password string) (models.User, bool, bool) {

	user, userExists, _ := CheckUserExists(email)

	var sucessfulLogin bool

	if !userExists {
		return models.User{}, userExists, sucessfulLogin
	}

	err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	)

	if err != nil {
		return user, userExists, sucessfulLogin
	} else {
		sucessfulLogin = true
	}

	return user, userExists, sucessfulLogin
}
