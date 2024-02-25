package route

import (
	"github.com/labstack/echo/v4"
	"github.com/namhq1989/maid-bots/config"
	"github.com/namhq1989/maid-bots/platform/web/handler"
	"github.com/namhq1989/maid-bots/util/echocontext"
	"github.com/namhq1989/maid-bots/util/jwt"
	"github.com/namhq1989/maid-bots/util/response"
)

// common ...
func common(e *echo.Echo) {
	var (
		g = e.Group("")
		h = handler.Common{}
	)

	// Ping ...
	g.GET("/ping", h.Ping)

	g.GET("/token/:id", func(c echo.Context) error {
		// disable if env is "release"
		if config.GetENV().Environment == "release" {
			return response.R404(c, "", echo.Map{})
		}

		var (
			ctx = echocontext.GetContext(c)
			id  = c.Param("id")
		)

		token, err := jwt.Signing(ctx, jwt.User{ID: id})
		if err != nil {
			return response.R400(c, err.Error(), echo.Map{})
		}

		return response.R200(c, "", echo.Map{
			"token": token,
		})
	})
}
