package repository

import (
	// 標準パッケージのimport･･･1
	"fmt"
	"io/ioutil"
	domain "mytestapp/domain/pokemon"
	"net/http"
)

type pokemonRepository struct {
}

func NewPokemonRepository() domain.PokemonRepository {
	return &pokemonRepository{}
}

func (s pokemonRepository) GetPokemonAPIData() {
	url := "https://pokeapi.co/api/v2/pokemon/ditto"

	req, _ := http.NewRequest("GET", url, nil)
	client := new(http.Client)

	resp, _ := client.Do(req)
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byteArray))
	return
}
