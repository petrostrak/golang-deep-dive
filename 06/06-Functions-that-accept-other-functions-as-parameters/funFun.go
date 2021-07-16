package main

import "fmt"

func func1(i int) int {
	return i + i
}

func func2(i int) int {
	return i * i
}

func funFun(f func(int) int, v int) int {
	return f(v)
}

func main() {

	fmt.Println("func1:", funFun(func1, 123))
	fmt.Println("func2:", funFun(func2, 123))
	fmt.Println("Inline func:", funFun(func(i int) int { return i + i/2 }, 123))

}
