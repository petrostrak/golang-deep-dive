package main

import (
	"fmt"
	"time"
)

func function() {
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}
}

// The output supports the fact that you cannot control the order in which our goroutines will be
// executed without taking extra care. This means writing extra code specifically for this to occur.
func main() {

	// The program starts by executing function() as a goroutine. After that, the program continues
	// its execution, while function() begins to run in the background.
	go function()

	go func() {
		for i := 10; i < 20; i++ {
			fmt.Print(i)
		}
	}()

	time.Sleep(1 * time.Second)
	fmt.Println()

}
