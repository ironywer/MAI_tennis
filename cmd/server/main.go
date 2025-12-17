package main

import (
	"net/http"

	"github.com/ironywer/MAI_tennis/internal/httpserver"
)

func main() {
	handler := httpserver.New()
	s := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}
	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
