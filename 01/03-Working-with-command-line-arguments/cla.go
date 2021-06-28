package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// os.Args is a Go slice of strings. The first element in the slice
	// is the name of the executable program. Therefore, in order to initialize
	// the min, max variables we need to use the second element of the slice
	if len(os.Args) == 1 {
		fmt.Println("Please give one or more floats.")
		os.Exit(1)
	}

	arguments := os.Args
	min, _ := strconv.ParseFloat(arguments[1], 64)
	max, _ := strconv.ParseFloat(arguments[1], 64)

	for i := 2; i < len(arguments); i++ {
		n, _ := strconv.ParseFloat(arguments[i], 64)

		if n < min {
			min = n
		}

		if n > max {
			max = n
		}
	}

	fmt.Println("Min is", min)
	fmt.Println("Max is", max)
}
