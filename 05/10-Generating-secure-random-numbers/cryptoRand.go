package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
)

func generateBytes(n int64) ([]byte, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return nil, err
	}

	return b, nil
}

func generatePass(s int64) (string, error) {
	b, err := generateBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}

// Should you wish to generate more secure pseudo-random numbers in Go, you should use the
// crypto/rand package, which implements a cryptographically secure pseudo-random number
// generator.
func main() {

	var LENGTH int64 = 8
	arguments := os.Args
	switch len(arguments) {
	case 2:
		LENGTH, _ = strconv.ParseInt(arguments[1], 10, 64)
		if LENGTH <= 0 {
			LENGTH = 8
		}
	default:
		fmt.Println("Using default values!")
	}

	myPass, err := generatePass(LENGTH)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(myPass[0:LENGTH])

}
