package main

import "C"
import "fmt"

// Each Go function that will be called by the C code, needs to be exported first.
// This means that you should put a comment line starting with //export before its
// implementation. After //export you will need to put the name of the function
// because this is what the C code will use.

//export PrintMessage
func PrintMessage() {
	fmt.Println("A Go function!")
}

//export Multiply
func Multiply(a, b int) int {
	return a * b
}

// The main() needs no code because it is not going to be exported and therefore
// used by the C program.
//
// After that you will need to generate a C shared library from the Go code by
// executing the following command:
//
// go build -o usedByC.o -buildmode=c-shared usedByC.go
//
// The preceding command will generate two files named usedByC.o and usedByC.h
func main() {}
