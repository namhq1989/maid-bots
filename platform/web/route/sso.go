package route

import (
	"github.com/labstack/echo/v4"
	"github.com/namhq1989/maid-bots/platform/web/handler"
	routevalidation "github.com/namhq1989/maid-bots/platform/web/route/validation"
)

// sso ...
func sso(e *echo.Echo) {
	var (
		g = e.Group("/sso")
		h = handler.SSO{}
		v = routevalidation.SSO{}
	)

	// Login with Google
	g.POST("/google", h.LoginWithGoogle, v.LoginWithGoogle)

	// Login with GitHub
	g.POST("/github", h.LoginWithGitHub, v.LoginWithGitHub)
}
