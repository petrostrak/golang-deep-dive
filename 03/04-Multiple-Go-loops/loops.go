package main

import "fmt"

func main() {

	// This shows a typical for loop as well as the use of continue and break
	// keywords.
	for i := 0; i < 100; i++ {
		if i%20 == 0 {
			continue
		}

		if i == 95 {
			break
		}

		fmt.Print(i, " ")
	}

	fmt.Println()

	// This code emulates the while loop. Note the break keyword to exit
	// the for loop.
	i := 10
	for {
		if i < 0 {
			break
		}
		fmt.Print(i, " ")
		i--
	}

	fmt.Println()

	// In this part, we see the use of a for loop that does the job of a do...while
	// loop.
	i = 0
	anExpression := true
	for ok := true; ok; ok = anExpression {
		if i > 10 {
			anExpression = false
		}

		fmt.Print(i, " ")
		i++
	}

	fmt.Println()

	// Applying the range keyword to an array variable returns two values, an array index
	// and the value of the element at that index, respectively.
	anArray := [5]int{0, 1, -1, 2, -2}
	for i, value := range anArray {
		fmt.Printf("index: %d, value: %d\n", i, value)
	}
}
