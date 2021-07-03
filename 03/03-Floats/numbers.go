package main

import "fmt"

func main() {
	x := 12
	var m, n float64
	m = 1.223
	fmt.Println("m, n:", m, n)

	y := 4 / 2.3
	fmt.Println("y:", y)

	divFloat := float64(x) / float64(y)
	fmt.Println("divFloat:", divFloat)
	fmt.Printf("Type of divFloat: %T\n", divFloat)
}
