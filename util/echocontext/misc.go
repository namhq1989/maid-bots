package echocontext

import (
	"strings"

	"github.com/labstack/echo/v4"
)

// GetAuthorizationToken ...
func GetAuthorizationToken(c echo.Context) string {
	token := c.Request().Header.Get(echo.HeaderAuthorization)

	split := strings.Split(token, " ")
	if len(split) == 2 {
		return split[1]
	}

	return ""
}

// GetIP ...
func GetIP(c echo.Context) string {
	return c.RealIP()
}
