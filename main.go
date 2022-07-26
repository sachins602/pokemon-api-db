package main

import (
	"fmt"

	"github.com/mtslzr/pokeapi-go"
)

func main() {
	a, err := pokeapi.Pokemon("1")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(a, err)
}
