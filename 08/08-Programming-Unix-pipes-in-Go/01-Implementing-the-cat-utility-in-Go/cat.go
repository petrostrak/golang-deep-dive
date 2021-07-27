package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// printFile's purpose is to print the contents of a file in the standard library
func printFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		io.WriteString(os.Stdout, scanner.Text())
		io.WriteString(os.Stdout, "\n")
	}
	return nil
}

func main() {

	filename := ""
	arguments := os.Args

	// if we execute the programm without any command-line arguments, then the
	// program will just copy standard input to standard output as defined by
	// io.Copy(os.Stdout, os.Stdin)
	if len(arguments) == 1 {
		io.Copy(os.Stdout, os.Stdin)
		return
	}

	for i := 1; i < len(arguments); i++ {
		filename = arguments[i]
		if err := printFile(filename); err != nil {
			fmt.Println(err)
		}
	}

}
