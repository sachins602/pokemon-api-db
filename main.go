package main

import (
	"fmt"
	"github.com/mtslzr/pokeapi-go"
	"net/http"
)

func main() {
	pokemonData, err := pokeapi.Pokemon("1")
	if err != nil {
		fmt.Println(err)
	}
	name := pokemonData.Name
	attack := pokemonData.Stats[4].BaseStat
	spriteUrl := fmt.Sprintf("https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/home/%d.png", 1)

	sprite, err := http.Get(spriteUrl)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(sprite)

	fmt.Println(name, attack, spriteUrl)
}
