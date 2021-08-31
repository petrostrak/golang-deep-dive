package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// A tiny program that removes the .NEF files from the SD of my camera.
func main() {

	dir, err := os.ReadDir(os.Args[1])
	if err != nil {
		fmt.Println("Cannot open directory")
	}

	for _, foto := range dir {
		if strings.HasSuffix(foto.Name(), ".NEF") {
			cmd := exec.Command("rm", fmt.Sprintf(os.Args[1]+"/"+foto.Name()))
			if err := cmd.Run(); err != nil {
				fmt.Printf("Could not remove %s\n", foto.Name())
			}
		}
	}

}
