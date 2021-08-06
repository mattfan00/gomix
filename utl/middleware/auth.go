package middleware

import (
	"net/http"

	"github.com/mattfan00/gomite/utl/entity"
	"github.com/mattfan00/gomite/utl/jwt"

	jwtGo "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func NewAuth(tp jwt.TokenParser) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token, err := tp.ParseToken(c.Request().Header.Get("Authorization"))
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
			}

			claims := (token.Claims.(jwtGo.MapClaims))

			user := entity.User{
				Username: claims["u"].(string),
			}

			c.Set("user", user)

			return next(c)
		}
	}
}
