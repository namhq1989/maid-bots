package modelmongodb

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CheckDomain struct {
	ID        primitive.ObjectID `bson:"_id"`
	Owner     primitive.ObjectID `bson:"user"`
	Name      string             `bson:"name"`
	Scheme    string             `bson:"scheme"`
	CreatedAt time.Time          `bson:"createdAt"`
}
