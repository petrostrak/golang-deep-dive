package main

import (
	"fmt"
	"os"
	"strconv"
)

// arr := []int{1,2,3,7,5}
// input := 12
// output := arr[2:7] and not arr[7:5] because
// {2,3,7} is the longest
func main() {

	arr := []int{1, 2, 3, 4, 5, 0, 0, 0, 6, 7, 8, 9, 10}
	x, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(findLongestSubarrayBySum(arr, x))

}

// O(n) Complexity
func findLongestSubarrayBySum(arr []int, s int) []int {
	result := []int{-1}

	sum := 0
	right := 0
	left := 0

	for right < len(arr) {
		sum += arr[right]

		// if sum goes over s we move the subarray window over
		for left < right && sum > s {
			left++
			sum -= arr[left]
		}

		if sum == s && len(result) == 1 {
			result = []int{left + 1, right + 1}
		}
		right++
	}

	return result
}
