package main

import (
	"github.com/mattfan00/gomite/auth/pkg/platform/memory"
	"github.com/mattfan00/gomite/auth/pkg/service"
	"github.com/mattfan00/gomite/auth/pkg/transport"
	"github.com/mattfan00/gomite/utl/jwt"
	"github.com/mattfan00/gomite/utl/middleware"
	"github.com/mattfan00/gomite/utl/transport/http"

	jwtGo "github.com/dgrijalva/jwt-go"
)

func main() {
	e := http.NewServer()
	mem := memory.New()

	accessTokenService := jwt.New("access", jwtGo.SigningMethodHS256)
	refreshTokenService := jwt.New("refresh", jwtGo.SigningMethodHS256)

	authService := service.New(mem, accessTokenService, refreshTokenService)
	authMiddleware := middleware.NewAuth(accessTokenService)

	transport.NewHTTP(e, authService, authMiddleware)

	e.Logger.Fatal(e.Start(":8080"))
}
