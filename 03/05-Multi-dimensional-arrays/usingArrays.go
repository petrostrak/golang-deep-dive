package main

import "fmt"

func main() {
	anArray := [4]int{1, 2, 3, 4}
	twoD := [4][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	threeD := [2][2][2]int{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}}

	fmt.Println("The length of", anArray, "is", len(anArray))
	fmt.Println("The length of", twoD, "is", len(twoD))
	fmt.Println("The length of", threeD, "is", len(threeD))

	// What we get from the 1st for loop is a 2d array (threeD[i])
	for i := 0; i < len(threeD); i++ {
		v := threeD[i]

		// What we get from the 2nd for loop is an array with one dimention(v[j])
		for j := 0; j < len(v); j++ {
			m := v[j]

			// The last for loop iterates over the elements of the array with
			// one dimention.
			for k := 0; k < len(m); k++ {
				fmt.Print(m[k], " ")
			}
		}
	}

	fmt.Println()

	for _, v := range threeD {
		for _, m := range v {
			for _, s := range m {
				fmt.Print(s, " ")
			}
		}
	}

	fmt.Println()
}
