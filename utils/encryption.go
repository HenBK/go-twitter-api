package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) (string, error) {
	cost := 8
	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)

	return string(passwordBytes), err
}
