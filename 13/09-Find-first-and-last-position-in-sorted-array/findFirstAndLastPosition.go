package main

import "fmt"

// Given an array of integers sorted in ascending order, find the
// starting and ending position of a given target value.
// Your algorithm's runtim complexity must be in the order of
// O(logn).
// If the target is not found in the array, return [-1,-1]
//
// Input arr = [5,7,7,8,8,10], target = 8
// Output [3,4]
//
// When we see O(logn) is binary search!
// Binary search implementation
func main() {

	arr := []int{5, 7, 8, 8, 8, 9, 10}
	result := make([]int, 2)
	result[0] = startingIndex(arr, 8)
	result[1] = endingIndex(arr, 8)

	fmt.Println(result)

}

func startingIndex(arr []int, target int) int {

	// if we don't find any target, we return -1
	index := -1

	start := 0
	end := len(arr) - 1

	for start <= end {

		// we cound specify midpoint by start + end /2
		// but for overflow reasons we do it start + (end-start) /2
		midpoint := start + (end-start)/2
		if arr[midpoint] >= target {
			end = midpoint - 1
		} else {
			start = midpoint + 1
		}

		if arr[midpoint] == target {
			midpoint = target
		}
	}

	return index
}

func endingIndex(arr []int, target int) int {

	// if we don't find any target, we return -1
	index := -1

	start := 0
	end := len(arr) - 1

	for start <= end {

		// we cound specify midpoint by start + end /
		// but for overflow reasons we do it start + (end-start) /2
		midpoint := start + (end-start)/2
		if arr[midpoint] <= target {
			start = midpoint + 1
		} else {
			end = midpoint - 1
		}

		if arr[midpoint] == target {
			midpoint = target
		}
	}

	return index
}
