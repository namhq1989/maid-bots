package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func Connect(uri, dbName string) {
	var (
		ctx = context.Background()
	)

	// use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}

	// send a ping to confirm a successful connection
	if err = client.Database("admin").RunCommand(ctx, bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}

	fmt.Printf("⚡️ [mongodb]: connected \n")

	db = client.Database(dbName)

	// index
	go colIndexes()
}

// UserCol ...
func UserCol() *mongo.Collection {
	return db.Collection("users")
}
