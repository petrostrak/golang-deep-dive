package main

import (
	"io"
	"os"
)

func main() {
	myStrings := " "
	arguments := os.Args
	if len(arguments) == 1 {
		myStrings = "Please give me one argument"
	} else {
		myStrings = arguments[1]
	}

	//  io.WriteString takes two parametes. The first parameter is the file
	// you want to write to, and the second parameter is a string variable
	//
	// io.WriteString sends the contents of its second parameter onto the screen
	// when its first parameter is os.Stdout
	io.WriteString(os.Stdout, myStrings)
	io.WriteString(os.Stdout, "\n")
}
