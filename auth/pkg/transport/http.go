package transport

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mattfan00/gomite/auth/pkg/service"
)

type routeHandler struct {
	svc service.Service
}

func NewHTTP(e *echo.Echo, svc service.Service) {
	a := e.Group("/auth")

	r := routeHandler{svc: svc}

	a.GET("/current", r.current)
	a.POST("/register", r.register)
	a.POST("/login", r.login)
}

func (r routeHandler) current(c echo.Context) error {
	str := r.svc.Current()

	return c.JSON(http.StatusOK, str)
}

func (r routeHandler) register(c echo.Context) error {
	body := registerRequest{}
	if err := c.Bind(&body); err != nil {
		return err
	}

	users, err := r.svc.Register(body.Username, body.Password)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, registerResponse{
		Users: users,
	})
}

func (r routeHandler) login(c echo.Context) error {
	body := loginRequest{}
	if err := c.Bind(&body); err != nil {
		return err
	}

	message, err := r.svc.Login(body.Username, body.Password)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, loginResponse{
		Message: message,
	})
}
