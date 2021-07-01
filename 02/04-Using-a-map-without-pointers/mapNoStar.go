package main

import "runtime"

func main() {

	var N = 40000000
	myMap := make(map[int]int)

	// for loop is used for putting all values into structures that are
	// stored in the map
	for i := 0; i < N; i++ {
		value := int(i)
		myMap[value] = value
	}

	runtime.GC()

	// - = myMap[0] is used for preventing the garbage collector from collectin the map
	// variable too early, as it is not referenced or used outside of the for loop.
	_ = myMap[0]

}
