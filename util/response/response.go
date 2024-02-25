package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func sendResponse(c echo.Context, httpCode int, message string, data map[string]interface{}) error {
	if data == nil {
		data = echo.Map{}
	}
	return c.JSON(httpCode, echo.Map{
		"data":    data,
		"message": message,
	})
}

// R200 response success
func R200(c echo.Context, message string, data map[string]interface{}) error {
	if message == "" {
		message = "successfully"
	}
	return sendResponse(c, http.StatusOK, message, data)
}

// R400 bad request
func R400(c echo.Context, message string, data map[string]interface{}) error {
	if message == "" {
		message = "bad request"
	}
	return sendResponse(c, http.StatusBadRequest, message, data)
}

// R401 unauthorized
func R401(c echo.Context, message string, data map[string]interface{}) error {
	if message == "" {
		message = "unauthorized"
	}
	return sendResponse(c, http.StatusUnauthorized, message, data)
}

// R403 forbidden
func R403(c echo.Context, message string, data map[string]interface{}) error {
	if message == "" {
		message = "forbidden"
	}
	return sendResponse(c, http.StatusUnauthorized, message, data)
}

// R404 not found
func R404(c echo.Context, message string, data map[string]interface{}) error {
	if message == "" {
		message = "not found"
	}
	return sendResponse(c, http.StatusNotFound, message, data)
}
