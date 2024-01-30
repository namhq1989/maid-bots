package routevalidation

import (
	"github.com/labstack/echo/v4"
	routepayload "github.com/namhq1989/maid-bots/platform/web/route/payload"
)

type Public struct{}

// CheckDomain ...
func (Public) CheckDomain(next echo.HandlerFunc) echo.HandlerFunc {
	return validateQuery[routepayload.PublicCheckDomain](next)
}
