package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "aaabbbcccdddeeefjjj"

	firstNonRepeatingChar(str)

}

// O(n^2) Complexity
func firstNonRepeatingChar(s string) byte {

	// we traverse through each element of the given string
	for i := 0; i < len(s); i++ {

		// we set a flag to keep track on whether we have found
		// duplicate character or not.
		seenDuplicate := false

		// j pointer checks if the given index is the same with the index of the
		// i pointer or not. (the i != j condition makes sure not to count the same
		// index as equal)
		for j := 0; j < len(s); j++ {
			if strings.IndexByte(s, s[i]) == strings.IndexByte(s, s[j]) && i != j {

				// if the above condition is satisfied, we set the flag to true and
				// break
				seenDuplicate = true
				break
			}
		}

		if !seenDuplicate {
			fmt.Println(string(s[i]))
			return s[i]
		}
	}
	fmt.Println("_")
	return '_'
}
