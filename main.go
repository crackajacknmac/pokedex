package main

import (
    "fmt"
    "bufio"
    "os"
)

type config struct {
	registry map[string]cliCommand
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

//var commands map[string]cliCommand

func commandExit(fig *config) error {
        fmt.Printf("Closing the Pokedex... Goodbye!\n")
        os.Exit(0)
	return nil
}

func commandHelp(fig *config) error {
        fmt.Printf("Welcome to the Pokedex!\n")
        fmt.Printf("Usage:\n\n")
	for _, a := range fig.registry {
		fmt.Printf("%s: %s\n", a.name, a.description)
	}
	return nil
}


func startREPL(fig *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > \n")
		scanner.Scan()
		text := scanner.Text()
		cleanText := cleanInput(text)
		if len(cleanText) == 0 {
			continue
			}
		command, ok := fig.registry[cleanText[0]]
		if ok {
			err := command.callback(fig)
			if err != nil {
				fmt.Println(err)
				}
		}else {
			fmt.Printf("Unknown command\n")
		}
	}
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
    }

	startREPL(newConfig)
}
