package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	arr := []int{-9, -3, 1, 2, 5, 8}

	fmt.Println(sortedSquaredArray(arr))
	fmt.Println(sortedSquaredArrayButBetter(arr))
}

// O(n) Complexity
func sortedSquaredArrayButBetter(arr []int) []int {
	newArr := make([]int, len(arr))

	// we keep track on the left and right indexes of the array
	// because, since the input array is sorted, these would be
	// the maximum values (afted the squared calculation).
	left := 0
	right := len(arr) - 1

	for i := len(arr) - 1; i >= 0; i-- {

		// if we determine the maximum value, we append it
		// at the end of the output array (the loop starts from
		// back to front)
		if math.Abs(float64(arr[left])) > float64(arr[right]) {
			newArr[i] = arr[left] * arr[left]

			// if the maximum value is on the left side, we
			// increment
			left++
		} else {
			newArr[i] = arr[right] * arr[right]

			// else we decrement
			right--
		}
	}

	return newArr
}

// O(nlogn) Complexity
func sortedSquaredArray(arr []int) []int {
	newArr := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {

		// for each index of the array, calculate the
		// squared value
		newArr[i] = arr[i] * arr[i]
	}

	// sort it
	sort.Slice(newArr, func(i, j int) bool {
		return newArr[i] < newArr[j]
	})
	return newArr
}
