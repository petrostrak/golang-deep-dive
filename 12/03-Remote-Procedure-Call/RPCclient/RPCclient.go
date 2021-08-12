package main

import (
	"fmt"
	"net/rpc"
	"os"

	sharedrpc "github.com/petrostrak/golang-deep-dive/12/03-Remote-Procedure-Call/sharedRPC"
)

func main() {

	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a host:port string")
		return
	}

	CONNECT := arguments[1]
	c, err := rpc.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}

	args := sharedrpc.MyFloats{16, -0.5}
	var reply float64

	if err := c.Call("MyInterface.Multiply", args, &reply); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Reply (Multiply): %f\n", reply)

	if err := c.Call("MyInterface.Power", args, &reply); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Reply (Power): %f\n", reply)

}
