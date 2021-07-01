package main

import "runtime"

func main() {

	var N = 40000000
	split := make([]map[int]int, 200)

	// For loop for creating a hash of hashes
	for i := range split {
		split[i] = make(map[int]int)
	}

	// For loop for storing the desired data in the hash of hashes
	for i := 0; i < N; i++ {
		value := int(i)
		split[i%200][value] = value
	}

	runtime.GC()
	_ = split[0][0]

}
