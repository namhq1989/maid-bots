package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/namhq1989/maid-bots/util/response"
)

// Common ...
type Common struct{}

// Ping godoc
// @tags Common
// @summary Ping
// @id common-ping
// @accept json
// @produce json
// @router /ping [get]
func (Common) Ping(c echo.Context) error {
	return response.R200(c, "", echo.Map{})
}
