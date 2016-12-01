package main

import (
	"bytes"
	"fmt"
	"log"
	"time"
)

var logger *log.Logger

func init() {
	var buf bytes.Buffer
	logger = log.New(&buf, "logger: ", log.Lshortfile)
	logger.Print("Hello, log file!")
}

//LogThis log messages to a file and writes them to the console
func LogThis(msg string) {
	t := time.Now()
	logger.Printf("%d %s\n", t, msg)
	fmt.Printf("%d %s\n", t, msg)
}
