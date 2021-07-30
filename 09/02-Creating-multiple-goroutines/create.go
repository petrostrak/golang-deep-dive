package main

import (
	"flag"
	"fmt"
	"time"
)

// go run create.go -n 300
func main() {

	n := flag.Int("n", 10, "Number of goroutines")
	flag.Parse()

	count := *n
	for i := 0; i < count; i++ {
		go func(x int) {
			fmt.Printf("%d ", x)
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Println("\nExiting...")

}
