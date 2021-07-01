package main

// #cgo CFLAGS: -I${SRCDIR}/callClib
// #cgo LDFLAGS: ${SRCDIR}/callC.a
// #include <stdlib.h>
// #include <callC.h>
import "C"
import (
	"fmt"
	"unsafe"
)

// Inside callClib directory we run gcc -c *.c to create callC.o
// file callC.o
// /usr/bin/ar rs callC.a *.o
// file callC.a
// rm callC.o
//
// Move callC.a to the directory of .go file and then
// go build callC.go
//
// If you are going to call a small amount of C code, then using a single Go file
// for both C and Go code is highly recommended because of its simplicity. However,
// if you are going to do something more complex and advanced, creating a static
// C library should be the preferred way.
func main() {

	fmt.Println("Going to call a C function!")
	C.cHello()

	fmt.Println("Going to call another C function!")
	myMessage := C.CString("This is Petros")
	defer C.free(unsafe.Pointer(myMessage))
	C.printMessage(myMessage)

	fmt.Println("All perfectly done!")
}
