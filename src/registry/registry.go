package registry

import (
	domain "mytestapp/domain/pokemon"
	"mytestapp/handler"
	"mytestapp/infrastructure/repository"
)

type pokemonInteractor struct {
}

type PokemonInteractor interface {
	NewPokemonHandler() *handler.PokemonHandler
}

func (i *pokemonInteractor) NewPokemonRepository() domain.PokemonRepository {
	return repository.NewPokemonRepository()
}

func NewPokemonInteractor() PokemonInteractor {
	return &pokemonInteractor{}
}

func (i *pokemonInteractor) NewPokemonHandler() *handler.PokemonHandler {
	return handler.NewPokemonHandler(i.NewPokemonRepository())
}
