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

const (
	MonitorInterval30Seconds = 30
	MonitorInterval60Seconds = 60
)

type Monitor struct {
	ID        primitive.ObjectID `bson:"_id"`
	Owner     primitive.ObjectID `bson:"owner"` // main factor
	Code      string             `bson:"code"`
	Type      MonitorType        `bson:"type"`
	Target    string             `bson:"target"`
	Data      MonitorMetadata    `bson:"data"`
	Interval  int                `json:"interval"`
	CreatedAt time.Time          `bson:"createdAt"`
}

type MonitorMetadata struct {
	Scheme string `bson:"scheme,omitempty"`
}
