package main

import "fmt"

// A rune is an int32 value, and therefore a Go type, that is used for representing a Unicode code
// point. A Unicode code point, or code position is a numerical value that is usually used for
// representing single Unicode characters, however, it can also have alternative meanings, such as
// providing formatting information.
//
// A run literal is a character in single quotes. We may also consider a rune literal as a rune
// constant. Behind the scenes, a rune literal is associated with a Unicode code point.
func main() {

	// € does not belong to the ASCII table of characters
	const r1 = '€'
	fmt.Println("(int32) r1:", r1)
	fmt.Printf("(HEX) r1: %x\n", r1)
	fmt.Printf("(as a String) r1: %s\n", r1)
	fmt.Printf("(as a character) r1: %c\n", r1)

	fmt.Println("A string is a collection of runes:", []byte("Petros"))
	aString := []byte("Petros")
	for x, y := range aString {
		fmt.Println(x, y)
		fmt.Printf("Char: %c\n", aString[x])
	}
	fmt.Printf("%s\n", aString)
}
