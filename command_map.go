package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config, _ ...string) error {
	areaURL := cfg.nextLocationAreaURL
	res, err := cfg.pokeapiClient.ListLocationAreas(areaURL)
	if err != nil {
		return err
	}
	cfg.nextLocationAreaURL = res.Next
	cfg.prevLocationAreaURL = res.Previous

	fmt.Println("Location Areas:")
	for _, area := range res.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	fmt.Println("")
	return nil
}

func commandMapB(cfg *config, _ ...string) error {
	areaURL := cfg.prevLocationAreaURL
	if areaURL == nil {
		return errors.New("you are on the first page")
	}
	res, err := cfg.pokeapiClient.ListLocationAreas(areaURL)
	if err != nil {
		return err
	}
	cfg.nextLocationAreaURL = res.Next
	cfg.prevLocationAreaURL = res.Previous

	fmt.Println("Location Areas:")
	for _, area := range res.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	fmt.Println("")
	return nil
}
