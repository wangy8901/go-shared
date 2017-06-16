package utils

import (
	"fmt"
	"log"
	"github.com/Sirupsen/logrus"
)

// An interface that accepts anything that could be logged.
type Logger interface {
	Info(text string, args ...interface{})
	Debug(text string, args ...interface{})
	Error(text string, args ...interface{})
	Fatal(text string, args ...interface{})
	Warn(text string, args ...interface{})
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

func (l defaultLogger) Warn(text string, args ...interface{}) {
	log.Printf("[WARN] %s\n", fmt.Sprintf(text, args...))
}

func (l defaultLogger) Custom(level, text string, args ...interface{}) {
	log.Printf("[%s] %s\n", level, fmt.Sprintf(text, args...))
}

func NewDefaultLogger() Logger {
	return defaultLogger{}
}

// Logrus logger
func NewLogrusLogger() Logger {
	return &logrusLogger{logger:logrus.New()}
}

func NewCustomLogrusLogger(logrus *logrus.Logger) Logger {
	return &logrusLogger{logger:logrus}
}

type logrusLogger struct {
	logger *logrus.Logger
}

func (l *logrusLogger) Info(text string, args ...interface{}) {
	l.logger.Infof(text, args...)
}

func (l *logrusLogger) Debug(text string, args ...interface{}) {
	l.logger.Debugf(text, args...)
}

func (l *logrusLogger) Error(text string, args ...interface{}) {
	l.logger.Errorf(text, args...)
}

func (l *logrusLogger) Fatal(text string, args ...interface{}) {
	l.logger.Fatalf(text, args...)
}

func (l *logrusLogger) Warn(text string, args ...interface{}) {
	l.logger.Warnf(text, args...)
}

func (l *logrusLogger) Custom(level, text string, args ...interface{}) {
	l.logger.Printf(fmt.Sprintf("[%s] %s", level, text), args...)
}