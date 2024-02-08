package mongodb

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID             `bson:"_id"`
	Name      string                         `bson:"name"`
	Username  string                         `bson:"username"`
	Avatar    string                         `bson:"avatar"`
	Google    *UserSocialProviderInformation `bson:"google"`
	GitHub    *UserSocialProviderInformation `bson:"github"`
	Telegram  *UserPlatform                  `bson:"telegram"`
	Slack     *UserPlatform                  `bson:"slack"`
	Discord   *UserPlatform                  `bson:"discord"`
	CreatedAt time.Time                      `bson:"createdAt"`
}

type UserSocialProviderInformation struct {
	ID     string `bson:"id"`
	Name   string `bson:"name"`
	Email  string `bson:"email"`
	Avatar string `bson:"avatar"`
}

type UserPlatform struct {
	UserID string `bson:"userId"`
	RoomID string `bson:"roomId"`
}
