package main

import (
	"fmt"

	"github.com/petrostrak/golang-deep-dive/06/08-The-init-function/a"
	"github.com/petrostrak/golang-deep-dive/06/08-The-init-function/b"
)

func init() {
	fmt.Println("init() from main")
}

func main() {
	a.FromA()
	b.FromB()
}
