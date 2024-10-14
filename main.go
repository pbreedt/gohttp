package main

import (
	"net/http"

	"github.com/pbreedt/gohttp/home"
	"github.com/pbreedt/gohttp/users"
)

func main() {
	s := http.NewServeMux()
	home.AddHandlers(s)

	users.AddHandlers(s)

	http.ListenAndServe(":8080", s)
}
