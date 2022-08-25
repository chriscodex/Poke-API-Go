package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/pokemon/pikachu").Methods(http.MethodGet)

	err := http.ListenAndServe(":8080", router)

	if err != nil {
		log.Println("Error found")
	}
}
