package route

import (
	"github.com/labstack/echo/v4"
	routemiddleware "github.com/namhq1989/maid-bots/platform/web/route/middleware"
)

// Init ...
func Init(e *echo.Echo) {
	// middlewares
	e.Use(routemiddleware.CORS())
	e.Use(routemiddleware.RateLimiter())
	e.Use(routemiddleware.SetContext)
	e.Use(routemiddleware.Auth)

	// components
	common(e)
	sso(e)
	monitor(e)
}
