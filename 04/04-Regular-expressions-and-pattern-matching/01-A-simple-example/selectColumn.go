package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {

	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Printf("usage: selectedColumn column <file1> [<file2> [<fileN]]\n")
		os.Exit(1)
	}

	temp, err := strconv.Atoi(arguments[1])
	if err != nil {
		fmt.Println("Column value is not an integer:", temp)
		return
	}

	column := temp
	if column < 0 {
		fmt.Println("Invalid Column number")
		os.Exit(1)
	}

	for _, filename := range arguments[2:] {
		fmt.Println("\t\t", filename)
		f, err := os.Open(filename)
		if err != nil {
			fmt.Printf("error opening file %s\n", err)
			continue
		}
		defer f.Close()

		r := bufio.NewReader(f)
		for {

			// r.ReadString reads a file until the first occurance of its parameter. As a result
			// bufio.ReadString('\n') tells Go to read a file line by line because \n is the UNIX
			// new line character.
			line, err := r.ReadString('\n')
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Printf("error reading file %s", err)
			}

			// strings.Fields function splits a string based on the whitespace characters that are defined
			// in the unicode.Inspace function and returns a slice of strings.
			data := strings.Fields(line)
			if len(data) >= column {
				fmt.Println(data[column-1])
			}
		}
	}

}
