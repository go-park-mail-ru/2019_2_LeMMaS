package logger

type CombinedLogger struct {
	loggers []Logger
}

func NewCombinedLogger(loggers ...Logger) Logger {
	return CombinedLogger{loggers}
}

func (l CombinedLogger) Error(err error) {
	for _, logger := range l.loggers {
		logger.Error(err)
	}
}

func (l CombinedLogger) Errorf(format string, args ...interface{}) {
	for _, logger := range l.loggers {
		logger.Errorf(format, args...)
	}
}

func (l CombinedLogger) Warn(err error) {
	for _, logger := range l.loggers {
		logger.Warn(err)
	}
}

func (l CombinedLogger) Warnf(format string, args ...interface{}) {
	for _, logger := range l.loggers {
		logger.Warnf(format, args...)
	}
}
