package models

type PokeApiPokemonResponse struct {
	Id          int           `json:"id"`
	Name        string        `json:"name"`
	PokemonType []PokemonType `json:"types"`
	Stats       []PokemonStat `json:"stats"`
}
