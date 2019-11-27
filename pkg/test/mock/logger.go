package mock

import "testing"

type MockLogger struct {
	t *testing.T
}

func NewMockLogger(t *testing.T) MockLogger {
	return MockLogger{t: t}
}

func (l MockLogger) Error(err error) {
	l.t.Logf("[logged error] %v", err)
}

func (l MockLogger) Errorf(format string, args ...interface{}) {
	l.t.Logf("[logged error] "+format, args)
}

func (l MockLogger) Warn(err error) {
	l.t.Logf("[logged warn] %v", err)
}

func (l MockLogger) Warnf(format string, args ...interface{}) {
	l.t.Logf("[logged warn] "+format, args)
}
