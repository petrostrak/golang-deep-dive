package main

import (
	"fmt"
	"os"
)

// The ... operator used as ...Type is called pack operator, whereas the unpack operator ends
// with ... and begins with a slice.
func varFunc(input ...string) {
	fmt.Println(input)
}

func oneByOne(msg string, s ...int) int {
	fmt.Println(msg)
	sum := 0
	for i, a := range s {
		fmt.Println(i, a)
		sum += a
	}
	s[0] = -1000
	return sum
}

func main() {

	intArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	arguments := os.Args
	varFunc(arguments...)
	sum := oneByOne("Adding numbers...", intArr...)
	fmt.Println("Sum:", sum)

}
