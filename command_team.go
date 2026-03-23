package main

import (
	"fmt"
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
			return fmt.Errorf("usage: team remove <pokemon>")
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

func teamRemove(cfg *config, name string) error {
	index := -1
	for i, p := range cfg.team {
		if p.Name == name {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Printf("'%s' not in current team\n", name)
		return nil
	}

	if len(cfg.team) == 1 {
		return fmt.Errorf("cannot have an empty team")
	}

	// pokemon removal
	cfg.team = append(cfg.team[:index], cfg.team[index+1:]...)

	fmt.Printf("%s removed from your team\n", name)
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