package routemiddleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// CORS ...
func CORS() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodOptions, http.MethodPut, http.MethodPatch, http.MethodDelete},
		AllowHeaders:     []string{"*"},
		AllowCredentials: false,
		MaxAge:           600,
	})
}
