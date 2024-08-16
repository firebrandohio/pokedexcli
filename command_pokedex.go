package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, _ ...string) error {
	if len(cfg.caughtPokemon) == 0 {
		return errors.New("you have not caught any pokemon yet")
	}
	fmt.Println("Pokemon you have caught:")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	fmt.Println("")
	return nil
}
