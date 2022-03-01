package db

import (
	ctx "context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoConnection = establishDatabaseConnection()

var serverAPIOptions = options.ServerAPI(options.ServerAPIVersion1)

var clientOptions = options.Client().
	ApplyURI("mongodb+srv://admindb:zjNTUYnm41cDYTSC@cluster0.jx2gd.mongodb.net/myFirstDatabase?retryWrites=true&w=majority").
	SetServerAPIOptions(serverAPIOptions)

func establishDatabaseConnection() *mongo.Client {
	client, err := mongo.Connect(ctx.TODO(), clientOptions)

	if err != nil {
		log.Fatalf("Error while trying to connect to MongoDB, error: %v", err)
	}

	err = client.Ping(ctx.TODO(), nil)

	if err != nil {
		log.Fatalf("Error while trying to ping to the database server, error: %v", err)
	}

	fmt.Println("MongoDB connection succeful!")

	return client
}

func GetConnectionStatus() bool {
	err := MongoConnection.Ping(ctx.TODO(), nil)

	return err != nil
}
