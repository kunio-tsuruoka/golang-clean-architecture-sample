package handler

import (
	domain "mytestapp/domain/pokemon" // 標準パッケージのimport･･･1
)

type PokemonHandler struct {
	pokemonRepository domain.PokemonRepository
}

func NewPokemonHandler(pokemonRepository domain.PokemonRepository) *PokemonHandler {
	return &PokemonHandler{pokemonRepository}
}

func (h PokemonHandler) FindPokemon() {
	h.pokemonRepository.GetPokemonAPIData()
	return
}
