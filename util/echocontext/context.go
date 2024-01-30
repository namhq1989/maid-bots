package echocontext

import (
	"github.com/labstack/echo/v4"
	"github.com/namhq1989/maid-bots/util/appcontext"
)

const KeyContext = "ctx"

// GetContext ...
func GetContext(c echo.Context) *appcontext.AppContext {
	return c.Get(KeyContext).(*appcontext.AppContext)
}

// SetContext ...
func SetContext(c echo.Context, value interface{}) {
	c.Set(KeyContext, value)
}
