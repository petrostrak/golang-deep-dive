package main

import (
	"fmt"
	"time"
)

func writeToChan(c chan int, x int) {
	fmt.Println(x)
	c <- x
	close(c)

	// This will not execute. That's because c <- x is blocking
	// the execution of the rest of the writeToChan() because
	// nobody is reading what was written to the c channel. Therefore
	// when the time.Sleep() statement finishes, the program terminates
	// without waiting for the writeToChan().
	fmt.Println(x)
}

func main() {

	c := make(chan int)
	go writeToChan(c, 10)
	time.Sleep(time.Second)

}
