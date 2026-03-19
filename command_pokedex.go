package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) < 1 {
		fmt.Println("You have not caught any Pokemon yet...")
		return nil
	}
	fmt.Println("Your pokedex:")
	for key := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", key)
	}
	return nil
}