package main

import "fmt"

// Loop through an array and calculate the product of each
// element of the array, except self element. Do not use any
// division.
//
// eg. arr := []int{1,2,3,4}
// answer []int{24, 12, 8, 6}
//
// The solution would be, one loop will calculate the products
// on the left side of the index, another one all the products
// on the right side of the index and the last one will multiply
// both results into the returning array.
//
// input =  [1,2,3,4]
// leftProduct = [1,1,2,6]
// rightProduct = [24,12,4,1] (filling backwords)
// output = [24,12,8,6]
func main() {

	arr := []int{1, 2, 3, 4}
	fmt.Println(productOfArrayExceptSelf(arr))
	fmt.Println(productOfArrayExceptSelfButBetter(arr))

}

// O(2n) Complexity
// Plus no unnecessary arrays
func productOfArrayExceptSelfButBetter(arr []int) []int {
	result := make([]int, len(arr))

	result[0] = 1

	// we are using the output array to store all the left products
	// as before
	for i := 1; i < len(arr); i++ {
		result[i] = arr[i-1] * result[i-1]
	}

	// and for the right producs we use the variable R and do the multipication
	// on the fly
	R := 1
	for i := len(arr) - 1; i >= 0; i-- {
		result[i] = result[i] * R
		R = R * arr[i]
	}

	return result
}

// O(3n) Complexity
func productOfArrayExceptSelf(arr []int) []int {
	result := make([]int, len(arr))
	leftProducts := make([]int, len(arr))
	rightProducts := make([]int, len(arr))

	leftProducts[0] = 1
	rightProducts[len(arr)-1] = 1

	// this loop calculates all the products to the
	// left of the given array.
	for i := 1; i < len(arr); i++ {
		leftProducts[i] = arr[i-1] * leftProducts[i-1]
	}

	// this loop calculates all the products to the
	// right of the given array.
	for i := len(arr) - 2; i >= 0; i-- {
		rightProducts[i] = arr[i+1] * rightProducts[i+1]
	}

	// this loop multiplies the left and right products
	for i := 0; i < len(arr); i++ {
		result[i] = leftProducts[i] * rightProducts[i]
	}

	return result
}
