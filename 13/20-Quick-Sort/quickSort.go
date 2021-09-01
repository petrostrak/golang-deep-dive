package main

import (
	"fmt"
	"math/rand"
	"time"
)

// The quick sort algorithm falls under the divide and conquer class
// of algorithms, where we break (devide) a problem into smaller chunks
// that are much simpler to solve (conquer). In this case, an unsorted
// array is broken into sub-arrays that are partially sorted, until all
// elements in the list are in the right position, by which time our
// unsorted list will have become sorted.
func main() {

	slice := generateSlice(20)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	quickSort(slice)
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

func quickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1
	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quickSort(a[:left])
	quickSort(a[left+1:])

	return a
}
