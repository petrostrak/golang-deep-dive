package main

import "fmt"

// Pancake sorting is the colloquial term for the mathematical
// problem of sorting a disordered stack of pancakes in order
// of size when a spatula can be inserted at any point in the
// stack and used to flip all the pancakes above it. A pancake
// number is the maximum number of flips required for a given
// number of pancakes.
//
// The idea of the following algorithm is to search for the maximum
// [19, 23, 6, 15, 45, 30, 14]
//                  ^
// then flip the slice from index[0] to index[max]
// [45, 15, 6, 23, 19, 30, 14]
//
// afterwards put the index[0] at the last index and lock it
// [15, 6, 23, 19, 30, 14, 45]
//                          ^
// then repeat the same procedure without the locked elements
func main() {

	list := data{19, 23, 6, 15, 45, 30, 14}
	fmt.Println("\n--- Unsorted --- \n\n", list)

	list.pancakeSort()
	fmt.Println("\n--- Sorted ---\n\n", list)

}

type data []int32

func (dataList data) flip(r int) {
	for l := 0; l < r; l, r = l+1, r-1 {
		dataList[l], dataList[r] = dataList[r], dataList[l]
	}
}

func (dataList data) pancakeSort() {
	for unsorted := len(dataList) - 1; unsorted > 0; unsorted-- {
		// find the largest in the unsorted range
		lx, lg := 0, dataList[0]
		for i := 1; i <= unsorted; i++ {
			if dataList[i] > lg {
				lx, lg = i, dataList[i]
			}
		}

		// move to final position in two flips
		dataList.flip(lx)
		dataList.flip(unsorted)
	}
}
