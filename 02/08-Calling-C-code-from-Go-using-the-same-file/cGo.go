package main

// All other packages should be imported separately
import "fmt"

// #include <stdio.h>
// void callC() {
// 	printf("Calling C code!\n");
// }
import "C"

func main() {
	fmt.Println("A Go statement")
	C.callC()
	fmt.Println("Another Go statement")
}
