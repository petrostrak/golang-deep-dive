package main

import (
	"fmt"
	"log"
	"os"
)

var (
	LOGFILE = "/home/petros_trak/github.com/petrostrak/golang-deep-dive/02/12-Using-defer-for-logging/mGo.log"
)

func one(aLog *log.Logger) {
	aLog.Println("--FUNCTION-one--")
	defer aLog.Println("--FUNCTION-one--")

	for i := 0; i < 10; i++ {
		aLog.Println(i)
	}
}

func two(aLog *log.Logger) {
	aLog.Println("--FUNCTION-two--")
	defer aLog.Println("--FUNCTION-two--")

	for i := 10; i > 0; i-- {
		aLog.Println(i)
	}
}

func main() {
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	iLog := log.New(f, "logDefer", log.LstdFlags)
	iLog.SetFlags(log.LstdFlags | log.Lshortfile)
	iLog.Println("Hello there!")
	iLog.Println("Another log entry!")

	one(iLog)
	two(iLog)
}
