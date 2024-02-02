package sso

import (
	"fmt"
	"io"
	"net/http"

	"github.com/namhq1989/maid-bots/util/appcontext"

	"github.com/goccy/go-json"
)

const googleTokenInfoURL = "https://oauth2.googleapis.com/tokeninfo?id_token=%s"

// GoogleUserData ...
type GoogleUserData struct {
	ID     string `json:"sub"`
	AUD    string `json:"aud"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	Avatar string `json:"picture"`
}

func LoginWithGoogle(ctx *appcontext.AppContext, token string) (*GoogleUserData, error) {
	ctx.AddLogData(appcontext.Fields{"token": token})

	// call api
	url := fmt.Sprintf(googleTokenInfoURL, token)
	r, err := http.Get(fmt.Sprintf(googleTokenInfoURL, token))
	if err != nil {
		ctx.Logger.Error("fetch token info failed", err, appcontext.Fields{"url": url})
		return nil, err
	}

	// parse body
	defer func() { _ = r.Body.Close() }()
	contents, err := io.ReadAll(r.Body)
	if err != nil {
		ctx.Logger.Error("parse response body failed", err, appcontext.Fields{})
		return nil, err
	}

	// map response to struct
	var result GoogleUserData
	if err = json.Unmarshal(contents, &result); err != nil {
		ctx.Logger.Error("map response to struct failed", err, appcontext.Fields{"contents": string(contents)})
		return nil, err
	}

	// verify that the token was issued to this application
	if result.AUD != providers.GoogleClientID {
		err = fmt.Errorf("wrong audience, got %v", result.AUD)
		ctx.Logger.Error(err.Error(), err, appcontext.Fields{"response": result})
		return nil, err
	}

	// return
	return &result, nil
}
