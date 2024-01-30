package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/namhq1989/maid-bots/internal/service"
	routepayload "github.com/namhq1989/maid-bots/platform/web/route/payload"
	"github.com/namhq1989/maid-bots/util/echocontext"
	"github.com/namhq1989/maid-bots/util/response"
)

// SSO ...
type SSO struct{}

// LoginWithGoogle godoc
// @tags SSO
// @summary Login with Google
// @id sso-login-with-google
// @accept json
// @produce json
// @router /auth/google [post]
func (SSO) LoginWithGoogle(c echo.Context) error {
	var (
		ctx  = echocontext.GetContext(c)
		body = echocontext.GetBody(c).(routepayload.SSOLoginWithGoogleBody)
		s    = service.SSO{}
	)

	token, err := s.LoginWithGoogle(ctx, body.Token)
	if err != nil {
		return response.R400(c, err.Error(), echo.Map{})
	}

	return response.R200(c, "", echo.Map{
		"token": token,
	})
}

// LoginWithGitHub godoc
// @tags SSO
// @summary Login with GitHub
// @id sso-login-with-github
// @accept json
// @produce json
// @router /auth/github [post]
func (SSO) LoginWithGitHub(c echo.Context) error {
	var (
		ctx  = echocontext.GetContext(c)
		body = echocontext.GetBody(c).(routepayload.SSOLoginWithGitHubBody)
		s    = service.SSO{}
	)

	token, err := s.LoginWithGitHub(ctx, body.Code)
	if err != nil {
		return response.R400(c, err.Error(), echo.Map{})
	}

	return response.R200(c, "", echo.Map{
		"token": token,
	})
}
