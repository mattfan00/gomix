package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type errorResponse struct {
	Message interface{} `json:"message"`
}

func customHTTPErrorHandler(err error, c echo.Context) {
	var (
		code int           = http.StatusInternalServerError
		resp errorResponse = errorResponse{
			Message: err.Error(),
		}
	)

	switch customError := err.(type) {
	case *echo.HTTPError:
		code = customError.Code
		resp = errorResponse{
			Message: customError.Message,
		}
	}

	c.JSON(code, resp)
}

func ErrNotFound() *echo.HTTPError {
	return echo.NewHTTPError(http.StatusNotFound, "Not found")
}

func ErrBadRequest(details ...string) *echo.HTTPError {
	msg := "Bad request"
	if len(details) > 0 {
		msg = details[0]
	}

	return echo.NewHTTPError(http.StatusBadRequest, msg)
}

func ErrUnauthorized(details ...string) *echo.HTTPError {
	msg := "Not authorized"
	if len(details) > 0 {
		msg = details[0]
	}

	return echo.NewHTTPError(http.StatusUnauthorized, msg)
}
