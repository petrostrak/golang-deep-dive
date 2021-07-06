package main

import "fmt"

// getPointer stores the result to the provided parameter
func getPointer(n *int) {
	*n = *n * *n
}

// returnPointer stores the result and requires a different variable for
// storing it.
func returnPointer(n int) *int {
	v := n * n
	return &v
}

func main() {

	i := -10
	j := 25

	// & gets the memory address of a non-pointer variable
	pI := &i
	pJ := &j

	fmt.Println("pI memory:", pI)
	fmt.Println("pJ memory:", pJ)

	// * gets the value of a pointer, which is called dereferencing the pointer
	fmt.Println("pI value:", *pI)
	fmt.Println("pJ value:", *pJ)

	*pI = 123456
	*pI--
	fmt.Println("i:", i)

	getPointer(pJ)
	fmt.Println("j:", j)
	k := returnPointer(12)
	fmt.Println(*k)
	fmt.Println(k)

}
