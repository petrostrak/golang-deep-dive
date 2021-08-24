package main

import "fmt"

// First we need to describe a node
type node struct {
	data int
	next *node
}

// then we need to describe the list
type linkedList struct {
	head   *node
	length int
}

// method that helps us insert values into the linkedList
//
// [Node|data, next] ->[Node|data, next] ->[Node|data, next] -> [Node|data, next]
// prepend func adds a Node at the begining of the linkList
// [temp. head] -> [Node|data, next] ->[Node|data, next] ->[Node|data, next] -> [Node|data, next]
func (l *linkedList) prepend(n *node) {
	// we make a temporary node and put the head
	second := l.head

	// then we insert the value in the head
	l.head = n

	// we point the new head to the second
	l.head.next = second

	// we increment the length of the linkedList
	l.length++
}

func (l linkedList) printListData() {
	toPrint := l.head
	for i := 0; i < l.length; i++ {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
	}
	fmt.Printf("\n")
}

// this func will delete a node from the linked list based on the
// the given value
//
// [Node|data, next] -> [Node|data, next]->   [XXX]   ->[Node|data, next] -> [Node|data, next]
//										  |__________|
func (l *linkedList) deleteWithValue(value int) {
	if l.length == 0 {
		fmt.Println("The list is empty")
		return
	}

	// if the value is in the head, then the second
	// node should be the head.
	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	previousToDelete := l.head
	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil {
			fmt.Printf("The value %d you are searching did not appear in the list\n", value)
			return
		}
		previousToDelete = previousToDelete.next
	}

	// if we find the node with the value, we simple reference the pointer
	// of the previous node to the next
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func main() {

	myList := linkedList{}
	n := &node{data: 34}
	n1 := &node{data: 12}
	n2 := &node{data: 5}
	n3 := &node{data: 1}
	n4 := &node{data: 93}
	n5 := &node{data: 76}
	n6 := &node{data: 32}
	myList.prepend(n)
	myList.prepend(n1)
	myList.prepend(n2)
	myList.prepend(n3)
	myList.prepend(n4)
	myList.prepend(n5)
	myList.prepend(n6)
	myList.deleteWithValue(32)
	myList.deleteWithValue(100)
	myList.printListData()
	emptyList := linkedList{}
	emptyList.deleteWithValue(10)

}
