package logger

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

var counter int

type TestLogger struct {
}

func (l TestLogger) Error(err error) {
	counter++
}

func (l TestLogger) Errorf(format string, args ...interface{}) {
	counter++
}

func (l TestLogger) Warn(err error) {
	counter++
}

func (l TestLogger) Warnf(format string, args ...interface{}) {
	counter++
}

func TestCombinedLogger(t *testing.T) {
	logger1 := TestLogger{}
	logger2 := TestLogger{}
	combined := NewCombinedLogger(logger1, logger2)
	combined.Error(errors.New("err"))
	combined.Errorf("err")
	combined.Warn(errors.New("err"))
	combined.Warnf("err")
	assert.Equal(t, 8, counter)
}
