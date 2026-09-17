package main

import (
"os"
"fmt"
)

func commandExit(fig *config) error {
        fmt.Printf("Closing the Pokedex... Goodbye!\n")
        os.Exit(0)
        return nil
}
