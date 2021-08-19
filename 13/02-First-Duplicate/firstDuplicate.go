package main

import (
	"fmt"
	"math"
)

// Find the first duplicate of the given array.
// First duplicate means the minimum index of
// duplicate within the array. The program returns
// -1 if no duplicates are found.
//
// Constrains:
// The integers of the array should not be bigger
// than the length of the array.
func main() {

	arr := []int{1, 2, 1, 4, 2, 3}

	fmt.Println(firstDup(arr))

}

// O(n^2) Complexity
func firstDup(arr []int) int {
	MIN_SECOND_INDEX := len(arr)

	// i would be the first pointer that traverses the array.
	for i := 0; i < len(arr); i++ {

		// j would be the second pointer that leads the i. Hence i+1
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {

				// if we find an occurance, we update MIN_SECOND_INDEX. We want
				// the minimum index, so for every occurance, we keep the min.
				MIN_SECOND_INDEX = int(math.Min(float64(MIN_SECOND_INDEX), float64(j)))
			}
		}
	}

	// if MIN_SECOND_INDEX remains the same, it means that no duplicates
	// are found, so we return -1.
	if MIN_SECOND_INDEX == len(arr) {
		return -1
	} else {

		// if duplicates are found, we return the value of MIN_SECOND_INDEX
		return arr[MIN_SECOND_INDEX]
	}

}
