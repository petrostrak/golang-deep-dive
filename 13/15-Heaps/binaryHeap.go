package main

import "fmt"

// binary heap is an array
// parent index * 2 + 1 = left child index
// parent index * 2 + 2 = right child index
// [0]  * 2 + 1 = [1] left chind of first index
// [0]  * 2 + 2 = [2] right chind of first index
//
// MaxHeapInsert adds an element at the bottom
// rifht of the heap meaning at the last index.
// After that, we compare that element with the
// parent index and swap it it necessary. This
// process is called heapify.
//
// HeapExtractMax removes the root node. Then the
// last node becomes the root, and wecompare it with
// its child nodes and swap if necessary. Repeat the
// process if needed.
func main() {
	m := &MaxHeap{}
	myHeap := []int{10, 20, 30}
	for _, v := range myHeap {
		m.Insert(v)
		fmt.Println(m)
	}
}

// MaxHeap struct has a slice that holds the array
type MaxHeap struct {
	array []int
}

// Insert adds an element to the heap
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

// We want to swap when the current index is
// larger than parent.
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// swap keys in the array
func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

// get the parent index
func parent(i int) int {
	return (i - 1) / 2
}

// get the left index
func left(i int) int {
	return i*2 + 1
}

// get the right index
func right(i int) int {
	return i*2 + 2
}
