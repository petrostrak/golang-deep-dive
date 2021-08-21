package main

import (
	"fmt"
	"math"
)

// Given an array of integers, 1<=a[i]<= n (n=size of array), some
// elements appear twice and others appear once.
//
// Find all the elements that appear twice in this array.
// Could you do it without extra space and in O(n) runtime?
func main() {
	arr := []int{4, 3, 2, 7, 8, 2, 3, 1}
	findDuplicates(arr)
}

/*
[4, 3, 2, 7, 8, 2, 3, 1]
input[i] = 4
index = 4-1 = 3
input[3] = 7
7<0 false
7 = -7

[4, 3, 2, -7, 8, 2, 3, 1]
input[i] = 3
index = 3-1 = 2
input[2] = 2
2<0 false
2 = -2

[4, 3, -2, -7, 8, 2, 3, 1]
input[i] = -2
index = 2-1 = 1
input[1] = 3
3<0 false
3 = -3

[4, -3, -2, -7, 8, 2, 3, 1]
input[i] = -7
index = 7-1 = 6
input[6] = 3
3<0 false
3 = -3

[4, -3, -2, -7, 8, 2, -3, 1]
input[i] = 8
index = 8-1 = 7
input[7] = -3
-3<0 true
append index + 1

.
.
.

*/
func findDuplicates(input []int) []int {
	output := make([]int, len(input))

	for i := 0; i < len(input); i++ {

		// every element of the array is a valid index
		index := math.Abs(float64(input[i])) - 1
		if input[int(index)] < 0 {
			output = append(output, int(index)+1)
		} else {
			input[int(index)] = -input[int(index)]
		}
	}

	fmt.Println(output)
	return output
}
