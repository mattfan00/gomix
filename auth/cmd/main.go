package main

import (
	"github.com/mattfan00/gomite/auth/pkg/service"
	"github.com/mattfan00/gomite/auth/pkg/transport"

	"github.com/mattfan00/gomite/utl/server"
)

func main() {
	e := server.New()

	authService := service.New()

	transport.NewHTTP(e, authService)

	e.Logger.Fatal(e.Start(":8080"))
}
