package main

import (
	"fmt"
	"os"
	"strconv"
)

// The namedMinMax function does not explicitly return any variables of values
// in its return statement. Nevertheless, as this function has named return values
// in its signature, the min and max parameters are automatically returned in the
// order in which they were put into the function definition.
func namedMinMax(x, y int) (min, max int) {
	if x > y {
		min = y
		max = x
	} else {
		min = x
		max = y
	}

	return
}

func minMax(x, y int) (min, max int) {
	if x > y {
		min = y
		max = x
	} else {
		min = x
		max = y
	}

	return min, max
}

func main() {
	arguments := os.Args
	if len(arguments) < 3 {
		fmt.Println("The program needs 2 arguments!")
		return
	}

	i, err := strconv.Atoi(arguments[1])
	if err != nil {
		fmt.Println("cannot convert given argument to integer")
		return
	}
	j, err := strconv.Atoi(arguments[2])
	if err != nil {
		fmt.Println("cannot convert given argument to integer")
		return
	}

	fmt.Println(minMax(i, j))
	min, max := minMax(i, j)
	fmt.Println(min, max)

	fmt.Println(namedMinMax(i, j))
	min, max = namedMinMax(i, j)
	fmt.Println(min, max)
}
