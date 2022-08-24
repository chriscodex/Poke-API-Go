package main

import (
	"net/http"

	"github.com/ChrisCodeX/Poke-API-Go/handler"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/pokemon/{id}", handler.HandlerGetPokemon).Methods(http.MethodGet)

}
