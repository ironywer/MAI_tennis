package httpserver

import (
	"net/http"
)

func New() http.Handler {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	h := &Handlers{}
	mux.HandleFunc("/", h.Index)
	mux.HandleFunc("/health", h.Health)
	return mux
}
