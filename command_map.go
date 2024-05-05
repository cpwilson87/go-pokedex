package main

import (
	"errors"
	"fmt"
	"log"
)

func callbackMap(cfg *config, args ...string) error {
	response, error := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if error != nil {
		return error
	}

	fmt.Println("Location areas:")
	for _, area := range response.Results {
		fmt.Printf(" - %s \n", area.Name)
	}
	cfg.nextLocationAreaURL = response.Next
	cfg.prevLocationAreaURL = response.Previous

	return nil
}
func callbackMapBack(cfg *config, args ...string) error {
	if cfg.prevLocationAreaURL == nil {
		return errors.New("you are on the first page")
	}
	response, error := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationAreaURL)
	if error != nil {
		log.Fatal(error)
	}

	fmt.Println("Location areas:")
	for _, area := range response.Results {
		fmt.Printf(" - %s \n", area.Name)
	}
	cfg.nextLocationAreaURL = response.Next
	cfg.prevLocationAreaURL = response.Previous

	return nil
}
