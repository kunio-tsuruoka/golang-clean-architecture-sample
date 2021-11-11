package main

import (
	"mytestapp/registry"
)

func main() {
	interactor := registry.NewPokemonInteractor()
	handler := interactor.NewPokemonHandler()
	handler.FindPokemon()
	return
}
