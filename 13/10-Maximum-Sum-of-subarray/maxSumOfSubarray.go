package main

import (
	"fmt"
	"math"
)

// Given an array of integers, find the maximum possible sum you can get
// from one of its contiguous subarrays. The subarray from which this sum
// comes must contain at least 1 element.
//
// For inputArray = [-2, 2, 5, -11, 6]
// the output should be 7
func main() {
	arr := []int{-2, 2, 5, -11, 6}
	kadanesAlgorithm(arr)
}

func kadanesAlgorithm(input []int) int {

	// we set the maxSum and currentSum to be the first element of the
	// given array.
	maxSum := input[0]
	currentSum := maxSum

	// for each element, starting from index 1, we check whether the currentSum
	// (which is also the maxSum) is greater than the sum of the currentSum plus
	// the next index.
	for i := 1; i < len(input); i++ {
		currentSum = int(math.Max(float64(input[i])+float64(currentSum), float64(input[i])))
		maxSum = int(math.Max(float64(currentSum), float64(maxSum)))
	}

	fmt.Println(maxSum)
	return maxSum
}
