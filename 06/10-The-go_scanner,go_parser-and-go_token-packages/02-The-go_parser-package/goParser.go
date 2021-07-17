package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

type Visitor int

func (v Visitor) Visit(n ast.Node) ast.Visitor {
	if n == nil {
		return nil
	}
	fmt.Printf("%s%T\n", strings.Repeat("\t", int(v)), n)
	return v + 1
}

func main() {

	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Not enough arguments")
		return
	}

	for _, file := range os.Args[1:] {
		fmt.Println("Processing:", file)
		one := token.NewFileSet()
		var v Visitor
		f, err := parser.ParseFile(one, file, nil, parser.AllErrors)
		if err != nil {
			fmt.Println(err)
			return
		}

		// The Walk() which is called recurcively, traverses an AST in depth-first order
		// in order to visit all of its nodes.
		ast.Walk(v, f)
	}

}
