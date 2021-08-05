package middleware

import (
	"fmt"
	"net/http"

	"github.com/mattfan00/gomite/utl/jwt"

	"github.com/labstack/echo/v4"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, err := jwt.ParseToken(c.Request().Header.Get("Authorization"))
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}

		fmt.Println(token.Claims)

		return next(c)
	}
}
