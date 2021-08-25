package main

import "fmt"

var (
	count int
)

type node struct {
	Key   int
	Left  *node
	Right *node
}

func (n *node) insert(k int) {
	if k > n.Key {
		if n.Right == nil {
			n.Right = &node{Key: k}
		} else {
			n.Right.insert(k)
		}
	} else if k < n.Key {
		if n.Left == nil {
			n.Left = &node{Key: k}
		} else {
			n.Left.insert(k)
		}
	}
}

func (n *node) search(k int) bool {
	count++

	// it means that we reached the end
	if n == nil {
		return false
	}

	if k > n.Key {
		return n.Right.search(k)
	} else if k < n.Key {
		return n.Left.search(k)
	}

	return true
}

func main() {

	myTree := &node{Key: 100}
	myTree.insert(160)
	myTree.insert(90)
	myTree.insert(23)
	myTree.insert(122)
	fmt.Println(myTree)
	fmt.Println(myTree.search(23))
	fmt.Println(count)

}
