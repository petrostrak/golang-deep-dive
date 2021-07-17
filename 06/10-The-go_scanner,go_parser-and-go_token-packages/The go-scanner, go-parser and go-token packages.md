### The go/scanner, go/parser and go/token packages
Parsing a language requires two phases. The first one is about breaking up the input into tokens (`lexical analysis`) and the second one is about feeding the parser with all these tokens in order to make sure that these tokens make sense and are in the right order
(`semantic analysis`).

### The go/ast package
An `abstract syntax tree (AST)` is a structured representation of the source code of a Go program. This tree is constructed according to some rules that are specified in the language specification. The `go/ast` package is used for declaring the data types required to represent ASTs in Go.

### The go/scanner package
A scanner is something that reads a program and generates tokens. The `go/scanner` package is used for reading Go programs and generating a series of tokens.