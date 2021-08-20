package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{-9, -3, 1, 2, 5, 8}

	fmt.Println(sortedSquaredArray(arr))

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
