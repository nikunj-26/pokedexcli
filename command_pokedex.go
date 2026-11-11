package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {

	fmt.Println("Your Pokedex:")
	for _, p := range cfg.caughtPokemon {
		fmt.Printf("  - %v\n", p.Name)
	}

	return nil
}
