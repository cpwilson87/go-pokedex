package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(" > ")

		scanner.Scan()
		text := scanner.Text()
		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}
		availableCommands := getCommands()
		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("invalid command")
			continue
		}
		error := command.callback(cfg, args...)
		if error != nil {
			fmt.Println(error)
		}

	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "prints the help menu",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "turns off pokedex",
			callback:    callbackExit,
		},
		"map": {
			name:        "map",
			description: "lists some location areas",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "lists previous page of location areas",
			callback:    callbackMapBack,
		},
		"explore": {
			name:        "explore",
			description: "lists pokemon in a location area",
			callback:    callbackExplore,
		},
		"catch": {
			name:        "catch",
			description: "attempt to catch a pokemon and add it to your pokedex",
			callback:    callbackCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "view information about caught pokemon",
			callback:    callbackInspect,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
