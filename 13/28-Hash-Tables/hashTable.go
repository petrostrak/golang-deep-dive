package main

import "fmt"

const (
	ArraySize = 7
)

// HashTable structure
type HashTable struct {
	array [ArraySize]*bucket
}

// bucketNode - linkedlist
type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	key  string
	next *bucketNode
}

// Insert will take in a key and add it to the hash table array
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

// Search will take in a key and return true if that key is stored in the hash table
func (h *HashTable) Search(key string) bool {
	// index := hash(key)
	return false
}

// Delete will take in a key and delete it from the hash table
func (h *HashTable) Delete(key string) {
	// index := hash(key)
}

// Insert will take in a key, create a node with the key and insert the node in the bucket
func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println(k, "already exists!")
	}

}

// Search will take in a key and return true if the bucket has a match
func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// Delete will take in a key and delete the node from the bucket
func (b *bucket) delete(k string) {

	if b.head.key == k {
		b.head = b.head.next
		return
	}

	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == k {
			// delete
			previousNode.next = previousNode.next.next
		}
		previousNode = previousNode.next
	}
}

// hash function
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}

	// hashCode is the returned value of the hash function
	return sum % ArraySize
}

// init with create a bucket in each slot of the hash table
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}

	return result
}

// Constant Complexity with separate chaining
func main() {

	// testHastTable := Init()
	// fmt.Println(testHastTable)
	// fmt.Println(hash("RANDY"))

	testBucket := &bucket{}
	testBucket.insert("RANDY")
	testBucket.delete("RANDY")

	fmt.Println(testBucket.search("RANDY"))

}
