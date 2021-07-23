package main

import "fmt"

func returnPtr(y int) *int {
	x := y * y
	return &x
}

func main() {
	sq := returnPtr(10)
	fmt.Println("sq value:", *sq)
	fmt.Println("sq memory address:", sq)
}
