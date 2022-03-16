package db

import (
	"context"
	"time"

	"github.com/henbk/go-twitter-api/models"
	"github.com/henbk/go-twitter-api/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertUser(user models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := MongoConnection.Database("go-api")

	userCollection := db.Collection("user")

	user.Password, _ = utils.EncryptPassword(user.Password)

	result, err := userCollection.InsertOne(ctx, user)

	if err != nil {
		return "", false, err
	}

	insertedObjID, _ := result.InsertedID.(primitive.ObjectID)

	return insertedObjID.String(), true, nil
}
