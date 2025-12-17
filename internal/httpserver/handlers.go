package httpserver

import (
	"net/http"
)

type Handlers struct{}

func (h *Handlers) Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("<html><head><link rel=\"stylesheet\" href=\"/static/css/style.css\"></head><body><h1>Welcome to the Index Page</h1></body></html>"))
}
func (h *Handlers) Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}
