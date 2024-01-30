package routevalidation

import (
	"github.com/labstack/echo/v4"
	routepayload "github.com/namhq1989/maid-bots/platform/web/route/payload"
)

type SSO struct{}

// LoginWithGoogle ...
func (SSO) LoginWithGoogle(next echo.HandlerFunc) echo.HandlerFunc {
	return validateBody[routepayload.SSOLoginWithGoogleBody](next)
}

// LoginWithGitHub ...
func (SSO) LoginWithGitHub(next echo.HandlerFunc) echo.HandlerFunc {
	return validateBody[routepayload.SSOLoginWithGitHubBody](next)
}
