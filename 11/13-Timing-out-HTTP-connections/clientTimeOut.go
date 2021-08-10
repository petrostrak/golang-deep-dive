package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var (
	timeout = time.Duration(time.Second)
)

func Timeout(network, host string) (net.Conn, error) {
	conn, err := net.DialTimeout(network, host, timeout)
	if err != nil {
		return nil, err
	}

	// The SetDeadline() is used by the net package to set the read and
	// write deadlines of a given network connection. Due to the way the
	// SetDeadline() works, you will need to call SetDeadline() before
	// any read or write operation.
	conn.SetDeadline(time.Now().Add(timeout))
	return conn, nil
}

func main() {

	if len(os.Args) == 1 {
		fmt.Printf("Usage: %s URL TIMEOUT\n", filepath.Base(os.Args[0]))
		return
	}

	if len(os.Args) == 3 {
		temp, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Using default timeout")
		} else {
			timeout = time.Duration(time.Duration(temp) * time.Second)
		}
	}

	URL := os.Args[1]
	t := http.Transport{Dial: Timeout}
	client := http.Client{
		Transport: &t,
	}

	data, err := client.Get(URL)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		defer data.Body.Close()
		if _, err := io.Copy(os.Stdout, data.Body); err != nil {
			fmt.Println(err)
			return
		}
	}

}
