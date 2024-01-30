package routemiddleware

import (
	"github.com/labstack/echo/v4"
)

// Auth api before move to handler
func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(c)
		//var (
		//	cfg       = config.GetENV()
		//	authToken = echocontext.GetToken(c)
		//)
		//
		//// parse token
		//id, _ := pstring.ParseJWTToken(authToken, cfg.Authentication.JWTSecret)
		//
		//// set data
		//echocontext.SetCurrentUserID(c, id)
		//
		//// next
		//return next(c)
	}
}

func LoginRequired(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//id := echocontext.GetCurrentUserID(c)
		//if id == "" {
		//	return response.R403(c, echo.Map{}, "")
		//}
		return next(c)
	}
}
