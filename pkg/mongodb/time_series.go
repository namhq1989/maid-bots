package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func colTimeSeries() {
	tsHealthCheckRecord()
}

func tsHealthCheckRecord() {
	var (
		ctx = context.Background()
	)

	// find
	cursor, err := db.ListCollections(ctx, bson.M{"name": collectionNames.HealthCheckRecord})
	if err != nil {
		panic(err)
	}
	defer func() { _ = cursor.Close(ctx) }()

	exists := false
	if cursor.Next(ctx) {
		exists = true
	}

	if !exists {
		metaField := "monitor"
		opts := options.CreateCollection().SetTimeSeriesOptions(
			&options.TimeSeriesOptions{
				TimeField: "createdAt",
				MetaField: &metaField,
			},
		)

		if err = db.CreateCollection(ctx, collectionNames.HealthCheckRecord, opts); err != nil {
			panic(err)
		}
	}
}
