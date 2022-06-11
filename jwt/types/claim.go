package jwt

import (
	jwt "github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Claim struct {
	Email string             `json:"email"`
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	jwt.StandardClaims
}
