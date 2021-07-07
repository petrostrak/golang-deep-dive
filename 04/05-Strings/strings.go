package main

import (
	"fmt"
)

func main() {
	const sLiteral = "\x99\x42\x32\x55\x50\x35\x23\x50\x29\x9c"
	fmt.Println(sLiteral)

	// %x returns the AB part of a \xAB
	fmt.Printf("x: %x\n", sLiteral)

	for i := 0; i < len(sLiteral); i++ {
		fmt.Printf("%x ", sLiteral)
	}
	fmt.Println()

	// %q will return a double-quoted string
	fmt.Printf("q: %q\n", sLiteral)

	// %+q will guarantee ASCII-only output
	fmt.Printf("+q: %+q\n", sLiteral)

	// % x (note the space between the % and the c) will put spaces between the
	// printed bytes.
	fmt.Printf(" x: % x\n", sLiteral)

	fmt.Printf("s: As a string: %s\n", sLiteral)

	s2 := "€£³"
	for x, y := range s2 {

		// %#U will print the characters in the U+0058 format
		fmt.Printf("%#U starts at byte position %d\n", y, x)
	}

	// As s2 variable contains Unicode characters, its size is larger than
	// the number of characters in it.
	fmt.Printf("s2 length: %d\n", len(s2))

	const s3 = "ab12AB"
	fmt.Println("s3:", s3)
	fmt.Printf("x: % x\n", s3)

	fmt.Printf("s3 length: %d\n", len(s3))

	for i := 0; i < len(s3); i++ {
		fmt.Printf("%x ", s3[i])
	}
	fmt.Println()

}
