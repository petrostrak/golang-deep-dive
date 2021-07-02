package main

import (
	"fmt"
	"os"
)

// Using the panic() function on its own will terminate the Go program without giving
// you the opportunity to recover.
func main() {
	if len(os.Args) == 1 {
		panic("Not enough arguments!")
	}

	fmt.Println("Thanks for the argument(s)!")
}
