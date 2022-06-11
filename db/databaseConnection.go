package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoConnection *mongo.Client = establishDatabaseConnection()

func establishDatabaseConnection() *mongo.Client {

	DatabasePassword := os.Getenv("MONGODBPASS")

	mongoDatabaseUri := fmt.Sprintf("mongodb+srv://admindb:%s@cluster0.jx2gd.mongodb.net/myFirstDatabase?retryWrites=true&w=majority", DatabasePassword)

	var serverAPIOptions *options.ServerAPIOptions = options.ServerAPI(options.ServerAPIVersion1)

	var clientOptions *options.ClientOptions = options.Client().
		ApplyURI(mongoDatabaseUri).
		SetServerAPIOptions(serverAPIOptions)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatalf("Error while trying to connect to MongoDB, error: %v", err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatalf("Error while trying to ping to the database server, error: %v", err)
	}

	fmt.Println("MongoDB connection succeful!")

	return client
}
