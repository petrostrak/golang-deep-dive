package main

import "fmt"

func main() {

	willClose := make(chan int, 10)

	willClose <- -1
	willClose <- 0
	willClose <- 2

	<-willClose
	<-willClose
	<-willClose
	close(willClose)

	read := <-willClose

	// Reading from a closed channel returns the zero value of its data type
	fmt.Println(read) // 0

}
