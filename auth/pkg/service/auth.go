package service

import (
	"net/http"

	"github.com/mattfan00/gomix/utl/entity"

	jwtGo "github.com/dgrijalva/jwt-go"
	"github.com/go-pg/pg/v10"
	"github.com/labstack/echo/v4"
)

var (
	ErrUsernameInUse      = echo.NewHTTPError(http.StatusBadRequest, "Username already in use")
	ErrInvalidCredentials = echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
)

func (s service) Register(username string, password string) (entity.User, entity.AuthToken, error) {
	_, err := s.pg.FindUserByUsername(username)
	if err == nil || err != pg.ErrNoRows {
		return entity.User{}, entity.AuthToken{}, ErrUsernameInUse
	}

	hashedPassword, err := s.bc.Hash(password)
	if err != nil {
		return entity.User{}, entity.AuthToken{}, err
	}

	newUser, err := s.pg.Register(entity.User{
		Username: username,
		Password: hashedPassword,
	})
	if err != nil {
		return entity.User{}, entity.AuthToken{}, err
	}

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

func (s service) Login(username string, plainPassword string) (entity.User, entity.AuthToken, error) {
	foundUser, err := s.pg.FindUserByUsername(username)
	// if the user does not exist
	if err != nil {
		return entity.User{}, entity.AuthToken{}, ErrInvalidCredentials
	}

	if correctPassword := s.bc.Compare(foundUser.Password, plainPassword); !correctPassword {
		return entity.User{}, entity.AuthToken{}, ErrInvalidCredentials
	}

	tokenClaims := jwtGo.MapClaims{"u": foundUser.Username}

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

	return foundUser, authToken, err
}
