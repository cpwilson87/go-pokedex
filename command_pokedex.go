package main

import (
	"fmt"
)

func callbackPokedex(cfg *config, _ ...string) error {
	fmt.Println("Pokemon in Pokedex:")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf("Name: %s\n", pokemon.Name)
	}
	return nil
}
