package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Given an unordered list, we compare adjacent elements in the list,
// each time, putting in the right order of magnitude, only two elements.
// The algorithm hinges on a swap procedure. Knowning how many times
// to swap is important when implementing a bubble sort algorithm. To
// sort a list of numbers such as [3,2,1], we need to swap the elements
// a maximum of twice. This is equal to the length of the list minus 1
func main() {

	slice := generateSlice(20)
	fmt.Println("\n--- Unsorted ---\n\n", slice)
	bubbleSort(slice)
	fmt.Println("\n--- Sorted ---\n\n", slice)

}

// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {
	min := -999
	max := 999
	slice := make([]int, size)
	rand.Seed(time.Now().Unix())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(max-min) + min
	}
	return slice
}

func bubbleSort(items []int) {
	var (
		n      = len(items)
		sorted = false
	)

	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if items[i] > items[i+1] {
				items[i+1], items[i] = items[i], items[i+1]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		n = n - 1
	}
}
