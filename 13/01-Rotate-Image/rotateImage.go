package main

import (
	"fmt"
)

func main() {

	img := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	N := len(img)

	for i := 0; i < N; i++ {
		fmt.Println(img[i])
	}

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
