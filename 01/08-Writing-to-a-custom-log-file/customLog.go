package main

import (
	"fmt"
	"log"
	"os"
)

var (
	LOGFILE = "/home/petros_trak/github.com/petrostrak/golang-deep-dive/01/08-Writing-to-a-custom-log-file/mGo.log"
)

// After executing twice, the log file will populate.
func main() {

	// Unix file permission 0644
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	iLog := log.New(f, "customLogLineNumber ", log.LstdFlags)

	// SetFlags allows you to set the output flags(options) for the current logger. The default values
	// as defined by LstdFlags are Ldate and Ltime, which means that we 'll get the current date and
	// time in each log entry we write in our log file.
	iLog.SetFlags(log.LstdFlags)
	iLog.Println("Hello there!")
	iLog.Println("Another log entry!")
}
