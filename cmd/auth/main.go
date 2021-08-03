package main

import (
	"github.com/mattfan00/gomite/pkg/server"
	"github.com/mattfan00/gomite/svc/auth"
	"github.com/mattfan00/gomite/svc/auth/transport"
)

func main() {
	e := server.New()

	authService := auth.New()

	transport.NewHTTP(e, authService)

	e.Logger.Fatal(e.Start(":8080"))
}
