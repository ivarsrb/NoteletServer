/*
Package logger sets up global variables of custom loggers
*/
package logger

import (
	"log"
	"os"
)

var (
	// Info is used to log information level messages
	Info *log.Logger
	// Error is used to log error level messages
	Error *log.Logger
)

func init() {
	// Currently configured to standard output
	Info = log.New(os.Stdout, "INFO ", log.Ldate|log.Ltime)
	// Currently configured to standard error output
	Error = log.New(os.Stderr, "ERROR ", log.Ldate|log.Ltime|log.Lshortfile)
}
