package main

import "fmt"

type square struct {
	X float64
}

type circle struct {
	R float64
}

type rectangle struct {
	X float64
	Y float64
}

func tellInterface(x interface{}) {

	// x.(type) returns the type of the x element
	switch v := x.(type) {
	case square:
		fmt.Println("This is a square")
	case circle:
		fmt.Printf("%v This is a cirlce!\n", v)
	case rectangle:
		fmt.Println("This is a rectangle")
	default:
		fmt.Printf("Unknown type %T\n", v)
	}
}

func main() {

	x := circle{10}
	tellInterface(x)
	y := rectangle{4, 1}
	tellInterface(y)
	z := square{4}
	tellInterface(z)
	tellInterface(10)

}
