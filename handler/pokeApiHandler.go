package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ChrisCodeX/Poke-API-Go/models"
	"github.com/gorilla/mux"
)

var (
	ErrPokemonNotFound = errors.New("pokemon not found")
	ErrPokeApiFailure  = errors.New("pokeapi failure")
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

func GetPokemonFromPokeApi(id string) (*models.PokeApiPokemonResponse, error) {
	request := fmt.Sprintf("http://pokeapi.co/api/v2/pokemon/%s", id)

	response, err := http.Get(request)
	if err != nil {
		return nil, err
	}

	/* 	Getting the status code from the pokeapi and validate*/
	if response.StatusCode == http.StatusNotFound {
		return nil, ErrPokemonNotFound
	}

	if response.StatusCode != http.StatusOK {
		return nil, ErrPokeApiFailure
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var apiPokemonResponse models.PokeApiPokemonResponse

	err = json.Unmarshal(body, &apiPokemonResponse)
	if err != nil {
		return nil, err
	}

	return &apiPokemonResponse, nil
}

func HandlerGetPokemon(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	apiPokemonResponse, err := GetPokemonFromPokeApi(id)

	if errors.Is(err, ErrPokemonNotFound) {
		respondWithJson(w, http.StatusNotFound, fmt.Sprintf("pokemon not found: %s"))
	}

	if err != nil {
		respondWithJson(w, http.StatusInternalServerError, fmt.Sprintf("error found: %s", err.Error()))
	}
}
