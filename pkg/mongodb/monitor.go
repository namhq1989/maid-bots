package mongodb

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MonitorType string

const (
	MonitorTypeDomain MonitorType = "domain"
	MonitorTypeHTTP   MonitorType = "http"
	MonitorTypeTCP    MonitorType = "tcp"
	MonitorTypeICMP   MonitorType = "icmp"
)

type Monitor struct {
	ID        primitive.ObjectID `bson:"_id"`
	Owner     primitive.ObjectID `bson:"user"` // main factor
	Code      string             `bson:"code"`
	Type      string             `bson:"type"`
	Data      MonitorMetadata    `bson:"data"`
	CreatedAt time.Time          `bson:"createdAt"`
}

type MonitorMetadata struct {
	Value  string `bson:"value"`
	Scheme string `bson:"scheme,omitempty"`
}
