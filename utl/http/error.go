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
		code = http.StatusInternalServerError
		resp = errorResponse{
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
