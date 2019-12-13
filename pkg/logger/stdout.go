package logger

import (
	"fmt"
)

type stdoutLogger struct {
}

func NewStdoutLogger() Logger {
	return stdoutLogger{}
}

func (l stdoutLogger) Error(err error) {
	fmt.Printf("[ERROR] %v", err)
}

func (l stdoutLogger) Errorf(format string, args ...interface{}) {
	fmt.Printf("[ERROR] "+format, args...)
}

func (l stdoutLogger) Warn(err error) {
	fmt.Printf("[WARN] %v", err)
}

func (l stdoutLogger) Warnf(format string, args ...interface{}) {
	fmt.Printf("[WARN] "+format, args...)
}
