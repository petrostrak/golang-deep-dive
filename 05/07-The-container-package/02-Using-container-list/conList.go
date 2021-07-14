package main

import (
	"container/list"
	"fmt"
	"strconv"
)

// printList allows us to print the contents of a list.List variable passed as a pointer. The Go code shows us
// how to print the elements of list.List starting from the first element and going to the last and vice versa.
//
// t.Prev() and t.Next() functions  allow you to iterate over the elements of a list backwards and forward.
func printList(l *list.List) {
	for t := l.Back(); t != nil; t = t.Prev() {
		fmt.Print(t.Value, " ")
	}
	fmt.Println()

	for t := l.Front(); t != nil; t = t.Next() {
		fmt.Print(t.Value, " ")
	}
	fmt.Println()
}

func main() {
	values := list.New()

	// PushBack allows you to insert an object at the back of a linked list.
	// The returned value is the element inserted in the list.
	e1 := values.PushBack("One")
	e2 := values.PushBack("Two")

	// PushFront allows you to insert an object at the front of a list.
	// The returned value is the element inserted in the list.
	values.PushFront("Three")

	// If you want to insert a new element before a specific object, use the InsertBefore().
	// If the object does not exist, the list will not change.
	values.InsertBefore("Four", e1)

	// If you want to insert a new element after a specific object, use the InsertAfter().
	// If the object does not exist, the list will not change.
	values.InsertAfter("Five", e2)
	values.Remove(e2)
	values.Remove(e2)
	values.InsertAfter("FiveFive", e2)

	// PushBackList puts a copy of an existing list at the front of another list.
	values.PushBackList(values)

	printList(values)

	// Init() either empties an existing list or initializes a new list.
	values.Init()

	fmt.Printf("After Init(): %v\n", values)

	for i := 0; i < 20; i++ {
		values.PushFront(strconv.Itoa(i))
	}

	printList(values)
}
