package sso

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/namhq1989/maid-bots/util/appcontext"

	"github.com/goccy/go-json"
)

const githubExchangeTokenURL = "https://github.com/login/oauth/access_token"
const githubFetchUserURL = "https://api.github.com/user"

type ExchangeTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

// GitHubUserData ...
type GitHubUserData struct {
	Login  string `json:"login"`
	IDRaw  int    `json:"id"`
	ID     string
	Name   string `json:"name"`
	Email  string `json:"email"`
	Avatar string `json:"avatar_url"`
}

func LoginWithGitHub(ctx *appcontext.AppContext, code string) (*GitHubUserData, error) {
	ctx.AddLogData(appcontext.Fields{"code": code})

	token, err := exchangeToken(ctx, code)
	if err != nil {
		return nil, err
	}

	user, err := fetchGitHubUser(ctx, token.AccessToken)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func exchangeToken(ctx *appcontext.AppContext, code string) (*ExchangeTokenResponse, error) {
	// prepare payload and request
	payload, _ := json.Marshal(map[string]string{
		"client_id":     providers.GitHubClientID,
		"client_secret": providers.GitHubClientSecret,
		"code":          code,
	})
	req, _ := http.NewRequest("POST", githubExchangeTokenURL, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// call api
	client := &http.Client{}
	r, err := client.Do(req)
	if err != nil {
		ctx.Logger.Error("exchange token failed", err, appcontext.Fields{"code": code})
		return nil, err
	}
	defer func() { _ = r.Body.Close() }()

	// parse body
	contents, _ := io.ReadAll(r.Body)
	var token ExchangeTokenResponse
	if err = json.Unmarshal(contents, &token); err != nil {
		ctx.Logger.Error("parse data failed", err, appcontext.Fields{"code": code, "body": string(contents)})
		return nil, err
	}

	return &token, nil
}

func fetchGitHubUser(ctx *appcontext.AppContext, accessToken string) (*GitHubUserData, error) {
	// prepare request
	req, _ := http.NewRequest("GET", githubFetchUserURL, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	// call api
	client := &http.Client{}
	r, err := client.Do(req)
	if err != nil {
		ctx.Logger.Error("fetch user failed", err, appcontext.Fields{"accessToken": accessToken})
		return nil, err
	}
	defer func() { _ = r.Body.Close() }()

	// parse body
	contents, _ := io.ReadAll(r.Body)
	var user GitHubUserData
	if err = json.Unmarshal(contents, &user); err != nil || user.IDRaw == 0 {
		ctx.Logger.Error("parse data failed", err, appcontext.Fields{"accessToken": accessToken, "body": string(contents)})
		return nil, errors.New("cannot fetch user data")
	}

	// parse id from int to string
	user.ID = strconv.Itoa(user.IDRaw)

	return &user, nil
}
