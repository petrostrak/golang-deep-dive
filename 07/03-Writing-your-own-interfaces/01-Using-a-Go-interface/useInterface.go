package main

import (
	"fmt"
	"math"

	myinterface "github.com/petrostrak/golang-deep-dive/07/03-Writing-your-own-interfaces/myInterface"
)

type square struct {
	X float64
}

type circle struct {
	R float64
}

func (s square) Area() float64 {
	return s.X * s.X
}

func (s square) Perimeter() float64 {
	return 4 * s.X
}

func (c circle) Area() float64 {
	return math.Pi * math.Pow(c.R, 2)
}

func (c circle) Perimeter() float64 {
	return 2 * math.Pi * c.R
}

func Calculate(x myinterface.Shape) {
	_, ok := x.(circle)
	if ok {
		fmt.Println("Its a cirlce")
	}

	_, ok = x.(square)
	if ok {
		fmt.Println("Its a square!")
	}

	fmt.Println(x.Area())
	fmt.Println(x.Perimeter())

}

func main() {
	x := square{10}
	fmt.Println("Perimeter:", x.Perimeter())
	Calculate(x)
	y := circle{5}
	Calculate(y)
}
