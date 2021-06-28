package main

import (
	"fmt"
	"log"
	"log/syslog"
)

// log.Fatal() is used when something really bad has happened and you just want to
// exit your program as fast as possible after reporting the bad situation.
func main() {
	sysLog, err := syslog.New(syslog.LOG_ALERT|syslog.LOG_MAIL, "some program!")
	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(sysLog)
	}

	// log.Fatal() terminates a Go programm at the point where log.Fatal() was called.
	log.Fatal(sysLog)
	fmt.Println("Will you see this?")
}
