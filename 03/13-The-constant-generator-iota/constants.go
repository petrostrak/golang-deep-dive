package main

import "fmt"

type Digit int
type Power2 int

const PI = 3.1415926

const (
	Zero Digit = iota
	One
	Two
	Three
	Four
)

// For p2_0, iota has the value of 0 and p2_0 is defined as 1. For p2_2, iota has the
// value of 2 and p2_2 is defined as the result of the expression 1 << 2, which is
// 00000100 in binary representation. The decimal value of 00000100 is 4, which is the
// value of p2_2. Analogously, the value of p2_4 is 16 and the value of p2_6 is 32.
//
// << left shift operator "times 2"
// n << x is n * 2, x times (1 << 5 = (1*2), 5 times => 32)
//
// >> right shift operator "devided by 2"
// n >> x is n / 2, x times (32 >> 5 = (32/2), 5 times => 1)
const (
	p2_0 Power2 = 1 << iota
	_
	p2_2
	_
	p2_4
	_
	p2_6
)

func main() {
	const s1 = 123
	var v1 float32 = s1 * 12
	fmt.Println(v1)
	fmt.Println(PI)
	fmt.Println(One)
	fmt.Println(Two)
	fmt.Println("2^0:", p2_0)
	fmt.Println("2^2:", p2_2)
	fmt.Println("2^4:", p2_4)
	fmt.Println("2^6:", p2_6)
}
