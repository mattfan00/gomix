package alpha

import (
	"github.com/mattfan00/gomite/pkg/utl/server"

	"github.com/labstack/echo/v4"
)

func Start() {
	e := server.New()

	e.GET("/hey", func(c echo.Context) error {
		return c.JSON(200, "hey from alpha")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
