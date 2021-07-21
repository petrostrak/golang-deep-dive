package main

import "fmt"

type a struct {
	XX int
	YY int
}

type b struct {
	AA string
	XX int
}

// Composition allows us to create a structure in our Go elements using multiple struc types. In
// this case, data type c groups an a variable and a b variable.
type c struct {
	A a
	B b
}

func (A a) A() {
	fmt.Println("Function A() for A")
}

func (B b) B() {
	fmt.Println("Function B() for B")
}

func main() {

	var i c
	i.A.A()
	i.B.B()

}
