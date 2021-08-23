package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arr := []int{-3, -1, 1, 3, 4, 5, 8, 10, 12, 13, 15, 17, 19}
	x, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Gimme an integer to (linear) search")
	}

	linearSearch(arr, x)
}

// Linear Search
// Start from the leftmost element and compare it to the element that we
// search for. If it is a match, stop and return the location / index of
// the element. Move to the next element. If there is no match in the array
// return -1
//
// O(n) Complexity
func linearSearch(arr []int, target int) int {

	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			fmt.Printf("Found target at arr[%d]\n", i)
			return i
		}
	}

	fmt.Printf("Target %d not found in array\n", target)
	return -1
}
