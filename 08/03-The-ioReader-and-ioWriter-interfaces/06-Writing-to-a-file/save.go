package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	s := []byte("Data to write\n")

	// Using fmt.Fprintf without buffer
	f1, err := os.Create("f1")
	if err != nil {
		fmt.Println("cannot create file")
		return
	}
	defer f1.Close()

	fmt.Fprintf(f1, string(s))

	// WriteString without buffer
	f2, err := os.Create("f2")
	if err != nil {
		fmt.Println("cannot create file")
		return
	}
	defer f2.Close()

	n, err := f2.WriteString(string(s))
	fmt.Printf("wrote %d bytes\n", n)

	// Using bufio.NewWriter
	f3, err := os.Create("f3")
	if err != nil {
		fmt.Println("cannot create file")
		return
	}
	defer f3.Close()

	w := bufio.NewWriter(f3)
	n, err = w.WriteString(string(s))
	if err != nil {
		fmt.Println("cannot write f3")
	}
	fmt.Printf("wrote %d bytes\n", n)
	w.Flush()

	// Using ioutil.Write
	f4 := "f4"
	if err = ioutil.WriteFile(f4, s, 0644); err != nil {
		fmt.Println("cannot write f4")
		return
	}

	// Using io.WriteString
	f5, err := os.Create("f5")
	if err != nil {
		fmt.Println("cannot create f5")
		return
	}
	defer f5.Close()

	n, err = io.WriteString(f5, string(s))
	if err != nil {
		fmt.Println("cannot write to f5")
		return
	}
	fmt.Printf("wrote %d bytes\n", n)

}
