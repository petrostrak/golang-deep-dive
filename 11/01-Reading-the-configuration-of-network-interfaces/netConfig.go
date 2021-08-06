package main

import (
	"fmt"
	"net"
)

func main() {

	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, i := range interfaces {
		fmt.Printf("Interface: %v\n", i.Name)
		byName, err := net.InterfaceByName(i.Name)
		if err != nil {
			fmt.Println(err)
		}
		addresss, err := byName.Addrs()
		if err != nil {
			fmt.Println(err)
		}
		for k, v := range addresss {
			fmt.Printf("Interface address #%v: %v\n", k, v.String())
		}
		fmt.Println()
	}

}
