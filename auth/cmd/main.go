package main

import (
	"log"
	"os"

	"github.com/mattfan00/gomite/auth/pkg/platform/pg"
	"github.com/mattfan00/gomite/auth/pkg/service"
	"github.com/mattfan00/gomite/auth/pkg/transport"
	"github.com/mattfan00/gomite/utl/jwt"
	"github.com/mattfan00/gomite/utl/middleware"
	"github.com/mattfan00/gomite/utl/store"
	"github.com/mattfan00/gomite/utl/transport/http"

	jwtGo "github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := http.NewServer()

	pgDB := store.NewPG(os.Getenv("PG_CONN"))
	pgStore := pg.New(pgDB)

	accessTokenService := jwt.New("access", jwtGo.SigningMethodHS256)
	refreshTokenService := jwt.New("refresh", jwtGo.SigningMethodHS256)

	authService := service.New(
		pgStore,
		accessTokenService,
		refreshTokenService,
	)

	authMiddleware := middleware.NewAuth(accessTokenService)

	transport.NewHTTP(e, authService, authMiddleware)

	e.Logger.Fatal(e.Start(":8080"))
}
