package main

import "fmt"

// A Go node is a struct with a large number of properties. Everything in a Go
// program is parsed and analyzed by the modules of the Go compiler according
// to the grammar of the Go programming language. The final product of the analysis
// is a tree that is specific to the provided Go code and represents the program
// in a different way that is suited to the compiler rather than to the developer.
//
// go tool compile -w nodeTree.go
//
// The -W parametertells the go tool compile command to print the debug parse tree
// after the type checking.
func main() {
	fmt.Println("Hello 世界!")
}
