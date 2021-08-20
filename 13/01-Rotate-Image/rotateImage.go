package main

import (
	"fmt"
)

func main() {

	img := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	N := len(img)

	// [1 2 3]
	// [4 5 6]
	// [7 8 9]
	// Initial array
	for i := 0; i < N; i++ {
		fmt.Println(img[i])
	}

	// [1 4 7]
	// [2 5 8]
	// [3 6 9]
	// This nested loop swaps the diagonal values
	// 2-4, 3-7 and 6-8
	for i := 0; i < N; i++ {
		for j := i; j < N; j++ {
			temp := img[i][j]
			img[i][j] = img[j][i]
			img[j][i] = temp
		}
	}

	fmt.Println()
	for i := 0; i < N; i++ {
		fmt.Println(img[i])
	}

	// [7 4 1]
	// [8 5 2]
	// [9 6 3]
	// We traveres through the half array
	// swapping the values at the eadges
	for i := 0; i < N; i++ {
		for j := 0; j < N/2; j++ {
			temp := img[i][j]
			img[i][j] = img[i][N-j-1]
			img[i][N-j-1] = temp
		}
	}

	fmt.Println()
	for i := 0; i < N; i++ {
		fmt.Println(img[i])
	}

}
