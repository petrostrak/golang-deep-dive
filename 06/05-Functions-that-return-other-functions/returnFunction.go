package main

import "fmt"

func funReturnFun() func() int {
	i := 0
	return func() int {
		i++
		return i * i
	}
}

func main() {

	// i and j are completely unrelated to each other
	i := funReturnFun()
	j := funReturnFun()
	fmt.Println(i())
	fmt.Println(i())
	fmt.Println(i())
	fmt.Println(j())
	fmt.Println(j())
}
