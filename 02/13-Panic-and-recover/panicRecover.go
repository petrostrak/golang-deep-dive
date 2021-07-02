package main

import "fmt"

// Stricly speaking, panic() is a build-in Go function that terminates the current flow of a Go
// program and starts panicking. On the other hand, the recover() function, which is also a
// build-in Go function, allows you to take back control of a goroutine that just panicked using
// panic().

func a() {
	fmt.Println("Inside a()")
	defer func() {
		if c := recover(); c != nil {
			fmt.Println("Recover inside a()!")
		}
	}()

	fmt.Println("About to call b()")
	b()
	fmt.Println("b() exited!")
	fmt.Println("Exiting a()")
}

func b() {
	fmt.Println("Inside b()")
	panic("Panic in b()!")
}

func main() {
	a()
	fmt.Println("main() ended")
}
