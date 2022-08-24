package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ChrisCodeX/Poke-API-Go/models"
	"github.com/ChrisCodeX/Poke-API-Go/util"
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

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var apiPokemonResponse models.PokeApiPokemonResponse

	err = json.Unmarshal(body, &apiPokemonResponse)
	if err != nil {
		log.Fatal(err)
	}

	parsedPokemon, err := util.ParsePokemon(apiPokemonResponse)
	if err != nil {
		respondWithJson(w, http.StatusInternalServerError, fmt.Sprintf("error found: %s", err.Error()))
	}

	respondWithJson(w, http.StatusOK, parsedPokemon)
}
