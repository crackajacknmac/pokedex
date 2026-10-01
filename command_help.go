package main

import (
"fmt"
)


func commandHelp(fig *config, _ ...string) error {
        fmt.Printf("Welcome to the Pokedex!\n")
        fmt.Printf("Usage:\n\n")
        for _, a := range fig.registry {
                fmt.Printf("%s: %s\n", a.name, a.description)
        }
        return nil
}
