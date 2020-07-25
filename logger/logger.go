package logger

import (
	"fmt"
	"log"
)

var logLevel = 0

/*
// SetOutput will se the logger output
func SetOutput(path string) (*os.File, error) {
	d, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	path = fmt.Sprintf("%v/%v", d, path)
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	log.SetOutput(f)
	return f, nil
}
*/

// ErrorErr will output an error to the log
func ErrorErr(err error) {
	Errorf("%v", err)
}

// Errorf will output an error to the log
func Errorf(format string, i ...interface{}) {
	msg := "[ERROR] " + fmt.Sprintf(format, i...)
	log.Println(msg)
}

// Debugf will output a debug message to the log
func Debugf(format string, i ...interface{}) {
	if logLevel > 0 {
		return
	}
	msg := "[DEBUG] " + fmt.Sprintf(format, i...)
	log.Println(msg)
}

// Message will write s to the log
func Message(s string) {
	log.Printf("[MESSAGE] %v\n", s)
}

// Messagef will write a formated message to the log
func Messagef(format string, i ...interface{}) {
	Message(fmt.Sprintf(format, i...))
}
