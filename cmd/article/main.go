package main

import (
	"github.com/mattfan00/gomite/pkg/server"
	"github.com/mattfan00/gomite/svc/article"
	"github.com/mattfan00/gomite/svc/article/transport"
)

func main() {
	e := server.New()

	articleService := article.New()

	transport.NewHTTP(e, articleService)

	e.Logger.Fatal(e.Start(":8081"))
}
