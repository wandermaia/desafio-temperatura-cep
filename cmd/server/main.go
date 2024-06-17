package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/wandermaia/desafio-temperatura-cep/internal/infra/webserver/handlers"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Get("/{cep}", handlers.BuscaTemperaturaHandler)
	//router.Get("/", handlers.BuscaCepHandler)

	log.Println("Servidor iniciado!")
	http.ListenAndServe(":8000", router)
}
