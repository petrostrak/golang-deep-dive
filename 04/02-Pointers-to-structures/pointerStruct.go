package main

import "fmt"

type myStruct struct {
	Name, Surname string
	Height        int32
}

// NewStructurePointer returns a memory address of a local variable (myStruct)
func NewStructurePointer(n, s string, h int32) *myStruct {
	if h > 300 {
		h = 0
	}
	return &myStruct{n, s, h}
}

// NewStructure is the no-pointer version of NewStructurePointer function
func NewStructure(n, s string, h int32) myStruct {
	if h > 300 {
		h = 0
	}
	return myStruct{n, s, h}
}

func main() {

	s1 := NewStructurePointer("Petros", "Trak", 173)
	s2 := NewStructure("Petros", "Trak", 173)
	fmt.Println((*s1).Name)
	fmt.Println(s1.Name)
	fmt.Println(s2.Name)

	// s1 is a pointer to a structure, which means that we will need to dereference
	// that pointer in order to use the object it points to
	fmt.Println(s1)

	// s2 is an entire structure object
	fmt.Println(s2)

	// new keyword allows you to allocate new objects. New returns the memory address of the
	// allocated object. New returns a pointer
	ps := new(myStruct)
	fmt.Println(ps)
	ps.Name = "petros"
	ps.Surname = "trak"
	ps.Height = 173
	fmt.Println(ps)

}
