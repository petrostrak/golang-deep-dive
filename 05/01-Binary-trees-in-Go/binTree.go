package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Here is the definition of the node of the tree
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// traverse() reveals how you can visit all of the nodes of a binary tree using
// recursion.
func traverse(t *Tree) {
	if t == nil {
		return
	}
	traverse(t.Left)
	fmt.Print(t.Value, " ")
	traverse(t.Right)
}

// create() is only used for populating the binary tree with random integers
func create(n int) *Tree {
	var t *Tree
	rand.Seed(time.Now().Unix())
	for i := 0; i < 2*n; i++ {
		temp := rand.Intn(n * 2)
		t = insert(t, temp)
	}

	return t
}

// insert() does many important things using if statements. The first if statement
// checks whether we are dealing with an empty tree or not. If it is indeed an empty
// tree, then the new node will be the root of the tree and will be created as
// &Tree{nil, v, nil}
//
// The second if statement determines whether the value you are trying to insert already
// exists in the binary tree or not. If it exists, the function returns without doing
// anything else.
//
// The third if statement determines whether the value you are trying to insert will go
// on the left or on the right of the node that is currently being examined and acts
// accordingly.
// This implementation create unbalanced trees
func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v == t.Value {
		return t
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
	}
	t.Right = insert(t.Right, v)
	return t
}

func main() {
	tree := create(10)
	fmt.Println("The value of the root of the tree is", tree.Value)

	traverse(tree)
	fmt.Println()
	tree = insert(tree, -10)
	tree = insert(tree, -2)
	traverse(tree)
	fmt.Println()
	fmt.Println("The value of the root of the tree is", tree.Value)
}
