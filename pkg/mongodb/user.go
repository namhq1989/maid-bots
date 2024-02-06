package mongodb

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID             `bson:"_id"`
	Name      string                         `bson:"name"`
	Avatar    string                         `bson:"avatar"`
	Platform  UserPlatform                   `bson:"platform"`
	Google    *UserSocialProviderInformation `bson:"google"`
	GitHub    *UserSocialProviderInformation `bson:"github"`
	CreatedAt time.Time                      `bson:"createdAt"`
}

type UserSocialProviderInformation struct {
	ID     string `bson:"id"`
	Name   string `bson:"name"`
	Email  string `bson:"email"`
	Avatar string `bson:"avatar"`
}

type UserPlatform struct {
	Telegram string `bson:"telegram"`
	Slack    string `bson:"slack"`
	Discord  string `bson:"discord"`
}
