package main

import "fmt"

// Buffered channels allow the Go scheduler to put jobs in the queue quickly in order to
// be able to deal with more requests.
//
// The technique presented here works  as follows: All incoming requests are forwarded to a
// channel, which processes them one by one. When the channel is don processing a request, it
// sends a message to the original caller saying that it is ready to process a new one. So the
// capacity of the buffer of the channel restricts the number of simultaneous requests that
// it can keep.
func main() {

	// This gives the channel a place to store up to five integers
	numbers := make(chan int, 5)
	counter := 10

	// We try to put 10 integers in the numbers channel, but as the channel has room for only five
	// integers, we will not be able to store all 10 integers in it.
	for i := 0; i < counter; i++ {
		select {
		case numbers <- i:
		default:
			fmt.Println("Not enough space for", i)
		}
	}

	// We try to read the contents of the numbers channel. As long as there is something to read from
	// the numbers channel, the first branch of the select statement will get executed. As long as the
	// numbers channel is empty, the default branch will be executed.
	for i := 0; i < counter+5; i++ {
		select {
		case num := <-numbers:
			fmt.Println(num)
		default:
			fmt.Println("Nothing more to be done")
		}
	}

}
