package main

import "fmt"

type Node struct {
	Value int
}

func (n *Node) String() string {
	return fmt.Sprint(n.Value)
}

type Stack struct {
	nodes []*Node
	count int
}

// Returns a new stack
func NewStack() *Stack {
	return &Stack{}
}

// Push adds a node to the Stack
func (s *Stack) Push(n *Node) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

// Pop removes and returns a node from the Stack in last to first order
func (s *Stack) Pop() *Node {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.nodes[s.count]
}

// A stack is an ordered list in which all insertions
// and deletions are made at one end, called the top.
// Stacks sometimes are referred to as Last In First
// Out (LIFO) lists.
func main() {

	s := NewStack()
	s.Push(&Node{Value: 34})
	s.Push(&Node{Value: 50})
	s.Push(&Node{Value: 61})
	s.Push(&Node{Value: 77})
	fmt.Println(s.Pop(), s.Pop(), s.Pop(), s.Pop())

}
