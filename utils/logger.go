package utils

import (
	"fmt"
	"log"
)

// An interface that accepts anything that could be logged.
type Logger interface {
	Info(text string, args ...interface{})
	Debug(text string, args ...interface{})
	Error(text string, args ...interface{})
	Fatal(text string, args ...interface{})
	Trace(text string, args ...interface{})
	// Method to accept a custom severity level
	Custom(level, text string, args ...interface{})
}

// A simple logger that simply prints to console via fmt
type defaultLogger struct{}

func (l defaultLogger) Info(text string, args ...interface{}) {
	log.Printf("[INFO] %s\n", fmt.Sprintf(text, args...))
}

func (l defaultLogger) Debug(text string, args ...interface{}) {
	log.Printf("[DEBUG] %s\n", fmt.Sprintf(text, args...))
}

func (l defaultLogger) Error(text string, args ...interface{}) {
	log.Printf("[ERROR] %s\n", fmt.Sprintf(text, args...))
}

func (l defaultLogger) Fatal(text string, args ...interface{}) {
	log.Panicf("[FATAL] %s\n", fmt.Sprintf(text, args...))
}

func (l defaultLogger) Trace(text string, args ...interface{}) {
	log.Printf("[TRACE] %s\n", fmt.Sprintf(text, args...))
}

func (l defaultLogger) Custom(level, text string, args ...interface{}) {
	log.Printf("[%s] %s\n", level, fmt.Sprintf(text, args...))
}

func NewDefaultLogger() Logger {
	return defaultLogger{}
}
