package main

import "fmt"

func getPtr(y *float64) float64 {
	return *y * *y
}

func main() {
	x := 12.2
	fmt.Println(getPtr(&x))
}
