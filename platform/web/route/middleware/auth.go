package routemiddleware

import (
	"github.com/labstack/echo/v4"
	"github.com/namhq1989/maid-bots/util/echocontext"
	"github.com/namhq1989/maid-bots/util/jwt"
	"github.com/namhq1989/maid-bots/util/response"
)

// Auth api before move to handler
func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(c)
		// var (
		//	cfg       = config.GetENV()
		//	authToken = echocontext.GetToken(c)
		// )
		//
		// // parse token
		// id, _ := pstring.ParseJWTToken(authToken, cfg.Authentication.JWTSecret)
		//
		// // set data
		// echocontext.SetCurrentUserID(c, id)
		//
		// // next
		// return next(c)
	}
}

func LoginRequired(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			ctx       = echocontext.GetContext(c)
			authToken = echocontext.GetAuthorizationToken(c)
		)

		// return if auth token is empty
		if authToken == "" {
			return response.R401(c, "", echo.Map{})
		}

		// parse user
		user, err := jwt.Parsing(ctx, authToken)

		// return if error
		if err != nil {
			return response.R401(c, err.Error(), echo.Map{})
		}

		// return if user not found
		if user == nil {
			return response.R401(c, "", echo.Map{})
		}

		// set user id
		echocontext.SetUserID(c, user.ID)

		return next(c)
	}
}
