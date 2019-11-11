package logger

type Logger interface {
	Error(err error)
	Errorf(format string, args ...interface{})
	Warn(err error)
	Warnf(format string, args ...interface{})
}
