package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/midorigreen/goapi-architecture/handler"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	r := chi.NewRouter()
	handler.Registe(r)
	return http.ListenAndServe(":3333", r)
}
