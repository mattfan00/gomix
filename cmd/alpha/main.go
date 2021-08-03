package main

import (
	"github.com/mattfan00/gomite/pkg/server"
	"github.com/mattfan00/gomite/svc/alpha"
	"github.com/mattfan00/gomite/svc/alpha/transport"
)

func main() {
	e := server.New()

	alphaService := alpha.New()

	transport.NewHTTP(e, alphaService)

	e.Logger.Fatal(e.Start(":8080"))
}
