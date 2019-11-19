package mock

type MockLogger struct {
}

func NewMockLogger() MockLogger {
	return MockLogger{}
}

func (l MockLogger) Error(err error) {
}

func (l MockLogger) Errorf(format string, args ...interface{}) {
}

func (l MockLogger) Warn(err error) {
}

func (l MockLogger) Warnf(format string, args ...interface{}) {
}
