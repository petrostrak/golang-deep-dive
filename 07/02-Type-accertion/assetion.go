package main

import "fmt"

// A type assertion is the x.(T) notation, where x is an interface type and T is a
// type. Additionally, the actual value stored in x is of type T and T must satisfly
// the interface type of x.
//
// Type assertions help us do two things. The first is checking whether an interface
// value keeps a particular type. When used this way, a type assertion returns two
// values: the underlyning value and a bool value. Although the underlyning value
// is what you might want to use, the bool value tells you whether the type was
// successful or not.
//
// The second thing that type assertions help us with is allowing us to use the concrete
// value stored in an interface or assign it to a new variable. This means that if there
// is an int variable in an interface, we can get that value using type assertion.
func main() {

	var myInt interface{} = 123

	k, ok := myInt.(int)
	if ok {
		fmt.Println("Success:", k)
	}

	v, ok := myInt.(float64)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("failed without panicking!")
	}

	i := myInt.(int)
	fmt.Println("No checking", i)

	j := myInt.(bool)
	fmt.Println(j)

}
