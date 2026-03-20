package main

import (
	"fmt"
)

func commandLocation(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: team <current/tbd>")
	}

	switch args[0] {

	case "current":
		if cfg.location.name == "" {
			return fmt.Errorf("You have not selected a location yet...")
		} else {
			return fmt.Errorf("Your current location is: %s", cfg.location.name)
		}
	default:
		return fmt.Errorf("unknown subcommand: %s", args[0])
	}
}