/*
Package logger sets up globals variables of custim loggers
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
	Info = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	// Currently configured to standard error output
	Error = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
}
