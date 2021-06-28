package main

import (
	"io"
	"os"
)

func main() {
	myString := ""
	arguments := os.Args

	if len(arguments) == 1 {
		myString = "Please give me one argument!"
	} else {
		myString = arguments[1]
	}

	// The preceding output cannot help you to differentiate between data written
	// to stdout and data written to stderr. However, if you are using the bash(1)
	// shell, there is a trick you can use in order to distinguish between stdout
	// and stderr data. Almost all UNIX shells offer this functionality in their
	// own way
	//
	// go run stdERR.go 2>/tmp/stdError
	// cat /tmp/stdError
	//
	// Similarly you can discard error output by redirecting it to the /dev/nill device
	// go run stdERR.go 2>/dev/null
	//
	// If you want to save both Stdout and Stderr to the same file, you can redirect the
	// file descriptor of Stderr(2) to the file descriptor of Stdout(1) like so
	//
	// go run stdERR.go >/tmp/output 2>&1
	io.WriteString(os.Stdout, "This is Standard output\n")
	io.WriteString(os.Stderr, myString)
	io.WriteString(os.Stderr, "\n")
}
