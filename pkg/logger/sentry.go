package logger

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	"os"
)

type SentryLogger struct {
}

func NewSentryLogger() (Logger, error) {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: os.Getenv("SENTRY_DSN"),
	})
	if err != nil {
		return nil, err
	}
	return SentryLogger{}, nil
}

func (l SentryLogger) Error(err error) {
	sentry.CaptureException(err)
}

func (l SentryLogger) Errorf(format string, args ...interface{}) {
	sentry.CaptureException(fmt.Errorf(format, args...))
}

func (l SentryLogger) Warn(err error) {
	sentry.CaptureMessage(err.Error())
}

func (l SentryLogger) Warnf(format string, args ...interface{}) {
	sentry.CaptureMessage(fmt.Sprintf(format, args...))
}
