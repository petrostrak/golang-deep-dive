package main

import "fmt"

// Node structure will be used for the nodes of the linked list
type Node struct {
	Value int
	Next  *Node
}

var (
	// root is a global variable that holds the first element of the linked list,
	// which will be available everywhere in the code.
	root = new(Node)
)

// addNode() is used for adding new nodes to the linked list.
// In the first case, we test whether we dealing with an empty linked list or not.
// In the second case we check whether the value that we want to add is already in the
// list (Due to the way that linked lists work, they do not normally contain duplicates)
// In the third case, we check whether you have reached the end of the linked list. In
// this case, we add a new code at the end of the list with the desired value using
// t.Next = &Node{v, nil}. If node of these conditions is true, we repeat the same process
// with the addNode() for the next node of the linked list using return addNode(t.Next, v)
func addNode(t *Node, v int) int {
	if t == nil {
		t = &Node{v, nil}
		root = t
		return 0
	}
	if v == t.Value {
		fmt.Println("Node already exists:", v)
		return -1
	}
	if t.Next == nil {
		t.Next = &Node{v, nil}
		return -2
	}

	return addNode(t.Next, v)
}

func deleteNode(t *Node, v int) {
	if !lookupNode(t, v) {
		fmt.Printf("Node with value %d does not exist!", v)
	}
	current := t
	for current.Next.Next != nil {
		current.Value = current.Next.Value
		current = current.Next
	}
	current.Value = current.Next.Value
	current.Next = nil
}

func traverse(t *Node) {
	if t == nil {
		fmt.Println("-> Empty list!")
		return
	}
	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}
	fmt.Println()
}

// lookupNode checks whether a given element exists in the linked list
func lookupNode(t *Node, v int) bool {
	if root == nil {
		t = &Node{v, nil}
		root = t
		return false
	}
	if v == t.Value {
		return true
	}
	if t.Next == nil {
		return false
	}

	return lookupNode(t.Next, v)
}

// size returns the size of the linked list, which is the number of nodes in the linked list.
func size(t *Node) int {
	if t == nil {
		fmt.Println("-> Empty list!")
		return 0
	}

	i := 0
	for t != nil {
		i++
		t = t.Next
	}
	return i
}

func main() {
	fmt.Println(root)
	root = nil
	traverse(root)
	addNode(root, 1)
	addNode(root, -1)
	traverse(root)
	addNode(root, 10)
	addNode(root, 5)
	addNode(root, 45)
	addNode(root, 5)
	addNode(root, 5)
	traverse(root)
	addNode(root, 100)
	traverse(root)

	if lookupNode(root, 100) {
		fmt.Println("Node exists!")
	} else {
		fmt.Println("Node does not exist!")
	}

	if lookupNode(root, -100) {
		fmt.Println("Node exists!")
	} else {
		fmt.Println("Node does not exist!")
	}
}
