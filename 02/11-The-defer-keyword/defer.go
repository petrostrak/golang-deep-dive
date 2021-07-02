package main

import "fmt"

// The defer keyword postpones the execution of a function until the surrounding function returns.
// It is very important to remember that deferred functions are executed in Last In, First Out (LIFO)
// order after the surrounding functions.
func main() {

	d1()
	d2()
	fmt.Println()
	d3()
	fmt.Println()

}

// 123
// d1() returns the three values of the i variable of the for loop in reverse order because deferred
// functions are executed in LIFO order.
func d1() {
	for i := 3; i > 0; i-- {
		defer fmt.Print(i, " ")
	}
}

// 000
// After the for loop has ended, the value of i is 0 because it is that value of i that made the loop
// terminate. However, the deferred anonymous function is evaluated after the for loop ends, because
// it has no parameters, which means that it is evaluated three times for an i value of 0.
func d2() {
	for i := 3; i > 0; i-- {
		defer func() {
			fmt.Print(i, " ")
		}()
		fmt.Println()
	}
}

// 123
// Due to the parameter of the anonymous function, each time the anonymous function is deferred, it gets
// and therefore uses the current value of i.
func d3() {
	for i := 3; i > 0; i-- {
		defer func(n int) {
			fmt.Print(n, " ")
		}(i)
		fmt.Println()
	}
}
