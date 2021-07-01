package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var value int64 = 5
	var p1 = &value

	// unsafe.Pointer allows us, at our own risk, to create an int32 pointer
	// named p2 that points to an int64 variable named value, which is accessed
	// using the p1 pointer. Any Go pointer can be converted to unsafe.Pointer
	var p2 = (*int32)(unsafe.Pointer(p1))

	fmt.Println("*p1: ", *p1)
	fmt.Println("*p2: ", *p2)
	*p1 = 123412341234
	fmt.Println(value)
	fmt.Println("*p2: ", *p2)
	*p1 = 4535435
	fmt.Println(value)
	fmt.Println("*p2: ", *p2)
}

// The output tells us that a 32-bit pointer cannot store a 64-bit integer
//
// *p1:  5
// *p2:  5
// 123412341234
// *p2:  -1141710350
// 4535435
// *p2:  4535435
