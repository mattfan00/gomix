package service

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	ErrUnauthorized = echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
)

func (s service) Current() string {
	return "this is me"
}

func (s service) Register(username string, password string) ([]string, error) {
	users, err := s.mem.Register(username, password)
	return users, err
}

func (s service) Login(username string, password string) (string, error) {
	if username != "matt" || password != "password" {
		return "", ErrUnauthorized
	}

	return "logged in", nil
}
