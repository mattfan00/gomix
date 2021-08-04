package transport

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mattfan00/gomite/svc/auth"
)

type routeHandler struct {
	svc auth.Service
}

func NewHTTP(e *echo.Echo, svc auth.Service) {
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
