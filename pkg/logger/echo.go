package logger

import "github.com/labstack/echo"

type EchoLogger struct {
	e *echo.Echo
}

func NewEchoLogger(e *echo.Echo) Logger {
	return EchoLogger{e}
}

func (l EchoLogger) Error(err error) {
	l.e.Logger.Error(err)
}

func (l EchoLogger) Errorf(format string, args ...interface{}) {
	l.e.Logger.Errorf(format, args...)
}

func (l EchoLogger) Warn(err error) {
	l.e.Logger.Warn(err)
}

func (l EchoLogger) Warnf(format string, args ...interface{}) {
	l.e.Logger.Warnf(format, args...)
}
