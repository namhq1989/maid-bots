package routemiddleware

import (
	"github.com/labstack/echo/v4"
	"github.com/namhq1989/maid-bots/util/appcontext"
	"github.com/namhq1989/maid-bots/util/echocontext"
)

func SetContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		appCtx := appcontext.New(c.Request().Context())
		echocontext.SetContext(c, appCtx)

		return next(c)
	}
}
