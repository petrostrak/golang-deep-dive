package main

import (
	"fmt"
	"os"
	"time"
)

// go run parseTime.go 12:10
func main() {

	var myTime string
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s string\n", myTime)
		os.Exit(1)
	}

	myTime = os.Args[1]

	d, err := time.Parse("15:04", myTime)
	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Time:", d.Hour(), d.Minute())
	} else {
		fmt.Println(err)
	}

}
