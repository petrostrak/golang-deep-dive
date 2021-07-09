package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// The idea behind a key-value store is modest, answer queries fast and work as fast
// as possible. This translates into using simple algorithms and simple data structures.
//
// This program will implement the four fundamental tasks of a key-value store
// Adding a new element (ADD)
// Deleting an existing element from the key-value store based on a key (DELETE)
// Looking up the value of a specific key in the store (LOOKUP)
// Changing the value of an existing key (CHANGE)

type myElement struct {
	Name, Surname, Id string
}

var (
	DATA = make(map[string]myElement)
)

func LOOKUP(k string) *myElement {

	// This allows us to determine whether a given key is in the map
	// or not. The bad thing is that if we try to get the value of
	// a map key that does not exist in the map, we end up getting
	// zero, which we dont know its the actual value or not. That is
	// why we have _, ok.
	_, ok := DATA[k]
	if ok {
		n := DATA[k]
		return &n
	} else {
		return nil
	}

}

func ADD(k string, n myElement) bool {

	if k == "" {
		return false
	}

	if LOOKUP(k) == nil {
		DATA[k] = n
		return true
	}

	return false
}

func DELETE(k string) bool {

	if LOOKUP(k) != nil {
		delete(DATA, k)
		return true
	}

	return false
}

func CHANGE(k string, n myElement) bool {
	DATA[k] = n
	return true
}

func PRINT() {
	for k, v := range DATA {
		fmt.Printf("key: %s, value: %v\n", k, v)
	}
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		text = strings.TrimSpace(text)
		tokens := strings.Fields(text)

		switch len(tokens) {
		case 0:
			continue
		case 1:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 2:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 3:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 4:
			tokens = append(tokens, "")
		}

		switch tokens[0] {
		case "PRINT":
			PRINT()
		case "STOP":
			return
		case "DELETE":
			if !DELETE(tokens[1]) {
				fmt.Println("Delete operation failed!")
			}
		case "ADD":
			n := myElement{tokens[2], tokens[3], tokens[4]}
			if !ADD(tokens[1], n) {
				fmt.Println("Add operation failed!")
			}
		case "LOOKUP":
			n := LOOKUP(tokens[1])
			if n != nil {
				fmt.Printf("%v\n", *n)
			}
		case "CHANGE":
			n := myElement{tokens[2], tokens[3], tokens[4]}
			if !CHANGE(tokens[1], n) {
				fmt.Println("Update operation failed!")
			}
		default:
			fmt.Println("Unknown command - please try again!")
		}
	}

}
