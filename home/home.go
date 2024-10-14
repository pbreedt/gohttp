package home

import "net/http"

func AddHandlers(mux *http.ServeMux) {
	mux.HandleFunc("GET /{$}", home)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HOME"))
}
