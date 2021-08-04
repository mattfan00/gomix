package transport

import (
	"net/http"

	"github.com/mattfan00/gomite/auth/pkg/service"

	"github.com/labstack/echo/v4"
)

type routeHandler struct {
	svc service.Service
}

func NewHTTP(e *echo.Echo, svc service.Service) {
	a := e.Group("/auth")

	r := routeHandler{
		svc: svc,
	}

	a.GET("/current", r.current)
	a.POST("/register", r.register)
}

func (r routeHandler) current(c echo.Context) error {
	str := r.svc.Current()

	return c.JSON(http.StatusOK, str)
}

func (r routeHandler) register(c echo.Context) error {
	r.svc.Register()

	return c.JSON(http.StatusOK, "registered!")
}
