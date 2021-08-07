package transport

import (
	"net/http"

	"github.com/mattfan00/gomix/auth/pkg/service"
	"github.com/mattfan00/gomix/utl/entity"

	"github.com/labstack/echo/v4"
)

type routeHandler struct {
	svc service.Service
}

func NewHTTP(e *echo.Echo, svc service.Service, authMiddleware echo.MiddlewareFunc) {
	a := e.Group("/auth")

	r := routeHandler{svc: svc}

	a.GET("/current", r.current, authMiddleware)
	a.POST("/register", r.register)
	a.POST("/login", r.login)
}

func (r routeHandler) current(c echo.Context) error {
	user := c.Get("user").(entity.User)

	return c.JSON(http.StatusOK, user)
}

func (r routeHandler) register(c echo.Context) error {
	body := registerRequest{}
	if err := c.Bind(&body); err != nil {
		return err
	}

	newUser, authToken, err := r.svc.Register(body.Username, body.Password)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, registerResponse{
		User:   newUser,
		Tokens: authToken,
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
