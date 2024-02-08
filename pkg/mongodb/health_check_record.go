package mongodb

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HealthCheckRecordStatus string

const (
	HealthCheckRecordStatusUp   HealthCheckRecordStatus = "up"
	HealthCheckRecordStatusDown HealthCheckRecordStatus = "down"
)

type HealthCheckRecord struct {
	ID               primitive.ObjectID      `bson:"_id"`
	Owner            primitive.ObjectID      `bson:"owner"`
	Type             MonitorType             `bson:"type"`
	Status           HealthCheckRecordStatus `bson:"status"`
	Code             string                  `bson:"code"`
	Target           string                  `bson:"target"`
	Description      string                  `bson:"description"`
	ResponseTimeInMs int64                   `bson:"responseTimeInMs"`
	CreatedAt        time.Time               `bson:"createdAt"`
}

type HealthCheckRecordMetadata struct {
	Value string `bson:"value"`
}
