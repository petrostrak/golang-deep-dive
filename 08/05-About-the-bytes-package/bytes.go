package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {

	var buffer bytes.Buffer
	buffer.Write([]byte("This is"))
	fmt.Fprintf(&buffer, " a string!\n")
	// The first buffer.WriteTo() call will print the contents of the buffer variable.
	buffer.WriteTo(os.Stdout)
	// However, the second call to buffer.WriteTo() has nothing to print because the
	// buffer variable will be empty after the first buffer.WriteTo() call.
	buffer.WriteTo(os.Stdout)

	// The buffer.Reset() method resets the buffer variable and the Write() method puts some
	// data into it again
	buffer.Reset()
	buffer.Write([]byte("Mastering Go!"))
	r := bytes.NewReader([]byte(buffer.String()))
	fmt.Println(buffer.String())
	for {
		b := make([]byte, 3)
		n, err := r.Read(b)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("Read %s bytes: %d\n", b, n)
	}

}
