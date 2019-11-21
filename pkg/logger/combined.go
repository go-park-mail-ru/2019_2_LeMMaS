package logger

type combinedLogger struct {
	Loggers []Logger
}

func NewCombinedLogger(loggers ...Logger) Logger {
	return combinedLogger{loggers}
}

func (l combinedLogger) Error(err error) {
	for _, logger := range l.Loggers {
		logger.Error(err)
	}
}

func (l combinedLogger) Errorf(format string, args ...interface{}) {
	for _, logger := range l.Loggers {
		logger.Errorf(format, args...)
	}
}

func (l combinedLogger) Warn(err error) {
	for _, logger := range l.Loggers {
		logger.Warn(err)
	}
}

func (l combinedLogger) Warnf(format string, args ...interface{}) {
	for _, logger := range l.Loggers {
		logger.Warnf(format, args...)
	}
}
