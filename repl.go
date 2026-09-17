package main

import (
"fmt"
"bufio"
"os"
)

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
