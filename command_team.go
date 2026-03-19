package main

import (
	"fmt"
	"strconv"
)

func commandTeam(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: team <add/remove/list>")
	}

	switch args[0] {

	case "add":
		if len(args) != 2 {
			return fmt.Errorf("usage: team add <pokemon>")
		}
		return teamAdd(cfg, args[1])

	case "remove":
		if len(args) != 2 {
			return fmt.Errorf("usage: team remove <index>")
		}
		return teamRemove(cfg, args[1])

	case "list":
		return teamList(cfg)

	default:
		return fmt.Errorf("unknown subcommand: %s", args[0])
	}
}

func teamAdd(cfg *config, name string) error {
	pokemon, ok := cfg.caughtPokemon[name]
	if !ok {
		return fmt.Errorf("you haven't caught %s", name)
	}

	if len(cfg.team) >= 6 {
		return fmt.Errorf("your team is full")
	}

	// prevent duplicates
	for _, p := range cfg.team {
		if p.Name == name {
			return fmt.Errorf("%s is already in your team", name)
		}
	}

	cfg.team = append(cfg.team, pokemon)
	fmt.Printf("%s added to your team!\n", name)
	return nil
}

func teamRemove(cfg *config, indexStr string) error {
	index, err := strconv.Atoi(indexStr)
	if err != nil {
		return fmt.Errorf("invalid index")
	}

	if index < 1 || index > len(cfg.team) {
		return fmt.Errorf("index out of range")
	}

	// convert to 0-based index
	i := index - 1
	removed := cfg.team[i]

	cfg.team = append(cfg.team[:i], cfg.team[i+1:]...)

	fmt.Printf("%s removed from your team\n", removed.Name)
	return nil
}

func teamList(cfg *config) error {
	if len(cfg.team) == 0 {
		fmt.Println("your team is empty")
		return nil
	}

	fmt.Println("Your team:")

	for i, p := range cfg.team {
		fmt.Printf("%d. %s\n", i+1, p.Name)
	}

	return nil
}