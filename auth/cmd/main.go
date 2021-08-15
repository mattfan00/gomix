package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mattfan00/gomix/auth/pkg/platform/pg"
	"github.com/mattfan00/gomix/auth/pkg/service"
	"github.com/mattfan00/gomix/auth/pkg/transport"
	"github.com/mattfan00/gomix/utl/bcrypt"
	"github.com/mattfan00/gomix/utl/jwt"
	"github.com/mattfan00/gomix/utl/middleware"
	"github.com/mattfan00/gomix/utl/store"
	"github.com/mattfan00/gomix/utl/transport/http"

	jwtGo "github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	bcryptGo "golang.org/x/crypto/bcrypt"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := http.NewServer()

	fmt.Println(os.Getenv("PG_CONN"))
	pgDB := store.NewPG(os.Getenv("PG_CONN"))
	pgStore := pg.New(pgDB)

	accessTokenService := jwt.New("access", jwtGo.SigningMethodHS256)
	refreshTokenService := jwt.New("refresh", jwtGo.SigningMethodHS256)

	bcryptService := bcrypt.New(bcryptGo.DefaultCost)

	authService := service.New(
		pgStore,
		accessTokenService,
		refreshTokenService,
		bcryptService,
	)

	authMiddleware := middleware.NewAuth(accessTokenService)

	transport.NewHTTP(e, authService, authMiddleware)

	e.Logger.Fatal(e.Start(":8080"))
}
