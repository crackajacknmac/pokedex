package main

import (
    "fmt"
    "bufio"
    "os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commands map[string]cliCommand

func commandExit() error {
        fmt.Printf("Closing the Pokedex... Goodbye!")
        os.Exit(0)
	return nil
}

func commandHelp() error {
        fmt.Printf("Welcome to the Pokedex!\n")
        fmt.Printf("Usage:\n\n")
	for _, a := range commands {
		fmt.Printf("%s: %s\n", a.name, a.description)
	}
	return nil
}


func main() {
	scanner := bufio.NewScanner(os.Stdin)

	commands = map[string]cliCommand{
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


	for {
		fmt.Print("Pokedex > \n")
		scanner.Scan()
		text := scanner.Text()
		cleanText := cleanInput(text)
		if len(cleanText) == 0 {
			continue
			}
		command, ok := commands[cleanText[0]]
		if ok {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
				}
		}else {
			fmt.Printf("Unknown command\n")
		}
	}
}
