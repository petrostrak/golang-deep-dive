package main

import (
	"flag"
	"fmt"
	"sync"
)

func main() {

	n := flag.Int("n", 10, "Number of goroutines")
	flag.Parse()

	count := *n
	fmt.Printf("Going to create %d goroutines\n", count)

	// type WaitGroup struct {
	// 	noCopy noCopy
	// 	state1 [3]uint32
	// }
	var waitGroup sync.WaitGroup
	fmt.Printf("%#v\n", waitGroup)

	for i := 0; i < count; i++ {
		// It is important to call sync.Add(1) before the go statement in order
		// to  avoid race condition.
		waitGroup.Add(1)
		go func(x int) {
			defer waitGroup.Done()
			fmt.Printf("%d ", x)
		}(i)
	}

	fmt.Printf("%#v\n", waitGroup)

	// The sync.Wait() call blocks until the counter in the relevant sync.WaitGroup
	// variable is zero, giving goroutines time to finish.
	waitGroup.Wait()
	fmt.Println("\nExiting...")
}
