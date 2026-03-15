package main

import (
	"time"
	"github.com/bootdotdev/pokedexcli/internal/pokeapi"
	"fmt"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second) //, time.Minute*5)
	fmt.Println("starting...")
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}