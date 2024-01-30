package echocontext

import "github.com/labstack/echo/v4"

const KeyQuery = "query"
const KeyBody = "body"

// GetQuery ...
func GetQuery(c echo.Context) interface{} {
	return c.Get(KeyQuery)
}

// SetQuery ...
func SetQuery(c echo.Context, value interface{}) {
	c.Set(KeyQuery, value)
}

// GetBody ...
func GetBody(c echo.Context) interface{} {
	return c.Get(KeyBody)
}

// SetBody ...
func SetBody(c echo.Context, value interface{}) {
	c.Set(KeyBody, value)
}
