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
	indexMonitor()
	indexHealthCheckRecord()
}

func indexUser() {
	indexes := []mongo.IndexModel{
		{
			Keys: bson.M{"google.id": 1},
		},
		{
			Keys: bson.M{"github.id": 1},
		},
		{
			Keys: bson.M{"platform.telegram": 1},
		},
		{
			Keys: bson.M{"platform.slack": 1},
		},
		{
			Keys: bson.M{"platform.discord": 1},
		},
	}
	processIndex(UserCol(), indexes)
}

func indexMonitor() {
	indexes := []mongo.IndexModel{
		{
			Keys: bson.D{{Key: "owner", Value: 1}, {Key: "code", Value: 1}, {Key: "createdAt", Value: -1}},
		},
		{
			Keys: bson.D{{Key: "owner", Value: 1}, {Key: "type", Value: 1}, {Key: "createdAt", Value: -1}},
		},
		{
			Keys: bson.D{{Key: "interval", Value: 1}, {Key: "createdAt", Value: 1}},
		},
	}
	processIndex(MonitorCol(), indexes)
}

func indexHealthCheckRecord() {
	indexes := []mongo.IndexModel{
		{
			Keys: bson.D{{Key: "owner", Value: 1}, {Key: "type", Value: 1}, {Key: "status", Value: 1}, {Key: "createdAt", Value: -1}},
		},
		{
			Keys: bson.D{{Key: "owner", Value: 1}, {Key: "code", Value: 1}, {Key: "createdAt", Value: -1}},
		},
	}
	processIndex(HealthCheckRecordCol(), indexes)
}

// process ...
func processIndex(col *mongo.Collection, indexes []mongo.IndexModel) {
	opts := options.CreateIndexes().SetMaxTime(time.Minute * 30)
	if _, err := col.Indexes().CreateMany(context.Background(), indexes, opts); err != nil {
		fmt.Printf("Index collection %s err: %v \n", col.Name(), err)
	}
}
