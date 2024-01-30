package modelmongodb

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Domain struct {
	ID        primitive.ObjectID `bson:"_id"`
	Owner     primitive.ObjectID `bson:"user"` // main factor
	Code      string             `bson:"code"`
	Root      string             `bson:"root"`
	Scheme    string             `bson:"scheme"`
	CreatedAt time.Time          `bson:"createdAt"`
}
