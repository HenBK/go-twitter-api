package db

import (
	"github.com/henbk/go-twitter-api/models"
	"golang.org/x/crypto/bcrypt"
)

func ValidateLogin(email string, password string) (models.User, bool) {

	user, userExists, _ := CheckUserExists(email)

	if !userExists {
		return models.User{}, false
	}

	err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	)

	if err != nil {
		return user, false
	}

	return user, true
}
