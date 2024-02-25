package routevalidation

import (
	"github.com/labstack/echo/v4"
	routepayload "github.com/namhq1989/maid-bots/platform/web/route/payload"
)

type Monitor struct{}

func (Monitor) List(next echo.HandlerFunc) echo.HandlerFunc {
	return validateQuery[routepayload.MonitorList](next)
}
