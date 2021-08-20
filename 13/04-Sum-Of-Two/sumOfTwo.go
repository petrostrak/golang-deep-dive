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
	sumOfTwoButBetter(a, b, c)

}

// O(2n) Complexity
func sumOfTwoButBetter(a, b []int, c int) bool {

	// we create a map to calculate and store the value we need in order
	// to have an addition.
	//
	// Interacting with maps have a constant time complexity
	complementArr := make(map[int]int)

	for i := 0; i < len(a); i++ {

		// for every index of the given array, we calculate the value that
		// we need and store it into the map
		complementArr[i] = c - a[i]
	}

	for j := 0; j < len(b); j++ {

		// then we check if there is a value from the second array
		// that matches for the addition.
		if complementArr[j] == b[j] {
			fmt.Println("True")
			return true
		}
	}

	fmt.Println("False")
	return false
}

// O(n^2) Complexity
func sumOfTwo(a, b []int, c int) bool {
	for i := 0; i < len(a); i++ {

		// we subtract from the given integer the index
		// so that we can find what we need.
		//
		// for example, if c = 7 and a[i] = 4 then we need
		// the value of 3 to have a positive result.
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
