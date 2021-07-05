package main

import "fmt"

func main() {
	iMap := make(map[string]int)
	iMap["k1"] = 12
	iMap["k2"] = 13
	fmt.Println("iMap:", iMap)

	anotherMap := map[string]int{
		"k1": 12,
		"k2": 13,
	}
	fmt.Println("anotherMap:", anotherMap)

	delete(anotherMap, "k1")
	delete(anotherMap, "k1")
	delete(anotherMap, "k1")
	fmt.Println("anotherMap:", anotherMap)

	// Here we see a technique that allows you to determine whether a given key is in
	// the map or not. This is vital technique because without it we would not know
	// whether a given map has the required information or not.
	//
	// The bad thing is that if you try to get the value of a map key that does not
	// exist in the map, you will end up getting zero, which gives you no way of
	// determining whether the result was zero because the key you requested was not
	// there, or because the element with the corresponding key actually had a zero
	// value. That is why we have _, ok in maps.
	_, ok := iMap["doesItExist"]
	if !ok {
		fmt.Println("Does not exist!")
	} else {
		fmt.Println("It exists!")
	}

	for key, value := range iMap {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
}
