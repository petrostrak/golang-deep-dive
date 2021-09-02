package main

import (
	"fmt"
	"math/rand"
	"time"
)

// This sorting algorithm begins by finding the smallest element
// in an array and interchanging it with data at, for instance
// index[0]. Starting from index 0, we search for the smallest
// item in the list that exists between index 1 and the index
// of the last element. When this element has been found, it is
// exchanged with the data found at index 0. We simply repeat
// this process until the list becomes sorthed.
func main() {

	slice := generateSlice(20)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	selectionSort(slice)
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

func selectionSort(items []int) {
	var n = len(items)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if items[j] < items[minIdx] {
				minIdx = j
			}

		}
		items[i], items[minIdx] = items[minIdx], items[i]
	}
}
