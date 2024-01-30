package jwt

import (
	"errors"
)

var (
	secret              []byte
	expireTimeInSeconds = 2592000 // 30 days
)

type User struct {
	ID string
}

func Init(s string) {
	if s == "" {
		panic(errors.New("invalid jwt secret value"))
	}

	secret = []byte(s)
}
