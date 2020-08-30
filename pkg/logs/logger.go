package logs

import (
	"log"
	"os"
)

var (
	debug     = false
	errLogger *log.Logger
)

// Init initializes the loggers
func Init(dbg bool) {
	errLogger = log.New(os.Stderr, "ERROR: ", log.LstdFlags)
	debug = dbg
	if debug {
		log.SetPrefix("DEBUG: ")
	}
}

// Debugf will output a debugging log if debug logs are configured
func Debugf(format string, args ...interface{}) {
	if debug {
		log.Printf(format, args...)
	}
}

// Error logs a message with an error
func Error(msg string, err error) {
	errLogger.Printf(msg+": %v", err)
}

// Panic logs a message with an error then panics
func Panic(msg string, err error) {
	log.Panicf(msg+": %v", err)
}

// Panicf logs a message then panics
func Panicf(format string, args ...interface{}) {
	log.Panicf(format, args...)
}