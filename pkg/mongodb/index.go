package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Indexes ...
func colIndexes() {
	indexUser()
}

func indexUser() {
	indexes := []mongo.IndexModel{
		{
			Keys: bson.M{"google.id": 1},
		},
		{
			Keys: bson.M{"github.id": 1},
		},
	}
	processIndex(UserCol(), indexes)
}

// process ...
func processIndex(col *mongo.Collection, indexes []mongo.IndexModel) {
	opts := options.CreateIndexes().SetMaxTime(time.Minute * 30)
	if _, err := col.Indexes().CreateMany(context.Background(), indexes, opts); err != nil {
		fmt.Printf("Index collection %s err: %v \n", col.Name(), err)
	}
}
