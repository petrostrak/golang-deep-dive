package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

func handleConnection(c net.Conn) {
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			os.Exit(100)
		}

		temp := strings.TrimSpace(string(netData))
		if temp == "STOP" {
			return
		}

		fibo := "-1\n"
		n, err := strconv.Atoi(temp)
		if err == nil {
			fibo = strconv.Itoa(f(n)) + "\n"
		}
		c.Write([]byte(string(fibo)))
	}
}

// The idea behind the algorithm used in f(), which uses a dynamic programming technique,
// is that whenever a Fibonacci number is computed, it is put into the fn map so that it
// will not be computed again. This simple idea saves a lot of time, especially when large
// Fibonacci numbers need to be calculated.
func f(n int) int {
	fn := make(map[int]int)
	for i := 0; i <= n; i++ {
		var f int
		if i <= 2 {
			f = 1
		} else {
			f = fn[i-1] + fn[i-2]
		}
		fn[i] = f
	}
	return fn[n]
}

// For each incoming connection to the TCP server, the program will start a new goroutine
// to handle that request.
//
// The job of this TCP conccurent server is to eccept a positive integer and return a natural
// number from fibonacci sequence.
func main() {

	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number")
		return
	}

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
	}

}
