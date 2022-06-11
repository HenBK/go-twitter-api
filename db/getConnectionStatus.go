package db

import "context"

func GetConnectionStatus() bool {
	err := MongoConnection.Ping(context.TODO(), nil)
	return err == nil
}
