package main

import (
	"fmt"
)

const (
	// SIZE constant holds the number of buckets of the hash table.
	SIZE = 15
)

type Node struct {
	Value int
	Next  *Node
}

// The actual hash is stored in a HashTable structure that has two fields. The first field is a
// map that associates an integer with a linked list (*Node), and the second field is the size of
// the hash table. As a result, this hash table will have as many linked lists as the number of
// its buckets. This also means that the node of each bucket of the hash table will be stored in
// linked lists.
type HashTable struct {
	Table map[int]*Node
	Size  int
}

// The hashFunction() uses the modulo operator. The main reason for choosing the
// modulo operator is because the particular hash table has to cope with integer
// values.
func hashFunction(i, size int) int {
	return (i % size)
}

// insert() is called for inserting elements into the hash table.
func insert(hash *HashTable, value int) int {
	index := hashFunction(value, hash.Size)
	element := Node{Value: value, Next: hash.Table[index]}
	hash.Table[index] = &element
	return index
}

// traverse() is used for printing all of the values in the hash table. The function visits
// each of the linked lists of the hash table and prints the stored values, linked list by
// linked list.
func traverse(hash *HashTable) {
	for k := range hash.Table {
		if hash.Table[k] != nil {
			t := hash.Table[k]
			for t != nil {
				fmt.Printf("%d -> ", t.Value)
				t = t.Next
			}
			fmt.Println()
		}
	}
}

// lookup allows us to determine whether a given element already exists in the hash table or not. It is
// based on the traverse()
func lookup(hash *HashTable, value int) bool {
	index := hashFunction(value, hash.Size)
	if hash.Table[index] != nil {
		t := hash.Table[index]
		for t != nil {
			if t.Value == value {
				return true
			}
			t = t.Next
		}
	}

	return false
}

func main() {

	// We create a new hash table usung the table variable which is a map that holds the buckets
	// of the hash table.
	table := make(map[int]*Node, SIZE)
	hash := &HashTable{Table: table, Size: SIZE}

	fmt.Println("Number of spaces:", hash.Size)
	for i := 0; i < 120; i++ {
		insert(hash, i)
	}
	traverse(hash)
}
