package route

import (
	"github.com/labstack/echo/v4"
	"github.com/namhq1989/maid-bots/platform/web/handler"
)

// common ...
func common(e *echo.Echo) {
	var (
		g = e.Group("")
		h = handler.Common{}
	)

	// Ping ...
	g.GET("/ping", h.Ping)
}
