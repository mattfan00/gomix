package transport

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mattfan00/gomite/svc/alpha"
)

type routeHandler struct {
	svc alpha.Service
}

func NewHTTP(e *echo.Echo, svc alpha.Service) {
	a := e.Group("/alpha")

	r := routeHandler{
		svc: svc,
	}

	a.GET("/hey", r.get)
}

func (r routeHandler) get(c echo.Context) error {
	str := r.svc.Hello()

	return c.JSON(http.StatusOK, str)
}
