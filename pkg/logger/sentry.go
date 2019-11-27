package logger

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	"os"
)

type sentryLogger struct {
}

func NewSentryLogger() (Logger, error) {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: os.Getenv("SENTRY_DSN"),
	})
	if err != nil {
		return nil, err
	}
	return sentryLogger{}, nil
}

func (l sentryLogger) Error(err error) {
	sentry.CaptureException(err)
}

func (l sentryLogger) Errorf(format string, args ...interface{}) {
	sentry.CaptureException(fmt.Errorf(format, args...))
}

func (l sentryLogger) Warn(err error) {
	sentry.CaptureMessage(err.Error())
}

func (l sentryLogger) Warnf(format string, args ...interface{}) {
	sentry.CaptureMessage(fmt.Sprintf(format, args...))
}
