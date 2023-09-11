package handler

import (
	"log"
	"net/http"
)

type HomeHandler struct{}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	if _, err := w.Write([]byte("Hello from container")); err != nil {
		log.Println(err)
	}
}
