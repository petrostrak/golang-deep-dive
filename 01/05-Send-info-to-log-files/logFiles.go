package main

import (
	"fmt"
	"log"
	"log/syslog"
	"os"
	"path/filepath"
)

// The Go code of logFiles.go will explain the use of the log and log/syslog packages
// to write to the system log files.
//
// grep LOG_MAIL /var/log/mail.log
// grep LOG_LOCAL7 /var/log/syslog
func main() {
	programName := filepath.Base(os.Args[0])

	// The first parameter to the syslog.New() is the priority, which is a combination of
	// the logging facility and the logging level. The second parameter is the name of the
	// process that will appear on the logs as the sender of the message
	sysLog, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL7, programName)
	if err != nil {
		log.Fatal(err)
	} else {

		// sets the output destination of the default logger, which is the logger we created
		// earlier on sysLog. Then we can use log.Println() to send information to the log
		// server.
		log.SetOutput(sysLog)
	}

	log.Println("LOG_INFO + LOG_LOCAL7: Logging in Go!")

	sysLog, err = syslog.New(syslog.LOG_MAIL, "some program!")
	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(sysLog)
	}

	log.Println("LOG_MAIL: Logging in Go!")
	fmt.Println("Will this be saved in log?")
}
