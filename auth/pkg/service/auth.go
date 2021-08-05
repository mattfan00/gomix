package service

import (
	"net/http"

	"github.com/mattfan00/gomite/utl/entity"
	"github.com/mattfan00/gomite/utl/jwt"

	jwtGo "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

var (
	ErrInvalidCredentials = echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
)

func (s service) Current() string {
	return "this is me"
}

func (s service) Register(username string, password string) (entity.User, string, error) {
	newUser, err := s.mem.Register(username, password)

	token, err := jwt.GenerateToken(jwtGo.MapClaims{
		"u": username,
	})
	if err != nil {
		return entity.User{}, "", err
	}

	return newUser, token, err
}

func (s service) Login(username string, password string) (string, error) {
	if username != "matt" || password != "password" {
		return "", ErrInvalidCredentials
	}

	return "logged in", nil
}
