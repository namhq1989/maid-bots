package route

import (
	"github.com/labstack/echo/v4"
	"github.com/namhq1989/maid-bots/platform/web/handler"
	routevalidation "github.com/namhq1989/maid-bots/platform/web/route/validation"
)

// public ...
func public(e *echo.Echo) {
	var (
		g = e.Group("/p")
		h = handler.Public{}
		v = routevalidation.Public{}
	)

	// Check domain
	g.GET("/check/domain", h.CheckDomain, v.CheckDomain)
}
