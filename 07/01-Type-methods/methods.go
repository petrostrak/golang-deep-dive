package main

import "fmt"

type twoInts struct {
	X int64
	Y int64
}

func regularFunction(a, b twoInts) twoInts {
	return twoInts{a.X + b.X, a.Y + b.Y}
}

func (a twoInts) method(b twoInts) twoInts {
	return twoInts{a.X + b.X, b.X + b.Y}
}

func main() {

	i := twoInts{1, 2}
	j := twoInts{-5, -2}

	fmt.Println(regularFunction(i, j))
	fmt.Println(i.method(j))

}
