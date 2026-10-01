package main

import (
"os"
"fmt"
)

func commandExit(fig *config, _ ...string) error {
        fmt.Printf("Closing the Pokedex... Goodbye!\n")
        os.Exit(0)
        return nil
}
