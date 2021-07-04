package main

import "fmt"

func main() {
	s1 := make([]int, 5)

	// We can access multiple continuous slice elements using the [:] notation. The
	// next statement selects the second and the third elements of the slice. Additionally
	// we can use the [excluding:including] notation for creating a new slice from an existing slice or array
	reSlice := s1[1:3]
	fmt.Println(s1)
	fmt.Println(reSlice)

	// We can access the first element of the slice as reSlice[0]
	reSlice[0] = -100
	reSlice[1] = 123456
	fmt.Println(s1)
	fmt.Println(reSlice)

	// We can access the last element of the slice as reSlice[len(reSlice)-1]
	fmt.Println(reSlice[len(reSlice)-1])
}
