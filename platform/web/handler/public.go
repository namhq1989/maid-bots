package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/namhq1989/maid-bots/util/response"
)

type Public struct{}

// CheckDomain godoc
// @tags Public
// @summary Check domain
// @id public-check-domain
// @accept json
// @produce json
// @router /p/check/domain [get]
func (Public) CheckDomain(c echo.Context) error {
	// var (
	// 	ctx   = echocontext.GetContext(c)
	// 	query = echocontext.GetQuery(c).(routepayload.PublicCheckDomain)
	// 	s     = monitor.Domain{Name: query.Domain}
	// )
	//
	// response, err := s.Check(ctx)
	// if err != nil {
	// 	return response.R400(c, err.Error(), echo.Map{})
	// }

	return response.R200(c, "", echo.Map{})
}
