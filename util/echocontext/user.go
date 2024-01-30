package echocontext

import "github.com/labstack/echo/v4"

const KeyUserID = "user_id"

// GetUserID ...
func GetUserID(c echo.Context) string {
	uid := c.Get(KeyUserID)
	if uid == nil {
		return ""
	}
	return uid.(string)
}

// SetUserID ...
func SetUserID(c echo.Context, value interface{}) {
	c.Set(KeyUserID, value)
}
