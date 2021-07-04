package main

import "fmt"

func printSlice(x []int) {
	for _, number := range x {
		fmt.Print(number, " ")
	}

	fmt.Println()
}

// This illustrates the concepts of capacity and length in more detail
func main() {
	aSlice := []int{-1, 0, 4}
	fmt.Printf("aSlice: ")
	printSlice(aSlice)

	fmt.Printf("Cap: %d, Len: %d\n", cap(aSlice), len(aSlice))
	aSlice = append(aSlice, -100)
	fmt.Printf("aSlice: ")
	printSlice(aSlice)
	fmt.Printf("Cap: %d, Len: %d\n", cap(aSlice), len(aSlice))

	aSlice = append(aSlice, -2)
	aSlice = append(aSlice, -3)
	aSlice = append(aSlice, -4)
	printSlice(aSlice)
	fmt.Printf("Cap: %d, Len: %d\n", cap(aSlice), len(aSlice))
}
