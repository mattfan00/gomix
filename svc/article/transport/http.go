package transport

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mattfan00/gomite/svc/article"
)

type routeHandler struct {
	svc article.Service
}

func NewHTTP(e *echo.Echo, svc article.Service) {
	a := e.Group("/article")

	r := routeHandler{
		svc: svc,
	}

	a.GET("/", r.get)
}

func (r routeHandler) get(c echo.Context) error {
	str := r.svc.Get()

	return c.JSON(http.StatusOK, str)
}
