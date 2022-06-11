package jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/henbk/go-twitter-api/models"
)

func GenerateJsonWebToken(user models.User) (string, error) {
	SignPass := os.Getenv("JWT_SECRET_PASS")

	claims := jwt.MapClaims{
		"email":            user.Email,
		"name":             user.Name,
		"lastname":         user.Lastname,
		"birthDate":        user.BirthDate,
		"bio":              user.Bio,
		"location":         user.Location,
		"website":          user.Website,
		"_id":              user.ID.Hex(),
		"token_expiration": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(SignPass))

	if err != nil {
		return "", err
	}

	return tokenString, nil

}
