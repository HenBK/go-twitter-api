package db

import (
	"context"
	"time"

	"github.com/henbk/go-twitter-api/models"
	"go.mongodb.org/mongo-driver/bson"
)

func CheckUserExists(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := MongoConnection.Database("go-api")
	userCollection := db.Collection("user")

	matchCondition := bson.M{"email": email}

	var foundUser models.User

	err := userCollection.FindOne(ctx, matchCondition).Decode(&foundUser)
	userID := foundUser.ID.Hex()

	if err != nil {
		return foundUser, false, userID
	}

	return foundUser, true, userID

}
