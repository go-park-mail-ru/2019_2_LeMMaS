package logger

import "github.com/labstack/echo"

type echoLogger struct {
	e *echo.Echo
}

func NewEchoLogger(e *echo.Echo) Logger {
	return echoLogger{e}
}

func (l echoLogger) Error(err error) {
	l.e.Logger.Error(err)
}

func (l echoLogger) Errorf(format string, args ...interface{}) {
	l.e.Logger.Errorf(format, args...)
}

func (l echoLogger) Warn(err error) {
	l.e.Logger.Warn(err)
}

func (l echoLogger) Warnf(format string, args ...interface{}) {
	l.e.Logger.Warnf(format, args...)
}
