package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

var (
	// Having a variable size for keeping the number of nodes that you have on the queue is handy
	// but not compulsory.
	size  = 0
	queue = new(Node)
)

func Push(t *Node, v int) bool {
	// if queue is empty, the the new node will become the queue
	if queue == nil {
		queue = &Node{v, nil}
		size++
		return true
	}

	// if the queue is not empty, we create a new node that is placed in front of the current queue.
	t = &Node{v, nil}
	t.Next = queue

	// After that, the head of the queue becomes the node that was just created.
	queue = t
	size++

	return true
}

func Pop(t *Node) (int, bool) {

	if size == 0 {
		return 0, false
	}

	// if the queue has only 1 node, we ectract the value of that node and the queue becomes empty.
	if size == 1 {
		queue = nil
		size--
		return t.Value, true
	}

	// otherwise we extract the last element of the queue...
	temp := t
	for t.Next != nil {
		temp = t
		t = t.Next
	}

	// ...remove the last node of the queue and fix the required pointers before returning the
	// desired value.
	v := temp.Next.Value
	temp.Next = nil

	size--
	return v, true

}

func traverse(t *Node) {
	if size == 0 {
		fmt.Println("Empty Queue!")
		return
	}

	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}
	fmt.Println()
}

func main() {
	queue = nil
	Push(queue, 10)
	fmt.Println("Size:", size)
	traverse(queue)

	v, b := Pop(queue)
	if b {
		fmt.Println("Pop:", v)
	}
	fmt.Println("Size:", size)

	for i := 0; i < 5; i++ {
		Push(queue, i)
	}

	traverse(queue)
	fmt.Println("Size:", size)

	v, b = Pop(queue)
	if b {
		fmt.Println("Pop:", v)
	}
	fmt.Println("Size:", size)

	v, b = Pop(queue)
	if b {
		fmt.Println("Pop:", v)
	}
	fmt.Println("Size:", size)
	traverse(queue)
}
