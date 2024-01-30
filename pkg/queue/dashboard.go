package queue

import (
	"github.com/hibiken/asynqmon"
)

func Dashboard(redisURL string) *asynqmon.HTTPHandler {
	// redis connection
	redisConn := getRedisConnFromURL(redisURL)

	return asynqmon.New(asynqmon.Options{
		RootPath:     "/q",
		RedisConnOpt: redisConn,
	})
}
