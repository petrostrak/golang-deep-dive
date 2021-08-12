package main

import (
	"fmt"
	"math"
	"net"
	"net/rpc"
	"os"

	sharedrpc "github.com/petrostrak/golang-deep-dive/12/03-Remote-Procedure-Call/sharedRPC"
)

type MyInterface struct{}

func Power(x, y float64) float64 {
	return math.Pow(x, y)
}

func (t *MyInterface) Multiply(arguments *sharedrpc.MyFloats, reply *float64) error {
	*reply = arguments.A1 * arguments.A2
	return nil
}

func (t *MyInterface) Power(arguments *sharedrpc.MyFloats, reply *float64) error {
	*reply = math.Pow(arguments.A1, arguments.A2)
	return nil
}

func main() {

	PORT := ":1234"
	arguments := os.Args
	if len(arguments) != 1 {
		PORT = ":" + arguments[1]
	}

	myInterface := new(MyInterface)

	// What makes this programm an RPC server is the use of the rpc.Register(). However,
	// as the RPC server uses TCP, it still needs to make function calls to net.ResolveTCPAddr().
	rpc.Register(myInterface)
	t, err := net.ResolveTCPAddr("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}

	l, err := net.ListenTCP("tcp4", t)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}
		fmt.Printf("%s\n", c.RemoteAddr())
		rpc.ServeConn(c)
	}

}
