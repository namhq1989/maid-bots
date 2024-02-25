package route

import (
	"github.com/labstack/echo/v4"
	"github.com/namhq1989/maid-bots/platform/web/handler"
	routemiddleware "github.com/namhq1989/maid-bots/platform/web/route/middleware"
	routevalidation "github.com/namhq1989/maid-bots/platform/web/route/validation"
)

func monitor(e *echo.Echo) {
	var (
		g = e.Group("/monitors")
		h = handler.Monitor{}
		v = routevalidation.Monitor{}
	)

	// List
	g.GET("/list", h.List, v.List, routemiddleware.LoginRequired)
}
