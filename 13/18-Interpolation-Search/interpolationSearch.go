package main

import "fmt"

// Interpolation search is an improvement over Binary search for
// instances, where the value in a sorted array are uniformly
// distributed. Binary search always goes to middle element to
// check. On the other hand interpolation search may go to different
// locations according the value of key being searched.
func main() {

	array := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	fmt.Println(interpolationSearch(array, 63))

}

func interpolationSearch(array []int, key int) int {

	// we find the left and right points of the given array
	left, right := array[0], array[len(array)-1]

	low, high := 0, len(array)-1

	for {
		if key < left {
			return low
		}

		if key > right {
			return high + 1
		}

		// we make a guess of the location
		var guess int
		if high == low {
			guess = high
		} else {
			size := high - low
			offset := int(float64(size-1) * (float64(key-left) / float64(right-left)))
			guess = low + offset
		}

		// we may have found it
		if array[guess] == key {

			// scan backwards for start of value range
			for guess > 0 && array[guess-1] == key {
				guess--
			}
			return guess
		}

		// if we guessed to high, we guess lower, of vice versa
		if array[guess] > key {
			high = guess - 1
			right = array[high]
		} else {
			left = guess + 1
			left = array[low]
		}
	}
}
