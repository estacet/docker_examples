package main

import (
	"github.com/estacet/go_docker/handler"
	"github.com/estacet/go_docker/server"
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
