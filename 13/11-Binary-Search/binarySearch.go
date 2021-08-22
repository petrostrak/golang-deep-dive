package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arr := []int{-4, -1, 2, 4, 6, 7, 9, 10, 14, 15, 16, 19, 21, 23, 24, 26, 27, 29, 30}
	x, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Gimme an integer to (binary) search", err)
	}
	binarySearch(arr, x)
}

// Find element from sorted arrays. It works by repeatedly dividing in half the portion
// of the list that could contain the target element, until we' ve narrowed down the
// possible locations to just one.
// arr = [-4, -1, 2, 4, 6, 7, 9, 10, 14, 15, 16, 19, 21, 23, 24, 26, 27, 29, 30]
//         L	 					   Midpoint								  R
func binarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		midpoint := left + (right-left)/2
		if arr[midpoint] == target {
			fmt.Printf("Target found at arr[%d]\n", midpoint)
			return midpoint
		} else if target < arr[midpoint] {
			right = midpoint - 1
		} else {
			left = midpoint + 1
		}
	}

	return -1
}
