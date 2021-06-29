package main

import (
	"fmt"
	"log"
	"os"
)

var (
	LOGFILE = "/home/petros_trak/github.com/petrostrak/golang-deep-dive/01/09-Printing-line-numbers-in-log-file/mGo.log"
)

func main() {
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	iLog := log.New(f, "customLogLineNumber", log.LstdFlags)

	// log.Lshortfile flag adds the full filename as well as the line number of the Go statement that printed
	// the log entry in the log entry itself.
	iLog.SetFlags(log.LstdFlags | log.Lshortfile)
	iLog.Println("Hello there!")
	iLog.Println("Another log entry!")
}
