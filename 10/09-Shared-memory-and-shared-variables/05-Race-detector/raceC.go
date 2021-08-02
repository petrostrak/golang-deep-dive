package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

func main() {

	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("Give me a natural number")
		return
	}

	numGR, err := strconv.Atoi(arguments[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	var waitGroup sync.WaitGroup
	// var m sync.Mutex
	var i int

	k := make(map[int]int)
	k[1] = 12

	for i = 0; i < numGR; i++ {
		waitGroup.Add(1)
		go func() { //	i int
			defer waitGroup.Done()
			// m.Lock()
			k[i] = i
			// m.Unlock()
		}() //	i
	}

	k[2] = 10
	waitGroup.Wait()
	fmt.Printf("k = %#v\n", k)

}
