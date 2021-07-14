package main

import (
	"container/ring"
	"fmt"
)

var (
	// size variable holds the size of the ring that is going to be created.
	size int = 10
)

func main() {

	// creates a new ring
	myRing := ring.New(size + 1)
	fmt.Println("Empty ring:", myRing)

	for i := 0; i < myRing.Len()-1; i++ {
		myRing.Value = i
		myRing = myRing.Next()
	}

	// adds 2 in the ring
	myRing.Value = 2

	sum := 0

	// Do() allows us to call a function for each element of a ring in chronological order.
	// However if the function makes any changes to the ring, then the behavior of the ring.Do()
	// is undefined.
	myRing.Do(func(i interface{}) {

		// type assertion.
		t := i.(int)
		sum = sum + t
	})

	fmt.Println("Sum:", sum)

	// The only problem with rings is that you can keep calling ring.Next() indefinitely, so you need
	// to find a way to put a stop to that. In this case, this is accomplished with the help of the
	// ring.Len() function. This can also be done the ring.Do()
	for i := 0; i < myRing.Len()+4; i++ {
		myRing = myRing.Next()
		fmt.Print(myRing.Value, " ")
	}
	fmt.Println()
}
