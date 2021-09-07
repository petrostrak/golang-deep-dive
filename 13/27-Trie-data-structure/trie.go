package main

import "fmt"

const (
	AlphabetSize = 26
)

// Node represents each node in the trie
type Node struct {
	children [AlphabetSize]*Node
	isEnd    bool
}

// Trie represent a trie and has a pointer to the root node
type Trie struct {
	root *Node
}

// InitTrie will create new Trie
func initTrie() *Trie {
	return &Trie{root: &Node{}}
}

// Insert

// Search

func main() {
	testTrie := initTrie()
	fmt.Println(testTrie.root)
}
