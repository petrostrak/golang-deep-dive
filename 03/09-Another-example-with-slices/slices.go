package main

import "fmt"

func main() {
	aSlice := []int{1, 2, 3, 4, 5}
	fmt.Println("aSlice:", aSlice)
	integers := make([]int, 2)
	fmt.Println("integers:", integers)
	integers = nil
	fmt.Println("integers:", integers)

	anArray := [5]int{-1, -2, -3, -4, -5}
	refArray := anArray

	fmt.Println("anArray:", anArray)
	fmt.Println("refArray:", refArray)
	anArray[4] = 100
	fmt.Println("refArray:", anArray)

	s := make([]byte, 5)
	fmt.Println("s:", s)
	twoD := make([][]int, 3)
	fmt.Println("twoD:", twoD)
	fmt.Println()

	// Manually initialize all the elements of a slice with two dimensions
	for i := 0; i < len(twoD); i++ {
		for j := 0; j < 2; j++ {
			twoD[i] = append(twoD[i], i*j)
		}
		fmt.Println()
	}

	// Visit and print all the elements of a slice with two dimentions
	for _, x := range twoD {
		for i, y := range x {
			fmt.Println("i:", i, "value:", y)
		}
		fmt.Println()
	}
}
