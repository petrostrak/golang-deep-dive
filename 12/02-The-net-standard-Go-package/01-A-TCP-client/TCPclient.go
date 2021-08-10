package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide host:port.")
		return
	}

	CONNECT := arguments[1]

	// The net.Dial() is used to connect to the remote server. The
	// first parameter of the net.Dial() defines the network that
	// will be used, while the second parameter defines the server
	// address.
	//
	// Valid values for the first parameter are tcp, tcp4(IPv4-only),
	// tcp6(IPv6-only), udp, udp4(IPv4-only), udp6(IPv6-only), ip,
	// ip4(IPv4-only), ip6(IPv4-only), unix(UNIX sockets), unixgram
	// and unixpacket.
	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(c, text+"\n")
		message, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Print("->: ", message)
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}
	}

}
