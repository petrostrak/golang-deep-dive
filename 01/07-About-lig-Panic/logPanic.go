package main

import (
	"fmt"
	"log"
	"log/syslog"
)

// There are situation where a program will fail for good and we want to have
// as much info about the failure as possible.
func main() {
	sysLog, err := syslog.New(syslog.LOG_ALERT|syslog.LOG_MAIL, "some program!")
	if err != nil {
		log.Panic(err)
	} else {
		log.SetOutput(sysLog)
	}

	// The output of log.Panic includes additional low-level info that will hopefully
	// help us to resolve difficult situations that happened in our Go code.
	//
	// Analogous to the log.Fatal, the use of log.Panic will add an entry to the proper
	// log file and will immediately terminate the Go program.
	log.Panic(sysLog)
	fmt.Println("Will you see this?")
}
