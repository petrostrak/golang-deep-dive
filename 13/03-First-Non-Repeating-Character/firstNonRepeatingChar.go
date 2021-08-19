package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	str := "aaabbbcccdddeffjjj"

	firstNonRepeatingChar(str)
	firstNonRepeatingCharButBetter(str)

}

// O(2n) Complexity
func firstNonRepeatingCharButBetter(s string) byte {

	// key is the character of the string
	// value is the total numbers of repetition.
	//
	// The idea behind this algorithm is to pass all the
	// characters into a map and increment the values of
	// them in every occurance.
	strMap := make(map[byte]int)

	for i := 0; i < len(s); i++ {

		// for shake of simplicity we are declaring the character of the index
		// into a variable c.
		c := byte(s[i])

		// first we check if the character of the index is already in the map
		if strMap[c] != 0 {

			// if it is already in the map, we increment the value of that particular index key
			strMap[c]++
		} else {

			// else we initialize the value to 1.
			strMap[c] = 1
		}
	}

	// since maps are unordered data structures, we need a second loop
	// to get the 1st single occurance.
	for i := 0; i < len(s); i++ {
		c := byte(s[i])
		if strMap[c] == 1 {
			fmt.Fprintln(os.Stdout, string(c))
			return c
		}
	}

	fmt.Fprintln(os.Stdout, "_")
	return '_'
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
