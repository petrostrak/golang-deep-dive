package main

import "fmt"

// The tower of Hanoi is a mathimatical puzzle which consists of three
// towers (pegs) and more thatn one rings. The puzzle starts with the
// disks in a neat stack in ascending order of size on one rod, the
// smallest at the top, thus making a conical shape.
//
// The objective of the puzzle is to move the entire stack to another
// rod.
//
// Moving of disks should follow the following rules:
//
// Only one disk can be moved at a time.
// A dist can only be moved if it is the uppermost disk on a stack
// No disk may be placed on top of a smaller disk
func main() {

	var t solver = new(towers)
	t.play(5)

}

type solver interface {
	play(int)
}

// towers is example of type satisfying solver interface
type towers struct{}

func (t *towers) play(n int) {
	t.moveN(n, 1, 2, 3)
}

func (t *towers) moveN(n, from, to, via int) {
	if n > 0 {
		t.moveN(n-1, from, via, to)
		t.moveM(from, to)
		t.moveN(n-1, via, to, from)
	}
}

func (t *towers) moveM(from, to int) {
	fmt.Println("Move disk from rod", from, "to rod", to)
}
