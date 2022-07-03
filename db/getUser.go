package db

import (
	"context"
	"log"
	"time"

	"github.com/henbk/go-twitter-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUserByID(ID string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	db := MongoConnection.Database("go-api")
	userCollection := db.Collection("user")

	var user models.User

	userID, err := primitive.ObjectIDFromHex(ID)

	if err != nil {
		return user, err
	}

	searchCondition := bson.M{
		"_id": userID,
	}

	err = userCollection.FindOne(ctx, searchCondition).Decode(&user)

	if err != nil {
		log.Printf("failure while trying to find user by ID: %s", err.Error())
		return user, err
	}

	return user, nil
}
