package main

import "runtime"

type data struct {
	i, j int
}

func main() {

	var N = 40000000
	var structure []data

	// for loop is used for putting all values into structures that are
	// stored in the slice
	for i := 0; i < N; i++ {
		value := int(i)
		structure = append(structure, data{value, value})
	}

	runtime.GC()

	// - = structure[0] is used for preventing the garbage collector from collectin the structure
	// variable too early, as it is not referenced or used outside of the for loop.
	_ = structure[0]
}
