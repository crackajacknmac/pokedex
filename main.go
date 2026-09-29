cpackage main

import (
	"time"
	"github.com/crackajacknmac/pokedex/internal/pokeapi"
)

type config struct {
	registry map[string]cliCommand
	pageNext *string
	pageBack *string
	pokeapiClient pokeapi.Client
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func main() {
	newConfig := &config{}
	newConfig.pokeapiClient = pokeapi.NewClient(5 * time.Second)
        newConfig.registry = map[string]cliCommand{
        "exit": {
                name:        "exit",
                description: "Exit the Pokedex",
                callback:    commandExit,
        },
        "help": {
                name:           "help",
                description: "Displays a help message",
                callback:       commandHelp,
        },
	"map": {
		name:		"map",
		description: "Displays location data",
		callback:	commandMap,
	},
	"mapb": {
		name:		"mapb",
		description: "Displays previous location data",
		callback:	commandMapb,
	},
	"explore": {
		name:		"explore",
		description:	"Displays more details about location",
		callback:	explore,
	},
    }

	startREPL(newConfig)
}
