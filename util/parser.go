package util

import (
	"errors"
)

var (
	// ErrNotFoundPokemonType occurs when the type array in pokeapi response it's not found
	ErrNotFoundPokemonType = errors.New("pokemon type array not found")
	// ErrNotFoundPokemonTypeName occurs when we found type struct but no name
	ErrNotFoundPokemonTypeName = errors.New("pokemon type name not found")
)
