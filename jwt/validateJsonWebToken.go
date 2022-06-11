package jwt

import (
	"errors"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/henbk/go-twitter-api/db"
	jwtTypes "github.com/henbk/go-twitter-api/jwt/types"
)

var TokenEmail string
var TokenUserID string

func ValidateToken(tokenString string) (*jwtTypes.Claim, bool, string, error) {
	SignPass := os.Getenv("JWT_SECRET_PASS")
	claims := &jwtTypes.Claim{}

	splitToken := strings.Split(tokenString, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, "", errors.New("badly formated json web token")
	}

	token := strings.TrimSpace(splitToken[1])

	parsedToken, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(SignPass), nil
	})

	if err != nil {
		return claims, false, "", err
	}

	if !parsedToken.Valid {
		return claims, false, "", errors.New("invalid token")
	}

	if _, userExists, _ := db.CheckUserExists(claims.Email); !userExists {
		return claims, false, "", errors.New("the user email found in the token could not be validated")
	}

	TokenEmail = claims.Email
	TokenUserID = claims.ID.Hex()

	return claims, true, claims.ID.Hex(), nil
}
