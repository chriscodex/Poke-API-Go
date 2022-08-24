package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Respond to client with json
func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	_, err = w.Write(response)
	if err != nil {
		log.Fatal(err)
	}
}

func HandlerGetPokemon(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	request := fmt.Sprintf("http://pokeapi.co/api/v2/pokemon/%s", id)

	response, err := http.Get(request)
	if err != nil {
		log.Fatal(err)
	}
}
