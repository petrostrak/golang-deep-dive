package main

import (
	"fmt"
	"os"
	"strconv"
)

func doubleSquare(x int) (int, int) {
	return x * 2, x * x
}

func main() {

	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("The program needs one argument!")
		return
	}

	y, err := strconv.Atoi(arguments[1])
	if err != nil {
		fmt.Println("cannot convert argument to integer")
		return
	}

	square := func(s int) int {
		return s * s
	}
	fmt.Printf("The square of %d is %d\n", y, square(y))

	double := func(s int) int {
		return s + s
	}
	fmt.Printf("The double of %d is %d\n", y, double(y))

	fmt.Println(doubleSquare(y))
	d, s := doubleSquare(y)
	fmt.Println(d, s)

}
