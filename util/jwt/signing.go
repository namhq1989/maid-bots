package jwt

import (
	"time"

	"github.com/namhq1989/maid-bots/util/appcontext"

	j "github.com/golang-jwt/jwt/v5"
)

func Signing(ctx *appcontext.AppContext, payload User) (string, error) {
	token := j.NewWithClaims(j.SigningMethodHS256, j.MapClaims{
		"id":  payload.ID,
		"exp": time.Now().Add(time.Second * time.Duration(expireTimeInSeconds)),
	})

	// sign and get the complete encoded token as a string
	s, err := token.SignedString(secret)
	if err != nil {
		ctx.Logger.Error("cannot sign jwt payload", err, appcontext.Fields{"payload": payload})
	}

	return s, err
}
