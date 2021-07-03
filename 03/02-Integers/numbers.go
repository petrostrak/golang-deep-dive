package main

import "fmt"

func main() {
	x := 12
	k := 5
	fmt.Println(x)
	fmt.Printf("Type of x: %T\n", x)

	// In this example, we are working with singed integers. Please note
	// that if you divide two integers, Go thinks that you want the result
	// of the integer division and will calculate and return the quotient
	// of the integer division.
	div := x / k
	fmt.Println("div:", div)
}
