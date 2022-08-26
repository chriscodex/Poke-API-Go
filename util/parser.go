package util

import (
	"errors"

	"github.com/ChrisCodeX/Poke-API-Go/models"
)

var (
	// ErrNotFoundPokemonType occurs when the type array in pokeapi response it's not found
	ErrNotFoundPokemonType = errors.New("pokemon type array not found")
	// ErrNotFoundPokemonTypeName occurs when we found type struct but no name
	ErrNotFoundPokemonTypeName = errors.New("pokemon type name not found")
)

func ParsePokemon(apiPokemonResponse models.PokeApiPokemonResponse) (models.Pokemon, error) {
	if len(apiPokemonResponse.PokemonType) < 1 {
		return models.Pokemon{}, ErrNotFoundPokemonType
	}

	if apiPokemonResponse.PokemonType[0].RefType.Name == "" {
		return models.Pokemon{}, ErrNotFoundPokemonTypeName
	}

	pokemonType := apiPokemonResponse.PokemonType[0].RefType.Name

	abilitiesMap := map[string]int{}

	for _, stat := range apiPokemonResponse.Stats {
		parsedAbilityName, ok := models.AllowedAbilities[stat.Stat.Name]
		if !ok {
			continue
		}

		abilitiesMap[parsedAbilityName] = stat.BaseStat
	}

	parsedPokemon := models.Pokemon{
		Id:        apiPokemonResponse.Id,
		Name:      apiPokemonResponse.Name,
		Power:     pokemonType,
		Abilities: abilitiesMap,
	}

	return parsedPokemon, nil
}
