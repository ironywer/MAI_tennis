package httpserver

import (
	"net/http"
)

func New() http.Handler {
	mux := http.NewServeMux()
	h := &Handlers{}
	mux.HandleFunc("/", h.Index)
	mux.HandleFunc("/health", h.Health)
	return mux
}
