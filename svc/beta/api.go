package beta

import (
	"github.com/mattfan00/gomite/pkg/server"

	"github.com/labstack/echo/v4"
)

func Start() {
	e := server.New()

	e.GET("/hey", func(c echo.Context) error {
		return c.JSON(200, "hey from beta")
	})

	e.Logger.Fatal(e.Start(":8081"))
}
