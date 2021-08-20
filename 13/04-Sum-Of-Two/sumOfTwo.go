package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	a := []int{2, 4, -5, 7}
	b := []int{3, -1, 5, 8}
	c, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Failed to parse input")
	}

	sumOfTwo(a, b, c)

}

func sumOfTwo(a, b []int, c int) bool {
	for i := 0; i < len(a); i++ {

		// we subtract from the given integer the index
		// so that we can find what we need.
		neededValue := c - a[i]
		for j := 0; j < len(b); j++ {
			if b[i] == neededValue {
				fmt.Println("True")
				return true
			}
		}
	}
	fmt.Println("False")
	return false
}
