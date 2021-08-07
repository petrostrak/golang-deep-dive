package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Need a domain name")
		return
	}

	domain := arguments[1]

	// All the work is done by the net.LookupNS() function, which returns the NS records of a
	// domain as a slice variable of the net.NS type. This is the reason for printing the Host field
	// of each net.NS
	NSs, err := net.LookupNS(domain)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, NS := range NSs {
		fmt.Println(NS.Host)
	}

}
