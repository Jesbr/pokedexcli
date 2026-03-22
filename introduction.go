package main

import (
	"bufio"
	"fmt"
	"os"
)

func getIntroduction(cfg *config) error {
	reader := bufio.NewScanner(os.Stdin)
	introduction := true

	for {
		if introduction == true {
			fmt.Println("Welcome to the Pokedexcli!")
		}
		introduction = false
		fmt.Println("Please choose a starter:")
		fmt.Println("Charmander")
		fmt.Println("Bulbasaur")
		fmt.Println("Squirtle")
		fmt.Print("Pokedex > ")
		reader.Scan()
		args := cleanInput(reader.Text())

		if len(args) != 1 {
			fmt.Println("You must provide a starter pokemon name...")
			continue
		}

		name := args[0]
		if name == "charmander" || name == "bulbasaur" || name == "squirtle" {
			pokemon, err := cfg.pokeapiClient.GetPokemon(name)
			if err != nil {
				return err
			}
			fmt.Printf("You have chosen: %s\n", pokemon.Name)
			cfg.caughtPokemon[pokemon.Name] = pokemon
			cfg.team = append(cfg.team, pokemon)
			fmt.Printf("%s has been added to your team!\n", pokemon.Name)
			return nil
		} else {
			fmt.Printf("'%s' is not a starter...\n", name)
			continue
		}
	}
}
