package b

import (
	"fmt"

	"github.com/petrostrak/golang-deep-dive/06/08-The-init-function/a"
)

func init() {
	fmt.Println("init() b")
}

func FromB() {
	fmt.Println("fromB()")
	a.FromA()
}
