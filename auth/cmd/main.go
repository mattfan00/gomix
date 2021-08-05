package main

import (
	"github.com/mattfan00/gomite/auth/pkg/platform/memory"
	"github.com/mattfan00/gomite/auth/pkg/service"
	"github.com/mattfan00/gomite/auth/pkg/transport"

	"github.com/mattfan00/gomite/utl/http"
)

func main() {
	e := http.NewServer()
	mem := memory.New()

	authService := service.New(mem)

	transport.NewHTTP(e, authService)

	e.Logger.Fatal(e.Start(":8080"))
}
