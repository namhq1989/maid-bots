package modelmongodb

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HealthCheckRecordType string

const (
	HealthCheckRecordTypeDomain   HealthCheckRecordType = "domain"
	HealthCheckRecordTypeIPPort   HealthCheckRecordType = "ip_port"
	HealthCheckRecordTypeAPI      HealthCheckRecordType = "api"
	HealthCheckRecordTypePingHost HealthCheckRecordType = "ping_host"
)

type HealthCheckRecordStatus string

const (
	HealthCheckRecordStatusUp   HealthCheckRecordStatus = "up"
	HealthCheckRecordStatusDown HealthCheckRecordStatus = "down"
)

type HealthCheckRecord struct {
	ID               primitive.ObjectID             `bson:"_id"`
	Owner            primitive.ObjectID             `bson:"owner"`
	Type             HealthCheckRecordType          `bson:"type"`
	Status           HealthCheckRecordStatus        `bson:"status"`
	Description      string                         `bson:"description"`
	ResponseTimeInMs int                            `bson:"responseTimeInMs"`
	CreatedAt        time.Time                      `bson:"createdAt"`
	Domain           *HealthCheckRecordDomainData   `bson:"domain"`
	IPPort           *HealthCheckRecordIPPortData   `bson:"ipPort"`
	API              *HealthCheckRecordAPIData      `bson:"api"`
	PingHost         *HealthCheckRecordPingHostData `bson:"pingHost"`
}

type HealthCheckRecordDomainData struct {
	Code   string `bson:"code"`
	Domain string `bson:"domain"`
}

type HealthCheckRecordIPPortData struct {
	Code string `bson:"code"`
	IP   string `bson:"ip"`
	Port int    `bson:"port"`
}

type HealthCheckRecordAPIData struct {
	Code string `bson:"code"`
	API  string `bson:"api"`
}

type HealthCheckRecordPingHostData struct {
	Code string `bson:"code"`
	Host string `bson:"host"`
}
