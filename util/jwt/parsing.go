package jwt

import (
	"errors"
	"fmt"

	"github.com/namhq1989/maid-bots/util/appcontext"

	j "github.com/golang-jwt/jwt/v5"
)

func Parsing(ctx *appcontext.AppContext, s string) (*User, error) {
	// parse token
	token, err := j.Parse(s, func(token *j.Token) (interface{}, error) {
		if _, ok := token.Method.(*j.SigningMethodHMAC); !ok {
			err := fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			ctx.Logger.Error("parsing jwt token failed", err, appcontext.Fields{"token": s})
			return nil, err
		}
		return secret, nil
	})

	// return if error
	if err != nil {
		return nil, err
	}

	// claim payload
	if claims, ok := token.Claims.(j.MapClaims); ok {
		user := User{
			ID: claims["id"].(string),
		}
		return &user, nil
	}

	ctx.Logger.Error("token payload structure is not valid", nil, appcontext.Fields{"token": s})
	return nil, errors.New("invalid token")
}
