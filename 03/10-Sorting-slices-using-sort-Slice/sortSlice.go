package main

import (
	"fmt"
	"sort"
)

type Person struct {
	name   string
	height int
	weight int
}

func main() {
	mySlice := make([]Person, 0)
	mySlice = append(mySlice, Person{name: "Petros", height: 173, weight: 70})
	mySlice = append(mySlice, Person{name: "Maggie", height: 170, weight: 60})
	mySlice = append(mySlice, Person{name: "John", height: 175, weight: 75})
	mySlice = append(mySlice, Person{name: "Deppy", height: 173, weight: 80})

	fmt.Println("0:", mySlice)

	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].height < mySlice[j].height
	})
	fmt.Println("height ascending:", mySlice)

	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].weight > mySlice[j].weight
	})
	fmt.Println("weight descending:", mySlice)
}
