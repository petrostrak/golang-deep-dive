package main

import (
	"fmt"
	"math/rand"
	"time"
)

// The idea of swapping adjacent elements to sort a list can
// be also be used to implement the insertion sort. In the
// insertion sort algorithm, we assume that a certain portion
// of the list has already been sorted, while the other portion
// remains unsorted. With this assumption, we move through the
// unsorted portion of the list, picking one element at a time.
// With this element, we go through the sorted portion of the
// list and insert it in the right order so that the sorted
// portion of the list remains sorted.
func main() {

	slice := generateSlice(20)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	insertionSort(slice)
	fmt.Println("\n--- Sorted ---\n\n", slice)

}

func generateSlice(size int) []int {
	slice := make([]int, size)
	max := 999
	min := -999
	rand.Seed(time.Now().Unix())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(max-min) + min
	}

	return slice
}

func insertionSort(items []int) {
	n := len(items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j--
		}
	}
}
