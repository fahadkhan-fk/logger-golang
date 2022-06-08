package main

import (
	// useful feature of Go’s import statement are aliases we used it as `log` in import here
	log "github.com/sirupsen/logrus"
)

func main() {

	// different variants of log as supported by this library `github.com/sirupsen/logrus`
	log.WithFields(log.Fields{
		"lion": "hunt",
	}).Warn("A lion will hunt you")

	// log.Info
	log.Info("This is a log information")

	// log.Warn
	log.Warn("This is a log warning")

	// log.Fatal
	log.Fatal("This is really bad, exiting ..")

	// just run this above code by `go run main.go` in your terminal and you will se the results.

	// we will see what now happens !
	//  just having these lines of code for testing purpose!
	// hehe new tang branch added !
	// hehe add new qa branch now!
}
