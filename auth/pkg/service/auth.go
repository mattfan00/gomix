package service

import (
	"net/http"

	"github.com/mattfan00/gomite/utl/entity"

	jwtGo "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

var (
	ErrInvalidCredentials = echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
)

func (s service) Register(username string, password string) (entity.User, entity.AuthToken, error) {
	newUser, err := s.pg.Register(entity.User{
		Username: username,
		Password: password,
	})

	tokenClaims := jwtGo.MapClaims{"u": username}

	accessToken, err := s.atg.GenerateToken(tokenClaims)
	if err != nil {
		return entity.User{}, entity.AuthToken{}, err
	}

	refreshToken, err := s.rtg.GenerateToken(tokenClaims)
	if err != nil {
		return entity.User{}, entity.AuthToken{}, err
	}

	authToken := entity.AuthToken{
		Access:  accessToken,
		Refresh: refreshToken,
	}

	return newUser, authToken, err
}

func (s service) Login(username string, password string) (string, error) {
	if username != "matt" || password != "password" {
		return "", ErrInvalidCredentials
	}

	return "logged in", nil
}
