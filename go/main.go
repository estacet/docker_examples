package main

import (
	"github.com/estacet/docker_examples/go/handler"
	"github.com/estacet/docker_examples/go/server"
	"log"
	"net/http"
)

func main() {
	homeHandler := handler.NewHomeHandler()

	mux := http.NewServeMux()
	mux.Handle("/", homeHandler)

	apiServer := server.NewAPIServer(mux)

	if err := apiServer.Start(); err != nil {
		log.Fatal(err)
	}
}
