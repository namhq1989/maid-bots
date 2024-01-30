package routevalidation

import (
	"github.com/gookit/validate"
	"github.com/labstack/echo/v4"
	"github.com/namhq1989/maid-bots/util/echocontext"
	"github.com/namhq1989/maid-bots/util/response"
)

func validateQuery[T any](next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			query T
		)

		if err := c.Bind(&query); err != nil {
			return response.R400(c, err.Error(), echo.Map{})
		}

		if v := validate.Struct(query); !v.Validate() {
			return response.R400(c, v.Errors.One(), echo.Map{})
		}

		// Assign to context
		echocontext.SetQuery(c, query)
		return next(c)
	}
}

func validateBody[T any](next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			body T
		)

		if err := c.Bind(&body); err != nil {
			return response.R400(c, err.Error(), echo.Map{})
		}

		if v := validate.Struct(body); !v.Validate() {
			return response.R400(c, v.Errors.One(), echo.Map{})
		}

		// Assign to context
		echocontext.SetBody(c, body)
		return next(c)
	}
}
