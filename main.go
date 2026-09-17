package main

type config struct {
	registry map[string]cliCommand
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func main() {
	newConfig := &config{}
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
    }

	startREPL(newConfig)
}
