package redis

import (
	"context"
	"fmt"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/redis/go-redis/v9"
)

var (
	client *redis.Client

	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

// Connect ...
func Connect(uri string) {
	var (
		ctx    = context.Background()
		opt, _ = redis.ParseURL(uri)
	)

	// new client
	client = redis.NewClient(opt)

	// Ping
	if _, err := client.Ping(ctx).Result(); err != nil {
		panic(err)
	}

	fmt.Printf("⚡️ [redis]: connected \n")
}

// SetKeyValue ...
func SetKeyValue(key string, value interface{}, expiration time.Duration) {
	b, _ := json.Marshal(value)
	client.Set(context.Background(), key, b, expiration)
}

// GetValueByKey ...
func GetValueByKey(key string) string {
	value, _ := client.Get(context.Background(), key).Result()
	return value
}

// DelKey ...
func DelKey(key string) {
	client.Del(context.Background(), key)
}

// GetJSON ...
func GetJSON(key string, result interface{}) (ok bool) {
	v := GetValueByKey(key)
	if v == "" {
		return false
	}
	if err := json.Unmarshal([]byte(v), result); err != nil {
		return false
	}
	return true
}
